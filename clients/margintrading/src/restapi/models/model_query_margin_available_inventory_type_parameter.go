/*
Margin REST API

Access account information, borrow and repay assets, and trade with Binance Margin.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// QueryMarginAvailableInventoryTypeParameter the model 'QueryMarginAvailableInventoryTypeParameter'
type QueryMarginAvailableInventoryTypeParameter string

// List of queryMarginAvailableInventory_type_parameter
const (
	QueryMarginAvailableInventoryTypeParameterMargin   QueryMarginAvailableInventoryTypeParameter = "MARGIN"
	QueryMarginAvailableInventoryTypeParameterIsolated QueryMarginAvailableInventoryTypeParameter = "ISOLATED"
)

// All allowed values of QueryMarginAvailableInventoryTypeParameter enum
var AllowedQueryMarginAvailableInventoryTypeParameterEnumValues = []QueryMarginAvailableInventoryTypeParameter{
	"MARGIN",
	"ISOLATED",
}

func (v *QueryMarginAvailableInventoryTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := QueryMarginAvailableInventoryTypeParameter(value)
	for _, existing := range AllowedQueryMarginAvailableInventoryTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid QueryMarginAvailableInventoryTypeParameter", value)
}

// NewQueryMarginAvailableInventoryTypeParameterFromValue returns a pointer to a valid QueryMarginAvailableInventoryTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQueryMarginAvailableInventoryTypeParameterFromValue(v string) (*QueryMarginAvailableInventoryTypeParameter, error) {
	ev := QueryMarginAvailableInventoryTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QueryMarginAvailableInventoryTypeParameter: valid values are %v", v, AllowedQueryMarginAvailableInventoryTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QueryMarginAvailableInventoryTypeParameter) IsValid() bool {
	for _, existing := range AllowedQueryMarginAvailableInventoryTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to queryMarginAvailableInventory_type_parameter value
func (v QueryMarginAvailableInventoryTypeParameter) Ptr() *QueryMarginAvailableInventoryTypeParameter {
	return &v
}

type NullableQueryMarginAvailableInventoryTypeParameter struct {
	value *QueryMarginAvailableInventoryTypeParameter
	isSet bool
}

func (v NullableQueryMarginAvailableInventoryTypeParameter) Get() *QueryMarginAvailableInventoryTypeParameter {
	return v.value
}

func (v *NullableQueryMarginAvailableInventoryTypeParameter) Set(val *QueryMarginAvailableInventoryTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryMarginAvailableInventoryTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryMarginAvailableInventoryTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryMarginAvailableInventoryTypeParameter(val *QueryMarginAvailableInventoryTypeParameter) *NullableQueryMarginAvailableInventoryTypeParameter {
	return &NullableQueryMarginAvailableInventoryTypeParameter{value: val, isSet: true}
}

func (v NullableQueryMarginAvailableInventoryTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryMarginAvailableInventoryTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
