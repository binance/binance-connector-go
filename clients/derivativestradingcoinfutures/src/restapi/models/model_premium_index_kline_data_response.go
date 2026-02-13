/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the PremiumIndexKlineDataResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PremiumIndexKlineDataResponse{}

// PremiumIndexKlineDataResponse struct for PremiumIndexKlineDataResponse
type PremiumIndexKlineDataResponse struct {
	Items []PremiumIndexKlineDataResponseItem
}

// NewPremiumIndexKlineDataResponse instantiates a new PremiumIndexKlineDataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPremiumIndexKlineDataResponse() *PremiumIndexKlineDataResponse {
	this := PremiumIndexKlineDataResponse{}
	return &this
}

// NewPremiumIndexKlineDataResponseWithDefaults instantiates a new PremiumIndexKlineDataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPremiumIndexKlineDataResponseWithDefaults() *PremiumIndexKlineDataResponse {
	this := PremiumIndexKlineDataResponse{}
	return &this
}

func (o PremiumIndexKlineDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PremiumIndexKlineDataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *PremiumIndexKlineDataResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullablePremiumIndexKlineDataResponse struct {
	value PremiumIndexKlineDataResponse
	isSet bool
}

func (v NullablePremiumIndexKlineDataResponse) Get() PremiumIndexKlineDataResponse {
	return v.value
}

func (v *NullablePremiumIndexKlineDataResponse) Set(val PremiumIndexKlineDataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePremiumIndexKlineDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePremiumIndexKlineDataResponse) Unset() {
	v.value = PremiumIndexKlineDataResponse{}
	v.isSet = false
}

func NewNullablePremiumIndexKlineDataResponse(val PremiumIndexKlineDataResponse) *NullablePremiumIndexKlineDataResponse {
	return &NullablePremiumIndexKlineDataResponse{value: val, isSet: true}
}

func (v NullablePremiumIndexKlineDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePremiumIndexKlineDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
