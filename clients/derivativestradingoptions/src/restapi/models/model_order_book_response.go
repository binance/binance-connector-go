/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OrderBookResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderBookResponse{}

// OrderBookResponse struct for OrderBookResponse
type OrderBookResponse struct {
	T                    *int64                      `json:"T,omitempty"`
	U                    *int64                      `json:"u,omitempty"`
	Bids                 []OrderBookResponseBidsItem `json:"bids,omitempty"`
	Asks                 []OrderBookResponseAsksItem `json:"asks,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderBookResponse OrderBookResponse

// NewOrderBookResponse instantiates a new OrderBookResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderBookResponse() *OrderBookResponse {
	this := OrderBookResponse{}
	return &this
}

// NewOrderBookResponseWithDefaults instantiates a new OrderBookResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderBookResponseWithDefaults() *OrderBookResponse {
	this := OrderBookResponse{}
	return &this
}

// GetT returns the T field value if set, zero value otherwise.
func (o *OrderBookResponse) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderBookResponse) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *OrderBookResponse) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *OrderBookResponse) SetT(v int64) {
	o.T = &v
}

// GetU returns the U field value if set, zero value otherwise.
func (o *OrderBookResponse) GetU() int64 {
	if o == nil || common.IsNil(o.U) {
		var ret int64
		return ret
	}
	return *o.U
}

// GetUOk returns a tuple with the U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderBookResponse) GetUOk() (*int64, bool) {
	if o == nil || common.IsNil(o.U) {
		return nil, false
	}
	return o.U, true
}

// HasU returns a boolean if a field has been set.
func (o *OrderBookResponse) HasU() bool {
	if o != nil && !common.IsNil(o.U) {
		return true
	}

	return false
}

// SetU gets a reference to the given int64 and assigns it to the U field.
func (o *OrderBookResponse) SetU(v int64) {
	o.U = &v
}

// GetBids returns the Bids field value if set, zero value otherwise.
func (o *OrderBookResponse) GetBids() []OrderBookResponseBidsItem {
	if o == nil || common.IsNil(o.Bids) {
		var ret []OrderBookResponseBidsItem
		return ret
	}
	return o.Bids
}

// GetBidsOk returns a tuple with the Bids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderBookResponse) GetBidsOk() ([]OrderBookResponseBidsItem, bool) {
	if o == nil || common.IsNil(o.Bids) {
		return nil, false
	}
	return o.Bids, true
}

// HasBids returns a boolean if a field has been set.
func (o *OrderBookResponse) HasBids() bool {
	if o != nil && !common.IsNil(o.Bids) {
		return true
	}

	return false
}

// SetBids gets a reference to the given []OrderBookResponseBidsItem and assigns it to the Bids field.
func (o *OrderBookResponse) SetBids(v []OrderBookResponseBidsItem) {
	o.Bids = v
}

// GetAsks returns the Asks field value if set, zero value otherwise.
func (o *OrderBookResponse) GetAsks() []OrderBookResponseAsksItem {
	if o == nil || common.IsNil(o.Asks) {
		var ret []OrderBookResponseAsksItem
		return ret
	}
	return o.Asks
}

// GetAsksOk returns a tuple with the Asks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderBookResponse) GetAsksOk() ([]OrderBookResponseAsksItem, bool) {
	if o == nil || common.IsNil(o.Asks) {
		return nil, false
	}
	return o.Asks, true
}

// HasAsks returns a boolean if a field has been set.
func (o *OrderBookResponse) HasAsks() bool {
	if o != nil && !common.IsNil(o.Asks) {
		return true
	}

	return false
}

// SetAsks gets a reference to the given []OrderBookResponseAsksItem and assigns it to the Asks field.
func (o *OrderBookResponse) SetAsks(v []OrderBookResponseAsksItem) {
	o.Asks = v
}

func (o OrderBookResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderBookResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}
	if !common.IsNil(o.U) {
		toSerialize["u"] = o.U
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

func (o *OrderBookResponse) UnmarshalJSON(data []byte) (err error) {
	varOrderBookResponse := _OrderBookResponse{}

	err = json.Unmarshal(data, &varOrderBookResponse)

	if err != nil {
		return err
	}

	*o = OrderBookResponse(varOrderBookResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "T")
		delete(additionalProperties, "u")
		delete(additionalProperties, "bids")
		delete(additionalProperties, "asks")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrderBookResponse struct {
	value *OrderBookResponse
	isSet bool
}

func (v NullableOrderBookResponse) Get() *OrderBookResponse {
	return v.value
}

func (v *NullableOrderBookResponse) Set(val *OrderBookResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderBookResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderBookResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderBookResponse(val *OrderBookResponse) *NullableOrderBookResponse {
	return &NullableOrderBookResponse{value: val, isSet: true}
}

func (v NullableOrderBookResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderBookResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
