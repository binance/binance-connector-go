# ReferencePriceResponse2Result

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**ReferencePrice** | Pointer to **interface{}** |  | [optional] 
**Timestamp** | Pointer to **int64** | Timestamp when the reference price was valid | [optional] 

## Methods

### NewReferencePriceResponse2Result

`func NewReferencePriceResponse2Result() *ReferencePriceResponse2Result`

NewReferencePriceResponse2Result instantiates a new ReferencePriceResponse2Result object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferencePriceResponse2ResultWithDefaults

`func NewReferencePriceResponse2ResultWithDefaults() *ReferencePriceResponse2Result`

NewReferencePriceResponse2ResultWithDefaults instantiates a new ReferencePriceResponse2Result object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *ReferencePriceResponse2Result) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ReferencePriceResponse2Result) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ReferencePriceResponse2Result) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *ReferencePriceResponse2Result) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetReferencePrice

`func (o *ReferencePriceResponse2Result) GetReferencePrice() interface{}`

GetReferencePrice returns the ReferencePrice field if non-nil, zero value otherwise.

### GetReferencePriceOk

`func (o *ReferencePriceResponse2Result) GetReferencePriceOk() (*interface{}, bool)`

GetReferencePriceOk returns a tuple with the ReferencePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencePrice

`func (o *ReferencePriceResponse2Result) SetReferencePrice(v interface{})`

SetReferencePrice sets ReferencePrice field to given value.

### HasReferencePrice

`func (o *ReferencePriceResponse2Result) HasReferencePrice() bool`

HasReferencePrice returns a boolean if a field has been set.

### SetReferencePriceNil

`func (o *ReferencePriceResponse2Result) SetReferencePriceNil(b bool)`

 SetReferencePriceNil sets the value for ReferencePrice to be an explicit nil

### UnsetReferencePrice
`func (o *ReferencePriceResponse2Result) UnsetReferencePrice()`

UnsetReferencePrice ensures that no value is present for ReferencePrice, not even an explicit nil
### GetTimestamp

`func (o *ReferencePriceResponse2Result) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ReferencePriceResponse2Result) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ReferencePriceResponse2Result) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ReferencePriceResponse2Result) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to README]](../README.md)


