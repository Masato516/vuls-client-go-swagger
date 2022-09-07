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

// ServerGetServerDetailByUUIDResponseBody ServerGetServerDetailByUUIDResponseBody
// Example: {"createdAt":"2018-07-14T08:13:28Z","defaultUserId":1,"defaultUserName":"vuls user","hostUuid":"141df30a-ecd0-39f4-a8f4-1ef216a4b5f2","id":1,"lastScannedAt":"2018-07-14T08:13:28Z","lastUploadedAt":"2018-07-14T08:13:28Z","needKernelRestart":true,"osFamily":"centos","osVersion":"6","platformInstanceId":"i-xxxxxxx","platformName":"aws","serverIpv4":"192.168.0.2","serverName":"server01","serverUuid":"abcdef12-ecd0-39f4-a8f4-1ef216a4b5f2","serverroleId":1,"serverroleName":"server_role01","tags":[{"id":1,"name":"tag"},{"id":1,"name":"tag"},{"id":1,"name":"tag"},{"id":1,"name":"tag"}],"tasks":[{"applyingPatchOn":"2018-07-13","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2017-6799","id":1,"ignore":true,"ignoreUntil":"vector","mainUserID":1,"mainUserName":"main-user-name","priority":"high","serverID":1,"status":"new","subUserID":1,"subUserName":"sub-user-name","updatedAt":"2018-07-14T08:13:28Z"},{"applyingPatchOn":"2018-07-13","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2017-6799","id":1,"ignore":true,"ignoreUntil":"vector","mainUserID":1,"mainUserName":"main-user-name","priority":"high","serverID":1,"status":"new","subUserID":1,"subUserName":"sub-user-name","updatedAt":"2018-07-14T08:13:28Z"},{"applyingPatchOn":"2018-07-13","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2017-6799","id":1,"ignore":true,"ignoreUntil":"vector","mainUserID":1,"mainUserName":"main-user-name","priority":"high","serverID":1,"status":"new","subUserID":1,"subUserName":"sub-user-name","updatedAt":"2018-07-14T08:13:28Z"},{"applyingPatchOn":"2018-07-13","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2017-6799","id":1,"ignore":true,"ignoreUntil":"vector","mainUserID":1,"mainUserName":"main-user-name","priority":"high","serverID":1,"status":"new","subUserID":1,"subUserName":"sub-user-name","updatedAt":"2018-07-14T08:13:28Z"}],"updatedAt":"2018-07-14T08:13:28Z"}
//
// swagger:model ServerGetServerDetailByUUIDResponseBody
type ServerGetServerDetailByUUIDResponseBody struct {

	// crated time of server
	// Example: 2018-07-14T08:13:28Z
	// Required: true
	CreatedAt *string `json:"createdAt"`

	// default user ID of server
	// Example: 1
	DefaultUserID int64 `json:"defaultUserId,omitempty"`

	// default user name of server
	// Example: vuls user
	DefaultUserName string `json:"defaultUserName,omitempty"`

	// UUID of server
	// Example: 141df30a-ecd0-39f4-a8f4-1ef216a4b5f2
	// Required: true
	// Format: uuid
	HostUUID *strfmt.UUID `json:"hostUuid"`

	// ID of server
	// Example: 1
	// Required: true
	ID *int64 `json:"id"`

	// last scanned time of server
	// Example: 2018-07-14T08:13:28Z
	LastScannedAt string `json:"lastScannedAt,omitempty"`

	// last uploaded time of server
	// Example: 2018-07-14T08:13:28Z
	LastUploadedAt string `json:"lastUploadedAt,omitempty"`

	// Whether server needs kernel restart
	// Example: true
	// Required: true
	NeedKernelRestart *bool `json:"needKernelRestart"`

	// OS Name of server
	// Example: centos
	// Required: true
	OsFamily *string `json:"osFamily"`

	// OS Version of server
	// Example: 6
	// Required: true
	OsVersion *string `json:"osVersion"`

	// platformInstanceId of server
	// Example: i-xxxxxxx
	// Required: true
	PlatformInstanceID *string `json:"platformInstanceId"`

	// platformName of server
	// Example: aws
	// Required: true
	PlatformName *string `json:"platformName"`

	// IPv4 of server
	// Example: 192.168.0.2
	// Required: true
	ServerIPV4 *string `json:"serverIpv4"`

	// Name of server
	// Example: server01
	// Required: true
	ServerName *string `json:"serverName"`

	// UUID of server
	// Example: abcdef12-ecd0-39f4-a8f4-1ef216a4b5f2
	// Required: true
	// Format: uuid
	ServerUUID *strfmt.UUID `json:"serverUuid"`

	// ID of server role
	// Example: 1
	// Required: true
	ServerroleID *int64 `json:"serverroleId"`

	// Name of server role
	// Example: server_role01
	// Required: true
	ServerroleName *string `json:"serverroleName"`

	// tags is list of server tag
	// Example: [{"id":1,"name":"tag"},{"id":1,"name":"tag"}]
	Tags []*ServerTagResponseBody `json:"tags"`

	// tasks of server
	// Example: [{"applyingPatchOn":"2018-07-13","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2017-6799","id":1,"ignore":true,"ignoreUntil":"vector","mainUserID":1,"mainUserName":"main-user-name","priority":"high","serverID":1,"status":"new","subUserID":1,"subUserName":"sub-user-name","updatedAt":"2018-07-14T08:13:28Z"},{"applyingPatchOn":"2018-07-13","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2017-6799","id":1,"ignore":true,"ignoreUntil":"vector","mainUserID":1,"mainUserName":"main-user-name","priority":"high","serverID":1,"status":"new","subUserID":1,"subUserName":"sub-user-name","updatedAt":"2018-07-14T08:13:28Z"}]
	Tasks []*ChildTaskResponseBody `json:"tasks"`

	// updated time of server
	// Example: 2018-07-14T08:13:28Z
	// Required: true
	UpdatedAt *string `json:"updatedAt"`
}

