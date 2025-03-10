// Code generated by go-swagger; DO NOT EDIT.

package template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ziru/openmock/swagger_gen/models"
)

// DeleteTemplateNoContentCode is the HTTP code returned for type DeleteTemplateNoContent
const DeleteTemplateNoContentCode int = 204

/*DeleteTemplateNoContent when successfully deleted

swagger:response deleteTemplateNoContent
*/
type DeleteTemplateNoContent struct {
}

// NewDeleteTemplateNoContent creates DeleteTemplateNoContent with default headers values
func NewDeleteTemplateNoContent() *DeleteTemplateNoContent {

	return &DeleteTemplateNoContent{}
}

// WriteResponse to the client
func (o *DeleteTemplateNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteTemplateBadRequestCode is the HTTP code returned for type DeleteTemplateBadRequest
const DeleteTemplateBadRequestCode int = 400

/*DeleteTemplateBadRequest invalid key

swagger:response deleteTemplateBadRequest
*/
type DeleteTemplateBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteTemplateBadRequest creates DeleteTemplateBadRequest with default headers values
func NewDeleteTemplateBadRequest() *DeleteTemplateBadRequest {

	return &DeleteTemplateBadRequest{}
}

// WithPayload adds the payload to the delete template bad request response
func (o *DeleteTemplateBadRequest) WithPayload(payload *models.Error) *DeleteTemplateBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete template bad request response
func (o *DeleteTemplateBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteTemplateBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteTemplateNotFoundCode is the HTTP code returned for type DeleteTemplateNotFound
const DeleteTemplateNotFoundCode int = 404

/*DeleteTemplateNotFound not found

swagger:response deleteTemplateNotFound
*/
type DeleteTemplateNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteTemplateNotFound creates DeleteTemplateNotFound with default headers values
func NewDeleteTemplateNotFound() *DeleteTemplateNotFound {

	return &DeleteTemplateNotFound{}
}

// WithPayload adds the payload to the delete template not found response
func (o *DeleteTemplateNotFound) WithPayload(payload *models.Error) *DeleteTemplateNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete template not found response
func (o *DeleteTemplateNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteTemplateNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteTemplateDefault server error

swagger:response deleteTemplateDefault
*/
type DeleteTemplateDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteTemplateDefault creates DeleteTemplateDefault with default headers values
func NewDeleteTemplateDefault(code int) *DeleteTemplateDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteTemplateDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete template default response
func (o *DeleteTemplateDefault) WithStatusCode(code int) *DeleteTemplateDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete template default response
func (o *DeleteTemplateDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete template default response
func (o *DeleteTemplateDefault) WithPayload(payload *models.Error) *DeleteTemplateDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete template default response
func (o *DeleteTemplateDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteTemplateDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
