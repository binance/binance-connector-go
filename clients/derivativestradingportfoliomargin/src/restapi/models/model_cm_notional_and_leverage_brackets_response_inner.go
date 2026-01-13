/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the CmNotionalAndLeverageBracketsResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CmNotionalAndLeverageBracketsResponseInner{}

// CmNotionalAndLeverageBracketsResponseInner struct for CmNotionalAndLeverageBracketsResponseInner
type CmNotionalAndLeverageBracketsResponseInner struct {
	Symbol               *string                                                   `json:"symbol,omitempty"`
	Brackets             []CmNotionalAndLeverageBracketsResponseInnerBracketsInner `json:"brackets,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CmNotionalAndLeverageBracketsResponseInner CmNotionalAndLeverageBracketsResponseInner

// NewCmNotionalAndLeverageBracketsResponseInner instantiates a new CmNotionalAndLeverageBracketsResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCmNotionalAndLeverageBracketsResponseInner() *CmNotionalAndLeverageBracketsResponseInner {
	this := CmNotionalAndLeverageBracketsResponseInner{}
	return &this
}

// NewCmNotionalAndLeverageBracketsResponseInnerWithDefaults instantiates a new CmNotionalAndLeverageBracketsResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCmNotionalAndLeverageBracketsResponseInnerWithDefaults() *CmNotionalAndLeverageBracketsResponseInner {
	this := CmNotionalAndLeverageBracketsResponseInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *CmNotionalAndLeverageBracketsResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmNotionalAndLeverageBracketsResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *CmNotionalAndLeverageBracketsResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *CmNotionalAndLeverageBracketsResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetBrackets returns the Brackets field value if set, zero value otherwise.
func (o *CmNotionalAndLeverageBracketsResponseInner) GetBrackets() []CmNotionalAndLeverageBracketsResponseInnerBracketsInner {
	if o == nil || common.IsNil(o.Brackets) {
		var ret []CmNotionalAndLeverageBracketsResponseInnerBracketsInner
		return ret
	}
	return o.Brackets
}

// GetBracketsOk returns a tuple with the Brackets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmNotionalAndLeverageBracketsResponseInner) GetBracketsOk() ([]CmNotionalAndLeverageBracketsResponseInnerBracketsInner, bool) {
	if o == nil || common.IsNil(o.Brackets) {
		return nil, false
	}
	return o.Brackets, true
}

// HasBrackets returns a boolean if a field has been set.
func (o *CmNotionalAndLeverageBracketsResponseInner) HasBrackets() bool {
	if o != nil && !common.IsNil(o.Brackets) {
		return true
	}

	return false
}

// SetBrackets gets a reference to the given []CmNotionalAndLeverageBracketsResponseInnerBracketsInner and assigns it to the Brackets field.
func (o *CmNotionalAndLeverageBracketsResponseInner) SetBrackets(v []CmNotionalAndLeverageBracketsResponseInnerBracketsInner) {
	o.Brackets = v
}

func (o CmNotionalAndLeverageBracketsResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CmNotionalAndLeverageBracketsResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Brackets) {
		toSerialize["brackets"] = o.Brackets
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CmNotionalAndLeverageBracketsResponseInner) UnmarshalJSON(data []byte) (err error) {
	varCmNotionalAndLeverageBracketsResponseInner := _CmNotionalAndLeverageBracketsResponseInner{}

	err = json.Unmarshal(data, &varCmNotionalAndLeverageBracketsResponseInner)

	if err != nil {
		return err
	}

	*o = CmNotionalAndLeverageBracketsResponseInner(varCmNotionalAndLeverageBracketsResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "brackets")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCmNotionalAndLeverageBracketsResponseInner struct {
	value *CmNotionalAndLeverageBracketsResponseInner
	isSet bool
}

func (v NullableCmNotionalAndLeverageBracketsResponseInner) Get() *CmNotionalAndLeverageBracketsResponseInner {
	return v.value
}

func (v *NullableCmNotionalAndLeverageBracketsResponseInner) Set(val *CmNotionalAndLeverageBracketsResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCmNotionalAndLeverageBracketsResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCmNotionalAndLeverageBracketsResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCmNotionalAndLeverageBracketsResponseInner(val *CmNotionalAndLeverageBracketsResponseInner) *NullableCmNotionalAndLeverageBracketsResponseInner {
	return &NullableCmNotionalAndLeverageBracketsResponseInner{value: val, isSet: true}
}

func (v NullableCmNotionalAndLeverageBracketsResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCmNotionalAndLeverageBracketsResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
