/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetDelistScheduleResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetDelistScheduleResponseInner{}

// GetDelistScheduleResponseInner struct for GetDelistScheduleResponseInner
type GetDelistScheduleResponseInner struct {
	DelistTime            *int64   `json:"delistTime,omitempty"`
	CrossMarginAssets     []string `json:"crossMarginAssets,omitempty"`
	IsolatedMarginSymbols []string `json:"isolatedMarginSymbols,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _GetDelistScheduleResponseInner GetDelistScheduleResponseInner

// NewGetDelistScheduleResponseInner instantiates a new GetDelistScheduleResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDelistScheduleResponseInner() *GetDelistScheduleResponseInner {
	this := GetDelistScheduleResponseInner{}
	return &this
}

// NewGetDelistScheduleResponseInnerWithDefaults instantiates a new GetDelistScheduleResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDelistScheduleResponseInnerWithDefaults() *GetDelistScheduleResponseInner {
	this := GetDelistScheduleResponseInner{}
	return &this
}

// GetDelistTime returns the DelistTime field value if set, zero value otherwise.
func (o *GetDelistScheduleResponseInner) GetDelistTime() int64 {
	if o == nil || common.IsNil(o.DelistTime) {
		var ret int64
		return ret
	}
	return *o.DelistTime
}

// GetDelistTimeOk returns a tuple with the DelistTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDelistScheduleResponseInner) GetDelistTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.DelistTime) {
		return nil, false
	}
	return o.DelistTime, true
}

// HasDelistTime returns a boolean if a field has been set.
func (o *GetDelistScheduleResponseInner) HasDelistTime() bool {
	if o != nil && !common.IsNil(o.DelistTime) {
		return true
	}

	return false
}

// SetDelistTime gets a reference to the given int64 and assigns it to the DelistTime field.
func (o *GetDelistScheduleResponseInner) SetDelistTime(v int64) {
	o.DelistTime = &v
}

// GetCrossMarginAssets returns the CrossMarginAssets field value if set, zero value otherwise.
func (o *GetDelistScheduleResponseInner) GetCrossMarginAssets() []string {
	if o == nil || common.IsNil(o.CrossMarginAssets) {
		var ret []string
		return ret
	}
	return o.CrossMarginAssets
}

// GetCrossMarginAssetsOk returns a tuple with the CrossMarginAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDelistScheduleResponseInner) GetCrossMarginAssetsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.CrossMarginAssets) {
		return nil, false
	}
	return o.CrossMarginAssets, true
}

// HasCrossMarginAssets returns a boolean if a field has been set.
func (o *GetDelistScheduleResponseInner) HasCrossMarginAssets() bool {
	if o != nil && !common.IsNil(o.CrossMarginAssets) {
		return true
	}

	return false
}

// SetCrossMarginAssets gets a reference to the given []string and assigns it to the CrossMarginAssets field.
func (o *GetDelistScheduleResponseInner) SetCrossMarginAssets(v []string) {
	o.CrossMarginAssets = v
}

// GetIsolatedMarginSymbols returns the IsolatedMarginSymbols field value if set, zero value otherwise.
func (o *GetDelistScheduleResponseInner) GetIsolatedMarginSymbols() []string {
	if o == nil || common.IsNil(o.IsolatedMarginSymbols) {
		var ret []string
		return ret
	}
	return o.IsolatedMarginSymbols
}

// GetIsolatedMarginSymbolsOk returns a tuple with the IsolatedMarginSymbols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDelistScheduleResponseInner) GetIsolatedMarginSymbolsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.IsolatedMarginSymbols) {
		return nil, false
	}
	return o.IsolatedMarginSymbols, true
}

// HasIsolatedMarginSymbols returns a boolean if a field has been set.
func (o *GetDelistScheduleResponseInner) HasIsolatedMarginSymbols() bool {
	if o != nil && !common.IsNil(o.IsolatedMarginSymbols) {
		return true
	}

	return false
}

// SetIsolatedMarginSymbols gets a reference to the given []string and assigns it to the IsolatedMarginSymbols field.
func (o *GetDelistScheduleResponseInner) SetIsolatedMarginSymbols(v []string) {
	o.IsolatedMarginSymbols = v
}

func (o GetDelistScheduleResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDelistScheduleResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.DelistTime) {
		toSerialize["delistTime"] = o.DelistTime
	}
	if !common.IsNil(o.CrossMarginAssets) {
		toSerialize["crossMarginAssets"] = o.CrossMarginAssets
	}
	if !common.IsNil(o.IsolatedMarginSymbols) {
		toSerialize["isolatedMarginSymbols"] = o.IsolatedMarginSymbols
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetDelistScheduleResponseInner) UnmarshalJSON(data []byte) (err error) {
	varGetDelistScheduleResponseInner := _GetDelistScheduleResponseInner{}

	err = json.Unmarshal(data, &varGetDelistScheduleResponseInner)

	if err != nil {
		return err
	}

	*o = GetDelistScheduleResponseInner(varGetDelistScheduleResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "delistTime")
		delete(additionalProperties, "crossMarginAssets")
		delete(additionalProperties, "isolatedMarginSymbols")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetDelistScheduleResponseInner struct {
	value *GetDelistScheduleResponseInner
	isSet bool
}

func (v NullableGetDelistScheduleResponseInner) Get() *GetDelistScheduleResponseInner {
	return v.value
}

func (v *NullableGetDelistScheduleResponseInner) Set(val *GetDelistScheduleResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDelistScheduleResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDelistScheduleResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDelistScheduleResponseInner(val *GetDelistScheduleResponseInner) *NullableGetDelistScheduleResponseInner {
	return &NullableGetDelistScheduleResponseInner{value: val, isSet: true}
}

func (v NullableGetDelistScheduleResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDelistScheduleResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
