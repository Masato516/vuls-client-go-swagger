// Code generated by go-swagger; DO NOT EDIT.

package health

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// HealthHealthReader is a Reader for the HealthHealth structure.
type HealthHealthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HealthHealthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHealthHealthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewHealthHealthOK creates a HealthHealthOK with default headers values
func NewHealthHealthOK() *HealthHealthOK {
	return &HealthHealthOK{}
}

/*
HealthHealthOK describes a response with status code 200, with default header values.

OK response.
*/
type HealthHealthOK struct {
	Payload bool
}

// IsSuccess returns true when this health health o k response has a 2xx status code
func (o *HealthHealthOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this health health o k response has a 3xx status code
func (o *HealthHealthOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this health health o k response has a 4xx status code
func (o *HealthHealthOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this health health o k response has a 5xx status code
func (o *HealthHealthOK) IsServerError() bool {
	return false
}

// IsCode returns true when this health health o k response a status code equal to that given
func (o *HealthHealthOK) IsCode(code int) bool {
	return code == 200
}

func (o *HealthHealthOK) Error() string {
	return fmt.Sprintf("[GET /healths][%d] healthHealthOK  %+v", 200, o.Payload)
}

func (o *HealthHealthOK) String() string {
	return fmt.Sprintf("[GET /healths][%d] healthHealthOK  %+v", 200, o.Payload)
}

func (o *HealthHealthOK) GetPayload() bool {
	return o.Payload
}

func (o *HealthHealthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