// Validate validates this server get server detail by UUID response body
func (m *ServerGetServerDetailByUUIDResponseBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNeedKernelRestart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOsFamily(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOsVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatformInstanceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatformName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerIPV4(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerroleID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerroleName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTasks(formats); err != nil {
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

func (m *ServerGetServerDetailByUUIDResponseBody) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *ServerGetServerDetailByUUIDResponseBody) validateHostUUID(formats strfmt.Registry) error {

	if err := validate.Required("hostUuid", "body", m.HostUUID); err != nil {
		return err
	}

	if err := validate.FormatOf("hostUuid", "body", "uuid", m.HostUUID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ServerGetServerDetailByUUIDResponseBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ServerGetServerDetailByUUIDResponseBody) validateNeedKernelRestart(formats strfmt.Registry) error {

	if err := validate.Required("needKernelRestart", "body", m.NeedKernelRestart); err != nil {
		return err
	}

	return nil
}

func (m *ServerGetServerDetailByUUIDResponseBody) validateOsFamily(formats strfmt.Registry) error {

	if err := validate.Required("osFamily", "body", m.OsFamily); err != nil {
		return err
	}

	return nil
}

func (m *ServerGetServerDetailByUUIDResponseBody) validateOsVersion(formats strfmt.Registry) error {

	if err := validate.Required("osVersion", "body", m.OsVersion); err != nil {
		return err
	}

	return nil
}

func (m *ServerGetServerDetailByUUIDResponseBody) validatePlatformInstanceID(formats strfmt.Registry) error {

	if err := validate.Required("platformInstanceId", "body", m.PlatformInstanceID); err != nil {
		return err
	}

	return nil
}

func (m *ServerGetServerDetailByUUIDResponseBody) validatePlatformName(formats strfmt.Registry) error {

	if err := validate.Required("platformName", "body", m.PlatformName); err != nil {
		return err
	}

	return nil
}

func (m *ServerGetServerDetailByUUIDResponseBody) validateServerIPV4(formats strfmt.Registry) error {

	if err := validate.Required("serverIpv4", "body", m.ServerIPV4); err != nil {
		return err
	}

	return nil
}

func (m *ServerGetServerDetailByUUIDResponseBody) validateServerName(formats strfmt.Registry) error {

	if err := validate.Required("serverName", "body", m.ServerName); err != nil {
		return err
	}

	return nil
}

func (m *ServerGetServerDetailByUUIDResponseBody) validateServerUUID(formats strfmt.Registry) error {

	if err := validate.Required("serverUuid", "body", m.ServerUUID); err != nil {
		return err
	}

	if err := validate.FormatOf("serverUuid", "body", "uuid", m.ServerUUID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ServerGetServerDetailByUUIDResponseBody) validateServerroleID(formats strfmt.Registry) error {

	if err := validate.Required("serverroleId", "body", m.ServerroleID); err != nil {
		return err
	}

	return nil
}

func (m *ServerGetServerDetailByUUIDResponseBody) validateServerroleName(formats strfmt.Registry) error {

	if err := validate.Required("serverroleName", "body", m.ServerroleName); err != nil {
		return err
	}

	return nil
}

func (m *ServerGetServerDetailByUUIDResponseBody) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServerGetServerDetailByUUIDResponseBody) validateTasks(formats strfmt.Registry) error {
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

func (m *ServerGetServerDetailByUUIDResponseBody) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updatedAt", "body", m.UpdatedAt); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this server get server detail by UUID response body based on the context it is used
func (m *ServerGetServerDetailByUUIDResponseBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTags(ctx, formats); err != nil {
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

func (m *ServerGetServerDetailByUUIDResponseBody) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServerGetServerDetailByUUIDResponseBody) contextValidateTasks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ServerGetServerDetailByUUIDResponseBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerGetServerDetailByUUIDResponseBody) UnmarshalBinary(b []byte) error {
	var res ServerGetServerDetailByUUIDResponseBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
