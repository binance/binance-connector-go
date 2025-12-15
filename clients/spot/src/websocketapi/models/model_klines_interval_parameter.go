/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)

API version: 1.0.0
*/

package models

import (
	"encoding/json"
	"fmt"
)

// KlinesIntervalParameter the model 'KlinesIntervalParameter'
type KlinesIntervalParameter string

// List of klines_interval_parameter
const (
	KlinesIntervalParameterInterval1s  KlinesIntervalParameter = "1s"
	KlinesIntervalParameterInterval1m  KlinesIntervalParameter = "1m"
	KlinesIntervalParameterInterval3m  KlinesIntervalParameter = "3m"
	KlinesIntervalParameterInterval5m  KlinesIntervalParameter = "5m"
	KlinesIntervalParameterInterval15m KlinesIntervalParameter = "15m"
	KlinesIntervalParameterInterval30m KlinesIntervalParameter = "30m"
	KlinesIntervalParameterInterval1h  KlinesIntervalParameter = "1h"
	KlinesIntervalParameterInterval2h  KlinesIntervalParameter = "2h"
	KlinesIntervalParameterInterval4h  KlinesIntervalParameter = "4h"
	KlinesIntervalParameterInterval6h  KlinesIntervalParameter = "6h"
	KlinesIntervalParameterInterval8h  KlinesIntervalParameter = "8h"
	KlinesIntervalParameterInterval12h KlinesIntervalParameter = "12h"
	KlinesIntervalParameterInterval1d  KlinesIntervalParameter = "1d"
	KlinesIntervalParameterInterval3d  KlinesIntervalParameter = "3d"
	KlinesIntervalParameterInterval1w  KlinesIntervalParameter = "1w"
	KlinesIntervalParameterInterval1M  KlinesIntervalParameter = "1M"
)

// All allowed values of KlinesIntervalParameter enum
var AllowedKlinesIntervalParameterEnumValues = []KlinesIntervalParameter{
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

func (v *KlinesIntervalParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := KlinesIntervalParameter(value)
	for _, existing := range AllowedKlinesIntervalParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid KlinesIntervalParameter", value)
}

// NewKlinesIntervalParameterFromValue returns a pointer to a valid KlinesIntervalParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKlinesIntervalParameterFromValue(v string) (*KlinesIntervalParameter, error) {
	ev := KlinesIntervalParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KlinesIntervalParameter: valid values are %v", v, AllowedKlinesIntervalParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KlinesIntervalParameter) IsValid() bool {
	for _, existing := range AllowedKlinesIntervalParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to klines_interval_parameter value
func (v KlinesIntervalParameter) Ptr() *KlinesIntervalParameter {
	return &v
}

type NullableKlinesIntervalParameter struct {
	value *KlinesIntervalParameter
	isSet bool
}

func (v NullableKlinesIntervalParameter) Get() *KlinesIntervalParameter {
	return v.value
}

func (v *NullableKlinesIntervalParameter) Set(val *KlinesIntervalParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableKlinesIntervalParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableKlinesIntervalParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKlinesIntervalParameter(val *KlinesIntervalParameter) *NullableKlinesIntervalParameter {
	return &NullableKlinesIntervalParameter{value: val, isSet: true}
}

func (v NullableKlinesIntervalParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlinesIntervalParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
