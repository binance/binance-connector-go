/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"
	"fmt"
)

// PlaceMultipleOrdersOrdersParameterInnerNewOrderRespType the model 'PlaceMultipleOrdersOrdersParameterInnerNewOrderRespType'
type PlaceMultipleOrdersOrdersParameterInnerNewOrderRespType string

// List of placeMultipleOrders_orders_parameter_inner_newOrderRespType
const (
	PlaceMultipleOrdersOrdersParameterInnerNewOrderRespTypeAck    PlaceMultipleOrdersOrdersParameterInnerNewOrderRespType = "ACK"
	PlaceMultipleOrdersOrdersParameterInnerNewOrderRespTypeResult PlaceMultipleOrdersOrdersParameterInnerNewOrderRespType = "RESULT"
)

// All allowed values of PlaceMultipleOrdersOrdersParameterInnerNewOrderRespType enum
var AllowedPlaceMultipleOrdersOrdersParameterInnerNewOrderRespTypeEnumValues = []PlaceMultipleOrdersOrdersParameterInnerNewOrderRespType{
	"ACK",
	"RESULT",
}

func (v *PlaceMultipleOrdersOrdersParameterInnerNewOrderRespType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PlaceMultipleOrdersOrdersParameterInnerNewOrderRespType(value)
	for _, existing := range AllowedPlaceMultipleOrdersOrdersParameterInnerNewOrderRespTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PlaceMultipleOrdersOrdersParameterInnerNewOrderRespType", value)
}

// NewPlaceMultipleOrdersOrdersParameterInnerNewOrderRespTypeFromValue returns a pointer to a valid PlaceMultipleOrdersOrdersParameterInnerNewOrderRespType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlaceMultipleOrdersOrdersParameterInnerNewOrderRespTypeFromValue(v string) (*PlaceMultipleOrdersOrdersParameterInnerNewOrderRespType, error) {
	ev := PlaceMultipleOrdersOrdersParameterInnerNewOrderRespType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PlaceMultipleOrdersOrdersParameterInnerNewOrderRespType: valid values are %v", v, AllowedPlaceMultipleOrdersOrdersParameterInnerNewOrderRespTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PlaceMultipleOrdersOrdersParameterInnerNewOrderRespType) IsValid() bool {
	for _, existing := range AllowedPlaceMultipleOrdersOrdersParameterInnerNewOrderRespTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to placeMultipleOrders_orders_parameter_inner_newOrderRespType value
func (v PlaceMultipleOrdersOrdersParameterInnerNewOrderRespType) Ptr() *PlaceMultipleOrdersOrdersParameterInnerNewOrderRespType {
	return &v
}

type NullablePlaceMultipleOrdersOrdersParameterInnerNewOrderRespType struct {
	value *PlaceMultipleOrdersOrdersParameterInnerNewOrderRespType
	isSet bool
}

func (v NullablePlaceMultipleOrdersOrdersParameterInnerNewOrderRespType) Get() *PlaceMultipleOrdersOrdersParameterInnerNewOrderRespType {
	return v.value
}

func (v *NullablePlaceMultipleOrdersOrdersParameterInnerNewOrderRespType) Set(val *PlaceMultipleOrdersOrdersParameterInnerNewOrderRespType) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceMultipleOrdersOrdersParameterInnerNewOrderRespType) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceMultipleOrdersOrdersParameterInnerNewOrderRespType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceMultipleOrdersOrdersParameterInnerNewOrderRespType(val *PlaceMultipleOrdersOrdersParameterInnerNewOrderRespType) *NullablePlaceMultipleOrdersOrdersParameterInnerNewOrderRespType {
	return &NullablePlaceMultipleOrdersOrdersParameterInnerNewOrderRespType{value: val, isSet: true}
}

func (v NullablePlaceMultipleOrdersOrdersParameterInnerNewOrderRespType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceMultipleOrdersOrdersParameterInnerNewOrderRespType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
