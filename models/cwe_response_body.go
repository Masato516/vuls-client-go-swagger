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

// CweResponseBody CweResponseBody
//
// Cwe describes a cwe
// Example: {"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"}
//
// swagger:model CweResponseBody
type CweResponseBody struct {

	// cweID of cwe
	// Example: CWE-416
	// Required: true
	CweID *string `json:"cweID"`

	// english summary of cwe
	// Example: english summary
	// Required: true
	English *string `json:"english"`

	// japanese summary of cwe
	// Example: japanese summary
	// Required: true
	Japanese *string `json:"japanese"`

	// owaspTopTen2017 of cwe
	// Example: 10
	OwaspTopTen2017 string `json:"owaspTopTen2017,omitempty"`

	// sourceType of cwe
	// Example: nvd
	// Required: true
	SourceType *string `json:"sourceType"`
}

// Validate validates this cwe response body
func (m *CweResponseBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCweID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnglish(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJapanese(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CweResponseBody) validateCweID(formats strfmt.Registry) error {

	if err := validate.Required("cweID", "body", m.CweID); err != nil {
		return err
	}

	return nil
}

func (m *CweResponseBody) validateEnglish(formats strfmt.Registry) error {

	if err := validate.Required("english", "body", m.English); err != nil {
		return err
	}

	return nil
}

func (m *CweResponseBody) validateJapanese(formats strfmt.Registry) error {

	if err := validate.Required("japanese", "body", m.Japanese); err != nil {
		return err
	}

	return nil
}

func (m *CweResponseBody) validateSourceType(formats strfmt.Registry) error {

	if err := validate.Required("sourceType", "body", m.SourceType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cwe response body based on context it is used
func (m *CweResponseBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CweResponseBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CweResponseBody) UnmarshalBinary(b []byte) error {
	var res CweResponseBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
