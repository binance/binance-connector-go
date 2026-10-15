/*
Prediction Trading Streams

Binance Web3 Prediction pushes 14 types of per-user business events in real time via the Binance SApi WebSocket channel, covering trading (market / limit), reward claim, fund transfer and market resolution, plus order book (bid/ask) quotes broadcast per topic. Business events are routed by exact `(topic, userId)` match; order book topics are broadcast to every subscriber. Delivery is at-most-once, ordering is not guaranteed, and offline messages are not cached — pull the latest state via the REST API after reconnecting and deduplicate on `pushId` (business events) or `marketId` + `updateTimestampMs` (order book).
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the MarketResolutionNotificationResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MarketResolutionNotificationResponse{}

// MarketResolutionNotificationResponse struct for MarketResolutionNotificationResponse
type MarketResolutionNotificationResponse struct {
	// Unique push ID in the form `pm_<refId>_<scenarioCode>_<8hex>`, present in every topic; `refId` is the `marketId`. Use it for client-side deduplication.
	PushId *string `json:"pushId,omitempty"`
	// Market title from `MarketTopic.title`, truncated to 12 chars with `..` appended when longer (max 14 chars); empty string when the title is missing.
	Topic                *string `json:"topic,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarketResolutionNotificationResponse MarketResolutionNotificationResponse

// NewMarketResolutionNotificationResponse instantiates a new MarketResolutionNotificationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketResolutionNotificationResponse() *MarketResolutionNotificationResponse {
	this := MarketResolutionNotificationResponse{}
	return &this
}

// NewMarketResolutionNotificationResponseWithDefaults instantiates a new MarketResolutionNotificationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketResolutionNotificationResponseWithDefaults() *MarketResolutionNotificationResponse {
	this := MarketResolutionNotificationResponse{}
	return &this
}

// GetPushId returns the PushId field value if set, zero value otherwise.
func (o *MarketResolutionNotificationResponse) GetPushId() string {
	if o == nil || common.IsNil(o.PushId) {
		var ret string
		return ret
	}
	return *o.PushId
}

// GetPushIdOk returns a tuple with the PushId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketResolutionNotificationResponse) GetPushIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.PushId) {
		return nil, false
	}
	return o.PushId, true
}

// HasPushId returns a boolean if a field has been set.
func (o *MarketResolutionNotificationResponse) HasPushId() bool {
	if o != nil && !common.IsNil(o.PushId) {
		return true
	}

	return false
}

// SetPushId gets a reference to the given string and assigns it to the PushId field.
func (o *MarketResolutionNotificationResponse) SetPushId(v string) {
	o.PushId = &v
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *MarketResolutionNotificationResponse) GetTopic() string {
	if o == nil || common.IsNil(o.Topic) {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketResolutionNotificationResponse) GetTopicOk() (*string, bool) {
	if o == nil || common.IsNil(o.Topic) {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *MarketResolutionNotificationResponse) HasTopic() bool {
	if o != nil && !common.IsNil(o.Topic) {
		return true
	}

	return false
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *MarketResolutionNotificationResponse) SetTopic(v string) {
	o.Topic = &v
}

func (o MarketResolutionNotificationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarketResolutionNotificationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.PushId) {
		toSerialize["pushId"] = o.PushId
	}
	if !common.IsNil(o.Topic) {
		toSerialize["topic"] = o.Topic
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MarketResolutionNotificationResponse) UnmarshalJSON(data []byte) (err error) {
	varMarketResolutionNotificationResponse := _MarketResolutionNotificationResponse{}

	err = json.Unmarshal(data, &varMarketResolutionNotificationResponse)

	if err != nil {
		return err
	}

	*o = MarketResolutionNotificationResponse(varMarketResolutionNotificationResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pushId")
		delete(additionalProperties, "topic")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarketResolutionNotificationResponse struct {
	value *MarketResolutionNotificationResponse
	isSet bool
}

func (v NullableMarketResolutionNotificationResponse) Get() *MarketResolutionNotificationResponse {
	return v.value
}

func (v *NullableMarketResolutionNotificationResponse) Set(val *MarketResolutionNotificationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketResolutionNotificationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketResolutionNotificationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketResolutionNotificationResponse(val *MarketResolutionNotificationResponse) *NullableMarketResolutionNotificationResponse {
	return &NullableMarketResolutionNotificationResponse{value: val, isSet: true}
}

func (v NullableMarketResolutionNotificationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketResolutionNotificationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
