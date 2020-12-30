package gateway

// BasicReponse ...
type BasicReponse struct {
	Error error       `json:"error"`
	Data  interface{} `json:"data"`
}
