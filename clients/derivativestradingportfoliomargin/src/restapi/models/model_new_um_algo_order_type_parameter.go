/*
Portfolio Margin REST API

Access account information, manage margin positions, and trade with Binance Portfolio Margin.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewUmAlgoOrderTypeParameter the model 'NewUmAlgoOrderTypeParameter'
type NewUmAlgoOrderTypeParameter string

// List of newUmAlgoOrder_type_parameter
const (
	NewUmAlgoOrderTypeParameterStop               NewUmAlgoOrderTypeParameter = "STOP"
	NewUmAlgoOrderTypeParameterTakeProfit         NewUmAlgoOrderTypeParameter = "TAKE_PROFIT"
	NewUmAlgoOrderTypeParameterStopMarket         NewUmAlgoOrderTypeParameter = "STOP_MARKET"
	NewUmAlgoOrderTypeParameterTakeProfitMarket   NewUmAlgoOrderTypeParameter = "TAKE_PROFIT_MARKET"
	NewUmAlgoOrderTypeParameterTrailingStopMarket NewUmAlgoOrderTypeParameter = "TRAILING_STOP_MARKET"
)

// All allowed values of NewUmAlgoOrderTypeParameter enum
var AllowedNewUmAlgoOrderTypeParameterEnumValues = []NewUmAlgoOrderTypeParameter{
	"STOP",
	"TAKE_PROFIT",
	"STOP_MARKET",
	"TAKE_PROFIT_MARKET",
	"TRAILING_STOP_MARKET",
}

func (v *NewUmAlgoOrderTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewUmAlgoOrderTypeParameter(value)
	for _, existing := range AllowedNewUmAlgoOrderTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewUmAlgoOrderTypeParameter", value)
}

// NewNewUmAlgoOrderTypeParameterFromValue returns a pointer to a valid NewUmAlgoOrderTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewUmAlgoOrderTypeParameterFromValue(v string) (*NewUmAlgoOrderTypeParameter, error) {
	ev := NewUmAlgoOrderTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewUmAlgoOrderTypeParameter: valid values are %v", v, AllowedNewUmAlgoOrderTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewUmAlgoOrderTypeParameter) IsValid() bool {
	for _, existing := range AllowedNewUmAlgoOrderTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newUmAlgoOrder_type_parameter value
func (v NewUmAlgoOrderTypeParameter) Ptr() *NewUmAlgoOrderTypeParameter {
	return &v
}

type NullableNewUmAlgoOrderTypeParameter struct {
	value *NewUmAlgoOrderTypeParameter
	isSet bool
}

func (v NullableNewUmAlgoOrderTypeParameter) Get() *NewUmAlgoOrderTypeParameter {
	return v.value
}

func (v *NullableNewUmAlgoOrderTypeParameter) Set(val *NewUmAlgoOrderTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewUmAlgoOrderTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewUmAlgoOrderTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewUmAlgoOrderTypeParameter(val *NewUmAlgoOrderTypeParameter) *NullableNewUmAlgoOrderTypeParameter {
	return &NullableNewUmAlgoOrderTypeParameter{value: val, isSet: true}
}

func (v NullableNewUmAlgoOrderTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewUmAlgoOrderTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
