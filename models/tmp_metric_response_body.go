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

// TmpMetricResponseBody TmpMetricResponseBody
//
// TmpMetric describes a tmpMetric
// Example: {"createdAt":"2018-07-14T08:13:28Z","e":"","rc":"","rl":"","updatedAt":"2018-07-14T08:13:28Z"}
//
// swagger:model TmpMetricResponseBody
type TmpMetricResponseBody struct {

	// created time
	// Example: 2018-07-14T08:13:28Z
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"createdAt"`

	// E of tmpMetric
	// Required: true
	E *string `json:"e"`

	// RC of tmpMetric
	// Required: true
	Rc *string `json:"rc"`

	// RL of tmpMetric
	// Required: true
	Rl *string `json:"rl"`

	// updated time
	// Example: 2018-07-14T08:13:28Z
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updatedAt"`
}

// Validate validates this tmp metric response body
func (m *TmpMetricResponseBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateE(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRl(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TmpMetricResponseBody) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TmpMetricResponseBody) validateE(formats strfmt.Registry) error {

	if err := validate.Required("e", "body", m.E); err != nil {
		return err
	}

	return nil
}

func (m *TmpMetricResponseBody) validateRc(formats strfmt.Registry) error {

	if err := validate.Required("rc", "body", m.Rc); err != nil {
		return err
	}

	return nil
}

func (m *TmpMetricResponseBody) validateRl(formats strfmt.Registry) error {

	if err := validate.Required("rl", "body", m.Rl); err != nil {
		return err
	}

	return nil
}

func (m *TmpMetricResponseBody) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updatedAt", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this tmp metric response body based on context it is used
func (m *TmpMetricResponseBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TmpMetricResponseBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TmpMetricResponseBody) UnmarshalBinary(b []byte) error {
	var res TmpMetricResponseBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
