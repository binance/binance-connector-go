/*
Margin REST API

Access account information, borrow and repay assets, and trade with Binance Margin.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the StartUserDataStreamV2Response type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &StartUserDataStreamV2Response{}

// StartUserDataStreamV2Response struct for StartUserDataStreamV2Response
type StartUserDataStreamV2Response struct {
	// listen Key.
	ListenKey            *string `json:"listenKey,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StartUserDataStreamV2Response StartUserDataStreamV2Response

// NewStartUserDataStreamV2Response instantiates a new StartUserDataStreamV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartUserDataStreamV2Response() *StartUserDataStreamV2Response {
	this := StartUserDataStreamV2Response{}
	return &this
}

// NewStartUserDataStreamV2ResponseWithDefaults instantiates a new StartUserDataStreamV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartUserDataStreamV2ResponseWithDefaults() *StartUserDataStreamV2Response {
	this := StartUserDataStreamV2Response{}
	return &this
}

// GetListenKey returns the ListenKey field value if set, zero value otherwise.
func (o *StartUserDataStreamV2Response) GetListenKey() string {
	if o == nil || common.IsNil(o.ListenKey) {
		var ret string
		return ret
	}
	return *o.ListenKey
}

// GetListenKeyOk returns a tuple with the ListenKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartUserDataStreamV2Response) GetListenKeyOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListenKey) {
		return nil, false
	}
	return o.ListenKey, true
}

// HasListenKey returns a boolean if a field has been set.
func (o *StartUserDataStreamV2Response) HasListenKey() bool {
	if o != nil && !common.IsNil(o.ListenKey) {
		return true
	}

	return false
}

// SetListenKey gets a reference to the given string and assigns it to the ListenKey field.
func (o *StartUserDataStreamV2Response) SetListenKey(v string) {
	o.ListenKey = &v
}

func (o StartUserDataStreamV2Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StartUserDataStreamV2Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ListenKey) {
		toSerialize["listenKey"] = o.ListenKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StartUserDataStreamV2Response) UnmarshalJSON(data []byte) (err error) {
	varStartUserDataStreamV2Response := _StartUserDataStreamV2Response{}

	err = json.Unmarshal(data, &varStartUserDataStreamV2Response)

	if err != nil {
		return err
	}

	*o = StartUserDataStreamV2Response(varStartUserDataStreamV2Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "listenKey")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStartUserDataStreamV2Response struct {
	value *StartUserDataStreamV2Response
	isSet bool
}

func (v NullableStartUserDataStreamV2Response) Get() *StartUserDataStreamV2Response {
	return v.value
}

func (v *NullableStartUserDataStreamV2Response) Set(val *StartUserDataStreamV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableStartUserDataStreamV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableStartUserDataStreamV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartUserDataStreamV2Response(val *StartUserDataStreamV2Response) *NullableStartUserDataStreamV2Response {
	return &NullableStartUserDataStreamV2Response{value: val, isSet: true}
}

func (v NullableStartUserDataStreamV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartUserDataStreamV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
