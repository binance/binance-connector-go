# OrderAmendKeepPriorityResponseListStatus

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**OrderListId** | Pointer to **int64** |  | [optional] 
**ContingencyType** | Pointer to **string** |  | [optional] 
**ListOrderStatus** | Pointer to **string** |  | [optional] 
**ListClientOrderId** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Orders** | Pointer to [**[]OrderAmendKeepPriorityResponseListStatusOrdersInner**](OrderAmendKeepPriorityResponseListStatusOrdersInner.md) |  | [optional] 

## Methods

### NewOrderAmendKeepPriorityResponseListStatus

`func NewOrderAmendKeepPriorityResponseListStatus() *OrderAmendKeepPriorityResponseListStatus`

NewOrderAmendKeepPriorityResponseListStatus instantiates a new OrderAmendKeepPriorityResponseListStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderAmendKeepPriorityResponseListStatusWithDefaults

`func NewOrderAmendKeepPriorityResponseListStatusWithDefaults() *OrderAmendKeepPriorityResponseListStatus`

NewOrderAmendKeepPriorityResponseListStatusWithDefaults instantiates a new OrderAmendKeepPriorityResponseListStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderListId

`func (o *OrderAmendKeepPriorityResponseListStatus) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *OrderAmendKeepPriorityResponseListStatus) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *OrderAmendKeepPriorityResponseListStatus) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *OrderAmendKeepPriorityResponseListStatus) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetContingencyType

`func (o *OrderAmendKeepPriorityResponseListStatus) GetContingencyType() string`

GetContingencyType returns the ContingencyType field if non-nil, zero value otherwise.

### GetContingencyTypeOk

`func (o *OrderAmendKeepPriorityResponseListStatus) GetContingencyTypeOk() (*string, bool)`

GetContingencyTypeOk returns a tuple with the ContingencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContingencyType

`func (o *OrderAmendKeepPriorityResponseListStatus) SetContingencyType(v string)`

SetContingencyType sets ContingencyType field to given value.

### HasContingencyType

`func (o *OrderAmendKeepPriorityResponseListStatus) HasContingencyType() bool`

HasContingencyType returns a boolean if a field has been set.

### GetListOrderStatus

`func (o *OrderAmendKeepPriorityResponseListStatus) GetListOrderStatus() string`

GetListOrderStatus returns the ListOrderStatus field if non-nil, zero value otherwise.

### GetListOrderStatusOk

`func (o *OrderAmendKeepPriorityResponseListStatus) GetListOrderStatusOk() (*string, bool)`

GetListOrderStatusOk returns a tuple with the ListOrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOrderStatus

`func (o *OrderAmendKeepPriorityResponseListStatus) SetListOrderStatus(v string)`

SetListOrderStatus sets ListOrderStatus field to given value.

### HasListOrderStatus

`func (o *OrderAmendKeepPriorityResponseListStatus) HasListOrderStatus() bool`

HasListOrderStatus returns a boolean if a field has been set.

### GetListClientOrderId

`func (o *OrderAmendKeepPriorityResponseListStatus) GetListClientOrderId() string`

GetListClientOrderId returns the ListClientOrderId field if non-nil, zero value otherwise.

### GetListClientOrderIdOk

`func (o *OrderAmendKeepPriorityResponseListStatus) GetListClientOrderIdOk() (*string, bool)`

GetListClientOrderIdOk returns a tuple with the ListClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListClientOrderId

`func (o *OrderAmendKeepPriorityResponseListStatus) SetListClientOrderId(v string)`

SetListClientOrderId sets ListClientOrderId field to given value.

### HasListClientOrderId

`func (o *OrderAmendKeepPriorityResponseListStatus) HasListClientOrderId() bool`

HasListClientOrderId returns a boolean if a field has been set.

### GetSymbol

`func (o *OrderAmendKeepPriorityResponseListStatus) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *OrderAmendKeepPriorityResponseListStatus) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *OrderAmendKeepPriorityResponseListStatus) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *OrderAmendKeepPriorityResponseListStatus) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetOrders

`func (o *OrderAmendKeepPriorityResponseListStatus) GetOrders() []OrderAmendKeepPriorityResponseListStatusOrdersInner`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *OrderAmendKeepPriorityResponseListStatus) GetOrdersOk() (*[]OrderAmendKeepPriorityResponseListStatusOrdersInner, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *OrderAmendKeepPriorityResponseListStatus) SetOrders(v []OrderAmendKeepPriorityResponseListStatusOrdersInner)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *OrderAmendKeepPriorityResponseListStatus) HasOrders() bool`

HasOrders returns a boolean if a field has been set.


[[Back to README]](../README.md)


