/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetOtcBlocktradeEventsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetOtcBlocktradeEventsResponse{}

// GetOtcBlocktradeEventsResponse struct for GetOtcBlocktradeEventsResponse
type GetOtcBlocktradeEventsResponse struct {
	Cursor               *string                                     `json:"cursor,omitempty"`
	Events               []GetOtcBlocktradeEventsResponseEventsInner `json:"events,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetOtcBlocktradeEventsResponse GetOtcBlocktradeEventsResponse

// NewGetOtcBlocktradeEventsResponse instantiates a new GetOtcBlocktradeEventsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOtcBlocktradeEventsResponse() *GetOtcBlocktradeEventsResponse {
	this := GetOtcBlocktradeEventsResponse{}
	return &this
}

// NewGetOtcBlocktradeEventsResponseWithDefaults instantiates a new GetOtcBlocktradeEventsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOtcBlocktradeEventsResponseWithDefaults() *GetOtcBlocktradeEventsResponse {
	this := GetOtcBlocktradeEventsResponse{}
	return &this
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *GetOtcBlocktradeEventsResponse) GetCursor() string {
	if o == nil || common.IsNil(o.Cursor) {
		var ret string
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOtcBlocktradeEventsResponse) GetCursorOk() (*string, bool) {
	if o == nil || common.IsNil(o.Cursor) {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *GetOtcBlocktradeEventsResponse) HasCursor() bool {
	if o != nil && !common.IsNil(o.Cursor) {
		return true
	}

	return false
}

// SetCursor gets a reference to the given string and assigns it to the Cursor field.
func (o *GetOtcBlocktradeEventsResponse) SetCursor(v string) {
	o.Cursor = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *GetOtcBlocktradeEventsResponse) GetEvents() []GetOtcBlocktradeEventsResponseEventsInner {
	if o == nil || common.IsNil(o.Events) {
		var ret []GetOtcBlocktradeEventsResponseEventsInner
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOtcBlocktradeEventsResponse) GetEventsOk() ([]GetOtcBlocktradeEventsResponseEventsInner, bool) {
	if o == nil || common.IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *GetOtcBlocktradeEventsResponse) HasEvents() bool {
	if o != nil && !common.IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []GetOtcBlocktradeEventsResponseEventsInner and assigns it to the Events field.
func (o *GetOtcBlocktradeEventsResponse) SetEvents(v []GetOtcBlocktradeEventsResponseEventsInner) {
	o.Events = v
}

func (o GetOtcBlocktradeEventsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOtcBlocktradeEventsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Cursor) {
		toSerialize["cursor"] = o.Cursor
	}
	if !common.IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetOtcBlocktradeEventsResponse) UnmarshalJSON(data []byte) (err error) {
	varGetOtcBlocktradeEventsResponse := _GetOtcBlocktradeEventsResponse{}

	err = json.Unmarshal(data, &varGetOtcBlocktradeEventsResponse)

	if err != nil {
		return err
	}

	*o = GetOtcBlocktradeEventsResponse(varGetOtcBlocktradeEventsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cursor")
		delete(additionalProperties, "events")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetOtcBlocktradeEventsResponse struct {
	value *GetOtcBlocktradeEventsResponse
	isSet bool
}

func (v NullableGetOtcBlocktradeEventsResponse) Get() *GetOtcBlocktradeEventsResponse {
	return v.value
}

func (v *NullableGetOtcBlocktradeEventsResponse) Set(val *GetOtcBlocktradeEventsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOtcBlocktradeEventsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOtcBlocktradeEventsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOtcBlocktradeEventsResponse(val *GetOtcBlocktradeEventsResponse) *NullableGetOtcBlocktradeEventsResponse {
	return &NullableGetOtcBlocktradeEventsResponse{value: val, isSet: true}
}

func (v NullableGetOtcBlocktradeEventsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOtcBlocktradeEventsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
