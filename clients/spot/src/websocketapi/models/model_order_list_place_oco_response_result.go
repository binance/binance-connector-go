/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the OrderListPlaceOcoResponseResult type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderListPlaceOcoResponseResult{}

// OrderListPlaceOcoResponseResult struct for OrderListPlaceOcoResponseResult
type OrderListPlaceOcoResponseResult struct {
	OrderListId          *int64                                             `json:"orderListId,omitempty"`
	ContingencyType      *string                                            `json:"contingencyType,omitempty"`
	ListStatusType       *string                                            `json:"listStatusType,omitempty"`
	ListOrderStatus      *string                                            `json:"listOrderStatus,omitempty"`
	ListClientOrderId    *string                                            `json:"listClientOrderId,omitempty"`
	TransactionTime      *int64                                             `json:"transactionTime,omitempty"`
	Symbol               *string                                            `json:"symbol,omitempty"`
	Orders               []OrderListPlaceOcoResponseResultOrdersInner       `json:"orders,omitempty"`
	OrderReports         []OrderListPlaceOcoResponseResultOrderReportsInner `json:"orderReports,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderListPlaceOcoResponseResult OrderListPlaceOcoResponseResult

// NewOrderListPlaceOcoResponseResult instantiates a new OrderListPlaceOcoResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderListPlaceOcoResponseResult() *OrderListPlaceOcoResponseResult {
	this := OrderListPlaceOcoResponseResult{}
	return &this
}

// NewOrderListPlaceOcoResponseResultWithDefaults instantiates a new OrderListPlaceOcoResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderListPlaceOcoResponseResultWithDefaults() *OrderListPlaceOcoResponseResult {
	this := OrderListPlaceOcoResponseResult{}
	return &this
}

// GetOrderListId returns the OrderListId field value if set, zero value otherwise.
func (o *OrderListPlaceOcoResponseResult) GetOrderListId() int64 {
	if o == nil || common.IsNil(o.OrderListId) {
		var ret int64
		return ret
	}
	return *o.OrderListId
}

// GetOrderListIdOk returns a tuple with the OrderListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOcoResponseResult) GetOrderListIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderListId) {
		return nil, false
	}
	return o.OrderListId, true
}

// HasOrderListId returns a boolean if a field has been set.
func (o *OrderListPlaceOcoResponseResult) HasOrderListId() bool {
	if o != nil && !common.IsNil(o.OrderListId) {
		return true
	}

	return false
}

// SetOrderListId gets a reference to the given int64 and assigns it to the OrderListId field.
func (o *OrderListPlaceOcoResponseResult) SetOrderListId(v int64) {
	o.OrderListId = &v
}

// GetContingencyType returns the ContingencyType field value if set, zero value otherwise.
func (o *OrderListPlaceOcoResponseResult) GetContingencyType() string {
	if o == nil || common.IsNil(o.ContingencyType) {
		var ret string
		return ret
	}
	return *o.ContingencyType
}

// GetContingencyTypeOk returns a tuple with the ContingencyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOcoResponseResult) GetContingencyTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ContingencyType) {
		return nil, false
	}
	return o.ContingencyType, true
}

// HasContingencyType returns a boolean if a field has been set.
func (o *OrderListPlaceOcoResponseResult) HasContingencyType() bool {
	if o != nil && !common.IsNil(o.ContingencyType) {
		return true
	}

	return false
}

// SetContingencyType gets a reference to the given string and assigns it to the ContingencyType field.
func (o *OrderListPlaceOcoResponseResult) SetContingencyType(v string) {
	o.ContingencyType = &v
}

// GetListStatusType returns the ListStatusType field value if set, zero value otherwise.
func (o *OrderListPlaceOcoResponseResult) GetListStatusType() string {
	if o == nil || common.IsNil(o.ListStatusType) {
		var ret string
		return ret
	}
	return *o.ListStatusType
}

// GetListStatusTypeOk returns a tuple with the ListStatusType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOcoResponseResult) GetListStatusTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListStatusType) {
		return nil, false
	}
	return o.ListStatusType, true
}

// HasListStatusType returns a boolean if a field has been set.
func (o *OrderListPlaceOcoResponseResult) HasListStatusType() bool {
	if o != nil && !common.IsNil(o.ListStatusType) {
		return true
	}

	return false
}

// SetListStatusType gets a reference to the given string and assigns it to the ListStatusType field.
func (o *OrderListPlaceOcoResponseResult) SetListStatusType(v string) {
	o.ListStatusType = &v
}

// GetListOrderStatus returns the ListOrderStatus field value if set, zero value otherwise.
func (o *OrderListPlaceOcoResponseResult) GetListOrderStatus() string {
	if o == nil || common.IsNil(o.ListOrderStatus) {
		var ret string
		return ret
	}
	return *o.ListOrderStatus
}

// GetListOrderStatusOk returns a tuple with the ListOrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOcoResponseResult) GetListOrderStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListOrderStatus) {
		return nil, false
	}
	return o.ListOrderStatus, true
}

// HasListOrderStatus returns a boolean if a field has been set.
func (o *OrderListPlaceOcoResponseResult) HasListOrderStatus() bool {
	if o != nil && !common.IsNil(o.ListOrderStatus) {
		return true
	}

	return false
}

// SetListOrderStatus gets a reference to the given string and assigns it to the ListOrderStatus field.
func (o *OrderListPlaceOcoResponseResult) SetListOrderStatus(v string) {
	o.ListOrderStatus = &v
}

// GetListClientOrderId returns the ListClientOrderId field value if set, zero value otherwise.
func (o *OrderListPlaceOcoResponseResult) GetListClientOrderId() string {
	if o == nil || common.IsNil(o.ListClientOrderId) {
		var ret string
		return ret
	}
	return *o.ListClientOrderId
}

// GetListClientOrderIdOk returns a tuple with the ListClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOcoResponseResult) GetListClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListClientOrderId) {
		return nil, false
	}
	return o.ListClientOrderId, true
}

// HasListClientOrderId returns a boolean if a field has been set.
func (o *OrderListPlaceOcoResponseResult) HasListClientOrderId() bool {
	if o != nil && !common.IsNil(o.ListClientOrderId) {
		return true
	}

	return false
}

// SetListClientOrderId gets a reference to the given string and assigns it to the ListClientOrderId field.
func (o *OrderListPlaceOcoResponseResult) SetListClientOrderId(v string) {
	o.ListClientOrderId = &v
}

// GetTransactionTime returns the TransactionTime field value if set, zero value otherwise.
func (o *OrderListPlaceOcoResponseResult) GetTransactionTime() int64 {
	if o == nil || common.IsNil(o.TransactionTime) {
		var ret int64
		return ret
	}
	return *o.TransactionTime
}

// GetTransactionTimeOk returns a tuple with the TransactionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOcoResponseResult) GetTransactionTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TransactionTime) {
		return nil, false
	}
	return o.TransactionTime, true
}

// HasTransactionTime returns a boolean if a field has been set.
func (o *OrderListPlaceOcoResponseResult) HasTransactionTime() bool {
	if o != nil && !common.IsNil(o.TransactionTime) {
		return true
	}

	return false
}

// SetTransactionTime gets a reference to the given int64 and assigns it to the TransactionTime field.
func (o *OrderListPlaceOcoResponseResult) SetTransactionTime(v int64) {
	o.TransactionTime = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *OrderListPlaceOcoResponseResult) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOcoResponseResult) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *OrderListPlaceOcoResponseResult) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *OrderListPlaceOcoResponseResult) SetSymbol(v string) {
	o.Symbol = &v
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *OrderListPlaceOcoResponseResult) GetOrders() []OrderListPlaceOcoResponseResultOrdersInner {
	if o == nil || common.IsNil(o.Orders) {
		var ret []OrderListPlaceOcoResponseResultOrdersInner
		return ret
	}
	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOcoResponseResult) GetOrdersOk() ([]OrderListPlaceOcoResponseResultOrdersInner, bool) {
	if o == nil || common.IsNil(o.Orders) {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *OrderListPlaceOcoResponseResult) HasOrders() bool {
	if o != nil && !common.IsNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []OrderListPlaceOcoResponseResultOrdersInner and assigns it to the Orders field.
func (o *OrderListPlaceOcoResponseResult) SetOrders(v []OrderListPlaceOcoResponseResultOrdersInner) {
	o.Orders = v
}

// GetOrderReports returns the OrderReports field value if set, zero value otherwise.
func (o *OrderListPlaceOcoResponseResult) GetOrderReports() []OrderListPlaceOcoResponseResultOrderReportsInner {
	if o == nil || common.IsNil(o.OrderReports) {
		var ret []OrderListPlaceOcoResponseResultOrderReportsInner
		return ret
	}
	return o.OrderReports
}

// GetOrderReportsOk returns a tuple with the OrderReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOcoResponseResult) GetOrderReportsOk() ([]OrderListPlaceOcoResponseResultOrderReportsInner, bool) {
	if o == nil || common.IsNil(o.OrderReports) {
		return nil, false
	}
	return o.OrderReports, true
}

// HasOrderReports returns a boolean if a field has been set.
func (o *OrderListPlaceOcoResponseResult) HasOrderReports() bool {
	if o != nil && !common.IsNil(o.OrderReports) {
		return true
	}

	return false
}

// SetOrderReports gets a reference to the given []OrderListPlaceOcoResponseResultOrderReportsInner and assigns it to the OrderReports field.
func (o *OrderListPlaceOcoResponseResult) SetOrderReports(v []OrderListPlaceOcoResponseResultOrderReportsInner) {
	o.OrderReports = v
}

func (o OrderListPlaceOcoResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderListPlaceOcoResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.OrderListId) {
		toSerialize["orderListId"] = o.OrderListId
	}
	if !common.IsNil(o.ContingencyType) {
		toSerialize["contingencyType"] = o.ContingencyType
	}
	if !common.IsNil(o.ListStatusType) {
		toSerialize["listStatusType"] = o.ListStatusType
	}
	if !common.IsNil(o.ListOrderStatus) {
		toSerialize["listOrderStatus"] = o.ListOrderStatus
	}
	if !common.IsNil(o.ListClientOrderId) {
		toSerialize["listClientOrderId"] = o.ListClientOrderId
	}
	if !common.IsNil(o.TransactionTime) {
		toSerialize["transactionTime"] = o.TransactionTime
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Orders) {
		toSerialize["orders"] = o.Orders
	}
	if !common.IsNil(o.OrderReports) {
		toSerialize["orderReports"] = o.OrderReports
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderListPlaceOcoResponseResult) UnmarshalJSON(data []byte) (err error) {
	varOrderListPlaceOcoResponseResult := _OrderListPlaceOcoResponseResult{}

	err = json.Unmarshal(data, &varOrderListPlaceOcoResponseResult)

	if err != nil {
		return err
	}

	*o = OrderListPlaceOcoResponseResult(varOrderListPlaceOcoResponseResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderListId")
		delete(additionalProperties, "contingencyType")
		delete(additionalProperties, "listStatusType")
		delete(additionalProperties, "listOrderStatus")
		delete(additionalProperties, "listClientOrderId")
		delete(additionalProperties, "transactionTime")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "orders")
		delete(additionalProperties, "orderReports")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrderListPlaceOcoResponseResult struct {
	value *OrderListPlaceOcoResponseResult
	isSet bool
}

func (v NullableOrderListPlaceOcoResponseResult) Get() *OrderListPlaceOcoResponseResult {
	return v.value
}

func (v *NullableOrderListPlaceOcoResponseResult) Set(val *OrderListPlaceOcoResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderListPlaceOcoResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderListPlaceOcoResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderListPlaceOcoResponseResult(val *OrderListPlaceOcoResponseResult) *NullableOrderListPlaceOcoResponseResult {
	return &NullableOrderListPlaceOcoResponseResult{value: val, isSet: true}
}

func (v NullableOrderListPlaceOcoResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderListPlaceOcoResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
