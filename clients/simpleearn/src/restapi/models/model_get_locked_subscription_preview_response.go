/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetLockedSubscriptionPreviewResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetLockedSubscriptionPreviewResponse{}

// GetLockedSubscriptionPreviewResponse struct for GetLockedSubscriptionPreviewResponse
type GetLockedSubscriptionPreviewResponse struct {
	Items []GetLockedSubscriptionPreviewResponseInner
}

// NewGetLockedSubscriptionPreviewResponse instantiates a new GetLockedSubscriptionPreviewResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLockedSubscriptionPreviewResponse() *GetLockedSubscriptionPreviewResponse {
	this := GetLockedSubscriptionPreviewResponse{}
	return &this
}

// NewGetLockedSubscriptionPreviewResponseWithDefaults instantiates a new GetLockedSubscriptionPreviewResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLockedSubscriptionPreviewResponseWithDefaults() *GetLockedSubscriptionPreviewResponse {
	this := GetLockedSubscriptionPreviewResponse{}
	return &this
}

func (o GetLockedSubscriptionPreviewResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLockedSubscriptionPreviewResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *GetLockedSubscriptionPreviewResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableGetLockedSubscriptionPreviewResponse struct {
	value GetLockedSubscriptionPreviewResponse
	isSet bool
}

func (v NullableGetLockedSubscriptionPreviewResponse) Get() GetLockedSubscriptionPreviewResponse {
	return v.value
}

func (v *NullableGetLockedSubscriptionPreviewResponse) Set(val GetLockedSubscriptionPreviewResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLockedSubscriptionPreviewResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLockedSubscriptionPreviewResponse) Unset() {
	v.value = GetLockedSubscriptionPreviewResponse{}
	v.isSet = false
}

func NewNullableGetLockedSubscriptionPreviewResponse(val GetLockedSubscriptionPreviewResponse) *NullableGetLockedSubscriptionPreviewResponse {
	return &NullableGetLockedSubscriptionPreviewResponse{value: val, isSet: true}
}

func (v NullableGetLockedSubscriptionPreviewResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLockedSubscriptionPreviewResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
