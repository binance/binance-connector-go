/*
Simple Earn REST API

Earn rewards by subscribing to flexible or locked Simple Earn products.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetYieldArenaActivitiesResponseActivitiesInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetYieldArenaActivitiesResponseActivitiesInner{}

// GetYieldArenaActivitiesResponseActivitiesInner struct for GetYieldArenaActivitiesResponseActivitiesInner
type GetYieldArenaActivitiesResponseActivitiesInner struct {
	// Activity ID.
	ActivityId *int64 `json:"activityId,omitempty"`
	// Activity category: `AIRDROP`, `LEADERBOARD`, or `EVENT`.
	ActivityType *string `json:"activityType,omitempty"`
	// Activity title, localized via the `lang` header.
	Title *string `json:"title,omitempty"`
	// Activity description, localized via the `lang` header.
	Description *string `json:"description,omitempty"`
	// USD value of the reward pool.
	RewardPoolInUsd *string `json:"rewardPoolInUsd,omitempty"`
	// Reward token symbols (e.g. `[\"BNB\"]`); may be empty.
	RewardToken []string `json:"rewardToken,omitempty"`
	// Web URL to the activity landing page.
	RedirectUrl *string `json:"redirectUrl,omitempty"`
	// Activity start time in milliseconds; may be null for activities that are immediately effective.
	StartTime *int64 `json:"startTime,omitempty"`
	// Activity end time in milliseconds; may be null for activities with no fixed end.
	EndTime              *int64 `json:"endTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetYieldArenaActivitiesResponseActivitiesInner GetYieldArenaActivitiesResponseActivitiesInner

// NewGetYieldArenaActivitiesResponseActivitiesInner instantiates a new GetYieldArenaActivitiesResponseActivitiesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetYieldArenaActivitiesResponseActivitiesInner() *GetYieldArenaActivitiesResponseActivitiesInner {
	this := GetYieldArenaActivitiesResponseActivitiesInner{}
	return &this
}

// NewGetYieldArenaActivitiesResponseActivitiesInnerWithDefaults instantiates a new GetYieldArenaActivitiesResponseActivitiesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetYieldArenaActivitiesResponseActivitiesInnerWithDefaults() *GetYieldArenaActivitiesResponseActivitiesInner {
	this := GetYieldArenaActivitiesResponseActivitiesInner{}
	return &this
}

// GetActivityId returns the ActivityId field value if set, zero value otherwise.
func (o *GetYieldArenaActivitiesResponseActivitiesInner) GetActivityId() int64 {
	if o == nil || common.IsNil(o.ActivityId) {
		var ret int64
		return ret
	}
	return *o.ActivityId
}

// GetActivityIdOk returns a tuple with the ActivityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetYieldArenaActivitiesResponseActivitiesInner) GetActivityIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ActivityId) {
		return nil, false
	}
	return o.ActivityId, true
}

// HasActivityId returns a boolean if a field has been set.
func (o *GetYieldArenaActivitiesResponseActivitiesInner) HasActivityId() bool {
	if o != nil && !common.IsNil(o.ActivityId) {
		return true
	}

	return false
}

// SetActivityId gets a reference to the given int64 and assigns it to the ActivityId field.
func (o *GetYieldArenaActivitiesResponseActivitiesInner) SetActivityId(v int64) {
	o.ActivityId = &v
}

// GetActivityType returns the ActivityType field value if set, zero value otherwise.
func (o *GetYieldArenaActivitiesResponseActivitiesInner) GetActivityType() string {
	if o == nil || common.IsNil(o.ActivityType) {
		var ret string
		return ret
	}
	return *o.ActivityType
}

// GetActivityTypeOk returns a tuple with the ActivityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetYieldArenaActivitiesResponseActivitiesInner) GetActivityTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ActivityType) {
		return nil, false
	}
	return o.ActivityType, true
}

// HasActivityType returns a boolean if a field has been set.
func (o *GetYieldArenaActivitiesResponseActivitiesInner) HasActivityType() bool {
	if o != nil && !common.IsNil(o.ActivityType) {
		return true
	}

	return false
}

// SetActivityType gets a reference to the given string and assigns it to the ActivityType field.
func (o *GetYieldArenaActivitiesResponseActivitiesInner) SetActivityType(v string) {
	o.ActivityType = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *GetYieldArenaActivitiesResponseActivitiesInner) GetTitle() string {
	if o == nil || common.IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetYieldArenaActivitiesResponseActivitiesInner) GetTitleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *GetYieldArenaActivitiesResponseActivitiesInner) HasTitle() bool {
	if o != nil && !common.IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *GetYieldArenaActivitiesResponseActivitiesInner) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetYieldArenaActivitiesResponseActivitiesInner) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetYieldArenaActivitiesResponseActivitiesInner) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GetYieldArenaActivitiesResponseActivitiesInner) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetYieldArenaActivitiesResponseActivitiesInner) SetDescription(v string) {
	o.Description = &v
}

// GetRewardPoolInUsd returns the RewardPoolInUsd field value if set, zero value otherwise.
func (o *GetYieldArenaActivitiesResponseActivitiesInner) GetRewardPoolInUsd() string {
	if o == nil || common.IsNil(o.RewardPoolInUsd) {
		var ret string
		return ret
	}
	return *o.RewardPoolInUsd
}

// GetRewardPoolInUsdOk returns a tuple with the RewardPoolInUsd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetYieldArenaActivitiesResponseActivitiesInner) GetRewardPoolInUsdOk() (*string, bool) {
	if o == nil || common.IsNil(o.RewardPoolInUsd) {
		return nil, false
	}
	return o.RewardPoolInUsd, true
}

// HasRewardPoolInUsd returns a boolean if a field has been set.
func (o *GetYieldArenaActivitiesResponseActivitiesInner) HasRewardPoolInUsd() bool {
	if o != nil && !common.IsNil(o.RewardPoolInUsd) {
		return true
	}

	return false
}

// SetRewardPoolInUsd gets a reference to the given string and assigns it to the RewardPoolInUsd field.
func (o *GetYieldArenaActivitiesResponseActivitiesInner) SetRewardPoolInUsd(v string) {
	o.RewardPoolInUsd = &v
}

// GetRewardToken returns the RewardToken field value if set, zero value otherwise.
func (o *GetYieldArenaActivitiesResponseActivitiesInner) GetRewardToken() []string {
	if o == nil || common.IsNil(o.RewardToken) {
		var ret []string
		return ret
	}
	return o.RewardToken
}

// GetRewardTokenOk returns a tuple with the RewardToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetYieldArenaActivitiesResponseActivitiesInner) GetRewardTokenOk() ([]string, bool) {
	if o == nil || common.IsNil(o.RewardToken) {
		return nil, false
	}
	return o.RewardToken, true
}

// HasRewardToken returns a boolean if a field has been set.
func (o *GetYieldArenaActivitiesResponseActivitiesInner) HasRewardToken() bool {
	if o != nil && !common.IsNil(o.RewardToken) {
		return true
	}

	return false
}

// SetRewardToken gets a reference to the given []string and assigns it to the RewardToken field.
func (o *GetYieldArenaActivitiesResponseActivitiesInner) SetRewardToken(v []string) {
	o.RewardToken = v
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise.
func (o *GetYieldArenaActivitiesResponseActivitiesInner) GetRedirectUrl() string {
	if o == nil || common.IsNil(o.RedirectUrl) {
		var ret string
		return ret
	}
	return *o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetYieldArenaActivitiesResponseActivitiesInner) GetRedirectUrlOk() (*string, bool) {
	if o == nil || common.IsNil(o.RedirectUrl) {
		return nil, false
	}
	return o.RedirectUrl, true
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *GetYieldArenaActivitiesResponseActivitiesInner) HasRedirectUrl() bool {
	if o != nil && !common.IsNil(o.RedirectUrl) {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given string and assigns it to the RedirectUrl field.
func (o *GetYieldArenaActivitiesResponseActivitiesInner) SetRedirectUrl(v string) {
	o.RedirectUrl = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *GetYieldArenaActivitiesResponseActivitiesInner) GetStartTime() int64 {
	if o == nil || common.IsNil(o.StartTime) {
		var ret int64
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetYieldArenaActivitiesResponseActivitiesInner) GetStartTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *GetYieldArenaActivitiesResponseActivitiesInner) HasStartTime() bool {
	if o != nil && !common.IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given int64 and assigns it to the StartTime field.
func (o *GetYieldArenaActivitiesResponseActivitiesInner) SetStartTime(v int64) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *GetYieldArenaActivitiesResponseActivitiesInner) GetEndTime() int64 {
	if o == nil || common.IsNil(o.EndTime) {
		var ret int64
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetYieldArenaActivitiesResponseActivitiesInner) GetEndTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *GetYieldArenaActivitiesResponseActivitiesInner) HasEndTime() bool {
	if o != nil && !common.IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given int64 and assigns it to the EndTime field.
func (o *GetYieldArenaActivitiesResponseActivitiesInner) SetEndTime(v int64) {
	o.EndTime = &v
}

func (o GetYieldArenaActivitiesResponseActivitiesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetYieldArenaActivitiesResponseActivitiesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ActivityId) {
		toSerialize["activityId"] = o.ActivityId
	}
	if !common.IsNil(o.ActivityType) {
		toSerialize["activityType"] = o.ActivityType
	}
	if !common.IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !common.IsNil(o.RewardPoolInUsd) {
		toSerialize["rewardPoolInUsd"] = o.RewardPoolInUsd
	}
	if !common.IsNil(o.RewardToken) {
		toSerialize["rewardToken"] = o.RewardToken
	}
	if !common.IsNil(o.RedirectUrl) {
		toSerialize["redirectUrl"] = o.RedirectUrl
	}
	if !common.IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !common.IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetYieldArenaActivitiesResponseActivitiesInner) UnmarshalJSON(data []byte) (err error) {
	varGetYieldArenaActivitiesResponseActivitiesInner := _GetYieldArenaActivitiesResponseActivitiesInner{}

	err = json.Unmarshal(data, &varGetYieldArenaActivitiesResponseActivitiesInner)

	if err != nil {
		return err
	}

	*o = GetYieldArenaActivitiesResponseActivitiesInner(varGetYieldArenaActivitiesResponseActivitiesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "activityId")
		delete(additionalProperties, "activityType")
		delete(additionalProperties, "title")
		delete(additionalProperties, "description")
		delete(additionalProperties, "rewardPoolInUsd")
		delete(additionalProperties, "rewardToken")
		delete(additionalProperties, "redirectUrl")
		delete(additionalProperties, "startTime")
		delete(additionalProperties, "endTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetYieldArenaActivitiesResponseActivitiesInner struct {
	value *GetYieldArenaActivitiesResponseActivitiesInner
	isSet bool
}

func (v NullableGetYieldArenaActivitiesResponseActivitiesInner) Get() *GetYieldArenaActivitiesResponseActivitiesInner {
	return v.value
}

func (v *NullableGetYieldArenaActivitiesResponseActivitiesInner) Set(val *GetYieldArenaActivitiesResponseActivitiesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetYieldArenaActivitiesResponseActivitiesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetYieldArenaActivitiesResponseActivitiesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetYieldArenaActivitiesResponseActivitiesInner(val *GetYieldArenaActivitiesResponseActivitiesInner) *NullableGetYieldArenaActivitiesResponseActivitiesInner {
	return &NullableGetYieldArenaActivitiesResponseActivitiesInner{value: val, isSet: true}
}

func (v NullableGetYieldArenaActivitiesResponseActivitiesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetYieldArenaActivitiesResponseActivitiesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
