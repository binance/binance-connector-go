/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the OrderListOcoResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderListOcoResponse{}

// OrderListOcoResponse struct for OrderListOcoResponse
type OrderListOcoResponse struct {
	OrderListId          *int64                                  `json:"orderListId,omitempty"`
	ContingencyType      *string                                 `json:"contingencyType,omitempty"`
	ListStatusType       *string                                 `json:"listStatusType,omitempty"`
	ListOrderStatus      *string                                 `json:"listOrderStatus,omitempty"`
	ListClientOrderId    *string                                 `json:"listClientOrderId,omitempty"`
	TransactionTime      *int64                                  `json:"transactionTime,omitempty"`
	Symbol               *string                                 `json:"symbol,omitempty"`
	Orders               []OrderListOcoResponseOrdersInner       `json:"orders,omitempty"`
	OrderReports         []OrderListOcoResponseOrderReportsInner `json:"orderReports,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderListOcoResponse OrderListOcoResponse

// NewOrderListOcoResponse instantiates a new OrderListOcoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderListOcoResponse() *OrderListOcoResponse {
	this := OrderListOcoResponse{}
	return &this
}

// NewOrderListOcoResponseWithDefaults instantiates a new OrderListOcoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderListOcoResponseWithDefaults() *OrderListOcoResponse {
	this := OrderListOcoResponse{}
	return &this
}

// GetOrderListId returns the OrderListId field value if set, zero value otherwise.
func (o *OrderListOcoResponse) GetOrderListId() int64 {
	if o == nil || common.IsNil(o.OrderListId) {
		var ret int64
		return ret
	}
	return *o.OrderListId
}

// GetOrderListIdOk returns a tuple with the OrderListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListOcoResponse) GetOrderListIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderListId) {
		return nil, false
	}
	return o.OrderListId, true
}

// HasOrderListId returns a boolean if a field has been set.
func (o *OrderListOcoResponse) HasOrderListId() bool {
	if o != nil && !common.IsNil(o.OrderListId) {
		return true
	}

	return false
}

// SetOrderListId gets a reference to the given int64 and assigns it to the OrderListId field.
func (o *OrderListOcoResponse) SetOrderListId(v int64) {
	o.OrderListId = &v
}

// GetContingencyType returns the ContingencyType field value if set, zero value otherwise.
func (o *OrderListOcoResponse) GetContingencyType() string {
	if o == nil || common.IsNil(o.ContingencyType) {
		var ret string
		return ret
	}
	return *o.ContingencyType
}

// GetContingencyTypeOk returns a tuple with the ContingencyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListOcoResponse) GetContingencyTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ContingencyType) {
		return nil, false
	}
	return o.ContingencyType, true
}

// HasContingencyType returns a boolean if a field has been set.
func (o *OrderListOcoResponse) HasContingencyType() bool {
	if o != nil && !common.IsNil(o.ContingencyType) {
		return true
	}

	return false
}

// SetContingencyType gets a reference to the given string and assigns it to the ContingencyType field.
func (o *OrderListOcoResponse) SetContingencyType(v string) {
	o.ContingencyType = &v
}

// GetListStatusType returns the ListStatusType field value if set, zero value otherwise.
func (o *OrderListOcoResponse) GetListStatusType() string {
	if o == nil || common.IsNil(o.ListStatusType) {
		var ret string
		return ret
	}
	return *o.ListStatusType
}

// GetListStatusTypeOk returns a tuple with the ListStatusType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListOcoResponse) GetListStatusTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListStatusType) {
		return nil, false
	}
	return o.ListStatusType, true
}

// HasListStatusType returns a boolean if a field has been set.
func (o *OrderListOcoResponse) HasListStatusType() bool {
	if o != nil && !common.IsNil(o.ListStatusType) {
		return true
	}

	return false
}

// SetListStatusType gets a reference to the given string and assigns it to the ListStatusType field.
func (o *OrderListOcoResponse) SetListStatusType(v string) {
	o.ListStatusType = &v
}

// GetListOrderStatus returns the ListOrderStatus field value if set, zero value otherwise.
func (o *OrderListOcoResponse) GetListOrderStatus() string {
	if o == nil || common.IsNil(o.ListOrderStatus) {
		var ret string
		return ret
	}
	return *o.ListOrderStatus
}

// GetListOrderStatusOk returns a tuple with the ListOrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListOcoResponse) GetListOrderStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListOrderStatus) {
		return nil, false
	}
	return o.ListOrderStatus, true
}

// HasListOrderStatus returns a boolean if a field has been set.
func (o *OrderListOcoResponse) HasListOrderStatus() bool {
	if o != nil && !common.IsNil(o.ListOrderStatus) {
		return true
	}

	return false
}

// SetListOrderStatus gets a reference to the given string and assigns it to the ListOrderStatus field.
func (o *OrderListOcoResponse) SetListOrderStatus(v string) {
	o.ListOrderStatus = &v
}

// GetListClientOrderId returns the ListClientOrderId field value if set, zero value otherwise.
func (o *OrderListOcoResponse) GetListClientOrderId() string {
	if o == nil || common.IsNil(o.ListClientOrderId) {
		var ret string
		return ret
	}
	return *o.ListClientOrderId
}

// GetListClientOrderIdOk returns a tuple with the ListClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListOcoResponse) GetListClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListClientOrderId) {
		return nil, false
	}
	return o.ListClientOrderId, true
}

// HasListClientOrderId returns a boolean if a field has been set.
func (o *OrderListOcoResponse) HasListClientOrderId() bool {
	if o != nil && !common.IsNil(o.ListClientOrderId) {
		return true
	}

	return false
}

// SetListClientOrderId gets a reference to the given string and assigns it to the ListClientOrderId field.
func (o *OrderListOcoResponse) SetListClientOrderId(v string) {
	o.ListClientOrderId = &v
}

// GetTransactionTime returns the TransactionTime field value if set, zero value otherwise.
func (o *OrderListOcoResponse) GetTransactionTime() int64 {
	if o == nil || common.IsNil(o.TransactionTime) {
		var ret int64
		return ret
	}
	return *o.TransactionTime
}

// GetTransactionTimeOk returns a tuple with the TransactionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListOcoResponse) GetTransactionTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TransactionTime) {
		return nil, false
	}
	return o.TransactionTime, true
}

// HasTransactionTime returns a boolean if a field has been set.
func (o *OrderListOcoResponse) HasTransactionTime() bool {
	if o != nil && !common.IsNil(o.TransactionTime) {
		return true
	}

	return false
}

// SetTransactionTime gets a reference to the given int64 and assigns it to the TransactionTime field.
func (o *OrderListOcoResponse) SetTransactionTime(v int64) {
	o.TransactionTime = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *OrderListOcoResponse) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListOcoResponse) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *OrderListOcoResponse) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *OrderListOcoResponse) SetSymbol(v string) {
	o.Symbol = &v
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *OrderListOcoResponse) GetOrders() []OrderListOcoResponseOrdersInner {
	if o == nil || common.IsNil(o.Orders) {
		var ret []OrderListOcoResponseOrdersInner
		return ret
	}
	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListOcoResponse) GetOrdersOk() ([]OrderListOcoResponseOrdersInner, bool) {
	if o == nil || common.IsNil(o.Orders) {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *OrderListOcoResponse) HasOrders() bool {
	if o != nil && !common.IsNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []OrderListOcoResponseOrdersInner and assigns it to the Orders field.
func (o *OrderListOcoResponse) SetOrders(v []OrderListOcoResponseOrdersInner) {
	o.Orders = v
}

// GetOrderReports returns the OrderReports field value if set, zero value otherwise.
func (o *OrderListOcoResponse) GetOrderReports() []OrderListOcoResponseOrderReportsInner {
	if o == nil || common.IsNil(o.OrderReports) {
		var ret []OrderListOcoResponseOrderReportsInner
		return ret
	}
	return o.OrderReports
}

// GetOrderReportsOk returns a tuple with the OrderReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListOcoResponse) GetOrderReportsOk() ([]OrderListOcoResponseOrderReportsInner, bool) {
	if o == nil || common.IsNil(o.OrderReports) {
		return nil, false
	}
	return o.OrderReports, true
}

// HasOrderReports returns a boolean if a field has been set.
func (o *OrderListOcoResponse) HasOrderReports() bool {
	if o != nil && !common.IsNil(o.OrderReports) {
		return true
	}

	return false
}

// SetOrderReports gets a reference to the given []OrderListOcoResponseOrderReportsInner and assigns it to the OrderReports field.
func (o *OrderListOcoResponse) SetOrderReports(v []OrderListOcoResponseOrderReportsInner) {
	o.OrderReports = v
}

func (o OrderListOcoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderListOcoResponse) ToMap() (map[string]interface{}, error) {
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

func (o *OrderListOcoResponse) UnmarshalJSON(data []byte) (err error) {
	varOrderListOcoResponse := _OrderListOcoResponse{}

	err = json.Unmarshal(data, &varOrderListOcoResponse)

	if err != nil {
		return err
	}

	*o = OrderListOcoResponse(varOrderListOcoResponse)

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

type NullableOrderListOcoResponse struct {
	value *OrderListOcoResponse
	isSet bool
}

func (v NullableOrderListOcoResponse) Get() *OrderListOcoResponse {
	return v.value
}

func (v *NullableOrderListOcoResponse) Set(val *OrderListOcoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderListOcoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderListOcoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderListOcoResponse(val *OrderListOcoResponse) *NullableOrderListOcoResponse {
	return &NullableOrderListOcoResponse{value: val, isSet: true}
}

func (v NullableOrderListOcoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderListOcoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
