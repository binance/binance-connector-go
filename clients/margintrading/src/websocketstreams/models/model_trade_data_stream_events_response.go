/*
Margin WebSocket Market Streams

Access account information, borrow and repay assets, and trade with Binance Margin.
*/

package models

import (
	"encoding/json"
	"fmt"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// TradeDataStreamEventsResponse - struct for TradeDataStreamEventsResponse
type TradeDataStreamEventsResponse struct {
	BalanceUpdate           *BalanceUpdate
	ExecutionReport         *ExecutionReport
	ListStatus              *ListStatus
	ListenKeyExpired        *ListenKeyExpired
	MarginLevelStatusChange *MarginLevelStatusChange
	OutboundAccountPosition *OutboundAccountPosition
	UserLiabilityChange     *UserLiabilityChange
}

// BalanceUpdateAsTradeDataStreamEventsResponse is a convenience function that returns BalanceUpdate wrapped in TradeDataStreamEventsResponse
func BalanceUpdateAsTradeDataStreamEventsResponse(v *BalanceUpdate) TradeDataStreamEventsResponse {
	return TradeDataStreamEventsResponse{
		BalanceUpdate: v,
	}
}

// ExecutionReportAsTradeDataStreamEventsResponse is a convenience function that returns ExecutionReport wrapped in TradeDataStreamEventsResponse
func ExecutionReportAsTradeDataStreamEventsResponse(v *ExecutionReport) TradeDataStreamEventsResponse {
	return TradeDataStreamEventsResponse{
		ExecutionReport: v,
	}
}

// ListStatusAsTradeDataStreamEventsResponse is a convenience function that returns ListStatus wrapped in TradeDataStreamEventsResponse
func ListStatusAsTradeDataStreamEventsResponse(v *ListStatus) TradeDataStreamEventsResponse {
	return TradeDataStreamEventsResponse{
		ListStatus: v,
	}
}

// ListenKeyExpiredAsTradeDataStreamEventsResponse is a convenience function that returns ListenKeyExpired wrapped in TradeDataStreamEventsResponse
func ListenKeyExpiredAsTradeDataStreamEventsResponse(v *ListenKeyExpired) TradeDataStreamEventsResponse {
	return TradeDataStreamEventsResponse{
		ListenKeyExpired: v,
	}
}

// MarginLevelStatusChangeAsTradeDataStreamEventsResponse is a convenience function that returns MarginLevelStatusChange wrapped in TradeDataStreamEventsResponse
func MarginLevelStatusChangeAsTradeDataStreamEventsResponse(v *MarginLevelStatusChange) TradeDataStreamEventsResponse {
	return TradeDataStreamEventsResponse{
		MarginLevelStatusChange: v,
	}
}

// OutboundAccountPositionAsTradeDataStreamEventsResponse is a convenience function that returns OutboundAccountPosition wrapped in TradeDataStreamEventsResponse
func OutboundAccountPositionAsTradeDataStreamEventsResponse(v *OutboundAccountPosition) TradeDataStreamEventsResponse {
	return TradeDataStreamEventsResponse{
		OutboundAccountPosition: v,
	}
}

// UserLiabilityChangeAsTradeDataStreamEventsResponse is a convenience function that returns UserLiabilityChange wrapped in TradeDataStreamEventsResponse
func UserLiabilityChangeAsTradeDataStreamEventsResponse(v *UserLiabilityChange) TradeDataStreamEventsResponse {
	return TradeDataStreamEventsResponse{
		UserLiabilityChange: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *TradeDataStreamEventsResponse) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'MARGIN_LEVEL_STATUS_CHANGE'
	if jsonDict["e"] == "MARGIN_LEVEL_STATUS_CHANGE" {
		// try to unmarshal JSON data into MarginLevelStatusChange
		err = json.Unmarshal(cleanedData, &dst.MarginLevelStatusChange)
		if err == nil {
			return nil // data stored in dst.MarginLevelStatusChange, return on the first match
		} else {
			dst.MarginLevelStatusChange = nil
			return fmt.Errorf("failed to unmarshal TradeDataStreamEventsResponse as MarginLevelStatusChange: %s", err.Error())
		}
	}

	// check if the discriminator value is 'USER_LIABILITY_CHANGE'
	if jsonDict["e"] == "USER_LIABILITY_CHANGE" {
		// try to unmarshal JSON data into UserLiabilityChange
		err = json.Unmarshal(cleanedData, &dst.UserLiabilityChange)
		if err == nil {
			return nil // data stored in dst.UserLiabilityChange, return on the first match
		} else {
			dst.UserLiabilityChange = nil
			return fmt.Errorf("failed to unmarshal TradeDataStreamEventsResponse as UserLiabilityChange: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal TradeDataStreamEventsResponse as BalanceUpdate: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal TradeDataStreamEventsResponse as ExecutionReport: %s", err.Error())
		}
	}

	// check if the discriminator value is 'listStatus'
	if jsonDict["e"] == "listStatus" {
		// try to unmarshal JSON data into ListStatus
		err = json.Unmarshal(cleanedData, &dst.ListStatus)
		if err == nil {
			return nil // data stored in dst.ListStatus, return on the first match
		} else {
			dst.ListStatus = nil
			return fmt.Errorf("failed to unmarshal TradeDataStreamEventsResponse as ListStatus: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal TradeDataStreamEventsResponse as ListenKeyExpired: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal TradeDataStreamEventsResponse as OutboundAccountPosition: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TradeDataStreamEventsResponse) MarshalJSON() ([]byte, error) {
	if src.BalanceUpdate != nil {
		return json.Marshal(&src.BalanceUpdate)
	}

	if src.ExecutionReport != nil {
		return json.Marshal(&src.ExecutionReport)
	}

	if src.ListStatus != nil {
		return json.Marshal(&src.ListStatus)
	}

	if src.ListenKeyExpired != nil {
		return json.Marshal(&src.ListenKeyExpired)
	}

	if src.MarginLevelStatusChange != nil {
		return json.Marshal(&src.MarginLevelStatusChange)
	}

	if src.OutboundAccountPosition != nil {
		return json.Marshal(&src.OutboundAccountPosition)
	}

	if src.UserLiabilityChange != nil {
		return json.Marshal(&src.UserLiabilityChange)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TradeDataStreamEventsResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.BalanceUpdate != nil {
		return obj.BalanceUpdate
	}

	if obj.ExecutionReport != nil {
		return obj.ExecutionReport
	}

	if obj.ListStatus != nil {
		return obj.ListStatus
	}

	if obj.ListenKeyExpired != nil {
		return obj.ListenKeyExpired
	}

	if obj.MarginLevelStatusChange != nil {
		return obj.MarginLevelStatusChange
	}

	if obj.OutboundAccountPosition != nil {
		return obj.OutboundAccountPosition
	}

	if obj.UserLiabilityChange != nil {
		return obj.UserLiabilityChange
	}

	// all schemas are nil
	return nil
}

type NullableTradeDataStreamEventsResponse struct {
	value *TradeDataStreamEventsResponse
	isSet bool
}

func (v NullableTradeDataStreamEventsResponse) Get() *TradeDataStreamEventsResponse {
	return v.value
}

func (v *NullableTradeDataStreamEventsResponse) Set(val *TradeDataStreamEventsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTradeDataStreamEventsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTradeDataStreamEventsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTradeDataStreamEventsResponse(val *TradeDataStreamEventsResponse) *NullableTradeDataStreamEventsResponse {
	return &NullableTradeDataStreamEventsResponse{value: val, isSet: true}
}

func (v NullableTradeDataStreamEventsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTradeDataStreamEventsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
