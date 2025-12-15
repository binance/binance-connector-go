# NotionalFilter

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | Pointer to **string** |  | [optional] 
**PriceExponent** | Pointer to **int32** |  | [optional] 
**MinNotional** | Pointer to **string** |  | [optional] 
**ApplyMinToMarket** | Pointer to **bool** |  | [optional] 
**MaxNotional** | Pointer to **string** |  | [optional] 
**ApplyMaxToMarket** | Pointer to **bool** |  | [optional] 
**AvgPriceMins** | Pointer to **int32** |  | [optional] 

## Methods

### NewNotionalFilter

`func NewNotionalFilter() *NotionalFilter`

NewNotionalFilter instantiates a new NotionalFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotionalFilterWithDefaults

`func NewNotionalFilterWithDefaults() *NotionalFilter`

NewNotionalFilterWithDefaults instantiates a new NotionalFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *NotionalFilter) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *NotionalFilter) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *NotionalFilter) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.

### HasFilterType

`func (o *NotionalFilter) HasFilterType() bool`

HasFilterType returns a boolean if a field has been set.

### GetPriceExponent

`func (o *NotionalFilter) GetPriceExponent() int32`

GetPriceExponent returns the PriceExponent field if non-nil, zero value otherwise.

### GetPriceExponentOk

`func (o *NotionalFilter) GetPriceExponentOk() (*int32, bool)`

GetPriceExponentOk returns a tuple with the PriceExponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceExponent

`func (o *NotionalFilter) SetPriceExponent(v int32)`

SetPriceExponent sets PriceExponent field to given value.

### HasPriceExponent

`func (o *NotionalFilter) HasPriceExponent() bool`

HasPriceExponent returns a boolean if a field has been set.

### GetMinNotional

`func (o *NotionalFilter) GetMinNotional() string`

GetMinNotional returns the MinNotional field if non-nil, zero value otherwise.

### GetMinNotionalOk

`func (o *NotionalFilter) GetMinNotionalOk() (*string, bool)`

GetMinNotionalOk returns a tuple with the MinNotional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNotional

`func (o *NotionalFilter) SetMinNotional(v string)`

SetMinNotional sets MinNotional field to given value.

### HasMinNotional

`func (o *NotionalFilter) HasMinNotional() bool`

HasMinNotional returns a boolean if a field has been set.

### GetApplyMinToMarket

`func (o *NotionalFilter) GetApplyMinToMarket() bool`

GetApplyMinToMarket returns the ApplyMinToMarket field if non-nil, zero value otherwise.

### GetApplyMinToMarketOk

`func (o *NotionalFilter) GetApplyMinToMarketOk() (*bool, bool)`

GetApplyMinToMarketOk returns a tuple with the ApplyMinToMarket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyMinToMarket

`func (o *NotionalFilter) SetApplyMinToMarket(v bool)`

SetApplyMinToMarket sets ApplyMinToMarket field to given value.

### HasApplyMinToMarket

`func (o *NotionalFilter) HasApplyMinToMarket() bool`

HasApplyMinToMarket returns a boolean if a field has been set.

### GetMaxNotional

`func (o *NotionalFilter) GetMaxNotional() string`

GetMaxNotional returns the MaxNotional field if non-nil, zero value otherwise.

### GetMaxNotionalOk

`func (o *NotionalFilter) GetMaxNotionalOk() (*string, bool)`

GetMaxNotionalOk returns a tuple with the MaxNotional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNotional

`func (o *NotionalFilter) SetMaxNotional(v string)`

SetMaxNotional sets MaxNotional field to given value.

### HasMaxNotional

`func (o *NotionalFilter) HasMaxNotional() bool`

HasMaxNotional returns a boolean if a field has been set.

### GetApplyMaxToMarket

`func (o *NotionalFilter) GetApplyMaxToMarket() bool`

GetApplyMaxToMarket returns the ApplyMaxToMarket field if non-nil, zero value otherwise.

### GetApplyMaxToMarketOk

`func (o *NotionalFilter) GetApplyMaxToMarketOk() (*bool, bool)`

GetApplyMaxToMarketOk returns a tuple with the ApplyMaxToMarket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyMaxToMarket

`func (o *NotionalFilter) SetApplyMaxToMarket(v bool)`

SetApplyMaxToMarket sets ApplyMaxToMarket field to given value.

### HasApplyMaxToMarket

`func (o *NotionalFilter) HasApplyMaxToMarket() bool`

HasApplyMaxToMarket returns a boolean if a field has been set.

### GetAvgPriceMins

`func (o *NotionalFilter) GetAvgPriceMins() int32`

GetAvgPriceMins returns the AvgPriceMins field if non-nil, zero value otherwise.

### GetAvgPriceMinsOk

`func (o *NotionalFilter) GetAvgPriceMinsOk() (*int32, bool)`

GetAvgPriceMinsOk returns a tuple with the AvgPriceMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPriceMins

`func (o *NotionalFilter) SetAvgPriceMins(v int32)`

SetAvgPriceMins sets AvgPriceMins field to given value.

### HasAvgPriceMins

`func (o *NotionalFilter) HasAvgPriceMins() bool`

HasAvgPriceMins returns a boolean if a field has been set.


[[Back to README]](../README.md)


