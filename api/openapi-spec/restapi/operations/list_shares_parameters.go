package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListSharesParams creates a new ListSharesParams object
// with the default values initialized.
func NewListSharesParams() ListSharesParams {
	var ()
	return ListSharesParams{}
}

// ListSharesParams contains all the bound params for the list shares operation
// typically these are obtained from a http.Request
//
// swagger:parameters ListShares
type ListSharesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*Type of specified share backend resource
	  Required: true
	  In: path
	*/
	ResourceType string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *ListSharesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
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

func (o *ListSharesParams) bindResourceType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.ResourceType = raw

	return nil
}
