/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AccountInformationResponsePositionsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountInformationResponsePositionsInner{}

// AccountInformationResponsePositionsInner struct for AccountInformationResponsePositionsInner
type AccountInformationResponsePositionsInner struct {
	Symbol                 *string `json:"symbol,omitempty"`
	PositionAmt            *string `json:"positionAmt,omitempty"`
	InitialMargin          *string `json:"initialMargin,omitempty"`
	MaintMargin            *string `json:"maintMargin,omitempty"`
	UnrealizedProfit       *string `json:"unrealizedProfit,omitempty"`
	PositionInitialMargin  *string `json:"positionInitialMargin,omitempty"`
	OpenOrderInitialMargin *string `json:"openOrderInitialMargin,omitempty"`
	Leverage               *string `json:"leverage,omitempty"`
	Isolated               *bool   `json:"isolated,omitempty"`
	PositionSide           *string `json:"positionSide,omitempty"`
	EntryPrice             *string `json:"entryPrice,omitempty"`
	BreakEvenPrice         *string `json:"breakEvenPrice,omitempty"`
	MaxQty                 *string `json:"maxQty,omitempty"`
	UpdateTime             *int64  `json:"updateTime,omitempty"`
	NotionalValue          *string `json:"notionalValue,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _AccountInformationResponsePositionsInner AccountInformationResponsePositionsInner

// NewAccountInformationResponsePositionsInner instantiates a new AccountInformationResponsePositionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountInformationResponsePositionsInner() *AccountInformationResponsePositionsInner {
	this := AccountInformationResponsePositionsInner{}
	return &this
}

// NewAccountInformationResponsePositionsInnerWithDefaults instantiates a new AccountInformationResponsePositionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountInformationResponsePositionsInnerWithDefaults() *AccountInformationResponsePositionsInner {
	this := AccountInformationResponsePositionsInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *AccountInformationResponsePositionsInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponsePositionsInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *AccountInformationResponsePositionsInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *AccountInformationResponsePositionsInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetPositionAmt returns the PositionAmt field value if set, zero value otherwise.
func (o *AccountInformationResponsePositionsInner) GetPositionAmt() string {
	if o == nil || common.IsNil(o.PositionAmt) {
		var ret string
		return ret
	}
	return *o.PositionAmt
}

// GetPositionAmtOk returns a tuple with the PositionAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponsePositionsInner) GetPositionAmtOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionAmt) {
		return nil, false
	}
	return o.PositionAmt, true
}

// HasPositionAmt returns a boolean if a field has been set.
func (o *AccountInformationResponsePositionsInner) HasPositionAmt() bool {
	if o != nil && !common.IsNil(o.PositionAmt) {
		return true
	}

	return false
}

// SetPositionAmt gets a reference to the given string and assigns it to the PositionAmt field.
func (o *AccountInformationResponsePositionsInner) SetPositionAmt(v string) {
	o.PositionAmt = &v
}

// GetInitialMargin returns the InitialMargin field value if set, zero value otherwise.
func (o *AccountInformationResponsePositionsInner) GetInitialMargin() string {
	if o == nil || common.IsNil(o.InitialMargin) {
		var ret string
		return ret
	}
	return *o.InitialMargin
}

// GetInitialMarginOk returns a tuple with the InitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponsePositionsInner) GetInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.InitialMargin) {
		return nil, false
	}
	return o.InitialMargin, true
}

// HasInitialMargin returns a boolean if a field has been set.
func (o *AccountInformationResponsePositionsInner) HasInitialMargin() bool {
	if o != nil && !common.IsNil(o.InitialMargin) {
		return true
	}

	return false
}

// SetInitialMargin gets a reference to the given string and assigns it to the InitialMargin field.
func (o *AccountInformationResponsePositionsInner) SetInitialMargin(v string) {
	o.InitialMargin = &v
}

// GetMaintMargin returns the MaintMargin field value if set, zero value otherwise.
func (o *AccountInformationResponsePositionsInner) GetMaintMargin() string {
	if o == nil || common.IsNil(o.MaintMargin) {
		var ret string
		return ret
	}
	return *o.MaintMargin
}

// GetMaintMarginOk returns a tuple with the MaintMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponsePositionsInner) GetMaintMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaintMargin) {
		return nil, false
	}
	return o.MaintMargin, true
}

// HasMaintMargin returns a boolean if a field has been set.
func (o *AccountInformationResponsePositionsInner) HasMaintMargin() bool {
	if o != nil && !common.IsNil(o.MaintMargin) {
		return true
	}

	return false
}

// SetMaintMargin gets a reference to the given string and assigns it to the MaintMargin field.
func (o *AccountInformationResponsePositionsInner) SetMaintMargin(v string) {
	o.MaintMargin = &v
}

// GetUnrealizedProfit returns the UnrealizedProfit field value if set, zero value otherwise.
func (o *AccountInformationResponsePositionsInner) GetUnrealizedProfit() string {
	if o == nil || common.IsNil(o.UnrealizedProfit) {
		var ret string
		return ret
	}
	return *o.UnrealizedProfit
}

// GetUnrealizedProfitOk returns a tuple with the UnrealizedProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponsePositionsInner) GetUnrealizedProfitOk() (*string, bool) {
	if o == nil || common.IsNil(o.UnrealizedProfit) {
		return nil, false
	}
	return o.UnrealizedProfit, true
}

// HasUnrealizedProfit returns a boolean if a field has been set.
func (o *AccountInformationResponsePositionsInner) HasUnrealizedProfit() bool {
	if o != nil && !common.IsNil(o.UnrealizedProfit) {
		return true
	}

	return false
}

// SetUnrealizedProfit gets a reference to the given string and assigns it to the UnrealizedProfit field.
func (o *AccountInformationResponsePositionsInner) SetUnrealizedProfit(v string) {
	o.UnrealizedProfit = &v
}

// GetPositionInitialMargin returns the PositionInitialMargin field value if set, zero value otherwise.
func (o *AccountInformationResponsePositionsInner) GetPositionInitialMargin() string {
	if o == nil || common.IsNil(o.PositionInitialMargin) {
		var ret string
		return ret
	}
	return *o.PositionInitialMargin
}

// GetPositionInitialMarginOk returns a tuple with the PositionInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponsePositionsInner) GetPositionInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionInitialMargin) {
		return nil, false
	}
	return o.PositionInitialMargin, true
}

// HasPositionInitialMargin returns a boolean if a field has been set.
func (o *AccountInformationResponsePositionsInner) HasPositionInitialMargin() bool {
	if o != nil && !common.IsNil(o.PositionInitialMargin) {
		return true
	}

	return false
}

// SetPositionInitialMargin gets a reference to the given string and assigns it to the PositionInitialMargin field.
func (o *AccountInformationResponsePositionsInner) SetPositionInitialMargin(v string) {
	o.PositionInitialMargin = &v
}

// GetOpenOrderInitialMargin returns the OpenOrderInitialMargin field value if set, zero value otherwise.
func (o *AccountInformationResponsePositionsInner) GetOpenOrderInitialMargin() string {
	if o == nil || common.IsNil(o.OpenOrderInitialMargin) {
		var ret string
		return ret
	}
	return *o.OpenOrderInitialMargin
}

// GetOpenOrderInitialMarginOk returns a tuple with the OpenOrderInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponsePositionsInner) GetOpenOrderInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.OpenOrderInitialMargin) {
		return nil, false
	}
	return o.OpenOrderInitialMargin, true
}

// HasOpenOrderInitialMargin returns a boolean if a field has been set.
func (o *AccountInformationResponsePositionsInner) HasOpenOrderInitialMargin() bool {
	if o != nil && !common.IsNil(o.OpenOrderInitialMargin) {
		return true
	}

	return false
}

// SetOpenOrderInitialMargin gets a reference to the given string and assigns it to the OpenOrderInitialMargin field.
func (o *AccountInformationResponsePositionsInner) SetOpenOrderInitialMargin(v string) {
	o.OpenOrderInitialMargin = &v
}

// GetLeverage returns the Leverage field value if set, zero value otherwise.
func (o *AccountInformationResponsePositionsInner) GetLeverage() string {
	if o == nil || common.IsNil(o.Leverage) {
		var ret string
		return ret
	}
	return *o.Leverage
}

// GetLeverageOk returns a tuple with the Leverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponsePositionsInner) GetLeverageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Leverage) {
		return nil, false
	}
	return o.Leverage, true
}

// HasLeverage returns a boolean if a field has been set.
func (o *AccountInformationResponsePositionsInner) HasLeverage() bool {
	if o != nil && !common.IsNil(o.Leverage) {
		return true
	}

	return false
}

// SetLeverage gets a reference to the given string and assigns it to the Leverage field.
func (o *AccountInformationResponsePositionsInner) SetLeverage(v string) {
	o.Leverage = &v
}

// GetIsolated returns the Isolated field value if set, zero value otherwise.
func (o *AccountInformationResponsePositionsInner) GetIsolated() bool {
	if o == nil || common.IsNil(o.Isolated) {
		var ret bool
		return ret
	}
	return *o.Isolated
}

// GetIsolatedOk returns a tuple with the Isolated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponsePositionsInner) GetIsolatedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Isolated) {
		return nil, false
	}
	return o.Isolated, true
}

// HasIsolated returns a boolean if a field has been set.
func (o *AccountInformationResponsePositionsInner) HasIsolated() bool {
	if o != nil && !common.IsNil(o.Isolated) {
		return true
	}

	return false
}

// SetIsolated gets a reference to the given bool and assigns it to the Isolated field.
func (o *AccountInformationResponsePositionsInner) SetIsolated(v bool) {
	o.Isolated = &v
}

// GetPositionSide returns the PositionSide field value if set, zero value otherwise.
func (o *AccountInformationResponsePositionsInner) GetPositionSide() string {
	if o == nil || common.IsNil(o.PositionSide) {
		var ret string
		return ret
	}
	return *o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponsePositionsInner) GetPositionSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionSide) {
		return nil, false
	}
	return o.PositionSide, true
}

// HasPositionSide returns a boolean if a field has been set.
func (o *AccountInformationResponsePositionsInner) HasPositionSide() bool {
	if o != nil && !common.IsNil(o.PositionSide) {
		return true
	}

	return false
}

// SetPositionSide gets a reference to the given string and assigns it to the PositionSide field.
func (o *AccountInformationResponsePositionsInner) SetPositionSide(v string) {
	o.PositionSide = &v
}

// GetEntryPrice returns the EntryPrice field value if set, zero value otherwise.
func (o *AccountInformationResponsePositionsInner) GetEntryPrice() string {
	if o == nil || common.IsNil(o.EntryPrice) {
		var ret string
		return ret
	}
	return *o.EntryPrice
}

// GetEntryPriceOk returns a tuple with the EntryPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponsePositionsInner) GetEntryPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.EntryPrice) {
		return nil, false
	}
	return o.EntryPrice, true
}

// HasEntryPrice returns a boolean if a field has been set.
func (o *AccountInformationResponsePositionsInner) HasEntryPrice() bool {
	if o != nil && !common.IsNil(o.EntryPrice) {
		return true
	}

	return false
}

// SetEntryPrice gets a reference to the given string and assigns it to the EntryPrice field.
func (o *AccountInformationResponsePositionsInner) SetEntryPrice(v string) {
	o.EntryPrice = &v
}

// GetBreakEvenPrice returns the BreakEvenPrice field value if set, zero value otherwise.
func (o *AccountInformationResponsePositionsInner) GetBreakEvenPrice() string {
	if o == nil || common.IsNil(o.BreakEvenPrice) {
		var ret string
		return ret
	}
	return *o.BreakEvenPrice
}

// GetBreakEvenPriceOk returns a tuple with the BreakEvenPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponsePositionsInner) GetBreakEvenPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.BreakEvenPrice) {
		return nil, false
	}
	return o.BreakEvenPrice, true
}

// HasBreakEvenPrice returns a boolean if a field has been set.
func (o *AccountInformationResponsePositionsInner) HasBreakEvenPrice() bool {
	if o != nil && !common.IsNil(o.BreakEvenPrice) {
		return true
	}

	return false
}

// SetBreakEvenPrice gets a reference to the given string and assigns it to the BreakEvenPrice field.
func (o *AccountInformationResponsePositionsInner) SetBreakEvenPrice(v string) {
	o.BreakEvenPrice = &v
}

// GetMaxQty returns the MaxQty field value if set, zero value otherwise.
func (o *AccountInformationResponsePositionsInner) GetMaxQty() string {
	if o == nil || common.IsNil(o.MaxQty) {
		var ret string
		return ret
	}
	return *o.MaxQty
}

// GetMaxQtyOk returns a tuple with the MaxQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponsePositionsInner) GetMaxQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxQty) {
		return nil, false
	}
	return o.MaxQty, true
}

// HasMaxQty returns a boolean if a field has been set.
func (o *AccountInformationResponsePositionsInner) HasMaxQty() bool {
	if o != nil && !common.IsNil(o.MaxQty) {
		return true
	}

	return false
}

// SetMaxQty gets a reference to the given string and assigns it to the MaxQty field.
func (o *AccountInformationResponsePositionsInner) SetMaxQty(v string) {
	o.MaxQty = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *AccountInformationResponsePositionsInner) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponsePositionsInner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *AccountInformationResponsePositionsInner) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *AccountInformationResponsePositionsInner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

// GetNotionalValue returns the NotionalValue field value if set, zero value otherwise.
func (o *AccountInformationResponsePositionsInner) GetNotionalValue() string {
	if o == nil || common.IsNil(o.NotionalValue) {
		var ret string
		return ret
	}
	return *o.NotionalValue
}

// GetNotionalValueOk returns a tuple with the NotionalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationResponsePositionsInner) GetNotionalValueOk() (*string, bool) {
	if o == nil || common.IsNil(o.NotionalValue) {
		return nil, false
	}
	return o.NotionalValue, true
}

// HasNotionalValue returns a boolean if a field has been set.
func (o *AccountInformationResponsePositionsInner) HasNotionalValue() bool {
	if o != nil && !common.IsNil(o.NotionalValue) {
		return true
	}

	return false
}

// SetNotionalValue gets a reference to the given string and assigns it to the NotionalValue field.
func (o *AccountInformationResponsePositionsInner) SetNotionalValue(v string) {
	o.NotionalValue = &v
}

func (o AccountInformationResponsePositionsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountInformationResponsePositionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.PositionAmt) {
		toSerialize["positionAmt"] = o.PositionAmt
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
	if !common.IsNil(o.BreakEvenPrice) {
		toSerialize["breakEvenPrice"] = o.BreakEvenPrice
	}
	if !common.IsNil(o.MaxQty) {
		toSerialize["maxQty"] = o.MaxQty
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if !common.IsNil(o.NotionalValue) {
		toSerialize["notionalValue"] = o.NotionalValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountInformationResponsePositionsInner) UnmarshalJSON(data []byte) (err error) {
	varAccountInformationResponsePositionsInner := _AccountInformationResponsePositionsInner{}

	err = json.Unmarshal(data, &varAccountInformationResponsePositionsInner)

	if err != nil {
		return err
	}

	*o = AccountInformationResponsePositionsInner(varAccountInformationResponsePositionsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "positionAmt")
		delete(additionalProperties, "initialMargin")
		delete(additionalProperties, "maintMargin")
		delete(additionalProperties, "unrealizedProfit")
		delete(additionalProperties, "positionInitialMargin")
		delete(additionalProperties, "openOrderInitialMargin")
		delete(additionalProperties, "leverage")
		delete(additionalProperties, "isolated")
		delete(additionalProperties, "positionSide")
		delete(additionalProperties, "entryPrice")
		delete(additionalProperties, "breakEvenPrice")
		delete(additionalProperties, "maxQty")
		delete(additionalProperties, "updateTime")
		delete(additionalProperties, "notionalValue")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountInformationResponsePositionsInner struct {
	value *AccountInformationResponsePositionsInner
	isSet bool
}

func (v NullableAccountInformationResponsePositionsInner) Get() *AccountInformationResponsePositionsInner {
	return v.value
}

func (v *NullableAccountInformationResponsePositionsInner) Set(val *AccountInformationResponsePositionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountInformationResponsePositionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountInformationResponsePositionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountInformationResponsePositionsInner(val *AccountInformationResponsePositionsInner) *NullableAccountInformationResponsePositionsInner {
	return &NullableAccountInformationResponsePositionsInner{value: val, isSet: true}
}

func (v NullableAccountInformationResponsePositionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountInformationResponsePositionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
