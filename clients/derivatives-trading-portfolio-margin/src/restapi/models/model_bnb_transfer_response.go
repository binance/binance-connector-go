/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the BnbTransferResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BnbTransferResponse{}

// BnbTransferResponse struct for BnbTransferResponse
type BnbTransferResponse struct {
	TranId               *int64 `json:"tranId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BnbTransferResponse BnbTransferResponse

// NewBnbTransferResponse instantiates a new BnbTransferResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBnbTransferResponse() *BnbTransferResponse {
	this := BnbTransferResponse{}
	return &this
}

// NewBnbTransferResponseWithDefaults instantiates a new BnbTransferResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBnbTransferResponseWithDefaults() *BnbTransferResponse {
	this := BnbTransferResponse{}
	return &this
}

// GetTranId returns the TranId field value if set, zero value otherwise.
func (o *BnbTransferResponse) GetTranId() int64 {
	if o == nil || common.IsNil(o.TranId) {
		var ret int64
		return ret
	}
	return *o.TranId
}

// GetTranIdOk returns a tuple with the TranId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BnbTransferResponse) GetTranIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TranId) {
		return nil, false
	}
	return o.TranId, true
}

// HasTranId returns a boolean if a field has been set.
func (o *BnbTransferResponse) HasTranId() bool {
	if o != nil && !common.IsNil(o.TranId) {
		return true
	}

	return false
}

// SetTranId gets a reference to the given int64 and assigns it to the TranId field.
func (o *BnbTransferResponse) SetTranId(v int64) {
	o.TranId = &v
}

func (o BnbTransferResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BnbTransferResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TranId) {
		toSerialize["tranId"] = o.TranId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BnbTransferResponse) UnmarshalJSON(data []byte) (err error) {
	varBnbTransferResponse := _BnbTransferResponse{}

	err = json.Unmarshal(data, &varBnbTransferResponse)

	if err != nil {
		return err
	}

	*o = BnbTransferResponse(varBnbTransferResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tranId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBnbTransferResponse struct {
	value *BnbTransferResponse
	isSet bool
}

func (v NullableBnbTransferResponse) Get() *BnbTransferResponse {
	return v.value
}

func (v *NullableBnbTransferResponse) Set(val *BnbTransferResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBnbTransferResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBnbTransferResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBnbTransferResponse(val *BnbTransferResponse) *NullableBnbTransferResponse {
	return &NullableBnbTransferResponse{value: val, isSet: true}
}

func (v NullableBnbTransferResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBnbTransferResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
