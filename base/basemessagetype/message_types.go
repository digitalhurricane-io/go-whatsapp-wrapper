package basemessagetype

type MessageType string

const (
	Template    MessageType = "template"
	Text        MessageType = "text"
	Reaction    MessageType = "reaction"
	Image       MessageType = "image"
	Location    MessageType = "location"
	Contacts    MessageType = "contacts"
	Interactive MessageType = "interactive"
)
