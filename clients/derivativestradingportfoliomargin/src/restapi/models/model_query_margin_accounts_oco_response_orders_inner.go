/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryMarginAccountsOcoResponseOrdersInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryMarginAccountsOcoResponseOrdersInner{}

// QueryMarginAccountsOcoResponseOrdersInner struct for QueryMarginAccountsOcoResponseOrdersInner
type QueryMarginAccountsOcoResponseOrdersInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	OrderId              *int64  `json:"orderId,omitempty"`
	ClientOrderId        *string `json:"clientOrderId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryMarginAccountsOcoResponseOrdersInner QueryMarginAccountsOcoResponseOrdersInner

// NewQueryMarginAccountsOcoResponseOrdersInner instantiates a new QueryMarginAccountsOcoResponseOrdersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryMarginAccountsOcoResponseOrdersInner() *QueryMarginAccountsOcoResponseOrdersInner {
	this := QueryMarginAccountsOcoResponseOrdersInner{}
	return &this
}

// NewQueryMarginAccountsOcoResponseOrdersInnerWithDefaults instantiates a new QueryMarginAccountsOcoResponseOrdersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryMarginAccountsOcoResponseOrdersInnerWithDefaults() *QueryMarginAccountsOcoResponseOrdersInner {
	this := QueryMarginAccountsOcoResponseOrdersInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *QueryMarginAccountsOcoResponseOrdersInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsOcoResponseOrdersInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *QueryMarginAccountsOcoResponseOrdersInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *QueryMarginAccountsOcoResponseOrdersInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *QueryMarginAccountsOcoResponseOrdersInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsOcoResponseOrdersInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *QueryMarginAccountsOcoResponseOrdersInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *QueryMarginAccountsOcoResponseOrdersInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *QueryMarginAccountsOcoResponseOrdersInner) GetClientOrderId() string {
	if o == nil || common.IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsOcoResponseOrdersInner) GetClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *QueryMarginAccountsOcoResponseOrdersInner) HasClientOrderId() bool {
	if o != nil && !common.IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *QueryMarginAccountsOcoResponseOrdersInner) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

func (o QueryMarginAccountsOcoResponseOrdersInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryMarginAccountsOcoResponseOrdersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.ClientOrderId) {
		toSerialize["clientOrderId"] = o.ClientOrderId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryMarginAccountsOcoResponseOrdersInner) UnmarshalJSON(data []byte) (err error) {
	varQueryMarginAccountsOcoResponseOrdersInner := _QueryMarginAccountsOcoResponseOrdersInner{}

	err = json.Unmarshal(data, &varQueryMarginAccountsOcoResponseOrdersInner)

	if err != nil {
		return err
	}

	*o = QueryMarginAccountsOcoResponseOrdersInner(varQueryMarginAccountsOcoResponseOrdersInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "clientOrderId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryMarginAccountsOcoResponseOrdersInner struct {
	value *QueryMarginAccountsOcoResponseOrdersInner
	isSet bool
}

func (v NullableQueryMarginAccountsOcoResponseOrdersInner) Get() *QueryMarginAccountsOcoResponseOrdersInner {
	return v.value
}

func (v *NullableQueryMarginAccountsOcoResponseOrdersInner) Set(val *QueryMarginAccountsOcoResponseOrdersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryMarginAccountsOcoResponseOrdersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryMarginAccountsOcoResponseOrdersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryMarginAccountsOcoResponseOrdersInner(val *QueryMarginAccountsOcoResponseOrdersInner) *NullableQueryMarginAccountsOcoResponseOrdersInner {
	return &NullableQueryMarginAccountsOcoResponseOrdersInner{value: val, isSet: true}
}

func (v NullableQueryMarginAccountsOcoResponseOrdersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryMarginAccountsOcoResponseOrdersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
