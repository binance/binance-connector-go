/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the FetchWithdrawQuotaResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FetchWithdrawQuotaResponse{}

// FetchWithdrawQuotaResponse struct for FetchWithdrawQuotaResponse
type FetchWithdrawQuotaResponse struct {
	WdQuota              *string `json:"wdQuota,omitempty"`
	UsedWdQuota          *string `json:"usedWdQuota,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FetchWithdrawQuotaResponse FetchWithdrawQuotaResponse

// NewFetchWithdrawQuotaResponse instantiates a new FetchWithdrawQuotaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFetchWithdrawQuotaResponse() *FetchWithdrawQuotaResponse {
	this := FetchWithdrawQuotaResponse{}
	return &this
}

// NewFetchWithdrawQuotaResponseWithDefaults instantiates a new FetchWithdrawQuotaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFetchWithdrawQuotaResponseWithDefaults() *FetchWithdrawQuotaResponse {
	this := FetchWithdrawQuotaResponse{}
	return &this
}

// GetWdQuota returns the WdQuota field value if set, zero value otherwise.
func (o *FetchWithdrawQuotaResponse) GetWdQuota() string {
	if o == nil || common.IsNil(o.WdQuota) {
		var ret string
		return ret
	}
	return *o.WdQuota
}

// GetWdQuotaOk returns a tuple with the WdQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchWithdrawQuotaResponse) GetWdQuotaOk() (*string, bool) {
	if o == nil || common.IsNil(o.WdQuota) {
		return nil, false
	}
	return o.WdQuota, true
}

// HasWdQuota returns a boolean if a field has been set.
func (o *FetchWithdrawQuotaResponse) HasWdQuota() bool {
	if o != nil && !common.IsNil(o.WdQuota) {
		return true
	}

	return false
}

// SetWdQuota gets a reference to the given string and assigns it to the WdQuota field.
func (o *FetchWithdrawQuotaResponse) SetWdQuota(v string) {
	o.WdQuota = &v
}

// GetUsedWdQuota returns the UsedWdQuota field value if set, zero value otherwise.
func (o *FetchWithdrawQuotaResponse) GetUsedWdQuota() string {
	if o == nil || common.IsNil(o.UsedWdQuota) {
		var ret string
		return ret
	}
	return *o.UsedWdQuota
}

// GetUsedWdQuotaOk returns a tuple with the UsedWdQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchWithdrawQuotaResponse) GetUsedWdQuotaOk() (*string, bool) {
	if o == nil || common.IsNil(o.UsedWdQuota) {
		return nil, false
	}
	return o.UsedWdQuota, true
}

// HasUsedWdQuota returns a boolean if a field has been set.
func (o *FetchWithdrawQuotaResponse) HasUsedWdQuota() bool {
	if o != nil && !common.IsNil(o.UsedWdQuota) {
		return true
	}

	return false
}

// SetUsedWdQuota gets a reference to the given string and assigns it to the UsedWdQuota field.
func (o *FetchWithdrawQuotaResponse) SetUsedWdQuota(v string) {
	o.UsedWdQuota = &v
}

func (o FetchWithdrawQuotaResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FetchWithdrawQuotaResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.WdQuota) {
		toSerialize["wdQuota"] = o.WdQuota
	}
	if !common.IsNil(o.UsedWdQuota) {
		toSerialize["usedWdQuota"] = o.UsedWdQuota
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FetchWithdrawQuotaResponse) UnmarshalJSON(data []byte) (err error) {
	varFetchWithdrawQuotaResponse := _FetchWithdrawQuotaResponse{}

	err = json.Unmarshal(data, &varFetchWithdrawQuotaResponse)

	if err != nil {
		return err
	}

	*o = FetchWithdrawQuotaResponse(varFetchWithdrawQuotaResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "wdQuota")
		delete(additionalProperties, "usedWdQuota")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFetchWithdrawQuotaResponse struct {
	value *FetchWithdrawQuotaResponse
	isSet bool
}

func (v NullableFetchWithdrawQuotaResponse) Get() *FetchWithdrawQuotaResponse {
	return v.value
}

func (v *NullableFetchWithdrawQuotaResponse) Set(val *FetchWithdrawQuotaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFetchWithdrawQuotaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFetchWithdrawQuotaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFetchWithdrawQuotaResponse(val *FetchWithdrawQuotaResponse) *NullableFetchWithdrawQuotaResponse {
	return &NullableFetchWithdrawQuotaResponse{value: val, isSet: true}
}

func (v NullableFetchWithdrawQuotaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFetchWithdrawQuotaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
