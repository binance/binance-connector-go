/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"
	"fmt"
)

// UsersForceOrdersAutoCloseTypeParameter the model 'UsersForceOrdersAutoCloseTypeParameter'
type UsersForceOrdersAutoCloseTypeParameter string

// List of usersForceOrders_autoCloseType_parameter
const (
	UsersForceOrdersAutoCloseTypeParameterLiquidation UsersForceOrdersAutoCloseTypeParameter = "LIQUIDATION"
	UsersForceOrdersAutoCloseTypeParameterAdl         UsersForceOrdersAutoCloseTypeParameter = "ADL"
)

// All allowed values of UsersForceOrdersAutoCloseTypeParameter enum
var AllowedUsersForceOrdersAutoCloseTypeParameterEnumValues = []UsersForceOrdersAutoCloseTypeParameter{
	"LIQUIDATION",
	"ADL",
}

func (v *UsersForceOrdersAutoCloseTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UsersForceOrdersAutoCloseTypeParameter(value)
	for _, existing := range AllowedUsersForceOrdersAutoCloseTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UsersForceOrdersAutoCloseTypeParameter", value)
}

// NewUsersForceOrdersAutoCloseTypeParameterFromValue returns a pointer to a valid UsersForceOrdersAutoCloseTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUsersForceOrdersAutoCloseTypeParameterFromValue(v string) (*UsersForceOrdersAutoCloseTypeParameter, error) {
	ev := UsersForceOrdersAutoCloseTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UsersForceOrdersAutoCloseTypeParameter: valid values are %v", v, AllowedUsersForceOrdersAutoCloseTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UsersForceOrdersAutoCloseTypeParameter) IsValid() bool {
	for _, existing := range AllowedUsersForceOrdersAutoCloseTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to usersForceOrders_autoCloseType_parameter value
func (v UsersForceOrdersAutoCloseTypeParameter) Ptr() *UsersForceOrdersAutoCloseTypeParameter {
	return &v
}

type NullableUsersForceOrdersAutoCloseTypeParameter struct {
	value *UsersForceOrdersAutoCloseTypeParameter
	isSet bool
}

func (v NullableUsersForceOrdersAutoCloseTypeParameter) Get() *UsersForceOrdersAutoCloseTypeParameter {
	return v.value
}

func (v *NullableUsersForceOrdersAutoCloseTypeParameter) Set(val *UsersForceOrdersAutoCloseTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersForceOrdersAutoCloseTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersForceOrdersAutoCloseTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersForceOrdersAutoCloseTypeParameter(val *UsersForceOrdersAutoCloseTypeParameter) *NullableUsersForceOrdersAutoCloseTypeParameter {
	return &NullableUsersForceOrdersAutoCloseTypeParameter{value: val, isSet: true}
}

func (v NullableUsersForceOrdersAutoCloseTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersForceOrdersAutoCloseTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
