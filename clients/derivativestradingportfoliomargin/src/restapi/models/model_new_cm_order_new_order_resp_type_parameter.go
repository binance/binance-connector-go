/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewCmOrderNewOrderRespTypeParameter the model 'NewCmOrderNewOrderRespTypeParameter'
type NewCmOrderNewOrderRespTypeParameter string

// List of newCmOrder_newOrderRespType_parameter
const (
	NewCmOrderNewOrderRespTypeParameterAck    NewCmOrderNewOrderRespTypeParameter = "ACK"
	NewCmOrderNewOrderRespTypeParameterResult NewCmOrderNewOrderRespTypeParameter = "RESULT"
)

// All allowed values of NewCmOrderNewOrderRespTypeParameter enum
var AllowedNewCmOrderNewOrderRespTypeParameterEnumValues = []NewCmOrderNewOrderRespTypeParameter{
	"ACK",
	"RESULT",
}

func (v *NewCmOrderNewOrderRespTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewCmOrderNewOrderRespTypeParameter(value)
	for _, existing := range AllowedNewCmOrderNewOrderRespTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewCmOrderNewOrderRespTypeParameter", value)
}

// NewNewCmOrderNewOrderRespTypeParameterFromValue returns a pointer to a valid NewCmOrderNewOrderRespTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewCmOrderNewOrderRespTypeParameterFromValue(v string) (*NewCmOrderNewOrderRespTypeParameter, error) {
	ev := NewCmOrderNewOrderRespTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewCmOrderNewOrderRespTypeParameter: valid values are %v", v, AllowedNewCmOrderNewOrderRespTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewCmOrderNewOrderRespTypeParameter) IsValid() bool {
	for _, existing := range AllowedNewCmOrderNewOrderRespTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newCmOrder_newOrderRespType_parameter value
func (v NewCmOrderNewOrderRespTypeParameter) Ptr() *NewCmOrderNewOrderRespTypeParameter {
	return &v
}

type NullableNewCmOrderNewOrderRespTypeParameter struct {
	value *NewCmOrderNewOrderRespTypeParameter
	isSet bool
}

func (v NullableNewCmOrderNewOrderRespTypeParameter) Get() *NewCmOrderNewOrderRespTypeParameter {
	return v.value
}

func (v *NullableNewCmOrderNewOrderRespTypeParameter) Set(val *NewCmOrderNewOrderRespTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewCmOrderNewOrderRespTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewCmOrderNewOrderRespTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewCmOrderNewOrderRespTypeParameter(val *NewCmOrderNewOrderRespTypeParameter) *NullableNewCmOrderNewOrderRespTypeParameter {
	return &NullableNewCmOrderNewOrderRespTypeParameter{value: val, isSet: true}
}

func (v NullableNewCmOrderNewOrderRespTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewCmOrderNewOrderRespTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
