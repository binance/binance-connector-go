/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the CompositeIndexSymbolInformationResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CompositeIndexSymbolInformationResponse{}

// CompositeIndexSymbolInformationResponse struct for CompositeIndexSymbolInformationResponse
type CompositeIndexSymbolInformationResponse struct {
	Items []CompositeIndexSymbolInformationResponseInner
}

// NewCompositeIndexSymbolInformationResponse instantiates a new CompositeIndexSymbolInformationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompositeIndexSymbolInformationResponse() *CompositeIndexSymbolInformationResponse {
	this := CompositeIndexSymbolInformationResponse{}
	return &this
}

// NewCompositeIndexSymbolInformationResponseWithDefaults instantiates a new CompositeIndexSymbolInformationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompositeIndexSymbolInformationResponseWithDefaults() *CompositeIndexSymbolInformationResponse {
	this := CompositeIndexSymbolInformationResponse{}
	return &this
}

func (o CompositeIndexSymbolInformationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompositeIndexSymbolInformationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *CompositeIndexSymbolInformationResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableCompositeIndexSymbolInformationResponse struct {
	value CompositeIndexSymbolInformationResponse
	isSet bool
}

func (v NullableCompositeIndexSymbolInformationResponse) Get() CompositeIndexSymbolInformationResponse {
	return v.value
}

func (v *NullableCompositeIndexSymbolInformationResponse) Set(val CompositeIndexSymbolInformationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCompositeIndexSymbolInformationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCompositeIndexSymbolInformationResponse) Unset() {
	v.value = CompositeIndexSymbolInformationResponse{}
	v.isSet = false
}

func NewNullableCompositeIndexSymbolInformationResponse(val CompositeIndexSymbolInformationResponse) *NullableCompositeIndexSymbolInformationResponse {
	return &NullableCompositeIndexSymbolInformationResponse{value: val, isSet: true}
}

func (v NullableCompositeIndexSymbolInformationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompositeIndexSymbolInformationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
