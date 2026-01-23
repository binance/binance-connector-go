/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"
	"fmt"
)

// PlaceMultipleOrdersOrdersParameterInnerTimeInForce the model 'PlaceMultipleOrdersOrdersParameterInnerTimeInForce'
type PlaceMultipleOrdersOrdersParameterInnerTimeInForce string

// List of placeMultipleOrders_orders_parameter_inner_timeInForce
const (
	PlaceMultipleOrdersOrdersParameterInnerTimeInForceGtc PlaceMultipleOrdersOrdersParameterInnerTimeInForce = "GTC"
	PlaceMultipleOrdersOrdersParameterInnerTimeInForceIoc PlaceMultipleOrdersOrdersParameterInnerTimeInForce = "IOC"
	PlaceMultipleOrdersOrdersParameterInnerTimeInForceFok PlaceMultipleOrdersOrdersParameterInnerTimeInForce = "FOK"
	PlaceMultipleOrdersOrdersParameterInnerTimeInForceGtx PlaceMultipleOrdersOrdersParameterInnerTimeInForce = "GTX"
)

// All allowed values of PlaceMultipleOrdersOrdersParameterInnerTimeInForce enum
var AllowedPlaceMultipleOrdersOrdersParameterInnerTimeInForceEnumValues = []PlaceMultipleOrdersOrdersParameterInnerTimeInForce{
	"GTC",
	"IOC",
	"FOK",
	"GTX",
}

func (v *PlaceMultipleOrdersOrdersParameterInnerTimeInForce) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PlaceMultipleOrdersOrdersParameterInnerTimeInForce(value)
	for _, existing := range AllowedPlaceMultipleOrdersOrdersParameterInnerTimeInForceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PlaceMultipleOrdersOrdersParameterInnerTimeInForce", value)
}

// NewPlaceMultipleOrdersOrdersParameterInnerTimeInForceFromValue returns a pointer to a valid PlaceMultipleOrdersOrdersParameterInnerTimeInForce
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlaceMultipleOrdersOrdersParameterInnerTimeInForceFromValue(v string) (*PlaceMultipleOrdersOrdersParameterInnerTimeInForce, error) {
	ev := PlaceMultipleOrdersOrdersParameterInnerTimeInForce(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PlaceMultipleOrdersOrdersParameterInnerTimeInForce: valid values are %v", v, AllowedPlaceMultipleOrdersOrdersParameterInnerTimeInForceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PlaceMultipleOrdersOrdersParameterInnerTimeInForce) IsValid() bool {
	for _, existing := range AllowedPlaceMultipleOrdersOrdersParameterInnerTimeInForceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to placeMultipleOrders_orders_parameter_inner_timeInForce value
func (v PlaceMultipleOrdersOrdersParameterInnerTimeInForce) Ptr() *PlaceMultipleOrdersOrdersParameterInnerTimeInForce {
	return &v
}

type NullablePlaceMultipleOrdersOrdersParameterInnerTimeInForce struct {
	value *PlaceMultipleOrdersOrdersParameterInnerTimeInForce
	isSet bool
}

func (v NullablePlaceMultipleOrdersOrdersParameterInnerTimeInForce) Get() *PlaceMultipleOrdersOrdersParameterInnerTimeInForce {
	return v.value
}

func (v *NullablePlaceMultipleOrdersOrdersParameterInnerTimeInForce) Set(val *PlaceMultipleOrdersOrdersParameterInnerTimeInForce) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceMultipleOrdersOrdersParameterInnerTimeInForce) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceMultipleOrdersOrdersParameterInnerTimeInForce) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceMultipleOrdersOrdersParameterInnerTimeInForce(val *PlaceMultipleOrdersOrdersParameterInnerTimeInForce) *NullablePlaceMultipleOrdersOrdersParameterInnerTimeInForce {
	return &NullablePlaceMultipleOrdersOrdersParameterInnerTimeInForce{value: val, isSet: true}
}

func (v NullablePlaceMultipleOrdersOrdersParameterInnerTimeInForce) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceMultipleOrdersOrdersParameterInnerTimeInForce) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
