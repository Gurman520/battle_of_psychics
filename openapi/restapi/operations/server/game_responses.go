// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

// GameOKCode is the HTTP code returned for type GameOK
const GameOKCode int = 200

/*
GameOK success and returns some html text

swagger:response gameOK
*/
type GameOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGameOK creates GameOK with default headers values
func NewGameOK() *GameOK {

	return &GameOK{}
}

// WithPayload adds the payload to the game o k response
func (o *GameOK) WithPayload(payload string) *GameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the game o k response
func (o *GameOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

func (o *GameOK) GameResponder() {}

// GameInternalServerErrorCode is the HTTP code returned for type GameInternalServerError
const GameInternalServerErrorCode int = 500

/*
GameInternalServerError Internal server error

swagger:response gameInternalServerError
*/
type GameInternalServerError struct {
}

// NewGameInternalServerError creates GameInternalServerError with default headers values
func NewGameInternalServerError() *GameInternalServerError {

	return &GameInternalServerError{}
}

// WriteResponse to the client
func (o *GameInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}

func (o *GameInternalServerError) GameResponder() {}

type GameNotImplementedResponder struct {
	middleware.Responder
}

func (*GameNotImplementedResponder) GameResponder() {}

func GameNotImplemented() GameResponder {
	return &GameNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.Game has not yet been implemented",
		),
	}
}

type GameResponder interface {
	middleware.Responder
	GameResponder()
}
