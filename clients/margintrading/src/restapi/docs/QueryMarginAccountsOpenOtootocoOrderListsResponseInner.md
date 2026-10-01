# QueryMarginAccountsOpenOtootocoOrderListsResponseInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**OrderListId** | Pointer to **int64** | order List Id. | [optional] 
**ContingencyType** | Pointer to **string** | contingency Type. | [optional] 
**ListStatusType** | Pointer to **string** | list Status Type. | [optional] 
**ListOrderStatus** | Pointer to **string** | list Order Status. | [optional] 
**ListClientOrderId** | Pointer to **string** | list Client Order Id. | [optional] 
**TransactionTime** | Pointer to **int64** | transaction Time. | [optional] 
**Symbol** | Pointer to **string** | symbol. | [optional] 
**Orders** | Pointer to [**[]QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner**](QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner.md) | orders list. | [optional] 

## Methods

### NewQueryMarginAccountsOpenOtootocoOrderListsResponseInner

`func NewQueryMarginAccountsOpenOtootocoOrderListsResponseInner() *QueryMarginAccountsOpenOtootocoOrderListsResponseInner`

NewQueryMarginAccountsOpenOtootocoOrderListsResponseInner instantiates a new QueryMarginAccountsOpenOtootocoOrderListsResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryMarginAccountsOpenOtootocoOrderListsResponseInnerWithDefaults

`func NewQueryMarginAccountsOpenOtootocoOrderListsResponseInnerWithDefaults() *QueryMarginAccountsOpenOtootocoOrderListsResponseInner`

NewQueryMarginAccountsOpenOtootocoOrderListsResponseInnerWithDefaults instantiates a new QueryMarginAccountsOpenOtootocoOrderListsResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderListId

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetContingencyType

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) GetContingencyType() string`

GetContingencyType returns the ContingencyType field if non-nil, zero value otherwise.

### GetContingencyTypeOk

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) GetContingencyTypeOk() (*string, bool)`

GetContingencyTypeOk returns a tuple with the ContingencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContingencyType

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) SetContingencyType(v string)`

SetContingencyType sets ContingencyType field to given value.

### HasContingencyType

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) HasContingencyType() bool`

HasContingencyType returns a boolean if a field has been set.

### GetListStatusType

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) GetListStatusType() string`

GetListStatusType returns the ListStatusType field if non-nil, zero value otherwise.

### GetListStatusTypeOk

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) GetListStatusTypeOk() (*string, bool)`

GetListStatusTypeOk returns a tuple with the ListStatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListStatusType

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) SetListStatusType(v string)`

SetListStatusType sets ListStatusType field to given value.

### HasListStatusType

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) HasListStatusType() bool`

HasListStatusType returns a boolean if a field has been set.

### GetListOrderStatus

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) GetListOrderStatus() string`

GetListOrderStatus returns the ListOrderStatus field if non-nil, zero value otherwise.

### GetListOrderStatusOk

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) GetListOrderStatusOk() (*string, bool)`

GetListOrderStatusOk returns a tuple with the ListOrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOrderStatus

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) SetListOrderStatus(v string)`

SetListOrderStatus sets ListOrderStatus field to given value.

### HasListOrderStatus

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) HasListOrderStatus() bool`

HasListOrderStatus returns a boolean if a field has been set.

### GetListClientOrderId

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) GetListClientOrderId() string`

GetListClientOrderId returns the ListClientOrderId field if non-nil, zero value otherwise.

### GetListClientOrderIdOk

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) GetListClientOrderIdOk() (*string, bool)`

GetListClientOrderIdOk returns a tuple with the ListClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListClientOrderId

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) SetListClientOrderId(v string)`

SetListClientOrderId sets ListClientOrderId field to given value.

### HasListClientOrderId

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) HasListClientOrderId() bool`

HasListClientOrderId returns a boolean if a field has been set.

### GetTransactionTime

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) GetTransactionTime() int64`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) GetTransactionTimeOk() (*int64, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) SetTransactionTime(v int64)`

SetTransactionTime sets TransactionTime field to given value.

### HasTransactionTime

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) HasTransactionTime() bool`

HasTransactionTime returns a boolean if a field has been set.

### GetSymbol

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetOrders

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) GetOrders() []QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) GetOrdersOk() (*[]QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) SetOrders(v []QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInner) HasOrders() bool`

HasOrders returns a boolean if a field has been set.


[[Back to README]](../README.md)


