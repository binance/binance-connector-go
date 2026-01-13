# MarginAccountTradeListResponseInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Commission** | Pointer to **string** |  | [optional] 
**CommissionAsset** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IsBestMatch** | Pointer to **bool** |  | [optional] 
**IsBuyer** | Pointer to **bool** |  | [optional] 
**IsMaker** | Pointer to **bool** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**Qty** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 

## Methods

### NewMarginAccountTradeListResponseInner

`func NewMarginAccountTradeListResponseInner() *MarginAccountTradeListResponseInner`

NewMarginAccountTradeListResponseInner instantiates a new MarginAccountTradeListResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarginAccountTradeListResponseInnerWithDefaults

`func NewMarginAccountTradeListResponseInnerWithDefaults() *MarginAccountTradeListResponseInner`

NewMarginAccountTradeListResponseInnerWithDefaults instantiates a new MarginAccountTradeListResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommission

`func (o *MarginAccountTradeListResponseInner) GetCommission() string`

GetCommission returns the Commission field if non-nil, zero value otherwise.

### GetCommissionOk

`func (o *MarginAccountTradeListResponseInner) GetCommissionOk() (*string, bool)`

GetCommissionOk returns a tuple with the Commission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommission

`func (o *MarginAccountTradeListResponseInner) SetCommission(v string)`

SetCommission sets Commission field to given value.

### HasCommission

`func (o *MarginAccountTradeListResponseInner) HasCommission() bool`

HasCommission returns a boolean if a field has been set.

### GetCommissionAsset

`func (o *MarginAccountTradeListResponseInner) GetCommissionAsset() string`

GetCommissionAsset returns the CommissionAsset field if non-nil, zero value otherwise.

### GetCommissionAssetOk

`func (o *MarginAccountTradeListResponseInner) GetCommissionAssetOk() (*string, bool)`

GetCommissionAssetOk returns a tuple with the CommissionAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionAsset

`func (o *MarginAccountTradeListResponseInner) SetCommissionAsset(v string)`

SetCommissionAsset sets CommissionAsset field to given value.

### HasCommissionAsset

`func (o *MarginAccountTradeListResponseInner) HasCommissionAsset() bool`

HasCommissionAsset returns a boolean if a field has been set.

### GetId

`func (o *MarginAccountTradeListResponseInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarginAccountTradeListResponseInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarginAccountTradeListResponseInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *MarginAccountTradeListResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsBestMatch

`func (o *MarginAccountTradeListResponseInner) GetIsBestMatch() bool`

GetIsBestMatch returns the IsBestMatch field if non-nil, zero value otherwise.

### GetIsBestMatchOk

`func (o *MarginAccountTradeListResponseInner) GetIsBestMatchOk() (*bool, bool)`

GetIsBestMatchOk returns a tuple with the IsBestMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBestMatch

`func (o *MarginAccountTradeListResponseInner) SetIsBestMatch(v bool)`

SetIsBestMatch sets IsBestMatch field to given value.

### HasIsBestMatch

`func (o *MarginAccountTradeListResponseInner) HasIsBestMatch() bool`

HasIsBestMatch returns a boolean if a field has been set.

### GetIsBuyer

`func (o *MarginAccountTradeListResponseInner) GetIsBuyer() bool`

GetIsBuyer returns the IsBuyer field if non-nil, zero value otherwise.

### GetIsBuyerOk

`func (o *MarginAccountTradeListResponseInner) GetIsBuyerOk() (*bool, bool)`

GetIsBuyerOk returns a tuple with the IsBuyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuyer

`func (o *MarginAccountTradeListResponseInner) SetIsBuyer(v bool)`

SetIsBuyer sets IsBuyer field to given value.

### HasIsBuyer

`func (o *MarginAccountTradeListResponseInner) HasIsBuyer() bool`

HasIsBuyer returns a boolean if a field has been set.

### GetIsMaker

`func (o *MarginAccountTradeListResponseInner) GetIsMaker() bool`

GetIsMaker returns the IsMaker field if non-nil, zero value otherwise.

### GetIsMakerOk

`func (o *MarginAccountTradeListResponseInner) GetIsMakerOk() (*bool, bool)`

GetIsMakerOk returns a tuple with the IsMaker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaker

`func (o *MarginAccountTradeListResponseInner) SetIsMaker(v bool)`

SetIsMaker sets IsMaker field to given value.

### HasIsMaker

`func (o *MarginAccountTradeListResponseInner) HasIsMaker() bool`

HasIsMaker returns a boolean if a field has been set.

### GetOrderId

`func (o *MarginAccountTradeListResponseInner) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *MarginAccountTradeListResponseInner) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *MarginAccountTradeListResponseInner) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *MarginAccountTradeListResponseInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetPrice

`func (o *MarginAccountTradeListResponseInner) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *MarginAccountTradeListResponseInner) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *MarginAccountTradeListResponseInner) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *MarginAccountTradeListResponseInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQty

`func (o *MarginAccountTradeListResponseInner) GetQty() string`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *MarginAccountTradeListResponseInner) GetQtyOk() (*string, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *MarginAccountTradeListResponseInner) SetQty(v string)`

SetQty sets Qty field to given value.

### HasQty

`func (o *MarginAccountTradeListResponseInner) HasQty() bool`

HasQty returns a boolean if a field has been set.

### GetSymbol

`func (o *MarginAccountTradeListResponseInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *MarginAccountTradeListResponseInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *MarginAccountTradeListResponseInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *MarginAccountTradeListResponseInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTime

`func (o *MarginAccountTradeListResponseInner) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *MarginAccountTradeListResponseInner) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *MarginAccountTradeListResponseInner) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *MarginAccountTradeListResponseInner) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


