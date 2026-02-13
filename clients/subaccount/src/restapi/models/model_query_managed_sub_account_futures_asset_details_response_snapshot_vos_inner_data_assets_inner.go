/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner{}

// QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner struct for QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner
type QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner struct {
	Asset                *string `json:"asset,omitempty"`
	MarginBalance        *int64  `json:"marginBalance,omitempty"`
	WalletBalance        *int64  `json:"walletBalance,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner

// NewQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner instantiates a new QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner() *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner {
	this := QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner{}
	return &this
}

// NewQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInnerWithDefaults instantiates a new QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInnerWithDefaults() *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner {
	this := QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetMarginBalance returns the MarginBalance field value if set, zero value otherwise.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner) GetMarginBalance() int64 {
	if o == nil || common.IsNil(o.MarginBalance) {
		var ret int64
		return ret
	}
	return *o.MarginBalance
}

// GetMarginBalanceOk returns a tuple with the MarginBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner) GetMarginBalanceOk() (*int64, bool) {
	if o == nil || common.IsNil(o.MarginBalance) {
		return nil, false
	}
	return o.MarginBalance, true
}

// HasMarginBalance returns a boolean if a field has been set.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner) HasMarginBalance() bool {
	if o != nil && !common.IsNil(o.MarginBalance) {
		return true
	}

	return false
}

// SetMarginBalance gets a reference to the given int64 and assigns it to the MarginBalance field.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner) SetMarginBalance(v int64) {
	o.MarginBalance = &v
}

// GetWalletBalance returns the WalletBalance field value if set, zero value otherwise.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner) GetWalletBalance() int64 {
	if o == nil || common.IsNil(o.WalletBalance) {
		var ret int64
		return ret
	}
	return *o.WalletBalance
}

// GetWalletBalanceOk returns a tuple with the WalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner) GetWalletBalanceOk() (*int64, bool) {
	if o == nil || common.IsNil(o.WalletBalance) {
		return nil, false
	}
	return o.WalletBalance, true
}

// HasWalletBalance returns a boolean if a field has been set.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner) HasWalletBalance() bool {
	if o != nil && !common.IsNil(o.WalletBalance) {
		return true
	}

	return false
}

// SetWalletBalance gets a reference to the given int64 and assigns it to the WalletBalance field.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner) SetWalletBalance(v int64) {
	o.WalletBalance = &v
}

func (o QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner) ToMap() (map[string]interface{}, error) {
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

func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner) UnmarshalJSON(data []byte) (err error) {
	varQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner := _QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner{}

	err = json.Unmarshal(data, &varQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner)

	if err != nil {
		return err
	}

	*o = QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner(varQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "marginBalance")
		delete(additionalProperties, "walletBalance")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner struct {
	value *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner
	isSet bool
}

func (v NullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner) Get() *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner {
	return v.value
}

func (v *NullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner) Set(val *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner(val *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner) *NullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner {
	return &NullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner{value: val, isSet: true}
}

func (v NullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
