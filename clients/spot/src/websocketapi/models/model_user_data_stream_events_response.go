/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"
	"fmt"

	"github.com/binance/binance-connector-go/common/common"
)

// UserDataStreamEventsResponse - struct for UserDataStreamEventsResponse
type UserDataStreamEventsResponse struct {
	BalanceUpdate           *BalanceUpdate
	EventStreamTerminated   *EventStreamTerminated
	ExecutionReport         *ExecutionReport
	ExternalLockUpdate      *ExternalLockUpdate
	ListStatus              *ListStatus
	OutboundAccountPosition *OutboundAccountPosition
}

// BalanceUpdateAsUserDataStreamEventsResponse is a convenience function that returns BalanceUpdate wrapped in UserDataStreamEventsResponse
func BalanceUpdateAsUserDataStreamEventsResponse(v *BalanceUpdate) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		BalanceUpdate: v,
	}
}

// EventStreamTerminatedAsUserDataStreamEventsResponse is a convenience function that returns EventStreamTerminated wrapped in UserDataStreamEventsResponse
func EventStreamTerminatedAsUserDataStreamEventsResponse(v *EventStreamTerminated) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		EventStreamTerminated: v,
	}
}

// ExecutionReportAsUserDataStreamEventsResponse is a convenience function that returns ExecutionReport wrapped in UserDataStreamEventsResponse
func ExecutionReportAsUserDataStreamEventsResponse(v *ExecutionReport) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		ExecutionReport: v,
	}
}

// ExternalLockUpdateAsUserDataStreamEventsResponse is a convenience function that returns ExternalLockUpdate wrapped in UserDataStreamEventsResponse
func ExternalLockUpdateAsUserDataStreamEventsResponse(v *ExternalLockUpdate) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		ExternalLockUpdate: v,
	}
}

// ListStatusAsUserDataStreamEventsResponse is a convenience function that returns ListStatus wrapped in UserDataStreamEventsResponse
func ListStatusAsUserDataStreamEventsResponse(v *ListStatus) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		ListStatus: v,
	}
}

// OutboundAccountPositionAsUserDataStreamEventsResponse is a convenience function that returns OutboundAccountPosition wrapped in UserDataStreamEventsResponse
func OutboundAccountPositionAsUserDataStreamEventsResponse(v *OutboundAccountPosition) UserDataStreamEventsResponse {
	return UserDataStreamEventsResponse{
		OutboundAccountPosition: v,
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

	eventMap, ok := jsonDict["event"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("event field is not a map")
	}
	eventType, ok := eventMap["e"].(string)
	if !ok {
		return fmt.Errorf("e field is not a string")
	}

	var modifiedData map[string]interface{}
	if err := json.Unmarshal(data, &modifiedData); err != nil {
		return fmt.Errorf("failed to unmarshal JSON for modification: %v", err)
	}

	// Remove the "e" field
	if event, ok := modifiedData["event"].(map[string]interface{}); ok {
		delete(event, "e")
	}

	// Marshal the modified data back to JSON
	cleanedData, err := json.Marshal(modifiedData["event"])
	if err != nil {
		return fmt.Errorf("failed to remarshal JSON: %v", err)
	}

	// check if the discriminator value is 'balanceUpdate'
	if eventType == "balanceUpdate" {
		// try to unmarshal JSON data into BalanceUpdate
		err = json.Unmarshal(cleanedData, &dst.BalanceUpdate)
		if err == nil {
			return nil // data stored in dst.BalanceUpdate, return on the first match
		} else {
			dst.BalanceUpdate = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as BalanceUpdate: %s", err.Error())
		}
	}

	// check if the discriminator value is 'eventStreamTerminated'
	if eventType == "eventStreamTerminated" {
		// try to unmarshal JSON data into EventStreamTerminated
		err = json.Unmarshal(cleanedData, &dst.EventStreamTerminated)
		if err == nil {
			return nil // data stored in dst.EventStreamTerminated, return on the first match
		} else {
			dst.EventStreamTerminated = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as EventStreamTerminated: %s", err.Error())
		}
	}

	// check if the discriminator value is 'executionReport'
	if eventType == "executionReport" {
		// try to unmarshal JSON data into ExecutionReport
		err = json.Unmarshal(cleanedData, &dst.ExecutionReport)
		if err == nil {
			return nil // data stored in dst.ExecutionReport, return on the first match
		} else {
			dst.ExecutionReport = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as ExecutionReport: %s", err.Error())
		}
	}

	// check if the discriminator value is 'externalLockUpdate'
	if eventType == "externalLockUpdate" {
		// try to unmarshal JSON data into ExternalLockUpdate
		err = json.Unmarshal(cleanedData, &dst.ExternalLockUpdate)
		if err == nil {
			return nil // data stored in dst.ExternalLockUpdate, return on the first match
		} else {
			dst.ExternalLockUpdate = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as ExternalLockUpdate: %s", err.Error())
		}
	}

	// check if the discriminator value is 'listStatus'
	if eventType == "listStatus" {
		// try to unmarshal JSON data into ListStatus
		err = json.Unmarshal(cleanedData, &dst.ListStatus)
		if err == nil {
			return nil // data stored in dst.ListStatus, return on the first match
		} else {
			dst.ListStatus = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as ListStatus: %s", err.Error())
		}
	}

	// check if the discriminator value is 'outboundAccountPosition'
	if eventType == "outboundAccountPosition" {
		// try to unmarshal JSON data into OutboundAccountPosition
		err = json.Unmarshal(cleanedData, &dst.OutboundAccountPosition)
		if err == nil {
			return nil // data stored in dst.OutboundAccountPosition, return on the first match
		} else {
			dst.OutboundAccountPosition = nil
			return fmt.Errorf("failed to unmarshal UserDataStreamEventsResponse as OutboundAccountPosition: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UserDataStreamEventsResponse) MarshalJSON() ([]byte, error) {
	if src.BalanceUpdate != nil {
		return json.Marshal(&src.BalanceUpdate)
	}

	if src.EventStreamTerminated != nil {
		return json.Marshal(&src.EventStreamTerminated)
	}

	if src.ExecutionReport != nil {
		return json.Marshal(&src.ExecutionReport)
	}

	if src.ExternalLockUpdate != nil {
		return json.Marshal(&src.ExternalLockUpdate)
	}

	if src.ListStatus != nil {
		return json.Marshal(&src.ListStatus)
	}

	if src.OutboundAccountPosition != nil {
		return json.Marshal(&src.OutboundAccountPosition)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UserDataStreamEventsResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.BalanceUpdate != nil {
		return obj.BalanceUpdate
	}

	if obj.EventStreamTerminated != nil {
		return obj.EventStreamTerminated
	}

	if obj.ExecutionReport != nil {
		return obj.ExecutionReport
	}

	if obj.ExternalLockUpdate != nil {
		return obj.ExternalLockUpdate
	}

	if obj.ListStatus != nil {
		return obj.ListStatus
	}

	if obj.OutboundAccountPosition != nil {
		return obj.OutboundAccountPosition
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
