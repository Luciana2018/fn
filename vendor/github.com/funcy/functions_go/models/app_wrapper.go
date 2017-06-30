package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AppWrapper app wrapper
// swagger:model AppWrapper
type AppWrapper struct {

	// app
	// Required: true
	App *App `json:"app"`

	// error
	Error *ErrorBody `json:"error,omitempty"`
}

// Validate validates this app wrapper
func (m *AppWrapper) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApp(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateError(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppWrapper) validateApp(formats strfmt.Registry) error {

	if err := validate.Required("app", "body", m.App); err != nil {
		return err
	}

	if m.App != nil {

		if err := m.App.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("app")
			}
			return err
		}
	}

	return nil
}

func (m *AppWrapper) validateError(formats strfmt.Registry) error {

	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {

		if err := m.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppWrapper) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppWrapper) UnmarshalBinary(b []byte) error {
	var res AppWrapper
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
