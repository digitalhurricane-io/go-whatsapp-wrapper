package base

import "github.com/digitalhurricane-io/go-whatsapp-wrapper/base/basemessagetype"

type BaseMessageParams struct {
	MessagingProduct string        `json:"messaging_product"`
	RecipientType    string        `json:"recipient_type"`
	To               string        `json:"to"`
	Type             string        `json:"type"`
	Context          *ReplyContext `json:"context,omitempty"`
}

// ReplyContext used when replying to a message
type ReplyContext struct {
	MessageID string `json:"message_id"`
}

func NewBaseMessageParams(messageType basemessagetype.MessageType, toPhoneNumber string, messageID ...string) BaseMessageParams {
	var mID *ReplyContext
	if len(messageID) > 0 {
		mID = &ReplyContext{MessageID: messageID[0]}
	}

	return BaseMessageParams{
		MessagingProduct: "whatsapp",
		RecipientType:    "individual",
		To:               toPhoneNumber,
		Type:             string(messageType),
		Context:          mID,
	}
}
