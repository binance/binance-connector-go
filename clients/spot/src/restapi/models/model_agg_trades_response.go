/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AggTradesResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AggTradesResponse{}

// AggTradesResponse struct for AggTradesResponse
type AggTradesResponse struct {
	Items []AggTradesResponseInner
}

// NewAggTradesResponse instantiates a new AggTradesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggTradesResponse() *AggTradesResponse {
	this := AggTradesResponse{}
	return &this
}

// NewAggTradesResponseWithDefaults instantiates a new AggTradesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggTradesResponseWithDefaults() *AggTradesResponse {
	this := AggTradesResponse{}
	return &this
}

func (o AggTradesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AggTradesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *AggTradesResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableAggTradesResponse struct {
	value AggTradesResponse
	isSet bool
}

func (v NullableAggTradesResponse) Get() AggTradesResponse {
	return v.value
}

func (v *NullableAggTradesResponse) Set(val AggTradesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAggTradesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAggTradesResponse) Unset() {
	v.value = AggTradesResponse{}
	v.isSet = false
}

func NewNullableAggTradesResponse(val AggTradesResponse) *NullableAggTradesResponse {
	return &NullableAggTradesResponse{value: val, isSet: true}
}

func (v NullableAggTradesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggTradesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
