/*
Options REST API

Access market data, manage accounts, and trade Binance Options.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewBlockTradeOrderLegsParameterInnerSide Buy/sell direction
type NewBlockTradeOrderLegsParameterInnerSide string

// List of newBlockTradeOrder_legs_parameter_inner_side
const (
	NewBlockTradeOrderLegsParameterInnerSideBuy  NewBlockTradeOrderLegsParameterInnerSide = "BUY"
	NewBlockTradeOrderLegsParameterInnerSideSell NewBlockTradeOrderLegsParameterInnerSide = "SELL"
)

// All allowed values of NewBlockTradeOrderLegsParameterInnerSide enum
var AllowedNewBlockTradeOrderLegsParameterInnerSideEnumValues = []NewBlockTradeOrderLegsParameterInnerSide{
	"BUY",
	"SELL",
}

func (v *NewBlockTradeOrderLegsParameterInnerSide) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewBlockTradeOrderLegsParameterInnerSide(value)
	for _, existing := range AllowedNewBlockTradeOrderLegsParameterInnerSideEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewBlockTradeOrderLegsParameterInnerSide", value)
}

// NewNewBlockTradeOrderLegsParameterInnerSideFromValue returns a pointer to a valid NewBlockTradeOrderLegsParameterInnerSide
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewBlockTradeOrderLegsParameterInnerSideFromValue(v string) (*NewBlockTradeOrderLegsParameterInnerSide, error) {
	ev := NewBlockTradeOrderLegsParameterInnerSide(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewBlockTradeOrderLegsParameterInnerSide: valid values are %v", v, AllowedNewBlockTradeOrderLegsParameterInnerSideEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewBlockTradeOrderLegsParameterInnerSide) IsValid() bool {
	for _, existing := range AllowedNewBlockTradeOrderLegsParameterInnerSideEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newBlockTradeOrder_legs_parameter_inner_side value
func (v NewBlockTradeOrderLegsParameterInnerSide) Ptr() *NewBlockTradeOrderLegsParameterInnerSide {
	return &v
}

type NullableNewBlockTradeOrderLegsParameterInnerSide struct {
	value *NewBlockTradeOrderLegsParameterInnerSide
	isSet bool
}

func (v NullableNewBlockTradeOrderLegsParameterInnerSide) Get() *NewBlockTradeOrderLegsParameterInnerSide {
	return v.value
}

func (v *NullableNewBlockTradeOrderLegsParameterInnerSide) Set(val *NewBlockTradeOrderLegsParameterInnerSide) {
	v.value = val
	v.isSet = true
}

func (v NullableNewBlockTradeOrderLegsParameterInnerSide) IsSet() bool {
	return v.isSet
}

func (v *NullableNewBlockTradeOrderLegsParameterInnerSide) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewBlockTradeOrderLegsParameterInnerSide(val *NewBlockTradeOrderLegsParameterInnerSide) *NullableNewBlockTradeOrderLegsParameterInnerSide {
	return &NullableNewBlockTradeOrderLegsParameterInnerSide{value: val, isSet: true}
}

func (v NullableNewBlockTradeOrderLegsParameterInnerSide) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewBlockTradeOrderLegsParameterInnerSide) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
