/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OrderListPlaceOtoResponseResultOrdersInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderListPlaceOtoResponseResultOrdersInner{}

// OrderListPlaceOtoResponseResultOrdersInner struct for OrderListPlaceOtoResponseResultOrdersInner
type OrderListPlaceOtoResponseResultOrdersInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	OrderId              *int64  `json:"orderId,omitempty"`
	ClientOrderId        *string `json:"clientOrderId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderListPlaceOtoResponseResultOrdersInner OrderListPlaceOtoResponseResultOrdersInner

// NewOrderListPlaceOtoResponseResultOrdersInner instantiates a new OrderListPlaceOtoResponseResultOrdersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderListPlaceOtoResponseResultOrdersInner() *OrderListPlaceOtoResponseResultOrdersInner {
	this := OrderListPlaceOtoResponseResultOrdersInner{}
	return &this
}

// NewOrderListPlaceOtoResponseResultOrdersInnerWithDefaults instantiates a new OrderListPlaceOtoResponseResultOrdersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderListPlaceOtoResponseResultOrdersInnerWithDefaults() *OrderListPlaceOtoResponseResultOrdersInner {
	this := OrderListPlaceOtoResponseResultOrdersInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *OrderListPlaceOtoResponseResultOrdersInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOtoResponseResultOrdersInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *OrderListPlaceOtoResponseResultOrdersInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *OrderListPlaceOtoResponseResultOrdersInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *OrderListPlaceOtoResponseResultOrdersInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOtoResponseResultOrdersInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *OrderListPlaceOtoResponseResultOrdersInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *OrderListPlaceOtoResponseResultOrdersInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *OrderListPlaceOtoResponseResultOrdersInner) GetClientOrderId() string {
	if o == nil || common.IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOtoResponseResultOrdersInner) GetClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *OrderListPlaceOtoResponseResultOrdersInner) HasClientOrderId() bool {
	if o != nil && !common.IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *OrderListPlaceOtoResponseResultOrdersInner) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

func (o OrderListPlaceOtoResponseResultOrdersInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderListPlaceOtoResponseResultOrdersInner) ToMap() (map[string]interface{}, error) {
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

func (o *OrderListPlaceOtoResponseResultOrdersInner) UnmarshalJSON(data []byte) (err error) {
	varOrderListPlaceOtoResponseResultOrdersInner := _OrderListPlaceOtoResponseResultOrdersInner{}

	err = json.Unmarshal(data, &varOrderListPlaceOtoResponseResultOrdersInner)

	if err != nil {
		return err
	}

	*o = OrderListPlaceOtoResponseResultOrdersInner(varOrderListPlaceOtoResponseResultOrdersInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "clientOrderId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrderListPlaceOtoResponseResultOrdersInner struct {
	value *OrderListPlaceOtoResponseResultOrdersInner
	isSet bool
}

func (v NullableOrderListPlaceOtoResponseResultOrdersInner) Get() *OrderListPlaceOtoResponseResultOrdersInner {
	return v.value
}

func (v *NullableOrderListPlaceOtoResponseResultOrdersInner) Set(val *OrderListPlaceOtoResponseResultOrdersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderListPlaceOtoResponseResultOrdersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderListPlaceOtoResponseResultOrdersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderListPlaceOtoResponseResultOrdersInner(val *OrderListPlaceOtoResponseResultOrdersInner) *NullableOrderListPlaceOtoResponseResultOrdersInner {
	return &NullableOrderListPlaceOtoResponseResultOrdersInner{value: val, isSet: true}
}

func (v NullableOrderListPlaceOtoResponseResultOrdersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderListPlaceOtoResponseResultOrdersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
