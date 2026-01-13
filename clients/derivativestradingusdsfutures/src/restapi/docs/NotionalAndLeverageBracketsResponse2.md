# NotionalAndLeverageBracketsResponse2

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**NotionalCoef** | Pointer to **float32** |  | [optional] 
**Brackets** | Pointer to [**[]NotionalAndLeverageBracketsResponse2BracketsInner**](NotionalAndLeverageBracketsResponse2BracketsInner.md) |  | [optional] 

## Methods

### NewNotionalAndLeverageBracketsResponse2

`func NewNotionalAndLeverageBracketsResponse2() *NotionalAndLeverageBracketsResponse2`

NewNotionalAndLeverageBracketsResponse2 instantiates a new NotionalAndLeverageBracketsResponse2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotionalAndLeverageBracketsResponse2WithDefaults

`func NewNotionalAndLeverageBracketsResponse2WithDefaults() *NotionalAndLeverageBracketsResponse2`

NewNotionalAndLeverageBracketsResponse2WithDefaults instantiates a new NotionalAndLeverageBracketsResponse2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *NotionalAndLeverageBracketsResponse2) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *NotionalAndLeverageBracketsResponse2) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *NotionalAndLeverageBracketsResponse2) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *NotionalAndLeverageBracketsResponse2) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetNotionalCoef

`func (o *NotionalAndLeverageBracketsResponse2) GetNotionalCoef() float32`

GetNotionalCoef returns the NotionalCoef field if non-nil, zero value otherwise.

### GetNotionalCoefOk

`func (o *NotionalAndLeverageBracketsResponse2) GetNotionalCoefOk() (*float32, bool)`

GetNotionalCoefOk returns a tuple with the NotionalCoef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotionalCoef

`func (o *NotionalAndLeverageBracketsResponse2) SetNotionalCoef(v float32)`

SetNotionalCoef sets NotionalCoef field to given value.

### HasNotionalCoef

`func (o *NotionalAndLeverageBracketsResponse2) HasNotionalCoef() bool`

HasNotionalCoef returns a boolean if a field has been set.

### GetBrackets

`func (o *NotionalAndLeverageBracketsResponse2) GetBrackets() []NotionalAndLeverageBracketsResponse2BracketsInner`

GetBrackets returns the Brackets field if non-nil, zero value otherwise.

### GetBracketsOk

`func (o *NotionalAndLeverageBracketsResponse2) GetBracketsOk() (*[]NotionalAndLeverageBracketsResponse2BracketsInner, bool)`

GetBracketsOk returns a tuple with the Brackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrackets

`func (o *NotionalAndLeverageBracketsResponse2) SetBrackets(v []NotionalAndLeverageBracketsResponse2BracketsInner)`

SetBrackets sets Brackets field to given value.

### HasBrackets

`func (o *NotionalAndLeverageBracketsResponse2) HasBrackets() bool`

HasBrackets returns a boolean if a field has been set.


[[Back to README]](../README.md)


