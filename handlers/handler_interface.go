package handlers

// HandlerInterface ...
type HandlerInterface interface {
	Init(config interface{}) error
	Run() error
}
