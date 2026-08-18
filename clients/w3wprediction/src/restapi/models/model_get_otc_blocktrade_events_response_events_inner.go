/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetOtcBlocktradeEventsResponseEventsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetOtcBlocktradeEventsResponseEventsInner{}

// GetOtcBlocktradeEventsResponseEventsInner struct for GetOtcBlocktradeEventsResponseEventsInner
type GetOtcBlocktradeEventsResponseEventsInner struct {
	Event                *string `json:"event,omitempty"`
	BlocktradeId         *string `json:"blocktradeId,omitempty"`
	OrderId              *string `json:"orderId,omitempty"`
	MarketId             *int32  `json:"marketId,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	Fee                  *string `json:"fee,omitempty"`
	EffectiveFee         *string `json:"effectiveFee,omitempty"`
	Price                *string `json:"price,omitempty"`
	TransactionHash      *string `json:"transactionHash,omitempty"`
	Reason               *string `json:"reason,omitempty"`
	IsMaker              *bool   `json:"isMaker,omitempty"`
	Timestamp            *string `json:"timestamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetOtcBlocktradeEventsResponseEventsInner GetOtcBlocktradeEventsResponseEventsInner

// NewGetOtcBlocktradeEventsResponseEventsInner instantiates a new GetOtcBlocktradeEventsResponseEventsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOtcBlocktradeEventsResponseEventsInner() *GetOtcBlocktradeEventsResponseEventsInner {
	this := GetOtcBlocktradeEventsResponseEventsInner{}
	return &this
}

// NewGetOtcBlocktradeEventsResponseEventsInnerWithDefaults instantiates a new GetOtcBlocktradeEventsResponseEventsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOtcBlocktradeEventsResponseEventsInnerWithDefaults() *GetOtcBlocktradeEventsResponseEventsInner {
	this := GetOtcBlocktradeEventsResponseEventsInner{}
	return &this
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *GetOtcBlocktradeEventsResponseEventsInner) GetEvent() string {
	if o == nil || common.IsNil(o.Event) {
		var ret string
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOtcBlocktradeEventsResponseEventsInner) GetEventOk() (*string, bool) {
	if o == nil || common.IsNil(o.Event) {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *GetOtcBlocktradeEventsResponseEventsInner) HasEvent() bool {
	if o != nil && !common.IsNil(o.Event) {
		return true
	}

	return false
}

// SetEvent gets a reference to the given string and assigns it to the Event field.
func (o *GetOtcBlocktradeEventsResponseEventsInner) SetEvent(v string) {
	o.Event = &v
}

// GetBlocktradeId returns the BlocktradeId field value if set, zero value otherwise.
func (o *GetOtcBlocktradeEventsResponseEventsInner) GetBlocktradeId() string {
	if o == nil || common.IsNil(o.BlocktradeId) {
		var ret string
		return ret
	}
	return *o.BlocktradeId
}

// GetBlocktradeIdOk returns a tuple with the BlocktradeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOtcBlocktradeEventsResponseEventsInner) GetBlocktradeIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.BlocktradeId) {
		return nil, false
	}
	return o.BlocktradeId, true
}

// HasBlocktradeId returns a boolean if a field has been set.
func (o *GetOtcBlocktradeEventsResponseEventsInner) HasBlocktradeId() bool {
	if o != nil && !common.IsNil(o.BlocktradeId) {
		return true
	}

	return false
}

// SetBlocktradeId gets a reference to the given string and assigns it to the BlocktradeId field.
func (o *GetOtcBlocktradeEventsResponseEventsInner) SetBlocktradeId(v string) {
	o.BlocktradeId = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *GetOtcBlocktradeEventsResponseEventsInner) GetOrderId() string {
	if o == nil || common.IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOtcBlocktradeEventsResponseEventsInner) GetOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *GetOtcBlocktradeEventsResponseEventsInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *GetOtcBlocktradeEventsResponseEventsInner) SetOrderId(v string) {
	o.OrderId = &v
}

// GetMarketId returns the MarketId field value if set, zero value otherwise.
func (o *GetOtcBlocktradeEventsResponseEventsInner) GetMarketId() int32 {
	if o == nil || common.IsNil(o.MarketId) {
		var ret int32
		return ret
	}
	return *o.MarketId
}

