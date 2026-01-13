/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"
	"fmt"
)

// PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch the model 'PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch'
type PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch string

// List of placeMultipleOrders_batchOrders_parameter_inner_priceMatch
const (
	PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatchNone       PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch = "NONE"
	PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatchOpponent   PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch = "OPPONENT"
	PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatchOpponent5  PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch = "OPPONENT_5"
	PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatchOpponent10 PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch = "OPPONENT_10"
	PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatchOpponent20 PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch = "OPPONENT_20"
	PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatchQueue      PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch = "QUEUE"
	PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatchQueue5     PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch = "QUEUE_5"
	PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatchQueue10    PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch = "QUEUE_10"
	PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatchQueue20    PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch = "QUEUE_20"
)

// All allowed values of PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch enum
var AllowedPlaceMultipleOrdersBatchOrdersParameterInnerPriceMatchEnumValues = []PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch{
	"NONE",
	"OPPONENT",
	"OPPONENT_5",
	"OPPONENT_10",
	"OPPONENT_20",
	"QUEUE",
	"QUEUE_5",
	"QUEUE_10",
	"QUEUE_20",
}

func (v *PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch(value)
	for _, existing := range AllowedPlaceMultipleOrdersBatchOrdersParameterInnerPriceMatchEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch", value)
}

// NewPlaceMultipleOrdersBatchOrdersParameterInnerPriceMatchFromValue returns a pointer to a valid PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlaceMultipleOrdersBatchOrdersParameterInnerPriceMatchFromValue(v string) (*PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch, error) {
	ev := PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch: valid values are %v", v, AllowedPlaceMultipleOrdersBatchOrdersParameterInnerPriceMatchEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch) IsValid() bool {
	for _, existing := range AllowedPlaceMultipleOrdersBatchOrdersParameterInnerPriceMatchEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to placeMultipleOrders_batchOrders_parameter_inner_priceMatch value
func (v PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch) Ptr() *PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch {
	return &v
}

type NullablePlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch struct {
	value *PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch
	isSet bool
}

func (v NullablePlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch) Get() *PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch {
	return v.value
}

func (v *NullablePlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch) Set(val *PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch(val *PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch) *NullablePlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch {
	return &NullablePlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch{value: val, isSet: true}
}

func (v NullablePlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
