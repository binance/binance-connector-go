/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"
	"fmt"
)

// ContinuousContractKlineCandlestickDataIntervalParameter the model 'ContinuousContractKlineCandlestickDataIntervalParameter'
type ContinuousContractKlineCandlestickDataIntervalParameter string

// List of continuousContractKlineCandlestickData_interval_parameter
const (
	ContinuousContractKlineCandlestickDataIntervalParameterInterval1m  ContinuousContractKlineCandlestickDataIntervalParameter = "1m"
	ContinuousContractKlineCandlestickDataIntervalParameterInterval3m  ContinuousContractKlineCandlestickDataIntervalParameter = "3m"
	ContinuousContractKlineCandlestickDataIntervalParameterInterval5m  ContinuousContractKlineCandlestickDataIntervalParameter = "5m"
	ContinuousContractKlineCandlestickDataIntervalParameterInterval15m ContinuousContractKlineCandlestickDataIntervalParameter = "15m"
	ContinuousContractKlineCandlestickDataIntervalParameterInterval30m ContinuousContractKlineCandlestickDataIntervalParameter = "30m"
	ContinuousContractKlineCandlestickDataIntervalParameterInterval1h  ContinuousContractKlineCandlestickDataIntervalParameter = "1h"
	ContinuousContractKlineCandlestickDataIntervalParameterInterval2h  ContinuousContractKlineCandlestickDataIntervalParameter = "2h"
	ContinuousContractKlineCandlestickDataIntervalParameterInterval4h  ContinuousContractKlineCandlestickDataIntervalParameter = "4h"
	ContinuousContractKlineCandlestickDataIntervalParameterInterval6h  ContinuousContractKlineCandlestickDataIntervalParameter = "6h"
	ContinuousContractKlineCandlestickDataIntervalParameterInterval8h  ContinuousContractKlineCandlestickDataIntervalParameter = "8h"
	ContinuousContractKlineCandlestickDataIntervalParameterInterval12h ContinuousContractKlineCandlestickDataIntervalParameter = "12h"
	ContinuousContractKlineCandlestickDataIntervalParameterInterval1d  ContinuousContractKlineCandlestickDataIntervalParameter = "1d"
	ContinuousContractKlineCandlestickDataIntervalParameterInterval3d  ContinuousContractKlineCandlestickDataIntervalParameter = "3d"
	ContinuousContractKlineCandlestickDataIntervalParameterInterval1w  ContinuousContractKlineCandlestickDataIntervalParameter = "1w"
	ContinuousContractKlineCandlestickDataIntervalParameterInterval1M  ContinuousContractKlineCandlestickDataIntervalParameter = "1M"
)

// All allowed values of ContinuousContractKlineCandlestickDataIntervalParameter enum
var AllowedContinuousContractKlineCandlestickDataIntervalParameterEnumValues = []ContinuousContractKlineCandlestickDataIntervalParameter{
	"1m",
	"3m",
	"5m",
	"15m",
	"30m",
	"1h",
	"2h",
	"4h",
	"6h",
	"8h",
	"12h",
	"1d",
	"3d",
	"1w",
	"1M",
}

func (v *ContinuousContractKlineCandlestickDataIntervalParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ContinuousContractKlineCandlestickDataIntervalParameter(value)
	for _, existing := range AllowedContinuousContractKlineCandlestickDataIntervalParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ContinuousContractKlineCandlestickDataIntervalParameter", value)
}

// NewContinuousContractKlineCandlestickDataIntervalParameterFromValue returns a pointer to a valid ContinuousContractKlineCandlestickDataIntervalParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewContinuousContractKlineCandlestickDataIntervalParameterFromValue(v string) (*ContinuousContractKlineCandlestickDataIntervalParameter, error) {
	ev := ContinuousContractKlineCandlestickDataIntervalParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ContinuousContractKlineCandlestickDataIntervalParameter: valid values are %v", v, AllowedContinuousContractKlineCandlestickDataIntervalParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ContinuousContractKlineCandlestickDataIntervalParameter) IsValid() bool {
	for _, existing := range AllowedContinuousContractKlineCandlestickDataIntervalParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to continuousContractKlineCandlestickData_interval_parameter value
func (v ContinuousContractKlineCandlestickDataIntervalParameter) Ptr() *ContinuousContractKlineCandlestickDataIntervalParameter {
	return &v
}

type NullableContinuousContractKlineCandlestickDataIntervalParameter struct {
	value *ContinuousContractKlineCandlestickDataIntervalParameter
	isSet bool
}

func (v NullableContinuousContractKlineCandlestickDataIntervalParameter) Get() *ContinuousContractKlineCandlestickDataIntervalParameter {
	return v.value
}

func (v *NullableContinuousContractKlineCandlestickDataIntervalParameter) Set(val *ContinuousContractKlineCandlestickDataIntervalParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableContinuousContractKlineCandlestickDataIntervalParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableContinuousContractKlineCandlestickDataIntervalParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContinuousContractKlineCandlestickDataIntervalParameter(val *ContinuousContractKlineCandlestickDataIntervalParameter) *NullableContinuousContractKlineCandlestickDataIntervalParameter {
	return &NullableContinuousContractKlineCandlestickDataIntervalParameter{value: val, isSet: true}
}

func (v NullableContinuousContractKlineCandlestickDataIntervalParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContinuousContractKlineCandlestickDataIntervalParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
