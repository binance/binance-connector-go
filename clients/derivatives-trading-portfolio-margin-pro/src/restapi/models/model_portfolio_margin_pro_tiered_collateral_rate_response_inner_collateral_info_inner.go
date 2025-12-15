/*
Binance Derivatives Trading Portfolio Margin Pro REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin Pro REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner{}

// PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner struct for PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner
type PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner struct {
	TierFloor            *string `json:"tierFloor,omitempty"`
	TierCap              *string `json:"tierCap,omitempty"`
	CollateralRate       *string `json:"collateralRate,omitempty"`
	Cum                  *string `json:"cum,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner

// NewPortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner instantiates a new PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner() *PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner {
	this := PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner{}
	return &this
}

// NewPortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInnerWithDefaults instantiates a new PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInnerWithDefaults() *PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner {
	this := PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner{}
	return &this
}

// GetTierFloor returns the TierFloor field value if set, zero value otherwise.
func (o *PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner) GetTierFloor() string {
	if o == nil || common.IsNil(o.TierFloor) {
		var ret string
		return ret
	}
	return *o.TierFloor
}

// GetTierFloorOk returns a tuple with the TierFloor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner) GetTierFloorOk() (*string, bool) {
	if o == nil || common.IsNil(o.TierFloor) {
		return nil, false
	}
	return o.TierFloor, true
}

// HasTierFloor returns a boolean if a field has been set.
func (o *PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner) HasTierFloor() bool {
	if o != nil && !common.IsNil(o.TierFloor) {
		return true
	}

	return false
}

// SetTierFloor gets a reference to the given string and assigns it to the TierFloor field.
func (o *PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner) SetTierFloor(v string) {
	o.TierFloor = &v
}

// GetTierCap returns the TierCap field value if set, zero value otherwise.
func (o *PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner) GetTierCap() string {
	if o == nil || common.IsNil(o.TierCap) {
		var ret string
		return ret
	}
	return *o.TierCap
}

// GetTierCapOk returns a tuple with the TierCap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner) GetTierCapOk() (*string, bool) {
	if o == nil || common.IsNil(o.TierCap) {
		return nil, false
	}
	return o.TierCap, true
}

// HasTierCap returns a boolean if a field has been set.
func (o *PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner) HasTierCap() bool {
	if o != nil && !common.IsNil(o.TierCap) {
		return true
	}

	return false
}

// SetTierCap gets a reference to the given string and assigns it to the TierCap field.
func (o *PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner) SetTierCap(v string) {
	o.TierCap = &v
}

// GetCollateralRate returns the CollateralRate field value if set, zero value otherwise.
func (o *PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner) GetCollateralRate() string {
	if o == nil || common.IsNil(o.CollateralRate) {
		var ret string
		return ret
	}
	return *o.CollateralRate
}

// GetCollateralRateOk returns a tuple with the CollateralRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner) GetCollateralRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralRate) {
		return nil, false
	}
	return o.CollateralRate, true
}

// HasCollateralRate returns a boolean if a field has been set.
func (o *PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner) HasCollateralRate() bool {
	if o != nil && !common.IsNil(o.CollateralRate) {
		return true
	}

	return false
}

// SetCollateralRate gets a reference to the given string and assigns it to the CollateralRate field.
func (o *PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner) SetCollateralRate(v string) {
	o.CollateralRate = &v
}

// GetCum returns the Cum field value if set, zero value otherwise.
func (o *PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner) GetCum() string {
	if o == nil || common.IsNil(o.Cum) {
		var ret string
		return ret
	}
	return *o.Cum
}

// GetCumOk returns a tuple with the Cum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner) GetCumOk() (*string, bool) {
	if o == nil || common.IsNil(o.Cum) {
		return nil, false
	}
	return o.Cum, true
}

// HasCum returns a boolean if a field has been set.
func (o *PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner) HasCum() bool {
	if o != nil && !common.IsNil(o.Cum) {
		return true
	}

	return false
}

// SetCum gets a reference to the given string and assigns it to the Cum field.
func (o *PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner) SetCum(v string) {
	o.Cum = &v
}

func (o PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner) ToMap() (map[string]interface{}, error) {
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

func (o *PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner) UnmarshalJSON(data []byte) (err error) {
	varPortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner := _PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner{}

	err = json.Unmarshal(data, &varPortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner)

	if err != nil {
		return err
	}

	*o = PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner(varPortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner)

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

type NullablePortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner struct {
	value *PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner
	isSet bool
}

func (v NullablePortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner) Get() *PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner {
	return v.value
}

func (v *NullablePortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner) Set(val *PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner(val *PortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner) *NullablePortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner {
	return &NullablePortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner{value: val, isSet: true}
}

func (v NullablePortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortfolioMarginProTieredCollateralRateResponseInnerCollateralInfoInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
