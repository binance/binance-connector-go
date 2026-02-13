/*
Binance Derivatives Trading Portfolio Margin Pro WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin Pro WebSocket Market Streams
*/

package models

import (
	"encoding/json"
	"fmt"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// UserDataStreamEventsResponse - struct for UserDataStreamEventsResponse
type UserDataStreamEventsResponse struct {
	Risklevelchange *Risklevelchange
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
