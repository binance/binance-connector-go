/*
Binance Derivatives Trading COIN Futures WebSocket API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures WebSocket API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the StartUserDataStreamResponseResult type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &StartUserDataStreamResponseResult{}

// StartUserDataStreamResponseResult struct for StartUserDataStreamResponseResult
type StartUserDataStreamResponseResult struct {
	ListenKey            *string `json:"listenKey,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StartUserDataStreamResponseResult StartUserDataStreamResponseResult

// NewStartUserDataStreamResponseResult instantiates a new StartUserDataStreamResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartUserDataStreamResponseResult() *StartUserDataStreamResponseResult {
	this := StartUserDataStreamResponseResult{}
	return &this
}

// NewStartUserDataStreamResponseResultWithDefaults instantiates a new StartUserDataStreamResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartUserDataStreamResponseResultWithDefaults() *StartUserDataStreamResponseResult {
	this := StartUserDataStreamResponseResult{}
	return &this
}

// GetListenKey returns the ListenKey field value if set, zero value otherwise.
func (o *StartUserDataStreamResponseResult) GetListenKey() string {
	if o == nil || common.IsNil(o.ListenKey) {
		var ret string
		return ret
	}
	return *o.ListenKey
}

// GetListenKeyOk returns a tuple with the ListenKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartUserDataStreamResponseResult) GetListenKeyOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListenKey) {
		return nil, false
	}
	return o.ListenKey, true
}

// HasListenKey returns a boolean if a field has been set.
func (o *StartUserDataStreamResponseResult) HasListenKey() bool {
	if o != nil && !common.IsNil(o.ListenKey) {
		return true
	}

	return false
}

// SetListenKey gets a reference to the given string and assigns it to the ListenKey field.
func (o *StartUserDataStreamResponseResult) SetListenKey(v string) {
	o.ListenKey = &v
}

func (o StartUserDataStreamResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StartUserDataStreamResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ListenKey) {
		toSerialize["listenKey"] = o.ListenKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StartUserDataStreamResponseResult) UnmarshalJSON(data []byte) (err error) {
	varStartUserDataStreamResponseResult := _StartUserDataStreamResponseResult{}

	err = json.Unmarshal(data, &varStartUserDataStreamResponseResult)

	if err != nil {
		return err
	}

	*o = StartUserDataStreamResponseResult(varStartUserDataStreamResponseResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "listenKey")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStartUserDataStreamResponseResult struct {
	value *StartUserDataStreamResponseResult
	isSet bool
}

func (v NullableStartUserDataStreamResponseResult) Get() *StartUserDataStreamResponseResult {
	return v.value
}

func (v *NullableStartUserDataStreamResponseResult) Set(val *StartUserDataStreamResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableStartUserDataStreamResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableStartUserDataStreamResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartUserDataStreamResponseResult(val *StartUserDataStreamResponseResult) *NullableStartUserDataStreamResponseResult {
	return &NullableStartUserDataStreamResponseResult{value: val, isSet: true}
}

func (v NullableStartUserDataStreamResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartUserDataStreamResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
