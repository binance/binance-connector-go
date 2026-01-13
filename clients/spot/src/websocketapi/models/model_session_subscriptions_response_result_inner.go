/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the SessionSubscriptionsResponseResultInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SessionSubscriptionsResponseResultInner{}

// SessionSubscriptionsResponseResultInner struct for SessionSubscriptionsResponseResultInner
type SessionSubscriptionsResponseResultInner struct {
	SubscriptionId       *int64 `json:"subscriptionId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SessionSubscriptionsResponseResultInner SessionSubscriptionsResponseResultInner

// NewSessionSubscriptionsResponseResultInner instantiates a new SessionSubscriptionsResponseResultInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionSubscriptionsResponseResultInner() *SessionSubscriptionsResponseResultInner {
	this := SessionSubscriptionsResponseResultInner{}
	return &this
}

// NewSessionSubscriptionsResponseResultInnerWithDefaults instantiates a new SessionSubscriptionsResponseResultInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionSubscriptionsResponseResultInnerWithDefaults() *SessionSubscriptionsResponseResultInner {
	this := SessionSubscriptionsResponseResultInner{}
	return &this
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *SessionSubscriptionsResponseResultInner) GetSubscriptionId() int64 {
	if o == nil || common.IsNil(o.SubscriptionId) {
		var ret int64
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionSubscriptionsResponseResultInner) GetSubscriptionIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *SessionSubscriptionsResponseResultInner) HasSubscriptionId() bool {
	if o != nil && !common.IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given int64 and assigns it to the SubscriptionId field.
func (o *SessionSubscriptionsResponseResultInner) SetSubscriptionId(v int64) {
	o.SubscriptionId = &v
}

func (o SessionSubscriptionsResponseResultInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionSubscriptionsResponseResultInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SessionSubscriptionsResponseResultInner) UnmarshalJSON(data []byte) (err error) {
	varSessionSubscriptionsResponseResultInner := _SessionSubscriptionsResponseResultInner{}

	err = json.Unmarshal(data, &varSessionSubscriptionsResponseResultInner)

	if err != nil {
		return err
	}

	*o = SessionSubscriptionsResponseResultInner(varSessionSubscriptionsResponseResultInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "subscriptionId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSessionSubscriptionsResponseResultInner struct {
	value *SessionSubscriptionsResponseResultInner
	isSet bool
}

func (v NullableSessionSubscriptionsResponseResultInner) Get() *SessionSubscriptionsResponseResultInner {
	return v.value
}

func (v *NullableSessionSubscriptionsResponseResultInner) Set(val *SessionSubscriptionsResponseResultInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionSubscriptionsResponseResultInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionSubscriptionsResponseResultInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionSubscriptionsResponseResultInner(val *SessionSubscriptionsResponseResultInner) *NullableSessionSubscriptionsResponseResultInner {
	return &NullableSessionSubscriptionsResponseResultInner{value: val, isSet: true}
}

func (v NullableSessionSubscriptionsResponseResultInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionSubscriptionsResponseResultInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
