/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the DepositHistoryTravelRuleResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DepositHistoryTravelRuleResponse{}

// DepositHistoryTravelRuleResponse struct for DepositHistoryTravelRuleResponse
type DepositHistoryTravelRuleResponse struct {
	Items []DepositHistoryTravelRuleResponseInner
}

// NewDepositHistoryTravelRuleResponse instantiates a new DepositHistoryTravelRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepositHistoryTravelRuleResponse() *DepositHistoryTravelRuleResponse {
	this := DepositHistoryTravelRuleResponse{}
	return &this
}

// NewDepositHistoryTravelRuleResponseWithDefaults instantiates a new DepositHistoryTravelRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositHistoryTravelRuleResponseWithDefaults() *DepositHistoryTravelRuleResponse {
	this := DepositHistoryTravelRuleResponse{}
	return &this
}

func (o DepositHistoryTravelRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DepositHistoryTravelRuleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *DepositHistoryTravelRuleResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableDepositHistoryTravelRuleResponse struct {
	value DepositHistoryTravelRuleResponse
	isSet bool
}

func (v NullableDepositHistoryTravelRuleResponse) Get() DepositHistoryTravelRuleResponse {
	return v.value
}

func (v *NullableDepositHistoryTravelRuleResponse) Set(val DepositHistoryTravelRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositHistoryTravelRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositHistoryTravelRuleResponse) Unset() {
	v.value = DepositHistoryTravelRuleResponse{}
	v.isSet = false
}

func NewNullableDepositHistoryTravelRuleResponse(val DepositHistoryTravelRuleResponse) *NullableDepositHistoryTravelRuleResponse {
	return &NullableDepositHistoryTravelRuleResponse{value: val, isSet: true}
}

func (v NullableDepositHistoryTravelRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositHistoryTravelRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
