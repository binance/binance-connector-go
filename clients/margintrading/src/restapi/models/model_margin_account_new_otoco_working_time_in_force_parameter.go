/*
Margin REST API

Access account information, borrow and repay assets, and trade with Binance Margin.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// MarginAccountNewOtocoWorkingTimeInForceParameter the model 'MarginAccountNewOtocoWorkingTimeInForceParameter'
type MarginAccountNewOtocoWorkingTimeInForceParameter string

// List of marginAccountNewOtoco_workingTimeInForce_parameter
const (
	MarginAccountNewOtocoWorkingTimeInForceParameterGtc MarginAccountNewOtocoWorkingTimeInForceParameter = "GTC"
	MarginAccountNewOtocoWorkingTimeInForceParameterIoc MarginAccountNewOtocoWorkingTimeInForceParameter = "IOC"
	MarginAccountNewOtocoWorkingTimeInForceParameterFok MarginAccountNewOtocoWorkingTimeInForceParameter = "FOK"
)

// All allowed values of MarginAccountNewOtocoWorkingTimeInForceParameter enum
var AllowedMarginAccountNewOtocoWorkingTimeInForceParameterEnumValues = []MarginAccountNewOtocoWorkingTimeInForceParameter{
	"GTC",
	"IOC",
	"FOK",
}

func (v *MarginAccountNewOtocoWorkingTimeInForceParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MarginAccountNewOtocoWorkingTimeInForceParameter(value)
	for _, existing := range AllowedMarginAccountNewOtocoWorkingTimeInForceParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MarginAccountNewOtocoWorkingTimeInForceParameter", value)
}

// NewMarginAccountNewOtocoWorkingTimeInForceParameterFromValue returns a pointer to a valid MarginAccountNewOtocoWorkingTimeInForceParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMarginAccountNewOtocoWorkingTimeInForceParameterFromValue(v string) (*MarginAccountNewOtocoWorkingTimeInForceParameter, error) {
	ev := MarginAccountNewOtocoWorkingTimeInForceParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MarginAccountNewOtocoWorkingTimeInForceParameter: valid values are %v", v, AllowedMarginAccountNewOtocoWorkingTimeInForceParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MarginAccountNewOtocoWorkingTimeInForceParameter) IsValid() bool {
	for _, existing := range AllowedMarginAccountNewOtocoWorkingTimeInForceParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to marginAccountNewOtoco_workingTimeInForce_parameter value
func (v MarginAccountNewOtocoWorkingTimeInForceParameter) Ptr() *MarginAccountNewOtocoWorkingTimeInForceParameter {
	return &v
}

type NullableMarginAccountNewOtocoWorkingTimeInForceParameter struct {
	value *MarginAccountNewOtocoWorkingTimeInForceParameter
	isSet bool
}

func (v NullableMarginAccountNewOtocoWorkingTimeInForceParameter) Get() *MarginAccountNewOtocoWorkingTimeInForceParameter {
	return v.value
}

func (v *NullableMarginAccountNewOtocoWorkingTimeInForceParameter) Set(val *MarginAccountNewOtocoWorkingTimeInForceParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginAccountNewOtocoWorkingTimeInForceParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginAccountNewOtocoWorkingTimeInForceParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginAccountNewOtocoWorkingTimeInForceParameter(val *MarginAccountNewOtocoWorkingTimeInForceParameter) *NullableMarginAccountNewOtocoWorkingTimeInForceParameter {
	return &NullableMarginAccountNewOtocoWorkingTimeInForceParameter{value: val, isSet: true}
}

func (v NullableMarginAccountNewOtocoWorkingTimeInForceParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginAccountNewOtocoWorkingTimeInForceParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
