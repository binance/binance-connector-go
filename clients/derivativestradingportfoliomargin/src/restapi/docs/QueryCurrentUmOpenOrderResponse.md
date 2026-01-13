# QueryCurrentUmOpenOrderResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**AvgPrice** | Pointer to **string** |  | [optional] 
**ClientOrderId** | Pointer to **string** |  | [optional] 
**CumQuote** | Pointer to **string** |  | [optional] 
**ExecutedQty** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**OrigQty** | Pointer to **string** |  | [optional] 
**OrigType** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**ReduceOnly** | Pointer to **bool** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**PositionSide** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 
**TimeInForce** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 
**SelfTradePreventionMode** | Pointer to **string** |  | [optional] 
**GoodTillDate** | Pointer to **int64** |  | [optional] 
**PriceMatch** | Pointer to **string** |  | [optional] 

## Methods

### NewQueryCurrentUmOpenOrderResponse

`func NewQueryCurrentUmOpenOrderResponse() *QueryCurrentUmOpenOrderResponse`

NewQueryCurrentUmOpenOrderResponse instantiates a new QueryCurrentUmOpenOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryCurrentUmOpenOrderResponseWithDefaults

`func NewQueryCurrentUmOpenOrderResponseWithDefaults() *QueryCurrentUmOpenOrderResponse`

NewQueryCurrentUmOpenOrderResponseWithDefaults instantiates a new QueryCurrentUmOpenOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgPrice

`func (o *QueryCurrentUmOpenOrderResponse) GetAvgPrice() string`

GetAvgPrice returns the AvgPrice field if non-nil, zero value otherwise.

### GetAvgPriceOk

`func (o *QueryCurrentUmOpenOrderResponse) GetAvgPriceOk() (*string, bool)`

GetAvgPriceOk returns a tuple with the AvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPrice

`func (o *QueryCurrentUmOpenOrderResponse) SetAvgPrice(v string)`

SetAvgPrice sets AvgPrice field to given value.

### HasAvgPrice

`func (o *QueryCurrentUmOpenOrderResponse) HasAvgPrice() bool`

HasAvgPrice returns a boolean if a field has been set.

### GetClientOrderId

`func (o *QueryCurrentUmOpenOrderResponse) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *QueryCurrentUmOpenOrderResponse) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *QueryCurrentUmOpenOrderResponse) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *QueryCurrentUmOpenOrderResponse) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetCumQuote

`func (o *QueryCurrentUmOpenOrderResponse) GetCumQuote() string`

GetCumQuote returns the CumQuote field if non-nil, zero value otherwise.

### GetCumQuoteOk

`func (o *QueryCurrentUmOpenOrderResponse) GetCumQuoteOk() (*string, bool)`

GetCumQuoteOk returns a tuple with the CumQuote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumQuote

`func (o *QueryCurrentUmOpenOrderResponse) SetCumQuote(v string)`

SetCumQuote sets CumQuote field to given value.

### HasCumQuote

`func (o *QueryCurrentUmOpenOrderResponse) HasCumQuote() bool`

HasCumQuote returns a boolean if a field has been set.

### GetExecutedQty

`func (o *QueryCurrentUmOpenOrderResponse) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *QueryCurrentUmOpenOrderResponse) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *QueryCurrentUmOpenOrderResponse) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *QueryCurrentUmOpenOrderResponse) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetOrderId

`func (o *QueryCurrentUmOpenOrderResponse) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *QueryCurrentUmOpenOrderResponse) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *QueryCurrentUmOpenOrderResponse) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *QueryCurrentUmOpenOrderResponse) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrigQty

`func (o *QueryCurrentUmOpenOrderResponse) GetOrigQty() string`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *QueryCurrentUmOpenOrderResponse) GetOrigQtyOk() (*string, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *QueryCurrentUmOpenOrderResponse) SetOrigQty(v string)`

SetOrigQty sets OrigQty field to given value.

### HasOrigQty

`func (o *QueryCurrentUmOpenOrderResponse) HasOrigQty() bool`

HasOrigQty returns a boolean if a field has been set.

### GetOrigType

`func (o *QueryCurrentUmOpenOrderResponse) GetOrigType() string`

GetOrigType returns the OrigType field if non-nil, zero value otherwise.

### GetOrigTypeOk

`func (o *QueryCurrentUmOpenOrderResponse) GetOrigTypeOk() (*string, bool)`

GetOrigTypeOk returns a tuple with the OrigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigType

`func (o *QueryCurrentUmOpenOrderResponse) SetOrigType(v string)`

SetOrigType sets OrigType field to given value.

### HasOrigType

`func (o *QueryCurrentUmOpenOrderResponse) HasOrigType() bool`

HasOrigType returns a boolean if a field has been set.

### GetPrice

