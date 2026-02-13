/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AllCoinsInformationResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AllCoinsInformationResponse{}

// AllCoinsInformationResponse struct for AllCoinsInformationResponse
type AllCoinsInformationResponse struct {
	Items []AllCoinsInformationResponseInner
}

// NewAllCoinsInformationResponse instantiates a new AllCoinsInformationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllCoinsInformationResponse() *AllCoinsInformationResponse {
	this := AllCoinsInformationResponse{}
	return &this
}

// NewAllCoinsInformationResponseWithDefaults instantiates a new AllCoinsInformationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllCoinsInformationResponseWithDefaults() *AllCoinsInformationResponse {
	this := AllCoinsInformationResponse{}
	return &this
}

func (o AllCoinsInformationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllCoinsInformationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *AllCoinsInformationResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableAllCoinsInformationResponse struct {
	value AllCoinsInformationResponse
	isSet bool
}

func (v NullableAllCoinsInformationResponse) Get() AllCoinsInformationResponse {
	return v.value
}

func (v *NullableAllCoinsInformationResponse) Set(val AllCoinsInformationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAllCoinsInformationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAllCoinsInformationResponse) Unset() {
	v.value = AllCoinsInformationResponse{}
	v.isSet = false
}

func NewNullableAllCoinsInformationResponse(val AllCoinsInformationResponse) *NullableAllCoinsInformationResponse {
	return &NullableAllCoinsInformationResponse{value: val, isSet: true}
}

func (v NullableAllCoinsInformationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllCoinsInformationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
