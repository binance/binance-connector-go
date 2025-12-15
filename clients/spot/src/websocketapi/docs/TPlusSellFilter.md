# TPlusSellFilter

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | Pointer to **string** |  | [optional] 
**EndTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewTPlusSellFilter

`func NewTPlusSellFilter() *TPlusSellFilter`

NewTPlusSellFilter instantiates a new TPlusSellFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTPlusSellFilterWithDefaults

`func NewTPlusSellFilterWithDefaults() *TPlusSellFilter`

NewTPlusSellFilterWithDefaults instantiates a new TPlusSellFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *TPlusSellFilter) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *TPlusSellFilter) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *TPlusSellFilter) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.

### HasFilterType

`func (o *TPlusSellFilter) HasFilterType() bool`

HasFilterType returns a boolean if a field has been set.

### GetEndTime

`func (o *TPlusSellFilter) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *TPlusSellFilter) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *TPlusSellFilter) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *TPlusSellFilter) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


