/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner{}

// DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner struct for DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner
type DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner struct {
	Asset                *string `json:"asset,omitempty"`
	Borrowed             *string `json:"borrowed,omitempty"`
	Free                 *string `json:"free,omitempty"`
	Interest             *string `json:"interest,omitempty"`
	Locked               *string `json:"locked,omitempty"`
	NetAsset             *string `json:"netAsset,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner

// NewDailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner instantiates a new DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner() *DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner {
	this := DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner{}
	return &this
}

// NewDailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInnerWithDefaults instantiates a new DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInnerWithDefaults() *DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner {
	this := DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetBorrowed returns the Borrowed field value if set, zero value otherwise.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner) GetBorrowed() string {
	if o == nil || common.IsNil(o.Borrowed) {
		var ret string
		return ret
	}
	return *o.Borrowed
}

// GetBorrowedOk returns a tuple with the Borrowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner) GetBorrowedOk() (*string, bool) {
	if o == nil || common.IsNil(o.Borrowed) {
		return nil, false
	}
	return o.Borrowed, true
}

// HasBorrowed returns a boolean if a field has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner) HasBorrowed() bool {
	if o != nil && !common.IsNil(o.Borrowed) {
		return true
	}

	return false
}

// SetBorrowed gets a reference to the given string and assigns it to the Borrowed field.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner) SetBorrowed(v string) {
	o.Borrowed = &v
}

// GetFree returns the Free field value if set, zero value otherwise.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner) GetFree() string {
	if o == nil || common.IsNil(o.Free) {
		var ret string
		return ret
	}
	return *o.Free
}

// GetFreeOk returns a tuple with the Free field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner) GetFreeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Free) {
		return nil, false
	}
	return o.Free, true
}

// HasFree returns a boolean if a field has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner) HasFree() bool {
	if o != nil && !common.IsNil(o.Free) {
		return true
	}

	return false
}

// SetFree gets a reference to the given string and assigns it to the Free field.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner) SetFree(v string) {
	o.Free = &v
}

// GetInterest returns the Interest field value if set, zero value otherwise.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner) GetInterest() string {
	if o == nil || common.IsNil(o.Interest) {
		var ret string
		return ret
	}
	return *o.Interest
}

// GetInterestOk returns a tuple with the Interest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner) GetInterestOk() (*string, bool) {
	if o == nil || common.IsNil(o.Interest) {
		return nil, false
	}
	return o.Interest, true
}

// HasInterest returns a boolean if a field has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner) HasInterest() bool {
	if o != nil && !common.IsNil(o.Interest) {
		return true
	}

	return false
}

// SetInterest gets a reference to the given string and assigns it to the Interest field.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner) SetInterest(v string) {
	o.Interest = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner) GetLocked() string {
	if o == nil || common.IsNil(o.Locked) {
		var ret string
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner) GetLockedOk() (*string, bool) {
	if o == nil || common.IsNil(o.Locked) {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner) HasLocked() bool {
	if o != nil && !common.IsNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given string and assigns it to the Locked field.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner) SetLocked(v string) {
	o.Locked = &v
}

// GetNetAsset returns the NetAsset field value if set, zero value otherwise.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner) GetNetAsset() string {
	if o == nil || common.IsNil(o.NetAsset) {
		var ret string
		return ret
	}
	return *o.NetAsset
}

// GetNetAssetOk returns a tuple with the NetAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner) GetNetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.NetAsset) {
		return nil, false
	}
	return o.NetAsset, true
}

// HasNetAsset returns a boolean if a field has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner) HasNetAsset() bool {
	if o != nil && !common.IsNil(o.NetAsset) {
		return true
	}

	return false
}

// SetNetAsset gets a reference to the given string and assigns it to the NetAsset field.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner) SetNetAsset(v string) {
	o.NetAsset = &v
}

func (o DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Borrowed) {
		toSerialize["borrowed"] = o.Borrowed
	}
	if !common.IsNil(o.Free) {
		toSerialize["free"] = o.Free
	}
	if !common.IsNil(o.Interest) {
		toSerialize["interest"] = o.Interest
	}
	if !common.IsNil(o.Locked) {
		toSerialize["locked"] = o.Locked
	}
	if !common.IsNil(o.NetAsset) {
		toSerialize["netAsset"] = o.NetAsset
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner) UnmarshalJSON(data []byte) (err error) {
	varDailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner := _DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner{}

	err = json.Unmarshal(data, &varDailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner)

	if err != nil {
		return err
	}

	*o = DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner(varDailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "borrowed")
		delete(additionalProperties, "free")
		delete(additionalProperties, "interest")
		delete(additionalProperties, "locked")
		delete(additionalProperties, "netAsset")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner struct {
	value *DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner
	isSet bool
}

func (v NullableDailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner) Get() *DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner {
	return v.value
}

func (v *NullableDailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner) Set(val *DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner(val *DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner) *NullableDailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner {
	return &NullableDailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner{value: val, isSet: true}
}

func (v NullableDailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
