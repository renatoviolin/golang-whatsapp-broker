package services

import (
	"broker/app/entity"
	"broker/dto"
	"broker/infra/custom_errors"
	"broker/infra/kafka"
	"broker/infra/redis"
	"encoding/json"
	"errors"
	"fmt"
)

const (
	ROUTER_ID = "router"
)

type RouterWebhook struct {
	redis    *redis.RedisClient
	producer *kafka.Producer
	sender   *Sender
	agents   []entity.AgentConfig
	decoder  WebhookDecoder
}

func NewRouterWebhook(producer *kafka.Producer, redis *redis.RedisClient, sender *Sender) RouterWebhook {
	return RouterWebhook{
		producer: producer,
		redis:    redis,
		sender:   sender,
		decoder:  NewWebhookDecoder(),
	}
}

func (h *RouterWebhook) RegisterAgent(agent entity.AgentConfig) error {
	h.agents = append(h.agents, agent)
	return nil
}

func (h *RouterWebhook) getAgent(agentID string) *entity.AgentConfig {
	for _, v := range h.agents {
		if v.ID == agentID {
			return &v
		}
	}
	return nil
}

func (h *RouterWebhook) Dispatch(inputBytes []byte) error {
	var input dto.WebhookInput
	if err := json.Unmarshal(inputBytes, &input); err != nil {
		return err
	}

	// 1. try to read all data from Webhook payload
	webhookDecoded, err := h.decoder.DecodeWebhook(input)
	if err != nil {
		return err
	}

	if webhookDecoded.Text == "#exit" {
		if err = h.UpdateStatus(webhookDecoded.WaID, ROUTER_ID, webhookDecoded.Origin); err != nil {
			return err
		}
		if err := h.sender.SendText("closing....", webhookDecoded.WaID); err != nil {
			return err
		}
		if err := h.SendMenuOptions(webhookDecoded.WaID); err != nil {
			return err
		}
		return nil
	}

	// 2. check in redis if the ID is associated with any agent
	// if expired or not found, means the first contact from the agent
	// send menu options and store payload in router topic
	agentID, err := h.redis.Get(webhookDecoded.WaID)
	if err != nil {
		if err = h.UpdateStatus(webhookDecoded.WaID, ROUTER_ID, webhookDecoded.Origin); err != nil {
			return err
		}
		if webhookDecoded.Origin == MESSAGE_FROM_USER {
			if err := h.SendMenuOptions(webhookDecoded.WaID); err != nil {
				return err
			}
		}
		return nil
	}

	// 3. check if the message is set as ROUTER owner
	if agentID == ROUTER_ID {
		desiredAgentID, err := h.ReadMenu(webhookDecoded.ListReply.Id)
		if err != nil {
			if webhookDecoded.Origin == MESSAGE_FROM_USER {
				if err := h.SendMenuOptions(webhookDecoded.WaID); err != nil {
					return err
				}
			}
			if err = h.UpdateStatus(webhookDecoded.WaID, ROUTER_ID, webhookDecoded.Origin); err != nil {
				return err
			}
			return nil
		}

		// has a valid desiredAgentID, set to that desired agentID
		if err = h.UpdateStatus(webhookDecoded.WaID, desiredAgentID, webhookDecoded.Origin); err == nil {
			if webhookDecoded.Origin == MESSAGE_FROM_USER {
				if err := h.sender.SendText("Transfering to "+h.getAgent(desiredAgentID).Name, webhookDecoded.WaID); err != nil {
					return err
				}
				kMsg := dto.KafkaOutputMessage{WaID: webhookDecoded.WaID, Type: "initial", Text: fmt.Sprintf("user %s want to talk to you", webhookDecoded.WaID)}
				_, _, err = h.producer.Produce(kMsg, h.getAgent(desiredAgentID).ReadTopic)
			}
		}
		return err
	}

	// 4. has a valid agent_id, send to specific topic
	if err = h.UpdateStatus(webhookDecoded.WaID, agentID, webhookDecoded.Origin); err != nil {
		return err
	}
	if webhookDecoded.Status == "failed" || webhookDecoded.Error != "" {
		kError := dto.KafkaErrorMessage{
			WaID:       webhookDecoded.WaID,
			Text:       webhookDecoded.Error,
			RawPayload: inputBytes,
		}
		_, _, err = h.producer.ProduceError(kError, h.getAgent(agentID).ErrorTopic)
	} else {
		kMsg := newKafkaOutputMessage(&webhookDecoded, inputBytes)
		_, _, err = h.producer.Produce(kMsg, h.getAgent(agentID).ReadTopic)
	}
	return err
}

func newKafkaOutputMessage(webhookDecoded *DecoderOutput, inputBytes []byte) dto.KafkaOutputMessage {
	var text string
	if webhookDecoded.Status != "" {
		text = webhookDecoded.Status
	} else if webhookDecoded.Text != "" {
		text = webhookDecoded.Text
	} else if webhookDecoded.ListReply.Id != "" {
		text = webhookDecoded.ListReply.Id
	}

	return dto.KafkaOutputMessage{
		WaID:       webhookDecoded.WaID,
		Type:       webhookDecoded.MessageType,
		Text:       text,
		RawPayload: inputBytes,
	}
}

func (h *RouterWebhook) SendMenuOptions(waID string) error {
	if len(h.agents) == 0 {
		e := custom_errors.New("router", "there is no agent registred in the Router")
		return e
	}
	actions := []entity.Row{}
	for _, v := range h.agents {
		actions = append(actions, entity.Row{ID: v.ID, Title: v.Name})
	}

	return h.sender.SendList("Choose an Agent:", "Available Agents", actions, waID)
}

func (h *RouterWebhook) ReadMenu(agentID string) (string, error) {
	agent := h.getAgent(agentID)
	if agent == nil {
		return "", errors.New("agent not found in menu response")
	}
	return agent.ID, nil
}

func (h *RouterWebhook) UpdateStatus(whatsappID string, agentID string, origin string) (err error) {
	if origin == MESSAGE_FROM_STATUS {
		return nil
	}
	return h.redis.Save(whatsappID, agentID)
}
