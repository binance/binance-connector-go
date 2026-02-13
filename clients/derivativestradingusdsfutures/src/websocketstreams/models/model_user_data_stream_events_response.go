/*
Binance Derivatives Trading USDS Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket Market Streams
*/

package models

import (
	"encoding/json"
	"fmt"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// UserDataStreamEventsResponse - struct for UserDataStreamEventsResponse
type UserDataStreamEventsResponse struct {
	AccountConfigUpdate           *AccountConfigUpdate
	AccountUpdate                 *AccountUpdate
	AlgoUpdate                    *AlgoUpdate
	ConditionalOrderTriggerReject *ConditionalOrderTriggerReject
	GridUpdate                    *GridUpdate
	Listenkeyexpired              *Listenkeyexpired
	MarginCall                    *MarginCall
	OrderTradeUpdate              *OrderTradeUpdate
	StrategyUpdate                *StrategyUpdate
	TradeLite                     *TradeLite
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

// AlgoUpdateAsUserDataStreamEventsResponse is a convenience function that returns AlgoUpdate wrapped in UserDataStreamEventsResponse
func AlgoUpdateAsUserDataStreamEventsResponse(v *AlgoUpdate) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		AlgoUpdate: v,
	}
}

// ConditionalOrderTriggerRejectAsUserDataStreamEventsResponse is a convenience function that returns ConditionalOrderTriggerReject wrapped in UserDataStreamEventsResponse
func ConditionalOrderTriggerRejectAsUserDataStreamEventsResponse(v *ConditionalOrderTriggerReject) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		ConditionalOrderTriggerReject: v,
	}
}

// GridUpdateAsUserDataStreamEventsResponse is a convenience function that returns GridUpdate wrapped in UserDataStreamEventsResponse
func GridUpdateAsUserDataStreamEventsResponse(v *GridUpdate) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		GridUpdate: v,
	}
}

// ListenkeyexpiredAsUserDataStreamEventsResponse is a convenience function that returns Listenkeyexpired wrapped in UserDataStreamEventsResponse
func ListenkeyexpiredAsUserDataStreamEventsResponse(v *Listenkeyexpired) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		Listenkeyexpired: v,
	}
}

// MarginCallAsUserDataStreamEventsResponse is a convenience function that returns MarginCall wrapped in UserDataStreamEventsResponse
func MarginCallAsUserDataStreamEventsResponse(v *MarginCall) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		MarginCall: v,
	}
}

// OrderTradeUpdateAsUserDataStreamEventsResponse is a convenience function that returns OrderTradeUpdate wrapped in UserDataStreamEventsResponse
func OrderTradeUpdateAsUserDataStreamEventsResponse(v *OrderTradeUpdate) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		OrderTradeUpdate: v,
	}
}

// StrategyUpdateAsUserDataStreamEventsResponse is a convenience function that returns StrategyUpdate wrapped in UserDataStreamEventsResponse
func StrategyUpdateAsUserDataStreamEventsResponse(v *StrategyUpdate) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		StrategyUpdate: v,
	}
}

