# OpenInterestResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**OpenInterest** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 

## Methods

### NewOpenInterestResponse

`func NewOpenInterestResponse() *OpenInterestResponse`

NewOpenInterestResponse instantiates a new OpenInterestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenInterestResponseWithDefaults

`func NewOpenInterestResponseWithDefaults() *OpenInterestResponse`

NewOpenInterestResponseWithDefaults instantiates a new OpenInterestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpenInterest

`func (o *OpenInterestResponse) GetOpenInterest() string`

GetOpenInterest returns the OpenInterest field if non-nil, zero value otherwise.

### GetOpenInterestOk

`func (o *OpenInterestResponse) GetOpenInterestOk() (*string, bool)`

GetOpenInterestOk returns a tuple with the OpenInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenInterest

`func (o *OpenInterestResponse) SetOpenInterest(v string)`

SetOpenInterest sets OpenInterest field to given value.

### HasOpenInterest

`func (o *OpenInterestResponse) HasOpenInterest() bool`

HasOpenInterest returns a boolean if a field has been set.

### GetSymbol

`func (o *OpenInterestResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *OpenInterestResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *OpenInterestResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *OpenInterestResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTime

`func (o *OpenInterestResponse) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *OpenInterestResponse) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *OpenInterestResponse) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *OpenInterestResponse) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


