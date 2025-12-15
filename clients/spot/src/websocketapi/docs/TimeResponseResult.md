# TimeResponseResult

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**ServerTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewTimeResponseResult

`func NewTimeResponseResult() *TimeResponseResult`

NewTimeResponseResult instantiates a new TimeResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeResponseResultWithDefaults

`func NewTimeResponseResultWithDefaults() *TimeResponseResult`

NewTimeResponseResultWithDefaults instantiates a new TimeResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerTime

`func (o *TimeResponseResult) GetServerTime() int64`

GetServerTime returns the ServerTime field if non-nil, zero value otherwise.

### GetServerTimeOk

`func (o *TimeResponseResult) GetServerTimeOk() (*int64, bool)`

GetServerTimeOk returns a tuple with the ServerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTime

`func (o *TimeResponseResult) SetServerTime(v int64)`

SetServerTime sets ServerTime field to given value.

### HasServerTime

`func (o *TimeResponseResult) HasServerTime() bool`

HasServerTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


