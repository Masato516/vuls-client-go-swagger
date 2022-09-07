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

// PkgCpeListResponseBody PkgCpeListResponseBody
// Example: {"applicationName":"Application","cpeFS":"cpe:2.3:a:clamav:clamav:*:*:*:*:*:*:*:*","cpeID":1,"cpeURI":"cpe:/a:cisco:ios:15.2","createdAt":"2018-07-14T08:13:28Z","githubPkgID":1,"libraryPath":"/FutureVuls/go.sum","libraryPkgID":1,"name":"package01","needRestartProcs":[{"initSystem":"initSystem","path":"path","pid":"12","serviceName":"serviceName"},{"initSystem":"initSystem","path":"path","pid":"12","serviceName":"serviceName"}],"newRelease":"new release","newVersion":"2.0","notFixedYet":true,"ossLicense":"ossLicense","pkgID":1,"release":"release","repository":"repository","serverID":1,"serverName":"server01","serverUuid":"abcdef12-ecd0-39f4-a8f4-1ef216a4b5f2","updatedAt":"2018-07-14T08:13:28Z","version":"1.0","wordpressPackageStatus":"wordpressPackageStatus","wordpressPkgID":1}
//
// swagger:model PkgCpeListResponseBody
type PkgCpeListResponseBody struct {

	// ApplicationName of library package
	// Example: Application
	ApplicationName string `json:"applicationName,omitempty"`

	// Cpe FS of cpe
	// Example: cpe:2.3:a:clamav:clamav:*:*:*:*:*:*:*:*
	CpeFS string `json:"cpeFS,omitempty"`

	// CpeID of cpe
	// Example: 1
	CpeID int64 `json:"cpeID,omitempty"`

	// Cpe URI of cpe
	// Example: cpe:/a:cisco:ios:15.2
	CpeURI string `json:"cpeURI,omitempty"`

	// crated time of package or cpe
	// Example: 2018-07-14T08:13:28Z
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"createdAt"`

	// githubPKGID of github pkg
	// Example: 1
	GithubPkgID int64 `json:"githubPkgID,omitempty"`

	// LibraryPath of library package
	// Example: /FutureVuls/go.sum
	LibraryPath string `json:"libraryPath,omitempty"`

	// libraryPKGID of library pkg
	// Example: 1
	LibraryPkgID int64 `json:"libraryPkgID,omitempty"`

	// Name of package or cpe
	// Example: package01
	// Required: true
	Name *string `json:"name"`

	// NeedRestartProcess list of package
	// Example: [{"initSystem":"initSystem","path":"path","pid":"12","serviceName":"serviceName"},{"initSystem":"initSystem","path":"path","pid":"12","serviceName":"serviceName"},{"initSystem":"initSystem","path":"path","pid":"12","serviceName":"serviceName"},{"initSystem":"initSystem","path":"path","pid":"12","serviceName":"serviceName"}]
	NeedRestartProcs []*NeedRestartProcResponseBody `json:"needRestartProcs"`

	// New Release of package
	// Example: new release
	NewRelease string `json:"newRelease,omitempty"`

	// New Version of package
	// Example: 2.0
	NewVersion string `json:"newVersion,omitempty"`

	// Flag of Not fixed yet of package
	// Example: true
	NotFixedYet bool `json:"notFixedYet,omitempty"`

	// ossLicense of library package
	// Example: ossLicense
	OssLicense string `json:"ossLicense,omitempty"`

	// Package ID of package
	// Example: 1
	PkgID int64 `json:"pkgID,omitempty"`

	// Release of package
	// Example: release
	Release string `json:"release,omitempty"`

	// Repository of package
	// Example: repository
	Repository string `json:"repository,omitempty"`

	// ServerID of package or cpe
	// Example: 1
	// Required: true
	ServerID *int64 `json:"serverID"`

	// ServerName of package or cpe
	// Example: server01
	// Required: true
	ServerName *string `json:"serverName"`

	// ServerUUID of package or cpe
	// Example: abcdef12-ecd0-39f4-a8f4-1ef216a4b5f2
	// Required: true
	// Format: uuid
	ServerUUID *strfmt.UUID `json:"serverUuid"`

	// updated time of package or cpe
	// Example: 2018-07-14T08:13:28Z
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updatedAt"`

	// Version of package or cpe
	// Example: 1.0
	// Required: true
	Version *string `json:"version"`

	// WordpressPackageStatus of wordpress package
	// Example: wordpressPackageStatus
	WordpressPackageStatus string `json:"wordpressPackageStatus,omitempty"`

	// wordpressPKGID of wordpress pkg
	// Example: 1
	WordpressPkgID int64 `json:"wordpressPkgID,omitempty"`
}

// Validate validates this pkg cpe list response body
func (m *PkgCpeListResponseBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNeedRestartProcs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PkgCpeListResponseBody) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PkgCpeListResponseBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PkgCpeListResponseBody) validateNeedRestartProcs(formats strfmt.Registry) error {
	if swag.IsZero(m.NeedRestartProcs) { // not required
		return nil
	}

	for i := 0; i < len(m.NeedRestartProcs); i++ {
		if swag.IsZero(m.NeedRestartProcs[i]) { // not required
			continue
		}

		if m.NeedRestartProcs[i] != nil {
			if err := m.NeedRestartProcs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("needRestartProcs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("needRestartProcs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PkgCpeListResponseBody) validateServerID(formats strfmt.Registry) error {

	if err := validate.Required("serverID", "body", m.ServerID); err != nil {
		return err
	}

	return nil
}

func (m *PkgCpeListResponseBody) validateServerName(formats strfmt.Registry) error {

	if err := validate.Required("serverName", "body", m.ServerName); err != nil {
		return err
	}

	return nil
}

func (m *PkgCpeListResponseBody) validateServerUUID(formats strfmt.Registry) error {

	if err := validate.Required("serverUuid", "body", m.ServerUUID); err != nil {
		return err
	}

	if err := validate.FormatOf("serverUuid", "body", "uuid", m.ServerUUID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PkgCpeListResponseBody) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updatedAt", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PkgCpeListResponseBody) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this pkg cpe list response body based on the context it is used
func (m *PkgCpeListResponseBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNeedRestartProcs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PkgCpeListResponseBody) contextValidateNeedRestartProcs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NeedRestartProcs); i++ {

		if m.NeedRestartProcs[i] != nil {
			if err := m.NeedRestartProcs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("needRestartProcs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("needRestartProcs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PkgCpeListResponseBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PkgCpeListResponseBody) UnmarshalBinary(b []byte) error {
	var res PkgCpeListResponseBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
