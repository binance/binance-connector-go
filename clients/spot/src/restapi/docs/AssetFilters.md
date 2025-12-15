# AssetFilters

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | Pointer to **string** |  | [optional] 
**QtyExponent** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **string** |  | [optional] 
**Asset** | Pointer to **string** |  | [optional] 

## Methods

### NewAssetFilters

`func NewAssetFilters() *AssetFilters`

NewAssetFilters instantiates a new AssetFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetFiltersWithDefaults

`func NewAssetFiltersWithDefaults() *AssetFilters`

NewAssetFiltersWithDefaults instantiates a new AssetFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *AssetFilters) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *AssetFilters) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *AssetFilters) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.

### HasFilterType

`func (o *AssetFilters) HasFilterType() bool`

HasFilterType returns a boolean if a field has been set.

### GetQtyExponent

`func (o *AssetFilters) GetQtyExponent() int32`

GetQtyExponent returns the QtyExponent field if non-nil, zero value otherwise.

### GetQtyExponentOk

`func (o *AssetFilters) GetQtyExponentOk() (*int32, bool)`

GetQtyExponentOk returns a tuple with the QtyExponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQtyExponent

`func (o *AssetFilters) SetQtyExponent(v int32)`

SetQtyExponent sets QtyExponent field to given value.

### HasQtyExponent

`func (o *AssetFilters) HasQtyExponent() bool`

HasQtyExponent returns a boolean if a field has been set.

### GetLimit

`func (o *AssetFilters) GetLimit() string`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AssetFilters) GetLimitOk() (*string, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AssetFilters) SetLimit(v string)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *AssetFilters) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetAsset

`func (o *AssetFilters) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *AssetFilters) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *AssetFilters) SetAsset(v string)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *AssetFilters) HasAsset() bool`

HasAsset returns a boolean if a field has been set.


[[Back to README]](../README.md)


