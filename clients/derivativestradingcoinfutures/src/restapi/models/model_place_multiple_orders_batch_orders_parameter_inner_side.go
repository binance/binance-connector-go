/*
Futures (COIN-M) REST API

Access market data, manage accounts, and trade COIN-M perpetual and delivery futures.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// PlaceMultipleOrdersBatchOrdersParameterInnerSide the model 'PlaceMultipleOrdersBatchOrdersParameterInnerSide'
type PlaceMultipleOrdersBatchOrdersParameterInnerSide string

// List of placeMultipleOrders_batchOrders_parameter_inner_side
const (
	PlaceMultipleOrdersBatchOrdersParameterInnerSideBuy  PlaceMultipleOrdersBatchOrdersParameterInnerSide = "BUY"
	PlaceMultipleOrdersBatchOrdersParameterInnerSideSell PlaceMultipleOrdersBatchOrdersParameterInnerSide = "SELL"
)

// All allowed values of PlaceMultipleOrdersBatchOrdersParameterInnerSide enum
var AllowedPlaceMultipleOrdersBatchOrdersParameterInnerSideEnumValues = []PlaceMultipleOrdersBatchOrdersParameterInnerSide{
	"BUY",
	"SELL",
}

func (v *PlaceMultipleOrdersBatchOrdersParameterInnerSide) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PlaceMultipleOrdersBatchOrdersParameterInnerSide(value)
	for _, existing := range AllowedPlaceMultipleOrdersBatchOrdersParameterInnerSideEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PlaceMultipleOrdersBatchOrdersParameterInnerSide", value)
}

// NewPlaceMultipleOrdersBatchOrdersParameterInnerSideFromValue returns a pointer to a valid PlaceMultipleOrdersBatchOrdersParameterInnerSide
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlaceMultipleOrdersBatchOrdersParameterInnerSideFromValue(v string) (*PlaceMultipleOrdersBatchOrdersParameterInnerSide, error) {
	ev := PlaceMultipleOrdersBatchOrdersParameterInnerSide(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PlaceMultipleOrdersBatchOrdersParameterInnerSide: valid values are %v", v, AllowedPlaceMultipleOrdersBatchOrdersParameterInnerSideEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PlaceMultipleOrdersBatchOrdersParameterInnerSide) IsValid() bool {
	for _, existing := range AllowedPlaceMultipleOrdersBatchOrdersParameterInnerSideEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to placeMultipleOrders_batchOrders_parameter_inner_side value
func (v PlaceMultipleOrdersBatchOrdersParameterInnerSide) Ptr() *PlaceMultipleOrdersBatchOrdersParameterInnerSide {
	return &v
}

type NullablePlaceMultipleOrdersBatchOrdersParameterInnerSide struct {
	value *PlaceMultipleOrdersBatchOrdersParameterInnerSide
	isSet bool
}

func (v NullablePlaceMultipleOrdersBatchOrdersParameterInnerSide) Get() *PlaceMultipleOrdersBatchOrdersParameterInnerSide {
	return v.value
}

func (v *NullablePlaceMultipleOrdersBatchOrdersParameterInnerSide) Set(val *PlaceMultipleOrdersBatchOrdersParameterInnerSide) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceMultipleOrdersBatchOrdersParameterInnerSide) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceMultipleOrdersBatchOrdersParameterInnerSide) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceMultipleOrdersBatchOrdersParameterInnerSide(val *PlaceMultipleOrdersBatchOrdersParameterInnerSide) *NullablePlaceMultipleOrdersBatchOrdersParameterInnerSide {
	return &NullablePlaceMultipleOrdersBatchOrdersParameterInnerSide{value: val, isSet: true}
}

func (v NullablePlaceMultipleOrdersBatchOrdersParameterInnerSide) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceMultipleOrdersBatchOrdersParameterInnerSide) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
