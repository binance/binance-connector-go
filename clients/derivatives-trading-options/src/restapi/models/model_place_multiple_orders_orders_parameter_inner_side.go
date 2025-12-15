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

// PlaceMultipleOrdersOrdersParameterInnerSide the model 'PlaceMultipleOrdersOrdersParameterInnerSide'
type PlaceMultipleOrdersOrdersParameterInnerSide string

// List of placeMultipleOrders_orders_parameter_inner_side
const (
	PlaceMultipleOrdersOrdersParameterInnerSideBuy  PlaceMultipleOrdersOrdersParameterInnerSide = "BUY"
	PlaceMultipleOrdersOrdersParameterInnerSideSell PlaceMultipleOrdersOrdersParameterInnerSide = "SELL"
)

// All allowed values of PlaceMultipleOrdersOrdersParameterInnerSide enum
var AllowedPlaceMultipleOrdersOrdersParameterInnerSideEnumValues = []PlaceMultipleOrdersOrdersParameterInnerSide{
	"BUY",
	"SELL",
}

func (v *PlaceMultipleOrdersOrdersParameterInnerSide) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PlaceMultipleOrdersOrdersParameterInnerSide(value)
	for _, existing := range AllowedPlaceMultipleOrdersOrdersParameterInnerSideEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PlaceMultipleOrdersOrdersParameterInnerSide", value)
}

// NewPlaceMultipleOrdersOrdersParameterInnerSideFromValue returns a pointer to a valid PlaceMultipleOrdersOrdersParameterInnerSide
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlaceMultipleOrdersOrdersParameterInnerSideFromValue(v string) (*PlaceMultipleOrdersOrdersParameterInnerSide, error) {
	ev := PlaceMultipleOrdersOrdersParameterInnerSide(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PlaceMultipleOrdersOrdersParameterInnerSide: valid values are %v", v, AllowedPlaceMultipleOrdersOrdersParameterInnerSideEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PlaceMultipleOrdersOrdersParameterInnerSide) IsValid() bool {
	for _, existing := range AllowedPlaceMultipleOrdersOrdersParameterInnerSideEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to placeMultipleOrders_orders_parameter_inner_side value
func (v PlaceMultipleOrdersOrdersParameterInnerSide) Ptr() *PlaceMultipleOrdersOrdersParameterInnerSide {
	return &v
}

type NullablePlaceMultipleOrdersOrdersParameterInnerSide struct {
	value *PlaceMultipleOrdersOrdersParameterInnerSide
	isSet bool
}

func (v NullablePlaceMultipleOrdersOrdersParameterInnerSide) Get() *PlaceMultipleOrdersOrdersParameterInnerSide {
	return v.value
}

func (v *NullablePlaceMultipleOrdersOrdersParameterInnerSide) Set(val *PlaceMultipleOrdersOrdersParameterInnerSide) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceMultipleOrdersOrdersParameterInnerSide) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceMultipleOrdersOrdersParameterInnerSide) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceMultipleOrdersOrdersParameterInnerSide(val *PlaceMultipleOrdersOrdersParameterInnerSide) *NullablePlaceMultipleOrdersOrdersParameterInnerSide {
	return &NullablePlaceMultipleOrdersOrdersParameterInnerSide{value: val, isSet: true}
}

func (v NullablePlaceMultipleOrdersOrdersParameterInnerSide) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceMultipleOrdersOrdersParameterInnerSide) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
