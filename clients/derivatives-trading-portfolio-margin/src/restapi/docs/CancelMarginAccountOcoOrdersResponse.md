# CancelMarginAccountOcoOrdersResponse

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
**Orders** | Pointer to [**[]CancelMarginAccountOcoOrdersResponseOrdersInner**](CancelMarginAccountOcoOrdersResponseOrdersInner.md) |  | [optional] 
**OrderReports** | Pointer to [**[]CancelMarginAccountOcoOrdersResponseOrderReportsInner**](CancelMarginAccountOcoOrdersResponseOrderReportsInner.md) |  | [optional] 

## Methods

### NewCancelMarginAccountOcoOrdersResponse

`func NewCancelMarginAccountOcoOrdersResponse() *CancelMarginAccountOcoOrdersResponse`

NewCancelMarginAccountOcoOrdersResponse instantiates a new CancelMarginAccountOcoOrdersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelMarginAccountOcoOrdersResponseWithDefaults

`func NewCancelMarginAccountOcoOrdersResponseWithDefaults() *CancelMarginAccountOcoOrdersResponse`

NewCancelMarginAccountOcoOrdersResponseWithDefaults instantiates a new CancelMarginAccountOcoOrdersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderListId

`func (o *CancelMarginAccountOcoOrdersResponse) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *CancelMarginAccountOcoOrdersResponse) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *CancelMarginAccountOcoOrdersResponse) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *CancelMarginAccountOcoOrdersResponse) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetContingencyType

`func (o *CancelMarginAccountOcoOrdersResponse) GetContingencyType() string`

GetContingencyType returns the ContingencyType field if non-nil, zero value otherwise.

### GetContingencyTypeOk

`func (o *CancelMarginAccountOcoOrdersResponse) GetContingencyTypeOk() (*string, bool)`

GetContingencyTypeOk returns a tuple with the ContingencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContingencyType

`func (o *CancelMarginAccountOcoOrdersResponse) SetContingencyType(v string)`

SetContingencyType sets ContingencyType field to given value.

### HasContingencyType

`func (o *CancelMarginAccountOcoOrdersResponse) HasContingencyType() bool`

HasContingencyType returns a boolean if a field has been set.

### GetListStatusType

`func (o *CancelMarginAccountOcoOrdersResponse) GetListStatusType() string`

GetListStatusType returns the ListStatusType field if non-nil, zero value otherwise.

### GetListStatusTypeOk

`func (o *CancelMarginAccountOcoOrdersResponse) GetListStatusTypeOk() (*string, bool)`

GetListStatusTypeOk returns a tuple with the ListStatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListStatusType

`func (o *CancelMarginAccountOcoOrdersResponse) SetListStatusType(v string)`

SetListStatusType sets ListStatusType field to given value.

### HasListStatusType

`func (o *CancelMarginAccountOcoOrdersResponse) HasListStatusType() bool`

HasListStatusType returns a boolean if a field has been set.

### GetListOrderStatus

`func (o *CancelMarginAccountOcoOrdersResponse) GetListOrderStatus() string`

GetListOrderStatus returns the ListOrderStatus field if non-nil, zero value otherwise.

### GetListOrderStatusOk

`func (o *CancelMarginAccountOcoOrdersResponse) GetListOrderStatusOk() (*string, bool)`

GetListOrderStatusOk returns a tuple with the ListOrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOrderStatus

`func (o *CancelMarginAccountOcoOrdersResponse) SetListOrderStatus(v string)`

SetListOrderStatus sets ListOrderStatus field to given value.

### HasListOrderStatus

`func (o *CancelMarginAccountOcoOrdersResponse) HasListOrderStatus() bool`

HasListOrderStatus returns a boolean if a field has been set.

### GetListClientOrderId

`func (o *CancelMarginAccountOcoOrdersResponse) GetListClientOrderId() string`

GetListClientOrderId returns the ListClientOrderId field if non-nil, zero value otherwise.

### GetListClientOrderIdOk

`func (o *CancelMarginAccountOcoOrdersResponse) GetListClientOrderIdOk() (*string, bool)`

GetListClientOrderIdOk returns a tuple with the ListClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListClientOrderId

`func (o *CancelMarginAccountOcoOrdersResponse) SetListClientOrderId(v string)`

SetListClientOrderId sets ListClientOrderId field to given value.

### HasListClientOrderId

`func (o *CancelMarginAccountOcoOrdersResponse) HasListClientOrderId() bool`

HasListClientOrderId returns a boolean if a field has been set.

### GetTransactionTime

`func (o *CancelMarginAccountOcoOrdersResponse) GetTransactionTime() int64`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *CancelMarginAccountOcoOrdersResponse) GetTransactionTimeOk() (*int64, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *CancelMarginAccountOcoOrdersResponse) SetTransactionTime(v int64)`

SetTransactionTime sets TransactionTime field to given value.

### HasTransactionTime

`func (o *CancelMarginAccountOcoOrdersResponse) HasTransactionTime() bool`

HasTransactionTime returns a boolean if a field has been set.

### GetSymbol

`func (o *CancelMarginAccountOcoOrdersResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CancelMarginAccountOcoOrdersResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CancelMarginAccountOcoOrdersResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *CancelMarginAccountOcoOrdersResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetOrders

`func (o *CancelMarginAccountOcoOrdersResponse) GetOrders() []CancelMarginAccountOcoOrdersResponseOrdersInner`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *CancelMarginAccountOcoOrdersResponse) GetOrdersOk() (*[]CancelMarginAccountOcoOrdersResponseOrdersInner, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *CancelMarginAccountOcoOrdersResponse) SetOrders(v []CancelMarginAccountOcoOrdersResponseOrdersInner)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *CancelMarginAccountOcoOrdersResponse) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetOrderReports

`func (o *CancelMarginAccountOcoOrdersResponse) GetOrderReports() []CancelMarginAccountOcoOrdersResponseOrderReportsInner`

GetOrderReports returns the OrderReports field if non-nil, zero value otherwise.

### GetOrderReportsOk

`func (o *CancelMarginAccountOcoOrdersResponse) GetOrderReportsOk() (*[]CancelMarginAccountOcoOrdersResponseOrderReportsInner, bool)`

GetOrderReportsOk returns a tuple with the OrderReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderReports

`func (o *CancelMarginAccountOcoOrdersResponse) SetOrderReports(v []CancelMarginAccountOcoOrdersResponseOrderReportsInner)`

SetOrderReports sets OrderReports field to given value.

### HasOrderReports

`func (o *CancelMarginAccountOcoOrdersResponse) HasOrderReports() bool`

HasOrderReports returns a boolean if a field has been set.


[[Back to README]](../README.md)


