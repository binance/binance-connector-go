# ReferencePriceResponse1

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**ReferencePrice** | Pointer to **string** | Reference price. Can be &#x60;null&#x60; if no reference price is set. | [optional] 
**Timestamp** | Pointer to **int64** | Timestamp when reference price was valid. | [optional] 

## Methods

### NewReferencePriceResponse1

`func NewReferencePriceResponse1() *ReferencePriceResponse1`

NewReferencePriceResponse1 instantiates a new ReferencePriceResponse1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferencePriceResponse1WithDefaults

`func NewReferencePriceResponse1WithDefaults() *ReferencePriceResponse1`

NewReferencePriceResponse1WithDefaults instantiates a new ReferencePriceResponse1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *ReferencePriceResponse1) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ReferencePriceResponse1) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ReferencePriceResponse1) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *ReferencePriceResponse1) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetReferencePrice

`func (o *ReferencePriceResponse1) GetReferencePrice() string`

GetReferencePrice returns the ReferencePrice field if non-nil, zero value otherwise.

### GetReferencePriceOk

`func (o *ReferencePriceResponse1) GetReferencePriceOk() (*string, bool)`

GetReferencePriceOk returns a tuple with the ReferencePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencePrice

`func (o *ReferencePriceResponse1) SetReferencePrice(v string)`

SetReferencePrice sets ReferencePrice field to given value.

### HasReferencePrice

`func (o *ReferencePriceResponse1) HasReferencePrice() bool`

HasReferencePrice returns a boolean if a field has been set.

### GetTimestamp

`func (o *ReferencePriceResponse1) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ReferencePriceResponse1) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ReferencePriceResponse1) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ReferencePriceResponse1) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to README]](../README.md)


