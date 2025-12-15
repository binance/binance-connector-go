# PriceFilter

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | Pointer to **string** |  | [optional] 
**PriceExponent** | Pointer to **int32** |  | [optional] 
**MinPrice** | Pointer to **string** |  | [optional] 
**MaxPrice** | Pointer to **string** |  | [optional] 
**TickSize** | Pointer to **string** |  | [optional] 

## Methods

### NewPriceFilter

`func NewPriceFilter() *PriceFilter`

NewPriceFilter instantiates a new PriceFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceFilterWithDefaults

`func NewPriceFilterWithDefaults() *PriceFilter`

NewPriceFilterWithDefaults instantiates a new PriceFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *PriceFilter) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *PriceFilter) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *PriceFilter) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.

### HasFilterType

`func (o *PriceFilter) HasFilterType() bool`

HasFilterType returns a boolean if a field has been set.

### GetPriceExponent

`func (o *PriceFilter) GetPriceExponent() int32`

GetPriceExponent returns the PriceExponent field if non-nil, zero value otherwise.

### GetPriceExponentOk

`func (o *PriceFilter) GetPriceExponentOk() (*int32, bool)`

GetPriceExponentOk returns a tuple with the PriceExponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceExponent

`func (o *PriceFilter) SetPriceExponent(v int32)`

SetPriceExponent sets PriceExponent field to given value.

### HasPriceExponent

`func (o *PriceFilter) HasPriceExponent() bool`

HasPriceExponent returns a boolean if a field has been set.

### GetMinPrice

`func (o *PriceFilter) GetMinPrice() string`

GetMinPrice returns the MinPrice field if non-nil, zero value otherwise.

### GetMinPriceOk

`func (o *PriceFilter) GetMinPriceOk() (*string, bool)`

GetMinPriceOk returns a tuple with the MinPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPrice

`func (o *PriceFilter) SetMinPrice(v string)`

SetMinPrice sets MinPrice field to given value.

### HasMinPrice

`func (o *PriceFilter) HasMinPrice() bool`

HasMinPrice returns a boolean if a field has been set.

### GetMaxPrice

`func (o *PriceFilter) GetMaxPrice() string`

GetMaxPrice returns the MaxPrice field if non-nil, zero value otherwise.

### GetMaxPriceOk

`func (o *PriceFilter) GetMaxPriceOk() (*string, bool)`

GetMaxPriceOk returns a tuple with the MaxPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPrice

`func (o *PriceFilter) SetMaxPrice(v string)`

SetMaxPrice sets MaxPrice field to given value.

### HasMaxPrice

`func (o *PriceFilter) HasMaxPrice() bool`

HasMaxPrice returns a boolean if a field has been set.

### GetTickSize

`func (o *PriceFilter) GetTickSize() string`

GetTickSize returns the TickSize field if non-nil, zero value otherwise.

### GetTickSizeOk

`func (o *PriceFilter) GetTickSizeOk() (*string, bool)`

GetTickSizeOk returns a tuple with the TickSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickSize

`func (o *PriceFilter) SetTickSize(v string)`

SetTickSize sets TickSize field to given value.

### HasTickSize

`func (o *PriceFilter) HasTickSize() bool`

HasTickSize returns a boolean if a field has been set.


[[Back to README]](../README.md)


