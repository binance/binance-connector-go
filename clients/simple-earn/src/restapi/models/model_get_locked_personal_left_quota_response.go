/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetLockedPersonalLeftQuotaResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetLockedPersonalLeftQuotaResponse{}

// GetLockedPersonalLeftQuotaResponse struct for GetLockedPersonalLeftQuotaResponse
type GetLockedPersonalLeftQuotaResponse struct {
	LeftPersonalQuota    *string `json:"leftPersonalQuota,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetLockedPersonalLeftQuotaResponse GetLockedPersonalLeftQuotaResponse

// NewGetLockedPersonalLeftQuotaResponse instantiates a new GetLockedPersonalLeftQuotaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLockedPersonalLeftQuotaResponse() *GetLockedPersonalLeftQuotaResponse {
	this := GetLockedPersonalLeftQuotaResponse{}
	return &this
}

// NewGetLockedPersonalLeftQuotaResponseWithDefaults instantiates a new GetLockedPersonalLeftQuotaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLockedPersonalLeftQuotaResponseWithDefaults() *GetLockedPersonalLeftQuotaResponse {
	this := GetLockedPersonalLeftQuotaResponse{}
	return &this
}

// GetLeftPersonalQuota returns the LeftPersonalQuota field value if set, zero value otherwise.
func (o *GetLockedPersonalLeftQuotaResponse) GetLeftPersonalQuota() string {
	if o == nil || common.IsNil(o.LeftPersonalQuota) {
		var ret string
		return ret
	}
	return *o.LeftPersonalQuota
}

// GetLeftPersonalQuotaOk returns a tuple with the LeftPersonalQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedPersonalLeftQuotaResponse) GetLeftPersonalQuotaOk() (*string, bool) {
	if o == nil || common.IsNil(o.LeftPersonalQuota) {
		return nil, false
	}
	return o.LeftPersonalQuota, true
}

// HasLeftPersonalQuota returns a boolean if a field has been set.
func (o *GetLockedPersonalLeftQuotaResponse) HasLeftPersonalQuota() bool {
	if o != nil && !common.IsNil(o.LeftPersonalQuota) {
		return true
	}

	return false
}

// SetLeftPersonalQuota gets a reference to the given string and assigns it to the LeftPersonalQuota field.
func (o *GetLockedPersonalLeftQuotaResponse) SetLeftPersonalQuota(v string) {
	o.LeftPersonalQuota = &v
}

func (o GetLockedPersonalLeftQuotaResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLockedPersonalLeftQuotaResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.LeftPersonalQuota) {
		toSerialize["leftPersonalQuota"] = o.LeftPersonalQuota
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetLockedPersonalLeftQuotaResponse) UnmarshalJSON(data []byte) (err error) {
	varGetLockedPersonalLeftQuotaResponse := _GetLockedPersonalLeftQuotaResponse{}

	err = json.Unmarshal(data, &varGetLockedPersonalLeftQuotaResponse)

	if err != nil {
		return err
	}

	*o = GetLockedPersonalLeftQuotaResponse(varGetLockedPersonalLeftQuotaResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "leftPersonalQuota")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetLockedPersonalLeftQuotaResponse struct {
	value *GetLockedPersonalLeftQuotaResponse
	isSet bool
}

func (v NullableGetLockedPersonalLeftQuotaResponse) Get() *GetLockedPersonalLeftQuotaResponse {
	return v.value
}

func (v *NullableGetLockedPersonalLeftQuotaResponse) Set(val *GetLockedPersonalLeftQuotaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLockedPersonalLeftQuotaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLockedPersonalLeftQuotaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLockedPersonalLeftQuotaResponse(val *GetLockedPersonalLeftQuotaResponse) *NullableGetLockedPersonalLeftQuotaResponse {
	return &NullableGetLockedPersonalLeftQuotaResponse{value: val, isSet: true}
}

func (v NullableGetLockedPersonalLeftQuotaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLockedPersonalLeftQuotaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
