/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the DeleteIpListForASubAccountApiKeyResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DeleteIpListForASubAccountApiKeyResponse{}

// DeleteIpListForASubAccountApiKeyResponse struct for DeleteIpListForASubAccountApiKeyResponse
type DeleteIpListForASubAccountApiKeyResponse struct {
	IpRestrict           *string  `json:"ipRestrict,omitempty"`
	IpList               []string `json:"ipList,omitempty"`
	UpdateTime           *int64   `json:"updateTime,omitempty"`
	ApiKey               *string  `json:"apiKey,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeleteIpListForASubAccountApiKeyResponse DeleteIpListForASubAccountApiKeyResponse

// NewDeleteIpListForASubAccountApiKeyResponse instantiates a new DeleteIpListForASubAccountApiKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteIpListForASubAccountApiKeyResponse() *DeleteIpListForASubAccountApiKeyResponse {
	this := DeleteIpListForASubAccountApiKeyResponse{}
	return &this
}

// NewDeleteIpListForASubAccountApiKeyResponseWithDefaults instantiates a new DeleteIpListForASubAccountApiKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteIpListForASubAccountApiKeyResponseWithDefaults() *DeleteIpListForASubAccountApiKeyResponse {
	this := DeleteIpListForASubAccountApiKeyResponse{}
	return &this
}

// GetIpRestrict returns the IpRestrict field value if set, zero value otherwise.
func (o *DeleteIpListForASubAccountApiKeyResponse) GetIpRestrict() string {
	if o == nil || common.IsNil(o.IpRestrict) {
		var ret string
		return ret
	}
	return *o.IpRestrict
}

// GetIpRestrictOk returns a tuple with the IpRestrict field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteIpListForASubAccountApiKeyResponse) GetIpRestrictOk() (*string, bool) {
	if o == nil || common.IsNil(o.IpRestrict) {
		return nil, false
	}
	return o.IpRestrict, true
}

// HasIpRestrict returns a boolean if a field has been set.
func (o *DeleteIpListForASubAccountApiKeyResponse) HasIpRestrict() bool {
	if o != nil && !common.IsNil(o.IpRestrict) {
		return true
	}

	return false
}

// SetIpRestrict gets a reference to the given string and assigns it to the IpRestrict field.
func (o *DeleteIpListForASubAccountApiKeyResponse) SetIpRestrict(v string) {
	o.IpRestrict = &v
}

// GetIpList returns the IpList field value if set, zero value otherwise.
func (o *DeleteIpListForASubAccountApiKeyResponse) GetIpList() []string {
	if o == nil || common.IsNil(o.IpList) {
		var ret []string
		return ret
	}
	return o.IpList
}

// GetIpListOk returns a tuple with the IpList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteIpListForASubAccountApiKeyResponse) GetIpListOk() ([]string, bool) {
	if o == nil || common.IsNil(o.IpList) {
		return nil, false
	}
	return o.IpList, true
}

// HasIpList returns a boolean if a field has been set.
func (o *DeleteIpListForASubAccountApiKeyResponse) HasIpList() bool {
	if o != nil && !common.IsNil(o.IpList) {
		return true
	}

	return false
}

// SetIpList gets a reference to the given []string and assigns it to the IpList field.
func (o *DeleteIpListForASubAccountApiKeyResponse) SetIpList(v []string) {
	o.IpList = v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *DeleteIpListForASubAccountApiKeyResponse) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteIpListForASubAccountApiKeyResponse) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *DeleteIpListForASubAccountApiKeyResponse) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *DeleteIpListForASubAccountApiKeyResponse) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *DeleteIpListForASubAccountApiKeyResponse) GetApiKey() string {
	if o == nil || common.IsNil(o.ApiKey) {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteIpListForASubAccountApiKeyResponse) GetApiKeyOk() (*string, bool) {
	if o == nil || common.IsNil(o.ApiKey) {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *DeleteIpListForASubAccountApiKeyResponse) HasApiKey() bool {
	if o != nil && !common.IsNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *DeleteIpListForASubAccountApiKeyResponse) SetApiKey(v string) {
	o.ApiKey = &v
}

func (o DeleteIpListForASubAccountApiKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteIpListForASubAccountApiKeyResponse) ToMap() (map[string]interface{}, error) {
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

func (o *DeleteIpListForASubAccountApiKeyResponse) UnmarshalJSON(data []byte) (err error) {
	varDeleteIpListForASubAccountApiKeyResponse := _DeleteIpListForASubAccountApiKeyResponse{}

	err = json.Unmarshal(data, &varDeleteIpListForASubAccountApiKeyResponse)

	if err != nil {
		return err
	}

	*o = DeleteIpListForASubAccountApiKeyResponse(varDeleteIpListForASubAccountApiKeyResponse)

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

type NullableDeleteIpListForASubAccountApiKeyResponse struct {
	value *DeleteIpListForASubAccountApiKeyResponse
	isSet bool
}

func (v NullableDeleteIpListForASubAccountApiKeyResponse) Get() *DeleteIpListForASubAccountApiKeyResponse {
	return v.value
}

func (v *NullableDeleteIpListForASubAccountApiKeyResponse) Set(val *DeleteIpListForASubAccountApiKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteIpListForASubAccountApiKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteIpListForASubAccountApiKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteIpListForASubAccountApiKeyResponse(val *DeleteIpListForASubAccountApiKeyResponse) *NullableDeleteIpListForASubAccountApiKeyResponse {
	return &NullableDeleteIpListForASubAccountApiKeyResponse{value: val, isSet: true}
}

func (v NullableDeleteIpListForASubAccountApiKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteIpListForASubAccountApiKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
