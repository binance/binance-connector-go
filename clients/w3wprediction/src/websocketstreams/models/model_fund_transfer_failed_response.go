/*
Prediction Trading Streams

Binance Web3 Prediction pushes 14 kinds of real-time events over the Binance SApi WebSocket channel, covering market and limit trading, reward claims, fund transfers and market resolution, plus order book (bid/ask) quotes per market. Trading events go only to the user they belong to; order book quotes go to everyone subscribed. A message may occasionally be missed, order is not guaranteed, and nothing is stored while you are offline — after reconnecting, fetch the latest state from the REST API and drop duplicates using `pushId` (business events) or `marketId` + `updateTimestampMs` (order book).
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the FundTransferFailedResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FundTransferFailedResponse{}

// FundTransferFailedResponse struct for FundTransferFailedResponse
type FundTransferFailedResponse struct {
	// Unique ID of this push. Use it to remove duplicates.
	PushId *string `json:"pushId,omitempty"`
	// Amount of the attempted transfer. Rounded down to 2 decimals; shown as `<0.01` if smaller.
	Amount               *string `json:"amount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FundTransferFailedResponse FundTransferFailedResponse

// NewFundTransferFailedResponse instantiates a new FundTransferFailedResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFundTransferFailedResponse() *FundTransferFailedResponse {
	this := FundTransferFailedResponse{}
	return &this
}

// NewFundTransferFailedResponseWithDefaults instantiates a new FundTransferFailedResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFundTransferFailedResponseWithDefaults() *FundTransferFailedResponse {
	this := FundTransferFailedResponse{}
	return &this
}

// GetPushId returns the PushId field value if set, zero value otherwise.
func (o *FundTransferFailedResponse) GetPushId() string {
	if o == nil || common.IsNil(o.PushId) {
		var ret string
		return ret
	}
	return *o.PushId
}

// GetPushIdOk returns a tuple with the PushId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundTransferFailedResponse) GetPushIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.PushId) {
		return nil, false
	}
	return o.PushId, true
}

// HasPushId returns a boolean if a field has been set.
func (o *FundTransferFailedResponse) HasPushId() bool {
	if o != nil && !common.IsNil(o.PushId) {
		return true
	}

	return false
}

// SetPushId gets a reference to the given string and assigns it to the PushId field.
func (o *FundTransferFailedResponse) SetPushId(v string) {
	o.PushId = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *FundTransferFailedResponse) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundTransferFailedResponse) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *FundTransferFailedResponse) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *FundTransferFailedResponse) SetAmount(v string) {
	o.Amount = &v
}

func (o FundTransferFailedResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FundTransferFailedResponse) ToMap() (map[string]interface{}, error) {
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

func (o *FundTransferFailedResponse) UnmarshalJSON(data []byte) (err error) {
	varFundTransferFailedResponse := _FundTransferFailedResponse{}

	err = json.Unmarshal(data, &varFundTransferFailedResponse)

	if err != nil {
		return err
	}

	*o = FundTransferFailedResponse(varFundTransferFailedResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pushId")
		delete(additionalProperties, "amount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFundTransferFailedResponse struct {
	value *FundTransferFailedResponse
	isSet bool
}

func (v NullableFundTransferFailedResponse) Get() *FundTransferFailedResponse {
	return v.value
}

func (v *NullableFundTransferFailedResponse) Set(val *FundTransferFailedResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFundTransferFailedResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFundTransferFailedResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFundTransferFailedResponse(val *FundTransferFailedResponse) *NullableFundTransferFailedResponse {
	return &NullableFundTransferFailedResponse{value: val, isSet: true}
}

func (v NullableFundTransferFailedResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFundTransferFailedResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
