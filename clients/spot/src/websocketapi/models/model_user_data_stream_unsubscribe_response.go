/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the UserDataStreamUnsubscribeResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UserDataStreamUnsubscribeResponse{}

// UserDataStreamUnsubscribeResponse struct for UserDataStreamUnsubscribeResponse
type UserDataStreamUnsubscribeResponse struct {
	Id                   *string       `json:"id,omitempty"`
	Status               *int64        `json:"status,omitempty"`
	Result               []interface{} `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserDataStreamUnsubscribeResponse UserDataStreamUnsubscribeResponse

// NewUserDataStreamUnsubscribeResponse instantiates a new UserDataStreamUnsubscribeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserDataStreamUnsubscribeResponse() *UserDataStreamUnsubscribeResponse {
	this := UserDataStreamUnsubscribeResponse{}
	return &this
}

// NewUserDataStreamUnsubscribeResponseWithDefaults instantiates a new UserDataStreamUnsubscribeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDataStreamUnsubscribeResponseWithDefaults() *UserDataStreamUnsubscribeResponse {
	this := UserDataStreamUnsubscribeResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserDataStreamUnsubscribeResponse) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDataStreamUnsubscribeResponse) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserDataStreamUnsubscribeResponse) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserDataStreamUnsubscribeResponse) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UserDataStreamUnsubscribeResponse) GetStatus() int64 {
	if o == nil || common.IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDataStreamUnsubscribeResponse) GetStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UserDataStreamUnsubscribeResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *UserDataStreamUnsubscribeResponse) SetStatus(v int64) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UserDataStreamUnsubscribeResponse) GetResult() interface{} {
	return nil
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDataStreamUnsubscribeResponse) GetResultOk() (interface{}, bool) {
	return nil, true
}

// HasResult returns a boolean if a field has been set.
func (o *UserDataStreamUnsubscribeResponse) HasResult() bool {
	if o != nil && !common.IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given map[string]interface{} and assigns it to the Result field.
func (o *UserDataStreamUnsubscribeResponse) SetResult(v []interface{}) {
	o.Result = v
}

func (o UserDataStreamUnsubscribeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserDataStreamUnsubscribeResponse) ToMap() (map[string]interface{}, error) {
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

func (o *UserDataStreamUnsubscribeResponse) UnmarshalJSON(data []byte) (err error) {
	varUserDataStreamUnsubscribeResponse := _UserDataStreamUnsubscribeResponse{}

	err = json.Unmarshal(data, &varUserDataStreamUnsubscribeResponse)

	if err != nil {
		return err
	}

	*o = UserDataStreamUnsubscribeResponse(varUserDataStreamUnsubscribeResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserDataStreamUnsubscribeResponse struct {
	value *UserDataStreamUnsubscribeResponse
	isSet bool
}

func (v NullableUserDataStreamUnsubscribeResponse) Get() *UserDataStreamUnsubscribeResponse {
	return v.value
}

func (v *NullableUserDataStreamUnsubscribeResponse) Set(val *UserDataStreamUnsubscribeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDataStreamUnsubscribeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserDataStreamUnsubscribeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDataStreamUnsubscribeResponse(val *UserDataStreamUnsubscribeResponse) *NullableUserDataStreamUnsubscribeResponse {
	return &NullableUserDataStreamUnsubscribeResponse{value: val, isSet: true}
}

func (v NullableUserDataStreamUnsubscribeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDataStreamUnsubscribeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
