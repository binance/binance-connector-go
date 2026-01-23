/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the SymbolConfigurationResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SymbolConfigurationResponseInner{}

// SymbolConfigurationResponseInner struct for SymbolConfigurationResponseInner
type SymbolConfigurationResponseInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	MarginType           *string `json:"marginType,omitempty"`
	IsAutoAddMargin      *bool   `json:"isAutoAddMargin,omitempty"`
	Leverage             *int64  `json:"leverage,omitempty"`
	MaxNotionalValue     *string `json:"maxNotionalValue,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SymbolConfigurationResponseInner SymbolConfigurationResponseInner

// NewSymbolConfigurationResponseInner instantiates a new SymbolConfigurationResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSymbolConfigurationResponseInner() *SymbolConfigurationResponseInner {
	this := SymbolConfigurationResponseInner{}
	return &this
}

// NewSymbolConfigurationResponseInnerWithDefaults instantiates a new SymbolConfigurationResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSymbolConfigurationResponseInnerWithDefaults() *SymbolConfigurationResponseInner {
	this := SymbolConfigurationResponseInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *SymbolConfigurationResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolConfigurationResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *SymbolConfigurationResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *SymbolConfigurationResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetMarginType returns the MarginType field value if set, zero value otherwise.
func (o *SymbolConfigurationResponseInner) GetMarginType() string {
	if o == nil || common.IsNil(o.MarginType) {
		var ret string
		return ret
	}
	return *o.MarginType
}

// GetMarginTypeOk returns a tuple with the MarginType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolConfigurationResponseInner) GetMarginTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarginType) {
		return nil, false
	}
	return o.MarginType, true
}

// HasMarginType returns a boolean if a field has been set.
func (o *SymbolConfigurationResponseInner) HasMarginType() bool {
	if o != nil && !common.IsNil(o.MarginType) {
		return true
	}

	return false
}

// SetMarginType gets a reference to the given string and assigns it to the MarginType field.
func (o *SymbolConfigurationResponseInner) SetMarginType(v string) {
	o.MarginType = &v
}

// GetIsAutoAddMargin returns the IsAutoAddMargin field value if set, zero value otherwise.
func (o *SymbolConfigurationResponseInner) GetIsAutoAddMargin() bool {
	if o == nil || common.IsNil(o.IsAutoAddMargin) {
		var ret bool
		return ret
	}
	return *o.IsAutoAddMargin
}

// GetIsAutoAddMarginOk returns a tuple with the IsAutoAddMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolConfigurationResponseInner) GetIsAutoAddMarginOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsAutoAddMargin) {
		return nil, false
	}
	return o.IsAutoAddMargin, true
}

// HasIsAutoAddMargin returns a boolean if a field has been set.
func (o *SymbolConfigurationResponseInner) HasIsAutoAddMargin() bool {
	if o != nil && !common.IsNil(o.IsAutoAddMargin) {
		return true
	}

	return false
}

// SetIsAutoAddMargin gets a reference to the given bool and assigns it to the IsAutoAddMargin field.
func (o *SymbolConfigurationResponseInner) SetIsAutoAddMargin(v bool) {
	o.IsAutoAddMargin = &v
}

// GetLeverage returns the Leverage field value if set, zero value otherwise.
func (o *SymbolConfigurationResponseInner) GetLeverage() int64 {
	if o == nil || common.IsNil(o.Leverage) {
		var ret int64
		return ret
	}
	return *o.Leverage
}

// GetLeverageOk returns a tuple with the Leverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolConfigurationResponseInner) GetLeverageOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Leverage) {
		return nil, false
	}
	return o.Leverage, true
}

// HasLeverage returns a boolean if a field has been set.
func (o *SymbolConfigurationResponseInner) HasLeverage() bool {
	if o != nil && !common.IsNil(o.Leverage) {
		return true
	}

	return false
}

// SetLeverage gets a reference to the given int64 and assigns it to the Leverage field.
func (o *SymbolConfigurationResponseInner) SetLeverage(v int64) {
	o.Leverage = &v
}

// GetMaxNotionalValue returns the MaxNotionalValue field value if set, zero value otherwise.
func (o *SymbolConfigurationResponseInner) GetMaxNotionalValue() string {
	if o == nil || common.IsNil(o.MaxNotionalValue) {
		var ret string
		return ret
	}
	return *o.MaxNotionalValue
}

// GetMaxNotionalValueOk returns a tuple with the MaxNotionalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolConfigurationResponseInner) GetMaxNotionalValueOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxNotionalValue) {
		return nil, false
	}
	return o.MaxNotionalValue, true
}

// HasMaxNotionalValue returns a boolean if a field has been set.
func (o *SymbolConfigurationResponseInner) HasMaxNotionalValue() bool {
	if o != nil && !common.IsNil(o.MaxNotionalValue) {
		return true
	}

	return false
}

// SetMaxNotionalValue gets a reference to the given string and assigns it to the MaxNotionalValue field.
func (o *SymbolConfigurationResponseInner) SetMaxNotionalValue(v string) {
	o.MaxNotionalValue = &v
}

func (o SymbolConfigurationResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SymbolConfigurationResponseInner) ToMap() (map[string]interface{}, error) {
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

func (o *SymbolConfigurationResponseInner) UnmarshalJSON(data []byte) (err error) {
	varSymbolConfigurationResponseInner := _SymbolConfigurationResponseInner{}

	err = json.Unmarshal(data, &varSymbolConfigurationResponseInner)

	if err != nil {
		return err
	}

	*o = SymbolConfigurationResponseInner(varSymbolConfigurationResponseInner)

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

type NullableSymbolConfigurationResponseInner struct {
	value *SymbolConfigurationResponseInner
	isSet bool
}

func (v NullableSymbolConfigurationResponseInner) Get() *SymbolConfigurationResponseInner {
	return v.value
}

func (v *NullableSymbolConfigurationResponseInner) Set(val *SymbolConfigurationResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSymbolConfigurationResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSymbolConfigurationResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSymbolConfigurationResponseInner(val *SymbolConfigurationResponseInner) *NullableSymbolConfigurationResponseInner {
	return &NullableSymbolConfigurationResponseInner{value: val, isSet: true}
}

func (v NullableSymbolConfigurationResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSymbolConfigurationResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
