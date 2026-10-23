# RewardClaimFailedResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**PushId** | Pointer to **string** | Unique ID of this push. Use it to remove duplicates. | [optional] 
**Amount** | Pointer to **string** | Amount of the **first failed claim**. Rounded down to 2 decimals; shown as &#x60;&lt;0.01&#x60; if smaller. | [optional] 
**Outcome** | Pointer to **string** | Outcome name of the **first failed claim**, such as &#x60;Yes&#x60; or &#x60;No&#x60;; empty if unavailable. | [optional] 

## Methods

### NewRewardClaimFailedResponse

`func NewRewardClaimFailedResponse() *RewardClaimFailedResponse`

NewRewardClaimFailedResponse instantiates a new RewardClaimFailedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardClaimFailedResponseWithDefaults

`func NewRewardClaimFailedResponseWithDefaults() *RewardClaimFailedResponse`

NewRewardClaimFailedResponseWithDefaults instantiates a new RewardClaimFailedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPushId

`func (o *RewardClaimFailedResponse) GetPushId() string`

GetPushId returns the PushId field if non-nil, zero value otherwise.

### GetPushIdOk

`func (o *RewardClaimFailedResponse) GetPushIdOk() (*string, bool)`

GetPushIdOk returns a tuple with the PushId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushId

`func (o *RewardClaimFailedResponse) SetPushId(v string)`

SetPushId sets PushId field to given value.

### HasPushId

`func (o *RewardClaimFailedResponse) HasPushId() bool`

HasPushId returns a boolean if a field has been set.

### GetAmount

`func (o *RewardClaimFailedResponse) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RewardClaimFailedResponse) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RewardClaimFailedResponse) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *RewardClaimFailedResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetOutcome

`func (o *RewardClaimFailedResponse) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *RewardClaimFailedResponse) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *RewardClaimFailedResponse) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *RewardClaimFailedResponse) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.


[[Back to README]](../README.md)


