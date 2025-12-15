/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetIpRestrictionForASubAccountApiKeyResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetIpRestrictionForASubAccountApiKeyResponse{}

// GetIpRestrictionForASubAccountApiKeyResponse struct for GetIpRestrictionForASubAccountApiKeyResponse
type GetIpRestrictionForASubAccountApiKeyResponse struct {
	IpRestrict           *string  `json:"ipRestrict,omitempty"`
	IpList               []string `json:"ipList,omitempty"`
	UpdateTime           *int64   `json:"updateTime,omitempty"`
	ApiKey               *string  `json:"apiKey,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetIpRestrictionForASubAccountApiKeyResponse GetIpRestrictionForASubAccountApiKeyResponse

// NewGetIpRestrictionForASubAccountApiKeyResponse instantiates a new GetIpRestrictionForASubAccountApiKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIpRestrictionForASubAccountApiKeyResponse() *GetIpRestrictionForASubAccountApiKeyResponse {
	this := GetIpRestrictionForASubAccountApiKeyResponse{}
	return &this
}

// NewGetIpRestrictionForASubAccountApiKeyResponseWithDefaults instantiates a new GetIpRestrictionForASubAccountApiKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIpRestrictionForASubAccountApiKeyResponseWithDefaults() *GetIpRestrictionForASubAccountApiKeyResponse {
	this := GetIpRestrictionForASubAccountApiKeyResponse{}
	return &this
}

// GetIpRestrict returns the IpRestrict field value if set, zero value otherwise.
func (o *GetIpRestrictionForASubAccountApiKeyResponse) GetIpRestrict() string {
	if o == nil || common.IsNil(o.IpRestrict) {
		var ret string
		return ret
	}
	return *o.IpRestrict
}

// GetIpRestrictOk returns a tuple with the IpRestrict field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIpRestrictionForASubAccountApiKeyResponse) GetIpRestrictOk() (*string, bool) {
	if o == nil || common.IsNil(o.IpRestrict) {
		return nil, false
	}
	return o.IpRestrict, true
}

// HasIpRestrict returns a boolean if a field has been set.
func (o *GetIpRestrictionForASubAccountApiKeyResponse) HasIpRestrict() bool {
	if o != nil && !common.IsNil(o.IpRestrict) {
		return true
	}

	return false
}

// SetIpRestrict gets a reference to the given string and assigns it to the IpRestrict field.
func (o *GetIpRestrictionForASubAccountApiKeyResponse) SetIpRestrict(v string) {
	o.IpRestrict = &v
}

// GetIpList returns the IpList field value if set, zero value otherwise.
func (o *GetIpRestrictionForASubAccountApiKeyResponse) GetIpList() []string {
	if o == nil || common.IsNil(o.IpList) {
		var ret []string
		return ret
	}
	return o.IpList
}

// GetIpListOk returns a tuple with the IpList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIpRestrictionForASubAccountApiKeyResponse) GetIpListOk() ([]string, bool) {
	if o == nil || common.IsNil(o.IpList) {
		return nil, false
	}
	return o.IpList, true
}

// HasIpList returns a boolean if a field has been set.
func (o *GetIpRestrictionForASubAccountApiKeyResponse) HasIpList() bool {
	if o != nil && !common.IsNil(o.IpList) {
		return true
	}

	return false
}

// SetIpList gets a reference to the given []string and assigns it to the IpList field.
func (o *GetIpRestrictionForASubAccountApiKeyResponse) SetIpList(v []string) {
	o.IpList = v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *GetIpRestrictionForASubAccountApiKeyResponse) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIpRestrictionForASubAccountApiKeyResponse) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *GetIpRestrictionForASubAccountApiKeyResponse) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *GetIpRestrictionForASubAccountApiKeyResponse) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *GetIpRestrictionForASubAccountApiKeyResponse) GetApiKey() string {
	if o == nil || common.IsNil(o.ApiKey) {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIpRestrictionForASubAccountApiKeyResponse) GetApiKeyOk() (*string, bool) {
	if o == nil || common.IsNil(o.ApiKey) {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *GetIpRestrictionForASubAccountApiKeyResponse) HasApiKey() bool {
	if o != nil && !common.IsNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *GetIpRestrictionForASubAccountApiKeyResponse) SetApiKey(v string) {
	o.ApiKey = &v
}

func (o GetIpRestrictionForASubAccountApiKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetIpRestrictionForASubAccountApiKeyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.IpRestrict) {
		toSerialize["ipRestrict"] = o.IpRestrict
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

func (o *GetIpRestrictionForASubAccountApiKeyResponse) UnmarshalJSON(data []byte) (err error) {
	varGetIpRestrictionForASubAccountApiKeyResponse := _GetIpRestrictionForASubAccountApiKeyResponse{}

	err = json.Unmarshal(data, &varGetIpRestrictionForASubAccountApiKeyResponse)

	if err != nil {
		return err
	}

	*o = GetIpRestrictionForASubAccountApiKeyResponse(varGetIpRestrictionForASubAccountApiKeyResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ipRestrict")
		delete(additionalProperties, "ipList")
		delete(additionalProperties, "updateTime")
		delete(additionalProperties, "apiKey")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetIpRestrictionForASubAccountApiKeyResponse struct {
	value *GetIpRestrictionForASubAccountApiKeyResponse
	isSet bool
}

func (v NullableGetIpRestrictionForASubAccountApiKeyResponse) Get() *GetIpRestrictionForASubAccountApiKeyResponse {
	return v.value
}

func (v *NullableGetIpRestrictionForASubAccountApiKeyResponse) Set(val *GetIpRestrictionForASubAccountApiKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIpRestrictionForASubAccountApiKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIpRestrictionForASubAccountApiKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIpRestrictionForASubAccountApiKeyResponse(val *GetIpRestrictionForASubAccountApiKeyResponse) *NullableGetIpRestrictionForASubAccountApiKeyResponse {
	return &NullableGetIpRestrictionForASubAccountApiKeyResponse{value: val, isSet: true}
}

func (v NullableGetIpRestrictionForASubAccountApiKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIpRestrictionForASubAccountApiKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
