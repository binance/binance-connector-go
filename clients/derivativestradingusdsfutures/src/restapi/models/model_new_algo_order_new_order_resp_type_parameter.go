/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewAlgoOrderNewOrderRespTypeParameter the model 'NewAlgoOrderNewOrderRespTypeParameter'
type NewAlgoOrderNewOrderRespTypeParameter string

// List of newAlgoOrder_newOrderRespType_parameter
const (
	NewAlgoOrderNewOrderRespTypeParameterAck    NewAlgoOrderNewOrderRespTypeParameter = "ACK"
	NewAlgoOrderNewOrderRespTypeParameterResult NewAlgoOrderNewOrderRespTypeParameter = "RESULT"
)

// All allowed values of NewAlgoOrderNewOrderRespTypeParameter enum
var AllowedNewAlgoOrderNewOrderRespTypeParameterEnumValues = []NewAlgoOrderNewOrderRespTypeParameter{
	"ACK",
	"RESULT",
}

func (v *NewAlgoOrderNewOrderRespTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewAlgoOrderNewOrderRespTypeParameter(value)
	for _, existing := range AllowedNewAlgoOrderNewOrderRespTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewAlgoOrderNewOrderRespTypeParameter", value)
}

// NewNewAlgoOrderNewOrderRespTypeParameterFromValue returns a pointer to a valid NewAlgoOrderNewOrderRespTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewAlgoOrderNewOrderRespTypeParameterFromValue(v string) (*NewAlgoOrderNewOrderRespTypeParameter, error) {
	ev := NewAlgoOrderNewOrderRespTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewAlgoOrderNewOrderRespTypeParameter: valid values are %v", v, AllowedNewAlgoOrderNewOrderRespTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewAlgoOrderNewOrderRespTypeParameter) IsValid() bool {
	for _, existing := range AllowedNewAlgoOrderNewOrderRespTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newAlgoOrder_newOrderRespType_parameter value
func (v NewAlgoOrderNewOrderRespTypeParameter) Ptr() *NewAlgoOrderNewOrderRespTypeParameter {
	return &v
}

type NullableNewAlgoOrderNewOrderRespTypeParameter struct {
	value *NewAlgoOrderNewOrderRespTypeParameter
	isSet bool
}

func (v NullableNewAlgoOrderNewOrderRespTypeParameter) Get() *NewAlgoOrderNewOrderRespTypeParameter {
	return v.value
}

func (v *NullableNewAlgoOrderNewOrderRespTypeParameter) Set(val *NewAlgoOrderNewOrderRespTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewAlgoOrderNewOrderRespTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewAlgoOrderNewOrderRespTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewAlgoOrderNewOrderRespTypeParameter(val *NewAlgoOrderNewOrderRespTypeParameter) *NullableNewAlgoOrderNewOrderRespTypeParameter {
	return &NullableNewAlgoOrderNewOrderRespTypeParameter{value: val, isSet: true}
}

func (v NullableNewAlgoOrderNewOrderRespTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewAlgoOrderNewOrderRespTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
