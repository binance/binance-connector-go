# AccountInformationResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**AccountInformationResponseResult**](AccountInformationResponseResult.md) |  | [optional] 
**RateLimits** | Pointer to [**[]AccountInformationV2ResponseRateLimitsInner**](AccountInformationV2ResponseRateLimitsInner.md) |  | [optional] 

## Methods

### NewAccountInformationResponse

`func NewAccountInformationResponse() *AccountInformationResponse`

NewAccountInformationResponse instantiates a new AccountInformationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountInformationResponseWithDefaults

`func NewAccountInformationResponseWithDefaults() *AccountInformationResponse`

NewAccountInformationResponseWithDefaults instantiates a new AccountInformationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountInformationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountInformationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountInformationResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountInformationResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *AccountInformationResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountInformationResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountInformationResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccountInformationResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *AccountInformationResponse) GetResult() AccountInformationResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *AccountInformationResponse) GetResultOk() (*AccountInformationResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *AccountInformationResponse) SetResult(v AccountInformationResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *AccountInformationResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRateLimits

`func (o *AccountInformationResponse) GetRateLimits() []AccountInformationV2ResponseRateLimitsInner`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *AccountInformationResponse) GetRateLimitsOk() (*[]AccountInformationV2ResponseRateLimitsInner, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *AccountInformationResponse) SetRateLimits(v []AccountInformationV2ResponseRateLimitsInner)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *AccountInformationResponse) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


