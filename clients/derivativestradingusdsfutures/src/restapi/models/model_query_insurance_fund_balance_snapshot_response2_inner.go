/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryInsuranceFundBalanceSnapshotResponse2Inner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryInsuranceFundBalanceSnapshotResponse2Inner{}

// QueryInsuranceFundBalanceSnapshotResponse2Inner struct for QueryInsuranceFundBalanceSnapshotResponse2Inner
type QueryInsuranceFundBalanceSnapshotResponse2Inner struct {
	Symbols              []string                                                     `json:"symbols,omitempty"`
	Assets               []QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner `json:"assets,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryInsuranceFundBalanceSnapshotResponse2Inner QueryInsuranceFundBalanceSnapshotResponse2Inner

// NewQueryInsuranceFundBalanceSnapshotResponse2Inner instantiates a new QueryInsuranceFundBalanceSnapshotResponse2Inner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryInsuranceFundBalanceSnapshotResponse2Inner() *QueryInsuranceFundBalanceSnapshotResponse2Inner {
	this := QueryInsuranceFundBalanceSnapshotResponse2Inner{}
	return &this
}

// NewQueryInsuranceFundBalanceSnapshotResponse2InnerWithDefaults instantiates a new QueryInsuranceFundBalanceSnapshotResponse2Inner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryInsuranceFundBalanceSnapshotResponse2InnerWithDefaults() *QueryInsuranceFundBalanceSnapshotResponse2Inner {
	this := QueryInsuranceFundBalanceSnapshotResponse2Inner{}
	return &this
}

// GetSymbols returns the Symbols field value if set, zero value otherwise.
func (o *QueryInsuranceFundBalanceSnapshotResponse2Inner) GetSymbols() []string {
	if o == nil || common.IsNil(o.Symbols) {
		var ret []string
		return ret
	}
	return o.Symbols
}

// GetSymbolsOk returns a tuple with the Symbols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryInsuranceFundBalanceSnapshotResponse2Inner) GetSymbolsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Symbols) {
		return nil, false
	}
	return o.Symbols, true
}

// HasSymbols returns a boolean if a field has been set.
func (o *QueryInsuranceFundBalanceSnapshotResponse2Inner) HasSymbols() bool {
	if o != nil && !common.IsNil(o.Symbols) {
		return true
	}

	return false
}

// SetSymbols gets a reference to the given []string and assigns it to the Symbols field.
func (o *QueryInsuranceFundBalanceSnapshotResponse2Inner) SetSymbols(v []string) {
	o.Symbols = v
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *QueryInsuranceFundBalanceSnapshotResponse2Inner) GetAssets() []QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner {
	if o == nil || common.IsNil(o.Assets) {
		var ret []QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner
		return ret
	}
	return o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryInsuranceFundBalanceSnapshotResponse2Inner) GetAssetsOk() ([]QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner, bool) {
	if o == nil || common.IsNil(o.Assets) {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *QueryInsuranceFundBalanceSnapshotResponse2Inner) HasAssets() bool {
	if o != nil && !common.IsNil(o.Assets) {
		return true
	}

	return false
}

// SetAssets gets a reference to the given []QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner and assigns it to the Assets field.
func (o *QueryInsuranceFundBalanceSnapshotResponse2Inner) SetAssets(v []QueryInsuranceFundBalanceSnapshotResponse2InnerAssetsInner) {
	o.Assets = v
}

func (o QueryInsuranceFundBalanceSnapshotResponse2Inner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryInsuranceFundBalanceSnapshotResponse2Inner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbols) {
		toSerialize["symbols"] = o.Symbols
	}
	if !common.IsNil(o.Assets) {
		toSerialize["assets"] = o.Assets
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryInsuranceFundBalanceSnapshotResponse2Inner) UnmarshalJSON(data []byte) (err error) {
	varQueryInsuranceFundBalanceSnapshotResponse2Inner := _QueryInsuranceFundBalanceSnapshotResponse2Inner{}

	err = json.Unmarshal(data, &varQueryInsuranceFundBalanceSnapshotResponse2Inner)

	if err != nil {
		return err
	}

	*o = QueryInsuranceFundBalanceSnapshotResponse2Inner(varQueryInsuranceFundBalanceSnapshotResponse2Inner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbols")
		delete(additionalProperties, "assets")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryInsuranceFundBalanceSnapshotResponse2Inner struct {
	value *QueryInsuranceFundBalanceSnapshotResponse2Inner
	isSet bool
}

func (v NullableQueryInsuranceFundBalanceSnapshotResponse2Inner) Get() *QueryInsuranceFundBalanceSnapshotResponse2Inner {
	return v.value
}

func (v *NullableQueryInsuranceFundBalanceSnapshotResponse2Inner) Set(val *QueryInsuranceFundBalanceSnapshotResponse2Inner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryInsuranceFundBalanceSnapshotResponse2Inner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryInsuranceFundBalanceSnapshotResponse2Inner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryInsuranceFundBalanceSnapshotResponse2Inner(val *QueryInsuranceFundBalanceSnapshotResponse2Inner) *NullableQueryInsuranceFundBalanceSnapshotResponse2Inner {
	return &NullableQueryInsuranceFundBalanceSnapshotResponse2Inner{value: val, isSet: true}
}

func (v NullableQueryInsuranceFundBalanceSnapshotResponse2Inner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryInsuranceFundBalanceSnapshotResponse2Inner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
