/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryIsolatedMarginFeeDataResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryIsolatedMarginFeeDataResponseInner{}

// QueryIsolatedMarginFeeDataResponseInner struct for QueryIsolatedMarginFeeDataResponseInner
type QueryIsolatedMarginFeeDataResponseInner struct {
	VipLevel             *int64                                             `json:"vipLevel,omitempty"`
	Symbol               *string                                            `json:"symbol,omitempty"`
	Leverage             *string                                            `json:"leverage,omitempty"`
	Data                 []QueryIsolatedMarginFeeDataResponseInnerDataInner `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryIsolatedMarginFeeDataResponseInner QueryIsolatedMarginFeeDataResponseInner

// NewQueryIsolatedMarginFeeDataResponseInner instantiates a new QueryIsolatedMarginFeeDataResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryIsolatedMarginFeeDataResponseInner() *QueryIsolatedMarginFeeDataResponseInner {
	this := QueryIsolatedMarginFeeDataResponseInner{}
	return &this
}

// NewQueryIsolatedMarginFeeDataResponseInnerWithDefaults instantiates a new QueryIsolatedMarginFeeDataResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryIsolatedMarginFeeDataResponseInnerWithDefaults() *QueryIsolatedMarginFeeDataResponseInner {
	this := QueryIsolatedMarginFeeDataResponseInner{}
	return &this
}

// GetVipLevel returns the VipLevel field value if set, zero value otherwise.
func (o *QueryIsolatedMarginFeeDataResponseInner) GetVipLevel() int64 {
	if o == nil || common.IsNil(o.VipLevel) {
		var ret int64
		return ret
	}
	return *o.VipLevel
}

// GetVipLevelOk returns a tuple with the VipLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginFeeDataResponseInner) GetVipLevelOk() (*int64, bool) {
	if o == nil || common.IsNil(o.VipLevel) {
		return nil, false
	}
	return o.VipLevel, true
}

// HasVipLevel returns a boolean if a field has been set.
func (o *QueryIsolatedMarginFeeDataResponseInner) HasVipLevel() bool {
	if o != nil && !common.IsNil(o.VipLevel) {
		return true
	}

	return false
}

// SetVipLevel gets a reference to the given int64 and assigns it to the VipLevel field.
func (o *QueryIsolatedMarginFeeDataResponseInner) SetVipLevel(v int64) {
	o.VipLevel = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *QueryIsolatedMarginFeeDataResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginFeeDataResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *QueryIsolatedMarginFeeDataResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *QueryIsolatedMarginFeeDataResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetLeverage returns the Leverage field value if set, zero value otherwise.
func (o *QueryIsolatedMarginFeeDataResponseInner) GetLeverage() string {
	if o == nil || common.IsNil(o.Leverage) {
		var ret string
		return ret
	}
	return *o.Leverage
}

// GetLeverageOk returns a tuple with the Leverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginFeeDataResponseInner) GetLeverageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Leverage) {
		return nil, false
	}
	return o.Leverage, true
}

// HasLeverage returns a boolean if a field has been set.
func (o *QueryIsolatedMarginFeeDataResponseInner) HasLeverage() bool {
	if o != nil && !common.IsNil(o.Leverage) {
		return true
	}

	return false
}

// SetLeverage gets a reference to the given string and assigns it to the Leverage field.
func (o *QueryIsolatedMarginFeeDataResponseInner) SetLeverage(v string) {
	o.Leverage = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *QueryIsolatedMarginFeeDataResponseInner) GetData() []QueryIsolatedMarginFeeDataResponseInnerDataInner {
	if o == nil || common.IsNil(o.Data) {
		var ret []QueryIsolatedMarginFeeDataResponseInnerDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIsolatedMarginFeeDataResponseInner) GetDataOk() ([]QueryIsolatedMarginFeeDataResponseInnerDataInner, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *QueryIsolatedMarginFeeDataResponseInner) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []QueryIsolatedMarginFeeDataResponseInnerDataInner and assigns it to the Data field.
func (o *QueryIsolatedMarginFeeDataResponseInner) SetData(v []QueryIsolatedMarginFeeDataResponseInnerDataInner) {
	o.Data = v
}

func (o QueryIsolatedMarginFeeDataResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryIsolatedMarginFeeDataResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.VipLevel) {
		toSerialize["vipLevel"] = o.VipLevel
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Leverage) {
		toSerialize["leverage"] = o.Leverage
	}
	if !common.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryIsolatedMarginFeeDataResponseInner) UnmarshalJSON(data []byte) (err error) {
	varQueryIsolatedMarginFeeDataResponseInner := _QueryIsolatedMarginFeeDataResponseInner{}

	err = json.Unmarshal(data, &varQueryIsolatedMarginFeeDataResponseInner)

	if err != nil {
		return err
	}

	*o = QueryIsolatedMarginFeeDataResponseInner(varQueryIsolatedMarginFeeDataResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "vipLevel")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "leverage")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryIsolatedMarginFeeDataResponseInner struct {
	value *QueryIsolatedMarginFeeDataResponseInner
	isSet bool
}

func (v NullableQueryIsolatedMarginFeeDataResponseInner) Get() *QueryIsolatedMarginFeeDataResponseInner {
	return v.value
}

func (v *NullableQueryIsolatedMarginFeeDataResponseInner) Set(val *QueryIsolatedMarginFeeDataResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryIsolatedMarginFeeDataResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryIsolatedMarginFeeDataResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryIsolatedMarginFeeDataResponseInner(val *QueryIsolatedMarginFeeDataResponseInner) *NullableQueryIsolatedMarginFeeDataResponseInner {
	return &NullableQueryIsolatedMarginFeeDataResponseInner{value: val, isSet: true}
}

func (v NullableQueryIsolatedMarginFeeDataResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryIsolatedMarginFeeDataResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
