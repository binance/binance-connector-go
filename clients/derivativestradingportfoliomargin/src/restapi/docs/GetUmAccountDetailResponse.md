# GetUmAccountDetailResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Assets** | Pointer to [**[]GetUmAccountDetailV2ResponseAssetsInner**](GetUmAccountDetailV2ResponseAssetsInner.md) |  | [optional] 
**Positions** | Pointer to [**[]GetUmAccountDetailResponsePositionsInner**](GetUmAccountDetailResponsePositionsInner.md) |  | [optional] 

## Methods

### NewGetUmAccountDetailResponse

`func NewGetUmAccountDetailResponse() *GetUmAccountDetailResponse`

NewGetUmAccountDetailResponse instantiates a new GetUmAccountDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUmAccountDetailResponseWithDefaults

`func NewGetUmAccountDetailResponseWithDefaults() *GetUmAccountDetailResponse`

NewGetUmAccountDetailResponseWithDefaults instantiates a new GetUmAccountDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *GetUmAccountDetailResponse) GetAssets() []GetUmAccountDetailV2ResponseAssetsInner`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *GetUmAccountDetailResponse) GetAssetsOk() (*[]GetUmAccountDetailV2ResponseAssetsInner, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *GetUmAccountDetailResponse) SetAssets(v []GetUmAccountDetailV2ResponseAssetsInner)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *GetUmAccountDetailResponse) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetPositions

`func (o *GetUmAccountDetailResponse) GetPositions() []GetUmAccountDetailResponsePositionsInner`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *GetUmAccountDetailResponse) GetPositionsOk() (*[]GetUmAccountDetailResponsePositionsInner, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *GetUmAccountDetailResponse) SetPositions(v []GetUmAccountDetailResponsePositionsInner)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *GetUmAccountDetailResponse) HasPositions() bool`

HasPositions returns a boolean if a field has been set.


[[Back to README]](../README.md)


