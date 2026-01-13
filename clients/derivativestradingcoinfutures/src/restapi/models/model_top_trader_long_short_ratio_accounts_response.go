/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the TopTraderLongShortRatioAccountsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TopTraderLongShortRatioAccountsResponse{}

// TopTraderLongShortRatioAccountsResponse struct for TopTraderLongShortRatioAccountsResponse
type TopTraderLongShortRatioAccountsResponse struct {
	Items []TopTraderLongShortRatioAccountsResponseInner
}

// NewTopTraderLongShortRatioAccountsResponse instantiates a new TopTraderLongShortRatioAccountsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopTraderLongShortRatioAccountsResponse() *TopTraderLongShortRatioAccountsResponse {
	this := TopTraderLongShortRatioAccountsResponse{}
	return &this
}

// NewTopTraderLongShortRatioAccountsResponseWithDefaults instantiates a new TopTraderLongShortRatioAccountsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopTraderLongShortRatioAccountsResponseWithDefaults() *TopTraderLongShortRatioAccountsResponse {
	this := TopTraderLongShortRatioAccountsResponse{}
	return &this
}

func (o TopTraderLongShortRatioAccountsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TopTraderLongShortRatioAccountsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *TopTraderLongShortRatioAccountsResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableTopTraderLongShortRatioAccountsResponse struct {
	value TopTraderLongShortRatioAccountsResponse
	isSet bool
}

func (v NullableTopTraderLongShortRatioAccountsResponse) Get() TopTraderLongShortRatioAccountsResponse {
	return v.value
}

func (v *NullableTopTraderLongShortRatioAccountsResponse) Set(val TopTraderLongShortRatioAccountsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTopTraderLongShortRatioAccountsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTopTraderLongShortRatioAccountsResponse) Unset() {
	v.value = TopTraderLongShortRatioAccountsResponse{}
	v.isSet = false
}

func NewNullableTopTraderLongShortRatioAccountsResponse(val TopTraderLongShortRatioAccountsResponse) *NullableTopTraderLongShortRatioAccountsResponse {
	return &NullableTopTraderLongShortRatioAccountsResponse{value: val, isSet: true}
}

func (v NullableTopTraderLongShortRatioAccountsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopTraderLongShortRatioAccountsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
