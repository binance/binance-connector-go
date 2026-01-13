/*
Binance Derivatives Trading USDS Futures WebSocket API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the KeepaliveUserDataStreamResponseResult type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &KeepaliveUserDataStreamResponseResult{}

// KeepaliveUserDataStreamResponseResult struct for KeepaliveUserDataStreamResponseResult
type KeepaliveUserDataStreamResponseResult struct {
	ListenKey            *string `json:"listenKey,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KeepaliveUserDataStreamResponseResult KeepaliveUserDataStreamResponseResult

// NewKeepaliveUserDataStreamResponseResult instantiates a new KeepaliveUserDataStreamResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeepaliveUserDataStreamResponseResult() *KeepaliveUserDataStreamResponseResult {
	this := KeepaliveUserDataStreamResponseResult{}
	return &this
}

// NewKeepaliveUserDataStreamResponseResultWithDefaults instantiates a new KeepaliveUserDataStreamResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeepaliveUserDataStreamResponseResultWithDefaults() *KeepaliveUserDataStreamResponseResult {
	this := KeepaliveUserDataStreamResponseResult{}
	return &this
}

// GetListenKey returns the ListenKey field value if set, zero value otherwise.
func (o *KeepaliveUserDataStreamResponseResult) GetListenKey() string {
	if o == nil || common.IsNil(o.ListenKey) {
		var ret string
		return ret
	}
	return *o.ListenKey
}

// GetListenKeyOk returns a tuple with the ListenKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeepaliveUserDataStreamResponseResult) GetListenKeyOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListenKey) {
		return nil, false
	}
	return o.ListenKey, true
}

// HasListenKey returns a boolean if a field has been set.
func (o *KeepaliveUserDataStreamResponseResult) HasListenKey() bool {
	if o != nil && !common.IsNil(o.ListenKey) {
		return true
	}

	return false
}

// SetListenKey gets a reference to the given string and assigns it to the ListenKey field.
func (o *KeepaliveUserDataStreamResponseResult) SetListenKey(v string) {
	o.ListenKey = &v
}

func (o KeepaliveUserDataStreamResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeepaliveUserDataStreamResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ListenKey) {
		toSerialize["listenKey"] = o.ListenKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KeepaliveUserDataStreamResponseResult) UnmarshalJSON(data []byte) (err error) {
	varKeepaliveUserDataStreamResponseResult := _KeepaliveUserDataStreamResponseResult{}

	err = json.Unmarshal(data, &varKeepaliveUserDataStreamResponseResult)

	if err != nil {
		return err
	}

	*o = KeepaliveUserDataStreamResponseResult(varKeepaliveUserDataStreamResponseResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "listenKey")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKeepaliveUserDataStreamResponseResult struct {
	value *KeepaliveUserDataStreamResponseResult
	isSet bool
}

func (v NullableKeepaliveUserDataStreamResponseResult) Get() *KeepaliveUserDataStreamResponseResult {
	return v.value
}

func (v *NullableKeepaliveUserDataStreamResponseResult) Set(val *KeepaliveUserDataStreamResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableKeepaliveUserDataStreamResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableKeepaliveUserDataStreamResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeepaliveUserDataStreamResponseResult(val *KeepaliveUserDataStreamResponseResult) *NullableKeepaliveUserDataStreamResponseResult {
	return &NullableKeepaliveUserDataStreamResponseResult{value: val, isSet: true}
}

func (v NullableKeepaliveUserDataStreamResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeepaliveUserDataStreamResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
