/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the FuturesTransferForSubAccountResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FuturesTransferForSubAccountResponse{}

// FuturesTransferForSubAccountResponse struct for FuturesTransferForSubAccountResponse
type FuturesTransferForSubAccountResponse struct {
	TxnId                *string `json:"txnId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FuturesTransferForSubAccountResponse FuturesTransferForSubAccountResponse

// NewFuturesTransferForSubAccountResponse instantiates a new FuturesTransferForSubAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFuturesTransferForSubAccountResponse() *FuturesTransferForSubAccountResponse {
	this := FuturesTransferForSubAccountResponse{}
	return &this
}

// NewFuturesTransferForSubAccountResponseWithDefaults instantiates a new FuturesTransferForSubAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFuturesTransferForSubAccountResponseWithDefaults() *FuturesTransferForSubAccountResponse {
	this := FuturesTransferForSubAccountResponse{}
	return &this
}

// GetTxnId returns the TxnId field value if set, zero value otherwise.
func (o *FuturesTransferForSubAccountResponse) GetTxnId() string {
	if o == nil || common.IsNil(o.TxnId) {
		var ret string
		return ret
	}
	return *o.TxnId
}

// GetTxnIdOk returns a tuple with the TxnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesTransferForSubAccountResponse) GetTxnIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TxnId) {
		return nil, false
	}
	return o.TxnId, true
}

// HasTxnId returns a boolean if a field has been set.
func (o *FuturesTransferForSubAccountResponse) HasTxnId() bool {
	if o != nil && !common.IsNil(o.TxnId) {
		return true
	}

	return false
}

// SetTxnId gets a reference to the given string and assigns it to the TxnId field.
func (o *FuturesTransferForSubAccountResponse) SetTxnId(v string) {
	o.TxnId = &v
}

func (o FuturesTransferForSubAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FuturesTransferForSubAccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TxnId) {
		toSerialize["txnId"] = o.TxnId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FuturesTransferForSubAccountResponse) UnmarshalJSON(data []byte) (err error) {
	varFuturesTransferForSubAccountResponse := _FuturesTransferForSubAccountResponse{}

	err = json.Unmarshal(data, &varFuturesTransferForSubAccountResponse)

	if err != nil {
		return err
	}

	*o = FuturesTransferForSubAccountResponse(varFuturesTransferForSubAccountResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "txnId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFuturesTransferForSubAccountResponse struct {
	value *FuturesTransferForSubAccountResponse
	isSet bool
}

func (v NullableFuturesTransferForSubAccountResponse) Get() *FuturesTransferForSubAccountResponse {
	return v.value
}

func (v *NullableFuturesTransferForSubAccountResponse) Set(val *FuturesTransferForSubAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFuturesTransferForSubAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFuturesTransferForSubAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFuturesTransferForSubAccountResponse(val *FuturesTransferForSubAccountResponse) *NullableFuturesTransferForSubAccountResponse {
	return &NullableFuturesTransferForSubAccountResponse{value: val, isSet: true}
}

func (v NullableFuturesTransferForSubAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFuturesTransferForSubAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
