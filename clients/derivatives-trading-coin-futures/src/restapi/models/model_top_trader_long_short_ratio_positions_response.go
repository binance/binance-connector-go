/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the TopTraderLongShortRatioPositionsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TopTraderLongShortRatioPositionsResponse{}

// TopTraderLongShortRatioPositionsResponse struct for TopTraderLongShortRatioPositionsResponse
type TopTraderLongShortRatioPositionsResponse struct {
	Items []TopTraderLongShortRatioPositionsResponseInner
}

// NewTopTraderLongShortRatioPositionsResponse instantiates a new TopTraderLongShortRatioPositionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopTraderLongShortRatioPositionsResponse() *TopTraderLongShortRatioPositionsResponse {
	this := TopTraderLongShortRatioPositionsResponse{}
	return &this
}

// NewTopTraderLongShortRatioPositionsResponseWithDefaults instantiates a new TopTraderLongShortRatioPositionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopTraderLongShortRatioPositionsResponseWithDefaults() *TopTraderLongShortRatioPositionsResponse {
	this := TopTraderLongShortRatioPositionsResponse{}
	return &this
}

func (o TopTraderLongShortRatioPositionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TopTraderLongShortRatioPositionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *TopTraderLongShortRatioPositionsResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableTopTraderLongShortRatioPositionsResponse struct {
	value TopTraderLongShortRatioPositionsResponse
	isSet bool
}

func (v NullableTopTraderLongShortRatioPositionsResponse) Get() TopTraderLongShortRatioPositionsResponse {
	return v.value
}

func (v *NullableTopTraderLongShortRatioPositionsResponse) Set(val TopTraderLongShortRatioPositionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTopTraderLongShortRatioPositionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTopTraderLongShortRatioPositionsResponse) Unset() {
	v.value = TopTraderLongShortRatioPositionsResponse{}
	v.isSet = false
}

func NewNullableTopTraderLongShortRatioPositionsResponse(val TopTraderLongShortRatioPositionsResponse) *NullableTopTraderLongShortRatioPositionsResponse {
	return &NullableTopTraderLongShortRatioPositionsResponse{value: val, isSet: true}
}

func (v NullableTopTraderLongShortRatioPositionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopTraderLongShortRatioPositionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
