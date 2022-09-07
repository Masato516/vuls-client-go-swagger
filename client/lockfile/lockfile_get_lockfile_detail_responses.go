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

// LockfileGetLockfileDetailReader is a Reader for the LockfileGetLockfileDetail structure.
type LockfileGetLockfileDetailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LockfileGetLockfileDetailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLockfileGetLockfileDetailOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLockfileGetLockfileDetailOK creates a LockfileGetLockfileDetailOK with default headers values
func NewLockfileGetLockfileDetailOK() *LockfileGetLockfileDetailOK {
	return &LockfileGetLockfileDetailOK{}
}

/*
LockfileGetLockfileDetailOK describes a response with status code 200, with default header values.

OK response.
*/
type LockfileGetLockfileDetailOK struct {
	Payload *models.LockfileGetLockfileDetailResponseBody
}

// IsSuccess returns true when this lockfile get lockfile detail o k response has a 2xx status code
func (o *LockfileGetLockfileDetailOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this lockfile get lockfile detail o k response has a 3xx status code
func (o *LockfileGetLockfileDetailOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this lockfile get lockfile detail o k response has a 4xx status code
func (o *LockfileGetLockfileDetailOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this lockfile get lockfile detail o k response has a 5xx status code
func (o *LockfileGetLockfileDetailOK) IsServerError() bool {
	return false
}

// IsCode returns true when this lockfile get lockfile detail o k response a status code equal to that given
func (o *LockfileGetLockfileDetailOK) IsCode(code int) bool {
	return code == 200
}

func (o *LockfileGetLockfileDetailOK) Error() string {
	return fmt.Sprintf("[GET /v1/lockfile/{lockfileID}][%d] lockfileGetLockfileDetailOK  %+v", 200, o.Payload)
}

func (o *LockfileGetLockfileDetailOK) String() string {
	return fmt.Sprintf("[GET /v1/lockfile/{lockfileID}][%d] lockfileGetLockfileDetailOK  %+v", 200, o.Payload)
}

func (o *LockfileGetLockfileDetailOK) GetPayload() *models.LockfileGetLockfileDetailResponseBody {
	return o.Payload
}

func (o *LockfileGetLockfileDetailOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LockfileGetLockfileDetailResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
