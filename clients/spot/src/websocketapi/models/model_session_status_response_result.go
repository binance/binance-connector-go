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

// checks if the SessionStatusResponseResult type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SessionStatusResponseResult{}

// SessionStatusResponseResult struct for SessionStatusResponseResult
type SessionStatusResponseResult struct {
	ApiKey               *string `json:"apiKey,omitempty"`
	AuthorizedSince      *int64  `json:"authorizedSince,omitempty"`
	ConnectedSince       *int64  `json:"connectedSince,omitempty"`
	ReturnRateLimits     *bool   `json:"returnRateLimits,omitempty"`
	ServerTime           *int64  `json:"serverTime,omitempty"`
	UserDataStream       *bool   `json:"userDataStream,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SessionStatusResponseResult SessionStatusResponseResult

// NewSessionStatusResponseResult instantiates a new SessionStatusResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionStatusResponseResult() *SessionStatusResponseResult {
	this := SessionStatusResponseResult{}
	return &this
}

// NewSessionStatusResponseResultWithDefaults instantiates a new SessionStatusResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionStatusResponseResultWithDefaults() *SessionStatusResponseResult {
	this := SessionStatusResponseResult{}
	return &this
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *SessionStatusResponseResult) GetApiKey() string {
	if o == nil || common.IsNil(o.ApiKey) {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionStatusResponseResult) GetApiKeyOk() (*string, bool) {
	if o == nil || common.IsNil(o.ApiKey) {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *SessionStatusResponseResult) HasApiKey() bool {
	if o != nil && !common.IsNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *SessionStatusResponseResult) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetAuthorizedSince returns the AuthorizedSince field value if set, zero value otherwise.
func (o *SessionStatusResponseResult) GetAuthorizedSince() int64 {
	if o == nil || common.IsNil(o.AuthorizedSince) {
		var ret int64
		return ret
	}
	return *o.AuthorizedSince
}

// GetAuthorizedSinceOk returns a tuple with the AuthorizedSince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionStatusResponseResult) GetAuthorizedSinceOk() (*int64, bool) {
	if o == nil || common.IsNil(o.AuthorizedSince) {
		return nil, false
	}
	return o.AuthorizedSince, true
}

// HasAuthorizedSince returns a boolean if a field has been set.
func (o *SessionStatusResponseResult) HasAuthorizedSince() bool {
	if o != nil && !common.IsNil(o.AuthorizedSince) {
		return true
	}

	return false
}

// SetAuthorizedSince gets a reference to the given int64 and assigns it to the AuthorizedSince field.
func (o *SessionStatusResponseResult) SetAuthorizedSince(v int64) {
	o.AuthorizedSince = &v
}

// GetConnectedSince returns the ConnectedSince field value if set, zero value otherwise.
func (o *SessionStatusResponseResult) GetConnectedSince() int64 {
	if o == nil || common.IsNil(o.ConnectedSince) {
		var ret int64
		return ret
	}
	return *o.ConnectedSince
}

// GetConnectedSinceOk returns a tuple with the ConnectedSince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionStatusResponseResult) GetConnectedSinceOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ConnectedSince) {
		return nil, false
	}
	return o.ConnectedSince, true
}

// HasConnectedSince returns a boolean if a field has been set.
func (o *SessionStatusResponseResult) HasConnectedSince() bool {
	if o != nil && !common.IsNil(o.ConnectedSince) {
		return true
	}

	return false
}

// SetConnectedSince gets a reference to the given int64 and assigns it to the ConnectedSince field.
func (o *SessionStatusResponseResult) SetConnectedSince(v int64) {
	o.ConnectedSince = &v
}

// GetReturnRateLimits returns the ReturnRateLimits field value if set, zero value otherwise.
func (o *SessionStatusResponseResult) GetReturnRateLimits() bool {
	if o == nil || common.IsNil(o.ReturnRateLimits) {
		var ret bool
		return ret
	}
	return *o.ReturnRateLimits
}

// GetReturnRateLimitsOk returns a tuple with the ReturnRateLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionStatusResponseResult) GetReturnRateLimitsOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ReturnRateLimits) {
		return nil, false
	}
	return o.ReturnRateLimits, true
}

// HasReturnRateLimits returns a boolean if a field has been set.
func (o *SessionStatusResponseResult) HasReturnRateLimits() bool {
	if o != nil && !common.IsNil(o.ReturnRateLimits) {
		return true
	}

	return false
}

// SetReturnRateLimits gets a reference to the given bool and assigns it to the ReturnRateLimits field.
func (o *SessionStatusResponseResult) SetReturnRateLimits(v bool) {
	o.ReturnRateLimits = &v
}

// GetServerTime returns the ServerTime field value if set, zero value otherwise.
func (o *SessionStatusResponseResult) GetServerTime() int64 {
	if o == nil || common.IsNil(o.ServerTime) {
		var ret int64
		return ret
	}
	return *o.ServerTime
}

// GetServerTimeOk returns a tuple with the ServerTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionStatusResponseResult) GetServerTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ServerTime) {
		return nil, false
	}
	return o.ServerTime, true
}

// HasServerTime returns a boolean if a field has been set.
func (o *SessionStatusResponseResult) HasServerTime() bool {
	if o != nil && !common.IsNil(o.ServerTime) {
		return true
	}

	return false
}

// SetServerTime gets a reference to the given int64 and assigns it to the ServerTime field.
func (o *SessionStatusResponseResult) SetServerTime(v int64) {
	o.ServerTime = &v
}

// GetUserDataStream returns the UserDataStream field value if set, zero value otherwise.
func (o *SessionStatusResponseResult) GetUserDataStream() bool {
	if o == nil || common.IsNil(o.UserDataStream) {
		var ret bool
		return ret
	}
	return *o.UserDataStream
}

// GetUserDataStreamOk returns a tuple with the UserDataStream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionStatusResponseResult) GetUserDataStreamOk() (*bool, bool) {
	if o == nil || common.IsNil(o.UserDataStream) {
		return nil, false
	}
	return o.UserDataStream, true
}

// HasUserDataStream returns a boolean if a field has been set.
func (o *SessionStatusResponseResult) HasUserDataStream() bool {
	if o != nil && !common.IsNil(o.UserDataStream) {
		return true
	}

	return false
}

// SetUserDataStream gets a reference to the given bool and assigns it to the UserDataStream field.
func (o *SessionStatusResponseResult) SetUserDataStream(v bool) {
	o.UserDataStream = &v
}

func (o SessionStatusResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionStatusResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ApiKey) {
		toSerialize["apiKey"] = o.ApiKey
	}
	if !common.IsNil(o.AuthorizedSince) {
		toSerialize["authorizedSince"] = o.AuthorizedSince
	}
	if !common.IsNil(o.ConnectedSince) {
		toSerialize["connectedSince"] = o.ConnectedSince
	}
	if !common.IsNil(o.ReturnRateLimits) {
		toSerialize["returnRateLimits"] = o.ReturnRateLimits
	}
	if !common.IsNil(o.ServerTime) {
		toSerialize["serverTime"] = o.ServerTime
	}
	if !common.IsNil(o.UserDataStream) {
		toSerialize["userDataStream"] = o.UserDataStream
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SessionStatusResponseResult) UnmarshalJSON(data []byte) (err error) {
	varSessionStatusResponseResult := _SessionStatusResponseResult{}

	err = json.Unmarshal(data, &varSessionStatusResponseResult)

	if err != nil {
		return err
	}

	*o = SessionStatusResponseResult(varSessionStatusResponseResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "apiKey")
		delete(additionalProperties, "authorizedSince")
		delete(additionalProperties, "connectedSince")
		delete(additionalProperties, "returnRateLimits")
		delete(additionalProperties, "serverTime")
		delete(additionalProperties, "userDataStream")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSessionStatusResponseResult struct {
	value *SessionStatusResponseResult
	isSet bool
}

func (v NullableSessionStatusResponseResult) Get() *SessionStatusResponseResult {
	return v.value
}

func (v *NullableSessionStatusResponseResult) Set(val *SessionStatusResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionStatusResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionStatusResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionStatusResponseResult(val *SessionStatusResponseResult) *NullableSessionStatusResponseResult {
	return &NullableSessionStatusResponseResult{value: val, isSet: true}
}

func (v NullableSessionStatusResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionStatusResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
