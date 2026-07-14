/*
Portfolio Margin Pro REST API

Access advanced account management and high-frequency trading with Binance Portfolio Margin Pro.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// BnbTransferTransferSideParameter the model 'BnbTransferTransferSideParameter'
type BnbTransferTransferSideParameter string

// List of bnbTransfer_transferSide_parameter
const (
	BnbTransferTransferSideParameterToUm   BnbTransferTransferSideParameter = "TO_UM"
	BnbTransferTransferSideParameterFromUm BnbTransferTransferSideParameter = "FROM_UM"
)

// All allowed values of BnbTransferTransferSideParameter enum
var AllowedBnbTransferTransferSideParameterEnumValues = []BnbTransferTransferSideParameter{
	"TO_UM",
	"FROM_UM",
}

func (v *BnbTransferTransferSideParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BnbTransferTransferSideParameter(value)
	for _, existing := range AllowedBnbTransferTransferSideParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BnbTransferTransferSideParameter", value)
}

// NewBnbTransferTransferSideParameterFromValue returns a pointer to a valid BnbTransferTransferSideParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBnbTransferTransferSideParameterFromValue(v string) (*BnbTransferTransferSideParameter, error) {
	ev := BnbTransferTransferSideParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BnbTransferTransferSideParameter: valid values are %v", v, AllowedBnbTransferTransferSideParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BnbTransferTransferSideParameter) IsValid() bool {
	for _, existing := range AllowedBnbTransferTransferSideParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to bnbTransfer_transferSide_parameter value
func (v BnbTransferTransferSideParameter) Ptr() *BnbTransferTransferSideParameter {
	return &v
}

type NullableBnbTransferTransferSideParameter struct {
	value *BnbTransferTransferSideParameter
	isSet bool
}

func (v NullableBnbTransferTransferSideParameter) Get() *BnbTransferTransferSideParameter {
	return v.value
}

func (v *NullableBnbTransferTransferSideParameter) Set(val *BnbTransferTransferSideParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableBnbTransferTransferSideParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableBnbTransferTransferSideParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBnbTransferTransferSideParameter(val *BnbTransferTransferSideParameter) *NullableBnbTransferTransferSideParameter {
	return &NullableBnbTransferTransferSideParameter{value: val, isSet: true}
}

func (v NullableBnbTransferTransferSideParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBnbTransferTransferSideParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
