/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AccountBalanceResponse2 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountBalanceResponse2{}

// AccountBalanceResponse2 struct for AccountBalanceResponse2
type AccountBalanceResponse2 struct {
	Asset                *string `json:"asset,omitempty"`
	TotalWalletBalance   *string `json:"totalWalletBalance,omitempty"`
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

type _AccountBalanceResponse2 AccountBalanceResponse2

// NewAccountBalanceResponse2 instantiates a new AccountBalanceResponse2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountBalanceResponse2() *AccountBalanceResponse2 {
	this := AccountBalanceResponse2{}
	return &this
}

// NewAccountBalanceResponse2WithDefaults instantiates a new AccountBalanceResponse2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountBalanceResponse2WithDefaults() *AccountBalanceResponse2 {
	this := AccountBalanceResponse2{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *AccountBalanceResponse2) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBalanceResponse2) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *AccountBalanceResponse2) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *AccountBalanceResponse2) SetAsset(v string) {
	o.Asset = &v
}

// GetTotalWalletBalance returns the TotalWalletBalance field value if set, zero value otherwise.
func (o *AccountBalanceResponse2) GetTotalWalletBalance() string {
	if o == nil || common.IsNil(o.TotalWalletBalance) {
		var ret string
		return ret
	}
	return *o.TotalWalletBalance
}

// GetTotalWalletBalanceOk returns a tuple with the TotalWalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBalanceResponse2) GetTotalWalletBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalWalletBalance) {
		return nil, false
	}
	return o.TotalWalletBalance, true
}

// HasTotalWalletBalance returns a boolean if a field has been set.
func (o *AccountBalanceResponse2) HasTotalWalletBalance() bool {
	if o != nil && !common.IsNil(o.TotalWalletBalance) {
		return true
	}

	return false
}

// SetTotalWalletBalance gets a reference to the given string and assigns it to the TotalWalletBalance field.
func (o *AccountBalanceResponse2) SetTotalWalletBalance(v string) {
	o.TotalWalletBalance = &v
}

// GetCrossMarginBorrowed returns the CrossMarginBorrowed field value if set, zero value otherwise.
func (o *AccountBalanceResponse2) GetCrossMarginBorrowed() string {
	if o == nil || common.IsNil(o.CrossMarginBorrowed) {
		var ret string
		return ret
	}
	return *o.CrossMarginBorrowed
}

// GetCrossMarginBorrowedOk returns a tuple with the CrossMarginBorrowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBalanceResponse2) GetCrossMarginBorrowedOk() (*string, bool) {
	if o == nil || common.IsNil(o.CrossMarginBorrowed) {
		return nil, false
	}
	return o.CrossMarginBorrowed, true
}

// HasCrossMarginBorrowed returns a boolean if a field has been set.
func (o *AccountBalanceResponse2) HasCrossMarginBorrowed() bool {
	if o != nil && !common.IsNil(o.CrossMarginBorrowed) {
		return true
	}

	return false
}

// SetCrossMarginBorrowed gets a reference to the given string and assigns it to the CrossMarginBorrowed field.
func (o *AccountBalanceResponse2) SetCrossMarginBorrowed(v string) {
	o.CrossMarginBorrowed = &v
}

// GetCrossMarginFree returns the CrossMarginFree field value if set, zero value otherwise.
func (o *AccountBalanceResponse2) GetCrossMarginFree() string {
	if o == nil || common.IsNil(o.CrossMarginFree) {
		var ret string
		return ret
	}
	return *o.CrossMarginFree
}

// GetCrossMarginFreeOk returns a tuple with the CrossMarginFree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBalanceResponse2) GetCrossMarginFreeOk() (*string, bool) {
	if o == nil || common.IsNil(o.CrossMarginFree) {
		return nil, false
	}
	return o.CrossMarginFree, true
}

// HasCrossMarginFree returns a boolean if a field has been set.
func (o *AccountBalanceResponse2) HasCrossMarginFree() bool {
	if o != nil && !common.IsNil(o.CrossMarginFree) {
		return true
	}

	return false
}

// SetCrossMarginFree gets a reference to the given string and assigns it to the CrossMarginFree field.
func (o *AccountBalanceResponse2) SetCrossMarginFree(v string) {
	o.CrossMarginFree = &v
}

