# MinNotionalFilter

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | Pointer to **string** |  | [optional] 
**PriceExponent** | Pointer to **int32** |  | [optional] 
**MinNotional** | Pointer to **string** |  | [optional] 
**ApplyToMarket** | Pointer to **bool** |  | [optional] 
**AvgPriceMins** | Pointer to **int32** |  | [optional] 

## Methods

### NewMinNotionalFilter

`func NewMinNotionalFilter() *MinNotionalFilter`

NewMinNotionalFilter instantiates a new MinNotionalFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMinNotionalFilterWithDefaults

`func NewMinNotionalFilterWithDefaults() *MinNotionalFilter`

NewMinNotionalFilterWithDefaults instantiates a new MinNotionalFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *MinNotionalFilter) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *MinNotionalFilter) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *MinNotionalFilter) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.

### HasFilterType

`func (o *MinNotionalFilter) HasFilterType() bool`

HasFilterType returns a boolean if a field has been set.

### GetPriceExponent

`func (o *MinNotionalFilter) GetPriceExponent() int32`

GetPriceExponent returns the PriceExponent field if non-nil, zero value otherwise.

### GetPriceExponentOk

`func (o *MinNotionalFilter) GetPriceExponentOk() (*int32, bool)`

GetPriceExponentOk returns a tuple with the PriceExponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceExponent

`func (o *MinNotionalFilter) SetPriceExponent(v int32)`

SetPriceExponent sets PriceExponent field to given value.

### HasPriceExponent

`func (o *MinNotionalFilter) HasPriceExponent() bool`

HasPriceExponent returns a boolean if a field has been set.

### GetMinNotional

`func (o *MinNotionalFilter) GetMinNotional() string`

GetMinNotional returns the MinNotional field if non-nil, zero value otherwise.

### GetMinNotionalOk

`func (o *MinNotionalFilter) GetMinNotionalOk() (*string, bool)`

GetMinNotionalOk returns a tuple with the MinNotional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNotional

`func (o *MinNotionalFilter) SetMinNotional(v string)`

SetMinNotional sets MinNotional field to given value.

### HasMinNotional

`func (o *MinNotionalFilter) HasMinNotional() bool`

HasMinNotional returns a boolean if a field has been set.

### GetApplyToMarket

`func (o *MinNotionalFilter) GetApplyToMarket() bool`

GetApplyToMarket returns the ApplyToMarket field if non-nil, zero value otherwise.

### GetApplyToMarketOk

`func (o *MinNotionalFilter) GetApplyToMarketOk() (*bool, bool)`

GetApplyToMarketOk returns a tuple with the ApplyToMarket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyToMarket

`func (o *MinNotionalFilter) SetApplyToMarket(v bool)`

SetApplyToMarket sets ApplyToMarket field to given value.

### HasApplyToMarket

`func (o *MinNotionalFilter) HasApplyToMarket() bool`

HasApplyToMarket returns a boolean if a field has been set.

### GetAvgPriceMins

`func (o *MinNotionalFilter) GetAvgPriceMins() int32`

GetAvgPriceMins returns the AvgPriceMins field if non-nil, zero value otherwise.

### GetAvgPriceMinsOk

`func (o *MinNotionalFilter) GetAvgPriceMinsOk() (*int32, bool)`

GetAvgPriceMinsOk returns a tuple with the AvgPriceMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPriceMins

`func (o *MinNotionalFilter) SetAvgPriceMins(v int32)`

SetAvgPriceMins sets AvgPriceMins field to given value.

### HasAvgPriceMins

`func (o *MinNotionalFilter) HasAvgPriceMins() bool`

HasAvgPriceMins returns a boolean if a field has been set.


[[Back to README]](../README.md)


