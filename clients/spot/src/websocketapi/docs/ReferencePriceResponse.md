# ReferencePriceResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**ReferencePriceResponseResult**](ReferencePriceResponseResult.md) |  | [optional] 

## Methods

### NewReferencePriceResponse

`func NewReferencePriceResponse() *ReferencePriceResponse`

NewReferencePriceResponse instantiates a new ReferencePriceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferencePriceResponseWithDefaults

`func NewReferencePriceResponseWithDefaults() *ReferencePriceResponse`

NewReferencePriceResponseWithDefaults instantiates a new ReferencePriceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReferencePriceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReferencePriceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReferencePriceResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReferencePriceResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *ReferencePriceResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReferencePriceResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReferencePriceResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReferencePriceResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *ReferencePriceResponse) GetResult() ReferencePriceResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ReferencePriceResponse) GetResultOk() (*ReferencePriceResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ReferencePriceResponse) SetResult(v ReferencePriceResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *ReferencePriceResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to README]](../README.md)


