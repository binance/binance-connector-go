/*
Binance Margin Trading WebSocket Market Streams

OpenAPI Specification for the Binance Margin Trading WebSocket Market Streams

API version: 1.0.0
*/

package models

import (
	"encoding/json"
	"fmt"

	"github.com/binance/binance-connector-go/common/common"
)

// TradeDataStreamEventsResponse - struct for TradeDataStreamEventsResponse
type TradeDataStreamEventsResponse struct {
	Balanceupdate           *Balanceupdate
	Executionreport         *Executionreport
	Listenkeyexpired        *Listenkeyexpired
	Liststatus              *Liststatus
	Outboundaccountposition *Outboundaccountposition
}

// BalanceupdateAsTradeDataStreamEventsResponse is a convenience function that returns Balanceupdate wrapped in TradeDataStreamEventsResponse
func BalanceupdateAsTradeDataStreamEventsResponse(v *Balanceupdate) TradeDataStreamEventsResponse {
	return TradeDataStreamEventsResponse{
		Balanceupdate: v,
	}
}

// ExecutionreportAsTradeDataStreamEventsResponse is a convenience function that returns Executionreport wrapped in TradeDataStreamEventsResponse
func ExecutionreportAsTradeDataStreamEventsResponse(v *Executionreport) TradeDataStreamEventsResponse {
	return TradeDataStreamEventsResponse{
		Executionreport: v,
	}
}

// ListenkeyexpiredAsTradeDataStreamEventsResponse is a convenience function that returns Listenkeyexpired wrapped in TradeDataStreamEventsResponse
func ListenkeyexpiredAsTradeDataStreamEventsResponse(v *Listenkeyexpired) TradeDataStreamEventsResponse {
	return TradeDataStreamEventsResponse{
		Listenkeyexpired: v,
	}
}

// ListstatusAsTradeDataStreamEventsResponse is a convenience function that returns Liststatus wrapped in TradeDataStreamEventsResponse
func ListstatusAsTradeDataStreamEventsResponse(v *Liststatus) TradeDataStreamEventsResponse {
	return TradeDataStreamEventsResponse{
		Liststatus: v,
	}
}

// OutboundaccountpositionAsTradeDataStreamEventsResponse is a convenience function that returns Outboundaccountposition wrapped in TradeDataStreamEventsResponse
func OutboundaccountpositionAsTradeDataStreamEventsResponse(v *Outboundaccountposition) TradeDataStreamEventsResponse {
	return TradeDataStreamEventsResponse{
		Outboundaccountposition: v,
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

	// check if the discriminator value is 'balanceUpdate'
	if jsonDict["e"] == "balanceUpdate" {
		// try to unmarshal JSON data into Balanceupdate
		err = json.Unmarshal(cleanedData, &dst.Balanceupdate)
		if err == nil {
			return nil // data stored in dst.Balanceupdate, return on the first match
		} else {
			dst.Balanceupdate = nil
			return fmt.Errorf("failed to unmarshal TradeDataStreamEventsResponse as Balanceupdate: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal TradeDataStreamEventsResponse as Executionreport: %s", err.Error())
		}
	}

	// check if the discriminator value is 'listStatus'
	if jsonDict["e"] == "listStatus" {
		// try to unmarshal JSON data into Liststatus
		err = json.Unmarshal(cleanedData, &dst.Liststatus)
		if err == nil {
			return nil // data stored in dst.Liststatus, return on the first match
		} else {
			dst.Liststatus = nil
			return fmt.Errorf("failed to unmarshal TradeDataStreamEventsResponse as Liststatus: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal TradeDataStreamEventsResponse as Listenkeyexpired: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal TradeDataStreamEventsResponse as Outboundaccountposition: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal TradeDataStreamEventsResponse as Balanceupdate: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal TradeDataStreamEventsResponse as Executionreport: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal TradeDataStreamEventsResponse as Listenkeyexpired: %s", err.Error())
		}
	}

	// check if the discriminator value is 'liststatus'
	if jsonDict["e"] == "liststatus" {
		// try to unmarshal JSON data into Liststatus
		err = json.Unmarshal(cleanedData, &dst.Liststatus)
		if err == nil {
			return nil // data stored in dst.Liststatus, return on the first match
		} else {
			dst.Liststatus = nil
			return fmt.Errorf("failed to unmarshal TradeDataStreamEventsResponse as Liststatus: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal TradeDataStreamEventsResponse as Outboundaccountposition: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TradeDataStreamEventsResponse) MarshalJSON() ([]byte, error) {
	if src.Balanceupdate != nil {
		return json.Marshal(&src.Balanceupdate)
	}

	if src.Executionreport != nil {
		return json.Marshal(&src.Executionreport)
	}

	if src.Listenkeyexpired != nil {
		return json.Marshal(&src.Listenkeyexpired)
	}

	if src.Liststatus != nil {
		return json.Marshal(&src.Liststatus)
	}

	if src.Outboundaccountposition != nil {
		return json.Marshal(&src.Outboundaccountposition)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TradeDataStreamEventsResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Balanceupdate != nil {
		return obj.Balanceupdate
	}

	if obj.Executionreport != nil {
		return obj.Executionreport
	}

	if obj.Listenkeyexpired != nil {
		return obj.Listenkeyexpired
	}

	if obj.Liststatus != nil {
		return obj.Liststatus
	}

	if obj.Outboundaccountposition != nil {
		return obj.Outboundaccountposition
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