// GetMarketIdOk returns a tuple with the MarketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOtcBlocktradeEventsResponseEventsInner) GetMarketIdOk() (*int32, bool) {
	if o == nil || common.IsNil(o.MarketId) {
		return nil, false
	}
	return o.MarketId, true
}

// HasMarketId returns a boolean if a field has been set.
func (o *GetOtcBlocktradeEventsResponseEventsInner) HasMarketId() bool {
	if o != nil && !common.IsNil(o.MarketId) {
		return true
	}

	return false
}

// SetMarketId gets a reference to the given int32 and assigns it to the MarketId field.
func (o *GetOtcBlocktradeEventsResponseEventsInner) SetMarketId(v int32) {
	o.MarketId = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetOtcBlocktradeEventsResponseEventsInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOtcBlocktradeEventsResponseEventsInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetOtcBlocktradeEventsResponseEventsInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *GetOtcBlocktradeEventsResponseEventsInner) SetAmount(v string) {
	o.Amount = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *GetOtcBlocktradeEventsResponseEventsInner) GetFee() string {
	if o == nil || common.IsNil(o.Fee) {
		var ret string
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOtcBlocktradeEventsResponseEventsInner) GetFeeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *GetOtcBlocktradeEventsResponseEventsInner) HasFee() bool {
	if o != nil && !common.IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given string and assigns it to the Fee field.
func (o *GetOtcBlocktradeEventsResponseEventsInner) SetFee(v string) {
	o.Fee = &v
}

// GetEffectiveFee returns the EffectiveFee field value if set, zero value otherwise.
func (o *GetOtcBlocktradeEventsResponseEventsInner) GetEffectiveFee() string {
	if o == nil || common.IsNil(o.EffectiveFee) {
		var ret string
		return ret
	}
	return *o.EffectiveFee
}

// GetEffectiveFeeOk returns a tuple with the EffectiveFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOtcBlocktradeEventsResponseEventsInner) GetEffectiveFeeOk() (*string, bool) {
	if o == nil || common.IsNil(o.EffectiveFee) {
		return nil, false
	}
	return o.EffectiveFee, true
}

// HasEffectiveFee returns a boolean if a field has been set.
func (o *GetOtcBlocktradeEventsResponseEventsInner) HasEffectiveFee() bool {
	if o != nil && !common.IsNil(o.EffectiveFee) {
		return true
	}

	return false
}

// SetEffectiveFee gets a reference to the given string and assigns it to the EffectiveFee field.
func (o *GetOtcBlocktradeEventsResponseEventsInner) SetEffectiveFee(v string) {
	o.EffectiveFee = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *GetOtcBlocktradeEventsResponseEventsInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOtcBlocktradeEventsResponseEventsInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *GetOtcBlocktradeEventsResponseEventsInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *GetOtcBlocktradeEventsResponseEventsInner) SetPrice(v string) {
	o.Price = &v
}

// GetTransactionHash returns the TransactionHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetOtcBlocktradeEventsResponseEventsInner) GetTransactionHash() string {
	if o == nil || common.IsNil(o.TransactionHash) {
		var ret string
		return ret
	}
	return *o.TransactionHash
}

// GetTransactionHashOk returns a tuple with the TransactionHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetOtcBlocktradeEventsResponseEventsInner) GetTransactionHashOk() (*string, bool) {
	if o == nil || common.IsNil(o.TransactionHash) {
		return nil, false
	}
	return o.TransactionHash, true
}

// HasTransactionHash returns a boolean if a field has been set.
func (o *GetOtcBlocktradeEventsResponseEventsInner) HasTransactionHash() bool {
	if o != nil && !common.IsNil(o.TransactionHash) {
		return true
	}

	return false
}

// SetTransactionHash gets a reference to the given NullableString and assigns it to the TransactionHash field.
func (o *GetOtcBlocktradeEventsResponseEventsInner) SetTransactionHash(v string) {
	o.TransactionHash = &v
}

