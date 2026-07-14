/*
Algo Trading REST API

Programmatic access to Binance’s execution algorithms for creating and managing Spot and Futures algo orders.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// TimeWeightedAveragePriceFutureAlgoPositionSideParameter the model 'TimeWeightedAveragePriceFutureAlgoPositionSideParameter'
type TimeWeightedAveragePriceFutureAlgoPositionSideParameter string

// List of timeWeightedAveragePriceFutureAlgo_positionSide_parameter
const (
	TimeWeightedAveragePriceFutureAlgoPositionSideParameterBoth  TimeWeightedAveragePriceFutureAlgoPositionSideParameter = "BOTH"
	TimeWeightedAveragePriceFutureAlgoPositionSideParameterLong  TimeWeightedAveragePriceFutureAlgoPositionSideParameter = "LONG"
	TimeWeightedAveragePriceFutureAlgoPositionSideParameterShort TimeWeightedAveragePriceFutureAlgoPositionSideParameter = "SHORT"
)

// All allowed values of TimeWeightedAveragePriceFutureAlgoPositionSideParameter enum
var AllowedTimeWeightedAveragePriceFutureAlgoPositionSideParameterEnumValues = []TimeWeightedAveragePriceFutureAlgoPositionSideParameter{
	"BOTH",
	"LONG",
	"SHORT",
}

func (v *TimeWeightedAveragePriceFutureAlgoPositionSideParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TimeWeightedAveragePriceFutureAlgoPositionSideParameter(value)
	for _, existing := range AllowedTimeWeightedAveragePriceFutureAlgoPositionSideParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TimeWeightedAveragePriceFutureAlgoPositionSideParameter", value)
}

// NewTimeWeightedAveragePriceFutureAlgoPositionSideParameterFromValue returns a pointer to a valid TimeWeightedAveragePriceFutureAlgoPositionSideParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTimeWeightedAveragePriceFutureAlgoPositionSideParameterFromValue(v string) (*TimeWeightedAveragePriceFutureAlgoPositionSideParameter, error) {
	ev := TimeWeightedAveragePriceFutureAlgoPositionSideParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TimeWeightedAveragePriceFutureAlgoPositionSideParameter: valid values are %v", v, AllowedTimeWeightedAveragePriceFutureAlgoPositionSideParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TimeWeightedAveragePriceFutureAlgoPositionSideParameter) IsValid() bool {
	for _, existing := range AllowedTimeWeightedAveragePriceFutureAlgoPositionSideParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to timeWeightedAveragePriceFutureAlgo_positionSide_parameter value
func (v TimeWeightedAveragePriceFutureAlgoPositionSideParameter) Ptr() *TimeWeightedAveragePriceFutureAlgoPositionSideParameter {
	return &v
}

type NullableTimeWeightedAveragePriceFutureAlgoPositionSideParameter struct {
	value *TimeWeightedAveragePriceFutureAlgoPositionSideParameter
	isSet bool
}

func (v NullableTimeWeightedAveragePriceFutureAlgoPositionSideParameter) Get() *TimeWeightedAveragePriceFutureAlgoPositionSideParameter {
	return v.value
}

func (v *NullableTimeWeightedAveragePriceFutureAlgoPositionSideParameter) Set(val *TimeWeightedAveragePriceFutureAlgoPositionSideParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeWeightedAveragePriceFutureAlgoPositionSideParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeWeightedAveragePriceFutureAlgoPositionSideParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeWeightedAveragePriceFutureAlgoPositionSideParameter(val *TimeWeightedAveragePriceFutureAlgoPositionSideParameter) *NullableTimeWeightedAveragePriceFutureAlgoPositionSideParameter {
	return &NullableTimeWeightedAveragePriceFutureAlgoPositionSideParameter{value: val, isSet: true}
}

func (v NullableTimeWeightedAveragePriceFutureAlgoPositionSideParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeWeightedAveragePriceFutureAlgoPositionSideParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
