/*
Spot WebSocket API

Access market data, manage accounts, and trade on Binance Spot.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// OrderCancelCancelRestrictionsParameter the model 'OrderCancelCancelRestrictionsParameter'
type OrderCancelCancelRestrictionsParameter string

// List of orderCancel_cancelRestrictions_parameter
const (
	OrderCancelCancelRestrictionsParameterOnlyNew             OrderCancelCancelRestrictionsParameter = "ONLY_NEW"
	OrderCancelCancelRestrictionsParameterOnlyPartiallyFilled OrderCancelCancelRestrictionsParameter = "ONLY_PARTIALLY_FILLED"
)

// All allowed values of OrderCancelCancelRestrictionsParameter enum
var AllowedOrderCancelCancelRestrictionsParameterEnumValues = []OrderCancelCancelRestrictionsParameter{
	"ONLY_NEW",
	"ONLY_PARTIALLY_FILLED",
}

func (v *OrderCancelCancelRestrictionsParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderCancelCancelRestrictionsParameter(value)
	for _, existing := range AllowedOrderCancelCancelRestrictionsParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderCancelCancelRestrictionsParameter", value)
}

// NewOrderCancelCancelRestrictionsParameterFromValue returns a pointer to a valid OrderCancelCancelRestrictionsParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderCancelCancelRestrictionsParameterFromValue(v string) (*OrderCancelCancelRestrictionsParameter, error) {
	ev := OrderCancelCancelRestrictionsParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderCancelCancelRestrictionsParameter: valid values are %v", v, AllowedOrderCancelCancelRestrictionsParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderCancelCancelRestrictionsParameter) IsValid() bool {
	for _, existing := range AllowedOrderCancelCancelRestrictionsParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to orderCancel_cancelRestrictions_parameter value
func (v OrderCancelCancelRestrictionsParameter) Ptr() *OrderCancelCancelRestrictionsParameter {
	return &v
}

type NullableOrderCancelCancelRestrictionsParameter struct {
	value *OrderCancelCancelRestrictionsParameter
	isSet bool
}

func (v NullableOrderCancelCancelRestrictionsParameter) Get() *OrderCancelCancelRestrictionsParameter {
	return v.value
}

func (v *NullableOrderCancelCancelRestrictionsParameter) Set(val *OrderCancelCancelRestrictionsParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCancelCancelRestrictionsParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCancelCancelRestrictionsParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCancelCancelRestrictionsParameter(val *OrderCancelCancelRestrictionsParameter) *NullableOrderCancelCancelRestrictionsParameter {
	return &NullableOrderCancelCancelRestrictionsParameter{value: val, isSet: true}
}

func (v NullableOrderCancelCancelRestrictionsParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCancelCancelRestrictionsParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
