# IndexPriceResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **int64** |  | [optional] 
**IndexPrice** | Pointer to **string** |  | [optional] 

## Methods

### NewIndexPriceResponse

`func NewIndexPriceResponse() *IndexPriceResponse`

NewIndexPriceResponse instantiates a new IndexPriceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexPriceResponseWithDefaults

`func NewIndexPriceResponseWithDefaults() *IndexPriceResponse`

NewIndexPriceResponseWithDefaults instantiates a new IndexPriceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *IndexPriceResponse) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *IndexPriceResponse) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *IndexPriceResponse) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *IndexPriceResponse) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetIndexPrice

`func (o *IndexPriceResponse) GetIndexPrice() string`

GetIndexPrice returns the IndexPrice field if non-nil, zero value otherwise.

### GetIndexPriceOk

`func (o *IndexPriceResponse) GetIndexPriceOk() (*string, bool)`

GetIndexPriceOk returns a tuple with the IndexPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexPrice

`func (o *IndexPriceResponse) SetIndexPrice(v string)`

SetIndexPrice sets IndexPrice field to given value.

### HasIndexPrice

`func (o *IndexPriceResponse) HasIndexPrice() bool`

HasIndexPrice returns a boolean if a field has been set.


[[Back to README]](../README.md)


