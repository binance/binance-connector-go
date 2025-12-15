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

// checks if the QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner{}

// QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner struct for QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner
type QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner struct {
	AlgoId               *int64  `json:"algoId,omitempty"`
	Symbol               *string `json:"symbol,omitempty"`
	Side                 *string `json:"side,omitempty"`
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

type _QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner

// NewQueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner instantiates a new QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner() *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner {
	this := QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner{}
	return &this
}

// NewQueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInnerWithDefaults instantiates a new QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInnerWithDefaults() *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner {
	this := QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner{}
	return &this
}

// GetAlgoId returns the AlgoId field value if set, zero value otherwise.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) GetAlgoId() int64 {
	if o == nil || common.IsNil(o.AlgoId) {
		var ret int64
		return ret
	}
	return *o.AlgoId
}

// GetAlgoIdOk returns a tuple with the AlgoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) GetAlgoIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.AlgoId) {
		return nil, false
	}
	return o.AlgoId, true
}

// HasAlgoId returns a boolean if a field has been set.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) HasAlgoId() bool {
	if o != nil && !common.IsNil(o.AlgoId) {
		return true
	}

	return false
}

// SetAlgoId gets a reference to the given int64 and assigns it to the AlgoId field.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) SetAlgoId(v int64) {
	o.AlgoId = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) SetSide(v string) {
	o.Side = &v
}

// GetTotalQty returns the TotalQty field value if set, zero value otherwise.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) GetTotalQty() string {
	if o == nil || common.IsNil(o.TotalQty) {
		var ret string
		return ret
	}
	return *o.TotalQty
}

// GetTotalQtyOk returns a tuple with the TotalQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) GetTotalQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalQty) {
		return nil, false
	}
	return o.TotalQty, true
}

// HasTotalQty returns a boolean if a field has been set.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) HasTotalQty() bool {
	if o != nil && !common.IsNil(o.TotalQty) {
		return true
	}

	return false
}

// SetTotalQty gets a reference to the given string and assigns it to the TotalQty field.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) SetTotalQty(v string) {
	o.TotalQty = &v
}

// GetExecutedQty returns the ExecutedQty field value if set, zero value otherwise.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) GetExecutedQty() string {
	if o == nil || common.IsNil(o.ExecutedQty) {
		var ret string
		return ret
	}
	return *o.ExecutedQty
}

// GetExecutedQtyOk returns a tuple with the ExecutedQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) GetExecutedQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExecutedQty) {
		return nil, false
	}
	return o.ExecutedQty, true
}

// HasExecutedQty returns a boolean if a field has been set.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) HasExecutedQty() bool {
	if o != nil && !common.IsNil(o.ExecutedQty) {
		return true
	}

	return false
}

// SetExecutedQty gets a reference to the given string and assigns it to the ExecutedQty field.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) SetExecutedQty(v string) {
	o.ExecutedQty = &v
}

// GetExecutedAmt returns the ExecutedAmt field value if set, zero value otherwise.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) GetExecutedAmt() string {
	if o == nil || common.IsNil(o.ExecutedAmt) {
		var ret string
		return ret
	}
	return *o.ExecutedAmt
}

// GetExecutedAmtOk returns a tuple with the ExecutedAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) GetExecutedAmtOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExecutedAmt) {
		return nil, false
	}
	return o.ExecutedAmt, true
}

// HasExecutedAmt returns a boolean if a field has been set.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) HasExecutedAmt() bool {
	if o != nil && !common.IsNil(o.ExecutedAmt) {
		return true
	}

	return false
}

// SetExecutedAmt gets a reference to the given string and assigns it to the ExecutedAmt field.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) SetExecutedAmt(v string) {
	o.ExecutedAmt = &v
}

// GetAvgPrice returns the AvgPrice field value if set, zero value otherwise.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) GetAvgPrice() string {
	if o == nil || common.IsNil(o.AvgPrice) {
		var ret string
		return ret
	}
	return *o.AvgPrice
}

// GetAvgPriceOk returns a tuple with the AvgPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) GetAvgPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.AvgPrice) {
		return nil, false
	}
	return o.AvgPrice, true
}

// HasAvgPrice returns a boolean if a field has been set.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) HasAvgPrice() bool {
	if o != nil && !common.IsNil(o.AvgPrice) {
		return true
	}

	return false
}

// SetAvgPrice gets a reference to the given string and assigns it to the AvgPrice field.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) SetAvgPrice(v string) {
	o.AvgPrice = &v
}

