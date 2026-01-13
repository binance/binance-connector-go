/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the PositionInformationResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PositionInformationResponseInner{}

// PositionInformationResponseInner struct for PositionInformationResponseInner
type PositionInformationResponseInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	PositionAmt          *string `json:"positionAmt,omitempty"`
	EntryPrice           *string `json:"entryPrice,omitempty"`
	BreakEvenPrice       *string `json:"breakEvenPrice,omitempty"`
	MarkPrice            *string `json:"markPrice,omitempty"`
	UnRealizedProfit     *string `json:"unRealizedProfit,omitempty"`
	LiquidationPrice     *string `json:"liquidationPrice,omitempty"`
	Leverage             *string `json:"leverage,omitempty"`
	MaxQty               *string `json:"maxQty,omitempty"`
	MarginType           *string `json:"marginType,omitempty"`
	IsolatedMargin       *string `json:"isolatedMargin,omitempty"`
	IsAutoAddMargin      *string `json:"isAutoAddMargin,omitempty"`
	PositionSide         *string `json:"positionSide,omitempty"`
	UpdateTime           *int64  `json:"updateTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PositionInformationResponseInner PositionInformationResponseInner

// NewPositionInformationResponseInner instantiates a new PositionInformationResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPositionInformationResponseInner() *PositionInformationResponseInner {
	this := PositionInformationResponseInner{}
	return &this
}

// NewPositionInformationResponseInnerWithDefaults instantiates a new PositionInformationResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPositionInformationResponseInnerWithDefaults() *PositionInformationResponseInner {
	this := PositionInformationResponseInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *PositionInformationResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *PositionInformationResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *PositionInformationResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetPositionAmt returns the PositionAmt field value if set, zero value otherwise.
func (o *PositionInformationResponseInner) GetPositionAmt() string {
	if o == nil || common.IsNil(o.PositionAmt) {
		var ret string
		return ret
	}
	return *o.PositionAmt
}

// GetPositionAmtOk returns a tuple with the PositionAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationResponseInner) GetPositionAmtOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionAmt) {
		return nil, false
	}
	return o.PositionAmt, true
}

// HasPositionAmt returns a boolean if a field has been set.
func (o *PositionInformationResponseInner) HasPositionAmt() bool {
	if o != nil && !common.IsNil(o.PositionAmt) {
		return true
	}

	return false
}

// SetPositionAmt gets a reference to the given string and assigns it to the PositionAmt field.
func (o *PositionInformationResponseInner) SetPositionAmt(v string) {
	o.PositionAmt = &v
}

// GetEntryPrice returns the EntryPrice field value if set, zero value otherwise.
func (o *PositionInformationResponseInner) GetEntryPrice() string {
	if o == nil || common.IsNil(o.EntryPrice) {
		var ret string
		return ret
	}
	return *o.EntryPrice
}

// GetEntryPriceOk returns a tuple with the EntryPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationResponseInner) GetEntryPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.EntryPrice) {
		return nil, false
	}
	return o.EntryPrice, true
}

// HasEntryPrice returns a boolean if a field has been set.
func (o *PositionInformationResponseInner) HasEntryPrice() bool {
	if o != nil && !common.IsNil(o.EntryPrice) {
		return true
	}

	return false
}

// SetEntryPrice gets a reference to the given string and assigns it to the EntryPrice field.
func (o *PositionInformationResponseInner) SetEntryPrice(v string) {
	o.EntryPrice = &v
}

// GetBreakEvenPrice returns the BreakEvenPrice field value if set, zero value otherwise.
func (o *PositionInformationResponseInner) GetBreakEvenPrice() string {
	if o == nil || common.IsNil(o.BreakEvenPrice) {
		var ret string
		return ret
	}
	return *o.BreakEvenPrice
}

// GetBreakEvenPriceOk returns a tuple with the BreakEvenPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationResponseInner) GetBreakEvenPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.BreakEvenPrice) {
		return nil, false
	}
	return o.BreakEvenPrice, true
}

// HasBreakEvenPrice returns a boolean if a field has been set.
func (o *PositionInformationResponseInner) HasBreakEvenPrice() bool {
	if o != nil && !common.IsNil(o.BreakEvenPrice) {
		return true
	}

	return false
}

// SetBreakEvenPrice gets a reference to the given string and assigns it to the BreakEvenPrice field.
func (o *PositionInformationResponseInner) SetBreakEvenPrice(v string) {
	o.BreakEvenPrice = &v
}

// GetMarkPrice returns the MarkPrice field value if set, zero value otherwise.
func (o *PositionInformationResponseInner) GetMarkPrice() string {
	if o == nil || common.IsNil(o.MarkPrice) {
		var ret string
		return ret
	}
	return *o.MarkPrice
}

// GetMarkPriceOk returns a tuple with the MarkPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationResponseInner) GetMarkPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarkPrice) {
		return nil, false
	}
	return o.MarkPrice, true
}

// HasMarkPrice returns a boolean if a field has been set.
func (o *PositionInformationResponseInner) HasMarkPrice() bool {
	if o != nil && !common.IsNil(o.MarkPrice) {
		return true
	}

	return false
}

// SetMarkPrice gets a reference to the given string and assigns it to the MarkPrice field.
func (o *PositionInformationResponseInner) SetMarkPrice(v string) {
	o.MarkPrice = &v
}

// GetUnRealizedProfit returns the UnRealizedProfit field value if set, zero value otherwise.
func (o *PositionInformationResponseInner) GetUnRealizedProfit() string {
	if o == nil || common.IsNil(o.UnRealizedProfit) {
		var ret string
		return ret
	}
	return *o.UnRealizedProfit
}

// GetUnRealizedProfitOk returns a tuple with the UnRealizedProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationResponseInner) GetUnRealizedProfitOk() (*string, bool) {
	if o == nil || common.IsNil(o.UnRealizedProfit) {
		return nil, false
	}
	return o.UnRealizedProfit, true
}

// HasUnRealizedProfit returns a boolean if a field has been set.
func (o *PositionInformationResponseInner) HasUnRealizedProfit() bool {
	if o != nil && !common.IsNil(o.UnRealizedProfit) {
		return true
	}

	return false
}

// SetUnRealizedProfit gets a reference to the given string and assigns it to the UnRealizedProfit field.
func (o *PositionInformationResponseInner) SetUnRealizedProfit(v string) {
	o.UnRealizedProfit = &v
}

// GetLiquidationPrice returns the LiquidationPrice field value if set, zero value otherwise.
func (o *PositionInformationResponseInner) GetLiquidationPrice() string {
	if o == nil || common.IsNil(o.LiquidationPrice) {
		var ret string
		return ret
	}
	return *o.LiquidationPrice
}

// GetLiquidationPriceOk returns a tuple with the LiquidationPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationResponseInner) GetLiquidationPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.LiquidationPrice) {
		return nil, false
	}
	return o.LiquidationPrice, true
}

// HasLiquidationPrice returns a boolean if a field has been set.
func (o *PositionInformationResponseInner) HasLiquidationPrice() bool {
	if o != nil && !common.IsNil(o.LiquidationPrice) {
		return true
	}

	return false
}

// SetLiquidationPrice gets a reference to the given string and assigns it to the LiquidationPrice field.
func (o *PositionInformationResponseInner) SetLiquidationPrice(v string) {
	o.LiquidationPrice = &v
}

// GetLeverage returns the Leverage field value if set, zero value otherwise.
func (o *PositionInformationResponseInner) GetLeverage() string {
	if o == nil || common.IsNil(o.Leverage) {
		var ret string
		return ret
	}
	return *o.Leverage
}

// GetLeverageOk returns a tuple with the Leverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationResponseInner) GetLeverageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Leverage) {
		return nil, false
	}
	return o.Leverage, true
}

// HasLeverage returns a boolean if a field has been set.
func (o *PositionInformationResponseInner) HasLeverage() bool {
	if o != nil && !common.IsNil(o.Leverage) {
		return true
	}

	return false
}

// SetLeverage gets a reference to the given string and assigns it to the Leverage field.
func (o *PositionInformationResponseInner) SetLeverage(v string) {
	o.Leverage = &v
}

// GetMaxQty returns the MaxQty field value if set, zero value otherwise.
func (o *PositionInformationResponseInner) GetMaxQty() string {
	if o == nil || common.IsNil(o.MaxQty) {
		var ret string
		return ret
	}
	return *o.MaxQty
}

// GetMaxQtyOk returns a tuple with the MaxQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationResponseInner) GetMaxQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxQty) {
		return nil, false
	}
	return o.MaxQty, true
}

// HasMaxQty returns a boolean if a field has been set.
func (o *PositionInformationResponseInner) HasMaxQty() bool {
	if o != nil && !common.IsNil(o.MaxQty) {
		return true
	}

	return false
}

// SetMaxQty gets a reference to the given string and assigns it to the MaxQty field.
func (o *PositionInformationResponseInner) SetMaxQty(v string) {
	o.MaxQty = &v
}

// GetMarginType returns the MarginType field value if set, zero value otherwise.
func (o *PositionInformationResponseInner) GetMarginType() string {
	if o == nil || common.IsNil(o.MarginType) {
		var ret string
		return ret
	}
	return *o.MarginType
}

// GetMarginTypeOk returns a tuple with the MarginType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationResponseInner) GetMarginTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarginType) {
		return nil, false
	}
	return o.MarginType, true
}

// HasMarginType returns a boolean if a field has been set.
func (o *PositionInformationResponseInner) HasMarginType() bool {
	if o != nil && !common.IsNil(o.MarginType) {
		return true
	}

	return false
}

// SetMarginType gets a reference to the given string and assigns it to the MarginType field.
func (o *PositionInformationResponseInner) SetMarginType(v string) {
	o.MarginType = &v
}

// GetIsolatedMargin returns the IsolatedMargin field value if set, zero value otherwise.
func (o *PositionInformationResponseInner) GetIsolatedMargin() string {
	if o == nil || common.IsNil(o.IsolatedMargin) {
		var ret string
		return ret
	}
	return *o.IsolatedMargin
}

// GetIsolatedMarginOk returns a tuple with the IsolatedMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationResponseInner) GetIsolatedMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.IsolatedMargin) {
		return nil, false
	}
	return o.IsolatedMargin, true
}

// HasIsolatedMargin returns a boolean if a field has been set.
func (o *PositionInformationResponseInner) HasIsolatedMargin() bool {
	if o != nil && !common.IsNil(o.IsolatedMargin) {
		return true
	}

	return false
}

// SetIsolatedMargin gets a reference to the given string and assigns it to the IsolatedMargin field.
func (o *PositionInformationResponseInner) SetIsolatedMargin(v string) {
	o.IsolatedMargin = &v
}

// GetIsAutoAddMargin returns the IsAutoAddMargin field value if set, zero value otherwise.
func (o *PositionInformationResponseInner) GetIsAutoAddMargin() string {
	if o == nil || common.IsNil(o.IsAutoAddMargin) {
		var ret string
		return ret
	}
	return *o.IsAutoAddMargin
}

// GetIsAutoAddMarginOk returns a tuple with the IsAutoAddMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationResponseInner) GetIsAutoAddMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.IsAutoAddMargin) {
		return nil, false
	}
	return o.IsAutoAddMargin, true
}

// HasIsAutoAddMargin returns a boolean if a field has been set.
func (o *PositionInformationResponseInner) HasIsAutoAddMargin() bool {
	if o != nil && !common.IsNil(o.IsAutoAddMargin) {
		return true
	}

	return false
}

// SetIsAutoAddMargin gets a reference to the given string and assigns it to the IsAutoAddMargin field.
func (o *PositionInformationResponseInner) SetIsAutoAddMargin(v string) {
	o.IsAutoAddMargin = &v
}

// GetPositionSide returns the PositionSide field value if set, zero value otherwise.
func (o *PositionInformationResponseInner) GetPositionSide() string {
	if o == nil || common.IsNil(o.PositionSide) {
		var ret string
		return ret
	}
	return *o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationResponseInner) GetPositionSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionSide) {
		return nil, false
	}
	return o.PositionSide, true
}

// HasPositionSide returns a boolean if a field has been set.
func (o *PositionInformationResponseInner) HasPositionSide() bool {
	if o != nil && !common.IsNil(o.PositionSide) {
		return true
	}

	return false
}

// SetPositionSide gets a reference to the given string and assigns it to the PositionSide field.
func (o *PositionInformationResponseInner) SetPositionSide(v string) {
	o.PositionSide = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *PositionInformationResponseInner) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionInformationResponseInner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *PositionInformationResponseInner) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *PositionInformationResponseInner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o PositionInformationResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PositionInformationResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.PositionAmt) {
		toSerialize["positionAmt"] = o.PositionAmt
	}
	if !common.IsNil(o.EntryPrice) {
		toSerialize["entryPrice"] = o.EntryPrice
	}
	if !common.IsNil(o.BreakEvenPrice) {
		toSerialize["breakEvenPrice"] = o.BreakEvenPrice
	}
	if !common.IsNil(o.MarkPrice) {
		toSerialize["markPrice"] = o.MarkPrice
	}
	if !common.IsNil(o.UnRealizedProfit) {
		toSerialize["unRealizedProfit"] = o.UnRealizedProfit
	}
	if !common.IsNil(o.LiquidationPrice) {
		toSerialize["liquidationPrice"] = o.LiquidationPrice
	}
	if !common.IsNil(o.Leverage) {
		toSerialize["leverage"] = o.Leverage
	}
	if !common.IsNil(o.MaxQty) {
		toSerialize["maxQty"] = o.MaxQty
	}
	if !common.IsNil(o.MarginType) {
		toSerialize["marginType"] = o.MarginType
	}
	if !common.IsNil(o.IsolatedMargin) {
		toSerialize["isolatedMargin"] = o.IsolatedMargin
	}
	if !common.IsNil(o.IsAutoAddMargin) {
		toSerialize["isAutoAddMargin"] = o.IsAutoAddMargin
	}
	if !common.IsNil(o.PositionSide) {
		toSerialize["positionSide"] = o.PositionSide
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PositionInformationResponseInner) UnmarshalJSON(data []byte) (err error) {
	varPositionInformationResponseInner := _PositionInformationResponseInner{}

	err = json.Unmarshal(data, &varPositionInformationResponseInner)

	if err != nil {
		return err
	}

	*o = PositionInformationResponseInner(varPositionInformationResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "positionAmt")
		delete(additionalProperties, "entryPrice")
		delete(additionalProperties, "breakEvenPrice")
		delete(additionalProperties, "markPrice")
		delete(additionalProperties, "unRealizedProfit")
		delete(additionalProperties, "liquidationPrice")
		delete(additionalProperties, "leverage")
		delete(additionalProperties, "maxQty")
		delete(additionalProperties, "marginType")
		delete(additionalProperties, "isolatedMargin")
		delete(additionalProperties, "isAutoAddMargin")
		delete(additionalProperties, "positionSide")
		delete(additionalProperties, "updateTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePositionInformationResponseInner struct {
	value *PositionInformationResponseInner
	isSet bool
}

func (v NullablePositionInformationResponseInner) Get() *PositionInformationResponseInner {
	return v.value
}

func (v *NullablePositionInformationResponseInner) Set(val *PositionInformationResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePositionInformationResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePositionInformationResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePositionInformationResponseInner(val *PositionInformationResponseInner) *NullablePositionInformationResponseInner {
	return &NullablePositionInformationResponseInner{value: val, isSet: true}
}

func (v NullablePositionInformationResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePositionInformationResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
