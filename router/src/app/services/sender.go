package services

import (
	"broker/app/entity"
	"broker/infra/custom_errors"
	"broker/infra/http_client"
	"encoding/json"
	"errors"
	"os"
)

var (
	ErrHttpRequest = errors.New("error while request whatsapp URL")
)

type Sender struct {
	httpClient *http_client.Client
	Url        string
}

func NewSendService(httpClient *http_client.Client) *Sender {
	return &Sender{
		httpClient: httpClient,
		Url:        "https://graph.facebook.com/v14.0/102261592636695/messages?access_token=" + os.Getenv("WHATSAPP_ACCESS_TOKEN"),
	}
}

func (h *Sender) SendText(text string, waID string) error {
	payload, err := entity.NewTextPayload(text, waID)
	if err != nil {
		return custom_errors.New("text-payload", err.Error())
	}
	bytePayload, _ := json.Marshal(payload)
	_, statusCode, err := h.httpClient.Post(h.Url, bytePayload)
	if err != nil {
		return custom_errors.New("post", err.Error())
	}
	if statusCode != 200 {
		return custom_errors.New("status-code not 200", ErrHttpRequest.Error())
	}
	return nil
}

func (h *Sender) SendList(title string, body string, rows []entity.Row, waID string) error {
	payload, err := entity.NewListPayload(title, body, rows, waID)
	if err != nil {
		return custom_errors.New("list-payload", err.Error())
	}
	bytePayload, _ := json.Marshal(payload)
	_, statusCode, err := h.httpClient.Post(h.Url, bytePayload)
	if err != nil {
		return custom_errors.New("list-payload", err.Error())
	}
	if statusCode != 200 {
		return custom_errors.New("list-payload", ErrHttpRequest.Error())
	}
	return nil
}
