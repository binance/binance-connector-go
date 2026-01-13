# UserDataStreamEventsResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **string** |  | [optional] 
**T** | Pointer to **int64** |  | [optional] 
**Ac** | Pointer to [**AccountConfigUpdateAc**](AccountConfigUpdateAc.md) |  | [optional] 
**Ai** | Pointer to [**AccountConfigUpdateAi**](AccountConfigUpdateAi.md) |  | [optional] 
**O** | Pointer to [**OrderTradeUpdateO**](OrderTradeUpdateO.md) |  | [optional] 
**A** | Pointer to [**AccountUpdateA**](AccountUpdateA.md) |  | [optional] 
**Or** | Pointer to [**ConditionalOrderTriggerRejectOr**](ConditionalOrderTriggerRejectOr.md) |  | [optional] 
**Gu** | Pointer to [**GridUpdateGu**](GridUpdateGu.md) |  | [optional] 
**Cw** | Pointer to **string** |  | [optional] 
**P** | Pointer to **string** |  | [optional] 
**Su** | Pointer to [**StrategyUpdateSu**](StrategyUpdateSu.md) |  | [optional] 
**S** | Pointer to **string** |  | [optional] 
**Q** | Pointer to **string** |  | [optional] 
**M** | Pointer to **bool** |  | [optional] 
**C** | Pointer to **string** |  | [optional] 
**S** | Pointer to **string** |  | [optional] 
**L** | Pointer to **string** |  | [optional] 
**L** | Pointer to **string** |  | [optional] 
**T** | Pointer to **int64** |  | [optional] 
**I** | Pointer to **int64** |  | [optional] 
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

`func (o *UserDataStreamEventsResponse) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *UserDataStreamEventsResponse) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *UserDataStreamEventsResponse) SetE(v string)`

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

### GetAi

`func (o *UserDataStreamEventsResponse) GetAi() AccountConfigUpdateAi`

GetAi returns the Ai field if non-nil, zero value otherwise.

### GetAiOk

`func (o *UserDataStreamEventsResponse) GetAiOk() (*AccountConfigUpdateAi, bool)`

GetAiOk returns a tuple with the Ai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAi

`func (o *UserDataStreamEventsResponse) SetAi(v AccountConfigUpdateAi)`

SetAi sets Ai field to given value.

### HasAi

`func (o *UserDataStreamEventsResponse) HasAi() bool`

HasAi returns a boolean if a field has been set.

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

### GetOr

`func (o *UserDataStreamEventsResponse) GetOr() ConditionalOrderTriggerRejectOr`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *UserDataStreamEventsResponse) GetOrOk() (*ConditionalOrderTriggerRejectOr, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *UserDataStreamEventsResponse) SetOr(v ConditionalOrderTriggerRejectOr)`

SetOr sets Or field to given value.

### HasOr

`func (o *UserDataStreamEventsResponse) HasOr() bool`

HasOr returns a boolean if a field has been set.

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

`func (o *UserDataStreamEventsResponse) GetP() string`

GetP returns the P field if non-nil, zero value otherwise.

### GetPOk

`func (o *UserDataStreamEventsResponse) GetPOk() (*string, bool)`

GetPOk returns a tuple with the P field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP

`func (o *UserDataStreamEventsResponse) SetP(v string)`

SetP sets P field to given value.

### HasP

`func (o *UserDataStreamEventsResponse) HasP() bool`

HasP returns a boolean if a field has been set.

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

### GetS

`func (o *UserDataStreamEventsResponse) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *UserDataStreamEventsResponse) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *UserDataStreamEventsResponse) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *UserDataStreamEventsResponse) HasS() bool`

HasS returns a boolean if a field has been set.

### GetQ

`func (o *UserDataStreamEventsResponse) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *UserDataStreamEventsResponse) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *UserDataStreamEventsResponse) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *UserDataStreamEventsResponse) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetM

`func (o *UserDataStreamEventsResponse) GetM() bool`

GetM returns the M field if non-nil, zero value otherwise.

### GetMOk

`func (o *UserDataStreamEventsResponse) GetMOk() (*bool, bool)`

GetMOk returns a tuple with the M field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetM

`func (o *UserDataStreamEventsResponse) SetM(v bool)`

SetM sets M field to given value.

### HasM

`func (o *UserDataStreamEventsResponse) HasM() bool`

HasM returns a boolean if a field has been set.

### GetC

`func (o *UserDataStreamEventsResponse) GetC() string`

GetC returns the C field if non-nil, zero value otherwise.

### GetCOk

`func (o *UserDataStreamEventsResponse) GetCOk() (*string, bool)`

GetCOk returns a tuple with the C field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetC

`func (o *UserDataStreamEventsResponse) SetC(v string)`

SetC sets C field to given value.

### HasC

`func (o *UserDataStreamEventsResponse) HasC() bool`

HasC returns a boolean if a field has been set.

### GetS

`func (o *UserDataStreamEventsResponse) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *UserDataStreamEventsResponse) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *UserDataStreamEventsResponse) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *UserDataStreamEventsResponse) HasS() bool`

HasS returns a boolean if a field has been set.

### GetL

`func (o *UserDataStreamEventsResponse) GetL() string`

GetL returns the L field if non-nil, zero value otherwise.

### GetLOk

`func (o *UserDataStreamEventsResponse) GetLOk() (*string, bool)`

GetLOk returns a tuple with the L field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL

`func (o *UserDataStreamEventsResponse) SetL(v string)`

SetL sets L field to given value.

### HasL

`func (o *UserDataStreamEventsResponse) HasL() bool`

HasL returns a boolean if a field has been set.

### GetL

`func (o *UserDataStreamEventsResponse) GetL() string`

GetL returns the L field if non-nil, zero value otherwise.

### GetLOk

`func (o *UserDataStreamEventsResponse) GetLOk() (*string, bool)`

GetLOk returns a tuple with the L field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL

`func (o *UserDataStreamEventsResponse) SetL(v string)`

SetL sets L field to given value.

### HasL

`func (o *UserDataStreamEventsResponse) HasL() bool`

HasL returns a boolean if a field has been set.

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

### GetI

`func (o *UserDataStreamEventsResponse) GetI() int64`

GetI returns the I field if non-nil, zero value otherwise.

### GetIOk

`func (o *UserDataStreamEventsResponse) GetIOk() (*int64, bool)`

GetIOk returns a tuple with the I field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetI

`func (o *UserDataStreamEventsResponse) SetI(v int64)`

SetI sets I field to given value.

### HasI

`func (o *UserDataStreamEventsResponse) HasI() bool`

HasI returns a boolean if a field has been set.

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


