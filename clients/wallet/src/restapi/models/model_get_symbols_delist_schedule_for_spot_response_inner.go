/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetSymbolsDelistScheduleForSpotResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSymbolsDelistScheduleForSpotResponseInner{}

// GetSymbolsDelistScheduleForSpotResponseInner struct for GetSymbolsDelistScheduleForSpotResponseInner
type GetSymbolsDelistScheduleForSpotResponseInner struct {
	DelistTime           *int64   `json:"delistTime,omitempty"`
	Symbols              []string `json:"symbols,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetSymbolsDelistScheduleForSpotResponseInner GetSymbolsDelistScheduleForSpotResponseInner

// NewGetSymbolsDelistScheduleForSpotResponseInner instantiates a new GetSymbolsDelistScheduleForSpotResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSymbolsDelistScheduleForSpotResponseInner() *GetSymbolsDelistScheduleForSpotResponseInner {
	this := GetSymbolsDelistScheduleForSpotResponseInner{}
	return &this
}

// NewGetSymbolsDelistScheduleForSpotResponseInnerWithDefaults instantiates a new GetSymbolsDelistScheduleForSpotResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSymbolsDelistScheduleForSpotResponseInnerWithDefaults() *GetSymbolsDelistScheduleForSpotResponseInner {
	this := GetSymbolsDelistScheduleForSpotResponseInner{}
	return &this
}

// GetDelistTime returns the DelistTime field value if set, zero value otherwise.
func (o *GetSymbolsDelistScheduleForSpotResponseInner) GetDelistTime() int64 {
	if o == nil || common.IsNil(o.DelistTime) {
		var ret int64
		return ret
	}
	return *o.DelistTime
}

// GetDelistTimeOk returns a tuple with the DelistTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSymbolsDelistScheduleForSpotResponseInner) GetDelistTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.DelistTime) {
		return nil, false
	}
	return o.DelistTime, true
}

// HasDelistTime returns a boolean if a field has been set.
func (o *GetSymbolsDelistScheduleForSpotResponseInner) HasDelistTime() bool {
	if o != nil && !common.IsNil(o.DelistTime) {
		return true
	}

	return false
}

// SetDelistTime gets a reference to the given int64 and assigns it to the DelistTime field.
func (o *GetSymbolsDelistScheduleForSpotResponseInner) SetDelistTime(v int64) {
	o.DelistTime = &v
}

// GetSymbols returns the Symbols field value if set, zero value otherwise.
func (o *GetSymbolsDelistScheduleForSpotResponseInner) GetSymbols() []string {
	if o == nil || common.IsNil(o.Symbols) {
		var ret []string
		return ret
	}
	return o.Symbols
}

// GetSymbolsOk returns a tuple with the Symbols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSymbolsDelistScheduleForSpotResponseInner) GetSymbolsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Symbols) {
		return nil, false
	}
	return o.Symbols, true
}

// HasSymbols returns a boolean if a field has been set.
func (o *GetSymbolsDelistScheduleForSpotResponseInner) HasSymbols() bool {
	if o != nil && !common.IsNil(o.Symbols) {
		return true
	}

	return false
}

// SetSymbols gets a reference to the given []string and assigns it to the Symbols field.
func (o *GetSymbolsDelistScheduleForSpotResponseInner) SetSymbols(v []string) {
	o.Symbols = v
}

func (o GetSymbolsDelistScheduleForSpotResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSymbolsDelistScheduleForSpotResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.DelistTime) {
		toSerialize["delistTime"] = o.DelistTime
	}
	if !common.IsNil(o.Symbols) {
		toSerialize["symbols"] = o.Symbols
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetSymbolsDelistScheduleForSpotResponseInner) UnmarshalJSON(data []byte) (err error) {
	varGetSymbolsDelistScheduleForSpotResponseInner := _GetSymbolsDelistScheduleForSpotResponseInner{}

	err = json.Unmarshal(data, &varGetSymbolsDelistScheduleForSpotResponseInner)

	if err != nil {
		return err
	}

	*o = GetSymbolsDelistScheduleForSpotResponseInner(varGetSymbolsDelistScheduleForSpotResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "delistTime")
		delete(additionalProperties, "symbols")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSymbolsDelistScheduleForSpotResponseInner struct {
	value *GetSymbolsDelistScheduleForSpotResponseInner
	isSet bool
}

func (v NullableGetSymbolsDelistScheduleForSpotResponseInner) Get() *GetSymbolsDelistScheduleForSpotResponseInner {
	return v.value
}

func (v *NullableGetSymbolsDelistScheduleForSpotResponseInner) Set(val *GetSymbolsDelistScheduleForSpotResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSymbolsDelistScheduleForSpotResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSymbolsDelistScheduleForSpotResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSymbolsDelistScheduleForSpotResponseInner(val *GetSymbolsDelistScheduleForSpotResponseInner) *NullableGetSymbolsDelistScheduleForSpotResponseInner {
	return &NullableGetSymbolsDelistScheduleForSpotResponseInner{value: val, isSet: true}
}

func (v NullableGetSymbolsDelistScheduleForSpotResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSymbolsDelistScheduleForSpotResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
