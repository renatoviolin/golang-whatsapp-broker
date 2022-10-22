package repository

import (
	"os"
	"testing"
	"whatsapp-client/dto"
	"whatsapp-client/util"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func init() {
	util.LoadVars()
}

func Test_Create(t *testing.T) {
	mongoClient, err := NewMongoClient(os.Getenv("MONGO_CONNECTION"), os.Getenv("MONGO_DATABASE"), 5)
	require.NoError(t, err)

	repository := NewMessageRepository(mongoClient, os.Getenv("DIALOG_COLLECTION"))
	require.NotNil(t, repository)
	id := uuid.New().String()
	waID := "1234"
	payload := dto.MessageRepository{
		Id:     id,
		WaID:   waID,
		Type:   "response",
		Body:   "body",
		IsRead: false,
	}

	err = repository.Create(payload)
	require.NoError(t, err)
}

func Test_FindByWaID(t *testing.T) {
	mongoClient, err := NewMongoClient(os.Getenv("MONGO_CONNECTION"), os.Getenv("MONGO_DATABASE"), 5)
	require.NoError(t, err)

	repository := NewMessageRepository(mongoClient, os.Getenv("DIALOG_COLLECTION"))
	require.NotNil(t, repository)

	allMessages, err := repository.FindAll("1234")
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(allMessages), 1)
}

func Test_MarkAsRead(t *testing.T) {
	mongoClient, err := NewMongoClient(os.Getenv("MONGO_CONNECTION"), os.Getenv("MONGO_DATABASE"), 5)
	require.NoError(t, err)

	repository := NewMessageRepository(mongoClient, os.Getenv("DIALOG_COLLECTION"))
	require.NotNil(t, repository)
	id := uuid.New().String()
	waID := "12345"
	payload := dto.MessageRepository{
		Id:     id,
		WaID:   waID,
		Type:   "response",
		Body:   "body",
		IsRead: false,
	}

	err = repository.Create(payload)
	require.NoError(t, err)

	err = repository.MarkAsRead(id)
	require.NoError(t, err)

	err = repository.MarkAsRead("id")
	require.Error(t, err)
}
