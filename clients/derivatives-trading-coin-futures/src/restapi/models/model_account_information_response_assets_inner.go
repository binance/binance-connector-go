/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AccountInformationResponseAssetsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountInformationResponseAssetsInner{}

// AccountInformationResponseAssetsInner struct for AccountInformationResponseAssetsInner
type AccountInformationResponseAssetsInner struct {
	Asset                  *string `json:"asset,omitempty"`
	WalletBalance          *string `json:"walletBalance,omitempty"`
	UnrealizedProfit       *string `json:"unrealizedProfit,omitempty"`
	MarginBalance          *string `json:"marginBalance,omitempty"`
	MaintMargin            *string `json:"maintMargin,omitempty"`
	InitialMargin          *string `json:"initialMargin,omitempty"`
	PositionInitialMargin  *string `json:"positionInitialMargin,omitempty"`
	OpenOrderInitialMargin *string `json:"openOrderInitialMargin,omitempty"`
	MaxWithdrawAmount      *string `json:"maxWithdrawAmount,omitempty"`
	CrossWalletBalance     *string `json:"crossWalletBalance,omitempty"`
	CrossUnPnl             *string `json:"crossUnPnl,omitempty"`
	AvailableBalance       *string `json:"availableBalance,omitempty"`
	UpdateTime             *int64  `json:"updateTime,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _AccountInformationResponseAssetsInner AccountInformationResponseAssetsInner

// NewAccountInformationResponseAssetsInner instantiates a new AccountInformationResponseAssetsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountInformationResponseAssetsInner() *AccountInformationResponseAssetsInner {
	this := AccountInformationResponseAssetsInner{}
	return &this
}

// NewAccountInformationResponseAssetsInnerWithDefaults instantiates a new AccountInformationResponseAssetsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountInformationResponseAssetsInnerWithDefaults() *AccountInformationResponseAssetsInner {
	this := AccountInformationResponseAssetsInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *AccountInformationResponseAssetsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseAssetsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *AccountInformationResponseAssetsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *AccountInformationResponseAssetsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetWalletBalance returns the WalletBalance field value if set, zero value otherwise.
func (o *AccountInformationResponseAssetsInner) GetWalletBalance() string {
	if o == nil || common.IsNil(o.WalletBalance) {
		var ret string
		return ret
	}
	return *o.WalletBalance
}

// GetWalletBalanceOk returns a tuple with the WalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseAssetsInner) GetWalletBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.WalletBalance) {
		return nil, false
	}
	return o.WalletBalance, true
}

// HasWalletBalance returns a boolean if a field has been set.
func (o *AccountInformationResponseAssetsInner) HasWalletBalance() bool {
	if o != nil && !common.IsNil(o.WalletBalance) {
		return true
	}

	return false
}

// SetWalletBalance gets a reference to the given string and assigns it to the WalletBalance field.
func (o *AccountInformationResponseAssetsInner) SetWalletBalance(v string) {
	o.WalletBalance = &v
}

// GetUnrealizedProfit returns the UnrealizedProfit field value if set, zero value otherwise.
func (o *AccountInformationResponseAssetsInner) GetUnrealizedProfit() string {
	if o == nil || common.IsNil(o.UnrealizedProfit) {
		var ret string
		return ret
	}
	return *o.UnrealizedProfit
}

// GetUnrealizedProfitOk returns a tuple with the UnrealizedProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseAssetsInner) GetUnrealizedProfitOk() (*string, bool) {
	if o == nil || common.IsNil(o.UnrealizedProfit) {
		return nil, false
	}
	return o.UnrealizedProfit, true
}

// HasUnrealizedProfit returns a boolean if a field has been set.
func (o *AccountInformationResponseAssetsInner) HasUnrealizedProfit() bool {
	if o != nil && !common.IsNil(o.UnrealizedProfit) {
		return true
	}

	return false
}

// SetUnrealizedProfit gets a reference to the given string and assigns it to the UnrealizedProfit field.
func (o *AccountInformationResponseAssetsInner) SetUnrealizedProfit(v string) {
	o.UnrealizedProfit = &v
}

// GetMarginBalance returns the MarginBalance field value if set, zero value otherwise.
func (o *AccountInformationResponseAssetsInner) GetMarginBalance() string {
	if o == nil || common.IsNil(o.MarginBalance) {
		var ret string
		return ret
	}
	return *o.MarginBalance
}

// GetMarginBalanceOk returns a tuple with the MarginBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseAssetsInner) GetMarginBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarginBalance) {
		return nil, false
	}
	return o.MarginBalance, true
}

// HasMarginBalance returns a boolean if a field has been set.
func (o *AccountInformationResponseAssetsInner) HasMarginBalance() bool {
	if o != nil && !common.IsNil(o.MarginBalance) {
		return true
	}

	return false
}

// SetMarginBalance gets a reference to the given string and assigns it to the MarginBalance field.
func (o *AccountInformationResponseAssetsInner) SetMarginBalance(v string) {
	o.MarginBalance = &v
}

// GetMaintMargin returns the MaintMargin field value if set, zero value otherwise.
func (o *AccountInformationResponseAssetsInner) GetMaintMargin() string {
	if o == nil || common.IsNil(o.MaintMargin) {
		var ret string
		return ret
	}
	return *o.MaintMargin
}

// GetMaintMarginOk returns a tuple with the MaintMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseAssetsInner) GetMaintMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaintMargin) {
		return nil, false
	}
	return o.MaintMargin, true
}

// HasMaintMargin returns a boolean if a field has been set.
func (o *AccountInformationResponseAssetsInner) HasMaintMargin() bool {
	if o != nil && !common.IsNil(o.MaintMargin) {
		return true
	}

	return false
}

// SetMaintMargin gets a reference to the given string and assigns it to the MaintMargin field.
func (o *AccountInformationResponseAssetsInner) SetMaintMargin(v string) {
	o.MaintMargin = &v
}

// GetInitialMargin returns the InitialMargin field value if set, zero value otherwise.
func (o *AccountInformationResponseAssetsInner) GetInitialMargin() string {
	if o == nil || common.IsNil(o.InitialMargin) {
		var ret string
		return ret
	}
	return *o.InitialMargin
}

// GetInitialMarginOk returns a tuple with the InitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseAssetsInner) GetInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.InitialMargin) {
		return nil, false
	}
	return o.InitialMargin, true
}

// HasInitialMargin returns a boolean if a field has been set.
func (o *AccountInformationResponseAssetsInner) HasInitialMargin() bool {
	if o != nil && !common.IsNil(o.InitialMargin) {
		return true
	}

	return false
}

// SetInitialMargin gets a reference to the given string and assigns it to the InitialMargin field.
func (o *AccountInformationResponseAssetsInner) SetInitialMargin(v string) {
	o.InitialMargin = &v
}

// GetPositionInitialMargin returns the PositionInitialMargin field value if set, zero value otherwise.
func (o *AccountInformationResponseAssetsInner) GetPositionInitialMargin() string {
	if o == nil || common.IsNil(o.PositionInitialMargin) {
		var ret string
		return ret
	}
	return *o.PositionInitialMargin
}

// GetPositionInitialMarginOk returns a tuple with the PositionInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseAssetsInner) GetPositionInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionInitialMargin) {
		return nil, false
	}
	return o.PositionInitialMargin, true
}

// HasPositionInitialMargin returns a boolean if a field has been set.
func (o *AccountInformationResponseAssetsInner) HasPositionInitialMargin() bool {
	if o != nil && !common.IsNil(o.PositionInitialMargin) {
		return true
	}

	return false
}

// SetPositionInitialMargin gets a reference to the given string and assigns it to the PositionInitialMargin field.
func (o *AccountInformationResponseAssetsInner) SetPositionInitialMargin(v string) {
	o.PositionInitialMargin = &v
}

// GetOpenOrderInitialMargin returns the OpenOrderInitialMargin field value if set, zero value otherwise.
func (o *AccountInformationResponseAssetsInner) GetOpenOrderInitialMargin() string {
	if o == nil || common.IsNil(o.OpenOrderInitialMargin) {
		var ret string
		return ret
	}
	return *o.OpenOrderInitialMargin
}

// GetOpenOrderInitialMarginOk returns a tuple with the OpenOrderInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseAssetsInner) GetOpenOrderInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.OpenOrderInitialMargin) {
		return nil, false
	}
	return o.OpenOrderInitialMargin, true
}

// HasOpenOrderInitialMargin returns a boolean if a field has been set.
func (o *AccountInformationResponseAssetsInner) HasOpenOrderInitialMargin() bool {
	if o != nil && !common.IsNil(o.OpenOrderInitialMargin) {
		return true
	}

	return false
}

// SetOpenOrderInitialMargin gets a reference to the given string and assigns it to the OpenOrderInitialMargin field.
func (o *AccountInformationResponseAssetsInner) SetOpenOrderInitialMargin(v string) {
	o.OpenOrderInitialMargin = &v
}

// GetMaxWithdrawAmount returns the MaxWithdrawAmount field value if set, zero value otherwise.
func (o *AccountInformationResponseAssetsInner) GetMaxWithdrawAmount() string {
	if o == nil || common.IsNil(o.MaxWithdrawAmount) {
		var ret string
		return ret
	}
	return *o.MaxWithdrawAmount
}

// GetMaxWithdrawAmountOk returns a tuple with the MaxWithdrawAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseAssetsInner) GetMaxWithdrawAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxWithdrawAmount) {
		return nil, false
	}
	return o.MaxWithdrawAmount, true
}

// HasMaxWithdrawAmount returns a boolean if a field has been set.
func (o *AccountInformationResponseAssetsInner) HasMaxWithdrawAmount() bool {
	if o != nil && !common.IsNil(o.MaxWithdrawAmount) {
		return true
	}

	return false
}

// SetMaxWithdrawAmount gets a reference to the given string and assigns it to the MaxWithdrawAmount field.
func (o *AccountInformationResponseAssetsInner) SetMaxWithdrawAmount(v string) {
	o.MaxWithdrawAmount = &v
}

// GetCrossWalletBalance returns the CrossWalletBalance field value if set, zero value otherwise.
func (o *AccountInformationResponseAssetsInner) GetCrossWalletBalance() string {
	if o == nil || common.IsNil(o.CrossWalletBalance) {
		var ret string
		return ret
	}
	return *o.CrossWalletBalance
}

// GetCrossWalletBalanceOk returns a tuple with the CrossWalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseAssetsInner) GetCrossWalletBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.CrossWalletBalance) {
		return nil, false
	}
	return o.CrossWalletBalance, true
}

// HasCrossWalletBalance returns a boolean if a field has been set.
func (o *AccountInformationResponseAssetsInner) HasCrossWalletBalance() bool {
	if o != nil && !common.IsNil(o.CrossWalletBalance) {
		return true
	}

	return false
}

// SetCrossWalletBalance gets a reference to the given string and assigns it to the CrossWalletBalance field.
func (o *AccountInformationResponseAssetsInner) SetCrossWalletBalance(v string) {
	o.CrossWalletBalance = &v
}

// GetCrossUnPnl returns the CrossUnPnl field value if set, zero value otherwise.
func (o *AccountInformationResponseAssetsInner) GetCrossUnPnl() string {
	if o == nil || common.IsNil(o.CrossUnPnl) {
		var ret string
		return ret
	}
	return *o.CrossUnPnl
}

// GetCrossUnPnlOk returns a tuple with the CrossUnPnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseAssetsInner) GetCrossUnPnlOk() (*string, bool) {
	if o == nil || common.IsNil(o.CrossUnPnl) {
		return nil, false
	}
	return o.CrossUnPnl, true
}

// HasCrossUnPnl returns a boolean if a field has been set.
func (o *AccountInformationResponseAssetsInner) HasCrossUnPnl() bool {
	if o != nil && !common.IsNil(o.CrossUnPnl) {
		return true
	}

	return false
}

// SetCrossUnPnl gets a reference to the given string and assigns it to the CrossUnPnl field.
func (o *AccountInformationResponseAssetsInner) SetCrossUnPnl(v string) {
	o.CrossUnPnl = &v
}

// GetAvailableBalance returns the AvailableBalance field value if set, zero value otherwise.
func (o *AccountInformationResponseAssetsInner) GetAvailableBalance() string {
	if o == nil || common.IsNil(o.AvailableBalance) {
		var ret string
		return ret
	}
	return *o.AvailableBalance
}

// GetAvailableBalanceOk returns a tuple with the AvailableBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseAssetsInner) GetAvailableBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.AvailableBalance) {
		return nil, false
	}
	return o.AvailableBalance, true
}

// HasAvailableBalance returns a boolean if a field has been set.
func (o *AccountInformationResponseAssetsInner) HasAvailableBalance() bool {
	if o != nil && !common.IsNil(o.AvailableBalance) {
		return true
	}

	return false
}

// SetAvailableBalance gets a reference to the given string and assigns it to the AvailableBalance field.
func (o *AccountInformationResponseAssetsInner) SetAvailableBalance(v string) {
	o.AvailableBalance = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *AccountInformationResponseAssetsInner) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseAssetsInner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *AccountInformationResponseAssetsInner) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *AccountInformationResponseAssetsInner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o AccountInformationResponseAssetsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountInformationResponseAssetsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.WalletBalance) {
		toSerialize["walletBalance"] = o.WalletBalance
	}
	if !common.IsNil(o.UnrealizedProfit) {
		toSerialize["unrealizedProfit"] = o.UnrealizedProfit
	}
	if !common.IsNil(o.MarginBalance) {
		toSerialize["marginBalance"] = o.MarginBalance
	}
	if !common.IsNil(o.MaintMargin) {
		toSerialize["maintMargin"] = o.MaintMargin
	}
	if !common.IsNil(o.InitialMargin) {
		toSerialize["initialMargin"] = o.InitialMargin
	}
	if !common.IsNil(o.PositionInitialMargin) {
		toSerialize["positionInitialMargin"] = o.PositionInitialMargin
	}
	if !common.IsNil(o.OpenOrderInitialMargin) {
		toSerialize["openOrderInitialMargin"] = o.OpenOrderInitialMargin
	}
	if !common.IsNil(o.MaxWithdrawAmount) {
		toSerialize["maxWithdrawAmount"] = o.MaxWithdrawAmount
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

func (o *AccountInformationResponseAssetsInner) UnmarshalJSON(data []byte) (err error) {
	varAccountInformationResponseAssetsInner := _AccountInformationResponseAssetsInner{}

	err = json.Unmarshal(data, &varAccountInformationResponseAssetsInner)

	if err != nil {
		return err
	}

	*o = AccountInformationResponseAssetsInner(varAccountInformationResponseAssetsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "walletBalance")
		delete(additionalProperties, "unrealizedProfit")
		delete(additionalProperties, "marginBalance")
		delete(additionalProperties, "maintMargin")
		delete(additionalProperties, "initialMargin")
		delete(additionalProperties, "positionInitialMargin")
		delete(additionalProperties, "openOrderInitialMargin")
		delete(additionalProperties, "maxWithdrawAmount")
		delete(additionalProperties, "crossWalletBalance")
		delete(additionalProperties, "crossUnPnl")
		delete(additionalProperties, "availableBalance")
		delete(additionalProperties, "updateTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountInformationResponseAssetsInner struct {
	value *AccountInformationResponseAssetsInner
	isSet bool
}

func (v NullableAccountInformationResponseAssetsInner) Get() *AccountInformationResponseAssetsInner {
	return v.value
}

func (v *NullableAccountInformationResponseAssetsInner) Set(val *AccountInformationResponseAssetsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountInformationResponseAssetsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountInformationResponseAssetsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountInformationResponseAssetsInner(val *AccountInformationResponseAssetsInner) *NullableAccountInformationResponseAssetsInner {
	return &NullableAccountInformationResponseAssetsInner{value: val, isSet: true}
}

func (v NullableAccountInformationResponseAssetsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountInformationResponseAssetsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
