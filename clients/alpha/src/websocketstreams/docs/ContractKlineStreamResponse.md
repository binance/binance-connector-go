# ContractKlineStreamResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Ca** | Pointer to **string** | Contract address@Chain ID | [optional] 
**E** | Pointer to **string** | Event type | [optional] 
**K** | Pointer to [**ContractKlineStreamResponseK**](ContractKlineStreamResponseK.md) |  | [optional] 

## Methods

### NewContractKlineStreamResponse

`func NewContractKlineStreamResponse() *ContractKlineStreamResponse`

NewContractKlineStreamResponse instantiates a new ContractKlineStreamResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractKlineStreamResponseWithDefaults

`func NewContractKlineStreamResponseWithDefaults() *ContractKlineStreamResponse`

NewContractKlineStreamResponseWithDefaults instantiates a new ContractKlineStreamResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCa

`func (o *ContractKlineStreamResponse) GetCa() string`

GetCa returns the Ca field if non-nil, zero value otherwise.

### GetCaOk

`func (o *ContractKlineStreamResponse) GetCaOk() (*string, bool)`

GetCaOk returns a tuple with the Ca field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCa

`func (o *ContractKlineStreamResponse) SetCa(v string)`

SetCa sets Ca field to given value.

### HasCa

`func (o *ContractKlineStreamResponse) HasCa() bool`

HasCa returns a boolean if a field has been set.

### GetE

`func (o *ContractKlineStreamResponse) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *ContractKlineStreamResponse) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *ContractKlineStreamResponse) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *ContractKlineStreamResponse) HasE() bool`

HasE returns a boolean if a field has been set.

### GetK

`func (o *ContractKlineStreamResponse) GetK() ContractKlineStreamResponseK`

GetK returns the K field if non-nil, zero value otherwise.

### GetKOk

`func (o *ContractKlineStreamResponse) GetKOk() (*ContractKlineStreamResponseK, bool)`

GetKOk returns a tuple with the K field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK

`func (o *ContractKlineStreamResponse) SetK(v ContractKlineStreamResponseK)`

SetK sets K field to given value.

### HasK

`func (o *ContractKlineStreamResponse) HasK() bool`

HasK returns a boolean if a field has been set.


[[Back to README]](../README.md)


