# AccountApiTradingStatusResponseData

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**IsLocked** | Pointer to **bool** |  | [optional] 
**PlannedRecoverTime** | Pointer to **int64** |  | [optional] 
**TriggerCondition** | Pointer to [**AccountApiTradingStatusResponseDataTriggerCondition**](AccountApiTradingStatusResponseDataTriggerCondition.md) |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewAccountApiTradingStatusResponseData

`func NewAccountApiTradingStatusResponseData() *AccountApiTradingStatusResponseData`

NewAccountApiTradingStatusResponseData instantiates a new AccountApiTradingStatusResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountApiTradingStatusResponseDataWithDefaults

`func NewAccountApiTradingStatusResponseDataWithDefaults() *AccountApiTradingStatusResponseData`

NewAccountApiTradingStatusResponseDataWithDefaults instantiates a new AccountApiTradingStatusResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsLocked

`func (o *AccountApiTradingStatusResponseData) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *AccountApiTradingStatusResponseData) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *AccountApiTradingStatusResponseData) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *AccountApiTradingStatusResponseData) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetPlannedRecoverTime

`func (o *AccountApiTradingStatusResponseData) GetPlannedRecoverTime() int64`

GetPlannedRecoverTime returns the PlannedRecoverTime field if non-nil, zero value otherwise.

### GetPlannedRecoverTimeOk

`func (o *AccountApiTradingStatusResponseData) GetPlannedRecoverTimeOk() (*int64, bool)`

GetPlannedRecoverTimeOk returns a tuple with the PlannedRecoverTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedRecoverTime

`func (o *AccountApiTradingStatusResponseData) SetPlannedRecoverTime(v int64)`

SetPlannedRecoverTime sets PlannedRecoverTime field to given value.

### HasPlannedRecoverTime

`func (o *AccountApiTradingStatusResponseData) HasPlannedRecoverTime() bool`

HasPlannedRecoverTime returns a boolean if a field has been set.

### GetTriggerCondition

`func (o *AccountApiTradingStatusResponseData) GetTriggerCondition() AccountApiTradingStatusResponseDataTriggerCondition`

GetTriggerCondition returns the TriggerCondition field if non-nil, zero value otherwise.

### GetTriggerConditionOk

`func (o *AccountApiTradingStatusResponseData) GetTriggerConditionOk() (*AccountApiTradingStatusResponseDataTriggerCondition, bool)`

GetTriggerConditionOk returns a tuple with the TriggerCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerCondition

`func (o *AccountApiTradingStatusResponseData) SetTriggerCondition(v AccountApiTradingStatusResponseDataTriggerCondition)`

SetTriggerCondition sets TriggerCondition field to given value.

### HasTriggerCondition

`func (o *AccountApiTradingStatusResponseData) HasTriggerCondition() bool`

HasTriggerCondition returns a boolean if a field has been set.

### GetUpdateTime

`func (o *AccountApiTradingStatusResponseData) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *AccountApiTradingStatusResponseData) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *AccountApiTradingStatusResponseData) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *AccountApiTradingStatusResponseData) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


