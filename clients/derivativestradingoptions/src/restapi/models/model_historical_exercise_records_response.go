/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the HistoricalExerciseRecordsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &HistoricalExerciseRecordsResponse{}

// HistoricalExerciseRecordsResponse struct for HistoricalExerciseRecordsResponse
type HistoricalExerciseRecordsResponse struct {
	Items []HistoricalExerciseRecordsResponseInner
}

// NewHistoricalExerciseRecordsResponse instantiates a new HistoricalExerciseRecordsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoricalExerciseRecordsResponse() *HistoricalExerciseRecordsResponse {
	this := HistoricalExerciseRecordsResponse{}
	return &this
}

// NewHistoricalExerciseRecordsResponseWithDefaults instantiates a new HistoricalExerciseRecordsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoricalExerciseRecordsResponseWithDefaults() *HistoricalExerciseRecordsResponse {
	this := HistoricalExerciseRecordsResponse{}
	return &this
}

func (o HistoricalExerciseRecordsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoricalExerciseRecordsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *HistoricalExerciseRecordsResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableHistoricalExerciseRecordsResponse struct {
	value HistoricalExerciseRecordsResponse
	isSet bool
}

func (v NullableHistoricalExerciseRecordsResponse) Get() HistoricalExerciseRecordsResponse {
	return v.value
}

func (v *NullableHistoricalExerciseRecordsResponse) Set(val HistoricalExerciseRecordsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoricalExerciseRecordsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoricalExerciseRecordsResponse) Unset() {
	v.value = HistoricalExerciseRecordsResponse{}
	v.isSet = false
}

func NewNullableHistoricalExerciseRecordsResponse(val HistoricalExerciseRecordsResponse) *NullableHistoricalExerciseRecordsResponse {
	return &NullableHistoricalExerciseRecordsResponse{value: val, isSet: true}
}

func (v NullableHistoricalExerciseRecordsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoricalExerciseRecordsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
