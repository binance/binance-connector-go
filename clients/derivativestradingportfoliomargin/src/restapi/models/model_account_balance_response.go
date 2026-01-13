/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"
	"fmt"
)

// AccountBalanceResponse - struct for AccountBalanceResponse
type AccountBalanceResponse struct {
	AccountBalanceResponse1 *AccountBalanceResponse1
	AccountBalanceResponse2 *AccountBalanceResponse2
}

// AccountBalanceResponse1AsAccountBalanceResponse is a convenience function that returns AccountBalanceResponse1 wrapped in AccountBalanceResponse
func AccountBalanceResponse1AsAccountBalanceResponse(v *AccountBalanceResponse1) AccountBalanceResponse {
	return AccountBalanceResponse{
		AccountBalanceResponse1: v,
	}
}

// AccountBalanceResponse2AsAccountBalanceResponse is a convenience function that returns AccountBalanceResponse2 wrapped in AccountBalanceResponse
func AccountBalanceResponse2AsAccountBalanceResponse(v *AccountBalanceResponse2) AccountBalanceResponse {
	return AccountBalanceResponse{
		AccountBalanceResponse2: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AccountBalanceResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AccountBalanceResponse1
	err = json.Unmarshal(data, &dst.AccountBalanceResponse1)
	if err == nil {
		jsonAccountBalanceResponse1, _ := json.Marshal(dst.AccountBalanceResponse1)
		if string(jsonAccountBalanceResponse1) == "{}" { // empty struct
			dst.AccountBalanceResponse1 = nil
		} else {
			match++
		}
	} else {
		dst.AccountBalanceResponse1 = nil
	}

	// try to unmarshal data into AccountBalanceResponse2
	err = json.Unmarshal(data, &dst.AccountBalanceResponse2)
	if err == nil {
		jsonAccountBalanceResponse2, _ := json.Marshal(dst.AccountBalanceResponse2)
		if string(jsonAccountBalanceResponse2) == "{}" { // empty struct
			dst.AccountBalanceResponse2 = nil
		} else {
			match++
		}
	} else {
		dst.AccountBalanceResponse2 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AccountBalanceResponse1 = nil
		dst.AccountBalanceResponse2 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AccountBalanceResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AccountBalanceResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AccountBalanceResponse) MarshalJSON() ([]byte, error) {
	if src.AccountBalanceResponse1 != nil {
		return json.Marshal(&src.AccountBalanceResponse1)
	}

	if src.AccountBalanceResponse2 != nil {
		return json.Marshal(&src.AccountBalanceResponse2)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AccountBalanceResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AccountBalanceResponse1 != nil {
		return obj.AccountBalanceResponse1
	}

	if obj.AccountBalanceResponse2 != nil {
		return obj.AccountBalanceResponse2
	}

	// all schemas are nil
	return nil
}

type NullableAccountBalanceResponse struct {
	value *AccountBalanceResponse
	isSet bool
}

func (v NullableAccountBalanceResponse) Get() *AccountBalanceResponse {
	return v.value
}

func (v *NullableAccountBalanceResponse) Set(val *AccountBalanceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountBalanceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountBalanceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountBalanceResponse(val *AccountBalanceResponse) *NullableAccountBalanceResponse {
	return &NullableAccountBalanceResponse{value: val, isSet: true}
}

func (v NullableAccountBalanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountBalanceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
