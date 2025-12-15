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

// checks if the QuerySubAccountAssetsAssetManagementResponseBalancesInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QuerySubAccountAssetsAssetManagementResponseBalancesInner{}

// QuerySubAccountAssetsAssetManagementResponseBalancesInner struct for QuerySubAccountAssetsAssetManagementResponseBalancesInner
type QuerySubAccountAssetsAssetManagementResponseBalancesInner struct {
	Freeze               *string `json:"freeze,omitempty"`
	Withdrawing          *string `json:"withdrawing,omitempty"`
	Asset                *string `json:"asset,omitempty"`
	Free                 *string `json:"free,omitempty"`
	Locked               *string `json:"locked,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuerySubAccountAssetsAssetManagementResponseBalancesInner QuerySubAccountAssetsAssetManagementResponseBalancesInner

// NewQuerySubAccountAssetsAssetManagementResponseBalancesInner instantiates a new QuerySubAccountAssetsAssetManagementResponseBalancesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuerySubAccountAssetsAssetManagementResponseBalancesInner() *QuerySubAccountAssetsAssetManagementResponseBalancesInner {
	this := QuerySubAccountAssetsAssetManagementResponseBalancesInner{}
	return &this
}

// NewQuerySubAccountAssetsAssetManagementResponseBalancesInnerWithDefaults instantiates a new QuerySubAccountAssetsAssetManagementResponseBalancesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuerySubAccountAssetsAssetManagementResponseBalancesInnerWithDefaults() *QuerySubAccountAssetsAssetManagementResponseBalancesInner {
	this := QuerySubAccountAssetsAssetManagementResponseBalancesInner{}
	return &this
}

// GetFreeze returns the Freeze field value if set, zero value otherwise.
func (o *QuerySubAccountAssetsAssetManagementResponseBalancesInner) GetFreeze() string {
	if o == nil || common.IsNil(o.Freeze) {
		var ret string
		return ret
	}
	return *o.Freeze
}

// GetFreezeOk returns a tuple with the Freeze field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountAssetsAssetManagementResponseBalancesInner) GetFreezeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Freeze) {
		return nil, false
	}
	return o.Freeze, true
}

// HasFreeze returns a boolean if a field has been set.
func (o *QuerySubAccountAssetsAssetManagementResponseBalancesInner) HasFreeze() bool {
	if o != nil && !common.IsNil(o.Freeze) {
		return true
	}

	return false
}

// SetFreeze gets a reference to the given string and assigns it to the Freeze field.
func (o *QuerySubAccountAssetsAssetManagementResponseBalancesInner) SetFreeze(v string) {
	o.Freeze = &v
}

// GetWithdrawing returns the Withdrawing field value if set, zero value otherwise.
func (o *QuerySubAccountAssetsAssetManagementResponseBalancesInner) GetWithdrawing() string {
	if o == nil || common.IsNil(o.Withdrawing) {
		var ret string
		return ret
	}
	return *o.Withdrawing
}

// GetWithdrawingOk returns a tuple with the Withdrawing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountAssetsAssetManagementResponseBalancesInner) GetWithdrawingOk() (*string, bool) {
	if o == nil || common.IsNil(o.Withdrawing) {
		return nil, false
	}
	return o.Withdrawing, true
}

// HasWithdrawing returns a boolean if a field has been set.
func (o *QuerySubAccountAssetsAssetManagementResponseBalancesInner) HasWithdrawing() bool {
	if o != nil && !common.IsNil(o.Withdrawing) {
		return true
	}

	return false
}

// SetWithdrawing gets a reference to the given string and assigns it to the Withdrawing field.
func (o *QuerySubAccountAssetsAssetManagementResponseBalancesInner) SetWithdrawing(v string) {
	o.Withdrawing = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *QuerySubAccountAssetsAssetManagementResponseBalancesInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountAssetsAssetManagementResponseBalancesInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *QuerySubAccountAssetsAssetManagementResponseBalancesInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *QuerySubAccountAssetsAssetManagementResponseBalancesInner) SetAsset(v string) {
	o.Asset = &v
}

// GetFree returns the Free field value if set, zero value otherwise.
func (o *QuerySubAccountAssetsAssetManagementResponseBalancesInner) GetFree() string {
	if o == nil || common.IsNil(o.Free) {
		var ret string
		return ret
	}
	return *o.Free
}

// GetFreeOk returns a tuple with the Free field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountAssetsAssetManagementResponseBalancesInner) GetFreeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Free) {
		return nil, false
	}
	return o.Free, true
}

// HasFree returns a boolean if a field has been set.
func (o *QuerySubAccountAssetsAssetManagementResponseBalancesInner) HasFree() bool {
	if o != nil && !common.IsNil(o.Free) {
		return true
	}

	return false
}

// SetFree gets a reference to the given string and assigns it to the Free field.
func (o *QuerySubAccountAssetsAssetManagementResponseBalancesInner) SetFree(v string) {
	o.Free = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *QuerySubAccountAssetsAssetManagementResponseBalancesInner) GetLocked() string {
	if o == nil || common.IsNil(o.Locked) {
		var ret string
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountAssetsAssetManagementResponseBalancesInner) GetLockedOk() (*string, bool) {
	if o == nil || common.IsNil(o.Locked) {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *QuerySubAccountAssetsAssetManagementResponseBalancesInner) HasLocked() bool {
	if o != nil && !common.IsNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given string and assigns it to the Locked field.
func (o *QuerySubAccountAssetsAssetManagementResponseBalancesInner) SetLocked(v string) {
	o.Locked = &v
}

func (o QuerySubAccountAssetsAssetManagementResponseBalancesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuerySubAccountAssetsAssetManagementResponseBalancesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Freeze) {
		toSerialize["freeze"] = o.Freeze
	}
	if !common.IsNil(o.Withdrawing) {
		toSerialize["withdrawing"] = o.Withdrawing
	}
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

func (o *QuerySubAccountAssetsAssetManagementResponseBalancesInner) UnmarshalJSON(data []byte) (err error) {
	varQuerySubAccountAssetsAssetManagementResponseBalancesInner := _QuerySubAccountAssetsAssetManagementResponseBalancesInner{}

	err = json.Unmarshal(data, &varQuerySubAccountAssetsAssetManagementResponseBalancesInner)

	if err != nil {
		return err
	}

	*o = QuerySubAccountAssetsAssetManagementResponseBalancesInner(varQuerySubAccountAssetsAssetManagementResponseBalancesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "freeze")
		delete(additionalProperties, "withdrawing")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "free")
		delete(additionalProperties, "locked")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuerySubAccountAssetsAssetManagementResponseBalancesInner struct {
	value *QuerySubAccountAssetsAssetManagementResponseBalancesInner
	isSet bool
}

func (v NullableQuerySubAccountAssetsAssetManagementResponseBalancesInner) Get() *QuerySubAccountAssetsAssetManagementResponseBalancesInner {
	return v.value
}

func (v *NullableQuerySubAccountAssetsAssetManagementResponseBalancesInner) Set(val *QuerySubAccountAssetsAssetManagementResponseBalancesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQuerySubAccountAssetsAssetManagementResponseBalancesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQuerySubAccountAssetsAssetManagementResponseBalancesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuerySubAccountAssetsAssetManagementResponseBalancesInner(val *QuerySubAccountAssetsAssetManagementResponseBalancesInner) *NullableQuerySubAccountAssetsAssetManagementResponseBalancesInner {
	return &NullableQuerySubAccountAssetsAssetManagementResponseBalancesInner{value: val, isSet: true}
}

func (v NullableQuerySubAccountAssetsAssetManagementResponseBalancesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuerySubAccountAssetsAssetManagementResponseBalancesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
