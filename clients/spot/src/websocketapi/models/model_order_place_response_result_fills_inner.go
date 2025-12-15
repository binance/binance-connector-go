/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OrderPlaceResponseResultFillsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderPlaceResponseResultFillsInner{}

// OrderPlaceResponseResultFillsInner struct for OrderPlaceResponseResultFillsInner
type OrderPlaceResponseResultFillsInner struct {
	Price                *string `json:"price,omitempty"`
	Qty                  *string `json:"qty,omitempty"`
	Commission           *string `json:"commission,omitempty"`
	CommissionAsset      *string `json:"commissionAsset,omitempty"`
	TradeId              *int64  `json:"tradeId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderPlaceResponseResultFillsInner OrderPlaceResponseResultFillsInner

// NewOrderPlaceResponseResultFillsInner instantiates a new OrderPlaceResponseResultFillsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderPlaceResponseResultFillsInner() *OrderPlaceResponseResultFillsInner {
	this := OrderPlaceResponseResultFillsInner{}
	return &this
}

// NewOrderPlaceResponseResultFillsInnerWithDefaults instantiates a new OrderPlaceResponseResultFillsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderPlaceResponseResultFillsInnerWithDefaults() *OrderPlaceResponseResultFillsInner {
	this := OrderPlaceResponseResultFillsInner{}
	return &this
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *OrderPlaceResponseResultFillsInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderPlaceResponseResultFillsInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *OrderPlaceResponseResultFillsInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *OrderPlaceResponseResultFillsInner) SetPrice(v string) {
	o.Price = &v
}

// GetQty returns the Qty field value if set, zero value otherwise.
func (o *OrderPlaceResponseResultFillsInner) GetQty() string {
	if o == nil || common.IsNil(o.Qty) {
		var ret string
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderPlaceResponseResultFillsInner) GetQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Qty) {
		return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *OrderPlaceResponseResultFillsInner) HasQty() bool {
	if o != nil && !common.IsNil(o.Qty) {
		return true
	}

	return false
}

// SetQty gets a reference to the given string and assigns it to the Qty field.
func (o *OrderPlaceResponseResultFillsInner) SetQty(v string) {
	o.Qty = &v
}

// GetCommission returns the Commission field value if set, zero value otherwise.
func (o *OrderPlaceResponseResultFillsInner) GetCommission() string {
	if o == nil || common.IsNil(o.Commission) {
		var ret string
		return ret
	}
	return *o.Commission
}

// GetCommissionOk returns a tuple with the Commission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderPlaceResponseResultFillsInner) GetCommissionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Commission) {
		return nil, false
	}
	return o.Commission, true
}

// HasCommission returns a boolean if a field has been set.
func (o *OrderPlaceResponseResultFillsInner) HasCommission() bool {
	if o != nil && !common.IsNil(o.Commission) {
		return true
	}

	return false
}

// SetCommission gets a reference to the given string and assigns it to the Commission field.
func (o *OrderPlaceResponseResultFillsInner) SetCommission(v string) {
	o.Commission = &v
}

// GetCommissionAsset returns the CommissionAsset field value if set, zero value otherwise.
func (o *OrderPlaceResponseResultFillsInner) GetCommissionAsset() string {
	if o == nil || common.IsNil(o.CommissionAsset) {
		var ret string
		return ret
	}
	return *o.CommissionAsset
}

// GetCommissionAssetOk returns a tuple with the CommissionAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderPlaceResponseResultFillsInner) GetCommissionAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.CommissionAsset) {
		return nil, false
	}
	return o.CommissionAsset, true
}

// HasCommissionAsset returns a boolean if a field has been set.
func (o *OrderPlaceResponseResultFillsInner) HasCommissionAsset() bool {
	if o != nil && !common.IsNil(o.CommissionAsset) {
		return true
	}

	return false
}

// SetCommissionAsset gets a reference to the given string and assigns it to the CommissionAsset field.
func (o *OrderPlaceResponseResultFillsInner) SetCommissionAsset(v string) {
	o.CommissionAsset = &v
}

// GetTradeId returns the TradeId field value if set, zero value otherwise.
func (o *OrderPlaceResponseResultFillsInner) GetTradeId() int64 {
	if o == nil || common.IsNil(o.TradeId) {
		var ret int64
		return ret
	}
	return *o.TradeId
}

// GetTradeIdOk returns a tuple with the TradeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderPlaceResponseResultFillsInner) GetTradeIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TradeId) {
		return nil, false
	}
	return o.TradeId, true
}

// HasTradeId returns a boolean if a field has been set.
func (o *OrderPlaceResponseResultFillsInner) HasTradeId() bool {
	if o != nil && !common.IsNil(o.TradeId) {
		return true
	}

	return false
}

// SetTradeId gets a reference to the given int64 and assigns it to the TradeId field.
func (o *OrderPlaceResponseResultFillsInner) SetTradeId(v int64) {
	o.TradeId = &v
}

func (o OrderPlaceResponseResultFillsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderPlaceResponseResultFillsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.Qty) {
		toSerialize["qty"] = o.Qty
	}
	if !common.IsNil(o.Commission) {
		toSerialize["commission"] = o.Commission
	}
	if !common.IsNil(o.CommissionAsset) {
		toSerialize["commissionAsset"] = o.CommissionAsset
	}
	if !common.IsNil(o.TradeId) {
		toSerialize["tradeId"] = o.TradeId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderPlaceResponseResultFillsInner) UnmarshalJSON(data []byte) (err error) {
	varOrderPlaceResponseResultFillsInner := _OrderPlaceResponseResultFillsInner{}

	err = json.Unmarshal(data, &varOrderPlaceResponseResultFillsInner)

	if err != nil {
		return err
	}

	*o = OrderPlaceResponseResultFillsInner(varOrderPlaceResponseResultFillsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "price")
		delete(additionalProperties, "qty")
		delete(additionalProperties, "commission")
		delete(additionalProperties, "commissionAsset")
		delete(additionalProperties, "tradeId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrderPlaceResponseResultFillsInner struct {
	value *OrderPlaceResponseResultFillsInner
	isSet bool
}

func (v NullableOrderPlaceResponseResultFillsInner) Get() *OrderPlaceResponseResultFillsInner {
	return v.value
}

func (v *NullableOrderPlaceResponseResultFillsInner) Set(val *OrderPlaceResponseResultFillsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderPlaceResponseResultFillsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderPlaceResponseResultFillsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderPlaceResponseResultFillsInner(val *OrderPlaceResponseResultFillsInner) *NullableOrderPlaceResponseResultFillsInner {
	return &NullableOrderPlaceResponseResultFillsInner{value: val, isSet: true}
}

func (v NullableOrderPlaceResponseResultFillsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderPlaceResponseResultFillsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
