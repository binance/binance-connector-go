# FullDepthStreamResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **string** | Event type | [optional] 
**E** | Pointer to **int64** | Event time (ms) | [optional] 
**T** | Pointer to **int64** | Matching time (ms) | [optional] 
**U** | Pointer to **int64** | First updateId in this event | [optional] 
**U** | Pointer to **int64** | Last updateId in this event | [optional] 
**Pu** | Pointer to **int64** | Previous updateId from the last push | [optional] 
**S** | Pointer to **string** | Symbol | [optional] 
**B** | Pointer to **[][]string** | bids to be updated | [optional] 
**A** | Pointer to **[][]string** | asks to be updated | [optional] 

## Methods

### NewFullDepthStreamResponse

`func NewFullDepthStreamResponse() *FullDepthStreamResponse`

NewFullDepthStreamResponse instantiates a new FullDepthStreamResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullDepthStreamResponseWithDefaults

`func NewFullDepthStreamResponseWithDefaults() *FullDepthStreamResponse`

NewFullDepthStreamResponseWithDefaults instantiates a new FullDepthStreamResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *FullDepthStreamResponse) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *FullDepthStreamResponse) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *FullDepthStreamResponse) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *FullDepthStreamResponse) HasE() bool`

HasE returns a boolean if a field has been set.

### GetE

`func (o *FullDepthStreamResponse) GetE() int64`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *FullDepthStreamResponse) GetEOk() (*int64, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *FullDepthStreamResponse) SetE(v int64)`

SetE sets E field to given value.

### HasE

`func (o *FullDepthStreamResponse) HasE() bool`

HasE returns a boolean if a field has been set.

### GetT

`func (o *FullDepthStreamResponse) GetT() int64`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *FullDepthStreamResponse) GetTOk() (*int64, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *FullDepthStreamResponse) SetT(v int64)`

SetT sets T field to given value.

### HasT

`func (o *FullDepthStreamResponse) HasT() bool`

HasT returns a boolean if a field has been set.

### GetU

`func (o *FullDepthStreamResponse) GetU() int64`

GetU returns the U field if non-nil, zero value otherwise.

### GetUOk

`func (o *FullDepthStreamResponse) GetUOk() (*int64, bool)`

GetUOk returns a tuple with the U field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetU

`func (o *FullDepthStreamResponse) SetU(v int64)`

SetU sets U field to given value.

### HasU

`func (o *FullDepthStreamResponse) HasU() bool`

HasU returns a boolean if a field has been set.

### GetU

`func (o *FullDepthStreamResponse) GetU() int64`

GetU returns the U field if non-nil, zero value otherwise.

### GetUOk

`func (o *FullDepthStreamResponse) GetUOk() (*int64, bool)`

GetUOk returns a tuple with the U field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetU

`func (o *FullDepthStreamResponse) SetU(v int64)`

SetU sets U field to given value.

### HasU

`func (o *FullDepthStreamResponse) HasU() bool`

HasU returns a boolean if a field has been set.

### GetPu

`func (o *FullDepthStreamResponse) GetPu() int64`

GetPu returns the Pu field if non-nil, zero value otherwise.

### GetPuOk

`func (o *FullDepthStreamResponse) GetPuOk() (*int64, bool)`

GetPuOk returns a tuple with the Pu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPu

`func (o *FullDepthStreamResponse) SetPu(v int64)`

SetPu sets Pu field to given value.

### HasPu

`func (o *FullDepthStreamResponse) HasPu() bool`

HasPu returns a boolean if a field has been set.

### GetS

`func (o *FullDepthStreamResponse) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *FullDepthStreamResponse) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *FullDepthStreamResponse) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *FullDepthStreamResponse) HasS() bool`

HasS returns a boolean if a field has been set.

### GetB

`func (o *FullDepthStreamResponse) GetB() [][]string`

GetB returns the B field if non-nil, zero value otherwise.

### GetBOk

`func (o *FullDepthStreamResponse) GetBOk() (*[][]string, bool)`

GetBOk returns a tuple with the B field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB

`func (o *FullDepthStreamResponse) SetB(v [][]string)`

SetB sets B field to given value.

### HasB

`func (o *FullDepthStreamResponse) HasB() bool`

HasB returns a boolean if a field has been set.

### GetA

`func (o *FullDepthStreamResponse) GetA() [][]string`

GetA returns the A field if non-nil, zero value otherwise.

### GetAOk

`func (o *FullDepthStreamResponse) GetAOk() (*[][]string, bool)`

GetAOk returns a tuple with the A field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetA

`func (o *FullDepthStreamResponse) SetA(v [][]string)`

SetA sets A field to given value.

### HasA

`func (o *FullDepthStreamResponse) HasA() bool`

HasA returns a boolean if a field has been set.


[[Back to README]](../README.md)


