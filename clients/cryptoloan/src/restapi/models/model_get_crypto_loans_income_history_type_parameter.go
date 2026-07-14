/*
Crypto Loan REST API

Access Binance Crypto Loans to query assets, subscribe to loans, and manage loan positions.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// GetCryptoLoansIncomeHistoryTypeParameter the model 'GetCryptoLoansIncomeHistoryTypeParameter'
type GetCryptoLoansIncomeHistoryTypeParameter string

// List of getCryptoLoansIncomeHistory_type_parameter
const (
	GetCryptoLoansIncomeHistoryTypeParameterBorrowin                         GetCryptoLoansIncomeHistoryTypeParameter = "borrowIn"
	GetCryptoLoansIncomeHistoryTypeParameterCollateralspent                  GetCryptoLoansIncomeHistoryTypeParameter = "collateralSpent"
	GetCryptoLoansIncomeHistoryTypeParameterRepayamount                      GetCryptoLoansIncomeHistoryTypeParameter = "repayAmount"
	GetCryptoLoansIncomeHistoryTypeParameterCollateralreturn                 GetCryptoLoansIncomeHistoryTypeParameter = "collateralReturn"
	GetCryptoLoansIncomeHistoryTypeParameterAddcollateral                    GetCryptoLoansIncomeHistoryTypeParameter = "addCollateral"
	GetCryptoLoansIncomeHistoryTypeParameterRemovecollateral                 GetCryptoLoansIncomeHistoryTypeParameter = "removeCollateral"
	GetCryptoLoansIncomeHistoryTypeParameterCollateralreturnafterliquidation GetCryptoLoansIncomeHistoryTypeParameter = "collateralReturnAfterLiquidation"
)

// All allowed values of GetCryptoLoansIncomeHistoryTypeParameter enum
var AllowedGetCryptoLoansIncomeHistoryTypeParameterEnumValues = []GetCryptoLoansIncomeHistoryTypeParameter{
	"borrowIn",
	"collateralSpent",
	"repayAmount",
	"collateralReturn",
	"addCollateral",
	"removeCollateral",
	"collateralReturnAfterLiquidation",
}

func (v *GetCryptoLoansIncomeHistoryTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GetCryptoLoansIncomeHistoryTypeParameter(value)
	for _, existing := range AllowedGetCryptoLoansIncomeHistoryTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GetCryptoLoansIncomeHistoryTypeParameter", value)
}

// NewGetCryptoLoansIncomeHistoryTypeParameterFromValue returns a pointer to a valid GetCryptoLoansIncomeHistoryTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGetCryptoLoansIncomeHistoryTypeParameterFromValue(v string) (*GetCryptoLoansIncomeHistoryTypeParameter, error) {
	ev := GetCryptoLoansIncomeHistoryTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GetCryptoLoansIncomeHistoryTypeParameter: valid values are %v", v, AllowedGetCryptoLoansIncomeHistoryTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GetCryptoLoansIncomeHistoryTypeParameter) IsValid() bool {
	for _, existing := range AllowedGetCryptoLoansIncomeHistoryTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to getCryptoLoansIncomeHistory_type_parameter value
func (v GetCryptoLoansIncomeHistoryTypeParameter) Ptr() *GetCryptoLoansIncomeHistoryTypeParameter {
	return &v
}

type NullableGetCryptoLoansIncomeHistoryTypeParameter struct {
	value *GetCryptoLoansIncomeHistoryTypeParameter
	isSet bool
}

func (v NullableGetCryptoLoansIncomeHistoryTypeParameter) Get() *GetCryptoLoansIncomeHistoryTypeParameter {
	return v.value
}

func (v *NullableGetCryptoLoansIncomeHistoryTypeParameter) Set(val *GetCryptoLoansIncomeHistoryTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCryptoLoansIncomeHistoryTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCryptoLoansIncomeHistoryTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCryptoLoansIncomeHistoryTypeParameter(val *GetCryptoLoansIncomeHistoryTypeParameter) *NullableGetCryptoLoansIncomeHistoryTypeParameter {
	return &NullableGetCryptoLoansIncomeHistoryTypeParameter{value: val, isSet: true}
}

func (v NullableGetCryptoLoansIncomeHistoryTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCryptoLoansIncomeHistoryTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
