/*
Binance Derivatives Trading Options WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Options WebSocket Market Streams
*/

package models

import (
	"encoding/json"
	"fmt"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// UserDataStreamEventsResponse - struct for UserDataStreamEventsResponse
type UserDataStreamEventsResponse struct {
	BalancePositionUpdate *BalancePositionUpdate
	GreekUpdate           *GreekUpdate
	Listenkeyexpired      *Listenkeyexpired
	OrderTradeUpdate      *OrderTradeUpdate
	RiskLevelChange       *RiskLevelChange
}

// BalancePositionUpdateAsUserDataStreamEventsResponse is a convenience function that returns BalancePositionUpdate wrapped in UserDataStreamEventsResponse
func BalancePositionUpdateAsUserDataStreamEventsResponse(v *BalancePositionUpdate) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		BalancePositionUpdate: v,
	}
}

// GreekUpdateAsUserDataStreamEventsResponse is a convenience function that returns GreekUpdate wrapped in UserDataStreamEventsResponse
func GreekUpdateAsUserDataStreamEventsResponse(v *GreekUpdate) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		GreekUpdate: v,
	}
}

// ListenkeyexpiredAsUserDataStreamEventsResponse is a convenience function that returns Listenkeyexpired wrapped in UserDataStreamEventsResponse
func ListenkeyexpiredAsUserDataStreamEventsResponse(v *Listenkeyexpired) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		Listenkeyexpired: v,
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

	// check if the discriminator value is 'BALANCE_POSITION_UPDATE'
	if jsonDict["e"] == "BALANCE_POSITION_UPDATE" {
		// try to unmarshal JSON data into BalancePositionUpdate
		err = json.Unmarshal(cleanedData, &dst.BalancePositionUpdate)
		if err == nil {
			return nil // data stored in dst.BalancePositionUpdate, return on the first match
		} else {
			dst.BalancePositionUpdate = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as BalancePositionUpdate: %s", err.Error())
		}
	}

	// check if the discriminator value is 'GREEK_UPDATE'
	if jsonDict["e"] == "GREEK_UPDATE" {
		// try to unmarshal JSON data into GreekUpdate
		err = json.Unmarshal(cleanedData, &dst.GreekUpdate)
		if err == nil {
			return nil // data stored in dst.GreekUpdate, return on the first match
		} else {
			dst.GreekUpdate = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as GreekUpdate: %s", err.Error())
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

	// check if the discriminator value is 'listenKeyExpired'
	if jsonDict["e"] == "listenKeyExpired" {
		// try to unmarshal JSON data into Listenkeyexpired
		err = json.Unmarshal(cleanedData, &dst.Listenkeyexpired)
		if err == nil {
			return nil // data stored in dst.Listenkeyexpired, return on the first match
		} else {
			dst.Listenkeyexpired = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as Listenkeyexpired: %s", err.Error())
		}
	}

	// check if the discriminator value is 'balancePositionUpdate'
	if jsonDict["e"] == "balancePositionUpdate" {
		// try to unmarshal JSON data into BalancePositionUpdate
		err = json.Unmarshal(cleanedData, &dst.BalancePositionUpdate)
		if err == nil {
			return nil // data stored in dst.BalancePositionUpdate, return on the first match
		} else {
			dst.BalancePositionUpdate = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as BalancePositionUpdate: %s", err.Error())
		}
	}

	// check if the discriminator value is 'greekUpdate'
	if jsonDict["e"] == "greekUpdate" {
		// try to unmarshal JSON data into GreekUpdate
		err = json.Unmarshal(cleanedData, &dst.GreekUpdate)
		if err == nil {
			return nil // data stored in dst.GreekUpdate, return on the first match
		} else {
			dst.GreekUpdate = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as GreekUpdate: %s", err.Error())
		}
	}

	// check if the discriminator value is 'listenkeyexpired'
	if jsonDict["e"] == "listenkeyexpired" {
		// try to unmarshal JSON data into Listenkeyexpired
		err = json.Unmarshal(cleanedData, &dst.Listenkeyexpired)
		if err == nil {
			return nil // data stored in dst.Listenkeyexpired, return on the first match
		} else {
			dst.Listenkeyexpired = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as Listenkeyexpired: %s", err.Error())
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
	if src.BalancePositionUpdate != nil {
		return json.Marshal(&src.BalancePositionUpdate)
	}

	if src.GreekUpdate != nil {
		return json.Marshal(&src.GreekUpdate)
	}

	if src.Listenkeyexpired != nil {
		return json.Marshal(&src.Listenkeyexpired)
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
	if obj.BalancePositionUpdate != nil {
		return obj.BalancePositionUpdate
	}

	if obj.GreekUpdate != nil {
		return obj.GreekUpdate
	}

	if obj.Listenkeyexpired != nil {
		return obj.Listenkeyexpired
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
