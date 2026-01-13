/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryMarginAccountsOpenOcoResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryMarginAccountsOpenOcoResponseInner{}

// QueryMarginAccountsOpenOcoResponseInner struct for QueryMarginAccountsOpenOcoResponseInner
type QueryMarginAccountsOpenOcoResponseInner struct {
	OrderListId          *int64                                               `json:"orderListId,omitempty"`
	ContingencyType      *string                                              `json:"contingencyType,omitempty"`
	ListStatusType       *string                                              `json:"listStatusType,omitempty"`
	ListOrderStatus      *string                                              `json:"listOrderStatus,omitempty"`
	ListClientOrderId    *string                                              `json:"listClientOrderId,omitempty"`
	TransactionTime      *int64                                               `json:"transactionTime,omitempty"`
	Symbol               *string                                              `json:"symbol,omitempty"`
	Orders               []QueryMarginAccountsOpenOcoResponseInnerOrdersInner `json:"orders,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryMarginAccountsOpenOcoResponseInner QueryMarginAccountsOpenOcoResponseInner

// NewQueryMarginAccountsOpenOcoResponseInner instantiates a new QueryMarginAccountsOpenOcoResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryMarginAccountsOpenOcoResponseInner() *QueryMarginAccountsOpenOcoResponseInner {
	this := QueryMarginAccountsOpenOcoResponseInner{}
	return &this
}

// NewQueryMarginAccountsOpenOcoResponseInnerWithDefaults instantiates a new QueryMarginAccountsOpenOcoResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryMarginAccountsOpenOcoResponseInnerWithDefaults() *QueryMarginAccountsOpenOcoResponseInner {
	this := QueryMarginAccountsOpenOcoResponseInner{}
	return &this
}

// GetOrderListId returns the OrderListId field value if set, zero value otherwise.
func (o *QueryMarginAccountsOpenOcoResponseInner) GetOrderListId() int64 {
	if o == nil || common.IsNil(o.OrderListId) {
		var ret int64
		return ret
	}
	return *o.OrderListId
}

// GetOrderListIdOk returns a tuple with the OrderListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsOpenOcoResponseInner) GetOrderListIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderListId) {
		return nil, false
	}
	return o.OrderListId, true
}

// HasOrderListId returns a boolean if a field has been set.
func (o *QueryMarginAccountsOpenOcoResponseInner) HasOrderListId() bool {
	if o != nil && !common.IsNil(o.OrderListId) {
		return true
	}

	return false
}

// SetOrderListId gets a reference to the given int64 and assigns it to the OrderListId field.
func (o *QueryMarginAccountsOpenOcoResponseInner) SetOrderListId(v int64) {
	o.OrderListId = &v
}

// GetContingencyType returns the ContingencyType field value if set, zero value otherwise.
func (o *QueryMarginAccountsOpenOcoResponseInner) GetContingencyType() string {
	if o == nil || common.IsNil(o.ContingencyType) {
		var ret string
		return ret
	}
	return *o.ContingencyType
}

// GetContingencyTypeOk returns a tuple with the ContingencyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsOpenOcoResponseInner) GetContingencyTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ContingencyType) {
		return nil, false
	}
	return o.ContingencyType, true
}

// HasContingencyType returns a boolean if a field has been set.
func (o *QueryMarginAccountsOpenOcoResponseInner) HasContingencyType() bool {
	if o != nil && !common.IsNil(o.ContingencyType) {
		return true
	}

	return false
}

// SetContingencyType gets a reference to the given string and assigns it to the ContingencyType field.
func (o *QueryMarginAccountsOpenOcoResponseInner) SetContingencyType(v string) {
	o.ContingencyType = &v
}

// GetListStatusType returns the ListStatusType field value if set, zero value otherwise.
func (o *QueryMarginAccountsOpenOcoResponseInner) GetListStatusType() string {
	if o == nil || common.IsNil(o.ListStatusType) {
		var ret string
		return ret
	}
	return *o.ListStatusType
}

// GetListStatusTypeOk returns a tuple with the ListStatusType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsOpenOcoResponseInner) GetListStatusTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListStatusType) {
		return nil, false
	}
	return o.ListStatusType, true
}

// HasListStatusType returns a boolean if a field has been set.
func (o *QueryMarginAccountsOpenOcoResponseInner) HasListStatusType() bool {
	if o != nil && !common.IsNil(o.ListStatusType) {
		return true
	}

	return false
}

// SetListStatusType gets a reference to the given string and assigns it to the ListStatusType field.
func (o *QueryMarginAccountsOpenOcoResponseInner) SetListStatusType(v string) {
	o.ListStatusType = &v
}

// GetListOrderStatus returns the ListOrderStatus field value if set, zero value otherwise.
func (o *QueryMarginAccountsOpenOcoResponseInner) GetListOrderStatus() string {
	if o == nil || common.IsNil(o.ListOrderStatus) {
		var ret string
		return ret
	}
	return *o.ListOrderStatus
}

// GetListOrderStatusOk returns a tuple with the ListOrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsOpenOcoResponseInner) GetListOrderStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListOrderStatus) {
		return nil, false
	}
	return o.ListOrderStatus, true
}

// HasListOrderStatus returns a boolean if a field has been set.
func (o *QueryMarginAccountsOpenOcoResponseInner) HasListOrderStatus() bool {
	if o != nil && !common.IsNil(o.ListOrderStatus) {
		return true
	}

	return false
}

// SetListOrderStatus gets a reference to the given string and assigns it to the ListOrderStatus field.
func (o *QueryMarginAccountsOpenOcoResponseInner) SetListOrderStatus(v string) {
	o.ListOrderStatus = &v
}

// GetListClientOrderId returns the ListClientOrderId field value if set, zero value otherwise.
func (o *QueryMarginAccountsOpenOcoResponseInner) GetListClientOrderId() string {
	if o == nil || common.IsNil(o.ListClientOrderId) {
		var ret string
		return ret
	}
	return *o.ListClientOrderId
}

// GetListClientOrderIdOk returns a tuple with the ListClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsOpenOcoResponseInner) GetListClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListClientOrderId) {
		return nil, false
	}
	return o.ListClientOrderId, true
}

// HasListClientOrderId returns a boolean if a field has been set.
func (o *QueryMarginAccountsOpenOcoResponseInner) HasListClientOrderId() bool {
	if o != nil && !common.IsNil(o.ListClientOrderId) {
		return true
	}

	return false
}

// SetListClientOrderId gets a reference to the given string and assigns it to the ListClientOrderId field.
func (o *QueryMarginAccountsOpenOcoResponseInner) SetListClientOrderId(v string) {
	o.ListClientOrderId = &v
}

// GetTransactionTime returns the TransactionTime field value if set, zero value otherwise.
func (o *QueryMarginAccountsOpenOcoResponseInner) GetTransactionTime() int64 {
	if o == nil || common.IsNil(o.TransactionTime) {
		var ret int64
		return ret
	}
	return *o.TransactionTime
}

// GetTransactionTimeOk returns a tuple with the TransactionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsOpenOcoResponseInner) GetTransactionTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TransactionTime) {
		return nil, false
	}
	return o.TransactionTime, true
}

// HasTransactionTime returns a boolean if a field has been set.
func (o *QueryMarginAccountsOpenOcoResponseInner) HasTransactionTime() bool {
	if o != nil && !common.IsNil(o.TransactionTime) {
		return true
	}

	return false
}

// SetTransactionTime gets a reference to the given int64 and assigns it to the TransactionTime field.
func (o *QueryMarginAccountsOpenOcoResponseInner) SetTransactionTime(v int64) {
	o.TransactionTime = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *QueryMarginAccountsOpenOcoResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsOpenOcoResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *QueryMarginAccountsOpenOcoResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *QueryMarginAccountsOpenOcoResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *QueryMarginAccountsOpenOcoResponseInner) GetOrders() []QueryMarginAccountsOpenOcoResponseInnerOrdersInner {
	if o == nil || common.IsNil(o.Orders) {
		var ret []QueryMarginAccountsOpenOcoResponseInnerOrdersInner
		return ret
	}
	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsOpenOcoResponseInner) GetOrdersOk() ([]QueryMarginAccountsOpenOcoResponseInnerOrdersInner, bool) {
	if o == nil || common.IsNil(o.Orders) {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *QueryMarginAccountsOpenOcoResponseInner) HasOrders() bool {
	if o != nil && !common.IsNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []QueryMarginAccountsOpenOcoResponseInnerOrdersInner and assigns it to the Orders field.
func (o *QueryMarginAccountsOpenOcoResponseInner) SetOrders(v []QueryMarginAccountsOpenOcoResponseInnerOrdersInner) {
	o.Orders = v
}

func (o QueryMarginAccountsOpenOcoResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryMarginAccountsOpenOcoResponseInner) ToMap() (map[string]interface{}, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryMarginAccountsOpenOcoResponseInner) UnmarshalJSON(data []byte) (err error) {
	varQueryMarginAccountsOpenOcoResponseInner := _QueryMarginAccountsOpenOcoResponseInner{}

	err = json.Unmarshal(data, &varQueryMarginAccountsOpenOcoResponseInner)

	if err != nil {
		return err
	}

	*o = QueryMarginAccountsOpenOcoResponseInner(varQueryMarginAccountsOpenOcoResponseInner)

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
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryMarginAccountsOpenOcoResponseInner struct {
	value *QueryMarginAccountsOpenOcoResponseInner
	isSet bool
}

func (v NullableQueryMarginAccountsOpenOcoResponseInner) Get() *QueryMarginAccountsOpenOcoResponseInner {
	return v.value
}

func (v *NullableQueryMarginAccountsOpenOcoResponseInner) Set(val *QueryMarginAccountsOpenOcoResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryMarginAccountsOpenOcoResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryMarginAccountsOpenOcoResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryMarginAccountsOpenOcoResponseInner(val *QueryMarginAccountsOpenOcoResponseInner) *NullableQueryMarginAccountsOpenOcoResponseInner {
	return &NullableQueryMarginAccountsOpenOcoResponseInner{value: val, isSet: true}
}

func (v NullableQueryMarginAccountsOpenOcoResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryMarginAccountsOpenOcoResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
