/*
Futures (USDⓈ-M) REST API

Access market data, manage accounts, and trade USDⓈ-M perpetual futures.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// ModifyMultipleOrdersBatchOrdersParameterInnerPriceMatch Only avaliable for LIMIT/STOP/TAKE_PROFIT order; Cannot be sent together with `price`.
type ModifyMultipleOrdersBatchOrdersParameterInnerPriceMatch string

// List of modifyMultipleOrders_batchOrders_parameter_inner_priceMatch
const (
	ModifyMultipleOrdersBatchOrdersParameterInnerPriceMatchOpponent   ModifyMultipleOrdersBatchOrdersParameterInnerPriceMatch = "OPPONENT"
	ModifyMultipleOrdersBatchOrdersParameterInnerPriceMatchOpponent5  ModifyMultipleOrdersBatchOrdersParameterInnerPriceMatch = "OPPONENT_5"
	ModifyMultipleOrdersBatchOrdersParameterInnerPriceMatchOpponent10 ModifyMultipleOrdersBatchOrdersParameterInnerPriceMatch = "OPPONENT_10"
	ModifyMultipleOrdersBatchOrdersParameterInnerPriceMatchOpponent20 ModifyMultipleOrdersBatchOrdersParameterInnerPriceMatch = "OPPONENT_20"
	ModifyMultipleOrdersBatchOrdersParameterInnerPriceMatchQueue      ModifyMultipleOrdersBatchOrdersParameterInnerPriceMatch = "QUEUE"
	ModifyMultipleOrdersBatchOrdersParameterInnerPriceMatchQueue5     ModifyMultipleOrdersBatchOrdersParameterInnerPriceMatch = "QUEUE_5"
	ModifyMultipleOrdersBatchOrdersParameterInnerPriceMatchQueue10    ModifyMultipleOrdersBatchOrdersParameterInnerPriceMatch = "QUEUE_10"
	ModifyMultipleOrdersBatchOrdersParameterInnerPriceMatchQueue20    ModifyMultipleOrdersBatchOrdersParameterInnerPriceMatch = "QUEUE_20"
)

// All allowed values of ModifyMultipleOrdersBatchOrdersParameterInnerPriceMatch enum
var AllowedModifyMultipleOrdersBatchOrdersParameterInnerPriceMatchEnumValues = []ModifyMultipleOrdersBatchOrdersParameterInnerPriceMatch{
	"OPPONENT",
	"OPPONENT_5",
	"OPPONENT_10",
	"OPPONENT_20",
	"QUEUE",
	"QUEUE_5",
	"QUEUE_10",
	"QUEUE_20",
}

func (v *ModifyMultipleOrdersBatchOrdersParameterInnerPriceMatch) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ModifyMultipleOrdersBatchOrdersParameterInnerPriceMatch(value)
	for _, existing := range AllowedModifyMultipleOrdersBatchOrdersParameterInnerPriceMatchEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ModifyMultipleOrdersBatchOrdersParameterInnerPriceMatch", value)
}

// NewModifyMultipleOrdersBatchOrdersParameterInnerPriceMatchFromValue returns a pointer to a valid ModifyMultipleOrdersBatchOrdersParameterInnerPriceMatch
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModifyMultipleOrdersBatchOrdersParameterInnerPriceMatchFromValue(v string) (*ModifyMultipleOrdersBatchOrdersParameterInnerPriceMatch, error) {
	ev := ModifyMultipleOrdersBatchOrdersParameterInnerPriceMatch(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ModifyMultipleOrdersBatchOrdersParameterInnerPriceMatch: valid values are %v", v, AllowedModifyMultipleOrdersBatchOrdersParameterInnerPriceMatchEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ModifyMultipleOrdersBatchOrdersParameterInnerPriceMatch) IsValid() bool {
	for _, existing := range AllowedModifyMultipleOrdersBatchOrdersParameterInnerPriceMatchEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to modifyMultipleOrders_batchOrders_parameter_inner_priceMatch value
func (v ModifyMultipleOrdersBatchOrdersParameterInnerPriceMatch) Ptr() *ModifyMultipleOrdersBatchOrdersParameterInnerPriceMatch {
	return &v
}

type NullableModifyMultipleOrdersBatchOrdersParameterInnerPriceMatch struct {
	value *ModifyMultipleOrdersBatchOrdersParameterInnerPriceMatch
	isSet bool
}

func (v NullableModifyMultipleOrdersBatchOrdersParameterInnerPriceMatch) Get() *ModifyMultipleOrdersBatchOrdersParameterInnerPriceMatch {
	return v.value
}

func (v *NullableModifyMultipleOrdersBatchOrdersParameterInnerPriceMatch) Set(val *ModifyMultipleOrdersBatchOrdersParameterInnerPriceMatch) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyMultipleOrdersBatchOrdersParameterInnerPriceMatch) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyMultipleOrdersBatchOrdersParameterInnerPriceMatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyMultipleOrdersBatchOrdersParameterInnerPriceMatch(val *ModifyMultipleOrdersBatchOrdersParameterInnerPriceMatch) *NullableModifyMultipleOrdersBatchOrdersParameterInnerPriceMatch {
	return &NullableModifyMultipleOrdersBatchOrdersParameterInnerPriceMatch{value: val, isSet: true}
}

func (v NullableModifyMultipleOrdersBatchOrdersParameterInnerPriceMatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyMultipleOrdersBatchOrdersParameterInnerPriceMatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
