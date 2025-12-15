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

// checks if the QueryMarginPriceindexResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryMarginPriceindexResponse{}

// QueryMarginPriceindexResponse struct for QueryMarginPriceindexResponse
type QueryMarginPriceindexResponse struct {
	CalcTime             *int64  `json:"calcTime,omitempty"`
	Price                *string `json:"price,omitempty"`
	Symbol               *string `json:"symbol,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryMarginPriceindexResponse QueryMarginPriceindexResponse

// NewQueryMarginPriceindexResponse instantiates a new QueryMarginPriceindexResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryMarginPriceindexResponse() *QueryMarginPriceindexResponse {
	this := QueryMarginPriceindexResponse{}
	return &this
}

// NewQueryMarginPriceindexResponseWithDefaults instantiates a new QueryMarginPriceindexResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryMarginPriceindexResponseWithDefaults() *QueryMarginPriceindexResponse {
	this := QueryMarginPriceindexResponse{}
	return &this
}

// GetCalcTime returns the CalcTime field value if set, zero value otherwise.
func (o *QueryMarginPriceindexResponse) GetCalcTime() int64 {
	if o == nil || common.IsNil(o.CalcTime) {
		var ret int64
		return ret
	}
	return *o.CalcTime
}

// GetCalcTimeOk returns a tuple with the CalcTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginPriceindexResponse) GetCalcTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CalcTime) {
		return nil, false
	}
	return o.CalcTime, true
}

// HasCalcTime returns a boolean if a field has been set.
func (o *QueryMarginPriceindexResponse) HasCalcTime() bool {
	if o != nil && !common.IsNil(o.CalcTime) {
		return true
	}

	return false
}

// SetCalcTime gets a reference to the given int64 and assigns it to the CalcTime field.
func (o *QueryMarginPriceindexResponse) SetCalcTime(v int64) {
	o.CalcTime = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *QueryMarginPriceindexResponse) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginPriceindexResponse) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *QueryMarginPriceindexResponse) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *QueryMarginPriceindexResponse) SetPrice(v string) {
	o.Price = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *QueryMarginPriceindexResponse) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginPriceindexResponse) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *QueryMarginPriceindexResponse) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *QueryMarginPriceindexResponse) SetSymbol(v string) {
	o.Symbol = &v
}

func (o QueryMarginPriceindexResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryMarginPriceindexResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CalcTime) {
		toSerialize["calcTime"] = o.CalcTime
	}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryMarginPriceindexResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryMarginPriceindexResponse := _QueryMarginPriceindexResponse{}

	err = json.Unmarshal(data, &varQueryMarginPriceindexResponse)

	if err != nil {
		return err
	}

	*o = QueryMarginPriceindexResponse(varQueryMarginPriceindexResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "calcTime")
		delete(additionalProperties, "price")
		delete(additionalProperties, "symbol")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryMarginPriceindexResponse struct {
	value *QueryMarginPriceindexResponse
	isSet bool
}

func (v NullableQueryMarginPriceindexResponse) Get() *QueryMarginPriceindexResponse {
	return v.value
}

func (v *NullableQueryMarginPriceindexResponse) Set(val *QueryMarginPriceindexResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryMarginPriceindexResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryMarginPriceindexResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryMarginPriceindexResponse(val *QueryMarginPriceindexResponse) *NullableQueryMarginPriceindexResponse {
	return &NullableQueryMarginPriceindexResponse{value: val, isSet: true}
}

func (v NullableQueryMarginPriceindexResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryMarginPriceindexResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
