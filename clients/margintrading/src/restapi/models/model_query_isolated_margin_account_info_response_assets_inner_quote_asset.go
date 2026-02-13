/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset{}

// QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset struct for QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset
type QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset struct {
	Asset                *string `json:"asset,omitempty"`
	BorrowEnabled        *bool   `json:"borrowEnabled,omitempty"`
	Borrowed             *string `json:"borrowed,omitempty"`
	Free                 *string `json:"free,omitempty"`
	Interest             *string `json:"interest,omitempty"`
	Locked               *string `json:"locked,omitempty"`
	NetAsset             *string `json:"netAsset,omitempty"`
	NetAssetOfBtc        *string `json:"netAssetOfBtc,omitempty"`
	RepayEnabled         *bool   `json:"repayEnabled,omitempty"`
	TotalAsset           *string `json:"totalAsset,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset

// NewQueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset instantiates a new QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset() *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset {
	this := QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset{}
	return &this
}

// NewQueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAssetWithDefaults instantiates a new QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAssetWithDefaults() *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset {
	this := QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) SetAsset(v string) {
	o.Asset = &v
}

// GetBorrowEnabled returns the BorrowEnabled field value if set, zero value otherwise.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) GetBorrowEnabled() bool {
	if o == nil || common.IsNil(o.BorrowEnabled) {
		var ret bool
		return ret
	}
	return *o.BorrowEnabled
}

// GetBorrowEnabledOk returns a tuple with the BorrowEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) GetBorrowEnabledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.BorrowEnabled) {
		return nil, false
	}
	return o.BorrowEnabled, true
}

// HasBorrowEnabled returns a boolean if a field has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) HasBorrowEnabled() bool {
	if o != nil && !common.IsNil(o.BorrowEnabled) {
		return true
	}

	return false
}

// SetBorrowEnabled gets a reference to the given bool and assigns it to the BorrowEnabled field.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) SetBorrowEnabled(v bool) {
	o.BorrowEnabled = &v
}

// GetBorrowed returns the Borrowed field value if set, zero value otherwise.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) GetBorrowed() string {
	if o == nil || common.IsNil(o.Borrowed) {
		var ret string
		return ret
	}
	return *o.Borrowed
}

// GetBorrowedOk returns a tuple with the Borrowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) GetBorrowedOk() (*string, bool) {
	if o == nil || common.IsNil(o.Borrowed) {
		return nil, false
	}
	return o.Borrowed, true
}

// HasBorrowed returns a boolean if a field has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) HasBorrowed() bool {
	if o != nil && !common.IsNil(o.Borrowed) {
		return true
	}

	return false
}

// SetBorrowed gets a reference to the given string and assigns it to the Borrowed field.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) SetBorrowed(v string) {
	o.Borrowed = &v
}

// GetFree returns the Free field value if set, zero value otherwise.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) GetFree() string {
	if o == nil || common.IsNil(o.Free) {
		var ret string
		return ret
	}
	return *o.Free
}

// GetFreeOk returns a tuple with the Free field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) GetFreeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Free) {
		return nil, false
	}
	return o.Free, true
}

// HasFree returns a boolean if a field has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) HasFree() bool {
	if o != nil && !common.IsNil(o.Free) {
		return true
	}

	return false
}

// SetFree gets a reference to the given string and assigns it to the Free field.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) SetFree(v string) {
	o.Free = &v
}

// GetInterest returns the Interest field value if set, zero value otherwise.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) GetInterest() string {
	if o == nil || common.IsNil(o.Interest) {
		var ret string
		return ret
	}
	return *o.Interest
}

// GetInterestOk returns a tuple with the Interest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) GetInterestOk() (*string, bool) {
	if o == nil || common.IsNil(o.Interest) {
		return nil, false
	}
	return o.Interest, true
}

// HasInterest returns a boolean if a field has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) HasInterest() bool {
	if o != nil && !common.IsNil(o.Interest) {
		return true
	}

	return false
}

// SetInterest gets a reference to the given string and assigns it to the Interest field.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) SetInterest(v string) {
	o.Interest = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) GetLocked() string {
	if o == nil || common.IsNil(o.Locked) {
		var ret string
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) GetLockedOk() (*string, bool) {
	if o == nil || common.IsNil(o.Locked) {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) HasLocked() bool {
	if o != nil && !common.IsNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given string and assigns it to the Locked field.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) SetLocked(v string) {
	o.Locked = &v
}

// GetNetAsset returns the NetAsset field value if set, zero value otherwise.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) GetNetAsset() string {
	if o == nil || common.IsNil(o.NetAsset) {
		var ret string
		return ret
	}
	return *o.NetAsset
}

// GetNetAssetOk returns a tuple with the NetAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) GetNetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.NetAsset) {
		return nil, false
	}
	return o.NetAsset, true
}

// HasNetAsset returns a boolean if a field has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) HasNetAsset() bool {
	if o != nil && !common.IsNil(o.NetAsset) {
		return true
	}

	return false
}

// SetNetAsset gets a reference to the given string and assigns it to the NetAsset field.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) SetNetAsset(v string) {
	o.NetAsset = &v
}

// GetNetAssetOfBtc returns the NetAssetOfBtc field value if set, zero value otherwise.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) GetNetAssetOfBtc() string {
	if o == nil || common.IsNil(o.NetAssetOfBtc) {
		var ret string
		return ret
	}
	return *o.NetAssetOfBtc
}

// GetNetAssetOfBtcOk returns a tuple with the NetAssetOfBtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) GetNetAssetOfBtcOk() (*string, bool) {
	if o == nil || common.IsNil(o.NetAssetOfBtc) {
		return nil, false
	}
	return o.NetAssetOfBtc, true
}

// HasNetAssetOfBtc returns a boolean if a field has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) HasNetAssetOfBtc() bool {
	if o != nil && !common.IsNil(o.NetAssetOfBtc) {
		return true
	}

	return false
}

// SetNetAssetOfBtc gets a reference to the given string and assigns it to the NetAssetOfBtc field.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) SetNetAssetOfBtc(v string) {
	o.NetAssetOfBtc = &v
}

// GetRepayEnabled returns the RepayEnabled field value if set, zero value otherwise.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) GetRepayEnabled() bool {
	if o == nil || common.IsNil(o.RepayEnabled) {
		var ret bool
		return ret
	}
	return *o.RepayEnabled
}

// GetRepayEnabledOk returns a tuple with the RepayEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) GetRepayEnabledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.RepayEnabled) {
		return nil, false
	}
	return o.RepayEnabled, true
}

// HasRepayEnabled returns a boolean if a field has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) HasRepayEnabled() bool {
	if o != nil && !common.IsNil(o.RepayEnabled) {
		return true
	}

	return false
}

// SetRepayEnabled gets a reference to the given bool and assigns it to the RepayEnabled field.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) SetRepayEnabled(v bool) {
	o.RepayEnabled = &v
}

// GetTotalAsset returns the TotalAsset field value if set, zero value otherwise.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) GetTotalAsset() string {
	if o == nil || common.IsNil(o.TotalAsset) {
		var ret string
		return ret
	}
	return *o.TotalAsset
}

// GetTotalAssetOk returns a tuple with the TotalAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) GetTotalAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalAsset) {
		return nil, false
	}
	return o.TotalAsset, true
}

// HasTotalAsset returns a boolean if a field has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) HasTotalAsset() bool {
	if o != nil && !common.IsNil(o.TotalAsset) {
		return true
	}

	return false
}

// SetTotalAsset gets a reference to the given string and assigns it to the TotalAsset field.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) SetTotalAsset(v string) {
	o.TotalAsset = &v
}

func (o QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.BorrowEnabled) {
		toSerialize["borrowEnabled"] = o.BorrowEnabled
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
	if !common.IsNil(o.NetAssetOfBtc) {
		toSerialize["netAssetOfBtc"] = o.NetAssetOfBtc
	}
	if !common.IsNil(o.RepayEnabled) {
		toSerialize["repayEnabled"] = o.RepayEnabled
	}
	if !common.IsNil(o.TotalAsset) {
		toSerialize["totalAsset"] = o.TotalAsset
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) UnmarshalJSON(data []byte) (err error) {
	varQueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset := _QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset{}

	err = json.Unmarshal(data, &varQueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset)

	if err != nil {
		return err
	}

	*o = QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset(varQueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "borrowEnabled")
		delete(additionalProperties, "borrowed")
		delete(additionalProperties, "free")
		delete(additionalProperties, "interest")
		delete(additionalProperties, "locked")
		delete(additionalProperties, "netAsset")
		delete(additionalProperties, "netAssetOfBtc")
		delete(additionalProperties, "repayEnabled")
		delete(additionalProperties, "totalAsset")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset struct {
	value *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset
	isSet bool
}

func (v NullableQueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) Get() *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset {
	return v.value
}

func (v *NullableQueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) Set(val *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset(val *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) *NullableQueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset {
	return &NullableQueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset{value: val, isSet: true}
}

func (v NullableQueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
