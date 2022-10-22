package controller

import (
	"broker/app/services"
)

type RouterController struct {
	routerWebhook *services.RouterWebhook
}

func NewRouterController(routerWebhook *services.RouterWebhook) *RouterController {
	return &RouterController{routerWebhook: routerWebhook}
}

func (h *RouterController) HandleMessageFromKafka(input []byte) error {
	return h.routerWebhook.Dispatch(input)
}
