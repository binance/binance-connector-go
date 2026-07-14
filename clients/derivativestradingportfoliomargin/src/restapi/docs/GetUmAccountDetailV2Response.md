# GetUmAccountDetailV2Response

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Assets** | Pointer to [**[]GetUmAccountDetailResponseAssetsInner**](GetUmAccountDetailResponseAssetsInner.md) | Assets. | [optional] 
**Positions** | Pointer to [**[]GetUmAccountDetailV2ResponsePositionsInner**](GetUmAccountDetailV2ResponsePositionsInner.md) | positions of all symbols in the market are returned | [optional] 

## Methods

### NewGetUmAccountDetailV2Response

`func NewGetUmAccountDetailV2Response() *GetUmAccountDetailV2Response`

NewGetUmAccountDetailV2Response instantiates a new GetUmAccountDetailV2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUmAccountDetailV2ResponseWithDefaults

`func NewGetUmAccountDetailV2ResponseWithDefaults() *GetUmAccountDetailV2Response`

NewGetUmAccountDetailV2ResponseWithDefaults instantiates a new GetUmAccountDetailV2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *GetUmAccountDetailV2Response) GetAssets() []GetUmAccountDetailResponseAssetsInner`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *GetUmAccountDetailV2Response) GetAssetsOk() (*[]GetUmAccountDetailResponseAssetsInner, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *GetUmAccountDetailV2Response) SetAssets(v []GetUmAccountDetailResponseAssetsInner)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *GetUmAccountDetailV2Response) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetPositions

`func (o *GetUmAccountDetailV2Response) GetPositions() []GetUmAccountDetailV2ResponsePositionsInner`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *GetUmAccountDetailV2Response) GetPositionsOk() (*[]GetUmAccountDetailV2ResponsePositionsInner, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *GetUmAccountDetailV2Response) SetPositions(v []GetUmAccountDetailV2ResponsePositionsInner)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *GetUmAccountDetailV2Response) HasPositions() bool`

HasPositions returns a boolean if a field has been set.


[[Back to README]](../README.md)


