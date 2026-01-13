/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetSpotDelistScheduleResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSpotDelistScheduleResponseInner{}

// GetSpotDelistScheduleResponseInner struct for GetSpotDelistScheduleResponseInner
type GetSpotDelistScheduleResponseInner struct {
	DelistTime           *int64   `json:"delistTime,omitempty"`
	Symbols              []string `json:"symbols,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetSpotDelistScheduleResponseInner GetSpotDelistScheduleResponseInner

// NewGetSpotDelistScheduleResponseInner instantiates a new GetSpotDelistScheduleResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSpotDelistScheduleResponseInner() *GetSpotDelistScheduleResponseInner {
	this := GetSpotDelistScheduleResponseInner{}
	return &this
}

// NewGetSpotDelistScheduleResponseInnerWithDefaults instantiates a new GetSpotDelistScheduleResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSpotDelistScheduleResponseInnerWithDefaults() *GetSpotDelistScheduleResponseInner {
	this := GetSpotDelistScheduleResponseInner{}
	return &this
}

// GetDelistTime returns the DelistTime field value if set, zero value otherwise.
func (o *GetSpotDelistScheduleResponseInner) GetDelistTime() int64 {
	if o == nil || common.IsNil(o.DelistTime) {
		var ret int64
		return ret
	}
	return *o.DelistTime
}

// GetDelistTimeOk returns a tuple with the DelistTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSpotDelistScheduleResponseInner) GetDelistTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.DelistTime) {
		return nil, false
	}
	return o.DelistTime, true
}

// HasDelistTime returns a boolean if a field has been set.
func (o *GetSpotDelistScheduleResponseInner) HasDelistTime() bool {
	if o != nil && !common.IsNil(o.DelistTime) {
		return true
	}

	return false
}

// SetDelistTime gets a reference to the given int64 and assigns it to the DelistTime field.
func (o *GetSpotDelistScheduleResponseInner) SetDelistTime(v int64) {
	o.DelistTime = &v
}

// GetSymbols returns the Symbols field value if set, zero value otherwise.
func (o *GetSpotDelistScheduleResponseInner) GetSymbols() []string {
	if o == nil || common.IsNil(o.Symbols) {
		var ret []string
		return ret
	}
	return o.Symbols
}

// GetSymbolsOk returns a tuple with the Symbols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSpotDelistScheduleResponseInner) GetSymbolsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Symbols) {
		return nil, false
	}
	return o.Symbols, true
}

// HasSymbols returns a boolean if a field has been set.
func (o *GetSpotDelistScheduleResponseInner) HasSymbols() bool {
	if o != nil && !common.IsNil(o.Symbols) {
		return true
	}

	return false
}

// SetSymbols gets a reference to the given []string and assigns it to the Symbols field.
func (o *GetSpotDelistScheduleResponseInner) SetSymbols(v []string) {
	o.Symbols = v
}

func (o GetSpotDelistScheduleResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSpotDelistScheduleResponseInner) ToMap() (map[string]interface{}, error) {
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

func (o *GetSpotDelistScheduleResponseInner) UnmarshalJSON(data []byte) (err error) {
	varGetSpotDelistScheduleResponseInner := _GetSpotDelistScheduleResponseInner{}

	err = json.Unmarshal(data, &varGetSpotDelistScheduleResponseInner)

	if err != nil {
		return err
	}

	*o = GetSpotDelistScheduleResponseInner(varGetSpotDelistScheduleResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "delistTime")
		delete(additionalProperties, "symbols")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSpotDelistScheduleResponseInner struct {
	value *GetSpotDelistScheduleResponseInner
	isSet bool
}

func (v NullableGetSpotDelistScheduleResponseInner) Get() *GetSpotDelistScheduleResponseInner {
	return v.value
}

func (v *NullableGetSpotDelistScheduleResponseInner) Set(val *GetSpotDelistScheduleResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSpotDelistScheduleResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSpotDelistScheduleResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSpotDelistScheduleResponseInner(val *GetSpotDelistScheduleResponseInner) *NullableGetSpotDelistScheduleResponseInner {
	return &NullableGetSpotDelistScheduleResponseInner{value: val, isSet: true}
}

func (v NullableGetSpotDelistScheduleResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSpotDelistScheduleResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
