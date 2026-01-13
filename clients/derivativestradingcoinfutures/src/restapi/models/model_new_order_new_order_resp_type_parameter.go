/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewOrderNewOrderRespTypeParameter the model 'NewOrderNewOrderRespTypeParameter'
type NewOrderNewOrderRespTypeParameter string

// List of newOrder_newOrderRespType_parameter
const (
	NewOrderNewOrderRespTypeParameterAck    NewOrderNewOrderRespTypeParameter = "ACK"
	NewOrderNewOrderRespTypeParameterResult NewOrderNewOrderRespTypeParameter = "RESULT"
)

// All allowed values of NewOrderNewOrderRespTypeParameter enum
var AllowedNewOrderNewOrderRespTypeParameterEnumValues = []NewOrderNewOrderRespTypeParameter{
	"ACK",
	"RESULT",
}

func (v *NewOrderNewOrderRespTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewOrderNewOrderRespTypeParameter(value)
	for _, existing := range AllowedNewOrderNewOrderRespTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewOrderNewOrderRespTypeParameter", value)
}

// NewNewOrderNewOrderRespTypeParameterFromValue returns a pointer to a valid NewOrderNewOrderRespTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewOrderNewOrderRespTypeParameterFromValue(v string) (*NewOrderNewOrderRespTypeParameter, error) {
	ev := NewOrderNewOrderRespTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewOrderNewOrderRespTypeParameter: valid values are %v", v, AllowedNewOrderNewOrderRespTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewOrderNewOrderRespTypeParameter) IsValid() bool {
	for _, existing := range AllowedNewOrderNewOrderRespTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newOrder_newOrderRespType_parameter value
func (v NewOrderNewOrderRespTypeParameter) Ptr() *NewOrderNewOrderRespTypeParameter {
	return &v
}

type NullableNewOrderNewOrderRespTypeParameter struct {
	value *NewOrderNewOrderRespTypeParameter
	isSet bool
}

func (v NullableNewOrderNewOrderRespTypeParameter) Get() *NewOrderNewOrderRespTypeParameter {
	return v.value
}

func (v *NullableNewOrderNewOrderRespTypeParameter) Set(val *NewOrderNewOrderRespTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewOrderNewOrderRespTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewOrderNewOrderRespTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewOrderNewOrderRespTypeParameter(val *NewOrderNewOrderRespTypeParameter) *NullableNewOrderNewOrderRespTypeParameter {
	return &NullableNewOrderNewOrderRespTypeParameter{value: val, isSet: true}
}

func (v NullableNewOrderNewOrderRespTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewOrderNewOrderRespTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
