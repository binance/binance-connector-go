# GetOpenSymbolListResponseInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**OpenTime** | Pointer to **int64** |  | [optional] 
**Symbols** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetOpenSymbolListResponseInner

`func NewGetOpenSymbolListResponseInner() *GetOpenSymbolListResponseInner`

NewGetOpenSymbolListResponseInner instantiates a new GetOpenSymbolListResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOpenSymbolListResponseInnerWithDefaults

`func NewGetOpenSymbolListResponseInnerWithDefaults() *GetOpenSymbolListResponseInner`

NewGetOpenSymbolListResponseInnerWithDefaults instantiates a new GetOpenSymbolListResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpenTime

`func (o *GetOpenSymbolListResponseInner) GetOpenTime() int64`

GetOpenTime returns the OpenTime field if non-nil, zero value otherwise.

### GetOpenTimeOk

`func (o *GetOpenSymbolListResponseInner) GetOpenTimeOk() (*int64, bool)`

GetOpenTimeOk returns a tuple with the OpenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenTime

`func (o *GetOpenSymbolListResponseInner) SetOpenTime(v int64)`

SetOpenTime sets OpenTime field to given value.

### HasOpenTime

`func (o *GetOpenSymbolListResponseInner) HasOpenTime() bool`

HasOpenTime returns a boolean if a field has been set.

### GetSymbols

`func (o *GetOpenSymbolListResponseInner) GetSymbols() []string`

GetSymbols returns the Symbols field if non-nil, zero value otherwise.

### GetSymbolsOk

`func (o *GetOpenSymbolListResponseInner) GetSymbolsOk() (*[]string, bool)`

GetSymbolsOk returns a tuple with the Symbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbols

`func (o *GetOpenSymbolListResponseInner) SetSymbols(v []string)`

SetSymbols sets Symbols field to given value.

### HasSymbols

`func (o *GetOpenSymbolListResponseInner) HasSymbols() bool`

HasSymbols returns a boolean if a field has been set.


[[Back to README]](../README.md)


