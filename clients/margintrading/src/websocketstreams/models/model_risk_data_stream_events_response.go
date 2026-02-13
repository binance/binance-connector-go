/*
Binance Margin Trading WebSocket Market Streams

OpenAPI Specification for the Binance Margin Trading WebSocket Market Streams
*/

package models

import (
	"encoding/json"
	"fmt"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// RiskDataStreamEventsResponse - struct for RiskDataStreamEventsResponse
type RiskDataStreamEventsResponse struct {
	MarginLevelStatusChange *MarginLevelStatusChange
	UserLiabilityChange     *UserLiabilityChange
}

// MarginLevelStatusChangeAsRiskDataStreamEventsResponse is a convenience function that returns MarginLevelStatusChange wrapped in RiskDataStreamEventsResponse
func MarginLevelStatusChangeAsRiskDataStreamEventsResponse(v *MarginLevelStatusChange) RiskDataStreamEventsResponse {
	return RiskDataStreamEventsResponse{
		MarginLevelStatusChange: v,
	}
}

// UserLiabilityChangeAsRiskDataStreamEventsResponse is a convenience function that returns UserLiabilityChange wrapped in RiskDataStreamEventsResponse
func UserLiabilityChangeAsRiskDataStreamEventsResponse(v *UserLiabilityChange) RiskDataStreamEventsResponse {
	return RiskDataStreamEventsResponse{
		UserLiabilityChange: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *RiskDataStreamEventsResponse) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("failed to unmarshal RiskDataStreamEventsResponse as MarginLevelStatusChange: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal RiskDataStreamEventsResponse as UserLiabilityChange: %s", err.Error())
		}
	}

	// check if the discriminator value is 'marginLevelStatusChange'
	if jsonDict["e"] == "marginLevelStatusChange" {
		// try to unmarshal JSON data into MarginLevelStatusChange
		err = json.Unmarshal(cleanedData, &dst.MarginLevelStatusChange)
		if err == nil {
			return nil // data stored in dst.MarginLevelStatusChange, return on the first match
		} else {
			dst.MarginLevelStatusChange = nil
			return fmt.Errorf("failed to unmarshal RiskDataStreamEventsResponse as MarginLevelStatusChange: %s", err.Error())
		}
	}

	// check if the discriminator value is 'userLiabilityChange'
	if jsonDict["e"] == "userLiabilityChange" {
		// try to unmarshal JSON data into UserLiabilityChange
		err = json.Unmarshal(cleanedData, &dst.UserLiabilityChange)
		if err == nil {
			return nil // data stored in dst.UserLiabilityChange, return on the first match
		} else {
			dst.UserLiabilityChange = nil
			return fmt.Errorf("failed to unmarshal RiskDataStreamEventsResponse as UserLiabilityChange: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RiskDataStreamEventsResponse) MarshalJSON() ([]byte, error) {
	if src.MarginLevelStatusChange != nil {
		return json.Marshal(&src.MarginLevelStatusChange)
	}

	if src.UserLiabilityChange != nil {
		return json.Marshal(&src.UserLiabilityChange)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RiskDataStreamEventsResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MarginLevelStatusChange != nil {
		return obj.MarginLevelStatusChange
	}

	if obj.UserLiabilityChange != nil {
		return obj.UserLiabilityChange
	}

	// all schemas are nil
	return nil
}

type NullableRiskDataStreamEventsResponse struct {
	value *RiskDataStreamEventsResponse
	isSet bool
}

func (v NullableRiskDataStreamEventsResponse) Get() *RiskDataStreamEventsResponse {
	return v.value
}

func (v *NullableRiskDataStreamEventsResponse) Set(val *RiskDataStreamEventsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskDataStreamEventsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskDataStreamEventsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskDataStreamEventsResponse(val *RiskDataStreamEventsResponse) *NullableRiskDataStreamEventsResponse {
	return &NullableRiskDataStreamEventsResponse{value: val, isSet: true}
}

func (v NullableRiskDataStreamEventsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskDataStreamEventsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
