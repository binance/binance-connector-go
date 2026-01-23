# PlaceMultipleOrdersResponseInner

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

### NewPlaceMultipleOrdersResponseInner

`func NewPlaceMultipleOrdersResponseInner() *PlaceMultipleOrdersResponseInner`

NewPlaceMultipleOrdersResponseInner instantiates a new PlaceMultipleOrdersResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaceMultipleOrdersResponseInnerWithDefaults

`func NewPlaceMultipleOrdersResponseInnerWithDefaults() *PlaceMultipleOrdersResponseInner`

NewPlaceMultipleOrdersResponseInnerWithDefaults instantiates a new PlaceMultipleOrdersResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *PlaceMultipleOrdersResponseInner) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *PlaceMultipleOrdersResponseInner) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *PlaceMultipleOrdersResponseInner) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *PlaceMultipleOrdersResponseInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetSymbol

`func (o *PlaceMultipleOrdersResponseInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *PlaceMultipleOrdersResponseInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *PlaceMultipleOrdersResponseInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *PlaceMultipleOrdersResponseInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetPrice

`func (o *PlaceMultipleOrdersResponseInner) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PlaceMultipleOrdersResponseInner) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PlaceMultipleOrdersResponseInner) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *PlaceMultipleOrdersResponseInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQuantity

`func (o *PlaceMultipleOrdersResponseInner) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PlaceMultipleOrdersResponseInner) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PlaceMultipleOrdersResponseInner) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *PlaceMultipleOrdersResponseInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetExecutedQty

`func (o *PlaceMultipleOrdersResponseInner) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *PlaceMultipleOrdersResponseInner) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *PlaceMultipleOrdersResponseInner) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *PlaceMultipleOrdersResponseInner) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetSide

`func (o *PlaceMultipleOrdersResponseInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *PlaceMultipleOrdersResponseInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *PlaceMultipleOrdersResponseInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *PlaceMultipleOrdersResponseInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetType

`func (o *PlaceMultipleOrdersResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlaceMultipleOrdersResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlaceMultipleOrdersResponseInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PlaceMultipleOrdersResponseInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTimeInForce

`func (o *PlaceMultipleOrdersResponseInner) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *PlaceMultipleOrdersResponseInner) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *PlaceMultipleOrdersResponseInner) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *PlaceMultipleOrdersResponseInner) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetReduceOnly

`func (o *PlaceMultipleOrdersResponseInner) GetReduceOnly() bool`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *PlaceMultipleOrdersResponseInner) GetReduceOnlyOk() (*bool, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *PlaceMultipleOrdersResponseInner) SetReduceOnly(v bool)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *PlaceMultipleOrdersResponseInner) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetCreateTime

`func (o *PlaceMultipleOrdersResponseInner) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *PlaceMultipleOrdersResponseInner) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *PlaceMultipleOrdersResponseInner) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *PlaceMultipleOrdersResponseInner) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *PlaceMultipleOrdersResponseInner) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *PlaceMultipleOrdersResponseInner) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *PlaceMultipleOrdersResponseInner) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *PlaceMultipleOrdersResponseInner) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetStatus

`func (o *PlaceMultipleOrdersResponseInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PlaceMultipleOrdersResponseInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PlaceMultipleOrdersResponseInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PlaceMultipleOrdersResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAvgPrice

`func (o *PlaceMultipleOrdersResponseInner) GetAvgPrice() string`

GetAvgPrice returns the AvgPrice field if non-nil, zero value otherwise.

### GetAvgPriceOk

`func (o *PlaceMultipleOrdersResponseInner) GetAvgPriceOk() (*string, bool)`

GetAvgPriceOk returns a tuple with the AvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPrice

`func (o *PlaceMultipleOrdersResponseInner) SetAvgPrice(v string)`

SetAvgPrice sets AvgPrice field to given value.

### HasAvgPrice

`func (o *PlaceMultipleOrdersResponseInner) HasAvgPrice() bool`

HasAvgPrice returns a boolean if a field has been set.

### GetSource

`func (o *PlaceMultipleOrdersResponseInner) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PlaceMultipleOrdersResponseInner) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PlaceMultipleOrdersResponseInner) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *PlaceMultipleOrdersResponseInner) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetClientOrderId

