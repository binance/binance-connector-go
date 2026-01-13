/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner{}

// QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner struct for QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner
type QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner struct {
	Asset                *string `json:"asset,omitempty"`
	MarginBalance        *string `json:"marginBalance,omitempty"`
	WalletBalance        *string `json:"walletBalance,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner

// NewQueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner instantiates a new QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner() *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner {
	this := QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner{}
	return &this
}

// NewQueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInnerWithDefaults instantiates a new QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInnerWithDefaults() *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner {
	this := QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetMarginBalance returns the MarginBalance field value if set, zero value otherwise.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) GetMarginBalance() string {
	if o == nil || common.IsNil(o.MarginBalance) {
		var ret string
		return ret
	}
	return *o.MarginBalance
}

// GetMarginBalanceOk returns a tuple with the MarginBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) GetMarginBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarginBalance) {
		return nil, false
	}
	return o.MarginBalance, true
}

// HasMarginBalance returns a boolean if a field has been set.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) HasMarginBalance() bool {
	if o != nil && !common.IsNil(o.MarginBalance) {
		return true
	}

	return false
}

// SetMarginBalance gets a reference to the given string and assigns it to the MarginBalance field.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) SetMarginBalance(v string) {
	o.MarginBalance = &v
}

// GetWalletBalance returns the WalletBalance field value if set, zero value otherwise.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) GetWalletBalance() string {
	if o == nil || common.IsNil(o.WalletBalance) {
		var ret string
		return ret
	}
	return *o.WalletBalance
}

// GetWalletBalanceOk returns a tuple with the WalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) GetWalletBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.WalletBalance) {
		return nil, false
	}
	return o.WalletBalance, true
}

// HasWalletBalance returns a boolean if a field has been set.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) HasWalletBalance() bool {
	if o != nil && !common.IsNil(o.WalletBalance) {
		return true
	}

	return false
}

// SetWalletBalance gets a reference to the given string and assigns it to the WalletBalance field.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) SetWalletBalance(v string) {
	o.WalletBalance = &v
}

func (o QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) ToMap() (map[string]interface{}, error) {
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

func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) UnmarshalJSON(data []byte) (err error) {
	varQueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner := _QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner{}

	err = json.Unmarshal(data, &varQueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner)

	if err != nil {
		return err
	}

	*o = QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner(varQueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "marginBalance")
		delete(additionalProperties, "walletBalance")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner struct {
	value *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner
	isSet bool
}

func (v NullableQueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) Get() *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner {
	return v.value
}

func (v *NullableQueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) Set(val *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner(val *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) *NullableQueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner {
	return &NullableQueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner{value: val, isSet: true}
}

func (v NullableQueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
