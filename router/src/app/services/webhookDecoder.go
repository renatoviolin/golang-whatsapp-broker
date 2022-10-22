package services

import (
	"broker/dto"
	"broker/infra/logger"
	"errors"
)

const (
	MESSAGE_FROM_USER   = "from-user"
	MESSAGE_FROM_STATUS = "from-status"
)

type WebhookDecoder struct{}

type DecoderOutput struct {
	WaID        string
	Text        string
	MessageType string // status, text, list_reply, error
	Origin      string
	Status      string
	Error       string
	ListReply   ListReply
}

type ListReply struct {
	Id    string
	Title string
}

func NewWebhookDecoder() WebhookDecoder {
	return WebhookDecoder{}
}

func (h *WebhookDecoder) DecodeWebhook(input dto.WebhookInput) (DecoderOutput, error) {
	waID, origin, err := getMetadata(input)
	if err != nil {
		logger.Error("get-metadata", err.Error())
		return DecoderOutput{}, err
	}

	status, _ := getStatus(input)
	text, _ := getResponseText(input)
	errorMessage, _ := getError(input)
	listReply, _ := getListReply(input)

	var messageType string
	if text != "" {
		messageType = "text"
	} else if listReply.Id != "" {
		messageType = "list_reply"
	} else if status != "" {
		messageType = "status"
	} else if errorMessage != "" {
		messageType = "error"
	}

	output := DecoderOutput{
		WaID:        waID,
		Origin:      origin,
		Text:        text,
		Status:      status,
		Error:       errorMessage,
		ListReply:   listReply,
		MessageType: messageType,
	}
	return output, nil
}

func getMetadata(input dto.WebhookInput) (waID string, origin string, err error) {
	var tempWaID string
	if len(input.Entry) > 0 {
		if len(input.Entry[0].Changes) > 0 {
			if len(input.Entry[0].Changes[0].Value.Contacts) > 0 {
				tempWaID = input.Entry[0].Changes[0].Value.Contacts[0].WaID
				if tempWaID != "" {
					waID = tempWaID
				}
				origin = MESSAGE_FROM_USER
			} else if len(input.Entry[0].Changes[0].Value.Statuses) > 0 {
				tempWaID = input.Entry[0].Changes[0].Value.Statuses[0].RecipientID
				if tempWaID != "" {
					waID = tempWaID
				}
				origin = MESSAGE_FROM_STATUS
			}
		}
	}
	if waID == "" || origin == "" {
		return "", "", errors.New("unable to find whatsapp_number/origin in payload")
	}
	return waID, origin, nil
}

func getResponseText(input dto.WebhookInput) (text string, err error) {
	if len(input.Entry) > 0 {
		if len(input.Entry[0].Changes) > 0 {
			if len(input.Entry[0].Changes[0].Value.Messages) > 0 {
				text = input.Entry[0].Changes[0].Value.Messages[0].Text.Body
			}
		}
	}
	if text == "" {
		return "", errors.New("response text not found")
	}
	return text, nil
}

func getStatus(input dto.WebhookInput) (status string, err error) {
	if len(input.Entry) > 0 {
		if len(input.Entry[0].Changes) > 0 {
			if len(input.Entry[0].Changes[0].Value.Statuses) > 0 {
				status = input.Entry[0].Changes[0].Value.Statuses[0].Status
			}
		}
	}
	if status == "" {
		return "", errors.New("status not found")
	}
	return status, nil
}

func getError(input dto.WebhookInput) (errorMessage string, err error) {
	if len(input.Entry) > 0 {
		if len(input.Entry[0].Changes) > 0 {
			if len(input.Entry[0].Changes[0].Value.Statuses) > 0 {
				if len(input.Entry[0].Changes[0].Value.Statuses[0].Errors) > 0 {
					errorMessage = input.Entry[0].Changes[0].Value.Statuses[0].Errors[0].Title
				}
			}
		}
	}
	if errorMessage == "" {
		return "", errors.New("errorMessage not found")
	}

	return errorMessage, nil
}

func getListReply(input dto.WebhookInput) (listReply ListReply, err error) {
	if len(input.Entry) > 0 {
		if len(input.Entry[0].Changes) > 0 {
			if len(input.Entry[0].Changes[0].Value.Messages) > 0 {
				listReply.Id = input.Entry[0].Changes[0].Value.Messages[0].Interactive.ListReply.ID
				listReply.Title = input.Entry[0].Changes[0].Value.Messages[0].Interactive.ListReply.Title
			}
		}
	}
	if listReply.Id == "" {
		return listReply, errors.New("list-reply not found")
	}
	return listReply, nil
}
