# QuerySingleOrderResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **int64** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **string** |  | [optional] 
**ExecutedQty** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**TimeInForce** | Pointer to **string** |  | [optional] 
**ReduceOnly** | Pointer to **bool** |  | [optional] 
**CreateTime** | Pointer to **int64** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**AvgPrice** | Pointer to **string** |  | [optional] 
**ClientOrderId** | Pointer to **string** |  | [optional] 
**PriceScale** | Pointer to **int64** |  | [optional] 
**QuantityScale** | Pointer to **int64** |  | [optional] 
**OptionSide** | Pointer to **string** |  | [optional] 
**QuoteAsset** | Pointer to **string** |  | [optional] 
**Mmp** | Pointer to **bool** |  | [optional] 

## Methods

### NewQuerySingleOrderResponse

`func NewQuerySingleOrderResponse() *QuerySingleOrderResponse`

NewQuerySingleOrderResponse instantiates a new QuerySingleOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuerySingleOrderResponseWithDefaults

`func NewQuerySingleOrderResponseWithDefaults() *QuerySingleOrderResponse`

NewQuerySingleOrderResponseWithDefaults instantiates a new QuerySingleOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *QuerySingleOrderResponse) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *QuerySingleOrderResponse) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *QuerySingleOrderResponse) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *QuerySingleOrderResponse) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetSymbol

`func (o *QuerySingleOrderResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *QuerySingleOrderResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *QuerySingleOrderResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *QuerySingleOrderResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetPrice

`func (o *QuerySingleOrderResponse) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *QuerySingleOrderResponse) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *QuerySingleOrderResponse) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *QuerySingleOrderResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQuantity

`func (o *QuerySingleOrderResponse) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *QuerySingleOrderResponse) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *QuerySingleOrderResponse) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *QuerySingleOrderResponse) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetExecutedQty

`func (o *QuerySingleOrderResponse) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *QuerySingleOrderResponse) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *QuerySingleOrderResponse) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *QuerySingleOrderResponse) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetSide

`func (o *QuerySingleOrderResponse) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *QuerySingleOrderResponse) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *QuerySingleOrderResponse) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *QuerySingleOrderResponse) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetType

`func (o *QuerySingleOrderResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QuerySingleOrderResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QuerySingleOrderResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *QuerySingleOrderResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTimeInForce

`func (o *QuerySingleOrderResponse) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *QuerySingleOrderResponse) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *QuerySingleOrderResponse) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *QuerySingleOrderResponse) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetReduceOnly

`func (o *QuerySingleOrderResponse) GetReduceOnly() bool`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *QuerySingleOrderResponse) GetReduceOnlyOk() (*bool, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *QuerySingleOrderResponse) SetReduceOnly(v bool)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *QuerySingleOrderResponse) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetCreateTime

`func (o *QuerySingleOrderResponse) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *QuerySingleOrderResponse) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *QuerySingleOrderResponse) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *QuerySingleOrderResponse) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *QuerySingleOrderResponse) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *QuerySingleOrderResponse) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *QuerySingleOrderResponse) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *QuerySingleOrderResponse) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetStatus

`func (o *QuerySingleOrderResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QuerySingleOrderResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QuerySingleOrderResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *QuerySingleOrderResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAvgPrice

`func (o *QuerySingleOrderResponse) GetAvgPrice() string`

GetAvgPrice returns the AvgPrice field if non-nil, zero value otherwise.

### GetAvgPriceOk

`func (o *QuerySingleOrderResponse) GetAvgPriceOk() (*string, bool)`

GetAvgPriceOk returns a tuple with the AvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPrice

`func (o *QuerySingleOrderResponse) SetAvgPrice(v string)`

SetAvgPrice sets AvgPrice field to given value.

### HasAvgPrice

`func (o *QuerySingleOrderResponse) HasAvgPrice() bool`

HasAvgPrice returns a boolean if a field has been set.

### GetClientOrderId

`func (o *QuerySingleOrderResponse) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *QuerySingleOrderResponse) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *QuerySingleOrderResponse) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *QuerySingleOrderResponse) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetPriceScale

`func (o *QuerySingleOrderResponse) GetPriceScale() int64`

GetPriceScale returns the PriceScale field if non-nil, zero value otherwise.

### GetPriceScaleOk

`func (o *QuerySingleOrderResponse) GetPriceScaleOk() (*int64, bool)`

GetPriceScaleOk returns a tuple with the PriceScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceScale

`func (o *QuerySingleOrderResponse) SetPriceScale(v int64)`

SetPriceScale sets PriceScale field to given value.

### HasPriceScale

`func (o *QuerySingleOrderResponse) HasPriceScale() bool`

HasPriceScale returns a boolean if a field has been set.

### GetQuantityScale

`func (o *QuerySingleOrderResponse) GetQuantityScale() int64`

GetQuantityScale returns the QuantityScale field if non-nil, zero value otherwise.

### GetQuantityScaleOk

`func (o *QuerySingleOrderResponse) GetQuantityScaleOk() (*int64, bool)`

GetQuantityScaleOk returns a tuple with the QuantityScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityScale

`func (o *QuerySingleOrderResponse) SetQuantityScale(v int64)`

SetQuantityScale sets QuantityScale field to given value.

### HasQuantityScale

`func (o *QuerySingleOrderResponse) HasQuantityScale() bool`

HasQuantityScale returns a boolean if a field has been set.

### GetOptionSide

`func (o *QuerySingleOrderResponse) GetOptionSide() string`

GetOptionSide returns the OptionSide field if non-nil, zero value otherwise.

### GetOptionSideOk

`func (o *QuerySingleOrderResponse) GetOptionSideOk() (*string, bool)`

GetOptionSideOk returns a tuple with the OptionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSide

`func (o *QuerySingleOrderResponse) SetOptionSide(v string)`

SetOptionSide sets OptionSide field to given value.

### HasOptionSide

`func (o *QuerySingleOrderResponse) HasOptionSide() bool`

HasOptionSide returns a boolean if a field has been set.

### GetQuoteAsset

`func (o *QuerySingleOrderResponse) GetQuoteAsset() string`

GetQuoteAsset returns the QuoteAsset field if non-nil, zero value otherwise.

### GetQuoteAssetOk

`func (o *QuerySingleOrderResponse) GetQuoteAssetOk() (*string, bool)`

GetQuoteAssetOk returns a tuple with the QuoteAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAsset

`func (o *QuerySingleOrderResponse) SetQuoteAsset(v string)`

SetQuoteAsset sets QuoteAsset field to given value.

### HasQuoteAsset

`func (o *QuerySingleOrderResponse) HasQuoteAsset() bool`

HasQuoteAsset returns a boolean if a field has been set.

### GetMmp

`func (o *QuerySingleOrderResponse) GetMmp() bool`

GetMmp returns the Mmp field if non-nil, zero value otherwise.

### GetMmpOk

`func (o *QuerySingleOrderResponse) GetMmpOk() (*bool, bool)`

GetMmpOk returns a tuple with the Mmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmp

`func (o *QuerySingleOrderResponse) SetMmp(v bool)`

SetMmp sets Mmp field to given value.

### HasMmp

`func (o *QuerySingleOrderResponse) HasMmp() bool`

HasMmp returns a boolean if a field has been set.


[[Back to README]](../README.md)


