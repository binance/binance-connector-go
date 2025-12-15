/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the FuturesAccountBalanceV2ResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FuturesAccountBalanceV2ResponseInner{}

// FuturesAccountBalanceV2ResponseInner struct for FuturesAccountBalanceV2ResponseInner
type FuturesAccountBalanceV2ResponseInner struct {
	AccountAlias         *string `json:"accountAlias,omitempty"`
	Asset                *string `json:"asset,omitempty"`
	Balance              *string `json:"balance,omitempty"`
	CrossWalletBalance   *string `json:"crossWalletBalance,omitempty"`
	CrossUnPnl           *string `json:"crossUnPnl,omitempty"`
	AvailableBalance     *string `json:"availableBalance,omitempty"`
	MaxWithdrawAmount    *string `json:"maxWithdrawAmount,omitempty"`
	MarginAvailable      *bool   `json:"marginAvailable,omitempty"`
	UpdateTime           *int64  `json:"updateTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FuturesAccountBalanceV2ResponseInner FuturesAccountBalanceV2ResponseInner

// NewFuturesAccountBalanceV2ResponseInner instantiates a new FuturesAccountBalanceV2ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFuturesAccountBalanceV2ResponseInner() *FuturesAccountBalanceV2ResponseInner {
	this := FuturesAccountBalanceV2ResponseInner{}
	return &this
}

// NewFuturesAccountBalanceV2ResponseInnerWithDefaults instantiates a new FuturesAccountBalanceV2ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFuturesAccountBalanceV2ResponseInnerWithDefaults() *FuturesAccountBalanceV2ResponseInner {
	this := FuturesAccountBalanceV2ResponseInner{}
	return &this
}

// GetAccountAlias returns the AccountAlias field value if set, zero value otherwise.
func (o *FuturesAccountBalanceV2ResponseInner) GetAccountAlias() string {
	if o == nil || common.IsNil(o.AccountAlias) {
		var ret string
		return ret
	}
	return *o.AccountAlias
}

// GetAccountAliasOk returns a tuple with the AccountAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesAccountBalanceV2ResponseInner) GetAccountAliasOk() (*string, bool) {
	if o == nil || common.IsNil(o.AccountAlias) {
		return nil, false
	}
	return o.AccountAlias, true
}

// HasAccountAlias returns a boolean if a field has been set.
func (o *FuturesAccountBalanceV2ResponseInner) HasAccountAlias() bool {
	if o != nil && !common.IsNil(o.AccountAlias) {
		return true
	}

	return false
}

// SetAccountAlias gets a reference to the given string and assigns it to the AccountAlias field.
func (o *FuturesAccountBalanceV2ResponseInner) SetAccountAlias(v string) {
	o.AccountAlias = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *FuturesAccountBalanceV2ResponseInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesAccountBalanceV2ResponseInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *FuturesAccountBalanceV2ResponseInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *FuturesAccountBalanceV2ResponseInner) SetAsset(v string) {
	o.Asset = &v
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *FuturesAccountBalanceV2ResponseInner) GetBalance() string {
	if o == nil || common.IsNil(o.Balance) {
		var ret string
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesAccountBalanceV2ResponseInner) GetBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *FuturesAccountBalanceV2ResponseInner) HasBalance() bool {
	if o != nil && !common.IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given string and assigns it to the Balance field.
func (o *FuturesAccountBalanceV2ResponseInner) SetBalance(v string) {
	o.Balance = &v
}

// GetCrossWalletBalance returns the CrossWalletBalance field value if set, zero value otherwise.
func (o *FuturesAccountBalanceV2ResponseInner) GetCrossWalletBalance() string {
	if o == nil || common.IsNil(o.CrossWalletBalance) {
		var ret string
		return ret
	}
	return *o.CrossWalletBalance
}

// GetCrossWalletBalanceOk returns a tuple with the CrossWalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesAccountBalanceV2ResponseInner) GetCrossWalletBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.CrossWalletBalance) {
		return nil, false
	}
	return o.CrossWalletBalance, true
}

// HasCrossWalletBalance returns a boolean if a field has been set.
func (o *FuturesAccountBalanceV2ResponseInner) HasCrossWalletBalance() bool {
	if o != nil && !common.IsNil(o.CrossWalletBalance) {
		return true
	}

	return false
}

// SetCrossWalletBalance gets a reference to the given string and assigns it to the CrossWalletBalance field.
func (o *FuturesAccountBalanceV2ResponseInner) SetCrossWalletBalance(v string) {
	o.CrossWalletBalance = &v
}

// GetCrossUnPnl returns the CrossUnPnl field value if set, zero value otherwise.
func (o *FuturesAccountBalanceV2ResponseInner) GetCrossUnPnl() string {
	if o == nil || common.IsNil(o.CrossUnPnl) {
		var ret string
		return ret
	}
	return *o.CrossUnPnl
}

// GetCrossUnPnlOk returns a tuple with the CrossUnPnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesAccountBalanceV2ResponseInner) GetCrossUnPnlOk() (*string, bool) {
	if o == nil || common.IsNil(o.CrossUnPnl) {
		return nil, false
	}
	return o.CrossUnPnl, true
}

// HasCrossUnPnl returns a boolean if a field has been set.
func (o *FuturesAccountBalanceV2ResponseInner) HasCrossUnPnl() bool {
	if o != nil && !common.IsNil(o.CrossUnPnl) {
		return true
	}

	return false
}

// SetCrossUnPnl gets a reference to the given string and assigns it to the CrossUnPnl field.
func (o *FuturesAccountBalanceV2ResponseInner) SetCrossUnPnl(v string) {
	o.CrossUnPnl = &v
}

// GetAvailableBalance returns the AvailableBalance field value if set, zero value otherwise.
func (o *FuturesAccountBalanceV2ResponseInner) GetAvailableBalance() string {
	if o == nil || common.IsNil(o.AvailableBalance) {
		var ret string
		return ret
	}
	return *o.AvailableBalance
}

// GetAvailableBalanceOk returns a tuple with the AvailableBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesAccountBalanceV2ResponseInner) GetAvailableBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.AvailableBalance) {
		return nil, false
	}
	return o.AvailableBalance, true
}

// HasAvailableBalance returns a boolean if a field has been set.
func (o *FuturesAccountBalanceV2ResponseInner) HasAvailableBalance() bool {
	if o != nil && !common.IsNil(o.AvailableBalance) {
		return true
	}

	return false
}

// SetAvailableBalance gets a reference to the given string and assigns it to the AvailableBalance field.
func (o *FuturesAccountBalanceV2ResponseInner) SetAvailableBalance(v string) {
	o.AvailableBalance = &v
}

// GetMaxWithdrawAmount returns the MaxWithdrawAmount field value if set, zero value otherwise.
func (o *FuturesAccountBalanceV2ResponseInner) GetMaxWithdrawAmount() string {
	if o == nil || common.IsNil(o.MaxWithdrawAmount) {
		var ret string
		return ret
	}
	return *o.MaxWithdrawAmount
}

// GetMaxWithdrawAmountOk returns a tuple with the MaxWithdrawAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesAccountBalanceV2ResponseInner) GetMaxWithdrawAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxWithdrawAmount) {
		return nil, false
	}
	return o.MaxWithdrawAmount, true
}

// HasMaxWithdrawAmount returns a boolean if a field has been set.
func (o *FuturesAccountBalanceV2ResponseInner) HasMaxWithdrawAmount() bool {
	if o != nil && !common.IsNil(o.MaxWithdrawAmount) {
		return true
	}

	return false
}

// SetMaxWithdrawAmount gets a reference to the given string and assigns it to the MaxWithdrawAmount field.
func (o *FuturesAccountBalanceV2ResponseInner) SetMaxWithdrawAmount(v string) {
	o.MaxWithdrawAmount = &v
}

// GetMarginAvailable returns the MarginAvailable field value if set, zero value otherwise.
func (o *FuturesAccountBalanceV2ResponseInner) GetMarginAvailable() bool {
	if o == nil || common.IsNil(o.MarginAvailable) {
		var ret bool
		return ret
	}
	return *o.MarginAvailable
}

// GetMarginAvailableOk returns a tuple with the MarginAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesAccountBalanceV2ResponseInner) GetMarginAvailableOk() (*bool, bool) {
	if o == nil || common.IsNil(o.MarginAvailable) {
		return nil, false
	}
	return o.MarginAvailable, true
}

// HasMarginAvailable returns a boolean if a field has been set.
func (o *FuturesAccountBalanceV2ResponseInner) HasMarginAvailable() bool {
	if o != nil && !common.IsNil(o.MarginAvailable) {
		return true
	}

	return false
}

// SetMarginAvailable gets a reference to the given bool and assigns it to the MarginAvailable field.
func (o *FuturesAccountBalanceV2ResponseInner) SetMarginAvailable(v bool) {
	o.MarginAvailable = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *FuturesAccountBalanceV2ResponseInner) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesAccountBalanceV2ResponseInner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *FuturesAccountBalanceV2ResponseInner) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *FuturesAccountBalanceV2ResponseInner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o FuturesAccountBalanceV2ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FuturesAccountBalanceV2ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AccountAlias) {
		toSerialize["accountAlias"] = o.AccountAlias
	}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Balance) {
		toSerialize["balance"] = o.Balance
	}
	if !common.IsNil(o.CrossWalletBalance) {
		toSerialize["crossWalletBalance"] = o.CrossWalletBalance
	}
	if !common.IsNil(o.CrossUnPnl) {
		toSerialize["crossUnPnl"] = o.CrossUnPnl
	}
	if !common.IsNil(o.AvailableBalance) {
		toSerialize["availableBalance"] = o.AvailableBalance
	}
	if !common.IsNil(o.MaxWithdrawAmount) {
		toSerialize["maxWithdrawAmount"] = o.MaxWithdrawAmount
	}
	if !common.IsNil(o.MarginAvailable) {
		toSerialize["marginAvailable"] = o.MarginAvailable
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FuturesAccountBalanceV2ResponseInner) UnmarshalJSON(data []byte) (err error) {
	varFuturesAccountBalanceV2ResponseInner := _FuturesAccountBalanceV2ResponseInner{}

	err = json.Unmarshal(data, &varFuturesAccountBalanceV2ResponseInner)

	if err != nil {
		return err
	}

	*o = FuturesAccountBalanceV2ResponseInner(varFuturesAccountBalanceV2ResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accountAlias")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "balance")
		delete(additionalProperties, "crossWalletBalance")
		delete(additionalProperties, "crossUnPnl")
		delete(additionalProperties, "availableBalance")
		delete(additionalProperties, "maxWithdrawAmount")
		delete(additionalProperties, "marginAvailable")
		delete(additionalProperties, "updateTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFuturesAccountBalanceV2ResponseInner struct {
	value *FuturesAccountBalanceV2ResponseInner
	isSet bool
}

func (v NullableFuturesAccountBalanceV2ResponseInner) Get() *FuturesAccountBalanceV2ResponseInner {
	return v.value
}

func (v *NullableFuturesAccountBalanceV2ResponseInner) Set(val *FuturesAccountBalanceV2ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableFuturesAccountBalanceV2ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableFuturesAccountBalanceV2ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFuturesAccountBalanceV2ResponseInner(val *FuturesAccountBalanceV2ResponseInner) *NullableFuturesAccountBalanceV2ResponseInner {
	return &NullableFuturesAccountBalanceV2ResponseInner{value: val, isSet: true}
}

func (v NullableFuturesAccountBalanceV2ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFuturesAccountBalanceV2ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
