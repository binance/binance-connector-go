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

// checks if the AccountInformationV2ResponsePositionsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountInformationV2ResponsePositionsInner{}

// AccountInformationV2ResponsePositionsInner struct for AccountInformationV2ResponsePositionsInner
type AccountInformationV2ResponsePositionsInner struct {
	Symbol                 *string `json:"symbol,omitempty"`
	InitialMargin          *string `json:"initialMargin,omitempty"`
	MaintMargin            *string `json:"maintMargin,omitempty"`
	UnrealizedProfit       *string `json:"unrealizedProfit,omitempty"`
	PositionInitialMargin  *string `json:"positionInitialMargin,omitempty"`
	OpenOrderInitialMargin *string `json:"openOrderInitialMargin,omitempty"`
	Leverage               *string `json:"leverage,omitempty"`
	Isolated               *bool   `json:"isolated,omitempty"`
	EntryPrice             *string `json:"entryPrice,omitempty"`
	MaxNotional            *string `json:"maxNotional,omitempty"`
	BidNotional            *string `json:"bidNotional,omitempty"`
	AskNotional            *string `json:"askNotional,omitempty"`
	PositionSide           *string `json:"positionSide,omitempty"`
	PositionAmt            *string `json:"positionAmt,omitempty"`
	UpdateTime             *int64  `json:"updateTime,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _AccountInformationV2ResponsePositionsInner AccountInformationV2ResponsePositionsInner

// NewAccountInformationV2ResponsePositionsInner instantiates a new AccountInformationV2ResponsePositionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountInformationV2ResponsePositionsInner() *AccountInformationV2ResponsePositionsInner {
	this := AccountInformationV2ResponsePositionsInner{}
	return &this
}

// NewAccountInformationV2ResponsePositionsInnerWithDefaults instantiates a new AccountInformationV2ResponsePositionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountInformationV2ResponsePositionsInnerWithDefaults() *AccountInformationV2ResponsePositionsInner {
	this := AccountInformationV2ResponsePositionsInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *AccountInformationV2ResponsePositionsInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2ResponsePositionsInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *AccountInformationV2ResponsePositionsInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *AccountInformationV2ResponsePositionsInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetInitialMargin returns the InitialMargin field value if set, zero value otherwise.
func (o *AccountInformationV2ResponsePositionsInner) GetInitialMargin() string {
	if o == nil || common.IsNil(o.InitialMargin) {
		var ret string
		return ret
	}
	return *o.InitialMargin
}

// GetInitialMarginOk returns a tuple with the InitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2ResponsePositionsInner) GetInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.InitialMargin) {
		return nil, false
	}
	return o.InitialMargin, true
}

// HasInitialMargin returns a boolean if a field has been set.
func (o *AccountInformationV2ResponsePositionsInner) HasInitialMargin() bool {
	if o != nil && !common.IsNil(o.InitialMargin) {
		return true
	}

	return false
}

// SetInitialMargin gets a reference to the given string and assigns it to the InitialMargin field.
func (o *AccountInformationV2ResponsePositionsInner) SetInitialMargin(v string) {
	o.InitialMargin = &v
}

// GetMaintMargin returns the MaintMargin field value if set, zero value otherwise.
func (o *AccountInformationV2ResponsePositionsInner) GetMaintMargin() string {
	if o == nil || common.IsNil(o.MaintMargin) {
		var ret string
		return ret
	}
	return *o.MaintMargin
}

// GetMaintMarginOk returns a tuple with the MaintMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2ResponsePositionsInner) GetMaintMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaintMargin) {
		return nil, false
	}
	return o.MaintMargin, true
}

// HasMaintMargin returns a boolean if a field has been set.
func (o *AccountInformationV2ResponsePositionsInner) HasMaintMargin() bool {
	if o != nil && !common.IsNil(o.MaintMargin) {
		return true
	}

	return false
}

// SetMaintMargin gets a reference to the given string and assigns it to the MaintMargin field.
func (o *AccountInformationV2ResponsePositionsInner) SetMaintMargin(v string) {
	o.MaintMargin = &v
}

// GetUnrealizedProfit returns the UnrealizedProfit field value if set, zero value otherwise.
func (o *AccountInformationV2ResponsePositionsInner) GetUnrealizedProfit() string {
	if o == nil || common.IsNil(o.UnrealizedProfit) {
		var ret string
		return ret
	}
	return *o.UnrealizedProfit
}

// GetUnrealizedProfitOk returns a tuple with the UnrealizedProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2ResponsePositionsInner) GetUnrealizedProfitOk() (*string, bool) {
	if o == nil || common.IsNil(o.UnrealizedProfit) {
		return nil, false
	}
	return o.UnrealizedProfit, true
}

// HasUnrealizedProfit returns a boolean if a field has been set.
func (o *AccountInformationV2ResponsePositionsInner) HasUnrealizedProfit() bool {
	if o != nil && !common.IsNil(o.UnrealizedProfit) {
		return true
	}

	return false
}

// SetUnrealizedProfit gets a reference to the given string and assigns it to the UnrealizedProfit field.
func (o *AccountInformationV2ResponsePositionsInner) SetUnrealizedProfit(v string) {
	o.UnrealizedProfit = &v
}

// GetPositionInitialMargin returns the PositionInitialMargin field value if set, zero value otherwise.
func (o *AccountInformationV2ResponsePositionsInner) GetPositionInitialMargin() string {
	if o == nil || common.IsNil(o.PositionInitialMargin) {
		var ret string
		return ret
	}
	return *o.PositionInitialMargin
}

// GetPositionInitialMarginOk returns a tuple with the PositionInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2ResponsePositionsInner) GetPositionInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionInitialMargin) {
		return nil, false
	}
	return o.PositionInitialMargin, true
}

// HasPositionInitialMargin returns a boolean if a field has been set.
func (o *AccountInformationV2ResponsePositionsInner) HasPositionInitialMargin() bool {
	if o != nil && !common.IsNil(o.PositionInitialMargin) {
		return true
	}

	return false
}

// SetPositionInitialMargin gets a reference to the given string and assigns it to the PositionInitialMargin field.
func (o *AccountInformationV2ResponsePositionsInner) SetPositionInitialMargin(v string) {
	o.PositionInitialMargin = &v
}

// GetOpenOrderInitialMargin returns the OpenOrderInitialMargin field value if set, zero value otherwise.
func (o *AccountInformationV2ResponsePositionsInner) GetOpenOrderInitialMargin() string {
	if o == nil || common.IsNil(o.OpenOrderInitialMargin) {
		var ret string
		return ret
	}
	return *o.OpenOrderInitialMargin
}

// GetOpenOrderInitialMarginOk returns a tuple with the OpenOrderInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2ResponsePositionsInner) GetOpenOrderInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.OpenOrderInitialMargin) {
		return nil, false
	}
	return o.OpenOrderInitialMargin, true
}

// HasOpenOrderInitialMargin returns a boolean if a field has been set.
func (o *AccountInformationV2ResponsePositionsInner) HasOpenOrderInitialMargin() bool {
	if o != nil && !common.IsNil(o.OpenOrderInitialMargin) {
		return true
	}

	return false
}

// SetOpenOrderInitialMargin gets a reference to the given string and assigns it to the OpenOrderInitialMargin field.
func (o *AccountInformationV2ResponsePositionsInner) SetOpenOrderInitialMargin(v string) {
	o.OpenOrderInitialMargin = &v
}

// GetLeverage returns the Leverage field value if set, zero value otherwise.
func (o *AccountInformationV2ResponsePositionsInner) GetLeverage() string {
	if o == nil || common.IsNil(o.Leverage) {
		var ret string
		return ret
	}
	return *o.Leverage
}

// GetLeverageOk returns a tuple with the Leverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2ResponsePositionsInner) GetLeverageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Leverage) {
		return nil, false
	}
	return o.Leverage, true
}

// HasLeverage returns a boolean if a field has been set.
func (o *AccountInformationV2ResponsePositionsInner) HasLeverage() bool {
	if o != nil && !common.IsNil(o.Leverage) {
		return true
	}

	return false
}

// SetLeverage gets a reference to the given string and assigns it to the Leverage field.
func (o *AccountInformationV2ResponsePositionsInner) SetLeverage(v string) {
	o.Leverage = &v
}

// GetIsolated returns the Isolated field value if set, zero value otherwise.
func (o *AccountInformationV2ResponsePositionsInner) GetIsolated() bool {
	if o == nil || common.IsNil(o.Isolated) {
		var ret bool
		return ret
	}
	return *o.Isolated
}

// GetIsolatedOk returns a tuple with the Isolated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2ResponsePositionsInner) GetIsolatedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Isolated) {
		return nil, false
	}
	return o.Isolated, true
}

// HasIsolated returns a boolean if a field has been set.
func (o *AccountInformationV2ResponsePositionsInner) HasIsolated() bool {
	if o != nil && !common.IsNil(o.Isolated) {
		return true
	}

	return false
}

// SetIsolated gets a reference to the given bool and assigns it to the Isolated field.
func (o *AccountInformationV2ResponsePositionsInner) SetIsolated(v bool) {
	o.Isolated = &v
}

// GetEntryPrice returns the EntryPrice field value if set, zero value otherwise.
func (o *AccountInformationV2ResponsePositionsInner) GetEntryPrice() string {
	if o == nil || common.IsNil(o.EntryPrice) {
		var ret string
		return ret
	}
	return *o.EntryPrice
}

// GetEntryPriceOk returns a tuple with the EntryPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2ResponsePositionsInner) GetEntryPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.EntryPrice) {
		return nil, false
	}
	return o.EntryPrice, true
}

// HasEntryPrice returns a boolean if a field has been set.
func (o *AccountInformationV2ResponsePositionsInner) HasEntryPrice() bool {
	if o != nil && !common.IsNil(o.EntryPrice) {
		return true
	}

	return false
}

// SetEntryPrice gets a reference to the given string and assigns it to the EntryPrice field.
func (o *AccountInformationV2ResponsePositionsInner) SetEntryPrice(v string) {
	o.EntryPrice = &v
}

// GetMaxNotional returns the MaxNotional field value if set, zero value otherwise.
func (o *AccountInformationV2ResponsePositionsInner) GetMaxNotional() string {
	if o == nil || common.IsNil(o.MaxNotional) {
		var ret string
		return ret
	}
	return *o.MaxNotional
}

// GetMaxNotionalOk returns a tuple with the MaxNotional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2ResponsePositionsInner) GetMaxNotionalOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxNotional) {
		return nil, false
	}
	return o.MaxNotional, true
}

// HasMaxNotional returns a boolean if a field has been set.
func (o *AccountInformationV2ResponsePositionsInner) HasMaxNotional() bool {
	if o != nil && !common.IsNil(o.MaxNotional) {
		return true
	}

	return false
}

// SetMaxNotional gets a reference to the given string and assigns it to the MaxNotional field.
func (o *AccountInformationV2ResponsePositionsInner) SetMaxNotional(v string) {
	o.MaxNotional = &v
}

// GetBidNotional returns the BidNotional field value if set, zero value otherwise.
func (o *AccountInformationV2ResponsePositionsInner) GetBidNotional() string {
	if o == nil || common.IsNil(o.BidNotional) {
		var ret string
		return ret
	}
	return *o.BidNotional
}

// GetBidNotionalOk returns a tuple with the BidNotional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2ResponsePositionsInner) GetBidNotionalOk() (*string, bool) {
	if o == nil || common.IsNil(o.BidNotional) {
		return nil, false
	}
	return o.BidNotional, true
}

// HasBidNotional returns a boolean if a field has been set.
func (o *AccountInformationV2ResponsePositionsInner) HasBidNotional() bool {
	if o != nil && !common.IsNil(o.BidNotional) {
		return true
	}

	return false
}

// SetBidNotional gets a reference to the given string and assigns it to the BidNotional field.
func (o *AccountInformationV2ResponsePositionsInner) SetBidNotional(v string) {
	o.BidNotional = &v
}

// GetAskNotional returns the AskNotional field value if set, zero value otherwise.
func (o *AccountInformationV2ResponsePositionsInner) GetAskNotional() string {
	if o == nil || common.IsNil(o.AskNotional) {
		var ret string
		return ret
	}
	return *o.AskNotional
}

// GetAskNotionalOk returns a tuple with the AskNotional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2ResponsePositionsInner) GetAskNotionalOk() (*string, bool) {
	if o == nil || common.IsNil(o.AskNotional) {
		return nil, false
	}
	return o.AskNotional, true
}

// HasAskNotional returns a boolean if a field has been set.
func (o *AccountInformationV2ResponsePositionsInner) HasAskNotional() bool {
	if o != nil && !common.IsNil(o.AskNotional) {
		return true
	}

	return false
}

// SetAskNotional gets a reference to the given string and assigns it to the AskNotional field.
func (o *AccountInformationV2ResponsePositionsInner) SetAskNotional(v string) {
	o.AskNotional = &v
}

// GetPositionSide returns the PositionSide field value if set, zero value otherwise.
func (o *AccountInformationV2ResponsePositionsInner) GetPositionSide() string {
	if o == nil || common.IsNil(o.PositionSide) {
		var ret string
		return ret
	}
	return *o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2ResponsePositionsInner) GetPositionSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionSide) {
		return nil, false
	}
	return o.PositionSide, true
}

// HasPositionSide returns a boolean if a field has been set.
func (o *AccountInformationV2ResponsePositionsInner) HasPositionSide() bool {
	if o != nil && !common.IsNil(o.PositionSide) {
		return true
	}

	return false
}

// SetPositionSide gets a reference to the given string and assigns it to the PositionSide field.
func (o *AccountInformationV2ResponsePositionsInner) SetPositionSide(v string) {
	o.PositionSide = &v
}

// GetPositionAmt returns the PositionAmt field value if set, zero value otherwise.
func (o *AccountInformationV2ResponsePositionsInner) GetPositionAmt() string {
	if o == nil || common.IsNil(o.PositionAmt) {
		var ret string
		return ret
	}
	return *o.PositionAmt
}

// GetPositionAmtOk returns a tuple with the PositionAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2ResponsePositionsInner) GetPositionAmtOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionAmt) {
		return nil, false
	}
	return o.PositionAmt, true
}

// HasPositionAmt returns a boolean if a field has been set.
func (o *AccountInformationV2ResponsePositionsInner) HasPositionAmt() bool {
	if o != nil && !common.IsNil(o.PositionAmt) {
		return true
	}

	return false
}

// SetPositionAmt gets a reference to the given string and assigns it to the PositionAmt field.
func (o *AccountInformationV2ResponsePositionsInner) SetPositionAmt(v string) {
	o.PositionAmt = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *AccountInformationV2ResponsePositionsInner) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInformationV2ResponsePositionsInner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *AccountInformationV2ResponsePositionsInner) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *AccountInformationV2ResponsePositionsInner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o AccountInformationV2ResponsePositionsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountInformationV2ResponsePositionsInner) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.EntryPrice) {
		toSerialize["entryPrice"] = o.EntryPrice
	}
	if !common.IsNil(o.MaxNotional) {
		toSerialize["maxNotional"] = o.MaxNotional
	}
	if !common.IsNil(o.BidNotional) {
		toSerialize["bidNotional"] = o.BidNotional
	}
	if !common.IsNil(o.AskNotional) {
		toSerialize["askNotional"] = o.AskNotional
	}
	if !common.IsNil(o.PositionSide) {
		toSerialize["positionSide"] = o.PositionSide
	}
	if !common.IsNil(o.PositionAmt) {
		toSerialize["positionAmt"] = o.PositionAmt
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountInformationV2ResponsePositionsInner) UnmarshalJSON(data []byte) (err error) {
	varAccountInformationV2ResponsePositionsInner := _AccountInformationV2ResponsePositionsInner{}

	err = json.Unmarshal(data, &varAccountInformationV2ResponsePositionsInner)

	if err != nil {
		return err
	}

	*o = AccountInformationV2ResponsePositionsInner(varAccountInformationV2ResponsePositionsInner)

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
		delete(additionalProperties, "entryPrice")
		delete(additionalProperties, "maxNotional")
		delete(additionalProperties, "bidNotional")
		delete(additionalProperties, "askNotional")
		delete(additionalProperties, "positionSide")
		delete(additionalProperties, "positionAmt")
		delete(additionalProperties, "updateTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountInformationV2ResponsePositionsInner struct {
	value *AccountInformationV2ResponsePositionsInner
	isSet bool
}

func (v NullableAccountInformationV2ResponsePositionsInner) Get() *AccountInformationV2ResponsePositionsInner {
	return v.value
}

func (v *NullableAccountInformationV2ResponsePositionsInner) Set(val *AccountInformationV2ResponsePositionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountInformationV2ResponsePositionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountInformationV2ResponsePositionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountInformationV2ResponsePositionsInner(val *AccountInformationV2ResponsePositionsInner) *NullableAccountInformationV2ResponsePositionsInner {
	return &NullableAccountInformationV2ResponsePositionsInner{value: val, isSet: true}
}

func (v NullableAccountInformationV2ResponsePositionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountInformationV2ResponsePositionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
