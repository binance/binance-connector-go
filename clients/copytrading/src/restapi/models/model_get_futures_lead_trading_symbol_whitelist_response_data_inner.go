/*
Binance Copy Trading REST API

OpenAPI Specification for the Binance Copy Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetFuturesLeadTradingSymbolWhitelistResponseDataInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFuturesLeadTradingSymbolWhitelistResponseDataInner{}

// GetFuturesLeadTradingSymbolWhitelistResponseDataInner struct for GetFuturesLeadTradingSymbolWhitelistResponseDataInner
type GetFuturesLeadTradingSymbolWhitelistResponseDataInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	BaseAsset            *string `json:"baseAsset,omitempty"`
	QuoteAsset           *string `json:"quoteAsset,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetFuturesLeadTradingSymbolWhitelistResponseDataInner GetFuturesLeadTradingSymbolWhitelistResponseDataInner

// NewGetFuturesLeadTradingSymbolWhitelistResponseDataInner instantiates a new GetFuturesLeadTradingSymbolWhitelistResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFuturesLeadTradingSymbolWhitelistResponseDataInner() *GetFuturesLeadTradingSymbolWhitelistResponseDataInner {
	this := GetFuturesLeadTradingSymbolWhitelistResponseDataInner{}
	return &this
}

// NewGetFuturesLeadTradingSymbolWhitelistResponseDataInnerWithDefaults instantiates a new GetFuturesLeadTradingSymbolWhitelistResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFuturesLeadTradingSymbolWhitelistResponseDataInnerWithDefaults() *GetFuturesLeadTradingSymbolWhitelistResponseDataInner {
	this := GetFuturesLeadTradingSymbolWhitelistResponseDataInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetFuturesLeadTradingSymbolWhitelistResponseDataInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesLeadTradingSymbolWhitelistResponseDataInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetFuturesLeadTradingSymbolWhitelistResponseDataInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetFuturesLeadTradingSymbolWhitelistResponseDataInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetBaseAsset returns the BaseAsset field value if set, zero value otherwise.
func (o *GetFuturesLeadTradingSymbolWhitelistResponseDataInner) GetBaseAsset() string {
	if o == nil || common.IsNil(o.BaseAsset) {
		var ret string
		return ret
	}
	return *o.BaseAsset
}

// GetBaseAssetOk returns a tuple with the BaseAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesLeadTradingSymbolWhitelistResponseDataInner) GetBaseAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.BaseAsset) {
		return nil, false
	}
	return o.BaseAsset, true
}

// HasBaseAsset returns a boolean if a field has been set.
func (o *GetFuturesLeadTradingSymbolWhitelistResponseDataInner) HasBaseAsset() bool {
	if o != nil && !common.IsNil(o.BaseAsset) {
		return true
	}

	return false
}

// SetBaseAsset gets a reference to the given string and assigns it to the BaseAsset field.
func (o *GetFuturesLeadTradingSymbolWhitelistResponseDataInner) SetBaseAsset(v string) {
	o.BaseAsset = &v
}

// GetQuoteAsset returns the QuoteAsset field value if set, zero value otherwise.
func (o *GetFuturesLeadTradingSymbolWhitelistResponseDataInner) GetQuoteAsset() string {
	if o == nil || common.IsNil(o.QuoteAsset) {
		var ret string
		return ret
	}
	return *o.QuoteAsset
}

// GetQuoteAssetOk returns a tuple with the QuoteAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesLeadTradingSymbolWhitelistResponseDataInner) GetQuoteAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.QuoteAsset) {
		return nil, false
	}
	return o.QuoteAsset, true
}

// HasQuoteAsset returns a boolean if a field has been set.
func (o *GetFuturesLeadTradingSymbolWhitelistResponseDataInner) HasQuoteAsset() bool {
	if o != nil && !common.IsNil(o.QuoteAsset) {
		return true
	}

	return false
}

// SetQuoteAsset gets a reference to the given string and assigns it to the QuoteAsset field.
func (o *GetFuturesLeadTradingSymbolWhitelistResponseDataInner) SetQuoteAsset(v string) {
	o.QuoteAsset = &v
}

func (o GetFuturesLeadTradingSymbolWhitelistResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFuturesLeadTradingSymbolWhitelistResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.BaseAsset) {
		toSerialize["baseAsset"] = o.BaseAsset
	}
	if !common.IsNil(o.QuoteAsset) {
		toSerialize["quoteAsset"] = o.QuoteAsset
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetFuturesLeadTradingSymbolWhitelistResponseDataInner) UnmarshalJSON(data []byte) (err error) {
	varGetFuturesLeadTradingSymbolWhitelistResponseDataInner := _GetFuturesLeadTradingSymbolWhitelistResponseDataInner{}

	err = json.Unmarshal(data, &varGetFuturesLeadTradingSymbolWhitelistResponseDataInner)

	if err != nil {
		return err
	}

	*o = GetFuturesLeadTradingSymbolWhitelistResponseDataInner(varGetFuturesLeadTradingSymbolWhitelistResponseDataInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "baseAsset")
		delete(additionalProperties, "quoteAsset")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFuturesLeadTradingSymbolWhitelistResponseDataInner struct {
	value *GetFuturesLeadTradingSymbolWhitelistResponseDataInner
	isSet bool
}

func (v NullableGetFuturesLeadTradingSymbolWhitelistResponseDataInner) Get() *GetFuturesLeadTradingSymbolWhitelistResponseDataInner {
	return v.value
}

func (v *NullableGetFuturesLeadTradingSymbolWhitelistResponseDataInner) Set(val *GetFuturesLeadTradingSymbolWhitelistResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFuturesLeadTradingSymbolWhitelistResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFuturesLeadTradingSymbolWhitelistResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFuturesLeadTradingSymbolWhitelistResponseDataInner(val *GetFuturesLeadTradingSymbolWhitelistResponseDataInner) *NullableGetFuturesLeadTradingSymbolWhitelistResponseDataInner {
	return &NullableGetFuturesLeadTradingSymbolWhitelistResponseDataInner{value: val, isSet: true}
}

func (v NullableGetFuturesLeadTradingSymbolWhitelistResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFuturesLeadTradingSymbolWhitelistResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
