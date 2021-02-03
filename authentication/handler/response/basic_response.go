package response

// BasicResponse ...
type BasicResponse struct {
	Code  int         `json:"code"`
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}
