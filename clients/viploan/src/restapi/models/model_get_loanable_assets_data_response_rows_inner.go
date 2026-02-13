/*
Binance VIP Loan REST API

OpenAPI Specification for the Binance VIP Loan REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetLoanableAssetsDataResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetLoanableAssetsDataResponseRowsInner{}

// GetLoanableAssetsDataResponseRowsInner struct for GetLoanableAssetsDataResponseRowsInner
type GetLoanableAssetsDataResponseRowsInner struct {
	LoanCoin                   *string `json:"loanCoin,omitempty"`
	FlexibleDailyInterestRate  *string `json:"_flexibleDailyInterestRate,omitempty"`
	FlexibleYearlyInterestRate *string `json:"_flexibleYearlyInterestRate,omitempty"`
	Var30dDailyInterestRate    *string `json:"_30dDailyInterestRate,omitempty"`
	Var30dYearlyInterestRate   *string `json:"_30dYearlyInterestRate,omitempty"`
	Var60dDailyInterestRate    *string `json:"_60dDailyInterestRate,omitempty"`
	Var60dYearlyInterestRate   *string `json:"_60dYearlyInterestRate,omitempty"`
	MinLimit                   *string `json:"minLimit,omitempty"`
	MaxLimit                   *string `json:"maxLimit,omitempty"`
	VipLevel                   *int64  `json:"vipLevel,omitempty"`
	AdditionalProperties       map[string]interface{}
}

type _GetLoanableAssetsDataResponseRowsInner GetLoanableAssetsDataResponseRowsInner

// NewGetLoanableAssetsDataResponseRowsInner instantiates a new GetLoanableAssetsDataResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLoanableAssetsDataResponseRowsInner() *GetLoanableAssetsDataResponseRowsInner {
	this := GetLoanableAssetsDataResponseRowsInner{}
	return &this
}

// NewGetLoanableAssetsDataResponseRowsInnerWithDefaults instantiates a new GetLoanableAssetsDataResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLoanableAssetsDataResponseRowsInnerWithDefaults() *GetLoanableAssetsDataResponseRowsInner {
	this := GetLoanableAssetsDataResponseRowsInner{}
	return &this
}

// GetLoanCoin returns the LoanCoin field value if set, zero value otherwise.
func (o *GetLoanableAssetsDataResponseRowsInner) GetLoanCoin() string {
	if o == nil || common.IsNil(o.LoanCoin) {
		var ret string
		return ret
	}
	return *o.LoanCoin
}

// GetLoanCoinOk returns a tuple with the LoanCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanableAssetsDataResponseRowsInner) GetLoanCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoanCoin) {
		return nil, false
	}
	return o.LoanCoin, true
}

// HasLoanCoin returns a boolean if a field has been set.
func (o *GetLoanableAssetsDataResponseRowsInner) HasLoanCoin() bool {
	if o != nil && !common.IsNil(o.LoanCoin) {
		return true
	}

	return false
}

// SetLoanCoin gets a reference to the given string and assigns it to the LoanCoin field.
func (o *GetLoanableAssetsDataResponseRowsInner) SetLoanCoin(v string) {
	o.LoanCoin = &v
}

// GetFlexibleDailyInterestRate returns the FlexibleDailyInterestRate field value if set, zero value otherwise.
func (o *GetLoanableAssetsDataResponseRowsInner) GetFlexibleDailyInterestRate() string {
	if o == nil || common.IsNil(o.FlexibleDailyInterestRate) {
		var ret string
		return ret
	}
	return *o.FlexibleDailyInterestRate
}

// GetFlexibleDailyInterestRateOk returns a tuple with the FlexibleDailyInterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanableAssetsDataResponseRowsInner) GetFlexibleDailyInterestRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.FlexibleDailyInterestRate) {
		return nil, false
	}
	return o.FlexibleDailyInterestRate, true
}

// HasFlexibleDailyInterestRate returns a boolean if a field has been set.
func (o *GetLoanableAssetsDataResponseRowsInner) HasFlexibleDailyInterestRate() bool {
	if o != nil && !common.IsNil(o.FlexibleDailyInterestRate) {
		return true
	}

	return false
}

// SetFlexibleDailyInterestRate gets a reference to the given string and assigns it to the FlexibleDailyInterestRate field.
func (o *GetLoanableAssetsDataResponseRowsInner) SetFlexibleDailyInterestRate(v string) {
	o.FlexibleDailyInterestRate = &v
}

// GetFlexibleYearlyInterestRate returns the FlexibleYearlyInterestRate field value if set, zero value otherwise.
func (o *GetLoanableAssetsDataResponseRowsInner) GetFlexibleYearlyInterestRate() string {
	if o == nil || common.IsNil(o.FlexibleYearlyInterestRate) {
		var ret string
		return ret
	}
	return *o.FlexibleYearlyInterestRate
}

// GetFlexibleYearlyInterestRateOk returns a tuple with the FlexibleYearlyInterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanableAssetsDataResponseRowsInner) GetFlexibleYearlyInterestRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.FlexibleYearlyInterestRate) {
		return nil, false
	}
	return o.FlexibleYearlyInterestRate, true
}

// HasFlexibleYearlyInterestRate returns a boolean if a field has been set.
func (o *GetLoanableAssetsDataResponseRowsInner) HasFlexibleYearlyInterestRate() bool {
	if o != nil && !common.IsNil(o.FlexibleYearlyInterestRate) {
		return true
	}

	return false
}

// SetFlexibleYearlyInterestRate gets a reference to the given string and assigns it to the FlexibleYearlyInterestRate field.
func (o *GetLoanableAssetsDataResponseRowsInner) SetFlexibleYearlyInterestRate(v string) {
	o.FlexibleYearlyInterestRate = &v
}

// GetVar30dDailyInterestRate returns the Var30dDailyInterestRate field value if set, zero value otherwise.
func (o *GetLoanableAssetsDataResponseRowsInner) GetVar30dDailyInterestRate() string {
	if o == nil || common.IsNil(o.Var30dDailyInterestRate) {
		var ret string
		return ret
	}
	return *o.Var30dDailyInterestRate
}

// GetVar30dDailyInterestRateOk returns a tuple with the Var30dDailyInterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanableAssetsDataResponseRowsInner) GetVar30dDailyInterestRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.Var30dDailyInterestRate) {
		return nil, false
	}
	return o.Var30dDailyInterestRate, true
}

// HasVar30dDailyInterestRate returns a boolean if a field has been set.
func (o *GetLoanableAssetsDataResponseRowsInner) HasVar30dDailyInterestRate() bool {
	if o != nil && !common.IsNil(o.Var30dDailyInterestRate) {
		return true
	}

	return false
}

// SetVar30dDailyInterestRate gets a reference to the given string and assigns it to the Var30dDailyInterestRate field.
func (o *GetLoanableAssetsDataResponseRowsInner) SetVar30dDailyInterestRate(v string) {
	o.Var30dDailyInterestRate = &v
}

// GetVar30dYearlyInterestRate returns the Var30dYearlyInterestRate field value if set, zero value otherwise.
func (o *GetLoanableAssetsDataResponseRowsInner) GetVar30dYearlyInterestRate() string {
	if o == nil || common.IsNil(o.Var30dYearlyInterestRate) {
		var ret string
		return ret
	}
	return *o.Var30dYearlyInterestRate
}

// GetVar30dYearlyInterestRateOk returns a tuple with the Var30dYearlyInterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanableAssetsDataResponseRowsInner) GetVar30dYearlyInterestRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.Var30dYearlyInterestRate) {
		return nil, false
	}
	return o.Var30dYearlyInterestRate, true
}

// HasVar30dYearlyInterestRate returns a boolean if a field has been set.
func (o *GetLoanableAssetsDataResponseRowsInner) HasVar30dYearlyInterestRate() bool {
	if o != nil && !common.IsNil(o.Var30dYearlyInterestRate) {
		return true
	}

	return false
}

// SetVar30dYearlyInterestRate gets a reference to the given string and assigns it to the Var30dYearlyInterestRate field.
func (o *GetLoanableAssetsDataResponseRowsInner) SetVar30dYearlyInterestRate(v string) {
	o.Var30dYearlyInterestRate = &v
}

// GetVar60dDailyInterestRate returns the Var60dDailyInterestRate field value if set, zero value otherwise.
func (o *GetLoanableAssetsDataResponseRowsInner) GetVar60dDailyInterestRate() string {
	if o == nil || common.IsNil(o.Var60dDailyInterestRate) {
		var ret string
		return ret
	}
	return *o.Var60dDailyInterestRate
}

// GetVar60dDailyInterestRateOk returns a tuple with the Var60dDailyInterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanableAssetsDataResponseRowsInner) GetVar60dDailyInterestRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.Var60dDailyInterestRate) {
		return nil, false
	}
	return o.Var60dDailyInterestRate, true
}

// HasVar60dDailyInterestRate returns a boolean if a field has been set.
func (o *GetLoanableAssetsDataResponseRowsInner) HasVar60dDailyInterestRate() bool {
	if o != nil && !common.IsNil(o.Var60dDailyInterestRate) {
		return true
	}

	return false
}

// SetVar60dDailyInterestRate gets a reference to the given string and assigns it to the Var60dDailyInterestRate field.
func (o *GetLoanableAssetsDataResponseRowsInner) SetVar60dDailyInterestRate(v string) {
	o.Var60dDailyInterestRate = &v
}

// GetVar60dYearlyInterestRate returns the Var60dYearlyInterestRate field value if set, zero value otherwise.
func (o *GetLoanableAssetsDataResponseRowsInner) GetVar60dYearlyInterestRate() string {
	if o == nil || common.IsNil(o.Var60dYearlyInterestRate) {
		var ret string
		return ret
	}
	return *o.Var60dYearlyInterestRate
}

// GetVar60dYearlyInterestRateOk returns a tuple with the Var60dYearlyInterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanableAssetsDataResponseRowsInner) GetVar60dYearlyInterestRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.Var60dYearlyInterestRate) {
		return nil, false
	}
	return o.Var60dYearlyInterestRate, true
}

// HasVar60dYearlyInterestRate returns a boolean if a field has been set.
func (o *GetLoanableAssetsDataResponseRowsInner) HasVar60dYearlyInterestRate() bool {
	if o != nil && !common.IsNil(o.Var60dYearlyInterestRate) {
		return true
	}

	return false
}

// SetVar60dYearlyInterestRate gets a reference to the given string and assigns it to the Var60dYearlyInterestRate field.
func (o *GetLoanableAssetsDataResponseRowsInner) SetVar60dYearlyInterestRate(v string) {
	o.Var60dYearlyInterestRate = &v
}

// GetMinLimit returns the MinLimit field value if set, zero value otherwise.
func (o *GetLoanableAssetsDataResponseRowsInner) GetMinLimit() string {
	if o == nil || common.IsNil(o.MinLimit) {
		var ret string
		return ret
	}
	return *o.MinLimit
}

// GetMinLimitOk returns a tuple with the MinLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanableAssetsDataResponseRowsInner) GetMinLimitOk() (*string, bool) {
	if o == nil || common.IsNil(o.MinLimit) {
		return nil, false
	}
	return o.MinLimit, true
}

// HasMinLimit returns a boolean if a field has been set.
func (o *GetLoanableAssetsDataResponseRowsInner) HasMinLimit() bool {
	if o != nil && !common.IsNil(o.MinLimit) {
		return true
	}

	return false
}

// SetMinLimit gets a reference to the given string and assigns it to the MinLimit field.
func (o *GetLoanableAssetsDataResponseRowsInner) SetMinLimit(v string) {
	o.MinLimit = &v
}

// GetMaxLimit returns the MaxLimit field value if set, zero value otherwise.
func (o *GetLoanableAssetsDataResponseRowsInner) GetMaxLimit() string {
	if o == nil || common.IsNil(o.MaxLimit) {
		var ret string
		return ret
	}
	return *o.MaxLimit
}

// GetMaxLimitOk returns a tuple with the MaxLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanableAssetsDataResponseRowsInner) GetMaxLimitOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxLimit) {
		return nil, false
	}
	return o.MaxLimit, true
}

// HasMaxLimit returns a boolean if a field has been set.
func (o *GetLoanableAssetsDataResponseRowsInner) HasMaxLimit() bool {
	if o != nil && !common.IsNil(o.MaxLimit) {
		return true
	}

	return false
}

// SetMaxLimit gets a reference to the given string and assigns it to the MaxLimit field.
func (o *GetLoanableAssetsDataResponseRowsInner) SetMaxLimit(v string) {
	o.MaxLimit = &v
}

// GetVipLevel returns the VipLevel field value if set, zero value otherwise.
func (o *GetLoanableAssetsDataResponseRowsInner) GetVipLevel() int64 {
	if o == nil || common.IsNil(o.VipLevel) {
		var ret int64
		return ret
	}
	return *o.VipLevel
}

// GetVipLevelOk returns a tuple with the VipLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanableAssetsDataResponseRowsInner) GetVipLevelOk() (*int64, bool) {
	if o == nil || common.IsNil(o.VipLevel) {
		return nil, false
	}
	return o.VipLevel, true
}

// HasVipLevel returns a boolean if a field has been set.
func (o *GetLoanableAssetsDataResponseRowsInner) HasVipLevel() bool {
	if o != nil && !common.IsNil(o.VipLevel) {
		return true
	}

	return false
}

// SetVipLevel gets a reference to the given int64 and assigns it to the VipLevel field.
func (o *GetLoanableAssetsDataResponseRowsInner) SetVipLevel(v int64) {
	o.VipLevel = &v
}

func (o GetLoanableAssetsDataResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLoanableAssetsDataResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.LoanCoin) {
		toSerialize["loanCoin"] = o.LoanCoin
	}
	if !common.IsNil(o.FlexibleDailyInterestRate) {
		toSerialize["_flexibleDailyInterestRate"] = o.FlexibleDailyInterestRate
	}
	if !common.IsNil(o.FlexibleYearlyInterestRate) {
		toSerialize["_flexibleYearlyInterestRate"] = o.FlexibleYearlyInterestRate
	}
	if !common.IsNil(o.Var30dDailyInterestRate) {
		toSerialize["_30dDailyInterestRate"] = o.Var30dDailyInterestRate
	}
	if !common.IsNil(o.Var30dYearlyInterestRate) {
		toSerialize["_30dYearlyInterestRate"] = o.Var30dYearlyInterestRate
	}
	if !common.IsNil(o.Var60dDailyInterestRate) {
		toSerialize["_60dDailyInterestRate"] = o.Var60dDailyInterestRate
	}
	if !common.IsNil(o.Var60dYearlyInterestRate) {
		toSerialize["_60dYearlyInterestRate"] = o.Var60dYearlyInterestRate
	}
	if !common.IsNil(o.MinLimit) {
		toSerialize["minLimit"] = o.MinLimit
	}
	if !common.IsNil(o.MaxLimit) {
		toSerialize["maxLimit"] = o.MaxLimit
	}
	if !common.IsNil(o.VipLevel) {
		toSerialize["vipLevel"] = o.VipLevel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetLoanableAssetsDataResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetLoanableAssetsDataResponseRowsInner := _GetLoanableAssetsDataResponseRowsInner{}

	err = json.Unmarshal(data, &varGetLoanableAssetsDataResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetLoanableAssetsDataResponseRowsInner(varGetLoanableAssetsDataResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "loanCoin")
		delete(additionalProperties, "_flexibleDailyInterestRate")
		delete(additionalProperties, "_flexibleYearlyInterestRate")
		delete(additionalProperties, "_30dDailyInterestRate")
		delete(additionalProperties, "_30dYearlyInterestRate")
		delete(additionalProperties, "_60dDailyInterestRate")
		delete(additionalProperties, "_60dYearlyInterestRate")
		delete(additionalProperties, "minLimit")
		delete(additionalProperties, "maxLimit")
		delete(additionalProperties, "vipLevel")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetLoanableAssetsDataResponseRowsInner struct {
	value *GetLoanableAssetsDataResponseRowsInner
	isSet bool
}

func (v NullableGetLoanableAssetsDataResponseRowsInner) Get() *GetLoanableAssetsDataResponseRowsInner {
	return v.value
}

func (v *NullableGetLoanableAssetsDataResponseRowsInner) Set(val *GetLoanableAssetsDataResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLoanableAssetsDataResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLoanableAssetsDataResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLoanableAssetsDataResponseRowsInner(val *GetLoanableAssetsDataResponseRowsInner) *NullableGetLoanableAssetsDataResponseRowsInner {
	return &NullableGetLoanableAssetsDataResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetLoanableAssetsDataResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLoanableAssetsDataResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
