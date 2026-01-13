# QueryMarginAccountOrderResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**ClientOrderId** | Pointer to **string** |  | [optional] 
**CummulativeQuoteQty** | Pointer to **string** |  | [optional] 
**ExecutedQty** | Pointer to **string** |  | [optional] 
**IcebergQty** | Pointer to **string** |  | [optional] 
**IsWorking** | Pointer to **bool** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**OrigQty** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StopPrice** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 
**TimeInForce** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**SelfTradePreventionMode** | Pointer to **string** |  | [optional] 
**PreventedMatchId** | Pointer to **string** |  | [optional] 
**PreventedQuantity** | Pointer to **string** |  | [optional] 

## Methods

### NewQueryMarginAccountOrderResponse

`func NewQueryMarginAccountOrderResponse() *QueryMarginAccountOrderResponse`

NewQueryMarginAccountOrderResponse instantiates a new QueryMarginAccountOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryMarginAccountOrderResponseWithDefaults

`func NewQueryMarginAccountOrderResponseWithDefaults() *QueryMarginAccountOrderResponse`

NewQueryMarginAccountOrderResponseWithDefaults instantiates a new QueryMarginAccountOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientOrderId

`func (o *QueryMarginAccountOrderResponse) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *QueryMarginAccountOrderResponse) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *QueryMarginAccountOrderResponse) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *QueryMarginAccountOrderResponse) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetCummulativeQuoteQty

`func (o *QueryMarginAccountOrderResponse) GetCummulativeQuoteQty() string`

GetCummulativeQuoteQty returns the CummulativeQuoteQty field if non-nil, zero value otherwise.

### GetCummulativeQuoteQtyOk

`func (o *QueryMarginAccountOrderResponse) GetCummulativeQuoteQtyOk() (*string, bool)`

GetCummulativeQuoteQtyOk returns a tuple with the CummulativeQuoteQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCummulativeQuoteQty

`func (o *QueryMarginAccountOrderResponse) SetCummulativeQuoteQty(v string)`

SetCummulativeQuoteQty sets CummulativeQuoteQty field to given value.

### HasCummulativeQuoteQty

`func (o *QueryMarginAccountOrderResponse) HasCummulativeQuoteQty() bool`

HasCummulativeQuoteQty returns a boolean if a field has been set.

### GetExecutedQty

`func (o *QueryMarginAccountOrderResponse) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *QueryMarginAccountOrderResponse) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *QueryMarginAccountOrderResponse) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *QueryMarginAccountOrderResponse) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetIcebergQty

`func (o *QueryMarginAccountOrderResponse) GetIcebergQty() string`

GetIcebergQty returns the IcebergQty field if non-nil, zero value otherwise.

### GetIcebergQtyOk

`func (o *QueryMarginAccountOrderResponse) GetIcebergQtyOk() (*string, bool)`

GetIcebergQtyOk returns a tuple with the IcebergQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcebergQty

`func (o *QueryMarginAccountOrderResponse) SetIcebergQty(v string)`

SetIcebergQty sets IcebergQty field to given value.

### HasIcebergQty

`func (o *QueryMarginAccountOrderResponse) HasIcebergQty() bool`

HasIcebergQty returns a boolean if a field has been set.

### GetIsWorking

`func (o *QueryMarginAccountOrderResponse) GetIsWorking() bool`

GetIsWorking returns the IsWorking field if non-nil, zero value otherwise.

### GetIsWorkingOk

`func (o *QueryMarginAccountOrderResponse) GetIsWorkingOk() (*bool, bool)`

GetIsWorkingOk returns a tuple with the IsWorking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWorking

`func (o *QueryMarginAccountOrderResponse) SetIsWorking(v bool)`

SetIsWorking sets IsWorking field to given value.

### HasIsWorking

`func (o *QueryMarginAccountOrderResponse) HasIsWorking() bool`

HasIsWorking returns a boolean if a field has been set.

### GetOrderId

`func (o *QueryMarginAccountOrderResponse) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *QueryMarginAccountOrderResponse) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *QueryMarginAccountOrderResponse) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *QueryMarginAccountOrderResponse) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrigQty

`func (o *QueryMarginAccountOrderResponse) GetOrigQty() string`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *QueryMarginAccountOrderResponse) GetOrigQtyOk() (*string, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *QueryMarginAccountOrderResponse) SetOrigQty(v string)`

SetOrigQty sets OrigQty field to given value.

### HasOrigQty

