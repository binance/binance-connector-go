# UserDataStreamEventsResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **string** |  | [optional] 
**T** | Pointer to **int64** |  | [optional] 
**M** | Pointer to **string** |  | [optional] 
**B** | Pointer to [**[]BalancePositionUpdateBInner**](BalancePositionUpdateBInner.md) |  | [optional] 
**P** | Pointer to [**[]BalancePositionUpdatePInner**](BalancePositionUpdatePInner.md) |  | [optional] 
**G** | Pointer to [**[]GreekUpdateGInner**](GreekUpdateGInner.md) |  | [optional] 
**O** | Pointer to [**OrderTradeUpdateO**](OrderTradeUpdateO.md) |  | [optional] 
**S** | Pointer to **string** |  | [optional] 
**Mb** | Pointer to **string** |  | [optional] 
**Mm** | Pointer to **string** |  | [optional] 
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

### GetM

`func (o *UserDataStreamEventsResponse) GetM() string`

GetM returns the M field if non-nil, zero value otherwise.

### GetMOk

`func (o *UserDataStreamEventsResponse) GetMOk() (*string, bool)`

GetMOk returns a tuple with the M field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetM

`func (o *UserDataStreamEventsResponse) SetM(v string)`

SetM sets M field to given value.

### HasM

`func (o *UserDataStreamEventsResponse) HasM() bool`

HasM returns a boolean if a field has been set.

### GetB

`func (o *UserDataStreamEventsResponse) GetB() []BalancePositionUpdateBInner`

GetB returns the B field if non-nil, zero value otherwise.

### GetBOk

`func (o *UserDataStreamEventsResponse) GetBOk() (*[]BalancePositionUpdateBInner, bool)`

GetBOk returns a tuple with the B field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB

`func (o *UserDataStreamEventsResponse) SetB(v []BalancePositionUpdateBInner)`

SetB sets B field to given value.

### HasB

`func (o *UserDataStreamEventsResponse) HasB() bool`

HasB returns a boolean if a field has been set.

### GetP

`func (o *UserDataStreamEventsResponse) GetP() []BalancePositionUpdatePInner`

GetP returns the P field if non-nil, zero value otherwise.

### GetPOk

`func (o *UserDataStreamEventsResponse) GetPOk() (*[]BalancePositionUpdatePInner, bool)`

GetPOk returns a tuple with the P field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP

`func (o *UserDataStreamEventsResponse) SetP(v []BalancePositionUpdatePInner)`

SetP sets P field to given value.

### HasP

`func (o *UserDataStreamEventsResponse) HasP() bool`

HasP returns a boolean if a field has been set.

### GetG

`func (o *UserDataStreamEventsResponse) GetG() []GreekUpdateGInner`

GetG returns the G field if non-nil, zero value otherwise.

### GetGOk

`func (o *UserDataStreamEventsResponse) GetGOk() (*[]GreekUpdateGInner, bool)`

GetGOk returns a tuple with the G field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetG

`func (o *UserDataStreamEventsResponse) SetG(v []GreekUpdateGInner)`

SetG sets G field to given value.

### HasG

`func (o *UserDataStreamEventsResponse) HasG() bool`

HasG returns a boolean if a field has been set.

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

### GetMb

`func (o *UserDataStreamEventsResponse) GetMb() string`

GetMb returns the Mb field if non-nil, zero value otherwise.

### GetMbOk

`func (o *UserDataStreamEventsResponse) GetMbOk() (*string, bool)`

GetMbOk returns a tuple with the Mb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMb

`func (o *UserDataStreamEventsResponse) SetMb(v string)`

SetMb sets Mb field to given value.

### HasMb

`func (o *UserDataStreamEventsResponse) HasMb() bool`

HasMb returns a boolean if a field has been set.

### GetMm

`func (o *UserDataStreamEventsResponse) GetMm() string`

GetMm returns the Mm field if non-nil, zero value otherwise.

### GetMmOk

`func (o *UserDataStreamEventsResponse) GetMmOk() (*string, bool)`

GetMmOk returns a tuple with the Mm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMm

`func (o *UserDataStreamEventsResponse) SetMm(v string)`

SetMm sets Mm field to given value.

### HasMm

`func (o *UserDataStreamEventsResponse) HasMm() bool`

HasMm returns a boolean if a field has been set.

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


