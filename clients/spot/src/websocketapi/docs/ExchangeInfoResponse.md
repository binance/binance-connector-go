# ExchangeInfoResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**ExchangeInfoResponseResult**](ExchangeInfoResponseResult.md) |  | [optional] 
**RateLimits** | Pointer to [**[]RateLimits**](RateLimits.md) |  | [optional] 

## Methods

### NewExchangeInfoResponse

`func NewExchangeInfoResponse() *ExchangeInfoResponse`

NewExchangeInfoResponse instantiates a new ExchangeInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeInfoResponseWithDefaults

`func NewExchangeInfoResponseWithDefaults() *ExchangeInfoResponse`

NewExchangeInfoResponseWithDefaults instantiates a new ExchangeInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExchangeInfoResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExchangeInfoResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExchangeInfoResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExchangeInfoResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *ExchangeInfoResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExchangeInfoResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExchangeInfoResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExchangeInfoResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *ExchangeInfoResponse) GetResult() ExchangeInfoResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ExchangeInfoResponse) GetResultOk() (*ExchangeInfoResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ExchangeInfoResponse) SetResult(v ExchangeInfoResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *ExchangeInfoResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRateLimits

`func (o *ExchangeInfoResponse) GetRateLimits() []RateLimits`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *ExchangeInfoResponse) GetRateLimitsOk() (*[]RateLimits, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *ExchangeInfoResponse) SetRateLimits(v []RateLimits)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *ExchangeInfoResponse) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


