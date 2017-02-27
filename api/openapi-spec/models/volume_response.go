package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// VolumeResponse volume response
// swagger:model VolumeResponse
type VolumeResponse struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// resource type
	ResourceType string `json:"resourceType,omitempty"`

	// size
	Size int32 `json:"size,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// volume type
	VolumeType string `json:"volumeType,omitempty"`
}

// Validate validates this volume response
func (m *VolumeResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
