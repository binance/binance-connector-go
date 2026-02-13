/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ExchangeInformationResponseOptionAssetsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ExchangeInformationResponseOptionAssetsInner{}

// ExchangeInformationResponseOptionAssetsInner struct for ExchangeInformationResponseOptionAssetsInner
type ExchangeInformationResponseOptionAssetsInner struct {
	Name                 *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExchangeInformationResponseOptionAssetsInner ExchangeInformationResponseOptionAssetsInner

// NewExchangeInformationResponseOptionAssetsInner instantiates a new ExchangeInformationResponseOptionAssetsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeInformationResponseOptionAssetsInner() *ExchangeInformationResponseOptionAssetsInner {
	this := ExchangeInformationResponseOptionAssetsInner{}
	return &this
}

// NewExchangeInformationResponseOptionAssetsInnerWithDefaults instantiates a new ExchangeInformationResponseOptionAssetsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeInformationResponseOptionAssetsInnerWithDefaults() *ExchangeInformationResponseOptionAssetsInner {
	this := ExchangeInformationResponseOptionAssetsInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExchangeInformationResponseOptionAssetsInner) GetName() string {
	if o == nil || common.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseOptionAssetsInner) GetNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExchangeInformationResponseOptionAssetsInner) HasName() bool {
	if o != nil && !common.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExchangeInformationResponseOptionAssetsInner) SetName(v string) {
	o.Name = &v
}

func (o ExchangeInformationResponseOptionAssetsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeInformationResponseOptionAssetsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExchangeInformationResponseOptionAssetsInner) UnmarshalJSON(data []byte) (err error) {
	varExchangeInformationResponseOptionAssetsInner := _ExchangeInformationResponseOptionAssetsInner{}

	err = json.Unmarshal(data, &varExchangeInformationResponseOptionAssetsInner)

	if err != nil {
		return err
	}

	*o = ExchangeInformationResponseOptionAssetsInner(varExchangeInformationResponseOptionAssetsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExchangeInformationResponseOptionAssetsInner struct {
	value *ExchangeInformationResponseOptionAssetsInner
	isSet bool
}

func (v NullableExchangeInformationResponseOptionAssetsInner) Get() *ExchangeInformationResponseOptionAssetsInner {
	return v.value
}

func (v *NullableExchangeInformationResponseOptionAssetsInner) Set(val *ExchangeInformationResponseOptionAssetsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeInformationResponseOptionAssetsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeInformationResponseOptionAssetsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeInformationResponseOptionAssetsInner(val *ExchangeInformationResponseOptionAssetsInner) *NullableExchangeInformationResponseOptionAssetsInner {
	return &NullableExchangeInformationResponseOptionAssetsInner{value: val, isSet: true}
}

func (v NullableExchangeInformationResponseOptionAssetsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeInformationResponseOptionAssetsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
