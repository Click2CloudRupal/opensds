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

// NewOperateVolumeParams creates a new OperateVolumeParams object
// with the default values initialized.
func NewOperateVolumeParams() OperateVolumeParams {
	var ()
	return OperateVolumeParams{}
}

// OperateVolumeParams contains all the bound params for the operate volume operation
// typically these are obtained from a http.Request
//
// swagger:parameters OperateVolume
type OperateVolumeParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*ID of specified volume
	  Required: true
	  In: path
	*/
	ID string
	/*Type of specified volume backend resource
	  Required: true
	  In: path
	*/
	ResourceType string
	/*Volume request object
	  Required: true
	  In: body
	*/
	VolumeRequest *models.VolumeRequest
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *OperateVolumeParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	rResourceType, rhkResourceType, _ := route.Params.GetOK("resourceType")
	if err := o.bindResourceType(rResourceType, rhkResourceType, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.VolumeRequest
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("volumeRequest", "body"))
			} else {
				res = append(res, errors.NewParseError("volumeRequest", "body", "", err))
			}

		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.VolumeRequest = &body
			}
		}

	} else {
		res = append(res, errors.Required("volumeRequest", "body"))
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *OperateVolumeParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.ID = raw

	return nil
}

func (o *OperateVolumeParams) bindResourceType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.ResourceType = raw

	return nil
}
