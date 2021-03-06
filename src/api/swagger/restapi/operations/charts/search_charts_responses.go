package charts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kubernetes-helm/monocular/src/api/swagger/models"
)

/*SearchChartsOK repositories response

swagger:response searchChartsOK
*/
type SearchChartsOK struct {

	// In: body
	Payload *models.ResourceArrayData `json:"body,omitempty"`
}

// NewSearchChartsOK creates SearchChartsOK with default headers values
func NewSearchChartsOK() *SearchChartsOK {
	return &SearchChartsOK{}
}

// WithPayload adds the payload to the search charts o k response
func (o *SearchChartsOK) WithPayload(payload *models.ResourceArrayData) *SearchChartsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the search charts o k response
func (o *SearchChartsOK) SetPayload(payload *models.ResourceArrayData) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SearchChartsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*SearchChartsDefault unexpected error

swagger:response searchChartsDefault
*/
type SearchChartsDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewSearchChartsDefault creates SearchChartsDefault with default headers values
func NewSearchChartsDefault(code int) *SearchChartsDefault {
	if code <= 0 {
		code = 500
	}

	return &SearchChartsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the search charts default response
func (o *SearchChartsDefault) WithStatusCode(code int) *SearchChartsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the search charts default response
func (o *SearchChartsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the search charts default response
func (o *SearchChartsDefault) WithPayload(payload *models.Error) *SearchChartsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the search charts default response
func (o *SearchChartsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SearchChartsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
