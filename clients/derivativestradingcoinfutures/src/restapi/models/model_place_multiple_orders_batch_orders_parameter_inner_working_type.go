/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"
	"fmt"
)

// PlaceMultipleOrdersBatchOrdersParameterInnerWorkingType the model 'PlaceMultipleOrdersBatchOrdersParameterInnerWorkingType'
type PlaceMultipleOrdersBatchOrdersParameterInnerWorkingType string

// List of placeMultipleOrders_batchOrders_parameter_inner_workingType
const (
	PlaceMultipleOrdersBatchOrdersParameterInnerWorkingTypeMarkPrice     PlaceMultipleOrdersBatchOrdersParameterInnerWorkingType = "MARK_PRICE"
	PlaceMultipleOrdersBatchOrdersParameterInnerWorkingTypeContractPrice PlaceMultipleOrdersBatchOrdersParameterInnerWorkingType = "CONTRACT_PRICE"
)

// All allowed values of PlaceMultipleOrdersBatchOrdersParameterInnerWorkingType enum
var AllowedPlaceMultipleOrdersBatchOrdersParameterInnerWorkingTypeEnumValues = []PlaceMultipleOrdersBatchOrdersParameterInnerWorkingType{
	"MARK_PRICE",
	"CONTRACT_PRICE",
}

func (v *PlaceMultipleOrdersBatchOrdersParameterInnerWorkingType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PlaceMultipleOrdersBatchOrdersParameterInnerWorkingType(value)
	for _, existing := range AllowedPlaceMultipleOrdersBatchOrdersParameterInnerWorkingTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PlaceMultipleOrdersBatchOrdersParameterInnerWorkingType", value)
}

// NewPlaceMultipleOrdersBatchOrdersParameterInnerWorkingTypeFromValue returns a pointer to a valid PlaceMultipleOrdersBatchOrdersParameterInnerWorkingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlaceMultipleOrdersBatchOrdersParameterInnerWorkingTypeFromValue(v string) (*PlaceMultipleOrdersBatchOrdersParameterInnerWorkingType, error) {
	ev := PlaceMultipleOrdersBatchOrdersParameterInnerWorkingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PlaceMultipleOrdersBatchOrdersParameterInnerWorkingType: valid values are %v", v, AllowedPlaceMultipleOrdersBatchOrdersParameterInnerWorkingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PlaceMultipleOrdersBatchOrdersParameterInnerWorkingType) IsValid() bool {
	for _, existing := range AllowedPlaceMultipleOrdersBatchOrdersParameterInnerWorkingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to placeMultipleOrders_batchOrders_parameter_inner_workingType value
func (v PlaceMultipleOrdersBatchOrdersParameterInnerWorkingType) Ptr() *PlaceMultipleOrdersBatchOrdersParameterInnerWorkingType {
	return &v
}

type NullablePlaceMultipleOrdersBatchOrdersParameterInnerWorkingType struct {
	value *PlaceMultipleOrdersBatchOrdersParameterInnerWorkingType
	isSet bool
}

func (v NullablePlaceMultipleOrdersBatchOrdersParameterInnerWorkingType) Get() *PlaceMultipleOrdersBatchOrdersParameterInnerWorkingType {
	return v.value
}

func (v *NullablePlaceMultipleOrdersBatchOrdersParameterInnerWorkingType) Set(val *PlaceMultipleOrdersBatchOrdersParameterInnerWorkingType) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceMultipleOrdersBatchOrdersParameterInnerWorkingType) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceMultipleOrdersBatchOrdersParameterInnerWorkingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceMultipleOrdersBatchOrdersParameterInnerWorkingType(val *PlaceMultipleOrdersBatchOrdersParameterInnerWorkingType) *NullablePlaceMultipleOrdersBatchOrdersParameterInnerWorkingType {
	return &NullablePlaceMultipleOrdersBatchOrdersParameterInnerWorkingType{value: val, isSet: true}
}

func (v NullablePlaceMultipleOrdersBatchOrdersParameterInnerWorkingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceMultipleOrdersBatchOrdersParameterInnerWorkingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
