# RewardClaimFailedResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**PushId** | Pointer to **string** | Unique push ID in the form &#x60;pm_&lt;refId&gt;_&lt;scenarioCode&gt;_&lt;8hex&gt;&#x60;, present in every topic; &#x60;refId&#x60; is the claim &#x60;batchId&#x60;. Use it for client-side deduplication. | [optional] 
**Amount** | Pointer to **string** | Amount of the **first failed claim**, &#x60;dataItems[0].amount&#x60;. Truncated to 2 decimals (towards zero); rendered as &#x60;&lt;0.01&#x60; when the value is below 0.01. | [optional] 
**Outcome** | Pointer to **string** | Outcome name of the **first failed claim**, &#x60;dataItems[0].outcome.name&#x60; (e.g. &#x60;Yes&#x60; / &#x60;No&#x60; for binary markets). Passed through as-is, never truncated; may be an empty string when the data is missing. | [optional] 

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


