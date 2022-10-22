package services

import (
	"broker/app/entity"
	"broker/dto"
	"broker/infra/custom_errors"
	"broker/infra/http_client"
	"broker/infra/kafka"
	"broker/infra/redis"
	"encoding/json"
	"fmt"
)

type RouteAgentMessages struct {
	kafkaProducer kafka.Producer
	sender        *Sender
	redisCache    *redis.RedisClient
	agentConfig   entity.AgentConfig
}

func NewRouteAgentMessages(kafkaProducer *kafka.Producer, httpClient *http_client.Client, redisCache *redis.RedisClient, agentConfig entity.AgentConfig) RouteAgentMessages {
	return RouteAgentMessages{
		kafkaProducer: *kafkaProducer,
		sender:        NewSendService(httpClient),
		redisCache:    redisCache,
		agentConfig:   agentConfig,
	}
}

func (h *RouteAgentMessages) Dispatch(inputBytes []byte) (err error) {
	var agentInput dto.AgentMessageInput
	defer func() {
		if err != nil {
			kError := dto.KafkaErrorMessage{WaID: agentInput.WaID, Text: err.Error(), RawPayload: []byte(err.Error())}
			h.kafkaProducer.ProduceError(kError, h.agentConfig.ErrorTopic)
		}
	}()

	err = json.Unmarshal(inputBytes, &agentInput)
	if err != nil {
		return custom_errors.New("agent-dispatch", err.Error())
	}
	if agentInput.WaID == "" {
		return custom_errors.New("agent-dispatch", "waID not found")
	}

	agentClientID, err := h.redisCache.Get(agentInput.WaID)
	if err != nil || agentClientID != h.agentConfig.ID {
		return custom_errors.New("agent-dispatch", fmt.Sprintf("agentID %s not associated with waID: %s", h.agentConfig.ID, agentInput.WaID))
	}
	if agentInput.MessageType == "text" {
		if err := h.sender.SendText(agentInput.Text, agentInput.WaID); err != nil {
			return custom_errors.New("agent-dispatch", err.Error())
		}
		return nil
	}
	if agentInput.MessageType == "list" {
		var rows []entity.Row
		for _, v := range agentInput.List {
			rows = append(rows, entity.Row{ID: v.ID, Title: v.Title})
		}
		if err := h.sender.SendList(agentInput.ListTitle, agentInput.ListBody, rows, agentInput.WaID); err != nil {
			return custom_errors.New("agent-dispatch", err.Error())
		}
		return nil
	}
	return custom_errors.New("agent-dispatch", "invalid message type")
}
