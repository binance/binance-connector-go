/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QuerySubAccountAssetsResponseBalancesInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QuerySubAccountAssetsResponseBalancesInner{}

// QuerySubAccountAssetsResponseBalancesInner struct for QuerySubAccountAssetsResponseBalancesInner
type QuerySubAccountAssetsResponseBalancesInner struct {
	Freeze               *int64   `json:"freeze,omitempty"`
	Withdrawing          *int64   `json:"withdrawing,omitempty"`
	Asset                *string  `json:"asset,omitempty"`
	Free                 *float32 `json:"free,omitempty"`
	Locked               *int64   `json:"locked,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuerySubAccountAssetsResponseBalancesInner QuerySubAccountAssetsResponseBalancesInner

// NewQuerySubAccountAssetsResponseBalancesInner instantiates a new QuerySubAccountAssetsResponseBalancesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuerySubAccountAssetsResponseBalancesInner() *QuerySubAccountAssetsResponseBalancesInner {
	this := QuerySubAccountAssetsResponseBalancesInner{}
	return &this
}

// NewQuerySubAccountAssetsResponseBalancesInnerWithDefaults instantiates a new QuerySubAccountAssetsResponseBalancesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuerySubAccountAssetsResponseBalancesInnerWithDefaults() *QuerySubAccountAssetsResponseBalancesInner {
	this := QuerySubAccountAssetsResponseBalancesInner{}
	return &this
}

// GetFreeze returns the Freeze field value if set, zero value otherwise.
func (o *QuerySubAccountAssetsResponseBalancesInner) GetFreeze() int64 {
	if o == nil || common.IsNil(o.Freeze) {
		var ret int64
		return ret
	}
	return *o.Freeze
}

// GetFreezeOk returns a tuple with the Freeze field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountAssetsResponseBalancesInner) GetFreezeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Freeze) {
		return nil, false
	}
	return o.Freeze, true
}

// HasFreeze returns a boolean if a field has been set.
func (o *QuerySubAccountAssetsResponseBalancesInner) HasFreeze() bool {
	if o != nil && !common.IsNil(o.Freeze) {
		return true
	}

	return false
}

// SetFreeze gets a reference to the given int64 and assigns it to the Freeze field.
func (o *QuerySubAccountAssetsResponseBalancesInner) SetFreeze(v int64) {
	o.Freeze = &v
}

// GetWithdrawing returns the Withdrawing field value if set, zero value otherwise.
func (o *QuerySubAccountAssetsResponseBalancesInner) GetWithdrawing() int64 {
	if o == nil || common.IsNil(o.Withdrawing) {
		var ret int64
		return ret
	}
	return *o.Withdrawing
}

// GetWithdrawingOk returns a tuple with the Withdrawing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountAssetsResponseBalancesInner) GetWithdrawingOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Withdrawing) {
		return nil, false
	}
	return o.Withdrawing, true
}

// HasWithdrawing returns a boolean if a field has been set.
func (o *QuerySubAccountAssetsResponseBalancesInner) HasWithdrawing() bool {
	if o != nil && !common.IsNil(o.Withdrawing) {
		return true
	}

	return false
}

// SetWithdrawing gets a reference to the given int64 and assigns it to the Withdrawing field.
func (o *QuerySubAccountAssetsResponseBalancesInner) SetWithdrawing(v int64) {
	o.Withdrawing = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *QuerySubAccountAssetsResponseBalancesInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountAssetsResponseBalancesInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *QuerySubAccountAssetsResponseBalancesInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *QuerySubAccountAssetsResponseBalancesInner) SetAsset(v string) {
	o.Asset = &v
}

// GetFree returns the Free field value if set, zero value otherwise.
func (o *QuerySubAccountAssetsResponseBalancesInner) GetFree() float32 {
	if o == nil || common.IsNil(o.Free) {
		var ret float32
		return ret
	}
	return *o.Free
}

// GetFreeOk returns a tuple with the Free field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountAssetsResponseBalancesInner) GetFreeOk() (*float32, bool) {
	if o == nil || common.IsNil(o.Free) {
		return nil, false
	}
	return o.Free, true
}

// HasFree returns a boolean if a field has been set.
func (o *QuerySubAccountAssetsResponseBalancesInner) HasFree() bool {
	if o != nil && !common.IsNil(o.Free) {
		return true
	}

	return false
}

// SetFree gets a reference to the given float32 and assigns it to the Free field.
func (o *QuerySubAccountAssetsResponseBalancesInner) SetFree(v float32) {
	o.Free = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *QuerySubAccountAssetsResponseBalancesInner) GetLocked() int64 {
	if o == nil || common.IsNil(o.Locked) {
		var ret int64
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountAssetsResponseBalancesInner) GetLockedOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Locked) {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *QuerySubAccountAssetsResponseBalancesInner) HasLocked() bool {
	if o != nil && !common.IsNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given int64 and assigns it to the Locked field.
func (o *QuerySubAccountAssetsResponseBalancesInner) SetLocked(v int64) {
	o.Locked = &v
}

func (o QuerySubAccountAssetsResponseBalancesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuerySubAccountAssetsResponseBalancesInner) ToMap() (map[string]interface{}, error) {
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

func (o *QuerySubAccountAssetsResponseBalancesInner) UnmarshalJSON(data []byte) (err error) {
	varQuerySubAccountAssetsResponseBalancesInner := _QuerySubAccountAssetsResponseBalancesInner{}

	err = json.Unmarshal(data, &varQuerySubAccountAssetsResponseBalancesInner)

	if err != nil {
		return err
	}

	*o = QuerySubAccountAssetsResponseBalancesInner(varQuerySubAccountAssetsResponseBalancesInner)

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

type NullableQuerySubAccountAssetsResponseBalancesInner struct {
	value *QuerySubAccountAssetsResponseBalancesInner
	isSet bool
}

func (v NullableQuerySubAccountAssetsResponseBalancesInner) Get() *QuerySubAccountAssetsResponseBalancesInner {
	return v.value
}

func (v *NullableQuerySubAccountAssetsResponseBalancesInner) Set(val *QuerySubAccountAssetsResponseBalancesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQuerySubAccountAssetsResponseBalancesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQuerySubAccountAssetsResponseBalancesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuerySubAccountAssetsResponseBalancesInner(val *QuerySubAccountAssetsResponseBalancesInner) *NullableQuerySubAccountAssetsResponseBalancesInner {
	return &NullableQuerySubAccountAssetsResponseBalancesInner{value: val, isSet: true}
}

func (v NullableQuerySubAccountAssetsResponseBalancesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuerySubAccountAssetsResponseBalancesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
