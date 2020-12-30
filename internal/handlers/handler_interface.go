package handlers

// HandlerInterface ...
type HandlerInterface interface {
	Init(cfg interface{}) error
	Run() error
}
