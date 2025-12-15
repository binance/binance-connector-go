/*
Binance Derivatives Trading COIN Futures WebSocket API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures WebSocket API

API version: 1.0.0
*/

package models

import (
	"encoding/json"
	"fmt"
)

// ModifyOrderPriceMatchParameter the model 'ModifyOrderPriceMatchParameter'
type ModifyOrderPriceMatchParameter string

// List of modifyOrder_priceMatch_parameter
const (
	ModifyOrderPriceMatchParameterNone       ModifyOrderPriceMatchParameter = "NONE"
	ModifyOrderPriceMatchParameterOpponent   ModifyOrderPriceMatchParameter = "OPPONENT"
	ModifyOrderPriceMatchParameterOpponent5  ModifyOrderPriceMatchParameter = "OPPONENT_5"
	ModifyOrderPriceMatchParameterOpponent10 ModifyOrderPriceMatchParameter = "OPPONENT_10"
	ModifyOrderPriceMatchParameterOpponent20 ModifyOrderPriceMatchParameter = "OPPONENT_20"
	ModifyOrderPriceMatchParameterQueue      ModifyOrderPriceMatchParameter = "QUEUE"
	ModifyOrderPriceMatchParameterQueue5     ModifyOrderPriceMatchParameter = "QUEUE_5"
	ModifyOrderPriceMatchParameterQueue10    ModifyOrderPriceMatchParameter = "QUEUE_10"
	ModifyOrderPriceMatchParameterQueue20    ModifyOrderPriceMatchParameter = "QUEUE_20"
)

// All allowed values of ModifyOrderPriceMatchParameter enum
var AllowedModifyOrderPriceMatchParameterEnumValues = []ModifyOrderPriceMatchParameter{
	"NONE",
	"OPPONENT",
	"OPPONENT_5",
	"OPPONENT_10",
	"OPPONENT_20",
	"QUEUE",
	"QUEUE_5",
	"QUEUE_10",
	"QUEUE_20",
}

func (v *ModifyOrderPriceMatchParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ModifyOrderPriceMatchParameter(value)
	for _, existing := range AllowedModifyOrderPriceMatchParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ModifyOrderPriceMatchParameter", value)
}

// NewModifyOrderPriceMatchParameterFromValue returns a pointer to a valid ModifyOrderPriceMatchParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModifyOrderPriceMatchParameterFromValue(v string) (*ModifyOrderPriceMatchParameter, error) {
	ev := ModifyOrderPriceMatchParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ModifyOrderPriceMatchParameter: valid values are %v", v, AllowedModifyOrderPriceMatchParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ModifyOrderPriceMatchParameter) IsValid() bool {
	for _, existing := range AllowedModifyOrderPriceMatchParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to modifyOrder_priceMatch_parameter value
func (v ModifyOrderPriceMatchParameter) Ptr() *ModifyOrderPriceMatchParameter {
	return &v
}

type NullableModifyOrderPriceMatchParameter struct {
	value *ModifyOrderPriceMatchParameter
	isSet bool
}

func (v NullableModifyOrderPriceMatchParameter) Get() *ModifyOrderPriceMatchParameter {
	return v.value
}

func (v *NullableModifyOrderPriceMatchParameter) Set(val *ModifyOrderPriceMatchParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyOrderPriceMatchParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyOrderPriceMatchParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyOrderPriceMatchParameter(val *ModifyOrderPriceMatchParameter) *NullableModifyOrderPriceMatchParameter {
	return &NullableModifyOrderPriceMatchParameter{value: val, isSet: true}
}

func (v NullableModifyOrderPriceMatchParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyOrderPriceMatchParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
