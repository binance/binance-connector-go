/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryIndexPriceConstituentsResponseConstituentsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryIndexPriceConstituentsResponseConstituentsInner{}

// QueryIndexPriceConstituentsResponseConstituentsInner struct for QueryIndexPriceConstituentsResponseConstituentsInner
type QueryIndexPriceConstituentsResponseConstituentsInner struct {
	Exchange             *string `json:"exchange,omitempty"`
	Symbol               *string `json:"symbol,omitempty"`
	Price                *string `json:"price,omitempty"`
	Weight               *string `json:"weight,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryIndexPriceConstituentsResponseConstituentsInner QueryIndexPriceConstituentsResponseConstituentsInner

// NewQueryIndexPriceConstituentsResponseConstituentsInner instantiates a new QueryIndexPriceConstituentsResponseConstituentsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryIndexPriceConstituentsResponseConstituentsInner() *QueryIndexPriceConstituentsResponseConstituentsInner {
	this := QueryIndexPriceConstituentsResponseConstituentsInner{}
	return &this
}

// NewQueryIndexPriceConstituentsResponseConstituentsInnerWithDefaults instantiates a new QueryIndexPriceConstituentsResponseConstituentsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryIndexPriceConstituentsResponseConstituentsInnerWithDefaults() *QueryIndexPriceConstituentsResponseConstituentsInner {
	this := QueryIndexPriceConstituentsResponseConstituentsInner{}
	return &this
}

// GetExchange returns the Exchange field value if set, zero value otherwise.
func (o *QueryIndexPriceConstituentsResponseConstituentsInner) GetExchange() string {
	if o == nil || common.IsNil(o.Exchange) {
		var ret string
		return ret
	}
	return *o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIndexPriceConstituentsResponseConstituentsInner) GetExchangeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Exchange) {
		return nil, false
	}
	return o.Exchange, true
}

// HasExchange returns a boolean if a field has been set.
func (o *QueryIndexPriceConstituentsResponseConstituentsInner) HasExchange() bool {
	if o != nil && !common.IsNil(o.Exchange) {
		return true
	}

	return false
}

// SetExchange gets a reference to the given string and assigns it to the Exchange field.
func (o *QueryIndexPriceConstituentsResponseConstituentsInner) SetExchange(v string) {
	o.Exchange = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *QueryIndexPriceConstituentsResponseConstituentsInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIndexPriceConstituentsResponseConstituentsInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *QueryIndexPriceConstituentsResponseConstituentsInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *QueryIndexPriceConstituentsResponseConstituentsInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *QueryIndexPriceConstituentsResponseConstituentsInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIndexPriceConstituentsResponseConstituentsInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *QueryIndexPriceConstituentsResponseConstituentsInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *QueryIndexPriceConstituentsResponseConstituentsInner) SetPrice(v string) {
	o.Price = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *QueryIndexPriceConstituentsResponseConstituentsInner) GetWeight() string {
	if o == nil || common.IsNil(o.Weight) {
		var ret string
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIndexPriceConstituentsResponseConstituentsInner) GetWeightOk() (*string, bool) {
	if o == nil || common.IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *QueryIndexPriceConstituentsResponseConstituentsInner) HasWeight() bool {
	if o != nil && !common.IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given string and assigns it to the Weight field.
func (o *QueryIndexPriceConstituentsResponseConstituentsInner) SetWeight(v string) {
	o.Weight = &v
}

func (o QueryIndexPriceConstituentsResponseConstituentsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryIndexPriceConstituentsResponseConstituentsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Exchange) {
		toSerialize["exchange"] = o.Exchange
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryIndexPriceConstituentsResponseConstituentsInner) UnmarshalJSON(data []byte) (err error) {
	varQueryIndexPriceConstituentsResponseConstituentsInner := _QueryIndexPriceConstituentsResponseConstituentsInner{}

	err = json.Unmarshal(data, &varQueryIndexPriceConstituentsResponseConstituentsInner)

	if err != nil {
		return err
	}

	*o = QueryIndexPriceConstituentsResponseConstituentsInner(varQueryIndexPriceConstituentsResponseConstituentsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "exchange")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "price")
		delete(additionalProperties, "weight")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryIndexPriceConstituentsResponseConstituentsInner struct {
	value *QueryIndexPriceConstituentsResponseConstituentsInner
	isSet bool
}

func (v NullableQueryIndexPriceConstituentsResponseConstituentsInner) Get() *QueryIndexPriceConstituentsResponseConstituentsInner {
	return v.value
}

func (v *NullableQueryIndexPriceConstituentsResponseConstituentsInner) Set(val *QueryIndexPriceConstituentsResponseConstituentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryIndexPriceConstituentsResponseConstituentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryIndexPriceConstituentsResponseConstituentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryIndexPriceConstituentsResponseConstituentsInner(val *QueryIndexPriceConstituentsResponseConstituentsInner) *NullableQueryIndexPriceConstituentsResponseConstituentsInner {
	return &NullableQueryIndexPriceConstituentsResponseConstituentsInner{value: val, isSet: true}
}

func (v NullableQueryIndexPriceConstituentsResponseConstituentsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryIndexPriceConstituentsResponseConstituentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
