/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the OrderListPlaceOpocoResponseResult type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderListPlaceOpocoResponseResult{}

// OrderListPlaceOpocoResponseResult struct for OrderListPlaceOpocoResponseResult
type OrderListPlaceOpocoResponseResult struct {
	OrderListId          *int64                                               `json:"orderListId,omitempty"`
	ContingencyType      *string                                              `json:"contingencyType,omitempty"`
	ListStatusType       *string                                              `json:"listStatusType,omitempty"`
	ListOrderStatus      *string                                              `json:"listOrderStatus,omitempty"`
	ListClientOrderId    *string                                              `json:"listClientOrderId,omitempty"`
	TransactionTime      *int64                                               `json:"transactionTime,omitempty"`
	Symbol               *string                                              `json:"symbol,omitempty"`
	Orders               []OrderListPlaceOpocoResponseResultOrdersInner       `json:"orders,omitempty"`
	OrderReports         []OrderListPlaceOpocoResponseResultOrderReportsInner `json:"orderReports,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderListPlaceOpocoResponseResult OrderListPlaceOpocoResponseResult

// NewOrderListPlaceOpocoResponseResult instantiates a new OrderListPlaceOpocoResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderListPlaceOpocoResponseResult() *OrderListPlaceOpocoResponseResult {
	this := OrderListPlaceOpocoResponseResult{}
	return &this
}

// NewOrderListPlaceOpocoResponseResultWithDefaults instantiates a new OrderListPlaceOpocoResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderListPlaceOpocoResponseResultWithDefaults() *OrderListPlaceOpocoResponseResult {
	this := OrderListPlaceOpocoResponseResult{}
	return &this
}

// GetOrderListId returns the OrderListId field value if set, zero value otherwise.
func (o *OrderListPlaceOpocoResponseResult) GetOrderListId() int64 {
	if o == nil || common.IsNil(o.OrderListId) {
		var ret int64
		return ret
	}
	return *o.OrderListId
}

// GetOrderListIdOk returns a tuple with the OrderListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOpocoResponseResult) GetOrderListIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderListId) {
		return nil, false
	}
	return o.OrderListId, true
}

// HasOrderListId returns a boolean if a field has been set.
func (o *OrderListPlaceOpocoResponseResult) HasOrderListId() bool {
	if o != nil && !common.IsNil(o.OrderListId) {
		return true
	}

	return false
}

// SetOrderListId gets a reference to the given int64 and assigns it to the OrderListId field.
func (o *OrderListPlaceOpocoResponseResult) SetOrderListId(v int64) {
	o.OrderListId = &v
}

// GetContingencyType returns the ContingencyType field value if set, zero value otherwise.
func (o *OrderListPlaceOpocoResponseResult) GetContingencyType() string {
	if o == nil || common.IsNil(o.ContingencyType) {
		var ret string
		return ret
	}
	return *o.ContingencyType
}

// GetContingencyTypeOk returns a tuple with the ContingencyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOpocoResponseResult) GetContingencyTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ContingencyType) {
		return nil, false
	}
	return o.ContingencyType, true
}

// HasContingencyType returns a boolean if a field has been set.
func (o *OrderListPlaceOpocoResponseResult) HasContingencyType() bool {
	if o != nil && !common.IsNil(o.ContingencyType) {
		return true
	}

	return false
}

// SetContingencyType gets a reference to the given string and assigns it to the ContingencyType field.
func (o *OrderListPlaceOpocoResponseResult) SetContingencyType(v string) {
	o.ContingencyType = &v
}

// GetListStatusType returns the ListStatusType field value if set, zero value otherwise.
func (o *OrderListPlaceOpocoResponseResult) GetListStatusType() string {
	if o == nil || common.IsNil(o.ListStatusType) {
		var ret string
		return ret
	}
	return *o.ListStatusType
}

// GetListStatusTypeOk returns a tuple with the ListStatusType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOpocoResponseResult) GetListStatusTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListStatusType) {
		return nil, false
	}
	return o.ListStatusType, true
}

// HasListStatusType returns a boolean if a field has been set.
func (o *OrderListPlaceOpocoResponseResult) HasListStatusType() bool {
	if o != nil && !common.IsNil(o.ListStatusType) {
		return true
	}

	return false
}

// SetListStatusType gets a reference to the given string and assigns it to the ListStatusType field.
func (o *OrderListPlaceOpocoResponseResult) SetListStatusType(v string) {
	o.ListStatusType = &v
}

// GetListOrderStatus returns the ListOrderStatus field value if set, zero value otherwise.
func (o *OrderListPlaceOpocoResponseResult) GetListOrderStatus() string {
	if o == nil || common.IsNil(o.ListOrderStatus) {
		var ret string
		return ret
	}
	return *o.ListOrderStatus
}

// GetListOrderStatusOk returns a tuple with the ListOrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOpocoResponseResult) GetListOrderStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListOrderStatus) {
		return nil, false
	}
	return o.ListOrderStatus, true
}

// HasListOrderStatus returns a boolean if a field has been set.
func (o *OrderListPlaceOpocoResponseResult) HasListOrderStatus() bool {
	if o != nil && !common.IsNil(o.ListOrderStatus) {
		return true
	}

	return false
}

// SetListOrderStatus gets a reference to the given string and assigns it to the ListOrderStatus field.
func (o *OrderListPlaceOpocoResponseResult) SetListOrderStatus(v string) {
	o.ListOrderStatus = &v
}

// GetListClientOrderId returns the ListClientOrderId field value if set, zero value otherwise.
func (o *OrderListPlaceOpocoResponseResult) GetListClientOrderId() string {
	if o == nil || common.IsNil(o.ListClientOrderId) {
		var ret string
		return ret
	}
	return *o.ListClientOrderId
}

// GetListClientOrderIdOk returns a tuple with the ListClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOpocoResponseResult) GetListClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListClientOrderId) {
		return nil, false
	}
	return o.ListClientOrderId, true
}

// HasListClientOrderId returns a boolean if a field has been set.
func (o *OrderListPlaceOpocoResponseResult) HasListClientOrderId() bool {
	if o != nil && !common.IsNil(o.ListClientOrderId) {
		return true
	}

	return false
}

// SetListClientOrderId gets a reference to the given string and assigns it to the ListClientOrderId field.
func (o *OrderListPlaceOpocoResponseResult) SetListClientOrderId(v string) {
	o.ListClientOrderId = &v
}

// GetTransactionTime returns the TransactionTime field value if set, zero value otherwise.
func (o *OrderListPlaceOpocoResponseResult) GetTransactionTime() int64 {
	if o == nil || common.IsNil(o.TransactionTime) {
		var ret int64
		return ret
	}
	return *o.TransactionTime
}

// GetTransactionTimeOk returns a tuple with the TransactionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOpocoResponseResult) GetTransactionTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TransactionTime) {
		return nil, false
	}
	return o.TransactionTime, true
}

// HasTransactionTime returns a boolean if a field has been set.
func (o *OrderListPlaceOpocoResponseResult) HasTransactionTime() bool {
	if o != nil && !common.IsNil(o.TransactionTime) {
		return true
	}

	return false
}

// SetTransactionTime gets a reference to the given int64 and assigns it to the TransactionTime field.
func (o *OrderListPlaceOpocoResponseResult) SetTransactionTime(v int64) {
	o.TransactionTime = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *OrderListPlaceOpocoResponseResult) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOpocoResponseResult) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *OrderListPlaceOpocoResponseResult) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *OrderListPlaceOpocoResponseResult) SetSymbol(v string) {
	o.Symbol = &v
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *OrderListPlaceOpocoResponseResult) GetOrders() []OrderListPlaceOpocoResponseResultOrdersInner {
	if o == nil || common.IsNil(o.Orders) {
		var ret []OrderListPlaceOpocoResponseResultOrdersInner
		return ret
	}
	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOpocoResponseResult) GetOrdersOk() ([]OrderListPlaceOpocoResponseResultOrdersInner, bool) {
	if o == nil || common.IsNil(o.Orders) {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *OrderListPlaceOpocoResponseResult) HasOrders() bool {
	if o != nil && !common.IsNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []OrderListPlaceOpocoResponseResultOrdersInner and assigns it to the Orders field.
func (o *OrderListPlaceOpocoResponseResult) SetOrders(v []OrderListPlaceOpocoResponseResultOrdersInner) {
	o.Orders = v
}

// GetOrderReports returns the OrderReports field value if set, zero value otherwise.
func (o *OrderListPlaceOpocoResponseResult) GetOrderReports() []OrderListPlaceOpocoResponseResultOrderReportsInner {
	if o == nil || common.IsNil(o.OrderReports) {
		var ret []OrderListPlaceOpocoResponseResultOrderReportsInner
		return ret
	}
	return o.OrderReports
}

// GetOrderReportsOk returns a tuple with the OrderReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOpocoResponseResult) GetOrderReportsOk() ([]OrderListPlaceOpocoResponseResultOrderReportsInner, bool) {
	if o == nil || common.IsNil(o.OrderReports) {
		return nil, false
	}
	return o.OrderReports, true
}

// HasOrderReports returns a boolean if a field has been set.
func (o *OrderListPlaceOpocoResponseResult) HasOrderReports() bool {
	if o != nil && !common.IsNil(o.OrderReports) {
		return true
	}

	return false
}

// SetOrderReports gets a reference to the given []OrderListPlaceOpocoResponseResultOrderReportsInner and assigns it to the OrderReports field.
func (o *OrderListPlaceOpocoResponseResult) SetOrderReports(v []OrderListPlaceOpocoResponseResultOrderReportsInner) {
	o.OrderReports = v
}

func (o OrderListPlaceOpocoResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderListPlaceOpocoResponseResult) ToMap() (map[string]interface{}, error) {
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

func (o *OrderListPlaceOpocoResponseResult) UnmarshalJSON(data []byte) (err error) {
	varOrderListPlaceOpocoResponseResult := _OrderListPlaceOpocoResponseResult{}

	err = json.Unmarshal(data, &varOrderListPlaceOpocoResponseResult)

	if err != nil {
		return err
	}

	*o = OrderListPlaceOpocoResponseResult(varOrderListPlaceOpocoResponseResult)

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

type NullableOrderListPlaceOpocoResponseResult struct {
	value *OrderListPlaceOpocoResponseResult
	isSet bool
}

func (v NullableOrderListPlaceOpocoResponseResult) Get() *OrderListPlaceOpocoResponseResult {
	return v.value
}

func (v *NullableOrderListPlaceOpocoResponseResult) Set(val *OrderListPlaceOpocoResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderListPlaceOpocoResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderListPlaceOpocoResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderListPlaceOpocoResponseResult(val *OrderListPlaceOpocoResponseResult) *NullableOrderListPlaceOpocoResponseResult {
	return &NullableOrderListPlaceOpocoResponseResult{value: val, isSet: true}
}

func (v NullableOrderListPlaceOpocoResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderListPlaceOpocoResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
