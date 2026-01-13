/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OrderAmendKeepPriorityResponseResultListStatusOrdersInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderAmendKeepPriorityResponseResultListStatusOrdersInner{}

// OrderAmendKeepPriorityResponseResultListStatusOrdersInner struct for OrderAmendKeepPriorityResponseResultListStatusOrdersInner
type OrderAmendKeepPriorityResponseResultListStatusOrdersInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	OrderId              *int64  `json:"orderId,omitempty"`
	ClientOrderId        *string `json:"clientOrderId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderAmendKeepPriorityResponseResultListStatusOrdersInner OrderAmendKeepPriorityResponseResultListStatusOrdersInner

// NewOrderAmendKeepPriorityResponseResultListStatusOrdersInner instantiates a new OrderAmendKeepPriorityResponseResultListStatusOrdersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderAmendKeepPriorityResponseResultListStatusOrdersInner() *OrderAmendKeepPriorityResponseResultListStatusOrdersInner {
	this := OrderAmendKeepPriorityResponseResultListStatusOrdersInner{}
	return &this
}

// NewOrderAmendKeepPriorityResponseResultListStatusOrdersInnerWithDefaults instantiates a new OrderAmendKeepPriorityResponseResultListStatusOrdersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderAmendKeepPriorityResponseResultListStatusOrdersInnerWithDefaults() *OrderAmendKeepPriorityResponseResultListStatusOrdersInner {
	this := OrderAmendKeepPriorityResponseResultListStatusOrdersInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseResultListStatusOrdersInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseResultListStatusOrdersInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseResultListStatusOrdersInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *OrderAmendKeepPriorityResponseResultListStatusOrdersInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseResultListStatusOrdersInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseResultListStatusOrdersInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseResultListStatusOrdersInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *OrderAmendKeepPriorityResponseResultListStatusOrdersInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseResultListStatusOrdersInner) GetClientOrderId() string {
	if o == nil || common.IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseResultListStatusOrdersInner) GetClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseResultListStatusOrdersInner) HasClientOrderId() bool {
	if o != nil && !common.IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *OrderAmendKeepPriorityResponseResultListStatusOrdersInner) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

func (o OrderAmendKeepPriorityResponseResultListStatusOrdersInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderAmendKeepPriorityResponseResultListStatusOrdersInner) ToMap() (map[string]interface{}, error) {
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

func (o *OrderAmendKeepPriorityResponseResultListStatusOrdersInner) UnmarshalJSON(data []byte) (err error) {
	varOrderAmendKeepPriorityResponseResultListStatusOrdersInner := _OrderAmendKeepPriorityResponseResultListStatusOrdersInner{}

	err = json.Unmarshal(data, &varOrderAmendKeepPriorityResponseResultListStatusOrdersInner)

	if err != nil {
		return err
	}

	*o = OrderAmendKeepPriorityResponseResultListStatusOrdersInner(varOrderAmendKeepPriorityResponseResultListStatusOrdersInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "clientOrderId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrderAmendKeepPriorityResponseResultListStatusOrdersInner struct {
	value *OrderAmendKeepPriorityResponseResultListStatusOrdersInner
	isSet bool
}

func (v NullableOrderAmendKeepPriorityResponseResultListStatusOrdersInner) Get() *OrderAmendKeepPriorityResponseResultListStatusOrdersInner {
	return v.value
}

func (v *NullableOrderAmendKeepPriorityResponseResultListStatusOrdersInner) Set(val *OrderAmendKeepPriorityResponseResultListStatusOrdersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderAmendKeepPriorityResponseResultListStatusOrdersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderAmendKeepPriorityResponseResultListStatusOrdersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderAmendKeepPriorityResponseResultListStatusOrdersInner(val *OrderAmendKeepPriorityResponseResultListStatusOrdersInner) *NullableOrderAmendKeepPriorityResponseResultListStatusOrdersInner {
	return &NullableOrderAmendKeepPriorityResponseResultListStatusOrdersInner{value: val, isSet: true}
}

func (v NullableOrderAmendKeepPriorityResponseResultListStatusOrdersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderAmendKeepPriorityResponseResultListStatusOrdersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
