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

// TickerTypeParameter the model 'TickerTypeParameter'
type TickerTypeParameter string

// List of ticker_type_parameter
const (
	TickerTypeParameterFull TickerTypeParameter = "FULL"
	TickerTypeParameterMini TickerTypeParameter = "MINI"
)

// All allowed values of TickerTypeParameter enum
var AllowedTickerTypeParameterEnumValues = []TickerTypeParameter{
	"FULL",
	"MINI",
}

func (v *TickerTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TickerTypeParameter(value)
	for _, existing := range AllowedTickerTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TickerTypeParameter", value)
}

// NewTickerTypeParameterFromValue returns a pointer to a valid TickerTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTickerTypeParameterFromValue(v string) (*TickerTypeParameter, error) {
	ev := TickerTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TickerTypeParameter: valid values are %v", v, AllowedTickerTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TickerTypeParameter) IsValid() bool {
	for _, existing := range AllowedTickerTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ticker_type_parameter value
func (v TickerTypeParameter) Ptr() *TickerTypeParameter {
	return &v
}

type NullableTickerTypeParameter struct {
	value *TickerTypeParameter
	isSet bool
}

func (v NullableTickerTypeParameter) Get() *TickerTypeParameter {
	return v.value
}

func (v *NullableTickerTypeParameter) Set(val *TickerTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableTickerTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableTickerTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTickerTypeParameter(val *TickerTypeParameter) *NullableTickerTypeParameter {
	return &NullableTickerTypeParameter{value: val, isSet: true}
}

func (v NullableTickerTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTickerTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
