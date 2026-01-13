/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryInsuranceFundBalanceSnapshotResponse2 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryInsuranceFundBalanceSnapshotResponse2{}

// QueryInsuranceFundBalanceSnapshotResponse2 struct for QueryInsuranceFundBalanceSnapshotResponse2
type QueryInsuranceFundBalanceSnapshotResponse2 struct {
	Items []QueryInsuranceFundBalanceSnapshotResponse2Inner
}

// NewQueryInsuranceFundBalanceSnapshotResponse2 instantiates a new QueryInsuranceFundBalanceSnapshotResponse2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryInsuranceFundBalanceSnapshotResponse2() *QueryInsuranceFundBalanceSnapshotResponse2 {
	this := QueryInsuranceFundBalanceSnapshotResponse2{}
	return &this
}

// NewQueryInsuranceFundBalanceSnapshotResponse2WithDefaults instantiates a new QueryInsuranceFundBalanceSnapshotResponse2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryInsuranceFundBalanceSnapshotResponse2WithDefaults() *QueryInsuranceFundBalanceSnapshotResponse2 {
	this := QueryInsuranceFundBalanceSnapshotResponse2{}
	return &this
}

func (o QueryInsuranceFundBalanceSnapshotResponse2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryInsuranceFundBalanceSnapshotResponse2) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryInsuranceFundBalanceSnapshotResponse2) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryInsuranceFundBalanceSnapshotResponse2 struct {
	value QueryInsuranceFundBalanceSnapshotResponse2
	isSet bool
}

func (v NullableQueryInsuranceFundBalanceSnapshotResponse2) Get() QueryInsuranceFundBalanceSnapshotResponse2 {
	return v.value
}

func (v *NullableQueryInsuranceFundBalanceSnapshotResponse2) Set(val QueryInsuranceFundBalanceSnapshotResponse2) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryInsuranceFundBalanceSnapshotResponse2) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryInsuranceFundBalanceSnapshotResponse2) Unset() {
	v.value = QueryInsuranceFundBalanceSnapshotResponse2{}
	v.isSet = false
}

func NewNullableQueryInsuranceFundBalanceSnapshotResponse2(val QueryInsuranceFundBalanceSnapshotResponse2) *NullableQueryInsuranceFundBalanceSnapshotResponse2 {
	return &NullableQueryInsuranceFundBalanceSnapshotResponse2{value: val, isSet: true}
}

func (v NullableQueryInsuranceFundBalanceSnapshotResponse2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryInsuranceFundBalanceSnapshotResponse2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
