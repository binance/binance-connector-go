/*
Portfolio Margin REST API

Access account information, manage margin positions, and trade with Binance Portfolio Margin.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewMarginOrderTypeParameter the model 'NewMarginOrderTypeParameter'
type NewMarginOrderTypeParameter string

// List of newMarginOrder_type_parameter
const (
	NewMarginOrderTypeParameterLimit           NewMarginOrderTypeParameter = "LIMIT"
	NewMarginOrderTypeParameterMarket          NewMarginOrderTypeParameter = "MARKET"
	NewMarginOrderTypeParameterStopLoss        NewMarginOrderTypeParameter = "STOP_LOSS"
	NewMarginOrderTypeParameterStopLossLimit   NewMarginOrderTypeParameter = "STOP_LOSS_LIMIT"
	NewMarginOrderTypeParameterTakeProfit      NewMarginOrderTypeParameter = "TAKE_PROFIT"
	NewMarginOrderTypeParameterTakeProfitLimit NewMarginOrderTypeParameter = "TAKE_PROFIT_LIMIT"
	NewMarginOrderTypeParameterLimitMaker      NewMarginOrderTypeParameter = "LIMIT_MAKER"
)

// All allowed values of NewMarginOrderTypeParameter enum
var AllowedNewMarginOrderTypeParameterEnumValues = []NewMarginOrderTypeParameter{
	"LIMIT",
	"MARKET",
	"STOP_LOSS",
	"STOP_LOSS_LIMIT",
	"TAKE_PROFIT",
	"TAKE_PROFIT_LIMIT",
	"LIMIT_MAKER",
}

func (v *NewMarginOrderTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewMarginOrderTypeParameter(value)
	for _, existing := range AllowedNewMarginOrderTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewMarginOrderTypeParameter", value)
}

// NewNewMarginOrderTypeParameterFromValue returns a pointer to a valid NewMarginOrderTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewMarginOrderTypeParameterFromValue(v string) (*NewMarginOrderTypeParameter, error) {
	ev := NewMarginOrderTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewMarginOrderTypeParameter: valid values are %v", v, AllowedNewMarginOrderTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewMarginOrderTypeParameter) IsValid() bool {
	for _, existing := range AllowedNewMarginOrderTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newMarginOrder_type_parameter value
func (v NewMarginOrderTypeParameter) Ptr() *NewMarginOrderTypeParameter {
	return &v
}

type NullableNewMarginOrderTypeParameter struct {
	value *NewMarginOrderTypeParameter
	isSet bool
}

func (v NullableNewMarginOrderTypeParameter) Get() *NewMarginOrderTypeParameter {
	return v.value
}

func (v *NullableNewMarginOrderTypeParameter) Set(val *NewMarginOrderTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewMarginOrderTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewMarginOrderTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewMarginOrderTypeParameter(val *NewMarginOrderTypeParameter) *NullableNewMarginOrderTypeParameter {
	return &NullableNewMarginOrderTypeParameter{value: val, isSet: true}
}

func (v NullableNewMarginOrderTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewMarginOrderTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
