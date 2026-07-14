/*
Portfolio Margin Pro REST API

Access advanced account management and high-frequency trading with Binance Portfolio Margin Pro.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// TransferLdusdtRwusdForPortfolioMarginAssetParameter the model 'TransferLdusdtRwusdForPortfolioMarginAssetParameter'
type TransferLdusdtRwusdForPortfolioMarginAssetParameter string

// List of transferLdusdtRwusdForPortfolioMargin_asset_parameter
const (
	TransferLdusdtRwusdForPortfolioMarginAssetParameterLdusdt TransferLdusdtRwusdForPortfolioMarginAssetParameter = "LDUSDT"
	TransferLdusdtRwusdForPortfolioMarginAssetParameterRwusd  TransferLdusdtRwusdForPortfolioMarginAssetParameter = "RWUSD"
)

// All allowed values of TransferLdusdtRwusdForPortfolioMarginAssetParameter enum
var AllowedTransferLdusdtRwusdForPortfolioMarginAssetParameterEnumValues = []TransferLdusdtRwusdForPortfolioMarginAssetParameter{
	"LDUSDT",
	"RWUSD",
}

func (v *TransferLdusdtRwusdForPortfolioMarginAssetParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransferLdusdtRwusdForPortfolioMarginAssetParameter(value)
	for _, existing := range AllowedTransferLdusdtRwusdForPortfolioMarginAssetParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TransferLdusdtRwusdForPortfolioMarginAssetParameter", value)
}

// NewTransferLdusdtRwusdForPortfolioMarginAssetParameterFromValue returns a pointer to a valid TransferLdusdtRwusdForPortfolioMarginAssetParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransferLdusdtRwusdForPortfolioMarginAssetParameterFromValue(v string) (*TransferLdusdtRwusdForPortfolioMarginAssetParameter, error) {
	ev := TransferLdusdtRwusdForPortfolioMarginAssetParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransferLdusdtRwusdForPortfolioMarginAssetParameter: valid values are %v", v, AllowedTransferLdusdtRwusdForPortfolioMarginAssetParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransferLdusdtRwusdForPortfolioMarginAssetParameter) IsValid() bool {
	for _, existing := range AllowedTransferLdusdtRwusdForPortfolioMarginAssetParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to transferLdusdtRwusdForPortfolioMargin_asset_parameter value
func (v TransferLdusdtRwusdForPortfolioMarginAssetParameter) Ptr() *TransferLdusdtRwusdForPortfolioMarginAssetParameter {
	return &v
}

type NullableTransferLdusdtRwusdForPortfolioMarginAssetParameter struct {
	value *TransferLdusdtRwusdForPortfolioMarginAssetParameter
	isSet bool
}

func (v NullableTransferLdusdtRwusdForPortfolioMarginAssetParameter) Get() *TransferLdusdtRwusdForPortfolioMarginAssetParameter {
	return v.value
}

func (v *NullableTransferLdusdtRwusdForPortfolioMarginAssetParameter) Set(val *TransferLdusdtRwusdForPortfolioMarginAssetParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferLdusdtRwusdForPortfolioMarginAssetParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferLdusdtRwusdForPortfolioMarginAssetParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferLdusdtRwusdForPortfolioMarginAssetParameter(val *TransferLdusdtRwusdForPortfolioMarginAssetParameter) *NullableTransferLdusdtRwusdForPortfolioMarginAssetParameter {
	return &NullableTransferLdusdtRwusdForPortfolioMarginAssetParameter{value: val, isSet: true}
}

func (v NullableTransferLdusdtRwusdForPortfolioMarginAssetParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferLdusdtRwusdForPortfolioMarginAssetParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
