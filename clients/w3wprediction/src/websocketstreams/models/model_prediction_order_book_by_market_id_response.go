/*
Prediction Trading Streams

Binance Web3 Prediction pushes 14 kinds of real-time events over the Binance SApi WebSocket channel, covering market and limit trading, reward claims, fund transfers and market resolution, plus order book (bid/ask) quotes per market. Trading events go only to the user they belong to; order book quotes go to everyone subscribed. A message may occasionally be missed, order is not guaranteed, and nothing is stored while you are offline — after reconnecting, fetch the latest state from the REST API and drop duplicates using `pushId` (business events) or `marketId` + `updateTimestampMs` (order book).
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the PredictionOrderBookByMarketIdResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PredictionOrderBookByMarketIdResponse{}

// PredictionOrderBookByMarketIdResponse struct for PredictionOrderBookByMarketIdResponse
type PredictionOrderBookByMarketIdResponse struct {
	// Message type, always `orderbook`.
	MsgType *string `json:"msgType,omitempty"`
	// Market ID of this order book, the same `markets[].marketId` used by the REST API.
	MarketId *int64 `json:"marketId,omitempty"`
	// Quote update time in milliseconds (UTC). Messages for the same market may arrive out of order, so discard anything older than the snapshot you already have.
	UpdateTimestampMs *int64 `json:"updateTimestampMs,omitempty"`
	// Ask levels sorted by price **ascending**, usually the top 50 or fewer; an empty array means no asks. Values may carry floating-point noise (e.g. `6865.999999999999`), so round them before display.
	Asks [][]float64 `json:"asks,omitempty"`
	// Bid levels sorted by price **descending**, usually the top 50 or fewer; an empty array means no bids. Values may carry floating-point noise (e.g. `6865.999999999999`), so round them before display.
	Bids                 [][]float64 `json:"bids,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PredictionOrderBookByMarketIdResponse PredictionOrderBookByMarketIdResponse

// NewPredictionOrderBookByMarketIdResponse instantiates a new PredictionOrderBookByMarketIdResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPredictionOrderBookByMarketIdResponse() *PredictionOrderBookByMarketIdResponse {
	this := PredictionOrderBookByMarketIdResponse{}
	return &this
}

// NewPredictionOrderBookByMarketIdResponseWithDefaults instantiates a new PredictionOrderBookByMarketIdResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPredictionOrderBookByMarketIdResponseWithDefaults() *PredictionOrderBookByMarketIdResponse {
	this := PredictionOrderBookByMarketIdResponse{}
	return &this
}

// GetMsgType returns the MsgType field value if set, zero value otherwise.
func (o *PredictionOrderBookByMarketIdResponse) GetMsgType() string {
	if o == nil || common.IsNil(o.MsgType) {
		var ret string
		return ret
	}
	return *o.MsgType
}

// GetMsgTypeOk returns a tuple with the MsgType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PredictionOrderBookByMarketIdResponse) GetMsgTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.MsgType) {
		return nil, false
	}
	return o.MsgType, true
}

// HasMsgType returns a boolean if a field has been set.
func (o *PredictionOrderBookByMarketIdResponse) HasMsgType() bool {
	if o != nil && !common.IsNil(o.MsgType) {
		return true
	}

	return false
}

// SetMsgType gets a reference to the given string and assigns it to the MsgType field.
func (o *PredictionOrderBookByMarketIdResponse) SetMsgType(v string) {
	o.MsgType = &v
}

// GetMarketId returns the MarketId field value if set, zero value otherwise.
func (o *PredictionOrderBookByMarketIdResponse) GetMarketId() int64 {
	if o == nil || common.IsNil(o.MarketId) {
		var ret int64
		return ret
	}
	return *o.MarketId
}

// GetMarketIdOk returns a tuple with the MarketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PredictionOrderBookByMarketIdResponse) GetMarketIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.MarketId) {
		return nil, false
	}
	return o.MarketId, true
}

// HasMarketId returns a boolean if a field has been set.
func (o *PredictionOrderBookByMarketIdResponse) HasMarketId() bool {
	if o != nil && !common.IsNil(o.MarketId) {
		return true
	}

	return false
}

// SetMarketId gets a reference to the given int64 and assigns it to the MarketId field.
func (o *PredictionOrderBookByMarketIdResponse) SetMarketId(v int64) {
	o.MarketId = &v
}

// GetUpdateTimestampMs returns the UpdateTimestampMs field value if set, zero value otherwise.
func (o *PredictionOrderBookByMarketIdResponse) GetUpdateTimestampMs() int64 {
	if o == nil || common.IsNil(o.UpdateTimestampMs) {
		var ret int64
		return ret
	}
	return *o.UpdateTimestampMs
}

// GetUpdateTimestampMsOk returns a tuple with the UpdateTimestampMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PredictionOrderBookByMarketIdResponse) GetUpdateTimestampMsOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTimestampMs) {
		return nil, false
	}
	return o.UpdateTimestampMs, true
}

// HasUpdateTimestampMs returns a boolean if a field has been set.
func (o *PredictionOrderBookByMarketIdResponse) HasUpdateTimestampMs() bool {
	if o != nil && !common.IsNil(o.UpdateTimestampMs) {
		return true
	}

	return false
}

// SetUpdateTimestampMs gets a reference to the given int64 and assigns it to the UpdateTimestampMs field.
func (o *PredictionOrderBookByMarketIdResponse) SetUpdateTimestampMs(v int64) {
	o.UpdateTimestampMs = &v
}

// GetAsks returns the Asks field value if set, zero value otherwise.
func (o *PredictionOrderBookByMarketIdResponse) GetAsks() [][]float64 {
	if o == nil || common.IsNil(o.Asks) {
		var ret [][]float64
		return ret
	}
	return o.Asks
}

// GetAsksOk returns a tuple with the Asks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PredictionOrderBookByMarketIdResponse) GetAsksOk() ([][]float64, bool) {
	if o == nil || common.IsNil(o.Asks) {
		return nil, false
	}
	return o.Asks, true
}

// HasAsks returns a boolean if a field has been set.
func (o *PredictionOrderBookByMarketIdResponse) HasAsks() bool {
	if o != nil && !common.IsNil(o.Asks) {
		return true
	}

	return false
}

// SetAsks gets a reference to the given [][]float64 and assigns it to the Asks field.
func (o *PredictionOrderBookByMarketIdResponse) SetAsks(v [][]float64) {
	o.Asks = v
}

// GetBids returns the Bids field value if set, zero value otherwise.
func (o *PredictionOrderBookByMarketIdResponse) GetBids() [][]float64 {
	if o == nil || common.IsNil(o.Bids) {
		var ret [][]float64
		return ret
	}
	return o.Bids
}

// GetBidsOk returns a tuple with the Bids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PredictionOrderBookByMarketIdResponse) GetBidsOk() ([][]float64, bool) {
	if o == nil || common.IsNil(o.Bids) {
		return nil, false
	}
	return o.Bids, true
}

// HasBids returns a boolean if a field has been set.
func (o *PredictionOrderBookByMarketIdResponse) HasBids() bool {
	if o != nil && !common.IsNil(o.Bids) {
		return true
	}

	return false
}

// SetBids gets a reference to the given [][]float64 and assigns it to the Bids field.
func (o *PredictionOrderBookByMarketIdResponse) SetBids(v [][]float64) {
	o.Bids = v
}

func (o PredictionOrderBookByMarketIdResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PredictionOrderBookByMarketIdResponse) ToMap() (map[string]interface{}, error) {
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

func (o *PredictionOrderBookByMarketIdResponse) UnmarshalJSON(data []byte) (err error) {
	varPredictionOrderBookByMarketIdResponse := _PredictionOrderBookByMarketIdResponse{}

	err = json.Unmarshal(data, &varPredictionOrderBookByMarketIdResponse)

	if err != nil {
		return err
	}

	*o = PredictionOrderBookByMarketIdResponse(varPredictionOrderBookByMarketIdResponse)

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

type NullablePredictionOrderBookByMarketIdResponse struct {
	value *PredictionOrderBookByMarketIdResponse
	isSet bool
}

func (v NullablePredictionOrderBookByMarketIdResponse) Get() *PredictionOrderBookByMarketIdResponse {
	return v.value
}

func (v *NullablePredictionOrderBookByMarketIdResponse) Set(val *PredictionOrderBookByMarketIdResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePredictionOrderBookByMarketIdResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePredictionOrderBookByMarketIdResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePredictionOrderBookByMarketIdResponse(val *PredictionOrderBookByMarketIdResponse) *NullablePredictionOrderBookByMarketIdResponse {
	return &NullablePredictionOrderBookByMarketIdResponse{value: val, isSet: true}
}

func (v NullablePredictionOrderBookByMarketIdResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePredictionOrderBookByMarketIdResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
