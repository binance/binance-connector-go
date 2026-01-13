/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetForceLiquidationRecordResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetForceLiquidationRecordResponseRowsInner{}

// GetForceLiquidationRecordResponseRowsInner struct for GetForceLiquidationRecordResponseRowsInner
type GetForceLiquidationRecordResponseRowsInner struct {
	AvgPrice             *string `json:"avgPrice,omitempty"`
	ExecutedQty          *string `json:"executedQty,omitempty"`
	OrderId              *int64  `json:"orderId,omitempty"`
	Price                *string `json:"price,omitempty"`
	Qty                  *string `json:"qty,omitempty"`
	Side                 *string `json:"side,omitempty"`
	Symbol               *string `json:"symbol,omitempty"`
	TimeInForce          *string `json:"timeInForce,omitempty"`
	IsIsolated           *bool   `json:"isIsolated,omitempty"`
	UpdatedTime          *int64  `json:"updatedTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetForceLiquidationRecordResponseRowsInner GetForceLiquidationRecordResponseRowsInner

// NewGetForceLiquidationRecordResponseRowsInner instantiates a new GetForceLiquidationRecordResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetForceLiquidationRecordResponseRowsInner() *GetForceLiquidationRecordResponseRowsInner {
	this := GetForceLiquidationRecordResponseRowsInner{}
	return &this
}

// NewGetForceLiquidationRecordResponseRowsInnerWithDefaults instantiates a new GetForceLiquidationRecordResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetForceLiquidationRecordResponseRowsInnerWithDefaults() *GetForceLiquidationRecordResponseRowsInner {
	this := GetForceLiquidationRecordResponseRowsInner{}
	return &this
}

// GetAvgPrice returns the AvgPrice field value if set, zero value otherwise.
func (o *GetForceLiquidationRecordResponseRowsInner) GetAvgPrice() string {
	if o == nil || common.IsNil(o.AvgPrice) {
		var ret string
		return ret
	}
	return *o.AvgPrice
}

// GetAvgPriceOk returns a tuple with the AvgPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetForceLiquidationRecordResponseRowsInner) GetAvgPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.AvgPrice) {
		return nil, false
	}
	return o.AvgPrice, true
}

// HasAvgPrice returns a boolean if a field has been set.
func (o *GetForceLiquidationRecordResponseRowsInner) HasAvgPrice() bool {
	if o != nil && !common.IsNil(o.AvgPrice) {
		return true
	}

	return false
}

// SetAvgPrice gets a reference to the given string and assigns it to the AvgPrice field.
func (o *GetForceLiquidationRecordResponseRowsInner) SetAvgPrice(v string) {
	o.AvgPrice = &v
}

// GetExecutedQty returns the ExecutedQty field value if set, zero value otherwise.
func (o *GetForceLiquidationRecordResponseRowsInner) GetExecutedQty() string {
	if o == nil || common.IsNil(o.ExecutedQty) {
		var ret string
		return ret
	}
	return *o.ExecutedQty
}

// GetExecutedQtyOk returns a tuple with the ExecutedQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetForceLiquidationRecordResponseRowsInner) GetExecutedQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExecutedQty) {
		return nil, false
	}
	return o.ExecutedQty, true
}

// HasExecutedQty returns a boolean if a field has been set.
func (o *GetForceLiquidationRecordResponseRowsInner) HasExecutedQty() bool {
	if o != nil && !common.IsNil(o.ExecutedQty) {
		return true
	}

	return false
}

// SetExecutedQty gets a reference to the given string and assigns it to the ExecutedQty field.
func (o *GetForceLiquidationRecordResponseRowsInner) SetExecutedQty(v string) {
	o.ExecutedQty = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *GetForceLiquidationRecordResponseRowsInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetForceLiquidationRecordResponseRowsInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *GetForceLiquidationRecordResponseRowsInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *GetForceLiquidationRecordResponseRowsInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *GetForceLiquidationRecordResponseRowsInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetForceLiquidationRecordResponseRowsInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *GetForceLiquidationRecordResponseRowsInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *GetForceLiquidationRecordResponseRowsInner) SetPrice(v string) {
	o.Price = &v
}

// GetQty returns the Qty field value if set, zero value otherwise.
func (o *GetForceLiquidationRecordResponseRowsInner) GetQty() string {
	if o == nil || common.IsNil(o.Qty) {
		var ret string
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetForceLiquidationRecordResponseRowsInner) GetQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Qty) {
		return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *GetForceLiquidationRecordResponseRowsInner) HasQty() bool {
	if o != nil && !common.IsNil(o.Qty) {
		return true
	}

	return false
}

// SetQty gets a reference to the given string and assigns it to the Qty field.
func (o *GetForceLiquidationRecordResponseRowsInner) SetQty(v string) {
	o.Qty = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *GetForceLiquidationRecordResponseRowsInner) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetForceLiquidationRecordResponseRowsInner) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *GetForceLiquidationRecordResponseRowsInner) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *GetForceLiquidationRecordResponseRowsInner) SetSide(v string) {
	o.Side = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetForceLiquidationRecordResponseRowsInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetForceLiquidationRecordResponseRowsInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetForceLiquidationRecordResponseRowsInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetForceLiquidationRecordResponseRowsInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *GetForceLiquidationRecordResponseRowsInner) GetTimeInForce() string {
	if o == nil || common.IsNil(o.TimeInForce) {
		var ret string
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetForceLiquidationRecordResponseRowsInner) GetTimeInForceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *GetForceLiquidationRecordResponseRowsInner) HasTimeInForce() bool {
	if o != nil && !common.IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given string and assigns it to the TimeInForce field.
func (o *GetForceLiquidationRecordResponseRowsInner) SetTimeInForce(v string) {
	o.TimeInForce = &v
}

// GetIsIsolated returns the IsIsolated field value if set, zero value otherwise.
func (o *GetForceLiquidationRecordResponseRowsInner) GetIsIsolated() bool {
	if o == nil || common.IsNil(o.IsIsolated) {
		var ret bool
		return ret
	}
	return *o.IsIsolated
}

// GetIsIsolatedOk returns a tuple with the IsIsolated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetForceLiquidationRecordResponseRowsInner) GetIsIsolatedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsIsolated) {
		return nil, false
	}
	return o.IsIsolated, true
}

// HasIsIsolated returns a boolean if a field has been set.
func (o *GetForceLiquidationRecordResponseRowsInner) HasIsIsolated() bool {
	if o != nil && !common.IsNil(o.IsIsolated) {
		return true
	}

	return false
}

// SetIsIsolated gets a reference to the given bool and assigns it to the IsIsolated field.
func (o *GetForceLiquidationRecordResponseRowsInner) SetIsIsolated(v bool) {
	o.IsIsolated = &v
}

// GetUpdatedTime returns the UpdatedTime field value if set, zero value otherwise.
func (o *GetForceLiquidationRecordResponseRowsInner) GetUpdatedTime() int64 {
	if o == nil || common.IsNil(o.UpdatedTime) {
		var ret int64
		return ret
	}
	return *o.UpdatedTime
}

// GetUpdatedTimeOk returns a tuple with the UpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetForceLiquidationRecordResponseRowsInner) GetUpdatedTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdatedTime) {
		return nil, false
	}
	return o.UpdatedTime, true
}

// HasUpdatedTime returns a boolean if a field has been set.
func (o *GetForceLiquidationRecordResponseRowsInner) HasUpdatedTime() bool {
	if o != nil && !common.IsNil(o.UpdatedTime) {
		return true
	}

	return false
}

// SetUpdatedTime gets a reference to the given int64 and assigns it to the UpdatedTime field.
func (o *GetForceLiquidationRecordResponseRowsInner) SetUpdatedTime(v int64) {
	o.UpdatedTime = &v
}

func (o GetForceLiquidationRecordResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetForceLiquidationRecordResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AvgPrice) {
		toSerialize["avgPrice"] = o.AvgPrice
	}
	if !common.IsNil(o.ExecutedQty) {
		toSerialize["executedQty"] = o.ExecutedQty
	}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.Qty) {
		toSerialize["qty"] = o.Qty
	}
	if !common.IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.TimeInForce) {
		toSerialize["timeInForce"] = o.TimeInForce
	}
	if !common.IsNil(o.IsIsolated) {
		toSerialize["isIsolated"] = o.IsIsolated
	}
	if !common.IsNil(o.UpdatedTime) {
		toSerialize["updatedTime"] = o.UpdatedTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetForceLiquidationRecordResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetForceLiquidationRecordResponseRowsInner := _GetForceLiquidationRecordResponseRowsInner{}

	err = json.Unmarshal(data, &varGetForceLiquidationRecordResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetForceLiquidationRecordResponseRowsInner(varGetForceLiquidationRecordResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "avgPrice")
		delete(additionalProperties, "executedQty")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "price")
		delete(additionalProperties, "qty")
		delete(additionalProperties, "side")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "timeInForce")
		delete(additionalProperties, "isIsolated")
		delete(additionalProperties, "updatedTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetForceLiquidationRecordResponseRowsInner struct {
	value *GetForceLiquidationRecordResponseRowsInner
	isSet bool
}

func (v NullableGetForceLiquidationRecordResponseRowsInner) Get() *GetForceLiquidationRecordResponseRowsInner {
	return v.value
}

func (v *NullableGetForceLiquidationRecordResponseRowsInner) Set(val *GetForceLiquidationRecordResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetForceLiquidationRecordResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetForceLiquidationRecordResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetForceLiquidationRecordResponseRowsInner(val *GetForceLiquidationRecordResponseRowsInner) *NullableGetForceLiquidationRecordResponseRowsInner {
	return &NullableGetForceLiquidationRecordResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetForceLiquidationRecordResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetForceLiquidationRecordResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
