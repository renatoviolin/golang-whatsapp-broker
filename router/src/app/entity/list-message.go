package entity

import (
	"errors"
)

type ListPayload struct {
	MessagingProduct string      `json:"messaging_product"`
	To               string      `json:"to"`
	Type             string      `json:"type"`
	Interactive      Interactive `json:"interactive"`
}

type Interactive struct {
	Type   string `json:"type"`
	Body   Body   `json:"body"`
	Action Action `json:"action"`
}

type Action struct {
	Button   string    `json:"button"`
	Sections []Section `json:"sections"`
}

type Section struct {
	Title string `json:"title"`
	Rows  []Row  `json:"rows"`
}

type Row struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type Body struct {
	Text string `json:"text"`
}

func NewListPayload(title string, body string, rows []Row, waID string) (payload ListPayload, err error) {
	if title == "" || body == "" || len(rows) == 0 || waID == "" {
		return payload, ErrInvalidListPayload
	}
	var ids = []string{}

	for _, v := range rows {
		if v.ID == "" || v.Title == "" {
			return payload, errors.New("empty ID or Title")
		}
		for _, storeId := range ids {
			if storeId == v.ID {
				return payload, errors.New("list_id duplicated")
			}
		}
		ids = append(ids, v.ID)
	}

	payload.MessagingProduct = "whatsapp"
	payload.Type = "interactive"
	payload.To = waID
	payload.Interactive.Type = "list"
	payload.Interactive.Body.Text = "Choose an option:"
	payload.Interactive.Action.Button = title
	payload.Interactive.Action.Sections = []Section{
		{Title: body, Rows: rows},
	}
	return payload, nil
}