// GetClientAlgoId returns the ClientAlgoId field value if set, zero value otherwise.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) GetClientAlgoId() string {
	if o == nil || common.IsNil(o.ClientAlgoId) {
		var ret string
		return ret
	}
	return *o.ClientAlgoId
}

// GetClientAlgoIdOk returns a tuple with the ClientAlgoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) GetClientAlgoIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientAlgoId) {
		return nil, false
	}
	return o.ClientAlgoId, true
}

// HasClientAlgoId returns a boolean if a field has been set.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) HasClientAlgoId() bool {
	if o != nil && !common.IsNil(o.ClientAlgoId) {
		return true
	}

	return false
}

// SetClientAlgoId gets a reference to the given string and assigns it to the ClientAlgoId field.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) SetClientAlgoId(v string) {
	o.ClientAlgoId = &v
}

// GetBookTime returns the BookTime field value if set, zero value otherwise.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) GetBookTime() int64 {
	if o == nil || common.IsNil(o.BookTime) {
		var ret int64
		return ret
	}
	return *o.BookTime
}

// GetBookTimeOk returns a tuple with the BookTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) GetBookTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.BookTime) {
		return nil, false
	}
	return o.BookTime, true
}

// HasBookTime returns a boolean if a field has been set.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) HasBookTime() bool {
	if o != nil && !common.IsNil(o.BookTime) {
		return true
	}

	return false
}

// SetBookTime gets a reference to the given int64 and assigns it to the BookTime field.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) SetBookTime(v int64) {
	o.BookTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) GetEndTime() int64 {
	if o == nil || common.IsNil(o.EndTime) {
		var ret int64
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) GetEndTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) HasEndTime() bool {
	if o != nil && !common.IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given int64 and assigns it to the EndTime field.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) SetEndTime(v int64) {
	o.EndTime = &v
}

// GetAlgoStatus returns the AlgoStatus field value if set, zero value otherwise.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) GetAlgoStatus() string {
	if o == nil || common.IsNil(o.AlgoStatus) {
		var ret string
		return ret
	}
	return *o.AlgoStatus
}

// GetAlgoStatusOk returns a tuple with the AlgoStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) GetAlgoStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.AlgoStatus) {
		return nil, false
	}
	return o.AlgoStatus, true
}

// HasAlgoStatus returns a boolean if a field has been set.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) HasAlgoStatus() bool {
	if o != nil && !common.IsNil(o.AlgoStatus) {
		return true
	}

	return false
}

// SetAlgoStatus gets a reference to the given string and assigns it to the AlgoStatus field.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) SetAlgoStatus(v string) {
	o.AlgoStatus = &v
}

// GetAlgoType returns the AlgoType field value if set, zero value otherwise.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) GetAlgoType() string {
	if o == nil || common.IsNil(o.AlgoType) {
		var ret string
		return ret
	}
	return *o.AlgoType
}

// GetAlgoTypeOk returns a tuple with the AlgoType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) GetAlgoTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.AlgoType) {
		return nil, false
	}
	return o.AlgoType, true
}

// HasAlgoType returns a boolean if a field has been set.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) HasAlgoType() bool {
	if o != nil && !common.IsNil(o.AlgoType) {
		return true
	}

	return false
}

// SetAlgoType gets a reference to the given string and assigns it to the AlgoType field.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) SetAlgoType(v string) {
	o.AlgoType = &v
}

// GetUrgency returns the Urgency field value if set, zero value otherwise.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) GetUrgency() string {
	if o == nil || common.IsNil(o.Urgency) {
		var ret string
		return ret
	}
	return *o.Urgency
}

// GetUrgencyOk returns a tuple with the Urgency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) GetUrgencyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Urgency) {
		return nil, false
	}
	return o.Urgency, true
}

// HasUrgency returns a boolean if a field has been set.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) HasUrgency() bool {
	if o != nil && !common.IsNil(o.Urgency) {
		return true
	}

	return false
}

// SetUrgency gets a reference to the given string and assigns it to the Urgency field.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) SetUrgency(v string) {
	o.Urgency = &v
}

func (o QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) ToMap() (map[string]interface{}, error) {
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

func (o *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) UnmarshalJSON(data []byte) (err error) {
	varQueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner := _QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner{}

	err = json.Unmarshal(data, &varQueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner)

	if err != nil {
		return err
	}

	*o = QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner(varQueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "algoId")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "side")
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

type NullableQueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner struct {
	value *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner
	isSet bool
}

func (v NullableQueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) Get() *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner {
	return v.value
}

func (v *NullableQueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) Set(val *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner(val *QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) *NullableQueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner {
	return &NullableQueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner{value: val, isSet: true}
}

func (v NullableQueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
