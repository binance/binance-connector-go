# MarginAccountNewOrderResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**ClientOrderId** | Pointer to **string** |  | [optional] 
**IsIsolated** | Pointer to **bool** |  | [optional] 
**TransactTime** | Pointer to **int64** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**OrigQty** | Pointer to **string** |  | [optional] 
**ExecutedQty** | Pointer to **string** |  | [optional] 
**CummulativeQuoteQty** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TimeInForce** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**SelfTradePreventionMode** | Pointer to **string** |  | [optional] 
**MarginBuyBorrowAmount** | Pointer to **int64** |  | [optional] 
**MarginBuyBorrowAsset** | Pointer to **string** |  | [optional] 
**Fills** | Pointer to [**[]MarginAccountNewOrderResponseFillsInner**](MarginAccountNewOrderResponseFillsInner.md) |  | [optional] 

## Methods

### NewMarginAccountNewOrderResponse

`func NewMarginAccountNewOrderResponse() *MarginAccountNewOrderResponse`

NewMarginAccountNewOrderResponse instantiates a new MarginAccountNewOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarginAccountNewOrderResponseWithDefaults

`func NewMarginAccountNewOrderResponseWithDefaults() *MarginAccountNewOrderResponse`

NewMarginAccountNewOrderResponseWithDefaults instantiates a new MarginAccountNewOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *MarginAccountNewOrderResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *MarginAccountNewOrderResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *MarginAccountNewOrderResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *MarginAccountNewOrderResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetOrderId

`func (o *MarginAccountNewOrderResponse) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *MarginAccountNewOrderResponse) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *MarginAccountNewOrderResponse) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *MarginAccountNewOrderResponse) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetClientOrderId

`func (o *MarginAccountNewOrderResponse) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *MarginAccountNewOrderResponse) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *MarginAccountNewOrderResponse) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *MarginAccountNewOrderResponse) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetIsIsolated

`func (o *MarginAccountNewOrderResponse) GetIsIsolated() bool`

GetIsIsolated returns the IsIsolated field if non-nil, zero value otherwise.

### GetIsIsolatedOk

`func (o *MarginAccountNewOrderResponse) GetIsIsolatedOk() (*bool, bool)`

GetIsIsolatedOk returns a tuple with the IsIsolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIsolated

`func (o *MarginAccountNewOrderResponse) SetIsIsolated(v bool)`

SetIsIsolated sets IsIsolated field to given value.

### HasIsIsolated

`func (o *MarginAccountNewOrderResponse) HasIsIsolated() bool`

HasIsIsolated returns a boolean if a field has been set.

### GetTransactTime

`func (o *MarginAccountNewOrderResponse) GetTransactTime() int64`

GetTransactTime returns the TransactTime field if non-nil, zero value otherwise.

### GetTransactTimeOk

`func (o *MarginAccountNewOrderResponse) GetTransactTimeOk() (*int64, bool)`

GetTransactTimeOk returns a tuple with the TransactTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactTime

`func (o *MarginAccountNewOrderResponse) SetTransactTime(v int64)`

SetTransactTime sets TransactTime field to given value.

### HasTransactTime

`func (o *MarginAccountNewOrderResponse) HasTransactTime() bool`

HasTransactTime returns a boolean if a field has been set.

### GetPrice

`func (o *MarginAccountNewOrderResponse) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *MarginAccountNewOrderResponse) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *MarginAccountNewOrderResponse) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *MarginAccountNewOrderResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetOrigQty

`func (o *MarginAccountNewOrderResponse) GetOrigQty() string`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *MarginAccountNewOrderResponse) GetOrigQtyOk() (*string, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *MarginAccountNewOrderResponse) SetOrigQty(v string)`

SetOrigQty sets OrigQty field to given value.

### HasOrigQty

`func (o *MarginAccountNewOrderResponse) HasOrigQty() bool`

HasOrigQty returns a boolean if a field has been set.

### GetExecutedQty

