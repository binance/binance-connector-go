# DustlogResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int64** |  | [optional] 
**UserAssetDribblets** | Pointer to [**[]DustlogResponseUserAssetDribbletsInner**](DustlogResponseUserAssetDribbletsInner.md) |  | [optional] 

## Methods

### NewDustlogResponse

`func NewDustlogResponse() *DustlogResponse`

NewDustlogResponse instantiates a new DustlogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDustlogResponseWithDefaults

`func NewDustlogResponseWithDefaults() *DustlogResponse`

NewDustlogResponseWithDefaults instantiates a new DustlogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *DustlogResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *DustlogResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *DustlogResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *DustlogResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUserAssetDribblets

`func (o *DustlogResponse) GetUserAssetDribblets() []DustlogResponseUserAssetDribbletsInner`

GetUserAssetDribblets returns the UserAssetDribblets field if non-nil, zero value otherwise.

### GetUserAssetDribbletsOk

`func (o *DustlogResponse) GetUserAssetDribbletsOk() (*[]DustlogResponseUserAssetDribbletsInner, bool)`

GetUserAssetDribbletsOk returns a tuple with the UserAssetDribblets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAssetDribblets

`func (o *DustlogResponse) SetUserAssetDribblets(v []DustlogResponseUserAssetDribbletsInner)`

SetUserAssetDribblets sets UserAssetDribblets field to given value.

### HasUserAssetDribblets

`func (o *DustlogResponse) HasUserAssetDribblets() bool`

HasUserAssetDribblets returns a boolean if a field has been set.


[[Back to README]](../README.md)


