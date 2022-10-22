package entity

import "errors"

var (
	ErrEmptyPayload       = errors.New("invalid payload: text/whatsappID is empty")
	ErrInvalidListPayload = errors.New("invalid list payload: fill all fields")
)

type TextPayload struct {
	Text             TextBody `json:"text"`
	MessagingProduct string   `json:"messaging_product"`
	To               string   `json:"to"`
}

type TextBody struct {
	Body string `json:"body"`
}

func NewTextPayload(text string, to string) (payload TextPayload, err error) {
	if text == "" || to == "" {
		return payload, ErrEmptyPayload
	}

	payload.MessagingProduct = "whatsapp"
	payload.Text = TextBody{text}
	payload.To = to
	return payload, nil
}
