/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the DustTransferResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DustTransferResponse{}

// DustTransferResponse struct for DustTransferResponse
type DustTransferResponse struct {
	TotalServiceCharge   *string                                   `json:"totalServiceCharge,omitempty"`
	TotalTransfered      *string                                   `json:"totalTransfered,omitempty"`
	TransferResult       []DustTransferResponseTransferResultInner `json:"transferResult,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DustTransferResponse DustTransferResponse

// NewDustTransferResponse instantiates a new DustTransferResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDustTransferResponse() *DustTransferResponse {
	this := DustTransferResponse{}
	return &this
}

// NewDustTransferResponseWithDefaults instantiates a new DustTransferResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDustTransferResponseWithDefaults() *DustTransferResponse {
	this := DustTransferResponse{}
	return &this
}

// GetTotalServiceCharge returns the TotalServiceCharge field value if set, zero value otherwise.
func (o *DustTransferResponse) GetTotalServiceCharge() string {
	if o == nil || common.IsNil(o.TotalServiceCharge) {
		var ret string
		return ret
	}
	return *o.TotalServiceCharge
}

// GetTotalServiceChargeOk returns a tuple with the TotalServiceCharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustTransferResponse) GetTotalServiceChargeOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalServiceCharge) {
		return nil, false
	}
	return o.TotalServiceCharge, true
}

// HasTotalServiceCharge returns a boolean if a field has been set.
func (o *DustTransferResponse) HasTotalServiceCharge() bool {
	if o != nil && !common.IsNil(o.TotalServiceCharge) {
		return true
	}

	return false
}

// SetTotalServiceCharge gets a reference to the given string and assigns it to the TotalServiceCharge field.
func (o *DustTransferResponse) SetTotalServiceCharge(v string) {
	o.TotalServiceCharge = &v
}

// GetTotalTransfered returns the TotalTransfered field value if set, zero value otherwise.
func (o *DustTransferResponse) GetTotalTransfered() string {
	if o == nil || common.IsNil(o.TotalTransfered) {
		var ret string
		return ret
	}
	return *o.TotalTransfered
}

// GetTotalTransferedOk returns a tuple with the TotalTransfered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustTransferResponse) GetTotalTransferedOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalTransfered) {
		return nil, false
	}
	return o.TotalTransfered, true
}

// HasTotalTransfered returns a boolean if a field has been set.
func (o *DustTransferResponse) HasTotalTransfered() bool {
	if o != nil && !common.IsNil(o.TotalTransfered) {
		return true
	}

	return false
}

// SetTotalTransfered gets a reference to the given string and assigns it to the TotalTransfered field.
func (o *DustTransferResponse) SetTotalTransfered(v string) {
	o.TotalTransfered = &v
}

// GetTransferResult returns the TransferResult field value if set, zero value otherwise.
func (o *DustTransferResponse) GetTransferResult() []DustTransferResponseTransferResultInner {
	if o == nil || common.IsNil(o.TransferResult) {
		var ret []DustTransferResponseTransferResultInner
		return ret
	}
	return o.TransferResult
}

// GetTransferResultOk returns a tuple with the TransferResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DustTransferResponse) GetTransferResultOk() ([]DustTransferResponseTransferResultInner, bool) {
	if o == nil || common.IsNil(o.TransferResult) {
		return nil, false
	}
	return o.TransferResult, true
}

// HasTransferResult returns a boolean if a field has been set.
func (o *DustTransferResponse) HasTransferResult() bool {
	if o != nil && !common.IsNil(o.TransferResult) {
		return true
	}

	return false
}

// SetTransferResult gets a reference to the given []DustTransferResponseTransferResultInner and assigns it to the TransferResult field.
func (o *DustTransferResponse) SetTransferResult(v []DustTransferResponseTransferResultInner) {
	o.TransferResult = v
}

func (o DustTransferResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DustTransferResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TotalServiceCharge) {
		toSerialize["totalServiceCharge"] = o.TotalServiceCharge
	}
	if !common.IsNil(o.TotalTransfered) {
		toSerialize["totalTransfered"] = o.TotalTransfered
	}
	if !common.IsNil(o.TransferResult) {
		toSerialize["transferResult"] = o.TransferResult
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DustTransferResponse) UnmarshalJSON(data []byte) (err error) {
	varDustTransferResponse := _DustTransferResponse{}

	err = json.Unmarshal(data, &varDustTransferResponse)

	if err != nil {
		return err
	}

	*o = DustTransferResponse(varDustTransferResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "totalServiceCharge")
		delete(additionalProperties, "totalTransfered")
		delete(additionalProperties, "transferResult")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDustTransferResponse struct {
	value *DustTransferResponse
	isSet bool
}

func (v NullableDustTransferResponse) Get() *DustTransferResponse {
	return v.value
}

func (v *NullableDustTransferResponse) Set(val *DustTransferResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDustTransferResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDustTransferResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDustTransferResponse(val *DustTransferResponse) *NullableDustTransferResponse {
	return &NullableDustTransferResponse{value: val, isSet: true}
}

func (v NullableDustTransferResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDustTransferResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
