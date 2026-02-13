/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the TradeFeeResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TradeFeeResponseInner{}

// TradeFeeResponseInner struct for TradeFeeResponseInner
type TradeFeeResponseInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	MakerCommission      *string `json:"makerCommission,omitempty"`
	TakerCommission      *string `json:"takerCommission,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TradeFeeResponseInner TradeFeeResponseInner

// NewTradeFeeResponseInner instantiates a new TradeFeeResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTradeFeeResponseInner() *TradeFeeResponseInner {
	this := TradeFeeResponseInner{}
	return &this
}

// NewTradeFeeResponseInnerWithDefaults instantiates a new TradeFeeResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTradeFeeResponseInnerWithDefaults() *TradeFeeResponseInner {
	this := TradeFeeResponseInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *TradeFeeResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeFeeResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *TradeFeeResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *TradeFeeResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetMakerCommission returns the MakerCommission field value if set, zero value otherwise.
func (o *TradeFeeResponseInner) GetMakerCommission() string {
	if o == nil || common.IsNil(o.MakerCommission) {
		var ret string
		return ret
	}
	return *o.MakerCommission
}

// GetMakerCommissionOk returns a tuple with the MakerCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeFeeResponseInner) GetMakerCommissionOk() (*string, bool) {
	if o == nil || common.IsNil(o.MakerCommission) {
		return nil, false
	}
	return o.MakerCommission, true
}

// HasMakerCommission returns a boolean if a field has been set.
func (o *TradeFeeResponseInner) HasMakerCommission() bool {
	if o != nil && !common.IsNil(o.MakerCommission) {
		return true
	}

	return false
}

// SetMakerCommission gets a reference to the given string and assigns it to the MakerCommission field.
func (o *TradeFeeResponseInner) SetMakerCommission(v string) {
	o.MakerCommission = &v
}

// GetTakerCommission returns the TakerCommission field value if set, zero value otherwise.
func (o *TradeFeeResponseInner) GetTakerCommission() string {
	if o == nil || common.IsNil(o.TakerCommission) {
		var ret string
		return ret
	}
	return *o.TakerCommission
}

// GetTakerCommissionOk returns a tuple with the TakerCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeFeeResponseInner) GetTakerCommissionOk() (*string, bool) {
	if o == nil || common.IsNil(o.TakerCommission) {
		return nil, false
	}
	return o.TakerCommission, true
}

// HasTakerCommission returns a boolean if a field has been set.
func (o *TradeFeeResponseInner) HasTakerCommission() bool {
	if o != nil && !common.IsNil(o.TakerCommission) {
		return true
	}

	return false
}

// SetTakerCommission gets a reference to the given string and assigns it to the TakerCommission field.
func (o *TradeFeeResponseInner) SetTakerCommission(v string) {
	o.TakerCommission = &v
}

func (o TradeFeeResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TradeFeeResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.MakerCommission) {
		toSerialize["makerCommission"] = o.MakerCommission
	}
	if !common.IsNil(o.TakerCommission) {
		toSerialize["takerCommission"] = o.TakerCommission
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TradeFeeResponseInner) UnmarshalJSON(data []byte) (err error) {
	varTradeFeeResponseInner := _TradeFeeResponseInner{}

	err = json.Unmarshal(data, &varTradeFeeResponseInner)

	if err != nil {
		return err
	}

	*o = TradeFeeResponseInner(varTradeFeeResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "makerCommission")
		delete(additionalProperties, "takerCommission")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTradeFeeResponseInner struct {
	value *TradeFeeResponseInner
	isSet bool
}

func (v NullableTradeFeeResponseInner) Get() *TradeFeeResponseInner {
	return v.value
}

func (v *NullableTradeFeeResponseInner) Set(val *TradeFeeResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTradeFeeResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTradeFeeResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTradeFeeResponseInner(val *TradeFeeResponseInner) *NullableTradeFeeResponseInner {
	return &NullableTradeFeeResponseInner{value: val, isSet: true}
}

func (v NullableTradeFeeResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTradeFeeResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
