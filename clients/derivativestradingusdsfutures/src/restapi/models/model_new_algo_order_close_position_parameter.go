/*
Futures (USDⓈ-M) REST API

Access market data, manage accounts, and trade USDⓈ-M perpetual futures.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewAlgoOrderClosePositionParameter the model 'NewAlgoOrderClosePositionParameter'
type NewAlgoOrderClosePositionParameter string

// List of newAlgoOrder_closePosition_parameter
const (
	NewAlgoOrderClosePositionParameterTrue  NewAlgoOrderClosePositionParameter = "true"
	NewAlgoOrderClosePositionParameterFalse NewAlgoOrderClosePositionParameter = "false"
)

// All allowed values of NewAlgoOrderClosePositionParameter enum
var AllowedNewAlgoOrderClosePositionParameterEnumValues = []NewAlgoOrderClosePositionParameter{
	"true",
	"false",
}

func (v *NewAlgoOrderClosePositionParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewAlgoOrderClosePositionParameter(value)
	for _, existing := range AllowedNewAlgoOrderClosePositionParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewAlgoOrderClosePositionParameter", value)
}

// NewNewAlgoOrderClosePositionParameterFromValue returns a pointer to a valid NewAlgoOrderClosePositionParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewAlgoOrderClosePositionParameterFromValue(v string) (*NewAlgoOrderClosePositionParameter, error) {
	ev := NewAlgoOrderClosePositionParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewAlgoOrderClosePositionParameter: valid values are %v", v, AllowedNewAlgoOrderClosePositionParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewAlgoOrderClosePositionParameter) IsValid() bool {
	for _, existing := range AllowedNewAlgoOrderClosePositionParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newAlgoOrder_closePosition_parameter value
func (v NewAlgoOrderClosePositionParameter) Ptr() *NewAlgoOrderClosePositionParameter {
	return &v
}

type NullableNewAlgoOrderClosePositionParameter struct {
	value *NewAlgoOrderClosePositionParameter
	isSet bool
}

func (v NullableNewAlgoOrderClosePositionParameter) Get() *NewAlgoOrderClosePositionParameter {
	return v.value
}

func (v *NullableNewAlgoOrderClosePositionParameter) Set(val *NewAlgoOrderClosePositionParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewAlgoOrderClosePositionParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewAlgoOrderClosePositionParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewAlgoOrderClosePositionParameter(val *NewAlgoOrderClosePositionParameter) *NullableNewAlgoOrderClosePositionParameter {
	return &NullableNewAlgoOrderClosePositionParameter{value: val, isSet: true}
}

func (v NullableNewAlgoOrderClosePositionParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewAlgoOrderClosePositionParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
