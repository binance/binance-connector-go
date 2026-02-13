/*
Binance Alpha REST API

OpenAPI Specification for the Binance Alpha REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetExchangeInfoResponseData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetExchangeInfoResponseData{}

// GetExchangeInfoResponseData struct for GetExchangeInfoResponseData
type GetExchangeInfoResponseData struct {
	Timezone             *string                                   `json:"timezone,omitempty"`
	Assets               []GetExchangeInfoResponseDataAssetsInner  `json:"assets,omitempty"`
	Symbols              []GetExchangeInfoResponseDataSymbolsInner `json:"symbols,omitempty"`
	OrderTypes           *string                                   `json:"orderTypes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetExchangeInfoResponseData GetExchangeInfoResponseData

// NewGetExchangeInfoResponseData instantiates a new GetExchangeInfoResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetExchangeInfoResponseData() *GetExchangeInfoResponseData {
	this := GetExchangeInfoResponseData{}
	return &this
}

// NewGetExchangeInfoResponseDataWithDefaults instantiates a new GetExchangeInfoResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetExchangeInfoResponseDataWithDefaults() *GetExchangeInfoResponseData {
	this := GetExchangeInfoResponseData{}
	return &this
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *GetExchangeInfoResponseData) GetTimezone() string {
	if o == nil || common.IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeInfoResponseData) GetTimezoneOk() (*string, bool) {
	if o == nil || common.IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *GetExchangeInfoResponseData) HasTimezone() bool {
	if o != nil && !common.IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *GetExchangeInfoResponseData) SetTimezone(v string) {
	o.Timezone = &v
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *GetExchangeInfoResponseData) GetAssets() []GetExchangeInfoResponseDataAssetsInner {
	if o == nil || common.IsNil(o.Assets) {
		var ret []GetExchangeInfoResponseDataAssetsInner
		return ret
	}
	return o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeInfoResponseData) GetAssetsOk() ([]GetExchangeInfoResponseDataAssetsInner, bool) {
	if o == nil || common.IsNil(o.Assets) {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *GetExchangeInfoResponseData) HasAssets() bool {
	if o != nil && !common.IsNil(o.Assets) {
		return true
	}

	return false
}

// SetAssets gets a reference to the given []GetExchangeInfoResponseDataAssetsInner and assigns it to the Assets field.
func (o *GetExchangeInfoResponseData) SetAssets(v []GetExchangeInfoResponseDataAssetsInner) {
	o.Assets = v
}

// GetSymbols returns the Symbols field value if set, zero value otherwise.
func (o *GetExchangeInfoResponseData) GetSymbols() []GetExchangeInfoResponseDataSymbolsInner {
	if o == nil || common.IsNil(o.Symbols) {
		var ret []GetExchangeInfoResponseDataSymbolsInner
		return ret
	}
	return o.Symbols
}

// GetSymbolsOk returns a tuple with the Symbols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeInfoResponseData) GetSymbolsOk() ([]GetExchangeInfoResponseDataSymbolsInner, bool) {
	if o == nil || common.IsNil(o.Symbols) {
		return nil, false
	}
	return o.Symbols, true
}

// HasSymbols returns a boolean if a field has been set.
func (o *GetExchangeInfoResponseData) HasSymbols() bool {
	if o != nil && !common.IsNil(o.Symbols) {
		return true
	}

	return false
}

// SetSymbols gets a reference to the given []GetExchangeInfoResponseDataSymbolsInner and assigns it to the Symbols field.
func (o *GetExchangeInfoResponseData) SetSymbols(v []GetExchangeInfoResponseDataSymbolsInner) {
	o.Symbols = v
}

// GetOrderTypes returns the OrderTypes field value if set, zero value otherwise.
func (o *GetExchangeInfoResponseData) GetOrderTypes() string {
	if o == nil || common.IsNil(o.OrderTypes) {
		var ret string
		return ret
	}
	return *o.OrderTypes
}

// GetOrderTypesOk returns a tuple with the OrderTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeInfoResponseData) GetOrderTypesOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderTypes) {
		return nil, false
	}
	return o.OrderTypes, true
}

// HasOrderTypes returns a boolean if a field has been set.
func (o *GetExchangeInfoResponseData) HasOrderTypes() bool {
	if o != nil && !common.IsNil(o.OrderTypes) {
		return true
	}

	return false
}

// SetOrderTypes gets a reference to the given string and assigns it to the OrderTypes field.
func (o *GetExchangeInfoResponseData) SetOrderTypes(v string) {
	o.OrderTypes = &v
}

func (o GetExchangeInfoResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetExchangeInfoResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	if !common.IsNil(o.Assets) {
		toSerialize["assets"] = o.Assets
	}
	if !common.IsNil(o.Symbols) {
		toSerialize["symbols"] = o.Symbols
	}
	if !common.IsNil(o.OrderTypes) {
		toSerialize["orderTypes"] = o.OrderTypes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetExchangeInfoResponseData) UnmarshalJSON(data []byte) (err error) {
	varGetExchangeInfoResponseData := _GetExchangeInfoResponseData{}

	err = json.Unmarshal(data, &varGetExchangeInfoResponseData)

	if err != nil {
		return err
	}

	*o = GetExchangeInfoResponseData(varGetExchangeInfoResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "timezone")
		delete(additionalProperties, "assets")
		delete(additionalProperties, "symbols")
		delete(additionalProperties, "orderTypes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetExchangeInfoResponseData struct {
	value *GetExchangeInfoResponseData
	isSet bool
}

func (v NullableGetExchangeInfoResponseData) Get() *GetExchangeInfoResponseData {
	return v.value
}

func (v *NullableGetExchangeInfoResponseData) Set(val *GetExchangeInfoResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetExchangeInfoResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetExchangeInfoResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetExchangeInfoResponseData(val *GetExchangeInfoResponseData) *NullableGetExchangeInfoResponseData {
	return &NullableGetExchangeInfoResponseData{value: val, isSet: true}
}

func (v NullableGetExchangeInfoResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetExchangeInfoResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
