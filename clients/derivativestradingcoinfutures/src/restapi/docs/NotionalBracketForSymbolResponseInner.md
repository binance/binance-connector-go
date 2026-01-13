# NotionalBracketForSymbolResponseInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**NotionalCoef** | Pointer to **float32** |  | [optional] 
**Brackets** | Pointer to [**[]NotionalBracketForPairResponseInnerBracketsInner**](NotionalBracketForPairResponseInnerBracketsInner.md) |  | [optional] 

## Methods

### NewNotionalBracketForSymbolResponseInner

`func NewNotionalBracketForSymbolResponseInner() *NotionalBracketForSymbolResponseInner`

NewNotionalBracketForSymbolResponseInner instantiates a new NotionalBracketForSymbolResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotionalBracketForSymbolResponseInnerWithDefaults

`func NewNotionalBracketForSymbolResponseInnerWithDefaults() *NotionalBracketForSymbolResponseInner`

NewNotionalBracketForSymbolResponseInnerWithDefaults instantiates a new NotionalBracketForSymbolResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *NotionalBracketForSymbolResponseInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *NotionalBracketForSymbolResponseInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *NotionalBracketForSymbolResponseInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *NotionalBracketForSymbolResponseInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetNotionalCoef

`func (o *NotionalBracketForSymbolResponseInner) GetNotionalCoef() float32`

GetNotionalCoef returns the NotionalCoef field if non-nil, zero value otherwise.

### GetNotionalCoefOk

`func (o *NotionalBracketForSymbolResponseInner) GetNotionalCoefOk() (*float32, bool)`

GetNotionalCoefOk returns a tuple with the NotionalCoef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotionalCoef

`func (o *NotionalBracketForSymbolResponseInner) SetNotionalCoef(v float32)`

SetNotionalCoef sets NotionalCoef field to given value.

### HasNotionalCoef

`func (o *NotionalBracketForSymbolResponseInner) HasNotionalCoef() bool`

HasNotionalCoef returns a boolean if a field has been set.

### GetBrackets

`func (o *NotionalBracketForSymbolResponseInner) GetBrackets() []NotionalBracketForPairResponseInnerBracketsInner`

GetBrackets returns the Brackets field if non-nil, zero value otherwise.

### GetBracketsOk

`func (o *NotionalBracketForSymbolResponseInner) GetBracketsOk() (*[]NotionalBracketForPairResponseInnerBracketsInner, bool)`

GetBracketsOk returns a tuple with the Brackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrackets

`func (o *NotionalBracketForSymbolResponseInner) SetBrackets(v []NotionalBracketForPairResponseInnerBracketsInner)`

SetBrackets sets Brackets field to given value.

### HasBrackets

`func (o *NotionalBracketForSymbolResponseInner) HasBrackets() bool`

HasBrackets returns a boolean if a field has been set.


[[Back to README]](../README.md)


