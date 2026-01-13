/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the UmPositionAdlQuantileEstimationResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UmPositionAdlQuantileEstimationResponseInner{}

// UmPositionAdlQuantileEstimationResponseInner struct for UmPositionAdlQuantileEstimationResponseInner
type UmPositionAdlQuantileEstimationResponseInner struct {
	Symbol               *string                                                  `json:"symbol,omitempty"`
	AdlQuantile          *UmPositionAdlQuantileEstimationResponseInnerAdlQuantile `json:"adlQuantile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UmPositionAdlQuantileEstimationResponseInner UmPositionAdlQuantileEstimationResponseInner

// NewUmPositionAdlQuantileEstimationResponseInner instantiates a new UmPositionAdlQuantileEstimationResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUmPositionAdlQuantileEstimationResponseInner() *UmPositionAdlQuantileEstimationResponseInner {
	this := UmPositionAdlQuantileEstimationResponseInner{}
	return &this
}

// NewUmPositionAdlQuantileEstimationResponseInnerWithDefaults instantiates a new UmPositionAdlQuantileEstimationResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUmPositionAdlQuantileEstimationResponseInnerWithDefaults() *UmPositionAdlQuantileEstimationResponseInner {
	this := UmPositionAdlQuantileEstimationResponseInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *UmPositionAdlQuantileEstimationResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmPositionAdlQuantileEstimationResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *UmPositionAdlQuantileEstimationResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *UmPositionAdlQuantileEstimationResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetAdlQuantile returns the AdlQuantile field value if set, zero value otherwise.
func (o *UmPositionAdlQuantileEstimationResponseInner) GetAdlQuantile() UmPositionAdlQuantileEstimationResponseInnerAdlQuantile {
	if o == nil || common.IsNil(o.AdlQuantile) {
		var ret UmPositionAdlQuantileEstimationResponseInnerAdlQuantile
		return ret
	}
	return *o.AdlQuantile
}

// GetAdlQuantileOk returns a tuple with the AdlQuantile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmPositionAdlQuantileEstimationResponseInner) GetAdlQuantileOk() (*UmPositionAdlQuantileEstimationResponseInnerAdlQuantile, bool) {
	if o == nil || common.IsNil(o.AdlQuantile) {
		return nil, false
	}
	return o.AdlQuantile, true
}

// HasAdlQuantile returns a boolean if a field has been set.
func (o *UmPositionAdlQuantileEstimationResponseInner) HasAdlQuantile() bool {
	if o != nil && !common.IsNil(o.AdlQuantile) {
		return true
	}

	return false
}

// SetAdlQuantile gets a reference to the given UmPositionAdlQuantileEstimationResponseInnerAdlQuantile and assigns it to the AdlQuantile field.
func (o *UmPositionAdlQuantileEstimationResponseInner) SetAdlQuantile(v UmPositionAdlQuantileEstimationResponseInnerAdlQuantile) {
	o.AdlQuantile = &v
}

func (o UmPositionAdlQuantileEstimationResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UmPositionAdlQuantileEstimationResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.AdlQuantile) {
		toSerialize["adlQuantile"] = o.AdlQuantile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UmPositionAdlQuantileEstimationResponseInner) UnmarshalJSON(data []byte) (err error) {
	varUmPositionAdlQuantileEstimationResponseInner := _UmPositionAdlQuantileEstimationResponseInner{}

	err = json.Unmarshal(data, &varUmPositionAdlQuantileEstimationResponseInner)

	if err != nil {
		return err
	}

	*o = UmPositionAdlQuantileEstimationResponseInner(varUmPositionAdlQuantileEstimationResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "adlQuantile")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUmPositionAdlQuantileEstimationResponseInner struct {
	value *UmPositionAdlQuantileEstimationResponseInner
	isSet bool
}

func (v NullableUmPositionAdlQuantileEstimationResponseInner) Get() *UmPositionAdlQuantileEstimationResponseInner {
	return v.value
}

func (v *NullableUmPositionAdlQuantileEstimationResponseInner) Set(val *UmPositionAdlQuantileEstimationResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUmPositionAdlQuantileEstimationResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUmPositionAdlQuantileEstimationResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUmPositionAdlQuantileEstimationResponseInner(val *UmPositionAdlQuantileEstimationResponseInner) *NullableUmPositionAdlQuantileEstimationResponseInner {
	return &NullableUmPositionAdlQuantileEstimationResponseInner{value: val, isSet: true}
}

func (v NullableUmPositionAdlQuantileEstimationResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUmPositionAdlQuantileEstimationResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