`func (o *PlaceMultipleOrdersResponseInner) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *PlaceMultipleOrdersResponseInner) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *PlaceMultipleOrdersResponseInner) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *PlaceMultipleOrdersResponseInner) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetPriceScale

`func (o *PlaceMultipleOrdersResponseInner) GetPriceScale() int64`

GetPriceScale returns the PriceScale field if non-nil, zero value otherwise.

### GetPriceScaleOk

`func (o *PlaceMultipleOrdersResponseInner) GetPriceScaleOk() (*int64, bool)`

GetPriceScaleOk returns a tuple with the PriceScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceScale

`func (o *PlaceMultipleOrdersResponseInner) SetPriceScale(v int64)`

SetPriceScale sets PriceScale field to given value.

### HasPriceScale

`func (o *PlaceMultipleOrdersResponseInner) HasPriceScale() bool`

HasPriceScale returns a boolean if a field has been set.

### GetQuantityScale

`func (o *PlaceMultipleOrdersResponseInner) GetQuantityScale() int64`

GetQuantityScale returns the QuantityScale field if non-nil, zero value otherwise.

### GetQuantityScaleOk

`func (o *PlaceMultipleOrdersResponseInner) GetQuantityScaleOk() (*int64, bool)`

GetQuantityScaleOk returns a tuple with the QuantityScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityScale

`func (o *PlaceMultipleOrdersResponseInner) SetQuantityScale(v int64)`

SetQuantityScale sets QuantityScale field to given value.

### HasQuantityScale

`func (o *PlaceMultipleOrdersResponseInner) HasQuantityScale() bool`

HasQuantityScale returns a boolean if a field has been set.

### GetOptionSide

`func (o *PlaceMultipleOrdersResponseInner) GetOptionSide() string`

GetOptionSide returns the OptionSide field if non-nil, zero value otherwise.

### GetOptionSideOk

`func (o *PlaceMultipleOrdersResponseInner) GetOptionSideOk() (*string, bool)`

GetOptionSideOk returns a tuple with the OptionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSide

`func (o *PlaceMultipleOrdersResponseInner) SetOptionSide(v string)`

SetOptionSide sets OptionSide field to given value.

### HasOptionSide

`func (o *PlaceMultipleOrdersResponseInner) HasOptionSide() bool`

HasOptionSide returns a boolean if a field has been set.

### GetQuoteAsset

`func (o *PlaceMultipleOrdersResponseInner) GetQuoteAsset() string`

GetQuoteAsset returns the QuoteAsset field if non-nil, zero value otherwise.

### GetQuoteAssetOk

`func (o *PlaceMultipleOrdersResponseInner) GetQuoteAssetOk() (*string, bool)`

GetQuoteAssetOk returns a tuple with the QuoteAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAsset

`func (o *PlaceMultipleOrdersResponseInner) SetQuoteAsset(v string)`

SetQuoteAsset sets QuoteAsset field to given value.

### HasQuoteAsset

`func (o *PlaceMultipleOrdersResponseInner) HasQuoteAsset() bool`

HasQuoteAsset returns a boolean if a field has been set.

### GetMmp

`func (o *PlaceMultipleOrdersResponseInner) GetMmp() bool`

GetMmp returns the Mmp field if non-nil, zero value otherwise.

### GetMmpOk

`func (o *PlaceMultipleOrdersResponseInner) GetMmpOk() (*bool, bool)`

GetMmpOk returns a tuple with the Mmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmp

`func (o *PlaceMultipleOrdersResponseInner) SetMmp(v bool)`

SetMmp sets Mmp field to given value.

### HasMmp

`func (o *PlaceMultipleOrdersResponseInner) HasMmp() bool`

HasMmp returns a boolean if a field has been set.


[[Back to README]](../README.md)


