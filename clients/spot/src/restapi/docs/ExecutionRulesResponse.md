# ExecutionRulesResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**SymbolRules** | Pointer to [**[]ExecutionRulesResponseSymbolRulesInner**](ExecutionRulesResponseSymbolRulesInner.md) |  | [optional] 

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

### GetSymbolRules

`func (o *ExecutionRulesResponse) GetSymbolRules() []ExecutionRulesResponseSymbolRulesInner`

GetSymbolRules returns the SymbolRules field if non-nil, zero value otherwise.

### GetSymbolRulesOk

`func (o *ExecutionRulesResponse) GetSymbolRulesOk() (*[]ExecutionRulesResponseSymbolRulesInner, bool)`

GetSymbolRulesOk returns a tuple with the SymbolRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbolRules

`func (o *ExecutionRulesResponse) SetSymbolRules(v []ExecutionRulesResponseSymbolRulesInner)`

SetSymbolRules sets SymbolRules field to given value.

### HasSymbolRules

`func (o *ExecutionRulesResponse) HasSymbolRules() bool`

HasSymbolRules returns a boolean if a field has been set.


[[Back to README]](../README.md)


