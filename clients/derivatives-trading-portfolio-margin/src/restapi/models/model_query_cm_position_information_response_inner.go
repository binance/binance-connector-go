/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryCmPositionInformationResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryCmPositionInformationResponseInner{}

// QueryCmPositionInformationResponseInner struct for QueryCmPositionInformationResponseInner
type QueryCmPositionInformationResponseInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	PositionAmt          *string `json:"positionAmt,omitempty"`
	EntryPrice           *string `json:"entryPrice,omitempty"`
	MarkPrice            *string `json:"markPrice,omitempty"`
	UnRealizedProfit     *string `json:"unRealizedProfit,omitempty"`
	LiquidationPrice     *string `json:"liquidationPrice,omitempty"`
	Leverage             *string `json:"leverage,omitempty"`
	PositionSide         *string `json:"positionSide,omitempty"`
	UpdateTime           *int64  `json:"updateTime,omitempty"`
	MaxQty               *string `json:"maxQty,omitempty"`
	NotionalValue        *string `json:"notionalValue,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryCmPositionInformationResponseInner QueryCmPositionInformationResponseInner

// NewQueryCmPositionInformationResponseInner instantiates a new QueryCmPositionInformationResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryCmPositionInformationResponseInner() *QueryCmPositionInformationResponseInner {
	this := QueryCmPositionInformationResponseInner{}
	return &this
}

// NewQueryCmPositionInformationResponseInnerWithDefaults instantiates a new QueryCmPositionInformationResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryCmPositionInformationResponseInnerWithDefaults() *QueryCmPositionInformationResponseInner {
	this := QueryCmPositionInformationResponseInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *QueryCmPositionInformationResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCmPositionInformationResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *QueryCmPositionInformationResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *QueryCmPositionInformationResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetPositionAmt returns the PositionAmt field value if set, zero value otherwise.
func (o *QueryCmPositionInformationResponseInner) GetPositionAmt() string {
	if o == nil || common.IsNil(o.PositionAmt) {
		var ret string
		return ret
	}
	return *o.PositionAmt
}

// GetPositionAmtOk returns a tuple with the PositionAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCmPositionInformationResponseInner) GetPositionAmtOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionAmt) {
		return nil, false
	}
	return o.PositionAmt, true
}

// HasPositionAmt returns a boolean if a field has been set.
func (o *QueryCmPositionInformationResponseInner) HasPositionAmt() bool {
	if o != nil && !common.IsNil(o.PositionAmt) {
		return true
	}

	return false
}

// SetPositionAmt gets a reference to the given string and assigns it to the PositionAmt field.
func (o *QueryCmPositionInformationResponseInner) SetPositionAmt(v string) {
	o.PositionAmt = &v
}

// GetEntryPrice returns the EntryPrice field value if set, zero value otherwise.
func (o *QueryCmPositionInformationResponseInner) GetEntryPrice() string {
	if o == nil || common.IsNil(o.EntryPrice) {
		var ret string
		return ret
	}
	return *o.EntryPrice
}

// GetEntryPriceOk returns a tuple with the EntryPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCmPositionInformationResponseInner) GetEntryPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.EntryPrice) {
		return nil, false
	}
	return o.EntryPrice, true
}

// HasEntryPrice returns a boolean if a field has been set.
func (o *QueryCmPositionInformationResponseInner) HasEntryPrice() bool {
	if o != nil && !common.IsNil(o.EntryPrice) {
		return true
	}

	return false
}

// SetEntryPrice gets a reference to the given string and assigns it to the EntryPrice field.
func (o *QueryCmPositionInformationResponseInner) SetEntryPrice(v string) {
	o.EntryPrice = &v
}

// GetMarkPrice returns the MarkPrice field value if set, zero value otherwise.
func (o *QueryCmPositionInformationResponseInner) GetMarkPrice() string {
	if o == nil || common.IsNil(o.MarkPrice) {
		var ret string
		return ret
	}
	return *o.MarkPrice
}

// GetMarkPriceOk returns a tuple with the MarkPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCmPositionInformationResponseInner) GetMarkPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarkPrice) {
		return nil, false
	}
	return o.MarkPrice, true
}

// HasMarkPrice returns a boolean if a field has been set.
func (o *QueryCmPositionInformationResponseInner) HasMarkPrice() bool {
	if o != nil && !common.IsNil(o.MarkPrice) {
		return true
	}

	return false
}

// SetMarkPrice gets a reference to the given string and assigns it to the MarkPrice field.
func (o *QueryCmPositionInformationResponseInner) SetMarkPrice(v string) {
	o.MarkPrice = &v
}

// GetUnRealizedProfit returns the UnRealizedProfit field value if set, zero value otherwise.
func (o *QueryCmPositionInformationResponseInner) GetUnRealizedProfit() string {
	if o == nil || common.IsNil(o.UnRealizedProfit) {
		var ret string
		return ret
	}
	return *o.UnRealizedProfit
}

// GetUnRealizedProfitOk returns a tuple with the UnRealizedProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCmPositionInformationResponseInner) GetUnRealizedProfitOk() (*string, bool) {
	if o == nil || common.IsNil(o.UnRealizedProfit) {
		return nil, false
	}
	return o.UnRealizedProfit, true
}

// HasUnRealizedProfit returns a boolean if a field has been set.
func (o *QueryCmPositionInformationResponseInner) HasUnRealizedProfit() bool {
	if o != nil && !common.IsNil(o.UnRealizedProfit) {
		return true
	}

	return false
}

// SetUnRealizedProfit gets a reference to the given string and assigns it to the UnRealizedProfit field.
func (o *QueryCmPositionInformationResponseInner) SetUnRealizedProfit(v string) {
	o.UnRealizedProfit = &v
}

// GetLiquidationPrice returns the LiquidationPrice field value if set, zero value otherwise.
func (o *QueryCmPositionInformationResponseInner) GetLiquidationPrice() string {
	if o == nil || common.IsNil(o.LiquidationPrice) {
		var ret string
		return ret
	}
	return *o.LiquidationPrice
}

// GetLiquidationPriceOk returns a tuple with the LiquidationPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCmPositionInformationResponseInner) GetLiquidationPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.LiquidationPrice) {
		return nil, false
	}
	return o.LiquidationPrice, true
}

// HasLiquidationPrice returns a boolean if a field has been set.
func (o *QueryCmPositionInformationResponseInner) HasLiquidationPrice() bool {
	if o != nil && !common.IsNil(o.LiquidationPrice) {
		return true
	}

	return false
}

// SetLiquidationPrice gets a reference to the given string and assigns it to the LiquidationPrice field.
func (o *QueryCmPositionInformationResponseInner) SetLiquidationPrice(v string) {
	o.LiquidationPrice = &v
}

// GetLeverage returns the Leverage field value if set, zero value otherwise.
func (o *QueryCmPositionInformationResponseInner) GetLeverage() string {
	if o == nil || common.IsNil(o.Leverage) {
		var ret string
		return ret
	}
	return *o.Leverage
}

// GetLeverageOk returns a tuple with the Leverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCmPositionInformationResponseInner) GetLeverageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Leverage) {
		return nil, false
	}
	return o.Leverage, true
}

// HasLeverage returns a boolean if a field has been set.
func (o *QueryCmPositionInformationResponseInner) HasLeverage() bool {
	if o != nil && !common.IsNil(o.Leverage) {
		return true
	}

	return false
}

// SetLeverage gets a reference to the given string and assigns it to the Leverage field.
func (o *QueryCmPositionInformationResponseInner) SetLeverage(v string) {
	o.Leverage = &v
}

// GetPositionSide returns the PositionSide field value if set, zero value otherwise.
func (o *QueryCmPositionInformationResponseInner) GetPositionSide() string {
	if o == nil || common.IsNil(o.PositionSide) {
		var ret string
		return ret
	}
	return *o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCmPositionInformationResponseInner) GetPositionSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionSide) {
		return nil, false
	}
	return o.PositionSide, true
}

// HasPositionSide returns a boolean if a field has been set.
func (o *QueryCmPositionInformationResponseInner) HasPositionSide() bool {
	if o != nil && !common.IsNil(o.PositionSide) {
		return true
	}

	return false
}

// SetPositionSide gets a reference to the given string and assigns it to the PositionSide field.
func (o *QueryCmPositionInformationResponseInner) SetPositionSide(v string) {
	o.PositionSide = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *QueryCmPositionInformationResponseInner) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCmPositionInformationResponseInner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *QueryCmPositionInformationResponseInner) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *QueryCmPositionInformationResponseInner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

// GetMaxQty returns the MaxQty field value if set, zero value otherwise.
func (o *QueryCmPositionInformationResponseInner) GetMaxQty() string {
	if o == nil || common.IsNil(o.MaxQty) {
		var ret string
		return ret
	}
	return *o.MaxQty
}

// GetMaxQtyOk returns a tuple with the MaxQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCmPositionInformationResponseInner) GetMaxQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxQty) {
		return nil, false
	}
	return o.MaxQty, true
}

// HasMaxQty returns a boolean if a field has been set.
func (o *QueryCmPositionInformationResponseInner) HasMaxQty() bool {
	if o != nil && !common.IsNil(o.MaxQty) {
		return true
	}

	return false
}

// SetMaxQty gets a reference to the given string and assigns it to the MaxQty field.
func (o *QueryCmPositionInformationResponseInner) SetMaxQty(v string) {
	o.MaxQty = &v
}

// GetNotionalValue returns the NotionalValue field value if set, zero value otherwise.
func (o *QueryCmPositionInformationResponseInner) GetNotionalValue() string {
	if o == nil || common.IsNil(o.NotionalValue) {
		var ret string
		return ret
	}
	return *o.NotionalValue
}

// GetNotionalValueOk returns a tuple with the NotionalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCmPositionInformationResponseInner) GetNotionalValueOk() (*string, bool) {
	if o == nil || common.IsNil(o.NotionalValue) {
		return nil, false
	}
	return o.NotionalValue, true
}

// HasNotionalValue returns a boolean if a field has been set.
func (o *QueryCmPositionInformationResponseInner) HasNotionalValue() bool {
	if o != nil && !common.IsNil(o.NotionalValue) {
		return true
	}

	return false
}

// SetNotionalValue gets a reference to the given string and assigns it to the NotionalValue field.
func (o *QueryCmPositionInformationResponseInner) SetNotionalValue(v string) {
	o.NotionalValue = &v
}

func (o QueryCmPositionInformationResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryCmPositionInformationResponseInner) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.PositionSide) {
		toSerialize["positionSide"] = o.PositionSide
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if !common.IsNil(o.MaxQty) {
		toSerialize["maxQty"] = o.MaxQty
	}
	if !common.IsNil(o.NotionalValue) {
		toSerialize["notionalValue"] = o.NotionalValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryCmPositionInformationResponseInner) UnmarshalJSON(data []byte) (err error) {
	varQueryCmPositionInformationResponseInner := _QueryCmPositionInformationResponseInner{}

	err = json.Unmarshal(data, &varQueryCmPositionInformationResponseInner)

	if err != nil {
		return err
	}

	*o = QueryCmPositionInformationResponseInner(varQueryCmPositionInformationResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "positionAmt")
		delete(additionalProperties, "entryPrice")
		delete(additionalProperties, "markPrice")
		delete(additionalProperties, "unRealizedProfit")
		delete(additionalProperties, "liquidationPrice")
		delete(additionalProperties, "leverage")
		delete(additionalProperties, "positionSide")
		delete(additionalProperties, "updateTime")
		delete(additionalProperties, "maxQty")
		delete(additionalProperties, "notionalValue")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryCmPositionInformationResponseInner struct {
	value *QueryCmPositionInformationResponseInner
	isSet bool
}

func (v NullableQueryCmPositionInformationResponseInner) Get() *QueryCmPositionInformationResponseInner {
	return v.value
}

func (v *NullableQueryCmPositionInformationResponseInner) Set(val *QueryCmPositionInformationResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryCmPositionInformationResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryCmPositionInformationResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryCmPositionInformationResponseInner(val *QueryCmPositionInformationResponseInner) *NullableQueryCmPositionInformationResponseInner {
	return &NullableQueryCmPositionInformationResponseInner{value: val, isSet: true}
}

func (v NullableQueryCmPositionInformationResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryCmPositionInformationResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
