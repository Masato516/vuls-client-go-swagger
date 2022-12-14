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

// LockfileAddLockfileRequestBody LockfileAddLockfileRequestBody
// Example: {"fileContent":"","path":"/FutureVuls/package.json","serverID":1}
//
// swagger:model LockfileAddLockfileRequestBody
type LockfileAddLockfileRequestBody struct {

	// fileContent of the lockfile
	// Required: true
	FileContent *string `json:"fileContent"`

	// Path of lockfile
	// Example: /FutureVuls/package.json
	// Required: true
	Path *string `json:"path"`

	// ServerID
	// Example: 1
	// Required: true
	ServerID *int64 `json:"serverID"`
}

// Validate validates this lockfile add lockfile request body
func (m *LockfileAddLockfileRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFileContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LockfileAddLockfileRequestBody) validateFileContent(formats strfmt.Registry) error {

	if err := validate.Required("fileContent", "body", m.FileContent); err != nil {
		return err
	}

	return nil
}

func (m *LockfileAddLockfileRequestBody) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

func (m *LockfileAddLockfileRequestBody) validateServerID(formats strfmt.Registry) error {

	if err := validate.Required("serverID", "body", m.ServerID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this lockfile add lockfile request body based on context it is used
func (m *LockfileAddLockfileRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LockfileAddLockfileRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LockfileAddLockfileRequestBody) UnmarshalBinary(b []byte) error {
	var res LockfileAddLockfileRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
