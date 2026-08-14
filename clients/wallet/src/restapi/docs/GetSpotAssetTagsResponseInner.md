# GetSpotAssetTagsResponseInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**AssetCode** | Pointer to **string** | Asset code. | [optional] 
**AssetName** | Pointer to **string** | Asset name. | [optional] 
**Trading** | Pointer to **bool** | Asset trading status. This endpoint only returns &#x60;true&#x60;. | [optional] 
**Tags** | Pointer to **[]string** | Tags configured for the asset. Returns an empty array when the asset has no tags. | [optional] 

## Methods

### NewGetSpotAssetTagsResponseInner

`func NewGetSpotAssetTagsResponseInner() *GetSpotAssetTagsResponseInner`

NewGetSpotAssetTagsResponseInner instantiates a new GetSpotAssetTagsResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSpotAssetTagsResponseInnerWithDefaults

`func NewGetSpotAssetTagsResponseInnerWithDefaults() *GetSpotAssetTagsResponseInner`

NewGetSpotAssetTagsResponseInnerWithDefaults instantiates a new GetSpotAssetTagsResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetCode

`func (o *GetSpotAssetTagsResponseInner) GetAssetCode() string`

GetAssetCode returns the AssetCode field if non-nil, zero value otherwise.

### GetAssetCodeOk

`func (o *GetSpotAssetTagsResponseInner) GetAssetCodeOk() (*string, bool)`

GetAssetCodeOk returns a tuple with the AssetCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetCode

`func (o *GetSpotAssetTagsResponseInner) SetAssetCode(v string)`

SetAssetCode sets AssetCode field to given value.

### HasAssetCode

`func (o *GetSpotAssetTagsResponseInner) HasAssetCode() bool`

HasAssetCode returns a boolean if a field has been set.

### GetAssetName

`func (o *GetSpotAssetTagsResponseInner) GetAssetName() string`

GetAssetName returns the AssetName field if non-nil, zero value otherwise.

### GetAssetNameOk

`func (o *GetSpotAssetTagsResponseInner) GetAssetNameOk() (*string, bool)`

GetAssetNameOk returns a tuple with the AssetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetName

`func (o *GetSpotAssetTagsResponseInner) SetAssetName(v string)`

SetAssetName sets AssetName field to given value.

### HasAssetName

`func (o *GetSpotAssetTagsResponseInner) HasAssetName() bool`

HasAssetName returns a boolean if a field has been set.

### GetTrading

`func (o *GetSpotAssetTagsResponseInner) GetTrading() bool`

GetTrading returns the Trading field if non-nil, zero value otherwise.

### GetTradingOk

`func (o *GetSpotAssetTagsResponseInner) GetTradingOk() (*bool, bool)`

GetTradingOk returns a tuple with the Trading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrading

`func (o *GetSpotAssetTagsResponseInner) SetTrading(v bool)`

SetTrading sets Trading field to given value.

### HasTrading

`func (o *GetSpotAssetTagsResponseInner) HasTrading() bool`

HasTrading returns a boolean if a field has been set.

### GetTags

`func (o *GetSpotAssetTagsResponseInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetSpotAssetTagsResponseInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetSpotAssetTagsResponseInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetSpotAssetTagsResponseInner) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to README]](../README.md)


