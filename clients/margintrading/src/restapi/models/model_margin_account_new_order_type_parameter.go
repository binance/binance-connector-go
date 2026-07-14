/*
Margin REST API

Access account information, borrow and repay assets, and trade with Binance Margin.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// MarginAccountNewOrderTypeParameter the model 'MarginAccountNewOrderTypeParameter'
type MarginAccountNewOrderTypeParameter string

// List of marginAccountNewOrder_type_parameter
const (
	MarginAccountNewOrderTypeParameterLimit           MarginAccountNewOrderTypeParameter = "LIMIT"
	MarginAccountNewOrderTypeParameterMarket          MarginAccountNewOrderTypeParameter = "MARKET"
	MarginAccountNewOrderTypeParameterStopLoss        MarginAccountNewOrderTypeParameter = "STOP_LOSS"
	MarginAccountNewOrderTypeParameterStopLossLimit   MarginAccountNewOrderTypeParameter = "STOP_LOSS_LIMIT"
	MarginAccountNewOrderTypeParameterTakeProfit      MarginAccountNewOrderTypeParameter = "TAKE_PROFIT"
	MarginAccountNewOrderTypeParameterTakeProfitLimit MarginAccountNewOrderTypeParameter = "TAKE_PROFIT_LIMIT"
	MarginAccountNewOrderTypeParameterLimitMaker      MarginAccountNewOrderTypeParameter = "LIMIT_MAKER"
)

// All allowed values of MarginAccountNewOrderTypeParameter enum
var AllowedMarginAccountNewOrderTypeParameterEnumValues = []MarginAccountNewOrderTypeParameter{
	"LIMIT",
	"MARKET",
	"STOP_LOSS",
	"STOP_LOSS_LIMIT",
	"TAKE_PROFIT",
	"TAKE_PROFIT_LIMIT",
	"LIMIT_MAKER",
}

func (v *MarginAccountNewOrderTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MarginAccountNewOrderTypeParameter(value)
	for _, existing := range AllowedMarginAccountNewOrderTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MarginAccountNewOrderTypeParameter", value)
}

// NewMarginAccountNewOrderTypeParameterFromValue returns a pointer to a valid MarginAccountNewOrderTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMarginAccountNewOrderTypeParameterFromValue(v string) (*MarginAccountNewOrderTypeParameter, error) {
	ev := MarginAccountNewOrderTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MarginAccountNewOrderTypeParameter: valid values are %v", v, AllowedMarginAccountNewOrderTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MarginAccountNewOrderTypeParameter) IsValid() bool {
	for _, existing := range AllowedMarginAccountNewOrderTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to marginAccountNewOrder_type_parameter value
func (v MarginAccountNewOrderTypeParameter) Ptr() *MarginAccountNewOrderTypeParameter {
	return &v
}

type NullableMarginAccountNewOrderTypeParameter struct {
	value *MarginAccountNewOrderTypeParameter
	isSet bool
}

func (v NullableMarginAccountNewOrderTypeParameter) Get() *MarginAccountNewOrderTypeParameter {
	return v.value
}

func (v *NullableMarginAccountNewOrderTypeParameter) Set(val *MarginAccountNewOrderTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginAccountNewOrderTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginAccountNewOrderTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginAccountNewOrderTypeParameter(val *MarginAccountNewOrderTypeParameter) *NullableMarginAccountNewOrderTypeParameter {
	return &NullableMarginAccountNewOrderTypeParameter{value: val, isSet: true}
}

func (v NullableMarginAccountNewOrderTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginAccountNewOrderTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
