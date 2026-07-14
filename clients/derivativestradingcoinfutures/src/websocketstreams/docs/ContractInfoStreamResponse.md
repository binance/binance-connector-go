# ContractInfoStreamResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **string** | Event Type | [optional] 
**E** | Pointer to **int64** | Event Time | [optional] 
**S** | Pointer to **string** | Symbol | [optional] 
**Ps** | Pointer to **string** | Pair | [optional] 
**Ct** | Pointer to **string** | Contract type | [optional] 
**Dt** | Pointer to **int64** | Delivery date time | [optional] 
**Ot** | Pointer to **int64** | onboard date time | [optional] 
**Cs** | Pointer to **string** | Contract status | [optional] 
**Bks** | Pointer to [**[]ContractInfoStreamResponseBksInner**](ContractInfoStreamResponseBksInner.md) | Bracket list. | [optional] 
**St** | Pointer to **int32** | (After CM migration) Symbol type: 1 &#x3D; UM, 2 &#x3D; CM | [optional] 

## Methods

### NewContractInfoStreamResponse

`func NewContractInfoStreamResponse() *ContractInfoStreamResponse`

NewContractInfoStreamResponse instantiates a new ContractInfoStreamResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractInfoStreamResponseWithDefaults

`func NewContractInfoStreamResponseWithDefaults() *ContractInfoStreamResponse`

NewContractInfoStreamResponseWithDefaults instantiates a new ContractInfoStreamResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *ContractInfoStreamResponse) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *ContractInfoStreamResponse) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *ContractInfoStreamResponse) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *ContractInfoStreamResponse) HasE() bool`

HasE returns a boolean if a field has been set.

### GetE

`func (o *ContractInfoStreamResponse) GetE() int64`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *ContractInfoStreamResponse) GetEOk() (*int64, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *ContractInfoStreamResponse) SetE(v int64)`

SetE sets E field to given value.

### HasE

`func (o *ContractInfoStreamResponse) HasE() bool`

HasE returns a boolean if a field has been set.

### GetS

`func (o *ContractInfoStreamResponse) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *ContractInfoStreamResponse) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *ContractInfoStreamResponse) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *ContractInfoStreamResponse) HasS() bool`

HasS returns a boolean if a field has been set.

### GetPs

`func (o *ContractInfoStreamResponse) GetPs() string`

GetPs returns the Ps field if non-nil, zero value otherwise.

### GetPsOk

`func (o *ContractInfoStreamResponse) GetPsOk() (*string, bool)`

GetPsOk returns a tuple with the Ps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPs

`func (o *ContractInfoStreamResponse) SetPs(v string)`

SetPs sets Ps field to given value.

### HasPs

`func (o *ContractInfoStreamResponse) HasPs() bool`

HasPs returns a boolean if a field has been set.

### GetCt

`func (o *ContractInfoStreamResponse) GetCt() string`

GetCt returns the Ct field if non-nil, zero value otherwise.

### GetCtOk

`func (o *ContractInfoStreamResponse) GetCtOk() (*string, bool)`

GetCtOk returns a tuple with the Ct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCt

`func (o *ContractInfoStreamResponse) SetCt(v string)`

SetCt sets Ct field to given value.

### HasCt

`func (o *ContractInfoStreamResponse) HasCt() bool`

HasCt returns a boolean if a field has been set.

### GetDt

`func (o *ContractInfoStreamResponse) GetDt() int64`

GetDt returns the Dt field if non-nil, zero value otherwise.

### GetDtOk

`func (o *ContractInfoStreamResponse) GetDtOk() (*int64, bool)`

GetDtOk returns a tuple with the Dt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDt

`func (o *ContractInfoStreamResponse) SetDt(v int64)`

SetDt sets Dt field to given value.

### HasDt

`func (o *ContractInfoStreamResponse) HasDt() bool`

HasDt returns a boolean if a field has been set.

### GetOt

`func (o *ContractInfoStreamResponse) GetOt() int64`

GetOt returns the Ot field if non-nil, zero value otherwise.

### GetOtOk

`func (o *ContractInfoStreamResponse) GetOtOk() (*int64, bool)`

GetOtOk returns a tuple with the Ot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOt

`func (o *ContractInfoStreamResponse) SetOt(v int64)`

SetOt sets Ot field to given value.

### HasOt

`func (o *ContractInfoStreamResponse) HasOt() bool`

HasOt returns a boolean if a field has been set.

### GetCs

`func (o *ContractInfoStreamResponse) GetCs() string`

GetCs returns the Cs field if non-nil, zero value otherwise.

### GetCsOk

`func (o *ContractInfoStreamResponse) GetCsOk() (*string, bool)`

GetCsOk returns a tuple with the Cs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCs

`func (o *ContractInfoStreamResponse) SetCs(v string)`

SetCs sets Cs field to given value.

### HasCs

`func (o *ContractInfoStreamResponse) HasCs() bool`

HasCs returns a boolean if a field has been set.

### GetBks

`func (o *ContractInfoStreamResponse) GetBks() []ContractInfoStreamResponseBksInner`

GetBks returns the Bks field if non-nil, zero value otherwise.

### GetBksOk

`func (o *ContractInfoStreamResponse) GetBksOk() (*[]ContractInfoStreamResponseBksInner, bool)`

GetBksOk returns a tuple with the Bks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBks

`func (o *ContractInfoStreamResponse) SetBks(v []ContractInfoStreamResponseBksInner)`

SetBks sets Bks field to given value.

### HasBks

`func (o *ContractInfoStreamResponse) HasBks() bool`

HasBks returns a boolean if a field has been set.

### GetSt

`func (o *ContractInfoStreamResponse) GetSt() int32`

GetSt returns the St field if non-nil, zero value otherwise.

### GetStOk

`func (o *ContractInfoStreamResponse) GetStOk() (*int32, bool)`

GetStOk returns a tuple with the St field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSt

`func (o *ContractInfoStreamResponse) SetSt(v int32)`

SetSt sets St field to given value.

### HasSt

`func (o *ContractInfoStreamResponse) HasSt() bool`

HasSt returns a boolean if a field has been set.


[[Back to README]](../README.md)


