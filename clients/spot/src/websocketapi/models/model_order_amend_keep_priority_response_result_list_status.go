/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the OrderAmendKeepPriorityResponseResultListStatus type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderAmendKeepPriorityResponseResultListStatus{}

// OrderAmendKeepPriorityResponseResultListStatus struct for OrderAmendKeepPriorityResponseResultListStatus
type OrderAmendKeepPriorityResponseResultListStatus struct {
	OrderListId          *int64                                                      `json:"orderListId,omitempty"`
	ContingencyType      *string                                                     `json:"contingencyType,omitempty"`
	ListOrderStatus      *string                                                     `json:"listOrderStatus,omitempty"`
	ListClientOrderId    *string                                                     `json:"listClientOrderId,omitempty"`
	Symbol               *string                                                     `json:"symbol,omitempty"`
	Orders               []OrderAmendKeepPriorityResponseResultListStatusOrdersInner `json:"orders,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderAmendKeepPriorityResponseResultListStatus OrderAmendKeepPriorityResponseResultListStatus

// NewOrderAmendKeepPriorityResponseResultListStatus instantiates a new OrderAmendKeepPriorityResponseResultListStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderAmendKeepPriorityResponseResultListStatus() *OrderAmendKeepPriorityResponseResultListStatus {
	this := OrderAmendKeepPriorityResponseResultListStatus{}
	return &this
}

// NewOrderAmendKeepPriorityResponseResultListStatusWithDefaults instantiates a new OrderAmendKeepPriorityResponseResultListStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderAmendKeepPriorityResponseResultListStatusWithDefaults() *OrderAmendKeepPriorityResponseResultListStatus {
	this := OrderAmendKeepPriorityResponseResultListStatus{}
	return &this
}

// GetOrderListId returns the OrderListId field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseResultListStatus) GetOrderListId() int64 {
	if o == nil || common.IsNil(o.OrderListId) {
		var ret int64
		return ret
	}
	return *o.OrderListId
}

// GetOrderListIdOk returns a tuple with the OrderListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseResultListStatus) GetOrderListIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderListId) {
		return nil, false
	}
	return o.OrderListId, true
}

// HasOrderListId returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseResultListStatus) HasOrderListId() bool {
	if o != nil && !common.IsNil(o.OrderListId) {
		return true
	}

	return false
}

// SetOrderListId gets a reference to the given int64 and assigns it to the OrderListId field.
func (o *OrderAmendKeepPriorityResponseResultListStatus) SetOrderListId(v int64) {
	o.OrderListId = &v
}

// GetContingencyType returns the ContingencyType field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseResultListStatus) GetContingencyType() string {
	if o == nil || common.IsNil(o.ContingencyType) {
		var ret string
		return ret
	}
	return *o.ContingencyType
}

// GetContingencyTypeOk returns a tuple with the ContingencyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseResultListStatus) GetContingencyTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ContingencyType) {
		return nil, false
	}
	return o.ContingencyType, true
}

// HasContingencyType returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseResultListStatus) HasContingencyType() bool {
	if o != nil && !common.IsNil(o.ContingencyType) {
		return true
	}

	return false
}

// SetContingencyType gets a reference to the given string and assigns it to the ContingencyType field.
func (o *OrderAmendKeepPriorityResponseResultListStatus) SetContingencyType(v string) {
	o.ContingencyType = &v
}

// GetListOrderStatus returns the ListOrderStatus field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseResultListStatus) GetListOrderStatus() string {
	if o == nil || common.IsNil(o.ListOrderStatus) {
		var ret string
		return ret
	}
	return *o.ListOrderStatus
}

// GetListOrderStatusOk returns a tuple with the ListOrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseResultListStatus) GetListOrderStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListOrderStatus) {
		return nil, false
	}
	return o.ListOrderStatus, true
}

// HasListOrderStatus returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseResultListStatus) HasListOrderStatus() bool {
	if o != nil && !common.IsNil(o.ListOrderStatus) {
		return true
	}

	return false
}

// SetListOrderStatus gets a reference to the given string and assigns it to the ListOrderStatus field.
func (o *OrderAmendKeepPriorityResponseResultListStatus) SetListOrderStatus(v string) {
	o.ListOrderStatus = &v
}

// GetListClientOrderId returns the ListClientOrderId field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseResultListStatus) GetListClientOrderId() string {
	if o == nil || common.IsNil(o.ListClientOrderId) {
		var ret string
		return ret
	}
	return *o.ListClientOrderId
}

// GetListClientOrderIdOk returns a tuple with the ListClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseResultListStatus) GetListClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListClientOrderId) {
		return nil, false
	}
	return o.ListClientOrderId, true
}

// HasListClientOrderId returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseResultListStatus) HasListClientOrderId() bool {
	if o != nil && !common.IsNil(o.ListClientOrderId) {
		return true
	}

	return false
}

// SetListClientOrderId gets a reference to the given string and assigns it to the ListClientOrderId field.
func (o *OrderAmendKeepPriorityResponseResultListStatus) SetListClientOrderId(v string) {
	o.ListClientOrderId = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseResultListStatus) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseResultListStatus) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseResultListStatus) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *OrderAmendKeepPriorityResponseResultListStatus) SetSymbol(v string) {
	o.Symbol = &v
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseResultListStatus) GetOrders() []OrderAmendKeepPriorityResponseResultListStatusOrdersInner {
	if o == nil || common.IsNil(o.Orders) {
		var ret []OrderAmendKeepPriorityResponseResultListStatusOrdersInner
		return ret
	}
	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseResultListStatus) GetOrdersOk() ([]OrderAmendKeepPriorityResponseResultListStatusOrdersInner, bool) {
	if o == nil || common.IsNil(o.Orders) {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseResultListStatus) HasOrders() bool {
	if o != nil && !common.IsNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []OrderAmendKeepPriorityResponseResultListStatusOrdersInner and assigns it to the Orders field.
func (o *OrderAmendKeepPriorityResponseResultListStatus) SetOrders(v []OrderAmendKeepPriorityResponseResultListStatusOrdersInner) {
	o.Orders = v
}

func (o OrderAmendKeepPriorityResponseResultListStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderAmendKeepPriorityResponseResultListStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.OrderListId) {
		toSerialize["orderListId"] = o.OrderListId
	}
	if !common.IsNil(o.ContingencyType) {
		toSerialize["contingencyType"] = o.ContingencyType
	}
	if !common.IsNil(o.ListOrderStatus) {
		toSerialize["listOrderStatus"] = o.ListOrderStatus
	}
	if !common.IsNil(o.ListClientOrderId) {
		toSerialize["listClientOrderId"] = o.ListClientOrderId
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Orders) {
		toSerialize["orders"] = o.Orders
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderAmendKeepPriorityResponseResultListStatus) UnmarshalJSON(data []byte) (err error) {
	varOrderAmendKeepPriorityResponseResultListStatus := _OrderAmendKeepPriorityResponseResultListStatus{}

	err = json.Unmarshal(data, &varOrderAmendKeepPriorityResponseResultListStatus)

	if err != nil {
		return err
	}

	*o = OrderAmendKeepPriorityResponseResultListStatus(varOrderAmendKeepPriorityResponseResultListStatus)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderListId")
		delete(additionalProperties, "contingencyType")
		delete(additionalProperties, "listOrderStatus")
		delete(additionalProperties, "listClientOrderId")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "orders")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrderAmendKeepPriorityResponseResultListStatus struct {
	value *OrderAmendKeepPriorityResponseResultListStatus
	isSet bool
}

func (v NullableOrderAmendKeepPriorityResponseResultListStatus) Get() *OrderAmendKeepPriorityResponseResultListStatus {
	return v.value
}

func (v *NullableOrderAmendKeepPriorityResponseResultListStatus) Set(val *OrderAmendKeepPriorityResponseResultListStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderAmendKeepPriorityResponseResultListStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderAmendKeepPriorityResponseResultListStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderAmendKeepPriorityResponseResultListStatus(val *OrderAmendKeepPriorityResponseResultListStatus) *NullableOrderAmendKeepPriorityResponseResultListStatus {
	return &NullableOrderAmendKeepPriorityResponseResultListStatus{value: val, isSet: true}
}

func (v NullableOrderAmendKeepPriorityResponseResultListStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderAmendKeepPriorityResponseResultListStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
