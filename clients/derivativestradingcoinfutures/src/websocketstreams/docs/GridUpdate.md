# GridUpdate

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**T** | Pointer to **int64** |  | [optional] 
**E** | Pointer to **int64** |  | [optional] 
**Gu** | Pointer to [**GridUpdateGu**](GridUpdateGu.md) |  | [optional] 

## Methods

### NewGridUpdate

`func NewGridUpdate() *GridUpdate`

NewGridUpdate instantiates a new GridUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridUpdateWithDefaults

`func NewGridUpdateWithDefaults() *GridUpdate`

NewGridUpdateWithDefaults instantiates a new GridUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetT

`func (o *GridUpdate) GetT() int64`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *GridUpdate) GetTOk() (*int64, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *GridUpdate) SetT(v int64)`

SetT sets T field to given value.

### HasT

`func (o *GridUpdate) HasT() bool`

HasT returns a boolean if a field has been set.

### GetE

`func (o *GridUpdate) GetE() int64`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *GridUpdate) GetEOk() (*int64, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *GridUpdate) SetE(v int64)`

SetE sets E field to given value.

### HasE

`func (o *GridUpdate) HasE() bool`

HasE returns a boolean if a field has been set.

### GetGu

`func (o *GridUpdate) GetGu() GridUpdateGu`

GetGu returns the Gu field if non-nil, zero value otherwise.

### GetGuOk

`func (o *GridUpdate) GetGuOk() (*GridUpdateGu, bool)`

GetGuOk returns a tuple with the Gu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGu

`func (o *GridUpdate) SetGu(v GridUpdateGu)`

SetGu sets Gu field to given value.

### HasGu

`func (o *GridUpdate) HasGu() bool`

HasGu returns a boolean if a field has been set.


[[Back to README]](../README.md)


