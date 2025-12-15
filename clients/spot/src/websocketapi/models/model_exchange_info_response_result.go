/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the ExchangeInfoResponseResult type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ExchangeInfoResponseResult{}

// ExchangeInfoResponseResult struct for ExchangeInfoResponseResult
type ExchangeInfoResponseResult struct {
	Timezone             *string                                  `json:"timezone,omitempty"`
	ServerTime           *int64                                   `json:"serverTime,omitempty"`
	RateLimits           []RateLimits                             `json:"rateLimits,omitempty"`
	ExchangeFilters      []ExchangeFilters                        `json:"exchangeFilters,omitempty"`
	Symbols              []ExchangeInfoResponseResultSymbolsInner `json:"symbols,omitempty"`
	Sors                 []ExchangeInfoResponseResultSorsInner    `json:"sors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExchangeInfoResponseResult ExchangeInfoResponseResult

// NewExchangeInfoResponseResult instantiates a new ExchangeInfoResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeInfoResponseResult() *ExchangeInfoResponseResult {
	this := ExchangeInfoResponseResult{}
	return &this
}

// NewExchangeInfoResponseResultWithDefaults instantiates a new ExchangeInfoResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeInfoResponseResultWithDefaults() *ExchangeInfoResponseResult {
	this := ExchangeInfoResponseResult{}
	return &this
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *ExchangeInfoResponseResult) GetTimezone() string {
	if o == nil || common.IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseResult) GetTimezoneOk() (*string, bool) {
	if o == nil || common.IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *ExchangeInfoResponseResult) HasTimezone() bool {
	if o != nil && !common.IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *ExchangeInfoResponseResult) SetTimezone(v string) {
	o.Timezone = &v
}

// GetServerTime returns the ServerTime field value if set, zero value otherwise.
func (o *ExchangeInfoResponseResult) GetServerTime() int64 {
	if o == nil || common.IsNil(o.ServerTime) {
		var ret int64
		return ret
	}
	return *o.ServerTime
}

// GetServerTimeOk returns a tuple with the ServerTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseResult) GetServerTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ServerTime) {
		return nil, false
	}
	return o.ServerTime, true
}

// HasServerTime returns a boolean if a field has been set.
func (o *ExchangeInfoResponseResult) HasServerTime() bool {
	if o != nil && !common.IsNil(o.ServerTime) {
		return true
	}

	return false
}

// SetServerTime gets a reference to the given int64 and assigns it to the ServerTime field.
func (o *ExchangeInfoResponseResult) SetServerTime(v int64) {
	o.ServerTime = &v
}

// GetRateLimits returns the RateLimits field value if set, zero value otherwise.
func (o *ExchangeInfoResponseResult) GetRateLimits() []RateLimits {
	if o == nil || common.IsNil(o.RateLimits) {
		var ret []RateLimits
		return ret
	}
	return o.RateLimits
}

// GetRateLimitsOk returns a tuple with the RateLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseResult) GetRateLimitsOk() ([]RateLimits, bool) {
	if o == nil || common.IsNil(o.RateLimits) {
		return nil, false
	}
	return o.RateLimits, true
}

// HasRateLimits returns a boolean if a field has been set.
func (o *ExchangeInfoResponseResult) HasRateLimits() bool {
	if o != nil && !common.IsNil(o.RateLimits) {
		return true
	}

	return false
}

// SetRateLimits gets a reference to the given []RateLimits and assigns it to the RateLimits field.
func (o *ExchangeInfoResponseResult) SetRateLimits(v []RateLimits) {
	o.RateLimits = v
}

// GetExchangeFilters returns the ExchangeFilters field value if set, zero value otherwise.
func (o *ExchangeInfoResponseResult) GetExchangeFilters() []ExchangeFilters {
	if o == nil || common.IsNil(o.ExchangeFilters) {
		var ret []ExchangeFilters
		return ret
	}
	return o.ExchangeFilters
}

// GetExchangeFiltersOk returns a tuple with the ExchangeFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseResult) GetExchangeFiltersOk() ([]ExchangeFilters, bool) {
	if o == nil || common.IsNil(o.ExchangeFilters) {
		return nil, false
	}
	return o.ExchangeFilters, true
}

// HasExchangeFilters returns a boolean if a field has been set.
func (o *ExchangeInfoResponseResult) HasExchangeFilters() bool {
	if o != nil && !common.IsNil(o.ExchangeFilters) {
		return true
	}

	return false
}

// SetExchangeFilters gets a reference to the given []ExchangeFilters and assigns it to the ExchangeFilters field.
func (o *ExchangeInfoResponseResult) SetExchangeFilters(v []ExchangeFilters) {
	o.ExchangeFilters = v
}

// GetSymbols returns the Symbols field value if set, zero value otherwise.
func (o *ExchangeInfoResponseResult) GetSymbols() []ExchangeInfoResponseResultSymbolsInner {
	if o == nil || common.IsNil(o.Symbols) {
		var ret []ExchangeInfoResponseResultSymbolsInner
		return ret
	}
	return o.Symbols
}

// GetSymbolsOk returns a tuple with the Symbols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseResult) GetSymbolsOk() ([]ExchangeInfoResponseResultSymbolsInner, bool) {
	if o == nil || common.IsNil(o.Symbols) {
		return nil, false
	}
	return o.Symbols, true
}

// HasSymbols returns a boolean if a field has been set.
func (o *ExchangeInfoResponseResult) HasSymbols() bool {
	if o != nil && !common.IsNil(o.Symbols) {
		return true
	}

	return false
}

// SetSymbols gets a reference to the given []ExchangeInfoResponseResultSymbolsInner and assigns it to the Symbols field.
func (o *ExchangeInfoResponseResult) SetSymbols(v []ExchangeInfoResponseResultSymbolsInner) {
	o.Symbols = v
}

// GetSors returns the Sors field value if set, zero value otherwise.
func (o *ExchangeInfoResponseResult) GetSors() []ExchangeInfoResponseResultSorsInner {
	if o == nil || common.IsNil(o.Sors) {
		var ret []ExchangeInfoResponseResultSorsInner
		return ret
	}
	return o.Sors
}

// GetSorsOk returns a tuple with the Sors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponseResult) GetSorsOk() ([]ExchangeInfoResponseResultSorsInner, bool) {
	if o == nil || common.IsNil(o.Sors) {
		return nil, false
	}
	return o.Sors, true
}

// HasSors returns a boolean if a field has been set.
func (o *ExchangeInfoResponseResult) HasSors() bool {
	if o != nil && !common.IsNil(o.Sors) {
		return true
	}

	return false
}

// SetSors gets a reference to the given []ExchangeInfoResponseResultSorsInner and assigns it to the Sors field.
func (o *ExchangeInfoResponseResult) SetSors(v []ExchangeInfoResponseResultSorsInner) {
	o.Sors = v
}

func (o ExchangeInfoResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeInfoResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	if !common.IsNil(o.ServerTime) {
		toSerialize["serverTime"] = o.ServerTime
	}
	if !common.IsNil(o.RateLimits) {
		toSerialize["rateLimits"] = o.RateLimits
	}
	if !common.IsNil(o.ExchangeFilters) {
		toSerialize["exchangeFilters"] = o.ExchangeFilters
	}
	if !common.IsNil(o.Symbols) {
		toSerialize["symbols"] = o.Symbols
	}
	if !common.IsNil(o.Sors) {
		toSerialize["sors"] = o.Sors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExchangeInfoResponseResult) UnmarshalJSON(data []byte) (err error) {
	varExchangeInfoResponseResult := _ExchangeInfoResponseResult{}

	err = json.Unmarshal(data, &varExchangeInfoResponseResult)

	if err != nil {
		return err
	}

	*o = ExchangeInfoResponseResult(varExchangeInfoResponseResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "timezone")
		delete(additionalProperties, "serverTime")
		delete(additionalProperties, "rateLimits")
		delete(additionalProperties, "exchangeFilters")
		delete(additionalProperties, "symbols")
		delete(additionalProperties, "sors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExchangeInfoResponseResult struct {
	value *ExchangeInfoResponseResult
	isSet bool
}

func (v NullableExchangeInfoResponseResult) Get() *ExchangeInfoResponseResult {
	return v.value
}

func (v *NullableExchangeInfoResponseResult) Set(val *ExchangeInfoResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeInfoResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeInfoResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeInfoResponseResult(val *ExchangeInfoResponseResult) *NullableExchangeInfoResponseResult {
	return &NullableExchangeInfoResponseResult{value: val, isSet: true}
}

func (v NullableExchangeInfoResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeInfoResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