// GetCrossMarginInterest returns the CrossMarginInterest field value if set, zero value otherwise.
func (o *AccountBalanceResponse2) GetCrossMarginInterest() string {
	if o == nil || common.IsNil(o.CrossMarginInterest) {
		var ret string
		return ret
	}
	return *o.CrossMarginInterest
}

// GetCrossMarginInterestOk returns a tuple with the CrossMarginInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBalanceResponse2) GetCrossMarginInterestOk() (*string, bool) {
	if o == nil || common.IsNil(o.CrossMarginInterest) {
		return nil, false
	}
	return o.CrossMarginInterest, true
}

// HasCrossMarginInterest returns a boolean if a field has been set.
func (o *AccountBalanceResponse2) HasCrossMarginInterest() bool {
	if o != nil && !common.IsNil(o.CrossMarginInterest) {
		return true
	}

	return false
}

// SetCrossMarginInterest gets a reference to the given string and assigns it to the CrossMarginInterest field.
func (o *AccountBalanceResponse2) SetCrossMarginInterest(v string) {
	o.CrossMarginInterest = &v
}

// GetCrossMarginLocked returns the CrossMarginLocked field value if set, zero value otherwise.
func (o *AccountBalanceResponse2) GetCrossMarginLocked() string {
	if o == nil || common.IsNil(o.CrossMarginLocked) {
		var ret string
		return ret
	}
	return *o.CrossMarginLocked
}

// GetCrossMarginLockedOk returns a tuple with the CrossMarginLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBalanceResponse2) GetCrossMarginLockedOk() (*string, bool) {
	if o == nil || common.IsNil(o.CrossMarginLocked) {
		return nil, false
	}
	return o.CrossMarginLocked, true
}

// HasCrossMarginLocked returns a boolean if a field has been set.
func (o *AccountBalanceResponse2) HasCrossMarginLocked() bool {
	if o != nil && !common.IsNil(o.CrossMarginLocked) {
		return true
	}

	return false
}

// SetCrossMarginLocked gets a reference to the given string and assigns it to the CrossMarginLocked field.
func (o *AccountBalanceResponse2) SetCrossMarginLocked(v string) {
	o.CrossMarginLocked = &v
}

// GetUmWalletBalance returns the UmWalletBalance field value if set, zero value otherwise.
func (o *AccountBalanceResponse2) GetUmWalletBalance() string {
	if o == nil || common.IsNil(o.UmWalletBalance) {
		var ret string
		return ret
	}
	return *o.UmWalletBalance
}

// GetUmWalletBalanceOk returns a tuple with the UmWalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBalanceResponse2) GetUmWalletBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.UmWalletBalance) {
		return nil, false
	}
	return o.UmWalletBalance, true
}

// HasUmWalletBalance returns a boolean if a field has been set.
func (o *AccountBalanceResponse2) HasUmWalletBalance() bool {
	if o != nil && !common.IsNil(o.UmWalletBalance) {
		return true
	}

	return false
}

// SetUmWalletBalance gets a reference to the given string and assigns it to the UmWalletBalance field.
func (o *AccountBalanceResponse2) SetUmWalletBalance(v string) {
	o.UmWalletBalance = &v
}

// GetUmUnrealizedPNL returns the UmUnrealizedPNL field value if set, zero value otherwise.
func (o *AccountBalanceResponse2) GetUmUnrealizedPNL() string {
	if o == nil || common.IsNil(o.UmUnrealizedPNL) {
		var ret string
		return ret
	}
	return *o.UmUnrealizedPNL
}

// GetUmUnrealizedPNLOk returns a tuple with the UmUnrealizedPNL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBalanceResponse2) GetUmUnrealizedPNLOk() (*string, bool) {
	if o == nil || common.IsNil(o.UmUnrealizedPNL) {
		return nil, false
	}
	return o.UmUnrealizedPNL, true
}

// HasUmUnrealizedPNL returns a boolean if a field has been set.
func (o *AccountBalanceResponse2) HasUmUnrealizedPNL() bool {
	if o != nil && !common.IsNil(o.UmUnrealizedPNL) {
		return true
	}

	return false
}

// SetUmUnrealizedPNL gets a reference to the given string and assigns it to the UmUnrealizedPNL field.
func (o *AccountBalanceResponse2) SetUmUnrealizedPNL(v string) {
	o.UmUnrealizedPNL = &v
}

