/*
Margin REST API

Access account information, borrow and repay assets, and trade with Binance Margin.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// CreateSpecialKeyPermissionModeParameter the model 'CreateSpecialKeyPermissionModeParameter'
type CreateSpecialKeyPermissionModeParameter string

// List of createSpecialKey_permissionMode_parameter
const (
	CreateSpecialKeyPermissionModeParameterTrade CreateSpecialKeyPermissionModeParameter = "TRADE"
	CreateSpecialKeyPermissionModeParameterRead  CreateSpecialKeyPermissionModeParameter = "READ"
)

// All allowed values of CreateSpecialKeyPermissionModeParameter enum
var AllowedCreateSpecialKeyPermissionModeParameterEnumValues = []CreateSpecialKeyPermissionModeParameter{
	"TRADE",
	"READ",
}

func (v *CreateSpecialKeyPermissionModeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CreateSpecialKeyPermissionModeParameter(value)
	for _, existing := range AllowedCreateSpecialKeyPermissionModeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CreateSpecialKeyPermissionModeParameter", value)
}

// NewCreateSpecialKeyPermissionModeParameterFromValue returns a pointer to a valid CreateSpecialKeyPermissionModeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreateSpecialKeyPermissionModeParameterFromValue(v string) (*CreateSpecialKeyPermissionModeParameter, error) {
	ev := CreateSpecialKeyPermissionModeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreateSpecialKeyPermissionModeParameter: valid values are %v", v, AllowedCreateSpecialKeyPermissionModeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreateSpecialKeyPermissionModeParameter) IsValid() bool {
	for _, existing := range AllowedCreateSpecialKeyPermissionModeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to createSpecialKey_permissionMode_parameter value
func (v CreateSpecialKeyPermissionModeParameter) Ptr() *CreateSpecialKeyPermissionModeParameter {
	return &v
}

type NullableCreateSpecialKeyPermissionModeParameter struct {
	value *CreateSpecialKeyPermissionModeParameter
	isSet bool
}

func (v NullableCreateSpecialKeyPermissionModeParameter) Get() *CreateSpecialKeyPermissionModeParameter {
	return v.value
}

func (v *NullableCreateSpecialKeyPermissionModeParameter) Set(val *CreateSpecialKeyPermissionModeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSpecialKeyPermissionModeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSpecialKeyPermissionModeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSpecialKeyPermissionModeParameter(val *CreateSpecialKeyPermissionModeParameter) *NullableCreateSpecialKeyPermissionModeParameter {
	return &NullableCreateSpecialKeyPermissionModeParameter{value: val, isSet: true}
}

func (v NullableCreateSpecialKeyPermissionModeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSpecialKeyPermissionModeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
