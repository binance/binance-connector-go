/*
Wallet REST API

Query balances, manage assets, and perform wallet operations via the Binance Wallet API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the BrokerWithdrawResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BrokerWithdrawResponse{}

// BrokerWithdrawResponse struct for BrokerWithdrawResponse
type BrokerWithdrawResponse struct {
	TrId                 *int64  `json:"trId,omitempty"`
	Accepted             *bool   `json:"accepted,omitempty"`
	Info                 *string `json:"info,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BrokerWithdrawResponse BrokerWithdrawResponse

// NewBrokerWithdrawResponse instantiates a new BrokerWithdrawResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrokerWithdrawResponse() *BrokerWithdrawResponse {
	this := BrokerWithdrawResponse{}
	return &this
}

// NewBrokerWithdrawResponseWithDefaults instantiates a new BrokerWithdrawResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrokerWithdrawResponseWithDefaults() *BrokerWithdrawResponse {
	this := BrokerWithdrawResponse{}
	return &this
}

// GetTrId returns the TrId field value if set, zero value otherwise.
func (o *BrokerWithdrawResponse) GetTrId() int64 {
	if o == nil || common.IsNil(o.TrId) {
		var ret int64
		return ret
	}
	return *o.TrId
}

// GetTrIdOk returns a tuple with the TrId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerWithdrawResponse) GetTrIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TrId) {
		return nil, false
	}
	return o.TrId, true
}

// HasTrId returns a boolean if a field has been set.
func (o *BrokerWithdrawResponse) HasTrId() bool {
	if o != nil && !common.IsNil(o.TrId) {
		return true
	}

	return false
}

// SetTrId gets a reference to the given int64 and assigns it to the TrId field.
func (o *BrokerWithdrawResponse) SetTrId(v int64) {
	o.TrId = &v
}

// GetAccepted returns the Accepted field value if set, zero value otherwise.
func (o *BrokerWithdrawResponse) GetAccepted() bool {
	if o == nil || common.IsNil(o.Accepted) {
		var ret bool
		return ret
	}
	return *o.Accepted
}

// GetAcceptedOk returns a tuple with the Accepted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerWithdrawResponse) GetAcceptedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Accepted) {
		return nil, false
	}
	return o.Accepted, true
}

// HasAccepted returns a boolean if a field has been set.
func (o *BrokerWithdrawResponse) HasAccepted() bool {
	if o != nil && !common.IsNil(o.Accepted) {
		return true
	}

	return false
}

// SetAccepted gets a reference to the given bool and assigns it to the Accepted field.
func (o *BrokerWithdrawResponse) SetAccepted(v bool) {
	o.Accepted = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *BrokerWithdrawResponse) GetInfo() string {
	if o == nil || common.IsNil(o.Info) {
		var ret string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerWithdrawResponse) GetInfoOk() (*string, bool) {
	if o == nil || common.IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *BrokerWithdrawResponse) HasInfo() bool {
	if o != nil && !common.IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given string and assigns it to the Info field.
func (o *BrokerWithdrawResponse) SetInfo(v string) {
	o.Info = &v
}

func (o BrokerWithdrawResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BrokerWithdrawResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TrId) {
		toSerialize["trId"] = o.TrId
	}
	if !common.IsNil(o.Accepted) {
		toSerialize["accepted"] = o.Accepted
	}
	if !common.IsNil(o.Info) {
		toSerialize["info"] = o.Info
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BrokerWithdrawResponse) UnmarshalJSON(data []byte) (err error) {
	varBrokerWithdrawResponse := _BrokerWithdrawResponse{}

	err = json.Unmarshal(data, &varBrokerWithdrawResponse)

	if err != nil {
		return err
	}

	*o = BrokerWithdrawResponse(varBrokerWithdrawResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "trId")
		delete(additionalProperties, "accepted")
		delete(additionalProperties, "info")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBrokerWithdrawResponse struct {
	value *BrokerWithdrawResponse
	isSet bool
}

func (v NullableBrokerWithdrawResponse) Get() *BrokerWithdrawResponse {
	return v.value
}

func (v *NullableBrokerWithdrawResponse) Set(val *BrokerWithdrawResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBrokerWithdrawResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBrokerWithdrawResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrokerWithdrawResponse(val *BrokerWithdrawResponse) *NullableBrokerWithdrawResponse {
	return &NullableBrokerWithdrawResponse{value: val, isSet: true}
}

func (v NullableBrokerWithdrawResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrokerWithdrawResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
