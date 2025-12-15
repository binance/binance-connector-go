/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryMarginAvailableInventoryResponseAssets type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryMarginAvailableInventoryResponseAssets{}

// QueryMarginAvailableInventoryResponseAssets struct for QueryMarginAvailableInventoryResponseAssets
type QueryMarginAvailableInventoryResponseAssets struct {
	MATIC                *string `json:"MATIC,omitempty"`
	STPT                 *string `json:"STPT,omitempty"`
	TVK                  *string `json:"TVK,omitempty"`
	SHIB                 *string `json:"SHIB,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryMarginAvailableInventoryResponseAssets QueryMarginAvailableInventoryResponseAssets

// NewQueryMarginAvailableInventoryResponseAssets instantiates a new QueryMarginAvailableInventoryResponseAssets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryMarginAvailableInventoryResponseAssets() *QueryMarginAvailableInventoryResponseAssets {
	this := QueryMarginAvailableInventoryResponseAssets{}
	return &this
}

// NewQueryMarginAvailableInventoryResponseAssetsWithDefaults instantiates a new QueryMarginAvailableInventoryResponseAssets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryMarginAvailableInventoryResponseAssetsWithDefaults() *QueryMarginAvailableInventoryResponseAssets {
	this := QueryMarginAvailableInventoryResponseAssets{}
	return &this
}

// GetMATIC returns the MATIC field value if set, zero value otherwise.
func (o *QueryMarginAvailableInventoryResponseAssets) GetMATIC() string {
	if o == nil || common.IsNil(o.MATIC) {
		var ret string
		return ret
	}
	return *o.MATIC
}

// GetMATICOk returns a tuple with the MATIC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAvailableInventoryResponseAssets) GetMATICOk() (*string, bool) {
	if o == nil || common.IsNil(o.MATIC) {
		return nil, false
	}
	return o.MATIC, true
}

// HasMATIC returns a boolean if a field has been set.
func (o *QueryMarginAvailableInventoryResponseAssets) HasMATIC() bool {
	if o != nil && !common.IsNil(o.MATIC) {
		return true
	}

	return false
}

// SetMATIC gets a reference to the given string and assigns it to the MATIC field.
func (o *QueryMarginAvailableInventoryResponseAssets) SetMATIC(v string) {
	o.MATIC = &v
}

// GetSTPT returns the STPT field value if set, zero value otherwise.
func (o *QueryMarginAvailableInventoryResponseAssets) GetSTPT() string {
	if o == nil || common.IsNil(o.STPT) {
		var ret string
		return ret
	}
	return *o.STPT
}

// GetSTPTOk returns a tuple with the STPT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAvailableInventoryResponseAssets) GetSTPTOk() (*string, bool) {
	if o == nil || common.IsNil(o.STPT) {
		return nil, false
	}
	return o.STPT, true
}

// HasSTPT returns a boolean if a field has been set.
func (o *QueryMarginAvailableInventoryResponseAssets) HasSTPT() bool {
	if o != nil && !common.IsNil(o.STPT) {
		return true
	}

	return false
}

// SetSTPT gets a reference to the given string and assigns it to the STPT field.
func (o *QueryMarginAvailableInventoryResponseAssets) SetSTPT(v string) {
	o.STPT = &v
}

// GetTVK returns the TVK field value if set, zero value otherwise.
func (o *QueryMarginAvailableInventoryResponseAssets) GetTVK() string {
	if o == nil || common.IsNil(o.TVK) {
		var ret string
		return ret
	}
	return *o.TVK
}

// GetTVKOk returns a tuple with the TVK field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAvailableInventoryResponseAssets) GetTVKOk() (*string, bool) {
	if o == nil || common.IsNil(o.TVK) {
		return nil, false
	}
	return o.TVK, true
}

// HasTVK returns a boolean if a field has been set.
func (o *QueryMarginAvailableInventoryResponseAssets) HasTVK() bool {
	if o != nil && !common.IsNil(o.TVK) {
		return true
	}

	return false
}

// SetTVK gets a reference to the given string and assigns it to the TVK field.
func (o *QueryMarginAvailableInventoryResponseAssets) SetTVK(v string) {
	o.TVK = &v
}

// GetSHIB returns the SHIB field value if set, zero value otherwise.
func (o *QueryMarginAvailableInventoryResponseAssets) GetSHIB() string {
	if o == nil || common.IsNil(o.SHIB) {
		var ret string
		return ret
	}
	return *o.SHIB
}

// GetSHIBOk returns a tuple with the SHIB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAvailableInventoryResponseAssets) GetSHIBOk() (*string, bool) {
	if o == nil || common.IsNil(o.SHIB) {
		return nil, false
	}
	return o.SHIB, true
}

// HasSHIB returns a boolean if a field has been set.
func (o *QueryMarginAvailableInventoryResponseAssets) HasSHIB() bool {
	if o != nil && !common.IsNil(o.SHIB) {
		return true
	}

	return false
}

// SetSHIB gets a reference to the given string and assigns it to the SHIB field.
func (o *QueryMarginAvailableInventoryResponseAssets) SetSHIB(v string) {
	o.SHIB = &v
}

func (o QueryMarginAvailableInventoryResponseAssets) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryMarginAvailableInventoryResponseAssets) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.MATIC) {
		toSerialize["MATIC"] = o.MATIC
	}
	if !common.IsNil(o.STPT) {
		toSerialize["STPT"] = o.STPT
	}
	if !common.IsNil(o.TVK) {
		toSerialize["TVK"] = o.TVK
	}
	if !common.IsNil(o.SHIB) {
		toSerialize["SHIB"] = o.SHIB
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryMarginAvailableInventoryResponseAssets) UnmarshalJSON(data []byte) (err error) {
	varQueryMarginAvailableInventoryResponseAssets := _QueryMarginAvailableInventoryResponseAssets{}

	err = json.Unmarshal(data, &varQueryMarginAvailableInventoryResponseAssets)

	if err != nil {
		return err
	}

	*o = QueryMarginAvailableInventoryResponseAssets(varQueryMarginAvailableInventoryResponseAssets)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "MATIC")
		delete(additionalProperties, "STPT")
		delete(additionalProperties, "TVK")
		delete(additionalProperties, "SHIB")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryMarginAvailableInventoryResponseAssets struct {
	value *QueryMarginAvailableInventoryResponseAssets
	isSet bool
}

func (v NullableQueryMarginAvailableInventoryResponseAssets) Get() *QueryMarginAvailableInventoryResponseAssets {
	return v.value
}

func (v *NullableQueryMarginAvailableInventoryResponseAssets) Set(val *QueryMarginAvailableInventoryResponseAssets) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryMarginAvailableInventoryResponseAssets) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryMarginAvailableInventoryResponseAssets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryMarginAvailableInventoryResponseAssets(val *QueryMarginAvailableInventoryResponseAssets) *NullableQueryMarginAvailableInventoryResponseAssets {
	return &NullableQueryMarginAvailableInventoryResponseAssets{value: val, isSet: true}
}

func (v NullableQueryMarginAvailableInventoryResponseAssets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryMarginAvailableInventoryResponseAssets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
