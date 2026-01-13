/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"
	"fmt"
)

// PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode the model 'PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode'
type PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode string

// List of placeMultipleOrders_batchOrders_parameter_inner_selfTradePreventionMode
const (
	PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionModeNone        PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode = "NONE"
	PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionModeExpireTaker PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode = "EXPIRE_TAKER"
	PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionModeExpireBoth  PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode = "EXPIRE_BOTH"
	PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionModeExpireMaker PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode = "EXPIRE_MAKER"
)

// All allowed values of PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode enum
var AllowedPlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionModeEnumValues = []PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode{
	"NONE",
	"EXPIRE_TAKER",
	"EXPIRE_BOTH",
	"EXPIRE_MAKER",
}

func (v *PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode(value)
	for _, existing := range AllowedPlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode", value)
}

// NewPlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionModeFromValue returns a pointer to a valid PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionModeFromValue(v string) (*PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode, error) {
	ev := PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode: valid values are %v", v, AllowedPlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode) IsValid() bool {
	for _, existing := range AllowedPlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to placeMultipleOrders_batchOrders_parameter_inner_selfTradePreventionMode value
func (v PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode) Ptr() *PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode {
	return &v
}

type NullablePlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode struct {
	value *PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode
	isSet bool
}

func (v NullablePlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode) Get() *PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode {
	return v.value
}

func (v *NullablePlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode) Set(val *PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode(val *PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode) *NullablePlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode {
	return &NullablePlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode{value: val, isSet: true}
}

func (v NullablePlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
