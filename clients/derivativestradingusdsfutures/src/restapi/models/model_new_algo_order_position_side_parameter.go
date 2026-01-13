/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewAlgoOrderPositionSideParameter the model 'NewAlgoOrderPositionSideParameter'
type NewAlgoOrderPositionSideParameter string

// List of newAlgoOrder_positionSide_parameter
const (
	NewAlgoOrderPositionSideParameterBoth  NewAlgoOrderPositionSideParameter = "BOTH"
	NewAlgoOrderPositionSideParameterLong  NewAlgoOrderPositionSideParameter = "LONG"
	NewAlgoOrderPositionSideParameterShort NewAlgoOrderPositionSideParameter = "SHORT"
)

// All allowed values of NewAlgoOrderPositionSideParameter enum
var AllowedNewAlgoOrderPositionSideParameterEnumValues = []NewAlgoOrderPositionSideParameter{
	"BOTH",
	"LONG",
	"SHORT",
}

func (v *NewAlgoOrderPositionSideParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewAlgoOrderPositionSideParameter(value)
	for _, existing := range AllowedNewAlgoOrderPositionSideParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewAlgoOrderPositionSideParameter", value)
}

// NewNewAlgoOrderPositionSideParameterFromValue returns a pointer to a valid NewAlgoOrderPositionSideParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewAlgoOrderPositionSideParameterFromValue(v string) (*NewAlgoOrderPositionSideParameter, error) {
	ev := NewAlgoOrderPositionSideParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewAlgoOrderPositionSideParameter: valid values are %v", v, AllowedNewAlgoOrderPositionSideParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewAlgoOrderPositionSideParameter) IsValid() bool {
	for _, existing := range AllowedNewAlgoOrderPositionSideParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newAlgoOrder_positionSide_parameter value
func (v NewAlgoOrderPositionSideParameter) Ptr() *NewAlgoOrderPositionSideParameter {
	return &v
}

type NullableNewAlgoOrderPositionSideParameter struct {
	value *NewAlgoOrderPositionSideParameter
	isSet bool
}

func (v NullableNewAlgoOrderPositionSideParameter) Get() *NewAlgoOrderPositionSideParameter {
	return v.value
}

func (v *NullableNewAlgoOrderPositionSideParameter) Set(val *NewAlgoOrderPositionSideParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewAlgoOrderPositionSideParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewAlgoOrderPositionSideParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewAlgoOrderPositionSideParameter(val *NewAlgoOrderPositionSideParameter) *NullableNewAlgoOrderPositionSideParameter {
	return &NullableNewAlgoOrderPositionSideParameter{value: val, isSet: true}
}

func (v NullableNewAlgoOrderPositionSideParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewAlgoOrderPositionSideParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
