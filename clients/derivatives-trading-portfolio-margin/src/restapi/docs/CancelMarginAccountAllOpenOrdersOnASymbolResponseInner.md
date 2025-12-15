# CancelMarginAccountAllOpenOrdersOnASymbolResponseInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**OrigClientOrderId** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**OrderListId** | Pointer to **int64** |  | [optional] 
**ClientOrderId** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**OrigQty** | Pointer to **string** |  | [optional] 
**ExecutedQty** | Pointer to **string** |  | [optional] 
**CummulativeQuoteQty** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TimeInForce** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**ContingencyType** | Pointer to **string** |  | [optional] 
**ListStatusType** | Pointer to **string** |  | [optional] 
**ListOrderStatus** | Pointer to **string** |  | [optional] 
**ListClientOrderId** | Pointer to **string** |  | [optional] 
**TransactionTime** | Pointer to **int64** |  | [optional] 
**Orders** | Pointer to [**[]CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner**](CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner.md) |  | [optional] 
**OrderReports** | Pointer to [**[]CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrderReportsInner**](CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrderReportsInner.md) |  | [optional] 

## Methods

### NewCancelMarginAccountAllOpenOrdersOnASymbolResponseInner

`func NewCancelMarginAccountAllOpenOrdersOnASymbolResponseInner() *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner`

NewCancelMarginAccountAllOpenOrdersOnASymbolResponseInner instantiates a new CancelMarginAccountAllOpenOrdersOnASymbolResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelMarginAccountAllOpenOrdersOnASymbolResponseInnerWithDefaults

`func NewCancelMarginAccountAllOpenOrdersOnASymbolResponseInnerWithDefaults() *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner`

NewCancelMarginAccountAllOpenOrdersOnASymbolResponseInnerWithDefaults instantiates a new CancelMarginAccountAllOpenOrdersOnASymbolResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetOrigClientOrderId

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetOrigClientOrderId() string`

GetOrigClientOrderId returns the OrigClientOrderId field if non-nil, zero value otherwise.

### GetOrigClientOrderIdOk

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetOrigClientOrderIdOk() (*string, bool)`

GetOrigClientOrderIdOk returns a tuple with the OrigClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigClientOrderId

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) SetOrigClientOrderId(v string)`

SetOrigClientOrderId sets OrigClientOrderId field to given value.

### HasOrigClientOrderId

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) HasOrigClientOrderId() bool`

HasOrigClientOrderId returns a boolean if a field has been set.

### GetOrderId

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderListId

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetClientOrderId

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetPrice

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetOrigQty

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetOrigQty() string`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetOrigQtyOk() (*string, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) SetOrigQty(v string)`

SetOrigQty sets OrigQty field to given value.

### HasOrigQty

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) HasOrigQty() bool`

HasOrigQty returns a boolean if a field has been set.

### GetExecutedQty

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetCummulativeQuoteQty

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetCummulativeQuoteQty() string`

GetCummulativeQuoteQty returns the CummulativeQuoteQty field if non-nil, zero value otherwise.

### GetCummulativeQuoteQtyOk

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetCummulativeQuoteQtyOk() (*string, bool)`

GetCummulativeQuoteQtyOk returns a tuple with the CummulativeQuoteQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCummulativeQuoteQty

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) SetCummulativeQuoteQty(v string)`

SetCummulativeQuoteQty sets CummulativeQuoteQty field to given value.

### HasCummulativeQuoteQty

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) HasCummulativeQuoteQty() bool`

HasCummulativeQuoteQty returns a boolean if a field has been set.

### GetStatus

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeInForce

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetType

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSide

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetContingencyType

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetContingencyType() string`

GetContingencyType returns the ContingencyType field if non-nil, zero value otherwise.

### GetContingencyTypeOk

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetContingencyTypeOk() (*string, bool)`

GetContingencyTypeOk returns a tuple with the ContingencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContingencyType

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) SetContingencyType(v string)`

SetContingencyType sets ContingencyType field to given value.

### HasContingencyType

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) HasContingencyType() bool`

HasContingencyType returns a boolean if a field has been set.

### GetListStatusType

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetListStatusType() string`

GetListStatusType returns the ListStatusType field if non-nil, zero value otherwise.

### GetListStatusTypeOk

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetListStatusTypeOk() (*string, bool)`

GetListStatusTypeOk returns a tuple with the ListStatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListStatusType

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) SetListStatusType(v string)`

SetListStatusType sets ListStatusType field to given value.

### HasListStatusType

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) HasListStatusType() bool`

HasListStatusType returns a boolean if a field has been set.

### GetListOrderStatus

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetListOrderStatus() string`

GetListOrderStatus returns the ListOrderStatus field if non-nil, zero value otherwise.

### GetListOrderStatusOk

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetListOrderStatusOk() (*string, bool)`

GetListOrderStatusOk returns a tuple with the ListOrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOrderStatus

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) SetListOrderStatus(v string)`

SetListOrderStatus sets ListOrderStatus field to given value.

### HasListOrderStatus

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) HasListOrderStatus() bool`

HasListOrderStatus returns a boolean if a field has been set.

### GetListClientOrderId

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetListClientOrderId() string`

GetListClientOrderId returns the ListClientOrderId field if non-nil, zero value otherwise.

### GetListClientOrderIdOk

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetListClientOrderIdOk() (*string, bool)`

GetListClientOrderIdOk returns a tuple with the ListClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListClientOrderId

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) SetListClientOrderId(v string)`

SetListClientOrderId sets ListClientOrderId field to given value.

### HasListClientOrderId

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) HasListClientOrderId() bool`

HasListClientOrderId returns a boolean if a field has been set.

### GetTransactionTime

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetTransactionTime() int64`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetTransactionTimeOk() (*int64, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) SetTransactionTime(v int64)`

SetTransactionTime sets TransactionTime field to given value.

### HasTransactionTime

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) HasTransactionTime() bool`

HasTransactionTime returns a boolean if a field has been set.

### GetOrders

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetOrders() []CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetOrdersOk() (*[]CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) SetOrders(v []CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetOrderReports

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetOrderReports() []CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrderReportsInner`

GetOrderReports returns the OrderReports field if non-nil, zero value otherwise.

### GetOrderReportsOk

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) GetOrderReportsOk() (*[]CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrderReportsInner, bool)`

GetOrderReportsOk returns a tuple with the OrderReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderReports

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) SetOrderReports(v []CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrderReportsInner)`

SetOrderReports sets OrderReports field to given value.

### HasOrderReports

`func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInner) HasOrderReports() bool`

HasOrderReports returns a boolean if a field has been set.


[[Back to README]](../README.md)


