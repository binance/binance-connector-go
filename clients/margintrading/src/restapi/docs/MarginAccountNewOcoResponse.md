# MarginAccountNewOcoResponse

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
**MarginBuyBorrowAmount** | Pointer to **string** |  | [optional] 
**MarginBuyBorrowAsset** | Pointer to **string** |  | [optional] 
**IsIsolated** | Pointer to **bool** |  | [optional] 
**Orders** | Pointer to [**[]MarginAccountNewOcoResponseOrdersInner**](MarginAccountNewOcoResponseOrdersInner.md) |  | [optional] 
**OrderReports** | Pointer to [**[]MarginAccountNewOcoResponseOrderReportsInner**](MarginAccountNewOcoResponseOrderReportsInner.md) |  | [optional] 

## Methods

### NewMarginAccountNewOcoResponse

`func NewMarginAccountNewOcoResponse() *MarginAccountNewOcoResponse`

NewMarginAccountNewOcoResponse instantiates a new MarginAccountNewOcoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarginAccountNewOcoResponseWithDefaults

`func NewMarginAccountNewOcoResponseWithDefaults() *MarginAccountNewOcoResponse`

NewMarginAccountNewOcoResponseWithDefaults instantiates a new MarginAccountNewOcoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderListId

`func (o *MarginAccountNewOcoResponse) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *MarginAccountNewOcoResponse) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *MarginAccountNewOcoResponse) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *MarginAccountNewOcoResponse) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetContingencyType

`func (o *MarginAccountNewOcoResponse) GetContingencyType() string`

GetContingencyType returns the ContingencyType field if non-nil, zero value otherwise.

### GetContingencyTypeOk

`func (o *MarginAccountNewOcoResponse) GetContingencyTypeOk() (*string, bool)`

GetContingencyTypeOk returns a tuple with the ContingencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContingencyType

`func (o *MarginAccountNewOcoResponse) SetContingencyType(v string)`

SetContingencyType sets ContingencyType field to given value.

### HasContingencyType

`func (o *MarginAccountNewOcoResponse) HasContingencyType() bool`

HasContingencyType returns a boolean if a field has been set.

### GetListStatusType

`func (o *MarginAccountNewOcoResponse) GetListStatusType() string`

GetListStatusType returns the ListStatusType field if non-nil, zero value otherwise.

### GetListStatusTypeOk

`func (o *MarginAccountNewOcoResponse) GetListStatusTypeOk() (*string, bool)`

GetListStatusTypeOk returns a tuple with the ListStatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListStatusType

`func (o *MarginAccountNewOcoResponse) SetListStatusType(v string)`

SetListStatusType sets ListStatusType field to given value.

### HasListStatusType

`func (o *MarginAccountNewOcoResponse) HasListStatusType() bool`

HasListStatusType returns a boolean if a field has been set.

### GetListOrderStatus

`func (o *MarginAccountNewOcoResponse) GetListOrderStatus() string`

GetListOrderStatus returns the ListOrderStatus field if non-nil, zero value otherwise.

### GetListOrderStatusOk

`func (o *MarginAccountNewOcoResponse) GetListOrderStatusOk() (*string, bool)`

GetListOrderStatusOk returns a tuple with the ListOrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOrderStatus

`func (o *MarginAccountNewOcoResponse) SetListOrderStatus(v string)`

SetListOrderStatus sets ListOrderStatus field to given value.

### HasListOrderStatus

`func (o *MarginAccountNewOcoResponse) HasListOrderStatus() bool`

HasListOrderStatus returns a boolean if a field has been set.

### GetListClientOrderId

`func (o *MarginAccountNewOcoResponse) GetListClientOrderId() string`

GetListClientOrderId returns the ListClientOrderId field if non-nil, zero value otherwise.

### GetListClientOrderIdOk

`func (o *MarginAccountNewOcoResponse) GetListClientOrderIdOk() (*string, bool)`

GetListClientOrderIdOk returns a tuple with the ListClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListClientOrderId

`func (o *MarginAccountNewOcoResponse) SetListClientOrderId(v string)`

SetListClientOrderId sets ListClientOrderId field to given value.

### HasListClientOrderId

`func (o *MarginAccountNewOcoResponse) HasListClientOrderId() bool`

HasListClientOrderId returns a boolean if a field has been set.

### GetTransactionTime

`func (o *MarginAccountNewOcoResponse) GetTransactionTime() int64`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *MarginAccountNewOcoResponse) GetTransactionTimeOk() (*int64, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *MarginAccountNewOcoResponse) SetTransactionTime(v int64)`

