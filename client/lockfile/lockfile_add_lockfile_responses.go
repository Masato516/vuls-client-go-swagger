// Code generated by go-swagger; DO NOT EDIT.

package lockfile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Masato516/vuls-client-go-swagger/models"
)

// LockfileAddLockfileReader is a Reader for the LockfileAddLockfile structure.
type LockfileAddLockfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LockfileAddLockfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLockfileAddLockfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLockfileAddLockfileOK creates a LockfileAddLockfileOK with default headers values
func NewLockfileAddLockfileOK() *LockfileAddLockfileOK {
	return &LockfileAddLockfileOK{}
}

/*
LockfileAddLockfileOK describes a response with status code 200, with default header values.

OK response.
*/
type LockfileAddLockfileOK struct {
	Payload *models.LockfileAddLockfileResponseBody
}

// IsSuccess returns true when this lockfile add lockfile o k response has a 2xx status code
func (o *LockfileAddLockfileOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this lockfile add lockfile o k response has a 3xx status code
func (o *LockfileAddLockfileOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this lockfile add lockfile o k response has a 4xx status code
func (o *LockfileAddLockfileOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this lockfile add lockfile o k response has a 5xx status code
func (o *LockfileAddLockfileOK) IsServerError() bool {
	return false
}

// IsCode returns true when this lockfile add lockfile o k response a status code equal to that given
func (o *LockfileAddLockfileOK) IsCode(code int) bool {
	return code == 200
}

func (o *LockfileAddLockfileOK) Error() string {
	return fmt.Sprintf("[POST /v1/lockfile][%d] lockfileAddLockfileOK  %+v", 200, o.Payload)
}

func (o *LockfileAddLockfileOK) String() string {
	return fmt.Sprintf("[POST /v1/lockfile][%d] lockfileAddLockfileOK  %+v", 200, o.Payload)
}

func (o *LockfileAddLockfileOK) GetPayload() *models.LockfileAddLockfileResponseBody {
	return o.Payload
}

func (o *LockfileAddLockfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LockfileAddLockfileResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
