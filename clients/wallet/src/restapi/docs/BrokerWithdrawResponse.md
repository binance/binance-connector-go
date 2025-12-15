# BrokerWithdrawResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**TrId** | Pointer to **int64** |  | [optional] 
**Accpted** | Pointer to **bool** |  | [optional] 
**Info** | Pointer to **string** |  | [optional] 

## Methods

### NewBrokerWithdrawResponse

`func NewBrokerWithdrawResponse() *BrokerWithdrawResponse`

NewBrokerWithdrawResponse instantiates a new BrokerWithdrawResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrokerWithdrawResponseWithDefaults

`func NewBrokerWithdrawResponseWithDefaults() *BrokerWithdrawResponse`

NewBrokerWithdrawResponseWithDefaults instantiates a new BrokerWithdrawResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrId

`func (o *BrokerWithdrawResponse) GetTrId() int64`

GetTrId returns the TrId field if non-nil, zero value otherwise.

### GetTrIdOk

`func (o *BrokerWithdrawResponse) GetTrIdOk() (*int64, bool)`

GetTrIdOk returns a tuple with the TrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrId

`func (o *BrokerWithdrawResponse) SetTrId(v int64)`

SetTrId sets TrId field to given value.

### HasTrId

`func (o *BrokerWithdrawResponse) HasTrId() bool`

HasTrId returns a boolean if a field has been set.

### GetAccpted

`func (o *BrokerWithdrawResponse) GetAccpted() bool`

GetAccpted returns the Accpted field if non-nil, zero value otherwise.

### GetAccptedOk

`func (o *BrokerWithdrawResponse) GetAccptedOk() (*bool, bool)`

GetAccptedOk returns a tuple with the Accpted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccpted

`func (o *BrokerWithdrawResponse) SetAccpted(v bool)`

SetAccpted sets Accpted field to given value.

### HasAccpted

`func (o *BrokerWithdrawResponse) HasAccpted() bool`

HasAccpted returns a boolean if a field has been set.

### GetInfo

`func (o *BrokerWithdrawResponse) GetInfo() string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *BrokerWithdrawResponse) GetInfoOk() (*string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *BrokerWithdrawResponse) SetInfo(v string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *BrokerWithdrawResponse) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to README]](../README.md)