`func (o *QueryCurrentUmOpenOrderResponse) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *QueryCurrentUmOpenOrderResponse) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *QueryCurrentUmOpenOrderResponse) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *QueryCurrentUmOpenOrderResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetReduceOnly

`func (o *QueryCurrentUmOpenOrderResponse) GetReduceOnly() bool`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *QueryCurrentUmOpenOrderResponse) GetReduceOnlyOk() (*bool, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *QueryCurrentUmOpenOrderResponse) SetReduceOnly(v bool)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *QueryCurrentUmOpenOrderResponse) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetSide

`func (o *QueryCurrentUmOpenOrderResponse) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *QueryCurrentUmOpenOrderResponse) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *QueryCurrentUmOpenOrderResponse) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *QueryCurrentUmOpenOrderResponse) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetPositionSide

`func (o *QueryCurrentUmOpenOrderResponse) GetPositionSide() string`

GetPositionSide returns the PositionSide field if non-nil, zero value otherwise.

### GetPositionSideOk

`func (o *QueryCurrentUmOpenOrderResponse) GetPositionSideOk() (*string, bool)`

GetPositionSideOk returns a tuple with the PositionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionSide

`func (o *QueryCurrentUmOpenOrderResponse) SetPositionSide(v string)`

SetPositionSide sets PositionSide field to given value.

### HasPositionSide

`func (o *QueryCurrentUmOpenOrderResponse) HasPositionSide() bool`

HasPositionSide returns a boolean if a field has been set.

### GetStatus

`func (o *QueryCurrentUmOpenOrderResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QueryCurrentUmOpenOrderResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QueryCurrentUmOpenOrderResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *QueryCurrentUmOpenOrderResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSymbol

`func (o *QueryCurrentUmOpenOrderResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *QueryCurrentUmOpenOrderResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *QueryCurrentUmOpenOrderResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *QueryCurrentUmOpenOrderResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTime

`func (o *QueryCurrentUmOpenOrderResponse) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *QueryCurrentUmOpenOrderResponse) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *QueryCurrentUmOpenOrderResponse) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *QueryCurrentUmOpenOrderResponse) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTimeInForce

`func (o *QueryCurrentUmOpenOrderResponse) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *QueryCurrentUmOpenOrderResponse) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *QueryCurrentUmOpenOrderResponse) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *QueryCurrentUmOpenOrderResponse) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetType

`func (o *QueryCurrentUmOpenOrderResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QueryCurrentUmOpenOrderResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QueryCurrentUmOpenOrderResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *QueryCurrentUmOpenOrderResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdateTime

`func (o *QueryCurrentUmOpenOrderResponse) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *QueryCurrentUmOpenOrderResponse) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *QueryCurrentUmOpenOrderResponse) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *QueryCurrentUmOpenOrderResponse) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetSelfTradePreventionMode

`func (o *QueryCurrentUmOpenOrderResponse) GetSelfTradePreventionMode() string`

GetSelfTradePreventionMode returns the SelfTradePreventionMode field if non-nil, zero value otherwise.

### GetSelfTradePreventionModeOk

`func (o *QueryCurrentUmOpenOrderResponse) GetSelfTradePreventionModeOk() (*string, bool)`

GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfTradePreventionMode

`func (o *QueryCurrentUmOpenOrderResponse) SetSelfTradePreventionMode(v string)`

SetSelfTradePreventionMode sets SelfTradePreventionMode field to given value.

### HasSelfTradePreventionMode

`func (o *QueryCurrentUmOpenOrderResponse) HasSelfTradePreventionMode() bool`

HasSelfTradePreventionMode returns a boolean if a field has been set.

### GetGoodTillDate

`func (o *QueryCurrentUmOpenOrderResponse) GetGoodTillDate() int64`

GetGoodTillDate returns the GoodTillDate field if non-nil, zero value otherwise.

### GetGoodTillDateOk

`func (o *QueryCurrentUmOpenOrderResponse) GetGoodTillDateOk() (*int64, bool)`

GetGoodTillDateOk returns a tuple with the GoodTillDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodTillDate

`func (o *QueryCurrentUmOpenOrderResponse) SetGoodTillDate(v int64)`

SetGoodTillDate sets GoodTillDate field to given value.

### HasGoodTillDate

`func (o *QueryCurrentUmOpenOrderResponse) HasGoodTillDate() bool`

HasGoodTillDate returns a boolean if a field has been set.

### GetPriceMatch

`func (o *QueryCurrentUmOpenOrderResponse) GetPriceMatch() string`

GetPriceMatch returns the PriceMatch field if non-nil, zero value otherwise.

### GetPriceMatchOk

`func (o *QueryCurrentUmOpenOrderResponse) GetPriceMatchOk() (*string, bool)`

GetPriceMatchOk returns a tuple with the PriceMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMatch

`func (o *QueryCurrentUmOpenOrderResponse) SetPriceMatch(v string)`

SetPriceMatch sets PriceMatch field to given value.

### HasPriceMatch

`func (o *QueryCurrentUmOpenOrderResponse) HasPriceMatch() bool`

HasPriceMatch returns a boolean if a field has been set.


[[Back to README]](../README.md)


