package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// DefaultResponse default response
// swagger:model DefaultResponse
type DefaultResponse struct {

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this default response
func (m *DefaultResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
