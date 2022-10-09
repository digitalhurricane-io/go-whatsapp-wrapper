package templatemessage

// TemplateParams used when sending a template message
type TemplateParams struct {
	MessagingProduct string   `json:"messaging_product"`
	RecipientType    string   `json:"recipient_type"`
	To               string   `json:"to"`
	Type     string   `json:"type"`
	Template Template `json:"template"`
}

// NewTemplateParams paramText strings are what fills in the placeholders in the whatsapp template {{1}}
func NewTemplateParams(templateName, toPhoneNumber, langCode string) TemplateParams {
	return TemplateParams{
		MessagingProduct: "whatsapp",
		RecipientType:    "individual",
		To:               toPhoneNumber,
		Type:             "template",
		Template: Template{
			Name: templateName,
			Language: Language{
				Code: langCode,
			},
		},
	}
}

type Template struct {
	Name       string      `json:"name"`
	Language   Language    `json:"language"`
	Components []Component `json:"components"`
}

type Language struct {
	Code string `json:"code"`
}
