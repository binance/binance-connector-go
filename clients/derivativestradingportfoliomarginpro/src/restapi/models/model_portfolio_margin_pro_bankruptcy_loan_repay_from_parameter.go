/*
Portfolio Margin Pro REST API

Access advanced account management and high-frequency trading with Binance Portfolio Margin Pro.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// PortfolioMarginProBankruptcyLoanRepayFromParameter the model 'PortfolioMarginProBankruptcyLoanRepayFromParameter'
type PortfolioMarginProBankruptcyLoanRepayFromParameter string

// List of portfolioMarginProBankruptcyLoanRepay_from_parameter
const (
	PortfolioMarginProBankruptcyLoanRepayFromParameterSpot   PortfolioMarginProBankruptcyLoanRepayFromParameter = "SPOT"
	PortfolioMarginProBankruptcyLoanRepayFromParameterMargin PortfolioMarginProBankruptcyLoanRepayFromParameter = "MARGIN"
)

// All allowed values of PortfolioMarginProBankruptcyLoanRepayFromParameter enum
var AllowedPortfolioMarginProBankruptcyLoanRepayFromParameterEnumValues = []PortfolioMarginProBankruptcyLoanRepayFromParameter{
	"SPOT",
	"MARGIN",
}

func (v *PortfolioMarginProBankruptcyLoanRepayFromParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PortfolioMarginProBankruptcyLoanRepayFromParameter(value)
	for _, existing := range AllowedPortfolioMarginProBankruptcyLoanRepayFromParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PortfolioMarginProBankruptcyLoanRepayFromParameter", value)
}

// NewPortfolioMarginProBankruptcyLoanRepayFromParameterFromValue returns a pointer to a valid PortfolioMarginProBankruptcyLoanRepayFromParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPortfolioMarginProBankruptcyLoanRepayFromParameterFromValue(v string) (*PortfolioMarginProBankruptcyLoanRepayFromParameter, error) {
	ev := PortfolioMarginProBankruptcyLoanRepayFromParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PortfolioMarginProBankruptcyLoanRepayFromParameter: valid values are %v", v, AllowedPortfolioMarginProBankruptcyLoanRepayFromParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PortfolioMarginProBankruptcyLoanRepayFromParameter) IsValid() bool {
	for _, existing := range AllowedPortfolioMarginProBankruptcyLoanRepayFromParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to portfolioMarginProBankruptcyLoanRepay_from_parameter value
func (v PortfolioMarginProBankruptcyLoanRepayFromParameter) Ptr() *PortfolioMarginProBankruptcyLoanRepayFromParameter {
	return &v
}

type NullablePortfolioMarginProBankruptcyLoanRepayFromParameter struct {
	value *PortfolioMarginProBankruptcyLoanRepayFromParameter
	isSet bool
}

func (v NullablePortfolioMarginProBankruptcyLoanRepayFromParameter) Get() *PortfolioMarginProBankruptcyLoanRepayFromParameter {
	return v.value
}

func (v *NullablePortfolioMarginProBankruptcyLoanRepayFromParameter) Set(val *PortfolioMarginProBankruptcyLoanRepayFromParameter) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioMarginProBankruptcyLoanRepayFromParameter) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioMarginProBankruptcyLoanRepayFromParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioMarginProBankruptcyLoanRepayFromParameter(val *PortfolioMarginProBankruptcyLoanRepayFromParameter) *NullablePortfolioMarginProBankruptcyLoanRepayFromParameter {
	return &NullablePortfolioMarginProBankruptcyLoanRepayFromParameter{value: val, isSet: true}
}

func (v NullablePortfolioMarginProBankruptcyLoanRepayFromParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortfolioMarginProBankruptcyLoanRepayFromParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
