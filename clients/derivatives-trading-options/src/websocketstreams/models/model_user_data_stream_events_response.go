/*
Binance Derivatives Trading Options WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Options WebSocket Market Streams

API version: 1.0.0
*/

package models

import (
	"encoding/json"
	"fmt"

	"github.com/binance/binance-connector-go/common/common"
)

// UserDataStreamEventsResponse - struct for UserDataStreamEventsResponse
type UserDataStreamEventsResponse struct {
	AccountUpdate    *AccountUpdate
	OrderTradeUpdate *OrderTradeUpdate
	RiskLevelChange  *RiskLevelChange
}

// AccountUpdateAsUserDataStreamEventsResponse is a convenience function that returns AccountUpdate wrapped in UserDataStreamEventsResponse
func AccountUpdateAsUserDataStreamEventsResponse(v *AccountUpdate) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		AccountUpdate: v,
	}
}

// OrderTradeUpdateAsUserDataStreamEventsResponse is a convenience function that returns OrderTradeUpdate wrapped in UserDataStreamEventsResponse
func OrderTradeUpdateAsUserDataStreamEventsResponse(v *OrderTradeUpdate) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		OrderTradeUpdate: v,
	}
}

// RiskLevelChangeAsUserDataStreamEventsResponse is a convenience function that returns RiskLevelChange wrapped in UserDataStreamEventsResponse
func RiskLevelChangeAsUserDataStreamEventsResponse(v *RiskLevelChange) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		RiskLevelChange: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UserDataStreamEventsResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = common.NewStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	var modifiedData map[string]interface{}
	if err := json.Unmarshal(data, &modifiedData); err != nil {
		return fmt.Errorf("failed to unmarshal JSON for modification: %v", err)
	}

	// Remove the "e" field
	delete(modifiedData, "e")

	// Marshal the modified data back to JSON
	cleanedData, err := json.Marshal(modifiedData)
	if err != nil {
		return fmt.Errorf("failed to remarshal JSON: %v", err)
	}

	// check if the discriminator value is 'ACCOUNT_UPDATE'
	if jsonDict["e"] == "ACCOUNT_UPDATE" {
		// try to unmarshal JSON data into AccountUpdate
		err = json.Unmarshal(cleanedData, &dst.AccountUpdate)
		if err == nil {
			return nil // data stored in dst.AccountUpdate, return on the first match
		} else {
			dst.AccountUpdate = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as AccountUpdate: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ORDER_TRADE_UPDATE'
	if jsonDict["e"] == "ORDER_TRADE_UPDATE" {
		// try to unmarshal JSON data into OrderTradeUpdate
		err = json.Unmarshal(cleanedData, &dst.OrderTradeUpdate)
		if err == nil {
			return nil // data stored in dst.OrderTradeUpdate, return on the first match
		} else {
			dst.OrderTradeUpdate = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as OrderTradeUpdate: %s", err.Error())
		}
	}

	// check if the discriminator value is 'RISK_LEVEL_CHANGE'
	if jsonDict["e"] == "RISK_LEVEL_CHANGE" {
		// try to unmarshal JSON data into RiskLevelChange
		err = json.Unmarshal(cleanedData, &dst.RiskLevelChange)
		if err == nil {
			return nil // data stored in dst.RiskLevelChange, return on the first match
		} else {
			dst.RiskLevelChange = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as RiskLevelChange: %s", err.Error())
		}
	}

	// check if the discriminator value is 'accountUpdate'
	if jsonDict["e"] == "accountUpdate" {
		// try to unmarshal JSON data into AccountUpdate
		err = json.Unmarshal(cleanedData, &dst.AccountUpdate)
		if err == nil {
			return nil // data stored in dst.AccountUpdate, return on the first match
		} else {
			dst.AccountUpdate = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as AccountUpdate: %s", err.Error())
		}
	}

	// check if the discriminator value is 'orderTradeUpdate'
	if jsonDict["e"] == "orderTradeUpdate" {
		// try to unmarshal JSON data into OrderTradeUpdate
		err = json.Unmarshal(cleanedData, &dst.OrderTradeUpdate)
		if err == nil {
			return nil // data stored in dst.OrderTradeUpdate, return on the first match
		} else {
			dst.OrderTradeUpdate = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as OrderTradeUpdate: %s", err.Error())
		}
	}

	// check if the discriminator value is 'riskLevelChange'
	if jsonDict["e"] == "riskLevelChange" {
		// try to unmarshal JSON data into RiskLevelChange
		err = json.Unmarshal(cleanedData, &dst.RiskLevelChange)
		if err == nil {
			return nil // data stored in dst.RiskLevelChange, return on the first match
		} else {
			dst.RiskLevelChange = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as RiskLevelChange: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UserDataStreamEventsResponse) MarshalJSON() ([]byte, error) {
	if src.AccountUpdate != nil {
		return json.Marshal(&src.AccountUpdate)
	}

	if src.OrderTradeUpdate != nil {
		return json.Marshal(&src.OrderTradeUpdate)
	}

	if src.RiskLevelChange != nil {
		return json.Marshal(&src.RiskLevelChange)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UserDataStreamEventsResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AccountUpdate != nil {
		return obj.AccountUpdate
	}

	if obj.OrderTradeUpdate != nil {
		return obj.OrderTradeUpdate
	}

	if obj.RiskLevelChange != nil {
		return obj.RiskLevelChange
	}

	// all schemas are nil
	return nil
}

type NullableUserDataStreamEventsResponse struct {
	value *UserDataStreamEventsResponse
	isSet bool
}

func (v NullableUserDataStreamEventsResponse) Get() *UserDataStreamEventsResponse {
	return v.value
}

func (v *NullableUserDataStreamEventsResponse) Set(val *UserDataStreamEventsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDataStreamEventsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserDataStreamEventsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDataStreamEventsResponse(val *UserDataStreamEventsResponse) *NullableUserDataStreamEventsResponse {
	return &NullableUserDataStreamEventsResponse{value: val, isSet: true}
}

func (v NullableUserDataStreamEventsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDataStreamEventsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
