/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the CmPositionAdlQuantileEstimationResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CmPositionAdlQuantileEstimationResponseInner{}

// CmPositionAdlQuantileEstimationResponseInner struct for CmPositionAdlQuantileEstimationResponseInner
type CmPositionAdlQuantileEstimationResponseInner struct {
	Symbol               *string                                                  `json:"symbol,omitempty"`
	AdlQuantile          *CmPositionAdlQuantileEstimationResponseInnerAdlQuantile `json:"adlQuantile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CmPositionAdlQuantileEstimationResponseInner CmPositionAdlQuantileEstimationResponseInner

// NewCmPositionAdlQuantileEstimationResponseInner instantiates a new CmPositionAdlQuantileEstimationResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCmPositionAdlQuantileEstimationResponseInner() *CmPositionAdlQuantileEstimationResponseInner {
	this := CmPositionAdlQuantileEstimationResponseInner{}
	return &this
}

// NewCmPositionAdlQuantileEstimationResponseInnerWithDefaults instantiates a new CmPositionAdlQuantileEstimationResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCmPositionAdlQuantileEstimationResponseInnerWithDefaults() *CmPositionAdlQuantileEstimationResponseInner {
	this := CmPositionAdlQuantileEstimationResponseInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *CmPositionAdlQuantileEstimationResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmPositionAdlQuantileEstimationResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *CmPositionAdlQuantileEstimationResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *CmPositionAdlQuantileEstimationResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetAdlQuantile returns the AdlQuantile field value if set, zero value otherwise.
func (o *CmPositionAdlQuantileEstimationResponseInner) GetAdlQuantile() CmPositionAdlQuantileEstimationResponseInnerAdlQuantile {
	if o == nil || common.IsNil(o.AdlQuantile) {
		var ret CmPositionAdlQuantileEstimationResponseInnerAdlQuantile
		return ret
	}
	return *o.AdlQuantile
}

// GetAdlQuantileOk returns a tuple with the AdlQuantile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmPositionAdlQuantileEstimationResponseInner) GetAdlQuantileOk() (*CmPositionAdlQuantileEstimationResponseInnerAdlQuantile, bool) {
	if o == nil || common.IsNil(o.AdlQuantile) {
		return nil, false
	}
	return o.AdlQuantile, true
}

// HasAdlQuantile returns a boolean if a field has been set.
func (o *CmPositionAdlQuantileEstimationResponseInner) HasAdlQuantile() bool {
	if o != nil && !common.IsNil(o.AdlQuantile) {
		return true
	}

	return false
}

// SetAdlQuantile gets a reference to the given CmPositionAdlQuantileEstimationResponseInnerAdlQuantile and assigns it to the AdlQuantile field.
func (o *CmPositionAdlQuantileEstimationResponseInner) SetAdlQuantile(v CmPositionAdlQuantileEstimationResponseInnerAdlQuantile) {
	o.AdlQuantile = &v
}

func (o CmPositionAdlQuantileEstimationResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CmPositionAdlQuantileEstimationResponseInner) ToMap() (map[string]interface{}, error) {
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

func (o *CmPositionAdlQuantileEstimationResponseInner) UnmarshalJSON(data []byte) (err error) {
	varCmPositionAdlQuantileEstimationResponseInner := _CmPositionAdlQuantileEstimationResponseInner{}

	err = json.Unmarshal(data, &varCmPositionAdlQuantileEstimationResponseInner)

	if err != nil {
		return err
	}

	*o = CmPositionAdlQuantileEstimationResponseInner(varCmPositionAdlQuantileEstimationResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "adlQuantile")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCmPositionAdlQuantileEstimationResponseInner struct {
	value *CmPositionAdlQuantileEstimationResponseInner
	isSet bool
}

func (v NullableCmPositionAdlQuantileEstimationResponseInner) Get() *CmPositionAdlQuantileEstimationResponseInner {
	return v.value
}

func (v *NullableCmPositionAdlQuantileEstimationResponseInner) Set(val *CmPositionAdlQuantileEstimationResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCmPositionAdlQuantileEstimationResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCmPositionAdlQuantileEstimationResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCmPositionAdlQuantileEstimationResponseInner(val *CmPositionAdlQuantileEstimationResponseInner) *NullableCmPositionAdlQuantileEstimationResponseInner {
	return &NullableCmPositionAdlQuantileEstimationResponseInner{value: val, isSet: true}
}

func (v NullableCmPositionAdlQuantileEstimationResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCmPositionAdlQuantileEstimationResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
