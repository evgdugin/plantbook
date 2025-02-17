// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kaatinga/plantbook/internal/api/models"
)

// LogoutUserOKCode is the HTTP code returned for type LogoutUserOK
const LogoutUserOKCode int = 200

/*LogoutUserOK successful operation, set expired cookie

swagger:response logoutUserOK
*/
type LogoutUserOK struct {
	/*set expired cookie with name plantbook_token

	 */
	SetCookie string `json:"Set-Cookie"`
	/*The request id this is a response to

	 */
	XRequestID string `json:"X-Request-Id"`
}

// NewLogoutUserOK creates LogoutUserOK with default headers values
func NewLogoutUserOK() *LogoutUserOK {

	return &LogoutUserOK{}
}

// WithSetCookie adds the setCookie to the logout user o k response
func (o *LogoutUserOK) WithSetCookie(setCookie string) *LogoutUserOK {
	o.SetCookie = setCookie
	return o
}

// SetSetCookie sets the setCookie to the logout user o k response
func (o *LogoutUserOK) SetSetCookie(setCookie string) {
	o.SetCookie = setCookie
}

// WithXRequestID adds the xRequestId to the logout user o k response
func (o *LogoutUserOK) WithXRequestID(xRequestID string) *LogoutUserOK {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the logout user o k response
func (o *LogoutUserOK) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WriteResponse to the client
func (o *LogoutUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Set-Cookie

	setCookie := o.SetCookie
	if setCookie != "" {
		rw.Header().Set("Set-Cookie", setCookie)
	}

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*LogoutUserDefault unexpected error

swagger:response logoutUserDefault
*/
type LogoutUserDefault struct {
	_statusCode int
	/*The request id this is a response to

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewLogoutUserDefault creates LogoutUserDefault with default headers values
func NewLogoutUserDefault(code int) *LogoutUserDefault {
	if code <= 0 {
		code = 500
	}

	return &LogoutUserDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the logout user default response
func (o *LogoutUserDefault) WithStatusCode(code int) *LogoutUserDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the logout user default response
func (o *LogoutUserDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithXRequestID adds the xRequestId to the logout user default response
func (o *LogoutUserDefault) WithXRequestID(xRequestID string) *LogoutUserDefault {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the logout user default response
func (o *LogoutUserDefault) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the logout user default response
func (o *LogoutUserDefault) WithPayload(payload *models.ErrorResponse) *LogoutUserDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the logout user default response
func (o *LogoutUserDefault) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LogoutUserDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
