package kafka

import (
	"broker/dto"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

var servers = []string{
	"localhost:9093",
}
var topic = "test-topic"

// func Test_Producer_InvalidConnection(t *testing.T) {
// 	var servers = []string{
// 		"localhost:9094",
// 	}
// 	producer := NewProducer(servers)
// 	require.Empty(t, producer, topic)
// }

func Test_Producer_ValidConnection(t *testing.T) {
	producer := NewProducer(servers)

	for i := 0; i < 10; i++ {
		msg := dto.KafkaOutputMessage{
			WaID: "123",
			Text: "teste",
		}
		partition, offset, err := producer.Produce(msg, topic)
		log.Printf("partition: %d  offset: %d", partition, offset)
		require.NoError(t, err)
	}
}
