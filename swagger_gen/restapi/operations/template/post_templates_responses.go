// Code generated by go-swagger; DO NOT EDIT.

package template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/checkr/openmock/swagger_gen/models"
)

// PostTemplatesOKCode is the HTTP code returned for type PostTemplatesOK
const PostTemplatesOKCode int = 200

/*PostTemplatesOK returns the successfully posted templates

swagger:response postTemplatesOK
*/
type PostTemplatesOK struct {

	/*
	  In: Body
	*/
	Payload models.Mocks `json:"body,omitempty"`
}

// NewPostTemplatesOK creates PostTemplatesOK with default headers values
func NewPostTemplatesOK() *PostTemplatesOK {

	return &PostTemplatesOK{}
}

// WithPayload adds the payload to the post templates o k response
func (o *PostTemplatesOK) WithPayload(payload models.Mocks) *PostTemplatesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post templates o k response
func (o *PostTemplatesOK) SetPayload(payload models.Mocks) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostTemplatesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// PostTemplatesBadRequestCode is the HTTP code returned for type PostTemplatesBadRequest
const PostTemplatesBadRequestCode int = 400

/*PostTemplatesBadRequest if incoming templates were invalid

swagger:response postTemplatesBadRequest
*/
type PostTemplatesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostTemplatesBadRequest creates PostTemplatesBadRequest with default headers values
func NewPostTemplatesBadRequest() *PostTemplatesBadRequest {

	return &PostTemplatesBadRequest{}
}

// WithPayload adds the payload to the post templates bad request response
func (o *PostTemplatesBadRequest) WithPayload(payload *models.Error) *PostTemplatesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post templates bad request response
func (o *PostTemplatesBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostTemplatesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PostTemplatesDefault generic error response

swagger:response postTemplatesDefault
*/
type PostTemplatesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostTemplatesDefault creates PostTemplatesDefault with default headers values
func NewPostTemplatesDefault(code int) *PostTemplatesDefault {
	if code <= 0 {
		code = 500
	}

	return &PostTemplatesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post templates default response
func (o *PostTemplatesDefault) WithStatusCode(code int) *PostTemplatesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post templates default response
func (o *PostTemplatesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post templates default response
func (o *PostTemplatesDefault) WithPayload(payload *models.Error) *PostTemplatesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post templates default response
func (o *PostTemplatesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostTemplatesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
