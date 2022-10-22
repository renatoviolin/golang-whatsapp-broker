package redis

import (
	"broker/util"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func init() {
	util.LoadVars()
}

func Test_Connect(t *testing.T) {
	client := NewRedisClient(os.Getenv("REDIS_CONNECTION"))
	require.NotNil(t, client)
}

func Test_Save(t *testing.T) {
	client := NewRedisClient(os.Getenv("REDIS_CONNECTION"))
	require.NotNil(t, client)

	for i := 0; i <= 10; i++ {
		key := "test-key-" + fmt.Sprint(i*3)
		value := "test-value-2"
		err := client.Save(key, value)
		require.NoError(t, err)

	}
}

func Test_ClearAll(t *testing.T) {
	client := NewRedisClient(os.Getenv("REDIS_CONNECTION"))
	require.NotNil(t, client)

	err := client.ClearAll()
	require.NoError(t, err)
}

func Test_Get_NotFound(t *testing.T) {
	client := NewRedisClient(os.Getenv("REDIS_CONNECTION"))
	require.NotNil(t, client)

	_, err := client.Get("none")
	require.Error(t, err)
}

func Test_Get_Found(t *testing.T) {
	client := NewRedisClient(os.Getenv("REDIS_CONNECTION"))
	require.NotNil(t, client)

	err := client.Save("teste-key-2", "teste-value")
	require.NoError(t, err)

	res, err := client.Get("teste-key-2")
	require.NoError(t, err)
	require.Equal(t, "teste-value", res)
}

func Test_Touch(t *testing.T) {
	client := NewRedisClient(os.Getenv("REDIS_CONNECTION"))
	require.NotNil(t, client)

	err := client.Save("teste-key-touch", "teste-value")
	require.NoError(t, err)

	duration, err := client.GetTTL("teste-key-touch")
	require.NoError(t, err)
	require.Equal(t, int64(600), duration)
	time.Sleep(time.Second * 5)

	err = client.Touch("teste-key-touch")
	require.NoError(t, err)

	duration, err = client.GetTTL("teste-key-touch")
	require.NoError(t, err)
	require.Equal(t, int64(600), duration)
}

func Test_GetTTL_Found(t *testing.T) {
	client := NewRedisClient(os.Getenv("REDIS_CONNECTION"))
	require.NotNil(t, client)

	err := client.Save("teste-key-2", "teste-value")
	require.NoError(t, err)

	res, err := client.GetTTL("teste-key-2")
	require.NoError(t, err)
	require.Equal(t, int64(600), res)
}

func Test_GetTTL_Expired(t *testing.T) {
	client := NewRedisClient(os.Getenv("REDIS_CONNECTION"))
	require.NotNil(t, client)

	err := client.Save("teste-key-2", "teste-value")
	require.NoError(t, err)

	_, err = client.GetTTL("teste-key-invalid")
	require.Error(t, err)
}
