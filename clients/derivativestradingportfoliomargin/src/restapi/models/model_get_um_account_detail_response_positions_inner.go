/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetUmAccountDetailResponsePositionsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetUmAccountDetailResponsePositionsInner{}

// GetUmAccountDetailResponsePositionsInner struct for GetUmAccountDetailResponsePositionsInner
type GetUmAccountDetailResponsePositionsInner struct {
	Symbol                 *string `json:"symbol,omitempty"`
	InitialMargin          *string `json:"initialMargin,omitempty"`
	MaintMargin            *string `json:"maintMargin,omitempty"`
	UnrealizedProfit       *string `json:"unrealizedProfit,omitempty"`
	PositionInitialMargin  *string `json:"positionInitialMargin,omitempty"`
	OpenOrderInitialMargin *string `json:"openOrderInitialMargin,omitempty"`
	Leverage               *string `json:"leverage,omitempty"`
	EntryPrice             *string `json:"entryPrice,omitempty"`
	MaxNotional            *string `json:"maxNotional,omitempty"`
	BidNotional            *string `json:"bidNotional,omitempty"`
	AskNotional            *string `json:"askNotional,omitempty"`
	PositionSide           *string `json:"positionSide,omitempty"`
	PositionAmt            *string `json:"positionAmt,omitempty"`
	UpdateTime             *int64  `json:"updateTime,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _GetUmAccountDetailResponsePositionsInner GetUmAccountDetailResponsePositionsInner

// NewGetUmAccountDetailResponsePositionsInner instantiates a new GetUmAccountDetailResponsePositionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUmAccountDetailResponsePositionsInner() *GetUmAccountDetailResponsePositionsInner {
	this := GetUmAccountDetailResponsePositionsInner{}
	return &this
}

// NewGetUmAccountDetailResponsePositionsInnerWithDefaults instantiates a new GetUmAccountDetailResponsePositionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUmAccountDetailResponsePositionsInnerWithDefaults() *GetUmAccountDetailResponsePositionsInner {
	this := GetUmAccountDetailResponsePositionsInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetUmAccountDetailResponsePositionsInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailResponsePositionsInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetUmAccountDetailResponsePositionsInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetUmAccountDetailResponsePositionsInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetInitialMargin returns the InitialMargin field value if set, zero value otherwise.
func (o *GetUmAccountDetailResponsePositionsInner) GetInitialMargin() string {
	if o == nil || common.IsNil(o.InitialMargin) {
		var ret string
		return ret
	}
	return *o.InitialMargin
}

// GetInitialMarginOk returns a tuple with the InitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailResponsePositionsInner) GetInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.InitialMargin) {
		return nil, false
	}
	return o.InitialMargin, true
}

// HasInitialMargin returns a boolean if a field has been set.
func (o *GetUmAccountDetailResponsePositionsInner) HasInitialMargin() bool {
	if o != nil && !common.IsNil(o.InitialMargin) {
		return true
	}

	return false
}

// SetInitialMargin gets a reference to the given string and assigns it to the InitialMargin field.
func (o *GetUmAccountDetailResponsePositionsInner) SetInitialMargin(v string) {
	o.InitialMargin = &v
}

// GetMaintMargin returns the MaintMargin field value if set, zero value otherwise.
func (o *GetUmAccountDetailResponsePositionsInner) GetMaintMargin() string {
	if o == nil || common.IsNil(o.MaintMargin) {
		var ret string
		return ret
	}
	return *o.MaintMargin
}

// GetMaintMarginOk returns a tuple with the MaintMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailResponsePositionsInner) GetMaintMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaintMargin) {
		return nil, false
	}
	return o.MaintMargin, true
}

// HasMaintMargin returns a boolean if a field has been set.
func (o *GetUmAccountDetailResponsePositionsInner) HasMaintMargin() bool {
	if o != nil && !common.IsNil(o.MaintMargin) {
		return true
	}

	return false
}

// SetMaintMargin gets a reference to the given string and assigns it to the MaintMargin field.
func (o *GetUmAccountDetailResponsePositionsInner) SetMaintMargin(v string) {
	o.MaintMargin = &v
}

// GetUnrealizedProfit returns the UnrealizedProfit field value if set, zero value otherwise.
func (o *GetUmAccountDetailResponsePositionsInner) GetUnrealizedProfit() string {
	if o == nil || common.IsNil(o.UnrealizedProfit) {
		var ret string
		return ret
	}
	return *o.UnrealizedProfit
}

// GetUnrealizedProfitOk returns a tuple with the UnrealizedProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailResponsePositionsInner) GetUnrealizedProfitOk() (*string, bool) {
	if o == nil || common.IsNil(o.UnrealizedProfit) {
		return nil, false
	}
	return o.UnrealizedProfit, true
}

// HasUnrealizedProfit returns a boolean if a field has been set.
func (o *GetUmAccountDetailResponsePositionsInner) HasUnrealizedProfit() bool {
	if o != nil && !common.IsNil(o.UnrealizedProfit) {
		return true
	}

	return false
}

// SetUnrealizedProfit gets a reference to the given string and assigns it to the UnrealizedProfit field.
func (o *GetUmAccountDetailResponsePositionsInner) SetUnrealizedProfit(v string) {
	o.UnrealizedProfit = &v
}

// GetPositionInitialMargin returns the PositionInitialMargin field value if set, zero value otherwise.
func (o *GetUmAccountDetailResponsePositionsInner) GetPositionInitialMargin() string {
	if o == nil || common.IsNil(o.PositionInitialMargin) {
		var ret string
		return ret
	}
	return *o.PositionInitialMargin
}

// GetPositionInitialMarginOk returns a tuple with the PositionInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailResponsePositionsInner) GetPositionInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionInitialMargin) {
		return nil, false
	}
	return o.PositionInitialMargin, true
}

// HasPositionInitialMargin returns a boolean if a field has been set.
func (o *GetUmAccountDetailResponsePositionsInner) HasPositionInitialMargin() bool {
	if o != nil && !common.IsNil(o.PositionInitialMargin) {
		return true
	}

	return false
}

// SetPositionInitialMargin gets a reference to the given string and assigns it to the PositionInitialMargin field.
func (o *GetUmAccountDetailResponsePositionsInner) SetPositionInitialMargin(v string) {
	o.PositionInitialMargin = &v
}

// GetOpenOrderInitialMargin returns the OpenOrderInitialMargin field value if set, zero value otherwise.
func (o *GetUmAccountDetailResponsePositionsInner) GetOpenOrderInitialMargin() string {
	if o == nil || common.IsNil(o.OpenOrderInitialMargin) {
		var ret string
		return ret
	}
	return *o.OpenOrderInitialMargin
}

// GetOpenOrderInitialMarginOk returns a tuple with the OpenOrderInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailResponsePositionsInner) GetOpenOrderInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.OpenOrderInitialMargin) {
		return nil, false
	}
	return o.OpenOrderInitialMargin, true
}

// HasOpenOrderInitialMargin returns a boolean if a field has been set.
func (o *GetUmAccountDetailResponsePositionsInner) HasOpenOrderInitialMargin() bool {
	if o != nil && !common.IsNil(o.OpenOrderInitialMargin) {
		return true
	}

	return false
}

// SetOpenOrderInitialMargin gets a reference to the given string and assigns it to the OpenOrderInitialMargin field.
func (o *GetUmAccountDetailResponsePositionsInner) SetOpenOrderInitialMargin(v string) {
	o.OpenOrderInitialMargin = &v
}

// GetLeverage returns the Leverage field value if set, zero value otherwise.
func (o *GetUmAccountDetailResponsePositionsInner) GetLeverage() string {
	if o == nil || common.IsNil(o.Leverage) {
		var ret string
		return ret
	}
	return *o.Leverage
}

// GetLeverageOk returns a tuple with the Leverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailResponsePositionsInner) GetLeverageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Leverage) {
		return nil, false
	}
	return o.Leverage, true
}

// HasLeverage returns a boolean if a field has been set.
func (o *GetUmAccountDetailResponsePositionsInner) HasLeverage() bool {
	if o != nil && !common.IsNil(o.Leverage) {
		return true
	}

	return false
}

// SetLeverage gets a reference to the given string and assigns it to the Leverage field.
func (o *GetUmAccountDetailResponsePositionsInner) SetLeverage(v string) {
	o.Leverage = &v
}

// GetEntryPrice returns the EntryPrice field value if set, zero value otherwise.
func (o *GetUmAccountDetailResponsePositionsInner) GetEntryPrice() string {
	if o == nil || common.IsNil(o.EntryPrice) {
		var ret string
		return ret
	}
	return *o.EntryPrice
}

// GetEntryPriceOk returns a tuple with the EntryPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailResponsePositionsInner) GetEntryPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.EntryPrice) {
		return nil, false
	}
	return o.EntryPrice, true
}

// HasEntryPrice returns a boolean if a field has been set.
func (o *GetUmAccountDetailResponsePositionsInner) HasEntryPrice() bool {
	if o != nil && !common.IsNil(o.EntryPrice) {
		return true
	}

	return false
}

// SetEntryPrice gets a reference to the given string and assigns it to the EntryPrice field.
func (o *GetUmAccountDetailResponsePositionsInner) SetEntryPrice(v string) {
	o.EntryPrice = &v
}

// GetMaxNotional returns the MaxNotional field value if set, zero value otherwise.
func (o *GetUmAccountDetailResponsePositionsInner) GetMaxNotional() string {
	if o == nil || common.IsNil(o.MaxNotional) {
		var ret string
		return ret
	}
	return *o.MaxNotional
}

// GetMaxNotionalOk returns a tuple with the MaxNotional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailResponsePositionsInner) GetMaxNotionalOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxNotional) {
		return nil, false
	}
	return o.MaxNotional, true
}

// HasMaxNotional returns a boolean if a field has been set.
func (o *GetUmAccountDetailResponsePositionsInner) HasMaxNotional() bool {
	if o != nil && !common.IsNil(o.MaxNotional) {
		return true
	}

	return false
}

// SetMaxNotional gets a reference to the given string and assigns it to the MaxNotional field.
func (o *GetUmAccountDetailResponsePositionsInner) SetMaxNotional(v string) {
	o.MaxNotional = &v
}

// GetBidNotional returns the BidNotional field value if set, zero value otherwise.
func (o *GetUmAccountDetailResponsePositionsInner) GetBidNotional() string {
	if o == nil || common.IsNil(o.BidNotional) {
		var ret string
		return ret
	}
	return *o.BidNotional
}

// GetBidNotionalOk returns a tuple with the BidNotional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailResponsePositionsInner) GetBidNotionalOk() (*string, bool) {
	if o == nil || common.IsNil(o.BidNotional) {
		return nil, false
	}
	return o.BidNotional, true
}

// HasBidNotional returns a boolean if a field has been set.
func (o *GetUmAccountDetailResponsePositionsInner) HasBidNotional() bool {
	if o != nil && !common.IsNil(o.BidNotional) {
		return true
	}

	return false
}

// SetBidNotional gets a reference to the given string and assigns it to the BidNotional field.
func (o *GetUmAccountDetailResponsePositionsInner) SetBidNotional(v string) {
	o.BidNotional = &v
}

// GetAskNotional returns the AskNotional field value if set, zero value otherwise.
func (o *GetUmAccountDetailResponsePositionsInner) GetAskNotional() string {
	if o == nil || common.IsNil(o.AskNotional) {
		var ret string
		return ret
	}
	return *o.AskNotional
}

// GetAskNotionalOk returns a tuple with the AskNotional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailResponsePositionsInner) GetAskNotionalOk() (*string, bool) {
	if o == nil || common.IsNil(o.AskNotional) {
		return nil, false
	}
	return o.AskNotional, true
}

// HasAskNotional returns a boolean if a field has been set.
func (o *GetUmAccountDetailResponsePositionsInner) HasAskNotional() bool {
	if o != nil && !common.IsNil(o.AskNotional) {
		return true
	}

	return false
}

// SetAskNotional gets a reference to the given string and assigns it to the AskNotional field.
func (o *GetUmAccountDetailResponsePositionsInner) SetAskNotional(v string) {
	o.AskNotional = &v
}

// GetPositionSide returns the PositionSide field value if set, zero value otherwise.
func (o *GetUmAccountDetailResponsePositionsInner) GetPositionSide() string {
	if o == nil || common.IsNil(o.PositionSide) {
		var ret string
		return ret
	}
	return *o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailResponsePositionsInner) GetPositionSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionSide) {
		return nil, false
	}
	return o.PositionSide, true
}

// HasPositionSide returns a boolean if a field has been set.
func (o *GetUmAccountDetailResponsePositionsInner) HasPositionSide() bool {
	if o != nil && !common.IsNil(o.PositionSide) {
		return true
	}

	return false
}

// SetPositionSide gets a reference to the given string and assigns it to the PositionSide field.
func (o *GetUmAccountDetailResponsePositionsInner) SetPositionSide(v string) {
	o.PositionSide = &v
}

// GetPositionAmt returns the PositionAmt field value if set, zero value otherwise.
func (o *GetUmAccountDetailResponsePositionsInner) GetPositionAmt() string {
	if o == nil || common.IsNil(o.PositionAmt) {
		var ret string
		return ret
	}
	return *o.PositionAmt
}

// GetPositionAmtOk returns a tuple with the PositionAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailResponsePositionsInner) GetPositionAmtOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionAmt) {
		return nil, false
	}
	return o.PositionAmt, true
}

// HasPositionAmt returns a boolean if a field has been set.
func (o *GetUmAccountDetailResponsePositionsInner) HasPositionAmt() bool {
	if o != nil && !common.IsNil(o.PositionAmt) {
		return true
	}

	return false
}

// SetPositionAmt gets a reference to the given string and assigns it to the PositionAmt field.
func (o *GetUmAccountDetailResponsePositionsInner) SetPositionAmt(v string) {
	o.PositionAmt = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *GetUmAccountDetailResponsePositionsInner) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailResponsePositionsInner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *GetUmAccountDetailResponsePositionsInner) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *GetUmAccountDetailResponsePositionsInner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o GetUmAccountDetailResponsePositionsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUmAccountDetailResponsePositionsInner) ToMap() (map[string]interface{}, error) {
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

func (o *GetUmAccountDetailResponsePositionsInner) UnmarshalJSON(data []byte) (err error) {
	varGetUmAccountDetailResponsePositionsInner := _GetUmAccountDetailResponsePositionsInner{}

	err = json.Unmarshal(data, &varGetUmAccountDetailResponsePositionsInner)

	if err != nil {
		return err
	}

	*o = GetUmAccountDetailResponsePositionsInner(varGetUmAccountDetailResponsePositionsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "initialMargin")
		delete(additionalProperties, "maintMargin")
		delete(additionalProperties, "unrealizedProfit")
		delete(additionalProperties, "positionInitialMargin")
		delete(additionalProperties, "openOrderInitialMargin")
		delete(additionalProperties, "leverage")
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

type NullableGetUmAccountDetailResponsePositionsInner struct {
	value *GetUmAccountDetailResponsePositionsInner
	isSet bool
}

func (v NullableGetUmAccountDetailResponsePositionsInner) Get() *GetUmAccountDetailResponsePositionsInner {
	return v.value
}

func (v *NullableGetUmAccountDetailResponsePositionsInner) Set(val *GetUmAccountDetailResponsePositionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUmAccountDetailResponsePositionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUmAccountDetailResponsePositionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUmAccountDetailResponsePositionsInner(val *GetUmAccountDetailResponsePositionsInner) *NullableGetUmAccountDetailResponsePositionsInner {
	return &NullableGetUmAccountDetailResponsePositionsInner{value: val, isSet: true}
}

func (v NullableGetUmAccountDetailResponsePositionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUmAccountDetailResponsePositionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
