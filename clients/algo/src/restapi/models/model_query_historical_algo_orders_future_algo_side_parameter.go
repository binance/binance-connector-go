/*
Algo Trading REST API

Programmatic access to Binance’s execution algorithms for creating and managing Spot and Futures algo orders.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// QueryHistoricalAlgoOrdersFutureAlgoSideParameter the model 'QueryHistoricalAlgoOrdersFutureAlgoSideParameter'
type QueryHistoricalAlgoOrdersFutureAlgoSideParameter string

// List of queryHistoricalAlgoOrdersFutureAlgo_side_parameter
const (
	QueryHistoricalAlgoOrdersFutureAlgoSideParameterBuy  QueryHistoricalAlgoOrdersFutureAlgoSideParameter = "BUY"
	QueryHistoricalAlgoOrdersFutureAlgoSideParameterSell QueryHistoricalAlgoOrdersFutureAlgoSideParameter = "SELL"
)

// All allowed values of QueryHistoricalAlgoOrdersFutureAlgoSideParameter enum
var AllowedQueryHistoricalAlgoOrdersFutureAlgoSideParameterEnumValues = []QueryHistoricalAlgoOrdersFutureAlgoSideParameter{
	"BUY",
	"SELL",
}

func (v *QueryHistoricalAlgoOrdersFutureAlgoSideParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := QueryHistoricalAlgoOrdersFutureAlgoSideParameter(value)
	for _, existing := range AllowedQueryHistoricalAlgoOrdersFutureAlgoSideParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid QueryHistoricalAlgoOrdersFutureAlgoSideParameter", value)
}

// NewQueryHistoricalAlgoOrdersFutureAlgoSideParameterFromValue returns a pointer to a valid QueryHistoricalAlgoOrdersFutureAlgoSideParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQueryHistoricalAlgoOrdersFutureAlgoSideParameterFromValue(v string) (*QueryHistoricalAlgoOrdersFutureAlgoSideParameter, error) {
	ev := QueryHistoricalAlgoOrdersFutureAlgoSideParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QueryHistoricalAlgoOrdersFutureAlgoSideParameter: valid values are %v", v, AllowedQueryHistoricalAlgoOrdersFutureAlgoSideParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QueryHistoricalAlgoOrdersFutureAlgoSideParameter) IsValid() bool {
	for _, existing := range AllowedQueryHistoricalAlgoOrdersFutureAlgoSideParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to queryHistoricalAlgoOrdersFutureAlgo_side_parameter value
func (v QueryHistoricalAlgoOrdersFutureAlgoSideParameter) Ptr() *QueryHistoricalAlgoOrdersFutureAlgoSideParameter {
	return &v
}

type NullableQueryHistoricalAlgoOrdersFutureAlgoSideParameter struct {
	value *QueryHistoricalAlgoOrdersFutureAlgoSideParameter
	isSet bool
}

func (v NullableQueryHistoricalAlgoOrdersFutureAlgoSideParameter) Get() *QueryHistoricalAlgoOrdersFutureAlgoSideParameter {
	return v.value
}

func (v *NullableQueryHistoricalAlgoOrdersFutureAlgoSideParameter) Set(val *QueryHistoricalAlgoOrdersFutureAlgoSideParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryHistoricalAlgoOrdersFutureAlgoSideParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryHistoricalAlgoOrdersFutureAlgoSideParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryHistoricalAlgoOrdersFutureAlgoSideParameter(val *QueryHistoricalAlgoOrdersFutureAlgoSideParameter) *NullableQueryHistoricalAlgoOrdersFutureAlgoSideParameter {
	return &NullableQueryHistoricalAlgoOrdersFutureAlgoSideParameter{value: val, isSet: true}
}

func (v NullableQueryHistoricalAlgoOrdersFutureAlgoSideParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryHistoricalAlgoOrdersFutureAlgoSideParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
