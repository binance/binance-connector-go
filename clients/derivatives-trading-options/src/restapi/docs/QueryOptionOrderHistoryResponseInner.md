# QueryOptionOrderHistoryResponseInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **int64** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **string** |  | [optional] 
**ExecutedQty** | Pointer to **string** |  | [optional] 
**Fee** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**TimeInForce** | Pointer to **string** |  | [optional] 
**ReduceOnly** | Pointer to **bool** |  | [optional] 
**PostOnly** | Pointer to **bool** |  | [optional] 
**CreateTime** | Pointer to **int64** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**AvgPrice** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**ClientOrderId** | Pointer to **string** |  | [optional] 
**PriceScale** | Pointer to **int64** |  | [optional] 
**QuantityScale** | Pointer to **int64** |  | [optional] 
**OptionSide** | Pointer to **string** |  | [optional] 
**QuoteAsset** | Pointer to **string** |  | [optional] 
**Mmp** | Pointer to **bool** |  | [optional] 

## Methods

### NewQueryOptionOrderHistoryResponseInner

`func NewQueryOptionOrderHistoryResponseInner() *QueryOptionOrderHistoryResponseInner`

NewQueryOptionOrderHistoryResponseInner instantiates a new QueryOptionOrderHistoryResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryOptionOrderHistoryResponseInnerWithDefaults

`func NewQueryOptionOrderHistoryResponseInnerWithDefaults() *QueryOptionOrderHistoryResponseInner`

NewQueryOptionOrderHistoryResponseInnerWithDefaults instantiates a new QueryOptionOrderHistoryResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *QueryOptionOrderHistoryResponseInner) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *QueryOptionOrderHistoryResponseInner) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *QueryOptionOrderHistoryResponseInner) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *QueryOptionOrderHistoryResponseInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetSymbol

`func (o *QueryOptionOrderHistoryResponseInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *QueryOptionOrderHistoryResponseInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *QueryOptionOrderHistoryResponseInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *QueryOptionOrderHistoryResponseInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetPrice

`func (o *QueryOptionOrderHistoryResponseInner) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *QueryOptionOrderHistoryResponseInner) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *QueryOptionOrderHistoryResponseInner) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *QueryOptionOrderHistoryResponseInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQuantity

`func (o *QueryOptionOrderHistoryResponseInner) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *QueryOptionOrderHistoryResponseInner) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *QueryOptionOrderHistoryResponseInner) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *QueryOptionOrderHistoryResponseInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetExecutedQty

`func (o *QueryOptionOrderHistoryResponseInner) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *QueryOptionOrderHistoryResponseInner) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *QueryOptionOrderHistoryResponseInner) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *QueryOptionOrderHistoryResponseInner) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetFee

`func (o *QueryOptionOrderHistoryResponseInner) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *QueryOptionOrderHistoryResponseInner) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *QueryOptionOrderHistoryResponseInner) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *QueryOptionOrderHistoryResponseInner) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetSide

`func (o *QueryOptionOrderHistoryResponseInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *QueryOptionOrderHistoryResponseInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *QueryOptionOrderHistoryResponseInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *QueryOptionOrderHistoryResponseInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetType

`func (o *QueryOptionOrderHistoryResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QueryOptionOrderHistoryResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QueryOptionOrderHistoryResponseInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *QueryOptionOrderHistoryResponseInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTimeInForce

`func (o *QueryOptionOrderHistoryResponseInner) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *QueryOptionOrderHistoryResponseInner) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *QueryOptionOrderHistoryResponseInner) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *QueryOptionOrderHistoryResponseInner) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetReduceOnly

`func (o *QueryOptionOrderHistoryResponseInner) GetReduceOnly() bool`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *QueryOptionOrderHistoryResponseInner) GetReduceOnlyOk() (*bool, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *QueryOptionOrderHistoryResponseInner) SetReduceOnly(v bool)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *QueryOptionOrderHistoryResponseInner) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetPostOnly

`func (o *QueryOptionOrderHistoryResponseInner) GetPostOnly() bool`

GetPostOnly returns the PostOnly field if non-nil, zero value otherwise.

### GetPostOnlyOk

`func (o *QueryOptionOrderHistoryResponseInner) GetPostOnlyOk() (*bool, bool)`

GetPostOnlyOk returns a tuple with the PostOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostOnly

`func (o *QueryOptionOrderHistoryResponseInner) SetPostOnly(v bool)`

SetPostOnly sets PostOnly field to given value.

### HasPostOnly

`func (o *QueryOptionOrderHistoryResponseInner) HasPostOnly() bool`

HasPostOnly returns a boolean if a field has been set.

### GetCreateTime

