package templatemessage

type Header struct {
	params []ComponentParam
}

func newHeader() *Header {
	return &Header{params: []ComponentParam{}}
}

func (h *Header) AddText(text string) {
	p := ComponentParam{
		Type: "text",
		Text: text,
	}
	h.params = append(h.params, p)
}

func (h *Header) HasParams() bool {
	return len(h.params) > 0
}

func (h *Header) Component() Component {
	return Component{
		Type:       "header",
		Parameters: h.params,
	}
}
