# GetYieldArenaActivitiesResponseActivitiesInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**ActivityId** | Pointer to **int64** | Activity ID. | [optional] 
**ActivityType** | Pointer to **string** | Activity category: &#x60;AIRDROP&#x60;, &#x60;LEADERBOARD&#x60;, or &#x60;EVENT&#x60;. | [optional] 
**Title** | Pointer to **string** | Activity title, localized via the &#x60;lang&#x60; header. | [optional] 
**Description** | Pointer to **string** | Activity description, localized via the &#x60;lang&#x60; header. | [optional] 
**RewardPoolInUsd** | Pointer to **string** | USD value of the reward pool. | [optional] 
**RewardToken** | Pointer to **[]string** | Reward token symbols (e.g. &#x60;[\&quot;BNB\&quot;]&#x60;); may be empty. | [optional] 
**RedirectUrl** | Pointer to **string** | Web URL to the activity landing page. | [optional] 
**StartTime** | Pointer to **int64** | Activity start time in milliseconds; may be null for activities that are immediately effective. | [optional] 
**EndTime** | Pointer to **int64** | Activity end time in milliseconds; may be null for activities with no fixed end. | [optional] 

## Methods

### NewGetYieldArenaActivitiesResponseActivitiesInner

`func NewGetYieldArenaActivitiesResponseActivitiesInner() *GetYieldArenaActivitiesResponseActivitiesInner`

NewGetYieldArenaActivitiesResponseActivitiesInner instantiates a new GetYieldArenaActivitiesResponseActivitiesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetYieldArenaActivitiesResponseActivitiesInnerWithDefaults

`func NewGetYieldArenaActivitiesResponseActivitiesInnerWithDefaults() *GetYieldArenaActivitiesResponseActivitiesInner`

NewGetYieldArenaActivitiesResponseActivitiesInnerWithDefaults instantiates a new GetYieldArenaActivitiesResponseActivitiesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityId

`func (o *GetYieldArenaActivitiesResponseActivitiesInner) GetActivityId() int64`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *GetYieldArenaActivitiesResponseActivitiesInner) GetActivityIdOk() (*int64, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *GetYieldArenaActivitiesResponseActivitiesInner) SetActivityId(v int64)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *GetYieldArenaActivitiesResponseActivitiesInner) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### GetActivityType

`func (o *GetYieldArenaActivitiesResponseActivitiesInner) GetActivityType() string`

GetActivityType returns the ActivityType field if non-nil, zero value otherwise.

### GetActivityTypeOk

`func (o *GetYieldArenaActivitiesResponseActivitiesInner) GetActivityTypeOk() (*string, bool)`

GetActivityTypeOk returns a tuple with the ActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityType

`func (o *GetYieldArenaActivitiesResponseActivitiesInner) SetActivityType(v string)`

SetActivityType sets ActivityType field to given value.

### HasActivityType

`func (o *GetYieldArenaActivitiesResponseActivitiesInner) HasActivityType() bool`

HasActivityType returns a boolean if a field has been set.

### GetTitle

`func (o *GetYieldArenaActivitiesResponseActivitiesInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetYieldArenaActivitiesResponseActivitiesInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetYieldArenaActivitiesResponseActivitiesInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GetYieldArenaActivitiesResponseActivitiesInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *GetYieldArenaActivitiesResponseActivitiesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetYieldArenaActivitiesResponseActivitiesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetYieldArenaActivitiesResponseActivitiesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetYieldArenaActivitiesResponseActivitiesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRewardPoolInUsd

`func (o *GetYieldArenaActivitiesResponseActivitiesInner) GetRewardPoolInUsd() string`

GetRewardPoolInUsd returns the RewardPoolInUsd field if non-nil, zero value otherwise.

### GetRewardPoolInUsdOk

`func (o *GetYieldArenaActivitiesResponseActivitiesInner) GetRewardPoolInUsdOk() (*string, bool)`

GetRewardPoolInUsdOk returns a tuple with the RewardPoolInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardPoolInUsd

`func (o *GetYieldArenaActivitiesResponseActivitiesInner) SetRewardPoolInUsd(v string)`

SetRewardPoolInUsd sets RewardPoolInUsd field to given value.

### HasRewardPoolInUsd

`func (o *GetYieldArenaActivitiesResponseActivitiesInner) HasRewardPoolInUsd() bool`

HasRewardPoolInUsd returns a boolean if a field has been set.

### GetRewardToken

`func (o *GetYieldArenaActivitiesResponseActivitiesInner) GetRewardToken() []string`

GetRewardToken returns the RewardToken field if non-nil, zero value otherwise.

### GetRewardTokenOk

`func (o *GetYieldArenaActivitiesResponseActivitiesInner) GetRewardTokenOk() (*[]string, bool)`

GetRewardTokenOk returns a tuple with the RewardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardToken

`func (o *GetYieldArenaActivitiesResponseActivitiesInner) SetRewardToken(v []string)`

SetRewardToken sets RewardToken field to given value.

### HasRewardToken

`func (o *GetYieldArenaActivitiesResponseActivitiesInner) HasRewardToken() bool`

HasRewardToken returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *GetYieldArenaActivitiesResponseActivitiesInner) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *GetYieldArenaActivitiesResponseActivitiesInner) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *GetYieldArenaActivitiesResponseActivitiesInner) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *GetYieldArenaActivitiesResponseActivitiesInner) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetStartTime

`func (o *GetYieldArenaActivitiesResponseActivitiesInner) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *GetYieldArenaActivitiesResponseActivitiesInner) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *GetYieldArenaActivitiesResponseActivitiesInner) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *GetYieldArenaActivitiesResponseActivitiesInner) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *GetYieldArenaActivitiesResponseActivitiesInner) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *GetYieldArenaActivitiesResponseActivitiesInner) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *GetYieldArenaActivitiesResponseActivitiesInner) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *GetYieldArenaActivitiesResponseActivitiesInner) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


