# QueryCurrentCmOpenOrderResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**AvgPrice** | Pointer to **string** |  | [optional] 
**ClientOrderId** | Pointer to **string** |  | [optional] 
**CumBase** | Pointer to **string** |  | [optional] 
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
**Pair** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 
**TimeInForce** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewQueryCurrentCmOpenOrderResponse

`func NewQueryCurrentCmOpenOrderResponse() *QueryCurrentCmOpenOrderResponse`

NewQueryCurrentCmOpenOrderResponse instantiates a new QueryCurrentCmOpenOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryCurrentCmOpenOrderResponseWithDefaults

`func NewQueryCurrentCmOpenOrderResponseWithDefaults() *QueryCurrentCmOpenOrderResponse`

NewQueryCurrentCmOpenOrderResponseWithDefaults instantiates a new QueryCurrentCmOpenOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgPrice

`func (o *QueryCurrentCmOpenOrderResponse) GetAvgPrice() string`

GetAvgPrice returns the AvgPrice field if non-nil, zero value otherwise.

### GetAvgPriceOk

`func (o *QueryCurrentCmOpenOrderResponse) GetAvgPriceOk() (*string, bool)`

GetAvgPriceOk returns a tuple with the AvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPrice

`func (o *QueryCurrentCmOpenOrderResponse) SetAvgPrice(v string)`

SetAvgPrice sets AvgPrice field to given value.

### HasAvgPrice

`func (o *QueryCurrentCmOpenOrderResponse) HasAvgPrice() bool`

HasAvgPrice returns a boolean if a field has been set.

### GetClientOrderId

`func (o *QueryCurrentCmOpenOrderResponse) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *QueryCurrentCmOpenOrderResponse) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *QueryCurrentCmOpenOrderResponse) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *QueryCurrentCmOpenOrderResponse) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetCumBase

`func (o *QueryCurrentCmOpenOrderResponse) GetCumBase() string`

GetCumBase returns the CumBase field if non-nil, zero value otherwise.

### GetCumBaseOk

`func (o *QueryCurrentCmOpenOrderResponse) GetCumBaseOk() (*string, bool)`

GetCumBaseOk returns a tuple with the CumBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumBase

`func (o *QueryCurrentCmOpenOrderResponse) SetCumBase(v string)`

SetCumBase sets CumBase field to given value.

### HasCumBase

`func (o *QueryCurrentCmOpenOrderResponse) HasCumBase() bool`

HasCumBase returns a boolean if a field has been set.

### GetExecutedQty

`func (o *QueryCurrentCmOpenOrderResponse) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *QueryCurrentCmOpenOrderResponse) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *QueryCurrentCmOpenOrderResponse) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *QueryCurrentCmOpenOrderResponse) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetOrderId

`func (o *QueryCurrentCmOpenOrderResponse) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *QueryCurrentCmOpenOrderResponse) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *QueryCurrentCmOpenOrderResponse) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *QueryCurrentCmOpenOrderResponse) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrigQty

`func (o *QueryCurrentCmOpenOrderResponse) GetOrigQty() string`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *QueryCurrentCmOpenOrderResponse) GetOrigQtyOk() (*string, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *QueryCurrentCmOpenOrderResponse) SetOrigQty(v string)`

SetOrigQty sets OrigQty field to given value.

### HasOrigQty

`func (o *QueryCurrentCmOpenOrderResponse) HasOrigQty() bool`

HasOrigQty returns a boolean if a field has been set.

### GetOrigType

`func (o *QueryCurrentCmOpenOrderResponse) GetOrigType() string`

GetOrigType returns the OrigType field if non-nil, zero value otherwise.

### GetOrigTypeOk

`func (o *QueryCurrentCmOpenOrderResponse) GetOrigTypeOk() (*string, bool)`

GetOrigTypeOk returns a tuple with the OrigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigType

`func (o *QueryCurrentCmOpenOrderResponse) SetOrigType(v string)`

SetOrigType sets OrigType field to given value.

### HasOrigType

`func (o *QueryCurrentCmOpenOrderResponse) HasOrigType() bool`

HasOrigType returns a boolean if a field has been set.

### GetPrice

`func (o *QueryCurrentCmOpenOrderResponse) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *QueryCurrentCmOpenOrderResponse) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *QueryCurrentCmOpenOrderResponse) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *QueryCurrentCmOpenOrderResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetReduceOnly

`func (o *QueryCurrentCmOpenOrderResponse) GetReduceOnly() bool`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *QueryCurrentCmOpenOrderResponse) GetReduceOnlyOk() (*bool, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *QueryCurrentCmOpenOrderResponse) SetReduceOnly(v bool)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *QueryCurrentCmOpenOrderResponse) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetSide

`func (o *QueryCurrentCmOpenOrderResponse) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *QueryCurrentCmOpenOrderResponse) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *QueryCurrentCmOpenOrderResponse) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *QueryCurrentCmOpenOrderResponse) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetPositionSide

`func (o *QueryCurrentCmOpenOrderResponse) GetPositionSide() string`

GetPositionSide returns the PositionSide field if non-nil, zero value otherwise.

### GetPositionSideOk

`func (o *QueryCurrentCmOpenOrderResponse) GetPositionSideOk() (*string, bool)`

GetPositionSideOk returns a tuple with the PositionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionSide

`func (o *QueryCurrentCmOpenOrderResponse) SetPositionSide(v string)`

SetPositionSide sets PositionSide field to given value.

### HasPositionSide

`func (o *QueryCurrentCmOpenOrderResponse) HasPositionSide() bool`

HasPositionSide returns a boolean if a field has been set.

### GetStatus

`func (o *QueryCurrentCmOpenOrderResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QueryCurrentCmOpenOrderResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QueryCurrentCmOpenOrderResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *QueryCurrentCmOpenOrderResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSymbol

`func (o *QueryCurrentCmOpenOrderResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *QueryCurrentCmOpenOrderResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *QueryCurrentCmOpenOrderResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *QueryCurrentCmOpenOrderResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetPair

`func (o *QueryCurrentCmOpenOrderResponse) GetPair() string`

GetPair returns the Pair field if non-nil, zero value otherwise.

### GetPairOk

`func (o *QueryCurrentCmOpenOrderResponse) GetPairOk() (*string, bool)`

GetPairOk returns a tuple with the Pair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPair

`func (o *QueryCurrentCmOpenOrderResponse) SetPair(v string)`

SetPair sets Pair field to given value.

### HasPair

`func (o *QueryCurrentCmOpenOrderResponse) HasPair() bool`

HasPair returns a boolean if a field has been set.

### GetTime

`func (o *QueryCurrentCmOpenOrderResponse) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *QueryCurrentCmOpenOrderResponse) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *QueryCurrentCmOpenOrderResponse) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *QueryCurrentCmOpenOrderResponse) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTimeInForce

`func (o *QueryCurrentCmOpenOrderResponse) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *QueryCurrentCmOpenOrderResponse) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *QueryCurrentCmOpenOrderResponse) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *QueryCurrentCmOpenOrderResponse) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetType

`func (o *QueryCurrentCmOpenOrderResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QueryCurrentCmOpenOrderResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QueryCurrentCmOpenOrderResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *QueryCurrentCmOpenOrderResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdateTime

`func (o *QueryCurrentCmOpenOrderResponse) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *QueryCurrentCmOpenOrderResponse) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *QueryCurrentCmOpenOrderResponse) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *QueryCurrentCmOpenOrderResponse) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


