/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the UmFuturesSymbolConfigurationResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UmFuturesSymbolConfigurationResponseInner{}

// UmFuturesSymbolConfigurationResponseInner struct for UmFuturesSymbolConfigurationResponseInner
type UmFuturesSymbolConfigurationResponseInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	MarginType           *string `json:"marginType,omitempty"`
	IsAutoAddMargin      *string `json:"isAutoAddMargin,omitempty"`
	Leverage             *int64  `json:"leverage,omitempty"`
	MaxNotionalValue     *string `json:"maxNotionalValue,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UmFuturesSymbolConfigurationResponseInner UmFuturesSymbolConfigurationResponseInner

// NewUmFuturesSymbolConfigurationResponseInner instantiates a new UmFuturesSymbolConfigurationResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUmFuturesSymbolConfigurationResponseInner() *UmFuturesSymbolConfigurationResponseInner {
	this := UmFuturesSymbolConfigurationResponseInner{}
	return &this
}

// NewUmFuturesSymbolConfigurationResponseInnerWithDefaults instantiates a new UmFuturesSymbolConfigurationResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUmFuturesSymbolConfigurationResponseInnerWithDefaults() *UmFuturesSymbolConfigurationResponseInner {
	this := UmFuturesSymbolConfigurationResponseInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *UmFuturesSymbolConfigurationResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmFuturesSymbolConfigurationResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *UmFuturesSymbolConfigurationResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *UmFuturesSymbolConfigurationResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetMarginType returns the MarginType field value if set, zero value otherwise.
func (o *UmFuturesSymbolConfigurationResponseInner) GetMarginType() string {
	if o == nil || common.IsNil(o.MarginType) {
		var ret string
		return ret
	}
	return *o.MarginType
}

// GetMarginTypeOk returns a tuple with the MarginType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmFuturesSymbolConfigurationResponseInner) GetMarginTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarginType) {
		return nil, false
	}
	return o.MarginType, true
}

// HasMarginType returns a boolean if a field has been set.
func (o *UmFuturesSymbolConfigurationResponseInner) HasMarginType() bool {
	if o != nil && !common.IsNil(o.MarginType) {
		return true
	}

	return false
}

// SetMarginType gets a reference to the given string and assigns it to the MarginType field.
func (o *UmFuturesSymbolConfigurationResponseInner) SetMarginType(v string) {
	o.MarginType = &v
}

// GetIsAutoAddMargin returns the IsAutoAddMargin field value if set, zero value otherwise.
func (o *UmFuturesSymbolConfigurationResponseInner) GetIsAutoAddMargin() string {
	if o == nil || common.IsNil(o.IsAutoAddMargin) {
		var ret string
		return ret
	}
	return *o.IsAutoAddMargin
}

// GetIsAutoAddMarginOk returns a tuple with the IsAutoAddMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmFuturesSymbolConfigurationResponseInner) GetIsAutoAddMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.IsAutoAddMargin) {
		return nil, false
	}
	return o.IsAutoAddMargin, true
}

// HasIsAutoAddMargin returns a boolean if a field has been set.
func (o *UmFuturesSymbolConfigurationResponseInner) HasIsAutoAddMargin() bool {
	if o != nil && !common.IsNil(o.IsAutoAddMargin) {
		return true
	}

	return false
}

// SetIsAutoAddMargin gets a reference to the given string and assigns it to the IsAutoAddMargin field.
func (o *UmFuturesSymbolConfigurationResponseInner) SetIsAutoAddMargin(v string) {
	o.IsAutoAddMargin = &v
}

// GetLeverage returns the Leverage field value if set, zero value otherwise.
func (o *UmFuturesSymbolConfigurationResponseInner) GetLeverage() int64 {
	if o == nil || common.IsNil(o.Leverage) {
		var ret int64
		return ret
	}
	return *o.Leverage
}

// GetLeverageOk returns a tuple with the Leverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmFuturesSymbolConfigurationResponseInner) GetLeverageOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Leverage) {
		return nil, false
	}
	return o.Leverage, true
}

// HasLeverage returns a boolean if a field has been set.
func (o *UmFuturesSymbolConfigurationResponseInner) HasLeverage() bool {
	if o != nil && !common.IsNil(o.Leverage) {
		return true
	}

	return false
}

// SetLeverage gets a reference to the given int64 and assigns it to the Leverage field.
func (o *UmFuturesSymbolConfigurationResponseInner) SetLeverage(v int64) {
	o.Leverage = &v
}

// GetMaxNotionalValue returns the MaxNotionalValue field value if set, zero value otherwise.
func (o *UmFuturesSymbolConfigurationResponseInner) GetMaxNotionalValue() string {
	if o == nil || common.IsNil(o.MaxNotionalValue) {
		var ret string
		return ret
	}
	return *o.MaxNotionalValue
}

// GetMaxNotionalValueOk returns a tuple with the MaxNotionalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmFuturesSymbolConfigurationResponseInner) GetMaxNotionalValueOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxNotionalValue) {
		return nil, false
	}
	return o.MaxNotionalValue, true
}

// HasMaxNotionalValue returns a boolean if a field has been set.
func (o *UmFuturesSymbolConfigurationResponseInner) HasMaxNotionalValue() bool {
	if o != nil && !common.IsNil(o.MaxNotionalValue) {
		return true
	}

	return false
}

// SetMaxNotionalValue gets a reference to the given string and assigns it to the MaxNotionalValue field.
func (o *UmFuturesSymbolConfigurationResponseInner) SetMaxNotionalValue(v string) {
	o.MaxNotionalValue = &v
}

func (o UmFuturesSymbolConfigurationResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UmFuturesSymbolConfigurationResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.MarginType) {
		toSerialize["marginType"] = o.MarginType
	}
	if !common.IsNil(o.IsAutoAddMargin) {
		toSerialize["isAutoAddMargin"] = o.IsAutoAddMargin
	}
	if !common.IsNil(o.Leverage) {
		toSerialize["leverage"] = o.Leverage
	}
	if !common.IsNil(o.MaxNotionalValue) {
		toSerialize["maxNotionalValue"] = o.MaxNotionalValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UmFuturesSymbolConfigurationResponseInner) UnmarshalJSON(data []byte) (err error) {
	varUmFuturesSymbolConfigurationResponseInner := _UmFuturesSymbolConfigurationResponseInner{}

	err = json.Unmarshal(data, &varUmFuturesSymbolConfigurationResponseInner)

	if err != nil {
		return err
	}

	*o = UmFuturesSymbolConfigurationResponseInner(varUmFuturesSymbolConfigurationResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "marginType")
		delete(additionalProperties, "isAutoAddMargin")
		delete(additionalProperties, "leverage")
		delete(additionalProperties, "maxNotionalValue")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUmFuturesSymbolConfigurationResponseInner struct {
	value *UmFuturesSymbolConfigurationResponseInner
	isSet bool
}

func (v NullableUmFuturesSymbolConfigurationResponseInner) Get() *UmFuturesSymbolConfigurationResponseInner {
	return v.value
}

func (v *NullableUmFuturesSymbolConfigurationResponseInner) Set(val *UmFuturesSymbolConfigurationResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUmFuturesSymbolConfigurationResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUmFuturesSymbolConfigurationResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUmFuturesSymbolConfigurationResponseInner(val *UmFuturesSymbolConfigurationResponseInner) *NullableUmFuturesSymbolConfigurationResponseInner {
	return &NullableUmFuturesSymbolConfigurationResponseInner{value: val, isSet: true}
}

func (v NullableUmFuturesSymbolConfigurationResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUmFuturesSymbolConfigurationResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
