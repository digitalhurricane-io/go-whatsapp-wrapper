package templatemessage

type Body struct {
	params []ComponentParam
}

func newBody() *Body {
	return &Body{params: []ComponentParam{}}
}

func (b *Body) AddText(text string) {
	p := ComponentParam{
		Type: "text",
		Text: text,
	}
	b.params = append(b.params, p)
}

//func (b *Body) AddCurrency() {
//	log.Println("AddCurrency unimplemented")
//	// todo
//}
//
//func (b *Body) AddDateTime() {
//	log.Println("AddDateTime unimplemented")
//	// todo
//}

func (b *Body) HasParams() bool {
	return len(b.params) > 0
}

func (b *Body) Component() Component {
	return Component{
		Type:       "body",
		Parameters: b.params,
	}
}
