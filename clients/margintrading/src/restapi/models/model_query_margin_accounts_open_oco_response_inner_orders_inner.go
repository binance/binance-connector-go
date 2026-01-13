/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryMarginAccountsOpenOcoResponseInnerOrdersInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryMarginAccountsOpenOcoResponseInnerOrdersInner{}

// QueryMarginAccountsOpenOcoResponseInnerOrdersInner struct for QueryMarginAccountsOpenOcoResponseInnerOrdersInner
type QueryMarginAccountsOpenOcoResponseInnerOrdersInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	OrderId              *int64  `json:"orderId,omitempty"`
	ClientOrderId        *string `json:"clientOrderId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryMarginAccountsOpenOcoResponseInnerOrdersInner QueryMarginAccountsOpenOcoResponseInnerOrdersInner

// NewQueryMarginAccountsOpenOcoResponseInnerOrdersInner instantiates a new QueryMarginAccountsOpenOcoResponseInnerOrdersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryMarginAccountsOpenOcoResponseInnerOrdersInner() *QueryMarginAccountsOpenOcoResponseInnerOrdersInner {
	this := QueryMarginAccountsOpenOcoResponseInnerOrdersInner{}
	return &this
}

// NewQueryMarginAccountsOpenOcoResponseInnerOrdersInnerWithDefaults instantiates a new QueryMarginAccountsOpenOcoResponseInnerOrdersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryMarginAccountsOpenOcoResponseInnerOrdersInnerWithDefaults() *QueryMarginAccountsOpenOcoResponseInnerOrdersInner {
	this := QueryMarginAccountsOpenOcoResponseInnerOrdersInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *QueryMarginAccountsOpenOcoResponseInnerOrdersInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsOpenOcoResponseInnerOrdersInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *QueryMarginAccountsOpenOcoResponseInnerOrdersInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *QueryMarginAccountsOpenOcoResponseInnerOrdersInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *QueryMarginAccountsOpenOcoResponseInnerOrdersInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsOpenOcoResponseInnerOrdersInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *QueryMarginAccountsOpenOcoResponseInnerOrdersInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *QueryMarginAccountsOpenOcoResponseInnerOrdersInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *QueryMarginAccountsOpenOcoResponseInnerOrdersInner) GetClientOrderId() string {
	if o == nil || common.IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsOpenOcoResponseInnerOrdersInner) GetClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *QueryMarginAccountsOpenOcoResponseInnerOrdersInner) HasClientOrderId() bool {
	if o != nil && !common.IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *QueryMarginAccountsOpenOcoResponseInnerOrdersInner) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

func (o QueryMarginAccountsOpenOcoResponseInnerOrdersInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryMarginAccountsOpenOcoResponseInnerOrdersInner) ToMap() (map[string]interface{}, error) {
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

func (o *QueryMarginAccountsOpenOcoResponseInnerOrdersInner) UnmarshalJSON(data []byte) (err error) {
	varQueryMarginAccountsOpenOcoResponseInnerOrdersInner := _QueryMarginAccountsOpenOcoResponseInnerOrdersInner{}

	err = json.Unmarshal(data, &varQueryMarginAccountsOpenOcoResponseInnerOrdersInner)

	if err != nil {
		return err
	}

	*o = QueryMarginAccountsOpenOcoResponseInnerOrdersInner(varQueryMarginAccountsOpenOcoResponseInnerOrdersInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "clientOrderId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryMarginAccountsOpenOcoResponseInnerOrdersInner struct {
	value *QueryMarginAccountsOpenOcoResponseInnerOrdersInner
	isSet bool
}

func (v NullableQueryMarginAccountsOpenOcoResponseInnerOrdersInner) Get() *QueryMarginAccountsOpenOcoResponseInnerOrdersInner {
	return v.value
}

func (v *NullableQueryMarginAccountsOpenOcoResponseInnerOrdersInner) Set(val *QueryMarginAccountsOpenOcoResponseInnerOrdersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryMarginAccountsOpenOcoResponseInnerOrdersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryMarginAccountsOpenOcoResponseInnerOrdersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryMarginAccountsOpenOcoResponseInnerOrdersInner(val *QueryMarginAccountsOpenOcoResponseInnerOrdersInner) *NullableQueryMarginAccountsOpenOcoResponseInnerOrdersInner {
	return &NullableQueryMarginAccountsOpenOcoResponseInnerOrdersInner{value: val, isSet: true}
}

func (v NullableQueryMarginAccountsOpenOcoResponseInnerOrdersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryMarginAccountsOpenOcoResponseInnerOrdersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
