/*
Prediction Trading Streams

Binance Web3 Prediction pushes 14 types of per-user business events in real time via the Binance SApi WebSocket channel, covering trading (market / limit), reward claim, fund transfer and market resolution, plus order book (bid/ask) quotes broadcast per topic. Business events are routed by exact `(topic, userId)` match; order book topics are broadcast to every subscriber. Delivery is at-most-once, ordering is not guaranteed, and offline messages are not cached — pull the latest state via the REST API after reconnecting and deduplicate on `pushId` (business events) or `marketId` + `updateTimestampMs` (order book).
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the RewardClaimFailedResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RewardClaimFailedResponse{}

// RewardClaimFailedResponse struct for RewardClaimFailedResponse
type RewardClaimFailedResponse struct {
	// Unique push ID in the form `pm_<refId>_<scenarioCode>_<8hex>`, present in every topic; `refId` is the claim `batchId`. Use it for client-side deduplication.
	PushId *string `json:"pushId,omitempty"`
	// Amount of the **first failed claim**, `dataItems[0].amount`. Truncated to 2 decimals (towards zero); rendered as `<0.01` when the value is below 0.01.
	Amount *string `json:"amount,omitempty"`
	// Outcome name of the **first failed claim**, `dataItems[0].outcome.name` (e.g. `Yes` / `No` for binary markets). Passed through as-is, never truncated; may be an empty string when the data is missing.
	Outcome              *string `json:"outcome,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RewardClaimFailedResponse RewardClaimFailedResponse

// NewRewardClaimFailedResponse instantiates a new RewardClaimFailedResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRewardClaimFailedResponse() *RewardClaimFailedResponse {
	this := RewardClaimFailedResponse{}
	return &this
}

// NewRewardClaimFailedResponseWithDefaults instantiates a new RewardClaimFailedResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRewardClaimFailedResponseWithDefaults() *RewardClaimFailedResponse {
	this := RewardClaimFailedResponse{}
	return &this
}

// GetPushId returns the PushId field value if set, zero value otherwise.
func (o *RewardClaimFailedResponse) GetPushId() string {
	if o == nil || common.IsNil(o.PushId) {
		var ret string
		return ret
	}
	return *o.PushId
}

// GetPushIdOk returns a tuple with the PushId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RewardClaimFailedResponse) GetPushIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.PushId) {
		return nil, false
	}
	return o.PushId, true
}

// HasPushId returns a boolean if a field has been set.
func (o *RewardClaimFailedResponse) HasPushId() bool {
	if o != nil && !common.IsNil(o.PushId) {
		return true
	}

	return false
}

// SetPushId gets a reference to the given string and assigns it to the PushId field.
func (o *RewardClaimFailedResponse) SetPushId(v string) {
	o.PushId = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *RewardClaimFailedResponse) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RewardClaimFailedResponse) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *RewardClaimFailedResponse) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *RewardClaimFailedResponse) SetAmount(v string) {
	o.Amount = &v
}

// GetOutcome returns the Outcome field value if set, zero value otherwise.
func (o *RewardClaimFailedResponse) GetOutcome() string {
	if o == nil || common.IsNil(o.Outcome) {
		var ret string
		return ret
	}
	return *o.Outcome
}

// GetOutcomeOk returns a tuple with the Outcome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RewardClaimFailedResponse) GetOutcomeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Outcome) {
		return nil, false
	}
	return o.Outcome, true
}

// HasOutcome returns a boolean if a field has been set.
func (o *RewardClaimFailedResponse) HasOutcome() bool {
	if o != nil && !common.IsNil(o.Outcome) {
		return true
	}

	return false
}

// SetOutcome gets a reference to the given string and assigns it to the Outcome field.
func (o *RewardClaimFailedResponse) SetOutcome(v string) {
	o.Outcome = &v
}

func (o RewardClaimFailedResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RewardClaimFailedResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.PushId) {
		toSerialize["pushId"] = o.PushId
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.Outcome) {
		toSerialize["outcome"] = o.Outcome
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RewardClaimFailedResponse) UnmarshalJSON(data []byte) (err error) {
	varRewardClaimFailedResponse := _RewardClaimFailedResponse{}

	err = json.Unmarshal(data, &varRewardClaimFailedResponse)

	if err != nil {
		return err
	}

	*o = RewardClaimFailedResponse(varRewardClaimFailedResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pushId")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "outcome")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRewardClaimFailedResponse struct {
	value *RewardClaimFailedResponse
	isSet bool
}

func (v NullableRewardClaimFailedResponse) Get() *RewardClaimFailedResponse {
	return v.value
}

func (v *NullableRewardClaimFailedResponse) Set(val *RewardClaimFailedResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRewardClaimFailedResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRewardClaimFailedResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRewardClaimFailedResponse(val *RewardClaimFailedResponse) *NullableRewardClaimFailedResponse {
	return &NullableRewardClaimFailedResponse{value: val, isSet: true}
}

func (v NullableRewardClaimFailedResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRewardClaimFailedResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
