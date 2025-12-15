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

// checks if the GetFuturesPositionRiskOfSubAccountResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFuturesPositionRiskOfSubAccountResponse{}

// GetFuturesPositionRiskOfSubAccountResponse struct for GetFuturesPositionRiskOfSubAccountResponse
type GetFuturesPositionRiskOfSubAccountResponse struct {
	Items []GetFuturesPositionRiskOfSubAccountV2ResponseFuturePositionRiskVosInner
}

// NewGetFuturesPositionRiskOfSubAccountResponse instantiates a new GetFuturesPositionRiskOfSubAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFuturesPositionRiskOfSubAccountResponse() *GetFuturesPositionRiskOfSubAccountResponse {
	this := GetFuturesPositionRiskOfSubAccountResponse{}
	return &this
}

// NewGetFuturesPositionRiskOfSubAccountResponseWithDefaults instantiates a new GetFuturesPositionRiskOfSubAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFuturesPositionRiskOfSubAccountResponseWithDefaults() *GetFuturesPositionRiskOfSubAccountResponse {
	this := GetFuturesPositionRiskOfSubAccountResponse{}
	return &this
}

func (o GetFuturesPositionRiskOfSubAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFuturesPositionRiskOfSubAccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *GetFuturesPositionRiskOfSubAccountResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableGetFuturesPositionRiskOfSubAccountResponse struct {
	value GetFuturesPositionRiskOfSubAccountResponse
	isSet bool
}

func (v NullableGetFuturesPositionRiskOfSubAccountResponse) Get() GetFuturesPositionRiskOfSubAccountResponse {
	return v.value
}

func (v *NullableGetFuturesPositionRiskOfSubAccountResponse) Set(val GetFuturesPositionRiskOfSubAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFuturesPositionRiskOfSubAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFuturesPositionRiskOfSubAccountResponse) Unset() {
	v.value = GetFuturesPositionRiskOfSubAccountResponse{}
	v.isSet = false
}

func NewNullableGetFuturesPositionRiskOfSubAccountResponse(val GetFuturesPositionRiskOfSubAccountResponse) *NullableGetFuturesPositionRiskOfSubAccountResponse {
	return &NullableGetFuturesPositionRiskOfSubAccountResponse{value: val, isSet: true}
}

func (v NullableGetFuturesPositionRiskOfSubAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFuturesPositionRiskOfSubAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
