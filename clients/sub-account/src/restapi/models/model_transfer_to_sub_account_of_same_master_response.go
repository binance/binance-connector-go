/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the TransferToSubAccountOfSameMasterResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TransferToSubAccountOfSameMasterResponse{}

// TransferToSubAccountOfSameMasterResponse struct for TransferToSubAccountOfSameMasterResponse
type TransferToSubAccountOfSameMasterResponse struct {
	TxnId                *string `json:"txnId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TransferToSubAccountOfSameMasterResponse TransferToSubAccountOfSameMasterResponse

// NewTransferToSubAccountOfSameMasterResponse instantiates a new TransferToSubAccountOfSameMasterResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferToSubAccountOfSameMasterResponse() *TransferToSubAccountOfSameMasterResponse {
	this := TransferToSubAccountOfSameMasterResponse{}
	return &this
}

// NewTransferToSubAccountOfSameMasterResponseWithDefaults instantiates a new TransferToSubAccountOfSameMasterResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferToSubAccountOfSameMasterResponseWithDefaults() *TransferToSubAccountOfSameMasterResponse {
	this := TransferToSubAccountOfSameMasterResponse{}
	return &this
}

// GetTxnId returns the TxnId field value if set, zero value otherwise.
func (o *TransferToSubAccountOfSameMasterResponse) GetTxnId() string {
	if o == nil || common.IsNil(o.TxnId) {
		var ret string
		return ret
	}
	return *o.TxnId
}

// GetTxnIdOk returns a tuple with the TxnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferToSubAccountOfSameMasterResponse) GetTxnIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TxnId) {
		return nil, false
	}
	return o.TxnId, true
}

// HasTxnId returns a boolean if a field has been set.
func (o *TransferToSubAccountOfSameMasterResponse) HasTxnId() bool {
	if o != nil && !common.IsNil(o.TxnId) {
		return true
	}

	return false
}

// SetTxnId gets a reference to the given string and assigns it to the TxnId field.
func (o *TransferToSubAccountOfSameMasterResponse) SetTxnId(v string) {
	o.TxnId = &v
}

func (o TransferToSubAccountOfSameMasterResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferToSubAccountOfSameMasterResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TxnId) {
		toSerialize["txnId"] = o.TxnId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TransferToSubAccountOfSameMasterResponse) UnmarshalJSON(data []byte) (err error) {
	varTransferToSubAccountOfSameMasterResponse := _TransferToSubAccountOfSameMasterResponse{}

	err = json.Unmarshal(data, &varTransferToSubAccountOfSameMasterResponse)

	if err != nil {
		return err
	}

	*o = TransferToSubAccountOfSameMasterResponse(varTransferToSubAccountOfSameMasterResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "txnId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferToSubAccountOfSameMasterResponse struct {
	value *TransferToSubAccountOfSameMasterResponse
	isSet bool
}

func (v NullableTransferToSubAccountOfSameMasterResponse) Get() *TransferToSubAccountOfSameMasterResponse {
	return v.value
}

func (v *NullableTransferToSubAccountOfSameMasterResponse) Set(val *TransferToSubAccountOfSameMasterResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferToSubAccountOfSameMasterResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferToSubAccountOfSameMasterResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferToSubAccountOfSameMasterResponse(val *TransferToSubAccountOfSameMasterResponse) *NullableTransferToSubAccountOfSameMasterResponse {
	return &NullableTransferToSubAccountOfSameMasterResponse{value: val, isSet: true}
}

func (v NullableTransferToSubAccountOfSameMasterResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferToSubAccountOfSameMasterResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
