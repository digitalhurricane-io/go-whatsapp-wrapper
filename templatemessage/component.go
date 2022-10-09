package templatemessage

type Component struct {
	// Required.
	Type string `json:"type"`

	// Required when type=button. Not used for the other templatemessage.
	Subtype string `json:"sub_type,omitempty"`

	// Required when type=button.
	Parameters []ComponentParam `json:"parameters,omitempty"`

	// Required when type=button. Not used for the other templatemessage
	Index *ButtonIndex `json:"index,omitempty"`
}

type ComponentParam struct {
	Type ComponentType `json:"type"`

	// One of the following fields must not be empty
	Text     string                  `json:"text,omitempty"`
	Currency *ComponentParamCurrency `json:"currency,omitempty"`
	DateTime *ComponentParamDateTime `json:"date_time,omitempty"`
}

type ComponentType string

const (
	ComponentTypeText     ComponentType = "text"
	ComponentTypeCurrency ComponentType = "currency"
	ComponentTypeDateTime ComponentType = "date_time"
)

type ComponentParamCurrency struct {
	// todo
}
type ComponentParamDateTime struct {
	// todo
}
