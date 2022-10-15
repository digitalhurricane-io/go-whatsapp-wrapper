package templatemessage

import (
	"github.com/digitalhurricane-io/go-whatsapp-wrapper/base"
	"github.com/digitalhurricane-io/go-whatsapp-wrapper/base/basemessagetype"
)

// TemplateParams used when sending a template message
type TemplateParams struct {
	base.BaseMessageParams
	Template Template `json:"template"`
}

// NewTemplateParams paramText strings are what fills in the placeholders in the whatsapp template {{1}}
func NewTemplateParams(templateName, toPhoneNumber, langCode string) TemplateParams {
	return TemplateParams{
		BaseMessageParams: base.NewBaseMessageParams(basemessagetype.Template, toPhoneNumber),
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
