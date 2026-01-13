/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the UsersForceOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UsersForceOrdersResponse{}

// UsersForceOrdersResponse struct for UsersForceOrdersResponse
type UsersForceOrdersResponse struct {
	Items []UsersForceOrdersResponseInner
}

// NewUsersForceOrdersResponse instantiates a new UsersForceOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersForceOrdersResponse() *UsersForceOrdersResponse {
	this := UsersForceOrdersResponse{}
	return &this
}

// NewUsersForceOrdersResponseWithDefaults instantiates a new UsersForceOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersForceOrdersResponseWithDefaults() *UsersForceOrdersResponse {
	this := UsersForceOrdersResponse{}
	return &this
}

func (o UsersForceOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsersForceOrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *UsersForceOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableUsersForceOrdersResponse struct {
	value UsersForceOrdersResponse
	isSet bool
}

func (v NullableUsersForceOrdersResponse) Get() UsersForceOrdersResponse {
	return v.value
}

func (v *NullableUsersForceOrdersResponse) Set(val UsersForceOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersForceOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersForceOrdersResponse) Unset() {
	v.value = UsersForceOrdersResponse{}
	v.isSet = false
}

func NewNullableUsersForceOrdersResponse(val UsersForceOrdersResponse) *NullableUsersForceOrdersResponse {
	return &NullableUsersForceOrdersResponse{value: val, isSet: true}
}

func (v NullableUsersForceOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersForceOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
