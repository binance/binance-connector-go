/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetOnChainYieldsLockedPersonalLeftQuotaResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetOnChainYieldsLockedPersonalLeftQuotaResponse{}

// GetOnChainYieldsLockedPersonalLeftQuotaResponse struct for GetOnChainYieldsLockedPersonalLeftQuotaResponse
type GetOnChainYieldsLockedPersonalLeftQuotaResponse struct {
	LeftPersonalQuota    *string `json:"leftPersonalQuota,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetOnChainYieldsLockedPersonalLeftQuotaResponse GetOnChainYieldsLockedPersonalLeftQuotaResponse

// NewGetOnChainYieldsLockedPersonalLeftQuotaResponse instantiates a new GetOnChainYieldsLockedPersonalLeftQuotaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOnChainYieldsLockedPersonalLeftQuotaResponse() *GetOnChainYieldsLockedPersonalLeftQuotaResponse {
	this := GetOnChainYieldsLockedPersonalLeftQuotaResponse{}
	return &this
}

// NewGetOnChainYieldsLockedPersonalLeftQuotaResponseWithDefaults instantiates a new GetOnChainYieldsLockedPersonalLeftQuotaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOnChainYieldsLockedPersonalLeftQuotaResponseWithDefaults() *GetOnChainYieldsLockedPersonalLeftQuotaResponse {
	this := GetOnChainYieldsLockedPersonalLeftQuotaResponse{}
	return &this
}

// GetLeftPersonalQuota returns the LeftPersonalQuota field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedPersonalLeftQuotaResponse) GetLeftPersonalQuota() string {
	if o == nil || common.IsNil(o.LeftPersonalQuota) {
		var ret string
		return ret
	}
	return *o.LeftPersonalQuota
}

// GetLeftPersonalQuotaOk returns a tuple with the LeftPersonalQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedPersonalLeftQuotaResponse) GetLeftPersonalQuotaOk() (*string, bool) {
	if o == nil || common.IsNil(o.LeftPersonalQuota) {
		return nil, false
	}
	return o.LeftPersonalQuota, true
}

// HasLeftPersonalQuota returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedPersonalLeftQuotaResponse) HasLeftPersonalQuota() bool {
	if o != nil && !common.IsNil(o.LeftPersonalQuota) {
		return true
	}

	return false
}

// SetLeftPersonalQuota gets a reference to the given string and assigns it to the LeftPersonalQuota field.
func (o *GetOnChainYieldsLockedPersonalLeftQuotaResponse) SetLeftPersonalQuota(v string) {
	o.LeftPersonalQuota = &v
}

func (o GetOnChainYieldsLockedPersonalLeftQuotaResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOnChainYieldsLockedPersonalLeftQuotaResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.LeftPersonalQuota) {
		toSerialize["leftPersonalQuota"] = o.LeftPersonalQuota
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetOnChainYieldsLockedPersonalLeftQuotaResponse) UnmarshalJSON(data []byte) (err error) {
	varGetOnChainYieldsLockedPersonalLeftQuotaResponse := _GetOnChainYieldsLockedPersonalLeftQuotaResponse{}

	err = json.Unmarshal(data, &varGetOnChainYieldsLockedPersonalLeftQuotaResponse)

	if err != nil {
		return err
	}

	*o = GetOnChainYieldsLockedPersonalLeftQuotaResponse(varGetOnChainYieldsLockedPersonalLeftQuotaResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "leftPersonalQuota")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetOnChainYieldsLockedPersonalLeftQuotaResponse struct {
	value *GetOnChainYieldsLockedPersonalLeftQuotaResponse
	isSet bool
}

func (v NullableGetOnChainYieldsLockedPersonalLeftQuotaResponse) Get() *GetOnChainYieldsLockedPersonalLeftQuotaResponse {
	return v.value
}

func (v *NullableGetOnChainYieldsLockedPersonalLeftQuotaResponse) Set(val *GetOnChainYieldsLockedPersonalLeftQuotaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOnChainYieldsLockedPersonalLeftQuotaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOnChainYieldsLockedPersonalLeftQuotaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOnChainYieldsLockedPersonalLeftQuotaResponse(val *GetOnChainYieldsLockedPersonalLeftQuotaResponse) *NullableGetOnChainYieldsLockedPersonalLeftQuotaResponse {
	return &NullableGetOnChainYieldsLockedPersonalLeftQuotaResponse{value: val, isSet: true}
}

func (v NullableGetOnChainYieldsLockedPersonalLeftQuotaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOnChainYieldsLockedPersonalLeftQuotaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
