/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the HistoricalExerciseRecordsResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &HistoricalExerciseRecordsResponseInner{}

// HistoricalExerciseRecordsResponseInner struct for HistoricalExerciseRecordsResponseInner
type HistoricalExerciseRecordsResponseInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	StrikePrice          *string `json:"strikePrice,omitempty"`
	RealStrikePrice      *string `json:"realStrikePrice,omitempty"`
	ExpiryDate           *int64  `json:"expiryDate,omitempty"`
	StrikeResult         *string `json:"strikeResult,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HistoricalExerciseRecordsResponseInner HistoricalExerciseRecordsResponseInner

// NewHistoricalExerciseRecordsResponseInner instantiates a new HistoricalExerciseRecordsResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoricalExerciseRecordsResponseInner() *HistoricalExerciseRecordsResponseInner {
	this := HistoricalExerciseRecordsResponseInner{}
	return &this
}

// NewHistoricalExerciseRecordsResponseInnerWithDefaults instantiates a new HistoricalExerciseRecordsResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoricalExerciseRecordsResponseInnerWithDefaults() *HistoricalExerciseRecordsResponseInner {
	this := HistoricalExerciseRecordsResponseInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *HistoricalExerciseRecordsResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalExerciseRecordsResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *HistoricalExerciseRecordsResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *HistoricalExerciseRecordsResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetStrikePrice returns the StrikePrice field value if set, zero value otherwise.
func (o *HistoricalExerciseRecordsResponseInner) GetStrikePrice() string {
	if o == nil || common.IsNil(o.StrikePrice) {
		var ret string
		return ret
	}
	return *o.StrikePrice
}

// GetStrikePriceOk returns a tuple with the StrikePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalExerciseRecordsResponseInner) GetStrikePriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.StrikePrice) {
		return nil, false
	}
	return o.StrikePrice, true
}

// HasStrikePrice returns a boolean if a field has been set.
func (o *HistoricalExerciseRecordsResponseInner) HasStrikePrice() bool {
	if o != nil && !common.IsNil(o.StrikePrice) {
		return true
	}

	return false
}

// SetStrikePrice gets a reference to the given string and assigns it to the StrikePrice field.
func (o *HistoricalExerciseRecordsResponseInner) SetStrikePrice(v string) {
	o.StrikePrice = &v
}

// GetRealStrikePrice returns the RealStrikePrice field value if set, zero value otherwise.
func (o *HistoricalExerciseRecordsResponseInner) GetRealStrikePrice() string {
	if o == nil || common.IsNil(o.RealStrikePrice) {
		var ret string
		return ret
	}
	return *o.RealStrikePrice
}

// GetRealStrikePriceOk returns a tuple with the RealStrikePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalExerciseRecordsResponseInner) GetRealStrikePriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.RealStrikePrice) {
		return nil, false
	}
	return o.RealStrikePrice, true
}

// HasRealStrikePrice returns a boolean if a field has been set.
func (o *HistoricalExerciseRecordsResponseInner) HasRealStrikePrice() bool {
	if o != nil && !common.IsNil(o.RealStrikePrice) {
		return true
	}

	return false
}

// SetRealStrikePrice gets a reference to the given string and assigns it to the RealStrikePrice field.
func (o *HistoricalExerciseRecordsResponseInner) SetRealStrikePrice(v string) {
	o.RealStrikePrice = &v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *HistoricalExerciseRecordsResponseInner) GetExpiryDate() int64 {
	if o == nil || common.IsNil(o.ExpiryDate) {
		var ret int64
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalExerciseRecordsResponseInner) GetExpiryDateOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ExpiryDate) {
		return nil, false
	}
	return o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *HistoricalExerciseRecordsResponseInner) HasExpiryDate() bool {
	if o != nil && !common.IsNil(o.ExpiryDate) {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given int64 and assigns it to the ExpiryDate field.
func (o *HistoricalExerciseRecordsResponseInner) SetExpiryDate(v int64) {
	o.ExpiryDate = &v
}

// GetStrikeResult returns the StrikeResult field value if set, zero value otherwise.
func (o *HistoricalExerciseRecordsResponseInner) GetStrikeResult() string {
	if o == nil || common.IsNil(o.StrikeResult) {
		var ret string
		return ret
	}
	return *o.StrikeResult
}

// GetStrikeResultOk returns a tuple with the StrikeResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalExerciseRecordsResponseInner) GetStrikeResultOk() (*string, bool) {
	if o == nil || common.IsNil(o.StrikeResult) {
		return nil, false
	}
	return o.StrikeResult, true
}

// HasStrikeResult returns a boolean if a field has been set.
func (o *HistoricalExerciseRecordsResponseInner) HasStrikeResult() bool {
	if o != nil && !common.IsNil(o.StrikeResult) {
		return true
	}

	return false
}

// SetStrikeResult gets a reference to the given string and assigns it to the StrikeResult field.
func (o *HistoricalExerciseRecordsResponseInner) SetStrikeResult(v string) {
	o.StrikeResult = &v
}

func (o HistoricalExerciseRecordsResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoricalExerciseRecordsResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.StrikePrice) {
		toSerialize["strikePrice"] = o.StrikePrice
	}
	if !common.IsNil(o.RealStrikePrice) {
		toSerialize["realStrikePrice"] = o.RealStrikePrice
	}
	if !common.IsNil(o.ExpiryDate) {
		toSerialize["expiryDate"] = o.ExpiryDate
	}
	if !common.IsNil(o.StrikeResult) {
		toSerialize["strikeResult"] = o.StrikeResult
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HistoricalExerciseRecordsResponseInner) UnmarshalJSON(data []byte) (err error) {
	varHistoricalExerciseRecordsResponseInner := _HistoricalExerciseRecordsResponseInner{}

	err = json.Unmarshal(data, &varHistoricalExerciseRecordsResponseInner)

	if err != nil {
		return err
	}

	*o = HistoricalExerciseRecordsResponseInner(varHistoricalExerciseRecordsResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "strikePrice")
		delete(additionalProperties, "realStrikePrice")
		delete(additionalProperties, "expiryDate")
		delete(additionalProperties, "strikeResult")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHistoricalExerciseRecordsResponseInner struct {
	value *HistoricalExerciseRecordsResponseInner
	isSet bool
}

func (v NullableHistoricalExerciseRecordsResponseInner) Get() *HistoricalExerciseRecordsResponseInner {
	return v.value
}

func (v *NullableHistoricalExerciseRecordsResponseInner) Set(val *HistoricalExerciseRecordsResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoricalExerciseRecordsResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoricalExerciseRecordsResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoricalExerciseRecordsResponseInner(val *HistoricalExerciseRecordsResponseInner) *NullableHistoricalExerciseRecordsResponseInner {
	return &NullableHistoricalExerciseRecordsResponseInner{value: val, isSet: true}
}

func (v NullableHistoricalExerciseRecordsResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoricalExerciseRecordsResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
