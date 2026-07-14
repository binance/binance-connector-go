/*
Margin REST API

Access account information, borrow and repay assets, and trade with Binance Margin.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// QueryMarginAccountsOpenOrdersIsIsolatedParameter the model 'QueryMarginAccountsOpenOrdersIsIsolatedParameter'
type QueryMarginAccountsOpenOrdersIsIsolatedParameter string

// List of queryMarginAccountsOpenOrders_isIsolated_parameter
const (
	QueryMarginAccountsOpenOrdersIsIsolatedParameterTrue  QueryMarginAccountsOpenOrdersIsIsolatedParameter = "TRUE"
	QueryMarginAccountsOpenOrdersIsIsolatedParameterFalse QueryMarginAccountsOpenOrdersIsIsolatedParameter = "FALSE"
)

// All allowed values of QueryMarginAccountsOpenOrdersIsIsolatedParameter enum
var AllowedQueryMarginAccountsOpenOrdersIsIsolatedParameterEnumValues = []QueryMarginAccountsOpenOrdersIsIsolatedParameter{
	"TRUE",
	"FALSE",
}

func (v *QueryMarginAccountsOpenOrdersIsIsolatedParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := QueryMarginAccountsOpenOrdersIsIsolatedParameter(value)
	for _, existing := range AllowedQueryMarginAccountsOpenOrdersIsIsolatedParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid QueryMarginAccountsOpenOrdersIsIsolatedParameter", value)
}

// NewQueryMarginAccountsOpenOrdersIsIsolatedParameterFromValue returns a pointer to a valid QueryMarginAccountsOpenOrdersIsIsolatedParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQueryMarginAccountsOpenOrdersIsIsolatedParameterFromValue(v string) (*QueryMarginAccountsOpenOrdersIsIsolatedParameter, error) {
	ev := QueryMarginAccountsOpenOrdersIsIsolatedParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QueryMarginAccountsOpenOrdersIsIsolatedParameter: valid values are %v", v, AllowedQueryMarginAccountsOpenOrdersIsIsolatedParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QueryMarginAccountsOpenOrdersIsIsolatedParameter) IsValid() bool {
	for _, existing := range AllowedQueryMarginAccountsOpenOrdersIsIsolatedParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to queryMarginAccountsOpenOrders_isIsolated_parameter value
func (v QueryMarginAccountsOpenOrdersIsIsolatedParameter) Ptr() *QueryMarginAccountsOpenOrdersIsIsolatedParameter {
	return &v
}

type NullableQueryMarginAccountsOpenOrdersIsIsolatedParameter struct {
	value *QueryMarginAccountsOpenOrdersIsIsolatedParameter
	isSet bool
}

func (v NullableQueryMarginAccountsOpenOrdersIsIsolatedParameter) Get() *QueryMarginAccountsOpenOrdersIsIsolatedParameter {
	return v.value
}

func (v *NullableQueryMarginAccountsOpenOrdersIsIsolatedParameter) Set(val *QueryMarginAccountsOpenOrdersIsIsolatedParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryMarginAccountsOpenOrdersIsIsolatedParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryMarginAccountsOpenOrdersIsIsolatedParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryMarginAccountsOpenOrdersIsIsolatedParameter(val *QueryMarginAccountsOpenOrdersIsIsolatedParameter) *NullableQueryMarginAccountsOpenOrdersIsIsolatedParameter {
	return &NullableQueryMarginAccountsOpenOrdersIsIsolatedParameter{value: val, isSet: true}
}

func (v NullableQueryMarginAccountsOpenOrdersIsIsolatedParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryMarginAccountsOpenOrdersIsIsolatedParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
