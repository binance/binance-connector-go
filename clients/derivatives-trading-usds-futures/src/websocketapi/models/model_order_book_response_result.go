/*
Binance Derivatives Trading USDS Futures WebSocket API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OrderBookResponseResult type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderBookResponseResult{}

// OrderBookResponseResult struct for OrderBookResponseResult
type OrderBookResponseResult struct {
	LastUpdateId         *int64                            `json:"lastUpdateId,omitempty"`
	E                    *int64                            `json:"E,omitempty"`
	T                    *int64                            `json:"T,omitempty"`
	Bids                 []OrderBookResponseResultBidsItem `json:"bids,omitempty"`
	Asks                 []OrderBookResponseResultAsksItem `json:"asks,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderBookResponseResult OrderBookResponseResult

// NewOrderBookResponseResult instantiates a new OrderBookResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderBookResponseResult() *OrderBookResponseResult {
	this := OrderBookResponseResult{}
	return &this
}

// NewOrderBookResponseResultWithDefaults instantiates a new OrderBookResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderBookResponseResultWithDefaults() *OrderBookResponseResult {
	this := OrderBookResponseResult{}
	return &this
}

// GetLastUpdateId returns the LastUpdateId field value if set, zero value otherwise.
func (o *OrderBookResponseResult) GetLastUpdateId() int64 {
	if o == nil || common.IsNil(o.LastUpdateId) {
		var ret int64
		return ret
	}
	return *o.LastUpdateId
}

// GetLastUpdateIdOk returns a tuple with the LastUpdateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderBookResponseResult) GetLastUpdateIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.LastUpdateId) {
		return nil, false
	}
	return o.LastUpdateId, true
}

// HasLastUpdateId returns a boolean if a field has been set.
func (o *OrderBookResponseResult) HasLastUpdateId() bool {
	if o != nil && !common.IsNil(o.LastUpdateId) {
		return true
	}

	return false
}

// SetLastUpdateId gets a reference to the given int64 and assigns it to the LastUpdateId field.
func (o *OrderBookResponseResult) SetLastUpdateId(v int64) {
	o.LastUpdateId = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *OrderBookResponseResult) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderBookResponseResult) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *OrderBookResponseResult) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *OrderBookResponseResult) SetE(v int64) {
	o.E = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *OrderBookResponseResult) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderBookResponseResult) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *OrderBookResponseResult) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *OrderBookResponseResult) SetT(v int64) {
	o.T = &v
}

// GetBids returns the Bids field value if set, zero value otherwise.
func (o *OrderBookResponseResult) GetBids() []OrderBookResponseResultBidsItem {
	if o == nil || common.IsNil(o.Bids) {
		var ret []OrderBookResponseResultBidsItem
		return ret
	}
	return o.Bids
}

// GetBidsOk returns a tuple with the Bids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderBookResponseResult) GetBidsOk() ([]OrderBookResponseResultBidsItem, bool) {
	if o == nil || common.IsNil(o.Bids) {
		return nil, false
	}
	return o.Bids, true
}

// HasBids returns a boolean if a field has been set.
func (o *OrderBookResponseResult) HasBids() bool {
	if o != nil && !common.IsNil(o.Bids) {
		return true
	}

	return false
}

// SetBids gets a reference to the given []OrderBookResponseResultBidsItem and assigns it to the Bids field.
func (o *OrderBookResponseResult) SetBids(v []OrderBookResponseResultBidsItem) {
	o.Bids = v
}

// GetAsks returns the Asks field value if set, zero value otherwise.
func (o *OrderBookResponseResult) GetAsks() []OrderBookResponseResultAsksItem {
	if o == nil || common.IsNil(o.Asks) {
		var ret []OrderBookResponseResultAsksItem
		return ret
	}
	return o.Asks
}

// GetAsksOk returns a tuple with the Asks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderBookResponseResult) GetAsksOk() ([]OrderBookResponseResultAsksItem, bool) {
	if o == nil || common.IsNil(o.Asks) {
		return nil, false
	}
	return o.Asks, true
}

// HasAsks returns a boolean if a field has been set.
func (o *OrderBookResponseResult) HasAsks() bool {
	if o != nil && !common.IsNil(o.Asks) {
		return true
	}

	return false
}

// SetAsks gets a reference to the given []OrderBookResponseResultAsksItem and assigns it to the Asks field.
func (o *OrderBookResponseResult) SetAsks(v []OrderBookResponseResultAsksItem) {
	o.Asks = v
}

func (o OrderBookResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderBookResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.LastUpdateId) {
		toSerialize["lastUpdateId"] = o.LastUpdateId
	}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}
	if !common.IsNil(o.Bids) {
		toSerialize["bids"] = o.Bids
	}
	if !common.IsNil(o.Asks) {
		toSerialize["asks"] = o.Asks
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderBookResponseResult) UnmarshalJSON(data []byte) (err error) {
	varOrderBookResponseResult := _OrderBookResponseResult{}

	err = json.Unmarshal(data, &varOrderBookResponseResult)

	if err != nil {
		return err
	}

	*o = OrderBookResponseResult(varOrderBookResponseResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "lastUpdateId")
		delete(additionalProperties, "E")
		delete(additionalProperties, "T")
		delete(additionalProperties, "bids")
		delete(additionalProperties, "asks")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrderBookResponseResult struct {
	value *OrderBookResponseResult
	isSet bool
}

func (v NullableOrderBookResponseResult) Get() *OrderBookResponseResult {
	return v.value
}

func (v *NullableOrderBookResponseResult) Set(val *OrderBookResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderBookResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderBookResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderBookResponseResult(val *OrderBookResponseResult) *NullableOrderBookResponseResult {
	return &NullableOrderBookResponseResult{value: val, isSet: true}
}

func (v NullableOrderBookResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderBookResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
