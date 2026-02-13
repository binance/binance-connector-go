/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the DepositAssetsIntoTheManagedSubAccountResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DepositAssetsIntoTheManagedSubAccountResponse{}

// DepositAssetsIntoTheManagedSubAccountResponse struct for DepositAssetsIntoTheManagedSubAccountResponse
type DepositAssetsIntoTheManagedSubAccountResponse struct {
	TranId               *int64 `json:"tranId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DepositAssetsIntoTheManagedSubAccountResponse DepositAssetsIntoTheManagedSubAccountResponse

// NewDepositAssetsIntoTheManagedSubAccountResponse instantiates a new DepositAssetsIntoTheManagedSubAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepositAssetsIntoTheManagedSubAccountResponse() *DepositAssetsIntoTheManagedSubAccountResponse {
	this := DepositAssetsIntoTheManagedSubAccountResponse{}
	return &this
}

// NewDepositAssetsIntoTheManagedSubAccountResponseWithDefaults instantiates a new DepositAssetsIntoTheManagedSubAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositAssetsIntoTheManagedSubAccountResponseWithDefaults() *DepositAssetsIntoTheManagedSubAccountResponse {
	this := DepositAssetsIntoTheManagedSubAccountResponse{}
	return &this
}

// GetTranId returns the TranId field value if set, zero value otherwise.
func (o *DepositAssetsIntoTheManagedSubAccountResponse) GetTranId() int64 {
	if o == nil || common.IsNil(o.TranId) {
		var ret int64
		return ret
	}
	return *o.TranId
}

// GetTranIdOk returns a tuple with the TranId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositAssetsIntoTheManagedSubAccountResponse) GetTranIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TranId) {
		return nil, false
	}
	return o.TranId, true
}

// HasTranId returns a boolean if a field has been set.
func (o *DepositAssetsIntoTheManagedSubAccountResponse) HasTranId() bool {
	if o != nil && !common.IsNil(o.TranId) {
		return true
	}

	return false
}

// SetTranId gets a reference to the given int64 and assigns it to the TranId field.
func (o *DepositAssetsIntoTheManagedSubAccountResponse) SetTranId(v int64) {
	o.TranId = &v
}

func (o DepositAssetsIntoTheManagedSubAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DepositAssetsIntoTheManagedSubAccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TranId) {
		toSerialize["tranId"] = o.TranId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DepositAssetsIntoTheManagedSubAccountResponse) UnmarshalJSON(data []byte) (err error) {
	varDepositAssetsIntoTheManagedSubAccountResponse := _DepositAssetsIntoTheManagedSubAccountResponse{}

	err = json.Unmarshal(data, &varDepositAssetsIntoTheManagedSubAccountResponse)

	if err != nil {
		return err
	}

	*o = DepositAssetsIntoTheManagedSubAccountResponse(varDepositAssetsIntoTheManagedSubAccountResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tranId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDepositAssetsIntoTheManagedSubAccountResponse struct {
	value *DepositAssetsIntoTheManagedSubAccountResponse
	isSet bool
}

func (v NullableDepositAssetsIntoTheManagedSubAccountResponse) Get() *DepositAssetsIntoTheManagedSubAccountResponse {
	return v.value
}

func (v *NullableDepositAssetsIntoTheManagedSubAccountResponse) Set(val *DepositAssetsIntoTheManagedSubAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositAssetsIntoTheManagedSubAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositAssetsIntoTheManagedSubAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepositAssetsIntoTheManagedSubAccountResponse(val *DepositAssetsIntoTheManagedSubAccountResponse) *NullableDepositAssetsIntoTheManagedSubAccountResponse {
	return &NullableDepositAssetsIntoTheManagedSubAccountResponse{value: val, isSet: true}
}

func (v NullableDepositAssetsIntoTheManagedSubAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositAssetsIntoTheManagedSubAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
