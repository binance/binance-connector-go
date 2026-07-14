/*
Wallet REST API

Query balances, manage assets, and perform wallet operations via the Binance Wallet API.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// QueryUserUniversalTransferHistoryFromSymbolParameter the model 'QueryUserUniversalTransferHistoryFromSymbolParameter'
type QueryUserUniversalTransferHistoryFromSymbolParameter string

// List of queryUserUniversalTransferHistory_fromSymbol_parameter
const (
	QueryUserUniversalTransferHistoryFromSymbolParameterIsolatedmarginMargin         QueryUserUniversalTransferHistoryFromSymbolParameter = "ISOLATEDMARGIN_MARGIN"
	QueryUserUniversalTransferHistoryFromSymbolParameterIsolatedmarginIsolatedmargin QueryUserUniversalTransferHistoryFromSymbolParameter = "ISOLATEDMARGIN_ISOLATEDMARGIN"
)

// All allowed values of QueryUserUniversalTransferHistoryFromSymbolParameter enum
var AllowedQueryUserUniversalTransferHistoryFromSymbolParameterEnumValues = []QueryUserUniversalTransferHistoryFromSymbolParameter{
	"ISOLATEDMARGIN_MARGIN",
	"ISOLATEDMARGIN_ISOLATEDMARGIN",
}

func (v *QueryUserUniversalTransferHistoryFromSymbolParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := QueryUserUniversalTransferHistoryFromSymbolParameter(value)
	for _, existing := range AllowedQueryUserUniversalTransferHistoryFromSymbolParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid QueryUserUniversalTransferHistoryFromSymbolParameter", value)
}

// NewQueryUserUniversalTransferHistoryFromSymbolParameterFromValue returns a pointer to a valid QueryUserUniversalTransferHistoryFromSymbolParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQueryUserUniversalTransferHistoryFromSymbolParameterFromValue(v string) (*QueryUserUniversalTransferHistoryFromSymbolParameter, error) {
	ev := QueryUserUniversalTransferHistoryFromSymbolParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QueryUserUniversalTransferHistoryFromSymbolParameter: valid values are %v", v, AllowedQueryUserUniversalTransferHistoryFromSymbolParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QueryUserUniversalTransferHistoryFromSymbolParameter) IsValid() bool {
	for _, existing := range AllowedQueryUserUniversalTransferHistoryFromSymbolParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to queryUserUniversalTransferHistory_fromSymbol_parameter value
func (v QueryUserUniversalTransferHistoryFromSymbolParameter) Ptr() *QueryUserUniversalTransferHistoryFromSymbolParameter {
	return &v
}

type NullableQueryUserUniversalTransferHistoryFromSymbolParameter struct {
	value *QueryUserUniversalTransferHistoryFromSymbolParameter
	isSet bool
}

func (v NullableQueryUserUniversalTransferHistoryFromSymbolParameter) Get() *QueryUserUniversalTransferHistoryFromSymbolParameter {
	return v.value
}

func (v *NullableQueryUserUniversalTransferHistoryFromSymbolParameter) Set(val *QueryUserUniversalTransferHistoryFromSymbolParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryUserUniversalTransferHistoryFromSymbolParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryUserUniversalTransferHistoryFromSymbolParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryUserUniversalTransferHistoryFromSymbolParameter(val *QueryUserUniversalTransferHistoryFromSymbolParameter) *NullableQueryUserUniversalTransferHistoryFromSymbolParameter {
	return &NullableQueryUserUniversalTransferHistoryFromSymbolParameter{value: val, isSet: true}
}

func (v NullableQueryUserUniversalTransferHistoryFromSymbolParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryUserUniversalTransferHistoryFromSymbolParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
