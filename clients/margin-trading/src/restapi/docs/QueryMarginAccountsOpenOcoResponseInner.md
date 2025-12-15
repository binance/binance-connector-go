# QueryMarginAccountsOpenOcoResponseInner

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
**Orders** | Pointer to [**[]QueryMarginAccountsOpenOcoResponseInnerOrdersInner**](QueryMarginAccountsOpenOcoResponseInnerOrdersInner.md) |  | [optional] 

## Methods

### NewQueryMarginAccountsOpenOcoResponseInner

`func NewQueryMarginAccountsOpenOcoResponseInner() *QueryMarginAccountsOpenOcoResponseInner`

NewQueryMarginAccountsOpenOcoResponseInner instantiates a new QueryMarginAccountsOpenOcoResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryMarginAccountsOpenOcoResponseInnerWithDefaults

`func NewQueryMarginAccountsOpenOcoResponseInnerWithDefaults() *QueryMarginAccountsOpenOcoResponseInner`

NewQueryMarginAccountsOpenOcoResponseInnerWithDefaults instantiates a new QueryMarginAccountsOpenOcoResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderListId

`func (o *QueryMarginAccountsOpenOcoResponseInner) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *QueryMarginAccountsOpenOcoResponseInner) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *QueryMarginAccountsOpenOcoResponseInner) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *QueryMarginAccountsOpenOcoResponseInner) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetContingencyType

`func (o *QueryMarginAccountsOpenOcoResponseInner) GetContingencyType() string`

GetContingencyType returns the ContingencyType field if non-nil, zero value otherwise.

### GetContingencyTypeOk

`func (o *QueryMarginAccountsOpenOcoResponseInner) GetContingencyTypeOk() (*string, bool)`

GetContingencyTypeOk returns a tuple with the ContingencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContingencyType

`func (o *QueryMarginAccountsOpenOcoResponseInner) SetContingencyType(v string)`

SetContingencyType sets ContingencyType field to given value.

### HasContingencyType

`func (o *QueryMarginAccountsOpenOcoResponseInner) HasContingencyType() bool`

HasContingencyType returns a boolean if a field has been set.

### GetListStatusType

`func (o *QueryMarginAccountsOpenOcoResponseInner) GetListStatusType() string`

GetListStatusType returns the ListStatusType field if non-nil, zero value otherwise.

### GetListStatusTypeOk

`func (o *QueryMarginAccountsOpenOcoResponseInner) GetListStatusTypeOk() (*string, bool)`

GetListStatusTypeOk returns a tuple with the ListStatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListStatusType

`func (o *QueryMarginAccountsOpenOcoResponseInner) SetListStatusType(v string)`

SetListStatusType sets ListStatusType field to given value.

### HasListStatusType

`func (o *QueryMarginAccountsOpenOcoResponseInner) HasListStatusType() bool`

HasListStatusType returns a boolean if a field has been set.

### GetListOrderStatus

`func (o *QueryMarginAccountsOpenOcoResponseInner) GetListOrderStatus() string`

GetListOrderStatus returns the ListOrderStatus field if non-nil, zero value otherwise.

### GetListOrderStatusOk

`func (o *QueryMarginAccountsOpenOcoResponseInner) GetListOrderStatusOk() (*string, bool)`

GetListOrderStatusOk returns a tuple with the ListOrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOrderStatus

`func (o *QueryMarginAccountsOpenOcoResponseInner) SetListOrderStatus(v string)`

SetListOrderStatus sets ListOrderStatus field to given value.

### HasListOrderStatus

`func (o *QueryMarginAccountsOpenOcoResponseInner) HasListOrderStatus() bool`

HasListOrderStatus returns a boolean if a field has been set.

### GetListClientOrderId

`func (o *QueryMarginAccountsOpenOcoResponseInner) GetListClientOrderId() string`

GetListClientOrderId returns the ListClientOrderId field if non-nil, zero value otherwise.

### GetListClientOrderIdOk

`func (o *QueryMarginAccountsOpenOcoResponseInner) GetListClientOrderIdOk() (*string, bool)`

GetListClientOrderIdOk returns a tuple with the ListClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListClientOrderId

`func (o *QueryMarginAccountsOpenOcoResponseInner) SetListClientOrderId(v string)`

SetListClientOrderId sets ListClientOrderId field to given value.

### HasListClientOrderId

`func (o *QueryMarginAccountsOpenOcoResponseInner) HasListClientOrderId() bool`

HasListClientOrderId returns a boolean if a field has been set.

### GetTransactionTime

`func (o *QueryMarginAccountsOpenOcoResponseInner) GetTransactionTime() int64`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *QueryMarginAccountsOpenOcoResponseInner) GetTransactionTimeOk() (*int64, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *QueryMarginAccountsOpenOcoResponseInner) SetTransactionTime(v int64)`

SetTransactionTime sets TransactionTime field to given value.

### HasTransactionTime

`func (o *QueryMarginAccountsOpenOcoResponseInner) HasTransactionTime() bool`

HasTransactionTime returns a boolean if a field has been set.

### GetSymbol

`func (o *QueryMarginAccountsOpenOcoResponseInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *QueryMarginAccountsOpenOcoResponseInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *QueryMarginAccountsOpenOcoResponseInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *QueryMarginAccountsOpenOcoResponseInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetIsIsolated

`func (o *QueryMarginAccountsOpenOcoResponseInner) GetIsIsolated() bool`

GetIsIsolated returns the IsIsolated field if non-nil, zero value otherwise.

### GetIsIsolatedOk

`func (o *QueryMarginAccountsOpenOcoResponseInner) GetIsIsolatedOk() (*bool, bool)`

GetIsIsolatedOk returns a tuple with the IsIsolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIsolated

`func (o *QueryMarginAccountsOpenOcoResponseInner) SetIsIsolated(v bool)`

SetIsIsolated sets IsIsolated field to given value.

### HasIsIsolated

`func (o *QueryMarginAccountsOpenOcoResponseInner) HasIsIsolated() bool`

HasIsIsolated returns a boolean if a field has been set.

### GetOrders

`func (o *QueryMarginAccountsOpenOcoResponseInner) GetOrders() []QueryMarginAccountsOpenOcoResponseInnerOrdersInner`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *QueryMarginAccountsOpenOcoResponseInner) GetOrdersOk() (*[]QueryMarginAccountsOpenOcoResponseInnerOrdersInner, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *QueryMarginAccountsOpenOcoResponseInner) SetOrders(v []QueryMarginAccountsOpenOcoResponseInnerOrdersInner)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *QueryMarginAccountsOpenOcoResponseInner) HasOrders() bool`

HasOrders returns a boolean if a field has been set.


[[Back to README]](../README.md)


