# UmAccountTradeListResponseInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**Qty** | Pointer to **string** |  | [optional] 
**RealizedPnl** | Pointer to **string** |  | [optional] 
**QuoteQty** | Pointer to **string** |  | [optional] 
**Commission** | Pointer to **string** |  | [optional] 
**CommissionAsset** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 
**Buyer** | Pointer to **bool** |  | [optional] 
**Maker** | Pointer to **bool** |  | [optional] 
**PositionSide** | Pointer to **string** |  | [optional] 

## Methods

### NewUmAccountTradeListResponseInner

`func NewUmAccountTradeListResponseInner() *UmAccountTradeListResponseInner`

NewUmAccountTradeListResponseInner instantiates a new UmAccountTradeListResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUmAccountTradeListResponseInnerWithDefaults

`func NewUmAccountTradeListResponseInnerWithDefaults() *UmAccountTradeListResponseInner`

NewUmAccountTradeListResponseInnerWithDefaults instantiates a new UmAccountTradeListResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *UmAccountTradeListResponseInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *UmAccountTradeListResponseInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *UmAccountTradeListResponseInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *UmAccountTradeListResponseInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetId

`func (o *UmAccountTradeListResponseInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UmAccountTradeListResponseInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UmAccountTradeListResponseInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UmAccountTradeListResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrderId

`func (o *UmAccountTradeListResponseInner) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *UmAccountTradeListResponseInner) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *UmAccountTradeListResponseInner) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *UmAccountTradeListResponseInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetSide

`func (o *UmAccountTradeListResponseInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *UmAccountTradeListResponseInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *UmAccountTradeListResponseInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *UmAccountTradeListResponseInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetPrice

`func (o *UmAccountTradeListResponseInner) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *UmAccountTradeListResponseInner) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *UmAccountTradeListResponseInner) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *UmAccountTradeListResponseInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQty

`func (o *UmAccountTradeListResponseInner) GetQty() string`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *UmAccountTradeListResponseInner) GetQtyOk() (*string, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *UmAccountTradeListResponseInner) SetQty(v string)`

SetQty sets Qty field to given value.

### HasQty

`func (o *UmAccountTradeListResponseInner) HasQty() bool`

HasQty returns a boolean if a field has been set.

### GetRealizedPnl

`func (o *UmAccountTradeListResponseInner) GetRealizedPnl() string`

GetRealizedPnl returns the RealizedPnl field if non-nil, zero value otherwise.

### GetRealizedPnlOk

`func (o *UmAccountTradeListResponseInner) GetRealizedPnlOk() (*string, bool)`

GetRealizedPnlOk returns a tuple with the RealizedPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedPnl

`func (o *UmAccountTradeListResponseInner) SetRealizedPnl(v string)`

SetRealizedPnl sets RealizedPnl field to given value.

### HasRealizedPnl

`func (o *UmAccountTradeListResponseInner) HasRealizedPnl() bool`

HasRealizedPnl returns a boolean if a field has been set.

### GetQuoteQty

`func (o *UmAccountTradeListResponseInner) GetQuoteQty() string`

GetQuoteQty returns the QuoteQty field if non-nil, zero value otherwise.

### GetQuoteQtyOk

`func (o *UmAccountTradeListResponseInner) GetQuoteQtyOk() (*string, bool)`

GetQuoteQtyOk returns a tuple with the QuoteQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteQty

`func (o *UmAccountTradeListResponseInner) SetQuoteQty(v string)`

SetQuoteQty sets QuoteQty field to given value.

### HasQuoteQty

`func (o *UmAccountTradeListResponseInner) HasQuoteQty() bool`

HasQuoteQty returns a boolean if a field has been set.

### GetCommission

`func (o *UmAccountTradeListResponseInner) GetCommission() string`

GetCommission returns the Commission field if non-nil, zero value otherwise.

### GetCommissionOk

`func (o *UmAccountTradeListResponseInner) GetCommissionOk() (*string, bool)`

GetCommissionOk returns a tuple with the Commission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommission

`func (o *UmAccountTradeListResponseInner) SetCommission(v string)`

SetCommission sets Commission field to given value.

### HasCommission

`func (o *UmAccountTradeListResponseInner) HasCommission() bool`

HasCommission returns a boolean if a field has been set.

### GetCommissionAsset

`func (o *UmAccountTradeListResponseInner) GetCommissionAsset() string`

GetCommissionAsset returns the CommissionAsset field if non-nil, zero value otherwise.

### GetCommissionAssetOk

`func (o *UmAccountTradeListResponseInner) GetCommissionAssetOk() (*string, bool)`

GetCommissionAssetOk returns a tuple with the CommissionAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionAsset

`func (o *UmAccountTradeListResponseInner) SetCommissionAsset(v string)`

SetCommissionAsset sets CommissionAsset field to given value.

### HasCommissionAsset

`func (o *UmAccountTradeListResponseInner) HasCommissionAsset() bool`

HasCommissionAsset returns a boolean if a field has been set.

### GetTime

`func (o *UmAccountTradeListResponseInner) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *UmAccountTradeListResponseInner) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *UmAccountTradeListResponseInner) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *UmAccountTradeListResponseInner) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetBuyer

`func (o *UmAccountTradeListResponseInner) GetBuyer() bool`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *UmAccountTradeListResponseInner) GetBuyerOk() (*bool, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *UmAccountTradeListResponseInner) SetBuyer(v bool)`

SetBuyer sets Buyer field to given value.

### HasBuyer

`func (o *UmAccountTradeListResponseInner) HasBuyer() bool`

HasBuyer returns a boolean if a field has been set.

### GetMaker

`func (o *UmAccountTradeListResponseInner) GetMaker() bool`

GetMaker returns the Maker field if non-nil, zero value otherwise.

### GetMakerOk

`func (o *UmAccountTradeListResponseInner) GetMakerOk() (*bool, bool)`

GetMakerOk returns a tuple with the Maker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaker

`func (o *UmAccountTradeListResponseInner) SetMaker(v bool)`

SetMaker sets Maker field to given value.

### HasMaker

`func (o *UmAccountTradeListResponseInner) HasMaker() bool`

HasMaker returns a boolean if a field has been set.

### GetPositionSide

`func (o *UmAccountTradeListResponseInner) GetPositionSide() string`

GetPositionSide returns the PositionSide field if non-nil, zero value otherwise.

### GetPositionSideOk

`func (o *UmAccountTradeListResponseInner) GetPositionSideOk() (*string, bool)`

GetPositionSideOk returns a tuple with the PositionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionSide

`func (o *UmAccountTradeListResponseInner) SetPositionSide(v string)`

SetPositionSide sets PositionSide field to given value.

### HasPositionSide

`func (o *UmAccountTradeListResponseInner) HasPositionSide() bool`

HasPositionSide returns a boolean if a field has been set.


[[Back to README]](../README.md)


