package templatemessage

type Buttons struct {
	buttons []Component
}

func newButtons() *Buttons {
	return &Buttons{}
}

func (b *Buttons) AddUrlButton(index ButtonIndex, urlSuffix string) {
	c := Component{
		Type:    "button",
		Subtype: "url",
		Index:   &index,
		Parameters: []ComponentParam{
			{
				Type: "text",
				Text: urlSuffix,
			},
		},
	}

	// if component with index already exists, overwrite it
	var existingIndex = -1
	for i, item := range b.buttons {
		if *item.Index == index {
			existingIndex = i
			break
		}
	}

	if existingIndex != -1 {
		b.buttons[existingIndex] = c
	} else {
		b.buttons = append(b.buttons, c)
	}
}

func (b *Buttons) HasComponents() bool {
	return len(b.buttons) > 0
}

type ButtonIndex int

const (
	ButtonIndexZero ButtonIndex = 0
	ButtonIndexOne  ButtonIndex = 1
	ButtonIndexTwo  ButtonIndex = 2
)

type ButtonType string

const (
	ButtonTypePayload ButtonType = "payload"
	ButtonTypeText    ButtonType = "text"
)
