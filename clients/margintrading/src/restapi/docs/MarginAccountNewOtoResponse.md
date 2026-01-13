# MarginAccountNewOtoResponse

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
**Orders** | Pointer to [**[]MarginAccountNewOtoResponseOrdersInner**](MarginAccountNewOtoResponseOrdersInner.md) |  | [optional] 
**OrderReports** | Pointer to [**[]MarginAccountNewOtoResponseOrderReportsInner**](MarginAccountNewOtoResponseOrderReportsInner.md) |  | [optional] 

## Methods

### NewMarginAccountNewOtoResponse

`func NewMarginAccountNewOtoResponse() *MarginAccountNewOtoResponse`

NewMarginAccountNewOtoResponse instantiates a new MarginAccountNewOtoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarginAccountNewOtoResponseWithDefaults

`func NewMarginAccountNewOtoResponseWithDefaults() *MarginAccountNewOtoResponse`

NewMarginAccountNewOtoResponseWithDefaults instantiates a new MarginAccountNewOtoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderListId

`func (o *MarginAccountNewOtoResponse) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *MarginAccountNewOtoResponse) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *MarginAccountNewOtoResponse) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *MarginAccountNewOtoResponse) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetContingencyType

`func (o *MarginAccountNewOtoResponse) GetContingencyType() string`

GetContingencyType returns the ContingencyType field if non-nil, zero value otherwise.

### GetContingencyTypeOk

`func (o *MarginAccountNewOtoResponse) GetContingencyTypeOk() (*string, bool)`

GetContingencyTypeOk returns a tuple with the ContingencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContingencyType

`func (o *MarginAccountNewOtoResponse) SetContingencyType(v string)`

SetContingencyType sets ContingencyType field to given value.

### HasContingencyType

`func (o *MarginAccountNewOtoResponse) HasContingencyType() bool`

HasContingencyType returns a boolean if a field has been set.

### GetListStatusType

`func (o *MarginAccountNewOtoResponse) GetListStatusType() string`

GetListStatusType returns the ListStatusType field if non-nil, zero value otherwise.

### GetListStatusTypeOk

`func (o *MarginAccountNewOtoResponse) GetListStatusTypeOk() (*string, bool)`

GetListStatusTypeOk returns a tuple with the ListStatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListStatusType

`func (o *MarginAccountNewOtoResponse) SetListStatusType(v string)`

SetListStatusType sets ListStatusType field to given value.

### HasListStatusType

`func (o *MarginAccountNewOtoResponse) HasListStatusType() bool`

HasListStatusType returns a boolean if a field has been set.

### GetListOrderStatus

`func (o *MarginAccountNewOtoResponse) GetListOrderStatus() string`

GetListOrderStatus returns the ListOrderStatus field if non-nil, zero value otherwise.

### GetListOrderStatusOk

`func (o *MarginAccountNewOtoResponse) GetListOrderStatusOk() (*string, bool)`

GetListOrderStatusOk returns a tuple with the ListOrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOrderStatus

`func (o *MarginAccountNewOtoResponse) SetListOrderStatus(v string)`

SetListOrderStatus sets ListOrderStatus field to given value.

### HasListOrderStatus

`func (o *MarginAccountNewOtoResponse) HasListOrderStatus() bool`

HasListOrderStatus returns a boolean if a field has been set.

### GetListClientOrderId

`func (o *MarginAccountNewOtoResponse) GetListClientOrderId() string`

GetListClientOrderId returns the ListClientOrderId field if non-nil, zero value otherwise.

### GetListClientOrderIdOk

`func (o *MarginAccountNewOtoResponse) GetListClientOrderIdOk() (*string, bool)`

GetListClientOrderIdOk returns a tuple with the ListClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListClientOrderId

`func (o *MarginAccountNewOtoResponse) SetListClientOrderId(v string)`

SetListClientOrderId sets ListClientOrderId field to given value.

### HasListClientOrderId

`func (o *MarginAccountNewOtoResponse) HasListClientOrderId() bool`

HasListClientOrderId returns a boolean if a field has been set.

### GetTransactionTime

`func (o *MarginAccountNewOtoResponse) GetTransactionTime() int64`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *MarginAccountNewOtoResponse) GetTransactionTimeOk() (*int64, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *MarginAccountNewOtoResponse) SetTransactionTime(v int64)`

SetTransactionTime sets TransactionTime field to given value.

### HasTransactionTime

`func (o *MarginAccountNewOtoResponse) HasTransactionTime() bool`

HasTransactionTime returns a boolean if a field has been set.

### GetSymbol

`func (o *MarginAccountNewOtoResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *MarginAccountNewOtoResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *MarginAccountNewOtoResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *MarginAccountNewOtoResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetIsIsolated

`func (o *MarginAccountNewOtoResponse) GetIsIsolated() bool`

GetIsIsolated returns the IsIsolated field if non-nil, zero value otherwise.

### GetIsIsolatedOk

`func (o *MarginAccountNewOtoResponse) GetIsIsolatedOk() (*bool, bool)`

GetIsIsolatedOk returns a tuple with the IsIsolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIsolated

`func (o *MarginAccountNewOtoResponse) SetIsIsolated(v bool)`

SetIsIsolated sets IsIsolated field to given value.

### HasIsIsolated

`func (o *MarginAccountNewOtoResponse) HasIsIsolated() bool`

HasIsIsolated returns a boolean if a field has been set.

### GetOrders

`func (o *MarginAccountNewOtoResponse) GetOrders() []MarginAccountNewOtoResponseOrdersInner`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *MarginAccountNewOtoResponse) GetOrdersOk() (*[]MarginAccountNewOtoResponseOrdersInner, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *MarginAccountNewOtoResponse) SetOrders(v []MarginAccountNewOtoResponseOrdersInner)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *MarginAccountNewOtoResponse) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetOrderReports

`func (o *MarginAccountNewOtoResponse) GetOrderReports() []MarginAccountNewOtoResponseOrderReportsInner`

GetOrderReports returns the OrderReports field if non-nil, zero value otherwise.

### GetOrderReportsOk

`func (o *MarginAccountNewOtoResponse) GetOrderReportsOk() (*[]MarginAccountNewOtoResponseOrderReportsInner, bool)`

GetOrderReportsOk returns a tuple with the OrderReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderReports

`func (o *MarginAccountNewOtoResponse) SetOrderReports(v []MarginAccountNewOtoResponseOrderReportsInner)`

SetOrderReports sets OrderReports field to given value.

### HasOrderReports

`func (o *MarginAccountNewOtoResponse) HasOrderReports() bool`

HasOrderReports returns a boolean if a field has been set.


[[Back to README]](../README.md)


