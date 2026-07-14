/*
Portfolio Margin REST API

Access account information, manage margin positions, and trade with Binance Portfolio Margin.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewUmAlgoOrderPriceMatchParameter the model 'NewUmAlgoOrderPriceMatchParameter'
type NewUmAlgoOrderPriceMatchParameter string

// List of newUmAlgoOrder_priceMatch_parameter
const (
	NewUmAlgoOrderPriceMatchParameterOpponent   NewUmAlgoOrderPriceMatchParameter = "OPPONENT"
	NewUmAlgoOrderPriceMatchParameterOpponent5  NewUmAlgoOrderPriceMatchParameter = "OPPONENT_5"
	NewUmAlgoOrderPriceMatchParameterOpponent10 NewUmAlgoOrderPriceMatchParameter = "OPPONENT_10"
	NewUmAlgoOrderPriceMatchParameterOpponent20 NewUmAlgoOrderPriceMatchParameter = "OPPONENT_20"
	NewUmAlgoOrderPriceMatchParameterQueue      NewUmAlgoOrderPriceMatchParameter = "QUEUE"
	NewUmAlgoOrderPriceMatchParameterQueue5     NewUmAlgoOrderPriceMatchParameter = "QUEUE_5"
	NewUmAlgoOrderPriceMatchParameterQueue10    NewUmAlgoOrderPriceMatchParameter = "QUEUE_10"
	NewUmAlgoOrderPriceMatchParameterQueue20    NewUmAlgoOrderPriceMatchParameter = "QUEUE_20"
)

// All allowed values of NewUmAlgoOrderPriceMatchParameter enum
var AllowedNewUmAlgoOrderPriceMatchParameterEnumValues = []NewUmAlgoOrderPriceMatchParameter{
	"OPPONENT",
	"OPPONENT_5",
	"OPPONENT_10",
	"OPPONENT_20",
	"QUEUE",
	"QUEUE_5",
	"QUEUE_10",
	"QUEUE_20",
}

func (v *NewUmAlgoOrderPriceMatchParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewUmAlgoOrderPriceMatchParameter(value)
	for _, existing := range AllowedNewUmAlgoOrderPriceMatchParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewUmAlgoOrderPriceMatchParameter", value)
}

// NewNewUmAlgoOrderPriceMatchParameterFromValue returns a pointer to a valid NewUmAlgoOrderPriceMatchParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewUmAlgoOrderPriceMatchParameterFromValue(v string) (*NewUmAlgoOrderPriceMatchParameter, error) {
	ev := NewUmAlgoOrderPriceMatchParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewUmAlgoOrderPriceMatchParameter: valid values are %v", v, AllowedNewUmAlgoOrderPriceMatchParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewUmAlgoOrderPriceMatchParameter) IsValid() bool {
	for _, existing := range AllowedNewUmAlgoOrderPriceMatchParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newUmAlgoOrder_priceMatch_parameter value
func (v NewUmAlgoOrderPriceMatchParameter) Ptr() *NewUmAlgoOrderPriceMatchParameter {
	return &v
}

type NullableNewUmAlgoOrderPriceMatchParameter struct {
	value *NewUmAlgoOrderPriceMatchParameter
	isSet bool
}

func (v NullableNewUmAlgoOrderPriceMatchParameter) Get() *NewUmAlgoOrderPriceMatchParameter {
	return v.value
}

func (v *NullableNewUmAlgoOrderPriceMatchParameter) Set(val *NewUmAlgoOrderPriceMatchParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewUmAlgoOrderPriceMatchParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewUmAlgoOrderPriceMatchParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewUmAlgoOrderPriceMatchParameter(val *NewUmAlgoOrderPriceMatchParameter) *NullableNewUmAlgoOrderPriceMatchParameter {
	return &NullableNewUmAlgoOrderPriceMatchParameter{value: val, isSet: true}
}

func (v NullableNewUmAlgoOrderPriceMatchParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewUmAlgoOrderPriceMatchParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
