package kafka

import (
	"fmt"
	"testing"

	"github.com/Shopify/sarama"
)

type ConsumerHandler struct {
}

func NewConsumerHandler() *ConsumerHandler {
	return &ConsumerHandler{}
}

func CustomFunction(input []byte) error {
	fmt.Printf("\n%+v\n", string(input))
	return nil
}

func (h *ConsumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():
			err := CustomFunction(message.Value)
			if err == nil {
				session.MarkMessage(message, "")
			}
		case <-session.Context().Done():
			return nil
		}
	}
}

func (h *ConsumerHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *ConsumerHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func Test_Consumer(t *testing.T) {
	consumer := NewConsumer(servers, "from-whatsapp", "cg-1", NewConsumerHandler())
	consumer.StartConsumer()
}
