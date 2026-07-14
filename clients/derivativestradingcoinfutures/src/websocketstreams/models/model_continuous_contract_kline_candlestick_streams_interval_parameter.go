/*
Futures (COIN-M) WebSocket Market Streams

Access market data, manage accounts, and trade COIN-M perpetual and delivery futures.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// ContinuousContractKlineCandlestickStreamsIntervalParameter the model 'ContinuousContractKlineCandlestickStreamsIntervalParameter'
type ContinuousContractKlineCandlestickStreamsIntervalParameter string

// List of continuousContractKlineCandlestickStreams_interval_parameter
const (
	ContinuousContractKlineCandlestickStreamsIntervalParameterInterval1m  ContinuousContractKlineCandlestickStreamsIntervalParameter = "1m"
	ContinuousContractKlineCandlestickStreamsIntervalParameterInterval3m  ContinuousContractKlineCandlestickStreamsIntervalParameter = "3m"
	ContinuousContractKlineCandlestickStreamsIntervalParameterInterval5m  ContinuousContractKlineCandlestickStreamsIntervalParameter = "5m"
	ContinuousContractKlineCandlestickStreamsIntervalParameterInterval15m ContinuousContractKlineCandlestickStreamsIntervalParameter = "15m"
	ContinuousContractKlineCandlestickStreamsIntervalParameterInterval30m ContinuousContractKlineCandlestickStreamsIntervalParameter = "30m"
	ContinuousContractKlineCandlestickStreamsIntervalParameterInterval1h  ContinuousContractKlineCandlestickStreamsIntervalParameter = "1h"
	ContinuousContractKlineCandlestickStreamsIntervalParameterInterval2h  ContinuousContractKlineCandlestickStreamsIntervalParameter = "2h"
	ContinuousContractKlineCandlestickStreamsIntervalParameterInterval4h  ContinuousContractKlineCandlestickStreamsIntervalParameter = "4h"
	ContinuousContractKlineCandlestickStreamsIntervalParameterInterval6h  ContinuousContractKlineCandlestickStreamsIntervalParameter = "6h"
	ContinuousContractKlineCandlestickStreamsIntervalParameterInterval8h  ContinuousContractKlineCandlestickStreamsIntervalParameter = "8h"
	ContinuousContractKlineCandlestickStreamsIntervalParameterInterval12h ContinuousContractKlineCandlestickStreamsIntervalParameter = "12h"
	ContinuousContractKlineCandlestickStreamsIntervalParameterInterval1d  ContinuousContractKlineCandlestickStreamsIntervalParameter = "1d"
	ContinuousContractKlineCandlestickStreamsIntervalParameterInterval3d  ContinuousContractKlineCandlestickStreamsIntervalParameter = "3d"
	ContinuousContractKlineCandlestickStreamsIntervalParameterInterval1w  ContinuousContractKlineCandlestickStreamsIntervalParameter = "1w"
	ContinuousContractKlineCandlestickStreamsIntervalParameterInterval1M  ContinuousContractKlineCandlestickStreamsIntervalParameter = "1M"
)

// All allowed values of ContinuousContractKlineCandlestickStreamsIntervalParameter enum
var AllowedContinuousContractKlineCandlestickStreamsIntervalParameterEnumValues = []ContinuousContractKlineCandlestickStreamsIntervalParameter{
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

func (v *ContinuousContractKlineCandlestickStreamsIntervalParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ContinuousContractKlineCandlestickStreamsIntervalParameter(value)
	for _, existing := range AllowedContinuousContractKlineCandlestickStreamsIntervalParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ContinuousContractKlineCandlestickStreamsIntervalParameter", value)
}

// NewContinuousContractKlineCandlestickStreamsIntervalParameterFromValue returns a pointer to a valid ContinuousContractKlineCandlestickStreamsIntervalParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewContinuousContractKlineCandlestickStreamsIntervalParameterFromValue(v string) (*ContinuousContractKlineCandlestickStreamsIntervalParameter, error) {
	ev := ContinuousContractKlineCandlestickStreamsIntervalParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ContinuousContractKlineCandlestickStreamsIntervalParameter: valid values are %v", v, AllowedContinuousContractKlineCandlestickStreamsIntervalParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ContinuousContractKlineCandlestickStreamsIntervalParameter) IsValid() bool {
	for _, existing := range AllowedContinuousContractKlineCandlestickStreamsIntervalParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to continuousContractKlineCandlestickStreams_interval_parameter value
func (v ContinuousContractKlineCandlestickStreamsIntervalParameter) Ptr() *ContinuousContractKlineCandlestickStreamsIntervalParameter {
	return &v
}

type NullableContinuousContractKlineCandlestickStreamsIntervalParameter struct {
	value *ContinuousContractKlineCandlestickStreamsIntervalParameter
	isSet bool
}

func (v NullableContinuousContractKlineCandlestickStreamsIntervalParameter) Get() *ContinuousContractKlineCandlestickStreamsIntervalParameter {
	return v.value
}

func (v *NullableContinuousContractKlineCandlestickStreamsIntervalParameter) Set(val *ContinuousContractKlineCandlestickStreamsIntervalParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableContinuousContractKlineCandlestickStreamsIntervalParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableContinuousContractKlineCandlestickStreamsIntervalParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContinuousContractKlineCandlestickStreamsIntervalParameter(val *ContinuousContractKlineCandlestickStreamsIntervalParameter) *NullableContinuousContractKlineCandlestickStreamsIntervalParameter {
	return &NullableContinuousContractKlineCandlestickStreamsIntervalParameter{value: val, isSet: true}
}

func (v NullableContinuousContractKlineCandlestickStreamsIntervalParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContinuousContractKlineCandlestickStreamsIntervalParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
