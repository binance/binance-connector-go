# QueryMarginAccountsAllOcoResponseInner

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
**Orders** | Pointer to [**[]QueryMarginAccountsAllOcoResponseInnerOrdersInner**](QueryMarginAccountsAllOcoResponseInnerOrdersInner.md) |  | [optional] 

## Methods

### NewQueryMarginAccountsAllOcoResponseInner

`func NewQueryMarginAccountsAllOcoResponseInner() *QueryMarginAccountsAllOcoResponseInner`

NewQueryMarginAccountsAllOcoResponseInner instantiates a new QueryMarginAccountsAllOcoResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryMarginAccountsAllOcoResponseInnerWithDefaults

`func NewQueryMarginAccountsAllOcoResponseInnerWithDefaults() *QueryMarginAccountsAllOcoResponseInner`

NewQueryMarginAccountsAllOcoResponseInnerWithDefaults instantiates a new QueryMarginAccountsAllOcoResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderListId

`func (o *QueryMarginAccountsAllOcoResponseInner) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *QueryMarginAccountsAllOcoResponseInner) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *QueryMarginAccountsAllOcoResponseInner) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *QueryMarginAccountsAllOcoResponseInner) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetContingencyType

`func (o *QueryMarginAccountsAllOcoResponseInner) GetContingencyType() string`

GetContingencyType returns the ContingencyType field if non-nil, zero value otherwise.

### GetContingencyTypeOk

`func (o *QueryMarginAccountsAllOcoResponseInner) GetContingencyTypeOk() (*string, bool)`

GetContingencyTypeOk returns a tuple with the ContingencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContingencyType

`func (o *QueryMarginAccountsAllOcoResponseInner) SetContingencyType(v string)`

SetContingencyType sets ContingencyType field to given value.

### HasContingencyType

`func (o *QueryMarginAccountsAllOcoResponseInner) HasContingencyType() bool`

HasContingencyType returns a boolean if a field has been set.

### GetListStatusType

`func (o *QueryMarginAccountsAllOcoResponseInner) GetListStatusType() string`

GetListStatusType returns the ListStatusType field if non-nil, zero value otherwise.

### GetListStatusTypeOk

`func (o *QueryMarginAccountsAllOcoResponseInner) GetListStatusTypeOk() (*string, bool)`

GetListStatusTypeOk returns a tuple with the ListStatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListStatusType

`func (o *QueryMarginAccountsAllOcoResponseInner) SetListStatusType(v string)`

SetListStatusType sets ListStatusType field to given value.

### HasListStatusType

`func (o *QueryMarginAccountsAllOcoResponseInner) HasListStatusType() bool`

HasListStatusType returns a boolean if a field has been set.

### GetListOrderStatus

`func (o *QueryMarginAccountsAllOcoResponseInner) GetListOrderStatus() string`

GetListOrderStatus returns the ListOrderStatus field if non-nil, zero value otherwise.

### GetListOrderStatusOk

`func (o *QueryMarginAccountsAllOcoResponseInner) GetListOrderStatusOk() (*string, bool)`

GetListOrderStatusOk returns a tuple with the ListOrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOrderStatus

`func (o *QueryMarginAccountsAllOcoResponseInner) SetListOrderStatus(v string)`

SetListOrderStatus sets ListOrderStatus field to given value.

### HasListOrderStatus

`func (o *QueryMarginAccountsAllOcoResponseInner) HasListOrderStatus() bool`

HasListOrderStatus returns a boolean if a field has been set.

### GetListClientOrderId

`func (o *QueryMarginAccountsAllOcoResponseInner) GetListClientOrderId() string`

GetListClientOrderId returns the ListClientOrderId field if non-nil, zero value otherwise.

### GetListClientOrderIdOk

`func (o *QueryMarginAccountsAllOcoResponseInner) GetListClientOrderIdOk() (*string, bool)`

GetListClientOrderIdOk returns a tuple with the ListClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListClientOrderId

`func (o *QueryMarginAccountsAllOcoResponseInner) SetListClientOrderId(v string)`

SetListClientOrderId sets ListClientOrderId field to given value.

### HasListClientOrderId

`func (o *QueryMarginAccountsAllOcoResponseInner) HasListClientOrderId() bool`

HasListClientOrderId returns a boolean if a field has been set.

### GetTransactionTime

`func (o *QueryMarginAccountsAllOcoResponseInner) GetTransactionTime() int64`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *QueryMarginAccountsAllOcoResponseInner) GetTransactionTimeOk() (*int64, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *QueryMarginAccountsAllOcoResponseInner) SetTransactionTime(v int64)`

SetTransactionTime sets TransactionTime field to given value.

### HasTransactionTime

`func (o *QueryMarginAccountsAllOcoResponseInner) HasTransactionTime() bool`

HasTransactionTime returns a boolean if a field has been set.

### GetSymbol

`func (o *QueryMarginAccountsAllOcoResponseInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *QueryMarginAccountsAllOcoResponseInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *QueryMarginAccountsAllOcoResponseInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *QueryMarginAccountsAllOcoResponseInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetOrders

`func (o *QueryMarginAccountsAllOcoResponseInner) GetOrders() []QueryMarginAccountsAllOcoResponseInnerOrdersInner`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *QueryMarginAccountsAllOcoResponseInner) GetOrdersOk() (*[]QueryMarginAccountsAllOcoResponseInnerOrdersInner, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *QueryMarginAccountsAllOcoResponseInner) SetOrders(v []QueryMarginAccountsAllOcoResponseInnerOrdersInner)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *QueryMarginAccountsAllOcoResponseInner) HasOrders() bool`

HasOrders returns a boolean if a field has been set.


[[Back to README]](../README.md)


