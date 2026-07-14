# PartialDepthStreamResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **string** | eventType | [optional] 
**E** | Pointer to **int64** | eventTime | [optional] 
**T** | Pointer to **int64** | transactionTime | [optional] 
**U** | Pointer to **int64** | firstUpdateId | [optional] 
**U** | Pointer to **int64** | lastUpdateId | [optional] 
**Pu** | Pointer to **int64** | previousUpdateId | [optional] 
**S** | Pointer to **string** | symbol | [optional] 
**B** | Pointer to **[][]string** | bids to be updated | [optional] 
**A** | Pointer to **[][]string** | asks to be updated | [optional] 

## Methods

### NewPartialDepthStreamResponse

`func NewPartialDepthStreamResponse() *PartialDepthStreamResponse`

NewPartialDepthStreamResponse instantiates a new PartialDepthStreamResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartialDepthStreamResponseWithDefaults

`func NewPartialDepthStreamResponseWithDefaults() *PartialDepthStreamResponse`

NewPartialDepthStreamResponseWithDefaults instantiates a new PartialDepthStreamResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *PartialDepthStreamResponse) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *PartialDepthStreamResponse) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *PartialDepthStreamResponse) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *PartialDepthStreamResponse) HasE() bool`

HasE returns a boolean if a field has been set.

### GetE

`func (o *PartialDepthStreamResponse) GetE() int64`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *PartialDepthStreamResponse) GetEOk() (*int64, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *PartialDepthStreamResponse) SetE(v int64)`

SetE sets E field to given value.

### HasE

`func (o *PartialDepthStreamResponse) HasE() bool`

HasE returns a boolean if a field has been set.

### GetT

`func (o *PartialDepthStreamResponse) GetT() int64`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *PartialDepthStreamResponse) GetTOk() (*int64, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *PartialDepthStreamResponse) SetT(v int64)`

SetT sets T field to given value.

### HasT

`func (o *PartialDepthStreamResponse) HasT() bool`

HasT returns a boolean if a field has been set.

### GetU

`func (o *PartialDepthStreamResponse) GetU() int64`

GetU returns the U field if non-nil, zero value otherwise.

### GetUOk

`func (o *PartialDepthStreamResponse) GetUOk() (*int64, bool)`

GetUOk returns a tuple with the U field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetU

`func (o *PartialDepthStreamResponse) SetU(v int64)`

SetU sets U field to given value.

### HasU

`func (o *PartialDepthStreamResponse) HasU() bool`

HasU returns a boolean if a field has been set.

### GetU

`func (o *PartialDepthStreamResponse) GetU() int64`

GetU returns the U field if non-nil, zero value otherwise.

### GetUOk

`func (o *PartialDepthStreamResponse) GetUOk() (*int64, bool)`

GetUOk returns a tuple with the U field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetU

`func (o *PartialDepthStreamResponse) SetU(v int64)`

SetU sets U field to given value.

### HasU

`func (o *PartialDepthStreamResponse) HasU() bool`

HasU returns a boolean if a field has been set.

### GetPu

`func (o *PartialDepthStreamResponse) GetPu() int64`

GetPu returns the Pu field if non-nil, zero value otherwise.

### GetPuOk

`func (o *PartialDepthStreamResponse) GetPuOk() (*int64, bool)`

GetPuOk returns a tuple with the Pu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPu

`func (o *PartialDepthStreamResponse) SetPu(v int64)`

SetPu sets Pu field to given value.

### HasPu

`func (o *PartialDepthStreamResponse) HasPu() bool`

HasPu returns a boolean if a field has been set.

### GetS

`func (o *PartialDepthStreamResponse) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *PartialDepthStreamResponse) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *PartialDepthStreamResponse) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *PartialDepthStreamResponse) HasS() bool`

HasS returns a boolean if a field has been set.

### GetB

`func (o *PartialDepthStreamResponse) GetB() [][]string`

GetB returns the B field if non-nil, zero value otherwise.

### GetBOk

`func (o *PartialDepthStreamResponse) GetBOk() (*[][]string, bool)`

GetBOk returns a tuple with the B field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB

`func (o *PartialDepthStreamResponse) SetB(v [][]string)`

SetB sets B field to given value.

### HasB

`func (o *PartialDepthStreamResponse) HasB() bool`

HasB returns a boolean if a field has been set.

### GetA

`func (o *PartialDepthStreamResponse) GetA() [][]string`

GetA returns the A field if non-nil, zero value otherwise.

### GetAOk

`func (o *PartialDepthStreamResponse) GetAOk() (*[][]string, bool)`

GetAOk returns a tuple with the A field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetA

`func (o *PartialDepthStreamResponse) SetA(v [][]string)`

SetA sets A field to given value.

### HasA

`func (o *PartialDepthStreamResponse) HasA() bool`

HasA returns a boolean if a field has been set.


[[Back to README]](../README.md)


