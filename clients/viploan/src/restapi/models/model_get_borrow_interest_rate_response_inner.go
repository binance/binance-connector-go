/*
Binance VIP Loan REST API

OpenAPI Specification for the Binance VIP Loan REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetBorrowInterestRateResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetBorrowInterestRateResponseInner{}

// GetBorrowInterestRateResponseInner struct for GetBorrowInterestRateResponseInner
type GetBorrowInterestRateResponseInner struct {
	Asset                      *string `json:"asset,omitempty"`
	FlexibleDailyInterestRate  *string `json:"flexibleDailyInterestRate,omitempty"`
	FlexibleYearlyInterestRate *string `json:"flexibleYearlyInterestRate,omitempty"`
	Time                       *int64  `json:"time,omitempty"`
	AdditionalProperties       map[string]interface{}
}

type _GetBorrowInterestRateResponseInner GetBorrowInterestRateResponseInner

// NewGetBorrowInterestRateResponseInner instantiates a new GetBorrowInterestRateResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBorrowInterestRateResponseInner() *GetBorrowInterestRateResponseInner {
	this := GetBorrowInterestRateResponseInner{}
	return &this
}

// NewGetBorrowInterestRateResponseInnerWithDefaults instantiates a new GetBorrowInterestRateResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBorrowInterestRateResponseInnerWithDefaults() *GetBorrowInterestRateResponseInner {
	this := GetBorrowInterestRateResponseInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetBorrowInterestRateResponseInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBorrowInterestRateResponseInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetBorrowInterestRateResponseInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetBorrowInterestRateResponseInner) SetAsset(v string) {
	o.Asset = &v
}

// GetFlexibleDailyInterestRate returns the FlexibleDailyInterestRate field value if set, zero value otherwise.
func (o *GetBorrowInterestRateResponseInner) GetFlexibleDailyInterestRate() string {
	if o == nil || common.IsNil(o.FlexibleDailyInterestRate) {
		var ret string
		return ret
	}
	return *o.FlexibleDailyInterestRate
}

// GetFlexibleDailyInterestRateOk returns a tuple with the FlexibleDailyInterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBorrowInterestRateResponseInner) GetFlexibleDailyInterestRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.FlexibleDailyInterestRate) {
		return nil, false
	}
	return o.FlexibleDailyInterestRate, true
}

// HasFlexibleDailyInterestRate returns a boolean if a field has been set.
func (o *GetBorrowInterestRateResponseInner) HasFlexibleDailyInterestRate() bool {
	if o != nil && !common.IsNil(o.FlexibleDailyInterestRate) {
		return true
	}

	return false
}

// SetFlexibleDailyInterestRate gets a reference to the given string and assigns it to the FlexibleDailyInterestRate field.
func (o *GetBorrowInterestRateResponseInner) SetFlexibleDailyInterestRate(v string) {
	o.FlexibleDailyInterestRate = &v
}

// GetFlexibleYearlyInterestRate returns the FlexibleYearlyInterestRate field value if set, zero value otherwise.
func (o *GetBorrowInterestRateResponseInner) GetFlexibleYearlyInterestRate() string {
	if o == nil || common.IsNil(o.FlexibleYearlyInterestRate) {
		var ret string
		return ret
	}
	return *o.FlexibleYearlyInterestRate
}

// GetFlexibleYearlyInterestRateOk returns a tuple with the FlexibleYearlyInterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBorrowInterestRateResponseInner) GetFlexibleYearlyInterestRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.FlexibleYearlyInterestRate) {
		return nil, false
	}
	return o.FlexibleYearlyInterestRate, true
}

// HasFlexibleYearlyInterestRate returns a boolean if a field has been set.
func (o *GetBorrowInterestRateResponseInner) HasFlexibleYearlyInterestRate() bool {
	if o != nil && !common.IsNil(o.FlexibleYearlyInterestRate) {
		return true
	}

	return false
}

// SetFlexibleYearlyInterestRate gets a reference to the given string and assigns it to the FlexibleYearlyInterestRate field.
func (o *GetBorrowInterestRateResponseInner) SetFlexibleYearlyInterestRate(v string) {
	o.FlexibleYearlyInterestRate = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetBorrowInterestRateResponseInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBorrowInterestRateResponseInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetBorrowInterestRateResponseInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *GetBorrowInterestRateResponseInner) SetTime(v int64) {
	o.Time = &v
}

func (o GetBorrowInterestRateResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBorrowInterestRateResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.FlexibleDailyInterestRate) {
		toSerialize["flexibleDailyInterestRate"] = o.FlexibleDailyInterestRate
	}
	if !common.IsNil(o.FlexibleYearlyInterestRate) {
		toSerialize["flexibleYearlyInterestRate"] = o.FlexibleYearlyInterestRate
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetBorrowInterestRateResponseInner) UnmarshalJSON(data []byte) (err error) {
	varGetBorrowInterestRateResponseInner := _GetBorrowInterestRateResponseInner{}

	err = json.Unmarshal(data, &varGetBorrowInterestRateResponseInner)

	if err != nil {
		return err
	}

	*o = GetBorrowInterestRateResponseInner(varGetBorrowInterestRateResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "flexibleDailyInterestRate")
		delete(additionalProperties, "flexibleYearlyInterestRate")
		delete(additionalProperties, "time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetBorrowInterestRateResponseInner struct {
	value *GetBorrowInterestRateResponseInner
	isSet bool
}

func (v NullableGetBorrowInterestRateResponseInner) Get() *GetBorrowInterestRateResponseInner {
	return v.value
}

func (v *NullableGetBorrowInterestRateResponseInner) Set(val *GetBorrowInterestRateResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBorrowInterestRateResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBorrowInterestRateResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBorrowInterestRateResponseInner(val *GetBorrowInterestRateResponseInner) *NullableGetBorrowInterestRateResponseInner {
	return &NullableGetBorrowInterestRateResponseInner{value: val, isSet: true}
}

func (v NullableGetBorrowInterestRateResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBorrowInterestRateResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
