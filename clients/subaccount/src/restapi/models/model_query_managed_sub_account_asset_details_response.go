/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryManagedSubAccountAssetDetailsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryManagedSubAccountAssetDetailsResponse{}

// QueryManagedSubAccountAssetDetailsResponse struct for QueryManagedSubAccountAssetDetailsResponse
type QueryManagedSubAccountAssetDetailsResponse struct {
	Items []QueryManagedSubAccountAssetDetailsResponseInner
}

// NewQueryManagedSubAccountAssetDetailsResponse instantiates a new QueryManagedSubAccountAssetDetailsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryManagedSubAccountAssetDetailsResponse() *QueryManagedSubAccountAssetDetailsResponse {
	this := QueryManagedSubAccountAssetDetailsResponse{}
	return &this
}

// NewQueryManagedSubAccountAssetDetailsResponseWithDefaults instantiates a new QueryManagedSubAccountAssetDetailsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryManagedSubAccountAssetDetailsResponseWithDefaults() *QueryManagedSubAccountAssetDetailsResponse {
	this := QueryManagedSubAccountAssetDetailsResponse{}
	return &this
}

func (o QueryManagedSubAccountAssetDetailsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryManagedSubAccountAssetDetailsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryManagedSubAccountAssetDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryManagedSubAccountAssetDetailsResponse struct {
	value QueryManagedSubAccountAssetDetailsResponse
	isSet bool
}

func (v NullableQueryManagedSubAccountAssetDetailsResponse) Get() QueryManagedSubAccountAssetDetailsResponse {
	return v.value
}

func (v *NullableQueryManagedSubAccountAssetDetailsResponse) Set(val QueryManagedSubAccountAssetDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryManagedSubAccountAssetDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryManagedSubAccountAssetDetailsResponse) Unset() {
	v.value = QueryManagedSubAccountAssetDetailsResponse{}
	v.isSet = false
}

func NewNullableQueryManagedSubAccountAssetDetailsResponse(val QueryManagedSubAccountAssetDetailsResponse) *NullableQueryManagedSubAccountAssetDetailsResponse {
	return &NullableQueryManagedSubAccountAssetDetailsResponse{value: val, isSet: true}
}

func (v NullableQueryManagedSubAccountAssetDetailsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryManagedSubAccountAssetDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
