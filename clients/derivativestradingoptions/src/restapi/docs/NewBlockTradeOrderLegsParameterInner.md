# NewBlockTradeOrderLegsParameterInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | Option trading pair | 
**Side** | [**NewBlockTradeOrderLegsParameterInnerSide**](NewBlockTradeOrderLegsParameterInnerSide.md) |  | 
**Type** | [**NewBlockTradeOrderLegsParameterInnerType**](NewBlockTradeOrderLegsParameterInnerType.md) |  | 
**Quantity** | **string** | Order quantity | 
**Price** | Pointer to **string** | Order price | [optional] 

## Methods

### NewNewBlockTradeOrderLegsParameterInner

`func NewNewBlockTradeOrderLegsParameterInner(symbol string, side NewBlockTradeOrderLegsParameterInnerSide, type_ NewBlockTradeOrderLegsParameterInnerType, quantity string, ) *NewBlockTradeOrderLegsParameterInner`

NewNewBlockTradeOrderLegsParameterInner instantiates a new NewBlockTradeOrderLegsParameterInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewBlockTradeOrderLegsParameterInnerWithDefaults

`func NewNewBlockTradeOrderLegsParameterInnerWithDefaults() *NewBlockTradeOrderLegsParameterInner`

NewNewBlockTradeOrderLegsParameterInnerWithDefaults instantiates a new NewBlockTradeOrderLegsParameterInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *NewBlockTradeOrderLegsParameterInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *NewBlockTradeOrderLegsParameterInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *NewBlockTradeOrderLegsParameterInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetSide

`func (o *NewBlockTradeOrderLegsParameterInner) GetSide() NewBlockTradeOrderLegsParameterInnerSide`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *NewBlockTradeOrderLegsParameterInner) GetSideOk() (*NewBlockTradeOrderLegsParameterInnerSide, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *NewBlockTradeOrderLegsParameterInner) SetSide(v NewBlockTradeOrderLegsParameterInnerSide)`

SetSide sets Side field to given value.


### GetType

`func (o *NewBlockTradeOrderLegsParameterInner) GetType() NewBlockTradeOrderLegsParameterInnerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NewBlockTradeOrderLegsParameterInner) GetTypeOk() (*NewBlockTradeOrderLegsParameterInnerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NewBlockTradeOrderLegsParameterInner) SetType(v NewBlockTradeOrderLegsParameterInnerType)`

SetType sets Type field to given value.


### GetQuantity

`func (o *NewBlockTradeOrderLegsParameterInner) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *NewBlockTradeOrderLegsParameterInner) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *NewBlockTradeOrderLegsParameterInner) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.


### GetPrice

`func (o *NewBlockTradeOrderLegsParameterInner) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *NewBlockTradeOrderLegsParameterInner) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *NewBlockTradeOrderLegsParameterInner) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *NewBlockTradeOrderLegsParameterInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to README]](../README.md)


