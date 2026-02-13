/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetOpenSymbolListResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetOpenSymbolListResponseInner{}

// GetOpenSymbolListResponseInner struct for GetOpenSymbolListResponseInner
type GetOpenSymbolListResponseInner struct {
	OpenTime             *int64   `json:"openTime,omitempty"`
	Symbols              []string `json:"symbols,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetOpenSymbolListResponseInner GetOpenSymbolListResponseInner

// NewGetOpenSymbolListResponseInner instantiates a new GetOpenSymbolListResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOpenSymbolListResponseInner() *GetOpenSymbolListResponseInner {
	this := GetOpenSymbolListResponseInner{}
	return &this
}

// NewGetOpenSymbolListResponseInnerWithDefaults instantiates a new GetOpenSymbolListResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOpenSymbolListResponseInnerWithDefaults() *GetOpenSymbolListResponseInner {
	this := GetOpenSymbolListResponseInner{}
	return &this
}

// GetOpenTime returns the OpenTime field value if set, zero value otherwise.
func (o *GetOpenSymbolListResponseInner) GetOpenTime() int64 {
	if o == nil || common.IsNil(o.OpenTime) {
		var ret int64
		return ret
	}
	return *o.OpenTime
}

// GetOpenTimeOk returns a tuple with the OpenTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOpenSymbolListResponseInner) GetOpenTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OpenTime) {
		return nil, false
	}
	return o.OpenTime, true
}

// HasOpenTime returns a boolean if a field has been set.
func (o *GetOpenSymbolListResponseInner) HasOpenTime() bool {
	if o != nil && !common.IsNil(o.OpenTime) {
		return true
	}

	return false
}

// SetOpenTime gets a reference to the given int64 and assigns it to the OpenTime field.
func (o *GetOpenSymbolListResponseInner) SetOpenTime(v int64) {
	o.OpenTime = &v
}

// GetSymbols returns the Symbols field value if set, zero value otherwise.
func (o *GetOpenSymbolListResponseInner) GetSymbols() []string {
	if o == nil || common.IsNil(o.Symbols) {
		var ret []string
		return ret
	}
	return o.Symbols
}

// GetSymbolsOk returns a tuple with the Symbols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOpenSymbolListResponseInner) GetSymbolsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Symbols) {
		return nil, false
	}
	return o.Symbols, true
}

// HasSymbols returns a boolean if a field has been set.
func (o *GetOpenSymbolListResponseInner) HasSymbols() bool {
	if o != nil && !common.IsNil(o.Symbols) {
		return true
	}

	return false
}

// SetSymbols gets a reference to the given []string and assigns it to the Symbols field.
func (o *GetOpenSymbolListResponseInner) SetSymbols(v []string) {
	o.Symbols = v
}

func (o GetOpenSymbolListResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOpenSymbolListResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.OpenTime) {
		toSerialize["openTime"] = o.OpenTime
	}
	if !common.IsNil(o.Symbols) {
		toSerialize["symbols"] = o.Symbols
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetOpenSymbolListResponseInner) UnmarshalJSON(data []byte) (err error) {
	varGetOpenSymbolListResponseInner := _GetOpenSymbolListResponseInner{}

	err = json.Unmarshal(data, &varGetOpenSymbolListResponseInner)

	if err != nil {
		return err
	}

	*o = GetOpenSymbolListResponseInner(varGetOpenSymbolListResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "openTime")
		delete(additionalProperties, "symbols")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetOpenSymbolListResponseInner struct {
	value *GetOpenSymbolListResponseInner
	isSet bool
}

func (v NullableGetOpenSymbolListResponseInner) Get() *GetOpenSymbolListResponseInner {
	return v.value
}

func (v *NullableGetOpenSymbolListResponseInner) Set(val *GetOpenSymbolListResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOpenSymbolListResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOpenSymbolListResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOpenSymbolListResponseInner(val *GetOpenSymbolListResponseInner) *NullableGetOpenSymbolListResponseInner {
	return &NullableGetOpenSymbolListResponseInner{value: val, isSet: true}
}

func (v NullableGetOpenSymbolListResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOpenSymbolListResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
