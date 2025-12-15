/*
Binance Derivatives Trading USDS Futures WebSocket API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket API

API version: 1.0.0
*/

package models

import (
	"encoding/json"
	"fmt"
)

// ModifyOrderSideParameter the model 'ModifyOrderSideParameter'
type ModifyOrderSideParameter string

// List of modifyOrder_side_parameter
const (
	ModifyOrderSideParameterBuy  ModifyOrderSideParameter = "BUY"
	ModifyOrderSideParameterSell ModifyOrderSideParameter = "SELL"
)

// All allowed values of ModifyOrderSideParameter enum
var AllowedModifyOrderSideParameterEnumValues = []ModifyOrderSideParameter{
	"BUY",
	"SELL",
}

func (v *ModifyOrderSideParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ModifyOrderSideParameter(value)
	for _, existing := range AllowedModifyOrderSideParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ModifyOrderSideParameter", value)
}

// NewModifyOrderSideParameterFromValue returns a pointer to a valid ModifyOrderSideParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModifyOrderSideParameterFromValue(v string) (*ModifyOrderSideParameter, error) {
	ev := ModifyOrderSideParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ModifyOrderSideParameter: valid values are %v", v, AllowedModifyOrderSideParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ModifyOrderSideParameter) IsValid() bool {
	for _, existing := range AllowedModifyOrderSideParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to modifyOrder_side_parameter value
func (v ModifyOrderSideParameter) Ptr() *ModifyOrderSideParameter {
	return &v
}

type NullableModifyOrderSideParameter struct {
	value *ModifyOrderSideParameter
	isSet bool
}

func (v NullableModifyOrderSideParameter) Get() *ModifyOrderSideParameter {
	return v.value
}

func (v *NullableModifyOrderSideParameter) Set(val *ModifyOrderSideParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyOrderSideParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyOrderSideParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyOrderSideParameter(val *ModifyOrderSideParameter) *NullableModifyOrderSideParameter {
	return &NullableModifyOrderSideParameter{value: val, isSet: true}
}

func (v NullableModifyOrderSideParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyOrderSideParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
