/*
Prediction Trading Streams

Binance Web3 Prediction pushes 14 types of per-user business events in real time via the Binance SApi WebSocket channel, covering trading (market / limit), reward claim, fund transfer and market resolution, plus order book (bid/ask) quotes broadcast per topic. Business events are routed by exact `(topic, userId)` match; order book topics are broadcast to every subscriber. Delivery is at-most-once, ordering is not guaranteed, and offline messages are not cached — pull the latest state via the REST API after reconnecting and deduplicate on `pushId` (business events) or `marketId` + `updateTimestampMs` (order book).
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the MarketBuyOrderFullyFilledResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MarketBuyOrderFullyFilledResponse{}

// MarketBuyOrderFullyFilledResponse struct for MarketBuyOrderFullyFilledResponse
type MarketBuyOrderFullyFilledResponse struct {
	// Unique push ID in the form `pm_<refId>_<scenarioCode>_<8hex>`, present in every topic; `refId` is `OrderHistory.orderId`. Use it for client-side deduplication.
	PushId *string `json:"pushId,omitempty"`
	// USDT amount the user paid, from `OrderHistory.amountIn`. Truncated to 2 decimals (towards zero); rendered as `<0.01` when the value is below 0.01.
	Amount *string `json:"amount,omitempty"`
	// Market title from `MarketTopic.title`, truncated to 12 chars with `..` appended when longer (max 14 chars); empty string when the title is missing.
	Topic                *string `json:"topic,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarketBuyOrderFullyFilledResponse MarketBuyOrderFullyFilledResponse

// NewMarketBuyOrderFullyFilledResponse instantiates a new MarketBuyOrderFullyFilledResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketBuyOrderFullyFilledResponse() *MarketBuyOrderFullyFilledResponse {
	this := MarketBuyOrderFullyFilledResponse{}
	return &this
}

// NewMarketBuyOrderFullyFilledResponseWithDefaults instantiates a new MarketBuyOrderFullyFilledResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketBuyOrderFullyFilledResponseWithDefaults() *MarketBuyOrderFullyFilledResponse {
	this := MarketBuyOrderFullyFilledResponse{}
	return &this
}

// GetPushId returns the PushId field value if set, zero value otherwise.
func (o *MarketBuyOrderFullyFilledResponse) GetPushId() string {
	if o == nil || common.IsNil(o.PushId) {
		var ret string
		return ret
	}
	return *o.PushId
}

// GetPushIdOk returns a tuple with the PushId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketBuyOrderFullyFilledResponse) GetPushIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.PushId) {
		return nil, false
	}
	return o.PushId, true
}

// HasPushId returns a boolean if a field has been set.
func (o *MarketBuyOrderFullyFilledResponse) HasPushId() bool {
	if o != nil && !common.IsNil(o.PushId) {
		return true
	}

	return false
}

// SetPushId gets a reference to the given string and assigns it to the PushId field.
func (o *MarketBuyOrderFullyFilledResponse) SetPushId(v string) {
	o.PushId = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *MarketBuyOrderFullyFilledResponse) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketBuyOrderFullyFilledResponse) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *MarketBuyOrderFullyFilledResponse) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *MarketBuyOrderFullyFilledResponse) SetAmount(v string) {
	o.Amount = &v
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *MarketBuyOrderFullyFilledResponse) GetTopic() string {
	if o == nil || common.IsNil(o.Topic) {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketBuyOrderFullyFilledResponse) GetTopicOk() (*string, bool) {
	if o == nil || common.IsNil(o.Topic) {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *MarketBuyOrderFullyFilledResponse) HasTopic() bool {
	if o != nil && !common.IsNil(o.Topic) {
		return true
	}

	return false
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *MarketBuyOrderFullyFilledResponse) SetTopic(v string) {
	o.Topic = &v
}

func (o MarketBuyOrderFullyFilledResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarketBuyOrderFullyFilledResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.PushId) {
		toSerialize["pushId"] = o.PushId
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.Topic) {
		toSerialize["topic"] = o.Topic
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MarketBuyOrderFullyFilledResponse) UnmarshalJSON(data []byte) (err error) {
	varMarketBuyOrderFullyFilledResponse := _MarketBuyOrderFullyFilledResponse{}

	err = json.Unmarshal(data, &varMarketBuyOrderFullyFilledResponse)

	if err != nil {
		return err
	}

	*o = MarketBuyOrderFullyFilledResponse(varMarketBuyOrderFullyFilledResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pushId")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "topic")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarketBuyOrderFullyFilledResponse struct {
	value *MarketBuyOrderFullyFilledResponse
	isSet bool
}

func (v NullableMarketBuyOrderFullyFilledResponse) Get() *MarketBuyOrderFullyFilledResponse {
	return v.value
}

func (v *NullableMarketBuyOrderFullyFilledResponse) Set(val *MarketBuyOrderFullyFilledResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketBuyOrderFullyFilledResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketBuyOrderFullyFilledResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketBuyOrderFullyFilledResponse(val *MarketBuyOrderFullyFilledResponse) *NullableMarketBuyOrderFullyFilledResponse {
	return &NullableMarketBuyOrderFullyFilledResponse{value: val, isSet: true}
}

func (v NullableMarketBuyOrderFullyFilledResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketBuyOrderFullyFilledResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
