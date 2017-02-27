package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/opensds/opensds/api/openapi-spec/models"
)

// HTTP code for type CreateShareOK
const CreateShareOKCode int = 200

/*CreateShareOK OK

swagger:response createShareOK
*/
type CreateShareOK struct {

	/*
	  In: Body
	*/
	Payload *models.ShareResponse `json:"body,omitempty"`
}

// NewCreateShareOK creates CreateShareOK with default headers values
func NewCreateShareOK() *CreateShareOK {
	return &CreateShareOK{}
}

// WithPayload adds the payload to the create share o k response
func (o *CreateShareOK) WithPayload(payload *models.ShareResponse) *CreateShareOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create share o k response
func (o *CreateShareOK) SetPayload(payload *models.ShareResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateShareOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
