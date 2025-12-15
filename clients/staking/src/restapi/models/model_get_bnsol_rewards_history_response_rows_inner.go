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

// checks if the GetBnsolRewardsHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetBnsolRewardsHistoryResponseRowsInner{}

// GetBnsolRewardsHistoryResponseRowsInner struct for GetBnsolRewardsHistoryResponseRowsInner
type GetBnsolRewardsHistoryResponseRowsInner struct {
	Time                 *int64  `json:"time,omitempty"`
	AmountInSOL          *string `json:"amountInSOL,omitempty"`
	Holding              *string `json:"holding,omitempty"`
	HoldingInSOL         *string `json:"holdingInSOL,omitempty"`
	AnnualPercentageRate *string `json:"annualPercentageRate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetBnsolRewardsHistoryResponseRowsInner GetBnsolRewardsHistoryResponseRowsInner

// NewGetBnsolRewardsHistoryResponseRowsInner instantiates a new GetBnsolRewardsHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBnsolRewardsHistoryResponseRowsInner() *GetBnsolRewardsHistoryResponseRowsInner {
	this := GetBnsolRewardsHistoryResponseRowsInner{}
	return &this
}

// NewGetBnsolRewardsHistoryResponseRowsInnerWithDefaults instantiates a new GetBnsolRewardsHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBnsolRewardsHistoryResponseRowsInnerWithDefaults() *GetBnsolRewardsHistoryResponseRowsInner {
	this := GetBnsolRewardsHistoryResponseRowsInner{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetBnsolRewardsHistoryResponseRowsInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBnsolRewardsHistoryResponseRowsInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetBnsolRewardsHistoryResponseRowsInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *GetBnsolRewardsHistoryResponseRowsInner) SetTime(v int64) {
	o.Time = &v
}

// GetAmountInSOL returns the AmountInSOL field value if set, zero value otherwise.
func (o *GetBnsolRewardsHistoryResponseRowsInner) GetAmountInSOL() string {
	if o == nil || common.IsNil(o.AmountInSOL) {
		var ret string
		return ret
	}
	return *o.AmountInSOL
}

// GetAmountInSOLOk returns a tuple with the AmountInSOL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBnsolRewardsHistoryResponseRowsInner) GetAmountInSOLOk() (*string, bool) {
	if o == nil || common.IsNil(o.AmountInSOL) {
		return nil, false
	}
	return o.AmountInSOL, true
}

// HasAmountInSOL returns a boolean if a field has been set.
func (o *GetBnsolRewardsHistoryResponseRowsInner) HasAmountInSOL() bool {
	if o != nil && !common.IsNil(o.AmountInSOL) {
		return true
	}

	return false
}

// SetAmountInSOL gets a reference to the given string and assigns it to the AmountInSOL field.
func (o *GetBnsolRewardsHistoryResponseRowsInner) SetAmountInSOL(v string) {
	o.AmountInSOL = &v
}

// GetHolding returns the Holding field value if set, zero value otherwise.
func (o *GetBnsolRewardsHistoryResponseRowsInner) GetHolding() string {
	if o == nil || common.IsNil(o.Holding) {
		var ret string
		return ret
	}
	return *o.Holding
}

// GetHoldingOk returns a tuple with the Holding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBnsolRewardsHistoryResponseRowsInner) GetHoldingOk() (*string, bool) {
	if o == nil || common.IsNil(o.Holding) {
		return nil, false
	}
	return o.Holding, true
}

// HasHolding returns a boolean if a field has been set.
func (o *GetBnsolRewardsHistoryResponseRowsInner) HasHolding() bool {
	if o != nil && !common.IsNil(o.Holding) {
		return true
	}

	return false
}

// SetHolding gets a reference to the given string and assigns it to the Holding field.
func (o *GetBnsolRewardsHistoryResponseRowsInner) SetHolding(v string) {
	o.Holding = &v
}

// GetHoldingInSOL returns the HoldingInSOL field value if set, zero value otherwise.
func (o *GetBnsolRewardsHistoryResponseRowsInner) GetHoldingInSOL() string {
	if o == nil || common.IsNil(o.HoldingInSOL) {
		var ret string
		return ret
	}
	return *o.HoldingInSOL
}

// GetHoldingInSOLOk returns a tuple with the HoldingInSOL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBnsolRewardsHistoryResponseRowsInner) GetHoldingInSOLOk() (*string, bool) {
	if o == nil || common.IsNil(o.HoldingInSOL) {
		return nil, false
	}
	return o.HoldingInSOL, true
}

// HasHoldingInSOL returns a boolean if a field has been set.
func (o *GetBnsolRewardsHistoryResponseRowsInner) HasHoldingInSOL() bool {
	if o != nil && !common.IsNil(o.HoldingInSOL) {
		return true
	}

	return false
}

// SetHoldingInSOL gets a reference to the given string and assigns it to the HoldingInSOL field.
func (o *GetBnsolRewardsHistoryResponseRowsInner) SetHoldingInSOL(v string) {
	o.HoldingInSOL = &v
}

// GetAnnualPercentageRate returns the AnnualPercentageRate field value if set, zero value otherwise.
func (o *GetBnsolRewardsHistoryResponseRowsInner) GetAnnualPercentageRate() string {
	if o == nil || common.IsNil(o.AnnualPercentageRate) {
		var ret string
		return ret
	}
	return *o.AnnualPercentageRate
}

// GetAnnualPercentageRateOk returns a tuple with the AnnualPercentageRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBnsolRewardsHistoryResponseRowsInner) GetAnnualPercentageRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.AnnualPercentageRate) {
		return nil, false
	}
	return o.AnnualPercentageRate, true
}

// HasAnnualPercentageRate returns a boolean if a field has been set.
func (o *GetBnsolRewardsHistoryResponseRowsInner) HasAnnualPercentageRate() bool {
	if o != nil && !common.IsNil(o.AnnualPercentageRate) {
		return true
	}

	return false
}

// SetAnnualPercentageRate gets a reference to the given string and assigns it to the AnnualPercentageRate field.
func (o *GetBnsolRewardsHistoryResponseRowsInner) SetAnnualPercentageRate(v string) {
	o.AnnualPercentageRate = &v
}

func (o GetBnsolRewardsHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBnsolRewardsHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !common.IsNil(o.AmountInSOL) {
		toSerialize["amountInSOL"] = o.AmountInSOL
	}
	if !common.IsNil(o.Holding) {
		toSerialize["holding"] = o.Holding
	}
	if !common.IsNil(o.HoldingInSOL) {
		toSerialize["holdingInSOL"] = o.HoldingInSOL
	}
	if !common.IsNil(o.AnnualPercentageRate) {
		toSerialize["annualPercentageRate"] = o.AnnualPercentageRate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetBnsolRewardsHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetBnsolRewardsHistoryResponseRowsInner := _GetBnsolRewardsHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varGetBnsolRewardsHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetBnsolRewardsHistoryResponseRowsInner(varGetBnsolRewardsHistoryResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "time")
		delete(additionalProperties, "amountInSOL")
		delete(additionalProperties, "holding")
		delete(additionalProperties, "holdingInSOL")
		delete(additionalProperties, "annualPercentageRate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetBnsolRewardsHistoryResponseRowsInner struct {
	value *GetBnsolRewardsHistoryResponseRowsInner
	isSet bool
}

func (v NullableGetBnsolRewardsHistoryResponseRowsInner) Get() *GetBnsolRewardsHistoryResponseRowsInner {
	return v.value
}

func (v *NullableGetBnsolRewardsHistoryResponseRowsInner) Set(val *GetBnsolRewardsHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBnsolRewardsHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBnsolRewardsHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBnsolRewardsHistoryResponseRowsInner(val *GetBnsolRewardsHistoryResponseRowsInner) *NullableGetBnsolRewardsHistoryResponseRowsInner {
	return &NullableGetBnsolRewardsHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetBnsolRewardsHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBnsolRewardsHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
