package dto

type Agent struct {
	ID         int
	Name       string
	ReadTopic  string
	WriteTopic string
	ErrorTopic string
	Active     int
}
