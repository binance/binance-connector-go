/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the DustConvertResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DustConvertResponse{}

// DustConvertResponse struct for DustConvertResponse
type DustConvertResponse struct {
	TotalTransfered      *string                                  `json:"totalTransfered,omitempty"`
	TotalServiceCharge   *string                                  `json:"totalServiceCharge,omitempty"`
	TransferResult       []DustConvertResponseTransferResultInner `json:"transferResult,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DustConvertResponse DustConvertResponse

// NewDustConvertResponse instantiates a new DustConvertResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDustConvertResponse() *DustConvertResponse {
	this := DustConvertResponse{}
	return &this
}

// NewDustConvertResponseWithDefaults instantiates a new DustConvertResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDustConvertResponseWithDefaults() *DustConvertResponse {
	this := DustConvertResponse{}
	return &this
}

// GetTotalTransfered returns the TotalTransfered field value if set, zero value otherwise.
func (o *DustConvertResponse) GetTotalTransfered() string {
	if o == nil || common.IsNil(o.TotalTransfered) {
		var ret string
		return ret
	}
	return *o.TotalTransfered
}

// GetTotalTransferedOk returns a tuple with the TotalTransfered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustConvertResponse) GetTotalTransferedOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalTransfered) {
		return nil, false
	}
	return o.TotalTransfered, true
}

// HasTotalTransfered returns a boolean if a field has been set.
func (o *DustConvertResponse) HasTotalTransfered() bool {
	if o != nil && !common.IsNil(o.TotalTransfered) {
		return true
	}

	return false
}

// SetTotalTransfered gets a reference to the given string and assigns it to the TotalTransfered field.
func (o *DustConvertResponse) SetTotalTransfered(v string) {
	o.TotalTransfered = &v
}

// GetTotalServiceCharge returns the TotalServiceCharge field value if set, zero value otherwise.
func (o *DustConvertResponse) GetTotalServiceCharge() string {
	if o == nil || common.IsNil(o.TotalServiceCharge) {
		var ret string
		return ret
	}
	return *o.TotalServiceCharge
}

// GetTotalServiceChargeOk returns a tuple with the TotalServiceCharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustConvertResponse) GetTotalServiceChargeOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalServiceCharge) {
		return nil, false
	}
	return o.TotalServiceCharge, true
}

// HasTotalServiceCharge returns a boolean if a field has been set.
func (o *DustConvertResponse) HasTotalServiceCharge() bool {
	if o != nil && !common.IsNil(o.TotalServiceCharge) {
		return true
	}

	return false
}

// SetTotalServiceCharge gets a reference to the given string and assigns it to the TotalServiceCharge field.
func (o *DustConvertResponse) SetTotalServiceCharge(v string) {
	o.TotalServiceCharge = &v
}

// GetTransferResult returns the TransferResult field value if set, zero value otherwise.
func (o *DustConvertResponse) GetTransferResult() []DustConvertResponseTransferResultInner {
	if o == nil || common.IsNil(o.TransferResult) {
		var ret []DustConvertResponseTransferResultInner
		return ret
	}
	return o.TransferResult
}

// GetTransferResultOk returns a tuple with the TransferResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustConvertResponse) GetTransferResultOk() ([]DustConvertResponseTransferResultInner, bool) {
	if o == nil || common.IsNil(o.TransferResult) {
		return nil, false
	}
	return o.TransferResult, true
}

// HasTransferResult returns a boolean if a field has been set.
func (o *DustConvertResponse) HasTransferResult() bool {
	if o != nil && !common.IsNil(o.TransferResult) {
		return true
	}

	return false
}

// SetTransferResult gets a reference to the given []DustConvertResponseTransferResultInner and assigns it to the TransferResult field.
func (o *DustConvertResponse) SetTransferResult(v []DustConvertResponseTransferResultInner) {
	o.TransferResult = v
}

func (o DustConvertResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DustConvertResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TotalTransfered) {
		toSerialize["totalTransfered"] = o.TotalTransfered
	}
	if !common.IsNil(o.TotalServiceCharge) {
		toSerialize["totalServiceCharge"] = o.TotalServiceCharge
	}
	if !common.IsNil(o.TransferResult) {
		toSerialize["transferResult"] = o.TransferResult
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DustConvertResponse) UnmarshalJSON(data []byte) (err error) {
	varDustConvertResponse := _DustConvertResponse{}

	err = json.Unmarshal(data, &varDustConvertResponse)

	if err != nil {
		return err
	}

	*o = DustConvertResponse(varDustConvertResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "totalTransfered")
		delete(additionalProperties, "totalServiceCharge")
		delete(additionalProperties, "transferResult")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDustConvertResponse struct {
	value *DustConvertResponse
	isSet bool
}

func (v NullableDustConvertResponse) Get() *DustConvertResponse {
	return v.value
}

func (v *NullableDustConvertResponse) Set(val *DustConvertResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDustConvertResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDustConvertResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDustConvertResponse(val *DustConvertResponse) *NullableDustConvertResponse {
	return &NullableDustConvertResponse{value: val, isSet: true}
}

func (v NullableDustConvertResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDustConvertResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
