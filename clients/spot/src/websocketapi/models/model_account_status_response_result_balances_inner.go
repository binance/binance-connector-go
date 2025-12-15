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

// checks if the AccountStatusResponseResultBalancesInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountStatusResponseResultBalancesInner{}

// AccountStatusResponseResultBalancesInner struct for AccountStatusResponseResultBalancesInner
type AccountStatusResponseResultBalancesInner struct {
	Asset                *string `json:"asset,omitempty"`
	Free                 *string `json:"free,omitempty"`
	Locked               *string `json:"locked,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountStatusResponseResultBalancesInner AccountStatusResponseResultBalancesInner

// NewAccountStatusResponseResultBalancesInner instantiates a new AccountStatusResponseResultBalancesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountStatusResponseResultBalancesInner() *AccountStatusResponseResultBalancesInner {
	this := AccountStatusResponseResultBalancesInner{}
	return &this
}

// NewAccountStatusResponseResultBalancesInnerWithDefaults instantiates a new AccountStatusResponseResultBalancesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountStatusResponseResultBalancesInnerWithDefaults() *AccountStatusResponseResultBalancesInner {
	this := AccountStatusResponseResultBalancesInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *AccountStatusResponseResultBalancesInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountStatusResponseResultBalancesInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *AccountStatusResponseResultBalancesInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *AccountStatusResponseResultBalancesInner) SetAsset(v string) {
	o.Asset = &v
}

// GetFree returns the Free field value if set, zero value otherwise.
func (o *AccountStatusResponseResultBalancesInner) GetFree() string {
	if o == nil || common.IsNil(o.Free) {
		var ret string
		return ret
	}
	return *o.Free
}

// GetFreeOk returns a tuple with the Free field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountStatusResponseResultBalancesInner) GetFreeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Free) {
		return nil, false
	}
	return o.Free, true
}

// HasFree returns a boolean if a field has been set.
func (o *AccountStatusResponseResultBalancesInner) HasFree() bool {
	if o != nil && !common.IsNil(o.Free) {
		return true
	}

	return false
}

// SetFree gets a reference to the given string and assigns it to the Free field.
func (o *AccountStatusResponseResultBalancesInner) SetFree(v string) {
	o.Free = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *AccountStatusResponseResultBalancesInner) GetLocked() string {
	if o == nil || common.IsNil(o.Locked) {
		var ret string
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountStatusResponseResultBalancesInner) GetLockedOk() (*string, bool) {
	if o == nil || common.IsNil(o.Locked) {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *AccountStatusResponseResultBalancesInner) HasLocked() bool {
	if o != nil && !common.IsNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given string and assigns it to the Locked field.
func (o *AccountStatusResponseResultBalancesInner) SetLocked(v string) {
	o.Locked = &v
}

func (o AccountStatusResponseResultBalancesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountStatusResponseResultBalancesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Free) {
		toSerialize["free"] = o.Free
	}
	if !common.IsNil(o.Locked) {
		toSerialize["locked"] = o.Locked
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountStatusResponseResultBalancesInner) UnmarshalJSON(data []byte) (err error) {
	varAccountStatusResponseResultBalancesInner := _AccountStatusResponseResultBalancesInner{}

	err = json.Unmarshal(data, &varAccountStatusResponseResultBalancesInner)

	if err != nil {
		return err
	}

	*o = AccountStatusResponseResultBalancesInner(varAccountStatusResponseResultBalancesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "free")
		delete(additionalProperties, "locked")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountStatusResponseResultBalancesInner struct {
	value *AccountStatusResponseResultBalancesInner
	isSet bool
}

func (v NullableAccountStatusResponseResultBalancesInner) Get() *AccountStatusResponseResultBalancesInner {
	return v.value
}

func (v *NullableAccountStatusResponseResultBalancesInner) Set(val *AccountStatusResponseResultBalancesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountStatusResponseResultBalancesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountStatusResponseResultBalancesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountStatusResponseResultBalancesInner(val *AccountStatusResponseResultBalancesInner) *NullableAccountStatusResponseResultBalancesInner {
	return &NullableAccountStatusResponseResultBalancesInner{value: val, isSet: true}
}

func (v NullableAccountStatusResponseResultBalancesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountStatusResponseResultBalancesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