`func (o *QueryMarginAccountOrderResponse) HasOrigQty() bool`

HasOrigQty returns a boolean if a field has been set.

### GetPrice

`func (o *QueryMarginAccountOrderResponse) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *QueryMarginAccountOrderResponse) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *QueryMarginAccountOrderResponse) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *QueryMarginAccountOrderResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetSide

`func (o *QueryMarginAccountOrderResponse) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *QueryMarginAccountOrderResponse) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *QueryMarginAccountOrderResponse) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *QueryMarginAccountOrderResponse) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetStatus

`func (o *QueryMarginAccountOrderResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QueryMarginAccountOrderResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QueryMarginAccountOrderResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *QueryMarginAccountOrderResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStopPrice

`func (o *QueryMarginAccountOrderResponse) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *QueryMarginAccountOrderResponse) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *QueryMarginAccountOrderResponse) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *QueryMarginAccountOrderResponse) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetSymbol

`func (o *QueryMarginAccountOrderResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *QueryMarginAccountOrderResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *QueryMarginAccountOrderResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *QueryMarginAccountOrderResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTime

`func (o *QueryMarginAccountOrderResponse) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *QueryMarginAccountOrderResponse) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *QueryMarginAccountOrderResponse) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *QueryMarginAccountOrderResponse) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTimeInForce

`func (o *QueryMarginAccountOrderResponse) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *QueryMarginAccountOrderResponse) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *QueryMarginAccountOrderResponse) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *QueryMarginAccountOrderResponse) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetType

`func (o *QueryMarginAccountOrderResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QueryMarginAccountOrderResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QueryMarginAccountOrderResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *QueryMarginAccountOrderResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdateTime

`func (o *QueryMarginAccountOrderResponse) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *QueryMarginAccountOrderResponse) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *QueryMarginAccountOrderResponse) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *QueryMarginAccountOrderResponse) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetAccountId

`func (o *QueryMarginAccountOrderResponse) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *QueryMarginAccountOrderResponse) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *QueryMarginAccountOrderResponse) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *QueryMarginAccountOrderResponse) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetSelfTradePreventionMode

`func (o *QueryMarginAccountOrderResponse) GetSelfTradePreventionMode() string`

GetSelfTradePreventionMode returns the SelfTradePreventionMode field if non-nil, zero value otherwise.

### GetSelfTradePreventionModeOk

`func (o *QueryMarginAccountOrderResponse) GetSelfTradePreventionModeOk() (*string, bool)`

GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfTradePreventionMode

`func (o *QueryMarginAccountOrderResponse) SetSelfTradePreventionMode(v string)`

SetSelfTradePreventionMode sets SelfTradePreventionMode field to given value.

### HasSelfTradePreventionMode

`func (o *QueryMarginAccountOrderResponse) HasSelfTradePreventionMode() bool`

HasSelfTradePreventionMode returns a boolean if a field has been set.

### GetPreventedMatchId

`func (o *QueryMarginAccountOrderResponse) GetPreventedMatchId() string`

GetPreventedMatchId returns the PreventedMatchId field if non-nil, zero value otherwise.

### GetPreventedMatchIdOk

`func (o *QueryMarginAccountOrderResponse) GetPreventedMatchIdOk() (*string, bool)`

GetPreventedMatchIdOk returns a tuple with the PreventedMatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventedMatchId

`func (o *QueryMarginAccountOrderResponse) SetPreventedMatchId(v string)`

SetPreventedMatchId sets PreventedMatchId field to given value.

### HasPreventedMatchId

`func (o *QueryMarginAccountOrderResponse) HasPreventedMatchId() bool`

HasPreventedMatchId returns a boolean if a field has been set.

### GetPreventedQuantity

`func (o *QueryMarginAccountOrderResponse) GetPreventedQuantity() string`

GetPreventedQuantity returns the PreventedQuantity field if non-nil, zero value otherwise.

### GetPreventedQuantityOk

`func (o *QueryMarginAccountOrderResponse) GetPreventedQuantityOk() (*string, bool)`

GetPreventedQuantityOk returns a tuple with the PreventedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventedQuantity

`func (o *QueryMarginAccountOrderResponse) SetPreventedQuantity(v string)`

SetPreventedQuantity sets PreventedQuantity field to given value.

### HasPreventedQuantity

`func (o *QueryMarginAccountOrderResponse) HasPreventedQuantity() bool`

HasPreventedQuantity returns a boolean if a field has been set.


[[Back to README]](../README.md)


