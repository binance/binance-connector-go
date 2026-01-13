/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner{}

// DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner struct for DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner
type DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner struct {
	Asset                *string `json:"asset,omitempty"`
	MarginBalance        *string `json:"marginBalance,omitempty"`
	WalletBalance        *string `json:"walletBalance,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner

// NewDailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner instantiates a new DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner() *DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner {
	this := DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner{}
	return &this
}

// NewDailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInnerWithDefaults instantiates a new DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInnerWithDefaults() *DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner {
	this := DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetMarginBalance returns the MarginBalance field value if set, zero value otherwise.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) GetMarginBalance() string {
	if o == nil || common.IsNil(o.MarginBalance) {
		var ret string
		return ret
	}
	return *o.MarginBalance
}

// GetMarginBalanceOk returns a tuple with the MarginBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) GetMarginBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarginBalance) {
		return nil, false
	}
	return o.MarginBalance, true
}

// HasMarginBalance returns a boolean if a field has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) HasMarginBalance() bool {
	if o != nil && !common.IsNil(o.MarginBalance) {
		return true
	}

	return false
}

// SetMarginBalance gets a reference to the given string and assigns it to the MarginBalance field.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) SetMarginBalance(v string) {
	o.MarginBalance = &v
}

// GetWalletBalance returns the WalletBalance field value if set, zero value otherwise.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) GetWalletBalance() string {
	if o == nil || common.IsNil(o.WalletBalance) {
		var ret string
		return ret
	}
	return *o.WalletBalance
}

// GetWalletBalanceOk returns a tuple with the WalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) GetWalletBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.WalletBalance) {
		return nil, false
	}
	return o.WalletBalance, true
}

// HasWalletBalance returns a boolean if a field has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) HasWalletBalance() bool {
	if o != nil && !common.IsNil(o.WalletBalance) {
		return true
	}

	return false
}

// SetWalletBalance gets a reference to the given string and assigns it to the WalletBalance field.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) SetWalletBalance(v string) {
	o.WalletBalance = &v
}

func (o DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.MarginBalance) {
		toSerialize["marginBalance"] = o.MarginBalance
	}
	if !common.IsNil(o.WalletBalance) {
		toSerialize["walletBalance"] = o.WalletBalance
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) UnmarshalJSON(data []byte) (err error) {
	varDailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner := _DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner{}

	err = json.Unmarshal(data, &varDailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner)

	if err != nil {
		return err
	}

	*o = DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner(varDailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "marginBalance")
		delete(additionalProperties, "walletBalance")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner struct {
	value *DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner
	isSet bool
}

func (v NullableDailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) Get() *DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner {
	return v.value
}

func (v *NullableDailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) Set(val *DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner(val *DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) *NullableDailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner {
	return &NullableDailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner{value: val, isSet: true}
}

func (v NullableDailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
