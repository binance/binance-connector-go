/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetAllIsolatedMarginSymbolResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetAllIsolatedMarginSymbolResponseInner{}

// GetAllIsolatedMarginSymbolResponseInner struct for GetAllIsolatedMarginSymbolResponseInner
type GetAllIsolatedMarginSymbolResponseInner struct {
	Base                 *string `json:"base,omitempty"`
	IsBuyAllowed         *bool   `json:"isBuyAllowed,omitempty"`
	IsMarginTrade        *bool   `json:"isMarginTrade,omitempty"`
	IsSellAllowed        *bool   `json:"isSellAllowed,omitempty"`
	Quote                *string `json:"quote,omitempty"`
	Symbol               *string `json:"symbol,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetAllIsolatedMarginSymbolResponseInner GetAllIsolatedMarginSymbolResponseInner

// NewGetAllIsolatedMarginSymbolResponseInner instantiates a new GetAllIsolatedMarginSymbolResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllIsolatedMarginSymbolResponseInner() *GetAllIsolatedMarginSymbolResponseInner {
	this := GetAllIsolatedMarginSymbolResponseInner{}
	return &this
}

// NewGetAllIsolatedMarginSymbolResponseInnerWithDefaults instantiates a new GetAllIsolatedMarginSymbolResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllIsolatedMarginSymbolResponseInnerWithDefaults() *GetAllIsolatedMarginSymbolResponseInner {
	this := GetAllIsolatedMarginSymbolResponseInner{}
	return &this
}

// GetBase returns the Base field value if set, zero value otherwise.
func (o *GetAllIsolatedMarginSymbolResponseInner) GetBase() string {
	if o == nil || common.IsNil(o.Base) {
		var ret string
		return ret
	}
	return *o.Base
}

// GetBaseOk returns a tuple with the Base field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllIsolatedMarginSymbolResponseInner) GetBaseOk() (*string, bool) {
	if o == nil || common.IsNil(o.Base) {
		return nil, false
	}
	return o.Base, true
}

// HasBase returns a boolean if a field has been set.
func (o *GetAllIsolatedMarginSymbolResponseInner) HasBase() bool {
	if o != nil && !common.IsNil(o.Base) {
		return true
	}

	return false
}

// SetBase gets a reference to the given string and assigns it to the Base field.
func (o *GetAllIsolatedMarginSymbolResponseInner) SetBase(v string) {
	o.Base = &v
}

// GetIsBuyAllowed returns the IsBuyAllowed field value if set, zero value otherwise.
func (o *GetAllIsolatedMarginSymbolResponseInner) GetIsBuyAllowed() bool {
	if o == nil || common.IsNil(o.IsBuyAllowed) {
		var ret bool
		return ret
	}
	return *o.IsBuyAllowed
}

// GetIsBuyAllowedOk returns a tuple with the IsBuyAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllIsolatedMarginSymbolResponseInner) GetIsBuyAllowedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsBuyAllowed) {
		return nil, false
	}
	return o.IsBuyAllowed, true
}

// HasIsBuyAllowed returns a boolean if a field has been set.
func (o *GetAllIsolatedMarginSymbolResponseInner) HasIsBuyAllowed() bool {
	if o != nil && !common.IsNil(o.IsBuyAllowed) {
		return true
	}

	return false
}

// SetIsBuyAllowed gets a reference to the given bool and assigns it to the IsBuyAllowed field.
func (o *GetAllIsolatedMarginSymbolResponseInner) SetIsBuyAllowed(v bool) {
	o.IsBuyAllowed = &v
}

// GetIsMarginTrade returns the IsMarginTrade field value if set, zero value otherwise.
func (o *GetAllIsolatedMarginSymbolResponseInner) GetIsMarginTrade() bool {
	if o == nil || common.IsNil(o.IsMarginTrade) {
		var ret bool
		return ret
	}
	return *o.IsMarginTrade
}

// GetIsMarginTradeOk returns a tuple with the IsMarginTrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllIsolatedMarginSymbolResponseInner) GetIsMarginTradeOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsMarginTrade) {
		return nil, false
	}
	return o.IsMarginTrade, true
}

// HasIsMarginTrade returns a boolean if a field has been set.
func (o *GetAllIsolatedMarginSymbolResponseInner) HasIsMarginTrade() bool {
	if o != nil && !common.IsNil(o.IsMarginTrade) {
		return true
	}

	return false
}

// SetIsMarginTrade gets a reference to the given bool and assigns it to the IsMarginTrade field.
func (o *GetAllIsolatedMarginSymbolResponseInner) SetIsMarginTrade(v bool) {
	o.IsMarginTrade = &v
}

// GetIsSellAllowed returns the IsSellAllowed field value if set, zero value otherwise.
func (o *GetAllIsolatedMarginSymbolResponseInner) GetIsSellAllowed() bool {
	if o == nil || common.IsNil(o.IsSellAllowed) {
		var ret bool
		return ret
	}
	return *o.IsSellAllowed
}

// GetIsSellAllowedOk returns a tuple with the IsSellAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllIsolatedMarginSymbolResponseInner) GetIsSellAllowedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsSellAllowed) {
		return nil, false
	}
	return o.IsSellAllowed, true
}

// HasIsSellAllowed returns a boolean if a field has been set.
func (o *GetAllIsolatedMarginSymbolResponseInner) HasIsSellAllowed() bool {
	if o != nil && !common.IsNil(o.IsSellAllowed) {
		return true
	}

	return false
}

// SetIsSellAllowed gets a reference to the given bool and assigns it to the IsSellAllowed field.
func (o *GetAllIsolatedMarginSymbolResponseInner) SetIsSellAllowed(v bool) {
	o.IsSellAllowed = &v
}

// GetQuote returns the Quote field value if set, zero value otherwise.
func (o *GetAllIsolatedMarginSymbolResponseInner) GetQuote() string {
	if o == nil || common.IsNil(o.Quote) {
		var ret string
		return ret
	}
	return *o.Quote
}

// GetQuoteOk returns a tuple with the Quote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllIsolatedMarginSymbolResponseInner) GetQuoteOk() (*string, bool) {
	if o == nil || common.IsNil(o.Quote) {
		return nil, false
	}
	return o.Quote, true
}

// HasQuote returns a boolean if a field has been set.
func (o *GetAllIsolatedMarginSymbolResponseInner) HasQuote() bool {
	if o != nil && !common.IsNil(o.Quote) {
		return true
	}

	return false
}

// SetQuote gets a reference to the given string and assigns it to the Quote field.
func (o *GetAllIsolatedMarginSymbolResponseInner) SetQuote(v string) {
	o.Quote = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetAllIsolatedMarginSymbolResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllIsolatedMarginSymbolResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetAllIsolatedMarginSymbolResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetAllIsolatedMarginSymbolResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

func (o GetAllIsolatedMarginSymbolResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllIsolatedMarginSymbolResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Base) {
		toSerialize["base"] = o.Base
	}
	if !common.IsNil(o.IsBuyAllowed) {
		toSerialize["isBuyAllowed"] = o.IsBuyAllowed
	}
	if !common.IsNil(o.IsMarginTrade) {
		toSerialize["isMarginTrade"] = o.IsMarginTrade
	}
	if !common.IsNil(o.IsSellAllowed) {
		toSerialize["isSellAllowed"] = o.IsSellAllowed
	}
	if !common.IsNil(o.Quote) {
		toSerialize["quote"] = o.Quote
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetAllIsolatedMarginSymbolResponseInner) UnmarshalJSON(data []byte) (err error) {
	varGetAllIsolatedMarginSymbolResponseInner := _GetAllIsolatedMarginSymbolResponseInner{}

	err = json.Unmarshal(data, &varGetAllIsolatedMarginSymbolResponseInner)

	if err != nil {
		return err
	}

	*o = GetAllIsolatedMarginSymbolResponseInner(varGetAllIsolatedMarginSymbolResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "base")
		delete(additionalProperties, "isBuyAllowed")
		delete(additionalProperties, "isMarginTrade")
		delete(additionalProperties, "isSellAllowed")
		delete(additionalProperties, "quote")
		delete(additionalProperties, "symbol")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetAllIsolatedMarginSymbolResponseInner struct {
	value *GetAllIsolatedMarginSymbolResponseInner
	isSet bool
}

func (v NullableGetAllIsolatedMarginSymbolResponseInner) Get() *GetAllIsolatedMarginSymbolResponseInner {
	return v.value
}

func (v *NullableGetAllIsolatedMarginSymbolResponseInner) Set(val *GetAllIsolatedMarginSymbolResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllIsolatedMarginSymbolResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllIsolatedMarginSymbolResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllIsolatedMarginSymbolResponseInner(val *GetAllIsolatedMarginSymbolResponseInner) *NullableGetAllIsolatedMarginSymbolResponseInner {
	return &NullableGetAllIsolatedMarginSymbolResponseInner{value: val, isSet: true}
}

func (v NullableGetAllIsolatedMarginSymbolResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllIsolatedMarginSymbolResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
