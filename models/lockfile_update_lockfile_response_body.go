// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LockfileUpdateLockfileResponseBody LockfileUpdateLockfileResponseBody
// Example: {"createdAt":"2018-07-14T08:13:28Z","fileContent":"","id":1,"libraryPkgs":[{"createdAt":"2018-07-14T08:13:28Z","id":1,"name":"package01","updatedAt":"2018-07-14T08:13:28Z","version":"1.0"},{"createdAt":"2018-07-14T08:13:28Z","id":1,"name":"package01","updatedAt":"2018-07-14T08:13:28Z","version":"1.0"}],"path":"/FutureVuls/package.json","server":{"createdAt":"2018-07-14T08:13:28Z","defaultUserId":1,"defaultUserName":"vuls user","hostUuid":"141df30a-ecd0-39f4-a8f4-1ef216a4b5f2","id":1,"lastScannedAt":"2018-07-14T08:13:28Z","lastUploadedAt":"2018-07-14T08:13:28Z","needKernelRestart":true,"osFamily":"centos","osVersion":"6","serverName":"server01","serverUuid":"abcdef12-ecd0-39f4-a8f4-1ef216a4b5f2","serverroleId":1,"serverroleName":"server_role01","tags":["Fugit repellendus illo.","Aperiam ipsa voluptate autem unde.","Fuga accusamus aut."],"updatedAt":"2018-07-14T08:13:28Z"},"updatedAt":"2018-07-14T08:13:28Z"}
//
// swagger:model LockfileUpdateLockfileResponseBody
type LockfileUpdateLockfileResponseBody struct {

	// created time of lockfile
	// Example: 2018-07-14T08:13:28Z
	// Required: true
	CreatedAt *string `json:"createdAt"`

	// FileContent of lockfile
	// Required: true
	FileContent *string `json:"fileContent"`

	// ID of lockfile
	// Example: 1
	// Required: true
	ID *int64 `json:"id"`

	// LibraryPkgs of lockfile
	// Example: [{"createdAt":"2018-07-14T08:13:28Z","id":1,"name":"package01","updatedAt":"2018-07-14T08:13:28Z","version":"1.0"},{"createdAt":"2018-07-14T08:13:28Z","id":1,"name":"package01","updatedAt":"2018-07-14T08:13:28Z","version":"1.0"},{"createdAt":"2018-07-14T08:13:28Z","id":1,"name":"package01","updatedAt":"2018-07-14T08:13:28Z","version":"1.0"}]
	LibraryPkgs []*LibraryPkgChildResponseBody `json:"libraryPkgs"`

	// Path of lockfile
	// Example: /FutureVuls/package.json
	// Required: true
	Path *string `json:"path"`

	// server
	Server *ServerChildResponseBody `json:"server,omitempty"`

	// updated time of lockfile
	// Example: 2018-07-14T08:13:28Z
	// Required: true
	UpdatedAt *string `json:"updatedAt"`
}

// Validate validates this lockfile update lockfile response body
func (m *LockfileUpdateLockfileResponseBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLibraryPkgs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServer(formats); err != nil {
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

func (m *LockfileUpdateLockfileResponseBody) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *LockfileUpdateLockfileResponseBody) validateFileContent(formats strfmt.Registry) error {

	if err := validate.Required("fileContent", "body", m.FileContent); err != nil {
		return err
	}

	return nil
}

func (m *LockfileUpdateLockfileResponseBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *LockfileUpdateLockfileResponseBody) validateLibraryPkgs(formats strfmt.Registry) error {
	if swag.IsZero(m.LibraryPkgs) { // not required
		return nil
	}

	for i := 0; i < len(m.LibraryPkgs); i++ {
		if swag.IsZero(m.LibraryPkgs[i]) { // not required
			continue
		}

		if m.LibraryPkgs[i] != nil {
			if err := m.LibraryPkgs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("libraryPkgs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("libraryPkgs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LockfileUpdateLockfileResponseBody) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

func (m *LockfileUpdateLockfileResponseBody) validateServer(formats strfmt.Registry) error {
	if swag.IsZero(m.Server) { // not required
		return nil
	}

	if m.Server != nil {
		if err := m.Server.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("server")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("server")
			}
			return err
		}
	}

	return nil
}

func (m *LockfileUpdateLockfileResponseBody) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updatedAt", "body", m.UpdatedAt); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this lockfile update lockfile response body based on the context it is used
func (m *LockfileUpdateLockfileResponseBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLibraryPkgs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LockfileUpdateLockfileResponseBody) contextValidateLibraryPkgs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LibraryPkgs); i++ {

		if m.LibraryPkgs[i] != nil {
			if err := m.LibraryPkgs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("libraryPkgs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("libraryPkgs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LockfileUpdateLockfileResponseBody) contextValidateServer(ctx context.Context, formats strfmt.Registry) error {

	if m.Server != nil {
		if err := m.Server.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("server")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("server")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LockfileUpdateLockfileResponseBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LockfileUpdateLockfileResponseBody) UnmarshalBinary(b []byte) error {
	var res LockfileUpdateLockfileResponseBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
