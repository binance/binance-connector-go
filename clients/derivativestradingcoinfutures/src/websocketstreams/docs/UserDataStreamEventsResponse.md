# UserDataStreamEventsResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **int64** |  | [optional] 
**T** | Pointer to **int64** |  | [optional] 
**Ac** | Pointer to [**AccountConfigUpdateAc**](AccountConfigUpdateAc.md) |  | [optional] 
**I** | Pointer to **string** |  | [optional] 
**A** | Pointer to [**AccountUpdateA**](AccountUpdateA.md) |  | [optional] 
**Gu** | Pointer to [**GridUpdateGu**](GridUpdateGu.md) |  | [optional] 
**Cw** | Pointer to **string** |  | [optional] 
**P** | Pointer to [**[]MarginCallPInner**](MarginCallPInner.md) |  | [optional] 
**O** | Pointer to [**OrderTradeUpdateO**](OrderTradeUpdateO.md) |  | [optional] 
**Su** | Pointer to [**StrategyUpdateSu**](StrategyUpdateSu.md) |  | [optional] 
**ListenKey** | Pointer to **string** |  | [optional] 

## Methods

### NewUserDataStreamEventsResponse

`func NewUserDataStreamEventsResponse() *UserDataStreamEventsResponse`

NewUserDataStreamEventsResponse instantiates a new UserDataStreamEventsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDataStreamEventsResponseWithDefaults

`func NewUserDataStreamEventsResponseWithDefaults() *UserDataStreamEventsResponse`

NewUserDataStreamEventsResponseWithDefaults instantiates a new UserDataStreamEventsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *UserDataStreamEventsResponse) GetE() int64`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *UserDataStreamEventsResponse) GetEOk() (*int64, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *UserDataStreamEventsResponse) SetE(v int64)`

SetE sets E field to given value.

### HasE

`func (o *UserDataStreamEventsResponse) HasE() bool`

HasE returns a boolean if a field has been set.

### GetT

`func (o *UserDataStreamEventsResponse) GetT() int64`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *UserDataStreamEventsResponse) GetTOk() (*int64, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *UserDataStreamEventsResponse) SetT(v int64)`

SetT sets T field to given value.

### HasT

`func (o *UserDataStreamEventsResponse) HasT() bool`

HasT returns a boolean if a field has been set.

### GetAc

`func (o *UserDataStreamEventsResponse) GetAc() AccountConfigUpdateAc`

GetAc returns the Ac field if non-nil, zero value otherwise.

### GetAcOk

`func (o *UserDataStreamEventsResponse) GetAcOk() (*AccountConfigUpdateAc, bool)`

GetAcOk returns a tuple with the Ac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAc

`func (o *UserDataStreamEventsResponse) SetAc(v AccountConfigUpdateAc)`

SetAc sets Ac field to given value.

### HasAc

`func (o *UserDataStreamEventsResponse) HasAc() bool`

HasAc returns a boolean if a field has been set.

### GetI

`func (o *UserDataStreamEventsResponse) GetI() string`

GetI returns the I field if non-nil, zero value otherwise.

### GetIOk

`func (o *UserDataStreamEventsResponse) GetIOk() (*string, bool)`

GetIOk returns a tuple with the I field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetI

`func (o *UserDataStreamEventsResponse) SetI(v string)`

SetI sets I field to given value.

### HasI

`func (o *UserDataStreamEventsResponse) HasI() bool`

HasI returns a boolean if a field has been set.

### GetA

`func (o *UserDataStreamEventsResponse) GetA() AccountUpdateA`

GetA returns the A field if non-nil, zero value otherwise.

### GetAOk

`func (o *UserDataStreamEventsResponse) GetAOk() (*AccountUpdateA, bool)`

GetAOk returns a tuple with the A field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetA

`func (o *UserDataStreamEventsResponse) SetA(v AccountUpdateA)`

SetA sets A field to given value.

### HasA

`func (o *UserDataStreamEventsResponse) HasA() bool`

HasA returns a boolean if a field has been set.

### GetGu

`func (o *UserDataStreamEventsResponse) GetGu() GridUpdateGu`

GetGu returns the Gu field if non-nil, zero value otherwise.

### GetGuOk

`func (o *UserDataStreamEventsResponse) GetGuOk() (*GridUpdateGu, bool)`

GetGuOk returns a tuple with the Gu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGu

`func (o *UserDataStreamEventsResponse) SetGu(v GridUpdateGu)`

SetGu sets Gu field to given value.

### HasGu

`func (o *UserDataStreamEventsResponse) HasGu() bool`

HasGu returns a boolean if a field has been set.

### GetCw

`func (o *UserDataStreamEventsResponse) GetCw() string`

GetCw returns the Cw field if non-nil, zero value otherwise.

### GetCwOk

`func (o *UserDataStreamEventsResponse) GetCwOk() (*string, bool)`

GetCwOk returns a tuple with the Cw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCw

`func (o *UserDataStreamEventsResponse) SetCw(v string)`

SetCw sets Cw field to given value.

### HasCw

`func (o *UserDataStreamEventsResponse) HasCw() bool`

HasCw returns a boolean if a field has been set.

### GetP

`func (o *UserDataStreamEventsResponse) GetP() []MarginCallPInner`

GetP returns the P field if non-nil, zero value otherwise.

### GetPOk

`func (o *UserDataStreamEventsResponse) GetPOk() (*[]MarginCallPInner, bool)`

GetPOk returns a tuple with the P field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP

`func (o *UserDataStreamEventsResponse) SetP(v []MarginCallPInner)`

SetP sets P field to given value.

### HasP

`func (o *UserDataStreamEventsResponse) HasP() bool`

HasP returns a boolean if a field has been set.

### GetO

`func (o *UserDataStreamEventsResponse) GetO() OrderTradeUpdateO`

GetO returns the O field if non-nil, zero value otherwise.

### GetOOk

`func (o *UserDataStreamEventsResponse) GetOOk() (*OrderTradeUpdateO, bool)`

GetOOk returns a tuple with the O field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetO

`func (o *UserDataStreamEventsResponse) SetO(v OrderTradeUpdateO)`

SetO sets O field to given value.

### HasO

`func (o *UserDataStreamEventsResponse) HasO() bool`

HasO returns a boolean if a field has been set.

### GetSu

`func (o *UserDataStreamEventsResponse) GetSu() StrategyUpdateSu`

GetSu returns the Su field if non-nil, zero value otherwise.

### GetSuOk

`func (o *UserDataStreamEventsResponse) GetSuOk() (*StrategyUpdateSu, bool)`

GetSuOk returns a tuple with the Su field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSu

`func (o *UserDataStreamEventsResponse) SetSu(v StrategyUpdateSu)`

SetSu sets Su field to given value.

### HasSu

`func (o *UserDataStreamEventsResponse) HasSu() bool`

HasSu returns a boolean if a field has been set.

### GetListenKey

`func (o *UserDataStreamEventsResponse) GetListenKey() string`

GetListenKey returns the ListenKey field if non-nil, zero value otherwise.

### GetListenKeyOk

`func (o *UserDataStreamEventsResponse) GetListenKeyOk() (*string, bool)`

GetListenKeyOk returns a tuple with the ListenKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenKey

`func (o *UserDataStreamEventsResponse) SetListenKey(v string)`

SetListenKey sets ListenKey field to given value.

### HasListenKey

`func (o *UserDataStreamEventsResponse) HasListenKey() bool`

HasListenKey returns a boolean if a field has been set.


[[Back to README]](../README.md)


