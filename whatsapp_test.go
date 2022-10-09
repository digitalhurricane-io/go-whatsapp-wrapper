package whatsapp

import (
	"encoding/json"
	"log"
	"os"
	"testing"
)

type testConf struct {
	PhoneNumberID                     string                `json:"phone_number_id"`
	AuthToken                         string                `json:"auth_token"`
	TestSendTemplateMessageWUrlButton templateMessageConfig `json:"test_send_template_message_w_url_button"`
	TestSendOTPMessage                templateMessageConfig `json:"test_send_otp_message"`
}

type templateMessageConfig struct {
	TemplateName  string `json:"template_name"`
	ToPhoneNumber string `json:"to_phone_number"`
	LangCode      string `json:"lang_code"`
}

var testConfig testConf

func TestMain(m *testing.M) {
	file, err := os.ReadFile("test_config.json")
	if err != nil {
		log.Fatal("failed to read file test_config.json. please copy test_config.sample.json to "+
			"test_config.json and customize it with your own information for testing", err)
		return
	}

	err = json.Unmarshal(file, &testConfig)
	if err != nil {
		log.Fatal("failed to parse json from contents of test_config.json", err)
		return
	}

	code := m.Run()

	os.Exit(code)
}

func TestSendTemplateMessageWUrlButton(t *testing.T) {

	wa := NewWhatsApp(testConfig.PhoneNumberID, testConfig.AuthToken)

	msg := wa.NewTemplateMessage(
		testConfig.TestSendTemplateMessageWUrlButton.TemplateName,
		testConfig.TestSendTemplateMessageWUrlButton.ToPhoneNumber,
		testConfig.TestSendTemplateMessageWUrlButton.LangCode,
	)

	msg.Buttons().AddUrlButton(0, "myurlsuffix")

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}
}

func TestSendOTPMessage(t *testing.T) {

	wa := NewWhatsApp(testConfig.PhoneNumberID, testConfig.AuthToken)

	msg := wa.NewTemplateMessage(
		testConfig.TestSendOTPMessage.TemplateName,
		testConfig.TestSendOTPMessage.ToPhoneNumber,
		testConfig.TestSendOTPMessage.LangCode,
	)
	msg.Header().AddText("Bobs Burgers")
	msg.Body().AddText("111-111")

	_, err := msg.Send()
	if err != nil {
		t.Fatal(err)
	}

}
