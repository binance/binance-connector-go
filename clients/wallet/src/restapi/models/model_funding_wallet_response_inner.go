/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the FundingWalletResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FundingWalletResponseInner{}

// FundingWalletResponseInner struct for FundingWalletResponseInner
type FundingWalletResponseInner struct {
	Asset                *string `json:"asset,omitempty"`
	Free                 *string `json:"free,omitempty"`
	Locked               *string `json:"locked,omitempty"`
	Freeze               *string `json:"freeze,omitempty"`
	Withdrawing          *string `json:"withdrawing,omitempty"`
	BtcValuation         *string `json:"btcValuation,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FundingWalletResponseInner FundingWalletResponseInner

// NewFundingWalletResponseInner instantiates a new FundingWalletResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFundingWalletResponseInner() *FundingWalletResponseInner {
	this := FundingWalletResponseInner{}
	return &this
}

// NewFundingWalletResponseInnerWithDefaults instantiates a new FundingWalletResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFundingWalletResponseInnerWithDefaults() *FundingWalletResponseInner {
	this := FundingWalletResponseInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *FundingWalletResponseInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingWalletResponseInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *FundingWalletResponseInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *FundingWalletResponseInner) SetAsset(v string) {
	o.Asset = &v
}

// GetFree returns the Free field value if set, zero value otherwise.
func (o *FundingWalletResponseInner) GetFree() string {
	if o == nil || common.IsNil(o.Free) {
		var ret string
		return ret
	}
	return *o.Free
}

// GetFreeOk returns a tuple with the Free field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingWalletResponseInner) GetFreeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Free) {
		return nil, false
	}
	return o.Free, true
}

// HasFree returns a boolean if a field has been set.
func (o *FundingWalletResponseInner) HasFree() bool {
	if o != nil && !common.IsNil(o.Free) {
		return true
	}

	return false
}

// SetFree gets a reference to the given string and assigns it to the Free field.
func (o *FundingWalletResponseInner) SetFree(v string) {
	o.Free = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *FundingWalletResponseInner) GetLocked() string {
	if o == nil || common.IsNil(o.Locked) {
		var ret string
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingWalletResponseInner) GetLockedOk() (*string, bool) {
	if o == nil || common.IsNil(o.Locked) {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *FundingWalletResponseInner) HasLocked() bool {
	if o != nil && !common.IsNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given string and assigns it to the Locked field.
func (o *FundingWalletResponseInner) SetLocked(v string) {
	o.Locked = &v
}

// GetFreeze returns the Freeze field value if set, zero value otherwise.
func (o *FundingWalletResponseInner) GetFreeze() string {
	if o == nil || common.IsNil(o.Freeze) {
		var ret string
		return ret
	}
	return *o.Freeze
}

// GetFreezeOk returns a tuple with the Freeze field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingWalletResponseInner) GetFreezeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Freeze) {
		return nil, false
	}
	return o.Freeze, true
}

// HasFreeze returns a boolean if a field has been set.
func (o *FundingWalletResponseInner) HasFreeze() bool {
	if o != nil && !common.IsNil(o.Freeze) {
		return true
	}

	return false
}

// SetFreeze gets a reference to the given string and assigns it to the Freeze field.
func (o *FundingWalletResponseInner) SetFreeze(v string) {
	o.Freeze = &v
}

// GetWithdrawing returns the Withdrawing field value if set, zero value otherwise.
func (o *FundingWalletResponseInner) GetWithdrawing() string {
	if o == nil || common.IsNil(o.Withdrawing) {
		var ret string
		return ret
	}
	return *o.Withdrawing
}

// GetWithdrawingOk returns a tuple with the Withdrawing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingWalletResponseInner) GetWithdrawingOk() (*string, bool) {
	if o == nil || common.IsNil(o.Withdrawing) {
		return nil, false
	}
	return o.Withdrawing, true
}

// HasWithdrawing returns a boolean if a field has been set.
func (o *FundingWalletResponseInner) HasWithdrawing() bool {
	if o != nil && !common.IsNil(o.Withdrawing) {
		return true
	}

	return false
}

// SetWithdrawing gets a reference to the given string and assigns it to the Withdrawing field.
func (o *FundingWalletResponseInner) SetWithdrawing(v string) {
	o.Withdrawing = &v
}

// GetBtcValuation returns the BtcValuation field value if set, zero value otherwise.
func (o *FundingWalletResponseInner) GetBtcValuation() string {
	if o == nil || common.IsNil(o.BtcValuation) {
		var ret string
		return ret
	}
	return *o.BtcValuation
}

// GetBtcValuationOk returns a tuple with the BtcValuation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingWalletResponseInner) GetBtcValuationOk() (*string, bool) {
	if o == nil || common.IsNil(o.BtcValuation) {
		return nil, false
	}
	return o.BtcValuation, true
}

// HasBtcValuation returns a boolean if a field has been set.
func (o *FundingWalletResponseInner) HasBtcValuation() bool {
	if o != nil && !common.IsNil(o.BtcValuation) {
		return true
	}

	return false
}

// SetBtcValuation gets a reference to the given string and assigns it to the BtcValuation field.
func (o *FundingWalletResponseInner) SetBtcValuation(v string) {
	o.BtcValuation = &v
}

func (o FundingWalletResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FundingWalletResponseInner) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.BtcValuation) {
		toSerialize["btcValuation"] = o.BtcValuation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FundingWalletResponseInner) UnmarshalJSON(data []byte) (err error) {
	varFundingWalletResponseInner := _FundingWalletResponseInner{}

	err = json.Unmarshal(data, &varFundingWalletResponseInner)

	if err != nil {
		return err
	}

	*o = FundingWalletResponseInner(varFundingWalletResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "free")
		delete(additionalProperties, "locked")
		delete(additionalProperties, "freeze")
		delete(additionalProperties, "withdrawing")
		delete(additionalProperties, "btcValuation")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFundingWalletResponseInner struct {
	value *FundingWalletResponseInner
	isSet bool
}

func (v NullableFundingWalletResponseInner) Get() *FundingWalletResponseInner {
	return v.value
}

func (v *NullableFundingWalletResponseInner) Set(val *FundingWalletResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableFundingWalletResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableFundingWalletResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFundingWalletResponseInner(val *FundingWalletResponseInner) *NullableFundingWalletResponseInner {
	return &NullableFundingWalletResponseInner{value: val, isSet: true}
}

func (v NullableFundingWalletResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFundingWalletResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
