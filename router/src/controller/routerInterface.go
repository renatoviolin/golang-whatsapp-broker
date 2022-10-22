package controller

type IController interface {
	HandleMessageFromKafka(input []byte) error
}
