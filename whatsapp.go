package whatsapp

import (
	utils "github.com/digitalhurricane-io/go-web-utils"
	"github.com/digitalhurricane-io/go-whatsapp-wrapper/templatemessage"
	"net/http"
)

type WhatsApp struct {
	PhoneNumberID string
	AuthToken     string
	Client        *http.Client
}

func NewWhatsApp(phoneNumberID, authToken string, client ...*http.Client) WhatsApp {
	var c *http.Client

	if len(client) == 0 {
		c = utils.NewHttpClientWithTimeout()
	} else {
		c = client[0]
	}

	return WhatsApp{
		PhoneNumberID: phoneNumberID,
		AuthToken:     authToken,
		Client:        c,
	}
}

func (w *WhatsApp) NewTemplateMessage(templateName, toPhoneNumber, langCode string) templatemessage.TemplateMessage {
	return templatemessage.NewTemplateMessage(w.Client, w.PhoneNumberID, w.AuthToken, templateName, toPhoneNumber, langCode)
}
