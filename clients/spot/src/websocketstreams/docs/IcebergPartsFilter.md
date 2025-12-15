# IcebergPartsFilter

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 

## Methods

### NewIcebergPartsFilter

`func NewIcebergPartsFilter() *IcebergPartsFilter`

NewIcebergPartsFilter instantiates a new IcebergPartsFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIcebergPartsFilterWithDefaults

`func NewIcebergPartsFilterWithDefaults() *IcebergPartsFilter`

NewIcebergPartsFilterWithDefaults instantiates a new IcebergPartsFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *IcebergPartsFilter) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *IcebergPartsFilter) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *IcebergPartsFilter) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.

### HasFilterType

`func (o *IcebergPartsFilter) HasFilterType() bool`

HasFilterType returns a boolean if a field has been set.

### GetLimit

`func (o *IcebergPartsFilter) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *IcebergPartsFilter) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *IcebergPartsFilter) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *IcebergPartsFilter) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to README]](../README.md)


