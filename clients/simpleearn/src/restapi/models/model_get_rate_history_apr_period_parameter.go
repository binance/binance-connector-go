/*
Simple Earn REST API

Earn rewards by subscribing to flexible or locked Simple Earn products.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// GetRateHistoryAprPeriodParameter the model 'GetRateHistoryAprPeriodParameter'
type GetRateHistoryAprPeriodParameter string

// List of getRateHistory_aprPeriod_parameter
const (
	GetRateHistoryAprPeriodParameterDay  GetRateHistoryAprPeriodParameter = "DAY"
	GetRateHistoryAprPeriodParameterYear GetRateHistoryAprPeriodParameter = "YEAR"
)

// All allowed values of GetRateHistoryAprPeriodParameter enum
var AllowedGetRateHistoryAprPeriodParameterEnumValues = []GetRateHistoryAprPeriodParameter{
	"DAY",
	"YEAR",
}

func (v *GetRateHistoryAprPeriodParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GetRateHistoryAprPeriodParameter(value)
	for _, existing := range AllowedGetRateHistoryAprPeriodParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GetRateHistoryAprPeriodParameter", value)
}

// NewGetRateHistoryAprPeriodParameterFromValue returns a pointer to a valid GetRateHistoryAprPeriodParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGetRateHistoryAprPeriodParameterFromValue(v string) (*GetRateHistoryAprPeriodParameter, error) {
	ev := GetRateHistoryAprPeriodParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GetRateHistoryAprPeriodParameter: valid values are %v", v, AllowedGetRateHistoryAprPeriodParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GetRateHistoryAprPeriodParameter) IsValid() bool {
	for _, existing := range AllowedGetRateHistoryAprPeriodParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to getRateHistory_aprPeriod_parameter value
func (v GetRateHistoryAprPeriodParameter) Ptr() *GetRateHistoryAprPeriodParameter {
	return &v
}

type NullableGetRateHistoryAprPeriodParameter struct {
	value *GetRateHistoryAprPeriodParameter
	isSet bool
}

func (v NullableGetRateHistoryAprPeriodParameter) Get() *GetRateHistoryAprPeriodParameter {
	return v.value
}

func (v *NullableGetRateHistoryAprPeriodParameter) Set(val *GetRateHistoryAprPeriodParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRateHistoryAprPeriodParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRateHistoryAprPeriodParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRateHistoryAprPeriodParameter(val *GetRateHistoryAprPeriodParameter) *NullableGetRateHistoryAprPeriodParameter {
	return &NullableGetRateHistoryAprPeriodParameter{value: val, isSet: true}
}

func (v NullableGetRateHistoryAprPeriodParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRateHistoryAprPeriodParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
