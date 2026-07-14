/*
Spot REST API

Access market data, manage accounts, and trade on Binance Spot.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the KlinesResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &KlinesResponse{}

// KlinesResponse struct for KlinesResponse
type KlinesResponse struct {
	Items []KlinesItem
}

// NewKlinesResponse instantiates a new KlinesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKlinesResponse() *KlinesResponse {
	this := KlinesResponse{}
	return &this
}

// NewKlinesResponseWithDefaults instantiates a new KlinesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKlinesResponseWithDefaults() *KlinesResponse {
	this := KlinesResponse{}
	return &this
}

func (o KlinesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KlinesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *KlinesResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableKlinesResponse struct {
	value KlinesResponse
	isSet bool
}

func (v NullableKlinesResponse) Get() KlinesResponse {
	return v.value
}

func (v *NullableKlinesResponse) Set(val KlinesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKlinesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKlinesResponse) Unset() {
	v.value = KlinesResponse{}
	v.isSet = false
}

func NewNullableKlinesResponse(val KlinesResponse) *NullableKlinesResponse {
	return &NullableKlinesResponse{value: val, isSet: true}
}

func (v NullableKlinesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlinesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
