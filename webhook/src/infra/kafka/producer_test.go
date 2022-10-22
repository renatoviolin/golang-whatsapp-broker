package kafka

import (
	"log"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

var servers = []string{
	"localhost:9093",
}
var topic = "from-whatsapp"

func Test_Producer_InvalidConnection(t *testing.T) {
	var servers = []string{
		"localhost:9094",
	}
	producer := NewProducer(servers, topic)
	require.Empty(t, producer)
}

func Test_Producer_ValidConnection(t *testing.T) {
	producer := NewProducer(servers, topic)

	for i := 0; i < 50; i++ {
		message := []byte("teste message from test Producer: " + strconv.Itoa(i))
		partition, offset, err := producer.Produce(message)
		log.Printf("partition: %d  offset: %d", partition, offset)
		require.NoError(t, err)
	}
}
