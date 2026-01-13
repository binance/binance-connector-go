/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewOrderPositionSideParameter the model 'NewOrderPositionSideParameter'
type NewOrderPositionSideParameter string

// List of newOrder_positionSide_parameter
const (
	NewOrderPositionSideParameterBoth  NewOrderPositionSideParameter = "BOTH"
	NewOrderPositionSideParameterLong  NewOrderPositionSideParameter = "LONG"
	NewOrderPositionSideParameterShort NewOrderPositionSideParameter = "SHORT"
)

// All allowed values of NewOrderPositionSideParameter enum
var AllowedNewOrderPositionSideParameterEnumValues = []NewOrderPositionSideParameter{
	"BOTH",
	"LONG",
	"SHORT",
}

func (v *NewOrderPositionSideParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewOrderPositionSideParameter(value)
	for _, existing := range AllowedNewOrderPositionSideParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewOrderPositionSideParameter", value)
}

// NewNewOrderPositionSideParameterFromValue returns a pointer to a valid NewOrderPositionSideParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewOrderPositionSideParameterFromValue(v string) (*NewOrderPositionSideParameter, error) {
	ev := NewOrderPositionSideParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewOrderPositionSideParameter: valid values are %v", v, AllowedNewOrderPositionSideParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewOrderPositionSideParameter) IsValid() bool {
	for _, existing := range AllowedNewOrderPositionSideParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newOrder_positionSide_parameter value
func (v NewOrderPositionSideParameter) Ptr() *NewOrderPositionSideParameter {
	return &v
}

type NullableNewOrderPositionSideParameter struct {
	value *NewOrderPositionSideParameter
	isSet bool
}

func (v NullableNewOrderPositionSideParameter) Get() *NewOrderPositionSideParameter {
	return v.value
}

func (v *NullableNewOrderPositionSideParameter) Set(val *NewOrderPositionSideParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewOrderPositionSideParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewOrderPositionSideParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewOrderPositionSideParameter(val *NewOrderPositionSideParameter) *NullableNewOrderPositionSideParameter {
	return &NullableNewOrderPositionSideParameter{value: val, isSet: true}
}

func (v NullableNewOrderPositionSideParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewOrderPositionSideParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
