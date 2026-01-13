# PositionInformationV2Response

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**[]PositionInformationV2ResponseResultInner**](PositionInformationV2ResponseResultInner.md) |  | [optional] 
**RateLimits** | Pointer to [**[]AccountInformationV2ResponseRateLimitsInner**](AccountInformationV2ResponseRateLimitsInner.md) |  | [optional] 

## Methods

### NewPositionInformationV2Response

`func NewPositionInformationV2Response() *PositionInformationV2Response`

NewPositionInformationV2Response instantiates a new PositionInformationV2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPositionInformationV2ResponseWithDefaults

`func NewPositionInformationV2ResponseWithDefaults() *PositionInformationV2Response`

NewPositionInformationV2ResponseWithDefaults instantiates a new PositionInformationV2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PositionInformationV2Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PositionInformationV2Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PositionInformationV2Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PositionInformationV2Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *PositionInformationV2Response) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PositionInformationV2Response) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PositionInformationV2Response) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PositionInformationV2Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *PositionInformationV2Response) GetResult() []PositionInformationV2ResponseResultInner`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *PositionInformationV2Response) GetResultOk() (*[]PositionInformationV2ResponseResultInner, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *PositionInformationV2Response) SetResult(v []PositionInformationV2ResponseResultInner)`

SetResult sets Result field to given value.

### HasResult

`func (o *PositionInformationV2Response) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRateLimits

`func (o *PositionInformationV2Response) GetRateLimits() []AccountInformationV2ResponseRateLimitsInner`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *PositionInformationV2Response) GetRateLimitsOk() (*[]AccountInformationV2ResponseRateLimitsInner, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *PositionInformationV2Response) SetRateLimits(v []AccountInformationV2ResponseRateLimitsInner)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *PositionInformationV2Response) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


