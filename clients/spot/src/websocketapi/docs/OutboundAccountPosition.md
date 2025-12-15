# OutboundAccountPosition

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **int64** |  | [optional] 
**U** | Pointer to **int64** |  | [optional] 
**B** | Pointer to [**[]OutboundAccountPositionBInner**](OutboundAccountPositionBInner.md) |  | [optional] 

## Methods

### NewOutboundAccountPosition

`func NewOutboundAccountPosition() *OutboundAccountPosition`

NewOutboundAccountPosition instantiates a new OutboundAccountPosition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutboundAccountPositionWithDefaults

`func NewOutboundAccountPositionWithDefaults() *OutboundAccountPosition`

NewOutboundAccountPositionWithDefaults instantiates a new OutboundAccountPosition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *OutboundAccountPosition) GetE() int64`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *OutboundAccountPosition) GetEOk() (*int64, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *OutboundAccountPosition) SetE(v int64)`

SetE sets E field to given value.

### HasE

`func (o *OutboundAccountPosition) HasE() bool`

HasE returns a boolean if a field has been set.

### GetU

`func (o *OutboundAccountPosition) GetU() int64`

GetU returns the U field if non-nil, zero value otherwise.

### GetUOk

`func (o *OutboundAccountPosition) GetUOk() (*int64, bool)`

GetUOk returns a tuple with the U field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetU

`func (o *OutboundAccountPosition) SetU(v int64)`

SetU sets U field to given value.

### HasU

`func (o *OutboundAccountPosition) HasU() bool`

HasU returns a boolean if a field has been set.

### GetB

`func (o *OutboundAccountPosition) GetB() []OutboundAccountPositionBInner`

GetB returns the B field if non-nil, zero value otherwise.

### GetBOk

`func (o *OutboundAccountPosition) GetBOk() (*[]OutboundAccountPositionBInner, bool)`

GetBOk returns a tuple with the B field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB

`func (o *OutboundAccountPosition) SetB(v []OutboundAccountPositionBInner)`

SetB sets B field to given value.

### HasB

`func (o *OutboundAccountPosition) HasB() bool`

HasB returns a boolean if a field has been set.


[[Back to README]](../README.md)


