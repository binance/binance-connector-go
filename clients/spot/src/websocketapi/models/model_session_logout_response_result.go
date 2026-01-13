/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the SessionLogoutResponseResult type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SessionLogoutResponseResult{}

// SessionLogoutResponseResult struct for SessionLogoutResponseResult
type SessionLogoutResponseResult struct {
	ApiKey               *string `json:"apiKey,omitempty"`
	AuthorizedSince      *int64  `json:"authorizedSince,omitempty"`
	ConnectedSince       *int64  `json:"connectedSince,omitempty"`
	ReturnRateLimits     *bool   `json:"returnRateLimits,omitempty"`
	ServerTime           *int64  `json:"serverTime,omitempty"`
	UserDataStream       *bool   `json:"userDataStream,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SessionLogoutResponseResult SessionLogoutResponseResult

// NewSessionLogoutResponseResult instantiates a new SessionLogoutResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionLogoutResponseResult() *SessionLogoutResponseResult {
	this := SessionLogoutResponseResult{}
	return &this
}

// NewSessionLogoutResponseResultWithDefaults instantiates a new SessionLogoutResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionLogoutResponseResultWithDefaults() *SessionLogoutResponseResult {
	this := SessionLogoutResponseResult{}
	return &this
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *SessionLogoutResponseResult) GetApiKey() string {
	if o == nil || common.IsNil(o.ApiKey) {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionLogoutResponseResult) GetApiKeyOk() (*string, bool) {
	if o == nil || common.IsNil(o.ApiKey) {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *SessionLogoutResponseResult) HasApiKey() bool {
	if o != nil && !common.IsNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *SessionLogoutResponseResult) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetAuthorizedSince returns the AuthorizedSince field value if set, zero value otherwise.
func (o *SessionLogoutResponseResult) GetAuthorizedSince() int64 {
	if o == nil || common.IsNil(o.AuthorizedSince) {
		var ret int64
		return ret
	}
	return *o.AuthorizedSince
}

// GetAuthorizedSinceOk returns a tuple with the AuthorizedSince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionLogoutResponseResult) GetAuthorizedSinceOk() (*int64, bool) {
	if o == nil || common.IsNil(o.AuthorizedSince) {
		return nil, false
	}
	return o.AuthorizedSince, true
}

// HasAuthorizedSince returns a boolean if a field has been set.
func (o *SessionLogoutResponseResult) HasAuthorizedSince() bool {
	if o != nil && !common.IsNil(o.AuthorizedSince) {
		return true
	}

	return false
}

// SetAuthorizedSince gets a reference to the given int64 and assigns it to the AuthorizedSince field.
func (o *SessionLogoutResponseResult) SetAuthorizedSince(v int64) {
	o.AuthorizedSince = &v
}

// GetConnectedSince returns the ConnectedSince field value if set, zero value otherwise.
func (o *SessionLogoutResponseResult) GetConnectedSince() int64 {
	if o == nil || common.IsNil(o.ConnectedSince) {
		var ret int64
		return ret
	}
	return *o.ConnectedSince
}

// GetConnectedSinceOk returns a tuple with the ConnectedSince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionLogoutResponseResult) GetConnectedSinceOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ConnectedSince) {
		return nil, false
	}
	return o.ConnectedSince, true
}

// HasConnectedSince returns a boolean if a field has been set.
func (o *SessionLogoutResponseResult) HasConnectedSince() bool {
	if o != nil && !common.IsNil(o.ConnectedSince) {
		return true
	}

	return false
}

// SetConnectedSince gets a reference to the given int64 and assigns it to the ConnectedSince field.
func (o *SessionLogoutResponseResult) SetConnectedSince(v int64) {
	o.ConnectedSince = &v
}

// GetReturnRateLimits returns the ReturnRateLimits field value if set, zero value otherwise.
func (o *SessionLogoutResponseResult) GetReturnRateLimits() bool {
	if o == nil || common.IsNil(o.ReturnRateLimits) {
		var ret bool
		return ret
	}
	return *o.ReturnRateLimits
}

// GetReturnRateLimitsOk returns a tuple with the ReturnRateLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionLogoutResponseResult) GetReturnRateLimitsOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ReturnRateLimits) {
		return nil, false
	}
	return o.ReturnRateLimits, true
}

// HasReturnRateLimits returns a boolean if a field has been set.
func (o *SessionLogoutResponseResult) HasReturnRateLimits() bool {
	if o != nil && !common.IsNil(o.ReturnRateLimits) {
		return true
	}

	return false
}

// SetReturnRateLimits gets a reference to the given bool and assigns it to the ReturnRateLimits field.
func (o *SessionLogoutResponseResult) SetReturnRateLimits(v bool) {
	o.ReturnRateLimits = &v
}

// GetServerTime returns the ServerTime field value if set, zero value otherwise.
func (o *SessionLogoutResponseResult) GetServerTime() int64 {
	if o == nil || common.IsNil(o.ServerTime) {
		var ret int64
		return ret
	}
	return *o.ServerTime
}

// GetServerTimeOk returns a tuple with the ServerTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionLogoutResponseResult) GetServerTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ServerTime) {
		return nil, false
	}
	return o.ServerTime, true
}

// HasServerTime returns a boolean if a field has been set.
func (o *SessionLogoutResponseResult) HasServerTime() bool {
	if o != nil && !common.IsNil(o.ServerTime) {
		return true
	}

	return false
}

// SetServerTime gets a reference to the given int64 and assigns it to the ServerTime field.
func (o *SessionLogoutResponseResult) SetServerTime(v int64) {
	o.ServerTime = &v
}

// GetUserDataStream returns the UserDataStream field value if set, zero value otherwise.
func (o *SessionLogoutResponseResult) GetUserDataStream() bool {
	if o == nil || common.IsNil(o.UserDataStream) {
		var ret bool
		return ret
	}
	return *o.UserDataStream
}

// GetUserDataStreamOk returns a tuple with the UserDataStream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionLogoutResponseResult) GetUserDataStreamOk() (*bool, bool) {
	if o == nil || common.IsNil(o.UserDataStream) {
		return nil, false
	}
	return o.UserDataStream, true
}

// HasUserDataStream returns a boolean if a field has been set.
func (o *SessionLogoutResponseResult) HasUserDataStream() bool {
	if o != nil && !common.IsNil(o.UserDataStream) {
		return true
	}

	return false
}

// SetUserDataStream gets a reference to the given bool and assigns it to the UserDataStream field.
func (o *SessionLogoutResponseResult) SetUserDataStream(v bool) {
	o.UserDataStream = &v
}

func (o SessionLogoutResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionLogoutResponseResult) ToMap() (map[string]interface{}, error) {
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

func (o *SessionLogoutResponseResult) UnmarshalJSON(data []byte) (err error) {
	varSessionLogoutResponseResult := _SessionLogoutResponseResult{}

	err = json.Unmarshal(data, &varSessionLogoutResponseResult)

	if err != nil {
		return err
	}

	*o = SessionLogoutResponseResult(varSessionLogoutResponseResult)

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

type NullableSessionLogoutResponseResult struct {
	value *SessionLogoutResponseResult
	isSet bool
}

func (v NullableSessionLogoutResponseResult) Get() *SessionLogoutResponseResult {
	return v.value
}

func (v *NullableSessionLogoutResponseResult) Set(val *SessionLogoutResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionLogoutResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionLogoutResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionLogoutResponseResult(val *SessionLogoutResponseResult) *NullableSessionLogoutResponseResult {
	return &NullableSessionLogoutResponseResult{value: val, isSet: true}
}

func (v NullableSessionLogoutResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionLogoutResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
