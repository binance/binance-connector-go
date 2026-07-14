/*
Margin REST API

Access account information, borrow and repay assets, and trade with Binance Margin.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// GetFutureHourlyInterestRateIsIsolatedParameter for isolated margin or not
type GetFutureHourlyInterestRateIsIsolatedParameter string

// List of getFutureHourlyInterestRate_isIsolated_parameter
const (
	GetFutureHourlyInterestRateIsIsolatedParameterTrue  GetFutureHourlyInterestRateIsIsolatedParameter = "TRUE"
	GetFutureHourlyInterestRateIsIsolatedParameterFalse GetFutureHourlyInterestRateIsIsolatedParameter = "FALSE"
)

// All allowed values of GetFutureHourlyInterestRateIsIsolatedParameter enum
var AllowedGetFutureHourlyInterestRateIsIsolatedParameterEnumValues = []GetFutureHourlyInterestRateIsIsolatedParameter{
	"TRUE",
	"FALSE",
}

func (v *GetFutureHourlyInterestRateIsIsolatedParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GetFutureHourlyInterestRateIsIsolatedParameter(value)
	for _, existing := range AllowedGetFutureHourlyInterestRateIsIsolatedParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GetFutureHourlyInterestRateIsIsolatedParameter", value)
}

// NewGetFutureHourlyInterestRateIsIsolatedParameterFromValue returns a pointer to a valid GetFutureHourlyInterestRateIsIsolatedParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGetFutureHourlyInterestRateIsIsolatedParameterFromValue(v string) (*GetFutureHourlyInterestRateIsIsolatedParameter, error) {
	ev := GetFutureHourlyInterestRateIsIsolatedParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GetFutureHourlyInterestRateIsIsolatedParameter: valid values are %v", v, AllowedGetFutureHourlyInterestRateIsIsolatedParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GetFutureHourlyInterestRateIsIsolatedParameter) IsValid() bool {
	for _, existing := range AllowedGetFutureHourlyInterestRateIsIsolatedParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to getFutureHourlyInterestRate_isIsolated_parameter value
func (v GetFutureHourlyInterestRateIsIsolatedParameter) Ptr() *GetFutureHourlyInterestRateIsIsolatedParameter {
	return &v
}

type NullableGetFutureHourlyInterestRateIsIsolatedParameter struct {
	value *GetFutureHourlyInterestRateIsIsolatedParameter
	isSet bool
}

func (v NullableGetFutureHourlyInterestRateIsIsolatedParameter) Get() *GetFutureHourlyInterestRateIsIsolatedParameter {
	return v.value
}

func (v *NullableGetFutureHourlyInterestRateIsIsolatedParameter) Set(val *GetFutureHourlyInterestRateIsIsolatedParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFutureHourlyInterestRateIsIsolatedParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFutureHourlyInterestRateIsIsolatedParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFutureHourlyInterestRateIsIsolatedParameter(val *GetFutureHourlyInterestRateIsIsolatedParameter) *NullableGetFutureHourlyInterestRateIsIsolatedParameter {
	return &NullableGetFutureHourlyInterestRateIsIsolatedParameter{value: val, isSet: true}
}

func (v NullableGetFutureHourlyInterestRateIsIsolatedParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFutureHourlyInterestRateIsIsolatedParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
