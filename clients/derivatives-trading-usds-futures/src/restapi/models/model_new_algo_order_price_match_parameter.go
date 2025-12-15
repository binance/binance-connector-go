/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewAlgoOrderPriceMatchParameter the model 'NewAlgoOrderPriceMatchParameter'
type NewAlgoOrderPriceMatchParameter string

// List of newAlgoOrder_priceMatch_parameter
const (
	NewAlgoOrderPriceMatchParameterNone       NewAlgoOrderPriceMatchParameter = "NONE"
	NewAlgoOrderPriceMatchParameterOpponent   NewAlgoOrderPriceMatchParameter = "OPPONENT"
	NewAlgoOrderPriceMatchParameterOpponent5  NewAlgoOrderPriceMatchParameter = "OPPONENT_5"
	NewAlgoOrderPriceMatchParameterOpponent10 NewAlgoOrderPriceMatchParameter = "OPPONENT_10"
	NewAlgoOrderPriceMatchParameterOpponent20 NewAlgoOrderPriceMatchParameter = "OPPONENT_20"
	NewAlgoOrderPriceMatchParameterQueue      NewAlgoOrderPriceMatchParameter = "QUEUE"
	NewAlgoOrderPriceMatchParameterQueue5     NewAlgoOrderPriceMatchParameter = "QUEUE_5"
	NewAlgoOrderPriceMatchParameterQueue10    NewAlgoOrderPriceMatchParameter = "QUEUE_10"
	NewAlgoOrderPriceMatchParameterQueue20    NewAlgoOrderPriceMatchParameter = "QUEUE_20"
)

// All allowed values of NewAlgoOrderPriceMatchParameter enum
var AllowedNewAlgoOrderPriceMatchParameterEnumValues = []NewAlgoOrderPriceMatchParameter{
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

func (v *NewAlgoOrderPriceMatchParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewAlgoOrderPriceMatchParameter(value)
	for _, existing := range AllowedNewAlgoOrderPriceMatchParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewAlgoOrderPriceMatchParameter", value)
}

// NewNewAlgoOrderPriceMatchParameterFromValue returns a pointer to a valid NewAlgoOrderPriceMatchParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewAlgoOrderPriceMatchParameterFromValue(v string) (*NewAlgoOrderPriceMatchParameter, error) {
	ev := NewAlgoOrderPriceMatchParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewAlgoOrderPriceMatchParameter: valid values are %v", v, AllowedNewAlgoOrderPriceMatchParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewAlgoOrderPriceMatchParameter) IsValid() bool {
	for _, existing := range AllowedNewAlgoOrderPriceMatchParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newAlgoOrder_priceMatch_parameter value
func (v NewAlgoOrderPriceMatchParameter) Ptr() *NewAlgoOrderPriceMatchParameter {
	return &v
}

type NullableNewAlgoOrderPriceMatchParameter struct {
	value *NewAlgoOrderPriceMatchParameter
	isSet bool
}

func (v NullableNewAlgoOrderPriceMatchParameter) Get() *NewAlgoOrderPriceMatchParameter {
	return v.value
}

func (v *NullableNewAlgoOrderPriceMatchParameter) Set(val *NewAlgoOrderPriceMatchParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewAlgoOrderPriceMatchParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewAlgoOrderPriceMatchParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewAlgoOrderPriceMatchParameter(val *NewAlgoOrderPriceMatchParameter) *NullableNewAlgoOrderPriceMatchParameter {
	return &NullableNewAlgoOrderPriceMatchParameter{value: val, isSet: true}
}

func (v NullableNewAlgoOrderPriceMatchParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewAlgoOrderPriceMatchParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
