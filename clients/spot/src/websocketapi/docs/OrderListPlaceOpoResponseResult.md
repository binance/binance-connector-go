# OrderListPlaceOpoResponseResult

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**OrderListId** | Pointer to **int64** |  | [optional] 
**ContingencyType** | Pointer to **string** |  | [optional] 
**ListStatusType** | Pointer to **string** |  | [optional] 
**ListOrderStatus** | Pointer to **string** |  | [optional] 
**ListClientOrderId** | Pointer to **string** |  | [optional] 
**TransactionTime** | Pointer to **int64** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Orders** | Pointer to [**[]OrderListPlaceOpoResponseResultOrdersInner**](OrderListPlaceOpoResponseResultOrdersInner.md) |  | [optional] 
**OrderReports** | Pointer to [**[]OrderListPlaceOpoResponseResultOrderReportsInner**](OrderListPlaceOpoResponseResultOrderReportsInner.md) |  | [optional] 

## Methods

### NewOrderListPlaceOpoResponseResult

`func NewOrderListPlaceOpoResponseResult() *OrderListPlaceOpoResponseResult`

NewOrderListPlaceOpoResponseResult instantiates a new OrderListPlaceOpoResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderListPlaceOpoResponseResultWithDefaults

`func NewOrderListPlaceOpoResponseResultWithDefaults() *OrderListPlaceOpoResponseResult`

NewOrderListPlaceOpoResponseResultWithDefaults instantiates a new OrderListPlaceOpoResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderListId

`func (o *OrderListPlaceOpoResponseResult) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *OrderListPlaceOpoResponseResult) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *OrderListPlaceOpoResponseResult) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *OrderListPlaceOpoResponseResult) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetContingencyType

`func (o *OrderListPlaceOpoResponseResult) GetContingencyType() string`

GetContingencyType returns the ContingencyType field if non-nil, zero value otherwise.

### GetContingencyTypeOk

`func (o *OrderListPlaceOpoResponseResult) GetContingencyTypeOk() (*string, bool)`

GetContingencyTypeOk returns a tuple with the ContingencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContingencyType

`func (o *OrderListPlaceOpoResponseResult) SetContingencyType(v string)`

SetContingencyType sets ContingencyType field to given value.

### HasContingencyType

`func (o *OrderListPlaceOpoResponseResult) HasContingencyType() bool`

HasContingencyType returns a boolean if a field has been set.

### GetListStatusType

`func (o *OrderListPlaceOpoResponseResult) GetListStatusType() string`

GetListStatusType returns the ListStatusType field if non-nil, zero value otherwise.

### GetListStatusTypeOk

`func (o *OrderListPlaceOpoResponseResult) GetListStatusTypeOk() (*string, bool)`

GetListStatusTypeOk returns a tuple with the ListStatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListStatusType

`func (o *OrderListPlaceOpoResponseResult) SetListStatusType(v string)`

SetListStatusType sets ListStatusType field to given value.

### HasListStatusType

`func (o *OrderListPlaceOpoResponseResult) HasListStatusType() bool`

HasListStatusType returns a boolean if a field has been set.

### GetListOrderStatus

`func (o *OrderListPlaceOpoResponseResult) GetListOrderStatus() string`

GetListOrderStatus returns the ListOrderStatus field if non-nil, zero value otherwise.

### GetListOrderStatusOk

`func (o *OrderListPlaceOpoResponseResult) GetListOrderStatusOk() (*string, bool)`

GetListOrderStatusOk returns a tuple with the ListOrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOrderStatus

`func (o *OrderListPlaceOpoResponseResult) SetListOrderStatus(v string)`

SetListOrderStatus sets ListOrderStatus field to given value.

### HasListOrderStatus

`func (o *OrderListPlaceOpoResponseResult) HasListOrderStatus() bool`

HasListOrderStatus returns a boolean if a field has been set.

### GetListClientOrderId

`func (o *OrderListPlaceOpoResponseResult) GetListClientOrderId() string`

GetListClientOrderId returns the ListClientOrderId field if non-nil, zero value otherwise.

### GetListClientOrderIdOk

`func (o *OrderListPlaceOpoResponseResult) GetListClientOrderIdOk() (*string, bool)`

GetListClientOrderIdOk returns a tuple with the ListClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListClientOrderId

`func (o *OrderListPlaceOpoResponseResult) SetListClientOrderId(v string)`

SetListClientOrderId sets ListClientOrderId field to given value.

### HasListClientOrderId

`func (o *OrderListPlaceOpoResponseResult) HasListClientOrderId() bool`

HasListClientOrderId returns a boolean if a field has been set.

### GetTransactionTime

`func (o *OrderListPlaceOpoResponseResult) GetTransactionTime() int64`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *OrderListPlaceOpoResponseResult) GetTransactionTimeOk() (*int64, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *OrderListPlaceOpoResponseResult) SetTransactionTime(v int64)`

SetTransactionTime sets TransactionTime field to given value.

### HasTransactionTime

`func (o *OrderListPlaceOpoResponseResult) HasTransactionTime() bool`

HasTransactionTime returns a boolean if a field has been set.

### GetSymbol

`func (o *OrderListPlaceOpoResponseResult) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *OrderListPlaceOpoResponseResult) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *OrderListPlaceOpoResponseResult) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *OrderListPlaceOpoResponseResult) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetOrders

`func (o *OrderListPlaceOpoResponseResult) GetOrders() []OrderListPlaceOpoResponseResultOrdersInner`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *OrderListPlaceOpoResponseResult) GetOrdersOk() (*[]OrderListPlaceOpoResponseResultOrdersInner, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *OrderListPlaceOpoResponseResult) SetOrders(v []OrderListPlaceOpoResponseResultOrdersInner)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *OrderListPlaceOpoResponseResult) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetOrderReports

`func (o *OrderListPlaceOpoResponseResult) GetOrderReports() []OrderListPlaceOpoResponseResultOrderReportsInner`

GetOrderReports returns the OrderReports field if non-nil, zero value otherwise.

### GetOrderReportsOk

`func (o *OrderListPlaceOpoResponseResult) GetOrderReportsOk() (*[]OrderListPlaceOpoResponseResultOrderReportsInner, bool)`

GetOrderReportsOk returns a tuple with the OrderReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderReports

`func (o *OrderListPlaceOpoResponseResult) SetOrderReports(v []OrderListPlaceOpoResponseResultOrderReportsInner)`

SetOrderReports sets OrderReports field to given value.

### HasOrderReports

`func (o *OrderListPlaceOpoResponseResult) HasOrderReports() bool`

HasOrderReports returns a boolean if a field has been set.


[[Back to README]](../README.md)


