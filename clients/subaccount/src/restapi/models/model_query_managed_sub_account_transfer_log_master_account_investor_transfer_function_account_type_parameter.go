/*
Sub Account REST API

Create and manage sub-accounts, control permissions, and transfer assets via the Sub Account API.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter the model 'QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter'
type QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter string

// List of queryManagedSubAccountTransferLogMasterAccountInvestor_transferFunctionAccountType_parameter
const (
	QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameterSpot           QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter = "SPOT"
	QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameterMargin         QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter = "MARGIN"
	QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameterIsolatedMargin QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter = "ISOLATED_MARGIN"
	QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameterUsdtFuture     QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter = "USDT_FUTURE"
	QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameterCoinFuture     QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter = "COIN_FUTURE"
)

// All allowed values of QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter enum
var AllowedQueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameterEnumValues = []QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter{
	"SPOT",
	"MARGIN",
	"ISOLATED_MARGIN",
	"USDT_FUTURE",
	"COIN_FUTURE",
}

func (v *QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter(value)
	for _, existing := range AllowedQueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter", value)
}

// NewQueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameterFromValue returns a pointer to a valid QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameterFromValue(v string) (*QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter, error) {
	ev := QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter: valid values are %v", v, AllowedQueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter) IsValid() bool {
	for _, existing := range AllowedQueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to queryManagedSubAccountTransferLogMasterAccountInvestor_transferFunctionAccountType_parameter value
func (v QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter) Ptr() *QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter {
	return &v
}

type NullableQueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter struct {
	value *QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter
	isSet bool
}

func (v NullableQueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter) Get() *QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter {
	return v.value
}

func (v *NullableQueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter) Set(val *QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter(val *QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter) *NullableQueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter {
	return &NullableQueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter{value: val, isSet: true}
}

func (v NullableQueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
