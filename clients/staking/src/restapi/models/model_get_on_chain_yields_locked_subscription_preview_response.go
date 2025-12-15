/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetOnChainYieldsLockedSubscriptionPreviewResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetOnChainYieldsLockedSubscriptionPreviewResponse{}

// GetOnChainYieldsLockedSubscriptionPreviewResponse struct for GetOnChainYieldsLockedSubscriptionPreviewResponse
type GetOnChainYieldsLockedSubscriptionPreviewResponse struct {
	RewardAsset          *string `json:"rewardAsset,omitempty"`
	TotalRewardAmt       *string `json:"totalRewardAmt,omitempty"`
	NextPay              *string `json:"nextPay,omitempty"`
	NextPayDate          *string `json:"nextPayDate,omitempty"`
	RewardsPayDate       *string `json:"rewardsPayDate,omitempty"`
	ValueDate            *string `json:"valueDate,omitempty"`
	RewardsEndDate       *string `json:"rewardsEndDate,omitempty"`
	DeliverDate          *string `json:"deliverDate,omitempty"`
	NextSubscriptionDate *string `json:"nextSubscriptionDate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetOnChainYieldsLockedSubscriptionPreviewResponse GetOnChainYieldsLockedSubscriptionPreviewResponse

// NewGetOnChainYieldsLockedSubscriptionPreviewResponse instantiates a new GetOnChainYieldsLockedSubscriptionPreviewResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOnChainYieldsLockedSubscriptionPreviewResponse() *GetOnChainYieldsLockedSubscriptionPreviewResponse {
	this := GetOnChainYieldsLockedSubscriptionPreviewResponse{}
	return &this
}

// NewGetOnChainYieldsLockedSubscriptionPreviewResponseWithDefaults instantiates a new GetOnChainYieldsLockedSubscriptionPreviewResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOnChainYieldsLockedSubscriptionPreviewResponseWithDefaults() *GetOnChainYieldsLockedSubscriptionPreviewResponse {
	this := GetOnChainYieldsLockedSubscriptionPreviewResponse{}
	return &this
}

// GetRewardAsset returns the RewardAsset field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedSubscriptionPreviewResponse) GetRewardAsset() string {
	if o == nil || common.IsNil(o.RewardAsset) {
		var ret string
		return ret
	}
	return *o.RewardAsset
}

// GetRewardAssetOk returns a tuple with the RewardAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedSubscriptionPreviewResponse) GetRewardAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.RewardAsset) {
		return nil, false
	}
	return o.RewardAsset, true
}

// HasRewardAsset returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedSubscriptionPreviewResponse) HasRewardAsset() bool {
	if o != nil && !common.IsNil(o.RewardAsset) {
		return true
	}

	return false
}

// SetRewardAsset gets a reference to the given string and assigns it to the RewardAsset field.
func (o *GetOnChainYieldsLockedSubscriptionPreviewResponse) SetRewardAsset(v string) {
	o.RewardAsset = &v
}

// GetTotalRewardAmt returns the TotalRewardAmt field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedSubscriptionPreviewResponse) GetTotalRewardAmt() string {
	if o == nil || common.IsNil(o.TotalRewardAmt) {
		var ret string
		return ret
	}
	return *o.TotalRewardAmt
}

// GetTotalRewardAmtOk returns a tuple with the TotalRewardAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedSubscriptionPreviewResponse) GetTotalRewardAmtOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalRewardAmt) {
		return nil, false
	}
	return o.TotalRewardAmt, true
}

// HasTotalRewardAmt returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedSubscriptionPreviewResponse) HasTotalRewardAmt() bool {
	if o != nil && !common.IsNil(o.TotalRewardAmt) {
		return true
	}

	return false
}

// SetTotalRewardAmt gets a reference to the given string and assigns it to the TotalRewardAmt field.
func (o *GetOnChainYieldsLockedSubscriptionPreviewResponse) SetTotalRewardAmt(v string) {
	o.TotalRewardAmt = &v
}

// GetNextPay returns the NextPay field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedSubscriptionPreviewResponse) GetNextPay() string {
	if o == nil || common.IsNil(o.NextPay) {
		var ret string
		return ret
	}
	return *o.NextPay
}

// GetNextPayOk returns a tuple with the NextPay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedSubscriptionPreviewResponse) GetNextPayOk() (*string, bool) {
	if o == nil || common.IsNil(o.NextPay) {
		return nil, false
	}
	return o.NextPay, true
}

// HasNextPay returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedSubscriptionPreviewResponse) HasNextPay() bool {
	if o != nil && !common.IsNil(o.NextPay) {
		return true
	}

	return false
}

// SetNextPay gets a reference to the given string and assigns it to the NextPay field.
func (o *GetOnChainYieldsLockedSubscriptionPreviewResponse) SetNextPay(v string) {
	o.NextPay = &v
}

// GetNextPayDate returns the NextPayDate field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedSubscriptionPreviewResponse) GetNextPayDate() string {
	if o == nil || common.IsNil(o.NextPayDate) {
		var ret string
		return ret
	}
	return *o.NextPayDate
}

// GetNextPayDateOk returns a tuple with the NextPayDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedSubscriptionPreviewResponse) GetNextPayDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.NextPayDate) {
		return nil, false
	}
	return o.NextPayDate, true
}

// HasNextPayDate returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedSubscriptionPreviewResponse) HasNextPayDate() bool {
	if o != nil && !common.IsNil(o.NextPayDate) {
		return true
	}

	return false
}

// SetNextPayDate gets a reference to the given string and assigns it to the NextPayDate field.
func (o *GetOnChainYieldsLockedSubscriptionPreviewResponse) SetNextPayDate(v string) {
	o.NextPayDate = &v
}

// GetRewardsPayDate returns the RewardsPayDate field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedSubscriptionPreviewResponse) GetRewardsPayDate() string {
	if o == nil || common.IsNil(o.RewardsPayDate) {
		var ret string
		return ret
	}
	return *o.RewardsPayDate
}

// GetRewardsPayDateOk returns a tuple with the RewardsPayDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedSubscriptionPreviewResponse) GetRewardsPayDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.RewardsPayDate) {
		return nil, false
	}
	return o.RewardsPayDate, true
}

// HasRewardsPayDate returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedSubscriptionPreviewResponse) HasRewardsPayDate() bool {
	if o != nil && !common.IsNil(o.RewardsPayDate) {
		return true
	}

	return false
}

// SetRewardsPayDate gets a reference to the given string and assigns it to the RewardsPayDate field.
func (o *GetOnChainYieldsLockedSubscriptionPreviewResponse) SetRewardsPayDate(v string) {
	o.RewardsPayDate = &v
}

// GetValueDate returns the ValueDate field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedSubscriptionPreviewResponse) GetValueDate() string {
	if o == nil || common.IsNil(o.ValueDate) {
		var ret string
		return ret
	}
	return *o.ValueDate
}

// GetValueDateOk returns a tuple with the ValueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedSubscriptionPreviewResponse) GetValueDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.ValueDate) {
		return nil, false
	}
	return o.ValueDate, true
}

// HasValueDate returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedSubscriptionPreviewResponse) HasValueDate() bool {
	if o != nil && !common.IsNil(o.ValueDate) {
		return true
	}

	return false
}

// SetValueDate gets a reference to the given string and assigns it to the ValueDate field.
func (o *GetOnChainYieldsLockedSubscriptionPreviewResponse) SetValueDate(v string) {
	o.ValueDate = &v
}

// GetRewardsEndDate returns the RewardsEndDate field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedSubscriptionPreviewResponse) GetRewardsEndDate() string {
	if o == nil || common.IsNil(o.RewardsEndDate) {
		var ret string
		return ret
	}
	return *o.RewardsEndDate
}

// GetRewardsEndDateOk returns a tuple with the RewardsEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedSubscriptionPreviewResponse) GetRewardsEndDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.RewardsEndDate) {
		return nil, false
	}
	return o.RewardsEndDate, true
}

// HasRewardsEndDate returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedSubscriptionPreviewResponse) HasRewardsEndDate() bool {
	if o != nil && !common.IsNil(o.RewardsEndDate) {
		return true
	}

	return false
}

// SetRewardsEndDate gets a reference to the given string and assigns it to the RewardsEndDate field.
func (o *GetOnChainYieldsLockedSubscriptionPreviewResponse) SetRewardsEndDate(v string) {
	o.RewardsEndDate = &v
}

// GetDeliverDate returns the DeliverDate field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedSubscriptionPreviewResponse) GetDeliverDate() string {
	if o == nil || common.IsNil(o.DeliverDate) {
		var ret string
		return ret
	}
	return *o.DeliverDate
}

// GetDeliverDateOk returns a tuple with the DeliverDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedSubscriptionPreviewResponse) GetDeliverDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.DeliverDate) {
		return nil, false
	}
	return o.DeliverDate, true
}

// HasDeliverDate returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedSubscriptionPreviewResponse) HasDeliverDate() bool {
	if o != nil && !common.IsNil(o.DeliverDate) {
		return true
	}

	return false
}

// SetDeliverDate gets a reference to the given string and assigns it to the DeliverDate field.
func (o *GetOnChainYieldsLockedSubscriptionPreviewResponse) SetDeliverDate(v string) {
	o.DeliverDate = &v
}

// GetNextSubscriptionDate returns the NextSubscriptionDate field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedSubscriptionPreviewResponse) GetNextSubscriptionDate() string {
	if o == nil || common.IsNil(o.NextSubscriptionDate) {
		var ret string
		return ret
	}
	return *o.NextSubscriptionDate
}

// GetNextSubscriptionDateOk returns a tuple with the NextSubscriptionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedSubscriptionPreviewResponse) GetNextSubscriptionDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.NextSubscriptionDate) {
		return nil, false
	}
	return o.NextSubscriptionDate, true
}

// HasNextSubscriptionDate returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedSubscriptionPreviewResponse) HasNextSubscriptionDate() bool {
	if o != nil && !common.IsNil(o.NextSubscriptionDate) {
		return true
	}

	return false
}

// SetNextSubscriptionDate gets a reference to the given string and assigns it to the NextSubscriptionDate field.
func (o *GetOnChainYieldsLockedSubscriptionPreviewResponse) SetNextSubscriptionDate(v string) {
	o.NextSubscriptionDate = &v
}

func (o GetOnChainYieldsLockedSubscriptionPreviewResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOnChainYieldsLockedSubscriptionPreviewResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.RewardAsset) {
		toSerialize["rewardAsset"] = o.RewardAsset
	}
	if !common.IsNil(o.TotalRewardAmt) {
		toSerialize["totalRewardAmt"] = o.TotalRewardAmt
	}
	if !common.IsNil(o.NextPay) {
		toSerialize["nextPay"] = o.NextPay
	}
	if !common.IsNil(o.NextPayDate) {
		toSerialize["nextPayDate"] = o.NextPayDate
	}
	if !common.IsNil(o.RewardsPayDate) {
		toSerialize["rewardsPayDate"] = o.RewardsPayDate
	}
	if !common.IsNil(o.ValueDate) {
		toSerialize["valueDate"] = o.ValueDate
	}
	if !common.IsNil(o.RewardsEndDate) {
		toSerialize["rewardsEndDate"] = o.RewardsEndDate
	}
	if !common.IsNil(o.DeliverDate) {
		toSerialize["deliverDate"] = o.DeliverDate
	}
	if !common.IsNil(o.NextSubscriptionDate) {
		toSerialize["nextSubscriptionDate"] = o.NextSubscriptionDate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetOnChainYieldsLockedSubscriptionPreviewResponse) UnmarshalJSON(data []byte) (err error) {
	varGetOnChainYieldsLockedSubscriptionPreviewResponse := _GetOnChainYieldsLockedSubscriptionPreviewResponse{}

	err = json.Unmarshal(data, &varGetOnChainYieldsLockedSubscriptionPreviewResponse)

	if err != nil {
		return err
	}

	*o = GetOnChainYieldsLockedSubscriptionPreviewResponse(varGetOnChainYieldsLockedSubscriptionPreviewResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rewardAsset")
		delete(additionalProperties, "totalRewardAmt")
		delete(additionalProperties, "nextPay")
		delete(additionalProperties, "nextPayDate")
		delete(additionalProperties, "rewardsPayDate")
		delete(additionalProperties, "valueDate")
		delete(additionalProperties, "rewardsEndDate")
		delete(additionalProperties, "deliverDate")
		delete(additionalProperties, "nextSubscriptionDate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetOnChainYieldsLockedSubscriptionPreviewResponse struct {
	value *GetOnChainYieldsLockedSubscriptionPreviewResponse
	isSet bool
}

func (v NullableGetOnChainYieldsLockedSubscriptionPreviewResponse) Get() *GetOnChainYieldsLockedSubscriptionPreviewResponse {
	return v.value
}

func (v *NullableGetOnChainYieldsLockedSubscriptionPreviewResponse) Set(val *GetOnChainYieldsLockedSubscriptionPreviewResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOnChainYieldsLockedSubscriptionPreviewResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOnChainYieldsLockedSubscriptionPreviewResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOnChainYieldsLockedSubscriptionPreviewResponse(val *GetOnChainYieldsLockedSubscriptionPreviewResponse) *NullableGetOnChainYieldsLockedSubscriptionPreviewResponse {
	return &NullableGetOnChainYieldsLockedSubscriptionPreviewResponse{value: val, isSet: true}
}

func (v NullableGetOnChainYieldsLockedSubscriptionPreviewResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOnChainYieldsLockedSubscriptionPreviewResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
