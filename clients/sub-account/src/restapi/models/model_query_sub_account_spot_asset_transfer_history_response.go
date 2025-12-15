/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QuerySubAccountSpotAssetTransferHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QuerySubAccountSpotAssetTransferHistoryResponse{}

// QuerySubAccountSpotAssetTransferHistoryResponse struct for QuerySubAccountSpotAssetTransferHistoryResponse
type QuerySubAccountSpotAssetTransferHistoryResponse struct {
	Items []QuerySubAccountSpotAssetTransferHistoryResponseInner
}

// NewQuerySubAccountSpotAssetTransferHistoryResponse instantiates a new QuerySubAccountSpotAssetTransferHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuerySubAccountSpotAssetTransferHistoryResponse() *QuerySubAccountSpotAssetTransferHistoryResponse {
	this := QuerySubAccountSpotAssetTransferHistoryResponse{}
	return &this
}

// NewQuerySubAccountSpotAssetTransferHistoryResponseWithDefaults instantiates a new QuerySubAccountSpotAssetTransferHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuerySubAccountSpotAssetTransferHistoryResponseWithDefaults() *QuerySubAccountSpotAssetTransferHistoryResponse {
	this := QuerySubAccountSpotAssetTransferHistoryResponse{}
	return &this
}

func (o QuerySubAccountSpotAssetTransferHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuerySubAccountSpotAssetTransferHistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QuerySubAccountSpotAssetTransferHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQuerySubAccountSpotAssetTransferHistoryResponse struct {
	value QuerySubAccountSpotAssetTransferHistoryResponse
	isSet bool
}

func (v NullableQuerySubAccountSpotAssetTransferHistoryResponse) Get() QuerySubAccountSpotAssetTransferHistoryResponse {
	return v.value
}

func (v *NullableQuerySubAccountSpotAssetTransferHistoryResponse) Set(val QuerySubAccountSpotAssetTransferHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQuerySubAccountSpotAssetTransferHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQuerySubAccountSpotAssetTransferHistoryResponse) Unset() {
	v.value = QuerySubAccountSpotAssetTransferHistoryResponse{}
	v.isSet = false
}

func NewNullableQuerySubAccountSpotAssetTransferHistoryResponse(val QuerySubAccountSpotAssetTransferHistoryResponse) *NullableQuerySubAccountSpotAssetTransferHistoryResponse {
	return &NullableQuerySubAccountSpotAssetTransferHistoryResponse{value: val, isSet: true}
}

func (v NullableQuerySubAccountSpotAssetTransferHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuerySubAccountSpotAssetTransferHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
