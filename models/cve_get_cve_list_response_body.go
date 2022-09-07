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
)

// CveGetCveListResponseBody CveGetCveListResponseBody
// Example: {"cves":[{"advisoryIDs":["advisoryID"],"allTaskCount":100,"createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","cwes":[{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"}],"hasExploit":true,"hasMitigation":true,"hasWorkaround":true,"isNotActive":true,"isOwaspTopTen2017":true,"maxV2":9,"maxV3":9,"newTaskCount":10,"scoreV2s":{"nvd":9,"redhat":7},"scoreV3s":{"nvd":8,"redhat":9},"title":"title","topicCount":10,"topicLastUpdatedAt":"2018-07-14T08:13:28Z","updatedAt":"2018-07-14T08:13:28Z","vectorV2s":{"jvn":"AV:L/AC:M/Au:N/C:C/I:N/A:N","nvd":"AV:L/AC:M/Au:N/C:C/I:N/A:N"},"vectorV3s":{"jvn":"AV:L/AC:H/PR:L/UI:N/S:C/C:H/I:N/A:N","nvd":"AV:L/AC:H/PR:L/UI:N/S:C/C:H/I:N/A:N"}},{"advisoryIDs":["advisoryID"],"allTaskCount":100,"createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","cwes":[{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"}],"hasExploit":true,"hasMitigation":true,"hasWorkaround":true,"isNotActive":true,"isOwaspTopTen2017":true,"maxV2":9,"maxV3":9,"newTaskCount":10,"scoreV2s":{"nvd":9,"redhat":7},"scoreV3s":{"nvd":8,"redhat":9},"title":"title","topicCount":10,"topicLastUpdatedAt":"2018-07-14T08:13:28Z","updatedAt":"2018-07-14T08:13:28Z","vectorV2s":{"jvn":"AV:L/AC:M/Au:N/C:C/I:N/A:N","nvd":"AV:L/AC:M/Au:N/C:C/I:N/A:N"},"vectorV3s":{"jvn":"AV:L/AC:H/PR:L/UI:N/S:C/C:H/I:N/A:N","nvd":"AV:L/AC:H/PR:L/UI:N/S:C/C:H/I:N/A:N"}}],"paging":{"limit":20,"offset":10,"page":1,"totalCount":200,"totalPage":10}}
//
// swagger:model CveGetCveListResponseBody
type CveGetCveListResponseBody struct {

	// Cves list
	// Example: [{"advisoryIDs":["advisoryID"],"allTaskCount":100,"createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","cwes":[{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"}],"hasExploit":true,"hasMitigation":true,"hasWorkaround":true,"isNotActive":true,"isOwaspTopTen2017":true,"maxV2":9,"maxV3":9,"newTaskCount":10,"scoreV2s":{"nvd":9,"redhat":7},"scoreV3s":{"nvd":8,"redhat":9},"title":"title","topicCount":10,"topicLastUpdatedAt":"2018-07-14T08:13:28Z","updatedAt":"2018-07-14T08:13:28Z","vectorV2s":{"jvn":"AV:L/AC:M/Au:N/C:C/I:N/A:N","nvd":"AV:L/AC:M/Au:N/C:C/I:N/A:N"},"vectorV3s":{"jvn":"AV:L/AC:H/PR:L/UI:N/S:C/C:H/I:N/A:N","nvd":"AV:L/AC:H/PR:L/UI:N/S:C/C:H/I:N/A:N"}},{"advisoryIDs":["advisoryID"],"allTaskCount":100,"createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","cwes":[{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"}],"hasExploit":true,"hasMitigation":true,"hasWorkaround":true,"isNotActive":true,"isOwaspTopTen2017":true,"maxV2":9,"maxV3":9,"newTaskCount":10,"scoreV2s":{"nvd":9,"redhat":7},"scoreV3s":{"nvd":8,"redhat":9},"title":"title","topicCount":10,"topicLastUpdatedAt":"2018-07-14T08:13:28Z","updatedAt":"2018-07-14T08:13:28Z","vectorV2s":{"jvn":"AV:L/AC:M/Au:N/C:C/I:N/A:N","nvd":"AV:L/AC:M/Au:N/C:C/I:N/A:N"},"vectorV3s":{"jvn":"AV:L/AC:H/PR:L/UI:N/S:C/C:H/I:N/A:N","nvd":"AV:L/AC:H/PR:L/UI:N/S:C/C:H/I:N/A:N"}},{"advisoryIDs":["advisoryID"],"allTaskCount":100,"createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","cwes":[{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"}],"hasExploit":true,"hasMitigation":true,"hasWorkaround":true,"isNotActive":true,"isOwaspTopTen2017":true,"maxV2":9,"maxV3":9,"newTaskCount":10,"scoreV2s":{"nvd":9,"redhat":7},"scoreV3s":{"nvd":8,"redhat":9},"title":"title","topicCount":10,"topicLastUpdatedAt":"2018-07-14T08:13:28Z","updatedAt":"2018-07-14T08:13:28Z","vectorV2s":{"jvn":"AV:L/AC:M/Au:N/C:C/I:N/A:N","nvd":"AV:L/AC:M/Au:N/C:C/I:N/A:N"},"vectorV3s":{"jvn":"AV:L/AC:H/PR:L/UI:N/S:C/C:H/I:N/A:N","nvd":"AV:L/AC:H/PR:L/UI:N/S:C/C:H/I:N/A:N"}},{"advisoryIDs":["advisoryID"],"allTaskCount":100,"createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","cwes":[{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"}],"hasExploit":true,"hasMitigation":true,"hasWorkaround":true,"isNotActive":true,"isOwaspTopTen2017":true,"maxV2":9,"maxV3":9,"newTaskCount":10,"scoreV2s":{"nvd":9,"redhat":7},"scoreV3s":{"nvd":8,"redhat":9},"title":"title","topicCount":10,"topicLastUpdatedAt":"2018-07-14T08:13:28Z","updatedAt":"2018-07-14T08:13:28Z","vectorV2s":{"jvn":"AV:L/AC:M/Au:N/C:C/I:N/A:N","nvd":"AV:L/AC:M/Au:N/C:C/I:N/A:N"},"vectorV3s":{"jvn":"AV:L/AC:H/PR:L/UI:N/S:C/C:H/I:N/A:N","nvd":"AV:L/AC:H/PR:L/UI:N/S:C/C:H/I:N/A:N"}}]
	Cves []*CveListResponseBody `json:"cves"`

	// paging
	Paging *PagingResponseBody `json:"paging,omitempty"`
}

