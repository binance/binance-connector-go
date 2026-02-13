/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the OrderListPlaceOpoResponseResult type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderListPlaceOpoResponseResult{}

// OrderListPlaceOpoResponseResult struct for OrderListPlaceOpoResponseResult
type OrderListPlaceOpoResponseResult struct {
	OrderListId          *int64                                             `json:"orderListId,omitempty"`
	ContingencyType      *string                                            `json:"contingencyType,omitempty"`
	ListStatusType       *string                                            `json:"listStatusType,omitempty"`
	ListOrderStatus      *string                                            `json:"listOrderStatus,omitempty"`
	ListClientOrderId    *string                                            `json:"listClientOrderId,omitempty"`
	TransactionTime      *int64                                             `json:"transactionTime,omitempty"`
	Symbol               *string                                            `json:"symbol,omitempty"`
	Orders               []OrderListPlaceOpoResponseResultOrdersInner       `json:"orders,omitempty"`
	OrderReports         []OrderListPlaceOpoResponseResultOrderReportsInner `json:"orderReports,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderListPlaceOpoResponseResult OrderListPlaceOpoResponseResult

// NewOrderListPlaceOpoResponseResult instantiates a new OrderListPlaceOpoResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderListPlaceOpoResponseResult() *OrderListPlaceOpoResponseResult {
	this := OrderListPlaceOpoResponseResult{}
	return &this
}

// NewOrderListPlaceOpoResponseResultWithDefaults instantiates a new OrderListPlaceOpoResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderListPlaceOpoResponseResultWithDefaults() *OrderListPlaceOpoResponseResult {
	this := OrderListPlaceOpoResponseResult{}
	return &this
}

// GetOrderListId returns the OrderListId field value if set, zero value otherwise.
func (o *OrderListPlaceOpoResponseResult) GetOrderListId() int64 {
	if o == nil || common.IsNil(o.OrderListId) {
		var ret int64
		return ret
	}
	return *o.OrderListId
}

// GetOrderListIdOk returns a tuple with the OrderListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOpoResponseResult) GetOrderListIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderListId) {
		return nil, false
	}
	return o.OrderListId, true
}

// HasOrderListId returns a boolean if a field has been set.
func (o *OrderListPlaceOpoResponseResult) HasOrderListId() bool {
	if o != nil && !common.IsNil(o.OrderListId) {
		return true
	}

	return false
}

// SetOrderListId gets a reference to the given int64 and assigns it to the OrderListId field.
func (o *OrderListPlaceOpoResponseResult) SetOrderListId(v int64) {
	o.OrderListId = &v
}

// GetContingencyType returns the ContingencyType field value if set, zero value otherwise.
func (o *OrderListPlaceOpoResponseResult) GetContingencyType() string {
	if o == nil || common.IsNil(o.ContingencyType) {
		var ret string
		return ret
	}
	return *o.ContingencyType
}

// GetContingencyTypeOk returns a tuple with the ContingencyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOpoResponseResult) GetContingencyTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ContingencyType) {
		return nil, false
	}
	return o.ContingencyType, true
}

// HasContingencyType returns a boolean if a field has been set.
func (o *OrderListPlaceOpoResponseResult) HasContingencyType() bool {
	if o != nil && !common.IsNil(o.ContingencyType) {
		return true
	}

	return false
}

// SetContingencyType gets a reference to the given string and assigns it to the ContingencyType field.
func (o *OrderListPlaceOpoResponseResult) SetContingencyType(v string) {
	o.ContingencyType = &v
}

// GetListStatusType returns the ListStatusType field value if set, zero value otherwise.
func (o *OrderListPlaceOpoResponseResult) GetListStatusType() string {
	if o == nil || common.IsNil(o.ListStatusType) {
		var ret string
		return ret
	}
	return *o.ListStatusType
}

// GetListStatusTypeOk returns a tuple with the ListStatusType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOpoResponseResult) GetListStatusTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListStatusType) {
		return nil, false
	}
	return o.ListStatusType, true
}

// HasListStatusType returns a boolean if a field has been set.
func (o *OrderListPlaceOpoResponseResult) HasListStatusType() bool {
	if o != nil && !common.IsNil(o.ListStatusType) {
		return true
	}

	return false
}

// SetListStatusType gets a reference to the given string and assigns it to the ListStatusType field.
func (o *OrderListPlaceOpoResponseResult) SetListStatusType(v string) {
	o.ListStatusType = &v
}

// GetListOrderStatus returns the ListOrderStatus field value if set, zero value otherwise.
func (o *OrderListPlaceOpoResponseResult) GetListOrderStatus() string {
	if o == nil || common.IsNil(o.ListOrderStatus) {
		var ret string
		return ret
	}
	return *o.ListOrderStatus
}

// GetListOrderStatusOk returns a tuple with the ListOrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOpoResponseResult) GetListOrderStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListOrderStatus) {
		return nil, false
	}
	return o.ListOrderStatus, true
}

// HasListOrderStatus returns a boolean if a field has been set.
func (o *OrderListPlaceOpoResponseResult) HasListOrderStatus() bool {
	if o != nil && !common.IsNil(o.ListOrderStatus) {
		return true
	}

	return false
}

// SetListOrderStatus gets a reference to the given string and assigns it to the ListOrderStatus field.
func (o *OrderListPlaceOpoResponseResult) SetListOrderStatus(v string) {
	o.ListOrderStatus = &v
}

// GetListClientOrderId returns the ListClientOrderId field value if set, zero value otherwise.
func (o *OrderListPlaceOpoResponseResult) GetListClientOrderId() string {
	if o == nil || common.IsNil(o.ListClientOrderId) {
		var ret string
		return ret
	}
	return *o.ListClientOrderId
}

// GetListClientOrderIdOk returns a tuple with the ListClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOpoResponseResult) GetListClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListClientOrderId) {
		return nil, false
	}
	return o.ListClientOrderId, true
}

// HasListClientOrderId returns a boolean if a field has been set.
func (o *OrderListPlaceOpoResponseResult) HasListClientOrderId() bool {
	if o != nil && !common.IsNil(o.ListClientOrderId) {
		return true
	}

	return false
}

// SetListClientOrderId gets a reference to the given string and assigns it to the ListClientOrderId field.
func (o *OrderListPlaceOpoResponseResult) SetListClientOrderId(v string) {
	o.ListClientOrderId = &v
}

// GetTransactionTime returns the TransactionTime field value if set, zero value otherwise.
func (o *OrderListPlaceOpoResponseResult) GetTransactionTime() int64 {
	if o == nil || common.IsNil(o.TransactionTime) {
		var ret int64
		return ret
	}
	return *o.TransactionTime
}

// GetTransactionTimeOk returns a tuple with the TransactionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOpoResponseResult) GetTransactionTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TransactionTime) {
		return nil, false
	}
	return o.TransactionTime, true
}

// HasTransactionTime returns a boolean if a field has been set.
func (o *OrderListPlaceOpoResponseResult) HasTransactionTime() bool {
	if o != nil && !common.IsNil(o.TransactionTime) {
		return true
	}

	return false
}

// SetTransactionTime gets a reference to the given int64 and assigns it to the TransactionTime field.
func (o *OrderListPlaceOpoResponseResult) SetTransactionTime(v int64) {
	o.TransactionTime = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *OrderListPlaceOpoResponseResult) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOpoResponseResult) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *OrderListPlaceOpoResponseResult) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *OrderListPlaceOpoResponseResult) SetSymbol(v string) {
	o.Symbol = &v
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *OrderListPlaceOpoResponseResult) GetOrders() []OrderListPlaceOpoResponseResultOrdersInner {
	if o == nil || common.IsNil(o.Orders) {
		var ret []OrderListPlaceOpoResponseResultOrdersInner
		return ret
	}
	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOpoResponseResult) GetOrdersOk() ([]OrderListPlaceOpoResponseResultOrdersInner, bool) {
	if o == nil || common.IsNil(o.Orders) {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *OrderListPlaceOpoResponseResult) HasOrders() bool {
	if o != nil && !common.IsNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []OrderListPlaceOpoResponseResultOrdersInner and assigns it to the Orders field.
func (o *OrderListPlaceOpoResponseResult) SetOrders(v []OrderListPlaceOpoResponseResultOrdersInner) {
	o.Orders = v
}

// GetOrderReports returns the OrderReports field value if set, zero value otherwise.
func (o *OrderListPlaceOpoResponseResult) GetOrderReports() []OrderListPlaceOpoResponseResultOrderReportsInner {
	if o == nil || common.IsNil(o.OrderReports) {
		var ret []OrderListPlaceOpoResponseResultOrderReportsInner
		return ret
	}
	return o.OrderReports
}

// GetOrderReportsOk returns a tuple with the OrderReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOpoResponseResult) GetOrderReportsOk() ([]OrderListPlaceOpoResponseResultOrderReportsInner, bool) {
	if o == nil || common.IsNil(o.OrderReports) {
		return nil, false
	}
	return o.OrderReports, true
}

// HasOrderReports returns a boolean if a field has been set.
func (o *OrderListPlaceOpoResponseResult) HasOrderReports() bool {
	if o != nil && !common.IsNil(o.OrderReports) {
		return true
	}

	return false
}

// SetOrderReports gets a reference to the given []OrderListPlaceOpoResponseResultOrderReportsInner and assigns it to the OrderReports field.
func (o *OrderListPlaceOpoResponseResult) SetOrderReports(v []OrderListPlaceOpoResponseResultOrderReportsInner) {
	o.OrderReports = v
}

func (o OrderListPlaceOpoResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderListPlaceOpoResponseResult) ToMap() (map[string]interface{}, error) {
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

func (o *OrderListPlaceOpoResponseResult) UnmarshalJSON(data []byte) (err error) {
	varOrderListPlaceOpoResponseResult := _OrderListPlaceOpoResponseResult{}

	err = json.Unmarshal(data, &varOrderListPlaceOpoResponseResult)

	if err != nil {
		return err
	}

	*o = OrderListPlaceOpoResponseResult(varOrderListPlaceOpoResponseResult)

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

type NullableOrderListPlaceOpoResponseResult struct {
	value *OrderListPlaceOpoResponseResult
	isSet bool
}

func (v NullableOrderListPlaceOpoResponseResult) Get() *OrderListPlaceOpoResponseResult {
	return v.value
}

func (v *NullableOrderListPlaceOpoResponseResult) Set(val *OrderListPlaceOpoResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderListPlaceOpoResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderListPlaceOpoResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderListPlaceOpoResponseResult(val *OrderListPlaceOpoResponseResult) *NullableOrderListPlaceOpoResponseResult {
	return &NullableOrderListPlaceOpoResponseResult{value: val, isSet: true}
}

func (v NullableOrderListPlaceOpoResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderListPlaceOpoResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
