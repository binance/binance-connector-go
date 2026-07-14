/*
Options REST API

Access market data, manage accounts, and trade Binance Options.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// AccountFundingFlowCurrencyParameter the model 'AccountFundingFlowCurrencyParameter'
type AccountFundingFlowCurrencyParameter string

// List of accountFundingFlow_currency_parameter
const (
	AccountFundingFlowCurrencyParameterUsdt AccountFundingFlowCurrencyParameter = "USDT"
)

// All allowed values of AccountFundingFlowCurrencyParameter enum
var AllowedAccountFundingFlowCurrencyParameterEnumValues = []AccountFundingFlowCurrencyParameter{
	"USDT",
}

func (v *AccountFundingFlowCurrencyParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccountFundingFlowCurrencyParameter(value)
	for _, existing := range AllowedAccountFundingFlowCurrencyParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccountFundingFlowCurrencyParameter", value)
}

// NewAccountFundingFlowCurrencyParameterFromValue returns a pointer to a valid AccountFundingFlowCurrencyParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccountFundingFlowCurrencyParameterFromValue(v string) (*AccountFundingFlowCurrencyParameter, error) {
	ev := AccountFundingFlowCurrencyParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccountFundingFlowCurrencyParameter: valid values are %v", v, AllowedAccountFundingFlowCurrencyParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccountFundingFlowCurrencyParameter) IsValid() bool {
	for _, existing := range AllowedAccountFundingFlowCurrencyParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to accountFundingFlow_currency_parameter value
func (v AccountFundingFlowCurrencyParameter) Ptr() *AccountFundingFlowCurrencyParameter {
	return &v
}

type NullableAccountFundingFlowCurrencyParameter struct {
	value *AccountFundingFlowCurrencyParameter
	isSet bool
}

func (v NullableAccountFundingFlowCurrencyParameter) Get() *AccountFundingFlowCurrencyParameter {
	return v.value
}

func (v *NullableAccountFundingFlowCurrencyParameter) Set(val *AccountFundingFlowCurrencyParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountFundingFlowCurrencyParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountFundingFlowCurrencyParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountFundingFlowCurrencyParameter(val *AccountFundingFlowCurrencyParameter) *NullableAccountFundingFlowCurrencyParameter {
	return &NullableAccountFundingFlowCurrencyParameter{value: val, isSet: true}
}

func (v NullableAccountFundingFlowCurrencyParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountFundingFlowCurrencyParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
