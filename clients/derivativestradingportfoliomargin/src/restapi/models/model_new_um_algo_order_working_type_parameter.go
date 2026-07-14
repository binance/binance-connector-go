/*
Portfolio Margin REST API

Access account information, manage margin positions, and trade with Binance Portfolio Margin.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewUmAlgoOrderWorkingTypeParameter the model 'NewUmAlgoOrderWorkingTypeParameter'
type NewUmAlgoOrderWorkingTypeParameter string

// List of newUmAlgoOrder_workingType_parameter
const (
	NewUmAlgoOrderWorkingTypeParameterMarkPrice     NewUmAlgoOrderWorkingTypeParameter = "MARK_PRICE"
	NewUmAlgoOrderWorkingTypeParameterContractPrice NewUmAlgoOrderWorkingTypeParameter = "CONTRACT_PRICE"
)

// All allowed values of NewUmAlgoOrderWorkingTypeParameter enum
var AllowedNewUmAlgoOrderWorkingTypeParameterEnumValues = []NewUmAlgoOrderWorkingTypeParameter{
	"MARK_PRICE",
	"CONTRACT_PRICE",
}

func (v *NewUmAlgoOrderWorkingTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewUmAlgoOrderWorkingTypeParameter(value)
	for _, existing := range AllowedNewUmAlgoOrderWorkingTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewUmAlgoOrderWorkingTypeParameter", value)
}

// NewNewUmAlgoOrderWorkingTypeParameterFromValue returns a pointer to a valid NewUmAlgoOrderWorkingTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewUmAlgoOrderWorkingTypeParameterFromValue(v string) (*NewUmAlgoOrderWorkingTypeParameter, error) {
	ev := NewUmAlgoOrderWorkingTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewUmAlgoOrderWorkingTypeParameter: valid values are %v", v, AllowedNewUmAlgoOrderWorkingTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewUmAlgoOrderWorkingTypeParameter) IsValid() bool {
	for _, existing := range AllowedNewUmAlgoOrderWorkingTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newUmAlgoOrder_workingType_parameter value
func (v NewUmAlgoOrderWorkingTypeParameter) Ptr() *NewUmAlgoOrderWorkingTypeParameter {
	return &v
}

type NullableNewUmAlgoOrderWorkingTypeParameter struct {
	value *NewUmAlgoOrderWorkingTypeParameter
	isSet bool
}

func (v NullableNewUmAlgoOrderWorkingTypeParameter) Get() *NewUmAlgoOrderWorkingTypeParameter {
	return v.value
}

func (v *NullableNewUmAlgoOrderWorkingTypeParameter) Set(val *NewUmAlgoOrderWorkingTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewUmAlgoOrderWorkingTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewUmAlgoOrderWorkingTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewUmAlgoOrderWorkingTypeParameter(val *NewUmAlgoOrderWorkingTypeParameter) *NullableNewUmAlgoOrderWorkingTypeParameter {
	return &NullableNewUmAlgoOrderWorkingTypeParameter{value: val, isSet: true}
}

func (v NullableNewUmAlgoOrderWorkingTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewUmAlgoOrderWorkingTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
