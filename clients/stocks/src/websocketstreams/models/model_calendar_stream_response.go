/*
Stocks Trading WebSocket Streams

WebSocket stream definitions for Binance Stocks Trading. Base URL: wss://nbstream.binance.com/equity
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the CalendarStreamResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CalendarStreamResponse{}

// CalendarStreamResponse struct for CalendarStreamResponse
type CalendarStreamResponse struct {
	// Event type, always `\"calendar\"`.
	E *string `json:"e,omitempty"`
	// Previous market phase.
	From *string `json:"from,omitempty"`
	// New market phase.
	To *string `json:"to,omitempty"`
	// Event timestamp (epoch milliseconds UTC) when the transition was detected.
	Ts                   *int64 `json:"ts,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CalendarStreamResponse CalendarStreamResponse

// NewCalendarStreamResponse instantiates a new CalendarStreamResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCalendarStreamResponse() *CalendarStreamResponse {
	this := CalendarStreamResponse{}
	return &this
}

// NewCalendarStreamResponseWithDefaults instantiates a new CalendarStreamResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCalendarStreamResponseWithDefaults() *CalendarStreamResponse {
	this := CalendarStreamResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *CalendarStreamResponse) GetE() string {
	if o == nil || common.IsNil(o.E) {
		var ret string
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarStreamResponse) GetEOk() (*string, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *CalendarStreamResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *CalendarStreamResponse) SetE(v string) {
	o.E = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *CalendarStreamResponse) GetFrom() string {
	if o == nil || common.IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarStreamResponse) GetFromOk() (*string, bool) {
	if o == nil || common.IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *CalendarStreamResponse) HasFrom() bool {
	if o != nil && !common.IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *CalendarStreamResponse) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *CalendarStreamResponse) GetTo() string {
	if o == nil || common.IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarStreamResponse) GetToOk() (*string, bool) {
	if o == nil || common.IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *CalendarStreamResponse) HasTo() bool {
	if o != nil && !common.IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *CalendarStreamResponse) SetTo(v string) {
	o.To = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *CalendarStreamResponse) GetTs() int64 {
	if o == nil || common.IsNil(o.Ts) {
		var ret int64
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarStreamResponse) GetTsOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *CalendarStreamResponse) HasTs() bool {
	if o != nil && !common.IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given int64 and assigns it to the Ts field.
func (o *CalendarStreamResponse) SetTs(v int64) {
	o.Ts = &v
}

func (o CalendarStreamResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CalendarStreamResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.E) {
		toSerialize["e"] = o.E
	}
	if !common.IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !common.IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	if !common.IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CalendarStreamResponse) UnmarshalJSON(data []byte) (err error) {
	varCalendarStreamResponse := _CalendarStreamResponse{}

	err = json.Unmarshal(data, &varCalendarStreamResponse)

	if err != nil {
		return err
	}

	*o = CalendarStreamResponse(varCalendarStreamResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "from")
		delete(additionalProperties, "to")
		delete(additionalProperties, "ts")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCalendarStreamResponse struct {
	value *CalendarStreamResponse
	isSet bool
}

func (v NullableCalendarStreamResponse) Get() *CalendarStreamResponse {
	return v.value
}

func (v *NullableCalendarStreamResponse) Set(val *CalendarStreamResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCalendarStreamResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCalendarStreamResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCalendarStreamResponse(val *CalendarStreamResponse) *NullableCalendarStreamResponse {
	return &NullableCalendarStreamResponse{value: val, isSet: true}
}

func (v NullableCalendarStreamResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCalendarStreamResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
