/*
Portfolio Margin REST API

Access account information, manage margin positions, and trade with Binance Portfolio Margin.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewUmAlgoOrderAlgoTypeParameter the model 'NewUmAlgoOrderAlgoTypeParameter'
type NewUmAlgoOrderAlgoTypeParameter string

// List of newUmAlgoOrder_algoType_parameter
const (
	NewUmAlgoOrderAlgoTypeParameterConditional NewUmAlgoOrderAlgoTypeParameter = "CONDITIONAL"
)

// All allowed values of NewUmAlgoOrderAlgoTypeParameter enum
var AllowedNewUmAlgoOrderAlgoTypeParameterEnumValues = []NewUmAlgoOrderAlgoTypeParameter{
	"CONDITIONAL",
}

func (v *NewUmAlgoOrderAlgoTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewUmAlgoOrderAlgoTypeParameter(value)
	for _, existing := range AllowedNewUmAlgoOrderAlgoTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewUmAlgoOrderAlgoTypeParameter", value)
}

// NewNewUmAlgoOrderAlgoTypeParameterFromValue returns a pointer to a valid NewUmAlgoOrderAlgoTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewUmAlgoOrderAlgoTypeParameterFromValue(v string) (*NewUmAlgoOrderAlgoTypeParameter, error) {
	ev := NewUmAlgoOrderAlgoTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewUmAlgoOrderAlgoTypeParameter: valid values are %v", v, AllowedNewUmAlgoOrderAlgoTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewUmAlgoOrderAlgoTypeParameter) IsValid() bool {
	for _, existing := range AllowedNewUmAlgoOrderAlgoTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newUmAlgoOrder_algoType_parameter value
func (v NewUmAlgoOrderAlgoTypeParameter) Ptr() *NewUmAlgoOrderAlgoTypeParameter {
	return &v
}

type NullableNewUmAlgoOrderAlgoTypeParameter struct {
	value *NewUmAlgoOrderAlgoTypeParameter
	isSet bool
}

func (v NullableNewUmAlgoOrderAlgoTypeParameter) Get() *NewUmAlgoOrderAlgoTypeParameter {
	return v.value
}

func (v *NullableNewUmAlgoOrderAlgoTypeParameter) Set(val *NewUmAlgoOrderAlgoTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewUmAlgoOrderAlgoTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewUmAlgoOrderAlgoTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewUmAlgoOrderAlgoTypeParameter(val *NewUmAlgoOrderAlgoTypeParameter) *NullableNewUmAlgoOrderAlgoTypeParameter {
	return &NullableNewUmAlgoOrderAlgoTypeParameter{value: val, isSet: true}
}

func (v NullableNewUmAlgoOrderAlgoTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewUmAlgoOrderAlgoTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
