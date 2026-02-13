/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ModifyMultipleOrdersBatchOrdersParameterInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ModifyMultipleOrdersBatchOrdersParameterInner{}

// ModifyMultipleOrdersBatchOrdersParameterInner struct for ModifyMultipleOrdersBatchOrdersParameterInner
type ModifyMultipleOrdersBatchOrdersParameterInner struct {
	OrderId              *string                          `json:"orderId,omitempty"`
	OrigClientOrderId    *string                          `json:"origClientOrderId,omitempty"`
	Symbol               *string                          `json:"symbol,omitempty"`
	Side                 *NewAlgoOrderSideParameter       `json:"side,omitempty"`
	Quantity             *string                          `json:"quantity,omitempty"`
	Price                *string                          `json:"price,omitempty"`
	PriceMatch           *NewAlgoOrderPriceMatchParameter `json:"priceMatch,omitempty"`
	StopPrice            *string                          `json:"stopPrice,omitempty"`
	RecvWindow           *string                          `json:"recvWindow,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModifyMultipleOrdersBatchOrdersParameterInner ModifyMultipleOrdersBatchOrdersParameterInner

// NewModifyMultipleOrdersBatchOrdersParameterInner instantiates a new ModifyMultipleOrdersBatchOrdersParameterInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyMultipleOrdersBatchOrdersParameterInner() *ModifyMultipleOrdersBatchOrdersParameterInner {
	this := ModifyMultipleOrdersBatchOrdersParameterInner{}
	return &this
}

// NewModifyMultipleOrdersBatchOrdersParameterInnerWithDefaults instantiates a new ModifyMultipleOrdersBatchOrdersParameterInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyMultipleOrdersBatchOrdersParameterInnerWithDefaults() *ModifyMultipleOrdersBatchOrdersParameterInner {
	this := ModifyMultipleOrdersBatchOrdersParameterInner{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetOrderId() string {
	if o == nil || common.IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *ModifyMultipleOrdersBatchOrdersParameterInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *ModifyMultipleOrdersBatchOrdersParameterInner) SetOrderId(v string) {
	o.OrderId = &v
}

// GetOrigClientOrderId returns the OrigClientOrderId field value if set, zero value otherwise.
func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetOrigClientOrderId() string {
	if o == nil || common.IsNil(o.OrigClientOrderId) {
		var ret string
		return ret
	}
	return *o.OrigClientOrderId
}

// GetOrigClientOrderIdOk returns a tuple with the OrigClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetOrigClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrigClientOrderId) {
		return nil, false
	}
	return o.OrigClientOrderId, true
}

// HasOrigClientOrderId returns a boolean if a field has been set.
func (o *ModifyMultipleOrdersBatchOrdersParameterInner) HasOrigClientOrderId() bool {
	if o != nil && !common.IsNil(o.OrigClientOrderId) {
		return true
	}

	return false
}

// SetOrigClientOrderId gets a reference to the given string and assigns it to the OrigClientOrderId field.
func (o *ModifyMultipleOrdersBatchOrdersParameterInner) SetOrigClientOrderId(v string) {
	o.OrigClientOrderId = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *ModifyMultipleOrdersBatchOrdersParameterInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *ModifyMultipleOrdersBatchOrdersParameterInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetSide() NewAlgoOrderSideParameter {
	if o == nil || common.IsNil(o.Side) {
		var ret NewAlgoOrderSideParameter
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetSideOk() (*NewAlgoOrderSideParameter, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *ModifyMultipleOrdersBatchOrdersParameterInner) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given NewAlgoOrderSideParameter and assigns it to the Side field.
func (o *ModifyMultipleOrdersBatchOrdersParameterInner) SetSide(v NewAlgoOrderSideParameter) {
	o.Side = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetQuantity() string {
	if o == nil || common.IsNil(o.Quantity) {
		var ret string
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetQuantityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *ModifyMultipleOrdersBatchOrdersParameterInner) HasQuantity() bool {
	if o != nil && !common.IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given string and assigns it to the Quantity field.
func (o *ModifyMultipleOrdersBatchOrdersParameterInner) SetQuantity(v string) {
	o.Quantity = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *ModifyMultipleOrdersBatchOrdersParameterInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *ModifyMultipleOrdersBatchOrdersParameterInner) SetPrice(v string) {
	o.Price = &v
}

// GetPriceMatch returns the PriceMatch field value if set, zero value otherwise.
func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetPriceMatch() NewAlgoOrderPriceMatchParameter {
	if o == nil || common.IsNil(o.PriceMatch) {
		var ret NewAlgoOrderPriceMatchParameter
		return ret
	}
	return *o.PriceMatch
}

// GetPriceMatchOk returns a tuple with the PriceMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetPriceMatchOk() (*NewAlgoOrderPriceMatchParameter, bool) {
	if o == nil || common.IsNil(o.PriceMatch) {
		return nil, false
	}
	return o.PriceMatch, true
}

// HasPriceMatch returns a boolean if a field has been set.
func (o *ModifyMultipleOrdersBatchOrdersParameterInner) HasPriceMatch() bool {
	if o != nil && !common.IsNil(o.PriceMatch) {
		return true
	}

	return false
}

// SetPriceMatch gets a reference to the given NewAlgoOrderPriceMatchParameter and assigns it to the PriceMatch field.
func (o *ModifyMultipleOrdersBatchOrdersParameterInner) SetPriceMatch(v NewAlgoOrderPriceMatchParameter) {
	o.PriceMatch = &v
}

// GetStopPrice returns the StopPrice field value if set, zero value otherwise.
func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetStopPrice() string {
	if o == nil || common.IsNil(o.StopPrice) {
		var ret string
		return ret
	}
	return *o.StopPrice
}

// GetStopPriceOk returns a tuple with the StopPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetStopPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.StopPrice) {
		return nil, false
	}
	return o.StopPrice, true
}

// HasStopPrice returns a boolean if a field has been set.
func (o *ModifyMultipleOrdersBatchOrdersParameterInner) HasStopPrice() bool {
	if o != nil && !common.IsNil(o.StopPrice) {
		return true
	}

	return false
}

// SetStopPrice gets a reference to the given string and assigns it to the StopPrice field.
func (o *ModifyMultipleOrdersBatchOrdersParameterInner) SetStopPrice(v string) {
	o.StopPrice = &v
}

// GetRecvWindow returns the RecvWindow field value if set, zero value otherwise.
func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetRecvWindow() string {
	if o == nil || common.IsNil(o.RecvWindow) {
		var ret string
		return ret
	}
	return *o.RecvWindow
}

// GetRecvWindowOk returns a tuple with the RecvWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetRecvWindowOk() (*string, bool) {
	if o == nil || common.IsNil(o.RecvWindow) {
		return nil, false
	}
	return o.RecvWindow, true
}

// HasRecvWindow returns a boolean if a field has been set.
func (o *ModifyMultipleOrdersBatchOrdersParameterInner) HasRecvWindow() bool {
	if o != nil && !common.IsNil(o.RecvWindow) {
		return true
	}

	return false
}

// SetRecvWindow gets a reference to the given string and assigns it to the RecvWindow field.
func (o *ModifyMultipleOrdersBatchOrdersParameterInner) SetRecvWindow(v string) {
	o.RecvWindow = &v
}

func (o ModifyMultipleOrdersBatchOrdersParameterInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyMultipleOrdersBatchOrdersParameterInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.OrigClientOrderId) {
		toSerialize["origClientOrderId"] = o.OrigClientOrderId
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !common.IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.PriceMatch) {
		toSerialize["priceMatch"] = o.PriceMatch
	}
	if !common.IsNil(o.StopPrice) {
		toSerialize["stopPrice"] = o.StopPrice
	}
	if !common.IsNil(o.RecvWindow) {
		toSerialize["recvWindow"] = o.RecvWindow
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModifyMultipleOrdersBatchOrdersParameterInner) UnmarshalJSON(data []byte) (err error) {
	varModifyMultipleOrdersBatchOrdersParameterInner := _ModifyMultipleOrdersBatchOrdersParameterInner{}

	err = json.Unmarshal(data, &varModifyMultipleOrdersBatchOrdersParameterInner)

	if err != nil {
		return err
	}

	*o = ModifyMultipleOrdersBatchOrdersParameterInner(varModifyMultipleOrdersBatchOrdersParameterInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "origClientOrderId")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "side")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "price")
		delete(additionalProperties, "priceMatch")
		delete(additionalProperties, "stopPrice")
		delete(additionalProperties, "recvWindow")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModifyMultipleOrdersBatchOrdersParameterInner struct {
	value *ModifyMultipleOrdersBatchOrdersParameterInner
	isSet bool
}

func (v NullableModifyMultipleOrdersBatchOrdersParameterInner) Get() *ModifyMultipleOrdersBatchOrdersParameterInner {
	return v.value
}

func (v *NullableModifyMultipleOrdersBatchOrdersParameterInner) Set(val *ModifyMultipleOrdersBatchOrdersParameterInner) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyMultipleOrdersBatchOrdersParameterInner) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyMultipleOrdersBatchOrdersParameterInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyMultipleOrdersBatchOrdersParameterInner(val *ModifyMultipleOrdersBatchOrdersParameterInner) *NullableModifyMultipleOrdersBatchOrdersParameterInner {
	return &NullableModifyMultipleOrdersBatchOrdersParameterInner{value: val, isSet: true}
}

func (v NullableModifyMultipleOrdersBatchOrdersParameterInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyMultipleOrdersBatchOrdersParameterInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
