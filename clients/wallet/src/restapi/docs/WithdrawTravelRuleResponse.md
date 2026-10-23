# WithdrawTravelRuleResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**TrId** | Pointer to **int64** | The travel rule record Id | [optional] 
**Accepted** | Pointer to **bool** | Whether the withdraw request is accepted | [optional] 
**Info** | Pointer to **string** | The detailed infomation of the withdrawal result. | [optional] 

## Methods

### NewWithdrawTravelRuleResponse

`func NewWithdrawTravelRuleResponse() *WithdrawTravelRuleResponse`

NewWithdrawTravelRuleResponse instantiates a new WithdrawTravelRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWithdrawTravelRuleResponseWithDefaults

`func NewWithdrawTravelRuleResponseWithDefaults() *WithdrawTravelRuleResponse`

NewWithdrawTravelRuleResponseWithDefaults instantiates a new WithdrawTravelRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrId

`func (o *WithdrawTravelRuleResponse) GetTrId() int64`

GetTrId returns the TrId field if non-nil, zero value otherwise.

### GetTrIdOk

`func (o *WithdrawTravelRuleResponse) GetTrIdOk() (*int64, bool)`

GetTrIdOk returns a tuple with the TrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrId

`func (o *WithdrawTravelRuleResponse) SetTrId(v int64)`

SetTrId sets TrId field to given value.

### HasTrId

`func (o *WithdrawTravelRuleResponse) HasTrId() bool`

HasTrId returns a boolean if a field has been set.

### GetAccepted

`func (o *WithdrawTravelRuleResponse) GetAccepted() bool`

GetAccepted returns the Accepted field if non-nil, zero value otherwise.

### GetAcceptedOk

`func (o *WithdrawTravelRuleResponse) GetAcceptedOk() (*bool, bool)`

GetAcceptedOk returns a tuple with the Accepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepted

`func (o *WithdrawTravelRuleResponse) SetAccepted(v bool)`

SetAccepted sets Accepted field to given value.

### HasAccepted

`func (o *WithdrawTravelRuleResponse) HasAccepted() bool`

HasAccepted returns a boolean if a field has been set.

### GetInfo

`func (o *WithdrawTravelRuleResponse) GetInfo() string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *WithdrawTravelRuleResponse) GetInfoOk() (*string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *WithdrawTravelRuleResponse) SetInfo(v string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *WithdrawTravelRuleResponse) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to README]](../README.md)


