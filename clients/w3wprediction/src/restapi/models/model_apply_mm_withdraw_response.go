/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ApplyMmWithdrawResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ApplyMmWithdrawResponse{}

// ApplyMmWithdrawResponse struct for ApplyMmWithdrawResponse
type ApplyMmWithdrawResponse struct {
	Id                   *string `json:"id,omitempty"`
	WalletId             *string `json:"walletId,omitempty"`
	WalletAddress        *string `json:"walletAddress,omitempty"`
	TransferId           *string `json:"transferId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplyMmWithdrawResponse ApplyMmWithdrawResponse

// NewApplyMmWithdrawResponse instantiates a new ApplyMmWithdrawResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplyMmWithdrawResponse() *ApplyMmWithdrawResponse {
	this := ApplyMmWithdrawResponse{}
	return &this
}

// NewApplyMmWithdrawResponseWithDefaults instantiates a new ApplyMmWithdrawResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplyMmWithdrawResponseWithDefaults() *ApplyMmWithdrawResponse {
	this := ApplyMmWithdrawResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApplyMmWithdrawResponse) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyMmWithdrawResponse) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApplyMmWithdrawResponse) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApplyMmWithdrawResponse) SetId(v string) {
	o.Id = &v
}

// GetWalletId returns the WalletId field value if set, zero value otherwise.
func (o *ApplyMmWithdrawResponse) GetWalletId() string {
	if o == nil || common.IsNil(o.WalletId) {
		var ret string
		return ret
	}
	return *o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyMmWithdrawResponse) GetWalletIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.WalletId) {
		return nil, false
	}
	return o.WalletId, true
}

// HasWalletId returns a boolean if a field has been set.
func (o *ApplyMmWithdrawResponse) HasWalletId() bool {
	if o != nil && !common.IsNil(o.WalletId) {
		return true
	}

	return false
}

// SetWalletId gets a reference to the given string and assigns it to the WalletId field.
func (o *ApplyMmWithdrawResponse) SetWalletId(v string) {
	o.WalletId = &v
}

// GetWalletAddress returns the WalletAddress field value if set, zero value otherwise.
func (o *ApplyMmWithdrawResponse) GetWalletAddress() string {
	if o == nil || common.IsNil(o.WalletAddress) {
		var ret string
		return ret
	}
	return *o.WalletAddress
}

// GetWalletAddressOk returns a tuple with the WalletAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyMmWithdrawResponse) GetWalletAddressOk() (*string, bool) {
	if o == nil || common.IsNil(o.WalletAddress) {
		return nil, false
	}
	return o.WalletAddress, true
}

// HasWalletAddress returns a boolean if a field has been set.
func (o *ApplyMmWithdrawResponse) HasWalletAddress() bool {
	if o != nil && !common.IsNil(o.WalletAddress) {
		return true
	}

	return false
}

// SetWalletAddress gets a reference to the given string and assigns it to the WalletAddress field.
func (o *ApplyMmWithdrawResponse) SetWalletAddress(v string) {
	o.WalletAddress = &v
}

// GetTransferId returns the TransferId field value if set, zero value otherwise.
func (o *ApplyMmWithdrawResponse) GetTransferId() string {
	if o == nil || common.IsNil(o.TransferId) {
		var ret string
		return ret
	}
	return *o.TransferId
}

// GetTransferIdOk returns a tuple with the TransferId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyMmWithdrawResponse) GetTransferIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TransferId) {
		return nil, false
	}
	return o.TransferId, true
}

// HasTransferId returns a boolean if a field has been set.
func (o *ApplyMmWithdrawResponse) HasTransferId() bool {
	if o != nil && !common.IsNil(o.TransferId) {
		return true
	}

	return false
}

// SetTransferId gets a reference to the given string and assigns it to the TransferId field.
func (o *ApplyMmWithdrawResponse) SetTransferId(v string) {
	o.TransferId = &v
}

func (o ApplyMmWithdrawResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplyMmWithdrawResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.WalletId) {
		toSerialize["walletId"] = o.WalletId
	}
	if !common.IsNil(o.WalletAddress) {
		toSerialize["walletAddress"] = o.WalletAddress
	}
	if !common.IsNil(o.TransferId) {
		toSerialize["transferId"] = o.TransferId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApplyMmWithdrawResponse) UnmarshalJSON(data []byte) (err error) {
	varApplyMmWithdrawResponse := _ApplyMmWithdrawResponse{}

	err = json.Unmarshal(data, &varApplyMmWithdrawResponse)

	if err != nil {
		return err
	}

	*o = ApplyMmWithdrawResponse(varApplyMmWithdrawResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "walletId")
		delete(additionalProperties, "walletAddress")
		delete(additionalProperties, "transferId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplyMmWithdrawResponse struct {
	value *ApplyMmWithdrawResponse
	isSet bool
}

func (v NullableApplyMmWithdrawResponse) Get() *ApplyMmWithdrawResponse {
	return v.value
}

func (v *NullableApplyMmWithdrawResponse) Set(val *ApplyMmWithdrawResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApplyMmWithdrawResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApplyMmWithdrawResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplyMmWithdrawResponse(val *ApplyMmWithdrawResponse) *NullableApplyMmWithdrawResponse {
	return &NullableApplyMmWithdrawResponse{value: val, isSet: true}
}

func (v NullableApplyMmWithdrawResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplyMmWithdrawResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
