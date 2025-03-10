// Code generated by go-swagger; DO NOT EDIT.

package template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ziru/openmock/swagger_gen/models"
)

// DeleteTemplatesNoContentCode is the HTTP code returned for type DeleteTemplatesNoContent
const DeleteTemplatesNoContentCode int = 204

/*DeleteTemplatesNoContent when successfully deleted

swagger:response deleteTemplatesNoContent
*/
type DeleteTemplatesNoContent struct {
}

// NewDeleteTemplatesNoContent creates DeleteTemplatesNoContent with default headers values
func NewDeleteTemplatesNoContent() *DeleteTemplatesNoContent {

	return &DeleteTemplatesNoContent{}
}

// WriteResponse to the client
func (o *DeleteTemplatesNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

/*DeleteTemplatesDefault server error

swagger:response deleteTemplatesDefault
*/
type DeleteTemplatesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteTemplatesDefault creates DeleteTemplatesDefault with default headers values
func NewDeleteTemplatesDefault(code int) *DeleteTemplatesDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteTemplatesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete templates default response
func (o *DeleteTemplatesDefault) WithStatusCode(code int) *DeleteTemplatesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete templates default response
func (o *DeleteTemplatesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete templates default response
func (o *DeleteTemplatesDefault) WithPayload(payload *models.Error) *DeleteTemplatesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete templates default response
func (o *DeleteTemplatesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteTemplatesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
