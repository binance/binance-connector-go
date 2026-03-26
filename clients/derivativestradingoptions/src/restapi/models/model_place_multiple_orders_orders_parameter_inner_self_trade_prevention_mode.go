/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"
	"fmt"
)

// PlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionMode the model 'PlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionMode'
type PlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionMode string

// List of placeMultipleOrders_orders_parameter_inner_selfTradePreventionMode
const (
	PlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionModeExpireTaker PlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionMode = "EXPIRE_TAKER"
	PlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionModeExpireBoth  PlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionMode = "EXPIRE_BOTH"
	PlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionModeExpireMaker PlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionMode = "EXPIRE_MAKER"
)

// All allowed values of PlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionMode enum
var AllowedPlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionModeEnumValues = []PlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionMode{
	"EXPIRE_TAKER",
	"EXPIRE_BOTH",
	"EXPIRE_MAKER",
}

func (v *PlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionMode(value)
	for _, existing := range AllowedPlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionMode", value)
}

// NewPlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionModeFromValue returns a pointer to a valid PlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionModeFromValue(v string) (*PlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionMode, error) {
	ev := PlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionMode: valid values are %v", v, AllowedPlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionMode) IsValid() bool {
	for _, existing := range AllowedPlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to placeMultipleOrders_orders_parameter_inner_selfTradePreventionMode value
func (v PlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionMode) Ptr() *PlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionMode {
	return &v
}

type NullablePlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionMode struct {
	value *PlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionMode
	isSet bool
}

func (v NullablePlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionMode) Get() *PlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionMode {
	return v.value
}

func (v *NullablePlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionMode) Set(val *PlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionMode) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionMode) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionMode(val *PlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionMode) *NullablePlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionMode {
	return &NullablePlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionMode{value: val, isSet: true}
}

func (v NullablePlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
