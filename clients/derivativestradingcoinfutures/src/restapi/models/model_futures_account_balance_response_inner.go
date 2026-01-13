/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the FuturesAccountBalanceResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FuturesAccountBalanceResponseInner{}

// FuturesAccountBalanceResponseInner struct for FuturesAccountBalanceResponseInner
type FuturesAccountBalanceResponseInner struct {
	AccountAlias         *string `json:"accountAlias,omitempty"`
	Asset                *string `json:"asset,omitempty"`
	Balance              *string `json:"balance,omitempty"`
	WithdrawAvailable    *string `json:"withdrawAvailable,omitempty"`
	CrossWalletBalance   *string `json:"crossWalletBalance,omitempty"`
	CrossUnPnl           *string `json:"crossUnPnl,omitempty"`
	AvailableBalance     *string `json:"availableBalance,omitempty"`
	UpdateTime           *int64  `json:"updateTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FuturesAccountBalanceResponseInner FuturesAccountBalanceResponseInner

// NewFuturesAccountBalanceResponseInner instantiates a new FuturesAccountBalanceResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFuturesAccountBalanceResponseInner() *FuturesAccountBalanceResponseInner {
	this := FuturesAccountBalanceResponseInner{}
	return &this
}

// NewFuturesAccountBalanceResponseInnerWithDefaults instantiates a new FuturesAccountBalanceResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFuturesAccountBalanceResponseInnerWithDefaults() *FuturesAccountBalanceResponseInner {
	this := FuturesAccountBalanceResponseInner{}
	return &this
}

// GetAccountAlias returns the AccountAlias field value if set, zero value otherwise.
func (o *FuturesAccountBalanceResponseInner) GetAccountAlias() string {
	if o == nil || common.IsNil(o.AccountAlias) {
		var ret string
		return ret
	}
	return *o.AccountAlias
}

// GetAccountAliasOk returns a tuple with the AccountAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesAccountBalanceResponseInner) GetAccountAliasOk() (*string, bool) {
	if o == nil || common.IsNil(o.AccountAlias) {
		return nil, false
	}
	return o.AccountAlias, true
}

// HasAccountAlias returns a boolean if a field has been set.
func (o *FuturesAccountBalanceResponseInner) HasAccountAlias() bool {
	if o != nil && !common.IsNil(o.AccountAlias) {
		return true
	}

	return false
}

// SetAccountAlias gets a reference to the given string and assigns it to the AccountAlias field.
func (o *FuturesAccountBalanceResponseInner) SetAccountAlias(v string) {
	o.AccountAlias = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *FuturesAccountBalanceResponseInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesAccountBalanceResponseInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *FuturesAccountBalanceResponseInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *FuturesAccountBalanceResponseInner) SetAsset(v string) {
	o.Asset = &v
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *FuturesAccountBalanceResponseInner) GetBalance() string {
	if o == nil || common.IsNil(o.Balance) {
		var ret string
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesAccountBalanceResponseInner) GetBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *FuturesAccountBalanceResponseInner) HasBalance() bool {
	if o != nil && !common.IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given string and assigns it to the Balance field.
func (o *FuturesAccountBalanceResponseInner) SetBalance(v string) {
	o.Balance = &v
}

// GetWithdrawAvailable returns the WithdrawAvailable field value if set, zero value otherwise.
func (o *FuturesAccountBalanceResponseInner) GetWithdrawAvailable() string {
	if o == nil || common.IsNil(o.WithdrawAvailable) {
		var ret string
		return ret
	}
	return *o.WithdrawAvailable
}

// GetWithdrawAvailableOk returns a tuple with the WithdrawAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesAccountBalanceResponseInner) GetWithdrawAvailableOk() (*string, bool) {
	if o == nil || common.IsNil(o.WithdrawAvailable) {
		return nil, false
	}
	return o.WithdrawAvailable, true
}

// HasWithdrawAvailable returns a boolean if a field has been set.
func (o *FuturesAccountBalanceResponseInner) HasWithdrawAvailable() bool {
	if o != nil && !common.IsNil(o.WithdrawAvailable) {
		return true
	}

	return false
}

// SetWithdrawAvailable gets a reference to the given string and assigns it to the WithdrawAvailable field.
func (o *FuturesAccountBalanceResponseInner) SetWithdrawAvailable(v string) {
	o.WithdrawAvailable = &v
}

// GetCrossWalletBalance returns the CrossWalletBalance field value if set, zero value otherwise.
func (o *FuturesAccountBalanceResponseInner) GetCrossWalletBalance() string {
	if o == nil || common.IsNil(o.CrossWalletBalance) {
		var ret string
		return ret
	}
	return *o.CrossWalletBalance
}

// GetCrossWalletBalanceOk returns a tuple with the CrossWalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesAccountBalanceResponseInner) GetCrossWalletBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.CrossWalletBalance) {
		return nil, false
	}
	return o.CrossWalletBalance, true
}

// HasCrossWalletBalance returns a boolean if a field has been set.
func (o *FuturesAccountBalanceResponseInner) HasCrossWalletBalance() bool {
	if o != nil && !common.IsNil(o.CrossWalletBalance) {
		return true
	}

	return false
}

// SetCrossWalletBalance gets a reference to the given string and assigns it to the CrossWalletBalance field.
func (o *FuturesAccountBalanceResponseInner) SetCrossWalletBalance(v string) {
	o.CrossWalletBalance = &v
}

// GetCrossUnPnl returns the CrossUnPnl field value if set, zero value otherwise.
func (o *FuturesAccountBalanceResponseInner) GetCrossUnPnl() string {
	if o == nil || common.IsNil(o.CrossUnPnl) {
		var ret string
		return ret
	}
	return *o.CrossUnPnl
}

// GetCrossUnPnlOk returns a tuple with the CrossUnPnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesAccountBalanceResponseInner) GetCrossUnPnlOk() (*string, bool) {
	if o == nil || common.IsNil(o.CrossUnPnl) {
		return nil, false
	}
	return o.CrossUnPnl, true
}

// HasCrossUnPnl returns a boolean if a field has been set.
func (o *FuturesAccountBalanceResponseInner) HasCrossUnPnl() bool {
	if o != nil && !common.IsNil(o.CrossUnPnl) {
		return true
	}

	return false
}

// SetCrossUnPnl gets a reference to the given string and assigns it to the CrossUnPnl field.
func (o *FuturesAccountBalanceResponseInner) SetCrossUnPnl(v string) {
	o.CrossUnPnl = &v
}

// GetAvailableBalance returns the AvailableBalance field value if set, zero value otherwise.
func (o *FuturesAccountBalanceResponseInner) GetAvailableBalance() string {
	if o == nil || common.IsNil(o.AvailableBalance) {
		var ret string
		return ret
	}
	return *o.AvailableBalance
}

// GetAvailableBalanceOk returns a tuple with the AvailableBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesAccountBalanceResponseInner) GetAvailableBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.AvailableBalance) {
		return nil, false
	}
	return o.AvailableBalance, true
}

// HasAvailableBalance returns a boolean if a field has been set.
func (o *FuturesAccountBalanceResponseInner) HasAvailableBalance() bool {
	if o != nil && !common.IsNil(o.AvailableBalance) {
		return true
	}

	return false
}

// SetAvailableBalance gets a reference to the given string and assigns it to the AvailableBalance field.
func (o *FuturesAccountBalanceResponseInner) SetAvailableBalance(v string) {
	o.AvailableBalance = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *FuturesAccountBalanceResponseInner) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesAccountBalanceResponseInner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *FuturesAccountBalanceResponseInner) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *FuturesAccountBalanceResponseInner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o FuturesAccountBalanceResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FuturesAccountBalanceResponseInner) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.WithdrawAvailable) {
		toSerialize["withdrawAvailable"] = o.WithdrawAvailable
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
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FuturesAccountBalanceResponseInner) UnmarshalJSON(data []byte) (err error) {
	varFuturesAccountBalanceResponseInner := _FuturesAccountBalanceResponseInner{}

	err = json.Unmarshal(data, &varFuturesAccountBalanceResponseInner)

	if err != nil {
		return err
	}

	*o = FuturesAccountBalanceResponseInner(varFuturesAccountBalanceResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accountAlias")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "balance")
		delete(additionalProperties, "withdrawAvailable")
		delete(additionalProperties, "crossWalletBalance")
		delete(additionalProperties, "crossUnPnl")
		delete(additionalProperties, "availableBalance")
		delete(additionalProperties, "updateTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFuturesAccountBalanceResponseInner struct {
	value *FuturesAccountBalanceResponseInner
	isSet bool
}

func (v NullableFuturesAccountBalanceResponseInner) Get() *FuturesAccountBalanceResponseInner {
	return v.value
}

func (v *NullableFuturesAccountBalanceResponseInner) Set(val *FuturesAccountBalanceResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableFuturesAccountBalanceResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableFuturesAccountBalanceResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFuturesAccountBalanceResponseInner(val *FuturesAccountBalanceResponseInner) *NullableFuturesAccountBalanceResponseInner {
	return &NullableFuturesAccountBalanceResponseInner{value: val, isSet: true}
}

func (v NullableFuturesAccountBalanceResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFuturesAccountBalanceResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
