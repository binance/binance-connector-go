/*
Binance Derivatives Trading COIN Futures WebSocket API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures WebSocket API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AccountInformationResponseResultPositionsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountInformationResponseResultPositionsInner{}

// AccountInformationResponseResultPositionsInner struct for AccountInformationResponseResultPositionsInner
type AccountInformationResponseResultPositionsInner struct {
	Symbol                 *string `json:"symbol,omitempty"`
	InitialMargin          *string `json:"initialMargin,omitempty"`
	MaintMargin            *string `json:"maintMargin,omitempty"`
	UnrealizedProfit       *string `json:"unrealizedProfit,omitempty"`
	PositionInitialMargin  *string `json:"positionInitialMargin,omitempty"`
	OpenOrderInitialMargin *string `json:"openOrderInitialMargin,omitempty"`
	Leverage               *string `json:"leverage,omitempty"`
	Isolated               *bool   `json:"isolated,omitempty"`
	PositionSide           *string `json:"positionSide,omitempty"`
	EntryPrice             *string `json:"entryPrice,omitempty"`
	MaxQty                 *string `json:"maxQty,omitempty"`
	NotionalValue          *string `json:"notionalValue,omitempty"`
	IsolatedWallet         *string `json:"isolatedWallet,omitempty"`
	UpdateTime             *int64  `json:"updateTime,omitempty"`
	PositionAmt            *string `json:"positionAmt,omitempty"`
	BreakEvenPrice         *string `json:"breakEvenPrice,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _AccountInformationResponseResultPositionsInner AccountInformationResponseResultPositionsInner

// NewAccountInformationResponseResultPositionsInner instantiates a new AccountInformationResponseResultPositionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountInformationResponseResultPositionsInner() *AccountInformationResponseResultPositionsInner {
	this := AccountInformationResponseResultPositionsInner{}
	return &this
}

// NewAccountInformationResponseResultPositionsInnerWithDefaults instantiates a new AccountInformationResponseResultPositionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountInformationResponseResultPositionsInnerWithDefaults() *AccountInformationResponseResultPositionsInner {
	this := AccountInformationResponseResultPositionsInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *AccountInformationResponseResultPositionsInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResultPositionsInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *AccountInformationResponseResultPositionsInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *AccountInformationResponseResultPositionsInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetInitialMargin returns the InitialMargin field value if set, zero value otherwise.
func (o *AccountInformationResponseResultPositionsInner) GetInitialMargin() string {
	if o == nil || common.IsNil(o.InitialMargin) {
		var ret string
		return ret
	}
	return *o.InitialMargin
}

// GetInitialMarginOk returns a tuple with the InitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResultPositionsInner) GetInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.InitialMargin) {
		return nil, false
	}
	return o.InitialMargin, true
}

// HasInitialMargin returns a boolean if a field has been set.
func (o *AccountInformationResponseResultPositionsInner) HasInitialMargin() bool {
	if o != nil && !common.IsNil(o.InitialMargin) {
		return true
	}

	return false
}

// SetInitialMargin gets a reference to the given string and assigns it to the InitialMargin field.
func (o *AccountInformationResponseResultPositionsInner) SetInitialMargin(v string) {
	o.InitialMargin = &v
}

// GetMaintMargin returns the MaintMargin field value if set, zero value otherwise.
func (o *AccountInformationResponseResultPositionsInner) GetMaintMargin() string {
	if o == nil || common.IsNil(o.MaintMargin) {
		var ret string
		return ret
	}
	return *o.MaintMargin
}

// GetMaintMarginOk returns a tuple with the MaintMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResultPositionsInner) GetMaintMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaintMargin) {
		return nil, false
	}
	return o.MaintMargin, true
}

// HasMaintMargin returns a boolean if a field has been set.
func (o *AccountInformationResponseResultPositionsInner) HasMaintMargin() bool {
	if o != nil && !common.IsNil(o.MaintMargin) {
		return true
	}

	return false
}

// SetMaintMargin gets a reference to the given string and assigns it to the MaintMargin field.
func (o *AccountInformationResponseResultPositionsInner) SetMaintMargin(v string) {
	o.MaintMargin = &v
}

// GetUnrealizedProfit returns the UnrealizedProfit field value if set, zero value otherwise.
func (o *AccountInformationResponseResultPositionsInner) GetUnrealizedProfit() string {
	if o == nil || common.IsNil(o.UnrealizedProfit) {
		var ret string
		return ret
	}
	return *o.UnrealizedProfit
}

// GetUnrealizedProfitOk returns a tuple with the UnrealizedProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResultPositionsInner) GetUnrealizedProfitOk() (*string, bool) {
	if o == nil || common.IsNil(o.UnrealizedProfit) {
		return nil, false
	}
	return o.UnrealizedProfit, true
}

// HasUnrealizedProfit returns a boolean if a field has been set.
func (o *AccountInformationResponseResultPositionsInner) HasUnrealizedProfit() bool {
	if o != nil && !common.IsNil(o.UnrealizedProfit) {
		return true
	}

	return false
}

// SetUnrealizedProfit gets a reference to the given string and assigns it to the UnrealizedProfit field.
func (o *AccountInformationResponseResultPositionsInner) SetUnrealizedProfit(v string) {
	o.UnrealizedProfit = &v
}

// GetPositionInitialMargin returns the PositionInitialMargin field value if set, zero value otherwise.
func (o *AccountInformationResponseResultPositionsInner) GetPositionInitialMargin() string {
	if o == nil || common.IsNil(o.PositionInitialMargin) {
		var ret string
		return ret
	}
	return *o.PositionInitialMargin
}

// GetPositionInitialMarginOk returns a tuple with the PositionInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResultPositionsInner) GetPositionInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionInitialMargin) {
		return nil, false
	}
	return o.PositionInitialMargin, true
}

// HasPositionInitialMargin returns a boolean if a field has been set.
func (o *AccountInformationResponseResultPositionsInner) HasPositionInitialMargin() bool {
	if o != nil && !common.IsNil(o.PositionInitialMargin) {
		return true
	}

	return false
}

// SetPositionInitialMargin gets a reference to the given string and assigns it to the PositionInitialMargin field.
func (o *AccountInformationResponseResultPositionsInner) SetPositionInitialMargin(v string) {
	o.PositionInitialMargin = &v
}

// GetOpenOrderInitialMargin returns the OpenOrderInitialMargin field value if set, zero value otherwise.
func (o *AccountInformationResponseResultPositionsInner) GetOpenOrderInitialMargin() string {
	if o == nil || common.IsNil(o.OpenOrderInitialMargin) {
		var ret string
		return ret
	}
	return *o.OpenOrderInitialMargin
}

// GetOpenOrderInitialMarginOk returns a tuple with the OpenOrderInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResultPositionsInner) GetOpenOrderInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.OpenOrderInitialMargin) {
		return nil, false
	}
	return o.OpenOrderInitialMargin, true
}

// HasOpenOrderInitialMargin returns a boolean if a field has been set.
func (o *AccountInformationResponseResultPositionsInner) HasOpenOrderInitialMargin() bool {
	if o != nil && !common.IsNil(o.OpenOrderInitialMargin) {
		return true
	}

	return false
}

// SetOpenOrderInitialMargin gets a reference to the given string and assigns it to the OpenOrderInitialMargin field.
func (o *AccountInformationResponseResultPositionsInner) SetOpenOrderInitialMargin(v string) {
	o.OpenOrderInitialMargin = &v
}

// GetLeverage returns the Leverage field value if set, zero value otherwise.
func (o *AccountInformationResponseResultPositionsInner) GetLeverage() string {
	if o == nil || common.IsNil(o.Leverage) {
		var ret string
		return ret
	}
	return *o.Leverage
}

// GetLeverageOk returns a tuple with the Leverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResultPositionsInner) GetLeverageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Leverage) {
		return nil, false
	}
	return o.Leverage, true
}

// HasLeverage returns a boolean if a field has been set.
func (o *AccountInformationResponseResultPositionsInner) HasLeverage() bool {
	if o != nil && !common.IsNil(o.Leverage) {
		return true
	}

	return false
}

// SetLeverage gets a reference to the given string and assigns it to the Leverage field.
func (o *AccountInformationResponseResultPositionsInner) SetLeverage(v string) {
	o.Leverage = &v
}

// GetIsolated returns the Isolated field value if set, zero value otherwise.
func (o *AccountInformationResponseResultPositionsInner) GetIsolated() bool {
	if o == nil || common.IsNil(o.Isolated) {
		var ret bool
		return ret
	}
	return *o.Isolated
}

// GetIsolatedOk returns a tuple with the Isolated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResultPositionsInner) GetIsolatedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Isolated) {
		return nil, false
	}
	return o.Isolated, true
}

// HasIsolated returns a boolean if a field has been set.
func (o *AccountInformationResponseResultPositionsInner) HasIsolated() bool {
	if o != nil && !common.IsNil(o.Isolated) {
		return true
	}

	return false
}

// SetIsolated gets a reference to the given bool and assigns it to the Isolated field.
func (o *AccountInformationResponseResultPositionsInner) SetIsolated(v bool) {
	o.Isolated = &v
}

// GetPositionSide returns the PositionSide field value if set, zero value otherwise.
func (o *AccountInformationResponseResultPositionsInner) GetPositionSide() string {
	if o == nil || common.IsNil(o.PositionSide) {
		var ret string
		return ret
	}
	return *o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResultPositionsInner) GetPositionSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionSide) {
		return nil, false
	}
	return o.PositionSide, true
}

// HasPositionSide returns a boolean if a field has been set.
func (o *AccountInformationResponseResultPositionsInner) HasPositionSide() bool {
	if o != nil && !common.IsNil(o.PositionSide) {
		return true
	}

	return false
}

// SetPositionSide gets a reference to the given string and assigns it to the PositionSide field.
func (o *AccountInformationResponseResultPositionsInner) SetPositionSide(v string) {
	o.PositionSide = &v
}

// GetEntryPrice returns the EntryPrice field value if set, zero value otherwise.
func (o *AccountInformationResponseResultPositionsInner) GetEntryPrice() string {
	if o == nil || common.IsNil(o.EntryPrice) {
		var ret string
		return ret
	}
	return *o.EntryPrice
}

// GetEntryPriceOk returns a tuple with the EntryPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResultPositionsInner) GetEntryPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.EntryPrice) {
		return nil, false
	}
	return o.EntryPrice, true
}

// HasEntryPrice returns a boolean if a field has been set.
func (o *AccountInformationResponseResultPositionsInner) HasEntryPrice() bool {
	if o != nil && !common.IsNil(o.EntryPrice) {
		return true
	}

	return false
}

// SetEntryPrice gets a reference to the given string and assigns it to the EntryPrice field.
func (o *AccountInformationResponseResultPositionsInner) SetEntryPrice(v string) {
	o.EntryPrice = &v
}

// GetMaxQty returns the MaxQty field value if set, zero value otherwise.
func (o *AccountInformationResponseResultPositionsInner) GetMaxQty() string {
	if o == nil || common.IsNil(o.MaxQty) {
		var ret string
		return ret
	}
	return *o.MaxQty
}

// GetMaxQtyOk returns a tuple with the MaxQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResultPositionsInner) GetMaxQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxQty) {
		return nil, false
	}
	return o.MaxQty, true
}

// HasMaxQty returns a boolean if a field has been set.
func (o *AccountInformationResponseResultPositionsInner) HasMaxQty() bool {
	if o != nil && !common.IsNil(o.MaxQty) {
		return true
	}

	return false
}

// SetMaxQty gets a reference to the given string and assigns it to the MaxQty field.
func (o *AccountInformationResponseResultPositionsInner) SetMaxQty(v string) {
	o.MaxQty = &v
}

// GetNotionalValue returns the NotionalValue field value if set, zero value otherwise.
func (o *AccountInformationResponseResultPositionsInner) GetNotionalValue() string {
	if o == nil || common.IsNil(o.NotionalValue) {
		var ret string
		return ret
	}
	return *o.NotionalValue
}

// GetNotionalValueOk returns a tuple with the NotionalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResultPositionsInner) GetNotionalValueOk() (*string, bool) {
	if o == nil || common.IsNil(o.NotionalValue) {
		return nil, false
	}
	return o.NotionalValue, true
}

// HasNotionalValue returns a boolean if a field has been set.
func (o *AccountInformationResponseResultPositionsInner) HasNotionalValue() bool {
	if o != nil && !common.IsNil(o.NotionalValue) {
		return true
	}

	return false
}

// SetNotionalValue gets a reference to the given string and assigns it to the NotionalValue field.
func (o *AccountInformationResponseResultPositionsInner) SetNotionalValue(v string) {
	o.NotionalValue = &v
}

// GetIsolatedWallet returns the IsolatedWallet field value if set, zero value otherwise.
func (o *AccountInformationResponseResultPositionsInner) GetIsolatedWallet() string {
	if o == nil || common.IsNil(o.IsolatedWallet) {
		var ret string
		return ret
	}
	return *o.IsolatedWallet
}

// GetIsolatedWalletOk returns a tuple with the IsolatedWallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResultPositionsInner) GetIsolatedWalletOk() (*string, bool) {
	if o == nil || common.IsNil(o.IsolatedWallet) {
		return nil, false
	}
	return o.IsolatedWallet, true
}

// HasIsolatedWallet returns a boolean if a field has been set.
func (o *AccountInformationResponseResultPositionsInner) HasIsolatedWallet() bool {
	if o != nil && !common.IsNil(o.IsolatedWallet) {
		return true
	}

	return false
}

// SetIsolatedWallet gets a reference to the given string and assigns it to the IsolatedWallet field.
func (o *AccountInformationResponseResultPositionsInner) SetIsolatedWallet(v string) {
	o.IsolatedWallet = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *AccountInformationResponseResultPositionsInner) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResultPositionsInner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *AccountInformationResponseResultPositionsInner) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *AccountInformationResponseResultPositionsInner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

// GetPositionAmt returns the PositionAmt field value if set, zero value otherwise.
func (o *AccountInformationResponseResultPositionsInner) GetPositionAmt() string {
	if o == nil || common.IsNil(o.PositionAmt) {
		var ret string
		return ret
	}
	return *o.PositionAmt
}

// GetPositionAmtOk returns a tuple with the PositionAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResultPositionsInner) GetPositionAmtOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionAmt) {
		return nil, false
	}
	return o.PositionAmt, true
}

// HasPositionAmt returns a boolean if a field has been set.
func (o *AccountInformationResponseResultPositionsInner) HasPositionAmt() bool {
	if o != nil && !common.IsNil(o.PositionAmt) {
		return true
	}

	return false
}

// SetPositionAmt gets a reference to the given string and assigns it to the PositionAmt field.
func (o *AccountInformationResponseResultPositionsInner) SetPositionAmt(v string) {
	o.PositionAmt = &v
}

// GetBreakEvenPrice returns the BreakEvenPrice field value if set, zero value otherwise.
func (o *AccountInformationResponseResultPositionsInner) GetBreakEvenPrice() string {
	if o == nil || common.IsNil(o.BreakEvenPrice) {
		var ret string
		return ret
	}
	return *o.BreakEvenPrice
}

// GetBreakEvenPriceOk returns a tuple with the BreakEvenPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponseResultPositionsInner) GetBreakEvenPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.BreakEvenPrice) {
		return nil, false
	}
	return o.BreakEvenPrice, true
}

// HasBreakEvenPrice returns a boolean if a field has been set.
func (o *AccountInformationResponseResultPositionsInner) HasBreakEvenPrice() bool {
	if o != nil && !common.IsNil(o.BreakEvenPrice) {
		return true
	}

	return false
}

// SetBreakEvenPrice gets a reference to the given string and assigns it to the BreakEvenPrice field.
func (o *AccountInformationResponseResultPositionsInner) SetBreakEvenPrice(v string) {
	o.BreakEvenPrice = &v
}

func (o AccountInformationResponseResultPositionsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountInformationResponseResultPositionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.InitialMargin) {
		toSerialize["initialMargin"] = o.InitialMargin
	}
	if !common.IsNil(o.MaintMargin) {
		toSerialize["maintMargin"] = o.MaintMargin
	}
	if !common.IsNil(o.UnrealizedProfit) {
		toSerialize["unrealizedProfit"] = o.UnrealizedProfit
	}
	if !common.IsNil(o.PositionInitialMargin) {
		toSerialize["positionInitialMargin"] = o.PositionInitialMargin
	}
	if !common.IsNil(o.OpenOrderInitialMargin) {
		toSerialize["openOrderInitialMargin"] = o.OpenOrderInitialMargin
	}
	if !common.IsNil(o.Leverage) {
		toSerialize["leverage"] = o.Leverage
	}
	if !common.IsNil(o.Isolated) {
		toSerialize["isolated"] = o.Isolated
	}
	if !common.IsNil(o.PositionSide) {
		toSerialize["positionSide"] = o.PositionSide
	}
	if !common.IsNil(o.EntryPrice) {
		toSerialize["entryPrice"] = o.EntryPrice
	}
	if !common.IsNil(o.MaxQty) {
		toSerialize["maxQty"] = o.MaxQty
	}
	if !common.IsNil(o.NotionalValue) {
		toSerialize["notionalValue"] = o.NotionalValue
	}
	if !common.IsNil(o.IsolatedWallet) {
		toSerialize["isolatedWallet"] = o.IsolatedWallet
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if !common.IsNil(o.PositionAmt) {
		toSerialize["positionAmt"] = o.PositionAmt
	}
	if !common.IsNil(o.BreakEvenPrice) {
		toSerialize["breakEvenPrice"] = o.BreakEvenPrice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountInformationResponseResultPositionsInner) UnmarshalJSON(data []byte) (err error) {
	varAccountInformationResponseResultPositionsInner := _AccountInformationResponseResultPositionsInner{}

	err = json.Unmarshal(data, &varAccountInformationResponseResultPositionsInner)

	if err != nil {
		return err
	}

	*o = AccountInformationResponseResultPositionsInner(varAccountInformationResponseResultPositionsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "initialMargin")
		delete(additionalProperties, "maintMargin")
		delete(additionalProperties, "unrealizedProfit")
		delete(additionalProperties, "positionInitialMargin")
		delete(additionalProperties, "openOrderInitialMargin")
		delete(additionalProperties, "leverage")
		delete(additionalProperties, "isolated")
		delete(additionalProperties, "positionSide")
		delete(additionalProperties, "entryPrice")
		delete(additionalProperties, "maxQty")
		delete(additionalProperties, "notionalValue")
		delete(additionalProperties, "isolatedWallet")
		delete(additionalProperties, "updateTime")
		delete(additionalProperties, "positionAmt")
		delete(additionalProperties, "breakEvenPrice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountInformationResponseResultPositionsInner struct {
	value *AccountInformationResponseResultPositionsInner
	isSet bool
}

func (v NullableAccountInformationResponseResultPositionsInner) Get() *AccountInformationResponseResultPositionsInner {
	return v.value
}

func (v *NullableAccountInformationResponseResultPositionsInner) Set(val *AccountInformationResponseResultPositionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountInformationResponseResultPositionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountInformationResponseResultPositionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountInformationResponseResultPositionsInner(val *AccountInformationResponseResultPositionsInner) *NullableAccountInformationResponseResultPositionsInner {
	return &NullableAccountInformationResponseResultPositionsInner{value: val, isSet: true}
}

func (v NullableAccountInformationResponseResultPositionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountInformationResponseResultPositionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
