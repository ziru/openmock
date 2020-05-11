// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ExpectHTTP expect HTTP
//
// swagger:model ExpectHTTP
type ExpectHTTP struct {

	// HTTP method to expect to trigger this behavior
	// Enum: [POST GET DELETE PUT OPTIONS HEAD]
	Method string `json:"method,omitempty"`

	// HTTP path to expect to trigger this behavior
	Path string `json:"path,omitempty"`
}

// Validate validates this expect HTTP
func (m *ExpectHTTP) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMethod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var expectHttpTypeMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["POST","GET","DELETE","PUT","OPTIONS","HEAD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		expectHttpTypeMethodPropEnum = append(expectHttpTypeMethodPropEnum, v)
	}
}

const (

	// ExpectHTTPMethodPOST captures enum value "POST"
	ExpectHTTPMethodPOST string = "POST"

	// ExpectHTTPMethodGET captures enum value "GET"
	ExpectHTTPMethodGET string = "GET"

	// ExpectHTTPMethodDELETE captures enum value "DELETE"
	ExpectHTTPMethodDELETE string = "DELETE"

	// ExpectHTTPMethodPUT captures enum value "PUT"
	ExpectHTTPMethodPUT string = "PUT"

	// ExpectHTTPMethodOPTIONS captures enum value "OPTIONS"
	ExpectHTTPMethodOPTIONS string = "OPTIONS"

	// ExpectHTTPMethodHEAD captures enum value "HEAD"
	ExpectHTTPMethodHEAD string = "HEAD"
)

// prop value enum
func (m *ExpectHTTP) validateMethodEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, expectHttpTypeMethodPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ExpectHTTP) validateMethod(formats strfmt.Registry) error {

	if swag.IsZero(m.Method) { // not required
		return nil
	}

	// value enum
	if err := m.validateMethodEnum("method", "body", m.Method); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExpectHTTP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExpectHTTP) UnmarshalBinary(b []byte) error {
	var res ExpectHTTP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
