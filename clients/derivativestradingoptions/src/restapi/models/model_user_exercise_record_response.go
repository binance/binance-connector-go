/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the UserExerciseRecordResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UserExerciseRecordResponse{}

// UserExerciseRecordResponse struct for UserExerciseRecordResponse
type UserExerciseRecordResponse struct {
	Items []UserExerciseRecordResponseInner
}

// NewUserExerciseRecordResponse instantiates a new UserExerciseRecordResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserExerciseRecordResponse() *UserExerciseRecordResponse {
	this := UserExerciseRecordResponse{}
	return &this
}

// NewUserExerciseRecordResponseWithDefaults instantiates a new UserExerciseRecordResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserExerciseRecordResponseWithDefaults() *UserExerciseRecordResponse {
	this := UserExerciseRecordResponse{}
	return &this
}

func (o UserExerciseRecordResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserExerciseRecordResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *UserExerciseRecordResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableUserExerciseRecordResponse struct {
	value UserExerciseRecordResponse
	isSet bool
}

func (v NullableUserExerciseRecordResponse) Get() UserExerciseRecordResponse {
	return v.value
}

func (v *NullableUserExerciseRecordResponse) Set(val UserExerciseRecordResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserExerciseRecordResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserExerciseRecordResponse) Unset() {
	v.value = UserExerciseRecordResponse{}
	v.isSet = false
}

func NewNullableUserExerciseRecordResponse(val UserExerciseRecordResponse) *NullableUserExerciseRecordResponse {
	return &NullableUserExerciseRecordResponse{value: val, isSet: true}
}

func (v NullableUserExerciseRecordResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserExerciseRecordResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
