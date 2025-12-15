/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the MyFiltersResponseResult type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MyFiltersResponseResult{}

// MyFiltersResponseResult struct for MyFiltersResponseResult
type MyFiltersResponseResult struct {
	ExchangeFilters      []ExchangeFilters `json:"exchangeFilters,omitempty"`
	SymbolFilters        []SymbolFilters   `json:"symbolFilters,omitempty"`
	AssetFilters         []AssetFilters    `json:"assetFilters,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MyFiltersResponseResult MyFiltersResponseResult

// NewMyFiltersResponseResult instantiates a new MyFiltersResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMyFiltersResponseResult() *MyFiltersResponseResult {
	this := MyFiltersResponseResult{}
	return &this
}

// NewMyFiltersResponseResultWithDefaults instantiates a new MyFiltersResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMyFiltersResponseResultWithDefaults() *MyFiltersResponseResult {
	this := MyFiltersResponseResult{}
	return &this
}

// GetExchangeFilters returns the ExchangeFilters field value if set, zero value otherwise.
func (o *MyFiltersResponseResult) GetExchangeFilters() []ExchangeFilters {
	if o == nil || common.IsNil(o.ExchangeFilters) {
		var ret []ExchangeFilters
		return ret
	}
	return o.ExchangeFilters
}

// GetExchangeFiltersOk returns a tuple with the ExchangeFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyFiltersResponseResult) GetExchangeFiltersOk() ([]ExchangeFilters, bool) {
	if o == nil || common.IsNil(o.ExchangeFilters) {
		return nil, false
	}
	return o.ExchangeFilters, true
}

// HasExchangeFilters returns a boolean if a field has been set.
func (o *MyFiltersResponseResult) HasExchangeFilters() bool {
	if o != nil && !common.IsNil(o.ExchangeFilters) {
		return true
	}

	return false
}

// SetExchangeFilters gets a reference to the given []ExchangeFilters and assigns it to the ExchangeFilters field.
func (o *MyFiltersResponseResult) SetExchangeFilters(v []ExchangeFilters) {
	o.ExchangeFilters = v
}

// GetSymbolFilters returns the SymbolFilters field value if set, zero value otherwise.
func (o *MyFiltersResponseResult) GetSymbolFilters() []SymbolFilters {
	if o == nil || common.IsNil(o.SymbolFilters) {
		var ret []SymbolFilters
		return ret
	}
	return o.SymbolFilters
}

// GetSymbolFiltersOk returns a tuple with the SymbolFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyFiltersResponseResult) GetSymbolFiltersOk() ([]SymbolFilters, bool) {
	if o == nil || common.IsNil(o.SymbolFilters) {
		return nil, false
	}
	return o.SymbolFilters, true
}

// HasSymbolFilters returns a boolean if a field has been set.
func (o *MyFiltersResponseResult) HasSymbolFilters() bool {
	if o != nil && !common.IsNil(o.SymbolFilters) {
		return true
	}

	return false
}

// SetSymbolFilters gets a reference to the given []SymbolFilters and assigns it to the SymbolFilters field.
func (o *MyFiltersResponseResult) SetSymbolFilters(v []SymbolFilters) {
	o.SymbolFilters = v
}

// GetAssetFilters returns the AssetFilters field value if set, zero value otherwise.
func (o *MyFiltersResponseResult) GetAssetFilters() []AssetFilters {
	if o == nil || common.IsNil(o.AssetFilters) {
		var ret []AssetFilters
		return ret
	}
	return o.AssetFilters
}

// GetAssetFiltersOk returns a tuple with the AssetFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyFiltersResponseResult) GetAssetFiltersOk() ([]AssetFilters, bool) {
	if o == nil || common.IsNil(o.AssetFilters) {
		return nil, false
	}
	return o.AssetFilters, true
}

// HasAssetFilters returns a boolean if a field has been set.
func (o *MyFiltersResponseResult) HasAssetFilters() bool {
	if o != nil && !common.IsNil(o.AssetFilters) {
		return true
	}

	return false
}

// SetAssetFilters gets a reference to the given []AssetFilters and assigns it to the AssetFilters field.
func (o *MyFiltersResponseResult) SetAssetFilters(v []AssetFilters) {
	o.AssetFilters = v
}

func (o MyFiltersResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MyFiltersResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ExchangeFilters) {
		toSerialize["exchangeFilters"] = o.ExchangeFilters
	}
	if !common.IsNil(o.SymbolFilters) {
		toSerialize["symbolFilters"] = o.SymbolFilters
	}
	if !common.IsNil(o.AssetFilters) {
		toSerialize["assetFilters"] = o.AssetFilters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MyFiltersResponseResult) UnmarshalJSON(data []byte) (err error) {
	varMyFiltersResponseResult := _MyFiltersResponseResult{}

	err = json.Unmarshal(data, &varMyFiltersResponseResult)

	if err != nil {
		return err
	}

	*o = MyFiltersResponseResult(varMyFiltersResponseResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "exchangeFilters")
		delete(additionalProperties, "symbolFilters")
		delete(additionalProperties, "assetFilters")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMyFiltersResponseResult struct {
	value *MyFiltersResponseResult
	isSet bool
}

func (v NullableMyFiltersResponseResult) Get() *MyFiltersResponseResult {
	return v.value
}

func (v *NullableMyFiltersResponseResult) Set(val *MyFiltersResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableMyFiltersResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableMyFiltersResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMyFiltersResponseResult(val *MyFiltersResponseResult) *NullableMyFiltersResponseResult {
	return &NullableMyFiltersResponseResult{value: val, isSet: true}
}

func (v NullableMyFiltersResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMyFiltersResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
