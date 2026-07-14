/*
Portfolio Margin Pro REST API

Access advanced account management and high-frequency trading with Binance Portfolio Margin Pro.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// GetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameter the model 'GetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameter'
type GetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameter string

// List of getTransferableEarnAssetBalanceForPortfolioMargin_transferType_parameter
const (
	GetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameterEarnToFuture GetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameter = "EARN_TO_FUTURE"
	GetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameterFutureToEarn GetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameter = "FUTURE_TO_EARN"
)

// All allowed values of GetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameter enum
var AllowedGetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameterEnumValues = []GetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameter{
	"EARN_TO_FUTURE",
	"FUTURE_TO_EARN",
}

func (v *GetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameter(value)
	for _, existing := range AllowedGetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameter", value)
}

// NewGetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameterFromValue returns a pointer to a valid GetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameterFromValue(v string) (*GetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameter, error) {
	ev := GetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameter: valid values are %v", v, AllowedGetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameter) IsValid() bool {
	for _, existing := range AllowedGetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to getTransferableEarnAssetBalanceForPortfolioMargin_transferType_parameter value
func (v GetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameter) Ptr() *GetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameter {
	return &v
}

type NullableGetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameter struct {
	value *GetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameter
	isSet bool
}

func (v NullableGetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameter) Get() *GetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameter {
	return v.value
}

func (v *NullableGetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameter) Set(val *GetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameter(val *GetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameter) *NullableGetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameter {
	return &NullableGetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameter{value: val, isSet: true}
}

func (v NullableGetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
