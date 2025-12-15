# GetNFTAssetResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int64** |  | [optional] 
**List** | Pointer to [**[]GetNFTAssetResponseListInner**](GetNFTAssetResponseListInner.md) |  | [optional] 

## Methods

### NewGetNFTAssetResponse

`func NewGetNFTAssetResponse() *GetNFTAssetResponse`

NewGetNFTAssetResponse instantiates a new GetNFTAssetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNFTAssetResponseWithDefaults

`func NewGetNFTAssetResponseWithDefaults() *GetNFTAssetResponse`

NewGetNFTAssetResponseWithDefaults instantiates a new GetNFTAssetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *GetNFTAssetResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetNFTAssetResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetNFTAssetResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetNFTAssetResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetList

`func (o *GetNFTAssetResponse) GetList() []GetNFTAssetResponseListInner`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *GetNFTAssetResponse) GetListOk() (*[]GetNFTAssetResponseListInner, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *GetNFTAssetResponse) SetList(v []GetNFTAssetResponseListInner)`

SetList sets List field to given value.

### HasList

`func (o *GetNFTAssetResponse) HasList() bool`

HasList returns a boolean if a field has been set.


[[Back to README]](../README.md)


