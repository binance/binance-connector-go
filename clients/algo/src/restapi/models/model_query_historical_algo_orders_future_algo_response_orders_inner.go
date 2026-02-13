/*
Binance Algo REST API

OpenAPI Specification for the Binance Algo REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner{}

// QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner struct for QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner
type QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner struct {
	AlgoId               *int64  `json:"algoId,omitempty"`
	Symbol               *string `json:"symbol,omitempty"`
	Side                 *string `json:"side,omitempty"`
	PositionSide         *string `json:"positionSide,omitempty"`
	TotalQty             *string `json:"totalQty,omitempty"`
	ExecutedQty          *string `json:"executedQty,omitempty"`
	ExecutedAmt          *string `json:"executedAmt,omitempty"`
	AvgPrice             *string `json:"avgPrice,omitempty"`
	ClientAlgoId         *string `json:"clientAlgoId,omitempty"`
	BookTime             *int64  `json:"bookTime,omitempty"`
	EndTime              *int64  `json:"endTime,omitempty"`
	AlgoStatus           *string `json:"algoStatus,omitempty"`
	AlgoType             *string `json:"algoType,omitempty"`
	Urgency              *string `json:"urgency,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner

// NewQueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner instantiates a new QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner() *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner {
	this := QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner{}
	return &this
}

// NewQueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInnerWithDefaults instantiates a new QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInnerWithDefaults() *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner {
	this := QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner{}
	return &this
}

// GetAlgoId returns the AlgoId field value if set, zero value otherwise.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) GetAlgoId() int64 {
	if o == nil || common.IsNil(o.AlgoId) {
		var ret int64
		return ret
	}
	return *o.AlgoId
}

// GetAlgoIdOk returns a tuple with the AlgoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) GetAlgoIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.AlgoId) {
		return nil, false
	}
	return o.AlgoId, true
}

// HasAlgoId returns a boolean if a field has been set.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) HasAlgoId() bool {
	if o != nil && !common.IsNil(o.AlgoId) {
		return true
	}

	return false
}

// SetAlgoId gets a reference to the given int64 and assigns it to the AlgoId field.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) SetAlgoId(v int64) {
	o.AlgoId = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) SetSide(v string) {
	o.Side = &v
}

// GetPositionSide returns the PositionSide field value if set, zero value otherwise.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) GetPositionSide() string {
	if o == nil || common.IsNil(o.PositionSide) {
		var ret string
		return ret
	}
	return *o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) GetPositionSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionSide) {
		return nil, false
	}
	return o.PositionSide, true
}

// HasPositionSide returns a boolean if a field has been set.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) HasPositionSide() bool {
	if o != nil && !common.IsNil(o.PositionSide) {
		return true
	}

	return false
}

// SetPositionSide gets a reference to the given string and assigns it to the PositionSide field.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) SetPositionSide(v string) {
	o.PositionSide = &v
}

// GetTotalQty returns the TotalQty field value if set, zero value otherwise.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) GetTotalQty() string {
	if o == nil || common.IsNil(o.TotalQty) {
		var ret string
		return ret
	}
	return *o.TotalQty
}

// GetTotalQtyOk returns a tuple with the TotalQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) GetTotalQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalQty) {
		return nil, false
	}
	return o.TotalQty, true
}

// HasTotalQty returns a boolean if a field has been set.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) HasTotalQty() bool {
	if o != nil && !common.IsNil(o.TotalQty) {
		return true
	}

	return false
}

// SetTotalQty gets a reference to the given string and assigns it to the TotalQty field.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) SetTotalQty(v string) {
	o.TotalQty = &v
}

// GetExecutedQty returns the ExecutedQty field value if set, zero value otherwise.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) GetExecutedQty() string {
	if o == nil || common.IsNil(o.ExecutedQty) {
		var ret string
		return ret
	}
	return *o.ExecutedQty
}

// GetExecutedQtyOk returns a tuple with the ExecutedQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) GetExecutedQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExecutedQty) {
		return nil, false
	}
	return o.ExecutedQty, true
}

// HasExecutedQty returns a boolean if a field has been set.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) HasExecutedQty() bool {
	if o != nil && !common.IsNil(o.ExecutedQty) {
		return true
	}

	return false
}

// SetExecutedQty gets a reference to the given string and assigns it to the ExecutedQty field.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) SetExecutedQty(v string) {
	o.ExecutedQty = &v
}

// GetExecutedAmt returns the ExecutedAmt field value if set, zero value otherwise.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) GetExecutedAmt() string {
	if o == nil || common.IsNil(o.ExecutedAmt) {
		var ret string
		return ret
	}
	return *o.ExecutedAmt
}

// GetExecutedAmtOk returns a tuple with the ExecutedAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) GetExecutedAmtOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExecutedAmt) {
		return nil, false
	}
	return o.ExecutedAmt, true
}

// HasExecutedAmt returns a boolean if a field has been set.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) HasExecutedAmt() bool {
	if o != nil && !common.IsNil(o.ExecutedAmt) {
		return true
	}

	return false
}

// SetExecutedAmt gets a reference to the given string and assigns it to the ExecutedAmt field.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) SetExecutedAmt(v string) {
	o.ExecutedAmt = &v
}

// GetAvgPrice returns the AvgPrice field value if set, zero value otherwise.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) GetAvgPrice() string {
	if o == nil || common.IsNil(o.AvgPrice) {
		var ret string
		return ret
	}
	return *o.AvgPrice
}

// GetAvgPriceOk returns a tuple with the AvgPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) GetAvgPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.AvgPrice) {
		return nil, false
	}
	return o.AvgPrice, true
}

// HasAvgPrice returns a boolean if a field has been set.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) HasAvgPrice() bool {
	if o != nil && !common.IsNil(o.AvgPrice) {
		return true
	}

	return false
}

// SetAvgPrice gets a reference to the given string and assigns it to the AvgPrice field.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) SetAvgPrice(v string) {
	o.AvgPrice = &v
}

// GetClientAlgoId returns the ClientAlgoId field value if set, zero value otherwise.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) GetClientAlgoId() string {
	if o == nil || common.IsNil(o.ClientAlgoId) {
		var ret string
		return ret
	}
	return *o.ClientAlgoId
}

// GetClientAlgoIdOk returns a tuple with the ClientAlgoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) GetClientAlgoIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientAlgoId) {
		return nil, false
	}
	return o.ClientAlgoId, true
}

// HasClientAlgoId returns a boolean if a field has been set.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) HasClientAlgoId() bool {
	if o != nil && !common.IsNil(o.ClientAlgoId) {
		return true
	}

	return false
}

// SetClientAlgoId gets a reference to the given string and assigns it to the ClientAlgoId field.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) SetClientAlgoId(v string) {
	o.ClientAlgoId = &v
}

// GetBookTime returns the BookTime field value if set, zero value otherwise.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) GetBookTime() int64 {
	if o == nil || common.IsNil(o.BookTime) {
		var ret int64
		return ret
	}
	return *o.BookTime
}

// GetBookTimeOk returns a tuple with the BookTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) GetBookTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.BookTime) {
		return nil, false
	}
	return o.BookTime, true
}

// HasBookTime returns a boolean if a field has been set.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) HasBookTime() bool {
	if o != nil && !common.IsNil(o.BookTime) {
		return true
	}

	return false
}

// SetBookTime gets a reference to the given int64 and assigns it to the BookTime field.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) SetBookTime(v int64) {
	o.BookTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) GetEndTime() int64 {
	if o == nil || common.IsNil(o.EndTime) {
		var ret int64
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) GetEndTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) HasEndTime() bool {
	if o != nil && !common.IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given int64 and assigns it to the EndTime field.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) SetEndTime(v int64) {
	o.EndTime = &v
}

// GetAlgoStatus returns the AlgoStatus field value if set, zero value otherwise.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) GetAlgoStatus() string {
	if o == nil || common.IsNil(o.AlgoStatus) {
		var ret string
		return ret
	}
	return *o.AlgoStatus
}

// GetAlgoStatusOk returns a tuple with the AlgoStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) GetAlgoStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.AlgoStatus) {
		return nil, false
	}
	return o.AlgoStatus, true
}

// HasAlgoStatus returns a boolean if a field has been set.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) HasAlgoStatus() bool {
	if o != nil && !common.IsNil(o.AlgoStatus) {
		return true
	}

	return false
}

// SetAlgoStatus gets a reference to the given string and assigns it to the AlgoStatus field.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) SetAlgoStatus(v string) {
	o.AlgoStatus = &v
}

// GetAlgoType returns the AlgoType field value if set, zero value otherwise.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) GetAlgoType() string {
	if o == nil || common.IsNil(o.AlgoType) {
		var ret string
		return ret
	}
	return *o.AlgoType
}

// GetAlgoTypeOk returns a tuple with the AlgoType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) GetAlgoTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.AlgoType) {
		return nil, false
	}
	return o.AlgoType, true
}

// HasAlgoType returns a boolean if a field has been set.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) HasAlgoType() bool {
	if o != nil && !common.IsNil(o.AlgoType) {
		return true
	}

	return false
}

// SetAlgoType gets a reference to the given string and assigns it to the AlgoType field.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) SetAlgoType(v string) {
	o.AlgoType = &v
}

// GetUrgency returns the Urgency field value if set, zero value otherwise.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) GetUrgency() string {
	if o == nil || common.IsNil(o.Urgency) {
		var ret string
		return ret
	}
	return *o.Urgency
}

// GetUrgencyOk returns a tuple with the Urgency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) GetUrgencyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Urgency) {
		return nil, false
	}
	return o.Urgency, true
}

// HasUrgency returns a boolean if a field has been set.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) HasUrgency() bool {
	if o != nil && !common.IsNil(o.Urgency) {
		return true
	}

	return false
}

// SetUrgency gets a reference to the given string and assigns it to the Urgency field.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) SetUrgency(v string) {
	o.Urgency = &v
}

func (o QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AlgoId) {
		toSerialize["algoId"] = o.AlgoId
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !common.IsNil(o.PositionSide) {
		toSerialize["positionSide"] = o.PositionSide
	}
	if !common.IsNil(o.TotalQty) {
		toSerialize["totalQty"] = o.TotalQty
	}
	if !common.IsNil(o.ExecutedQty) {
		toSerialize["executedQty"] = o.ExecutedQty
	}
	if !common.IsNil(o.ExecutedAmt) {
		toSerialize["executedAmt"] = o.ExecutedAmt
	}
	if !common.IsNil(o.AvgPrice) {
		toSerialize["avgPrice"] = o.AvgPrice
	}
	if !common.IsNil(o.ClientAlgoId) {
		toSerialize["clientAlgoId"] = o.ClientAlgoId
	}
	if !common.IsNil(o.BookTime) {
		toSerialize["bookTime"] = o.BookTime
	}
	if !common.IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	if !common.IsNil(o.AlgoStatus) {
		toSerialize["algoStatus"] = o.AlgoStatus
	}
	if !common.IsNil(o.AlgoType) {
		toSerialize["algoType"] = o.AlgoType
	}
	if !common.IsNil(o.Urgency) {
		toSerialize["urgency"] = o.Urgency
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) UnmarshalJSON(data []byte) (err error) {
	varQueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner := _QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner{}

	err = json.Unmarshal(data, &varQueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner)

	if err != nil {
		return err
	}

	*o = QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner(varQueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "algoId")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "side")
		delete(additionalProperties, "positionSide")
		delete(additionalProperties, "totalQty")
		delete(additionalProperties, "executedQty")
		delete(additionalProperties, "executedAmt")
		delete(additionalProperties, "avgPrice")
		delete(additionalProperties, "clientAlgoId")
		delete(additionalProperties, "bookTime")
		delete(additionalProperties, "endTime")
		delete(additionalProperties, "algoStatus")
		delete(additionalProperties, "algoType")
		delete(additionalProperties, "urgency")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner struct {
	value *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner
	isSet bool
}

func (v NullableQueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) Get() *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner {
	return v.value
}

func (v *NullableQueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) Set(val *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner(val *QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) *NullableQueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner {
	return &NullableQueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner{value: val, isSet: true}
}

func (v NullableQueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
