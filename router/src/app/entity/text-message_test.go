package entity

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NewTextPayload_Valid(t *testing.T) {
	payload, err := NewTextPayload("message", "5516993259256")
	require.NoError(t, err)
	require.Equal(t, "message", payload.Text.Body)
}

func Test_NewTextPayload_Invalid(t *testing.T) {
	_, err := NewTextPayload("message", "")
	require.Error(t, err)

	_, err = NewTextPayload("", "5516993259256")
	require.Error(t, err)
}
