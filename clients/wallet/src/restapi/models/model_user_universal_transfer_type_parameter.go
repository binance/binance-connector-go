/*
Wallet REST API

Query balances, manage assets, and perform wallet operations via the Binance Wallet API.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// UserUniversalTransferTypeParameter the model 'UserUniversalTransferTypeParameter'
type UserUniversalTransferTypeParameter string

// List of userUniversalTransfer_type_parameter
const (
	UserUniversalTransferTypeParameterMainUmfuture                 UserUniversalTransferTypeParameter = "MAIN_UMFUTURE"
	UserUniversalTransferTypeParameterMainCmfuture                 UserUniversalTransferTypeParameter = "MAIN_CMFUTURE"
	UserUniversalTransferTypeParameterMainMargin                   UserUniversalTransferTypeParameter = "MAIN_MARGIN"
	UserUniversalTransferTypeParameterUmfutureMain                 UserUniversalTransferTypeParameter = "UMFUTURE_MAIN"
	UserUniversalTransferTypeParameterUmfutureMargin               UserUniversalTransferTypeParameter = "UMFUTURE_MARGIN"
	UserUniversalTransferTypeParameterCmfutureMain                 UserUniversalTransferTypeParameter = "CMFUTURE_MAIN"
	UserUniversalTransferTypeParameterCmfutureMargin               UserUniversalTransferTypeParameter = "CMFUTURE_MARGIN"
	UserUniversalTransferTypeParameterMarginMain                   UserUniversalTransferTypeParameter = "MARGIN_MAIN"
	UserUniversalTransferTypeParameterMarginUmfuture               UserUniversalTransferTypeParameter = "MARGIN_UMFUTURE"
	UserUniversalTransferTypeParameterMarginCmfuture               UserUniversalTransferTypeParameter = "MARGIN_CMFUTURE"
	UserUniversalTransferTypeParameterIsolatedmarginMargin         UserUniversalTransferTypeParameter = "ISOLATEDMARGIN_MARGIN"
	UserUniversalTransferTypeParameterMarginIsolatedmargin         UserUniversalTransferTypeParameter = "MARGIN_ISOLATEDMARGIN"
	UserUniversalTransferTypeParameterIsolatedmarginIsolatedmargin UserUniversalTransferTypeParameter = "ISOLATEDMARGIN_ISOLATEDMARGIN"
	UserUniversalTransferTypeParameterMainFunding                  UserUniversalTransferTypeParameter = "MAIN_FUNDING"
	UserUniversalTransferTypeParameterFundingMain                  UserUniversalTransferTypeParameter = "FUNDING_MAIN"
	UserUniversalTransferTypeParameterFundingUmfuture              UserUniversalTransferTypeParameter = "FUNDING_UMFUTURE"
	UserUniversalTransferTypeParameterUmfutureFunding              UserUniversalTransferTypeParameter = "UMFUTURE_FUNDING"
	UserUniversalTransferTypeParameterMarginFunding                UserUniversalTransferTypeParameter = "MARGIN_FUNDING"
	UserUniversalTransferTypeParameterFundingMargin                UserUniversalTransferTypeParameter = "FUNDING_MARGIN"
	UserUniversalTransferTypeParameterFundingCmfuture              UserUniversalTransferTypeParameter = "FUNDING_CMFUTURE"
	UserUniversalTransferTypeParameterCmfutureFunding              UserUniversalTransferTypeParameter = "CMFUTURE_FUNDING"
	UserUniversalTransferTypeParameterMainOption                   UserUniversalTransferTypeParameter = "MAIN_OPTION"
	UserUniversalTransferTypeParameterOptionMain                   UserUniversalTransferTypeParameter = "OPTION_MAIN"
	UserUniversalTransferTypeParameterUmfutureOption               UserUniversalTransferTypeParameter = "UMFUTURE_OPTION"
	UserUniversalTransferTypeParameterOptionUmfuture               UserUniversalTransferTypeParameter = "OPTION_UMFUTURE"
	UserUniversalTransferTypeParameterMarginOption                 UserUniversalTransferTypeParameter = "MARGIN_OPTION"
	UserUniversalTransferTypeParameterOptionMargin                 UserUniversalTransferTypeParameter = "OPTION_MARGIN"
	UserUniversalTransferTypeParameterFundingOption                UserUniversalTransferTypeParameter = "FUNDING_OPTION"
	UserUniversalTransferTypeParameterOptionFunding                UserUniversalTransferTypeParameter = "OPTION_FUNDING"
	UserUniversalTransferTypeParameterMainPortfolioMargin          UserUniversalTransferTypeParameter = "MAIN_PORTFOLIO_MARGIN"
	UserUniversalTransferTypeParameterPortfolioMarginMain          UserUniversalTransferTypeParameter = "PORTFOLIO_MARGIN_MAIN"
)

// All allowed values of UserUniversalTransferTypeParameter enum
var AllowedUserUniversalTransferTypeParameterEnumValues = []UserUniversalTransferTypeParameter{
	"MAIN_UMFUTURE",
	"MAIN_CMFUTURE",
	"MAIN_MARGIN",
	"UMFUTURE_MAIN",
	"UMFUTURE_MARGIN",
	"CMFUTURE_MAIN",
	"CMFUTURE_MARGIN",
	"MARGIN_MAIN",
	"MARGIN_UMFUTURE",
	"MARGIN_CMFUTURE",
	"ISOLATEDMARGIN_MARGIN",
	"MARGIN_ISOLATEDMARGIN",
	"ISOLATEDMARGIN_ISOLATEDMARGIN",
	"MAIN_FUNDING",
	"FUNDING_MAIN",
	"FUNDING_UMFUTURE",
	"UMFUTURE_FUNDING",
	"MARGIN_FUNDING",
	"FUNDING_MARGIN",
	"FUNDING_CMFUTURE",
	"CMFUTURE_FUNDING",
	"MAIN_OPTION",
	"OPTION_MAIN",
	"UMFUTURE_OPTION",
	"OPTION_UMFUTURE",
	"MARGIN_OPTION",
	"OPTION_MARGIN",
	"FUNDING_OPTION",
	"OPTION_FUNDING",
	"MAIN_PORTFOLIO_MARGIN",
	"PORTFOLIO_MARGIN_MAIN",
}

func (v *UserUniversalTransferTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UserUniversalTransferTypeParameter(value)
	for _, existing := range AllowedUserUniversalTransferTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UserUniversalTransferTypeParameter", value)
}

// NewUserUniversalTransferTypeParameterFromValue returns a pointer to a valid UserUniversalTransferTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUserUniversalTransferTypeParameterFromValue(v string) (*UserUniversalTransferTypeParameter, error) {
	ev := UserUniversalTransferTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UserUniversalTransferTypeParameter: valid values are %v", v, AllowedUserUniversalTransferTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UserUniversalTransferTypeParameter) IsValid() bool {
	for _, existing := range AllowedUserUniversalTransferTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to userUniversalTransfer_type_parameter value
func (v UserUniversalTransferTypeParameter) Ptr() *UserUniversalTransferTypeParameter {
	return &v
}

type NullableUserUniversalTransferTypeParameter struct {
	value *UserUniversalTransferTypeParameter
	isSet bool
}

func (v NullableUserUniversalTransferTypeParameter) Get() *UserUniversalTransferTypeParameter {
	return v.value
}

func (v *NullableUserUniversalTransferTypeParameter) Set(val *UserUniversalTransferTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableUserUniversalTransferTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableUserUniversalTransferTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserUniversalTransferTypeParameter(val *UserUniversalTransferTypeParameter) *NullableUserUniversalTransferTypeParameter {
	return &NullableUserUniversalTransferTypeParameter{value: val, isSet: true}
}

func (v NullableUserUniversalTransferTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserUniversalTransferTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
