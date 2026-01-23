# NewOrderResponse

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
**Source** | Pointer to **string** |  | [optional] 
**ClientOrderId** | Pointer to **string** |  | [optional] 
**PriceScale** | Pointer to **int64** |  | [optional] 
**QuantityScale** | Pointer to **int64** |  | [optional] 
**OptionSide** | Pointer to **string** |  | [optional] 
**QuoteAsset** | Pointer to **string** |  | [optional] 
**Mmp** | Pointer to **bool** |  | [optional] 

## Methods

### NewNewOrderResponse

`func NewNewOrderResponse() *NewOrderResponse`

NewNewOrderResponse instantiates a new NewOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewOrderResponseWithDefaults

`func NewNewOrderResponseWithDefaults() *NewOrderResponse`

NewNewOrderResponseWithDefaults instantiates a new NewOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *NewOrderResponse) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *NewOrderResponse) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *NewOrderResponse) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *NewOrderResponse) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetSymbol

`func (o *NewOrderResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *NewOrderResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *NewOrderResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *NewOrderResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetPrice

`func (o *NewOrderResponse) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *NewOrderResponse) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *NewOrderResponse) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *NewOrderResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQuantity

`func (o *NewOrderResponse) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *NewOrderResponse) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *NewOrderResponse) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *NewOrderResponse) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetExecutedQty

`func (o *NewOrderResponse) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *NewOrderResponse) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *NewOrderResponse) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *NewOrderResponse) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetSide

`func (o *NewOrderResponse) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *NewOrderResponse) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *NewOrderResponse) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *NewOrderResponse) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetType

`func (o *NewOrderResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NewOrderResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NewOrderResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NewOrderResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTimeInForce

`func (o *NewOrderResponse) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *NewOrderResponse) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *NewOrderResponse) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *NewOrderResponse) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetReduceOnly

`func (o *NewOrderResponse) GetReduceOnly() bool`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *NewOrderResponse) GetReduceOnlyOk() (*bool, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *NewOrderResponse) SetReduceOnly(v bool)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *NewOrderResponse) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetCreateTime

`func (o *NewOrderResponse) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *NewOrderResponse) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *NewOrderResponse) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *NewOrderResponse) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *NewOrderResponse) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *NewOrderResponse) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *NewOrderResponse) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *NewOrderResponse) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetStatus

`func (o *NewOrderResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NewOrderResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NewOrderResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NewOrderResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAvgPrice

`func (o *NewOrderResponse) GetAvgPrice() string`

GetAvgPrice returns the AvgPrice field if non-nil, zero value otherwise.

### GetAvgPriceOk

`func (o *NewOrderResponse) GetAvgPriceOk() (*string, bool)`

GetAvgPriceOk returns a tuple with the AvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPrice

`func (o *NewOrderResponse) SetAvgPrice(v string)`

SetAvgPrice sets AvgPrice field to given value.

### HasAvgPrice

`func (o *NewOrderResponse) HasAvgPrice() bool`

HasAvgPrice returns a boolean if a field has been set.

### GetSource

`func (o *NewOrderResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *NewOrderResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *NewOrderResponse) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *NewOrderResponse) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetClientOrderId

`func (o *NewOrderResponse) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *NewOrderResponse) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *NewOrderResponse) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *NewOrderResponse) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetPriceScale

`func (o *NewOrderResponse) GetPriceScale() int64`

GetPriceScale returns the PriceScale field if non-nil, zero value otherwise.

### GetPriceScaleOk

`func (o *NewOrderResponse) GetPriceScaleOk() (*int64, bool)`

GetPriceScaleOk returns a tuple with the PriceScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceScale

`func (o *NewOrderResponse) SetPriceScale(v int64)`

SetPriceScale sets PriceScale field to given value.

### HasPriceScale

`func (o *NewOrderResponse) HasPriceScale() bool`

HasPriceScale returns a boolean if a field has been set.

### GetQuantityScale

`func (o *NewOrderResponse) GetQuantityScale() int64`

GetQuantityScale returns the QuantityScale field if non-nil, zero value otherwise.

### GetQuantityScaleOk

`func (o *NewOrderResponse) GetQuantityScaleOk() (*int64, bool)`

GetQuantityScaleOk returns a tuple with the QuantityScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityScale

`func (o *NewOrderResponse) SetQuantityScale(v int64)`

SetQuantityScale sets QuantityScale field to given value.

### HasQuantityScale

`func (o *NewOrderResponse) HasQuantityScale() bool`

HasQuantityScale returns a boolean if a field has been set.

### GetOptionSide

`func (o *NewOrderResponse) GetOptionSide() string`

GetOptionSide returns the OptionSide field if non-nil, zero value otherwise.

### GetOptionSideOk

`func (o *NewOrderResponse) GetOptionSideOk() (*string, bool)`

GetOptionSideOk returns a tuple with the OptionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSide

`func (o *NewOrderResponse) SetOptionSide(v string)`

SetOptionSide sets OptionSide field to given value.

### HasOptionSide

`func (o *NewOrderResponse) HasOptionSide() bool`

HasOptionSide returns a boolean if a field has been set.

### GetQuoteAsset

`func (o *NewOrderResponse) GetQuoteAsset() string`

GetQuoteAsset returns the QuoteAsset field if non-nil, zero value otherwise.

### GetQuoteAssetOk

`func (o *NewOrderResponse) GetQuoteAssetOk() (*string, bool)`

GetQuoteAssetOk returns a tuple with the QuoteAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAsset

`func (o *NewOrderResponse) SetQuoteAsset(v string)`

SetQuoteAsset sets QuoteAsset field to given value.

### HasQuoteAsset

`func (o *NewOrderResponse) HasQuoteAsset() bool`

HasQuoteAsset returns a boolean if a field has been set.

### GetMmp

`func (o *NewOrderResponse) GetMmp() bool`

GetMmp returns the Mmp field if non-nil, zero value otherwise.

### GetMmpOk

`func (o *NewOrderResponse) GetMmpOk() (*bool, bool)`

GetMmpOk returns a tuple with the Mmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmp

`func (o *NewOrderResponse) SetMmp(v bool)`

SetMmp sets Mmp field to given value.

### HasMmp

`func (o *NewOrderResponse) HasMmp() bool`

HasMmp returns a boolean if a field has been set.


[[Back to README]](../README.md)


