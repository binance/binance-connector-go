/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"
	"fmt"
)

// PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType the model 'PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType'
type PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType string

// List of placeMultipleOrders_batchOrders_parameter_inner_newOrderRespType
const (
	PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespTypeAck    PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType = "ACK"
	PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespTypeResult PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType = "RESULT"
)

// All allowed values of PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType enum
var AllowedPlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespTypeEnumValues = []PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType{
	"ACK",
	"RESULT",
}

func (v *PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType(value)
	for _, existing := range AllowedPlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType", value)
}

// NewPlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespTypeFromValue returns a pointer to a valid PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespTypeFromValue(v string) (*PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType, error) {
	ev := PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType: valid values are %v", v, AllowedPlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType) IsValid() bool {
	for _, existing := range AllowedPlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to placeMultipleOrders_batchOrders_parameter_inner_newOrderRespType value
func (v PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType) Ptr() *PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType {
	return &v
}

type NullablePlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType struct {
	value *PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType
	isSet bool
}

func (v NullablePlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType) Get() *PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType {
	return v.value
}

func (v *NullablePlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType) Set(val *PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType(val *PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType) *NullablePlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType {
	return &NullablePlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType{value: val, isSet: true}
}

func (v NullablePlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
