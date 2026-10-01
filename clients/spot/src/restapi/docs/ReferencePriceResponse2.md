# ReferencePriceResponse2

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**ReferencePrice** | Pointer to **interface{}** | Reference price. Can be &#x60;null&#x60; if no reference price is set. | [optional] 
**Timestamp** | Pointer to **int64** | Timestamp when reference price was valid | [optional] 

## Methods

### NewReferencePriceResponse2

`func NewReferencePriceResponse2() *ReferencePriceResponse2`

NewReferencePriceResponse2 instantiates a new ReferencePriceResponse2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferencePriceResponse2WithDefaults

`func NewReferencePriceResponse2WithDefaults() *ReferencePriceResponse2`

NewReferencePriceResponse2WithDefaults instantiates a new ReferencePriceResponse2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *ReferencePriceResponse2) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ReferencePriceResponse2) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ReferencePriceResponse2) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *ReferencePriceResponse2) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetReferencePrice

`func (o *ReferencePriceResponse2) GetReferencePrice() interface{}`

GetReferencePrice returns the ReferencePrice field if non-nil, zero value otherwise.

### GetReferencePriceOk

`func (o *ReferencePriceResponse2) GetReferencePriceOk() (*interface{}, bool)`

GetReferencePriceOk returns a tuple with the ReferencePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencePrice

`func (o *ReferencePriceResponse2) SetReferencePrice(v interface{})`

SetReferencePrice sets ReferencePrice field to given value.

### HasReferencePrice

`func (o *ReferencePriceResponse2) HasReferencePrice() bool`

HasReferencePrice returns a boolean if a field has been set.

### SetReferencePriceNil

`func (o *ReferencePriceResponse2) SetReferencePriceNil(b bool)`

 SetReferencePriceNil sets the value for ReferencePrice to be an explicit nil

### UnsetReferencePrice
`func (o *ReferencePriceResponse2) UnsetReferencePrice()`

UnsetReferencePrice ensures that no value is present for ReferencePrice, not even an explicit nil
### GetTimestamp

`func (o *ReferencePriceResponse2) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ReferencePriceResponse2) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ReferencePriceResponse2) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ReferencePriceResponse2) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to README]](../README.md)


