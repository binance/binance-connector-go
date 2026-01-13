/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewOrderWorkingTypeParameter the model 'NewOrderWorkingTypeParameter'
type NewOrderWorkingTypeParameter string

// List of newOrder_workingType_parameter
const (
	NewOrderWorkingTypeParameterMarkPrice     NewOrderWorkingTypeParameter = "MARK_PRICE"
	NewOrderWorkingTypeParameterContractPrice NewOrderWorkingTypeParameter = "CONTRACT_PRICE"
)

// All allowed values of NewOrderWorkingTypeParameter enum
var AllowedNewOrderWorkingTypeParameterEnumValues = []NewOrderWorkingTypeParameter{
	"MARK_PRICE",
	"CONTRACT_PRICE",
}

func (v *NewOrderWorkingTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewOrderWorkingTypeParameter(value)
	for _, existing := range AllowedNewOrderWorkingTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewOrderWorkingTypeParameter", value)
}

// NewNewOrderWorkingTypeParameterFromValue returns a pointer to a valid NewOrderWorkingTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewOrderWorkingTypeParameterFromValue(v string) (*NewOrderWorkingTypeParameter, error) {
	ev := NewOrderWorkingTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewOrderWorkingTypeParameter: valid values are %v", v, AllowedNewOrderWorkingTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewOrderWorkingTypeParameter) IsValid() bool {
	for _, existing := range AllowedNewOrderWorkingTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newOrder_workingType_parameter value
func (v NewOrderWorkingTypeParameter) Ptr() *NewOrderWorkingTypeParameter {
	return &v
}

type NullableNewOrderWorkingTypeParameter struct {
	value *NewOrderWorkingTypeParameter
	isSet bool
}

func (v NullableNewOrderWorkingTypeParameter) Get() *NewOrderWorkingTypeParameter {
	return v.value
}

func (v *NullableNewOrderWorkingTypeParameter) Set(val *NewOrderWorkingTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewOrderWorkingTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewOrderWorkingTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewOrderWorkingTypeParameter(val *NewOrderWorkingTypeParameter) *NullableNewOrderWorkingTypeParameter {
	return &NullableNewOrderWorkingTypeParameter{value: val, isSet: true}
}

func (v NullableNewOrderWorkingTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewOrderWorkingTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
