package dto

type AgentMessageInput struct {
	WaID        string     `json:"waID"`
	MessageType string     `json:"message_type"`
	Text        string     `json:"text"`
	ListTitle   string     `json:"list_title"`
	ListBody    string     `json:"list_body"`
	List        []ListItem `json:"list_items"`
}

type ListItem struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
