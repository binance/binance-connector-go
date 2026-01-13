# OrderAmendKeepPriorityResponseResultListStatus

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**OrderListId** | Pointer to **int64** |  | [optional] 
**ContingencyType** | Pointer to **string** |  | [optional] 
**ListOrderStatus** | Pointer to **string** |  | [optional] 
**ListClientOrderId** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Orders** | Pointer to [**[]OrderAmendKeepPriorityResponseResultListStatusOrdersInner**](OrderAmendKeepPriorityResponseResultListStatusOrdersInner.md) |  | [optional] 

## Methods

### NewOrderAmendKeepPriorityResponseResultListStatus

`func NewOrderAmendKeepPriorityResponseResultListStatus() *OrderAmendKeepPriorityResponseResultListStatus`

NewOrderAmendKeepPriorityResponseResultListStatus instantiates a new OrderAmendKeepPriorityResponseResultListStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderAmendKeepPriorityResponseResultListStatusWithDefaults

`func NewOrderAmendKeepPriorityResponseResultListStatusWithDefaults() *OrderAmendKeepPriorityResponseResultListStatus`

NewOrderAmendKeepPriorityResponseResultListStatusWithDefaults instantiates a new OrderAmendKeepPriorityResponseResultListStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderListId

`func (o *OrderAmendKeepPriorityResponseResultListStatus) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *OrderAmendKeepPriorityResponseResultListStatus) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *OrderAmendKeepPriorityResponseResultListStatus) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *OrderAmendKeepPriorityResponseResultListStatus) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetContingencyType

`func (o *OrderAmendKeepPriorityResponseResultListStatus) GetContingencyType() string`

GetContingencyType returns the ContingencyType field if non-nil, zero value otherwise.

### GetContingencyTypeOk

`func (o *OrderAmendKeepPriorityResponseResultListStatus) GetContingencyTypeOk() (*string, bool)`

GetContingencyTypeOk returns a tuple with the ContingencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContingencyType

`func (o *OrderAmendKeepPriorityResponseResultListStatus) SetContingencyType(v string)`

SetContingencyType sets ContingencyType field to given value.

### HasContingencyType

`func (o *OrderAmendKeepPriorityResponseResultListStatus) HasContingencyType() bool`

HasContingencyType returns a boolean if a field has been set.

### GetListOrderStatus

`func (o *OrderAmendKeepPriorityResponseResultListStatus) GetListOrderStatus() string`

GetListOrderStatus returns the ListOrderStatus field if non-nil, zero value otherwise.

### GetListOrderStatusOk

`func (o *OrderAmendKeepPriorityResponseResultListStatus) GetListOrderStatusOk() (*string, bool)`

GetListOrderStatusOk returns a tuple with the ListOrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOrderStatus

`func (o *OrderAmendKeepPriorityResponseResultListStatus) SetListOrderStatus(v string)`

SetListOrderStatus sets ListOrderStatus field to given value.

### HasListOrderStatus

`func (o *OrderAmendKeepPriorityResponseResultListStatus) HasListOrderStatus() bool`

HasListOrderStatus returns a boolean if a field has been set.

### GetListClientOrderId

`func (o *OrderAmendKeepPriorityResponseResultListStatus) GetListClientOrderId() string`

GetListClientOrderId returns the ListClientOrderId field if non-nil, zero value otherwise.

### GetListClientOrderIdOk

`func (o *OrderAmendKeepPriorityResponseResultListStatus) GetListClientOrderIdOk() (*string, bool)`

GetListClientOrderIdOk returns a tuple with the ListClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListClientOrderId

`func (o *OrderAmendKeepPriorityResponseResultListStatus) SetListClientOrderId(v string)`

SetListClientOrderId sets ListClientOrderId field to given value.

### HasListClientOrderId

`func (o *OrderAmendKeepPriorityResponseResultListStatus) HasListClientOrderId() bool`

HasListClientOrderId returns a boolean if a field has been set.

### GetSymbol

`func (o *OrderAmendKeepPriorityResponseResultListStatus) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *OrderAmendKeepPriorityResponseResultListStatus) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *OrderAmendKeepPriorityResponseResultListStatus) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *OrderAmendKeepPriorityResponseResultListStatus) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetOrders

`func (o *OrderAmendKeepPriorityResponseResultListStatus) GetOrders() []OrderAmendKeepPriorityResponseResultListStatusOrdersInner`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *OrderAmendKeepPriorityResponseResultListStatus) GetOrdersOk() (*[]OrderAmendKeepPriorityResponseResultListStatusOrdersInner, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *OrderAmendKeepPriorityResponseResultListStatus) SetOrders(v []OrderAmendKeepPriorityResponseResultListStatusOrdersInner)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *OrderAmendKeepPriorityResponseResultListStatus) HasOrders() bool`

HasOrders returns a boolean if a field has been set.


[[Back to README]](../README.md)


