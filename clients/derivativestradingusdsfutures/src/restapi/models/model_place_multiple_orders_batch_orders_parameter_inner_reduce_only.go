/*
Futures (USDⓈ-M) REST API

Access market data, manage accounts, and trade USDⓈ-M perpetual futures.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// PlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly the model 'PlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly'
type PlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly string

// List of placeMultipleOrders_batchOrders_parameter_inner_reduceOnly
const (
	PlaceMultipleOrdersBatchOrdersParameterInnerReduceOnlyTrue  PlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly = "true"
	PlaceMultipleOrdersBatchOrdersParameterInnerReduceOnlyFalse PlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly = "false"
)

// All allowed values of PlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly enum
var AllowedPlaceMultipleOrdersBatchOrdersParameterInnerReduceOnlyEnumValues = []PlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly{
	"true",
	"false",
}

func (v *PlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly(value)
	for _, existing := range AllowedPlaceMultipleOrdersBatchOrdersParameterInnerReduceOnlyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly", value)
}

// NewPlaceMultipleOrdersBatchOrdersParameterInnerReduceOnlyFromValue returns a pointer to a valid PlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlaceMultipleOrdersBatchOrdersParameterInnerReduceOnlyFromValue(v string) (*PlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly, error) {
	ev := PlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly: valid values are %v", v, AllowedPlaceMultipleOrdersBatchOrdersParameterInnerReduceOnlyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly) IsValid() bool {
	for _, existing := range AllowedPlaceMultipleOrdersBatchOrdersParameterInnerReduceOnlyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to placeMultipleOrders_batchOrders_parameter_inner_reduceOnly value
func (v PlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly) Ptr() *PlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly {
	return &v
}

type NullablePlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly struct {
	value *PlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly
	isSet bool
}

func (v NullablePlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly) Get() *PlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly {
	return v.value
}

func (v *NullablePlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly) Set(val *PlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly(val *PlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly) *NullablePlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly {
	return &NullablePlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly{value: val, isSet: true}
}

func (v NullablePlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
