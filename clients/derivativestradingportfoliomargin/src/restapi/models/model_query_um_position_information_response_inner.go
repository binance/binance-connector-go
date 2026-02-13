/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryUmPositionInformationResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryUmPositionInformationResponseInner{}

// QueryUmPositionInformationResponseInner struct for QueryUmPositionInformationResponseInner
type QueryUmPositionInformationResponseInner struct {
	EntryPrice           *string `json:"entryPrice,omitempty"`
	Leverage             *string `json:"leverage,omitempty"`
	MarkPrice            *string `json:"markPrice,omitempty"`
	MaxNotionalValue     *string `json:"maxNotionalValue,omitempty"`
	PositionAmt          *string `json:"positionAmt,omitempty"`
	Notional             *string `json:"notional,omitempty"`
	Symbol               *string `json:"symbol,omitempty"`
	UnRealizedProfit     *string `json:"unRealizedProfit,omitempty"`
	LiquidationPrice     *string `json:"liquidationPrice,omitempty"`
	PositionSide         *string `json:"positionSide,omitempty"`
	UpdateTime           *int64  `json:"updateTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryUmPositionInformationResponseInner QueryUmPositionInformationResponseInner

// NewQueryUmPositionInformationResponseInner instantiates a new QueryUmPositionInformationResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryUmPositionInformationResponseInner() *QueryUmPositionInformationResponseInner {
	this := QueryUmPositionInformationResponseInner{}
	return &this
}

// NewQueryUmPositionInformationResponseInnerWithDefaults instantiates a new QueryUmPositionInformationResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryUmPositionInformationResponseInnerWithDefaults() *QueryUmPositionInformationResponseInner {
	this := QueryUmPositionInformationResponseInner{}
	return &this
}

// GetEntryPrice returns the EntryPrice field value if set, zero value otherwise.
func (o *QueryUmPositionInformationResponseInner) GetEntryPrice() string {
	if o == nil || common.IsNil(o.EntryPrice) {
		var ret string
		return ret
	}
	return *o.EntryPrice
}

// GetEntryPriceOk returns a tuple with the EntryPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmPositionInformationResponseInner) GetEntryPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.EntryPrice) {
		return nil, false
	}
	return o.EntryPrice, true
}

// HasEntryPrice returns a boolean if a field has been set.
func (o *QueryUmPositionInformationResponseInner) HasEntryPrice() bool {
	if o != nil && !common.IsNil(o.EntryPrice) {
		return true
	}

	return false
}

// SetEntryPrice gets a reference to the given string and assigns it to the EntryPrice field.
func (o *QueryUmPositionInformationResponseInner) SetEntryPrice(v string) {
	o.EntryPrice = &v
}

// GetLeverage returns the Leverage field value if set, zero value otherwise.
func (o *QueryUmPositionInformationResponseInner) GetLeverage() string {
	if o == nil || common.IsNil(o.Leverage) {
		var ret string
		return ret
	}
	return *o.Leverage
}

// GetLeverageOk returns a tuple with the Leverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmPositionInformationResponseInner) GetLeverageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Leverage) {
		return nil, false
	}
	return o.Leverage, true
}

// HasLeverage returns a boolean if a field has been set.
func (o *QueryUmPositionInformationResponseInner) HasLeverage() bool {
	if o != nil && !common.IsNil(o.Leverage) {
		return true
	}

	return false
}

// SetLeverage gets a reference to the given string and assigns it to the Leverage field.
func (o *QueryUmPositionInformationResponseInner) SetLeverage(v string) {
	o.Leverage = &v
}

// GetMarkPrice returns the MarkPrice field value if set, zero value otherwise.
func (o *QueryUmPositionInformationResponseInner) GetMarkPrice() string {
	if o == nil || common.IsNil(o.MarkPrice) {
		var ret string
		return ret
	}
	return *o.MarkPrice
}

// GetMarkPriceOk returns a tuple with the MarkPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmPositionInformationResponseInner) GetMarkPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarkPrice) {
		return nil, false
	}
	return o.MarkPrice, true
}

// HasMarkPrice returns a boolean if a field has been set.
func (o *QueryUmPositionInformationResponseInner) HasMarkPrice() bool {
	if o != nil && !common.IsNil(o.MarkPrice) {
		return true
	}

	return false
}

// SetMarkPrice gets a reference to the given string and assigns it to the MarkPrice field.
func (o *QueryUmPositionInformationResponseInner) SetMarkPrice(v string) {
	o.MarkPrice = &v
}

// GetMaxNotionalValue returns the MaxNotionalValue field value if set, zero value otherwise.
func (o *QueryUmPositionInformationResponseInner) GetMaxNotionalValue() string {
	if o == nil || common.IsNil(o.MaxNotionalValue) {
		var ret string
		return ret
	}
	return *o.MaxNotionalValue
}

// GetMaxNotionalValueOk returns a tuple with the MaxNotionalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmPositionInformationResponseInner) GetMaxNotionalValueOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxNotionalValue) {
		return nil, false
	}
	return o.MaxNotionalValue, true
}

// HasMaxNotionalValue returns a boolean if a field has been set.
func (o *QueryUmPositionInformationResponseInner) HasMaxNotionalValue() bool {
	if o != nil && !common.IsNil(o.MaxNotionalValue) {
		return true
	}

	return false
}

// SetMaxNotionalValue gets a reference to the given string and assigns it to the MaxNotionalValue field.
func (o *QueryUmPositionInformationResponseInner) SetMaxNotionalValue(v string) {
	o.MaxNotionalValue = &v
}

// GetPositionAmt returns the PositionAmt field value if set, zero value otherwise.
func (o *QueryUmPositionInformationResponseInner) GetPositionAmt() string {
	if o == nil || common.IsNil(o.PositionAmt) {
		var ret string
		return ret
	}
	return *o.PositionAmt
}

// GetPositionAmtOk returns a tuple with the PositionAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmPositionInformationResponseInner) GetPositionAmtOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionAmt) {
		return nil, false
	}
	return o.PositionAmt, true
}

// HasPositionAmt returns a boolean if a field has been set.
func (o *QueryUmPositionInformationResponseInner) HasPositionAmt() bool {
	if o != nil && !common.IsNil(o.PositionAmt) {
		return true
	}

	return false
}

// SetPositionAmt gets a reference to the given string and assigns it to the PositionAmt field.
func (o *QueryUmPositionInformationResponseInner) SetPositionAmt(v string) {
	o.PositionAmt = &v
}

// GetNotional returns the Notional field value if set, zero value otherwise.
func (o *QueryUmPositionInformationResponseInner) GetNotional() string {
	if o == nil || common.IsNil(o.Notional) {
		var ret string
		return ret
	}
	return *o.Notional
}

// GetNotionalOk returns a tuple with the Notional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmPositionInformationResponseInner) GetNotionalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Notional) {
		return nil, false
	}
	return o.Notional, true
}

// HasNotional returns a boolean if a field has been set.
func (o *QueryUmPositionInformationResponseInner) HasNotional() bool {
	if o != nil && !common.IsNil(o.Notional) {
		return true
	}

	return false
}

// SetNotional gets a reference to the given string and assigns it to the Notional field.
func (o *QueryUmPositionInformationResponseInner) SetNotional(v string) {
	o.Notional = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *QueryUmPositionInformationResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmPositionInformationResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *QueryUmPositionInformationResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *QueryUmPositionInformationResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetUnRealizedProfit returns the UnRealizedProfit field value if set, zero value otherwise.
func (o *QueryUmPositionInformationResponseInner) GetUnRealizedProfit() string {
	if o == nil || common.IsNil(o.UnRealizedProfit) {
		var ret string
		return ret
	}
	return *o.UnRealizedProfit
}

// GetUnRealizedProfitOk returns a tuple with the UnRealizedProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmPositionInformationResponseInner) GetUnRealizedProfitOk() (*string, bool) {
	if o == nil || common.IsNil(o.UnRealizedProfit) {
		return nil, false
	}
	return o.UnRealizedProfit, true
}

// HasUnRealizedProfit returns a boolean if a field has been set.
func (o *QueryUmPositionInformationResponseInner) HasUnRealizedProfit() bool {
	if o != nil && !common.IsNil(o.UnRealizedProfit) {
		return true
	}

	return false
}

// SetUnRealizedProfit gets a reference to the given string and assigns it to the UnRealizedProfit field.
func (o *QueryUmPositionInformationResponseInner) SetUnRealizedProfit(v string) {
	o.UnRealizedProfit = &v
}

// GetLiquidationPrice returns the LiquidationPrice field value if set, zero value otherwise.
func (o *QueryUmPositionInformationResponseInner) GetLiquidationPrice() string {
	if o == nil || common.IsNil(o.LiquidationPrice) {
		var ret string
		return ret
	}
	return *o.LiquidationPrice
}

// GetLiquidationPriceOk returns a tuple with the LiquidationPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmPositionInformationResponseInner) GetLiquidationPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.LiquidationPrice) {
		return nil, false
	}
	return o.LiquidationPrice, true
}

// HasLiquidationPrice returns a boolean if a field has been set.
func (o *QueryUmPositionInformationResponseInner) HasLiquidationPrice() bool {
	if o != nil && !common.IsNil(o.LiquidationPrice) {
		return true
	}

	return false
}

// SetLiquidationPrice gets a reference to the given string and assigns it to the LiquidationPrice field.
func (o *QueryUmPositionInformationResponseInner) SetLiquidationPrice(v string) {
	o.LiquidationPrice = &v
}

// GetPositionSide returns the PositionSide field value if set, zero value otherwise.
func (o *QueryUmPositionInformationResponseInner) GetPositionSide() string {
	if o == nil || common.IsNil(o.PositionSide) {
		var ret string
		return ret
	}
	return *o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmPositionInformationResponseInner) GetPositionSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionSide) {
		return nil, false
	}
	return o.PositionSide, true
}

// HasPositionSide returns a boolean if a field has been set.
func (o *QueryUmPositionInformationResponseInner) HasPositionSide() bool {
	if o != nil && !common.IsNil(o.PositionSide) {
		return true
	}

	return false
}

// SetPositionSide gets a reference to the given string and assigns it to the PositionSide field.
func (o *QueryUmPositionInformationResponseInner) SetPositionSide(v string) {
	o.PositionSide = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *QueryUmPositionInformationResponseInner) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmPositionInformationResponseInner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *QueryUmPositionInformationResponseInner) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *QueryUmPositionInformationResponseInner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o QueryUmPositionInformationResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryUmPositionInformationResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.EntryPrice) {
		toSerialize["entryPrice"] = o.EntryPrice
	}
	if !common.IsNil(o.Leverage) {
		toSerialize["leverage"] = o.Leverage
	}
	if !common.IsNil(o.MarkPrice) {
		toSerialize["markPrice"] = o.MarkPrice
	}
	if !common.IsNil(o.MaxNotionalValue) {
		toSerialize["maxNotionalValue"] = o.MaxNotionalValue
	}
	if !common.IsNil(o.PositionAmt) {
		toSerialize["positionAmt"] = o.PositionAmt
	}
	if !common.IsNil(o.Notional) {
		toSerialize["notional"] = o.Notional
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.UnRealizedProfit) {
		toSerialize["unRealizedProfit"] = o.UnRealizedProfit
	}
	if !common.IsNil(o.LiquidationPrice) {
		toSerialize["liquidationPrice"] = o.LiquidationPrice
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

func (o *QueryUmPositionInformationResponseInner) UnmarshalJSON(data []byte) (err error) {
	varQueryUmPositionInformationResponseInner := _QueryUmPositionInformationResponseInner{}

	err = json.Unmarshal(data, &varQueryUmPositionInformationResponseInner)

	if err != nil {
		return err
	}

	*o = QueryUmPositionInformationResponseInner(varQueryUmPositionInformationResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "entryPrice")
		delete(additionalProperties, "leverage")
		delete(additionalProperties, "markPrice")
		delete(additionalProperties, "maxNotionalValue")
		delete(additionalProperties, "positionAmt")
		delete(additionalProperties, "notional")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "unRealizedProfit")
		delete(additionalProperties, "liquidationPrice")
		delete(additionalProperties, "positionSide")
		delete(additionalProperties, "updateTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryUmPositionInformationResponseInner struct {
	value *QueryUmPositionInformationResponseInner
	isSet bool
}

func (v NullableQueryUmPositionInformationResponseInner) Get() *QueryUmPositionInformationResponseInner {
	return v.value
}

func (v *NullableQueryUmPositionInformationResponseInner) Set(val *QueryUmPositionInformationResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryUmPositionInformationResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryUmPositionInformationResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryUmPositionInformationResponseInner(val *QueryUmPositionInformationResponseInner) *NullableQueryUmPositionInformationResponseInner {
	return &NullableQueryUmPositionInformationResponseInner{value: val, isSet: true}
}

func (v NullableQueryUmPositionInformationResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryUmPositionInformationResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
