/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetFlexiblePersonalLeftQuotaResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFlexiblePersonalLeftQuotaResponse{}

// GetFlexiblePersonalLeftQuotaResponse struct for GetFlexiblePersonalLeftQuotaResponse
type GetFlexiblePersonalLeftQuotaResponse struct {
	LeftPersonalQuota    *string `json:"leftPersonalQuota,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetFlexiblePersonalLeftQuotaResponse GetFlexiblePersonalLeftQuotaResponse

// NewGetFlexiblePersonalLeftQuotaResponse instantiates a new GetFlexiblePersonalLeftQuotaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFlexiblePersonalLeftQuotaResponse() *GetFlexiblePersonalLeftQuotaResponse {
	this := GetFlexiblePersonalLeftQuotaResponse{}
	return &this
}

// NewGetFlexiblePersonalLeftQuotaResponseWithDefaults instantiates a new GetFlexiblePersonalLeftQuotaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFlexiblePersonalLeftQuotaResponseWithDefaults() *GetFlexiblePersonalLeftQuotaResponse {
	this := GetFlexiblePersonalLeftQuotaResponse{}
	return &this
}

// GetLeftPersonalQuota returns the LeftPersonalQuota field value if set, zero value otherwise.
func (o *GetFlexiblePersonalLeftQuotaResponse) GetLeftPersonalQuota() string {
	if o == nil || common.IsNil(o.LeftPersonalQuota) {
		var ret string
		return ret
	}
	return *o.LeftPersonalQuota
}

// GetLeftPersonalQuotaOk returns a tuple with the LeftPersonalQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexiblePersonalLeftQuotaResponse) GetLeftPersonalQuotaOk() (*string, bool) {
	if o == nil || common.IsNil(o.LeftPersonalQuota) {
		return nil, false
	}
	return o.LeftPersonalQuota, true
}

// HasLeftPersonalQuota returns a boolean if a field has been set.
func (o *GetFlexiblePersonalLeftQuotaResponse) HasLeftPersonalQuota() bool {
	if o != nil && !common.IsNil(o.LeftPersonalQuota) {
		return true
	}

	return false
}

// SetLeftPersonalQuota gets a reference to the given string and assigns it to the LeftPersonalQuota field.
func (o *GetFlexiblePersonalLeftQuotaResponse) SetLeftPersonalQuota(v string) {
	o.LeftPersonalQuota = &v
}

func (o GetFlexiblePersonalLeftQuotaResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFlexiblePersonalLeftQuotaResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.LeftPersonalQuota) {
		toSerialize["leftPersonalQuota"] = o.LeftPersonalQuota
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetFlexiblePersonalLeftQuotaResponse) UnmarshalJSON(data []byte) (err error) {
	varGetFlexiblePersonalLeftQuotaResponse := _GetFlexiblePersonalLeftQuotaResponse{}

	err = json.Unmarshal(data, &varGetFlexiblePersonalLeftQuotaResponse)

	if err != nil {
		return err
	}

	*o = GetFlexiblePersonalLeftQuotaResponse(varGetFlexiblePersonalLeftQuotaResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "leftPersonalQuota")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFlexiblePersonalLeftQuotaResponse struct {
	value *GetFlexiblePersonalLeftQuotaResponse
	isSet bool
}

func (v NullableGetFlexiblePersonalLeftQuotaResponse) Get() *GetFlexiblePersonalLeftQuotaResponse {
	return v.value
}

func (v *NullableGetFlexiblePersonalLeftQuotaResponse) Set(val *GetFlexiblePersonalLeftQuotaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFlexiblePersonalLeftQuotaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFlexiblePersonalLeftQuotaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFlexiblePersonalLeftQuotaResponse(val *GetFlexiblePersonalLeftQuotaResponse) *NullableGetFlexiblePersonalLeftQuotaResponse {
	return &NullableGetFlexiblePersonalLeftQuotaResponse{value: val, isSet: true}
}

func (v NullableGetFlexiblePersonalLeftQuotaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFlexiblePersonalLeftQuotaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
