/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetFlexibleSubscriptionPreviewResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFlexibleSubscriptionPreviewResponse{}

// GetFlexibleSubscriptionPreviewResponse struct for GetFlexibleSubscriptionPreviewResponse
type GetFlexibleSubscriptionPreviewResponse struct {
	TotalAmount             *string `json:"totalAmount,omitempty"`
	RewardAsset             *string `json:"rewardAsset,omitempty"`
	AirDropAsset            *string `json:"airDropAsset,omitempty"`
	EstDailyBonusRewards    *string `json:"estDailyBonusRewards,omitempty"`
	EstDailyRealTimeRewards *string `json:"estDailyRealTimeRewards,omitempty"`
	EstDailyAirdropRewards  *string `json:"estDailyAirdropRewards,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _GetFlexibleSubscriptionPreviewResponse GetFlexibleSubscriptionPreviewResponse

// NewGetFlexibleSubscriptionPreviewResponse instantiates a new GetFlexibleSubscriptionPreviewResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFlexibleSubscriptionPreviewResponse() *GetFlexibleSubscriptionPreviewResponse {
	this := GetFlexibleSubscriptionPreviewResponse{}
	return &this
}

// NewGetFlexibleSubscriptionPreviewResponseWithDefaults instantiates a new GetFlexibleSubscriptionPreviewResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFlexibleSubscriptionPreviewResponseWithDefaults() *GetFlexibleSubscriptionPreviewResponse {
	this := GetFlexibleSubscriptionPreviewResponse{}
	return &this
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *GetFlexibleSubscriptionPreviewResponse) GetTotalAmount() string {
	if o == nil || common.IsNil(o.TotalAmount) {
		var ret string
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleSubscriptionPreviewResponse) GetTotalAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalAmount) {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *GetFlexibleSubscriptionPreviewResponse) HasTotalAmount() bool {
	if o != nil && !common.IsNil(o.TotalAmount) {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given string and assigns it to the TotalAmount field.
func (o *GetFlexibleSubscriptionPreviewResponse) SetTotalAmount(v string) {
	o.TotalAmount = &v
}

// GetRewardAsset returns the RewardAsset field value if set, zero value otherwise.
func (o *GetFlexibleSubscriptionPreviewResponse) GetRewardAsset() string {
	if o == nil || common.IsNil(o.RewardAsset) {
		var ret string
		return ret
	}
	return *o.RewardAsset
}

// GetRewardAssetOk returns a tuple with the RewardAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleSubscriptionPreviewResponse) GetRewardAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.RewardAsset) {
		return nil, false
	}
	return o.RewardAsset, true
}

// HasRewardAsset returns a boolean if a field has been set.
func (o *GetFlexibleSubscriptionPreviewResponse) HasRewardAsset() bool {
	if o != nil && !common.IsNil(o.RewardAsset) {
		return true
	}

	return false
}

// SetRewardAsset gets a reference to the given string and assigns it to the RewardAsset field.
func (o *GetFlexibleSubscriptionPreviewResponse) SetRewardAsset(v string) {
	o.RewardAsset = &v
}

// GetAirDropAsset returns the AirDropAsset field value if set, zero value otherwise.
func (o *GetFlexibleSubscriptionPreviewResponse) GetAirDropAsset() string {
	if o == nil || common.IsNil(o.AirDropAsset) {
		var ret string
		return ret
	}
	return *o.AirDropAsset
}

// GetAirDropAssetOk returns a tuple with the AirDropAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleSubscriptionPreviewResponse) GetAirDropAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.AirDropAsset) {
		return nil, false
	}
	return o.AirDropAsset, true
}

// HasAirDropAsset returns a boolean if a field has been set.
func (o *GetFlexibleSubscriptionPreviewResponse) HasAirDropAsset() bool {
	if o != nil && !common.IsNil(o.AirDropAsset) {
		return true
	}

	return false
}

// SetAirDropAsset gets a reference to the given string and assigns it to the AirDropAsset field.
func (o *GetFlexibleSubscriptionPreviewResponse) SetAirDropAsset(v string) {
	o.AirDropAsset = &v
}

// GetEstDailyBonusRewards returns the EstDailyBonusRewards field value if set, zero value otherwise.
func (o *GetFlexibleSubscriptionPreviewResponse) GetEstDailyBonusRewards() string {
	if o == nil || common.IsNil(o.EstDailyBonusRewards) {
		var ret string
		return ret
	}
	return *o.EstDailyBonusRewards
}

// GetEstDailyBonusRewardsOk returns a tuple with the EstDailyBonusRewards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleSubscriptionPreviewResponse) GetEstDailyBonusRewardsOk() (*string, bool) {
	if o == nil || common.IsNil(o.EstDailyBonusRewards) {
		return nil, false
	}
	return o.EstDailyBonusRewards, true
}

// HasEstDailyBonusRewards returns a boolean if a field has been set.
func (o *GetFlexibleSubscriptionPreviewResponse) HasEstDailyBonusRewards() bool {
	if o != nil && !common.IsNil(o.EstDailyBonusRewards) {
		return true
	}

	return false
}

// SetEstDailyBonusRewards gets a reference to the given string and assigns it to the EstDailyBonusRewards field.
func (o *GetFlexibleSubscriptionPreviewResponse) SetEstDailyBonusRewards(v string) {
	o.EstDailyBonusRewards = &v
}

// GetEstDailyRealTimeRewards returns the EstDailyRealTimeRewards field value if set, zero value otherwise.
func (o *GetFlexibleSubscriptionPreviewResponse) GetEstDailyRealTimeRewards() string {
	if o == nil || common.IsNil(o.EstDailyRealTimeRewards) {
		var ret string
		return ret
	}
	return *o.EstDailyRealTimeRewards
}

// GetEstDailyRealTimeRewardsOk returns a tuple with the EstDailyRealTimeRewards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleSubscriptionPreviewResponse) GetEstDailyRealTimeRewardsOk() (*string, bool) {
	if o == nil || common.IsNil(o.EstDailyRealTimeRewards) {
		return nil, false
	}
	return o.EstDailyRealTimeRewards, true
}

// HasEstDailyRealTimeRewards returns a boolean if a field has been set.
func (o *GetFlexibleSubscriptionPreviewResponse) HasEstDailyRealTimeRewards() bool {
	if o != nil && !common.IsNil(o.EstDailyRealTimeRewards) {
		return true
	}

	return false
}

// SetEstDailyRealTimeRewards gets a reference to the given string and assigns it to the EstDailyRealTimeRewards field.
func (o *GetFlexibleSubscriptionPreviewResponse) SetEstDailyRealTimeRewards(v string) {
	o.EstDailyRealTimeRewards = &v
}

// GetEstDailyAirdropRewards returns the EstDailyAirdropRewards field value if set, zero value otherwise.
func (o *GetFlexibleSubscriptionPreviewResponse) GetEstDailyAirdropRewards() string {
	if o == nil || common.IsNil(o.EstDailyAirdropRewards) {
		var ret string
		return ret
	}
	return *o.EstDailyAirdropRewards
}

// GetEstDailyAirdropRewardsOk returns a tuple with the EstDailyAirdropRewards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleSubscriptionPreviewResponse) GetEstDailyAirdropRewardsOk() (*string, bool) {
	if o == nil || common.IsNil(o.EstDailyAirdropRewards) {
		return nil, false
	}
	return o.EstDailyAirdropRewards, true
}

// HasEstDailyAirdropRewards returns a boolean if a field has been set.
func (o *GetFlexibleSubscriptionPreviewResponse) HasEstDailyAirdropRewards() bool {
	if o != nil && !common.IsNil(o.EstDailyAirdropRewards) {
		return true
	}

	return false
}

// SetEstDailyAirdropRewards gets a reference to the given string and assigns it to the EstDailyAirdropRewards field.
func (o *GetFlexibleSubscriptionPreviewResponse) SetEstDailyAirdropRewards(v string) {
	o.EstDailyAirdropRewards = &v
}

func (o GetFlexibleSubscriptionPreviewResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFlexibleSubscriptionPreviewResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TotalAmount) {
		toSerialize["totalAmount"] = o.TotalAmount
	}
	if !common.IsNil(o.RewardAsset) {
		toSerialize["rewardAsset"] = o.RewardAsset
	}
	if !common.IsNil(o.AirDropAsset) {
		toSerialize["airDropAsset"] = o.AirDropAsset
	}
	if !common.IsNil(o.EstDailyBonusRewards) {
		toSerialize["estDailyBonusRewards"] = o.EstDailyBonusRewards
	}
	if !common.IsNil(o.EstDailyRealTimeRewards) {
		toSerialize["estDailyRealTimeRewards"] = o.EstDailyRealTimeRewards
	}
	if !common.IsNil(o.EstDailyAirdropRewards) {
		toSerialize["estDailyAirdropRewards"] = o.EstDailyAirdropRewards
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetFlexibleSubscriptionPreviewResponse) UnmarshalJSON(data []byte) (err error) {
	varGetFlexibleSubscriptionPreviewResponse := _GetFlexibleSubscriptionPreviewResponse{}

	err = json.Unmarshal(data, &varGetFlexibleSubscriptionPreviewResponse)

	if err != nil {
		return err
	}

	*o = GetFlexibleSubscriptionPreviewResponse(varGetFlexibleSubscriptionPreviewResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "totalAmount")
		delete(additionalProperties, "rewardAsset")
		delete(additionalProperties, "airDropAsset")
		delete(additionalProperties, "estDailyBonusRewards")
		delete(additionalProperties, "estDailyRealTimeRewards")
		delete(additionalProperties, "estDailyAirdropRewards")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFlexibleSubscriptionPreviewResponse struct {
	value *GetFlexibleSubscriptionPreviewResponse
	isSet bool
}

func (v NullableGetFlexibleSubscriptionPreviewResponse) Get() *GetFlexibleSubscriptionPreviewResponse {
	return v.value
}

func (v *NullableGetFlexibleSubscriptionPreviewResponse) Set(val *GetFlexibleSubscriptionPreviewResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFlexibleSubscriptionPreviewResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFlexibleSubscriptionPreviewResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFlexibleSubscriptionPreviewResponse(val *GetFlexibleSubscriptionPreviewResponse) *NullableGetFlexibleSubscriptionPreviewResponse {
	return &NullableGetFlexibleSubscriptionPreviewResponse{value: val, isSet: true}
}

func (v NullableGetFlexibleSubscriptionPreviewResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFlexibleSubscriptionPreviewResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
