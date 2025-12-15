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

// ExchangeInfoSymbolStatusParameter the model 'ExchangeInfoSymbolStatusParameter'
type ExchangeInfoSymbolStatusParameter string

// List of exchangeInfo_symbolStatus_parameter
const (
	ExchangeInfoSymbolStatusParameterTrading          ExchangeInfoSymbolStatusParameter = "TRADING"
	ExchangeInfoSymbolStatusParameterEndOfDay         ExchangeInfoSymbolStatusParameter = "END_OF_DAY"
	ExchangeInfoSymbolStatusParameterHalt             ExchangeInfoSymbolStatusParameter = "HALT"
	ExchangeInfoSymbolStatusParameterBreak            ExchangeInfoSymbolStatusParameter = "BREAK"
	ExchangeInfoSymbolStatusParameterNonRepresentable ExchangeInfoSymbolStatusParameter = "NON_REPRESENTABLE"
)

// All allowed values of ExchangeInfoSymbolStatusParameter enum
var AllowedExchangeInfoSymbolStatusParameterEnumValues = []ExchangeInfoSymbolStatusParameter{
	"TRADING",
	"END_OF_DAY",
	"HALT",
	"BREAK",
	"NON_REPRESENTABLE",
}

func (v *ExchangeInfoSymbolStatusParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExchangeInfoSymbolStatusParameter(value)
	for _, existing := range AllowedExchangeInfoSymbolStatusParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExchangeInfoSymbolStatusParameter", value)
}

// NewExchangeInfoSymbolStatusParameterFromValue returns a pointer to a valid ExchangeInfoSymbolStatusParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExchangeInfoSymbolStatusParameterFromValue(v string) (*ExchangeInfoSymbolStatusParameter, error) {
	ev := ExchangeInfoSymbolStatusParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExchangeInfoSymbolStatusParameter: valid values are %v", v, AllowedExchangeInfoSymbolStatusParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExchangeInfoSymbolStatusParameter) IsValid() bool {
	for _, existing := range AllowedExchangeInfoSymbolStatusParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to exchangeInfo_symbolStatus_parameter value
func (v ExchangeInfoSymbolStatusParameter) Ptr() *ExchangeInfoSymbolStatusParameter {
	return &v
}

type NullableExchangeInfoSymbolStatusParameter struct {
	value *ExchangeInfoSymbolStatusParameter
	isSet bool
}

func (v NullableExchangeInfoSymbolStatusParameter) Get() *ExchangeInfoSymbolStatusParameter {
	return v.value
}

func (v *NullableExchangeInfoSymbolStatusParameter) Set(val *ExchangeInfoSymbolStatusParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeInfoSymbolStatusParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeInfoSymbolStatusParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeInfoSymbolStatusParameter(val *ExchangeInfoSymbolStatusParameter) *NullableExchangeInfoSymbolStatusParameter {
	return &NullableExchangeInfoSymbolStatusParameter{value: val, isSet: true}
}

func (v NullableExchangeInfoSymbolStatusParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeInfoSymbolStatusParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
