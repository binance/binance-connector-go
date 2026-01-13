# MarginAccountCancelOcoResponse

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
**IsIsolated** | Pointer to **bool** |  | [optional] 
**Orders** | Pointer to [**[]MarginAccountCancelOcoResponseOrdersInner**](MarginAccountCancelOcoResponseOrdersInner.md) |  | [optional] 
**OrderReports** | Pointer to [**[]MarginAccountCancelOcoResponseOrderReportsInner**](MarginAccountCancelOcoResponseOrderReportsInner.md) |  | [optional] 

## Methods

### NewMarginAccountCancelOcoResponse

`func NewMarginAccountCancelOcoResponse() *MarginAccountCancelOcoResponse`

NewMarginAccountCancelOcoResponse instantiates a new MarginAccountCancelOcoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarginAccountCancelOcoResponseWithDefaults

`func NewMarginAccountCancelOcoResponseWithDefaults() *MarginAccountCancelOcoResponse`

NewMarginAccountCancelOcoResponseWithDefaults instantiates a new MarginAccountCancelOcoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderListId

`func (o *MarginAccountCancelOcoResponse) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *MarginAccountCancelOcoResponse) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *MarginAccountCancelOcoResponse) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *MarginAccountCancelOcoResponse) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetContingencyType

`func (o *MarginAccountCancelOcoResponse) GetContingencyType() string`

GetContingencyType returns the ContingencyType field if non-nil, zero value otherwise.

### GetContingencyTypeOk

`func (o *MarginAccountCancelOcoResponse) GetContingencyTypeOk() (*string, bool)`

GetContingencyTypeOk returns a tuple with the ContingencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContingencyType

`func (o *MarginAccountCancelOcoResponse) SetContingencyType(v string)`

SetContingencyType sets ContingencyType field to given value.

### HasContingencyType

`func (o *MarginAccountCancelOcoResponse) HasContingencyType() bool`

HasContingencyType returns a boolean if a field has been set.

### GetListStatusType

`func (o *MarginAccountCancelOcoResponse) GetListStatusType() string`

GetListStatusType returns the ListStatusType field if non-nil, zero value otherwise.

### GetListStatusTypeOk

`func (o *MarginAccountCancelOcoResponse) GetListStatusTypeOk() (*string, bool)`

GetListStatusTypeOk returns a tuple with the ListStatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListStatusType

`func (o *MarginAccountCancelOcoResponse) SetListStatusType(v string)`

SetListStatusType sets ListStatusType field to given value.

### HasListStatusType

`func (o *MarginAccountCancelOcoResponse) HasListStatusType() bool`

HasListStatusType returns a boolean if a field has been set.

### GetListOrderStatus

`func (o *MarginAccountCancelOcoResponse) GetListOrderStatus() string`

GetListOrderStatus returns the ListOrderStatus field if non-nil, zero value otherwise.

### GetListOrderStatusOk

`func (o *MarginAccountCancelOcoResponse) GetListOrderStatusOk() (*string, bool)`

GetListOrderStatusOk returns a tuple with the ListOrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOrderStatus

`func (o *MarginAccountCancelOcoResponse) SetListOrderStatus(v string)`

SetListOrderStatus sets ListOrderStatus field to given value.

### HasListOrderStatus

`func (o *MarginAccountCancelOcoResponse) HasListOrderStatus() bool`

HasListOrderStatus returns a boolean if a field has been set.

### GetListClientOrderId

`func (o *MarginAccountCancelOcoResponse) GetListClientOrderId() string`

GetListClientOrderId returns the ListClientOrderId field if non-nil, zero value otherwise.

### GetListClientOrderIdOk

`func (o *MarginAccountCancelOcoResponse) GetListClientOrderIdOk() (*string, bool)`

GetListClientOrderIdOk returns a tuple with the ListClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListClientOrderId

`func (o *MarginAccountCancelOcoResponse) SetListClientOrderId(v string)`

SetListClientOrderId sets ListClientOrderId field to given value.

### HasListClientOrderId

`func (o *MarginAccountCancelOcoResponse) HasListClientOrderId() bool`

HasListClientOrderId returns a boolean if a field has been set.

### GetTransactionTime

`func (o *MarginAccountCancelOcoResponse) GetTransactionTime() int64`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *MarginAccountCancelOcoResponse) GetTransactionTimeOk() (*int64, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *MarginAccountCancelOcoResponse) SetTransactionTime(v int64)`

SetTransactionTime sets TransactionTime field to given value.

### HasTransactionTime

`func (o *MarginAccountCancelOcoResponse) HasTransactionTime() bool`

HasTransactionTime returns a boolean if a field has been set.

### GetSymbol

`func (o *MarginAccountCancelOcoResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *MarginAccountCancelOcoResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *MarginAccountCancelOcoResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *MarginAccountCancelOcoResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetIsIsolated

`func (o *MarginAccountCancelOcoResponse) GetIsIsolated() bool`

GetIsIsolated returns the IsIsolated field if non-nil, zero value otherwise.

### GetIsIsolatedOk

`func (o *MarginAccountCancelOcoResponse) GetIsIsolatedOk() (*bool, bool)`

GetIsIsolatedOk returns a tuple with the IsIsolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIsolated

`func (o *MarginAccountCancelOcoResponse) SetIsIsolated(v bool)`

SetIsIsolated sets IsIsolated field to given value.

### HasIsIsolated

`func (o *MarginAccountCancelOcoResponse) HasIsIsolated() bool`

HasIsIsolated returns a boolean if a field has been set.

### GetOrders

`func (o *MarginAccountCancelOcoResponse) GetOrders() []MarginAccountCancelOcoResponseOrdersInner`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *MarginAccountCancelOcoResponse) GetOrdersOk() (*[]MarginAccountCancelOcoResponseOrdersInner, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *MarginAccountCancelOcoResponse) SetOrders(v []MarginAccountCancelOcoResponseOrdersInner)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *MarginAccountCancelOcoResponse) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetOrderReports

`func (o *MarginAccountCancelOcoResponse) GetOrderReports() []MarginAccountCancelOcoResponseOrderReportsInner`

GetOrderReports returns the OrderReports field if non-nil, zero value otherwise.

### GetOrderReportsOk

`func (o *MarginAccountCancelOcoResponse) GetOrderReportsOk() (*[]MarginAccountCancelOcoResponseOrderReportsInner, bool)`

GetOrderReportsOk returns a tuple with the OrderReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderReports

`func (o *MarginAccountCancelOcoResponse) SetOrderReports(v []MarginAccountCancelOcoResponseOrderReportsInner)`

SetOrderReports sets OrderReports field to given value.

### HasOrderReports

`func (o *MarginAccountCancelOcoResponse) HasOrderReports() bool`

HasOrderReports returns a boolean if a field has been set.


[[Back to README]](../README.md)


