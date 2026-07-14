/*
Spot REST API

Access market data, manage accounts, and trade on Binance Spot.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// DeleteOrderCancelRestrictionsParameter the model 'DeleteOrderCancelRestrictionsParameter'
type DeleteOrderCancelRestrictionsParameter string

// List of deleteOrder_cancelRestrictions_parameter
const (
	DeleteOrderCancelRestrictionsParameterOnlyNew             DeleteOrderCancelRestrictionsParameter = "ONLY_NEW"
	DeleteOrderCancelRestrictionsParameterOnlyPartiallyFilled DeleteOrderCancelRestrictionsParameter = "ONLY_PARTIALLY_FILLED"
)

// All allowed values of DeleteOrderCancelRestrictionsParameter enum
var AllowedDeleteOrderCancelRestrictionsParameterEnumValues = []DeleteOrderCancelRestrictionsParameter{
	"ONLY_NEW",
	"ONLY_PARTIALLY_FILLED",
}

func (v *DeleteOrderCancelRestrictionsParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeleteOrderCancelRestrictionsParameter(value)
	for _, existing := range AllowedDeleteOrderCancelRestrictionsParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeleteOrderCancelRestrictionsParameter", value)
}

// NewDeleteOrderCancelRestrictionsParameterFromValue returns a pointer to a valid DeleteOrderCancelRestrictionsParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeleteOrderCancelRestrictionsParameterFromValue(v string) (*DeleteOrderCancelRestrictionsParameter, error) {
	ev := DeleteOrderCancelRestrictionsParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeleteOrderCancelRestrictionsParameter: valid values are %v", v, AllowedDeleteOrderCancelRestrictionsParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeleteOrderCancelRestrictionsParameter) IsValid() bool {
	for _, existing := range AllowedDeleteOrderCancelRestrictionsParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to deleteOrder_cancelRestrictions_parameter value
func (v DeleteOrderCancelRestrictionsParameter) Ptr() *DeleteOrderCancelRestrictionsParameter {
	return &v
}

type NullableDeleteOrderCancelRestrictionsParameter struct {
	value *DeleteOrderCancelRestrictionsParameter
	isSet bool
}

func (v NullableDeleteOrderCancelRestrictionsParameter) Get() *DeleteOrderCancelRestrictionsParameter {
	return v.value
}

func (v *NullableDeleteOrderCancelRestrictionsParameter) Set(val *DeleteOrderCancelRestrictionsParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteOrderCancelRestrictionsParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteOrderCancelRestrictionsParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteOrderCancelRestrictionsParameter(val *DeleteOrderCancelRestrictionsParameter) *NullableDeleteOrderCancelRestrictionsParameter {
	return &NullableDeleteOrderCancelRestrictionsParameter{value: val, isSet: true}
}

func (v NullableDeleteOrderCancelRestrictionsParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteOrderCancelRestrictionsParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
