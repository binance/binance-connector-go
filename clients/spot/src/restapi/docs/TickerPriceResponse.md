# TickerPriceResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 

## Methods

### NewTickerPriceResponse

`func NewTickerPriceResponse() *TickerPriceResponse`

NewTickerPriceResponse instantiates a new TickerPriceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTickerPriceResponseWithDefaults

`func NewTickerPriceResponseWithDefaults() *TickerPriceResponse`

NewTickerPriceResponseWithDefaults instantiates a new TickerPriceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *TickerPriceResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *TickerPriceResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *TickerPriceResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *TickerPriceResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetPrice

`func (o *TickerPriceResponse) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *TickerPriceResponse) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *TickerPriceResponse) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *TickerPriceResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to README]](../README.md)


