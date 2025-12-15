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

// NewAlgoOrderWorkingTypeParameter the model 'NewAlgoOrderWorkingTypeParameter'
type NewAlgoOrderWorkingTypeParameter string

// List of newAlgoOrder_workingType_parameter
const (
	NewAlgoOrderWorkingTypeParameterMarkPrice     NewAlgoOrderWorkingTypeParameter = "MARK_PRICE"
	NewAlgoOrderWorkingTypeParameterContractPrice NewAlgoOrderWorkingTypeParameter = "CONTRACT_PRICE"
)

// All allowed values of NewAlgoOrderWorkingTypeParameter enum
var AllowedNewAlgoOrderWorkingTypeParameterEnumValues = []NewAlgoOrderWorkingTypeParameter{
	"MARK_PRICE",
	"CONTRACT_PRICE",
}

func (v *NewAlgoOrderWorkingTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewAlgoOrderWorkingTypeParameter(value)
	for _, existing := range AllowedNewAlgoOrderWorkingTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewAlgoOrderWorkingTypeParameter", value)
}

// NewNewAlgoOrderWorkingTypeParameterFromValue returns a pointer to a valid NewAlgoOrderWorkingTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewAlgoOrderWorkingTypeParameterFromValue(v string) (*NewAlgoOrderWorkingTypeParameter, error) {
	ev := NewAlgoOrderWorkingTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewAlgoOrderWorkingTypeParameter: valid values are %v", v, AllowedNewAlgoOrderWorkingTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewAlgoOrderWorkingTypeParameter) IsValid() bool {
	for _, existing := range AllowedNewAlgoOrderWorkingTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newAlgoOrder_workingType_parameter value
func (v NewAlgoOrderWorkingTypeParameter) Ptr() *NewAlgoOrderWorkingTypeParameter {
	return &v
}

type NullableNewAlgoOrderWorkingTypeParameter struct {
	value *NewAlgoOrderWorkingTypeParameter
	isSet bool
}

func (v NullableNewAlgoOrderWorkingTypeParameter) Get() *NewAlgoOrderWorkingTypeParameter {
	return v.value
}

func (v *NullableNewAlgoOrderWorkingTypeParameter) Set(val *NewAlgoOrderWorkingTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewAlgoOrderWorkingTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewAlgoOrderWorkingTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewAlgoOrderWorkingTypeParameter(val *NewAlgoOrderWorkingTypeParameter) *NullableNewAlgoOrderWorkingTypeParameter {
	return &NullableNewAlgoOrderWorkingTypeParameter{value: val, isSet: true}
}

func (v NullableNewAlgoOrderWorkingTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewAlgoOrderWorkingTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
