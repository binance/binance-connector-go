/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the EventStreamTerminated type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &EventStreamTerminated{}

// EventStreamTerminated struct for EventStreamTerminated
type EventStreamTerminated struct {
	E                    *int64 `json:"E,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EventStreamTerminated EventStreamTerminated

// NewEventStreamTerminated instantiates a new EventStreamTerminated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventStreamTerminated() *EventStreamTerminated {
	this := EventStreamTerminated{}
	return &this
}

// NewEventStreamTerminatedWithDefaults instantiates a new EventStreamTerminated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventStreamTerminatedWithDefaults() *EventStreamTerminated {
	this := EventStreamTerminated{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *EventStreamTerminated) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventStreamTerminated) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *EventStreamTerminated) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *EventStreamTerminated) SetE(v int64) {
	o.E = &v
}

func (o EventStreamTerminated) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventStreamTerminated) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EventStreamTerminated) UnmarshalJSON(data []byte) (err error) {
	varEventStreamTerminated := _EventStreamTerminated{}

	err = json.Unmarshal(data, &varEventStreamTerminated)

	if err != nil {
		return err
	}

	*o = EventStreamTerminated(varEventStreamTerminated)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "E")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEventStreamTerminated struct {
	value *EventStreamTerminated
	isSet bool
}

func (v NullableEventStreamTerminated) Get() *EventStreamTerminated {
	return v.value
}

func (v *NullableEventStreamTerminated) Set(val *EventStreamTerminated) {
	v.value = val
	v.isSet = true
}

func (v NullableEventStreamTerminated) IsSet() bool {
	return v.isSet
}

func (v *NullableEventStreamTerminated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventStreamTerminated(val *EventStreamTerminated) *NullableEventStreamTerminated {
	return &NullableEventStreamTerminated{value: val, isSet: true}
}

func (v NullableEventStreamTerminated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventStreamTerminated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
