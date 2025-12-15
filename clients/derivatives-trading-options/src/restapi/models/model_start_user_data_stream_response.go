/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the StartUserDataStreamResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &StartUserDataStreamResponse{}

// StartUserDataStreamResponse struct for StartUserDataStreamResponse
type StartUserDataStreamResponse struct {
	ListenKey            *string `json:"listenKey,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StartUserDataStreamResponse StartUserDataStreamResponse

// NewStartUserDataStreamResponse instantiates a new StartUserDataStreamResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartUserDataStreamResponse() *StartUserDataStreamResponse {
	this := StartUserDataStreamResponse{}
	return &this
}

// NewStartUserDataStreamResponseWithDefaults instantiates a new StartUserDataStreamResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartUserDataStreamResponseWithDefaults() *StartUserDataStreamResponse {
	this := StartUserDataStreamResponse{}
	return &this
}

// GetListenKey returns the ListenKey field value if set, zero value otherwise.
func (o *StartUserDataStreamResponse) GetListenKey() string {
	if o == nil || common.IsNil(o.ListenKey) {
		var ret string
		return ret
	}
	return *o.ListenKey
}

// GetListenKeyOk returns a tuple with the ListenKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartUserDataStreamResponse) GetListenKeyOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListenKey) {
		return nil, false
	}
	return o.ListenKey, true
}

// HasListenKey returns a boolean if a field has been set.
func (o *StartUserDataStreamResponse) HasListenKey() bool {
	if o != nil && !common.IsNil(o.ListenKey) {
		return true
	}

	return false
}

// SetListenKey gets a reference to the given string and assigns it to the ListenKey field.
func (o *StartUserDataStreamResponse) SetListenKey(v string) {
	o.ListenKey = &v
}

func (o StartUserDataStreamResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StartUserDataStreamResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ListenKey) {
		toSerialize["listenKey"] = o.ListenKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StartUserDataStreamResponse) UnmarshalJSON(data []byte) (err error) {
	varStartUserDataStreamResponse := _StartUserDataStreamResponse{}

	err = json.Unmarshal(data, &varStartUserDataStreamResponse)

	if err != nil {
		return err
	}

	*o = StartUserDataStreamResponse(varStartUserDataStreamResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "listenKey")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStartUserDataStreamResponse struct {
	value *StartUserDataStreamResponse
	isSet bool
}

func (v NullableStartUserDataStreamResponse) Get() *StartUserDataStreamResponse {
	return v.value
}

func (v *NullableStartUserDataStreamResponse) Set(val *StartUserDataStreamResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStartUserDataStreamResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStartUserDataStreamResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartUserDataStreamResponse(val *StartUserDataStreamResponse) *NullableStartUserDataStreamResponse {
	return &NullableStartUserDataStreamResponse{value: val, isSet: true}
}

func (v NullableStartUserDataStreamResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartUserDataStreamResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
