package interactivemessage

import (
	"github.com/digitalhurricane-io/go-whatsapp-wrapper/base"
	messagetype "github.com/digitalhurricane-io/go-whatsapp-wrapper/interactivemessage/interactivemessagetype"
)

type Params struct {
	base.BaseMessageParams
	Action Action                  `json:"action"`
	Header *Header                 `json:"header,omitempty"`
	Body   *Body                   `json:"body,omitempty"`
	Footer *Footer                 `json:"footer,omitempty"`
	Type   messagetype.MessageType `json:"type"`
}

type Header struct {
	Type     HeaderType `json:"type"`
	Document *Media     `json:"document,omitempty"`
	Image    *Media     `json:"image,omitempty"`
	Text     string     `json:"text,omitempty"`
	Video    *Media     `json:"video,omitempty"`
}

type HeaderType string

const (
	HeaderTypeText     HeaderType = "text"
	HeaderTypeVideo    HeaderType = "video"
	HeaderTypeImage    HeaderType = "image"
	HeaderTypeDocument HeaderType = "document"
)

type Media struct {
	ID       string `json:"id"`
	Link     string `json:"link"`
	Caption  string `json:"caption"`
	Filename string `json:"filename"`
}

type Body struct {
	Text string `json:"text"`
}

type Footer struct {
	Text string `json:"text"`
}

// Action https://developers.facebook.com/docs/whatsapp/cloud-api/reference/messages#action-object
type Action struct {
	Button            string    `json:"button"`
	Buttons           []Button  `json:"buttons"`
	CatalogID         string    `json:"catalog_id"`
	ProductRetailerID string    `json:"product_retailer_id"`
	Sections          []Section `json:"sections"`
}

type Section struct {
	Title        string    `json:"title,omitempty"` // required if message has more than 1 section
	Rows         []Row     `json:"rows,omitempty"`
	ProductItems []Product `json:"product_items,omitempty"`
}

type Product struct {
	ProductRetailerID string `json:"product_retailer_id"`
}

type Row struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Button https://developers.facebook.com/docs/whatsapp/cloud-api/reference/messages#button-parameter-object
type Button struct {
	Type    ButtonType  `json:"type"`
	Payload interface{} `json:"payload"`
	Text    string      `json:"text"`
}

type ButtonType string

const (
	ButtonTypePayload ButtonType = "payload"
	ButtonTypeText    ButtonType = "text"
)
