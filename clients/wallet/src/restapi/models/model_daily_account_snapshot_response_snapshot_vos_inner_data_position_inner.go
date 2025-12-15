/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner{}

// DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner struct for DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner
type DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner struct {
	EntryPrice           *string `json:"entryPrice,omitempty"`
	MarkPrice            *string `json:"markPrice,omitempty"`
	PositionAmt          *string `json:"positionAmt,omitempty"`
	Symbol               *string `json:"symbol,omitempty"`
	UnRealizedProfit     *string `json:"unRealizedProfit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner

// NewDailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner instantiates a new DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner() *DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner {
	this := DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner{}
	return &this
}

// NewDailyAccountSnapshotResponseSnapshotVosInnerDataPositionInnerWithDefaults instantiates a new DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDailyAccountSnapshotResponseSnapshotVosInnerDataPositionInnerWithDefaults() *DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner {
	this := DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner{}
	return &this
}

// GetEntryPrice returns the EntryPrice field value if set, zero value otherwise.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner) GetEntryPrice() string {
	if o == nil || common.IsNil(o.EntryPrice) {
		var ret string
		return ret
	}
	return *o.EntryPrice
}

// GetEntryPriceOk returns a tuple with the EntryPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner) GetEntryPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.EntryPrice) {
		return nil, false
	}
	return o.EntryPrice, true
}

// HasEntryPrice returns a boolean if a field has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner) HasEntryPrice() bool {
	if o != nil && !common.IsNil(o.EntryPrice) {
		return true
	}

	return false
}

// SetEntryPrice gets a reference to the given string and assigns it to the EntryPrice field.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner) SetEntryPrice(v string) {
	o.EntryPrice = &v
}

// GetMarkPrice returns the MarkPrice field value if set, zero value otherwise.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner) GetMarkPrice() string {
	if o == nil || common.IsNil(o.MarkPrice) {
		var ret string
		return ret
	}
	return *o.MarkPrice
}

// GetMarkPriceOk returns a tuple with the MarkPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner) GetMarkPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarkPrice) {
		return nil, false
	}
	return o.MarkPrice, true
}

// HasMarkPrice returns a boolean if a field has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner) HasMarkPrice() bool {
	if o != nil && !common.IsNil(o.MarkPrice) {
		return true
	}

	return false
}

// SetMarkPrice gets a reference to the given string and assigns it to the MarkPrice field.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner) SetMarkPrice(v string) {
	o.MarkPrice = &v
}

// GetPositionAmt returns the PositionAmt field value if set, zero value otherwise.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner) GetPositionAmt() string {
	if o == nil || common.IsNil(o.PositionAmt) {
		var ret string
		return ret
	}
	return *o.PositionAmt
}

// GetPositionAmtOk returns a tuple with the PositionAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner) GetPositionAmtOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionAmt) {
		return nil, false
	}
	return o.PositionAmt, true
}

// HasPositionAmt returns a boolean if a field has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner) HasPositionAmt() bool {
	if o != nil && !common.IsNil(o.PositionAmt) {
		return true
	}

	return false
}

// SetPositionAmt gets a reference to the given string and assigns it to the PositionAmt field.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner) SetPositionAmt(v string) {
	o.PositionAmt = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetUnRealizedProfit returns the UnRealizedProfit field value if set, zero value otherwise.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner) GetUnRealizedProfit() string {
	if o == nil || common.IsNil(o.UnRealizedProfit) {
		var ret string
		return ret
	}
	return *o.UnRealizedProfit
}

// GetUnRealizedProfitOk returns a tuple with the UnRealizedProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner) GetUnRealizedProfitOk() (*string, bool) {
	if o == nil || common.IsNil(o.UnRealizedProfit) {
		return nil, false
	}
	return o.UnRealizedProfit, true
}

// HasUnRealizedProfit returns a boolean if a field has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner) HasUnRealizedProfit() bool {
	if o != nil && !common.IsNil(o.UnRealizedProfit) {
		return true
	}

	return false
}

// SetUnRealizedProfit gets a reference to the given string and assigns it to the UnRealizedProfit field.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner) SetUnRealizedProfit(v string) {
	o.UnRealizedProfit = &v
}

func (o DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.EntryPrice) {
		toSerialize["entryPrice"] = o.EntryPrice
	}
	if !common.IsNil(o.MarkPrice) {
		toSerialize["markPrice"] = o.MarkPrice
	}
	if !common.IsNil(o.PositionAmt) {
		toSerialize["positionAmt"] = o.PositionAmt
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.UnRealizedProfit) {
		toSerialize["unRealizedProfit"] = o.UnRealizedProfit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner) UnmarshalJSON(data []byte) (err error) {
	varDailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner := _DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner{}

	err = json.Unmarshal(data, &varDailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner)

	if err != nil {
		return err
	}

	*o = DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner(varDailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "entryPrice")
		delete(additionalProperties, "markPrice")
		delete(additionalProperties, "positionAmt")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "unRealizedProfit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner struct {
	value *DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner
	isSet bool
}

func (v NullableDailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner) Get() *DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner {
	return v.value
}

func (v *NullableDailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner) Set(val *DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner(val *DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner) *NullableDailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner {
	return &NullableDailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner{value: val, isSet: true}
}

func (v NullableDailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
