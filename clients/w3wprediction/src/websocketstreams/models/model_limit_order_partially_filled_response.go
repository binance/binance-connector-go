/*
Prediction Trading Streams

Binance Web3 Prediction pushes 14 kinds of real-time events over the Binance SApi WebSocket channel, covering market and limit trading, reward claims, fund transfers and market resolution, plus order book (bid/ask) quotes per market. Trading events go only to the user they belong to; order book quotes go to everyone subscribed. A message may occasionally be missed, order is not guaranteed, and nothing is stored while you are offline — after reconnecting, fetch the latest state from the REST API and drop duplicates using `pushId` (business events) or `marketId` + `updateTimestampMs` (order book).
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the LimitOrderPartiallyFilledResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &LimitOrderPartiallyFilledResponse{}

// LimitOrderPartiallyFilledResponse struct for LimitOrderPartiallyFilledResponse
type LimitOrderPartiallyFilledResponse struct {
	// Unique ID of this push. Use it to remove duplicates.
	PushId *string `json:"pushId,omitempty"`
	// **Total** filled so far in USDT, not just this fill — paid for a buy, received after fees for a sell. Rounded down to 2 decimals; shown as `<0.01` if smaller.
	Amount *string `json:"amount,omitempty"`
	// Market title, cut to 12 characters with `..` added when longer; empty if unavailable.
	Topic                *string `json:"topic,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LimitOrderPartiallyFilledResponse LimitOrderPartiallyFilledResponse

// NewLimitOrderPartiallyFilledResponse instantiates a new LimitOrderPartiallyFilledResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLimitOrderPartiallyFilledResponse() *LimitOrderPartiallyFilledResponse {
	this := LimitOrderPartiallyFilledResponse{}
	return &this
}

// NewLimitOrderPartiallyFilledResponseWithDefaults instantiates a new LimitOrderPartiallyFilledResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLimitOrderPartiallyFilledResponseWithDefaults() *LimitOrderPartiallyFilledResponse {
	this := LimitOrderPartiallyFilledResponse{}
	return &this
}

// GetPushId returns the PushId field value if set, zero value otherwise.
func (o *LimitOrderPartiallyFilledResponse) GetPushId() string {
	if o == nil || common.IsNil(o.PushId) {
		var ret string
		return ret
	}
	return *o.PushId
}

// GetPushIdOk returns a tuple with the PushId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitOrderPartiallyFilledResponse) GetPushIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.PushId) {
		return nil, false
	}
	return o.PushId, true
}

// HasPushId returns a boolean if a field has been set.
func (o *LimitOrderPartiallyFilledResponse) HasPushId() bool {
	if o != nil && !common.IsNil(o.PushId) {
		return true
	}

	return false
}

// SetPushId gets a reference to the given string and assigns it to the PushId field.
func (o *LimitOrderPartiallyFilledResponse) SetPushId(v string) {
	o.PushId = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *LimitOrderPartiallyFilledResponse) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitOrderPartiallyFilledResponse) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *LimitOrderPartiallyFilledResponse) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *LimitOrderPartiallyFilledResponse) SetAmount(v string) {
	o.Amount = &v
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *LimitOrderPartiallyFilledResponse) GetTopic() string {
	if o == nil || common.IsNil(o.Topic) {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitOrderPartiallyFilledResponse) GetTopicOk() (*string, bool) {
	if o == nil || common.IsNil(o.Topic) {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *LimitOrderPartiallyFilledResponse) HasTopic() bool {
	if o != nil && !common.IsNil(o.Topic) {
		return true
	}

	return false
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *LimitOrderPartiallyFilledResponse) SetTopic(v string) {
	o.Topic = &v
}

func (o LimitOrderPartiallyFilledResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LimitOrderPartiallyFilledResponse) ToMap() (map[string]interface{}, error) {
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

func (o *LimitOrderPartiallyFilledResponse) UnmarshalJSON(data []byte) (err error) {
	varLimitOrderPartiallyFilledResponse := _LimitOrderPartiallyFilledResponse{}

	err = json.Unmarshal(data, &varLimitOrderPartiallyFilledResponse)

	if err != nil {
		return err
	}

	*o = LimitOrderPartiallyFilledResponse(varLimitOrderPartiallyFilledResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pushId")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "topic")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLimitOrderPartiallyFilledResponse struct {
	value *LimitOrderPartiallyFilledResponse
	isSet bool
}

func (v NullableLimitOrderPartiallyFilledResponse) Get() *LimitOrderPartiallyFilledResponse {
	return v.value
}

func (v *NullableLimitOrderPartiallyFilledResponse) Set(val *LimitOrderPartiallyFilledResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLimitOrderPartiallyFilledResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLimitOrderPartiallyFilledResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLimitOrderPartiallyFilledResponse(val *LimitOrderPartiallyFilledResponse) *NullableLimitOrderPartiallyFilledResponse {
	return &NullableLimitOrderPartiallyFilledResponse{value: val, isSet: true}
}

func (v NullableLimitOrderPartiallyFilledResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLimitOrderPartiallyFilledResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
