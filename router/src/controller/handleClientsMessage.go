package controller

import (
	"broker/app/services"
)

type AgentController struct {
	routeragent services.RouteAgentMessages
}

func NewAgentController(routeragent services.RouteAgentMessages) AgentController {
	return AgentController{routeragent: routeragent}
}

func (h *AgentController) HandleMessageFromKafka(input []byte) (err error) {
	_ = h.routeragent.Dispatch(input)
	return nil
}
