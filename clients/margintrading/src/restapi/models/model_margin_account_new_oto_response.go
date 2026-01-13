/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the MarginAccountNewOtoResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MarginAccountNewOtoResponse{}

// MarginAccountNewOtoResponse struct for MarginAccountNewOtoResponse
type MarginAccountNewOtoResponse struct {
	OrderListId          *int64                                         `json:"orderListId,omitempty"`
	ContingencyType      *string                                        `json:"contingencyType,omitempty"`
	ListStatusType       *string                                        `json:"listStatusType,omitempty"`
	ListOrderStatus      *string                                        `json:"listOrderStatus,omitempty"`
	ListClientOrderId    *string                                        `json:"listClientOrderId,omitempty"`
	TransactionTime      *int64                                         `json:"transactionTime,omitempty"`
	Symbol               *string                                        `json:"symbol,omitempty"`
	IsIsolated           *bool                                          `json:"isIsolated,omitempty"`
	Orders               []MarginAccountNewOtoResponseOrdersInner       `json:"orders,omitempty"`
	OrderReports         []MarginAccountNewOtoResponseOrderReportsInner `json:"orderReports,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarginAccountNewOtoResponse MarginAccountNewOtoResponse

// NewMarginAccountNewOtoResponse instantiates a new MarginAccountNewOtoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginAccountNewOtoResponse() *MarginAccountNewOtoResponse {
	this := MarginAccountNewOtoResponse{}
	return &this
}

// NewMarginAccountNewOtoResponseWithDefaults instantiates a new MarginAccountNewOtoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginAccountNewOtoResponseWithDefaults() *MarginAccountNewOtoResponse {
	this := MarginAccountNewOtoResponse{}
	return &this
}

// GetOrderListId returns the OrderListId field value if set, zero value otherwise.
func (o *MarginAccountNewOtoResponse) GetOrderListId() int64 {
	if o == nil || common.IsNil(o.OrderListId) {
		var ret int64
		return ret
	}
	return *o.OrderListId
}

// GetOrderListIdOk returns a tuple with the OrderListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountNewOtoResponse) GetOrderListIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderListId) {
		return nil, false
	}
	return o.OrderListId, true
}

// HasOrderListId returns a boolean if a field has been set.
func (o *MarginAccountNewOtoResponse) HasOrderListId() bool {
	if o != nil && !common.IsNil(o.OrderListId) {
		return true
	}

	return false
}

// SetOrderListId gets a reference to the given int64 and assigns it to the OrderListId field.
func (o *MarginAccountNewOtoResponse) SetOrderListId(v int64) {
	o.OrderListId = &v
}

// GetContingencyType returns the ContingencyType field value if set, zero value otherwise.
func (o *MarginAccountNewOtoResponse) GetContingencyType() string {
	if o == nil || common.IsNil(o.ContingencyType) {
		var ret string
		return ret
	}
	return *o.ContingencyType
}

// GetContingencyTypeOk returns a tuple with the ContingencyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountNewOtoResponse) GetContingencyTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ContingencyType) {
		return nil, false
	}
	return o.ContingencyType, true
}

// HasContingencyType returns a boolean if a field has been set.
func (o *MarginAccountNewOtoResponse) HasContingencyType() bool {
	if o != nil && !common.IsNil(o.ContingencyType) {
		return true
	}

	return false
}

// SetContingencyType gets a reference to the given string and assigns it to the ContingencyType field.
func (o *MarginAccountNewOtoResponse) SetContingencyType(v string) {
	o.ContingencyType = &v
}

// GetListStatusType returns the ListStatusType field value if set, zero value otherwise.
func (o *MarginAccountNewOtoResponse) GetListStatusType() string {
	if o == nil || common.IsNil(o.ListStatusType) {
		var ret string
		return ret
	}
	return *o.ListStatusType
}

// GetListStatusTypeOk returns a tuple with the ListStatusType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountNewOtoResponse) GetListStatusTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListStatusType) {
		return nil, false
	}
	return o.ListStatusType, true
}

// HasListStatusType returns a boolean if a field has been set.
func (o *MarginAccountNewOtoResponse) HasListStatusType() bool {
	if o != nil && !common.IsNil(o.ListStatusType) {
		return true
	}

	return false
}

// SetListStatusType gets a reference to the given string and assigns it to the ListStatusType field.
func (o *MarginAccountNewOtoResponse) SetListStatusType(v string) {
	o.ListStatusType = &v
}

// GetListOrderStatus returns the ListOrderStatus field value if set, zero value otherwise.
func (o *MarginAccountNewOtoResponse) GetListOrderStatus() string {
	if o == nil || common.IsNil(o.ListOrderStatus) {
		var ret string
		return ret
	}
	return *o.ListOrderStatus
}

// GetListOrderStatusOk returns a tuple with the ListOrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountNewOtoResponse) GetListOrderStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListOrderStatus) {
		return nil, false
	}
	return o.ListOrderStatus, true
}

// HasListOrderStatus returns a boolean if a field has been set.
func (o *MarginAccountNewOtoResponse) HasListOrderStatus() bool {
	if o != nil && !common.IsNil(o.ListOrderStatus) {
		return true
	}

	return false
}

// SetListOrderStatus gets a reference to the given string and assigns it to the ListOrderStatus field.
func (o *MarginAccountNewOtoResponse) SetListOrderStatus(v string) {
	o.ListOrderStatus = &v
}

// GetListClientOrderId returns the ListClientOrderId field value if set, zero value otherwise.
func (o *MarginAccountNewOtoResponse) GetListClientOrderId() string {
	if o == nil || common.IsNil(o.ListClientOrderId) {
		var ret string
		return ret
	}
	return *o.ListClientOrderId
}

// GetListClientOrderIdOk returns a tuple with the ListClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountNewOtoResponse) GetListClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListClientOrderId) {
		return nil, false
	}
	return o.ListClientOrderId, true
}

// HasListClientOrderId returns a boolean if a field has been set.
func (o *MarginAccountNewOtoResponse) HasListClientOrderId() bool {
	if o != nil && !common.IsNil(o.ListClientOrderId) {
		return true
	}

	return false
}

// SetListClientOrderId gets a reference to the given string and assigns it to the ListClientOrderId field.
func (o *MarginAccountNewOtoResponse) SetListClientOrderId(v string) {
	o.ListClientOrderId = &v
}

// GetTransactionTime returns the TransactionTime field value if set, zero value otherwise.
func (o *MarginAccountNewOtoResponse) GetTransactionTime() int64 {
	if o == nil || common.IsNil(o.TransactionTime) {
		var ret int64
		return ret
	}
	return *o.TransactionTime
}

// GetTransactionTimeOk returns a tuple with the TransactionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountNewOtoResponse) GetTransactionTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TransactionTime) {
		return nil, false
	}
	return o.TransactionTime, true
}

// HasTransactionTime returns a boolean if a field has been set.
func (o *MarginAccountNewOtoResponse) HasTransactionTime() bool {
	if o != nil && !common.IsNil(o.TransactionTime) {
		return true
	}

	return false
}

// SetTransactionTime gets a reference to the given int64 and assigns it to the TransactionTime field.
func (o *MarginAccountNewOtoResponse) SetTransactionTime(v int64) {
	o.TransactionTime = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *MarginAccountNewOtoResponse) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountNewOtoResponse) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *MarginAccountNewOtoResponse) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *MarginAccountNewOtoResponse) SetSymbol(v string) {
	o.Symbol = &v
}

// GetIsIsolated returns the IsIsolated field value if set, zero value otherwise.
func (o *MarginAccountNewOtoResponse) GetIsIsolated() bool {
	if o == nil || common.IsNil(o.IsIsolated) {
		var ret bool
		return ret
	}
	return *o.IsIsolated
}

// GetIsIsolatedOk returns a tuple with the IsIsolated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountNewOtoResponse) GetIsIsolatedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsIsolated) {
		return nil, false
	}
	return o.IsIsolated, true
}

// HasIsIsolated returns a boolean if a field has been set.
func (o *MarginAccountNewOtoResponse) HasIsIsolated() bool {
	if o != nil && !common.IsNil(o.IsIsolated) {
		return true
	}

	return false
}

// SetIsIsolated gets a reference to the given bool and assigns it to the IsIsolated field.
func (o *MarginAccountNewOtoResponse) SetIsIsolated(v bool) {
	o.IsIsolated = &v
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *MarginAccountNewOtoResponse) GetOrders() []MarginAccountNewOtoResponseOrdersInner {
	if o == nil || common.IsNil(o.Orders) {
		var ret []MarginAccountNewOtoResponseOrdersInner
		return ret
	}
	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountNewOtoResponse) GetOrdersOk() ([]MarginAccountNewOtoResponseOrdersInner, bool) {
	if o == nil || common.IsNil(o.Orders) {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *MarginAccountNewOtoResponse) HasOrders() bool {
	if o != nil && !common.IsNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []MarginAccountNewOtoResponseOrdersInner and assigns it to the Orders field.
func (o *MarginAccountNewOtoResponse) SetOrders(v []MarginAccountNewOtoResponseOrdersInner) {
	o.Orders = v
}

// GetOrderReports returns the OrderReports field value if set, zero value otherwise.
func (o *MarginAccountNewOtoResponse) GetOrderReports() []MarginAccountNewOtoResponseOrderReportsInner {
	if o == nil || common.IsNil(o.OrderReports) {
		var ret []MarginAccountNewOtoResponseOrderReportsInner
		return ret
	}
	return o.OrderReports
}

// GetOrderReportsOk returns a tuple with the OrderReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountNewOtoResponse) GetOrderReportsOk() ([]MarginAccountNewOtoResponseOrderReportsInner, bool) {
	if o == nil || common.IsNil(o.OrderReports) {
		return nil, false
	}
	return o.OrderReports, true
}

// HasOrderReports returns a boolean if a field has been set.
func (o *MarginAccountNewOtoResponse) HasOrderReports() bool {
	if o != nil && !common.IsNil(o.OrderReports) {
		return true
	}

	return false
}

// SetOrderReports gets a reference to the given []MarginAccountNewOtoResponseOrderReportsInner and assigns it to the OrderReports field.
func (o *MarginAccountNewOtoResponse) SetOrderReports(v []MarginAccountNewOtoResponseOrderReportsInner) {
	o.OrderReports = v
}

func (o MarginAccountNewOtoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarginAccountNewOtoResponse) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.IsIsolated) {
		toSerialize["isIsolated"] = o.IsIsolated
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

func (o *MarginAccountNewOtoResponse) UnmarshalJSON(data []byte) (err error) {
	varMarginAccountNewOtoResponse := _MarginAccountNewOtoResponse{}

	err = json.Unmarshal(data, &varMarginAccountNewOtoResponse)

	if err != nil {
		return err
	}

	*o = MarginAccountNewOtoResponse(varMarginAccountNewOtoResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderListId")
		delete(additionalProperties, "contingencyType")
		delete(additionalProperties, "listStatusType")
		delete(additionalProperties, "listOrderStatus")
		delete(additionalProperties, "listClientOrderId")
		delete(additionalProperties, "transactionTime")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "isIsolated")
		delete(additionalProperties, "orders")
		delete(additionalProperties, "orderReports")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarginAccountNewOtoResponse struct {
	value *MarginAccountNewOtoResponse
	isSet bool
}

func (v NullableMarginAccountNewOtoResponse) Get() *MarginAccountNewOtoResponse {
	return v.value
}

func (v *NullableMarginAccountNewOtoResponse) Set(val *MarginAccountNewOtoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginAccountNewOtoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginAccountNewOtoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginAccountNewOtoResponse(val *MarginAccountNewOtoResponse) *NullableMarginAccountNewOtoResponse {
	return &NullableMarginAccountNewOtoResponse{value: val, isSet: true}
}

func (v NullableMarginAccountNewOtoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginAccountNewOtoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
