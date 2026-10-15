/*
Prediction Trading Streams

Binance Web3 Prediction pushes 14 types of per-user business events in real time via the Binance SApi WebSocket channel, covering trading (market / limit), reward claim, fund transfer and market resolution, plus order book (bid/ask) quotes broadcast per topic. Business events are routed by exact `(topic, userId)` match; order book topics are broadcast to every subscriber. Delivery is at-most-once, ordering is not guaranteed, and offline messages are not cached — pull the latest state via the REST API after reconnecting and deduplicate on `pushId` (business events) or `marketId` + `updateTimestampMs` (order book).
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the PredictionOrderBookAllResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PredictionOrderBookAllResponse{}

// PredictionOrderBookAllResponse struct for PredictionOrderBookAllResponse
type PredictionOrderBookAllResponse struct {
	// Message type, fixed value `orderbook`; use it to dispatch message types.
	MsgType *string `json:"msgType,omitempty"`
	// Binance internal market ID (mapped from the upstream vendor ID), aligned with the REST API.
	MarketId *int64 `json:"marketId,omitempty"`
	// Quote update time, millisecond UTC timestamp. No strict ordering is guaranteed for the same `marketId`, so use this field for out-of-order detection and discard messages older than the current snapshot.
	UpdateTimestampMs *int64 `json:"updateTimestampMs,omitempty"`
	// Ask levels sorted by price **ascending**; an empty array means no asks. Typically the top N levels (generally <= 50, depending on upstream policy).
	Asks [][]string `json:"asks,omitempty"`
	// Bid levels sorted by price **descending**; an empty array means no bids. Typically the top N levels (generally <= 50, depending on upstream policy).
	Bids                 [][]string `json:"bids,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PredictionOrderBookAllResponse PredictionOrderBookAllResponse

// NewPredictionOrderBookAllResponse instantiates a new PredictionOrderBookAllResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPredictionOrderBookAllResponse() *PredictionOrderBookAllResponse {
	this := PredictionOrderBookAllResponse{}
	return &this
}

// NewPredictionOrderBookAllResponseWithDefaults instantiates a new PredictionOrderBookAllResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPredictionOrderBookAllResponseWithDefaults() *PredictionOrderBookAllResponse {
	this := PredictionOrderBookAllResponse{}
	return &this
}

// GetMsgType returns the MsgType field value if set, zero value otherwise.
func (o *PredictionOrderBookAllResponse) GetMsgType() string {
	if o == nil || common.IsNil(o.MsgType) {
		var ret string
		return ret
	}
	return *o.MsgType
}

// GetMsgTypeOk returns a tuple with the MsgType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PredictionOrderBookAllResponse) GetMsgTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.MsgType) {
		return nil, false
	}
	return o.MsgType, true
}

// HasMsgType returns a boolean if a field has been set.
func (o *PredictionOrderBookAllResponse) HasMsgType() bool {
	if o != nil && !common.IsNil(o.MsgType) {
		return true
	}

	return false
}

// SetMsgType gets a reference to the given string and assigns it to the MsgType field.
func (o *PredictionOrderBookAllResponse) SetMsgType(v string) {
	o.MsgType = &v
}

// GetMarketId returns the MarketId field value if set, zero value otherwise.
func (o *PredictionOrderBookAllResponse) GetMarketId() int64 {
	if o == nil || common.IsNil(o.MarketId) {
		var ret int64
		return ret
	}
	return *o.MarketId
}

// GetMarketIdOk returns a tuple with the MarketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PredictionOrderBookAllResponse) GetMarketIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.MarketId) {
		return nil, false
	}
	return o.MarketId, true
}

// HasMarketId returns a boolean if a field has been set.
func (o *PredictionOrderBookAllResponse) HasMarketId() bool {
	if o != nil && !common.IsNil(o.MarketId) {
		return true
	}

	return false
}

// SetMarketId gets a reference to the given int64 and assigns it to the MarketId field.
func (o *PredictionOrderBookAllResponse) SetMarketId(v int64) {
	o.MarketId = &v
}

// GetUpdateTimestampMs returns the UpdateTimestampMs field value if set, zero value otherwise.
func (o *PredictionOrderBookAllResponse) GetUpdateTimestampMs() int64 {
	if o == nil || common.IsNil(o.UpdateTimestampMs) {
		var ret int64
		return ret
	}
	return *o.UpdateTimestampMs
}

// GetUpdateTimestampMsOk returns a tuple with the UpdateTimestampMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PredictionOrderBookAllResponse) GetUpdateTimestampMsOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTimestampMs) {
		return nil, false
	}
	return o.UpdateTimestampMs, true
}

// HasUpdateTimestampMs returns a boolean if a field has been set.
func (o *PredictionOrderBookAllResponse) HasUpdateTimestampMs() bool {
	if o != nil && !common.IsNil(o.UpdateTimestampMs) {
		return true
	}

	return false
}

// SetUpdateTimestampMs gets a reference to the given int64 and assigns it to the UpdateTimestampMs field.
func (o *PredictionOrderBookAllResponse) SetUpdateTimestampMs(v int64) {
	o.UpdateTimestampMs = &v
}

// GetAsks returns the Asks field value if set, zero value otherwise.
func (o *PredictionOrderBookAllResponse) GetAsks() [][]string {
	if o == nil || common.IsNil(o.Asks) {
		var ret [][]string
		return ret
	}
	return o.Asks
}

// GetAsksOk returns a tuple with the Asks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PredictionOrderBookAllResponse) GetAsksOk() ([][]string, bool) {
	if o == nil || common.IsNil(o.Asks) {
		return nil, false
	}
	return o.Asks, true
}

// HasAsks returns a boolean if a field has been set.
func (o *PredictionOrderBookAllResponse) HasAsks() bool {
	if o != nil && !common.IsNil(o.Asks) {
		return true
	}

	return false
}

// SetAsks gets a reference to the given [][]string and assigns it to the Asks field.
func (o *PredictionOrderBookAllResponse) SetAsks(v [][]string) {
	o.Asks = v
}

// GetBids returns the Bids field value if set, zero value otherwise.
func (o *PredictionOrderBookAllResponse) GetBids() [][]string {
	if o == nil || common.IsNil(o.Bids) {
		var ret [][]string
		return ret
	}
	return o.Bids
}

// GetBidsOk returns a tuple with the Bids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PredictionOrderBookAllResponse) GetBidsOk() ([][]string, bool) {
	if o == nil || common.IsNil(o.Bids) {
		return nil, false
	}
	return o.Bids, true
}

// HasBids returns a boolean if a field has been set.
func (o *PredictionOrderBookAllResponse) HasBids() bool {
	if o != nil && !common.IsNil(o.Bids) {
		return true
	}

	return false
}

// SetBids gets a reference to the given [][]string and assigns it to the Bids field.
func (o *PredictionOrderBookAllResponse) SetBids(v [][]string) {
	o.Bids = v
}

func (o PredictionOrderBookAllResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PredictionOrderBookAllResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.MsgType) {
		toSerialize["msgType"] = o.MsgType
	}
	if !common.IsNil(o.MarketId) {
		toSerialize["marketId"] = o.MarketId
	}
	if !common.IsNil(o.UpdateTimestampMs) {
		toSerialize["updateTimestampMs"] = o.UpdateTimestampMs
	}
	if !common.IsNil(o.Asks) {
		toSerialize["asks"] = o.Asks
	}
	if !common.IsNil(o.Bids) {
		toSerialize["bids"] = o.Bids
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PredictionOrderBookAllResponse) UnmarshalJSON(data []byte) (err error) {
	varPredictionOrderBookAllResponse := _PredictionOrderBookAllResponse{}

	err = json.Unmarshal(data, &varPredictionOrderBookAllResponse)

	if err != nil {
		return err
	}

	*o = PredictionOrderBookAllResponse(varPredictionOrderBookAllResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "msgType")
		delete(additionalProperties, "marketId")
		delete(additionalProperties, "updateTimestampMs")
		delete(additionalProperties, "asks")
		delete(additionalProperties, "bids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePredictionOrderBookAllResponse struct {
	value *PredictionOrderBookAllResponse
	isSet bool
}

func (v NullablePredictionOrderBookAllResponse) Get() *PredictionOrderBookAllResponse {
	return v.value
}

func (v *NullablePredictionOrderBookAllResponse) Set(val *PredictionOrderBookAllResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePredictionOrderBookAllResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePredictionOrderBookAllResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePredictionOrderBookAllResponse(val *PredictionOrderBookAllResponse) *NullablePredictionOrderBookAllResponse {
	return &NullablePredictionOrderBookAllResponse{value: val, isSet: true}
}

func (v NullablePredictionOrderBookAllResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePredictionOrderBookAllResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
