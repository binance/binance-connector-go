/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetRwusdRewardsHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetRwusdRewardsHistoryResponseRowsInner{}

// GetRwusdRewardsHistoryResponseRowsInner struct for GetRwusdRewardsHistoryResponseRowsInner
type GetRwusdRewardsHistoryResponseRowsInner struct {
	Time                 *int64  `json:"time,omitempty"`
	RewardsAmount        *string `json:"rewardsAmount,omitempty"`
	RwusdPosition        *string `json:"rwusdPosition,omitempty"`
	AnnualPercentageRate *string `json:"annualPercentageRate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetRwusdRewardsHistoryResponseRowsInner GetRwusdRewardsHistoryResponseRowsInner

// NewGetRwusdRewardsHistoryResponseRowsInner instantiates a new GetRwusdRewardsHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRwusdRewardsHistoryResponseRowsInner() *GetRwusdRewardsHistoryResponseRowsInner {
	this := GetRwusdRewardsHistoryResponseRowsInner{}
	return &this
}

// NewGetRwusdRewardsHistoryResponseRowsInnerWithDefaults instantiates a new GetRwusdRewardsHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRwusdRewardsHistoryResponseRowsInnerWithDefaults() *GetRwusdRewardsHistoryResponseRowsInner {
	this := GetRwusdRewardsHistoryResponseRowsInner{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetRwusdRewardsHistoryResponseRowsInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRwusdRewardsHistoryResponseRowsInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetRwusdRewardsHistoryResponseRowsInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *GetRwusdRewardsHistoryResponseRowsInner) SetTime(v int64) {
	o.Time = &v
}

// GetRewardsAmount returns the RewardsAmount field value if set, zero value otherwise.
func (o *GetRwusdRewardsHistoryResponseRowsInner) GetRewardsAmount() string {
	if o == nil || common.IsNil(o.RewardsAmount) {
		var ret string
		return ret
	}
	return *o.RewardsAmount
}

// GetRewardsAmountOk returns a tuple with the RewardsAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRwusdRewardsHistoryResponseRowsInner) GetRewardsAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.RewardsAmount) {
		return nil, false
	}
	return o.RewardsAmount, true
}

// HasRewardsAmount returns a boolean if a field has been set.
func (o *GetRwusdRewardsHistoryResponseRowsInner) HasRewardsAmount() bool {
	if o != nil && !common.IsNil(o.RewardsAmount) {
		return true
	}

	return false
}

// SetRewardsAmount gets a reference to the given string and assigns it to the RewardsAmount field.
func (o *GetRwusdRewardsHistoryResponseRowsInner) SetRewardsAmount(v string) {
	o.RewardsAmount = &v
}

// GetRwusdPosition returns the RwusdPosition field value if set, zero value otherwise.
func (o *GetRwusdRewardsHistoryResponseRowsInner) GetRwusdPosition() string {
	if o == nil || common.IsNil(o.RwusdPosition) {
		var ret string
		return ret
	}
	return *o.RwusdPosition
}

// GetRwusdPositionOk returns a tuple with the RwusdPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRwusdRewardsHistoryResponseRowsInner) GetRwusdPositionOk() (*string, bool) {
	if o == nil || common.IsNil(o.RwusdPosition) {
		return nil, false
	}
	return o.RwusdPosition, true
}

// HasRwusdPosition returns a boolean if a field has been set.
func (o *GetRwusdRewardsHistoryResponseRowsInner) HasRwusdPosition() bool {
	if o != nil && !common.IsNil(o.RwusdPosition) {
		return true
	}

	return false
}

// SetRwusdPosition gets a reference to the given string and assigns it to the RwusdPosition field.
func (o *GetRwusdRewardsHistoryResponseRowsInner) SetRwusdPosition(v string) {
	o.RwusdPosition = &v
}

// GetAnnualPercentageRate returns the AnnualPercentageRate field value if set, zero value otherwise.
func (o *GetRwusdRewardsHistoryResponseRowsInner) GetAnnualPercentageRate() string {
	if o == nil || common.IsNil(o.AnnualPercentageRate) {
		var ret string
		return ret
	}
	return *o.AnnualPercentageRate
}

// GetAnnualPercentageRateOk returns a tuple with the AnnualPercentageRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRwusdRewardsHistoryResponseRowsInner) GetAnnualPercentageRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.AnnualPercentageRate) {
		return nil, false
	}
	return o.AnnualPercentageRate, true
}

// HasAnnualPercentageRate returns a boolean if a field has been set.
func (o *GetRwusdRewardsHistoryResponseRowsInner) HasAnnualPercentageRate() bool {
	if o != nil && !common.IsNil(o.AnnualPercentageRate) {
		return true
	}

	return false
}

// SetAnnualPercentageRate gets a reference to the given string and assigns it to the AnnualPercentageRate field.
func (o *GetRwusdRewardsHistoryResponseRowsInner) SetAnnualPercentageRate(v string) {
	o.AnnualPercentageRate = &v
}

func (o GetRwusdRewardsHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRwusdRewardsHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !common.IsNil(o.RewardsAmount) {
		toSerialize["rewardsAmount"] = o.RewardsAmount
	}
	if !common.IsNil(o.RwusdPosition) {
		toSerialize["rwusdPosition"] = o.RwusdPosition
	}
	if !common.IsNil(o.AnnualPercentageRate) {
		toSerialize["annualPercentageRate"] = o.AnnualPercentageRate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetRwusdRewardsHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetRwusdRewardsHistoryResponseRowsInner := _GetRwusdRewardsHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varGetRwusdRewardsHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetRwusdRewardsHistoryResponseRowsInner(varGetRwusdRewardsHistoryResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "time")
		delete(additionalProperties, "rewardsAmount")
		delete(additionalProperties, "rwusdPosition")
		delete(additionalProperties, "annualPercentageRate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetRwusdRewardsHistoryResponseRowsInner struct {
	value *GetRwusdRewardsHistoryResponseRowsInner
	isSet bool
}

func (v NullableGetRwusdRewardsHistoryResponseRowsInner) Get() *GetRwusdRewardsHistoryResponseRowsInner {
	return v.value
}

func (v *NullableGetRwusdRewardsHistoryResponseRowsInner) Set(val *GetRwusdRewardsHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRwusdRewardsHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRwusdRewardsHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRwusdRewardsHistoryResponseRowsInner(val *GetRwusdRewardsHistoryResponseRowsInner) *NullableGetRwusdRewardsHistoryResponseRowsInner {
	return &NullableGetRwusdRewardsHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetRwusdRewardsHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRwusdRewardsHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
