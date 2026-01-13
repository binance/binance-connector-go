/*
Binance Spot WebSocket Streams

OpenAPI Specifications for the Binance Spot WebSocket Streams  API documents:   - [Github web-socket-streams documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-streams.md)   - [General API information for web-socket-streams on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-streams)
*/

package models

import (
	"encoding/json"
	"fmt"
)

// AllMarketRollingWindowTickerWindowSizeParameter the model 'AllMarketRollingWindowTickerWindowSizeParameter'
type AllMarketRollingWindowTickerWindowSizeParameter string

// List of allMarketRollingWindowTicker_windowSize_parameter
const (
	AllMarketRollingWindowTickerWindowSizeParameterWindowSize1h AllMarketRollingWindowTickerWindowSizeParameter = "1h"
	AllMarketRollingWindowTickerWindowSizeParameterWindowSize4h AllMarketRollingWindowTickerWindowSizeParameter = "4h"
	AllMarketRollingWindowTickerWindowSizeParameterWindowSize1d AllMarketRollingWindowTickerWindowSizeParameter = "1d"
)

// All allowed values of AllMarketRollingWindowTickerWindowSizeParameter enum
var AllowedAllMarketRollingWindowTickerWindowSizeParameterEnumValues = []AllMarketRollingWindowTickerWindowSizeParameter{
	"1h",
	"4h",
	"1d",
}

func (v *AllMarketRollingWindowTickerWindowSizeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AllMarketRollingWindowTickerWindowSizeParameter(value)
	for _, existing := range AllowedAllMarketRollingWindowTickerWindowSizeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AllMarketRollingWindowTickerWindowSizeParameter", value)
}

// NewAllMarketRollingWindowTickerWindowSizeParameterFromValue returns a pointer to a valid AllMarketRollingWindowTickerWindowSizeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAllMarketRollingWindowTickerWindowSizeParameterFromValue(v string) (*AllMarketRollingWindowTickerWindowSizeParameter, error) {
	ev := AllMarketRollingWindowTickerWindowSizeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AllMarketRollingWindowTickerWindowSizeParameter: valid values are %v", v, AllowedAllMarketRollingWindowTickerWindowSizeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AllMarketRollingWindowTickerWindowSizeParameter) IsValid() bool {
	for _, existing := range AllowedAllMarketRollingWindowTickerWindowSizeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to allMarketRollingWindowTicker_windowSize_parameter value
func (v AllMarketRollingWindowTickerWindowSizeParameter) Ptr() *AllMarketRollingWindowTickerWindowSizeParameter {
	return &v
}

type NullableAllMarketRollingWindowTickerWindowSizeParameter struct {
	value *AllMarketRollingWindowTickerWindowSizeParameter
	isSet bool
}

func (v NullableAllMarketRollingWindowTickerWindowSizeParameter) Get() *AllMarketRollingWindowTickerWindowSizeParameter {
	return v.value
}

func (v *NullableAllMarketRollingWindowTickerWindowSizeParameter) Set(val *AllMarketRollingWindowTickerWindowSizeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableAllMarketRollingWindowTickerWindowSizeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableAllMarketRollingWindowTickerWindowSizeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllMarketRollingWindowTickerWindowSizeParameter(val *AllMarketRollingWindowTickerWindowSizeParameter) *NullableAllMarketRollingWindowTickerWindowSizeParameter {
	return &NullableAllMarketRollingWindowTickerWindowSizeParameter{value: val, isSet: true}
}

func (v NullableAllMarketRollingWindowTickerWindowSizeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllMarketRollingWindowTickerWindowSizeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
