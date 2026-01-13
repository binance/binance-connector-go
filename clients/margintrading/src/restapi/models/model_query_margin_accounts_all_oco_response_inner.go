/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryMarginAccountsAllOcoResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryMarginAccountsAllOcoResponseInner{}

// QueryMarginAccountsAllOcoResponseInner struct for QueryMarginAccountsAllOcoResponseInner
type QueryMarginAccountsAllOcoResponseInner struct {
	OrderListId          *int64                                              `json:"orderListId,omitempty"`
	ContingencyType      *string                                             `json:"contingencyType,omitempty"`
	ListStatusType       *string                                             `json:"listStatusType,omitempty"`
	ListOrderStatus      *string                                             `json:"listOrderStatus,omitempty"`
	ListClientOrderId    *string                                             `json:"listClientOrderId,omitempty"`
	TransactionTime      *int64                                              `json:"transactionTime,omitempty"`
	Symbol               *string                                             `json:"symbol,omitempty"`
	IsIsolated           *bool                                               `json:"isIsolated,omitempty"`
	Orders               []QueryMarginAccountsAllOcoResponseInnerOrdersInner `json:"orders,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryMarginAccountsAllOcoResponseInner QueryMarginAccountsAllOcoResponseInner

// NewQueryMarginAccountsAllOcoResponseInner instantiates a new QueryMarginAccountsAllOcoResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryMarginAccountsAllOcoResponseInner() *QueryMarginAccountsAllOcoResponseInner {
	this := QueryMarginAccountsAllOcoResponseInner{}
	return &this
}

// NewQueryMarginAccountsAllOcoResponseInnerWithDefaults instantiates a new QueryMarginAccountsAllOcoResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryMarginAccountsAllOcoResponseInnerWithDefaults() *QueryMarginAccountsAllOcoResponseInner {
	this := QueryMarginAccountsAllOcoResponseInner{}
	return &this
}

// GetOrderListId returns the OrderListId field value if set, zero value otherwise.
func (o *QueryMarginAccountsAllOcoResponseInner) GetOrderListId() int64 {
	if o == nil || common.IsNil(o.OrderListId) {
		var ret int64
		return ret
	}
	return *o.OrderListId
}

// GetOrderListIdOk returns a tuple with the OrderListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsAllOcoResponseInner) GetOrderListIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderListId) {
		return nil, false
	}
	return o.OrderListId, true
}

// HasOrderListId returns a boolean if a field has been set.
func (o *QueryMarginAccountsAllOcoResponseInner) HasOrderListId() bool {
	if o != nil && !common.IsNil(o.OrderListId) {
		return true
	}

	return false
}

// SetOrderListId gets a reference to the given int64 and assigns it to the OrderListId field.
func (o *QueryMarginAccountsAllOcoResponseInner) SetOrderListId(v int64) {
	o.OrderListId = &v
}

// GetContingencyType returns the ContingencyType field value if set, zero value otherwise.
func (o *QueryMarginAccountsAllOcoResponseInner) GetContingencyType() string {
	if o == nil || common.IsNil(o.ContingencyType) {
		var ret string
		return ret
	}
	return *o.ContingencyType
}

// GetContingencyTypeOk returns a tuple with the ContingencyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsAllOcoResponseInner) GetContingencyTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ContingencyType) {
		return nil, false
	}
	return o.ContingencyType, true
}

// HasContingencyType returns a boolean if a field has been set.
func (o *QueryMarginAccountsAllOcoResponseInner) HasContingencyType() bool {
	if o != nil && !common.IsNil(o.ContingencyType) {
		return true
	}

	return false
}

// SetContingencyType gets a reference to the given string and assigns it to the ContingencyType field.
func (o *QueryMarginAccountsAllOcoResponseInner) SetContingencyType(v string) {
	o.ContingencyType = &v
}

// GetListStatusType returns the ListStatusType field value if set, zero value otherwise.
func (o *QueryMarginAccountsAllOcoResponseInner) GetListStatusType() string {
	if o == nil || common.IsNil(o.ListStatusType) {
		var ret string
		return ret
	}
	return *o.ListStatusType
}

// GetListStatusTypeOk returns a tuple with the ListStatusType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsAllOcoResponseInner) GetListStatusTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListStatusType) {
		return nil, false
	}
	return o.ListStatusType, true
}

// HasListStatusType returns a boolean if a field has been set.
func (o *QueryMarginAccountsAllOcoResponseInner) HasListStatusType() bool {
	if o != nil && !common.IsNil(o.ListStatusType) {
		return true
	}

	return false
}

// SetListStatusType gets a reference to the given string and assigns it to the ListStatusType field.
func (o *QueryMarginAccountsAllOcoResponseInner) SetListStatusType(v string) {
	o.ListStatusType = &v
}

// GetListOrderStatus returns the ListOrderStatus field value if set, zero value otherwise.
func (o *QueryMarginAccountsAllOcoResponseInner) GetListOrderStatus() string {
	if o == nil || common.IsNil(o.ListOrderStatus) {
		var ret string
		return ret
	}
	return *o.ListOrderStatus
}

// GetListOrderStatusOk returns a tuple with the ListOrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsAllOcoResponseInner) GetListOrderStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListOrderStatus) {
		return nil, false
	}
	return o.ListOrderStatus, true
}

// HasListOrderStatus returns a boolean if a field has been set.
func (o *QueryMarginAccountsAllOcoResponseInner) HasListOrderStatus() bool {
	if o != nil && !common.IsNil(o.ListOrderStatus) {
		return true
	}

	return false
}

// SetListOrderStatus gets a reference to the given string and assigns it to the ListOrderStatus field.
func (o *QueryMarginAccountsAllOcoResponseInner) SetListOrderStatus(v string) {
	o.ListOrderStatus = &v
}

// GetListClientOrderId returns the ListClientOrderId field value if set, zero value otherwise.
func (o *QueryMarginAccountsAllOcoResponseInner) GetListClientOrderId() string {
	if o == nil || common.IsNil(o.ListClientOrderId) {
		var ret string
		return ret
	}
	return *o.ListClientOrderId
}

// GetListClientOrderIdOk returns a tuple with the ListClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsAllOcoResponseInner) GetListClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListClientOrderId) {
		return nil, false
	}
	return o.ListClientOrderId, true
}

// HasListClientOrderId returns a boolean if a field has been set.
func (o *QueryMarginAccountsAllOcoResponseInner) HasListClientOrderId() bool {
	if o != nil && !common.IsNil(o.ListClientOrderId) {
		return true
	}

	return false
}

// SetListClientOrderId gets a reference to the given string and assigns it to the ListClientOrderId field.
func (o *QueryMarginAccountsAllOcoResponseInner) SetListClientOrderId(v string) {
	o.ListClientOrderId = &v
}

// GetTransactionTime returns the TransactionTime field value if set, zero value otherwise.
func (o *QueryMarginAccountsAllOcoResponseInner) GetTransactionTime() int64 {
	if o == nil || common.IsNil(o.TransactionTime) {
		var ret int64
		return ret
	}
	return *o.TransactionTime
}

// GetTransactionTimeOk returns a tuple with the TransactionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsAllOcoResponseInner) GetTransactionTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TransactionTime) {
		return nil, false
	}
	return o.TransactionTime, true
}

// HasTransactionTime returns a boolean if a field has been set.
func (o *QueryMarginAccountsAllOcoResponseInner) HasTransactionTime() bool {
	if o != nil && !common.IsNil(o.TransactionTime) {
		return true
	}

	return false
}

// SetTransactionTime gets a reference to the given int64 and assigns it to the TransactionTime field.
func (o *QueryMarginAccountsAllOcoResponseInner) SetTransactionTime(v int64) {
	o.TransactionTime = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *QueryMarginAccountsAllOcoResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsAllOcoResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *QueryMarginAccountsAllOcoResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *QueryMarginAccountsAllOcoResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetIsIsolated returns the IsIsolated field value if set, zero value otherwise.
func (o *QueryMarginAccountsAllOcoResponseInner) GetIsIsolated() bool {
	if o == nil || common.IsNil(o.IsIsolated) {
		var ret bool
		return ret
	}
	return *o.IsIsolated
}

// GetIsIsolatedOk returns a tuple with the IsIsolated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsAllOcoResponseInner) GetIsIsolatedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsIsolated) {
		return nil, false
	}
	return o.IsIsolated, true
}

// HasIsIsolated returns a boolean if a field has been set.
func (o *QueryMarginAccountsAllOcoResponseInner) HasIsIsolated() bool {
	if o != nil && !common.IsNil(o.IsIsolated) {
		return true
	}

	return false
}

// SetIsIsolated gets a reference to the given bool and assigns it to the IsIsolated field.
func (o *QueryMarginAccountsAllOcoResponseInner) SetIsIsolated(v bool) {
	o.IsIsolated = &v
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *QueryMarginAccountsAllOcoResponseInner) GetOrders() []QueryMarginAccountsAllOcoResponseInnerOrdersInner {
	if o == nil || common.IsNil(o.Orders) {
		var ret []QueryMarginAccountsAllOcoResponseInnerOrdersInner
		return ret
	}
	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsAllOcoResponseInner) GetOrdersOk() ([]QueryMarginAccountsAllOcoResponseInnerOrdersInner, bool) {
	if o == nil || common.IsNil(o.Orders) {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *QueryMarginAccountsAllOcoResponseInner) HasOrders() bool {
	if o != nil && !common.IsNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []QueryMarginAccountsAllOcoResponseInnerOrdersInner and assigns it to the Orders field.
func (o *QueryMarginAccountsAllOcoResponseInner) SetOrders(v []QueryMarginAccountsAllOcoResponseInnerOrdersInner) {
	o.Orders = v
}

func (o QueryMarginAccountsAllOcoResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryMarginAccountsAllOcoResponseInner) ToMap() (map[string]interface{}, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryMarginAccountsAllOcoResponseInner) UnmarshalJSON(data []byte) (err error) {
	varQueryMarginAccountsAllOcoResponseInner := _QueryMarginAccountsAllOcoResponseInner{}

	err = json.Unmarshal(data, &varQueryMarginAccountsAllOcoResponseInner)

	if err != nil {
		return err
	}

	*o = QueryMarginAccountsAllOcoResponseInner(varQueryMarginAccountsAllOcoResponseInner)

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
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryMarginAccountsAllOcoResponseInner struct {
	value *QueryMarginAccountsAllOcoResponseInner
	isSet bool
}

func (v NullableQueryMarginAccountsAllOcoResponseInner) Get() *QueryMarginAccountsAllOcoResponseInner {
	return v.value
}

func (v *NullableQueryMarginAccountsAllOcoResponseInner) Set(val *QueryMarginAccountsAllOcoResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryMarginAccountsAllOcoResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryMarginAccountsAllOcoResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryMarginAccountsAllOcoResponseInner(val *QueryMarginAccountsAllOcoResponseInner) *NullableQueryMarginAccountsAllOcoResponseInner {
	return &NullableQueryMarginAccountsAllOcoResponseInner{value: val, isSet: true}
}

func (v NullableQueryMarginAccountsAllOcoResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryMarginAccountsAllOcoResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
