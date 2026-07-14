/*
Futures (USDⓈ-M) REST API

Access market data, manage accounts, and trade USDⓈ-M perpetual futures.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewAlgoOrderTypeParameter the model 'NewAlgoOrderTypeParameter'
type NewAlgoOrderTypeParameter string

// List of newAlgoOrder_type_parameter
const (
	NewAlgoOrderTypeParameterLimit              NewAlgoOrderTypeParameter = "LIMIT"
	NewAlgoOrderTypeParameterMarket             NewAlgoOrderTypeParameter = "MARKET"
	NewAlgoOrderTypeParameterStop               NewAlgoOrderTypeParameter = "STOP"
	NewAlgoOrderTypeParameterStopMarket         NewAlgoOrderTypeParameter = "STOP_MARKET"
	NewAlgoOrderTypeParameterTakeProfit         NewAlgoOrderTypeParameter = "TAKE_PROFIT"
	NewAlgoOrderTypeParameterTakeProfitMarket   NewAlgoOrderTypeParameter = "TAKE_PROFIT_MARKET"
	NewAlgoOrderTypeParameterTrailingStopMarket NewAlgoOrderTypeParameter = "TRAILING_STOP_MARKET"
)

// All allowed values of NewAlgoOrderTypeParameter enum
var AllowedNewAlgoOrderTypeParameterEnumValues = []NewAlgoOrderTypeParameter{
	"LIMIT",
	"MARKET",
	"STOP",
	"STOP_MARKET",
	"TAKE_PROFIT",
	"TAKE_PROFIT_MARKET",
	"TRAILING_STOP_MARKET",
}

func (v *NewAlgoOrderTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewAlgoOrderTypeParameter(value)
	for _, existing := range AllowedNewAlgoOrderTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewAlgoOrderTypeParameter", value)
}

// NewNewAlgoOrderTypeParameterFromValue returns a pointer to a valid NewAlgoOrderTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewAlgoOrderTypeParameterFromValue(v string) (*NewAlgoOrderTypeParameter, error) {
	ev := NewAlgoOrderTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewAlgoOrderTypeParameter: valid values are %v", v, AllowedNewAlgoOrderTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewAlgoOrderTypeParameter) IsValid() bool {
	for _, existing := range AllowedNewAlgoOrderTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newAlgoOrder_type_parameter value
func (v NewAlgoOrderTypeParameter) Ptr() *NewAlgoOrderTypeParameter {
	return &v
}

type NullableNewAlgoOrderTypeParameter struct {
	value *NewAlgoOrderTypeParameter
	isSet bool
}

func (v NullableNewAlgoOrderTypeParameter) Get() *NewAlgoOrderTypeParameter {
	return v.value
}

func (v *NullableNewAlgoOrderTypeParameter) Set(val *NewAlgoOrderTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewAlgoOrderTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewAlgoOrderTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewAlgoOrderTypeParameter(val *NewAlgoOrderTypeParameter) *NullableNewAlgoOrderTypeParameter {
	return &NullableNewAlgoOrderTypeParameter{value: val, isSet: true}
}

func (v NullableNewAlgoOrderTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewAlgoOrderTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
