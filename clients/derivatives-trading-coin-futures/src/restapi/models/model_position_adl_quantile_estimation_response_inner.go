/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the PositionAdlQuantileEstimationResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PositionAdlQuantileEstimationResponseInner{}

// PositionAdlQuantileEstimationResponseInner struct for PositionAdlQuantileEstimationResponseInner
type PositionAdlQuantileEstimationResponseInner struct {
	Symbol               *string                                                `json:"symbol,omitempty"`
	AdlQuantile          *PositionAdlQuantileEstimationResponseInnerAdlQuantile `json:"adlQuantile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PositionAdlQuantileEstimationResponseInner PositionAdlQuantileEstimationResponseInner

// NewPositionAdlQuantileEstimationResponseInner instantiates a new PositionAdlQuantileEstimationResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPositionAdlQuantileEstimationResponseInner() *PositionAdlQuantileEstimationResponseInner {
	this := PositionAdlQuantileEstimationResponseInner{}
	return &this
}

// NewPositionAdlQuantileEstimationResponseInnerWithDefaults instantiates a new PositionAdlQuantileEstimationResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPositionAdlQuantileEstimationResponseInnerWithDefaults() *PositionAdlQuantileEstimationResponseInner {
	this := PositionAdlQuantileEstimationResponseInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *PositionAdlQuantileEstimationResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionAdlQuantileEstimationResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *PositionAdlQuantileEstimationResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *PositionAdlQuantileEstimationResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetAdlQuantile returns the AdlQuantile field value if set, zero value otherwise.
func (o *PositionAdlQuantileEstimationResponseInner) GetAdlQuantile() PositionAdlQuantileEstimationResponseInnerAdlQuantile {
	if o == nil || common.IsNil(o.AdlQuantile) {
		var ret PositionAdlQuantileEstimationResponseInnerAdlQuantile
		return ret
	}
	return *o.AdlQuantile
}

// GetAdlQuantileOk returns a tuple with the AdlQuantile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionAdlQuantileEstimationResponseInner) GetAdlQuantileOk() (*PositionAdlQuantileEstimationResponseInnerAdlQuantile, bool) {
	if o == nil || common.IsNil(o.AdlQuantile) {
		return nil, false
	}
	return o.AdlQuantile, true
}

// HasAdlQuantile returns a boolean if a field has been set.
func (o *PositionAdlQuantileEstimationResponseInner) HasAdlQuantile() bool {
	if o != nil && !common.IsNil(o.AdlQuantile) {
		return true
	}

	return false
}

// SetAdlQuantile gets a reference to the given PositionAdlQuantileEstimationResponseInnerAdlQuantile and assigns it to the AdlQuantile field.
func (o *PositionAdlQuantileEstimationResponseInner) SetAdlQuantile(v PositionAdlQuantileEstimationResponseInnerAdlQuantile) {
	o.AdlQuantile = &v
}

func (o PositionAdlQuantileEstimationResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PositionAdlQuantileEstimationResponseInner) ToMap() (map[string]interface{}, error) {
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

func (o *PositionAdlQuantileEstimationResponseInner) UnmarshalJSON(data []byte) (err error) {
	varPositionAdlQuantileEstimationResponseInner := _PositionAdlQuantileEstimationResponseInner{}

	err = json.Unmarshal(data, &varPositionAdlQuantileEstimationResponseInner)

	if err != nil {
		return err
	}

	*o = PositionAdlQuantileEstimationResponseInner(varPositionAdlQuantileEstimationResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "adlQuantile")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePositionAdlQuantileEstimationResponseInner struct {
	value *PositionAdlQuantileEstimationResponseInner
	isSet bool
}

func (v NullablePositionAdlQuantileEstimationResponseInner) Get() *PositionAdlQuantileEstimationResponseInner {
	return v.value
}

func (v *NullablePositionAdlQuantileEstimationResponseInner) Set(val *PositionAdlQuantileEstimationResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePositionAdlQuantileEstimationResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePositionAdlQuantileEstimationResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePositionAdlQuantileEstimationResponseInner(val *PositionAdlQuantileEstimationResponseInner) *NullablePositionAdlQuantileEstimationResponseInner {
	return &NullablePositionAdlQuantileEstimationResponseInner{value: val, isSet: true}
}

func (v NullablePositionAdlQuantileEstimationResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePositionAdlQuantileEstimationResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
