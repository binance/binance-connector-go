/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AddIpRestrictionForSubAccountApiKeyResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AddIpRestrictionForSubAccountApiKeyResponse{}

// AddIpRestrictionForSubAccountApiKeyResponse struct for AddIpRestrictionForSubAccountApiKeyResponse
type AddIpRestrictionForSubAccountApiKeyResponse struct {
	Status               *string  `json:"status,omitempty"`
	IpList               []string `json:"ipList,omitempty"`
	UpdateTime           *int64   `json:"updateTime,omitempty"`
	ApiKey               *string  `json:"apiKey,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AddIpRestrictionForSubAccountApiKeyResponse AddIpRestrictionForSubAccountApiKeyResponse

// NewAddIpRestrictionForSubAccountApiKeyResponse instantiates a new AddIpRestrictionForSubAccountApiKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddIpRestrictionForSubAccountApiKeyResponse() *AddIpRestrictionForSubAccountApiKeyResponse {
	this := AddIpRestrictionForSubAccountApiKeyResponse{}
	return &this
}

// NewAddIpRestrictionForSubAccountApiKeyResponseWithDefaults instantiates a new AddIpRestrictionForSubAccountApiKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddIpRestrictionForSubAccountApiKeyResponseWithDefaults() *AddIpRestrictionForSubAccountApiKeyResponse {
	this := AddIpRestrictionForSubAccountApiKeyResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AddIpRestrictionForSubAccountApiKeyResponse) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIpRestrictionForSubAccountApiKeyResponse) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AddIpRestrictionForSubAccountApiKeyResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AddIpRestrictionForSubAccountApiKeyResponse) SetStatus(v string) {
	o.Status = &v
}

// GetIpList returns the IpList field value if set, zero value otherwise.
func (o *AddIpRestrictionForSubAccountApiKeyResponse) GetIpList() []string {
	if o == nil || common.IsNil(o.IpList) {
		var ret []string
		return ret
	}
	return o.IpList
}

// GetIpListOk returns a tuple with the IpList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIpRestrictionForSubAccountApiKeyResponse) GetIpListOk() ([]string, bool) {
	if o == nil || common.IsNil(o.IpList) {
		return nil, false
	}
	return o.IpList, true
}

// HasIpList returns a boolean if a field has been set.
func (o *AddIpRestrictionForSubAccountApiKeyResponse) HasIpList() bool {
	if o != nil && !common.IsNil(o.IpList) {
		return true
	}

	return false
}

// SetIpList gets a reference to the given []string and assigns it to the IpList field.
func (o *AddIpRestrictionForSubAccountApiKeyResponse) SetIpList(v []string) {
	o.IpList = v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *AddIpRestrictionForSubAccountApiKeyResponse) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIpRestrictionForSubAccountApiKeyResponse) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *AddIpRestrictionForSubAccountApiKeyResponse) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *AddIpRestrictionForSubAccountApiKeyResponse) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *AddIpRestrictionForSubAccountApiKeyResponse) GetApiKey() string {
	if o == nil || common.IsNil(o.ApiKey) {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIpRestrictionForSubAccountApiKeyResponse) GetApiKeyOk() (*string, bool) {
	if o == nil || common.IsNil(o.ApiKey) {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *AddIpRestrictionForSubAccountApiKeyResponse) HasApiKey() bool {
	if o != nil && !common.IsNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *AddIpRestrictionForSubAccountApiKeyResponse) SetApiKey(v string) {
	o.ApiKey = &v
}

func (o AddIpRestrictionForSubAccountApiKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddIpRestrictionForSubAccountApiKeyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.IpList) {
		toSerialize["ipList"] = o.IpList
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if !common.IsNil(o.ApiKey) {
		toSerialize["apiKey"] = o.ApiKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AddIpRestrictionForSubAccountApiKeyResponse) UnmarshalJSON(data []byte) (err error) {
	varAddIpRestrictionForSubAccountApiKeyResponse := _AddIpRestrictionForSubAccountApiKeyResponse{}

	err = json.Unmarshal(data, &varAddIpRestrictionForSubAccountApiKeyResponse)

	if err != nil {
		return err
	}

	*o = AddIpRestrictionForSubAccountApiKeyResponse(varAddIpRestrictionForSubAccountApiKeyResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "ipList")
		delete(additionalProperties, "updateTime")
		delete(additionalProperties, "apiKey")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAddIpRestrictionForSubAccountApiKeyResponse struct {
	value *AddIpRestrictionForSubAccountApiKeyResponse
	isSet bool
}

func (v NullableAddIpRestrictionForSubAccountApiKeyResponse) Get() *AddIpRestrictionForSubAccountApiKeyResponse {
	return v.value
}

func (v *NullableAddIpRestrictionForSubAccountApiKeyResponse) Set(val *AddIpRestrictionForSubAccountApiKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAddIpRestrictionForSubAccountApiKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAddIpRestrictionForSubAccountApiKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddIpRestrictionForSubAccountApiKeyResponse(val *AddIpRestrictionForSubAccountApiKeyResponse) *NullableAddIpRestrictionForSubAccountApiKeyResponse {
	return &NullableAddIpRestrictionForSubAccountApiKeyResponse{value: val, isSet: true}
}

func (v NullableAddIpRestrictionForSubAccountApiKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddIpRestrictionForSubAccountApiKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
