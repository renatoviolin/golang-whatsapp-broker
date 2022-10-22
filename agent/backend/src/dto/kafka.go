package dto

type KafkaMessageInput struct {
	WaID       string `json:"waID"`
	Type       string `json:"type"` // text, list_reply, inital, status
	Text       string `json:"text"`
	Media      string `json:"media"`
	RawPayload []byte `json:"raw_payload"`
}

type KafkaErrorInput struct {
	WaID       string `json:"waID"`
	Text       string `json:"text"`
	RawPayload []byte `json:"raw_payload"`
}

type KafkaMessageOutput struct {
	WaID        string     `json:"waID"`
	MessageType string     `json:"message_type"` // text, list
	Text        string     `json:"text"`
	ListTitle   string     `json:"list_title"`
	ListBody    string     `json:"list_body"`
	List        []ListItem `json:"list_items"`
}

type ListItem struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
