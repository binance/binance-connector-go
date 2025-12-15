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

// checks if the SubAccountFuturesAssetTransferResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SubAccountFuturesAssetTransferResponse{}

// SubAccountFuturesAssetTransferResponse struct for SubAccountFuturesAssetTransferResponse
type SubAccountFuturesAssetTransferResponse struct {
	Success              *bool   `json:"success,omitempty"`
	TxnId                *string `json:"txnId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SubAccountFuturesAssetTransferResponse SubAccountFuturesAssetTransferResponse

// NewSubAccountFuturesAssetTransferResponse instantiates a new SubAccountFuturesAssetTransferResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubAccountFuturesAssetTransferResponse() *SubAccountFuturesAssetTransferResponse {
	this := SubAccountFuturesAssetTransferResponse{}
	return &this
}

// NewSubAccountFuturesAssetTransferResponseWithDefaults instantiates a new SubAccountFuturesAssetTransferResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubAccountFuturesAssetTransferResponseWithDefaults() *SubAccountFuturesAssetTransferResponse {
	this := SubAccountFuturesAssetTransferResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SubAccountFuturesAssetTransferResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccountFuturesAssetTransferResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SubAccountFuturesAssetTransferResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *SubAccountFuturesAssetTransferResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetTxnId returns the TxnId field value if set, zero value otherwise.
func (o *SubAccountFuturesAssetTransferResponse) GetTxnId() string {
	if o == nil || common.IsNil(o.TxnId) {
		var ret string
		return ret
	}
	return *o.TxnId
}

// GetTxnIdOk returns a tuple with the TxnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccountFuturesAssetTransferResponse) GetTxnIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TxnId) {
		return nil, false
	}
	return o.TxnId, true
}

// HasTxnId returns a boolean if a field has been set.
func (o *SubAccountFuturesAssetTransferResponse) HasTxnId() bool {
	if o != nil && !common.IsNil(o.TxnId) {
		return true
	}

	return false
}

// SetTxnId gets a reference to the given string and assigns it to the TxnId field.
func (o *SubAccountFuturesAssetTransferResponse) SetTxnId(v string) {
	o.TxnId = &v
}

func (o SubAccountFuturesAssetTransferResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubAccountFuturesAssetTransferResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !common.IsNil(o.TxnId) {
		toSerialize["txnId"] = o.TxnId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SubAccountFuturesAssetTransferResponse) UnmarshalJSON(data []byte) (err error) {
	varSubAccountFuturesAssetTransferResponse := _SubAccountFuturesAssetTransferResponse{}

	err = json.Unmarshal(data, &varSubAccountFuturesAssetTransferResponse)

	if err != nil {
		return err
	}

	*o = SubAccountFuturesAssetTransferResponse(varSubAccountFuturesAssetTransferResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		delete(additionalProperties, "txnId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubAccountFuturesAssetTransferResponse struct {
	value *SubAccountFuturesAssetTransferResponse
	isSet bool
}

func (v NullableSubAccountFuturesAssetTransferResponse) Get() *SubAccountFuturesAssetTransferResponse {
	return v.value
}

func (v *NullableSubAccountFuturesAssetTransferResponse) Set(val *SubAccountFuturesAssetTransferResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubAccountFuturesAssetTransferResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubAccountFuturesAssetTransferResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubAccountFuturesAssetTransferResponse(val *SubAccountFuturesAssetTransferResponse) *NullableSubAccountFuturesAssetTransferResponse {
	return &NullableSubAccountFuturesAssetTransferResponse{value: val, isSet: true}
}

func (v NullableSubAccountFuturesAssetTransferResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubAccountFuturesAssetTransferResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
