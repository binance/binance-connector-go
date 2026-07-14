/*
Options REST API

Access market data, manage accounts, and trade Binance Options.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewBlockTradeOrderLiquidityParameter the model 'NewBlockTradeOrderLiquidityParameter'
type NewBlockTradeOrderLiquidityParameter string

// List of newBlockTradeOrder_liquidity_parameter
const (
	NewBlockTradeOrderLiquidityParameterMaker NewBlockTradeOrderLiquidityParameter = "MAKER"
	NewBlockTradeOrderLiquidityParameterTaker NewBlockTradeOrderLiquidityParameter = "TAKER"
)

// All allowed values of NewBlockTradeOrderLiquidityParameter enum
var AllowedNewBlockTradeOrderLiquidityParameterEnumValues = []NewBlockTradeOrderLiquidityParameter{
	"MAKER",
	"TAKER",
}

func (v *NewBlockTradeOrderLiquidityParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewBlockTradeOrderLiquidityParameter(value)
	for _, existing := range AllowedNewBlockTradeOrderLiquidityParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewBlockTradeOrderLiquidityParameter", value)
}

// NewNewBlockTradeOrderLiquidityParameterFromValue returns a pointer to a valid NewBlockTradeOrderLiquidityParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewBlockTradeOrderLiquidityParameterFromValue(v string) (*NewBlockTradeOrderLiquidityParameter, error) {
	ev := NewBlockTradeOrderLiquidityParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewBlockTradeOrderLiquidityParameter: valid values are %v", v, AllowedNewBlockTradeOrderLiquidityParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewBlockTradeOrderLiquidityParameter) IsValid() bool {
	for _, existing := range AllowedNewBlockTradeOrderLiquidityParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newBlockTradeOrder_liquidity_parameter value
func (v NewBlockTradeOrderLiquidityParameter) Ptr() *NewBlockTradeOrderLiquidityParameter {
	return &v
}

type NullableNewBlockTradeOrderLiquidityParameter struct {
	value *NewBlockTradeOrderLiquidityParameter
	isSet bool
}

func (v NullableNewBlockTradeOrderLiquidityParameter) Get() *NewBlockTradeOrderLiquidityParameter {
	return v.value
}

func (v *NullableNewBlockTradeOrderLiquidityParameter) Set(val *NewBlockTradeOrderLiquidityParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewBlockTradeOrderLiquidityParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewBlockTradeOrderLiquidityParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewBlockTradeOrderLiquidityParameter(val *NewBlockTradeOrderLiquidityParameter) *NullableNewBlockTradeOrderLiquidityParameter {
	return &NullableNewBlockTradeOrderLiquidityParameter{value: val, isSet: true}
}

func (v NullableNewBlockTradeOrderLiquidityParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewBlockTradeOrderLiquidityParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
