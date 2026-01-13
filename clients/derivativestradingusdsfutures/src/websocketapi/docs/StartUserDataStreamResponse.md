# StartUserDataStreamResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**StartUserDataStreamResponseResult**](StartUserDataStreamResponseResult.md) |  | [optional] 
**RateLimits** | Pointer to [**[]SymbolOrderBookTickerResponse1RateLimitsInner**](SymbolOrderBookTickerResponse1RateLimitsInner.md) |  | [optional] 

## Methods

### NewStartUserDataStreamResponse

`func NewStartUserDataStreamResponse() *StartUserDataStreamResponse`

NewStartUserDataStreamResponse instantiates a new StartUserDataStreamResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartUserDataStreamResponseWithDefaults

`func NewStartUserDataStreamResponseWithDefaults() *StartUserDataStreamResponse`

NewStartUserDataStreamResponseWithDefaults instantiates a new StartUserDataStreamResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StartUserDataStreamResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StartUserDataStreamResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StartUserDataStreamResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StartUserDataStreamResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *StartUserDataStreamResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StartUserDataStreamResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StartUserDataStreamResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StartUserDataStreamResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *StartUserDataStreamResponse) GetResult() StartUserDataStreamResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *StartUserDataStreamResponse) GetResultOk() (*StartUserDataStreamResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *StartUserDataStreamResponse) SetResult(v StartUserDataStreamResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *StartUserDataStreamResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRateLimits

`func (o *StartUserDataStreamResponse) GetRateLimits() []SymbolOrderBookTickerResponse1RateLimitsInner`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *StartUserDataStreamResponse) GetRateLimitsOk() (*[]SymbolOrderBookTickerResponse1RateLimitsInner, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *StartUserDataStreamResponse) SetRateLimits(v []SymbolOrderBookTickerResponse1RateLimitsInner)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *StartUserDataStreamResponse) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


