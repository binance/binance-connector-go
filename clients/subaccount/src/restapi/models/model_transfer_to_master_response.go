/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the TransferToMasterResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TransferToMasterResponse{}

// TransferToMasterResponse struct for TransferToMasterResponse
type TransferToMasterResponse struct {
	TxnId                *string `json:"txnId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TransferToMasterResponse TransferToMasterResponse

// NewTransferToMasterResponse instantiates a new TransferToMasterResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferToMasterResponse() *TransferToMasterResponse {
	this := TransferToMasterResponse{}
	return &this
}

// NewTransferToMasterResponseWithDefaults instantiates a new TransferToMasterResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferToMasterResponseWithDefaults() *TransferToMasterResponse {
	this := TransferToMasterResponse{}
	return &this
}

// GetTxnId returns the TxnId field value if set, zero value otherwise.
func (o *TransferToMasterResponse) GetTxnId() string {
	if o == nil || common.IsNil(o.TxnId) {
		var ret string
		return ret
	}
	return *o.TxnId
}

// GetTxnIdOk returns a tuple with the TxnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferToMasterResponse) GetTxnIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TxnId) {
		return nil, false
	}
	return o.TxnId, true
}

// HasTxnId returns a boolean if a field has been set.
func (o *TransferToMasterResponse) HasTxnId() bool {
	if o != nil && !common.IsNil(o.TxnId) {
		return true
	}

	return false
}

// SetTxnId gets a reference to the given string and assigns it to the TxnId field.
func (o *TransferToMasterResponse) SetTxnId(v string) {
	o.TxnId = &v
}

func (o TransferToMasterResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferToMasterResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TxnId) {
		toSerialize["txnId"] = o.TxnId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TransferToMasterResponse) UnmarshalJSON(data []byte) (err error) {
	varTransferToMasterResponse := _TransferToMasterResponse{}

	err = json.Unmarshal(data, &varTransferToMasterResponse)

	if err != nil {
		return err
	}

	*o = TransferToMasterResponse(varTransferToMasterResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "txnId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferToMasterResponse struct {
	value *TransferToMasterResponse
	isSet bool
}

func (v NullableTransferToMasterResponse) Get() *TransferToMasterResponse {
	return v.value
}

func (v *NullableTransferToMasterResponse) Set(val *TransferToMasterResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferToMasterResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferToMasterResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferToMasterResponse(val *TransferToMasterResponse) *NullableTransferToMasterResponse {
	return &NullableTransferToMasterResponse{value: val, isSet: true}
}

func (v NullableTransferToMasterResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferToMasterResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
