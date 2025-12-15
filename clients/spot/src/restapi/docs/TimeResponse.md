# TimeResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**ServerTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewTimeResponse

`func NewTimeResponse() *TimeResponse`

NewTimeResponse instantiates a new TimeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeResponseWithDefaults

`func NewTimeResponseWithDefaults() *TimeResponse`

NewTimeResponseWithDefaults instantiates a new TimeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerTime

`func (o *TimeResponse) GetServerTime() int64`

GetServerTime returns the ServerTime field if non-nil, zero value otherwise.

### GetServerTimeOk

`func (o *TimeResponse) GetServerTimeOk() (*int64, bool)`

GetServerTimeOk returns a tuple with the ServerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTime

`func (o *TimeResponse) SetServerTime(v int64)`

SetServerTime sets ServerTime field to given value.

### HasServerTime

`func (o *TimeResponse) HasServerTime() bool`

HasServerTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


