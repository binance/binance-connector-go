/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the UserAssetResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UserAssetResponseInner{}

// UserAssetResponseInner struct for UserAssetResponseInner
type UserAssetResponseInner struct {
	Asset                *string `json:"asset,omitempty"`
	Free                 *string `json:"free,omitempty"`
	Locked               *string `json:"locked,omitempty"`
	Freeze               *string `json:"freeze,omitempty"`
	Withdrawing          *string `json:"withdrawing,omitempty"`
	Ipoable              *string `json:"ipoable,omitempty"`
	BtcValuation         *string `json:"btcValuation,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserAssetResponseInner UserAssetResponseInner

// NewUserAssetResponseInner instantiates a new UserAssetResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAssetResponseInner() *UserAssetResponseInner {
	this := UserAssetResponseInner{}
	return &this
}

// NewUserAssetResponseInnerWithDefaults instantiates a new UserAssetResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAssetResponseInnerWithDefaults() *UserAssetResponseInner {
	this := UserAssetResponseInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *UserAssetResponseInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAssetResponseInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *UserAssetResponseInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *UserAssetResponseInner) SetAsset(v string) {
	o.Asset = &v
}

// GetFree returns the Free field value if set, zero value otherwise.
func (o *UserAssetResponseInner) GetFree() string {
	if o == nil || common.IsNil(o.Free) {
		var ret string
		return ret
	}
	return *o.Free
}

// GetFreeOk returns a tuple with the Free field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAssetResponseInner) GetFreeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Free) {
		return nil, false
	}
	return o.Free, true
}

// HasFree returns a boolean if a field has been set.
func (o *UserAssetResponseInner) HasFree() bool {
	if o != nil && !common.IsNil(o.Free) {
		return true
	}

	return false
}

// SetFree gets a reference to the given string and assigns it to the Free field.
func (o *UserAssetResponseInner) SetFree(v string) {
	o.Free = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *UserAssetResponseInner) GetLocked() string {
	if o == nil || common.IsNil(o.Locked) {
		var ret string
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAssetResponseInner) GetLockedOk() (*string, bool) {
	if o == nil || common.IsNil(o.Locked) {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *UserAssetResponseInner) HasLocked() bool {
	if o != nil && !common.IsNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given string and assigns it to the Locked field.
func (o *UserAssetResponseInner) SetLocked(v string) {
	o.Locked = &v
}

// GetFreeze returns the Freeze field value if set, zero value otherwise.
func (o *UserAssetResponseInner) GetFreeze() string {
	if o == nil || common.IsNil(o.Freeze) {
		var ret string
		return ret
	}
	return *o.Freeze
}

// GetFreezeOk returns a tuple with the Freeze field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAssetResponseInner) GetFreezeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Freeze) {
		return nil, false
	}
	return o.Freeze, true
}

// HasFreeze returns a boolean if a field has been set.
func (o *UserAssetResponseInner) HasFreeze() bool {
	if o != nil && !common.IsNil(o.Freeze) {
		return true
	}

	return false
}

// SetFreeze gets a reference to the given string and assigns it to the Freeze field.
func (o *UserAssetResponseInner) SetFreeze(v string) {
	o.Freeze = &v
}

// GetWithdrawing returns the Withdrawing field value if set, zero value otherwise.
func (o *UserAssetResponseInner) GetWithdrawing() string {
	if o == nil || common.IsNil(o.Withdrawing) {
		var ret string
		return ret
	}
	return *o.Withdrawing
}

// GetWithdrawingOk returns a tuple with the Withdrawing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAssetResponseInner) GetWithdrawingOk() (*string, bool) {
	if o == nil || common.IsNil(o.Withdrawing) {
		return nil, false
	}
	return o.Withdrawing, true
}

// HasWithdrawing returns a boolean if a field has been set.
func (o *UserAssetResponseInner) HasWithdrawing() bool {
	if o != nil && !common.IsNil(o.Withdrawing) {
		return true
	}

	return false
}

// SetWithdrawing gets a reference to the given string and assigns it to the Withdrawing field.
func (o *UserAssetResponseInner) SetWithdrawing(v string) {
	o.Withdrawing = &v
}

// GetIpoable returns the Ipoable field value if set, zero value otherwise.
func (o *UserAssetResponseInner) GetIpoable() string {
	if o == nil || common.IsNil(o.Ipoable) {
		var ret string
		return ret
	}
	return *o.Ipoable
}

// GetIpoableOk returns a tuple with the Ipoable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAssetResponseInner) GetIpoableOk() (*string, bool) {
	if o == nil || common.IsNil(o.Ipoable) {
		return nil, false
	}
	return o.Ipoable, true
}

// HasIpoable returns a boolean if a field has been set.
func (o *UserAssetResponseInner) HasIpoable() bool {
	if o != nil && !common.IsNil(o.Ipoable) {
		return true
	}

	return false
}

// SetIpoable gets a reference to the given string and assigns it to the Ipoable field.
func (o *UserAssetResponseInner) SetIpoable(v string) {
	o.Ipoable = &v
}

// GetBtcValuation returns the BtcValuation field value if set, zero value otherwise.
func (o *UserAssetResponseInner) GetBtcValuation() string {
	if o == nil || common.IsNil(o.BtcValuation) {
		var ret string
		return ret
	}
	return *o.BtcValuation
}

// GetBtcValuationOk returns a tuple with the BtcValuation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAssetResponseInner) GetBtcValuationOk() (*string, bool) {
	if o == nil || common.IsNil(o.BtcValuation) {
		return nil, false
	}
	return o.BtcValuation, true
}

// HasBtcValuation returns a boolean if a field has been set.
func (o *UserAssetResponseInner) HasBtcValuation() bool {
	if o != nil && !common.IsNil(o.BtcValuation) {
		return true
	}

	return false
}

// SetBtcValuation gets a reference to the given string and assigns it to the BtcValuation field.
func (o *UserAssetResponseInner) SetBtcValuation(v string) {
	o.BtcValuation = &v
}

func (o UserAssetResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserAssetResponseInner) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.Freeze) {
		toSerialize["freeze"] = o.Freeze
	}
	if !common.IsNil(o.Withdrawing) {
		toSerialize["withdrawing"] = o.Withdrawing
	}
	if !common.IsNil(o.Ipoable) {
		toSerialize["ipoable"] = o.Ipoable
	}
	if !common.IsNil(o.BtcValuation) {
		toSerialize["btcValuation"] = o.BtcValuation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserAssetResponseInner) UnmarshalJSON(data []byte) (err error) {
	varUserAssetResponseInner := _UserAssetResponseInner{}

	err = json.Unmarshal(data, &varUserAssetResponseInner)

	if err != nil {
		return err
	}

	*o = UserAssetResponseInner(varUserAssetResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "free")
		delete(additionalProperties, "locked")
		delete(additionalProperties, "freeze")
		delete(additionalProperties, "withdrawing")
		delete(additionalProperties, "ipoable")
		delete(additionalProperties, "btcValuation")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserAssetResponseInner struct {
	value *UserAssetResponseInner
	isSet bool
}

func (v NullableUserAssetResponseInner) Get() *UserAssetResponseInner {
	return v.value
}

func (v *NullableUserAssetResponseInner) Set(val *UserAssetResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAssetResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAssetResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAssetResponseInner(val *UserAssetResponseInner) *NullableUserAssetResponseInner {
	return &NullableUserAssetResponseInner{value: val, isSet: true}
}

func (v NullableUserAssetResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAssetResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
