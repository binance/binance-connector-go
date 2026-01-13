/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewOrderTypeParameter the model 'NewOrderTypeParameter'
type NewOrderTypeParameter string

// List of newOrder_type_parameter
const (
	NewOrderTypeParameterLimit              NewOrderTypeParameter = "LIMIT"
	NewOrderTypeParameterMarket             NewOrderTypeParameter = "MARKET"
	NewOrderTypeParameterStop               NewOrderTypeParameter = "STOP"
	NewOrderTypeParameterStopMarket         NewOrderTypeParameter = "STOP_MARKET"
	NewOrderTypeParameterTakeProfit         NewOrderTypeParameter = "TAKE_PROFIT"
	NewOrderTypeParameterTakeProfitMarket   NewOrderTypeParameter = "TAKE_PROFIT_MARKET"
	NewOrderTypeParameterTrailingStopMarket NewOrderTypeParameter = "TRAILING_STOP_MARKET"
)

// All allowed values of NewOrderTypeParameter enum
var AllowedNewOrderTypeParameterEnumValues = []NewOrderTypeParameter{
	"LIMIT",
	"MARKET",
	"STOP",
	"STOP_MARKET",
	"TAKE_PROFIT",
	"TAKE_PROFIT_MARKET",
	"TRAILING_STOP_MARKET",
}

func (v *NewOrderTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewOrderTypeParameter(value)
	for _, existing := range AllowedNewOrderTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewOrderTypeParameter", value)
}

// NewNewOrderTypeParameterFromValue returns a pointer to a valid NewOrderTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewOrderTypeParameterFromValue(v string) (*NewOrderTypeParameter, error) {
	ev := NewOrderTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewOrderTypeParameter: valid values are %v", v, AllowedNewOrderTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewOrderTypeParameter) IsValid() bool {
	for _, existing := range AllowedNewOrderTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newOrder_type_parameter value
func (v NewOrderTypeParameter) Ptr() *NewOrderTypeParameter {
	return &v
}

type NullableNewOrderTypeParameter struct {
	value *NewOrderTypeParameter
	isSet bool
}

func (v NullableNewOrderTypeParameter) Get() *NewOrderTypeParameter {
	return v.value
}

func (v *NullableNewOrderTypeParameter) Set(val *NewOrderTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewOrderTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewOrderTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewOrderTypeParameter(val *NewOrderTypeParameter) *NullableNewOrderTypeParameter {
	return &NullableNewOrderTypeParameter{value: val, isSet: true}
}

func (v NullableNewOrderTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewOrderTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
