# MarginAccountCancelAllOpenOrdersOnASymbolResponseInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**IsIsolated** | Pointer to **bool** |  | [optional] 
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
**SelfTradePreventionMode** | Pointer to **string** |  | [optional] 
**ContingencyType** | Pointer to **string** |  | [optional] 
**ListStatusType** | Pointer to **string** |  | [optional] 
**ListOrderStatus** | Pointer to **string** |  | [optional] 
**ListClientOrderId** | Pointer to **string** |  | [optional] 
**TransactionTime** | Pointer to **int64** |  | [optional] 
**Orders** | Pointer to [**[]MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrdersInner**](MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrdersInner.md) |  | [optional] 
**OrderReports** | Pointer to [**[]MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner**](MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner.md) |  | [optional] 

## Methods

### NewMarginAccountCancelAllOpenOrdersOnASymbolResponseInner

`func NewMarginAccountCancelAllOpenOrdersOnASymbolResponseInner() *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner`

NewMarginAccountCancelAllOpenOrdersOnASymbolResponseInner instantiates a new MarginAccountCancelAllOpenOrdersOnASymbolResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarginAccountCancelAllOpenOrdersOnASymbolResponseInnerWithDefaults

`func NewMarginAccountCancelAllOpenOrdersOnASymbolResponseInnerWithDefaults() *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner`

NewMarginAccountCancelAllOpenOrdersOnASymbolResponseInnerWithDefaults instantiates a new MarginAccountCancelAllOpenOrdersOnASymbolResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetIsIsolated

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetIsIsolated() bool`

GetIsIsolated returns the IsIsolated field if non-nil, zero value otherwise.

### GetIsIsolatedOk

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetIsIsolatedOk() (*bool, bool)`

GetIsIsolatedOk returns a tuple with the IsIsolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIsolated

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetIsIsolated(v bool)`

SetIsIsolated sets IsIsolated field to given value.

### HasIsIsolated

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasIsIsolated() bool`

HasIsIsolated returns a boolean if a field has been set.

### GetOrigClientOrderId

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetOrigClientOrderId() string`

GetOrigClientOrderId returns the OrigClientOrderId field if non-nil, zero value otherwise.

### GetOrigClientOrderIdOk

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetOrigClientOrderIdOk() (*string, bool)`

GetOrigClientOrderIdOk returns a tuple with the OrigClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigClientOrderId

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetOrigClientOrderId(v string)`

SetOrigClientOrderId sets OrigClientOrderId field to given value.

### HasOrigClientOrderId

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasOrigClientOrderId() bool`

HasOrigClientOrderId returns a boolean if a field has been set.

### GetOrderId

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderListId

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetClientOrderId

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetPrice

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetOrigQty

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetOrigQty() string`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetOrigQtyOk() (*string, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetOrigQty(v string)`

SetOrigQty sets OrigQty field to given value.

### HasOrigQty

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasOrigQty() bool`

HasOrigQty returns a boolean if a field has been set.

### GetExecutedQty

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetCummulativeQuoteQty

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetCummulativeQuoteQty() string`

GetCummulativeQuoteQty returns the CummulativeQuoteQty field if non-nil, zero value otherwise.

### GetCummulativeQuoteQtyOk

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetCummulativeQuoteQtyOk() (*string, bool)`

GetCummulativeQuoteQtyOk returns a tuple with the CummulativeQuoteQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCummulativeQuoteQty

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetCummulativeQuoteQty(v string)`

SetCummulativeQuoteQty sets CummulativeQuoteQty field to given value.

### HasCummulativeQuoteQty

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasCummulativeQuoteQty() bool`

HasCummulativeQuoteQty returns a boolean if a field has been set.

### GetStatus

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeInForce

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetType

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSide

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetSelfTradePreventionMode

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetSelfTradePreventionMode() string`

GetSelfTradePreventionMode returns the SelfTradePreventionMode field if non-nil, zero value otherwise.

### GetSelfTradePreventionModeOk

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetSelfTradePreventionModeOk() (*string, bool)`

GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfTradePreventionMode

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetSelfTradePreventionMode(v string)`

SetSelfTradePreventionMode sets SelfTradePreventionMode field to given value.

### HasSelfTradePreventionMode

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasSelfTradePreventionMode() bool`

HasSelfTradePreventionMode returns a boolean if a field has been set.

### GetContingencyType

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetContingencyType() string`

GetContingencyType returns the ContingencyType field if non-nil, zero value otherwise.

### GetContingencyTypeOk

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetContingencyTypeOk() (*string, bool)`

GetContingencyTypeOk returns a tuple with the ContingencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContingencyType

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetContingencyType(v string)`

SetContingencyType sets ContingencyType field to given value.

### HasContingencyType

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasContingencyType() bool`

HasContingencyType returns a boolean if a field has been set.

### GetListStatusType

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetListStatusType() string`

GetListStatusType returns the ListStatusType field if non-nil, zero value otherwise.

### GetListStatusTypeOk

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetListStatusTypeOk() (*string, bool)`

GetListStatusTypeOk returns a tuple with the ListStatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListStatusType

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetListStatusType(v string)`

SetListStatusType sets ListStatusType field to given value.

### HasListStatusType

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasListStatusType() bool`

HasListStatusType returns a boolean if a field has been set.

### GetListOrderStatus

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetListOrderStatus() string`

GetListOrderStatus returns the ListOrderStatus field if non-nil, zero value otherwise.

### GetListOrderStatusOk

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetListOrderStatusOk() (*string, bool)`

GetListOrderStatusOk returns a tuple with the ListOrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOrderStatus

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetListOrderStatus(v string)`

SetListOrderStatus sets ListOrderStatus field to given value.

### HasListOrderStatus

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasListOrderStatus() bool`

HasListOrderStatus returns a boolean if a field has been set.

### GetListClientOrderId

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetListClientOrderId() string`

GetListClientOrderId returns the ListClientOrderId field if non-nil, zero value otherwise.

### GetListClientOrderIdOk

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetListClientOrderIdOk() (*string, bool)`

GetListClientOrderIdOk returns a tuple with the ListClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListClientOrderId

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetListClientOrderId(v string)`

SetListClientOrderId sets ListClientOrderId field to given value.

### HasListClientOrderId

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasListClientOrderId() bool`

HasListClientOrderId returns a boolean if a field has been set.

### GetTransactionTime

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetTransactionTime() int64`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetTransactionTimeOk() (*int64, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetTransactionTime(v int64)`

SetTransactionTime sets TransactionTime field to given value.

### HasTransactionTime

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasTransactionTime() bool`

HasTransactionTime returns a boolean if a field has been set.

### GetOrders

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetOrders() []MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrdersInner`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetOrdersOk() (*[]MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrdersInner, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetOrders(v []MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrdersInner)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetOrderReports

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetOrderReports() []MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner`

GetOrderReports returns the OrderReports field if non-nil, zero value otherwise.

### GetOrderReportsOk

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) GetOrderReportsOk() (*[]MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner, bool)`

GetOrderReportsOk returns a tuple with the OrderReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderReports

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) SetOrderReports(v []MarginAccountCancelAllOpenOrdersOnASymbolResponseInnerOrderReportsInner)`

SetOrderReports sets OrderReports field to given value.

### HasOrderReports

`func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponseInner) HasOrderReports() bool`

HasOrderReports returns a boolean if a field has been set.


[[Back to README]](../README.md)


