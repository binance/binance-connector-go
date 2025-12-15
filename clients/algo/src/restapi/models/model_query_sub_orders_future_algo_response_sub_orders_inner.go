/*
Binance Algo REST API

OpenAPI Specification for the Binance Algo REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QuerySubOrdersFutureAlgoResponseSubOrdersInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QuerySubOrdersFutureAlgoResponseSubOrdersInner{}

// QuerySubOrdersFutureAlgoResponseSubOrdersInner struct for QuerySubOrdersFutureAlgoResponseSubOrdersInner
type QuerySubOrdersFutureAlgoResponseSubOrdersInner struct {
	AlgoId               *int64  `json:"algoId,omitempty"`
	OrderId              *int64  `json:"orderId,omitempty"`
	OrderStatus          *string `json:"orderStatus,omitempty"`
	ExecutedQty          *string `json:"executedQty,omitempty"`
	ExecutedAmt          *string `json:"executedAmt,omitempty"`
	FeeAmt               *string `json:"feeAmt,omitempty"`
	FeeAsset             *string `json:"feeAsset,omitempty"`
	BookTime             *int64  `json:"bookTime,omitempty"`
	AvgPrice             *string `json:"avgPrice,omitempty"`
	Side                 *string `json:"side,omitempty"`
	Symbol               *string `json:"symbol,omitempty"`
	SubId                *int64  `json:"subId,omitempty"`
	TimeInForce          *string `json:"timeInForce,omitempty"`
	OrigQty              *string `json:"origQty,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuerySubOrdersFutureAlgoResponseSubOrdersInner QuerySubOrdersFutureAlgoResponseSubOrdersInner

// NewQuerySubOrdersFutureAlgoResponseSubOrdersInner instantiates a new QuerySubOrdersFutureAlgoResponseSubOrdersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuerySubOrdersFutureAlgoResponseSubOrdersInner() *QuerySubOrdersFutureAlgoResponseSubOrdersInner {
	this := QuerySubOrdersFutureAlgoResponseSubOrdersInner{}
	return &this
}

// NewQuerySubOrdersFutureAlgoResponseSubOrdersInnerWithDefaults instantiates a new QuerySubOrdersFutureAlgoResponseSubOrdersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuerySubOrdersFutureAlgoResponseSubOrdersInnerWithDefaults() *QuerySubOrdersFutureAlgoResponseSubOrdersInner {
	this := QuerySubOrdersFutureAlgoResponseSubOrdersInner{}
	return &this
}

// GetAlgoId returns the AlgoId field value if set, zero value otherwise.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) GetAlgoId() int64 {
	if o == nil || common.IsNil(o.AlgoId) {
		var ret int64
		return ret
	}
	return *o.AlgoId
}

// GetAlgoIdOk returns a tuple with the AlgoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) GetAlgoIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.AlgoId) {
		return nil, false
	}
	return o.AlgoId, true
}

// HasAlgoId returns a boolean if a field has been set.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) HasAlgoId() bool {
	if o != nil && !common.IsNil(o.AlgoId) {
		return true
	}

	return false
}

// SetAlgoId gets a reference to the given int64 and assigns it to the AlgoId field.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) SetAlgoId(v int64) {
	o.AlgoId = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetOrderStatus returns the OrderStatus field value if set, zero value otherwise.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) GetOrderStatus() string {
	if o == nil || common.IsNil(o.OrderStatus) {
		var ret string
		return ret
	}
	return *o.OrderStatus
}

// GetOrderStatusOk returns a tuple with the OrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) GetOrderStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderStatus) {
		return nil, false
	}
	return o.OrderStatus, true
}

// HasOrderStatus returns a boolean if a field has been set.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) HasOrderStatus() bool {
	if o != nil && !common.IsNil(o.OrderStatus) {
		return true
	}

	return false
}

// SetOrderStatus gets a reference to the given string and assigns it to the OrderStatus field.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) SetOrderStatus(v string) {
	o.OrderStatus = &v
}

// GetExecutedQty returns the ExecutedQty field value if set, zero value otherwise.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) GetExecutedQty() string {
	if o == nil || common.IsNil(o.ExecutedQty) {
		var ret string
		return ret
	}
	return *o.ExecutedQty
}

// GetExecutedQtyOk returns a tuple with the ExecutedQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) GetExecutedQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExecutedQty) {
		return nil, false
	}
	return o.ExecutedQty, true
}

// HasExecutedQty returns a boolean if a field has been set.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) HasExecutedQty() bool {
	if o != nil && !common.IsNil(o.ExecutedQty) {
		return true
	}

	return false
}

// SetExecutedQty gets a reference to the given string and assigns it to the ExecutedQty field.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) SetExecutedQty(v string) {
	o.ExecutedQty = &v
}

// GetExecutedAmt returns the ExecutedAmt field value if set, zero value otherwise.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) GetExecutedAmt() string {
	if o == nil || common.IsNil(o.ExecutedAmt) {
		var ret string
		return ret
	}
	return *o.ExecutedAmt
}

// GetExecutedAmtOk returns a tuple with the ExecutedAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) GetExecutedAmtOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExecutedAmt) {
		return nil, false
	}
	return o.ExecutedAmt, true
}

// HasExecutedAmt returns a boolean if a field has been set.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) HasExecutedAmt() bool {
	if o != nil && !common.IsNil(o.ExecutedAmt) {
		return true
	}

	return false
}

// SetExecutedAmt gets a reference to the given string and assigns it to the ExecutedAmt field.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) SetExecutedAmt(v string) {
	o.ExecutedAmt = &v
}

// GetFeeAmt returns the FeeAmt field value if set, zero value otherwise.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) GetFeeAmt() string {
	if o == nil || common.IsNil(o.FeeAmt) {
		var ret string
		return ret
	}
	return *o.FeeAmt
}

// GetFeeAmtOk returns a tuple with the FeeAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) GetFeeAmtOk() (*string, bool) {
	if o == nil || common.IsNil(o.FeeAmt) {
		return nil, false
	}
	return o.FeeAmt, true
}

// HasFeeAmt returns a boolean if a field has been set.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) HasFeeAmt() bool {
	if o != nil && !common.IsNil(o.FeeAmt) {
		return true
	}

	return false
}

// SetFeeAmt gets a reference to the given string and assigns it to the FeeAmt field.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) SetFeeAmt(v string) {
	o.FeeAmt = &v
}

// GetFeeAsset returns the FeeAsset field value if set, zero value otherwise.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) GetFeeAsset() string {
	if o == nil || common.IsNil(o.FeeAsset) {
		var ret string
		return ret
	}
	return *o.FeeAsset
}

// GetFeeAssetOk returns a tuple with the FeeAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) GetFeeAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.FeeAsset) {
		return nil, false
	}
	return o.FeeAsset, true
}

// HasFeeAsset returns a boolean if a field has been set.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) HasFeeAsset() bool {
	if o != nil && !common.IsNil(o.FeeAsset) {
		return true
	}

	return false
}

// SetFeeAsset gets a reference to the given string and assigns it to the FeeAsset field.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) SetFeeAsset(v string) {
	o.FeeAsset = &v
}

// GetBookTime returns the BookTime field value if set, zero value otherwise.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) GetBookTime() int64 {
	if o == nil || common.IsNil(o.BookTime) {
		var ret int64
		return ret
	}
	return *o.BookTime
}

// GetBookTimeOk returns a tuple with the BookTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) GetBookTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.BookTime) {
		return nil, false
	}
	return o.BookTime, true
}

// HasBookTime returns a boolean if a field has been set.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) HasBookTime() bool {
	if o != nil && !common.IsNil(o.BookTime) {
		return true
	}

	return false
}

// SetBookTime gets a reference to the given int64 and assigns it to the BookTime field.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) SetBookTime(v int64) {
	o.BookTime = &v
}

// GetAvgPrice returns the AvgPrice field value if set, zero value otherwise.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) GetAvgPrice() string {
	if o == nil || common.IsNil(o.AvgPrice) {
		var ret string
		return ret
	}
	return *o.AvgPrice
}

// GetAvgPriceOk returns a tuple with the AvgPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) GetAvgPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.AvgPrice) {
		return nil, false
	}
	return o.AvgPrice, true
}

// HasAvgPrice returns a boolean if a field has been set.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) HasAvgPrice() bool {
	if o != nil && !common.IsNil(o.AvgPrice) {
		return true
	}

	return false
}

// SetAvgPrice gets a reference to the given string and assigns it to the AvgPrice field.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) SetAvgPrice(v string) {
	o.AvgPrice = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) SetSide(v string) {
	o.Side = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetSubId returns the SubId field value if set, zero value otherwise.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) GetSubId() int64 {
	if o == nil || common.IsNil(o.SubId) {
		var ret int64
		return ret
	}
	return *o.SubId
}

// GetSubIdOk returns a tuple with the SubId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) GetSubIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.SubId) {
		return nil, false
	}
	return o.SubId, true
}

// HasSubId returns a boolean if a field has been set.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) HasSubId() bool {
	if o != nil && !common.IsNil(o.SubId) {
		return true
	}

	return false
}

// SetSubId gets a reference to the given int64 and assigns it to the SubId field.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) SetSubId(v int64) {
	o.SubId = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) GetTimeInForce() string {
	if o == nil || common.IsNil(o.TimeInForce) {
		var ret string
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) GetTimeInForceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) HasTimeInForce() bool {
	if o != nil && !common.IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given string and assigns it to the TimeInForce field.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) SetTimeInForce(v string) {
	o.TimeInForce = &v
}

// GetOrigQty returns the OrigQty field value if set, zero value otherwise.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) GetOrigQty() string {
	if o == nil || common.IsNil(o.OrigQty) {
		var ret string
		return ret
	}
	return *o.OrigQty
}

// GetOrigQtyOk returns a tuple with the OrigQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) GetOrigQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrigQty) {
		return nil, false
	}
	return o.OrigQty, true
}

// HasOrigQty returns a boolean if a field has been set.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) HasOrigQty() bool {
	if o != nil && !common.IsNil(o.OrigQty) {
		return true
	}

	return false
}

// SetOrigQty gets a reference to the given string and assigns it to the OrigQty field.
func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) SetOrigQty(v string) {
	o.OrigQty = &v
}

func (o QuerySubOrdersFutureAlgoResponseSubOrdersInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuerySubOrdersFutureAlgoResponseSubOrdersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AlgoId) {
		toSerialize["algoId"] = o.AlgoId
	}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.OrderStatus) {
		toSerialize["orderStatus"] = o.OrderStatus
	}
	if !common.IsNil(o.ExecutedQty) {
		toSerialize["executedQty"] = o.ExecutedQty
	}
	if !common.IsNil(o.ExecutedAmt) {
		toSerialize["executedAmt"] = o.ExecutedAmt
	}
	if !common.IsNil(o.FeeAmt) {
		toSerialize["feeAmt"] = o.FeeAmt
	}
	if !common.IsNil(o.FeeAsset) {
		toSerialize["feeAsset"] = o.FeeAsset
	}
	if !common.IsNil(o.BookTime) {
		toSerialize["bookTime"] = o.BookTime
	}
	if !common.IsNil(o.AvgPrice) {
		toSerialize["avgPrice"] = o.AvgPrice
	}
	if !common.IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.SubId) {
		toSerialize["subId"] = o.SubId
	}
	if !common.IsNil(o.TimeInForce) {
		toSerialize["timeInForce"] = o.TimeInForce
	}
	if !common.IsNil(o.OrigQty) {
		toSerialize["origQty"] = o.OrigQty
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuerySubOrdersFutureAlgoResponseSubOrdersInner) UnmarshalJSON(data []byte) (err error) {
	varQuerySubOrdersFutureAlgoResponseSubOrdersInner := _QuerySubOrdersFutureAlgoResponseSubOrdersInner{}

	err = json.Unmarshal(data, &varQuerySubOrdersFutureAlgoResponseSubOrdersInner)

	if err != nil {
		return err
	}

	*o = QuerySubOrdersFutureAlgoResponseSubOrdersInner(varQuerySubOrdersFutureAlgoResponseSubOrdersInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "algoId")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "orderStatus")
		delete(additionalProperties, "executedQty")
		delete(additionalProperties, "executedAmt")
		delete(additionalProperties, "feeAmt")
		delete(additionalProperties, "feeAsset")
		delete(additionalProperties, "bookTime")
		delete(additionalProperties, "avgPrice")
		delete(additionalProperties, "side")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "subId")
		delete(additionalProperties, "timeInForce")
		delete(additionalProperties, "origQty")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuerySubOrdersFutureAlgoResponseSubOrdersInner struct {
	value *QuerySubOrdersFutureAlgoResponseSubOrdersInner
	isSet bool
}

func (v NullableQuerySubOrdersFutureAlgoResponseSubOrdersInner) Get() *QuerySubOrdersFutureAlgoResponseSubOrdersInner {
	return v.value
}

func (v *NullableQuerySubOrdersFutureAlgoResponseSubOrdersInner) Set(val *QuerySubOrdersFutureAlgoResponseSubOrdersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQuerySubOrdersFutureAlgoResponseSubOrdersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQuerySubOrdersFutureAlgoResponseSubOrdersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuerySubOrdersFutureAlgoResponseSubOrdersInner(val *QuerySubOrdersFutureAlgoResponseSubOrdersInner) *NullableQuerySubOrdersFutureAlgoResponseSubOrdersInner {
	return &NullableQuerySubOrdersFutureAlgoResponseSubOrdersInner{value: val, isSet: true}
}

func (v NullableQuerySubOrdersFutureAlgoResponseSubOrdersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuerySubOrdersFutureAlgoResponseSubOrdersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
