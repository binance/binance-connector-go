/*
Wallet REST API

Query balances, manage assets, and perform wallet operations via the Binance Wallet API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryUserWalletBalanceResponseInnerAssetBalancesInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryUserWalletBalanceResponseInnerAssetBalancesInner{}

// QueryUserWalletBalanceResponseInnerAssetBalancesInner struct for QueryUserWalletBalanceResponseInnerAssetBalancesInner
type QueryUserWalletBalanceResponseInnerAssetBalancesInner struct {
	Asset                *string `json:"asset,omitempty"`
	AssetName            *string `json:"assetName,omitempty"`
	Free                 *string `json:"free,omitempty"`
	Locked               *string `json:"locked,omitempty"`
	Freeze               *string `json:"freeze,omitempty"`
	Withdrawing          *string `json:"withdrawing,omitempty"`
	BtcValuation         *string `json:"btcValuation,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryUserWalletBalanceResponseInnerAssetBalancesInner QueryUserWalletBalanceResponseInnerAssetBalancesInner

// NewQueryUserWalletBalanceResponseInnerAssetBalancesInner instantiates a new QueryUserWalletBalanceResponseInnerAssetBalancesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryUserWalletBalanceResponseInnerAssetBalancesInner() *QueryUserWalletBalanceResponseInnerAssetBalancesInner {
	this := QueryUserWalletBalanceResponseInnerAssetBalancesInner{}
	return &this
}

// NewQueryUserWalletBalanceResponseInnerAssetBalancesInnerWithDefaults instantiates a new QueryUserWalletBalanceResponseInnerAssetBalancesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryUserWalletBalanceResponseInnerAssetBalancesInnerWithDefaults() *QueryUserWalletBalanceResponseInnerAssetBalancesInner {
	this := QueryUserWalletBalanceResponseInnerAssetBalancesInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *QueryUserWalletBalanceResponseInnerAssetBalancesInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUserWalletBalanceResponseInnerAssetBalancesInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *QueryUserWalletBalanceResponseInnerAssetBalancesInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *QueryUserWalletBalanceResponseInnerAssetBalancesInner) SetAsset(v string) {
	o.Asset = &v
}

// GetAssetName returns the AssetName field value if set, zero value otherwise.
func (o *QueryUserWalletBalanceResponseInnerAssetBalancesInner) GetAssetName() string {
	if o == nil || common.IsNil(o.AssetName) {
		var ret string
		return ret
	}
	return *o.AssetName
}

// GetAssetNameOk returns a tuple with the AssetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUserWalletBalanceResponseInnerAssetBalancesInner) GetAssetNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.AssetName) {
		return nil, false
	}
	return o.AssetName, true
}

// HasAssetName returns a boolean if a field has been set.
func (o *QueryUserWalletBalanceResponseInnerAssetBalancesInner) HasAssetName() bool {
	if o != nil && !common.IsNil(o.AssetName) {
		return true
	}

	return false
}

// SetAssetName gets a reference to the given string and assigns it to the AssetName field.
func (o *QueryUserWalletBalanceResponseInnerAssetBalancesInner) SetAssetName(v string) {
	o.AssetName = &v
}

// GetFree returns the Free field value if set, zero value otherwise.
func (o *QueryUserWalletBalanceResponseInnerAssetBalancesInner) GetFree() string {
	if o == nil || common.IsNil(o.Free) {
		var ret string
		return ret
	}
	return *o.Free
}

// GetFreeOk returns a tuple with the Free field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUserWalletBalanceResponseInnerAssetBalancesInner) GetFreeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Free) {
		return nil, false
	}
	return o.Free, true
}

// HasFree returns a boolean if a field has been set.
func (o *QueryUserWalletBalanceResponseInnerAssetBalancesInner) HasFree() bool {
	if o != nil && !common.IsNil(o.Free) {
		return true
	}

	return false
}

// SetFree gets a reference to the given string and assigns it to the Free field.
func (o *QueryUserWalletBalanceResponseInnerAssetBalancesInner) SetFree(v string) {
	o.Free = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *QueryUserWalletBalanceResponseInnerAssetBalancesInner) GetLocked() string {
	if o == nil || common.IsNil(o.Locked) {
		var ret string
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUserWalletBalanceResponseInnerAssetBalancesInner) GetLockedOk() (*string, bool) {
	if o == nil || common.IsNil(o.Locked) {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *QueryUserWalletBalanceResponseInnerAssetBalancesInner) HasLocked() bool {
	if o != nil && !common.IsNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given string and assigns it to the Locked field.
func (o *QueryUserWalletBalanceResponseInnerAssetBalancesInner) SetLocked(v string) {
	o.Locked = &v
}

// GetFreeze returns the Freeze field value if set, zero value otherwise.
func (o *QueryUserWalletBalanceResponseInnerAssetBalancesInner) GetFreeze() string {
	if o == nil || common.IsNil(o.Freeze) {
		var ret string
		return ret
	}
	return *o.Freeze
}

// GetFreezeOk returns a tuple with the Freeze field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUserWalletBalanceResponseInnerAssetBalancesInner) GetFreezeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Freeze) {
		return nil, false
	}
	return o.Freeze, true
}

// HasFreeze returns a boolean if a field has been set.
func (o *QueryUserWalletBalanceResponseInnerAssetBalancesInner) HasFreeze() bool {
	if o != nil && !common.IsNil(o.Freeze) {
		return true
	}

	return false
}

// SetFreeze gets a reference to the given string and assigns it to the Freeze field.
func (o *QueryUserWalletBalanceResponseInnerAssetBalancesInner) SetFreeze(v string) {
	o.Freeze = &v
}

// GetWithdrawing returns the Withdrawing field value if set, zero value otherwise.
func (o *QueryUserWalletBalanceResponseInnerAssetBalancesInner) GetWithdrawing() string {
	if o == nil || common.IsNil(o.Withdrawing) {
		var ret string
		return ret
	}
	return *o.Withdrawing
}

// GetWithdrawingOk returns a tuple with the Withdrawing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUserWalletBalanceResponseInnerAssetBalancesInner) GetWithdrawingOk() (*string, bool) {
	if o == nil || common.IsNil(o.Withdrawing) {
		return nil, false
	}
	return o.Withdrawing, true
}

// HasWithdrawing returns a boolean if a field has been set.
func (o *QueryUserWalletBalanceResponseInnerAssetBalancesInner) HasWithdrawing() bool {
	if o != nil && !common.IsNil(o.Withdrawing) {
		return true
	}

	return false
}

// SetWithdrawing gets a reference to the given string and assigns it to the Withdrawing field.
func (o *QueryUserWalletBalanceResponseInnerAssetBalancesInner) SetWithdrawing(v string) {
	o.Withdrawing = &v
}

// GetBtcValuation returns the BtcValuation field value if set, zero value otherwise.
func (o *QueryUserWalletBalanceResponseInnerAssetBalancesInner) GetBtcValuation() string {
	if o == nil || common.IsNil(o.BtcValuation) {
		var ret string
		return ret
	}
	return *o.BtcValuation
}

// GetBtcValuationOk returns a tuple with the BtcValuation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUserWalletBalanceResponseInnerAssetBalancesInner) GetBtcValuationOk() (*string, bool) {
	if o == nil || common.IsNil(o.BtcValuation) {
		return nil, false
	}
	return o.BtcValuation, true
}

// HasBtcValuation returns a boolean if a field has been set.
func (o *QueryUserWalletBalanceResponseInnerAssetBalancesInner) HasBtcValuation() bool {
	if o != nil && !common.IsNil(o.BtcValuation) {
		return true
	}

	return false
}

// SetBtcValuation gets a reference to the given string and assigns it to the BtcValuation field.
func (o *QueryUserWalletBalanceResponseInnerAssetBalancesInner) SetBtcValuation(v string) {
	o.BtcValuation = &v
}

func (o QueryUserWalletBalanceResponseInnerAssetBalancesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryUserWalletBalanceResponseInnerAssetBalancesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.AssetName) {
		toSerialize["assetName"] = o.AssetName
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
	if !common.IsNil(o.BtcValuation) {
		toSerialize["btcValuation"] = o.BtcValuation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryUserWalletBalanceResponseInnerAssetBalancesInner) UnmarshalJSON(data []byte) (err error) {
	varQueryUserWalletBalanceResponseInnerAssetBalancesInner := _QueryUserWalletBalanceResponseInnerAssetBalancesInner{}

	err = json.Unmarshal(data, &varQueryUserWalletBalanceResponseInnerAssetBalancesInner)

	if err != nil {
		return err
	}

	*o = QueryUserWalletBalanceResponseInnerAssetBalancesInner(varQueryUserWalletBalanceResponseInnerAssetBalancesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "assetName")
		delete(additionalProperties, "free")
		delete(additionalProperties, "locked")
		delete(additionalProperties, "freeze")
		delete(additionalProperties, "withdrawing")
		delete(additionalProperties, "btcValuation")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryUserWalletBalanceResponseInnerAssetBalancesInner struct {
	value *QueryUserWalletBalanceResponseInnerAssetBalancesInner
	isSet bool
}

func (v NullableQueryUserWalletBalanceResponseInnerAssetBalancesInner) Get() *QueryUserWalletBalanceResponseInnerAssetBalancesInner {
	return v.value
}

func (v *NullableQueryUserWalletBalanceResponseInnerAssetBalancesInner) Set(val *QueryUserWalletBalanceResponseInnerAssetBalancesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryUserWalletBalanceResponseInnerAssetBalancesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryUserWalletBalanceResponseInnerAssetBalancesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryUserWalletBalanceResponseInnerAssetBalancesInner(val *QueryUserWalletBalanceResponseInnerAssetBalancesInner) *NullableQueryUserWalletBalanceResponseInnerAssetBalancesInner {
	return &NullableQueryUserWalletBalanceResponseInnerAssetBalancesInner{value: val, isSet: true}
}

func (v NullableQueryUserWalletBalanceResponseInnerAssetBalancesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryUserWalletBalanceResponseInnerAssetBalancesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
