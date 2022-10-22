package agent

type AgentConfig struct {
	ReadTopic     string
	WriteTopic    string
	ErrorTopic    string
	ConsumerGroup string
}

func NewAgentConfig(readTopic, writeTopic, ErrorTopic, consumerGroup string) AgentConfig {
	return AgentConfig{
		ReadTopic:     readTopic,
		WriteTopic:    writeTopic,
		ErrorTopic:    ErrorTopic,
		ConsumerGroup: consumerGroup,
	}
}
