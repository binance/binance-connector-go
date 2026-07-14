/*
Futures (COIN-M) REST API

Access market data, manage accounts, and trade COIN-M perpetual and delivery futures.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// GetIncomeHistoryIncomeTypeParameter the model 'GetIncomeHistoryIncomeTypeParameter'
type GetIncomeHistoryIncomeTypeParameter string

// List of getIncomeHistory_incomeType_parameter
const (
	GetIncomeHistoryIncomeTypeParameterTransfer            GetIncomeHistoryIncomeTypeParameter = "TRANSFER"
	GetIncomeHistoryIncomeTypeParameterWelcomeBonus        GetIncomeHistoryIncomeTypeParameter = "WELCOME_BONUS"
	GetIncomeHistoryIncomeTypeParameterFundingFee          GetIncomeHistoryIncomeTypeParameter = "FUNDING_FEE"
	GetIncomeHistoryIncomeTypeParameterRealizedPnl         GetIncomeHistoryIncomeTypeParameter = "REALIZED_PNL"
	GetIncomeHistoryIncomeTypeParameterCommission          GetIncomeHistoryIncomeTypeParameter = "COMMISSION"
	GetIncomeHistoryIncomeTypeParameterInsuranceClear      GetIncomeHistoryIncomeTypeParameter = "INSURANCE_CLEAR"
	GetIncomeHistoryIncomeTypeParameterDeliveredSettelment GetIncomeHistoryIncomeTypeParameter = "DELIVERED_SETTELMENT"
)

// All allowed values of GetIncomeHistoryIncomeTypeParameter enum
var AllowedGetIncomeHistoryIncomeTypeParameterEnumValues = []GetIncomeHistoryIncomeTypeParameter{
	"TRANSFER",
	"WELCOME_BONUS",
	"FUNDING_FEE",
	"REALIZED_PNL",
	"COMMISSION",
	"INSURANCE_CLEAR",
	"DELIVERED_SETTELMENT",
}

func (v *GetIncomeHistoryIncomeTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GetIncomeHistoryIncomeTypeParameter(value)
	for _, existing := range AllowedGetIncomeHistoryIncomeTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GetIncomeHistoryIncomeTypeParameter", value)
}

// NewGetIncomeHistoryIncomeTypeParameterFromValue returns a pointer to a valid GetIncomeHistoryIncomeTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGetIncomeHistoryIncomeTypeParameterFromValue(v string) (*GetIncomeHistoryIncomeTypeParameter, error) {
	ev := GetIncomeHistoryIncomeTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GetIncomeHistoryIncomeTypeParameter: valid values are %v", v, AllowedGetIncomeHistoryIncomeTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GetIncomeHistoryIncomeTypeParameter) IsValid() bool {
	for _, existing := range AllowedGetIncomeHistoryIncomeTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to getIncomeHistory_incomeType_parameter value
func (v GetIncomeHistoryIncomeTypeParameter) Ptr() *GetIncomeHistoryIncomeTypeParameter {
	return &v
}

type NullableGetIncomeHistoryIncomeTypeParameter struct {
	value *GetIncomeHistoryIncomeTypeParameter
	isSet bool
}

func (v NullableGetIncomeHistoryIncomeTypeParameter) Get() *GetIncomeHistoryIncomeTypeParameter {
	return v.value
}

func (v *NullableGetIncomeHistoryIncomeTypeParameter) Set(val *GetIncomeHistoryIncomeTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIncomeHistoryIncomeTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIncomeHistoryIncomeTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIncomeHistoryIncomeTypeParameter(val *GetIncomeHistoryIncomeTypeParameter) *NullableGetIncomeHistoryIncomeTypeParameter {
	return &NullableGetIncomeHistoryIncomeTypeParameter{value: val, isSet: true}
}

func (v NullableGetIncomeHistoryIncomeTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIncomeHistoryIncomeTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
