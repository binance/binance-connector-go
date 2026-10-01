/*
Margin REST API

Access account information, borrow and repay assets, and trade with Binance Margin.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryMarginAccountsOpenOtootocoOrderListsResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryMarginAccountsOpenOtootocoOrderListsResponseInner{}

// QueryMarginAccountsOpenOtootocoOrderListsResponseInner struct for QueryMarginAccountsOpenOtootocoOrderListsResponseInner
type QueryMarginAccountsOpenOtootocoOrderListsResponseInner struct {
	// order List Id.
	OrderListId *int64 `json:"orderListId,omitempty"`
	// contingency Type.
	ContingencyType *string `json:"contingencyType,omitempty"`
	// list Status Type.
	ListStatusType *string `json:"listStatusType,omitempty"`
	// list Order Status.
	ListOrderStatus *string `json:"listOrderStatus,omitempty"`
	// list Client Order Id.
	ListClientOrderId *string `json:"listClientOrderId,omitempty"`
	// transaction Time.
	TransactionTime *int64 `json:"transactionTime,omitempty"`
	// symbol.
	Symbol *string `json:"symbol,omitempty"`
	// orders list.
	Orders               []QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner `json:"orders,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryMarginAccountsOpenOtootocoOrderListsResponseInner QueryMarginAccountsOpenOtootocoOrderListsResponseInner

// NewQueryMarginAccountsOpenOtootocoOrderListsResponseInner instantiates a new QueryMarginAccountsOpenOtootocoOrderListsResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryMarginAccountsOpenOtootocoOrderListsResponseInner() *QueryMarginAccountsOpenOtootocoOrderListsResponseInner {
	this := QueryMarginAccountsOpenOtootocoOrderListsResponseInner{}
	return &this
}

// NewQueryMarginAccountsOpenOtootocoOrderListsResponseInnerWithDefaults instantiates a new QueryMarginAccountsOpenOtootocoOrderListsResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryMarginAccountsOpenOtootocoOrderListsResponseInnerWithDefaults() *QueryMarginAccountsOpenOtootocoOrderListsResponseInner {
	this := QueryMarginAccountsOpenOtootocoOrderListsResponseInner{}
	return &this
}

// GetOrderListId returns the OrderListId field value if set, zero value otherwise.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) GetOrderListId() int64 {
	if o == nil || common.IsNil(o.OrderListId) {
		var ret int64
		return ret
	}
	return *o.OrderListId
}

// GetOrderListIdOk returns a tuple with the OrderListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) GetOrderListIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderListId) {
		return nil, false
	}
	return o.OrderListId, true
}

// HasOrderListId returns a boolean if a field has been set.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) HasOrderListId() bool {
	if o != nil && !common.IsNil(o.OrderListId) {
		return true
	}

	return false
}

// SetOrderListId gets a reference to the given int64 and assigns it to the OrderListId field.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) SetOrderListId(v int64) {
	o.OrderListId = &v
}

// GetContingencyType returns the ContingencyType field value if set, zero value otherwise.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) GetContingencyType() string {
	if o == nil || common.IsNil(o.ContingencyType) {
		var ret string
		return ret
	}
	return *o.ContingencyType
}

// GetContingencyTypeOk returns a tuple with the ContingencyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) GetContingencyTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ContingencyType) {
		return nil, false
	}
	return o.ContingencyType, true
}

// HasContingencyType returns a boolean if a field has been set.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) HasContingencyType() bool {
	if o != nil && !common.IsNil(o.ContingencyType) {
		return true
	}

	return false
}

// SetContingencyType gets a reference to the given string and assigns it to the ContingencyType field.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) SetContingencyType(v string) {
	o.ContingencyType = &v
}

// GetListStatusType returns the ListStatusType field value if set, zero value otherwise.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) GetListStatusType() string {
	if o == nil || common.IsNil(o.ListStatusType) {
		var ret string
		return ret
	}
	return *o.ListStatusType
}

// GetListStatusTypeOk returns a tuple with the ListStatusType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) GetListStatusTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListStatusType) {
		return nil, false
	}
	return o.ListStatusType, true
}

// HasListStatusType returns a boolean if a field has been set.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) HasListStatusType() bool {
	if o != nil && !common.IsNil(o.ListStatusType) {
		return true
	}

	return false
}

// SetListStatusType gets a reference to the given string and assigns it to the ListStatusType field.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) SetListStatusType(v string) {
	o.ListStatusType = &v
}

// GetListOrderStatus returns the ListOrderStatus field value if set, zero value otherwise.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) GetListOrderStatus() string {
	if o == nil || common.IsNil(o.ListOrderStatus) {
		var ret string
		return ret
	}
	return *o.ListOrderStatus
}

// GetListOrderStatusOk returns a tuple with the ListOrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) GetListOrderStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListOrderStatus) {
		return nil, false
	}
	return o.ListOrderStatus, true
}

// HasListOrderStatus returns a boolean if a field has been set.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) HasListOrderStatus() bool {
	if o != nil && !common.IsNil(o.ListOrderStatus) {
		return true
	}

	return false
}

// SetListOrderStatus gets a reference to the given string and assigns it to the ListOrderStatus field.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) SetListOrderStatus(v string) {
	o.ListOrderStatus = &v
}

// GetListClientOrderId returns the ListClientOrderId field value if set, zero value otherwise.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) GetListClientOrderId() string {
	if o == nil || common.IsNil(o.ListClientOrderId) {
		var ret string
		return ret
	}
	return *o.ListClientOrderId
}

// GetListClientOrderIdOk returns a tuple with the ListClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) GetListClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListClientOrderId) {
		return nil, false
	}
	return o.ListClientOrderId, true
}

// HasListClientOrderId returns a boolean if a field has been set.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) HasListClientOrderId() bool {
	if o != nil && !common.IsNil(o.ListClientOrderId) {
		return true
	}

	return false
}

// SetListClientOrderId gets a reference to the given string and assigns it to the ListClientOrderId field.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) SetListClientOrderId(v string) {
	o.ListClientOrderId = &v
}

// GetTransactionTime returns the TransactionTime field value if set, zero value otherwise.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) GetTransactionTime() int64 {
	if o == nil || common.IsNil(o.TransactionTime) {
		var ret int64
		return ret
	}
	return *o.TransactionTime
}

// GetTransactionTimeOk returns a tuple with the TransactionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) GetTransactionTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TransactionTime) {
		return nil, false
	}
	return o.TransactionTime, true
}

// HasTransactionTime returns a boolean if a field has been set.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) HasTransactionTime() bool {
	if o != nil && !common.IsNil(o.TransactionTime) {
		return true
	}

	return false
}

// SetTransactionTime gets a reference to the given int64 and assigns it to the TransactionTime field.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) SetTransactionTime(v int64) {
	o.TransactionTime = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) GetOrders() []QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner {
	if o == nil || common.IsNil(o.Orders) {
		var ret []QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner
		return ret
	}
	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) GetOrdersOk() ([]QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner, bool) {
	if o == nil || common.IsNil(o.Orders) {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) HasOrders() bool {
	if o != nil && !common.IsNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner and assigns it to the Orders field.
func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) SetOrders(v []QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) {
	o.Orders = v
}

func (o QueryMarginAccountsOpenOtootocoOrderListsResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryMarginAccountsOpenOtootocoOrderListsResponseInner) ToMap() (map[string]interface{}, error) {
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

func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) UnmarshalJSON(data []byte) (err error) {
	varQueryMarginAccountsOpenOtootocoOrderListsResponseInner := _QueryMarginAccountsOpenOtootocoOrderListsResponseInner{}

	err = json.Unmarshal(data, &varQueryMarginAccountsOpenOtootocoOrderListsResponseInner)

	if err != nil {
		return err
	}

	*o = QueryMarginAccountsOpenOtootocoOrderListsResponseInner(varQueryMarginAccountsOpenOtootocoOrderListsResponseInner)

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

type NullableQueryMarginAccountsOpenOtootocoOrderListsResponseInner struct {
	value *QueryMarginAccountsOpenOtootocoOrderListsResponseInner
	isSet bool
}

func (v NullableQueryMarginAccountsOpenOtootocoOrderListsResponseInner) Get() *QueryMarginAccountsOpenOtootocoOrderListsResponseInner {
	return v.value
}

func (v *NullableQueryMarginAccountsOpenOtootocoOrderListsResponseInner) Set(val *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryMarginAccountsOpenOtootocoOrderListsResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryMarginAccountsOpenOtootocoOrderListsResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryMarginAccountsOpenOtootocoOrderListsResponseInner(val *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) *NullableQueryMarginAccountsOpenOtootocoOrderListsResponseInner {
	return &NullableQueryMarginAccountsOpenOtootocoOrderListsResponseInner{value: val, isSet: true}
}

func (v NullableQueryMarginAccountsOpenOtootocoOrderListsResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryMarginAccountsOpenOtootocoOrderListsResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
