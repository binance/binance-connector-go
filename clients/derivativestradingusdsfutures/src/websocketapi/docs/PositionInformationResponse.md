# PositionInformationResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**[]PositionInformationResponseResultInner**](PositionInformationResponseResultInner.md) |  | [optional] 
**RateLimits** | Pointer to [**[]AccountInformationV2ResponseRateLimitsInner**](AccountInformationV2ResponseRateLimitsInner.md) |  | [optional] 

## Methods

### NewPositionInformationResponse

`func NewPositionInformationResponse() *PositionInformationResponse`

NewPositionInformationResponse instantiates a new PositionInformationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPositionInformationResponseWithDefaults

`func NewPositionInformationResponseWithDefaults() *PositionInformationResponse`

NewPositionInformationResponseWithDefaults instantiates a new PositionInformationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PositionInformationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PositionInformationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PositionInformationResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PositionInformationResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *PositionInformationResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PositionInformationResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PositionInformationResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PositionInformationResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *PositionInformationResponse) GetResult() []PositionInformationResponseResultInner`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *PositionInformationResponse) GetResultOk() (*[]PositionInformationResponseResultInner, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *PositionInformationResponse) SetResult(v []PositionInformationResponseResultInner)`

SetResult sets Result field to given value.

### HasResult

`func (o *PositionInformationResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRateLimits

`func (o *PositionInformationResponse) GetRateLimits() []AccountInformationV2ResponseRateLimitsInner`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *PositionInformationResponse) GetRateLimitsOk() (*[]AccountInformationV2ResponseRateLimitsInner, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *PositionInformationResponse) SetRateLimits(v []AccountInformationV2ResponseRateLimitsInner)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *PositionInformationResponse) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


