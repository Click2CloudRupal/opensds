package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/opensds/opensds/api/openapi-spec/models"
)

// HTTP code for type ListVolumesOK
const ListVolumesOKCode int = 200

/*ListVolumesOK OK

swagger:response listVolumesOK
*/
type ListVolumesOK struct {

	/*
	  In: Body
	*/
	Payload []*models.VolumeResponse `json:"body,omitempty"`
}

// NewListVolumesOK creates ListVolumesOK with default headers values
func NewListVolumesOK() *ListVolumesOK {
	return &ListVolumesOK{}
}

// WithPayload adds the payload to the list volumes o k response
func (o *ListVolumesOK) WithPayload(payload []*models.VolumeResponse) *ListVolumesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list volumes o k response
func (o *ListVolumesOK) SetPayload(payload []*models.VolumeResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListVolumesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.VolumeResponse, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
