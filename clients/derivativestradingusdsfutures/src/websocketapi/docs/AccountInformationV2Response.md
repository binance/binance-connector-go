# AccountInformationV2Response

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**AccountInformationV2ResponseResult**](AccountInformationV2ResponseResult.md) |  | [optional] 
**RateLimits** | Pointer to [**[]AccountInformationV2ResponseRateLimitsInner**](AccountInformationV2ResponseRateLimitsInner.md) |  | [optional] 

## Methods

### NewAccountInformationV2Response

`func NewAccountInformationV2Response() *AccountInformationV2Response`

NewAccountInformationV2Response instantiates a new AccountInformationV2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountInformationV2ResponseWithDefaults

`func NewAccountInformationV2ResponseWithDefaults() *AccountInformationV2Response`

NewAccountInformationV2ResponseWithDefaults instantiates a new AccountInformationV2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountInformationV2Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountInformationV2Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountInformationV2Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountInformationV2Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *AccountInformationV2Response) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountInformationV2Response) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountInformationV2Response) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccountInformationV2Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *AccountInformationV2Response) GetResult() AccountInformationV2ResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *AccountInformationV2Response) GetResultOk() (*AccountInformationV2ResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *AccountInformationV2Response) SetResult(v AccountInformationV2ResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *AccountInformationV2Response) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRateLimits

`func (o *AccountInformationV2Response) GetRateLimits() []AccountInformationV2ResponseRateLimitsInner`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *AccountInformationV2Response) GetRateLimitsOk() (*[]AccountInformationV2ResponseRateLimitsInner, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *AccountInformationV2Response) SetRateLimits(v []AccountInformationV2ResponseRateLimitsInner)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *AccountInformationV2Response) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


