/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"
	"fmt"
)

// PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide the model 'PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide'
type PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide string

// List of placeMultipleOrders_batchOrders_parameter_inner_positionSide
const (
	PlaceMultipleOrdersBatchOrdersParameterInnerPositionSideBoth  PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide = "BOTH"
	PlaceMultipleOrdersBatchOrdersParameterInnerPositionSideLong  PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide = "LONG"
	PlaceMultipleOrdersBatchOrdersParameterInnerPositionSideShort PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide = "SHORT"
)

// All allowed values of PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide enum
var AllowedPlaceMultipleOrdersBatchOrdersParameterInnerPositionSideEnumValues = []PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide{
	"BOTH",
	"LONG",
	"SHORT",
}

func (v *PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide(value)
	for _, existing := range AllowedPlaceMultipleOrdersBatchOrdersParameterInnerPositionSideEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide", value)
}

// NewPlaceMultipleOrdersBatchOrdersParameterInnerPositionSideFromValue returns a pointer to a valid PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlaceMultipleOrdersBatchOrdersParameterInnerPositionSideFromValue(v string) (*PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide, error) {
	ev := PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide: valid values are %v", v, AllowedPlaceMultipleOrdersBatchOrdersParameterInnerPositionSideEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide) IsValid() bool {
	for _, existing := range AllowedPlaceMultipleOrdersBatchOrdersParameterInnerPositionSideEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to placeMultipleOrders_batchOrders_parameter_inner_positionSide value
func (v PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide) Ptr() *PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide {
	return &v
}

type NullablePlaceMultipleOrdersBatchOrdersParameterInnerPositionSide struct {
	value *PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide
	isSet bool
}

func (v NullablePlaceMultipleOrdersBatchOrdersParameterInnerPositionSide) Get() *PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide {
	return v.value
}

func (v *NullablePlaceMultipleOrdersBatchOrdersParameterInnerPositionSide) Set(val *PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceMultipleOrdersBatchOrdersParameterInnerPositionSide) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceMultipleOrdersBatchOrdersParameterInnerPositionSide) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceMultipleOrdersBatchOrdersParameterInnerPositionSide(val *PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide) *NullablePlaceMultipleOrdersBatchOrdersParameterInnerPositionSide {
	return &NullablePlaceMultipleOrdersBatchOrdersParameterInnerPositionSide{value: val, isSet: true}
}

func (v NullablePlaceMultipleOrdersBatchOrdersParameterInnerPositionSide) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceMultipleOrdersBatchOrdersParameterInnerPositionSide) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
