/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryEnabledIsolatedMarginAccountLimitResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryEnabledIsolatedMarginAccountLimitResponse{}

// QueryEnabledIsolatedMarginAccountLimitResponse struct for QueryEnabledIsolatedMarginAccountLimitResponse
type QueryEnabledIsolatedMarginAccountLimitResponse struct {
	EnabledAccount       *int64 `json:"enabledAccount,omitempty"`
	MaxAccount           *int64 `json:"maxAccount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryEnabledIsolatedMarginAccountLimitResponse QueryEnabledIsolatedMarginAccountLimitResponse

// NewQueryEnabledIsolatedMarginAccountLimitResponse instantiates a new QueryEnabledIsolatedMarginAccountLimitResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryEnabledIsolatedMarginAccountLimitResponse() *QueryEnabledIsolatedMarginAccountLimitResponse {
	this := QueryEnabledIsolatedMarginAccountLimitResponse{}
	return &this
}

// NewQueryEnabledIsolatedMarginAccountLimitResponseWithDefaults instantiates a new QueryEnabledIsolatedMarginAccountLimitResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryEnabledIsolatedMarginAccountLimitResponseWithDefaults() *QueryEnabledIsolatedMarginAccountLimitResponse {
	this := QueryEnabledIsolatedMarginAccountLimitResponse{}
	return &this
}

// GetEnabledAccount returns the EnabledAccount field value if set, zero value otherwise.
func (o *QueryEnabledIsolatedMarginAccountLimitResponse) GetEnabledAccount() int64 {
	if o == nil || common.IsNil(o.EnabledAccount) {
		var ret int64
		return ret
	}
	return *o.EnabledAccount
}

// GetEnabledAccountOk returns a tuple with the EnabledAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryEnabledIsolatedMarginAccountLimitResponse) GetEnabledAccountOk() (*int64, bool) {
	if o == nil || common.IsNil(o.EnabledAccount) {
		return nil, false
	}
	return o.EnabledAccount, true
}

// HasEnabledAccount returns a boolean if a field has been set.
func (o *QueryEnabledIsolatedMarginAccountLimitResponse) HasEnabledAccount() bool {
	if o != nil && !common.IsNil(o.EnabledAccount) {
		return true
	}

	return false
}

// SetEnabledAccount gets a reference to the given int64 and assigns it to the EnabledAccount field.
func (o *QueryEnabledIsolatedMarginAccountLimitResponse) SetEnabledAccount(v int64) {
	o.EnabledAccount = &v
}

// GetMaxAccount returns the MaxAccount field value if set, zero value otherwise.
func (o *QueryEnabledIsolatedMarginAccountLimitResponse) GetMaxAccount() int64 {
	if o == nil || common.IsNil(o.MaxAccount) {
		var ret int64
		return ret
	}
	return *o.MaxAccount
}

// GetMaxAccountOk returns a tuple with the MaxAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryEnabledIsolatedMarginAccountLimitResponse) GetMaxAccountOk() (*int64, bool) {
	if o == nil || common.IsNil(o.MaxAccount) {
		return nil, false
	}
	return o.MaxAccount, true
}

// HasMaxAccount returns a boolean if a field has been set.
func (o *QueryEnabledIsolatedMarginAccountLimitResponse) HasMaxAccount() bool {
	if o != nil && !common.IsNil(o.MaxAccount) {
		return true
	}

	return false
}

// SetMaxAccount gets a reference to the given int64 and assigns it to the MaxAccount field.
func (o *QueryEnabledIsolatedMarginAccountLimitResponse) SetMaxAccount(v int64) {
	o.MaxAccount = &v
}

func (o QueryEnabledIsolatedMarginAccountLimitResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryEnabledIsolatedMarginAccountLimitResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.EnabledAccount) {
		toSerialize["enabledAccount"] = o.EnabledAccount
	}
	if !common.IsNil(o.MaxAccount) {
		toSerialize["maxAccount"] = o.MaxAccount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryEnabledIsolatedMarginAccountLimitResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryEnabledIsolatedMarginAccountLimitResponse := _QueryEnabledIsolatedMarginAccountLimitResponse{}

	err = json.Unmarshal(data, &varQueryEnabledIsolatedMarginAccountLimitResponse)

	if err != nil {
		return err
	}

	*o = QueryEnabledIsolatedMarginAccountLimitResponse(varQueryEnabledIsolatedMarginAccountLimitResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabledAccount")
		delete(additionalProperties, "maxAccount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryEnabledIsolatedMarginAccountLimitResponse struct {
	value *QueryEnabledIsolatedMarginAccountLimitResponse
	isSet bool
}

func (v NullableQueryEnabledIsolatedMarginAccountLimitResponse) Get() *QueryEnabledIsolatedMarginAccountLimitResponse {
	return v.value
}

func (v *NullableQueryEnabledIsolatedMarginAccountLimitResponse) Set(val *QueryEnabledIsolatedMarginAccountLimitResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryEnabledIsolatedMarginAccountLimitResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryEnabledIsolatedMarginAccountLimitResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryEnabledIsolatedMarginAccountLimitResponse(val *QueryEnabledIsolatedMarginAccountLimitResponse) *NullableQueryEnabledIsolatedMarginAccountLimitResponse {
	return &NullableQueryEnabledIsolatedMarginAccountLimitResponse{value: val, isSet: true}
}

func (v NullableQueryEnabledIsolatedMarginAccountLimitResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryEnabledIsolatedMarginAccountLimitResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
