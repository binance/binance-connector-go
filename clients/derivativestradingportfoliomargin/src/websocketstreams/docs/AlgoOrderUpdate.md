# AlgoOrderUpdate

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**T** | Pointer to **int64** | Transaction Time | [optional] 
**E** | Pointer to **int64** | Event Time | [optional] 
**Fs** | Pointer to **string** | Futures segment, &#x60;UM&#x60; for USDS-M Futures, &#x60;CM&#x60; for Coin-M Futures | [optional] 
**Ao** | Pointer to [**AlgoOrderUpdateAo**](AlgoOrderUpdateAo.md) |  | [optional] 

## Methods

### NewAlgoOrderUpdate

`func NewAlgoOrderUpdate() *AlgoOrderUpdate`

NewAlgoOrderUpdate instantiates a new AlgoOrderUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlgoOrderUpdateWithDefaults

`func NewAlgoOrderUpdateWithDefaults() *AlgoOrderUpdate`

NewAlgoOrderUpdateWithDefaults instantiates a new AlgoOrderUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetT

`func (o *AlgoOrderUpdate) GetT() int64`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *AlgoOrderUpdate) GetTOk() (*int64, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *AlgoOrderUpdate) SetT(v int64)`

SetT sets T field to given value.

### HasT

`func (o *AlgoOrderUpdate) HasT() bool`

HasT returns a boolean if a field has been set.

### GetE

`func (o *AlgoOrderUpdate) GetE() int64`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *AlgoOrderUpdate) GetEOk() (*int64, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *AlgoOrderUpdate) SetE(v int64)`

SetE sets E field to given value.

### HasE

`func (o *AlgoOrderUpdate) HasE() bool`

HasE returns a boolean if a field has been set.

### GetFs

`func (o *AlgoOrderUpdate) GetFs() string`

GetFs returns the Fs field if non-nil, zero value otherwise.

### GetFsOk

`func (o *AlgoOrderUpdate) GetFsOk() (*string, bool)`

GetFsOk returns a tuple with the Fs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFs

`func (o *AlgoOrderUpdate) SetFs(v string)`

SetFs sets Fs field to given value.

### HasFs

`func (o *AlgoOrderUpdate) HasFs() bool`

HasFs returns a boolean if a field has been set.

### GetAo

`func (o *AlgoOrderUpdate) GetAo() AlgoOrderUpdateAo`

GetAo returns the Ao field if non-nil, zero value otherwise.

### GetAoOk

`func (o *AlgoOrderUpdate) GetAoOk() (*AlgoOrderUpdateAo, bool)`

GetAoOk returns a tuple with the Ao field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAo

`func (o *AlgoOrderUpdate) SetAo(v AlgoOrderUpdateAo)`

SetAo sets Ao field to given value.

### HasAo

`func (o *AlgoOrderUpdate) HasAo() bool`

HasAo returns a boolean if a field has been set.


[[Back to README]](../README.md)


