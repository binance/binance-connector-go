/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QuerySpecialKeyListResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QuerySpecialKeyListResponseInner{}

// QuerySpecialKeyListResponseInner struct for QuerySpecialKeyListResponseInner
type QuerySpecialKeyListResponseInner struct {
	ApiName              *string `json:"apiName,omitempty"`
	ApiKey               *string `json:"apiKey,omitempty"`
	Ip                   *string `json:"ip,omitempty"`
	Type                 *string `json:"type,omitempty"`
	PermissionMode       *string `json:"permissionMode,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuerySpecialKeyListResponseInner QuerySpecialKeyListResponseInner

// NewQuerySpecialKeyListResponseInner instantiates a new QuerySpecialKeyListResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuerySpecialKeyListResponseInner() *QuerySpecialKeyListResponseInner {
	this := QuerySpecialKeyListResponseInner{}
	return &this
}

// NewQuerySpecialKeyListResponseInnerWithDefaults instantiates a new QuerySpecialKeyListResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuerySpecialKeyListResponseInnerWithDefaults() *QuerySpecialKeyListResponseInner {
	this := QuerySpecialKeyListResponseInner{}
	return &this
}

// GetApiName returns the ApiName field value if set, zero value otherwise.
func (o *QuerySpecialKeyListResponseInner) GetApiName() string {
	if o == nil || common.IsNil(o.ApiName) {
		var ret string
		return ret
	}
	return *o.ApiName
}

// GetApiNameOk returns a tuple with the ApiName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySpecialKeyListResponseInner) GetApiNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.ApiName) {
		return nil, false
	}
	return o.ApiName, true
}

// HasApiName returns a boolean if a field has been set.
func (o *QuerySpecialKeyListResponseInner) HasApiName() bool {
	if o != nil && !common.IsNil(o.ApiName) {
		return true
	}

	return false
}

// SetApiName gets a reference to the given string and assigns it to the ApiName field.
func (o *QuerySpecialKeyListResponseInner) SetApiName(v string) {
	o.ApiName = &v
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *QuerySpecialKeyListResponseInner) GetApiKey() string {
	if o == nil || common.IsNil(o.ApiKey) {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySpecialKeyListResponseInner) GetApiKeyOk() (*string, bool) {
	if o == nil || common.IsNil(o.ApiKey) {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *QuerySpecialKeyListResponseInner) HasApiKey() bool {
	if o != nil && !common.IsNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *QuerySpecialKeyListResponseInner) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *QuerySpecialKeyListResponseInner) GetIp() string {
	if o == nil || common.IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySpecialKeyListResponseInner) GetIpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *QuerySpecialKeyListResponseInner) HasIp() bool {
	if o != nil && !common.IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *QuerySpecialKeyListResponseInner) SetIp(v string) {
	o.Ip = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *QuerySpecialKeyListResponseInner) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySpecialKeyListResponseInner) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *QuerySpecialKeyListResponseInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *QuerySpecialKeyListResponseInner) SetType(v string) {
	o.Type = &v
}

// GetPermissionMode returns the PermissionMode field value if set, zero value otherwise.
func (o *QuerySpecialKeyListResponseInner) GetPermissionMode() string {
	if o == nil || common.IsNil(o.PermissionMode) {
		var ret string
		return ret
	}
	return *o.PermissionMode
}

// GetPermissionModeOk returns a tuple with the PermissionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySpecialKeyListResponseInner) GetPermissionModeOk() (*string, bool) {
	if o == nil || common.IsNil(o.PermissionMode) {
		return nil, false
	}
	return o.PermissionMode, true
}

// HasPermissionMode returns a boolean if a field has been set.
func (o *QuerySpecialKeyListResponseInner) HasPermissionMode() bool {
	if o != nil && !common.IsNil(o.PermissionMode) {
		return true
	}

	return false
}

// SetPermissionMode gets a reference to the given string and assigns it to the PermissionMode field.
func (o *QuerySpecialKeyListResponseInner) SetPermissionMode(v string) {
	o.PermissionMode = &v
}

func (o QuerySpecialKeyListResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuerySpecialKeyListResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ApiName) {
		toSerialize["apiName"] = o.ApiName
	}
	if !common.IsNil(o.ApiKey) {
		toSerialize["apiKey"] = o.ApiKey
	}
	if !common.IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.PermissionMode) {
		toSerialize["permissionMode"] = o.PermissionMode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuerySpecialKeyListResponseInner) UnmarshalJSON(data []byte) (err error) {
	varQuerySpecialKeyListResponseInner := _QuerySpecialKeyListResponseInner{}

	err = json.Unmarshal(data, &varQuerySpecialKeyListResponseInner)

	if err != nil {
		return err
	}

	*o = QuerySpecialKeyListResponseInner(varQuerySpecialKeyListResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "apiName")
		delete(additionalProperties, "apiKey")
		delete(additionalProperties, "ip")
		delete(additionalProperties, "type")
		delete(additionalProperties, "permissionMode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuerySpecialKeyListResponseInner struct {
	value *QuerySpecialKeyListResponseInner
	isSet bool
}

func (v NullableQuerySpecialKeyListResponseInner) Get() *QuerySpecialKeyListResponseInner {
	return v.value
}

func (v *NullableQuerySpecialKeyListResponseInner) Set(val *QuerySpecialKeyListResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQuerySpecialKeyListResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQuerySpecialKeyListResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuerySpecialKeyListResponseInner(val *QuerySpecialKeyListResponseInner) *NullableQuerySpecialKeyListResponseInner {
	return &NullableQuerySpecialKeyListResponseInner{value: val, isSet: true}
}

func (v NullableQuerySpecialKeyListResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuerySpecialKeyListResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
