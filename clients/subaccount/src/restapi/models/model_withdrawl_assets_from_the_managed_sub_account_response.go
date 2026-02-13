/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the WithdrawlAssetsFromTheManagedSubAccountResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &WithdrawlAssetsFromTheManagedSubAccountResponse{}

// WithdrawlAssetsFromTheManagedSubAccountResponse struct for WithdrawlAssetsFromTheManagedSubAccountResponse
type WithdrawlAssetsFromTheManagedSubAccountResponse struct {
	TranId               *int64 `json:"tranId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WithdrawlAssetsFromTheManagedSubAccountResponse WithdrawlAssetsFromTheManagedSubAccountResponse

// NewWithdrawlAssetsFromTheManagedSubAccountResponse instantiates a new WithdrawlAssetsFromTheManagedSubAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWithdrawlAssetsFromTheManagedSubAccountResponse() *WithdrawlAssetsFromTheManagedSubAccountResponse {
	this := WithdrawlAssetsFromTheManagedSubAccountResponse{}
	return &this
}

// NewWithdrawlAssetsFromTheManagedSubAccountResponseWithDefaults instantiates a new WithdrawlAssetsFromTheManagedSubAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWithdrawlAssetsFromTheManagedSubAccountResponseWithDefaults() *WithdrawlAssetsFromTheManagedSubAccountResponse {
	this := WithdrawlAssetsFromTheManagedSubAccountResponse{}
	return &this
}

// GetTranId returns the TranId field value if set, zero value otherwise.
func (o *WithdrawlAssetsFromTheManagedSubAccountResponse) GetTranId() int64 {
	if o == nil || common.IsNil(o.TranId) {
		var ret int64
		return ret
	}
	return *o.TranId
}

// GetTranIdOk returns a tuple with the TranId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawlAssetsFromTheManagedSubAccountResponse) GetTranIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TranId) {
		return nil, false
	}
	return o.TranId, true
}

// HasTranId returns a boolean if a field has been set.
func (o *WithdrawlAssetsFromTheManagedSubAccountResponse) HasTranId() bool {
	if o != nil && !common.IsNil(o.TranId) {
		return true
	}

	return false
}

// SetTranId gets a reference to the given int64 and assigns it to the TranId field.
func (o *WithdrawlAssetsFromTheManagedSubAccountResponse) SetTranId(v int64) {
	o.TranId = &v
}

func (o WithdrawlAssetsFromTheManagedSubAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WithdrawlAssetsFromTheManagedSubAccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TranId) {
		toSerialize["tranId"] = o.TranId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WithdrawlAssetsFromTheManagedSubAccountResponse) UnmarshalJSON(data []byte) (err error) {
	varWithdrawlAssetsFromTheManagedSubAccountResponse := _WithdrawlAssetsFromTheManagedSubAccountResponse{}

	err = json.Unmarshal(data, &varWithdrawlAssetsFromTheManagedSubAccountResponse)

	if err != nil {
		return err
	}

	*o = WithdrawlAssetsFromTheManagedSubAccountResponse(varWithdrawlAssetsFromTheManagedSubAccountResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tranId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWithdrawlAssetsFromTheManagedSubAccountResponse struct {
	value *WithdrawlAssetsFromTheManagedSubAccountResponse
	isSet bool
}

func (v NullableWithdrawlAssetsFromTheManagedSubAccountResponse) Get() *WithdrawlAssetsFromTheManagedSubAccountResponse {
	return v.value
}

func (v *NullableWithdrawlAssetsFromTheManagedSubAccountResponse) Set(val *WithdrawlAssetsFromTheManagedSubAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWithdrawlAssetsFromTheManagedSubAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWithdrawlAssetsFromTheManagedSubAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWithdrawlAssetsFromTheManagedSubAccountResponse(val *WithdrawlAssetsFromTheManagedSubAccountResponse) *NullableWithdrawlAssetsFromTheManagedSubAccountResponse {
	return &NullableWithdrawlAssetsFromTheManagedSubAccountResponse{value: val, isSet: true}
}

func (v NullableWithdrawlAssetsFromTheManagedSubAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWithdrawlAssetsFromTheManagedSubAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
