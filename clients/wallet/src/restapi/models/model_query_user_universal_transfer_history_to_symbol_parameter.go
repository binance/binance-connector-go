/*
Wallet REST API

Query balances, manage assets, and perform wallet operations via the Binance Wallet API.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// QueryUserUniversalTransferHistoryToSymbolParameter the model 'QueryUserUniversalTransferHistoryToSymbolParameter'
type QueryUserUniversalTransferHistoryToSymbolParameter string

// List of queryUserUniversalTransferHistory_toSymbol_parameter
const (
	QueryUserUniversalTransferHistoryToSymbolParameterMarginIsolatedmargin         QueryUserUniversalTransferHistoryToSymbolParameter = "MARGIN_ISOLATEDMARGIN"
	QueryUserUniversalTransferHistoryToSymbolParameterIsolatedmarginIsolatedmargin QueryUserUniversalTransferHistoryToSymbolParameter = "ISOLATEDMARGIN_ISOLATEDMARGIN"
)

// All allowed values of QueryUserUniversalTransferHistoryToSymbolParameter enum
var AllowedQueryUserUniversalTransferHistoryToSymbolParameterEnumValues = []QueryUserUniversalTransferHistoryToSymbolParameter{
	"MARGIN_ISOLATEDMARGIN",
	"ISOLATEDMARGIN_ISOLATEDMARGIN",
}

func (v *QueryUserUniversalTransferHistoryToSymbolParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := QueryUserUniversalTransferHistoryToSymbolParameter(value)
	for _, existing := range AllowedQueryUserUniversalTransferHistoryToSymbolParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid QueryUserUniversalTransferHistoryToSymbolParameter", value)
}

// NewQueryUserUniversalTransferHistoryToSymbolParameterFromValue returns a pointer to a valid QueryUserUniversalTransferHistoryToSymbolParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQueryUserUniversalTransferHistoryToSymbolParameterFromValue(v string) (*QueryUserUniversalTransferHistoryToSymbolParameter, error) {
	ev := QueryUserUniversalTransferHistoryToSymbolParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QueryUserUniversalTransferHistoryToSymbolParameter: valid values are %v", v, AllowedQueryUserUniversalTransferHistoryToSymbolParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QueryUserUniversalTransferHistoryToSymbolParameter) IsValid() bool {
	for _, existing := range AllowedQueryUserUniversalTransferHistoryToSymbolParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to queryUserUniversalTransferHistory_toSymbol_parameter value
func (v QueryUserUniversalTransferHistoryToSymbolParameter) Ptr() *QueryUserUniversalTransferHistoryToSymbolParameter {
	return &v
}

type NullableQueryUserUniversalTransferHistoryToSymbolParameter struct {
	value *QueryUserUniversalTransferHistoryToSymbolParameter
	isSet bool
}

func (v NullableQueryUserUniversalTransferHistoryToSymbolParameter) Get() *QueryUserUniversalTransferHistoryToSymbolParameter {
	return v.value
}

func (v *NullableQueryUserUniversalTransferHistoryToSymbolParameter) Set(val *QueryUserUniversalTransferHistoryToSymbolParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryUserUniversalTransferHistoryToSymbolParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryUserUniversalTransferHistoryToSymbolParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryUserUniversalTransferHistoryToSymbolParameter(val *QueryUserUniversalTransferHistoryToSymbolParameter) *NullableQueryUserUniversalTransferHistoryToSymbolParameter {
	return &NullableQueryUserUniversalTransferHistoryToSymbolParameter{value: val, isSet: true}
}

func (v NullableQueryUserUniversalTransferHistoryToSymbolParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryUserUniversalTransferHistoryToSymbolParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
