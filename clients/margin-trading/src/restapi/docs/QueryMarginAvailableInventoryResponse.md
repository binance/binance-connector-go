# QueryMarginAvailableInventoryResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Assets** | Pointer to [**QueryMarginAvailableInventoryResponseAssets**](QueryMarginAvailableInventoryResponseAssets.md) |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewQueryMarginAvailableInventoryResponse

`func NewQueryMarginAvailableInventoryResponse() *QueryMarginAvailableInventoryResponse`

NewQueryMarginAvailableInventoryResponse instantiates a new QueryMarginAvailableInventoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryMarginAvailableInventoryResponseWithDefaults

`func NewQueryMarginAvailableInventoryResponseWithDefaults() *QueryMarginAvailableInventoryResponse`

NewQueryMarginAvailableInventoryResponseWithDefaults instantiates a new QueryMarginAvailableInventoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *QueryMarginAvailableInventoryResponse) GetAssets() QueryMarginAvailableInventoryResponseAssets`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *QueryMarginAvailableInventoryResponse) GetAssetsOk() (*QueryMarginAvailableInventoryResponseAssets, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *QueryMarginAvailableInventoryResponse) SetAssets(v QueryMarginAvailableInventoryResponseAssets)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *QueryMarginAvailableInventoryResponse) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetUpdateTime

`func (o *QueryMarginAvailableInventoryResponse) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *QueryMarginAvailableInventoryResponse) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *QueryMarginAvailableInventoryResponse) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *QueryMarginAvailableInventoryResponse) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


