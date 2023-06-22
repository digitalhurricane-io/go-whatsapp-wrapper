package templatemessage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

const apiVersion = "v15.0"

type TemplateMessage struct {
	params        TemplateParams
	body          *Body
	header        *Header
	buttons       *Buttons
	phoneNumberID string
	authToken     string
	client        *http.Client
}

func NewTemplateMessage(client *http.Client, phoneNumberID, authToken, templateName, toPhoneNumber, langCode string) TemplateMessage {
	return TemplateMessage{
		client:        client,
		params:        NewTemplateParams(templateName, toPhoneNumber, langCode),
		body:          newBody(),
		header:        newHeader(),
		buttons:       newButtons(),
		phoneNumberID: phoneNumberID,
		authToken:     authToken,
	}
}

func (tm *TemplateMessage) Header() *Header {
	return tm.header
}

func (tm *TemplateMessage) Body() *Body {
	return tm.body
}

func (tm *TemplateMessage) Buttons() *Buttons {
	return tm.buttons
}

func (tm *TemplateMessage) MarshalJSON() ([]byte, error) {
	params := tm.params

	if tm.header.HasParams() {
		params.Template.Components = append(params.Template.Components, tm.header.Component())
	}

	if tm.body.HasParams() {
		params.Template.Components = append(params.Template.Components, tm.body.Component())
	}

	if tm.buttons.HasComponents() {
		params.Template.Components = append(params.Template.Components, tm.buttons.buttons...)
	}

	return json.Marshal(params)
}

// Send Sends the template message. An http client can be passed in, if not, the client from the parent WhatsApp
// will be used
func (tm *TemplateMessage) Send(client ...*http.Client) (reqPrettyJson string, err error) {
	var c = tm.client
	if len(client) > 0 {
		c = client[0]
	}

	var URL = fmt.Sprintf("https://graph.facebook.com/%s/%s/messages", apiVersion, tm.phoneNumberID)

	data, err := json.Marshal(&tm)
	if err != nil {
		return "", err
	}

	prettyJson, err := json.MarshalIndent(&tm, "", "    ")
	if err != nil {
		return "", err
	}

	reqPrettyJson = string(prettyJson)

	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(data))
	if err != nil {
		return reqPrettyJson, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tm.authToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Do(req)
	if err != nil {
		return reqPrettyJson, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return reqPrettyJson, nil
	}

	var respBodyJsonText = "failed to read body."

	respData, err := io.ReadAll(resp.Body)
	if err == nil {
		respBodyJsonText = string(respData)
	}

	var errRespData errResponseJson
	err = json.Unmarshal(respData, &errRespData)
	if err != nil {
		return reqPrettyJson, err
	}

	// List of codes: https://developers.facebook.com/docs/whatsapp/cloud-api/support/error-codes
	if errRespData.Error.Code == 100 { // code 100 actually means invalid parameter but we're going to take it to always mean invalid phone number
		return reqPrettyJson, ErrInvalidPhoneNumber
	} else if errRespData.Error.Code == 131026 {
		return reqPrettyJson, ErrNoWhatsappAccount
	}

	err = errors.Errorf("got status code %d when sending whatsapp "+
		"template message. resp body: %s", resp.StatusCode, respBodyJsonText)

	return reqPrettyJson, err
}

var ErrInvalidPhoneNumber = errors.New("Invalid phone number")
var ErrNoWhatsappAccount = errors.New("Phone number is not associated with a whatsapp account")

type errResponseJson struct {
	Error struct {
		Code int `json:"code"`
	} `json:"error"`
}
