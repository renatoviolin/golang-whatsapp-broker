package controller

type IController interface {
	HandleMessage(input []byte) error
}
