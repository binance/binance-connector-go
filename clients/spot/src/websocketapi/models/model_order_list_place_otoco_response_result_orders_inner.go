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

// checks if the OrderListPlaceOtocoResponseResultOrdersInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderListPlaceOtocoResponseResultOrdersInner{}

// OrderListPlaceOtocoResponseResultOrdersInner struct for OrderListPlaceOtocoResponseResultOrdersInner
type OrderListPlaceOtocoResponseResultOrdersInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	OrderId              *int64  `json:"orderId,omitempty"`
	ClientOrderId        *string `json:"clientOrderId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderListPlaceOtocoResponseResultOrdersInner OrderListPlaceOtocoResponseResultOrdersInner

// NewOrderListPlaceOtocoResponseResultOrdersInner instantiates a new OrderListPlaceOtocoResponseResultOrdersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderListPlaceOtocoResponseResultOrdersInner() *OrderListPlaceOtocoResponseResultOrdersInner {
	this := OrderListPlaceOtocoResponseResultOrdersInner{}
	return &this
}

// NewOrderListPlaceOtocoResponseResultOrdersInnerWithDefaults instantiates a new OrderListPlaceOtocoResponseResultOrdersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderListPlaceOtocoResponseResultOrdersInnerWithDefaults() *OrderListPlaceOtocoResponseResultOrdersInner {
	this := OrderListPlaceOtocoResponseResultOrdersInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *OrderListPlaceOtocoResponseResultOrdersInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOtocoResponseResultOrdersInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *OrderListPlaceOtocoResponseResultOrdersInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *OrderListPlaceOtocoResponseResultOrdersInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *OrderListPlaceOtocoResponseResultOrdersInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOtocoResponseResultOrdersInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *OrderListPlaceOtocoResponseResultOrdersInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *OrderListPlaceOtocoResponseResultOrdersInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *OrderListPlaceOtocoResponseResultOrdersInner) GetClientOrderId() string {
	if o == nil || common.IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOtocoResponseResultOrdersInner) GetClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *OrderListPlaceOtocoResponseResultOrdersInner) HasClientOrderId() bool {
	if o != nil && !common.IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *OrderListPlaceOtocoResponseResultOrdersInner) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

func (o OrderListPlaceOtocoResponseResultOrdersInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderListPlaceOtocoResponseResultOrdersInner) ToMap() (map[string]interface{}, error) {
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

func (o *OrderListPlaceOtocoResponseResultOrdersInner) UnmarshalJSON(data []byte) (err error) {
	varOrderListPlaceOtocoResponseResultOrdersInner := _OrderListPlaceOtocoResponseResultOrdersInner{}

	err = json.Unmarshal(data, &varOrderListPlaceOtocoResponseResultOrdersInner)

	if err != nil {
		return err
	}

	*o = OrderListPlaceOtocoResponseResultOrdersInner(varOrderListPlaceOtocoResponseResultOrdersInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "clientOrderId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrderListPlaceOtocoResponseResultOrdersInner struct {
	value *OrderListPlaceOtocoResponseResultOrdersInner
	isSet bool
}

func (v NullableOrderListPlaceOtocoResponseResultOrdersInner) Get() *OrderListPlaceOtocoResponseResultOrdersInner {
	return v.value
}

func (v *NullableOrderListPlaceOtocoResponseResultOrdersInner) Set(val *OrderListPlaceOtocoResponseResultOrdersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderListPlaceOtocoResponseResultOrdersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderListPlaceOtocoResponseResultOrdersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderListPlaceOtocoResponseResultOrdersInner(val *OrderListPlaceOtocoResponseResultOrdersInner) *NullableOrderListPlaceOtocoResponseResultOrdersInner {
	return &NullableOrderListPlaceOtocoResponseResultOrdersInner{value: val, isSet: true}
}

func (v NullableOrderListPlaceOtocoResponseResultOrdersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderListPlaceOtocoResponseResultOrdersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
