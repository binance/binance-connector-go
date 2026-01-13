/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryIsolatedMarginAccountInfoResponseAssetsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryIsolatedMarginAccountInfoResponseAssetsInner{}

// QueryIsolatedMarginAccountInfoResponseAssetsInner struct for QueryIsolatedMarginAccountInfoResponseAssetsInner
type QueryIsolatedMarginAccountInfoResponseAssetsInner struct {
	BaseAsset            *QueryIsolatedMarginAccountInfoResponseAssetsInnerBaseAsset  `json:"baseAsset,omitempty"`
	QuoteAsset           *QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset `json:"quoteAsset,omitempty"`
	Symbol               *string                                                      `json:"symbol,omitempty"`
	IsolatedCreated      *bool                                                        `json:"isolatedCreated,omitempty"`
	Enabled              *bool                                                        `json:"enabled,omitempty"`
	MarginLevel          *string                                                      `json:"marginLevel,omitempty"`
	MarginLevelStatus    *string                                                      `json:"marginLevelStatus,omitempty"`
	MarginRatio          *string                                                      `json:"marginRatio,omitempty"`
	IndexPrice           *string                                                      `json:"indexPrice,omitempty"`
	LiquidatePrice       *string                                                      `json:"liquidatePrice,omitempty"`
	LiquidateRate        *string                                                      `json:"liquidateRate,omitempty"`
	TradeEnabled         *bool                                                        `json:"tradeEnabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryIsolatedMarginAccountInfoResponseAssetsInner QueryIsolatedMarginAccountInfoResponseAssetsInner

// NewQueryIsolatedMarginAccountInfoResponseAssetsInner instantiates a new QueryIsolatedMarginAccountInfoResponseAssetsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryIsolatedMarginAccountInfoResponseAssetsInner() *QueryIsolatedMarginAccountInfoResponseAssetsInner {
	this := QueryIsolatedMarginAccountInfoResponseAssetsInner{}
	return &this
}

// NewQueryIsolatedMarginAccountInfoResponseAssetsInnerWithDefaults instantiates a new QueryIsolatedMarginAccountInfoResponseAssetsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryIsolatedMarginAccountInfoResponseAssetsInnerWithDefaults() *QueryIsolatedMarginAccountInfoResponseAssetsInner {
	this := QueryIsolatedMarginAccountInfoResponseAssetsInner{}
	return &this
}

// GetBaseAsset returns the BaseAsset field value if set, zero value otherwise.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetBaseAsset() QueryIsolatedMarginAccountInfoResponseAssetsInnerBaseAsset {
	if o == nil || common.IsNil(o.BaseAsset) {
		var ret QueryIsolatedMarginAccountInfoResponseAssetsInnerBaseAsset
		return ret
	}
	return *o.BaseAsset
}

// GetBaseAssetOk returns a tuple with the BaseAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetBaseAssetOk() (*QueryIsolatedMarginAccountInfoResponseAssetsInnerBaseAsset, bool) {
	if o == nil || common.IsNil(o.BaseAsset) {
		return nil, false
	}
	return o.BaseAsset, true
}

// HasBaseAsset returns a boolean if a field has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) HasBaseAsset() bool {
	if o != nil && !common.IsNil(o.BaseAsset) {
		return true
	}

	return false
}

// SetBaseAsset gets a reference to the given QueryIsolatedMarginAccountInfoResponseAssetsInnerBaseAsset and assigns it to the BaseAsset field.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) SetBaseAsset(v QueryIsolatedMarginAccountInfoResponseAssetsInnerBaseAsset) {
	o.BaseAsset = &v
}

// GetQuoteAsset returns the QuoteAsset field value if set, zero value otherwise.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetQuoteAsset() QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset {
	if o == nil || common.IsNil(o.QuoteAsset) {
		var ret QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset
		return ret
	}
	return *o.QuoteAsset
}

// GetQuoteAssetOk returns a tuple with the QuoteAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetQuoteAssetOk() (*QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset, bool) {
	if o == nil || common.IsNil(o.QuoteAsset) {
		return nil, false
	}
	return o.QuoteAsset, true
}

// HasQuoteAsset returns a boolean if a field has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) HasQuoteAsset() bool {
	if o != nil && !common.IsNil(o.QuoteAsset) {
		return true
	}

	return false
}

// SetQuoteAsset gets a reference to the given QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset and assigns it to the QuoteAsset field.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) SetQuoteAsset(v QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset) {
	o.QuoteAsset = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetIsolatedCreated returns the IsolatedCreated field value if set, zero value otherwise.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetIsolatedCreated() bool {
	if o == nil || common.IsNil(o.IsolatedCreated) {
		var ret bool
		return ret
	}
	return *o.IsolatedCreated
}

// GetIsolatedCreatedOk returns a tuple with the IsolatedCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetIsolatedCreatedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsolatedCreated) {
		return nil, false
	}
	return o.IsolatedCreated, true
}

// HasIsolatedCreated returns a boolean if a field has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) HasIsolatedCreated() bool {
	if o != nil && !common.IsNil(o.IsolatedCreated) {
		return true
	}

	return false
}

// SetIsolatedCreated gets a reference to the given bool and assigns it to the IsolatedCreated field.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) SetIsolatedCreated(v bool) {
	o.IsolatedCreated = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetEnabled() bool {
	if o == nil || common.IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetEnabledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) HasEnabled() bool {
	if o != nil && !common.IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMarginLevel returns the MarginLevel field value if set, zero value otherwise.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetMarginLevel() string {
	if o == nil || common.IsNil(o.MarginLevel) {
		var ret string
		return ret
	}
	return *o.MarginLevel
}

// GetMarginLevelOk returns a tuple with the MarginLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetMarginLevelOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarginLevel) {
		return nil, false
	}
	return o.MarginLevel, true
}

// HasMarginLevel returns a boolean if a field has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) HasMarginLevel() bool {
	if o != nil && !common.IsNil(o.MarginLevel) {
		return true
	}

	return false
}

// SetMarginLevel gets a reference to the given string and assigns it to the MarginLevel field.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) SetMarginLevel(v string) {
	o.MarginLevel = &v
}

// GetMarginLevelStatus returns the MarginLevelStatus field value if set, zero value otherwise.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetMarginLevelStatus() string {
	if o == nil || common.IsNil(o.MarginLevelStatus) {
		var ret string
		return ret
	}
	return *o.MarginLevelStatus
}

// GetMarginLevelStatusOk returns a tuple with the MarginLevelStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetMarginLevelStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarginLevelStatus) {
		return nil, false
	}
	return o.MarginLevelStatus, true
}

// HasMarginLevelStatus returns a boolean if a field has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) HasMarginLevelStatus() bool {
	if o != nil && !common.IsNil(o.MarginLevelStatus) {
		return true
	}

	return false
}

// SetMarginLevelStatus gets a reference to the given string and assigns it to the MarginLevelStatus field.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) SetMarginLevelStatus(v string) {
	o.MarginLevelStatus = &v
}

// GetMarginRatio returns the MarginRatio field value if set, zero value otherwise.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetMarginRatio() string {
	if o == nil || common.IsNil(o.MarginRatio) {
		var ret string
		return ret
	}
	return *o.MarginRatio
}

// GetMarginRatioOk returns a tuple with the MarginRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetMarginRatioOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarginRatio) {
		return nil, false
	}
	return o.MarginRatio, true
}

// HasMarginRatio returns a boolean if a field has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) HasMarginRatio() bool {
	if o != nil && !common.IsNil(o.MarginRatio) {
		return true
	}

	return false
}

// SetMarginRatio gets a reference to the given string and assigns it to the MarginRatio field.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) SetMarginRatio(v string) {
	o.MarginRatio = &v
}

// GetIndexPrice returns the IndexPrice field value if set, zero value otherwise.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetIndexPrice() string {
	if o == nil || common.IsNil(o.IndexPrice) {
		var ret string
		return ret
	}
	return *o.IndexPrice
}

// GetIndexPriceOk returns a tuple with the IndexPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetIndexPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.IndexPrice) {
		return nil, false
	}
	return o.IndexPrice, true
}

// HasIndexPrice returns a boolean if a field has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) HasIndexPrice() bool {
	if o != nil && !common.IsNil(o.IndexPrice) {
		return true
	}

	return false
}

// SetIndexPrice gets a reference to the given string and assigns it to the IndexPrice field.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) SetIndexPrice(v string) {
	o.IndexPrice = &v
}

// GetLiquidatePrice returns the LiquidatePrice field value if set, zero value otherwise.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetLiquidatePrice() string {
	if o == nil || common.IsNil(o.LiquidatePrice) {
		var ret string
		return ret
	}
	return *o.LiquidatePrice
}

// GetLiquidatePriceOk returns a tuple with the LiquidatePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetLiquidatePriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.LiquidatePrice) {
		return nil, false
	}
	return o.LiquidatePrice, true
}

// HasLiquidatePrice returns a boolean if a field has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) HasLiquidatePrice() bool {
	if o != nil && !common.IsNil(o.LiquidatePrice) {
		return true
	}

	return false
}

// SetLiquidatePrice gets a reference to the given string and assigns it to the LiquidatePrice field.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) SetLiquidatePrice(v string) {
	o.LiquidatePrice = &v
}

// GetLiquidateRate returns the LiquidateRate field value if set, zero value otherwise.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetLiquidateRate() string {
	if o == nil || common.IsNil(o.LiquidateRate) {
		var ret string
		return ret
	}
	return *o.LiquidateRate
}

// GetLiquidateRateOk returns a tuple with the LiquidateRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetLiquidateRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.LiquidateRate) {
		return nil, false
	}
	return o.LiquidateRate, true
}

// HasLiquidateRate returns a boolean if a field has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) HasLiquidateRate() bool {
	if o != nil && !common.IsNil(o.LiquidateRate) {
		return true
	}

	return false
}

// SetLiquidateRate gets a reference to the given string and assigns it to the LiquidateRate field.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) SetLiquidateRate(v string) {
	o.LiquidateRate = &v
}

// GetTradeEnabled returns the TradeEnabled field value if set, zero value otherwise.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetTradeEnabled() bool {
	if o == nil || common.IsNil(o.TradeEnabled) {
		var ret bool
		return ret
	}
	return *o.TradeEnabled
}

// GetTradeEnabledOk returns a tuple with the TradeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetTradeEnabledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.TradeEnabled) {
		return nil, false
	}
	return o.TradeEnabled, true
}

// HasTradeEnabled returns a boolean if a field has been set.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) HasTradeEnabled() bool {
	if o != nil && !common.IsNil(o.TradeEnabled) {
		return true
	}

	return false
}

// SetTradeEnabled gets a reference to the given bool and assigns it to the TradeEnabled field.
func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) SetTradeEnabled(v bool) {
	o.TradeEnabled = &v
}

func (o QueryIsolatedMarginAccountInfoResponseAssetsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryIsolatedMarginAccountInfoResponseAssetsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BaseAsset) {
		toSerialize["baseAsset"] = o.BaseAsset
	}
	if !common.IsNil(o.QuoteAsset) {
		toSerialize["quoteAsset"] = o.QuoteAsset
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.IsolatedCreated) {
		toSerialize["isolatedCreated"] = o.IsolatedCreated
	}
	if !common.IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !common.IsNil(o.MarginLevel) {
		toSerialize["marginLevel"] = o.MarginLevel
	}
	if !common.IsNil(o.MarginLevelStatus) {
		toSerialize["marginLevelStatus"] = o.MarginLevelStatus
	}
	if !common.IsNil(o.MarginRatio) {
		toSerialize["marginRatio"] = o.MarginRatio
	}
	if !common.IsNil(o.IndexPrice) {
		toSerialize["indexPrice"] = o.IndexPrice
	}
	if !common.IsNil(o.LiquidatePrice) {
		toSerialize["liquidatePrice"] = o.LiquidatePrice
	}
	if !common.IsNil(o.LiquidateRate) {
		toSerialize["liquidateRate"] = o.LiquidateRate
	}
	if !common.IsNil(o.TradeEnabled) {
		toSerialize["tradeEnabled"] = o.TradeEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) UnmarshalJSON(data []byte) (err error) {
	varQueryIsolatedMarginAccountInfoResponseAssetsInner := _QueryIsolatedMarginAccountInfoResponseAssetsInner{}

	err = json.Unmarshal(data, &varQueryIsolatedMarginAccountInfoResponseAssetsInner)

	if err != nil {
		return err
	}

	*o = QueryIsolatedMarginAccountInfoResponseAssetsInner(varQueryIsolatedMarginAccountInfoResponseAssetsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "baseAsset")
		delete(additionalProperties, "quoteAsset")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "isolatedCreated")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "marginLevel")
		delete(additionalProperties, "marginLevelStatus")
		delete(additionalProperties, "marginRatio")
		delete(additionalProperties, "indexPrice")
		delete(additionalProperties, "liquidatePrice")
		delete(additionalProperties, "liquidateRate")
		delete(additionalProperties, "tradeEnabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryIsolatedMarginAccountInfoResponseAssetsInner struct {
	value *QueryIsolatedMarginAccountInfoResponseAssetsInner
	isSet bool
}

func (v NullableQueryIsolatedMarginAccountInfoResponseAssetsInner) Get() *QueryIsolatedMarginAccountInfoResponseAssetsInner {
	return v.value
}

func (v *NullableQueryIsolatedMarginAccountInfoResponseAssetsInner) Set(val *QueryIsolatedMarginAccountInfoResponseAssetsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryIsolatedMarginAccountInfoResponseAssetsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryIsolatedMarginAccountInfoResponseAssetsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryIsolatedMarginAccountInfoResponseAssetsInner(val *QueryIsolatedMarginAccountInfoResponseAssetsInner) *NullableQueryIsolatedMarginAccountInfoResponseAssetsInner {
	return &NullableQueryIsolatedMarginAccountInfoResponseAssetsInner{value: val, isSet: true}
}

func (v NullableQueryIsolatedMarginAccountInfoResponseAssetsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryIsolatedMarginAccountInfoResponseAssetsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
