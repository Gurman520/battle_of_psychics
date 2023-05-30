// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

// StartOKCode is the HTTP code returned for type StartOK
const StartOKCode int = 200

/*
StartOK success and returns some html text

swagger:response startOK
*/
type StartOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewStartOK creates StartOK with default headers values
func NewStartOK() *StartOK {

	return &StartOK{}
}

// WithPayload adds the payload to the start o k response
func (o *StartOK) WithPayload(payload string) *StartOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the start o k response
func (o *StartOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StartOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

func (o *StartOK) StartResponder() {}

// StartInternalServerErrorCode is the HTTP code returned for type StartInternalServerError
const StartInternalServerErrorCode int = 500

/*
StartInternalServerError Internal server error

swagger:response startInternalServerError
*/
type StartInternalServerError struct {
}

// NewStartInternalServerError creates StartInternalServerError with default headers values
func NewStartInternalServerError() *StartInternalServerError {

	return &StartInternalServerError{}
}

// WriteResponse to the client
func (o *StartInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}

func (o *StartInternalServerError) StartResponder() {}

type StartNotImplementedResponder struct {
	middleware.Responder
}

func (*StartNotImplementedResponder) StartResponder() {}

func StartNotImplemented() StartResponder {
	return &StartNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.Start has not yet been implemented",
		),
	}
}

type StartResponder interface {
	middleware.Responder
	StartResponder()
}