// GetCmWalletBalance returns the CmWalletBalance field value if set, zero value otherwise.
func (o *AccountBalanceResponse2) GetCmWalletBalance() string {
	if o == nil || common.IsNil(o.CmWalletBalance) {
		var ret string
		return ret
	}
	return *o.CmWalletBalance
}

// GetCmWalletBalanceOk returns a tuple with the CmWalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBalanceResponse2) GetCmWalletBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.CmWalletBalance) {
		return nil, false
	}
	return o.CmWalletBalance, true
}

// HasCmWalletBalance returns a boolean if a field has been set.
func (o *AccountBalanceResponse2) HasCmWalletBalance() bool {
	if o != nil && !common.IsNil(o.CmWalletBalance) {
		return true
	}

	return false
}

// SetCmWalletBalance gets a reference to the given string and assigns it to the CmWalletBalance field.
func (o *AccountBalanceResponse2) SetCmWalletBalance(v string) {
	o.CmWalletBalance = &v
}

// GetCmUnrealizedPNL returns the CmUnrealizedPNL field value if set, zero value otherwise.
func (o *AccountBalanceResponse2) GetCmUnrealizedPNL() string {
	if o == nil || common.IsNil(o.CmUnrealizedPNL) {
		var ret string
		return ret
	}
	return *o.CmUnrealizedPNL
}

// GetCmUnrealizedPNLOk returns a tuple with the CmUnrealizedPNL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBalanceResponse2) GetCmUnrealizedPNLOk() (*string, bool) {
	if o == nil || common.IsNil(o.CmUnrealizedPNL) {
		return nil, false
	}
	return o.CmUnrealizedPNL, true
}

// HasCmUnrealizedPNL returns a boolean if a field has been set.
func (o *AccountBalanceResponse2) HasCmUnrealizedPNL() bool {
	if o != nil && !common.IsNil(o.CmUnrealizedPNL) {
		return true
	}

	return false
}

// SetCmUnrealizedPNL gets a reference to the given string and assigns it to the CmUnrealizedPNL field.
func (o *AccountBalanceResponse2) SetCmUnrealizedPNL(v string) {
	o.CmUnrealizedPNL = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *AccountBalanceResponse2) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBalanceResponse2) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *AccountBalanceResponse2) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *AccountBalanceResponse2) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

// GetNegativeBalance returns the NegativeBalance field value if set, zero value otherwise.
func (o *AccountBalanceResponse2) GetNegativeBalance() string {
	if o == nil || common.IsNil(o.NegativeBalance) {
		var ret string
		return ret
	}
	return *o.NegativeBalance
}

// GetNegativeBalanceOk returns a tuple with the NegativeBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBalanceResponse2) GetNegativeBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.NegativeBalance) {
		return nil, false
	}
	return o.NegativeBalance, true
}

// HasNegativeBalance returns a boolean if a field has been set.
func (o *AccountBalanceResponse2) HasNegativeBalance() bool {
	if o != nil && !common.IsNil(o.NegativeBalance) {
		return true
	}

	return false
}

// SetNegativeBalance gets a reference to the given string and assigns it to the NegativeBalance field.
func (o *AccountBalanceResponse2) SetNegativeBalance(v string) {
	o.NegativeBalance = &v
}

func (o AccountBalanceResponse2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountBalanceResponse2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.TotalWalletBalance) {
		toSerialize["totalWalletBalance"] = o.TotalWalletBalance
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

func (o *AccountBalanceResponse2) UnmarshalJSON(data []byte) (err error) {
	varAccountBalanceResponse2 := _AccountBalanceResponse2{}

	err = json.Unmarshal(data, &varAccountBalanceResponse2)

	if err != nil {
		return err
	}

	*o = AccountBalanceResponse2(varAccountBalanceResponse2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "totalWalletBalance")
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

type NullableAccountBalanceResponse2 struct {
	value *AccountBalanceResponse2
	isSet bool
}

func (v NullableAccountBalanceResponse2) Get() *AccountBalanceResponse2 {
	return v.value
}

func (v *NullableAccountBalanceResponse2) Set(val *AccountBalanceResponse2) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountBalanceResponse2) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountBalanceResponse2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountBalanceResponse2(val *AccountBalanceResponse2) *NullableAccountBalanceResponse2 {
	return &NullableAccountBalanceResponse2{value: val, isSet: true}
}

func (v NullableAccountBalanceResponse2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountBalanceResponse2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
