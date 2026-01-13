/*
Binance Derivatives Trading USDS Futures WebSocket API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the FuturesAccountBalanceV2ResponseResultInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FuturesAccountBalanceV2ResponseResultInner{}

// FuturesAccountBalanceV2ResponseResultInner struct for FuturesAccountBalanceV2ResponseResultInner
type FuturesAccountBalanceV2ResponseResultInner struct {
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

type _FuturesAccountBalanceV2ResponseResultInner FuturesAccountBalanceV2ResponseResultInner

// NewFuturesAccountBalanceV2ResponseResultInner instantiates a new FuturesAccountBalanceV2ResponseResultInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFuturesAccountBalanceV2ResponseResultInner() *FuturesAccountBalanceV2ResponseResultInner {
	this := FuturesAccountBalanceV2ResponseResultInner{}
	return &this
}

// NewFuturesAccountBalanceV2ResponseResultInnerWithDefaults instantiates a new FuturesAccountBalanceV2ResponseResultInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFuturesAccountBalanceV2ResponseResultInnerWithDefaults() *FuturesAccountBalanceV2ResponseResultInner {
	this := FuturesAccountBalanceV2ResponseResultInner{}
	return &this
}

// GetAccountAlias returns the AccountAlias field value if set, zero value otherwise.
func (o *FuturesAccountBalanceV2ResponseResultInner) GetAccountAlias() string {
	if o == nil || common.IsNil(o.AccountAlias) {
		var ret string
		return ret
	}
	return *o.AccountAlias
}

// GetAccountAliasOk returns a tuple with the AccountAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesAccountBalanceV2ResponseResultInner) GetAccountAliasOk() (*string, bool) {
	if o == nil || common.IsNil(o.AccountAlias) {
		return nil, false
	}
	return o.AccountAlias, true
}

// HasAccountAlias returns a boolean if a field has been set.
func (o *FuturesAccountBalanceV2ResponseResultInner) HasAccountAlias() bool {
	if o != nil && !common.IsNil(o.AccountAlias) {
		return true
	}

	return false
}

// SetAccountAlias gets a reference to the given string and assigns it to the AccountAlias field.
func (o *FuturesAccountBalanceV2ResponseResultInner) SetAccountAlias(v string) {
	o.AccountAlias = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *FuturesAccountBalanceV2ResponseResultInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesAccountBalanceV2ResponseResultInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *FuturesAccountBalanceV2ResponseResultInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *FuturesAccountBalanceV2ResponseResultInner) SetAsset(v string) {
	o.Asset = &v
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *FuturesAccountBalanceV2ResponseResultInner) GetBalance() string {
	if o == nil || common.IsNil(o.Balance) {
		var ret string
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesAccountBalanceV2ResponseResultInner) GetBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *FuturesAccountBalanceV2ResponseResultInner) HasBalance() bool {
	if o != nil && !common.IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given string and assigns it to the Balance field.
func (o *FuturesAccountBalanceV2ResponseResultInner) SetBalance(v string) {
	o.Balance = &v
}

// GetCrossWalletBalance returns the CrossWalletBalance field value if set, zero value otherwise.
func (o *FuturesAccountBalanceV2ResponseResultInner) GetCrossWalletBalance() string {
	if o == nil || common.IsNil(o.CrossWalletBalance) {
		var ret string
		return ret
	}
	return *o.CrossWalletBalance
}

// GetCrossWalletBalanceOk returns a tuple with the CrossWalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesAccountBalanceV2ResponseResultInner) GetCrossWalletBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.CrossWalletBalance) {
		return nil, false
	}
	return o.CrossWalletBalance, true
}

// HasCrossWalletBalance returns a boolean if a field has been set.
func (o *FuturesAccountBalanceV2ResponseResultInner) HasCrossWalletBalance() bool {
	if o != nil && !common.IsNil(o.CrossWalletBalance) {
		return true
	}

	return false
}

// SetCrossWalletBalance gets a reference to the given string and assigns it to the CrossWalletBalance field.
func (o *FuturesAccountBalanceV2ResponseResultInner) SetCrossWalletBalance(v string) {
	o.CrossWalletBalance = &v
}

// GetCrossUnPnl returns the CrossUnPnl field value if set, zero value otherwise.
func (o *FuturesAccountBalanceV2ResponseResultInner) GetCrossUnPnl() string {
	if o == nil || common.IsNil(o.CrossUnPnl) {
		var ret string
		return ret
	}
	return *o.CrossUnPnl
}

// GetCrossUnPnlOk returns a tuple with the CrossUnPnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesAccountBalanceV2ResponseResultInner) GetCrossUnPnlOk() (*string, bool) {
	if o == nil || common.IsNil(o.CrossUnPnl) {
		return nil, false
	}
	return o.CrossUnPnl, true
}

// HasCrossUnPnl returns a boolean if a field has been set.
func (o *FuturesAccountBalanceV2ResponseResultInner) HasCrossUnPnl() bool {
	if o != nil && !common.IsNil(o.CrossUnPnl) {
		return true
	}

	return false
}

// SetCrossUnPnl gets a reference to the given string and assigns it to the CrossUnPnl field.
func (o *FuturesAccountBalanceV2ResponseResultInner) SetCrossUnPnl(v string) {
	o.CrossUnPnl = &v
}

// GetAvailableBalance returns the AvailableBalance field value if set, zero value otherwise.
func (o *FuturesAccountBalanceV2ResponseResultInner) GetAvailableBalance() string {
	if o == nil || common.IsNil(o.AvailableBalance) {
		var ret string
		return ret
	}
	return *o.AvailableBalance
}

// GetAvailableBalanceOk returns a tuple with the AvailableBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesAccountBalanceV2ResponseResultInner) GetAvailableBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.AvailableBalance) {
		return nil, false
	}
	return o.AvailableBalance, true
}

// HasAvailableBalance returns a boolean if a field has been set.
func (o *FuturesAccountBalanceV2ResponseResultInner) HasAvailableBalance() bool {
	if o != nil && !common.IsNil(o.AvailableBalance) {
		return true
	}

	return false
}

// SetAvailableBalance gets a reference to the given string and assigns it to the AvailableBalance field.
func (o *FuturesAccountBalanceV2ResponseResultInner) SetAvailableBalance(v string) {
	o.AvailableBalance = &v
}

// GetMaxWithdrawAmount returns the MaxWithdrawAmount field value if set, zero value otherwise.
func (o *FuturesAccountBalanceV2ResponseResultInner) GetMaxWithdrawAmount() string {
	if o == nil || common.IsNil(o.MaxWithdrawAmount) {
		var ret string
		return ret
	}
	return *o.MaxWithdrawAmount
}

// GetMaxWithdrawAmountOk returns a tuple with the MaxWithdrawAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesAccountBalanceV2ResponseResultInner) GetMaxWithdrawAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxWithdrawAmount) {
		return nil, false
	}
	return o.MaxWithdrawAmount, true
}

// HasMaxWithdrawAmount returns a boolean if a field has been set.
func (o *FuturesAccountBalanceV2ResponseResultInner) HasMaxWithdrawAmount() bool {
	if o != nil && !common.IsNil(o.MaxWithdrawAmount) {
		return true
	}

	return false
}

// SetMaxWithdrawAmount gets a reference to the given string and assigns it to the MaxWithdrawAmount field.
func (o *FuturesAccountBalanceV2ResponseResultInner) SetMaxWithdrawAmount(v string) {
	o.MaxWithdrawAmount = &v
}

// GetMarginAvailable returns the MarginAvailable field value if set, zero value otherwise.
func (o *FuturesAccountBalanceV2ResponseResultInner) GetMarginAvailable() bool {
	if o == nil || common.IsNil(o.MarginAvailable) {
		var ret bool
		return ret
	}
	return *o.MarginAvailable
}

// GetMarginAvailableOk returns a tuple with the MarginAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesAccountBalanceV2ResponseResultInner) GetMarginAvailableOk() (*bool, bool) {
	if o == nil || common.IsNil(o.MarginAvailable) {
		return nil, false
	}
	return o.MarginAvailable, true
}

// HasMarginAvailable returns a boolean if a field has been set.
func (o *FuturesAccountBalanceV2ResponseResultInner) HasMarginAvailable() bool {
	if o != nil && !common.IsNil(o.MarginAvailable) {
		return true
	}

	return false
}

// SetMarginAvailable gets a reference to the given bool and assigns it to the MarginAvailable field.
func (o *FuturesAccountBalanceV2ResponseResultInner) SetMarginAvailable(v bool) {
	o.MarginAvailable = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *FuturesAccountBalanceV2ResponseResultInner) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesAccountBalanceV2ResponseResultInner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *FuturesAccountBalanceV2ResponseResultInner) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *FuturesAccountBalanceV2ResponseResultInner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o FuturesAccountBalanceV2ResponseResultInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FuturesAccountBalanceV2ResponseResultInner) ToMap() (map[string]interface{}, error) {
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

func (o *FuturesAccountBalanceV2ResponseResultInner) UnmarshalJSON(data []byte) (err error) {
	varFuturesAccountBalanceV2ResponseResultInner := _FuturesAccountBalanceV2ResponseResultInner{}

	err = json.Unmarshal(data, &varFuturesAccountBalanceV2ResponseResultInner)

	if err != nil {
		return err
	}

	*o = FuturesAccountBalanceV2ResponseResultInner(varFuturesAccountBalanceV2ResponseResultInner)

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

type NullableFuturesAccountBalanceV2ResponseResultInner struct {
	value *FuturesAccountBalanceV2ResponseResultInner
	isSet bool
}

func (v NullableFuturesAccountBalanceV2ResponseResultInner) Get() *FuturesAccountBalanceV2ResponseResultInner {
	return v.value
}

func (v *NullableFuturesAccountBalanceV2ResponseResultInner) Set(val *FuturesAccountBalanceV2ResponseResultInner) {
	v.value = val
	v.isSet = true
}

func (v NullableFuturesAccountBalanceV2ResponseResultInner) IsSet() bool {
	return v.isSet
}

func (v *NullableFuturesAccountBalanceV2ResponseResultInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFuturesAccountBalanceV2ResponseResultInner(val *FuturesAccountBalanceV2ResponseResultInner) *NullableFuturesAccountBalanceV2ResponseResultInner {
	return &NullableFuturesAccountBalanceV2ResponseResultInner{value: val, isSet: true}
}

func (v NullableFuturesAccountBalanceV2ResponseResultInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFuturesAccountBalanceV2ResponseResultInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
