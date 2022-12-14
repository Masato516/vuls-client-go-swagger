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

// PkgCpeGetPkgDetailResponseBody PkgCpeGetPkgDetailResponseBody
// Example: {"affectedProcs":[{"name":"apache","pid":"12"},{"name":"apache","pid":"12"},{"name":"apache","pid":"12"}],"createdAt":"2018-07-14T08:13:28Z","id":1,"name":"package01","needRestartProcs":[{"initSystem":"initSystem","path":"path","pid":"12","serviceName":"serviceName"},{"initSystem":"initSystem","path":"path","pid":"12","serviceName":"serviceName"},{"initSystem":"initSystem","path":"path","pid":"12","serviceName":"serviceName"},{"initSystem":"initSystem","path":"path","pid":"12","serviceName":"serviceName"}],"newRelease":"new release","newVersion":"2.0","packageStatuses":{"bash":"resolved"},"release":"release","repository":"repository","server":{"createdAt":"2018-07-14T08:13:28Z","defaultUserId":1,"defaultUserName":"vuls user","hostUuid":"141df30a-ecd0-39f4-a8f4-1ef216a4b5f2","id":1,"lastScannedAt":"2018-07-14T08:13:28Z","lastUploadedAt":"2018-07-14T08:13:28Z","needKernelRestart":true,"osFamily":"centos","osVersion":"6","serverName":"server01","serverUuid":"abcdef12-ecd0-39f4-a8f4-1ef216a4b5f2","serverroleId":1,"serverroleName":"server_role01","tags":["Fugit repellendus illo.","Aperiam ipsa voluptate autem unde.","Fuga accusamus aut."],"updatedAt":"2018-07-14T08:13:28Z"},"tasks":[{"applyingPatchOn":"2018-07-13","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2017-6799","id":1,"ignore":true,"ignoreUntil":"vector","mainUserID":1,"mainUserName":"main-user-name","priority":"high","serverID":1,"status":"new","subUserID":1,"subUserName":"sub-user-name","updatedAt":"2018-07-14T08:13:28Z"},{"applyingPatchOn":"2018-07-13","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2017-6799","id":1,"ignore":true,"ignoreUntil":"vector","mainUserID":1,"mainUserName":"main-user-name","priority":"high","serverID":1,"status":"new","subUserID":1,"subUserName":"sub-user-name","updatedAt":"2018-07-14T08:13:28Z"},{"applyingPatchOn":"2018-07-13","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2017-6799","id":1,"ignore":true,"ignoreUntil":"vector","mainUserID":1,"mainUserName":"main-user-name","priority":"high","serverID":1,"status":"new","subUserID":1,"subUserName":"sub-user-name","updatedAt":"2018-07-14T08:13:28Z"},{"applyingPatchOn":"2018-07-13","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2017-6799","id":1,"ignore":true,"ignoreUntil":"vector","mainUserID":1,"mainUserName":"main-user-name","priority":"high","serverID":1,"status":"new","subUserID":1,"subUserName":"sub-user-name","updatedAt":"2018-07-14T08:13:28Z"}],"updatedAt":"2018-07-14T08:13:28Z","version":"1.0"}
//
// swagger:model PkgCpeGetPkgDetailResponseBody
type PkgCpeGetPkgDetailResponseBody struct {

	// AffectedProcess list of package
	// Example: [{"name":"apache","pid":"12"},{"name":"apache","pid":"12"},{"name":"apache","pid":"12"},{"name":"apache","pid":"12"}]
	AffectedProcs []*AffectedProcResponseBody `json:"affectedProcs"`

	// crated time of package or cpe
	// Example: 2018-07-14T08:13:28Z
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"createdAt"`

	// ID of package
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// Name of package or cpe
	// Example: package01
	// Required: true
	Name *string `json:"name"`

	// NeedRestartProcess list of package
	// Example: [{"initSystem":"initSystem","path":"path","pid":"12","serviceName":"serviceName"},{"initSystem":"initSystem","path":"path","pid":"12","serviceName":"serviceName"},{"initSystem":"initSystem","path":"path","pid":"12","serviceName":"serviceName"}]
	NeedRestartProcs []*NeedRestartProcResponseBody `json:"needRestartProcs"`

	// New Release of package
	// Example: new release
	NewRelease string `json:"newRelease,omitempty"`

	// New Version of package
	// Example: 2.0
	NewVersion string `json:"newVersion,omitempty"`

	// package status of package
	// Example: {"bash":"resolved"}
	PackageStatuses map[string]string `json:"packageStatuses,omitempty"`

	// Release of package
	// Example: release
	// Required: true
	Release *string `json:"release"`

	// Repository of package
	// Example: repository
	Repository string `json:"repository,omitempty"`

	// server
	// Required: true
	Server *ServerChildResponseBody `json:"server"`

	// updated time of server
	// Example: [{"applyingPatchOn":"2018-07-13","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2017-6799","id":1,"ignore":true,"ignoreUntil":"vector","mainUserID":1,"mainUserName":"main-user-name","priority":"high","serverID":1,"status":"new","subUserID":1,"subUserName":"sub-user-name","updatedAt":"2018-07-14T08:13:28Z"},{"applyingPatchOn":"2018-07-13","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2017-6799","id":1,"ignore":true,"ignoreUntil":"vector","mainUserID":1,"mainUserName":"main-user-name","priority":"high","serverID":1,"status":"new","subUserID":1,"subUserName":"sub-user-name","updatedAt":"2018-07-14T08:13:28Z"},{"applyingPatchOn":"2018-07-13","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2017-6799","id":1,"ignore":true,"ignoreUntil":"vector","mainUserID":1,"mainUserName":"main-user-name","priority":"high","serverID":1,"status":"new","subUserID":1,"subUserName":"sub-user-name","updatedAt":"2018-07-14T08:13:28Z"}]
	Tasks []*ChildTaskResponseBody `json:"tasks"`

	// updated time of package or cpe
	// Example: 2018-07-14T08:13:28Z
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updatedAt"`

	// Version of package or cpe
	// Example: 1.0
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this pkg cpe get pkg detail response body
func (m *PkgCpeGetPkgDetailResponseBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAffectedProcs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNeedRestartProcs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelease(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTasks(formats); err != nil {
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

func (m *PkgCpeGetPkgDetailResponseBody) validateAffectedProcs(formats strfmt.Registry) error {
	if swag.IsZero(m.AffectedProcs) { // not required
		return nil
	}

	for i := 0; i < len(m.AffectedProcs); i++ {
		if swag.IsZero(m.AffectedProcs[i]) { // not required
			continue
		}

		if m.AffectedProcs[i] != nil {
			if err := m.AffectedProcs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("affectedProcs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("affectedProcs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PkgCpeGetPkgDetailResponseBody) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PkgCpeGetPkgDetailResponseBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PkgCpeGetPkgDetailResponseBody) validateNeedRestartProcs(formats strfmt.Registry) error {
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

func (m *PkgCpeGetPkgDetailResponseBody) validateRelease(formats strfmt.Registry) error {

	if err := validate.Required("release", "body", m.Release); err != nil {
		return err
	}

	return nil
}

func (m *PkgCpeGetPkgDetailResponseBody) validateServer(formats strfmt.Registry) error {

	if err := validate.Required("server", "body", m.Server); err != nil {
		return err
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

func (m *PkgCpeGetPkgDetailResponseBody) validateTasks(formats strfmt.Registry) error {
	if swag.IsZero(m.Tasks) { // not required
		return nil
	}

	for i := 0; i < len(m.Tasks); i++ {
		if swag.IsZero(m.Tasks[i]) { // not required
			continue
		}

		if m.Tasks[i] != nil {
			if err := m.Tasks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tasks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tasks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PkgCpeGetPkgDetailResponseBody) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updatedAt", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PkgCpeGetPkgDetailResponseBody) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this pkg cpe get pkg detail response body based on the context it is used
func (m *PkgCpeGetPkgDetailResponseBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAffectedProcs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNeedRestartProcs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTasks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PkgCpeGetPkgDetailResponseBody) contextValidateAffectedProcs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AffectedProcs); i++ {

		if m.AffectedProcs[i] != nil {
			if err := m.AffectedProcs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("affectedProcs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("affectedProcs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PkgCpeGetPkgDetailResponseBody) contextValidateNeedRestartProcs(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PkgCpeGetPkgDetailResponseBody) contextValidateServer(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PkgCpeGetPkgDetailResponseBody) contextValidateTasks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tasks); i++ {

		if m.Tasks[i] != nil {
			if err := m.Tasks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tasks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tasks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PkgCpeGetPkgDetailResponseBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PkgCpeGetPkgDetailResponseBody) UnmarshalBinary(b []byte) error {
	var res PkgCpeGetPkgDetailResponseBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
