package controllers

import (
	"broker/app/services"
)

type WebhookController struct {
	receiveService services.ReceiveService
	tokenService   services.TokenService
}

func NewWebhookController(receiveService services.ReceiveService, tokenService services.TokenService) WebhookController {
	return WebhookController{receiveService: receiveService, tokenService: tokenService}
}

func (h *WebhookController) ReceiveFromWhatsapp(input []byte) error {
	err := h.receiveService.Receive(input)
	if err != nil {
		return err
	}
	return nil
}

func (h *WebhookController) VerifyToken(token string) error {
	return h.tokenService.VerifyToken(token)
}
