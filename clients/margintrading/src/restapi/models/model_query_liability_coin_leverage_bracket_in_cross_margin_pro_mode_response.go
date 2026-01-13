/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponse{}

// QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponse struct for QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponse
type QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponse struct {
	Items []QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseInner
}

// NewQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponse instantiates a new QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponse() *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponse {
	this := QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponse{}
	return &this
}

// NewQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseWithDefaults instantiates a new QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponseWithDefaults() *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponse {
	this := QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponse{}
	return &this
}

func (o QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponse struct {
	value QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponse
	isSet bool
}

func (v NullableQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponse) Get() QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponse {
	return v.value
}

func (v *NullableQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponse) Set(val QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponse) Unset() {
	v.value = QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponse{}
	v.isSet = false
}

func NewNullableQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponse(val QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponse) *NullableQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponse {
	return &NullableQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponse{value: val, isSet: true}
}

func (v NullableQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryLiabilityCoinLeverageBracketInCrossMarginProModeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
