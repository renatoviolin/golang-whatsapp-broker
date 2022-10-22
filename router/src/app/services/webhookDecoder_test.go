package services

import (
	"broker/dto"
	"broker/util"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func init() {
	util.LoadVars()
}

func Test_InvalidPayload(t *testing.T) {
	decoder := NewWebhookDecoder()

	var webhookInput1 dto.WebhookInput
	bytes := util.GetPayloadInvalid()
	_ = json.Unmarshal(bytes, &webhookInput1)
	_, err := decoder.DecodeWebhook(webhookInput1)
	require.Error(t, err)
}

func Test_WaID(t *testing.T) {
	all_payloads := [][]byte{util.GetPayload1(), util.GetPayload2(), util.GetPayload3(), util.GetPayloadInvalidAgentID(), util.GetPayloadResponseButton(), util.GetPayloadSair(), util.GetPayloadStatus1(), util.GetPayloadStatus2()}
	decoder := NewWebhookDecoder()
	for _, v := range all_payloads {
		var webhookInput dto.WebhookInput
		err := json.Unmarshal(v, &webhookInput)
		require.NoError(t, err)

		output, err := decoder.DecodeWebhook(webhookInput)
		require.NoError(t, err)
		require.Equal(t, "5516993259256", output.WaID)

	}
}

func Test_Text(t *testing.T) {
	decoder := NewWebhookDecoder()

	var webhookInput1 dto.WebhookInput
	bytes := util.GetPayloadSair()
	_ = json.Unmarshal(bytes, &webhookInput1)
	output, err := decoder.DecodeWebhook(webhookInput1)
	require.NoError(t, err)
	require.Equal(t, "sair", output.Text)
	require.Equal(t, "text", output.MessageType)

	var webhookInput2 dto.WebhookInput
	bytes = util.GetPayload1()
	_ = json.Unmarshal(bytes, &webhookInput2)
	output, err = decoder.DecodeWebhook(webhookInput2)
	require.NoError(t, err)
	require.Equal(t, "olá, gostaria de ser atendido", output.Text)
	require.Equal(t, "text", output.MessageType)

	var webhookInput3 dto.WebhookInput
	bytes = util.GetPayload3()
	_ = json.Unmarshal(bytes, &webhookInput3)
	output, err = decoder.DecodeWebhook(webhookInput3)
	require.NoError(t, err)
	require.Equal(t, "mensagem do Renato para o agent_test", output.Text)
	require.Equal(t, "text", output.MessageType)

	var webhookInput4 dto.WebhookInput
	bytes = util.GetPayloadStatus1()
	_ = json.Unmarshal(bytes, &webhookInput4)
	output, err = decoder.DecodeWebhook(webhookInput4)
	require.NoError(t, err)
	require.Equal(t, "", output.Text)
	require.Equal(t, "status", output.MessageType)
}

func Test_Status(t *testing.T) {
	decoder := NewWebhookDecoder()

	var webhookInput1 dto.WebhookInput
	bytes := util.GetPayloadStatus1()
	_ = json.Unmarshal(bytes, &webhookInput1)
	output, err := decoder.DecodeWebhook(webhookInput1)
	require.NoError(t, err)
	require.Equal(t, "sent", output.Status)
	require.Equal(t, "", output.Text)
	require.Equal(t, "status", output.MessageType)

	var webhookInput2 dto.WebhookInput
	bytes = util.GetPayloadStatus2()
	_ = json.Unmarshal(bytes, &webhookInput2)
	output, err = decoder.DecodeWebhook(webhookInput2)
	require.NoError(t, err)
	require.Equal(t, "delivered", output.Status)
	require.Equal(t, "", output.Text)
	require.Equal(t, "status", output.MessageType)
}

func Test_ListReply(t *testing.T) {
	decoder := NewWebhookDecoder()

	var webhookInput1 dto.WebhookInput
	bytes := util.GetPayloadInvalidAgentID()
	_ = json.Unmarshal(bytes, &webhookInput1)
	output, err := decoder.DecodeWebhook(webhookInput1)
	require.NoError(t, err)
	require.Equal(t, "agent_xxx", output.ListReply.Id)
	require.Equal(t, "Agent X", output.ListReply.Title)
	require.Equal(t, "", output.Text)
	require.Equal(t, "list_reply", output.MessageType)

	var webhookInput2 dto.WebhookInput
	bytes = util.GetPayload2()
	_ = json.Unmarshal(bytes, &webhookInput2)
	output, err = decoder.DecodeWebhook(webhookInput2)
	require.NoError(t, err)
	require.Equal(t, "agent_test", output.ListReply.Id)
	require.Equal(t, "Agent A", output.ListReply.Title)
	require.Equal(t, "", output.Text)
	require.Equal(t, "list_reply", output.MessageType)
}

func Test_Error(t *testing.T) {
	decoder := NewWebhookDecoder()

	var webhookInput1 dto.WebhookInput
	bytes := util.GetPayloadError()
	_ = json.Unmarshal(bytes, &webhookInput1)
	output, err := decoder.DecodeWebhook(webhookInput1)
	require.NoError(t, err)
	require.Equal(t, "Message failed to send because more than 24 hours have passed since the customer last replied to this number.", output.Error)
}
