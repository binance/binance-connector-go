/*
Prediction Trading Streams

Binance Web3 Prediction pushes 14 types of per-user business events in real time via the Binance SApi WebSocket channel, covering trading (market / limit), reward claim, fund transfer and market resolution, plus order book (bid/ask) quotes broadcast per topic. Business events are routed by exact `(topic, userId)` match; order book topics are broadcast to every subscriber. Delivery is at-most-once, ordering is not guaranteed, and offline messages are not cached — pull the latest state via the REST API after reconnecting and deduplicate on `pushId` (business events) or `marketId` + `updateTimestampMs` (order book).
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the FundTransferSucceededResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FundTransferSucceededResponse{}

// FundTransferSucceededResponse struct for FundTransferSucceededResponse
type FundTransferSucceededResponse struct {
	// Unique push ID in the form `pm_<refId>_<scenarioCode>_<8hex>`, present in every topic; `refId` is `FundTransferOrder.transferId`. Use it for client-side deduplication.
	PushId *string `json:"pushId,omitempty"`
	// Token amount transferred, from `FundTransferOrder.fromTokenAmount`. Truncated to 2 decimals (towards zero); rendered as `<0.01` when the value is below 0.01.
	Amount               *string `json:"amount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FundTransferSucceededResponse FundTransferSucceededResponse

// NewFundTransferSucceededResponse instantiates a new FundTransferSucceededResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFundTransferSucceededResponse() *FundTransferSucceededResponse {
	this := FundTransferSucceededResponse{}
	return &this
}

// NewFundTransferSucceededResponseWithDefaults instantiates a new FundTransferSucceededResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFundTransferSucceededResponseWithDefaults() *FundTransferSucceededResponse {
	this := FundTransferSucceededResponse{}
	return &this
}

// GetPushId returns the PushId field value if set, zero value otherwise.
func (o *FundTransferSucceededResponse) GetPushId() string {
	if o == nil || common.IsNil(o.PushId) {
		var ret string
		return ret
	}
	return *o.PushId
}

// GetPushIdOk returns a tuple with the PushId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundTransferSucceededResponse) GetPushIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.PushId) {
		return nil, false
	}
	return o.PushId, true
}

// HasPushId returns a boolean if a field has been set.
func (o *FundTransferSucceededResponse) HasPushId() bool {
	if o != nil && !common.IsNil(o.PushId) {
		return true
	}

	return false
}

// SetPushId gets a reference to the given string and assigns it to the PushId field.
func (o *FundTransferSucceededResponse) SetPushId(v string) {
	o.PushId = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *FundTransferSucceededResponse) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundTransferSucceededResponse) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *FundTransferSucceededResponse) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *FundTransferSucceededResponse) SetAmount(v string) {
	o.Amount = &v
}

func (o FundTransferSucceededResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FundTransferSucceededResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.PushId) {
		toSerialize["pushId"] = o.PushId
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FundTransferSucceededResponse) UnmarshalJSON(data []byte) (err error) {
	varFundTransferSucceededResponse := _FundTransferSucceededResponse{}

	err = json.Unmarshal(data, &varFundTransferSucceededResponse)

	if err != nil {
		return err
	}

	*o = FundTransferSucceededResponse(varFundTransferSucceededResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pushId")
		delete(additionalProperties, "amount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFundTransferSucceededResponse struct {
	value *FundTransferSucceededResponse
	isSet bool
}

func (v NullableFundTransferSucceededResponse) Get() *FundTransferSucceededResponse {
	return v.value
}

func (v *NullableFundTransferSucceededResponse) Set(val *FundTransferSucceededResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFundTransferSucceededResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFundTransferSucceededResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFundTransferSucceededResponse(val *FundTransferSucceededResponse) *NullableFundTransferSucceededResponse {
	return &NullableFundTransferSucceededResponse{value: val, isSet: true}
}

func (v NullableFundTransferSucceededResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFundTransferSucceededResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
