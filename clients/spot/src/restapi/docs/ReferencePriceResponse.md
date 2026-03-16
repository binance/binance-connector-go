# ReferencePriceResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**ReferencePrice** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 

## Methods

### NewReferencePriceResponse

`func NewReferencePriceResponse() *ReferencePriceResponse`

NewReferencePriceResponse instantiates a new ReferencePriceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferencePriceResponseWithDefaults

`func NewReferencePriceResponseWithDefaults() *ReferencePriceResponse`

NewReferencePriceResponseWithDefaults instantiates a new ReferencePriceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *ReferencePriceResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ReferencePriceResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ReferencePriceResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *ReferencePriceResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetReferencePrice

`func (o *ReferencePriceResponse) GetReferencePrice() string`

GetReferencePrice returns the ReferencePrice field if non-nil, zero value otherwise.

### GetReferencePriceOk

`func (o *ReferencePriceResponse) GetReferencePriceOk() (*string, bool)`

GetReferencePriceOk returns a tuple with the ReferencePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencePrice

`func (o *ReferencePriceResponse) SetReferencePrice(v string)`

SetReferencePrice sets ReferencePrice field to given value.

### HasReferencePrice

`func (o *ReferencePriceResponse) HasReferencePrice() bool`

HasReferencePrice returns a boolean if a field has been set.

### GetTimestamp

`func (o *ReferencePriceResponse) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ReferencePriceResponse) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ReferencePriceResponse) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ReferencePriceResponse) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to README]](../README.md)


