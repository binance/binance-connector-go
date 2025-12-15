# TickerPriceResponse1

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**TickerPriceResponse1Result**](TickerPriceResponse1Result.md) |  | [optional] 
**RateLimits** | Pointer to [**[]RateLimits**](RateLimits.md) |  | [optional] 

## Methods

### NewTickerPriceResponse1

`func NewTickerPriceResponse1() *TickerPriceResponse1`

NewTickerPriceResponse1 instantiates a new TickerPriceResponse1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTickerPriceResponse1WithDefaults

`func NewTickerPriceResponse1WithDefaults() *TickerPriceResponse1`

NewTickerPriceResponse1WithDefaults instantiates a new TickerPriceResponse1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TickerPriceResponse1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TickerPriceResponse1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TickerPriceResponse1) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TickerPriceResponse1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *TickerPriceResponse1) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TickerPriceResponse1) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TickerPriceResponse1) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TickerPriceResponse1) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *TickerPriceResponse1) GetResult() TickerPriceResponse1Result`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TickerPriceResponse1) GetResultOk() (*TickerPriceResponse1Result, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TickerPriceResponse1) SetResult(v TickerPriceResponse1Result)`

SetResult sets Result field to given value.

### HasResult

`func (o *TickerPriceResponse1) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRateLimits

`func (o *TickerPriceResponse1) GetRateLimits() []RateLimits`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *TickerPriceResponse1) GetRateLimitsOk() (*[]RateLimits, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *TickerPriceResponse1) SetRateLimits(v []RateLimits)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *TickerPriceResponse1) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