// Validate validates this cve get cve list response body
func (m *CveGetCveListResponseBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCves(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaging(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CveGetCveListResponseBody) validateCves(formats strfmt.Registry) error {
	if swag.IsZero(m.Cves) { // not required
		return nil
	}

	for i := 0; i < len(m.Cves); i++ {
		if swag.IsZero(m.Cves[i]) { // not required
			continue
		}

		if m.Cves[i] != nil {
			if err := m.Cves[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cves" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cves" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CveGetCveListResponseBody) validatePaging(formats strfmt.Registry) error {
	if swag.IsZero(m.Paging) { // not required
		return nil
	}

	if m.Paging != nil {
		if err := m.Paging.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("paging")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("paging")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cve get cve list response body based on the context it is used
func (m *CveGetCveListResponseBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCves(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePaging(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CveGetCveListResponseBody) contextValidateCves(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Cves); i++ {

		if m.Cves[i] != nil {
			if err := m.Cves[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cves" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cves" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CveGetCveListResponseBody) contextValidatePaging(ctx context.Context, formats strfmt.Registry) error {

	if m.Paging != nil {
		if err := m.Paging.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("paging")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("paging")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CveGetCveListResponseBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CveGetCveListResponseBody) UnmarshalBinary(b []byte) error {
	var res CveGetCveListResponseBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
