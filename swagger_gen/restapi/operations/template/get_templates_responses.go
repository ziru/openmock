// Code generated by go-swagger; DO NOT EDIT.

package template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ziru/openmock/swagger_gen/models"
)

// GetTemplatesOKCode is the HTTP code returned for type GetTemplatesOK
const GetTemplatesOKCode int = 200

/*GetTemplatesOK all the templates

swagger:response getTemplatesOK
*/
type GetTemplatesOK struct {

	/*
	  In: Body
	*/
	Payload models.Mocks `json:"body,omitempty"`
}

// NewGetTemplatesOK creates GetTemplatesOK with default headers values
func NewGetTemplatesOK() *GetTemplatesOK {

	return &GetTemplatesOK{}
}

// WithPayload adds the payload to the get templates o k response
func (o *GetTemplatesOK) WithPayload(payload models.Mocks) *GetTemplatesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get templates o k response
func (o *GetTemplatesOK) SetPayload(payload models.Mocks) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTemplatesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.Mocks{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetTemplatesDefault generic error response

swagger:response getTemplatesDefault
*/
type GetTemplatesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTemplatesDefault creates GetTemplatesDefault with default headers values
func NewGetTemplatesDefault(code int) *GetTemplatesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetTemplatesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get templates default response
func (o *GetTemplatesDefault) WithStatusCode(code int) *GetTemplatesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get templates default response
func (o *GetTemplatesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get templates default response
func (o *GetTemplatesDefault) WithPayload(payload *models.Error) *GetTemplatesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get templates default response
func (o *GetTemplatesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTemplatesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
