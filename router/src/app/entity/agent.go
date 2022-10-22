package entity

type AgentConfig struct {
	Name          string
	ID            string
	ReadTopic     string
	WriteTopic    string
	ErrorTopic    string
	ConsumerGroup string
}

func NewAgentConfig(name, id, readTopic, writeTopic, ErrorTopic, consumerGroup string) AgentConfig {
	return AgentConfig{
		Name:          name,
		ID:            id,
		ReadTopic:     readTopic,
		WriteTopic:    writeTopic,
		ErrorTopic:    ErrorTopic,
		ConsumerGroup: consumerGroup,
	}
}
