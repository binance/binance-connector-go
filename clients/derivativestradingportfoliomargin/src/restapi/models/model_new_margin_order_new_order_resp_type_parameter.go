/*
Portfolio Margin REST API

Access account information, manage margin positions, and trade with Binance Portfolio Margin.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewMarginOrderNewOrderRespTypeParameter the model 'NewMarginOrderNewOrderRespTypeParameter'
type NewMarginOrderNewOrderRespTypeParameter string

// List of newMarginOrder_newOrderRespType_parameter
const (
	NewMarginOrderNewOrderRespTypeParameterAck    NewMarginOrderNewOrderRespTypeParameter = "ACK"
	NewMarginOrderNewOrderRespTypeParameterResult NewMarginOrderNewOrderRespTypeParameter = "RESULT"
	NewMarginOrderNewOrderRespTypeParameterFull   NewMarginOrderNewOrderRespTypeParameter = "FULL"
)

// All allowed values of NewMarginOrderNewOrderRespTypeParameter enum
var AllowedNewMarginOrderNewOrderRespTypeParameterEnumValues = []NewMarginOrderNewOrderRespTypeParameter{
	"ACK",
	"RESULT",
	"FULL",
}

func (v *NewMarginOrderNewOrderRespTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewMarginOrderNewOrderRespTypeParameter(value)
	for _, existing := range AllowedNewMarginOrderNewOrderRespTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewMarginOrderNewOrderRespTypeParameter", value)
}

// NewNewMarginOrderNewOrderRespTypeParameterFromValue returns a pointer to a valid NewMarginOrderNewOrderRespTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewMarginOrderNewOrderRespTypeParameterFromValue(v string) (*NewMarginOrderNewOrderRespTypeParameter, error) {
	ev := NewMarginOrderNewOrderRespTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewMarginOrderNewOrderRespTypeParameter: valid values are %v", v, AllowedNewMarginOrderNewOrderRespTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewMarginOrderNewOrderRespTypeParameter) IsValid() bool {
	for _, existing := range AllowedNewMarginOrderNewOrderRespTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newMarginOrder_newOrderRespType_parameter value
func (v NewMarginOrderNewOrderRespTypeParameter) Ptr() *NewMarginOrderNewOrderRespTypeParameter {
	return &v
}

type NullableNewMarginOrderNewOrderRespTypeParameter struct {
	value *NewMarginOrderNewOrderRespTypeParameter
	isSet bool
}

func (v NullableNewMarginOrderNewOrderRespTypeParameter) Get() *NewMarginOrderNewOrderRespTypeParameter {
	return v.value
}

func (v *NullableNewMarginOrderNewOrderRespTypeParameter) Set(val *NewMarginOrderNewOrderRespTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewMarginOrderNewOrderRespTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewMarginOrderNewOrderRespTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewMarginOrderNewOrderRespTypeParameter(val *NewMarginOrderNewOrderRespTypeParameter) *NullableNewMarginOrderNewOrderRespTypeParameter {
	return &NullableNewMarginOrderNewOrderRespTypeParameter{value: val, isSet: true}
}

func (v NullableNewMarginOrderNewOrderRespTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewMarginOrderNewOrderRespTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
