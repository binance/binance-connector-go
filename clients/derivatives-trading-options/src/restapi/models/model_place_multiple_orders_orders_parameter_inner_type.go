/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"
	"fmt"
)

// PlaceMultipleOrdersOrdersParameterInnerType the model 'PlaceMultipleOrdersOrdersParameterInnerType'
type PlaceMultipleOrdersOrdersParameterInnerType string

// List of placeMultipleOrders_orders_parameter_inner_type
const (
	PlaceMultipleOrdersOrdersParameterInnerTypeLimit PlaceMultipleOrdersOrdersParameterInnerType = "LIMIT"
)

// All allowed values of PlaceMultipleOrdersOrdersParameterInnerType enum
var AllowedPlaceMultipleOrdersOrdersParameterInnerTypeEnumValues = []PlaceMultipleOrdersOrdersParameterInnerType{
	"LIMIT",
}

func (v *PlaceMultipleOrdersOrdersParameterInnerType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PlaceMultipleOrdersOrdersParameterInnerType(value)
	for _, existing := range AllowedPlaceMultipleOrdersOrdersParameterInnerTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PlaceMultipleOrdersOrdersParameterInnerType", value)
}

// NewPlaceMultipleOrdersOrdersParameterInnerTypeFromValue returns a pointer to a valid PlaceMultipleOrdersOrdersParameterInnerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlaceMultipleOrdersOrdersParameterInnerTypeFromValue(v string) (*PlaceMultipleOrdersOrdersParameterInnerType, error) {
	ev := PlaceMultipleOrdersOrdersParameterInnerType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PlaceMultipleOrdersOrdersParameterInnerType: valid values are %v", v, AllowedPlaceMultipleOrdersOrdersParameterInnerTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PlaceMultipleOrdersOrdersParameterInnerType) IsValid() bool {
	for _, existing := range AllowedPlaceMultipleOrdersOrdersParameterInnerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to placeMultipleOrders_orders_parameter_inner_type value
func (v PlaceMultipleOrdersOrdersParameterInnerType) Ptr() *PlaceMultipleOrdersOrdersParameterInnerType {
	return &v
}

type NullablePlaceMultipleOrdersOrdersParameterInnerType struct {
	value *PlaceMultipleOrdersOrdersParameterInnerType
	isSet bool
}

func (v NullablePlaceMultipleOrdersOrdersParameterInnerType) Get() *PlaceMultipleOrdersOrdersParameterInnerType {
	return v.value
}

func (v *NullablePlaceMultipleOrdersOrdersParameterInnerType) Set(val *PlaceMultipleOrdersOrdersParameterInnerType) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceMultipleOrdersOrdersParameterInnerType) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceMultipleOrdersOrdersParameterInnerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceMultipleOrdersOrdersParameterInnerType(val *PlaceMultipleOrdersOrdersParameterInnerType) *NullablePlaceMultipleOrdersOrdersParameterInnerType {
	return &NullablePlaceMultipleOrdersOrdersParameterInnerType{value: val, isSet: true}
}

func (v NullablePlaceMultipleOrdersOrdersParameterInnerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceMultipleOrdersOrdersParameterInnerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
