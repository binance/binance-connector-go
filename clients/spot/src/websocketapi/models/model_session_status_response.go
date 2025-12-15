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

// checks if the SessionStatusResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SessionStatusResponse{}

// SessionStatusResponse struct for SessionStatusResponse
type SessionStatusResponse struct {
	Id                   *string                      `json:"id,omitempty"`
	Status               *int64                       `json:"status,omitempty"`
	Result               *SessionStatusResponseResult `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SessionStatusResponse SessionStatusResponse

// NewSessionStatusResponse instantiates a new SessionStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionStatusResponse() *SessionStatusResponse {
	this := SessionStatusResponse{}
	return &this
}

// NewSessionStatusResponseWithDefaults instantiates a new SessionStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionStatusResponseWithDefaults() *SessionStatusResponse {
	this := SessionStatusResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SessionStatusResponse) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionStatusResponse) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SessionStatusResponse) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SessionStatusResponse) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SessionStatusResponse) GetStatus() int64 {
	if o == nil || common.IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionStatusResponse) GetStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SessionStatusResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *SessionStatusResponse) SetStatus(v int64) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *SessionStatusResponse) GetResult() SessionStatusResponseResult {
	if o == nil || common.IsNil(o.Result) {
		var ret SessionStatusResponseResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionStatusResponse) GetResultOk() (*SessionStatusResponseResult, bool) {
	if o == nil || common.IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *SessionStatusResponse) HasResult() bool {
	if o != nil && !common.IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given SessionStatusResponseResult and assigns it to the Result field.
func (o *SessionStatusResponse) SetResult(v SessionStatusResponseResult) {
	o.Result = &v
}

func (o SessionStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SessionStatusResponse) UnmarshalJSON(data []byte) (err error) {
	varSessionStatusResponse := _SessionStatusResponse{}

	err = json.Unmarshal(data, &varSessionStatusResponse)

	if err != nil {
		return err
	}

	*o = SessionStatusResponse(varSessionStatusResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSessionStatusResponse struct {
	value *SessionStatusResponse
	isSet bool
}

func (v NullableSessionStatusResponse) Get() *SessionStatusResponse {
	return v.value
}

func (v *NullableSessionStatusResponse) Set(val *SessionStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionStatusResponse(val *SessionStatusResponse) *NullableSessionStatusResponse {
	return &NullableSessionStatusResponse{value: val, isSet: true}
}

func (v NullableSessionStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
