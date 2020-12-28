package authentication

import (
	"talktalk/handlers"
)

const (
	// ErrorAlreadySignedUp ...
	ErrorAlreadySignedUp int = 1000
	// ErrorFailedToParseSignUpRequest ...
	ErrorFailedToParseSignUpRequest int = 1001
)

var messages = map[int]string{
	0:                               "OK",
	ErrorAlreadySignedUp:            "user already signed up",
	ErrorFailedToParseSignUpRequest: "failed to parse sign-up request",
}

// UserSignUpResponse ...
type UserSignUpResponse struct {
	handlers.BasicJSONResponse
}
