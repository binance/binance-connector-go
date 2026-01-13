/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the KeepaliveUserDataStreamResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &KeepaliveUserDataStreamResponse{}

// KeepaliveUserDataStreamResponse struct for KeepaliveUserDataStreamResponse
type KeepaliveUserDataStreamResponse struct {
	ListenKey            *string `json:"listenKey,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KeepaliveUserDataStreamResponse KeepaliveUserDataStreamResponse

// NewKeepaliveUserDataStreamResponse instantiates a new KeepaliveUserDataStreamResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeepaliveUserDataStreamResponse() *KeepaliveUserDataStreamResponse {
	this := KeepaliveUserDataStreamResponse{}
	return &this
}

// NewKeepaliveUserDataStreamResponseWithDefaults instantiates a new KeepaliveUserDataStreamResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeepaliveUserDataStreamResponseWithDefaults() *KeepaliveUserDataStreamResponse {
	this := KeepaliveUserDataStreamResponse{}
	return &this
}

// GetListenKey returns the ListenKey field value if set, zero value otherwise.
func (o *KeepaliveUserDataStreamResponse) GetListenKey() string {
	if o == nil || common.IsNil(o.ListenKey) {
		var ret string
		return ret
	}
	return *o.ListenKey
}

// GetListenKeyOk returns a tuple with the ListenKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeepaliveUserDataStreamResponse) GetListenKeyOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListenKey) {
		return nil, false
	}
	return o.ListenKey, true
}

// HasListenKey returns a boolean if a field has been set.
func (o *KeepaliveUserDataStreamResponse) HasListenKey() bool {
	if o != nil && !common.IsNil(o.ListenKey) {
		return true
	}

	return false
}

// SetListenKey gets a reference to the given string and assigns it to the ListenKey field.
func (o *KeepaliveUserDataStreamResponse) SetListenKey(v string) {
	o.ListenKey = &v
}

func (o KeepaliveUserDataStreamResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeepaliveUserDataStreamResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ListenKey) {
		toSerialize["listenKey"] = o.ListenKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KeepaliveUserDataStreamResponse) UnmarshalJSON(data []byte) (err error) {
	varKeepaliveUserDataStreamResponse := _KeepaliveUserDataStreamResponse{}

	err = json.Unmarshal(data, &varKeepaliveUserDataStreamResponse)

	if err != nil {
		return err
	}

	*o = KeepaliveUserDataStreamResponse(varKeepaliveUserDataStreamResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "listenKey")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKeepaliveUserDataStreamResponse struct {
	value *KeepaliveUserDataStreamResponse
	isSet bool
}

func (v NullableKeepaliveUserDataStreamResponse) Get() *KeepaliveUserDataStreamResponse {
	return v.value
}

func (v *NullableKeepaliveUserDataStreamResponse) Set(val *KeepaliveUserDataStreamResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeepaliveUserDataStreamResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeepaliveUserDataStreamResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeepaliveUserDataStreamResponse(val *KeepaliveUserDataStreamResponse) *NullableKeepaliveUserDataStreamResponse {
	return &NullableKeepaliveUserDataStreamResponse{value: val, isSet: true}
}

func (v NullableKeepaliveUserDataStreamResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeepaliveUserDataStreamResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
