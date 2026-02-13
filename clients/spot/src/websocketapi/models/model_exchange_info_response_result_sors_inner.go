/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ExchangeInfoResponseResultSorsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ExchangeInfoResponseResultSorsInner{}

// ExchangeInfoResponseResultSorsInner struct for ExchangeInfoResponseResultSorsInner
type ExchangeInfoResponseResultSorsInner struct {
	BaseAsset            *string  `json:"baseAsset,omitempty"`
	Symbols              []string `json:"symbols,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExchangeInfoResponseResultSorsInner ExchangeInfoResponseResultSorsInner

// NewExchangeInfoResponseResultSorsInner instantiates a new ExchangeInfoResponseResultSorsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeInfoResponseResultSorsInner() *ExchangeInfoResponseResultSorsInner {
	this := ExchangeInfoResponseResultSorsInner{}
	return &this
}

// NewExchangeInfoResponseResultSorsInnerWithDefaults instantiates a new ExchangeInfoResponseResultSorsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeInfoResponseResultSorsInnerWithDefaults() *ExchangeInfoResponseResultSorsInner {
	this := ExchangeInfoResponseResultSorsInner{}
	return &this
}

// GetBaseAsset returns the BaseAsset field value if set, zero value otherwise.
func (o *ExchangeInfoResponseResultSorsInner) GetBaseAsset() string {
	if o == nil || common.IsNil(o.BaseAsset) {
		var ret string
		return ret
	}
	return *o.BaseAsset
}

// GetBaseAssetOk returns a tuple with the BaseAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseResultSorsInner) GetBaseAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.BaseAsset) {
		return nil, false
	}
	return o.BaseAsset, true
}

// HasBaseAsset returns a boolean if a field has been set.
func (o *ExchangeInfoResponseResultSorsInner) HasBaseAsset() bool {
	if o != nil && !common.IsNil(o.BaseAsset) {
		return true
	}

	return false
}

// SetBaseAsset gets a reference to the given string and assigns it to the BaseAsset field.
func (o *ExchangeInfoResponseResultSorsInner) SetBaseAsset(v string) {
	o.BaseAsset = &v
}

// GetSymbols returns the Symbols field value if set, zero value otherwise.
func (o *ExchangeInfoResponseResultSorsInner) GetSymbols() []string {
	if o == nil || common.IsNil(o.Symbols) {
		var ret []string
		return ret
	}
	return o.Symbols
}

// GetSymbolsOk returns a tuple with the Symbols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseResultSorsInner) GetSymbolsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Symbols) {
		return nil, false
	}
	return o.Symbols, true
}

// HasSymbols returns a boolean if a field has been set.
func (o *ExchangeInfoResponseResultSorsInner) HasSymbols() bool {
	if o != nil && !common.IsNil(o.Symbols) {
		return true
	}

	return false
}

// SetSymbols gets a reference to the given []string and assigns it to the Symbols field.
func (o *ExchangeInfoResponseResultSorsInner) SetSymbols(v []string) {
	o.Symbols = v
}

func (o ExchangeInfoResponseResultSorsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeInfoResponseResultSorsInner) ToMap() (map[string]interface{}, error) {
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

func (o *ExchangeInfoResponseResultSorsInner) UnmarshalJSON(data []byte) (err error) {
	varExchangeInfoResponseResultSorsInner := _ExchangeInfoResponseResultSorsInner{}

	err = json.Unmarshal(data, &varExchangeInfoResponseResultSorsInner)

	if err != nil {
		return err
	}

	*o = ExchangeInfoResponseResultSorsInner(varExchangeInfoResponseResultSorsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "baseAsset")
		delete(additionalProperties, "symbols")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExchangeInfoResponseResultSorsInner struct {
	value *ExchangeInfoResponseResultSorsInner
	isSet bool
}

func (v NullableExchangeInfoResponseResultSorsInner) Get() *ExchangeInfoResponseResultSorsInner {
	return v.value
}

func (v *NullableExchangeInfoResponseResultSorsInner) Set(val *ExchangeInfoResponseResultSorsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeInfoResponseResultSorsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeInfoResponseResultSorsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeInfoResponseResultSorsInner(val *ExchangeInfoResponseResultSorsInner) *NullableExchangeInfoResponseResultSorsInner {
	return &NullableExchangeInfoResponseResultSorsInner{value: val, isSet: true}
}

func (v NullableExchangeInfoResponseResultSorsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeInfoResponseResultSorsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
