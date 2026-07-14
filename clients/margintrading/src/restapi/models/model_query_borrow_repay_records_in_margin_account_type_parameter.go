/*
Margin REST API

Access account information, borrow and repay assets, and trade with Binance Margin.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// QueryBorrowRepayRecordsInMarginAccountTypeParameter the model 'QueryBorrowRepayRecordsInMarginAccountTypeParameter'
type QueryBorrowRepayRecordsInMarginAccountTypeParameter string

// List of queryBorrowRepayRecordsInMarginAccount_type_parameter
const (
	QueryBorrowRepayRecordsInMarginAccountTypeParameterBorrow QueryBorrowRepayRecordsInMarginAccountTypeParameter = "BORROW"
	QueryBorrowRepayRecordsInMarginAccountTypeParameterRepay  QueryBorrowRepayRecordsInMarginAccountTypeParameter = "REPAY"
)

// All allowed values of QueryBorrowRepayRecordsInMarginAccountTypeParameter enum
var AllowedQueryBorrowRepayRecordsInMarginAccountTypeParameterEnumValues = []QueryBorrowRepayRecordsInMarginAccountTypeParameter{
	"BORROW",
	"REPAY",
}

func (v *QueryBorrowRepayRecordsInMarginAccountTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := QueryBorrowRepayRecordsInMarginAccountTypeParameter(value)
	for _, existing := range AllowedQueryBorrowRepayRecordsInMarginAccountTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid QueryBorrowRepayRecordsInMarginAccountTypeParameter", value)
}

// NewQueryBorrowRepayRecordsInMarginAccountTypeParameterFromValue returns a pointer to a valid QueryBorrowRepayRecordsInMarginAccountTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQueryBorrowRepayRecordsInMarginAccountTypeParameterFromValue(v string) (*QueryBorrowRepayRecordsInMarginAccountTypeParameter, error) {
	ev := QueryBorrowRepayRecordsInMarginAccountTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QueryBorrowRepayRecordsInMarginAccountTypeParameter: valid values are %v", v, AllowedQueryBorrowRepayRecordsInMarginAccountTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QueryBorrowRepayRecordsInMarginAccountTypeParameter) IsValid() bool {
	for _, existing := range AllowedQueryBorrowRepayRecordsInMarginAccountTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to queryBorrowRepayRecordsInMarginAccount_type_parameter value
func (v QueryBorrowRepayRecordsInMarginAccountTypeParameter) Ptr() *QueryBorrowRepayRecordsInMarginAccountTypeParameter {
	return &v
}

type NullableQueryBorrowRepayRecordsInMarginAccountTypeParameter struct {
	value *QueryBorrowRepayRecordsInMarginAccountTypeParameter
	isSet bool
}

func (v NullableQueryBorrowRepayRecordsInMarginAccountTypeParameter) Get() *QueryBorrowRepayRecordsInMarginAccountTypeParameter {
	return v.value
}

func (v *NullableQueryBorrowRepayRecordsInMarginAccountTypeParameter) Set(val *QueryBorrowRepayRecordsInMarginAccountTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryBorrowRepayRecordsInMarginAccountTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryBorrowRepayRecordsInMarginAccountTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryBorrowRepayRecordsInMarginAccountTypeParameter(val *QueryBorrowRepayRecordsInMarginAccountTypeParameter) *NullableQueryBorrowRepayRecordsInMarginAccountTypeParameter {
	return &NullableQueryBorrowRepayRecordsInMarginAccountTypeParameter{value: val, isSet: true}
}

func (v NullableQueryBorrowRepayRecordsInMarginAccountTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryBorrowRepayRecordsInMarginAccountTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