`func (o *QueryOptionOrderHistoryResponseInner) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *QueryOptionOrderHistoryResponseInner) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *QueryOptionOrderHistoryResponseInner) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *QueryOptionOrderHistoryResponseInner) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *QueryOptionOrderHistoryResponseInner) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *QueryOptionOrderHistoryResponseInner) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *QueryOptionOrderHistoryResponseInner) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *QueryOptionOrderHistoryResponseInner) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetStatus

`func (o *QueryOptionOrderHistoryResponseInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QueryOptionOrderHistoryResponseInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QueryOptionOrderHistoryResponseInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *QueryOptionOrderHistoryResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReason

`func (o *QueryOptionOrderHistoryResponseInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *QueryOptionOrderHistoryResponseInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *QueryOptionOrderHistoryResponseInner) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *QueryOptionOrderHistoryResponseInner) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetAvgPrice

`func (o *QueryOptionOrderHistoryResponseInner) GetAvgPrice() string`

GetAvgPrice returns the AvgPrice field if non-nil, zero value otherwise.

### GetAvgPriceOk

`func (o *QueryOptionOrderHistoryResponseInner) GetAvgPriceOk() (*string, bool)`

GetAvgPriceOk returns a tuple with the AvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPrice

`func (o *QueryOptionOrderHistoryResponseInner) SetAvgPrice(v string)`

SetAvgPrice sets AvgPrice field to given value.

### HasAvgPrice

`func (o *QueryOptionOrderHistoryResponseInner) HasAvgPrice() bool`

HasAvgPrice returns a boolean if a field has been set.

### GetSource

`func (o *QueryOptionOrderHistoryResponseInner) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *QueryOptionOrderHistoryResponseInner) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *QueryOptionOrderHistoryResponseInner) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *QueryOptionOrderHistoryResponseInner) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetClientOrderId

`func (o *QueryOptionOrderHistoryResponseInner) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *QueryOptionOrderHistoryResponseInner) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *QueryOptionOrderHistoryResponseInner) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *QueryOptionOrderHistoryResponseInner) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetPriceScale

`func (o *QueryOptionOrderHistoryResponseInner) GetPriceScale() int64`

GetPriceScale returns the PriceScale field if non-nil, zero value otherwise.

### GetPriceScaleOk

`func (o *QueryOptionOrderHistoryResponseInner) GetPriceScaleOk() (*int64, bool)`

GetPriceScaleOk returns a tuple with the PriceScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceScale

`func (o *QueryOptionOrderHistoryResponseInner) SetPriceScale(v int64)`

SetPriceScale sets PriceScale field to given value.

### HasPriceScale

`func (o *QueryOptionOrderHistoryResponseInner) HasPriceScale() bool`

HasPriceScale returns a boolean if a field has been set.

### GetQuantityScale

`func (o *QueryOptionOrderHistoryResponseInner) GetQuantityScale() int64`

GetQuantityScale returns the QuantityScale field if non-nil, zero value otherwise.

### GetQuantityScaleOk

`func (o *QueryOptionOrderHistoryResponseInner) GetQuantityScaleOk() (*int64, bool)`

GetQuantityScaleOk returns a tuple with the QuantityScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityScale

`func (o *QueryOptionOrderHistoryResponseInner) SetQuantityScale(v int64)`

SetQuantityScale sets QuantityScale field to given value.

### HasQuantityScale

`func (o *QueryOptionOrderHistoryResponseInner) HasQuantityScale() bool`

HasQuantityScale returns a boolean if a field has been set.

### GetOptionSide

`func (o *QueryOptionOrderHistoryResponseInner) GetOptionSide() string`

GetOptionSide returns the OptionSide field if non-nil, zero value otherwise.

### GetOptionSideOk

`func (o *QueryOptionOrderHistoryResponseInner) GetOptionSideOk() (*string, bool)`

GetOptionSideOk returns a tuple with the OptionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSide

`func (o *QueryOptionOrderHistoryResponseInner) SetOptionSide(v string)`

SetOptionSide sets OptionSide field to given value.

### HasOptionSide

`func (o *QueryOptionOrderHistoryResponseInner) HasOptionSide() bool`

HasOptionSide returns a boolean if a field has been set.

### GetQuoteAsset

`func (o *QueryOptionOrderHistoryResponseInner) GetQuoteAsset() string`

GetQuoteAsset returns the QuoteAsset field if non-nil, zero value otherwise.

### GetQuoteAssetOk

`func (o *QueryOptionOrderHistoryResponseInner) GetQuoteAssetOk() (*string, bool)`

GetQuoteAssetOk returns a tuple with the QuoteAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAsset

`func (o *QueryOptionOrderHistoryResponseInner) SetQuoteAsset(v string)`

SetQuoteAsset sets QuoteAsset field to given value.

### HasQuoteAsset

`func (o *QueryOptionOrderHistoryResponseInner) HasQuoteAsset() bool`

HasQuoteAsset returns a boolean if a field has been set.

### GetMmp

`func (o *QueryOptionOrderHistoryResponseInner) GetMmp() bool`

GetMmp returns the Mmp field if non-nil, zero value otherwise.

### GetMmpOk

`func (o *QueryOptionOrderHistoryResponseInner) GetMmpOk() (*bool, bool)`

GetMmpOk returns a tuple with the Mmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmp

`func (o *QueryOptionOrderHistoryResponseInner) SetMmp(v bool)`

SetMmp sets Mmp field to given value.

### HasMmp

`func (o *QueryOptionOrderHistoryResponseInner) HasMmp() bool`

HasMmp returns a boolean if a field has been set.


[[Back to README]](../README.md)


