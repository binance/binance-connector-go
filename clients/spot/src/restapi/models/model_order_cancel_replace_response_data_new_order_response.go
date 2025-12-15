/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OrderCancelReplaceResponseDataNewOrderResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderCancelReplaceResponseDataNewOrderResponse{}

// OrderCancelReplaceResponseDataNewOrderResponse struct for OrderCancelReplaceResponseDataNewOrderResponse
type OrderCancelReplaceResponseDataNewOrderResponse struct {
	Code                 *int64  `json:"code,omitempty"`
	Msg                  *string `json:"msg,omitempty"`
	Symbol               *string `json:"symbol,omitempty"`
	OrderId              *int64  `json:"orderId,omitempty"`
	OrderListId          *int64  `json:"orderListId,omitempty"`
	ClientOrderId        *string `json:"clientOrderId,omitempty"`
	TransactTime         *int64  `json:"transactTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderCancelReplaceResponseDataNewOrderResponse OrderCancelReplaceResponseDataNewOrderResponse

// NewOrderCancelReplaceResponseDataNewOrderResponse instantiates a new OrderCancelReplaceResponseDataNewOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCancelReplaceResponseDataNewOrderResponse() *OrderCancelReplaceResponseDataNewOrderResponse {
	this := OrderCancelReplaceResponseDataNewOrderResponse{}
	return &this
}

// NewOrderCancelReplaceResponseDataNewOrderResponseWithDefaults instantiates a new OrderCancelReplaceResponseDataNewOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCancelReplaceResponseDataNewOrderResponseWithDefaults() *OrderCancelReplaceResponseDataNewOrderResponse {
	this := OrderCancelReplaceResponseDataNewOrderResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseDataNewOrderResponse) GetCode() int64 {
	if o == nil || common.IsNil(o.Code) {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseDataNewOrderResponse) GetCodeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseDataNewOrderResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *OrderCancelReplaceResponseDataNewOrderResponse) SetCode(v int64) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseDataNewOrderResponse) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseDataNewOrderResponse) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseDataNewOrderResponse) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *OrderCancelReplaceResponseDataNewOrderResponse) SetMsg(v string) {
	o.Msg = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseDataNewOrderResponse) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseDataNewOrderResponse) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseDataNewOrderResponse) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *OrderCancelReplaceResponseDataNewOrderResponse) SetSymbol(v string) {
	o.Symbol = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseDataNewOrderResponse) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseDataNewOrderResponse) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseDataNewOrderResponse) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *OrderCancelReplaceResponseDataNewOrderResponse) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetOrderListId returns the OrderListId field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseDataNewOrderResponse) GetOrderListId() int64 {
	if o == nil || common.IsNil(o.OrderListId) {
		var ret int64
		return ret
	}
	return *o.OrderListId
}

// GetOrderListIdOk returns a tuple with the OrderListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseDataNewOrderResponse) GetOrderListIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderListId) {
		return nil, false
	}
	return o.OrderListId, true
}

// HasOrderListId returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseDataNewOrderResponse) HasOrderListId() bool {
	if o != nil && !common.IsNil(o.OrderListId) {
		return true
	}

	return false
}

// SetOrderListId gets a reference to the given int64 and assigns it to the OrderListId field.
func (o *OrderCancelReplaceResponseDataNewOrderResponse) SetOrderListId(v int64) {
	o.OrderListId = &v
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseDataNewOrderResponse) GetClientOrderId() string {
	if o == nil || common.IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseDataNewOrderResponse) GetClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseDataNewOrderResponse) HasClientOrderId() bool {
	if o != nil && !common.IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *OrderCancelReplaceResponseDataNewOrderResponse) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

// GetTransactTime returns the TransactTime field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseDataNewOrderResponse) GetTransactTime() int64 {
	if o == nil || common.IsNil(o.TransactTime) {
		var ret int64
		return ret
	}
	return *o.TransactTime
}

// GetTransactTimeOk returns a tuple with the TransactTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseDataNewOrderResponse) GetTransactTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TransactTime) {
		return nil, false
	}
	return o.TransactTime, true
}

// HasTransactTime returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseDataNewOrderResponse) HasTransactTime() bool {
	if o != nil && !common.IsNil(o.TransactTime) {
		return true
	}

	return false
}

// SetTransactTime gets a reference to the given int64 and assigns it to the TransactTime field.
func (o *OrderCancelReplaceResponseDataNewOrderResponse) SetTransactTime(v int64) {
	o.TransactTime = &v
}

func (o OrderCancelReplaceResponseDataNewOrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCancelReplaceResponseDataNewOrderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !common.IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.OrderListId) {
		toSerialize["orderListId"] = o.OrderListId
	}
	if !common.IsNil(o.ClientOrderId) {
		toSerialize["clientOrderId"] = o.ClientOrderId
	}
	if !common.IsNil(o.TransactTime) {
		toSerialize["transactTime"] = o.TransactTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderCancelReplaceResponseDataNewOrderResponse) UnmarshalJSON(data []byte) (err error) {
	varOrderCancelReplaceResponseDataNewOrderResponse := _OrderCancelReplaceResponseDataNewOrderResponse{}

	err = json.Unmarshal(data, &varOrderCancelReplaceResponseDataNewOrderResponse)

	if err != nil {
		return err
	}

	*o = OrderCancelReplaceResponseDataNewOrderResponse(varOrderCancelReplaceResponseDataNewOrderResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "msg")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "orderListId")
		delete(additionalProperties, "clientOrderId")
		delete(additionalProperties, "transactTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrderCancelReplaceResponseDataNewOrderResponse struct {
	value *OrderCancelReplaceResponseDataNewOrderResponse
	isSet bool
}

func (v NullableOrderCancelReplaceResponseDataNewOrderResponse) Get() *OrderCancelReplaceResponseDataNewOrderResponse {
	return v.value
}

func (v *NullableOrderCancelReplaceResponseDataNewOrderResponse) Set(val *OrderCancelReplaceResponseDataNewOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCancelReplaceResponseDataNewOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCancelReplaceResponseDataNewOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCancelReplaceResponseDataNewOrderResponse(val *OrderCancelReplaceResponseDataNewOrderResponse) *NullableOrderCancelReplaceResponseDataNewOrderResponse {
	return &NullableOrderCancelReplaceResponseDataNewOrderResponse{value: val, isSet: true}
}

func (v NullableOrderCancelReplaceResponseDataNewOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCancelReplaceResponseDataNewOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
