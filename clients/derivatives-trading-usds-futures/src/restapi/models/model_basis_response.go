/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the BasisResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BasisResponse{}

// BasisResponse struct for BasisResponse
type BasisResponse struct {
	Items []BasisResponseInner
}

// NewBasisResponse instantiates a new BasisResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBasisResponse() *BasisResponse {
	this := BasisResponse{}
	return &this
}

// NewBasisResponseWithDefaults instantiates a new BasisResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBasisResponseWithDefaults() *BasisResponse {
	this := BasisResponse{}
	return &this
}

func (o BasisResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BasisResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *BasisResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableBasisResponse struct {
	value BasisResponse
	isSet bool
}

func (v NullableBasisResponse) Get() BasisResponse {
	return v.value
}

func (v *NullableBasisResponse) Set(val BasisResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBasisResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBasisResponse) Unset() {
	v.value = BasisResponse{}
	v.isSet = false
}

func NewNullableBasisResponse(val BasisResponse) *NullableBasisResponse {
	return &NullableBasisResponse{value: val, isSet: true}
}

func (v NullableBasisResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBasisResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
