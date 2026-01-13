/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"
	"fmt"
)

// PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce the model 'PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce'
type PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce string

// List of placeMultipleOrders_batchOrders_parameter_inner_timeInForce
const (
	PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForceGtc PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce = "GTC"
	PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForceIoc PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce = "IOC"
	PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForceFok PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce = "FOK"
	PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForceGtx PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce = "GTX"
)

// All allowed values of PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce enum
var AllowedPlaceMultipleOrdersBatchOrdersParameterInnerTimeInForceEnumValues = []PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce{
	"GTC",
	"IOC",
	"FOK",
	"GTX",
}

func (v *PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce(value)
	for _, existing := range AllowedPlaceMultipleOrdersBatchOrdersParameterInnerTimeInForceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce", value)
}

// NewPlaceMultipleOrdersBatchOrdersParameterInnerTimeInForceFromValue returns a pointer to a valid PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlaceMultipleOrdersBatchOrdersParameterInnerTimeInForceFromValue(v string) (*PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce, error) {
	ev := PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce: valid values are %v", v, AllowedPlaceMultipleOrdersBatchOrdersParameterInnerTimeInForceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce) IsValid() bool {
	for _, existing := range AllowedPlaceMultipleOrdersBatchOrdersParameterInnerTimeInForceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to placeMultipleOrders_batchOrders_parameter_inner_timeInForce value
func (v PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce) Ptr() *PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce {
	return &v
}

type NullablePlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce struct {
	value *PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce
	isSet bool
}

func (v NullablePlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce) Get() *PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce {
	return v.value
}

func (v *NullablePlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce) Set(val *PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce(val *PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce) *NullablePlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce {
	return &NullablePlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce{value: val, isSet: true}
}

func (v NullablePlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
