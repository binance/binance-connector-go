/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"
	"fmt"
)

// PlaceMultipleOrdersBatchOrdersParameterInnerType the model 'PlaceMultipleOrdersBatchOrdersParameterInnerType'
type PlaceMultipleOrdersBatchOrdersParameterInnerType string

// List of placeMultipleOrders_batchOrders_parameter_inner_type
const (
	PlaceMultipleOrdersBatchOrdersParameterInnerTypeLimit              PlaceMultipleOrdersBatchOrdersParameterInnerType = "LIMIT"
	PlaceMultipleOrdersBatchOrdersParameterInnerTypeMarket             PlaceMultipleOrdersBatchOrdersParameterInnerType = "MARKET"
	PlaceMultipleOrdersBatchOrdersParameterInnerTypeStop               PlaceMultipleOrdersBatchOrdersParameterInnerType = "STOP"
	PlaceMultipleOrdersBatchOrdersParameterInnerTypeStopMarket         PlaceMultipleOrdersBatchOrdersParameterInnerType = "STOP_MARKET"
	PlaceMultipleOrdersBatchOrdersParameterInnerTypeTakeProfit         PlaceMultipleOrdersBatchOrdersParameterInnerType = "TAKE_PROFIT"
	PlaceMultipleOrdersBatchOrdersParameterInnerTypeTakeProfitMarket   PlaceMultipleOrdersBatchOrdersParameterInnerType = "TAKE_PROFIT_MARKET"
	PlaceMultipleOrdersBatchOrdersParameterInnerTypeTrailingStopMarket PlaceMultipleOrdersBatchOrdersParameterInnerType = "TRAILING_STOP_MARKET"
)

// All allowed values of PlaceMultipleOrdersBatchOrdersParameterInnerType enum
var AllowedPlaceMultipleOrdersBatchOrdersParameterInnerTypeEnumValues = []PlaceMultipleOrdersBatchOrdersParameterInnerType{
	"LIMIT",
	"MARKET",
	"STOP",
	"STOP_MARKET",
	"TAKE_PROFIT",
	"TAKE_PROFIT_MARKET",
	"TRAILING_STOP_MARKET",
}

func (v *PlaceMultipleOrdersBatchOrdersParameterInnerType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PlaceMultipleOrdersBatchOrdersParameterInnerType(value)
	for _, existing := range AllowedPlaceMultipleOrdersBatchOrdersParameterInnerTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PlaceMultipleOrdersBatchOrdersParameterInnerType", value)
}

// NewPlaceMultipleOrdersBatchOrdersParameterInnerTypeFromValue returns a pointer to a valid PlaceMultipleOrdersBatchOrdersParameterInnerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlaceMultipleOrdersBatchOrdersParameterInnerTypeFromValue(v string) (*PlaceMultipleOrdersBatchOrdersParameterInnerType, error) {
	ev := PlaceMultipleOrdersBatchOrdersParameterInnerType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PlaceMultipleOrdersBatchOrdersParameterInnerType: valid values are %v", v, AllowedPlaceMultipleOrdersBatchOrdersParameterInnerTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PlaceMultipleOrdersBatchOrdersParameterInnerType) IsValid() bool {
	for _, existing := range AllowedPlaceMultipleOrdersBatchOrdersParameterInnerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to placeMultipleOrders_batchOrders_parameter_inner_type value
func (v PlaceMultipleOrdersBatchOrdersParameterInnerType) Ptr() *PlaceMultipleOrdersBatchOrdersParameterInnerType {
	return &v
}

type NullablePlaceMultipleOrdersBatchOrdersParameterInnerType struct {
	value *PlaceMultipleOrdersBatchOrdersParameterInnerType
	isSet bool
}

func (v NullablePlaceMultipleOrdersBatchOrdersParameterInnerType) Get() *PlaceMultipleOrdersBatchOrdersParameterInnerType {
	return v.value
}

func (v *NullablePlaceMultipleOrdersBatchOrdersParameterInnerType) Set(val *PlaceMultipleOrdersBatchOrdersParameterInnerType) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceMultipleOrdersBatchOrdersParameterInnerType) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceMultipleOrdersBatchOrdersParameterInnerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceMultipleOrdersBatchOrdersParameterInnerType(val *PlaceMultipleOrdersBatchOrdersParameterInnerType) *NullablePlaceMultipleOrdersBatchOrdersParameterInnerType {
	return &NullablePlaceMultipleOrdersBatchOrdersParameterInnerType{value: val, isSet: true}
}

func (v NullablePlaceMultipleOrdersBatchOrdersParameterInnerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceMultipleOrdersBatchOrdersParameterInnerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
