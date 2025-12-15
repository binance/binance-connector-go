/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OptionAccountInformationResponseGreekInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OptionAccountInformationResponseGreekInner{}

// OptionAccountInformationResponseGreekInner struct for OptionAccountInformationResponseGreekInner
type OptionAccountInformationResponseGreekInner struct {
	Underlying           *string `json:"underlying,omitempty"`
	Delta                *string `json:"delta,omitempty"`
	Gamma                *string `json:"gamma,omitempty"`
	Theta                *string `json:"theta,omitempty"`
	Vega                 *string `json:"vega,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OptionAccountInformationResponseGreekInner OptionAccountInformationResponseGreekInner

// NewOptionAccountInformationResponseGreekInner instantiates a new OptionAccountInformationResponseGreekInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionAccountInformationResponseGreekInner() *OptionAccountInformationResponseGreekInner {
	this := OptionAccountInformationResponseGreekInner{}
	return &this
}

// NewOptionAccountInformationResponseGreekInnerWithDefaults instantiates a new OptionAccountInformationResponseGreekInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionAccountInformationResponseGreekInnerWithDefaults() *OptionAccountInformationResponseGreekInner {
	this := OptionAccountInformationResponseGreekInner{}
	return &this
}

// GetUnderlying returns the Underlying field value if set, zero value otherwise.
func (o *OptionAccountInformationResponseGreekInner) GetUnderlying() string {
	if o == nil || common.IsNil(o.Underlying) {
		var ret string
		return ret
	}
	return *o.Underlying
}

// GetUnderlyingOk returns a tuple with the Underlying field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionAccountInformationResponseGreekInner) GetUnderlyingOk() (*string, bool) {
	if o == nil || common.IsNil(o.Underlying) {
		return nil, false
	}
	return o.Underlying, true
}

// HasUnderlying returns a boolean if a field has been set.
func (o *OptionAccountInformationResponseGreekInner) HasUnderlying() bool {
	if o != nil && !common.IsNil(o.Underlying) {
		return true
	}

	return false
}

// SetUnderlying gets a reference to the given string and assigns it to the Underlying field.
func (o *OptionAccountInformationResponseGreekInner) SetUnderlying(v string) {
	o.Underlying = &v
}

// GetDelta returns the Delta field value if set, zero value otherwise.
func (o *OptionAccountInformationResponseGreekInner) GetDelta() string {
	if o == nil || common.IsNil(o.Delta) {
		var ret string
		return ret
	}
	return *o.Delta
}

// GetDeltaOk returns a tuple with the Delta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionAccountInformationResponseGreekInner) GetDeltaOk() (*string, bool) {
	if o == nil || common.IsNil(o.Delta) {
		return nil, false
	}
	return o.Delta, true
}

// HasDelta returns a boolean if a field has been set.
func (o *OptionAccountInformationResponseGreekInner) HasDelta() bool {
	if o != nil && !common.IsNil(o.Delta) {
		return true
	}

	return false
}

// SetDelta gets a reference to the given string and assigns it to the Delta field.
func (o *OptionAccountInformationResponseGreekInner) SetDelta(v string) {
	o.Delta = &v
}

// GetGamma returns the Gamma field value if set, zero value otherwise.
func (o *OptionAccountInformationResponseGreekInner) GetGamma() string {
	if o == nil || common.IsNil(o.Gamma) {
		var ret string
		return ret
	}
	return *o.Gamma
}

// GetGammaOk returns a tuple with the Gamma field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionAccountInformationResponseGreekInner) GetGammaOk() (*string, bool) {
	if o == nil || common.IsNil(o.Gamma) {
		return nil, false
	}
	return o.Gamma, true
}

// HasGamma returns a boolean if a field has been set.
func (o *OptionAccountInformationResponseGreekInner) HasGamma() bool {
	if o != nil && !common.IsNil(o.Gamma) {
		return true
	}

	return false
}

// SetGamma gets a reference to the given string and assigns it to the Gamma field.
func (o *OptionAccountInformationResponseGreekInner) SetGamma(v string) {
	o.Gamma = &v
}

// GetTheta returns the Theta field value if set, zero value otherwise.
func (o *OptionAccountInformationResponseGreekInner) GetTheta() string {
	if o == nil || common.IsNil(o.Theta) {
		var ret string
		return ret
	}
	return *o.Theta
}

// GetThetaOk returns a tuple with the Theta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionAccountInformationResponseGreekInner) GetThetaOk() (*string, bool) {
	if o == nil || common.IsNil(o.Theta) {
		return nil, false
	}
	return o.Theta, true
}

// HasTheta returns a boolean if a field has been set.
func (o *OptionAccountInformationResponseGreekInner) HasTheta() bool {
	if o != nil && !common.IsNil(o.Theta) {
		return true
	}

	return false
}

// SetTheta gets a reference to the given string and assigns it to the Theta field.
func (o *OptionAccountInformationResponseGreekInner) SetTheta(v string) {
	o.Theta = &v
}

// GetVega returns the Vega field value if set, zero value otherwise.
func (o *OptionAccountInformationResponseGreekInner) GetVega() string {
	if o == nil || common.IsNil(o.Vega) {
		var ret string
		return ret
	}
	return *o.Vega
}

// GetVegaOk returns a tuple with the Vega field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionAccountInformationResponseGreekInner) GetVegaOk() (*string, bool) {
	if o == nil || common.IsNil(o.Vega) {
		return nil, false
	}
	return o.Vega, true
}

// HasVega returns a boolean if a field has been set.
func (o *OptionAccountInformationResponseGreekInner) HasVega() bool {
	if o != nil && !common.IsNil(o.Vega) {
		return true
	}

	return false
}

// SetVega gets a reference to the given string and assigns it to the Vega field.
func (o *OptionAccountInformationResponseGreekInner) SetVega(v string) {
	o.Vega = &v
}

func (o OptionAccountInformationResponseGreekInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptionAccountInformationResponseGreekInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Underlying) {
		toSerialize["underlying"] = o.Underlying
	}
	if !common.IsNil(o.Delta) {
		toSerialize["delta"] = o.Delta
	}
	if !common.IsNil(o.Gamma) {
		toSerialize["gamma"] = o.Gamma
	}
	if !common.IsNil(o.Theta) {
		toSerialize["theta"] = o.Theta
	}
	if !common.IsNil(o.Vega) {
		toSerialize["vega"] = o.Vega
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OptionAccountInformationResponseGreekInner) UnmarshalJSON(data []byte) (err error) {
	varOptionAccountInformationResponseGreekInner := _OptionAccountInformationResponseGreekInner{}

	err = json.Unmarshal(data, &varOptionAccountInformationResponseGreekInner)

	if err != nil {
		return err
	}

	*o = OptionAccountInformationResponseGreekInner(varOptionAccountInformationResponseGreekInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "underlying")
		delete(additionalProperties, "delta")
		delete(additionalProperties, "gamma")
		delete(additionalProperties, "theta")
		delete(additionalProperties, "vega")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOptionAccountInformationResponseGreekInner struct {
	value *OptionAccountInformationResponseGreekInner
	isSet bool
}

func (v NullableOptionAccountInformationResponseGreekInner) Get() *OptionAccountInformationResponseGreekInner {
	return v.value
}

func (v *NullableOptionAccountInformationResponseGreekInner) Set(val *OptionAccountInformationResponseGreekInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionAccountInformationResponseGreekInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionAccountInformationResponseGreekInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionAccountInformationResponseGreekInner(val *OptionAccountInformationResponseGreekInner) *NullableOptionAccountInformationResponseGreekInner {
	return &NullableOptionAccountInformationResponseGreekInner{value: val, isSet: true}
}

func (v NullableOptionAccountInformationResponseGreekInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionAccountInformationResponseGreekInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
