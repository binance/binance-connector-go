# GetCmAccountDetailResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Assets** | Pointer to [**[]GetCmAccountDetailResponseAssetsInner**](GetCmAccountDetailResponseAssetsInner.md) |  | [optional] 
**Positions** | Pointer to [**[]GetCmAccountDetailResponsePositionsInner**](GetCmAccountDetailResponsePositionsInner.md) |  | [optional] 

## Methods

### NewGetCmAccountDetailResponse

`func NewGetCmAccountDetailResponse() *GetCmAccountDetailResponse`

NewGetCmAccountDetailResponse instantiates a new GetCmAccountDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCmAccountDetailResponseWithDefaults

`func NewGetCmAccountDetailResponseWithDefaults() *GetCmAccountDetailResponse`

NewGetCmAccountDetailResponseWithDefaults instantiates a new GetCmAccountDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *GetCmAccountDetailResponse) GetAssets() []GetCmAccountDetailResponseAssetsInner`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *GetCmAccountDetailResponse) GetAssetsOk() (*[]GetCmAccountDetailResponseAssetsInner, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *GetCmAccountDetailResponse) SetAssets(v []GetCmAccountDetailResponseAssetsInner)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *GetCmAccountDetailResponse) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetPositions

`func (o *GetCmAccountDetailResponse) GetPositions() []GetCmAccountDetailResponsePositionsInner`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *GetCmAccountDetailResponse) GetPositionsOk() (*[]GetCmAccountDetailResponsePositionsInner, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *GetCmAccountDetailResponse) SetPositions(v []GetCmAccountDetailResponsePositionsInner)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *GetCmAccountDetailResponse) HasPositions() bool`

HasPositions returns a boolean if a field has been set.


[[Back to README]](../README.md)


