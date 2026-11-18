/*
Portfolio Margin Pro REST API

Access advanced account management and high-frequency trading with Binance Portfolio Margin Pro.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner{}

// GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner struct for GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner
type GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner struct {
	// tier Floor.
	TierFloor *string `json:"tierFloor,omitempty"`
	// tier Cap.
	TierCap *string `json:"tierCap,omitempty"`
	// collateral Rate.
	CollateralRate *string `json:"collateralRate,omitempty"`
	// account equity quick addition number.
	Cum                  *string `json:"cum,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner

// NewGetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner instantiates a new GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner() *GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner {
	this := GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner{}
	return &this
}

// NewGetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInnerWithDefaults instantiates a new GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInnerWithDefaults() *GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner {
	this := GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner{}
	return &this
}

// GetTierFloor returns the TierFloor field value if set, zero value otherwise.
func (o *GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner) GetTierFloor() string {
	if o == nil || common.IsNil(o.TierFloor) {
		var ret string
		return ret
	}
	return *o.TierFloor
}

// GetTierFloorOk returns a tuple with the TierFloor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner) GetTierFloorOk() (*string, bool) {
	if o == nil || common.IsNil(o.TierFloor) {
		return nil, false
	}
	return o.TierFloor, true
}

// HasTierFloor returns a boolean if a field has been set.
func (o *GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner) HasTierFloor() bool {
	if o != nil && !common.IsNil(o.TierFloor) {
		return true
	}

	return false
}

// SetTierFloor gets a reference to the given string and assigns it to the TierFloor field.
func (o *GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner) SetTierFloor(v string) {
	o.TierFloor = &v
}

// GetTierCap returns the TierCap field value if set, zero value otherwise.
func (o *GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner) GetTierCap() string {
	if o == nil || common.IsNil(o.TierCap) {
		var ret string
		return ret
	}
	return *o.TierCap
}

// GetTierCapOk returns a tuple with the TierCap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner) GetTierCapOk() (*string, bool) {
	if o == nil || common.IsNil(o.TierCap) {
		return nil, false
	}
	return o.TierCap, true
}

// HasTierCap returns a boolean if a field has been set.
func (o *GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner) HasTierCap() bool {
	if o != nil && !common.IsNil(o.TierCap) {
		return true
	}

	return false
}

// SetTierCap gets a reference to the given string and assigns it to the TierCap field.
func (o *GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner) SetTierCap(v string) {
	o.TierCap = &v
}

// GetCollateralRate returns the CollateralRate field value if set, zero value otherwise.
func (o *GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner) GetCollateralRate() string {
	if o == nil || common.IsNil(o.CollateralRate) {
		var ret string
		return ret
	}
	return *o.CollateralRate
}

// GetCollateralRateOk returns a tuple with the CollateralRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner) GetCollateralRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralRate) {
		return nil, false
	}
	return o.CollateralRate, true
}

// HasCollateralRate returns a boolean if a field has been set.
func (o *GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner) HasCollateralRate() bool {
	if o != nil && !common.IsNil(o.CollateralRate) {
		return true
	}

	return false
}

// SetCollateralRate gets a reference to the given string and assigns it to the CollateralRate field.
func (o *GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner) SetCollateralRate(v string) {
	o.CollateralRate = &v
}

// GetCum returns the Cum field value if set, zero value otherwise.
func (o *GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner) GetCum() string {
	if o == nil || common.IsNil(o.Cum) {
		var ret string
		return ret
	}
	return *o.Cum
}

// GetCumOk returns a tuple with the Cum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner) GetCumOk() (*string, bool) {
	if o == nil || common.IsNil(o.Cum) {
		return nil, false
	}
	return o.Cum, true
}

// HasCum returns a boolean if a field has been set.
func (o *GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner) HasCum() bool {
	if o != nil && !common.IsNil(o.Cum) {
		return true
	}

	return false
}

// SetCum gets a reference to the given string and assigns it to the Cum field.
func (o *GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner) SetCum(v string) {
	o.Cum = &v
}

func (o GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TierFloor) {
		toSerialize["tierFloor"] = o.TierFloor
	}
	if !common.IsNil(o.TierCap) {
		toSerialize["tierCap"] = o.TierCap
	}
	if !common.IsNil(o.CollateralRate) {
		toSerialize["collateralRate"] = o.CollateralRate
	}
	if !common.IsNil(o.Cum) {
		toSerialize["cum"] = o.Cum
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner) UnmarshalJSON(data []byte) (err error) {
	varGetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner := _GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner{}

	err = json.Unmarshal(data, &varGetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner)

	if err != nil {
		return err
	}

	*o = GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner(varGetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tierFloor")
		delete(additionalProperties, "tierCap")
		delete(additionalProperties, "collateralRate")
		delete(additionalProperties, "cum")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner struct {
	value *GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner
	isSet bool
}

func (v NullableGetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner) Get() *GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner {
	return v.value
}

func (v *NullableGetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner) Set(val *GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner(val *GetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner) *NullableGetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner {
	return &NullableGetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner{value: val, isSet: true}
}

func (v NullableGetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPortfolioMarginTieredCollateralRatesResponseInnerCollateralInfoInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
