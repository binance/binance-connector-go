/*
Futures (COIN-M) REST API

Access market data, manage accounts, and trade COIN-M perpetual and delivery futures.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// PlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect Used with `STOP/STOP_MARKET` or `TAKE_PROFIT/TAKE_PROFIT_MARKET` orders.
type PlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect string

// List of placeMultipleOrders_batchOrders_parameter_inner_priceProtect
const (
	PlaceMultipleOrdersBatchOrdersParameterInnerPriceProtectTrue  PlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect = "true"
	PlaceMultipleOrdersBatchOrdersParameterInnerPriceProtectFalse PlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect = "false"
)

// All allowed values of PlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect enum
var AllowedPlaceMultipleOrdersBatchOrdersParameterInnerPriceProtectEnumValues = []PlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect{
	"true",
	"false",
}

func (v *PlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect(value)
	for _, existing := range AllowedPlaceMultipleOrdersBatchOrdersParameterInnerPriceProtectEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect", value)
}

// NewPlaceMultipleOrdersBatchOrdersParameterInnerPriceProtectFromValue returns a pointer to a valid PlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlaceMultipleOrdersBatchOrdersParameterInnerPriceProtectFromValue(v string) (*PlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect, error) {
	ev := PlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect: valid values are %v", v, AllowedPlaceMultipleOrdersBatchOrdersParameterInnerPriceProtectEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect) IsValid() bool {
	for _, existing := range AllowedPlaceMultipleOrdersBatchOrdersParameterInnerPriceProtectEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to placeMultipleOrders_batchOrders_parameter_inner_priceProtect value
func (v PlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect) Ptr() *PlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect {
	return &v
}

type NullablePlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect struct {
	value *PlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect
	isSet bool
}

func (v NullablePlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect) Get() *PlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect {
	return v.value
}

func (v *NullablePlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect) Set(val *PlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect(val *PlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect) *NullablePlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect {
	return &NullablePlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect{value: val, isSet: true}
}

func (v NullablePlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
