// Code generated by go-swagger; DO NOT EDIT.

package gardens

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kaatinga/plantbook/internal/api/models"
)

// GetUserGardenByIDOKCode is the HTTP code returned for type GetUserGardenByIDOK
const GetUserGardenByIDOKCode int = 200

/*GetUserGardenByIDOK Exists garden

swagger:response getUserGardenByIdOK
*/
type GetUserGardenByIDOK struct {
	/*The request id this is a response to

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Garden `json:"body,omitempty"`
}

// NewGetUserGardenByIDOK creates GetUserGardenByIDOK with default headers values
func NewGetUserGardenByIDOK() *GetUserGardenByIDOK {

	return &GetUserGardenByIDOK{}
}

// WithXRequestID adds the xRequestId to the get user garden by Id o k response
func (o *GetUserGardenByIDOK) WithXRequestID(xRequestID string) *GetUserGardenByIDOK {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get user garden by Id o k response
func (o *GetUserGardenByIDOK) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get user garden by Id o k response
func (o *GetUserGardenByIDOK) WithPayload(payload *models.Garden) *GetUserGardenByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user garden by Id o k response
func (o *GetUserGardenByIDOK) SetPayload(payload *models.Garden) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserGardenByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetUserGardenByIDDefault unexpected error

swagger:response getUserGardenByIdDefault
*/
type GetUserGardenByIDDefault struct {
	_statusCode int
	/*The request id this is a response to

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetUserGardenByIDDefault creates GetUserGardenByIDDefault with default headers values
func NewGetUserGardenByIDDefault(code int) *GetUserGardenByIDDefault {
	if code <= 0 {
		code = 500
	}

	return &GetUserGardenByIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get user garden by ID default response
func (o *GetUserGardenByIDDefault) WithStatusCode(code int) *GetUserGardenByIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get user garden by ID default response
func (o *GetUserGardenByIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithXRequestID adds the xRequestId to the get user garden by ID default response
func (o *GetUserGardenByIDDefault) WithXRequestID(xRequestID string) *GetUserGardenByIDDefault {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get user garden by ID default response
func (o *GetUserGardenByIDDefault) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get user garden by ID default response
func (o *GetUserGardenByIDDefault) WithPayload(payload *models.ErrorResponse) *GetUserGardenByIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user garden by ID default response
func (o *GetUserGardenByIDDefault) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserGardenByIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
