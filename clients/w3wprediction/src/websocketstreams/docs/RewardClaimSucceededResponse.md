# RewardClaimSucceededResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**PushId** | Pointer to **string** | Unique push ID in the form &#x60;pm_&lt;refId&gt;_&lt;scenarioCode&gt;_&lt;8hex&gt;&#x60;, present in every topic; &#x60;refId&#x60; is the claim &#x60;batchId&#x60;. Use it for client-side deduplication. | [optional] 
**Amount** | Pointer to **string** | Total claimed USDT across the batch, &#x60;SUM(dataItems[].amount)&#x60;. Truncated to 2 decimals (towards zero); rendered as &#x60;&lt;0.01&#x60; when the value is below 0.01. | [optional] 

## Methods

### NewRewardClaimSucceededResponse

`func NewRewardClaimSucceededResponse() *RewardClaimSucceededResponse`

NewRewardClaimSucceededResponse instantiates a new RewardClaimSucceededResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardClaimSucceededResponseWithDefaults

`func NewRewardClaimSucceededResponseWithDefaults() *RewardClaimSucceededResponse`

NewRewardClaimSucceededResponseWithDefaults instantiates a new RewardClaimSucceededResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPushId

`func (o *RewardClaimSucceededResponse) GetPushId() string`

GetPushId returns the PushId field if non-nil, zero value otherwise.

### GetPushIdOk

`func (o *RewardClaimSucceededResponse) GetPushIdOk() (*string, bool)`

GetPushIdOk returns a tuple with the PushId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushId

`func (o *RewardClaimSucceededResponse) SetPushId(v string)`

SetPushId sets PushId field to given value.

### HasPushId

`func (o *RewardClaimSucceededResponse) HasPushId() bool`

HasPushId returns a boolean if a field has been set.

### GetAmount

`func (o *RewardClaimSucceededResponse) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RewardClaimSucceededResponse) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RewardClaimSucceededResponse) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *RewardClaimSucceededResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to README]](../README.md)


