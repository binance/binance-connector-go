/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"
	"fmt"
)

// ModifyMultipleOrdersBatchOrdersParameterInnerSide the model 'ModifyMultipleOrdersBatchOrdersParameterInnerSide'
type ModifyMultipleOrdersBatchOrdersParameterInnerSide string

// List of modifyMultipleOrders_batchOrders_parameter_inner_side
const (
	ModifyMultipleOrdersBatchOrdersParameterInnerSideBuy  ModifyMultipleOrdersBatchOrdersParameterInnerSide = "BUY"
	ModifyMultipleOrdersBatchOrdersParameterInnerSideSell ModifyMultipleOrdersBatchOrdersParameterInnerSide = "SELL"
)

// All allowed values of ModifyMultipleOrdersBatchOrdersParameterInnerSide enum
var AllowedModifyMultipleOrdersBatchOrdersParameterInnerSideEnumValues = []ModifyMultipleOrdersBatchOrdersParameterInnerSide{
	"BUY",
	"SELL",
}

func (v *ModifyMultipleOrdersBatchOrdersParameterInnerSide) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ModifyMultipleOrdersBatchOrdersParameterInnerSide(value)
	for _, existing := range AllowedModifyMultipleOrdersBatchOrdersParameterInnerSideEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ModifyMultipleOrdersBatchOrdersParameterInnerSide", value)
}

// NewModifyMultipleOrdersBatchOrdersParameterInnerSideFromValue returns a pointer to a valid ModifyMultipleOrdersBatchOrdersParameterInnerSide
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModifyMultipleOrdersBatchOrdersParameterInnerSideFromValue(v string) (*ModifyMultipleOrdersBatchOrdersParameterInnerSide, error) {
	ev := ModifyMultipleOrdersBatchOrdersParameterInnerSide(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ModifyMultipleOrdersBatchOrdersParameterInnerSide: valid values are %v", v, AllowedModifyMultipleOrdersBatchOrdersParameterInnerSideEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ModifyMultipleOrdersBatchOrdersParameterInnerSide) IsValid() bool {
	for _, existing := range AllowedModifyMultipleOrdersBatchOrdersParameterInnerSideEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to modifyMultipleOrders_batchOrders_parameter_inner_side value
func (v ModifyMultipleOrdersBatchOrdersParameterInnerSide) Ptr() *ModifyMultipleOrdersBatchOrdersParameterInnerSide {
	return &v
}

type NullableModifyMultipleOrdersBatchOrdersParameterInnerSide struct {
	value *ModifyMultipleOrdersBatchOrdersParameterInnerSide
	isSet bool
}

func (v NullableModifyMultipleOrdersBatchOrdersParameterInnerSide) Get() *ModifyMultipleOrdersBatchOrdersParameterInnerSide {
	return v.value
}

func (v *NullableModifyMultipleOrdersBatchOrdersParameterInnerSide) Set(val *ModifyMultipleOrdersBatchOrdersParameterInnerSide) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyMultipleOrdersBatchOrdersParameterInnerSide) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyMultipleOrdersBatchOrdersParameterInnerSide) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyMultipleOrdersBatchOrdersParameterInnerSide(val *ModifyMultipleOrdersBatchOrdersParameterInnerSide) *NullableModifyMultipleOrdersBatchOrdersParameterInnerSide {
	return &NullableModifyMultipleOrdersBatchOrdersParameterInnerSide{value: val, isSet: true}
}

func (v NullableModifyMultipleOrdersBatchOrdersParameterInnerSide) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyMultipleOrdersBatchOrdersParameterInnerSide) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
