/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryInsuranceFundBalanceSnapshotResponse1 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryInsuranceFundBalanceSnapshotResponse1{}

// QueryInsuranceFundBalanceSnapshotResponse1 struct for QueryInsuranceFundBalanceSnapshotResponse1
type QueryInsuranceFundBalanceSnapshotResponse1 struct {
	Symbols              []string                                                `json:"symbols,omitempty"`
	Assets               []QueryInsuranceFundBalanceSnapshotResponse1AssetsInner `json:"assets,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryInsuranceFundBalanceSnapshotResponse1 QueryInsuranceFundBalanceSnapshotResponse1

// NewQueryInsuranceFundBalanceSnapshotResponse1 instantiates a new QueryInsuranceFundBalanceSnapshotResponse1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryInsuranceFundBalanceSnapshotResponse1() *QueryInsuranceFundBalanceSnapshotResponse1 {
	this := QueryInsuranceFundBalanceSnapshotResponse1{}
	return &this
}

// NewQueryInsuranceFundBalanceSnapshotResponse1WithDefaults instantiates a new QueryInsuranceFundBalanceSnapshotResponse1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryInsuranceFundBalanceSnapshotResponse1WithDefaults() *QueryInsuranceFundBalanceSnapshotResponse1 {
	this := QueryInsuranceFundBalanceSnapshotResponse1{}
	return &this
}

// GetSymbols returns the Symbols field value if set, zero value otherwise.
func (o *QueryInsuranceFundBalanceSnapshotResponse1) GetSymbols() []string {
	if o == nil || common.IsNil(o.Symbols) {
		var ret []string
		return ret
	}
	return o.Symbols
}

// GetSymbolsOk returns a tuple with the Symbols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryInsuranceFundBalanceSnapshotResponse1) GetSymbolsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Symbols) {
		return nil, false
	}
	return o.Symbols, true
}

// HasSymbols returns a boolean if a field has been set.
func (o *QueryInsuranceFundBalanceSnapshotResponse1) HasSymbols() bool {
	if o != nil && !common.IsNil(o.Symbols) {
		return true
	}

	return false
}

// SetSymbols gets a reference to the given []string and assigns it to the Symbols field.
func (o *QueryInsuranceFundBalanceSnapshotResponse1) SetSymbols(v []string) {
	o.Symbols = v
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *QueryInsuranceFundBalanceSnapshotResponse1) GetAssets() []QueryInsuranceFundBalanceSnapshotResponse1AssetsInner {
	if o == nil || common.IsNil(o.Assets) {
		var ret []QueryInsuranceFundBalanceSnapshotResponse1AssetsInner
		return ret
	}
	return o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryInsuranceFundBalanceSnapshotResponse1) GetAssetsOk() ([]QueryInsuranceFundBalanceSnapshotResponse1AssetsInner, bool) {
	if o == nil || common.IsNil(o.Assets) {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *QueryInsuranceFundBalanceSnapshotResponse1) HasAssets() bool {
	if o != nil && !common.IsNil(o.Assets) {
		return true
	}

	return false
}

// SetAssets gets a reference to the given []QueryInsuranceFundBalanceSnapshotResponse1AssetsInner and assigns it to the Assets field.
func (o *QueryInsuranceFundBalanceSnapshotResponse1) SetAssets(v []QueryInsuranceFundBalanceSnapshotResponse1AssetsInner) {
	o.Assets = v
}

func (o QueryInsuranceFundBalanceSnapshotResponse1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryInsuranceFundBalanceSnapshotResponse1) ToMap() (map[string]interface{}, error) {
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

func (o *QueryInsuranceFundBalanceSnapshotResponse1) UnmarshalJSON(data []byte) (err error) {
	varQueryInsuranceFundBalanceSnapshotResponse1 := _QueryInsuranceFundBalanceSnapshotResponse1{}

	err = json.Unmarshal(data, &varQueryInsuranceFundBalanceSnapshotResponse1)

	if err != nil {
		return err
	}

	*o = QueryInsuranceFundBalanceSnapshotResponse1(varQueryInsuranceFundBalanceSnapshotResponse1)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbols")
		delete(additionalProperties, "assets")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryInsuranceFundBalanceSnapshotResponse1 struct {
	value *QueryInsuranceFundBalanceSnapshotResponse1
	isSet bool
}

func (v NullableQueryInsuranceFundBalanceSnapshotResponse1) Get() *QueryInsuranceFundBalanceSnapshotResponse1 {
	return v.value
}

func (v *NullableQueryInsuranceFundBalanceSnapshotResponse1) Set(val *QueryInsuranceFundBalanceSnapshotResponse1) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryInsuranceFundBalanceSnapshotResponse1) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryInsuranceFundBalanceSnapshotResponse1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryInsuranceFundBalanceSnapshotResponse1(val *QueryInsuranceFundBalanceSnapshotResponse1) *NullableQueryInsuranceFundBalanceSnapshotResponse1 {
	return &NullableQueryInsuranceFundBalanceSnapshotResponse1{value: val, isSet: true}
}

func (v NullableQueryInsuranceFundBalanceSnapshotResponse1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryInsuranceFundBalanceSnapshotResponse1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
