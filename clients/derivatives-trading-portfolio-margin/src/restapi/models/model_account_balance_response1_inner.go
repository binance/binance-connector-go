/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AccountBalanceResponse1Inner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountBalanceResponse1Inner{}

// AccountBalanceResponse1Inner struct for AccountBalanceResponse1Inner
type AccountBalanceResponse1Inner struct {
	Asset                *string `json:"asset,omitempty"`
	TotalWalletBalance   *string `json:"totalWalletBalance,omitempty"`
	CrossMarginAsset     *string `json:"crossMarginAsset,omitempty"`
	CrossMarginBorrowed  *string `json:"crossMarginBorrowed,omitempty"`
	CrossMarginFree      *string `json:"crossMarginFree,omitempty"`
	CrossMarginInterest  *string `json:"crossMarginInterest,omitempty"`
	CrossMarginLocked    *string `json:"crossMarginLocked,omitempty"`
	UmWalletBalance      *string `json:"umWalletBalance,omitempty"`
	UmUnrealizedPNL      *string `json:"umUnrealizedPNL,omitempty"`
	CmWalletBalance      *string `json:"cmWalletBalance,omitempty"`
	CmUnrealizedPNL      *string `json:"cmUnrealizedPNL,omitempty"`
	UpdateTime           *int64  `json:"updateTime,omitempty"`
	NegativeBalance      *string `json:"negativeBalance,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountBalanceResponse1Inner AccountBalanceResponse1Inner

// NewAccountBalanceResponse1Inner instantiates a new AccountBalanceResponse1Inner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountBalanceResponse1Inner() *AccountBalanceResponse1Inner {
	this := AccountBalanceResponse1Inner{}
	return &this
}

// NewAccountBalanceResponse1InnerWithDefaults instantiates a new AccountBalanceResponse1Inner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountBalanceResponse1InnerWithDefaults() *AccountBalanceResponse1Inner {
	this := AccountBalanceResponse1Inner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *AccountBalanceResponse1Inner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBalanceResponse1Inner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *AccountBalanceResponse1Inner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *AccountBalanceResponse1Inner) SetAsset(v string) {
	o.Asset = &v
}

// GetTotalWalletBalance returns the TotalWalletBalance field value if set, zero value otherwise.
func (o *AccountBalanceResponse1Inner) GetTotalWalletBalance() string {
	if o == nil || common.IsNil(o.TotalWalletBalance) {
		var ret string
		return ret
	}
	return *o.TotalWalletBalance
}

// GetTotalWalletBalanceOk returns a tuple with the TotalWalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBalanceResponse1Inner) GetTotalWalletBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalWalletBalance) {
		return nil, false
	}
	return o.TotalWalletBalance, true
}

// HasTotalWalletBalance returns a boolean if a field has been set.
func (o *AccountBalanceResponse1Inner) HasTotalWalletBalance() bool {
	if o != nil && !common.IsNil(o.TotalWalletBalance) {
		return true
	}

	return false
}

// SetTotalWalletBalance gets a reference to the given string and assigns it to the TotalWalletBalance field.
func (o *AccountBalanceResponse1Inner) SetTotalWalletBalance(v string) {
	o.TotalWalletBalance = &v
}

// GetCrossMarginAsset returns the CrossMarginAsset field value if set, zero value otherwise.
func (o *AccountBalanceResponse1Inner) GetCrossMarginAsset() string {
	if o == nil || common.IsNil(o.CrossMarginAsset) {
		var ret string
		return ret
	}
	return *o.CrossMarginAsset
}

// GetCrossMarginAssetOk returns a tuple with the CrossMarginAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBalanceResponse1Inner) GetCrossMarginAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.CrossMarginAsset) {
		return nil, false
	}
	return o.CrossMarginAsset, true
}

// HasCrossMarginAsset returns a boolean if a field has been set.
func (o *AccountBalanceResponse1Inner) HasCrossMarginAsset() bool {
	if o != nil && !common.IsNil(o.CrossMarginAsset) {
		return true
	}

	return false
}

// SetCrossMarginAsset gets a reference to the given string and assigns it to the CrossMarginAsset field.
func (o *AccountBalanceResponse1Inner) SetCrossMarginAsset(v string) {
	o.CrossMarginAsset = &v
}

// GetCrossMarginBorrowed returns the CrossMarginBorrowed field value if set, zero value otherwise.
func (o *AccountBalanceResponse1Inner) GetCrossMarginBorrowed() string {
	if o == nil || common.IsNil(o.CrossMarginBorrowed) {
		var ret string
		return ret
	}
	return *o.CrossMarginBorrowed
}

// GetCrossMarginBorrowedOk returns a tuple with the CrossMarginBorrowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBalanceResponse1Inner) GetCrossMarginBorrowedOk() (*string, bool) {
	if o == nil || common.IsNil(o.CrossMarginBorrowed) {
		return nil, false
	}
	return o.CrossMarginBorrowed, true
}

// HasCrossMarginBorrowed returns a boolean if a field has been set.
func (o *AccountBalanceResponse1Inner) HasCrossMarginBorrowed() bool {
	if o != nil && !common.IsNil(o.CrossMarginBorrowed) {
		return true
	}

	return false
}

// SetCrossMarginBorrowed gets a reference to the given string and assigns it to the CrossMarginBorrowed field.
func (o *AccountBalanceResponse1Inner) SetCrossMarginBorrowed(v string) {
	o.CrossMarginBorrowed = &v
}

// GetCrossMarginFree returns the CrossMarginFree field value if set, zero value otherwise.
func (o *AccountBalanceResponse1Inner) GetCrossMarginFree() string {
	if o == nil || common.IsNil(o.CrossMarginFree) {
		var ret string
		return ret
	}
	return *o.CrossMarginFree
}

// GetCrossMarginFreeOk returns a tuple with the CrossMarginFree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBalanceResponse1Inner) GetCrossMarginFreeOk() (*string, bool) {
	if o == nil || common.IsNil(o.CrossMarginFree) {
		return nil, false
	}
	return o.CrossMarginFree, true
}

// HasCrossMarginFree returns a boolean if a field has been set.
func (o *AccountBalanceResponse1Inner) HasCrossMarginFree() bool {
	if o != nil && !common.IsNil(o.CrossMarginFree) {
		return true
	}

	return false
}

// SetCrossMarginFree gets a reference to the given string and assigns it to the CrossMarginFree field.
func (o *AccountBalanceResponse1Inner) SetCrossMarginFree(v string) {
	o.CrossMarginFree = &v
}

// GetCrossMarginInterest returns the CrossMarginInterest field value if set, zero value otherwise.
func (o *AccountBalanceResponse1Inner) GetCrossMarginInterest() string {
	if o == nil || common.IsNil(o.CrossMarginInterest) {
		var ret string
		return ret
	}
	return *o.CrossMarginInterest
}

// GetCrossMarginInterestOk returns a tuple with the CrossMarginInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBalanceResponse1Inner) GetCrossMarginInterestOk() (*string, bool) {
	if o == nil || common.IsNil(o.CrossMarginInterest) {
		return nil, false
	}
	return o.CrossMarginInterest, true
}

// HasCrossMarginInterest returns a boolean if a field has been set.
func (o *AccountBalanceResponse1Inner) HasCrossMarginInterest() bool {
	if o != nil && !common.IsNil(o.CrossMarginInterest) {
		return true
	}

	return false
}

// SetCrossMarginInterest gets a reference to the given string and assigns it to the CrossMarginInterest field.
func (o *AccountBalanceResponse1Inner) SetCrossMarginInterest(v string) {
	o.CrossMarginInterest = &v
}

// GetCrossMarginLocked returns the CrossMarginLocked field value if set, zero value otherwise.
func (o *AccountBalanceResponse1Inner) GetCrossMarginLocked() string {
	if o == nil || common.IsNil(o.CrossMarginLocked) {
		var ret string
		return ret
	}
	return *o.CrossMarginLocked
}

// GetCrossMarginLockedOk returns a tuple with the CrossMarginLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBalanceResponse1Inner) GetCrossMarginLockedOk() (*string, bool) {
	if o == nil || common.IsNil(o.CrossMarginLocked) {
		return nil, false
	}
	return o.CrossMarginLocked, true
}

// HasCrossMarginLocked returns a boolean if a field has been set.
func (o *AccountBalanceResponse1Inner) HasCrossMarginLocked() bool {
	if o != nil && !common.IsNil(o.CrossMarginLocked) {
		return true
	}

	return false
}

// SetCrossMarginLocked gets a reference to the given string and assigns it to the CrossMarginLocked field.
func (o *AccountBalanceResponse1Inner) SetCrossMarginLocked(v string) {
	o.CrossMarginLocked = &v
}

// GetUmWalletBalance returns the UmWalletBalance field value if set, zero value otherwise.
func (o *AccountBalanceResponse1Inner) GetUmWalletBalance() string {
	if o == nil || common.IsNil(o.UmWalletBalance) {
		var ret string
		return ret
	}
	return *o.UmWalletBalance
}

// GetUmWalletBalanceOk returns a tuple with the UmWalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBalanceResponse1Inner) GetUmWalletBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.UmWalletBalance) {
		return nil, false
	}
	return o.UmWalletBalance, true
}

// HasUmWalletBalance returns a boolean if a field has been set.
func (o *AccountBalanceResponse1Inner) HasUmWalletBalance() bool {
	if o != nil && !common.IsNil(o.UmWalletBalance) {
		return true
	}

	return false
}

// SetUmWalletBalance gets a reference to the given string and assigns it to the UmWalletBalance field.
func (o *AccountBalanceResponse1Inner) SetUmWalletBalance(v string) {
	o.UmWalletBalance = &v
}

// GetUmUnrealizedPNL returns the UmUnrealizedPNL field value if set, zero value otherwise.
func (o *AccountBalanceResponse1Inner) GetUmUnrealizedPNL() string {
	if o == nil || common.IsNil(o.UmUnrealizedPNL) {
		var ret string
		return ret
	}
	return *o.UmUnrealizedPNL
}

// GetUmUnrealizedPNLOk returns a tuple with the UmUnrealizedPNL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBalanceResponse1Inner) GetUmUnrealizedPNLOk() (*string, bool) {
	if o == nil || common.IsNil(o.UmUnrealizedPNL) {
		return nil, false
	}
	return o.UmUnrealizedPNL, true
}

// HasUmUnrealizedPNL returns a boolean if a field has been set.
func (o *AccountBalanceResponse1Inner) HasUmUnrealizedPNL() bool {
	if o != nil && !common.IsNil(o.UmUnrealizedPNL) {
		return true
	}

	return false
}

// SetUmUnrealizedPNL gets a reference to the given string and assigns it to the UmUnrealizedPNL field.
func (o *AccountBalanceResponse1Inner) SetUmUnrealizedPNL(v string) {
	o.UmUnrealizedPNL = &v
}

// GetCmWalletBalance returns the CmWalletBalance field value if set, zero value otherwise.
func (o *AccountBalanceResponse1Inner) GetCmWalletBalance() string {
	if o == nil || common.IsNil(o.CmWalletBalance) {
		var ret string
		return ret
	}
	return *o.CmWalletBalance
}

// GetCmWalletBalanceOk returns a tuple with the CmWalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBalanceResponse1Inner) GetCmWalletBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.CmWalletBalance) {
		return nil, false
	}
	return o.CmWalletBalance, true
}

// HasCmWalletBalance returns a boolean if a field has been set.
func (o *AccountBalanceResponse1Inner) HasCmWalletBalance() bool {
	if o != nil && !common.IsNil(o.CmWalletBalance) {
		return true
	}

	return false
}

// SetCmWalletBalance gets a reference to the given string and assigns it to the CmWalletBalance field.
func (o *AccountBalanceResponse1Inner) SetCmWalletBalance(v string) {
	o.CmWalletBalance = &v
}

// GetCmUnrealizedPNL returns the CmUnrealizedPNL field value if set, zero value otherwise.
func (o *AccountBalanceResponse1Inner) GetCmUnrealizedPNL() string {
	if o == nil || common.IsNil(o.CmUnrealizedPNL) {
		var ret string
		return ret
	}
	return *o.CmUnrealizedPNL
}

// GetCmUnrealizedPNLOk returns a tuple with the CmUnrealizedPNL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBalanceResponse1Inner) GetCmUnrealizedPNLOk() (*string, bool) {
	if o == nil || common.IsNil(o.CmUnrealizedPNL) {
		return nil, false
	}
	return o.CmUnrealizedPNL, true
}

// HasCmUnrealizedPNL returns a boolean if a field has been set.
func (o *AccountBalanceResponse1Inner) HasCmUnrealizedPNL() bool {
	if o != nil && !common.IsNil(o.CmUnrealizedPNL) {
		return true
	}

	return false
}

// SetCmUnrealizedPNL gets a reference to the given string and assigns it to the CmUnrealizedPNL field.
func (o *AccountBalanceResponse1Inner) SetCmUnrealizedPNL(v string) {
	o.CmUnrealizedPNL = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *AccountBalanceResponse1Inner) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBalanceResponse1Inner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *AccountBalanceResponse1Inner) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *AccountBalanceResponse1Inner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

// GetNegativeBalance returns the NegativeBalance field value if set, zero value otherwise.
func (o *AccountBalanceResponse1Inner) GetNegativeBalance() string {
	if o == nil || common.IsNil(o.NegativeBalance) {
		var ret string
		return ret
	}
	return *o.NegativeBalance
}

// GetNegativeBalanceOk returns a tuple with the NegativeBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBalanceResponse1Inner) GetNegativeBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.NegativeBalance) {
		return nil, false
	}
	return o.NegativeBalance, true
}

// HasNegativeBalance returns a boolean if a field has been set.
func (o *AccountBalanceResponse1Inner) HasNegativeBalance() bool {
	if o != nil && !common.IsNil(o.NegativeBalance) {
		return true
	}

	return false
}

// SetNegativeBalance gets a reference to the given string and assigns it to the NegativeBalance field.
func (o *AccountBalanceResponse1Inner) SetNegativeBalance(v string) {
	o.NegativeBalance = &v
}

func (o AccountBalanceResponse1Inner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountBalanceResponse1Inner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.TotalWalletBalance) {
		toSerialize["totalWalletBalance"] = o.TotalWalletBalance
	}
	if !common.IsNil(o.CrossMarginAsset) {
		toSerialize["crossMarginAsset"] = o.CrossMarginAsset
	}
	if !common.IsNil(o.CrossMarginBorrowed) {
		toSerialize["crossMarginBorrowed"] = o.CrossMarginBorrowed
	}
	if !common.IsNil(o.CrossMarginFree) {
		toSerialize["crossMarginFree"] = o.CrossMarginFree
	}
	if !common.IsNil(o.CrossMarginInterest) {
		toSerialize["crossMarginInterest"] = o.CrossMarginInterest
	}
	if !common.IsNil(o.CrossMarginLocked) {
		toSerialize["crossMarginLocked"] = o.CrossMarginLocked
	}
	if !common.IsNil(o.UmWalletBalance) {
		toSerialize["umWalletBalance"] = o.UmWalletBalance
	}
	if !common.IsNil(o.UmUnrealizedPNL) {
		toSerialize["umUnrealizedPNL"] = o.UmUnrealizedPNL
	}
	if !common.IsNil(o.CmWalletBalance) {
		toSerialize["cmWalletBalance"] = o.CmWalletBalance
	}
	if !common.IsNil(o.CmUnrealizedPNL) {
		toSerialize["cmUnrealizedPNL"] = o.CmUnrealizedPNL
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if !common.IsNil(o.NegativeBalance) {
		toSerialize["negativeBalance"] = o.NegativeBalance
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountBalanceResponse1Inner) UnmarshalJSON(data []byte) (err error) {
	varAccountBalanceResponse1Inner := _AccountBalanceResponse1Inner{}

	err = json.Unmarshal(data, &varAccountBalanceResponse1Inner)

	if err != nil {
		return err
	}

	*o = AccountBalanceResponse1Inner(varAccountBalanceResponse1Inner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "totalWalletBalance")
		delete(additionalProperties, "crossMarginAsset")
		delete(additionalProperties, "crossMarginBorrowed")
		delete(additionalProperties, "crossMarginFree")
		delete(additionalProperties, "crossMarginInterest")
		delete(additionalProperties, "crossMarginLocked")
		delete(additionalProperties, "umWalletBalance")
		delete(additionalProperties, "umUnrealizedPNL")
		delete(additionalProperties, "cmWalletBalance")
		delete(additionalProperties, "cmUnrealizedPNL")
		delete(additionalProperties, "updateTime")
		delete(additionalProperties, "negativeBalance")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountBalanceResponse1Inner struct {
	value *AccountBalanceResponse1Inner
	isSet bool
}

func (v NullableAccountBalanceResponse1Inner) Get() *AccountBalanceResponse1Inner {
	return v.value
}

func (v *NullableAccountBalanceResponse1Inner) Set(val *AccountBalanceResponse1Inner) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountBalanceResponse1Inner) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountBalanceResponse1Inner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountBalanceResponse1Inner(val *AccountBalanceResponse1Inner) *NullableAccountBalanceResponse1Inner {
	return &NullableAccountBalanceResponse1Inner{value: val, isSet: true}
}

func (v NullableAccountBalanceResponse1Inner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountBalanceResponse1Inner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
