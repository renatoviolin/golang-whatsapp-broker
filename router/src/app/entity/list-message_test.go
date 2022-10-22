package entity

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NewListPayload_Valid(t *testing.T) {
	rows := []Row{
		{ID: "id1", Title: "title1"},
		{ID: "id2", Title: "title2"},
	}
	payload, err := NewListPayload("title", "body", rows, "5516993259256")
	require.NoError(t, err)
	require.Equal(t, "interactive", payload.Type)
	require.Equal(t, "Escolha uma opção:", payload.Interactive.Body.Text)
}

func Test_NewListPayload_Invalid(t *testing.T) {
	rows := []Row{
		{ID: "", Title: "title1"},
		{ID: "id2", Title: "title2"},
	}
	_, err := NewListPayload("title", "body", rows, "5516993259256")
	require.Error(t, err)

	rows2 := []Row{
		{ID: "id1", Title: "title1"},
		{ID: "id2", Title: "title2"},
	}
	_, err = NewListPayload("title", "", rows2, "5516993259256")
	require.Error(t, err)

	_, err = NewListPayload("title", "body", rows2, "")
	require.Error(t, err)
}
