package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/cilium/api/v1/models"
)

// GetEndpointIDOKCode is the HTTP code returned for type GetEndpointIDOK
const GetEndpointIDOKCode int = 200

/*GetEndpointIDOK Success

swagger:response getEndpointIdOK
*/
type GetEndpointIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Endpoint `json:"body,omitempty"`
}

// NewGetEndpointIDOK creates GetEndpointIDOK with default headers values
func NewGetEndpointIDOK() *GetEndpointIDOK {
	return &GetEndpointIDOK{}
}

// WithPayload adds the payload to the get endpoint Id o k response
func (o *GetEndpointIDOK) WithPayload(payload *models.Endpoint) *GetEndpointIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get endpoint Id o k response
func (o *GetEndpointIDOK) SetPayload(payload *models.Endpoint) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEndpointIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetEndpointIDInvalidCode is the HTTP code returned for type GetEndpointIDInvalid
const GetEndpointIDInvalidCode int = 400

/*GetEndpointIDInvalid Invalid endpoint ID format for specified type

swagger:response getEndpointIdInvalid
*/
type GetEndpointIDInvalid struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewGetEndpointIDInvalid creates GetEndpointIDInvalid with default headers values
func NewGetEndpointIDInvalid() *GetEndpointIDInvalid {
	return &GetEndpointIDInvalid{}
}

// WithPayload adds the payload to the get endpoint Id invalid response
func (o *GetEndpointIDInvalid) WithPayload(payload models.Error) *GetEndpointIDInvalid {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get endpoint Id invalid response
func (o *GetEndpointIDInvalid) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEndpointIDInvalid) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetEndpointIDNotFoundCode is the HTTP code returned for type GetEndpointIDNotFound
const GetEndpointIDNotFoundCode int = 404

/*GetEndpointIDNotFound Endpoint not found

swagger:response getEndpointIdNotFound
*/
type GetEndpointIDNotFound struct {
}

// NewGetEndpointIDNotFound creates GetEndpointIDNotFound with default headers values
func NewGetEndpointIDNotFound() *GetEndpointIDNotFound {
	return &GetEndpointIDNotFound{}
}

// WriteResponse to the client
func (o *GetEndpointIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}
