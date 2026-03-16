# ExecutionRulesResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**ExecutionRulesResponseResult**](ExecutionRulesResponseResult.md) |  | [optional] 

## Methods

### NewExecutionRulesResponse

`func NewExecutionRulesResponse() *ExecutionRulesResponse`

NewExecutionRulesResponse instantiates a new ExecutionRulesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionRulesResponseWithDefaults

`func NewExecutionRulesResponseWithDefaults() *ExecutionRulesResponse`

NewExecutionRulesResponseWithDefaults instantiates a new ExecutionRulesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExecutionRulesResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExecutionRulesResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExecutionRulesResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExecutionRulesResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *ExecutionRulesResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExecutionRulesResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExecutionRulesResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExecutionRulesResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *ExecutionRulesResponse) GetResult() ExecutionRulesResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ExecutionRulesResponse) GetResultOk() (*ExecutionRulesResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ExecutionRulesResponse) SetResult(v ExecutionRulesResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *ExecutionRulesResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to README]](../README.md)


