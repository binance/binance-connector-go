/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AllOrderListsResponseResultInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AllOrderListsResponseResultInner{}

// AllOrderListsResponseResultInner struct for AllOrderListsResponseResultInner
type AllOrderListsResponseResultInner struct {
	OrderListId          *int64                                     `json:"orderListId,omitempty"`
	ContingencyType      *string                                    `json:"contingencyType,omitempty"`
	ListStatusType       *string                                    `json:"listStatusType,omitempty"`
	ListOrderStatus      *string                                    `json:"listOrderStatus,omitempty"`
	ListClientOrderId    *string                                    `json:"listClientOrderId,omitempty"`
	TransactionTime      *int64                                     `json:"transactionTime,omitempty"`
	Symbol               *string                                    `json:"symbol,omitempty"`
	Orders               []OrderListCancelResponseResultOrdersInner `json:"orders,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AllOrderListsResponseResultInner AllOrderListsResponseResultInner

// NewAllOrderListsResponseResultInner instantiates a new AllOrderListsResponseResultInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllOrderListsResponseResultInner() *AllOrderListsResponseResultInner {
	this := AllOrderListsResponseResultInner{}
	return &this
}

// NewAllOrderListsResponseResultInnerWithDefaults instantiates a new AllOrderListsResponseResultInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllOrderListsResponseResultInnerWithDefaults() *AllOrderListsResponseResultInner {
	this := AllOrderListsResponseResultInner{}
	return &this
}

// GetOrderListId returns the OrderListId field value if set, zero value otherwise.
func (o *AllOrderListsResponseResultInner) GetOrderListId() int64 {
	if o == nil || common.IsNil(o.OrderListId) {
		var ret int64
		return ret
	}
	return *o.OrderListId
}

// GetOrderListIdOk returns a tuple with the OrderListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllOrderListsResponseResultInner) GetOrderListIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderListId) {
		return nil, false
	}
	return o.OrderListId, true
}

// HasOrderListId returns a boolean if a field has been set.
func (o *AllOrderListsResponseResultInner) HasOrderListId() bool {
	if o != nil && !common.IsNil(o.OrderListId) {
		return true
	}

	return false
}

// SetOrderListId gets a reference to the given int64 and assigns it to the OrderListId field.
func (o *AllOrderListsResponseResultInner) SetOrderListId(v int64) {
	o.OrderListId = &v
}

// GetContingencyType returns the ContingencyType field value if set, zero value otherwise.
func (o *AllOrderListsResponseResultInner) GetContingencyType() string {
	if o == nil || common.IsNil(o.ContingencyType) {
		var ret string
		return ret
	}
	return *o.ContingencyType
}

// GetContingencyTypeOk returns a tuple with the ContingencyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllOrderListsResponseResultInner) GetContingencyTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ContingencyType) {
		return nil, false
	}
	return o.ContingencyType, true
}

// HasContingencyType returns a boolean if a field has been set.
func (o *AllOrderListsResponseResultInner) HasContingencyType() bool {
	if o != nil && !common.IsNil(o.ContingencyType) {
		return true
	}

	return false
}

// SetContingencyType gets a reference to the given string and assigns it to the ContingencyType field.
func (o *AllOrderListsResponseResultInner) SetContingencyType(v string) {
	o.ContingencyType = &v
}

// GetListStatusType returns the ListStatusType field value if set, zero value otherwise.
func (o *AllOrderListsResponseResultInner) GetListStatusType() string {
	if o == nil || common.IsNil(o.ListStatusType) {
		var ret string
		return ret
	}
	return *o.ListStatusType
}

// GetListStatusTypeOk returns a tuple with the ListStatusType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllOrderListsResponseResultInner) GetListStatusTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListStatusType) {
		return nil, false
	}
	return o.ListStatusType, true
}

// HasListStatusType returns a boolean if a field has been set.
func (o *AllOrderListsResponseResultInner) HasListStatusType() bool {
	if o != nil && !common.IsNil(o.ListStatusType) {
		return true
	}

	return false
}

// SetListStatusType gets a reference to the given string and assigns it to the ListStatusType field.
func (o *AllOrderListsResponseResultInner) SetListStatusType(v string) {
	o.ListStatusType = &v
}

// GetListOrderStatus returns the ListOrderStatus field value if set, zero value otherwise.
func (o *AllOrderListsResponseResultInner) GetListOrderStatus() string {
	if o == nil || common.IsNil(o.ListOrderStatus) {
		var ret string
		return ret
	}
	return *o.ListOrderStatus
}

// GetListOrderStatusOk returns a tuple with the ListOrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllOrderListsResponseResultInner) GetListOrderStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListOrderStatus) {
		return nil, false
	}
	return o.ListOrderStatus, true
}

// HasListOrderStatus returns a boolean if a field has been set.
func (o *AllOrderListsResponseResultInner) HasListOrderStatus() bool {
	if o != nil && !common.IsNil(o.ListOrderStatus) {
		return true
	}

	return false
}

// SetListOrderStatus gets a reference to the given string and assigns it to the ListOrderStatus field.
func (o *AllOrderListsResponseResultInner) SetListOrderStatus(v string) {
	o.ListOrderStatus = &v
}

// GetListClientOrderId returns the ListClientOrderId field value if set, zero value otherwise.
func (o *AllOrderListsResponseResultInner) GetListClientOrderId() string {
	if o == nil || common.IsNil(o.ListClientOrderId) {
		var ret string
		return ret
	}
	return *o.ListClientOrderId
}

// GetListClientOrderIdOk returns a tuple with the ListClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllOrderListsResponseResultInner) GetListClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ListClientOrderId) {
		return nil, false
	}
	return o.ListClientOrderId, true
}

// HasListClientOrderId returns a boolean if a field has been set.
func (o *AllOrderListsResponseResultInner) HasListClientOrderId() bool {
	if o != nil && !common.IsNil(o.ListClientOrderId) {
		return true
	}

	return false
}

// SetListClientOrderId gets a reference to the given string and assigns it to the ListClientOrderId field.
func (o *AllOrderListsResponseResultInner) SetListClientOrderId(v string) {
	o.ListClientOrderId = &v
}

// GetTransactionTime returns the TransactionTime field value if set, zero value otherwise.
func (o *AllOrderListsResponseResultInner) GetTransactionTime() int64 {
	if o == nil || common.IsNil(o.TransactionTime) {
		var ret int64
		return ret
	}
	return *o.TransactionTime
}

// GetTransactionTimeOk returns a tuple with the TransactionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllOrderListsResponseResultInner) GetTransactionTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TransactionTime) {
		return nil, false
	}
	return o.TransactionTime, true
}

// HasTransactionTime returns a boolean if a field has been set.
func (o *AllOrderListsResponseResultInner) HasTransactionTime() bool {
	if o != nil && !common.IsNil(o.TransactionTime) {
		return true
	}

	return false
}

// SetTransactionTime gets a reference to the given int64 and assigns it to the TransactionTime field.
func (o *AllOrderListsResponseResultInner) SetTransactionTime(v int64) {
	o.TransactionTime = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *AllOrderListsResponseResultInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllOrderListsResponseResultInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *AllOrderListsResponseResultInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *AllOrderListsResponseResultInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *AllOrderListsResponseResultInner) GetOrders() []OrderListCancelResponseResultOrdersInner {
	if o == nil || common.IsNil(o.Orders) {
		var ret []OrderListCancelResponseResultOrdersInner
		return ret
	}
	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllOrderListsResponseResultInner) GetOrdersOk() ([]OrderListCancelResponseResultOrdersInner, bool) {
	if o == nil || common.IsNil(o.Orders) {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *AllOrderListsResponseResultInner) HasOrders() bool {
	if o != nil && !common.IsNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []OrderListCancelResponseResultOrdersInner and assigns it to the Orders field.
func (o *AllOrderListsResponseResultInner) SetOrders(v []OrderListCancelResponseResultOrdersInner) {
	o.Orders = v
}

func (o AllOrderListsResponseResultInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllOrderListsResponseResultInner) ToMap() (map[string]interface{}, error) {
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

func (o *AllOrderListsResponseResultInner) UnmarshalJSON(data []byte) (err error) {
	varAllOrderListsResponseResultInner := _AllOrderListsResponseResultInner{}

	err = json.Unmarshal(data, &varAllOrderListsResponseResultInner)

	if err != nil {
		return err
	}

	*o = AllOrderListsResponseResultInner(varAllOrderListsResponseResultInner)

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

type NullableAllOrderListsResponseResultInner struct {
	value *AllOrderListsResponseResultInner
	isSet bool
}

func (v NullableAllOrderListsResponseResultInner) Get() *AllOrderListsResponseResultInner {
	return v.value
}

func (v *NullableAllOrderListsResponseResultInner) Set(val *AllOrderListsResponseResultInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAllOrderListsResponseResultInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAllOrderListsResponseResultInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllOrderListsResponseResultInner(val *AllOrderListsResponseResultInner) *NullableAllOrderListsResponseResultInner {
	return &NullableAllOrderListsResponseResultInner{value: val, isSet: true}
}

func (v NullableAllOrderListsResponseResultInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllOrderListsResponseResultInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
