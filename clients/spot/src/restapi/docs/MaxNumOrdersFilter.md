# MaxNumOrdersFilter

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | Pointer to **string** |  | [optional] 
**MaxNumOrders** | Pointer to **int64** |  | [optional] 

## Methods

### NewMaxNumOrdersFilter

`func NewMaxNumOrdersFilter() *MaxNumOrdersFilter`

NewMaxNumOrdersFilter instantiates a new MaxNumOrdersFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaxNumOrdersFilterWithDefaults

`func NewMaxNumOrdersFilterWithDefaults() *MaxNumOrdersFilter`

NewMaxNumOrdersFilterWithDefaults instantiates a new MaxNumOrdersFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *MaxNumOrdersFilter) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *MaxNumOrdersFilter) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *MaxNumOrdersFilter) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.

### HasFilterType

`func (o *MaxNumOrdersFilter) HasFilterType() bool`

HasFilterType returns a boolean if a field has been set.

### GetMaxNumOrders

`func (o *MaxNumOrdersFilter) GetMaxNumOrders() int64`

GetMaxNumOrders returns the MaxNumOrders field if non-nil, zero value otherwise.

### GetMaxNumOrdersOk

`func (o *MaxNumOrdersFilter) GetMaxNumOrdersOk() (*int64, bool)`

GetMaxNumOrdersOk returns a tuple with the MaxNumOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumOrders

`func (o *MaxNumOrdersFilter) SetMaxNumOrders(v int64)`

SetMaxNumOrders sets MaxNumOrders field to given value.

### HasMaxNumOrders

`func (o *MaxNumOrdersFilter) HasMaxNumOrders() bool`

HasMaxNumOrders returns a boolean if a field has been set.


[[Back to README]](../README.md)