// GetReason returns the Reason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetOtcBlocktradeEventsResponseEventsInner) GetReason() string {
	if o == nil || common.IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetOtcBlocktradeEventsResponseEventsInner) GetReasonOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *GetOtcBlocktradeEventsResponseEventsInner) HasReason() bool {
	if o != nil && !common.IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given NullableString and assigns it to the Reason field.
func (o *GetOtcBlocktradeEventsResponseEventsInner) SetReason(v string) {
	o.Reason = &v
}

// GetIsMaker returns the IsMaker field value if set, zero value otherwise.
func (o *GetOtcBlocktradeEventsResponseEventsInner) GetIsMaker() bool {
	if o == nil || common.IsNil(o.IsMaker) {
		var ret bool
		return ret
	}
	return *o.IsMaker
}

// GetIsMakerOk returns a tuple with the IsMaker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOtcBlocktradeEventsResponseEventsInner) GetIsMakerOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsMaker) {
		return nil, false
	}
	return o.IsMaker, true
}

// HasIsMaker returns a boolean if a field has been set.
func (o *GetOtcBlocktradeEventsResponseEventsInner) HasIsMaker() bool {
	if o != nil && !common.IsNil(o.IsMaker) {
		return true
	}

	return false
}

// SetIsMaker gets a reference to the given bool and assigns it to the IsMaker field.
func (o *GetOtcBlocktradeEventsResponseEventsInner) SetIsMaker(v bool) {
	o.IsMaker = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *GetOtcBlocktradeEventsResponseEventsInner) GetTimestamp() string {
	if o == nil || common.IsNil(o.Timestamp) {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOtcBlocktradeEventsResponseEventsInner) GetTimestampOk() (*string, bool) {
	if o == nil || common.IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *GetOtcBlocktradeEventsResponseEventsInner) HasTimestamp() bool {
	if o != nil && !common.IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *GetOtcBlocktradeEventsResponseEventsInner) SetTimestamp(v string) {
	o.Timestamp = &v
}

func (o GetOtcBlocktradeEventsResponseEventsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOtcBlocktradeEventsResponseEventsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Event) {
		toSerialize["event"] = o.Event
	}
	if !common.IsNil(o.BlocktradeId) {
		toSerialize["blocktradeId"] = o.BlocktradeId
	}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.MarketId) {
		toSerialize["marketId"] = o.MarketId
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	if !common.IsNil(o.EffectiveFee) {
		toSerialize["effectiveFee"] = o.EffectiveFee
	}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.TransactionHash) {
		toSerialize["transactionHash"] = o.TransactionHash
	}
	if !common.IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !common.IsNil(o.IsMaker) {
		toSerialize["isMaker"] = o.IsMaker
	}
	if !common.IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetOtcBlocktradeEventsResponseEventsInner) UnmarshalJSON(data []byte) (err error) {
	varGetOtcBlocktradeEventsResponseEventsInner := _GetOtcBlocktradeEventsResponseEventsInner{}

	err = json.Unmarshal(data, &varGetOtcBlocktradeEventsResponseEventsInner)

	if err != nil {
		return err
	}

	*o = GetOtcBlocktradeEventsResponseEventsInner(varGetOtcBlocktradeEventsResponseEventsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "event")
		delete(additionalProperties, "blocktradeId")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "marketId")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "fee")
		delete(additionalProperties, "effectiveFee")
		delete(additionalProperties, "price")
		delete(additionalProperties, "transactionHash")
		delete(additionalProperties, "reason")
		delete(additionalProperties, "isMaker")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetOtcBlocktradeEventsResponseEventsInner struct {
	value *GetOtcBlocktradeEventsResponseEventsInner
	isSet bool
}

func (v NullableGetOtcBlocktradeEventsResponseEventsInner) Get() *GetOtcBlocktradeEventsResponseEventsInner {
	return v.value
}

func (v *NullableGetOtcBlocktradeEventsResponseEventsInner) Set(val *GetOtcBlocktradeEventsResponseEventsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOtcBlocktradeEventsResponseEventsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOtcBlocktradeEventsResponseEventsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOtcBlocktradeEventsResponseEventsInner(val *GetOtcBlocktradeEventsResponseEventsInner) *NullableGetOtcBlocktradeEventsResponseEventsInner {
	return &NullableGetOtcBlocktradeEventsResponseEventsInner{value: val, isSet: true}
}

func (v NullableGetOtcBlocktradeEventsResponseEventsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOtcBlocktradeEventsResponseEventsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
