# FuturesAccountBalanceResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**[]FuturesAccountBalanceResponseResultInner**](FuturesAccountBalanceResponseResultInner.md) |  | [optional] 
**RateLimits** | Pointer to [**[]AccountInformationResponseRateLimitsInner**](AccountInformationResponseRateLimitsInner.md) |  | [optional] 

## Methods

### NewFuturesAccountBalanceResponse

`func NewFuturesAccountBalanceResponse() *FuturesAccountBalanceResponse`

NewFuturesAccountBalanceResponse instantiates a new FuturesAccountBalanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFuturesAccountBalanceResponseWithDefaults

`func NewFuturesAccountBalanceResponseWithDefaults() *FuturesAccountBalanceResponse`

NewFuturesAccountBalanceResponseWithDefaults instantiates a new FuturesAccountBalanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FuturesAccountBalanceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FuturesAccountBalanceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FuturesAccountBalanceResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FuturesAccountBalanceResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *FuturesAccountBalanceResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FuturesAccountBalanceResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FuturesAccountBalanceResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FuturesAccountBalanceResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *FuturesAccountBalanceResponse) GetResult() []FuturesAccountBalanceResponseResultInner`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *FuturesAccountBalanceResponse) GetResultOk() (*[]FuturesAccountBalanceResponseResultInner, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *FuturesAccountBalanceResponse) SetResult(v []FuturesAccountBalanceResponseResultInner)`

SetResult sets Result field to given value.

### HasResult

`func (o *FuturesAccountBalanceResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRateLimits

`func (o *FuturesAccountBalanceResponse) GetRateLimits() []AccountInformationResponseRateLimitsInner`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *FuturesAccountBalanceResponse) GetRateLimitsOk() (*[]AccountInformationResponseRateLimitsInner, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *FuturesAccountBalanceResponse) SetRateLimits(v []AccountInformationResponseRateLimitsInner)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *FuturesAccountBalanceResponse) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


