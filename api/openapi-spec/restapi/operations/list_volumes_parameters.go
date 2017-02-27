package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListVolumesParams creates a new ListVolumesParams object
// with the default values initialized.
func NewListVolumesParams() ListVolumesParams {
	var ()
	return ListVolumesParams{}
}

// ListVolumesParams contains all the bound params for the list volumes operation
// typically these are obtained from a http.Request
//
// swagger:parameters ListVolumes
type ListVolumesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*Type of specified volume backend resource
	  Required: true
	  In: path
	*/
	ResourceType string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *ListVolumesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rResourceType, rhkResourceType, _ := route.Params.GetOK("resourceType")
	if err := o.bindResourceType(rResourceType, rhkResourceType, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListVolumesParams) bindResourceType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.ResourceType = raw

	return nil
}
