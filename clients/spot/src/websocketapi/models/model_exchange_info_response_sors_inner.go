/*
Spot WebSocket API

Access market data, manage accounts, and trade on Binance Spot.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ExchangeInfoResponseSorsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ExchangeInfoResponseSorsInner{}

// ExchangeInfoResponseSorsInner struct for ExchangeInfoResponseSorsInner
type ExchangeInfoResponseSorsInner struct {
	BaseAsset            *string  `json:"baseAsset,omitempty"`
	Symbols              []string `json:"symbols,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExchangeInfoResponseSorsInner ExchangeInfoResponseSorsInner

// NewExchangeInfoResponseSorsInner instantiates a new ExchangeInfoResponseSorsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeInfoResponseSorsInner() *ExchangeInfoResponseSorsInner {
	this := ExchangeInfoResponseSorsInner{}
	return &this
}

// NewExchangeInfoResponseSorsInnerWithDefaults instantiates a new ExchangeInfoResponseSorsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeInfoResponseSorsInnerWithDefaults() *ExchangeInfoResponseSorsInner {
	this := ExchangeInfoResponseSorsInner{}
	return &this
}

// GetBaseAsset returns the BaseAsset field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSorsInner) GetBaseAsset() string {
	if o == nil || common.IsNil(o.BaseAsset) {
		var ret string
		return ret
	}
	return *o.BaseAsset
}

// GetBaseAssetOk returns a tuple with the BaseAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSorsInner) GetBaseAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.BaseAsset) {
		return nil, false
	}
	return o.BaseAsset, true
}

// HasBaseAsset returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSorsInner) HasBaseAsset() bool {
	if o != nil && !common.IsNil(o.BaseAsset) {
		return true
	}

	return false
}

// SetBaseAsset gets a reference to the given string and assigns it to the BaseAsset field.
func (o *ExchangeInfoResponseSorsInner) SetBaseAsset(v string) {
	o.BaseAsset = &v
}

// GetSymbols returns the Symbols field value if set, zero value otherwise.
func (o *ExchangeInfoResponseSorsInner) GetSymbols() []string {
	if o == nil || common.IsNil(o.Symbols) {
		var ret []string
		return ret
	}
	return o.Symbols
}

// GetSymbolsOk returns a tuple with the Symbols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseSorsInner) GetSymbolsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Symbols) {
		return nil, false
	}
	return o.Symbols, true
}

// HasSymbols returns a boolean if a field has been set.
func (o *ExchangeInfoResponseSorsInner) HasSymbols() bool {
	if o != nil && !common.IsNil(o.Symbols) {
		return true
	}

	return false
}

// SetSymbols gets a reference to the given []string and assigns it to the Symbols field.
func (o *ExchangeInfoResponseSorsInner) SetSymbols(v []string) {
	o.Symbols = v
}

func (o ExchangeInfoResponseSorsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeInfoResponseSorsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BaseAsset) {
		toSerialize["baseAsset"] = o.BaseAsset
	}
	if !common.IsNil(o.Symbols) {
		toSerialize["symbols"] = o.Symbols
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExchangeInfoResponseSorsInner) UnmarshalJSON(data []byte) (err error) {
	varExchangeInfoResponseSorsInner := _ExchangeInfoResponseSorsInner{}

	err = json.Unmarshal(data, &varExchangeInfoResponseSorsInner)

	if err != nil {
		return err
	}

	*o = ExchangeInfoResponseSorsInner(varExchangeInfoResponseSorsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "baseAsset")
		delete(additionalProperties, "symbols")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExchangeInfoResponseSorsInner struct {
	value *ExchangeInfoResponseSorsInner
	isSet bool
}

func (v NullableExchangeInfoResponseSorsInner) Get() *ExchangeInfoResponseSorsInner {
	return v.value
}

func (v *NullableExchangeInfoResponseSorsInner) Set(val *ExchangeInfoResponseSorsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeInfoResponseSorsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeInfoResponseSorsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeInfoResponseSorsInner(val *ExchangeInfoResponseSorsInner) *NullableExchangeInfoResponseSorsInner {
	return &NullableExchangeInfoResponseSorsInner{value: val, isSet: true}
}

func (v NullableExchangeInfoResponseSorsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeInfoResponseSorsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
