package app

import "errors"

var (
	ErrNotFoundSession = errors.New("a session with such a token was not found")
	ErrIncorrectApiKey = errors.New("incorrect api key auth")
	ErrCallBan         = errors.New("cannot call a method without first calling /Conceive")
)
