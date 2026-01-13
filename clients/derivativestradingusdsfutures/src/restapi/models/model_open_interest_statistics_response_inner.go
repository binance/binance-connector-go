/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OpenInterestStatisticsResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OpenInterestStatisticsResponseInner{}

// OpenInterestStatisticsResponseInner struct for OpenInterestStatisticsResponseInner
type OpenInterestStatisticsResponseInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	SumOpenInterest      *string `json:"sumOpenInterest,omitempty"`
	SumOpenInterestValue *string `json:"sumOpenInterestValue,omitempty"`
	CMCCirculatingSupply *string `json:"CMCCirculatingSupply,omitempty"`
	Timestamp            *string `json:"timestamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OpenInterestStatisticsResponseInner OpenInterestStatisticsResponseInner

// NewOpenInterestStatisticsResponseInner instantiates a new OpenInterestStatisticsResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenInterestStatisticsResponseInner() *OpenInterestStatisticsResponseInner {
	this := OpenInterestStatisticsResponseInner{}
	return &this
}

// NewOpenInterestStatisticsResponseInnerWithDefaults instantiates a new OpenInterestStatisticsResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenInterestStatisticsResponseInnerWithDefaults() *OpenInterestStatisticsResponseInner {
	this := OpenInterestStatisticsResponseInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *OpenInterestStatisticsResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenInterestStatisticsResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *OpenInterestStatisticsResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *OpenInterestStatisticsResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetSumOpenInterest returns the SumOpenInterest field value if set, zero value otherwise.
func (o *OpenInterestStatisticsResponseInner) GetSumOpenInterest() string {
	if o == nil || common.IsNil(o.SumOpenInterest) {
		var ret string
		return ret
	}
	return *o.SumOpenInterest
}

// GetSumOpenInterestOk returns a tuple with the SumOpenInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenInterestStatisticsResponseInner) GetSumOpenInterestOk() (*string, bool) {
	if o == nil || common.IsNil(o.SumOpenInterest) {
		return nil, false
	}
	return o.SumOpenInterest, true
}

// HasSumOpenInterest returns a boolean if a field has been set.
func (o *OpenInterestStatisticsResponseInner) HasSumOpenInterest() bool {
	if o != nil && !common.IsNil(o.SumOpenInterest) {
		return true
	}

	return false
}

// SetSumOpenInterest gets a reference to the given string and assigns it to the SumOpenInterest field.
func (o *OpenInterestStatisticsResponseInner) SetSumOpenInterest(v string) {
	o.SumOpenInterest = &v
}

// GetSumOpenInterestValue returns the SumOpenInterestValue field value if set, zero value otherwise.
func (o *OpenInterestStatisticsResponseInner) GetSumOpenInterestValue() string {
	if o == nil || common.IsNil(o.SumOpenInterestValue) {
		var ret string
		return ret
	}
	return *o.SumOpenInterestValue
}

// GetSumOpenInterestValueOk returns a tuple with the SumOpenInterestValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenInterestStatisticsResponseInner) GetSumOpenInterestValueOk() (*string, bool) {
	if o == nil || common.IsNil(o.SumOpenInterestValue) {
		return nil, false
	}
	return o.SumOpenInterestValue, true
}

// HasSumOpenInterestValue returns a boolean if a field has been set.
func (o *OpenInterestStatisticsResponseInner) HasSumOpenInterestValue() bool {
	if o != nil && !common.IsNil(o.SumOpenInterestValue) {
		return true
	}

	return false
}

// SetSumOpenInterestValue gets a reference to the given string and assigns it to the SumOpenInterestValue field.
func (o *OpenInterestStatisticsResponseInner) SetSumOpenInterestValue(v string) {
	o.SumOpenInterestValue = &v
}

// GetCMCCirculatingSupply returns the CMCCirculatingSupply field value if set, zero value otherwise.
func (o *OpenInterestStatisticsResponseInner) GetCMCCirculatingSupply() string {
	if o == nil || common.IsNil(o.CMCCirculatingSupply) {
		var ret string
		return ret
	}
	return *o.CMCCirculatingSupply
}

// GetCMCCirculatingSupplyOk returns a tuple with the CMCCirculatingSupply field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenInterestStatisticsResponseInner) GetCMCCirculatingSupplyOk() (*string, bool) {
	if o == nil || common.IsNil(o.CMCCirculatingSupply) {
		return nil, false
	}
	return o.CMCCirculatingSupply, true
}

// HasCMCCirculatingSupply returns a boolean if a field has been set.
func (o *OpenInterestStatisticsResponseInner) HasCMCCirculatingSupply() bool {
	if o != nil && !common.IsNil(o.CMCCirculatingSupply) {
		return true
	}

	return false
}

// SetCMCCirculatingSupply gets a reference to the given string and assigns it to the CMCCirculatingSupply field.
func (o *OpenInterestStatisticsResponseInner) SetCMCCirculatingSupply(v string) {
	o.CMCCirculatingSupply = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *OpenInterestStatisticsResponseInner) GetTimestamp() string {
	if o == nil || common.IsNil(o.Timestamp) {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenInterestStatisticsResponseInner) GetTimestampOk() (*string, bool) {
	if o == nil || common.IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *OpenInterestStatisticsResponseInner) HasTimestamp() bool {
	if o != nil && !common.IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *OpenInterestStatisticsResponseInner) SetTimestamp(v string) {
	o.Timestamp = &v
}

func (o OpenInterestStatisticsResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenInterestStatisticsResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.SumOpenInterest) {
		toSerialize["sumOpenInterest"] = o.SumOpenInterest
	}
	if !common.IsNil(o.SumOpenInterestValue) {
		toSerialize["sumOpenInterestValue"] = o.SumOpenInterestValue
	}
	if !common.IsNil(o.CMCCirculatingSupply) {
		toSerialize["CMCCirculatingSupply"] = o.CMCCirculatingSupply
	}
	if !common.IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OpenInterestStatisticsResponseInner) UnmarshalJSON(data []byte) (err error) {
	varOpenInterestStatisticsResponseInner := _OpenInterestStatisticsResponseInner{}

	err = json.Unmarshal(data, &varOpenInterestStatisticsResponseInner)

	if err != nil {
		return err
	}

	*o = OpenInterestStatisticsResponseInner(varOpenInterestStatisticsResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "sumOpenInterest")
		delete(additionalProperties, "sumOpenInterestValue")
		delete(additionalProperties, "CMCCirculatingSupply")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOpenInterestStatisticsResponseInner struct {
	value *OpenInterestStatisticsResponseInner
	isSet bool
}

func (v NullableOpenInterestStatisticsResponseInner) Get() *OpenInterestStatisticsResponseInner {
	return v.value
}

func (v *NullableOpenInterestStatisticsResponseInner) Set(val *OpenInterestStatisticsResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenInterestStatisticsResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenInterestStatisticsResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenInterestStatisticsResponseInner(val *OpenInterestStatisticsResponseInner) *NullableOpenInterestStatisticsResponseInner {
	return &NullableOpenInterestStatisticsResponseInner{value: val, isSet: true}
}

func (v NullableOpenInterestStatisticsResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenInterestStatisticsResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
