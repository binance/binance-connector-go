/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the UserDataStreamSubscribeResponseResult type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UserDataStreamSubscribeResponseResult{}

// UserDataStreamSubscribeResponseResult struct for UserDataStreamSubscribeResponseResult
type UserDataStreamSubscribeResponseResult struct {
	SubscriptionId       *int64 `json:"subscriptionId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserDataStreamSubscribeResponseResult UserDataStreamSubscribeResponseResult

// NewUserDataStreamSubscribeResponseResult instantiates a new UserDataStreamSubscribeResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserDataStreamSubscribeResponseResult() *UserDataStreamSubscribeResponseResult {
	this := UserDataStreamSubscribeResponseResult{}
	return &this
}

// NewUserDataStreamSubscribeResponseResultWithDefaults instantiates a new UserDataStreamSubscribeResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDataStreamSubscribeResponseResultWithDefaults() *UserDataStreamSubscribeResponseResult {
	this := UserDataStreamSubscribeResponseResult{}
	return &this
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *UserDataStreamSubscribeResponseResult) GetSubscriptionId() int64 {
	if o == nil || common.IsNil(o.SubscriptionId) {
		var ret int64
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDataStreamSubscribeResponseResult) GetSubscriptionIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *UserDataStreamSubscribeResponseResult) HasSubscriptionId() bool {
	if o != nil && !common.IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given int64 and assigns it to the SubscriptionId field.
func (o *UserDataStreamSubscribeResponseResult) SetSubscriptionId(v int64) {
	o.SubscriptionId = &v
}

func (o UserDataStreamSubscribeResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserDataStreamSubscribeResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserDataStreamSubscribeResponseResult) UnmarshalJSON(data []byte) (err error) {
	varUserDataStreamSubscribeResponseResult := _UserDataStreamSubscribeResponseResult{}

	err = json.Unmarshal(data, &varUserDataStreamSubscribeResponseResult)

	if err != nil {
		return err
	}

	*o = UserDataStreamSubscribeResponseResult(varUserDataStreamSubscribeResponseResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "subscriptionId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserDataStreamSubscribeResponseResult struct {
	value *UserDataStreamSubscribeResponseResult
	isSet bool
}

func (v NullableUserDataStreamSubscribeResponseResult) Get() *UserDataStreamSubscribeResponseResult {
	return v.value
}

func (v *NullableUserDataStreamSubscribeResponseResult) Set(val *UserDataStreamSubscribeResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDataStreamSubscribeResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableUserDataStreamSubscribeResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDataStreamSubscribeResponseResult(val *UserDataStreamSubscribeResponseResult) *NullableUserDataStreamSubscribeResponseResult {
	return &NullableUserDataStreamSubscribeResponseResult{value: val, isSet: true}
}

func (v NullableUserDataStreamSubscribeResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDataStreamSubscribeResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
