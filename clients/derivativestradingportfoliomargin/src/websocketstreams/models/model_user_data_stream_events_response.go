/*
Portfolio Margin WebSocket Market Streams

Access account information, manage margin positions, and trade with Binance Portfolio Margin.
*/

package models

import (
	"encoding/json"
	"fmt"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// UserDataStreamEventsResponse - struct for UserDataStreamEventsResponse
type UserDataStreamEventsResponse struct {
	AccountConfigUpdate         *AccountConfigUpdate
	AccountUpdate               *AccountUpdate
	AlgoOrderUpdate             *AlgoOrderUpdate
	BalanceUpdate               *BalanceUpdate
	ConditionalOrderTradeUpdate *ConditionalOrderTradeUpdate
	ExecutionReport             *ExecutionReport
	LiabilityChange             *LiabilityChange
	ListenKeyExpired            *ListenKeyExpired
	OpenOrderLoss               *OpenOrderLoss
	OrderTradeUpdate            *OrderTradeUpdate
	OutboundAccountPosition     *OutboundAccountPosition
	RiskLevelChange             *RiskLevelChange
}

// AccountConfigUpdateAsUserDataStreamEventsResponse is a convenience function that returns AccountConfigUpdate wrapped in UserDataStreamEventsResponse
func AccountConfigUpdateAsUserDataStreamEventsResponse(v *AccountConfigUpdate) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		AccountConfigUpdate: v,
	}
}

// AccountUpdateAsUserDataStreamEventsResponse is a convenience function that returns AccountUpdate wrapped in UserDataStreamEventsResponse
func AccountUpdateAsUserDataStreamEventsResponse(v *AccountUpdate) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		AccountUpdate: v,
	}
}

// AlgoOrderUpdateAsUserDataStreamEventsResponse is a convenience function that returns AlgoOrderUpdate wrapped in UserDataStreamEventsResponse
func AlgoOrderUpdateAsUserDataStreamEventsResponse(v *AlgoOrderUpdate) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		AlgoOrderUpdate: v,
	}
}

// BalanceUpdateAsUserDataStreamEventsResponse is a convenience function that returns BalanceUpdate wrapped in UserDataStreamEventsResponse
func BalanceUpdateAsUserDataStreamEventsResponse(v *BalanceUpdate) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		BalanceUpdate: v,
	}
}

// ConditionalOrderTradeUpdateAsUserDataStreamEventsResponse is a convenience function that returns ConditionalOrderTradeUpdate wrapped in UserDataStreamEventsResponse
func ConditionalOrderTradeUpdateAsUserDataStreamEventsResponse(v *ConditionalOrderTradeUpdate) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		ConditionalOrderTradeUpdate: v,
	}
}

// ExecutionReportAsUserDataStreamEventsResponse is a convenience function that returns ExecutionReport wrapped in UserDataStreamEventsResponse
func ExecutionReportAsUserDataStreamEventsResponse(v *ExecutionReport) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		ExecutionReport: v,
	}
}

// LiabilityChangeAsUserDataStreamEventsResponse is a convenience function that returns LiabilityChange wrapped in UserDataStreamEventsResponse
func LiabilityChangeAsUserDataStreamEventsResponse(v *LiabilityChange) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		LiabilityChange: v,
	}
}

// ListenKeyExpiredAsUserDataStreamEventsResponse is a convenience function that returns ListenKeyExpired wrapped in UserDataStreamEventsResponse
func ListenKeyExpiredAsUserDataStreamEventsResponse(v *ListenKeyExpired) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		ListenKeyExpired: v,
	}
}

// OpenOrderLossAsUserDataStreamEventsResponse is a convenience function that returns OpenOrderLoss wrapped in UserDataStreamEventsResponse
func OpenOrderLossAsUserDataStreamEventsResponse(v *OpenOrderLoss) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		OpenOrderLoss: v,
	}
}

// OrderTradeUpdateAsUserDataStreamEventsResponse is a convenience function that returns OrderTradeUpdate wrapped in UserDataStreamEventsResponse
func OrderTradeUpdateAsUserDataStreamEventsResponse(v *OrderTradeUpdate) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		OrderTradeUpdate: v,
	}
}

