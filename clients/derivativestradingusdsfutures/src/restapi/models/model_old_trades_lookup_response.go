/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the OldTradesLookupResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OldTradesLookupResponse{}

// OldTradesLookupResponse struct for OldTradesLookupResponse
type OldTradesLookupResponse struct {
	Items []OldTradesLookupResponseInner
}

// NewOldTradesLookupResponse instantiates a new OldTradesLookupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOldTradesLookupResponse() *OldTradesLookupResponse {
	this := OldTradesLookupResponse{}
	return &this
}

// NewOldTradesLookupResponseWithDefaults instantiates a new OldTradesLookupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOldTradesLookupResponseWithDefaults() *OldTradesLookupResponse {
	this := OldTradesLookupResponse{}
	return &this
}

func (o OldTradesLookupResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OldTradesLookupResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *OldTradesLookupResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableOldTradesLookupResponse struct {
	value OldTradesLookupResponse
	isSet bool
}

func (v NullableOldTradesLookupResponse) Get() OldTradesLookupResponse {
	return v.value
}

func (v *NullableOldTradesLookupResponse) Set(val OldTradesLookupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOldTradesLookupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOldTradesLookupResponse) Unset() {
	v.value = OldTradesLookupResponse{}
	v.isSet = false
}

func NewNullableOldTradesLookupResponse(val OldTradesLookupResponse) *NullableOldTradesLookupResponse {
	return &NullableOldTradesLookupResponse{value: val, isSet: true}
}

func (v NullableOldTradesLookupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOldTradesLookupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
