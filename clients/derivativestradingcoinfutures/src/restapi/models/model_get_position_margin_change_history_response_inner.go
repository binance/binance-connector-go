/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetPositionMarginChangeHistoryResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetPositionMarginChangeHistoryResponseInner{}

// GetPositionMarginChangeHistoryResponseInner struct for GetPositionMarginChangeHistoryResponseInner
type GetPositionMarginChangeHistoryResponseInner struct {
	Amount               *string `json:"amount,omitempty"`
	Asset                *string `json:"asset,omitempty"`
	Symbol               *string `json:"symbol,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	Type                 *int64  `json:"type,omitempty"`
	PositionSide         *string `json:"positionSide,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetPositionMarginChangeHistoryResponseInner GetPositionMarginChangeHistoryResponseInner

// NewGetPositionMarginChangeHistoryResponseInner instantiates a new GetPositionMarginChangeHistoryResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPositionMarginChangeHistoryResponseInner() *GetPositionMarginChangeHistoryResponseInner {
	this := GetPositionMarginChangeHistoryResponseInner{}
	return &this
}

// NewGetPositionMarginChangeHistoryResponseInnerWithDefaults instantiates a new GetPositionMarginChangeHistoryResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPositionMarginChangeHistoryResponseInnerWithDefaults() *GetPositionMarginChangeHistoryResponseInner {
	this := GetPositionMarginChangeHistoryResponseInner{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetPositionMarginChangeHistoryResponseInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPositionMarginChangeHistoryResponseInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetPositionMarginChangeHistoryResponseInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *GetPositionMarginChangeHistoryResponseInner) SetAmount(v string) {
	o.Amount = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetPositionMarginChangeHistoryResponseInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPositionMarginChangeHistoryResponseInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetPositionMarginChangeHistoryResponseInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetPositionMarginChangeHistoryResponseInner) SetAsset(v string) {
	o.Asset = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetPositionMarginChangeHistoryResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPositionMarginChangeHistoryResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetPositionMarginChangeHistoryResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetPositionMarginChangeHistoryResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetPositionMarginChangeHistoryResponseInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPositionMarginChangeHistoryResponseInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetPositionMarginChangeHistoryResponseInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *GetPositionMarginChangeHistoryResponseInner) SetTime(v int64) {
	o.Time = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetPositionMarginChangeHistoryResponseInner) GetType() int64 {
	if o == nil || common.IsNil(o.Type) {
		var ret int64
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPositionMarginChangeHistoryResponseInner) GetTypeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetPositionMarginChangeHistoryResponseInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int64 and assigns it to the Type field.
func (o *GetPositionMarginChangeHistoryResponseInner) SetType(v int64) {
	o.Type = &v
}

// GetPositionSide returns the PositionSide field value if set, zero value otherwise.
func (o *GetPositionMarginChangeHistoryResponseInner) GetPositionSide() string {
	if o == nil || common.IsNil(o.PositionSide) {
		var ret string
		return ret
	}
	return *o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPositionMarginChangeHistoryResponseInner) GetPositionSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionSide) {
		return nil, false
	}
	return o.PositionSide, true
}

// HasPositionSide returns a boolean if a field has been set.
func (o *GetPositionMarginChangeHistoryResponseInner) HasPositionSide() bool {
	if o != nil && !common.IsNil(o.PositionSide) {
		return true
	}

	return false
}

// SetPositionSide gets a reference to the given string and assigns it to the PositionSide field.
func (o *GetPositionMarginChangeHistoryResponseInner) SetPositionSide(v string) {
	o.PositionSide = &v
}

func (o GetPositionMarginChangeHistoryResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPositionMarginChangeHistoryResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.PositionSide) {
		toSerialize["positionSide"] = o.PositionSide
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetPositionMarginChangeHistoryResponseInner) UnmarshalJSON(data []byte) (err error) {
	varGetPositionMarginChangeHistoryResponseInner := _GetPositionMarginChangeHistoryResponseInner{}

	err = json.Unmarshal(data, &varGetPositionMarginChangeHistoryResponseInner)

	if err != nil {
		return err
	}

	*o = GetPositionMarginChangeHistoryResponseInner(varGetPositionMarginChangeHistoryResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "time")
		delete(additionalProperties, "type")
		delete(additionalProperties, "positionSide")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetPositionMarginChangeHistoryResponseInner struct {
	value *GetPositionMarginChangeHistoryResponseInner
	isSet bool
}

func (v NullableGetPositionMarginChangeHistoryResponseInner) Get() *GetPositionMarginChangeHistoryResponseInner {
	return v.value
}

func (v *NullableGetPositionMarginChangeHistoryResponseInner) Set(val *GetPositionMarginChangeHistoryResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPositionMarginChangeHistoryResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPositionMarginChangeHistoryResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPositionMarginChangeHistoryResponseInner(val *GetPositionMarginChangeHistoryResponseInner) *NullableGetPositionMarginChangeHistoryResponseInner {
	return &NullableGetPositionMarginChangeHistoryResponseInner{value: val, isSet: true}
}

func (v NullableGetPositionMarginChangeHistoryResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPositionMarginChangeHistoryResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
