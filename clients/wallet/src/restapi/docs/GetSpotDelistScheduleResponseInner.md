# GetSpotDelistScheduleResponseInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**DelistTime** | Pointer to **int64** |  | [optional] 
**Symbols** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetSpotDelistScheduleResponseInner

`func NewGetSpotDelistScheduleResponseInner() *GetSpotDelistScheduleResponseInner`

NewGetSpotDelistScheduleResponseInner instantiates a new GetSpotDelistScheduleResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSpotDelistScheduleResponseInnerWithDefaults

`func NewGetSpotDelistScheduleResponseInnerWithDefaults() *GetSpotDelistScheduleResponseInner`

NewGetSpotDelistScheduleResponseInnerWithDefaults instantiates a new GetSpotDelistScheduleResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelistTime

`func (o *GetSpotDelistScheduleResponseInner) GetDelistTime() int64`

GetDelistTime returns the DelistTime field if non-nil, zero value otherwise.

### GetDelistTimeOk

`func (o *GetSpotDelistScheduleResponseInner) GetDelistTimeOk() (*int64, bool)`

GetDelistTimeOk returns a tuple with the DelistTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelistTime

`func (o *GetSpotDelistScheduleResponseInner) SetDelistTime(v int64)`

SetDelistTime sets DelistTime field to given value.

### HasDelistTime

`func (o *GetSpotDelistScheduleResponseInner) HasDelistTime() bool`

HasDelistTime returns a boolean if a field has been set.

### GetSymbols

`func (o *GetSpotDelistScheduleResponseInner) GetSymbols() []string`

GetSymbols returns the Symbols field if non-nil, zero value otherwise.

### GetSymbolsOk

`func (o *GetSpotDelistScheduleResponseInner) GetSymbolsOk() (*[]string, bool)`

GetSymbolsOk returns a tuple with the Symbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbols

`func (o *GetSpotDelistScheduleResponseInner) SetSymbols(v []string)`

SetSymbols sets Symbols field to given value.

### HasSymbols

`func (o *GetSpotDelistScheduleResponseInner) HasSymbols() bool`

HasSymbols returns a boolean if a field has been set.


[[Back to README]](../README.md)


