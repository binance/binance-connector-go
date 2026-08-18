/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ApplyMmDepositResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ApplyMmDepositResponse{}

// ApplyMmDepositResponse struct for ApplyMmDepositResponse
type ApplyMmDepositResponse struct {
	TransferId           *string `json:"transferId,omitempty"`
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplyMmDepositResponse ApplyMmDepositResponse

// NewApplyMmDepositResponse instantiates a new ApplyMmDepositResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplyMmDepositResponse() *ApplyMmDepositResponse {
	this := ApplyMmDepositResponse{}
	return &this
}

// NewApplyMmDepositResponseWithDefaults instantiates a new ApplyMmDepositResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplyMmDepositResponseWithDefaults() *ApplyMmDepositResponse {
	this := ApplyMmDepositResponse{}
	return &this
}

// GetTransferId returns the TransferId field value if set, zero value otherwise.
func (o *ApplyMmDepositResponse) GetTransferId() string {
	if o == nil || common.IsNil(o.TransferId) {
		var ret string
		return ret
	}
	return *o.TransferId
}

// GetTransferIdOk returns a tuple with the TransferId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyMmDepositResponse) GetTransferIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TransferId) {
		return nil, false
	}
	return o.TransferId, true
}

// HasTransferId returns a boolean if a field has been set.
func (o *ApplyMmDepositResponse) HasTransferId() bool {
	if o != nil && !common.IsNil(o.TransferId) {
		return true
	}

	return false
}

// SetTransferId gets a reference to the given string and assigns it to the TransferId field.
func (o *ApplyMmDepositResponse) SetTransferId(v string) {
	o.TransferId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApplyMmDepositResponse) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyMmDepositResponse) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApplyMmDepositResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApplyMmDepositResponse) SetStatus(v string) {
	o.Status = &v
}

func (o ApplyMmDepositResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplyMmDepositResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TransferId) {
		toSerialize["transferId"] = o.TransferId
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApplyMmDepositResponse) UnmarshalJSON(data []byte) (err error) {
	varApplyMmDepositResponse := _ApplyMmDepositResponse{}

	err = json.Unmarshal(data, &varApplyMmDepositResponse)

	if err != nil {
		return err
	}

	*o = ApplyMmDepositResponse(varApplyMmDepositResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "transferId")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplyMmDepositResponse struct {
	value *ApplyMmDepositResponse
	isSet bool
}

func (v NullableApplyMmDepositResponse) Get() *ApplyMmDepositResponse {
	return v.value
}

func (v *NullableApplyMmDepositResponse) Set(val *ApplyMmDepositResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApplyMmDepositResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApplyMmDepositResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplyMmDepositResponse(val *ApplyMmDepositResponse) *NullableApplyMmDepositResponse {
	return &NullableApplyMmDepositResponse{value: val, isSet: true}
}

func (v NullableApplyMmDepositResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplyMmDepositResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
