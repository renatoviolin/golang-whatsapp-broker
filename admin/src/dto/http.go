package dto

import "time"

type CreateAgentInput struct {
	Name       string `json:"name"`
	ReadTopic  string `json:"read_topic"`
	WriteTopic string `json:"write_topic"`
	ErrorTopic string `json:"error_topic"`
}

type UpdateAgentInput struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	ReadTopic  string `json:"read_topic"`
	WriteTopic string `json:"write_topic"`
	ErrorTopic string `json:"error_topic"`
}

type AgentOutput struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	ReadTopic  string    `json:"read_topic"`
	WriteTopic string    `json:"write_topic"`
	ErrorTopic string    `json:"error_topic"`
	CreatedAt  time.Time `json:"created_at"`
}
