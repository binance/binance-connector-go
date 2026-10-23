/*
Prediction Trading Streams

Binance Web3 Prediction pushes 14 kinds of real-time events over the Binance SApi WebSocket channel, covering market and limit trading, reward claims, fund transfers and market resolution, plus order book (bid/ask) quotes per market. Trading events go only to the user they belong to; order book quotes go to everyone subscribed. A message may occasionally be missed, order is not guaranteed, and nothing is stored while you are offline — after reconnecting, fetch the latest state from the REST API and drop duplicates using `pushId` (business events) or `marketId` + `updateTimestampMs` (order book).
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the BatchClaimPartiallySucceededResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BatchClaimPartiallySucceededResponse{}

// BatchClaimPartiallySucceededResponse struct for BatchClaimPartiallySucceededResponse
type BatchClaimPartiallySucceededResponse struct {
	// Unique ID of this push. Use it to remove duplicates.
	PushId               *string `json:"pushId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BatchClaimPartiallySucceededResponse BatchClaimPartiallySucceededResponse

// NewBatchClaimPartiallySucceededResponse instantiates a new BatchClaimPartiallySucceededResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchClaimPartiallySucceededResponse() *BatchClaimPartiallySucceededResponse {
	this := BatchClaimPartiallySucceededResponse{}
	return &this
}

// NewBatchClaimPartiallySucceededResponseWithDefaults instantiates a new BatchClaimPartiallySucceededResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchClaimPartiallySucceededResponseWithDefaults() *BatchClaimPartiallySucceededResponse {
	this := BatchClaimPartiallySucceededResponse{}
	return &this
}

// GetPushId returns the PushId field value if set, zero value otherwise.
func (o *BatchClaimPartiallySucceededResponse) GetPushId() string {
	if o == nil || common.IsNil(o.PushId) {
		var ret string
		return ret
	}
	return *o.PushId
}

// GetPushIdOk returns a tuple with the PushId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchClaimPartiallySucceededResponse) GetPushIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.PushId) {
		return nil, false
	}
	return o.PushId, true
}

// HasPushId returns a boolean if a field has been set.
func (o *BatchClaimPartiallySucceededResponse) HasPushId() bool {
	if o != nil && !common.IsNil(o.PushId) {
		return true
	}

	return false
}

// SetPushId gets a reference to the given string and assigns it to the PushId field.
func (o *BatchClaimPartiallySucceededResponse) SetPushId(v string) {
	o.PushId = &v
}

func (o BatchClaimPartiallySucceededResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchClaimPartiallySucceededResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.PushId) {
		toSerialize["pushId"] = o.PushId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BatchClaimPartiallySucceededResponse) UnmarshalJSON(data []byte) (err error) {
	varBatchClaimPartiallySucceededResponse := _BatchClaimPartiallySucceededResponse{}

	err = json.Unmarshal(data, &varBatchClaimPartiallySucceededResponse)

	if err != nil {
		return err
	}

	*o = BatchClaimPartiallySucceededResponse(varBatchClaimPartiallySucceededResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pushId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBatchClaimPartiallySucceededResponse struct {
	value *BatchClaimPartiallySucceededResponse
	isSet bool
}

func (v NullableBatchClaimPartiallySucceededResponse) Get() *BatchClaimPartiallySucceededResponse {
	return v.value
}

func (v *NullableBatchClaimPartiallySucceededResponse) Set(val *BatchClaimPartiallySucceededResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchClaimPartiallySucceededResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchClaimPartiallySucceededResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchClaimPartiallySucceededResponse(val *BatchClaimPartiallySucceededResponse) *NullableBatchClaimPartiallySucceededResponse {
	return &NullableBatchClaimPartiallySucceededResponse{value: val, isSet: true}
}

func (v NullableBatchClaimPartiallySucceededResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchClaimPartiallySucceededResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