// TradeLiteAsUserDataStreamEventsResponse is a convenience function that returns TradeLite wrapped in UserDataStreamEventsResponse
func TradeLiteAsUserDataStreamEventsResponse(v *TradeLite) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		TradeLite: v,
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
		// try to unmarshal JSON data into AlgoUpdate
		err = json.Unmarshal(cleanedData, &dst.AlgoUpdate)
		if err == nil {
			return nil // data stored in dst.AlgoUpdate, return on the first match
		} else {
			dst.AlgoUpdate = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as AlgoUpdate: %s", err.Error())
		}
	}

	// check if the discriminator value is 'CONDITIONAL_ORDER_TRIGGER_REJECT'
	if jsonDict["e"] == "CONDITIONAL_ORDER_TRIGGER_REJECT" {
		// try to unmarshal JSON data into ConditionalOrderTriggerReject
		err = json.Unmarshal(cleanedData, &dst.ConditionalOrderTriggerReject)
		if err == nil {
			return nil // data stored in dst.ConditionalOrderTriggerReject, return on the first match
		} else {
			dst.ConditionalOrderTriggerReject = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as ConditionalOrderTriggerReject: %s", err.Error())
		}
	}

	// check if the discriminator value is 'GRID_UPDATE'
	if jsonDict["e"] == "GRID_UPDATE" {
		// try to unmarshal JSON data into GridUpdate
		err = json.Unmarshal(cleanedData, &dst.GridUpdate)
		if err == nil {
			return nil // data stored in dst.GridUpdate, return on the first match
		} else {
			dst.GridUpdate = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as GridUpdate: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MARGIN_CALL'
	if jsonDict["e"] == "MARGIN_CALL" {
		// try to unmarshal JSON data into MarginCall
		err = json.Unmarshal(cleanedData, &dst.MarginCall)
		if err == nil {
			return nil // data stored in dst.MarginCall, return on the first match
		} else {
			dst.MarginCall = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as MarginCall: %s", err.Error())
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

	// check if the discriminator value is 'STRATEGY_UPDATE'
	if jsonDict["e"] == "STRATEGY_UPDATE" {
		// try to unmarshal JSON data into StrategyUpdate
		err = json.Unmarshal(cleanedData, &dst.StrategyUpdate)
		if err == nil {
			return nil // data stored in dst.StrategyUpdate, return on the first match
		} else {
			dst.StrategyUpdate = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as StrategyUpdate: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TRADE_LITE'
	if jsonDict["e"] == "TRADE_LITE" {
		// try to unmarshal JSON data into TradeLite
		err = json.Unmarshal(cleanedData, &dst.TradeLite)
		if err == nil {
			return nil // data stored in dst.TradeLite, return on the first match
		} else {
			dst.TradeLite = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as TradeLite: %s", err.Error())
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

	// check if the discriminator value is 'accountConfigUpdate'
	if jsonDict["e"] == "accountConfigUpdate" {
		// try to unmarshal JSON data into AccountConfigUpdate
		err = json.Unmarshal(cleanedData, &dst.AccountConfigUpdate)
		if err == nil {
			return nil // data stored in dst.AccountConfigUpdate, return on the first match
		} else {
			dst.AccountConfigUpdate = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as AccountConfigUpdate: %s", err.Error())
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

	// check if the discriminator value is 'algoUpdate'
	if jsonDict["e"] == "algoUpdate" {
		// try to unmarshal JSON data into AlgoUpdate
		err = json.Unmarshal(cleanedData, &dst.AlgoUpdate)
		if err == nil {
			return nil // data stored in dst.AlgoUpdate, return on the first match
		} else {
			dst.AlgoUpdate = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as AlgoUpdate: %s", err.Error())
		}
	}

	// check if the discriminator value is 'conditionalOrderTriggerReject'
	if jsonDict["e"] == "conditionalOrderTriggerReject" {
		// try to unmarshal JSON data into ConditionalOrderTriggerReject
		err = json.Unmarshal(cleanedData, &dst.ConditionalOrderTriggerReject)
		if err == nil {
			return nil // data stored in dst.ConditionalOrderTriggerReject, return on the first match
		} else {
			dst.ConditionalOrderTriggerReject = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as ConditionalOrderTriggerReject: %s", err.Error())
		}
	}

	// check if the discriminator value is 'gridUpdate'
	if jsonDict["e"] == "gridUpdate" {
		// try to unmarshal JSON data into GridUpdate
		err = json.Unmarshal(cleanedData, &dst.GridUpdate)
		if err == nil {
			return nil // data stored in dst.GridUpdate, return on the first match
		} else {
			dst.GridUpdate = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as GridUpdate: %s", err.Error())
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

	// check if the discriminator value is 'marginCall'
	if jsonDict["e"] == "marginCall" {
		// try to unmarshal JSON data into MarginCall
		err = json.Unmarshal(cleanedData, &dst.MarginCall)
		if err == nil {
			return nil // data stored in dst.MarginCall, return on the first match
		} else {
			dst.MarginCall = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as MarginCall: %s", err.Error())
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

	// check if the discriminator value is 'strategyUpdate'
	if jsonDict["e"] == "strategyUpdate" {
		// try to unmarshal JSON data into StrategyUpdate
		err = json.Unmarshal(cleanedData, &dst.StrategyUpdate)
		if err == nil {
			return nil // data stored in dst.StrategyUpdate, return on the first match
		} else {
			dst.StrategyUpdate = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as StrategyUpdate: %s", err.Error())
		}
	}

	// check if the discriminator value is 'tradeLite'
	if jsonDict["e"] == "tradeLite" {
		// try to unmarshal JSON data into TradeLite
		err = json.Unmarshal(cleanedData, &dst.TradeLite)
		if err == nil {
			return nil // data stored in dst.TradeLite, return on the first match
		} else {
			dst.TradeLite = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as TradeLite: %s", err.Error())
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

	if src.AlgoUpdate != nil {
		return json.Marshal(&src.AlgoUpdate)
	}

	if src.ConditionalOrderTriggerReject != nil {
		return json.Marshal(&src.ConditionalOrderTriggerReject)
	}

	if src.GridUpdate != nil {
		return json.Marshal(&src.GridUpdate)
	}

	if src.Listenkeyexpired != nil {
		return json.Marshal(&src.Listenkeyexpired)
	}

	if src.MarginCall != nil {
		return json.Marshal(&src.MarginCall)
	}

	if src.OrderTradeUpdate != nil {
		return json.Marshal(&src.OrderTradeUpdate)
	}

	if src.StrategyUpdate != nil {
		return json.Marshal(&src.StrategyUpdate)
	}

	if src.TradeLite != nil {
		return json.Marshal(&src.TradeLite)
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

	if obj.AlgoUpdate != nil {
		return obj.AlgoUpdate
	}

	if obj.ConditionalOrderTriggerReject != nil {
		return obj.ConditionalOrderTriggerReject
	}

	if obj.GridUpdate != nil {
		return obj.GridUpdate
	}

	if obj.Listenkeyexpired != nil {
		return obj.Listenkeyexpired
	}

	if obj.MarginCall != nil {
		return obj.MarginCall
	}

	if obj.OrderTradeUpdate != nil {
		return obj.OrderTradeUpdate
	}

	if obj.StrategyUpdate != nil {
		return obj.StrategyUpdate
	}

	if obj.TradeLite != nil {
		return obj.TradeLite
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
