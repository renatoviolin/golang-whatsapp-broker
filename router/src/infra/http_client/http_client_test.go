package http_client

import (
	"broker/util"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func init() {
	util.LoadVars()
}

func TestHttpClientGet(t *testing.T) {
	url := "https://reqres.in/api/users"
	client := NewClient()
	data := []byte(`{"name": "morpheus","job": "leader"}`)
	res, status, err := client.Post(url, data)
	require.Equal(t, 201, status)
	require.NoError(t, err)
	require.NotEmpty(t, res)
}

func TestHttpClientGetWhatsapp(t *testing.T) {
	url := "https://graph.facebook.com/v14.0/102261592636695/messages?access_token=" + os.Getenv("WHATSAPP_ACCESS_TOKEN")

	client := NewClient()
	data := []byte(`{"name": "morpheus","job": "leader"}`)
	_, status, err := client.Post(url, data)
	require.Equal(t, 400, status)
	require.Error(t, err)
}
