/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetWbethRewardsHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetWbethRewardsHistoryResponseRowsInner{}

// GetWbethRewardsHistoryResponseRowsInner struct for GetWbethRewardsHistoryResponseRowsInner
type GetWbethRewardsHistoryResponseRowsInner struct {
	Time                 *int64  `json:"time,omitempty"`
	AmountInETH          *string `json:"amountInETH,omitempty"`
	Holding              *string `json:"holding,omitempty"`
	HoldingInETH         *string `json:"holdingInETH,omitempty"`
	AnnualPercentageRate *string `json:"annualPercentageRate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetWbethRewardsHistoryResponseRowsInner GetWbethRewardsHistoryResponseRowsInner

// NewGetWbethRewardsHistoryResponseRowsInner instantiates a new GetWbethRewardsHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWbethRewardsHistoryResponseRowsInner() *GetWbethRewardsHistoryResponseRowsInner {
	this := GetWbethRewardsHistoryResponseRowsInner{}
	return &this
}

// NewGetWbethRewardsHistoryResponseRowsInnerWithDefaults instantiates a new GetWbethRewardsHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWbethRewardsHistoryResponseRowsInnerWithDefaults() *GetWbethRewardsHistoryResponseRowsInner {
	this := GetWbethRewardsHistoryResponseRowsInner{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetWbethRewardsHistoryResponseRowsInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWbethRewardsHistoryResponseRowsInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetWbethRewardsHistoryResponseRowsInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *GetWbethRewardsHistoryResponseRowsInner) SetTime(v int64) {
	o.Time = &v
}

// GetAmountInETH returns the AmountInETH field value if set, zero value otherwise.
func (o *GetWbethRewardsHistoryResponseRowsInner) GetAmountInETH() string {
	if o == nil || common.IsNil(o.AmountInETH) {
		var ret string
		return ret
	}
	return *o.AmountInETH
}

// GetAmountInETHOk returns a tuple with the AmountInETH field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWbethRewardsHistoryResponseRowsInner) GetAmountInETHOk() (*string, bool) {
	if o == nil || common.IsNil(o.AmountInETH) {
		return nil, false
	}
	return o.AmountInETH, true
}

// HasAmountInETH returns a boolean if a field has been set.
func (o *GetWbethRewardsHistoryResponseRowsInner) HasAmountInETH() bool {
	if o != nil && !common.IsNil(o.AmountInETH) {
		return true
	}

	return false
}

// SetAmountInETH gets a reference to the given string and assigns it to the AmountInETH field.
func (o *GetWbethRewardsHistoryResponseRowsInner) SetAmountInETH(v string) {
	o.AmountInETH = &v
}

// GetHolding returns the Holding field value if set, zero value otherwise.
func (o *GetWbethRewardsHistoryResponseRowsInner) GetHolding() string {
	if o == nil || common.IsNil(o.Holding) {
		var ret string
		return ret
	}
	return *o.Holding
}

// GetHoldingOk returns a tuple with the Holding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWbethRewardsHistoryResponseRowsInner) GetHoldingOk() (*string, bool) {
	if o == nil || common.IsNil(o.Holding) {
		return nil, false
	}
	return o.Holding, true
}

// HasHolding returns a boolean if a field has been set.
func (o *GetWbethRewardsHistoryResponseRowsInner) HasHolding() bool {
	if o != nil && !common.IsNil(o.Holding) {
		return true
	}

	return false
}

// SetHolding gets a reference to the given string and assigns it to the Holding field.
func (o *GetWbethRewardsHistoryResponseRowsInner) SetHolding(v string) {
	o.Holding = &v
}

// GetHoldingInETH returns the HoldingInETH field value if set, zero value otherwise.
func (o *GetWbethRewardsHistoryResponseRowsInner) GetHoldingInETH() string {
	if o == nil || common.IsNil(o.HoldingInETH) {
		var ret string
		return ret
	}
	return *o.HoldingInETH
}

// GetHoldingInETHOk returns a tuple with the HoldingInETH field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWbethRewardsHistoryResponseRowsInner) GetHoldingInETHOk() (*string, bool) {
	if o == nil || common.IsNil(o.HoldingInETH) {
		return nil, false
	}
	return o.HoldingInETH, true
}

// HasHoldingInETH returns a boolean if a field has been set.
func (o *GetWbethRewardsHistoryResponseRowsInner) HasHoldingInETH() bool {
	if o != nil && !common.IsNil(o.HoldingInETH) {
		return true
	}

	return false
}

// SetHoldingInETH gets a reference to the given string and assigns it to the HoldingInETH field.
func (o *GetWbethRewardsHistoryResponseRowsInner) SetHoldingInETH(v string) {
	o.HoldingInETH = &v
}

// GetAnnualPercentageRate returns the AnnualPercentageRate field value if set, zero value otherwise.
func (o *GetWbethRewardsHistoryResponseRowsInner) GetAnnualPercentageRate() string {
	if o == nil || common.IsNil(o.AnnualPercentageRate) {
		var ret string
		return ret
	}
	return *o.AnnualPercentageRate
}

// GetAnnualPercentageRateOk returns a tuple with the AnnualPercentageRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWbethRewardsHistoryResponseRowsInner) GetAnnualPercentageRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.AnnualPercentageRate) {
		return nil, false
	}
	return o.AnnualPercentageRate, true
}

// HasAnnualPercentageRate returns a boolean if a field has been set.
func (o *GetWbethRewardsHistoryResponseRowsInner) HasAnnualPercentageRate() bool {
	if o != nil && !common.IsNil(o.AnnualPercentageRate) {
		return true
	}

	return false
}

// SetAnnualPercentageRate gets a reference to the given string and assigns it to the AnnualPercentageRate field.
func (o *GetWbethRewardsHistoryResponseRowsInner) SetAnnualPercentageRate(v string) {
	o.AnnualPercentageRate = &v
}

func (o GetWbethRewardsHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetWbethRewardsHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !common.IsNil(o.AmountInETH) {
		toSerialize["amountInETH"] = o.AmountInETH
	}
	if !common.IsNil(o.Holding) {
		toSerialize["holding"] = o.Holding
	}
	if !common.IsNil(o.HoldingInETH) {
		toSerialize["holdingInETH"] = o.HoldingInETH
	}
	if !common.IsNil(o.AnnualPercentageRate) {
		toSerialize["annualPercentageRate"] = o.AnnualPercentageRate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetWbethRewardsHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetWbethRewardsHistoryResponseRowsInner := _GetWbethRewardsHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varGetWbethRewardsHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetWbethRewardsHistoryResponseRowsInner(varGetWbethRewardsHistoryResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "time")
		delete(additionalProperties, "amountInETH")
		delete(additionalProperties, "holding")
		delete(additionalProperties, "holdingInETH")
		delete(additionalProperties, "annualPercentageRate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetWbethRewardsHistoryResponseRowsInner struct {
	value *GetWbethRewardsHistoryResponseRowsInner
	isSet bool
}

func (v NullableGetWbethRewardsHistoryResponseRowsInner) Get() *GetWbethRewardsHistoryResponseRowsInner {
	return v.value
}

func (v *NullableGetWbethRewardsHistoryResponseRowsInner) Set(val *GetWbethRewardsHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWbethRewardsHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWbethRewardsHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWbethRewardsHistoryResponseRowsInner(val *GetWbethRewardsHistoryResponseRowsInner) *NullableGetWbethRewardsHistoryResponseRowsInner {
	return &NullableGetWbethRewardsHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetWbethRewardsHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWbethRewardsHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
