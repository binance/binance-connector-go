/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate{}

// GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate struct for GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate
type GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate struct {
	Var05BTC             *float32 `json:"0-5BTC,omitempty"`
	Var510BTC            *float32 `json:"5-10BTC,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate

// NewGetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate instantiates a new GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate() *GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate {
	this := GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate{}
	return &this
}

// NewGetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRateWithDefaults instantiates a new GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRateWithDefaults() *GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate {
	this := GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate{}
	return &this
}

// GetVar05BTC returns the Var05BTC field value if set, zero value otherwise.
func (o *GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate) GetVar05BTC() float32 {
	if o == nil || common.IsNil(o.Var05BTC) {
		var ret float32
		return ret
	}
	return *o.Var05BTC
}

// GetVar05BTCOk returns a tuple with the Var05BTC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate) GetVar05BTCOk() (*float32, bool) {
	if o == nil || common.IsNil(o.Var05BTC) {
		return nil, false
	}
	return o.Var05BTC, true
}

// HasVar05BTC returns a boolean if a field has been set.
func (o *GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate) HasVar05BTC() bool {
	if o != nil && !common.IsNil(o.Var05BTC) {
		return true
	}

	return false
}

// SetVar05BTC gets a reference to the given float32 and assigns it to the Var05BTC field.
func (o *GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate) SetVar05BTC(v float32) {
	o.Var05BTC = &v
}

// GetVar510BTC returns the Var510BTC field value if set, zero value otherwise.
func (o *GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate) GetVar510BTC() float32 {
	if o == nil || common.IsNil(o.Var510BTC) {
		var ret float32
		return ret
	}
	return *o.Var510BTC
}

// GetVar510BTCOk returns a tuple with the Var510BTC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate) GetVar510BTCOk() (*float32, bool) {
	if o == nil || common.IsNil(o.Var510BTC) {
		return nil, false
	}
	return o.Var510BTC, true
}

// HasVar510BTC returns a boolean if a field has been set.
func (o *GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate) HasVar510BTC() bool {
	if o != nil && !common.IsNil(o.Var510BTC) {
		return true
	}

	return false
}

// SetVar510BTC gets a reference to the given float32 and assigns it to the Var510BTC field.
func (o *GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate) SetVar510BTC(v float32) {
	o.Var510BTC = &v
}

func (o GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Var05BTC) {
		toSerialize["0-5BTC"] = o.Var05BTC
	}
	if !common.IsNil(o.Var510BTC) {
		toSerialize["5-10BTC"] = o.Var510BTC
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate) UnmarshalJSON(data []byte) (err error) {
	varGetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate := _GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate{}

	err = json.Unmarshal(data, &varGetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate)

	if err != nil {
		return err
	}

	*o = GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate(varGetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "0-5BTC")
		delete(additionalProperties, "5-10BTC")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate struct {
	value *GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate
	isSet bool
}

func (v NullableGetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate) Get() *GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate {
	return v.value
}

func (v *NullableGetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate) Set(val *GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate(val *GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate) *NullableGetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate {
	return &NullableGetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate{value: val, isSet: true}
}

func (v NullableGetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