`func (o *MarginAccountNewOrderResponse) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *MarginAccountNewOrderResponse) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *MarginAccountNewOrderResponse) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *MarginAccountNewOrderResponse) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetCummulativeQuoteQty

`func (o *MarginAccountNewOrderResponse) GetCummulativeQuoteQty() string`

GetCummulativeQuoteQty returns the CummulativeQuoteQty field if non-nil, zero value otherwise.

### GetCummulativeQuoteQtyOk

`func (o *MarginAccountNewOrderResponse) GetCummulativeQuoteQtyOk() (*string, bool)`

GetCummulativeQuoteQtyOk returns a tuple with the CummulativeQuoteQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCummulativeQuoteQty

`func (o *MarginAccountNewOrderResponse) SetCummulativeQuoteQty(v string)`

SetCummulativeQuoteQty sets CummulativeQuoteQty field to given value.

### HasCummulativeQuoteQty

`func (o *MarginAccountNewOrderResponse) HasCummulativeQuoteQty() bool`

HasCummulativeQuoteQty returns a boolean if a field has been set.

### GetStatus

`func (o *MarginAccountNewOrderResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MarginAccountNewOrderResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MarginAccountNewOrderResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MarginAccountNewOrderResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeInForce

`func (o *MarginAccountNewOrderResponse) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *MarginAccountNewOrderResponse) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *MarginAccountNewOrderResponse) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *MarginAccountNewOrderResponse) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetType

`func (o *MarginAccountNewOrderResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MarginAccountNewOrderResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MarginAccountNewOrderResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MarginAccountNewOrderResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSide

`func (o *MarginAccountNewOrderResponse) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *MarginAccountNewOrderResponse) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *MarginAccountNewOrderResponse) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *MarginAccountNewOrderResponse) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetSelfTradePreventionMode

`func (o *MarginAccountNewOrderResponse) GetSelfTradePreventionMode() string`

GetSelfTradePreventionMode returns the SelfTradePreventionMode field if non-nil, zero value otherwise.

### GetSelfTradePreventionModeOk

`func (o *MarginAccountNewOrderResponse) GetSelfTradePreventionModeOk() (*string, bool)`

GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfTradePreventionMode

`func (o *MarginAccountNewOrderResponse) SetSelfTradePreventionMode(v string)`

SetSelfTradePreventionMode sets SelfTradePreventionMode field to given value.

### HasSelfTradePreventionMode

`func (o *MarginAccountNewOrderResponse) HasSelfTradePreventionMode() bool`

HasSelfTradePreventionMode returns a boolean if a field has been set.

### GetMarginBuyBorrowAmount

`func (o *MarginAccountNewOrderResponse) GetMarginBuyBorrowAmount() int64`

GetMarginBuyBorrowAmount returns the MarginBuyBorrowAmount field if non-nil, zero value otherwise.

### GetMarginBuyBorrowAmountOk

`func (o *MarginAccountNewOrderResponse) GetMarginBuyBorrowAmountOk() (*int64, bool)`

GetMarginBuyBorrowAmountOk returns a tuple with the MarginBuyBorrowAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginBuyBorrowAmount

`func (o *MarginAccountNewOrderResponse) SetMarginBuyBorrowAmount(v int64)`

SetMarginBuyBorrowAmount sets MarginBuyBorrowAmount field to given value.

### HasMarginBuyBorrowAmount

`func (o *MarginAccountNewOrderResponse) HasMarginBuyBorrowAmount() bool`

HasMarginBuyBorrowAmount returns a boolean if a field has been set.

### GetMarginBuyBorrowAsset

`func (o *MarginAccountNewOrderResponse) GetMarginBuyBorrowAsset() string`

GetMarginBuyBorrowAsset returns the MarginBuyBorrowAsset field if non-nil, zero value otherwise.

### GetMarginBuyBorrowAssetOk

`func (o *MarginAccountNewOrderResponse) GetMarginBuyBorrowAssetOk() (*string, bool)`

GetMarginBuyBorrowAssetOk returns a tuple with the MarginBuyBorrowAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginBuyBorrowAsset

`func (o *MarginAccountNewOrderResponse) SetMarginBuyBorrowAsset(v string)`

SetMarginBuyBorrowAsset sets MarginBuyBorrowAsset field to given value.

### HasMarginBuyBorrowAsset

`func (o *MarginAccountNewOrderResponse) HasMarginBuyBorrowAsset() bool`

HasMarginBuyBorrowAsset returns a boolean if a field has been set.

### GetFills

`func (o *MarginAccountNewOrderResponse) GetFills() []MarginAccountNewOrderResponseFillsInner`

GetFills returns the Fills field if non-nil, zero value otherwise.

### GetFillsOk

`func (o *MarginAccountNewOrderResponse) GetFillsOk() (*[]MarginAccountNewOrderResponseFillsInner, bool)`

GetFillsOk returns a tuple with the Fills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFills

`func (o *MarginAccountNewOrderResponse) SetFills(v []MarginAccountNewOrderResponseFillsInner)`

SetFills sets Fills field to given value.

### HasFills

`func (o *MarginAccountNewOrderResponse) HasFills() bool`

HasFills returns a boolean if a field has been set.


[[Back to README]](../README.md)


