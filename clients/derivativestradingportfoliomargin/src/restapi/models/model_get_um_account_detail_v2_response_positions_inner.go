/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetUmAccountDetailV2ResponsePositionsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetUmAccountDetailV2ResponsePositionsInner{}

// GetUmAccountDetailV2ResponsePositionsInner struct for GetUmAccountDetailV2ResponsePositionsInner
type GetUmAccountDetailV2ResponsePositionsInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	InitialMargin        *string `json:"initialMargin,omitempty"`
	MaintMargin          *string `json:"maintMargin,omitempty"`
	UnrealizedProfit     *string `json:"unrealizedProfit,omitempty"`
	PositionSide         *string `json:"positionSide,omitempty"`
	PositionAmt          *string `json:"positionAmt,omitempty"`
	UpdateTime           *int64  `json:"updateTime,omitempty"`
	Notional             *string `json:"notional,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetUmAccountDetailV2ResponsePositionsInner GetUmAccountDetailV2ResponsePositionsInner

// NewGetUmAccountDetailV2ResponsePositionsInner instantiates a new GetUmAccountDetailV2ResponsePositionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUmAccountDetailV2ResponsePositionsInner() *GetUmAccountDetailV2ResponsePositionsInner {
	this := GetUmAccountDetailV2ResponsePositionsInner{}
	return &this
}

// NewGetUmAccountDetailV2ResponsePositionsInnerWithDefaults instantiates a new GetUmAccountDetailV2ResponsePositionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUmAccountDetailV2ResponsePositionsInnerWithDefaults() *GetUmAccountDetailV2ResponsePositionsInner {
	this := GetUmAccountDetailV2ResponsePositionsInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetUmAccountDetailV2ResponsePositionsInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailV2ResponsePositionsInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetUmAccountDetailV2ResponsePositionsInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetUmAccountDetailV2ResponsePositionsInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetInitialMargin returns the InitialMargin field value if set, zero value otherwise.
func (o *GetUmAccountDetailV2ResponsePositionsInner) GetInitialMargin() string {
	if o == nil || common.IsNil(o.InitialMargin) {
		var ret string
		return ret
	}
	return *o.InitialMargin
}

// GetInitialMarginOk returns a tuple with the InitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailV2ResponsePositionsInner) GetInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.InitialMargin) {
		return nil, false
	}
	return o.InitialMargin, true
}

// HasInitialMargin returns a boolean if a field has been set.
func (o *GetUmAccountDetailV2ResponsePositionsInner) HasInitialMargin() bool {
	if o != nil && !common.IsNil(o.InitialMargin) {
		return true
	}

	return false
}

// SetInitialMargin gets a reference to the given string and assigns it to the InitialMargin field.
func (o *GetUmAccountDetailV2ResponsePositionsInner) SetInitialMargin(v string) {
	o.InitialMargin = &v
}

// GetMaintMargin returns the MaintMargin field value if set, zero value otherwise.
func (o *GetUmAccountDetailV2ResponsePositionsInner) GetMaintMargin() string {
	if o == nil || common.IsNil(o.MaintMargin) {
		var ret string
		return ret
	}
	return *o.MaintMargin
}

// GetMaintMarginOk returns a tuple with the MaintMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailV2ResponsePositionsInner) GetMaintMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaintMargin) {
		return nil, false
	}
	return o.MaintMargin, true
}

// HasMaintMargin returns a boolean if a field has been set.
func (o *GetUmAccountDetailV2ResponsePositionsInner) HasMaintMargin() bool {
	if o != nil && !common.IsNil(o.MaintMargin) {
		return true
	}

	return false
}

// SetMaintMargin gets a reference to the given string and assigns it to the MaintMargin field.
func (o *GetUmAccountDetailV2ResponsePositionsInner) SetMaintMargin(v string) {
	o.MaintMargin = &v
}

// GetUnrealizedProfit returns the UnrealizedProfit field value if set, zero value otherwise.
func (o *GetUmAccountDetailV2ResponsePositionsInner) GetUnrealizedProfit() string {
	if o == nil || common.IsNil(o.UnrealizedProfit) {
		var ret string
		return ret
	}
	return *o.UnrealizedProfit
}

// GetUnrealizedProfitOk returns a tuple with the UnrealizedProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailV2ResponsePositionsInner) GetUnrealizedProfitOk() (*string, bool) {
	if o == nil || common.IsNil(o.UnrealizedProfit) {
		return nil, false
	}
	return o.UnrealizedProfit, true
}

// HasUnrealizedProfit returns a boolean if a field has been set.
func (o *GetUmAccountDetailV2ResponsePositionsInner) HasUnrealizedProfit() bool {
	if o != nil && !common.IsNil(o.UnrealizedProfit) {
		return true
	}

	return false
}

// SetUnrealizedProfit gets a reference to the given string and assigns it to the UnrealizedProfit field.
func (o *GetUmAccountDetailV2ResponsePositionsInner) SetUnrealizedProfit(v string) {
	o.UnrealizedProfit = &v
}

// GetPositionSide returns the PositionSide field value if set, zero value otherwise.
func (o *GetUmAccountDetailV2ResponsePositionsInner) GetPositionSide() string {
	if o == nil || common.IsNil(o.PositionSide) {
		var ret string
		return ret
	}
	return *o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailV2ResponsePositionsInner) GetPositionSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionSide) {
		return nil, false
	}
	return o.PositionSide, true
}

// HasPositionSide returns a boolean if a field has been set.
func (o *GetUmAccountDetailV2ResponsePositionsInner) HasPositionSide() bool {
	if o != nil && !common.IsNil(o.PositionSide) {
		return true
	}

	return false
}

// SetPositionSide gets a reference to the given string and assigns it to the PositionSide field.
func (o *GetUmAccountDetailV2ResponsePositionsInner) SetPositionSide(v string) {
	o.PositionSide = &v
}

// GetPositionAmt returns the PositionAmt field value if set, zero value otherwise.
func (o *GetUmAccountDetailV2ResponsePositionsInner) GetPositionAmt() string {
	if o == nil || common.IsNil(o.PositionAmt) {
		var ret string
		return ret
	}
	return *o.PositionAmt
}

// GetPositionAmtOk returns a tuple with the PositionAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailV2ResponsePositionsInner) GetPositionAmtOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionAmt) {
		return nil, false
	}
	return o.PositionAmt, true
}

// HasPositionAmt returns a boolean if a field has been set.
func (o *GetUmAccountDetailV2ResponsePositionsInner) HasPositionAmt() bool {
	if o != nil && !common.IsNil(o.PositionAmt) {
		return true
	}

	return false
}

// SetPositionAmt gets a reference to the given string and assigns it to the PositionAmt field.
func (o *GetUmAccountDetailV2ResponsePositionsInner) SetPositionAmt(v string) {
	o.PositionAmt = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *GetUmAccountDetailV2ResponsePositionsInner) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailV2ResponsePositionsInner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *GetUmAccountDetailV2ResponsePositionsInner) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *GetUmAccountDetailV2ResponsePositionsInner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

// GetNotional returns the Notional field value if set, zero value otherwise.
func (o *GetUmAccountDetailV2ResponsePositionsInner) GetNotional() string {
	if o == nil || common.IsNil(o.Notional) {
		var ret string
		return ret
	}
	return *o.Notional
}

// GetNotionalOk returns a tuple with the Notional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailV2ResponsePositionsInner) GetNotionalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Notional) {
		return nil, false
	}
	return o.Notional, true
}

// HasNotional returns a boolean if a field has been set.
func (o *GetUmAccountDetailV2ResponsePositionsInner) HasNotional() bool {
	if o != nil && !common.IsNil(o.Notional) {
		return true
	}

	return false
}

// SetNotional gets a reference to the given string and assigns it to the Notional field.
func (o *GetUmAccountDetailV2ResponsePositionsInner) SetNotional(v string) {
	o.Notional = &v
}

func (o GetUmAccountDetailV2ResponsePositionsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUmAccountDetailV2ResponsePositionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.InitialMargin) {
		toSerialize["initialMargin"] = o.InitialMargin
	}
	if !common.IsNil(o.MaintMargin) {
		toSerialize["maintMargin"] = o.MaintMargin
	}
	if !common.IsNil(o.UnrealizedProfit) {
		toSerialize["unrealizedProfit"] = o.UnrealizedProfit
	}
	if !common.IsNil(o.PositionSide) {
		toSerialize["positionSide"] = o.PositionSide
	}
	if !common.IsNil(o.PositionAmt) {
		toSerialize["positionAmt"] = o.PositionAmt
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if !common.IsNil(o.Notional) {
		toSerialize["notional"] = o.Notional
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetUmAccountDetailV2ResponsePositionsInner) UnmarshalJSON(data []byte) (err error) {
	varGetUmAccountDetailV2ResponsePositionsInner := _GetUmAccountDetailV2ResponsePositionsInner{}

	err = json.Unmarshal(data, &varGetUmAccountDetailV2ResponsePositionsInner)

	if err != nil {
		return err
	}

	*o = GetUmAccountDetailV2ResponsePositionsInner(varGetUmAccountDetailV2ResponsePositionsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "initialMargin")
		delete(additionalProperties, "maintMargin")
		delete(additionalProperties, "unrealizedProfit")
		delete(additionalProperties, "positionSide")
		delete(additionalProperties, "positionAmt")
		delete(additionalProperties, "updateTime")
		delete(additionalProperties, "notional")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetUmAccountDetailV2ResponsePositionsInner struct {
	value *GetUmAccountDetailV2ResponsePositionsInner
	isSet bool
}

func (v NullableGetUmAccountDetailV2ResponsePositionsInner) Get() *GetUmAccountDetailV2ResponsePositionsInner {
	return v.value
}

func (v *NullableGetUmAccountDetailV2ResponsePositionsInner) Set(val *GetUmAccountDetailV2ResponsePositionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUmAccountDetailV2ResponsePositionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUmAccountDetailV2ResponsePositionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUmAccountDetailV2ResponsePositionsInner(val *GetUmAccountDetailV2ResponsePositionsInner) *NullableGetUmAccountDetailV2ResponsePositionsInner {
	return &NullableGetUmAccountDetailV2ResponsePositionsInner{value: val, isSet: true}
}

func (v NullableGetUmAccountDetailV2ResponsePositionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUmAccountDetailV2ResponsePositionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
