/*
Binance Derivatives Trading Portfolio Margin WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin WebSocket Market Streams

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
	AccountConfigUpdate         *AccountConfigUpdate
	AccountUpdate               *AccountUpdate
	Balanceupdate               *Balanceupdate
	ConditionalOrderTradeUpdate *ConditionalOrderTradeUpdate
	Executionreport             *Executionreport
	Liabilitychange             *Liabilitychange
	Listenkeyexpired            *Listenkeyexpired
	Openorderloss               *Openorderloss
	OrderTradeUpdate            *OrderTradeUpdate
	Outboundaccountposition     *Outboundaccountposition
	Risklevelchange             *Risklevelchange
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

// BalanceupdateAsUserDataStreamEventsResponse is a convenience function that returns Balanceupdate wrapped in UserDataStreamEventsResponse
func BalanceupdateAsUserDataStreamEventsResponse(v *Balanceupdate) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		Balanceupdate: v,
	}
}

// ConditionalOrderTradeUpdateAsUserDataStreamEventsResponse is a convenience function that returns ConditionalOrderTradeUpdate wrapped in UserDataStreamEventsResponse
func ConditionalOrderTradeUpdateAsUserDataStreamEventsResponse(v *ConditionalOrderTradeUpdate) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		ConditionalOrderTradeUpdate: v,
	}
}

// ExecutionreportAsUserDataStreamEventsResponse is a convenience function that returns Executionreport wrapped in UserDataStreamEventsResponse
func ExecutionreportAsUserDataStreamEventsResponse(v *Executionreport) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		Executionreport: v,
	}
}

// LiabilitychangeAsUserDataStreamEventsResponse is a convenience function that returns Liabilitychange wrapped in UserDataStreamEventsResponse
func LiabilitychangeAsUserDataStreamEventsResponse(v *Liabilitychange) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		Liabilitychange: v,
	}
}

// ListenkeyexpiredAsUserDataStreamEventsResponse is a convenience function that returns Listenkeyexpired wrapped in UserDataStreamEventsResponse
func ListenkeyexpiredAsUserDataStreamEventsResponse(v *Listenkeyexpired) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		Listenkeyexpired: v,
	}
}

// OpenorderlossAsUserDataStreamEventsResponse is a convenience function that returns Openorderloss wrapped in UserDataStreamEventsResponse
func OpenorderlossAsUserDataStreamEventsResponse(v *Openorderloss) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		Openorderloss: v,
	}
}

// OrderTradeUpdateAsUserDataStreamEventsResponse is a convenience function that returns OrderTradeUpdate wrapped in UserDataStreamEventsResponse
func OrderTradeUpdateAsUserDataStreamEventsResponse(v *OrderTradeUpdate) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		OrderTradeUpdate: v,
	}
}

// OutboundaccountpositionAsUserDataStreamEventsResponse is a convenience function that returns Outboundaccountposition wrapped in UserDataStreamEventsResponse
func OutboundaccountpositionAsUserDataStreamEventsResponse(v *Outboundaccountposition) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		Outboundaccountposition: v,
	}
}

// RisklevelchangeAsUserDataStreamEventsResponse is a convenience function that returns Risklevelchange wrapped in UserDataStreamEventsResponse
func RisklevelchangeAsUserDataStreamEventsResponse(v *Risklevelchange) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		Risklevelchange: v,
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
		// try to unmarshal JSON data into Balanceupdate
		err = json.Unmarshal(cleanedData, &dst.Balanceupdate)
		if err == nil {
			return nil // data stored in dst.Balanceupdate, return on the first match
		} else {
			dst.Balanceupdate = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as Balanceupdate: %s", err.Error())
		}
	}

	// check if the discriminator value is 'executionReport'
	if jsonDict["e"] == "executionReport" {
		// try to unmarshal JSON data into Executionreport
		err = json.Unmarshal(cleanedData, &dst.Executionreport)
		if err == nil {
			return nil // data stored in dst.Executionreport, return on the first match
		} else {
			dst.Executionreport = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as Executionreport: %s", err.Error())
		}
	}

	// check if the discriminator value is 'liabilityChange'
	if jsonDict["e"] == "liabilityChange" {
		// try to unmarshal JSON data into Liabilitychange
		err = json.Unmarshal(cleanedData, &dst.Liabilitychange)
		if err == nil {
			return nil // data stored in dst.Liabilitychange, return on the first match
		} else {
			dst.Liabilitychange = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as Liabilitychange: %s", err.Error())
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

	// check if the discriminator value is 'openOrderLoss'
	if jsonDict["e"] == "openOrderLoss" {
		// try to unmarshal JSON data into Openorderloss
		err = json.Unmarshal(cleanedData, &dst.Openorderloss)
		if err == nil {
			return nil // data stored in dst.Openorderloss, return on the first match
		} else {
			dst.Openorderloss = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as Openorderloss: %s", err.Error())
		}
	}

	// check if the discriminator value is 'outboundAccountPosition'
	if jsonDict["e"] == "outboundAccountPosition" {
		// try to unmarshal JSON data into Outboundaccountposition
		err = json.Unmarshal(cleanedData, &dst.Outboundaccountposition)
		if err == nil {
			return nil // data stored in dst.Outboundaccountposition, return on the first match
		} else {
			dst.Outboundaccountposition = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as Outboundaccountposition: %s", err.Error())
		}
	}

	// check if the discriminator value is 'riskLevelChange'
	if jsonDict["e"] == "riskLevelChange" {
		// try to unmarshal JSON data into Risklevelchange
		err = json.Unmarshal(cleanedData, &dst.Risklevelchange)
		if err == nil {
			return nil // data stored in dst.Risklevelchange, return on the first match
		} else {
			dst.Risklevelchange = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as Risklevelchange: %s", err.Error())
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

	// check if the discriminator value is 'balanceupdate'
	if jsonDict["e"] == "balanceupdate" {
		// try to unmarshal JSON data into Balanceupdate
		err = json.Unmarshal(cleanedData, &dst.Balanceupdate)
		if err == nil {
			return nil // data stored in dst.Balanceupdate, return on the first match
		} else {
			dst.Balanceupdate = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as Balanceupdate: %s", err.Error())
		}
	}

	// check if the discriminator value is 'conditionalOrderTradeUpdate'
	if jsonDict["e"] == "conditionalOrderTradeUpdate" {
		// try to unmarshal JSON data into ConditionalOrderTradeUpdate
		err = json.Unmarshal(cleanedData, &dst.ConditionalOrderTradeUpdate)
		if err == nil {
			return nil // data stored in dst.ConditionalOrderTradeUpdate, return on the first match
		} else {
			dst.ConditionalOrderTradeUpdate = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as ConditionalOrderTradeUpdate: %s", err.Error())
		}
	}

	// check if the discriminator value is 'executionreport'
	if jsonDict["e"] == "executionreport" {
		// try to unmarshal JSON data into Executionreport
		err = json.Unmarshal(cleanedData, &dst.Executionreport)
		if err == nil {
			return nil // data stored in dst.Executionreport, return on the first match
		} else {
			dst.Executionreport = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as Executionreport: %s", err.Error())
		}
	}

	// check if the discriminator value is 'liabilitychange'
	if jsonDict["e"] == "liabilitychange" {
		// try to unmarshal JSON data into Liabilitychange
		err = json.Unmarshal(cleanedData, &dst.Liabilitychange)
		if err == nil {
			return nil // data stored in dst.Liabilitychange, return on the first match
		} else {
			dst.Liabilitychange = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as Liabilitychange: %s", err.Error())
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

	// check if the discriminator value is 'openorderloss'
	if jsonDict["e"] == "openorderloss" {
		// try to unmarshal JSON data into Openorderloss
		err = json.Unmarshal(cleanedData, &dst.Openorderloss)
		if err == nil {
			return nil // data stored in dst.Openorderloss, return on the first match
		} else {
			dst.Openorderloss = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as Openorderloss: %s", err.Error())
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

	// check if the discriminator value is 'outboundaccountposition'
	if jsonDict["e"] == "outboundaccountposition" {
		// try to unmarshal JSON data into Outboundaccountposition
		err = json.Unmarshal(cleanedData, &dst.Outboundaccountposition)
		if err == nil {
			return nil // data stored in dst.Outboundaccountposition, return on the first match
		} else {
			dst.Outboundaccountposition = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as Outboundaccountposition: %s", err.Error())
		}
	}

	// check if the discriminator value is 'risklevelchange'
	if jsonDict["e"] == "risklevelchange" {
		// try to unmarshal JSON data into Risklevelchange
		err = json.Unmarshal(cleanedData, &dst.Risklevelchange)
		if err == nil {
			return nil // data stored in dst.Risklevelchange, return on the first match
		} else {
			dst.Risklevelchange = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as Risklevelchange: %s", err.Error())
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

	if src.Balanceupdate != nil {
		return json.Marshal(&src.Balanceupdate)
	}

	if src.ConditionalOrderTradeUpdate != nil {
		return json.Marshal(&src.ConditionalOrderTradeUpdate)
	}

	if src.Executionreport != nil {
		return json.Marshal(&src.Executionreport)
	}

	if src.Liabilitychange != nil {
		return json.Marshal(&src.Liabilitychange)
	}

	if src.Listenkeyexpired != nil {
		return json.Marshal(&src.Listenkeyexpired)
	}

	if src.Openorderloss != nil {
		return json.Marshal(&src.Openorderloss)
	}

	if src.OrderTradeUpdate != nil {
		return json.Marshal(&src.OrderTradeUpdate)
	}

	if src.Outboundaccountposition != nil {
		return json.Marshal(&src.Outboundaccountposition)
	}

	if src.Risklevelchange != nil {
		return json.Marshal(&src.Risklevelchange)
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

	if obj.Balanceupdate != nil {
		return obj.Balanceupdate
	}

	if obj.ConditionalOrderTradeUpdate != nil {
		return obj.ConditionalOrderTradeUpdate
	}

	if obj.Executionreport != nil {
		return obj.Executionreport
	}

	if obj.Liabilitychange != nil {
		return obj.Liabilitychange
	}

	if obj.Listenkeyexpired != nil {
		return obj.Listenkeyexpired
	}

	if obj.Openorderloss != nil {
		return obj.Openorderloss
	}

	if obj.OrderTradeUpdate != nil {
		return obj.OrderTradeUpdate
	}

	if obj.Outboundaccountposition != nil {
		return obj.Outboundaccountposition
	}

	if obj.Risklevelchange != nil {
		return obj.Risklevelchange
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
