package dto

type KafkaOutputMessage struct {
	WaID       string `json:"waID"`
	Type       string `json:"type"` // text, list_reply, inital, status
	Text       string `json:"text"`
	RawPayload []byte `json:"raw_payload"`
}

type KafkaErrorMessage struct {
	WaID       string `json:"waID"`
	Text       string `json:"text"`
	RawPayload []byte `json:"raw_payload"`
}
