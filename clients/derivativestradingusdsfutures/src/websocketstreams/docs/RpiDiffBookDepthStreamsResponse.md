# RpiDiffBookDepthStreamsResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **string** | Event type | [optional] 
**E** | Pointer to **int64** | Event time | [optional] 
**T** | Pointer to **int64** | Transaction time | [optional] 
**S** | Pointer to **string** | Symbol | [optional] 
**U** | Pointer to **int64** | First update ID in event | [optional] 
**U** | Pointer to **int64** | Final update ID in event | [optional] 
**Pu** | Pointer to **int64** | Final update Id in last stream(ie &#x60;u&#x60; in last stream) | [optional] 
**B** | Pointer to **[][]string** | Bids to be updated | [optional] 
**A** | Pointer to **[][]string** | Asks to be updated | [optional] 
**Ps** | Pointer to **string** | (After CM migration) Pair symbol | [optional] 
**St** | Pointer to **int32** | (After CM migration) Symbol type: 1 &#x3D; UM, 2 &#x3D; CM | [optional] 

## Methods

### NewRpiDiffBookDepthStreamsResponse

`func NewRpiDiffBookDepthStreamsResponse() *RpiDiffBookDepthStreamsResponse`

NewRpiDiffBookDepthStreamsResponse instantiates a new RpiDiffBookDepthStreamsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpiDiffBookDepthStreamsResponseWithDefaults

`func NewRpiDiffBookDepthStreamsResponseWithDefaults() *RpiDiffBookDepthStreamsResponse`

NewRpiDiffBookDepthStreamsResponseWithDefaults instantiates a new RpiDiffBookDepthStreamsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *RpiDiffBookDepthStreamsResponse) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *RpiDiffBookDepthStreamsResponse) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *RpiDiffBookDepthStreamsResponse) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *RpiDiffBookDepthStreamsResponse) HasE() bool`

HasE returns a boolean if a field has been set.

### GetE

`func (o *RpiDiffBookDepthStreamsResponse) GetE() int64`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *RpiDiffBookDepthStreamsResponse) GetEOk() (*int64, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *RpiDiffBookDepthStreamsResponse) SetE(v int64)`

SetE sets E field to given value.

### HasE

`func (o *RpiDiffBookDepthStreamsResponse) HasE() bool`

HasE returns a boolean if a field has been set.

### GetT

`func (o *RpiDiffBookDepthStreamsResponse) GetT() int64`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *RpiDiffBookDepthStreamsResponse) GetTOk() (*int64, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *RpiDiffBookDepthStreamsResponse) SetT(v int64)`

SetT sets T field to given value.

### HasT

`func (o *RpiDiffBookDepthStreamsResponse) HasT() bool`

HasT returns a boolean if a field has been set.

### GetS

`func (o *RpiDiffBookDepthStreamsResponse) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *RpiDiffBookDepthStreamsResponse) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *RpiDiffBookDepthStreamsResponse) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *RpiDiffBookDepthStreamsResponse) HasS() bool`

HasS returns a boolean if a field has been set.

### GetU

`func (o *RpiDiffBookDepthStreamsResponse) GetU() int64`

GetU returns the U field if non-nil, zero value otherwise.

### GetUOk

`func (o *RpiDiffBookDepthStreamsResponse) GetUOk() (*int64, bool)`

GetUOk returns a tuple with the U field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetU

`func (o *RpiDiffBookDepthStreamsResponse) SetU(v int64)`

SetU sets U field to given value.

### HasU

`func (o *RpiDiffBookDepthStreamsResponse) HasU() bool`

HasU returns a boolean if a field has been set.

### GetU

`func (o *RpiDiffBookDepthStreamsResponse) GetU() int64`

GetU returns the U field if non-nil, zero value otherwise.

### GetUOk

`func (o *RpiDiffBookDepthStreamsResponse) GetUOk() (*int64, bool)`

GetUOk returns a tuple with the U field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetU

`func (o *RpiDiffBookDepthStreamsResponse) SetU(v int64)`

SetU sets U field to given value.

### HasU

`func (o *RpiDiffBookDepthStreamsResponse) HasU() bool`

HasU returns a boolean if a field has been set.

### GetPu

`func (o *RpiDiffBookDepthStreamsResponse) GetPu() int64`

GetPu returns the Pu field if non-nil, zero value otherwise.

### GetPuOk

`func (o *RpiDiffBookDepthStreamsResponse) GetPuOk() (*int64, bool)`

GetPuOk returns a tuple with the Pu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPu

`func (o *RpiDiffBookDepthStreamsResponse) SetPu(v int64)`

SetPu sets Pu field to given value.

### HasPu

`func (o *RpiDiffBookDepthStreamsResponse) HasPu() bool`

HasPu returns a boolean if a field has been set.

### GetB

`func (o *RpiDiffBookDepthStreamsResponse) GetB() [][]string`

GetB returns the B field if non-nil, zero value otherwise.

### GetBOk

`func (o *RpiDiffBookDepthStreamsResponse) GetBOk() (*[][]string, bool)`

GetBOk returns a tuple with the B field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB

`func (o *RpiDiffBookDepthStreamsResponse) SetB(v [][]string)`

SetB sets B field to given value.

### HasB

`func (o *RpiDiffBookDepthStreamsResponse) HasB() bool`

HasB returns a boolean if a field has been set.

### GetA

`func (o *RpiDiffBookDepthStreamsResponse) GetA() [][]string`

GetA returns the A field if non-nil, zero value otherwise.

### GetAOk

`func (o *RpiDiffBookDepthStreamsResponse) GetAOk() (*[][]string, bool)`

GetAOk returns a tuple with the A field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetA

`func (o *RpiDiffBookDepthStreamsResponse) SetA(v [][]string)`

SetA sets A field to given value.

### HasA

`func (o *RpiDiffBookDepthStreamsResponse) HasA() bool`

HasA returns a boolean if a field has been set.

### GetPs

`func (o *RpiDiffBookDepthStreamsResponse) GetPs() string`

GetPs returns the Ps field if non-nil, zero value otherwise.

### GetPsOk

`func (o *RpiDiffBookDepthStreamsResponse) GetPsOk() (*string, bool)`

GetPsOk returns a tuple with the Ps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPs

`func (o *RpiDiffBookDepthStreamsResponse) SetPs(v string)`

SetPs sets Ps field to given value.

### HasPs

`func (o *RpiDiffBookDepthStreamsResponse) HasPs() bool`

HasPs returns a boolean if a field has been set.

### GetSt

`func (o *RpiDiffBookDepthStreamsResponse) GetSt() int32`

GetSt returns the St field if non-nil, zero value otherwise.

### GetStOk

`func (o *RpiDiffBookDepthStreamsResponse) GetStOk() (*int32, bool)`

GetStOk returns a tuple with the St field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSt

`func (o *RpiDiffBookDepthStreamsResponse) SetSt(v int32)`

SetSt sets St field to given value.

### HasSt

`func (o *RpiDiffBookDepthStreamsResponse) HasSt() bool`

HasSt returns a boolean if a field has been set.


[[Back to README]](../README.md)


