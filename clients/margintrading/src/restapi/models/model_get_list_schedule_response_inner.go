/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetListScheduleResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetListScheduleResponseInner{}

// GetListScheduleResponseInner struct for GetListScheduleResponseInner
type GetListScheduleResponseInner struct {
	ListTime              *int64   `json:"listTime,omitempty"`
	CrossMarginAssets     []string `json:"crossMarginAssets,omitempty"`
	IsolatedMarginSymbols []string `json:"isolatedMarginSymbols,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _GetListScheduleResponseInner GetListScheduleResponseInner

// NewGetListScheduleResponseInner instantiates a new GetListScheduleResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetListScheduleResponseInner() *GetListScheduleResponseInner {
	this := GetListScheduleResponseInner{}
	return &this
}

// NewGetListScheduleResponseInnerWithDefaults instantiates a new GetListScheduleResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetListScheduleResponseInnerWithDefaults() *GetListScheduleResponseInner {
	this := GetListScheduleResponseInner{}
	return &this
}

// GetListTime returns the ListTime field value if set, zero value otherwise.
func (o *GetListScheduleResponseInner) GetListTime() int64 {
	if o == nil || common.IsNil(o.ListTime) {
		var ret int64
		return ret
	}
	return *o.ListTime
}

// GetListTimeOk returns a tuple with the ListTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetListScheduleResponseInner) GetListTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ListTime) {
		return nil, false
	}
	return o.ListTime, true
}

// HasListTime returns a boolean if a field has been set.
func (o *GetListScheduleResponseInner) HasListTime() bool {
	if o != nil && !common.IsNil(o.ListTime) {
		return true
	}

	return false
}

// SetListTime gets a reference to the given int64 and assigns it to the ListTime field.
func (o *GetListScheduleResponseInner) SetListTime(v int64) {
	o.ListTime = &v
}

// GetCrossMarginAssets returns the CrossMarginAssets field value if set, zero value otherwise.
func (o *GetListScheduleResponseInner) GetCrossMarginAssets() []string {
	if o == nil || common.IsNil(o.CrossMarginAssets) {
		var ret []string
		return ret
	}
	return o.CrossMarginAssets
}

// GetCrossMarginAssetsOk returns a tuple with the CrossMarginAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetListScheduleResponseInner) GetCrossMarginAssetsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.CrossMarginAssets) {
		return nil, false
	}
	return o.CrossMarginAssets, true
}

// HasCrossMarginAssets returns a boolean if a field has been set.
func (o *GetListScheduleResponseInner) HasCrossMarginAssets() bool {
	if o != nil && !common.IsNil(o.CrossMarginAssets) {
		return true
	}

	return false
}

// SetCrossMarginAssets gets a reference to the given []string and assigns it to the CrossMarginAssets field.
func (o *GetListScheduleResponseInner) SetCrossMarginAssets(v []string) {
	o.CrossMarginAssets = v
}

// GetIsolatedMarginSymbols returns the IsolatedMarginSymbols field value if set, zero value otherwise.
func (o *GetListScheduleResponseInner) GetIsolatedMarginSymbols() []string {
	if o == nil || common.IsNil(o.IsolatedMarginSymbols) {
		var ret []string
		return ret
	}
	return o.IsolatedMarginSymbols
}

// GetIsolatedMarginSymbolsOk returns a tuple with the IsolatedMarginSymbols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetListScheduleResponseInner) GetIsolatedMarginSymbolsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.IsolatedMarginSymbols) {
		return nil, false
	}
	return o.IsolatedMarginSymbols, true
}

// HasIsolatedMarginSymbols returns a boolean if a field has been set.
func (o *GetListScheduleResponseInner) HasIsolatedMarginSymbols() bool {
	if o != nil && !common.IsNil(o.IsolatedMarginSymbols) {
		return true
	}

	return false
}

// SetIsolatedMarginSymbols gets a reference to the given []string and assigns it to the IsolatedMarginSymbols field.
func (o *GetListScheduleResponseInner) SetIsolatedMarginSymbols(v []string) {
	o.IsolatedMarginSymbols = v
}

func (o GetListScheduleResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetListScheduleResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ListTime) {
		toSerialize["listTime"] = o.ListTime
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

func (o *GetListScheduleResponseInner) UnmarshalJSON(data []byte) (err error) {
	varGetListScheduleResponseInner := _GetListScheduleResponseInner{}

	err = json.Unmarshal(data, &varGetListScheduleResponseInner)

	if err != nil {
		return err
	}

	*o = GetListScheduleResponseInner(varGetListScheduleResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "listTime")
		delete(additionalProperties, "crossMarginAssets")
		delete(additionalProperties, "isolatedMarginSymbols")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetListScheduleResponseInner struct {
	value *GetListScheduleResponseInner
	isSet bool
}

func (v NullableGetListScheduleResponseInner) Get() *GetListScheduleResponseInner {
	return v.value
}

func (v *NullableGetListScheduleResponseInner) Set(val *GetListScheduleResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetListScheduleResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetListScheduleResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetListScheduleResponseInner(val *GetListScheduleResponseInner) *NullableGetListScheduleResponseInner {
	return &NullableGetListScheduleResponseInner{value: val, isSet: true}
}

func (v NullableGetListScheduleResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetListScheduleResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
