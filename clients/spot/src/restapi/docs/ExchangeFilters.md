# ExchangeFilters

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | Pointer to **string** |  | [optional] 
**MaxNumOrders** | Pointer to **int64** |  | [optional] 
**MaxNumAlgoOrders** | Pointer to **int64** |  | [optional] 
**MaxNumIcebergOrders** | Pointer to **int64** |  | [optional] 
**MaxNumOrderLists** | Pointer to **int64** |  | [optional] 

## Methods

### NewExchangeFilters

`func NewExchangeFilters() *ExchangeFilters`

NewExchangeFilters instantiates a new ExchangeFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeFiltersWithDefaults

`func NewExchangeFiltersWithDefaults() *ExchangeFilters`

NewExchangeFiltersWithDefaults instantiates a new ExchangeFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *ExchangeFilters) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *ExchangeFilters) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *ExchangeFilters) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.

### HasFilterType

`func (o *ExchangeFilters) HasFilterType() bool`

HasFilterType returns a boolean if a field has been set.

### GetMaxNumOrders

`func (o *ExchangeFilters) GetMaxNumOrders() int64`

GetMaxNumOrders returns the MaxNumOrders field if non-nil, zero value otherwise.

### GetMaxNumOrdersOk

`func (o *ExchangeFilters) GetMaxNumOrdersOk() (*int64, bool)`

GetMaxNumOrdersOk returns a tuple with the MaxNumOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumOrders

`func (o *ExchangeFilters) SetMaxNumOrders(v int64)`

SetMaxNumOrders sets MaxNumOrders field to given value.

### HasMaxNumOrders

`func (o *ExchangeFilters) HasMaxNumOrders() bool`

HasMaxNumOrders returns a boolean if a field has been set.

### GetMaxNumAlgoOrders

`func (o *ExchangeFilters) GetMaxNumAlgoOrders() int64`

GetMaxNumAlgoOrders returns the MaxNumAlgoOrders field if non-nil, zero value otherwise.

### GetMaxNumAlgoOrdersOk

`func (o *ExchangeFilters) GetMaxNumAlgoOrdersOk() (*int64, bool)`

GetMaxNumAlgoOrdersOk returns a tuple with the MaxNumAlgoOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumAlgoOrders

`func (o *ExchangeFilters) SetMaxNumAlgoOrders(v int64)`

SetMaxNumAlgoOrders sets MaxNumAlgoOrders field to given value.

### HasMaxNumAlgoOrders

`func (o *ExchangeFilters) HasMaxNumAlgoOrders() bool`

HasMaxNumAlgoOrders returns a boolean if a field has been set.

### GetMaxNumIcebergOrders

`func (o *ExchangeFilters) GetMaxNumIcebergOrders() int64`

GetMaxNumIcebergOrders returns the MaxNumIcebergOrders field if non-nil, zero value otherwise.

### GetMaxNumIcebergOrdersOk

`func (o *ExchangeFilters) GetMaxNumIcebergOrdersOk() (*int64, bool)`

GetMaxNumIcebergOrdersOk returns a tuple with the MaxNumIcebergOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumIcebergOrders

`func (o *ExchangeFilters) SetMaxNumIcebergOrders(v int64)`

SetMaxNumIcebergOrders sets MaxNumIcebergOrders field to given value.

### HasMaxNumIcebergOrders

`func (o *ExchangeFilters) HasMaxNumIcebergOrders() bool`

HasMaxNumIcebergOrders returns a boolean if a field has been set.

### GetMaxNumOrderLists

`func (o *ExchangeFilters) GetMaxNumOrderLists() int64`

GetMaxNumOrderLists returns the MaxNumOrderLists field if non-nil, zero value otherwise.

### GetMaxNumOrderListsOk

`func (o *ExchangeFilters) GetMaxNumOrderListsOk() (*int64, bool)`

GetMaxNumOrderListsOk returns a tuple with the MaxNumOrderLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumOrderLists

`func (o *ExchangeFilters) SetMaxNumOrderLists(v int64)`

SetMaxNumOrderLists sets MaxNumOrderLists field to given value.

### HasMaxNumOrderLists

`func (o *ExchangeFilters) HasMaxNumOrderLists() bool`

HasMaxNumOrderLists returns a boolean if a field has been set.


[[Back to README]](../README.md)