// OutboundAccountPositionAsUserDataStreamEventsResponse is a convenience function that returns OutboundAccountPosition wrapped in UserDataStreamEventsResponse
func OutboundAccountPositionAsUserDataStreamEventsResponse(v *OutboundAccountPosition) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		OutboundAccountPosition: v,
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

	// check if the discriminator value is 'ACCOUNT_CONFIG_UPDATE'
	if jsonDict["e"] == "ACCOUNT_CONFIG_UPDATE" {
		// try to unmarshal JSON data into AccountConfigUpdate
		err = json.Unmarshal(cleanedData, &dst.AccountConfigUpdate)
		if err == nil {
			return nil // data stored in dst.AccountConfigUpdate, return on the first match
		} else {
			dst.AccountConfigUpdate = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as AccountConfigUpdate: %s", err.Error())
		}
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

	// check if the discriminator value is 'ALGO_UPDATE'
	if jsonDict["e"] == "ALGO_UPDATE" {
		// try to unmarshal JSON data into AlgoOrderUpdate
		err = json.Unmarshal(cleanedData, &dst.AlgoOrderUpdate)
		if err == nil {
			return nil // data stored in dst.AlgoOrderUpdate, return on the first match
		} else {
			dst.AlgoOrderUpdate = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as AlgoOrderUpdate: %s", err.Error())
		}
	}

	// check if the discriminator value is 'CONDITIONAL_ORDER_TRADE_UPDATE'
	if jsonDict["e"] == "CONDITIONAL_ORDER_TRADE_UPDATE" {
		// try to unmarshal JSON data into ConditionalOrderTradeUpdate
		err = json.Unmarshal(cleanedData, &dst.ConditionalOrderTradeUpdate)
		if err == nil {
			return nil // data stored in dst.ConditionalOrderTradeUpdate, return on the first match
		} else {
			dst.ConditionalOrderTradeUpdate = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as ConditionalOrderTradeUpdate: %s", err.Error())
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

	// check if the discriminator value is 'balanceUpdate'
	if jsonDict["e"] == "balanceUpdate" {
		// try to unmarshal JSON data into BalanceUpdate
		err = json.Unmarshal(cleanedData, &dst.BalanceUpdate)
		if err == nil {
			return nil // data stored in dst.BalanceUpdate, return on the first match
		} else {
			dst.BalanceUpdate = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as BalanceUpdate: %s", err.Error())
		}
	}

	// check if the discriminator value is 'executionReport'
	if jsonDict["e"] == "executionReport" {
		// try to unmarshal JSON data into ExecutionReport
		err = json.Unmarshal(cleanedData, &dst.ExecutionReport)
		if err == nil {
			return nil // data stored in dst.ExecutionReport, return on the first match
		} else {
			dst.ExecutionReport = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as ExecutionReport: %s", err.Error())
		}
	}

	// check if the discriminator value is 'liabilityChange'
	if jsonDict["e"] == "liabilityChange" {
		// try to unmarshal JSON data into LiabilityChange
		err = json.Unmarshal(cleanedData, &dst.LiabilityChange)
		if err == nil {
			return nil // data stored in dst.LiabilityChange, return on the first match
		} else {
			dst.LiabilityChange = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as LiabilityChange: %s", err.Error())
		}
	}

	// check if the discriminator value is 'listenKeyExpired'
	if jsonDict["e"] == "listenKeyExpired" {
		// try to unmarshal JSON data into ListenKeyExpired
		err = json.Unmarshal(cleanedData, &dst.ListenKeyExpired)
		if err == nil {
			return nil // data stored in dst.ListenKeyExpired, return on the first match
		} else {
			dst.ListenKeyExpired = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as ListenKeyExpired: %s", err.Error())
		}
	}

	// check if the discriminator value is 'openOrderLoss'
	if jsonDict["e"] == "openOrderLoss" {
		// try to unmarshal JSON data into OpenOrderLoss
		err = json.Unmarshal(cleanedData, &dst.OpenOrderLoss)
		if err == nil {
			return nil // data stored in dst.OpenOrderLoss, return on the first match
		} else {
			dst.OpenOrderLoss = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as OpenOrderLoss: %s", err.Error())
		}
	}

	// check if the discriminator value is 'outboundAccountPosition'
	if jsonDict["e"] == "outboundAccountPosition" {
		// try to unmarshal JSON data into OutboundAccountPosition
		err = json.Unmarshal(cleanedData, &dst.OutboundAccountPosition)
		if err == nil {
			return nil // data stored in dst.OutboundAccountPosition, return on the first match
		} else {
			dst.OutboundAccountPosition = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as OutboundAccountPosition: %s", err.Error())
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
	if src.AccountConfigUpdate != nil {
		return json.Marshal(&src.AccountConfigUpdate)
	}

	if src.AccountUpdate != nil {
		return json.Marshal(&src.AccountUpdate)
	}

	if src.AlgoOrderUpdate != nil {
		return json.Marshal(&src.AlgoOrderUpdate)
	}

	if src.BalanceUpdate != nil {
		return json.Marshal(&src.BalanceUpdate)
	}

	if src.ConditionalOrderTradeUpdate != nil {
		return json.Marshal(&src.ConditionalOrderTradeUpdate)
	}

	if src.ExecutionReport != nil {
		return json.Marshal(&src.ExecutionReport)
	}

	if src.LiabilityChange != nil {
		return json.Marshal(&src.LiabilityChange)
	}

	if src.ListenKeyExpired != nil {
		return json.Marshal(&src.ListenKeyExpired)
	}

	if src.OpenOrderLoss != nil {
		return json.Marshal(&src.OpenOrderLoss)
	}

	if src.OrderTradeUpdate != nil {
		return json.Marshal(&src.OrderTradeUpdate)
	}

	if src.OutboundAccountPosition != nil {
		return json.Marshal(&src.OutboundAccountPosition)
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
	if obj.AccountConfigUpdate != nil {
		return obj.AccountConfigUpdate
	}

	if obj.AccountUpdate != nil {
		return obj.AccountUpdate
	}

	if obj.AlgoOrderUpdate != nil {
		return obj.AlgoOrderUpdate
	}

	if obj.BalanceUpdate != nil {
		return obj.BalanceUpdate
	}

	if obj.ConditionalOrderTradeUpdate != nil {
		return obj.ConditionalOrderTradeUpdate
	}

	if obj.ExecutionReport != nil {
		return obj.ExecutionReport
	}

	if obj.LiabilityChange != nil {
		return obj.LiabilityChange
	}

	if obj.ListenKeyExpired != nil {
		return obj.ListenKeyExpired
	}

	if obj.OpenOrderLoss != nil {
		return obj.OpenOrderLoss
	}

	if obj.OrderTradeUpdate != nil {
		return obj.OrderTradeUpdate
	}

	if obj.OutboundAccountPosition != nil {
		return obj.OutboundAccountPosition
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
