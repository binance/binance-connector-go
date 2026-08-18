/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ListOtcBlocktradesResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ListOtcBlocktradesResponse{}

// ListOtcBlocktradesResponse struct for ListOtcBlocktradesResponse
type ListOtcBlocktradesResponse struct {
	Cursor               *string                                      `json:"cursor,omitempty"`
	Blocktrades          []ListOtcBlocktradesResponseBlocktradesInner `json:"blocktrades,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListOtcBlocktradesResponse ListOtcBlocktradesResponse

// NewListOtcBlocktradesResponse instantiates a new ListOtcBlocktradesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOtcBlocktradesResponse() *ListOtcBlocktradesResponse {
	this := ListOtcBlocktradesResponse{}
	return &this
}

// NewListOtcBlocktradesResponseWithDefaults instantiates a new ListOtcBlocktradesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOtcBlocktradesResponseWithDefaults() *ListOtcBlocktradesResponse {
	this := ListOtcBlocktradesResponse{}
	return &this
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *ListOtcBlocktradesResponse) GetCursor() string {
	if o == nil || common.IsNil(o.Cursor) {
		var ret string
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOtcBlocktradesResponse) GetCursorOk() (*string, bool) {
	if o == nil || common.IsNil(o.Cursor) {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *ListOtcBlocktradesResponse) HasCursor() bool {
	if o != nil && !common.IsNil(o.Cursor) {
		return true
	}

	return false
}

// SetCursor gets a reference to the given string and assigns it to the Cursor field.
func (o *ListOtcBlocktradesResponse) SetCursor(v string) {
	o.Cursor = &v
}

// GetBlocktrades returns the Blocktrades field value if set, zero value otherwise.
func (o *ListOtcBlocktradesResponse) GetBlocktrades() []ListOtcBlocktradesResponseBlocktradesInner {
	if o == nil || common.IsNil(o.Blocktrades) {
		var ret []ListOtcBlocktradesResponseBlocktradesInner
		return ret
	}
	return o.Blocktrades
}

// GetBlocktradesOk returns a tuple with the Blocktrades field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOtcBlocktradesResponse) GetBlocktradesOk() ([]ListOtcBlocktradesResponseBlocktradesInner, bool) {
	if o == nil || common.IsNil(o.Blocktrades) {
		return nil, false
	}
	return o.Blocktrades, true
}

// HasBlocktrades returns a boolean if a field has been set.
func (o *ListOtcBlocktradesResponse) HasBlocktrades() bool {
	if o != nil && !common.IsNil(o.Blocktrades) {
		return true
	}

	return false
}

// SetBlocktrades gets a reference to the given []ListOtcBlocktradesResponseBlocktradesInner and assigns it to the Blocktrades field.
func (o *ListOtcBlocktradesResponse) SetBlocktrades(v []ListOtcBlocktradesResponseBlocktradesInner) {
	o.Blocktrades = v
}

func (o ListOtcBlocktradesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListOtcBlocktradesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Cursor) {
		toSerialize["cursor"] = o.Cursor
	}
	if !common.IsNil(o.Blocktrades) {
		toSerialize["blocktrades"] = o.Blocktrades
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListOtcBlocktradesResponse) UnmarshalJSON(data []byte) (err error) {
	varListOtcBlocktradesResponse := _ListOtcBlocktradesResponse{}

	err = json.Unmarshal(data, &varListOtcBlocktradesResponse)

	if err != nil {
		return err
	}

	*o = ListOtcBlocktradesResponse(varListOtcBlocktradesResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cursor")
		delete(additionalProperties, "blocktrades")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListOtcBlocktradesResponse struct {
	value *ListOtcBlocktradesResponse
	isSet bool
}

func (v NullableListOtcBlocktradesResponse) Get() *ListOtcBlocktradesResponse {
	return v.value
}

func (v *NullableListOtcBlocktradesResponse) Set(val *ListOtcBlocktradesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListOtcBlocktradesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListOtcBlocktradesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOtcBlocktradesResponse(val *ListOtcBlocktradesResponse) *NullableListOtcBlocktradesResponse {
	return &NullableListOtcBlocktradesResponse{value: val, isSet: true}
}

func (v NullableListOtcBlocktradesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOtcBlocktradesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
