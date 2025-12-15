# PercentPriceBySideFilter

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | Pointer to **string** |  | [optional] 
**MultiplierExponent** | Pointer to **int32** |  | [optional] 
**BidMultiplierUp** | Pointer to **string** |  | [optional] 
**BidMultiplierDown** | Pointer to **string** |  | [optional] 
**AskMultiplierUp** | Pointer to **string** |  | [optional] 
**AskMultiplierDown** | Pointer to **string** |  | [optional] 
**AvgPriceMins** | Pointer to **int32** |  | [optional] 

## Methods

### NewPercentPriceBySideFilter

`func NewPercentPriceBySideFilter() *PercentPriceBySideFilter`

NewPercentPriceBySideFilter instantiates a new PercentPriceBySideFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPercentPriceBySideFilterWithDefaults

`func NewPercentPriceBySideFilterWithDefaults() *PercentPriceBySideFilter`

NewPercentPriceBySideFilterWithDefaults instantiates a new PercentPriceBySideFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *PercentPriceBySideFilter) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *PercentPriceBySideFilter) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *PercentPriceBySideFilter) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.

### HasFilterType

`func (o *PercentPriceBySideFilter) HasFilterType() bool`

HasFilterType returns a boolean if a field has been set.

### GetMultiplierExponent

`func (o *PercentPriceBySideFilter) GetMultiplierExponent() int32`

GetMultiplierExponent returns the MultiplierExponent field if non-nil, zero value otherwise.

### GetMultiplierExponentOk

`func (o *PercentPriceBySideFilter) GetMultiplierExponentOk() (*int32, bool)`

GetMultiplierExponentOk returns a tuple with the MultiplierExponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplierExponent

`func (o *PercentPriceBySideFilter) SetMultiplierExponent(v int32)`

SetMultiplierExponent sets MultiplierExponent field to given value.

### HasMultiplierExponent

`func (o *PercentPriceBySideFilter) HasMultiplierExponent() bool`

HasMultiplierExponent returns a boolean if a field has been set.

### GetBidMultiplierUp

`func (o *PercentPriceBySideFilter) GetBidMultiplierUp() string`

GetBidMultiplierUp returns the BidMultiplierUp field if non-nil, zero value otherwise.

### GetBidMultiplierUpOk

`func (o *PercentPriceBySideFilter) GetBidMultiplierUpOk() (*string, bool)`

GetBidMultiplierUpOk returns a tuple with the BidMultiplierUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidMultiplierUp

`func (o *PercentPriceBySideFilter) SetBidMultiplierUp(v string)`

SetBidMultiplierUp sets BidMultiplierUp field to given value.

### HasBidMultiplierUp

`func (o *PercentPriceBySideFilter) HasBidMultiplierUp() bool`

HasBidMultiplierUp returns a boolean if a field has been set.

### GetBidMultiplierDown

`func (o *PercentPriceBySideFilter) GetBidMultiplierDown() string`

GetBidMultiplierDown returns the BidMultiplierDown field if non-nil, zero value otherwise.

### GetBidMultiplierDownOk

`func (o *PercentPriceBySideFilter) GetBidMultiplierDownOk() (*string, bool)`

GetBidMultiplierDownOk returns a tuple with the BidMultiplierDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidMultiplierDown

`func (o *PercentPriceBySideFilter) SetBidMultiplierDown(v string)`

SetBidMultiplierDown sets BidMultiplierDown field to given value.

### HasBidMultiplierDown

`func (o *PercentPriceBySideFilter) HasBidMultiplierDown() bool`

HasBidMultiplierDown returns a boolean if a field has been set.

### GetAskMultiplierUp

`func (o *PercentPriceBySideFilter) GetAskMultiplierUp() string`

GetAskMultiplierUp returns the AskMultiplierUp field if non-nil, zero value otherwise.

### GetAskMultiplierUpOk

`func (o *PercentPriceBySideFilter) GetAskMultiplierUpOk() (*string, bool)`

GetAskMultiplierUpOk returns a tuple with the AskMultiplierUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskMultiplierUp

`func (o *PercentPriceBySideFilter) SetAskMultiplierUp(v string)`

SetAskMultiplierUp sets AskMultiplierUp field to given value.

### HasAskMultiplierUp

`func (o *PercentPriceBySideFilter) HasAskMultiplierUp() bool`

HasAskMultiplierUp returns a boolean if a field has been set.

### GetAskMultiplierDown

`func (o *PercentPriceBySideFilter) GetAskMultiplierDown() string`

GetAskMultiplierDown returns the AskMultiplierDown field if non-nil, zero value otherwise.

### GetAskMultiplierDownOk

`func (o *PercentPriceBySideFilter) GetAskMultiplierDownOk() (*string, bool)`

GetAskMultiplierDownOk returns a tuple with the AskMultiplierDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskMultiplierDown

`func (o *PercentPriceBySideFilter) SetAskMultiplierDown(v string)`

SetAskMultiplierDown sets AskMultiplierDown field to given value.

### HasAskMultiplierDown

`func (o *PercentPriceBySideFilter) HasAskMultiplierDown() bool`

HasAskMultiplierDown returns a boolean if a field has been set.

### GetAvgPriceMins

`func (o *PercentPriceBySideFilter) GetAvgPriceMins() int32`

GetAvgPriceMins returns the AvgPriceMins field if non-nil, zero value otherwise.

### GetAvgPriceMinsOk

`func (o *PercentPriceBySideFilter) GetAvgPriceMinsOk() (*int32, bool)`

GetAvgPriceMinsOk returns a tuple with the AvgPriceMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPriceMins

`func (o *PercentPriceBySideFilter) SetAvgPriceMins(v int32)`

SetAvgPriceMins sets AvgPriceMins field to given value.

### HasAvgPriceMins

`func (o *PercentPriceBySideFilter) HasAvgPriceMins() bool`

HasAvgPriceMins returns a boolean if a field has been set.


[[Back to README]](../README.md)


