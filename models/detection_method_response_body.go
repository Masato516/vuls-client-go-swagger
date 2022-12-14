// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DetectionMethodResponseBody DetectionMethod ResponseBody
// Example: {"name":"vuls","reliabilityScore":100}
//
// swagger:model DetectionMethod ResponseBody
type DetectionMethodResponseBody struct {

	// Detection Method Name
	// Example: vuls
	// Required: true
	Name *string `json:"name"`

	// ReliabilityScore
	// Example: 100
	// Required: true
	ReliabilityScore *int64 `json:"reliabilityScore"`
}

// Validate validates this detection method response body
func (m *DetectionMethodResponseBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReliabilityScore(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DetectionMethodResponseBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *DetectionMethodResponseBody) validateReliabilityScore(formats strfmt.Registry) error {

	if err := validate.Required("reliabilityScore", "body", m.ReliabilityScore); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this detection method response body based on context it is used
func (m *DetectionMethodResponseBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DetectionMethodResponseBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DetectionMethodResponseBody) UnmarshalBinary(b []byte) error {
	var res DetectionMethodResponseBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
