# QueryOrderResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**QueryOrderResponseResult**](QueryOrderResponseResult.md) |  | [optional] 

## Methods

### NewQueryOrderResponse

`func NewQueryOrderResponse() *QueryOrderResponse`

NewQueryOrderResponse instantiates a new QueryOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryOrderResponseWithDefaults

`func NewQueryOrderResponseWithDefaults() *QueryOrderResponse`

NewQueryOrderResponseWithDefaults instantiates a new QueryOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QueryOrderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QueryOrderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QueryOrderResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *QueryOrderResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *QueryOrderResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QueryOrderResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QueryOrderResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *QueryOrderResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *QueryOrderResponse) GetResult() QueryOrderResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *QueryOrderResponse) GetResultOk() (*QueryOrderResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *QueryOrderResponse) SetResult(v QueryOrderResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *QueryOrderResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to README]](../README.md)


