/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the CompositeIndexSymbolInformationResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CompositeIndexSymbolInformationResponseInner{}

// CompositeIndexSymbolInformationResponseInner struct for CompositeIndexSymbolInformationResponseInner
type CompositeIndexSymbolInformationResponseInner struct {
	Symbol               *string                                                          `json:"symbol,omitempty"`
	Time                 *int64                                                           `json:"time,omitempty"`
	Component            *string                                                          `json:"component,omitempty"`
	BaseAssetList        []CompositeIndexSymbolInformationResponseInnerBaseAssetListInner `json:"baseAssetList,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CompositeIndexSymbolInformationResponseInner CompositeIndexSymbolInformationResponseInner

// NewCompositeIndexSymbolInformationResponseInner instantiates a new CompositeIndexSymbolInformationResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompositeIndexSymbolInformationResponseInner() *CompositeIndexSymbolInformationResponseInner {
	this := CompositeIndexSymbolInformationResponseInner{}
	return &this
}

// NewCompositeIndexSymbolInformationResponseInnerWithDefaults instantiates a new CompositeIndexSymbolInformationResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompositeIndexSymbolInformationResponseInnerWithDefaults() *CompositeIndexSymbolInformationResponseInner {
	this := CompositeIndexSymbolInformationResponseInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *CompositeIndexSymbolInformationResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeIndexSymbolInformationResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *CompositeIndexSymbolInformationResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *CompositeIndexSymbolInformationResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *CompositeIndexSymbolInformationResponseInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeIndexSymbolInformationResponseInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *CompositeIndexSymbolInformationResponseInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *CompositeIndexSymbolInformationResponseInner) SetTime(v int64) {
	o.Time = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *CompositeIndexSymbolInformationResponseInner) GetComponent() string {
	if o == nil || common.IsNil(o.Component) {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeIndexSymbolInformationResponseInner) GetComponentOk() (*string, bool) {
	if o == nil || common.IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *CompositeIndexSymbolInformationResponseInner) HasComponent() bool {
	if o != nil && !common.IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *CompositeIndexSymbolInformationResponseInner) SetComponent(v string) {
	o.Component = &v
}

// GetBaseAssetList returns the BaseAssetList field value if set, zero value otherwise.
func (o *CompositeIndexSymbolInformationResponseInner) GetBaseAssetList() []CompositeIndexSymbolInformationResponseInnerBaseAssetListInner {
	if o == nil || common.IsNil(o.BaseAssetList) {
		var ret []CompositeIndexSymbolInformationResponseInnerBaseAssetListInner
		return ret
	}
	return o.BaseAssetList
}

// GetBaseAssetListOk returns a tuple with the BaseAssetList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeIndexSymbolInformationResponseInner) GetBaseAssetListOk() ([]CompositeIndexSymbolInformationResponseInnerBaseAssetListInner, bool) {
	if o == nil || common.IsNil(o.BaseAssetList) {
		return nil, false
	}
	return o.BaseAssetList, true
}

// HasBaseAssetList returns a boolean if a field has been set.
func (o *CompositeIndexSymbolInformationResponseInner) HasBaseAssetList() bool {
	if o != nil && !common.IsNil(o.BaseAssetList) {
		return true
	}

	return false
}

// SetBaseAssetList gets a reference to the given []CompositeIndexSymbolInformationResponseInnerBaseAssetListInner and assigns it to the BaseAssetList field.
func (o *CompositeIndexSymbolInformationResponseInner) SetBaseAssetList(v []CompositeIndexSymbolInformationResponseInnerBaseAssetListInner) {
	o.BaseAssetList = v
}

func (o CompositeIndexSymbolInformationResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompositeIndexSymbolInformationResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !common.IsNil(o.Component) {
		toSerialize["component"] = o.Component
	}
	if !common.IsNil(o.BaseAssetList) {
		toSerialize["baseAssetList"] = o.BaseAssetList
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CompositeIndexSymbolInformationResponseInner) UnmarshalJSON(data []byte) (err error) {
	varCompositeIndexSymbolInformationResponseInner := _CompositeIndexSymbolInformationResponseInner{}

	err = json.Unmarshal(data, &varCompositeIndexSymbolInformationResponseInner)

	if err != nil {
		return err
	}

	*o = CompositeIndexSymbolInformationResponseInner(varCompositeIndexSymbolInformationResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "time")
		delete(additionalProperties, "component")
		delete(additionalProperties, "baseAssetList")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCompositeIndexSymbolInformationResponseInner struct {
	value *CompositeIndexSymbolInformationResponseInner
	isSet bool
}

func (v NullableCompositeIndexSymbolInformationResponseInner) Get() *CompositeIndexSymbolInformationResponseInner {
	return v.value
}

func (v *NullableCompositeIndexSymbolInformationResponseInner) Set(val *CompositeIndexSymbolInformationResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCompositeIndexSymbolInformationResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCompositeIndexSymbolInformationResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompositeIndexSymbolInformationResponseInner(val *CompositeIndexSymbolInformationResponseInner) *NullableCompositeIndexSymbolInformationResponseInner {
	return &NullableCompositeIndexSymbolInformationResponseInner{value: val, isSet: true}
}

func (v NullableCompositeIndexSymbolInformationResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompositeIndexSymbolInformationResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
