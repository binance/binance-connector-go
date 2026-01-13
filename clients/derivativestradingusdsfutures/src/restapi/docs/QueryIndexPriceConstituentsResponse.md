# QueryIndexPriceConstituentsResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 
**Constituents** | Pointer to [**[]QueryIndexPriceConstituentsResponseConstituentsInner**](QueryIndexPriceConstituentsResponseConstituentsInner.md) |  | [optional] 

## Methods

### NewQueryIndexPriceConstituentsResponse

`func NewQueryIndexPriceConstituentsResponse() *QueryIndexPriceConstituentsResponse`

NewQueryIndexPriceConstituentsResponse instantiates a new QueryIndexPriceConstituentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryIndexPriceConstituentsResponseWithDefaults

`func NewQueryIndexPriceConstituentsResponseWithDefaults() *QueryIndexPriceConstituentsResponse`

NewQueryIndexPriceConstituentsResponseWithDefaults instantiates a new QueryIndexPriceConstituentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *QueryIndexPriceConstituentsResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *QueryIndexPriceConstituentsResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *QueryIndexPriceConstituentsResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *QueryIndexPriceConstituentsResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTime

`func (o *QueryIndexPriceConstituentsResponse) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *QueryIndexPriceConstituentsResponse) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *QueryIndexPriceConstituentsResponse) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *QueryIndexPriceConstituentsResponse) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetConstituents

`func (o *QueryIndexPriceConstituentsResponse) GetConstituents() []QueryIndexPriceConstituentsResponseConstituentsInner`

GetConstituents returns the Constituents field if non-nil, zero value otherwise.

### GetConstituentsOk

`func (o *QueryIndexPriceConstituentsResponse) GetConstituentsOk() (*[]QueryIndexPriceConstituentsResponseConstituentsInner, bool)`

GetConstituentsOk returns a tuple with the Constituents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstituents

`func (o *QueryIndexPriceConstituentsResponse) SetConstituents(v []QueryIndexPriceConstituentsResponseConstituentsInner)`

SetConstituents sets Constituents field to given value.

### HasConstituents

`func (o *QueryIndexPriceConstituentsResponse) HasConstituents() bool`

HasConstituents returns a boolean if a field has been set.


[[Back to README]](../README.md)


