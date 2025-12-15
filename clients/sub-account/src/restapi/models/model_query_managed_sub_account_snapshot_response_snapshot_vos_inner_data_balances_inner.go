/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner{}

// QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner struct for QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner
type QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner struct {
	Asset                *string `json:"asset,omitempty"`
	Free                 *string `json:"free,omitempty"`
	Locked               *string `json:"locked,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner

// NewQueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner instantiates a new QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner() *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner {
	this := QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner{}
	return &this
}

// NewQueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInnerWithDefaults instantiates a new QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInnerWithDefaults() *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner {
	this := QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner) SetAsset(v string) {
	o.Asset = &v
}

// GetFree returns the Free field value if set, zero value otherwise.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner) GetFree() string {
	if o == nil || common.IsNil(o.Free) {
		var ret string
		return ret
	}
	return *o.Free
}

// GetFreeOk returns a tuple with the Free field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner) GetFreeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Free) {
		return nil, false
	}
	return o.Free, true
}

// HasFree returns a boolean if a field has been set.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner) HasFree() bool {
	if o != nil && !common.IsNil(o.Free) {
		return true
	}

	return false
}

// SetFree gets a reference to the given string and assigns it to the Free field.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner) SetFree(v string) {
	o.Free = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner) GetLocked() string {
	if o == nil || common.IsNil(o.Locked) {
		var ret string
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner) GetLockedOk() (*string, bool) {
	if o == nil || common.IsNil(o.Locked) {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner) HasLocked() bool {
	if o != nil && !common.IsNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given string and assigns it to the Locked field.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner) SetLocked(v string) {
	o.Locked = &v
}

func (o QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner) ToMap() (map[string]interface{}, error) {
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

func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner) UnmarshalJSON(data []byte) (err error) {
	varQueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner := _QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner{}

	err = json.Unmarshal(data, &varQueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner)

	if err != nil {
		return err
	}

	*o = QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner(varQueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "free")
		delete(additionalProperties, "locked")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner struct {
	value *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner
	isSet bool
}

func (v NullableQueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner) Get() *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner {
	return v.value
}

func (v *NullableQueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner) Set(val *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner(val *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner) *NullableQueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner {
	return &NullableQueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner{value: val, isSet: true}
}

func (v NullableQueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
