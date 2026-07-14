/*
Margin REST API

Access account information, borrow and repay assets, and trade with Binance Margin.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// MarginAccountNewOtoWorkingTypeParameter the model 'MarginAccountNewOtoWorkingTypeParameter'
type MarginAccountNewOtoWorkingTypeParameter string

// List of marginAccountNewOto_workingType_parameter
const (
	MarginAccountNewOtoWorkingTypeParameterLimit      MarginAccountNewOtoWorkingTypeParameter = "LIMIT"
	MarginAccountNewOtoWorkingTypeParameterLimitMaker MarginAccountNewOtoWorkingTypeParameter = "LIMIT_MAKER"
)

// All allowed values of MarginAccountNewOtoWorkingTypeParameter enum
var AllowedMarginAccountNewOtoWorkingTypeParameterEnumValues = []MarginAccountNewOtoWorkingTypeParameter{
	"LIMIT",
	"LIMIT_MAKER",
}

func (v *MarginAccountNewOtoWorkingTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MarginAccountNewOtoWorkingTypeParameter(value)
	for _, existing := range AllowedMarginAccountNewOtoWorkingTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MarginAccountNewOtoWorkingTypeParameter", value)
}

// NewMarginAccountNewOtoWorkingTypeParameterFromValue returns a pointer to a valid MarginAccountNewOtoWorkingTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMarginAccountNewOtoWorkingTypeParameterFromValue(v string) (*MarginAccountNewOtoWorkingTypeParameter, error) {
	ev := MarginAccountNewOtoWorkingTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MarginAccountNewOtoWorkingTypeParameter: valid values are %v", v, AllowedMarginAccountNewOtoWorkingTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MarginAccountNewOtoWorkingTypeParameter) IsValid() bool {
	for _, existing := range AllowedMarginAccountNewOtoWorkingTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to marginAccountNewOto_workingType_parameter value
func (v MarginAccountNewOtoWorkingTypeParameter) Ptr() *MarginAccountNewOtoWorkingTypeParameter {
	return &v
}

type NullableMarginAccountNewOtoWorkingTypeParameter struct {
	value *MarginAccountNewOtoWorkingTypeParameter
	isSet bool
}

func (v NullableMarginAccountNewOtoWorkingTypeParameter) Get() *MarginAccountNewOtoWorkingTypeParameter {
	return v.value
}

func (v *NullableMarginAccountNewOtoWorkingTypeParameter) Set(val *MarginAccountNewOtoWorkingTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginAccountNewOtoWorkingTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginAccountNewOtoWorkingTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginAccountNewOtoWorkingTypeParameter(val *MarginAccountNewOtoWorkingTypeParameter) *NullableMarginAccountNewOtoWorkingTypeParameter {
	return &NullableMarginAccountNewOtoWorkingTypeParameter{value: val, isSet: true}
}

func (v NullableMarginAccountNewOtoWorkingTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginAccountNewOtoWorkingTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
