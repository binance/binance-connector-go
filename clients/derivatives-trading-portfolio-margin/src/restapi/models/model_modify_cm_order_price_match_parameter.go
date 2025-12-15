/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"
	"fmt"
)

// ModifyCmOrderPriceMatchParameter the model 'ModifyCmOrderPriceMatchParameter'
type ModifyCmOrderPriceMatchParameter string

// List of modifyCmOrder_priceMatch_parameter
const (
	ModifyCmOrderPriceMatchParameterNone       ModifyCmOrderPriceMatchParameter = "NONE"
	ModifyCmOrderPriceMatchParameterOpponent   ModifyCmOrderPriceMatchParameter = "OPPONENT"
	ModifyCmOrderPriceMatchParameterOpponent5  ModifyCmOrderPriceMatchParameter = "OPPONENT_5"
	ModifyCmOrderPriceMatchParameterOpponent10 ModifyCmOrderPriceMatchParameter = "OPPONENT_10"
	ModifyCmOrderPriceMatchParameterOpponent20 ModifyCmOrderPriceMatchParameter = "OPPONENT_20"
	ModifyCmOrderPriceMatchParameterQueue      ModifyCmOrderPriceMatchParameter = "QUEUE"
	ModifyCmOrderPriceMatchParameterQueue5     ModifyCmOrderPriceMatchParameter = "QUEUE_5"
	ModifyCmOrderPriceMatchParameterQueue10    ModifyCmOrderPriceMatchParameter = "QUEUE_10"
	ModifyCmOrderPriceMatchParameterQueue20    ModifyCmOrderPriceMatchParameter = "QUEUE_20"
)

// All allowed values of ModifyCmOrderPriceMatchParameter enum
var AllowedModifyCmOrderPriceMatchParameterEnumValues = []ModifyCmOrderPriceMatchParameter{
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

func (v *ModifyCmOrderPriceMatchParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ModifyCmOrderPriceMatchParameter(value)
	for _, existing := range AllowedModifyCmOrderPriceMatchParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ModifyCmOrderPriceMatchParameter", value)
}

// NewModifyCmOrderPriceMatchParameterFromValue returns a pointer to a valid ModifyCmOrderPriceMatchParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModifyCmOrderPriceMatchParameterFromValue(v string) (*ModifyCmOrderPriceMatchParameter, error) {
	ev := ModifyCmOrderPriceMatchParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ModifyCmOrderPriceMatchParameter: valid values are %v", v, AllowedModifyCmOrderPriceMatchParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ModifyCmOrderPriceMatchParameter) IsValid() bool {
	for _, existing := range AllowedModifyCmOrderPriceMatchParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to modifyCmOrder_priceMatch_parameter value
func (v ModifyCmOrderPriceMatchParameter) Ptr() *ModifyCmOrderPriceMatchParameter {
	return &v
}

type NullableModifyCmOrderPriceMatchParameter struct {
	value *ModifyCmOrderPriceMatchParameter
	isSet bool
}

func (v NullableModifyCmOrderPriceMatchParameter) Get() *ModifyCmOrderPriceMatchParameter {
	return v.value
}

func (v *NullableModifyCmOrderPriceMatchParameter) Set(val *ModifyCmOrderPriceMatchParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyCmOrderPriceMatchParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyCmOrderPriceMatchParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyCmOrderPriceMatchParameter(val *ModifyCmOrderPriceMatchParameter) *NullableModifyCmOrderPriceMatchParameter {
	return &NullableModifyCmOrderPriceMatchParameter{value: val, isSet: true}
}

func (v NullableModifyCmOrderPriceMatchParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyCmOrderPriceMatchParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
