# UserDataStreamEventsResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **int64** |  | [optional] 
**B** | Pointer to [**[]AccountUpdateBInner**](AccountUpdateBInner.md) |  | [optional] 
**G** | Pointer to [**[]AccountUpdateGInner**](AccountUpdateGInner.md) |  | [optional] 
**P** | Pointer to [**[]AccountUpdatePInner**](AccountUpdatePInner.md) |  | [optional] 
**Uid** | Pointer to **int64** |  | [optional] 
**O** | Pointer to [**[]OrderTradeUpdateOInner**](OrderTradeUpdateOInner.md) |  | [optional] 
**S** | Pointer to **string** |  | [optional] 
**Mb** | Pointer to **string** |  | [optional] 
**Mm** | Pointer to **string** |  | [optional] 

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

### GetB

`func (o *UserDataStreamEventsResponse) GetB() []AccountUpdateBInner`

GetB returns the B field if non-nil, zero value otherwise.

### GetBOk

`func (o *UserDataStreamEventsResponse) GetBOk() (*[]AccountUpdateBInner, bool)`

GetBOk returns a tuple with the B field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB

`func (o *UserDataStreamEventsResponse) SetB(v []AccountUpdateBInner)`

SetB sets B field to given value.

### HasB

`func (o *UserDataStreamEventsResponse) HasB() bool`

HasB returns a boolean if a field has been set.

### GetG

`func (o *UserDataStreamEventsResponse) GetG() []AccountUpdateGInner`

GetG returns the G field if non-nil, zero value otherwise.

### GetGOk

`func (o *UserDataStreamEventsResponse) GetGOk() (*[]AccountUpdateGInner, bool)`

GetGOk returns a tuple with the G field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetG

`func (o *UserDataStreamEventsResponse) SetG(v []AccountUpdateGInner)`

SetG sets G field to given value.

### HasG

`func (o *UserDataStreamEventsResponse) HasG() bool`

HasG returns a boolean if a field has been set.

### GetP

`func (o *UserDataStreamEventsResponse) GetP() []AccountUpdatePInner`

GetP returns the P field if non-nil, zero value otherwise.

### GetPOk

`func (o *UserDataStreamEventsResponse) GetPOk() (*[]AccountUpdatePInner, bool)`

GetPOk returns a tuple with the P field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP

`func (o *UserDataStreamEventsResponse) SetP(v []AccountUpdatePInner)`

SetP sets P field to given value.

### HasP

`func (o *UserDataStreamEventsResponse) HasP() bool`

HasP returns a boolean if a field has been set.

### GetUid

`func (o *UserDataStreamEventsResponse) GetUid() int64`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *UserDataStreamEventsResponse) GetUidOk() (*int64, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *UserDataStreamEventsResponse) SetUid(v int64)`

SetUid sets Uid field to given value.

### HasUid

`func (o *UserDataStreamEventsResponse) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetO

`func (o *UserDataStreamEventsResponse) GetO() []OrderTradeUpdateOInner`

GetO returns the O field if non-nil, zero value otherwise.

### GetOOk

`func (o *UserDataStreamEventsResponse) GetOOk() (*[]OrderTradeUpdateOInner, bool)`

GetOOk returns a tuple with the O field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetO

`func (o *UserDataStreamEventsResponse) SetO(v []OrderTradeUpdateOInner)`

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


[[Back to README]](../README.md)


