/*
Portfolio Margin REST API

Access account information, manage margin positions, and trade with Binance Portfolio Margin.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// GetCmIncomeHistoryIncomeTypeParameter the model 'GetCmIncomeHistoryIncomeTypeParameter'
type GetCmIncomeHistoryIncomeTypeParameter string

// List of getCmIncomeHistory_incomeType_parameter
const (
	GetCmIncomeHistoryIncomeTypeParameterTransfer            GetCmIncomeHistoryIncomeTypeParameter = "TRANSFER"
	GetCmIncomeHistoryIncomeTypeParameterWelcomeBonus        GetCmIncomeHistoryIncomeTypeParameter = "WELCOME_BONUS"
	GetCmIncomeHistoryIncomeTypeParameterFundingFee          GetCmIncomeHistoryIncomeTypeParameter = "FUNDING_FEE"
	GetCmIncomeHistoryIncomeTypeParameterRealizedPnl         GetCmIncomeHistoryIncomeTypeParameter = "REALIZED_PNL"
	GetCmIncomeHistoryIncomeTypeParameterCommission          GetCmIncomeHistoryIncomeTypeParameter = "COMMISSION"
	GetCmIncomeHistoryIncomeTypeParameterInsuranceClear      GetCmIncomeHistoryIncomeTypeParameter = "INSURANCE_CLEAR"
	GetCmIncomeHistoryIncomeTypeParameterDeliveredSettelment GetCmIncomeHistoryIncomeTypeParameter = "DELIVERED_SETTELMENT"
)

// All allowed values of GetCmIncomeHistoryIncomeTypeParameter enum
var AllowedGetCmIncomeHistoryIncomeTypeParameterEnumValues = []GetCmIncomeHistoryIncomeTypeParameter{
	"TRANSFER",
	"WELCOME_BONUS",
	"FUNDING_FEE",
	"REALIZED_PNL",
	"COMMISSION",
	"INSURANCE_CLEAR",
	"DELIVERED_SETTELMENT",
}

func (v *GetCmIncomeHistoryIncomeTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GetCmIncomeHistoryIncomeTypeParameter(value)
	for _, existing := range AllowedGetCmIncomeHistoryIncomeTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GetCmIncomeHistoryIncomeTypeParameter", value)
}

// NewGetCmIncomeHistoryIncomeTypeParameterFromValue returns a pointer to a valid GetCmIncomeHistoryIncomeTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGetCmIncomeHistoryIncomeTypeParameterFromValue(v string) (*GetCmIncomeHistoryIncomeTypeParameter, error) {
	ev := GetCmIncomeHistoryIncomeTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GetCmIncomeHistoryIncomeTypeParameter: valid values are %v", v, AllowedGetCmIncomeHistoryIncomeTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GetCmIncomeHistoryIncomeTypeParameter) IsValid() bool {
	for _, existing := range AllowedGetCmIncomeHistoryIncomeTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to getCmIncomeHistory_incomeType_parameter value
func (v GetCmIncomeHistoryIncomeTypeParameter) Ptr() *GetCmIncomeHistoryIncomeTypeParameter {
	return &v
}

type NullableGetCmIncomeHistoryIncomeTypeParameter struct {
	value *GetCmIncomeHistoryIncomeTypeParameter
	isSet bool
}

func (v NullableGetCmIncomeHistoryIncomeTypeParameter) Get() *GetCmIncomeHistoryIncomeTypeParameter {
	return v.value
}

func (v *NullableGetCmIncomeHistoryIncomeTypeParameter) Set(val *GetCmIncomeHistoryIncomeTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCmIncomeHistoryIncomeTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCmIncomeHistoryIncomeTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCmIncomeHistoryIncomeTypeParameter(val *GetCmIncomeHistoryIncomeTypeParameter) *NullableGetCmIncomeHistoryIncomeTypeParameter {
	return &NullableGetCmIncomeHistoryIncomeTypeParameter{value: val, isSet: true}
}

func (v NullableGetCmIncomeHistoryIncomeTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCmIncomeHistoryIncomeTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
