/*
Binance Derivatives Trading Portfolio Margin Pro REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin Pro REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetPortfolioMarginProAccountBalanceResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetPortfolioMarginProAccountBalanceResponseInner{}

// GetPortfolioMarginProAccountBalanceResponseInner struct for GetPortfolioMarginProAccountBalanceResponseInner
type GetPortfolioMarginProAccountBalanceResponseInner struct {
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
	OptionWalletBalance  *string `json:"optionWalletBalance,omitempty"`
	OptionEquity         *string `json:"optionEquity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetPortfolioMarginProAccountBalanceResponseInner GetPortfolioMarginProAccountBalanceResponseInner

// NewGetPortfolioMarginProAccountBalanceResponseInner instantiates a new GetPortfolioMarginProAccountBalanceResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPortfolioMarginProAccountBalanceResponseInner() *GetPortfolioMarginProAccountBalanceResponseInner {
	this := GetPortfolioMarginProAccountBalanceResponseInner{}
	return &this
}

// NewGetPortfolioMarginProAccountBalanceResponseInnerWithDefaults instantiates a new GetPortfolioMarginProAccountBalanceResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPortfolioMarginProAccountBalanceResponseInnerWithDefaults() *GetPortfolioMarginProAccountBalanceResponseInner {
	this := GetPortfolioMarginProAccountBalanceResponseInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) SetAsset(v string) {
	o.Asset = &v
}

// GetTotalWalletBalance returns the TotalWalletBalance field value if set, zero value otherwise.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) GetTotalWalletBalance() string {
	if o == nil || common.IsNil(o.TotalWalletBalance) {
		var ret string
		return ret
	}
	return *o.TotalWalletBalance
}

// GetTotalWalletBalanceOk returns a tuple with the TotalWalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) GetTotalWalletBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalWalletBalance) {
		return nil, false
	}
	return o.TotalWalletBalance, true
}

// HasTotalWalletBalance returns a boolean if a field has been set.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) HasTotalWalletBalance() bool {
	if o != nil && !common.IsNil(o.TotalWalletBalance) {
		return true
	}

	return false
}

// SetTotalWalletBalance gets a reference to the given string and assigns it to the TotalWalletBalance field.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) SetTotalWalletBalance(v string) {
	o.TotalWalletBalance = &v
}

// GetCrossMarginAsset returns the CrossMarginAsset field value if set, zero value otherwise.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) GetCrossMarginAsset() string {
	if o == nil || common.IsNil(o.CrossMarginAsset) {
		var ret string
		return ret
	}
	return *o.CrossMarginAsset
}

// GetCrossMarginAssetOk returns a tuple with the CrossMarginAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) GetCrossMarginAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.CrossMarginAsset) {
		return nil, false
	}
	return o.CrossMarginAsset, true
}

// HasCrossMarginAsset returns a boolean if a field has been set.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) HasCrossMarginAsset() bool {
	if o != nil && !common.IsNil(o.CrossMarginAsset) {
		return true
	}

	return false
}

// SetCrossMarginAsset gets a reference to the given string and assigns it to the CrossMarginAsset field.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) SetCrossMarginAsset(v string) {
	o.CrossMarginAsset = &v
}

// GetCrossMarginBorrowed returns the CrossMarginBorrowed field value if set, zero value otherwise.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) GetCrossMarginBorrowed() string {
	if o == nil || common.IsNil(o.CrossMarginBorrowed) {
		var ret string
		return ret
	}
	return *o.CrossMarginBorrowed
}

// GetCrossMarginBorrowedOk returns a tuple with the CrossMarginBorrowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) GetCrossMarginBorrowedOk() (*string, bool) {
	if o == nil || common.IsNil(o.CrossMarginBorrowed) {
		return nil, false
	}
	return o.CrossMarginBorrowed, true
}

// HasCrossMarginBorrowed returns a boolean if a field has been set.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) HasCrossMarginBorrowed() bool {
	if o != nil && !common.IsNil(o.CrossMarginBorrowed) {
		return true
	}

	return false
}

// SetCrossMarginBorrowed gets a reference to the given string and assigns it to the CrossMarginBorrowed field.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) SetCrossMarginBorrowed(v string) {
	o.CrossMarginBorrowed = &v
}

// GetCrossMarginFree returns the CrossMarginFree field value if set, zero value otherwise.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) GetCrossMarginFree() string {
	if o == nil || common.IsNil(o.CrossMarginFree) {
		var ret string
		return ret
	}
	return *o.CrossMarginFree
}

// GetCrossMarginFreeOk returns a tuple with the CrossMarginFree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) GetCrossMarginFreeOk() (*string, bool) {
	if o == nil || common.IsNil(o.CrossMarginFree) {
		return nil, false
	}
	return o.CrossMarginFree, true
}

// HasCrossMarginFree returns a boolean if a field has been set.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) HasCrossMarginFree() bool {
	if o != nil && !common.IsNil(o.CrossMarginFree) {
		return true
	}

	return false
}

// SetCrossMarginFree gets a reference to the given string and assigns it to the CrossMarginFree field.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) SetCrossMarginFree(v string) {
	o.CrossMarginFree = &v
}

// GetCrossMarginInterest returns the CrossMarginInterest field value if set, zero value otherwise.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) GetCrossMarginInterest() string {
	if o == nil || common.IsNil(o.CrossMarginInterest) {
		var ret string
		return ret
	}
	return *o.CrossMarginInterest
}

// GetCrossMarginInterestOk returns a tuple with the CrossMarginInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) GetCrossMarginInterestOk() (*string, bool) {
	if o == nil || common.IsNil(o.CrossMarginInterest) {
		return nil, false
	}
	return o.CrossMarginInterest, true
}

// HasCrossMarginInterest returns a boolean if a field has been set.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) HasCrossMarginInterest() bool {
	if o != nil && !common.IsNil(o.CrossMarginInterest) {
		return true
	}

	return false
}

// SetCrossMarginInterest gets a reference to the given string and assigns it to the CrossMarginInterest field.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) SetCrossMarginInterest(v string) {
	o.CrossMarginInterest = &v
}

// GetCrossMarginLocked returns the CrossMarginLocked field value if set, zero value otherwise.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) GetCrossMarginLocked() string {
	if o == nil || common.IsNil(o.CrossMarginLocked) {
		var ret string
		return ret
	}
	return *o.CrossMarginLocked
}

// GetCrossMarginLockedOk returns a tuple with the CrossMarginLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) GetCrossMarginLockedOk() (*string, bool) {
	if o == nil || common.IsNil(o.CrossMarginLocked) {
		return nil, false
	}
	return o.CrossMarginLocked, true
}

// HasCrossMarginLocked returns a boolean if a field has been set.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) HasCrossMarginLocked() bool {
	if o != nil && !common.IsNil(o.CrossMarginLocked) {
		return true
	}

	return false
}

// SetCrossMarginLocked gets a reference to the given string and assigns it to the CrossMarginLocked field.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) SetCrossMarginLocked(v string) {
	o.CrossMarginLocked = &v
}

// GetUmWalletBalance returns the UmWalletBalance field value if set, zero value otherwise.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) GetUmWalletBalance() string {
	if o == nil || common.IsNil(o.UmWalletBalance) {
		var ret string
		return ret
	}
	return *o.UmWalletBalance
}

// GetUmWalletBalanceOk returns a tuple with the UmWalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) GetUmWalletBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.UmWalletBalance) {
		return nil, false
	}
	return o.UmWalletBalance, true
}

// HasUmWalletBalance returns a boolean if a field has been set.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) HasUmWalletBalance() bool {
	if o != nil && !common.IsNil(o.UmWalletBalance) {
		return true
	}

	return false
}

// SetUmWalletBalance gets a reference to the given string and assigns it to the UmWalletBalance field.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) SetUmWalletBalance(v string) {
	o.UmWalletBalance = &v
}

// GetUmUnrealizedPNL returns the UmUnrealizedPNL field value if set, zero value otherwise.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) GetUmUnrealizedPNL() string {
	if o == nil || common.IsNil(o.UmUnrealizedPNL) {
		var ret string
		return ret
	}
	return *o.UmUnrealizedPNL
}

// GetUmUnrealizedPNLOk returns a tuple with the UmUnrealizedPNL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) GetUmUnrealizedPNLOk() (*string, bool) {
	if o == nil || common.IsNil(o.UmUnrealizedPNL) {
		return nil, false
	}
	return o.UmUnrealizedPNL, true
}

// HasUmUnrealizedPNL returns a boolean if a field has been set.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) HasUmUnrealizedPNL() bool {
	if o != nil && !common.IsNil(o.UmUnrealizedPNL) {
		return true
	}

	return false
}

// SetUmUnrealizedPNL gets a reference to the given string and assigns it to the UmUnrealizedPNL field.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) SetUmUnrealizedPNL(v string) {
	o.UmUnrealizedPNL = &v
}

// GetCmWalletBalance returns the CmWalletBalance field value if set, zero value otherwise.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) GetCmWalletBalance() string {
	if o == nil || common.IsNil(o.CmWalletBalance) {
		var ret string
		return ret
	}
	return *o.CmWalletBalance
}

// GetCmWalletBalanceOk returns a tuple with the CmWalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) GetCmWalletBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.CmWalletBalance) {
		return nil, false
	}
	return o.CmWalletBalance, true
}

// HasCmWalletBalance returns a boolean if a field has been set.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) HasCmWalletBalance() bool {
	if o != nil && !common.IsNil(o.CmWalletBalance) {
		return true
	}

	return false
}

// SetCmWalletBalance gets a reference to the given string and assigns it to the CmWalletBalance field.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) SetCmWalletBalance(v string) {
	o.CmWalletBalance = &v
}

// GetCmUnrealizedPNL returns the CmUnrealizedPNL field value if set, zero value otherwise.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) GetCmUnrealizedPNL() string {
	if o == nil || common.IsNil(o.CmUnrealizedPNL) {
		var ret string
		return ret
	}
	return *o.CmUnrealizedPNL
}

// GetCmUnrealizedPNLOk returns a tuple with the CmUnrealizedPNL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) GetCmUnrealizedPNLOk() (*string, bool) {
	if o == nil || common.IsNil(o.CmUnrealizedPNL) {
		return nil, false
	}
	return o.CmUnrealizedPNL, true
}

// HasCmUnrealizedPNL returns a boolean if a field has been set.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) HasCmUnrealizedPNL() bool {
	if o != nil && !common.IsNil(o.CmUnrealizedPNL) {
		return true
	}

	return false
}

// SetCmUnrealizedPNL gets a reference to the given string and assigns it to the CmUnrealizedPNL field.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) SetCmUnrealizedPNL(v string) {
	o.CmUnrealizedPNL = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

// GetNegativeBalance returns the NegativeBalance field value if set, zero value otherwise.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) GetNegativeBalance() string {
	if o == nil || common.IsNil(o.NegativeBalance) {
		var ret string
		return ret
	}
	return *o.NegativeBalance
}

// GetNegativeBalanceOk returns a tuple with the NegativeBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) GetNegativeBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.NegativeBalance) {
		return nil, false
	}
	return o.NegativeBalance, true
}

// HasNegativeBalance returns a boolean if a field has been set.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) HasNegativeBalance() bool {
	if o != nil && !common.IsNil(o.NegativeBalance) {
		return true
	}

	return false
}

// SetNegativeBalance gets a reference to the given string and assigns it to the NegativeBalance field.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) SetNegativeBalance(v string) {
	o.NegativeBalance = &v
}

// GetOptionWalletBalance returns the OptionWalletBalance field value if set, zero value otherwise.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) GetOptionWalletBalance() string {
	if o == nil || common.IsNil(o.OptionWalletBalance) {
		var ret string
		return ret
	}
	return *o.OptionWalletBalance
}

// GetOptionWalletBalanceOk returns a tuple with the OptionWalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) GetOptionWalletBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.OptionWalletBalance) {
		return nil, false
	}
	return o.OptionWalletBalance, true
}

// HasOptionWalletBalance returns a boolean if a field has been set.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) HasOptionWalletBalance() bool {
	if o != nil && !common.IsNil(o.OptionWalletBalance) {
		return true
	}

	return false
}

// SetOptionWalletBalance gets a reference to the given string and assigns it to the OptionWalletBalance field.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) SetOptionWalletBalance(v string) {
	o.OptionWalletBalance = &v
}

// GetOptionEquity returns the OptionEquity field value if set, zero value otherwise.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) GetOptionEquity() string {
	if o == nil || common.IsNil(o.OptionEquity) {
		var ret string
		return ret
	}
	return *o.OptionEquity
}

// GetOptionEquityOk returns a tuple with the OptionEquity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) GetOptionEquityOk() (*string, bool) {
	if o == nil || common.IsNil(o.OptionEquity) {
		return nil, false
	}
	return o.OptionEquity, true
}

// HasOptionEquity returns a boolean if a field has been set.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) HasOptionEquity() bool {
	if o != nil && !common.IsNil(o.OptionEquity) {
		return true
	}

	return false
}

// SetOptionEquity gets a reference to the given string and assigns it to the OptionEquity field.
func (o *GetPortfolioMarginProAccountBalanceResponseInner) SetOptionEquity(v string) {
	o.OptionEquity = &v
}

func (o GetPortfolioMarginProAccountBalanceResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPortfolioMarginProAccountBalanceResponseInner) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.OptionWalletBalance) {
		toSerialize["optionWalletBalance"] = o.OptionWalletBalance
	}
	if !common.IsNil(o.OptionEquity) {
		toSerialize["optionEquity"] = o.OptionEquity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetPortfolioMarginProAccountBalanceResponseInner) UnmarshalJSON(data []byte) (err error) {
	varGetPortfolioMarginProAccountBalanceResponseInner := _GetPortfolioMarginProAccountBalanceResponseInner{}

	err = json.Unmarshal(data, &varGetPortfolioMarginProAccountBalanceResponseInner)

	if err != nil {
		return err
	}

	*o = GetPortfolioMarginProAccountBalanceResponseInner(varGetPortfolioMarginProAccountBalanceResponseInner)

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
		delete(additionalProperties, "optionWalletBalance")
		delete(additionalProperties, "optionEquity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetPortfolioMarginProAccountBalanceResponseInner struct {
	value *GetPortfolioMarginProAccountBalanceResponseInner
	isSet bool
}

func (v NullableGetPortfolioMarginProAccountBalanceResponseInner) Get() *GetPortfolioMarginProAccountBalanceResponseInner {
	return v.value
}

func (v *NullableGetPortfolioMarginProAccountBalanceResponseInner) Set(val *GetPortfolioMarginProAccountBalanceResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPortfolioMarginProAccountBalanceResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPortfolioMarginProAccountBalanceResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPortfolioMarginProAccountBalanceResponseInner(val *GetPortfolioMarginProAccountBalanceResponseInner) *NullableGetPortfolioMarginProAccountBalanceResponseInner {
	return &NullableGetPortfolioMarginProAccountBalanceResponseInner{value: val, isSet: true}
}

func (v NullableGetPortfolioMarginProAccountBalanceResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPortfolioMarginProAccountBalanceResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
