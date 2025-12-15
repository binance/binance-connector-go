/*
Binance Spot WebSocket Streams

OpenAPI Specifications for the Binance Spot WebSocket Streams  API documents:   - [Github web-socket-streams documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-streams.md)   - [General API information for web-socket-streams on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-streams)

API version: 1.0.0
*/

package models

import (
	"encoding/json"
	"fmt"
)

// KlineIntervalParameter the model 'KlineIntervalParameter'
type KlineIntervalParameter string

// List of kline_interval_parameter
const (
	KlineIntervalParameterInterval1s  KlineIntervalParameter = "1s"
	KlineIntervalParameterInterval1m  KlineIntervalParameter = "1m"
	KlineIntervalParameterInterval3m  KlineIntervalParameter = "3m"
	KlineIntervalParameterInterval5m  KlineIntervalParameter = "5m"
	KlineIntervalParameterInterval15m KlineIntervalParameter = "15m"
	KlineIntervalParameterInterval30m KlineIntervalParameter = "30m"
	KlineIntervalParameterInterval1h  KlineIntervalParameter = "1h"
	KlineIntervalParameterInterval2h  KlineIntervalParameter = "2h"
	KlineIntervalParameterInterval4h  KlineIntervalParameter = "4h"
	KlineIntervalParameterInterval6h  KlineIntervalParameter = "6h"
	KlineIntervalParameterInterval8h  KlineIntervalParameter = "8h"
	KlineIntervalParameterInterval12h KlineIntervalParameter = "12h"
	KlineIntervalParameterInterval1d  KlineIntervalParameter = "1d"
	KlineIntervalParameterInterval3d  KlineIntervalParameter = "3d"
	KlineIntervalParameterInterval1w  KlineIntervalParameter = "1w"
	KlineIntervalParameterInterval1M  KlineIntervalParameter = "1M"
)

// All allowed values of KlineIntervalParameter enum
var AllowedKlineIntervalParameterEnumValues = []KlineIntervalParameter{
	"1s",
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

func (v *KlineIntervalParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := KlineIntervalParameter(value)
	for _, existing := range AllowedKlineIntervalParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid KlineIntervalParameter", value)
}

// NewKlineIntervalParameterFromValue returns a pointer to a valid KlineIntervalParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKlineIntervalParameterFromValue(v string) (*KlineIntervalParameter, error) {
	ev := KlineIntervalParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KlineIntervalParameter: valid values are %v", v, AllowedKlineIntervalParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KlineIntervalParameter) IsValid() bool {
	for _, existing := range AllowedKlineIntervalParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to kline_interval_parameter value
func (v KlineIntervalParameter) Ptr() *KlineIntervalParameter {
	return &v
}

type NullableKlineIntervalParameter struct {
	value *KlineIntervalParameter
	isSet bool
}

func (v NullableKlineIntervalParameter) Get() *KlineIntervalParameter {
	return v.value
}

func (v *NullableKlineIntervalParameter) Set(val *KlineIntervalParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableKlineIntervalParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableKlineIntervalParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKlineIntervalParameter(val *KlineIntervalParameter) *NullableKlineIntervalParameter {
	return &NullableKlineIntervalParameter{value: val, isSet: true}
}

func (v NullableKlineIntervalParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlineIntervalParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
