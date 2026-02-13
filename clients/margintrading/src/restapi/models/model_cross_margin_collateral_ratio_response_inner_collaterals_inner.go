/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the CrossMarginCollateralRatioResponseInnerCollateralsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CrossMarginCollateralRatioResponseInnerCollateralsInner{}

// CrossMarginCollateralRatioResponseInnerCollateralsInner struct for CrossMarginCollateralRatioResponseInnerCollateralsInner
type CrossMarginCollateralRatioResponseInnerCollateralsInner struct {
	MinUsdValue          *string `json:"minUsdValue,omitempty"`
	MaxUsdValue          *string `json:"maxUsdValue,omitempty"`
	DiscountRate         *string `json:"discountRate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CrossMarginCollateralRatioResponseInnerCollateralsInner CrossMarginCollateralRatioResponseInnerCollateralsInner

// NewCrossMarginCollateralRatioResponseInnerCollateralsInner instantiates a new CrossMarginCollateralRatioResponseInnerCollateralsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCrossMarginCollateralRatioResponseInnerCollateralsInner() *CrossMarginCollateralRatioResponseInnerCollateralsInner {
	this := CrossMarginCollateralRatioResponseInnerCollateralsInner{}
	return &this
}

// NewCrossMarginCollateralRatioResponseInnerCollateralsInnerWithDefaults instantiates a new CrossMarginCollateralRatioResponseInnerCollateralsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCrossMarginCollateralRatioResponseInnerCollateralsInnerWithDefaults() *CrossMarginCollateralRatioResponseInnerCollateralsInner {
	this := CrossMarginCollateralRatioResponseInnerCollateralsInner{}
	return &this
}

// GetMinUsdValue returns the MinUsdValue field value if set, zero value otherwise.
func (o *CrossMarginCollateralRatioResponseInnerCollateralsInner) GetMinUsdValue() string {
	if o == nil || common.IsNil(o.MinUsdValue) {
		var ret string
		return ret
	}
	return *o.MinUsdValue
}

// GetMinUsdValueOk returns a tuple with the MinUsdValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrossMarginCollateralRatioResponseInnerCollateralsInner) GetMinUsdValueOk() (*string, bool) {
	if o == nil || common.IsNil(o.MinUsdValue) {
		return nil, false
	}
	return o.MinUsdValue, true
}

// HasMinUsdValue returns a boolean if a field has been set.
func (o *CrossMarginCollateralRatioResponseInnerCollateralsInner) HasMinUsdValue() bool {
	if o != nil && !common.IsNil(o.MinUsdValue) {
		return true
	}

	return false
}

// SetMinUsdValue gets a reference to the given string and assigns it to the MinUsdValue field.
func (o *CrossMarginCollateralRatioResponseInnerCollateralsInner) SetMinUsdValue(v string) {
	o.MinUsdValue = &v
}

// GetMaxUsdValue returns the MaxUsdValue field value if set, zero value otherwise.
func (o *CrossMarginCollateralRatioResponseInnerCollateralsInner) GetMaxUsdValue() string {
	if o == nil || common.IsNil(o.MaxUsdValue) {
		var ret string
		return ret
	}
	return *o.MaxUsdValue
}

// GetMaxUsdValueOk returns a tuple with the MaxUsdValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrossMarginCollateralRatioResponseInnerCollateralsInner) GetMaxUsdValueOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxUsdValue) {
		return nil, false
	}
	return o.MaxUsdValue, true
}

// HasMaxUsdValue returns a boolean if a field has been set.
func (o *CrossMarginCollateralRatioResponseInnerCollateralsInner) HasMaxUsdValue() bool {
	if o != nil && !common.IsNil(o.MaxUsdValue) {
		return true
	}

	return false
}

// SetMaxUsdValue gets a reference to the given string and assigns it to the MaxUsdValue field.
func (o *CrossMarginCollateralRatioResponseInnerCollateralsInner) SetMaxUsdValue(v string) {
	o.MaxUsdValue = &v
}

// GetDiscountRate returns the DiscountRate field value if set, zero value otherwise.
func (o *CrossMarginCollateralRatioResponseInnerCollateralsInner) GetDiscountRate() string {
	if o == nil || common.IsNil(o.DiscountRate) {
		var ret string
		return ret
	}
	return *o.DiscountRate
}

// GetDiscountRateOk returns a tuple with the DiscountRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrossMarginCollateralRatioResponseInnerCollateralsInner) GetDiscountRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.DiscountRate) {
		return nil, false
	}
	return o.DiscountRate, true
}

// HasDiscountRate returns a boolean if a field has been set.
func (o *CrossMarginCollateralRatioResponseInnerCollateralsInner) HasDiscountRate() bool {
	if o != nil && !common.IsNil(o.DiscountRate) {
		return true
	}

	return false
}

// SetDiscountRate gets a reference to the given string and assigns it to the DiscountRate field.
func (o *CrossMarginCollateralRatioResponseInnerCollateralsInner) SetDiscountRate(v string) {
	o.DiscountRate = &v
}

func (o CrossMarginCollateralRatioResponseInnerCollateralsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CrossMarginCollateralRatioResponseInnerCollateralsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.MinUsdValue) {
		toSerialize["minUsdValue"] = o.MinUsdValue
	}
	if !common.IsNil(o.MaxUsdValue) {
		toSerialize["maxUsdValue"] = o.MaxUsdValue
	}
	if !common.IsNil(o.DiscountRate) {
		toSerialize["discountRate"] = o.DiscountRate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CrossMarginCollateralRatioResponseInnerCollateralsInner) UnmarshalJSON(data []byte) (err error) {
	varCrossMarginCollateralRatioResponseInnerCollateralsInner := _CrossMarginCollateralRatioResponseInnerCollateralsInner{}

	err = json.Unmarshal(data, &varCrossMarginCollateralRatioResponseInnerCollateralsInner)

	if err != nil {
		return err
	}

	*o = CrossMarginCollateralRatioResponseInnerCollateralsInner(varCrossMarginCollateralRatioResponseInnerCollateralsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "minUsdValue")
		delete(additionalProperties, "maxUsdValue")
		delete(additionalProperties, "discountRate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCrossMarginCollateralRatioResponseInnerCollateralsInner struct {
	value *CrossMarginCollateralRatioResponseInnerCollateralsInner
	isSet bool
}

func (v NullableCrossMarginCollateralRatioResponseInnerCollateralsInner) Get() *CrossMarginCollateralRatioResponseInnerCollateralsInner {
	return v.value
}

func (v *NullableCrossMarginCollateralRatioResponseInnerCollateralsInner) Set(val *CrossMarginCollateralRatioResponseInnerCollateralsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCrossMarginCollateralRatioResponseInnerCollateralsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCrossMarginCollateralRatioResponseInnerCollateralsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCrossMarginCollateralRatioResponseInnerCollateralsInner(val *CrossMarginCollateralRatioResponseInnerCollateralsInner) *NullableCrossMarginCollateralRatioResponseInnerCollateralsInner {
	return &NullableCrossMarginCollateralRatioResponseInnerCollateralsInner{value: val, isSet: true}
}

func (v NullableCrossMarginCollateralRatioResponseInnerCollateralsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCrossMarginCollateralRatioResponseInnerCollateralsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
