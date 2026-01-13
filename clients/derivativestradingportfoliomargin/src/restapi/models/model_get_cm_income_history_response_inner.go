/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetCmIncomeHistoryResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetCmIncomeHistoryResponseInner{}

// GetCmIncomeHistoryResponseInner struct for GetCmIncomeHistoryResponseInner
type GetCmIncomeHistoryResponseInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	IncomeType           *string `json:"incomeType,omitempty"`
	Income               *string `json:"income,omitempty"`
	Asset                *string `json:"asset,omitempty"`
	Info                 *string `json:"info,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	TranId               *string `json:"tranId,omitempty"`
	TradeId              *string `json:"tradeId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetCmIncomeHistoryResponseInner GetCmIncomeHistoryResponseInner

// NewGetCmIncomeHistoryResponseInner instantiates a new GetCmIncomeHistoryResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCmIncomeHistoryResponseInner() *GetCmIncomeHistoryResponseInner {
	this := GetCmIncomeHistoryResponseInner{}
	return &this
}

// NewGetCmIncomeHistoryResponseInnerWithDefaults instantiates a new GetCmIncomeHistoryResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCmIncomeHistoryResponseInnerWithDefaults() *GetCmIncomeHistoryResponseInner {
	this := GetCmIncomeHistoryResponseInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetCmIncomeHistoryResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCmIncomeHistoryResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetCmIncomeHistoryResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetCmIncomeHistoryResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetIncomeType returns the IncomeType field value if set, zero value otherwise.
func (o *GetCmIncomeHistoryResponseInner) GetIncomeType() string {
	if o == nil || common.IsNil(o.IncomeType) {
		var ret string
		return ret
	}
	return *o.IncomeType
}

// GetIncomeTypeOk returns a tuple with the IncomeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCmIncomeHistoryResponseInner) GetIncomeTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.IncomeType) {
		return nil, false
	}
	return o.IncomeType, true
}

// HasIncomeType returns a boolean if a field has been set.
func (o *GetCmIncomeHistoryResponseInner) HasIncomeType() bool {
	if o != nil && !common.IsNil(o.IncomeType) {
		return true
	}

	return false
}

// SetIncomeType gets a reference to the given string and assigns it to the IncomeType field.
func (o *GetCmIncomeHistoryResponseInner) SetIncomeType(v string) {
	o.IncomeType = &v
}

// GetIncome returns the Income field value if set, zero value otherwise.
func (o *GetCmIncomeHistoryResponseInner) GetIncome() string {
	if o == nil || common.IsNil(o.Income) {
		var ret string
		return ret
	}
	return *o.Income
}

// GetIncomeOk returns a tuple with the Income field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCmIncomeHistoryResponseInner) GetIncomeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Income) {
		return nil, false
	}
	return o.Income, true
}

// HasIncome returns a boolean if a field has been set.
func (o *GetCmIncomeHistoryResponseInner) HasIncome() bool {
	if o != nil && !common.IsNil(o.Income) {
		return true
	}

	return false
}

// SetIncome gets a reference to the given string and assigns it to the Income field.
func (o *GetCmIncomeHistoryResponseInner) SetIncome(v string) {
	o.Income = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetCmIncomeHistoryResponseInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCmIncomeHistoryResponseInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetCmIncomeHistoryResponseInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetCmIncomeHistoryResponseInner) SetAsset(v string) {
	o.Asset = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *GetCmIncomeHistoryResponseInner) GetInfo() string {
	if o == nil || common.IsNil(o.Info) {
		var ret string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCmIncomeHistoryResponseInner) GetInfoOk() (*string, bool) {
	if o == nil || common.IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *GetCmIncomeHistoryResponseInner) HasInfo() bool {
	if o != nil && !common.IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given string and assigns it to the Info field.
func (o *GetCmIncomeHistoryResponseInner) SetInfo(v string) {
	o.Info = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetCmIncomeHistoryResponseInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCmIncomeHistoryResponseInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetCmIncomeHistoryResponseInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *GetCmIncomeHistoryResponseInner) SetTime(v int64) {
	o.Time = &v
}

// GetTranId returns the TranId field value if set, zero value otherwise.
func (o *GetCmIncomeHistoryResponseInner) GetTranId() string {
	if o == nil || common.IsNil(o.TranId) {
		var ret string
		return ret
	}
	return *o.TranId
}

// GetTranIdOk returns a tuple with the TranId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCmIncomeHistoryResponseInner) GetTranIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TranId) {
		return nil, false
	}
	return o.TranId, true
}

// HasTranId returns a boolean if a field has been set.
func (o *GetCmIncomeHistoryResponseInner) HasTranId() bool {
	if o != nil && !common.IsNil(o.TranId) {
		return true
	}

	return false
}

// SetTranId gets a reference to the given string and assigns it to the TranId field.
func (o *GetCmIncomeHistoryResponseInner) SetTranId(v string) {
	o.TranId = &v
}

// GetTradeId returns the TradeId field value if set, zero value otherwise.
func (o *GetCmIncomeHistoryResponseInner) GetTradeId() string {
	if o == nil || common.IsNil(o.TradeId) {
		var ret string
		return ret
	}
	return *o.TradeId
}

// GetTradeIdOk returns a tuple with the TradeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCmIncomeHistoryResponseInner) GetTradeIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TradeId) {
		return nil, false
	}
	return o.TradeId, true
}

// HasTradeId returns a boolean if a field has been set.
func (o *GetCmIncomeHistoryResponseInner) HasTradeId() bool {
	if o != nil && !common.IsNil(o.TradeId) {
		return true
	}

	return false
}

// SetTradeId gets a reference to the given string and assigns it to the TradeId field.
func (o *GetCmIncomeHistoryResponseInner) SetTradeId(v string) {
	o.TradeId = &v
}

func (o GetCmIncomeHistoryResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCmIncomeHistoryResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.IncomeType) {
		toSerialize["incomeType"] = o.IncomeType
	}
	if !common.IsNil(o.Income) {
		toSerialize["income"] = o.Income
	}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Info) {
		toSerialize["info"] = o.Info
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !common.IsNil(o.TranId) {
		toSerialize["tranId"] = o.TranId
	}
	if !common.IsNil(o.TradeId) {
		toSerialize["tradeId"] = o.TradeId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetCmIncomeHistoryResponseInner) UnmarshalJSON(data []byte) (err error) {
	varGetCmIncomeHistoryResponseInner := _GetCmIncomeHistoryResponseInner{}

	err = json.Unmarshal(data, &varGetCmIncomeHistoryResponseInner)

	if err != nil {
		return err
	}

	*o = GetCmIncomeHistoryResponseInner(varGetCmIncomeHistoryResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "incomeType")
		delete(additionalProperties, "income")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "info")
		delete(additionalProperties, "time")
		delete(additionalProperties, "tranId")
		delete(additionalProperties, "tradeId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetCmIncomeHistoryResponseInner struct {
	value *GetCmIncomeHistoryResponseInner
	isSet bool
}

func (v NullableGetCmIncomeHistoryResponseInner) Get() *GetCmIncomeHistoryResponseInner {
	return v.value
}

func (v *NullableGetCmIncomeHistoryResponseInner) Set(val *GetCmIncomeHistoryResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCmIncomeHistoryResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCmIncomeHistoryResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCmIncomeHistoryResponseInner(val *GetCmIncomeHistoryResponseInner) *NullableGetCmIncomeHistoryResponseInner {
	return &NullableGetCmIncomeHistoryResponseInner{value: val, isSet: true}
}

func (v NullableGetCmIncomeHistoryResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCmIncomeHistoryResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