SetTransactionTime sets TransactionTime field to given value.

### HasTransactionTime

`func (o *MarginAccountNewOcoResponse) HasTransactionTime() bool`

HasTransactionTime returns a boolean if a field has been set.

### GetSymbol

`func (o *MarginAccountNewOcoResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *MarginAccountNewOcoResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *MarginAccountNewOcoResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *MarginAccountNewOcoResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetMarginBuyBorrowAmount

`func (o *MarginAccountNewOcoResponse) GetMarginBuyBorrowAmount() string`

GetMarginBuyBorrowAmount returns the MarginBuyBorrowAmount field if non-nil, zero value otherwise.

### GetMarginBuyBorrowAmountOk

`func (o *MarginAccountNewOcoResponse) GetMarginBuyBorrowAmountOk() (*string, bool)`

GetMarginBuyBorrowAmountOk returns a tuple with the MarginBuyBorrowAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginBuyBorrowAmount

`func (o *MarginAccountNewOcoResponse) SetMarginBuyBorrowAmount(v string)`

SetMarginBuyBorrowAmount sets MarginBuyBorrowAmount field to given value.

### HasMarginBuyBorrowAmount

`func (o *MarginAccountNewOcoResponse) HasMarginBuyBorrowAmount() bool`

HasMarginBuyBorrowAmount returns a boolean if a field has been set.

### GetMarginBuyBorrowAsset

`func (o *MarginAccountNewOcoResponse) GetMarginBuyBorrowAsset() string`

GetMarginBuyBorrowAsset returns the MarginBuyBorrowAsset field if non-nil, zero value otherwise.

### GetMarginBuyBorrowAssetOk

`func (o *MarginAccountNewOcoResponse) GetMarginBuyBorrowAssetOk() (*string, bool)`

GetMarginBuyBorrowAssetOk returns a tuple with the MarginBuyBorrowAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginBuyBorrowAsset

`func (o *MarginAccountNewOcoResponse) SetMarginBuyBorrowAsset(v string)`

SetMarginBuyBorrowAsset sets MarginBuyBorrowAsset field to given value.

### HasMarginBuyBorrowAsset

`func (o *MarginAccountNewOcoResponse) HasMarginBuyBorrowAsset() bool`

HasMarginBuyBorrowAsset returns a boolean if a field has been set.

### GetIsIsolated

`func (o *MarginAccountNewOcoResponse) GetIsIsolated() bool`

GetIsIsolated returns the IsIsolated field if non-nil, zero value otherwise.

### GetIsIsolatedOk

`func (o *MarginAccountNewOcoResponse) GetIsIsolatedOk() (*bool, bool)`

GetIsIsolatedOk returns a tuple with the IsIsolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIsolated

`func (o *MarginAccountNewOcoResponse) SetIsIsolated(v bool)`

SetIsIsolated sets IsIsolated field to given value.

### HasIsIsolated

`func (o *MarginAccountNewOcoResponse) HasIsIsolated() bool`

HasIsIsolated returns a boolean if a field has been set.

### GetOrders

`func (o *MarginAccountNewOcoResponse) GetOrders() []MarginAccountNewOcoResponseOrdersInner`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *MarginAccountNewOcoResponse) GetOrdersOk() (*[]MarginAccountNewOcoResponseOrdersInner, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *MarginAccountNewOcoResponse) SetOrders(v []MarginAccountNewOcoResponseOrdersInner)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *MarginAccountNewOcoResponse) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetOrderReports

`func (o *MarginAccountNewOcoResponse) GetOrderReports() []MarginAccountNewOcoResponseOrderReportsInner`

GetOrderReports returns the OrderReports field if non-nil, zero value otherwise.

### GetOrderReportsOk

`func (o *MarginAccountNewOcoResponse) GetOrderReportsOk() (*[]MarginAccountNewOcoResponseOrderReportsInner, bool)`

GetOrderReportsOk returns a tuple with the OrderReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderReports

`func (o *MarginAccountNewOcoResponse) SetOrderReports(v []MarginAccountNewOcoResponseOrderReportsInner)`

SetOrderReports sets OrderReports field to given value.

### HasOrderReports

`func (o *MarginAccountNewOcoResponse) HasOrderReports() bool`

HasOrderReports returns a boolean if a field has been set.


[[Back to README]](../README.md)


