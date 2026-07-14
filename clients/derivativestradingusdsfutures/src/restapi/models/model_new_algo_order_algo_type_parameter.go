/*
Futures (USDⓈ-M) REST API

Access market data, manage accounts, and trade USDⓈ-M perpetual futures.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewAlgoOrderAlgoTypeParameter the model 'NewAlgoOrderAlgoTypeParameter'
type NewAlgoOrderAlgoTypeParameter string

// List of newAlgoOrder_algoType_parameter
const (
	NewAlgoOrderAlgoTypeParameterConditional NewAlgoOrderAlgoTypeParameter = "CONDITIONAL"
)

// All allowed values of NewAlgoOrderAlgoTypeParameter enum
var AllowedNewAlgoOrderAlgoTypeParameterEnumValues = []NewAlgoOrderAlgoTypeParameter{
	"CONDITIONAL",
}

func (v *NewAlgoOrderAlgoTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewAlgoOrderAlgoTypeParameter(value)
	for _, existing := range AllowedNewAlgoOrderAlgoTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewAlgoOrderAlgoTypeParameter", value)
}

// NewNewAlgoOrderAlgoTypeParameterFromValue returns a pointer to a valid NewAlgoOrderAlgoTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewAlgoOrderAlgoTypeParameterFromValue(v string) (*NewAlgoOrderAlgoTypeParameter, error) {
	ev := NewAlgoOrderAlgoTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewAlgoOrderAlgoTypeParameter: valid values are %v", v, AllowedNewAlgoOrderAlgoTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewAlgoOrderAlgoTypeParameter) IsValid() bool {
	for _, existing := range AllowedNewAlgoOrderAlgoTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newAlgoOrder_algoType_parameter value
func (v NewAlgoOrderAlgoTypeParameter) Ptr() *NewAlgoOrderAlgoTypeParameter {
	return &v
}

type NullableNewAlgoOrderAlgoTypeParameter struct {
	value *NewAlgoOrderAlgoTypeParameter
	isSet bool
}

func (v NullableNewAlgoOrderAlgoTypeParameter) Get() *NewAlgoOrderAlgoTypeParameter {
	return v.value
}

func (v *NullableNewAlgoOrderAlgoTypeParameter) Set(val *NewAlgoOrderAlgoTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewAlgoOrderAlgoTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewAlgoOrderAlgoTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewAlgoOrderAlgoTypeParameter(val *NewAlgoOrderAlgoTypeParameter) *NullableNewAlgoOrderAlgoTypeParameter {
	return &NullableNewAlgoOrderAlgoTypeParameter{value: val, isSet: true}
}

func (v NullableNewAlgoOrderAlgoTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewAlgoOrderAlgoTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
