package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/opensds/opensds/api/openapi-spec/models"
)

// NewCreateShareParams creates a new CreateShareParams object
// with the default values initialized.
func NewCreateShareParams() CreateShareParams {
	var ()
	return CreateShareParams{}
}

// CreateShareParams contains all the bound params for the create share operation
// typically these are obtained from a http.Request
//
// swagger:parameters CreateShare
type CreateShareParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*Type of specified share backend resource
	  Required: true
	  In: path
	*/
	ResourceType string
	/*Share request object
	  Required: true
	  In: body
	*/
	ShareRequest *models.ShareRequest
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *CreateShareParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rResourceType, rhkResourceType, _ := route.Params.GetOK("resourceType")
	if err := o.bindResourceType(rResourceType, rhkResourceType, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.ShareRequest
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("shareRequest", "body"))
			} else {
				res = append(res, errors.NewParseError("shareRequest", "body", "", err))
			}

		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.ShareRequest = &body
			}
		}

	} else {
		res = append(res, errors.Required("shareRequest", "body"))
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateShareParams) bindResourceType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.ResourceType = raw

	return nil
}
