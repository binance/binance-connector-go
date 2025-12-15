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

// checks if the MarginTransferForSubAccountResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MarginTransferForSubAccountResponse{}

// MarginTransferForSubAccountResponse struct for MarginTransferForSubAccountResponse
type MarginTransferForSubAccountResponse struct {
	TxnId                *string `json:"txnId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarginTransferForSubAccountResponse MarginTransferForSubAccountResponse

// NewMarginTransferForSubAccountResponse instantiates a new MarginTransferForSubAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginTransferForSubAccountResponse() *MarginTransferForSubAccountResponse {
	this := MarginTransferForSubAccountResponse{}
	return &this
}

// NewMarginTransferForSubAccountResponseWithDefaults instantiates a new MarginTransferForSubAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginTransferForSubAccountResponseWithDefaults() *MarginTransferForSubAccountResponse {
	this := MarginTransferForSubAccountResponse{}
	return &this
}

// GetTxnId returns the TxnId field value if set, zero value otherwise.
func (o *MarginTransferForSubAccountResponse) GetTxnId() string {
	if o == nil || common.IsNil(o.TxnId) {
		var ret string
		return ret
	}
	return *o.TxnId
}

// GetTxnIdOk returns a tuple with the TxnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginTransferForSubAccountResponse) GetTxnIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TxnId) {
		return nil, false
	}
	return o.TxnId, true
}

// HasTxnId returns a boolean if a field has been set.
func (o *MarginTransferForSubAccountResponse) HasTxnId() bool {
	if o != nil && !common.IsNil(o.TxnId) {
		return true
	}

	return false
}

// SetTxnId gets a reference to the given string and assigns it to the TxnId field.
func (o *MarginTransferForSubAccountResponse) SetTxnId(v string) {
	o.TxnId = &v
}

func (o MarginTransferForSubAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarginTransferForSubAccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TxnId) {
		toSerialize["txnId"] = o.TxnId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MarginTransferForSubAccountResponse) UnmarshalJSON(data []byte) (err error) {
	varMarginTransferForSubAccountResponse := _MarginTransferForSubAccountResponse{}

	err = json.Unmarshal(data, &varMarginTransferForSubAccountResponse)

	if err != nil {
		return err
	}

	*o = MarginTransferForSubAccountResponse(varMarginTransferForSubAccountResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "txnId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarginTransferForSubAccountResponse struct {
	value *MarginTransferForSubAccountResponse
	isSet bool
}

func (v NullableMarginTransferForSubAccountResponse) Get() *MarginTransferForSubAccountResponse {
	return v.value
}

func (v *NullableMarginTransferForSubAccountResponse) Set(val *MarginTransferForSubAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginTransferForSubAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginTransferForSubAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginTransferForSubAccountResponse(val *MarginTransferForSubAccountResponse) *NullableMarginTransferForSubAccountResponse {
	return &NullableMarginTransferForSubAccountResponse{value: val, isSet: true}
}

func (v NullableMarginTransferForSubAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginTransferForSubAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
