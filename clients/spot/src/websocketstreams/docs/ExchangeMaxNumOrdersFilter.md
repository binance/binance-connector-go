# ExchangeMaxNumOrdersFilter

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | Pointer to **string** |  | [optional] 
**MaxNumOrders** | Pointer to **int64** |  | [optional] 

## Methods

### NewExchangeMaxNumOrdersFilter

`func NewExchangeMaxNumOrdersFilter() *ExchangeMaxNumOrdersFilter`

NewExchangeMaxNumOrdersFilter instantiates a new ExchangeMaxNumOrdersFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeMaxNumOrdersFilterWithDefaults

`func NewExchangeMaxNumOrdersFilterWithDefaults() *ExchangeMaxNumOrdersFilter`

NewExchangeMaxNumOrdersFilterWithDefaults instantiates a new ExchangeMaxNumOrdersFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *ExchangeMaxNumOrdersFilter) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *ExchangeMaxNumOrdersFilter) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *ExchangeMaxNumOrdersFilter) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.

### HasFilterType

`func (o *ExchangeMaxNumOrdersFilter) HasFilterType() bool`

HasFilterType returns a boolean if a field has been set.

### GetMaxNumOrders

`func (o *ExchangeMaxNumOrdersFilter) GetMaxNumOrders() int64`

GetMaxNumOrders returns the MaxNumOrders field if non-nil, zero value otherwise.

### GetMaxNumOrdersOk

`func (o *ExchangeMaxNumOrdersFilter) GetMaxNumOrdersOk() (*int64, bool)`

GetMaxNumOrdersOk returns a tuple with the MaxNumOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumOrders

`func (o *ExchangeMaxNumOrdersFilter) SetMaxNumOrders(v int64)`

SetMaxNumOrders sets MaxNumOrders field to given value.

### HasMaxNumOrders

`func (o *ExchangeMaxNumOrdersFilter) HasMaxNumOrders() bool`

HasMaxNumOrders returns a boolean if a field has been set.


[[Back to README]](../README.md)


