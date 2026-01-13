/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewCmConditionalOrderStrategyTypeParameter the model 'NewCmConditionalOrderStrategyTypeParameter'
type NewCmConditionalOrderStrategyTypeParameter string

// List of newCmConditionalOrder_strategyType_parameter
const (
	NewCmConditionalOrderStrategyTypeParameterStop               NewCmConditionalOrderStrategyTypeParameter = "STOP"
	NewCmConditionalOrderStrategyTypeParameterStopMarket         NewCmConditionalOrderStrategyTypeParameter = "STOP_MARKET"
	NewCmConditionalOrderStrategyTypeParameterTakeProfit         NewCmConditionalOrderStrategyTypeParameter = "TAKE_PROFIT"
	NewCmConditionalOrderStrategyTypeParameterTakeProfitMarket   NewCmConditionalOrderStrategyTypeParameter = "TAKE_PROFIT_MARKET"
	NewCmConditionalOrderStrategyTypeParameterTrailingStopMarket NewCmConditionalOrderStrategyTypeParameter = "TRAILING_STOP_MARKET"
)

// All allowed values of NewCmConditionalOrderStrategyTypeParameter enum
var AllowedNewCmConditionalOrderStrategyTypeParameterEnumValues = []NewCmConditionalOrderStrategyTypeParameter{
	"STOP",
	"STOP_MARKET",
	"TAKE_PROFIT",
	"TAKE_PROFIT_MARKET",
	"TRAILING_STOP_MARKET",
}

func (v *NewCmConditionalOrderStrategyTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewCmConditionalOrderStrategyTypeParameter(value)
	for _, existing := range AllowedNewCmConditionalOrderStrategyTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewCmConditionalOrderStrategyTypeParameter", value)
}

// NewNewCmConditionalOrderStrategyTypeParameterFromValue returns a pointer to a valid NewCmConditionalOrderStrategyTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewCmConditionalOrderStrategyTypeParameterFromValue(v string) (*NewCmConditionalOrderStrategyTypeParameter, error) {
	ev := NewCmConditionalOrderStrategyTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewCmConditionalOrderStrategyTypeParameter: valid values are %v", v, AllowedNewCmConditionalOrderStrategyTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewCmConditionalOrderStrategyTypeParameter) IsValid() bool {
	for _, existing := range AllowedNewCmConditionalOrderStrategyTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newCmConditionalOrder_strategyType_parameter value
func (v NewCmConditionalOrderStrategyTypeParameter) Ptr() *NewCmConditionalOrderStrategyTypeParameter {
	return &v
}

type NullableNewCmConditionalOrderStrategyTypeParameter struct {
	value *NewCmConditionalOrderStrategyTypeParameter
	isSet bool
}

func (v NullableNewCmConditionalOrderStrategyTypeParameter) Get() *NewCmConditionalOrderStrategyTypeParameter {
	return v.value
}

func (v *NullableNewCmConditionalOrderStrategyTypeParameter) Set(val *NewCmConditionalOrderStrategyTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewCmConditionalOrderStrategyTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewCmConditionalOrderStrategyTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewCmConditionalOrderStrategyTypeParameter(val *NewCmConditionalOrderStrategyTypeParameter) *NullableNewCmConditionalOrderStrategyTypeParameter {
	return &NullableNewCmConditionalOrderStrategyTypeParameter{value: val, isSet: true}
}

func (v NullableNewCmConditionalOrderStrategyTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewCmConditionalOrderStrategyTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
