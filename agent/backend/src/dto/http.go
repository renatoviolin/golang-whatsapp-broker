package dto

type MessageOutput struct {
	Type         string `json:"type"`
	WaID         string `json:"waID"`
	Text         string `json:"text"`
	Status       string `json:"status"`
	MessageID    string `json:"message_id"`
	ErrorMessage string `json:"error_message"`
}

type HttpMessageInputPayload struct {
	ListItems []ListItem `json:"list_items"`
	Type      string     `json:"type"`
	WaID      string     `json:"waID"`
	Text      string     `json:"text"`
	Image     string     `json:"image"`
	WSAction  string     `json:"ws_action"`
}

type MessageError struct{}
