/*
Prediction Trading Streams

Binance Web3 Prediction pushes 14 kinds of real-time events over the Binance SApi WebSocket channel, covering market and limit trading, reward claims, fund transfers and market resolution, plus order book (bid/ask) quotes per market. Trading events go only to the user they belong to; order book quotes go to everyone subscribed. A message may occasionally be missed, order is not guaranteed, and nothing is stored while you are offline — after reconnecting, fetch the latest state from the REST API and drop duplicates using `pushId` (business events) or `marketId` + `updateTimestampMs` (order book).
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the LimitOrderFullyFilledResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &LimitOrderFullyFilledResponse{}

// LimitOrderFullyFilledResponse struct for LimitOrderFullyFilledResponse
type LimitOrderFullyFilledResponse struct {
	// Unique ID of this push. Use it to remove duplicates.
	PushId *string `json:"pushId,omitempty"`
	// Filled amount in USDT — paid for a buy, received after fees for a sell. Rounded down to 2 decimals; shown as `<0.01` if smaller.
	Amount *string `json:"amount,omitempty"`
	// Market title, cut to 12 characters with `..` added when longer; empty if unavailable.
	Topic                *string `json:"topic,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LimitOrderFullyFilledResponse LimitOrderFullyFilledResponse

// NewLimitOrderFullyFilledResponse instantiates a new LimitOrderFullyFilledResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLimitOrderFullyFilledResponse() *LimitOrderFullyFilledResponse {
	this := LimitOrderFullyFilledResponse{}
	return &this
}

// NewLimitOrderFullyFilledResponseWithDefaults instantiates a new LimitOrderFullyFilledResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLimitOrderFullyFilledResponseWithDefaults() *LimitOrderFullyFilledResponse {
	this := LimitOrderFullyFilledResponse{}
	return &this
}

// GetPushId returns the PushId field value if set, zero value otherwise.
func (o *LimitOrderFullyFilledResponse) GetPushId() string {
	if o == nil || common.IsNil(o.PushId) {
		var ret string
		return ret
	}
	return *o.PushId
}

// GetPushIdOk returns a tuple with the PushId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitOrderFullyFilledResponse) GetPushIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.PushId) {
		return nil, false
	}
	return o.PushId, true
}

// HasPushId returns a boolean if a field has been set.
func (o *LimitOrderFullyFilledResponse) HasPushId() bool {
	if o != nil && !common.IsNil(o.PushId) {
		return true
	}

	return false
}

// SetPushId gets a reference to the given string and assigns it to the PushId field.
func (o *LimitOrderFullyFilledResponse) SetPushId(v string) {
	o.PushId = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *LimitOrderFullyFilledResponse) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitOrderFullyFilledResponse) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *LimitOrderFullyFilledResponse) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *LimitOrderFullyFilledResponse) SetAmount(v string) {
	o.Amount = &v
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *LimitOrderFullyFilledResponse) GetTopic() string {
	if o == nil || common.IsNil(o.Topic) {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitOrderFullyFilledResponse) GetTopicOk() (*string, bool) {
	if o == nil || common.IsNil(o.Topic) {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *LimitOrderFullyFilledResponse) HasTopic() bool {
	if o != nil && !common.IsNil(o.Topic) {
		return true
	}

	return false
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *LimitOrderFullyFilledResponse) SetTopic(v string) {
	o.Topic = &v
}

func (o LimitOrderFullyFilledResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LimitOrderFullyFilledResponse) ToMap() (map[string]interface{}, error) {
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

func (o *LimitOrderFullyFilledResponse) UnmarshalJSON(data []byte) (err error) {
	varLimitOrderFullyFilledResponse := _LimitOrderFullyFilledResponse{}

	err = json.Unmarshal(data, &varLimitOrderFullyFilledResponse)

	if err != nil {
		return err
	}

	*o = LimitOrderFullyFilledResponse(varLimitOrderFullyFilledResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pushId")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "topic")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLimitOrderFullyFilledResponse struct {
	value *LimitOrderFullyFilledResponse
	isSet bool
}

func (v NullableLimitOrderFullyFilledResponse) Get() *LimitOrderFullyFilledResponse {
	return v.value
}

func (v *NullableLimitOrderFullyFilledResponse) Set(val *LimitOrderFullyFilledResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLimitOrderFullyFilledResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLimitOrderFullyFilledResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLimitOrderFullyFilledResponse(val *LimitOrderFullyFilledResponse) *NullableLimitOrderFullyFilledResponse {
	return &NullableLimitOrderFullyFilledResponse{value: val, isSet: true}
}

func (v NullableLimitOrderFullyFilledResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLimitOrderFullyFilledResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
