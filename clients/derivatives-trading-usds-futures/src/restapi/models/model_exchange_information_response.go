/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the ExchangeInformationResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ExchangeInformationResponse{}

// ExchangeInformationResponse struct for ExchangeInformationResponse
type ExchangeInformationResponse struct {
	ExchangeFilters      []string                                     `json:"exchangeFilters,omitempty"`
	RateLimits           []ExchangeInformationResponseRateLimitsInner `json:"rateLimits,omitempty"`
	ServerTime           *int64                                       `json:"serverTime,omitempty"`
	Assets               []ExchangeInformationResponseAssetsInner     `json:"assets,omitempty"`
	Symbols              []ExchangeInformationResponseSymbolsInner    `json:"symbols,omitempty"`
	Timezone             *string                                      `json:"timezone,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExchangeInformationResponse ExchangeInformationResponse

// NewExchangeInformationResponse instantiates a new ExchangeInformationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeInformationResponse() *ExchangeInformationResponse {
	this := ExchangeInformationResponse{}
	return &this
}

// NewExchangeInformationResponseWithDefaults instantiates a new ExchangeInformationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeInformationResponseWithDefaults() *ExchangeInformationResponse {
	this := ExchangeInformationResponse{}
	return &this
}

// GetExchangeFilters returns the ExchangeFilters field value if set, zero value otherwise.
func (o *ExchangeInformationResponse) GetExchangeFilters() []string {
	if o == nil || common.IsNil(o.ExchangeFilters) {
		var ret []string
		return ret
	}
	return o.ExchangeFilters
}

// GetExchangeFiltersOk returns a tuple with the ExchangeFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponse) GetExchangeFiltersOk() ([]string, bool) {
	if o == nil || common.IsNil(o.ExchangeFilters) {
		return nil, false
	}
	return o.ExchangeFilters, true
}

// HasExchangeFilters returns a boolean if a field has been set.
func (o *ExchangeInformationResponse) HasExchangeFilters() bool {
	if o != nil && !common.IsNil(o.ExchangeFilters) {
		return true
	}

	return false
}

// SetExchangeFilters gets a reference to the given []string and assigns it to the ExchangeFilters field.
func (o *ExchangeInformationResponse) SetExchangeFilters(v []string) {
	o.ExchangeFilters = v
}

// GetRateLimits returns the RateLimits field value if set, zero value otherwise.
func (o *ExchangeInformationResponse) GetRateLimits() []ExchangeInformationResponseRateLimitsInner {
	if o == nil || common.IsNil(o.RateLimits) {
		var ret []ExchangeInformationResponseRateLimitsInner
		return ret
	}
	return o.RateLimits
}

// GetRateLimitsOk returns a tuple with the RateLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponse) GetRateLimitsOk() ([]ExchangeInformationResponseRateLimitsInner, bool) {
	if o == nil || common.IsNil(o.RateLimits) {
		return nil, false
	}
	return o.RateLimits, true
}

// HasRateLimits returns a boolean if a field has been set.
func (o *ExchangeInformationResponse) HasRateLimits() bool {
	if o != nil && !common.IsNil(o.RateLimits) {
		return true
	}

	return false
}

// SetRateLimits gets a reference to the given []ExchangeInformationResponseRateLimitsInner and assigns it to the RateLimits field.
func (o *ExchangeInformationResponse) SetRateLimits(v []ExchangeInformationResponseRateLimitsInner) {
	o.RateLimits = v
}

// GetServerTime returns the ServerTime field value if set, zero value otherwise.
func (o *ExchangeInformationResponse) GetServerTime() int64 {
	if o == nil || common.IsNil(o.ServerTime) {
		var ret int64
		return ret
	}
	return *o.ServerTime
}

// GetServerTimeOk returns a tuple with the ServerTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponse) GetServerTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ServerTime) {
		return nil, false
	}
	return o.ServerTime, true
}

// HasServerTime returns a boolean if a field has been set.
func (o *ExchangeInformationResponse) HasServerTime() bool {
	if o != nil && !common.IsNil(o.ServerTime) {
		return true
	}

	return false
}

// SetServerTime gets a reference to the given int64 and assigns it to the ServerTime field.
func (o *ExchangeInformationResponse) SetServerTime(v int64) {
	o.ServerTime = &v
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *ExchangeInformationResponse) GetAssets() []ExchangeInformationResponseAssetsInner {
	if o == nil || common.IsNil(o.Assets) {
		var ret []ExchangeInformationResponseAssetsInner
		return ret
	}
	return o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponse) GetAssetsOk() ([]ExchangeInformationResponseAssetsInner, bool) {
	if o == nil || common.IsNil(o.Assets) {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *ExchangeInformationResponse) HasAssets() bool {
	if o != nil && !common.IsNil(o.Assets) {
		return true
	}

	return false
}

// SetAssets gets a reference to the given []ExchangeInformationResponseAssetsInner and assigns it to the Assets field.
func (o *ExchangeInformationResponse) SetAssets(v []ExchangeInformationResponseAssetsInner) {
	o.Assets = v
}

// GetSymbols returns the Symbols field value if set, zero value otherwise.
func (o *ExchangeInformationResponse) GetSymbols() []ExchangeInformationResponseSymbolsInner {
	if o == nil || common.IsNil(o.Symbols) {
		var ret []ExchangeInformationResponseSymbolsInner
		return ret
	}
	return o.Symbols
}

// GetSymbolsOk returns a tuple with the Symbols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponse) GetSymbolsOk() ([]ExchangeInformationResponseSymbolsInner, bool) {
	if o == nil || common.IsNil(o.Symbols) {
		return nil, false
	}
	return o.Symbols, true
}

// HasSymbols returns a boolean if a field has been set.
func (o *ExchangeInformationResponse) HasSymbols() bool {
	if o != nil && !common.IsNil(o.Symbols) {
		return true
	}

	return false
}

// SetSymbols gets a reference to the given []ExchangeInformationResponseSymbolsInner and assigns it to the Symbols field.
func (o *ExchangeInformationResponse) SetSymbols(v []ExchangeInformationResponseSymbolsInner) {
	o.Symbols = v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *ExchangeInformationResponse) GetTimezone() string {
	if o == nil || common.IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponse) GetTimezoneOk() (*string, bool) {
	if o == nil || common.IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *ExchangeInformationResponse) HasTimezone() bool {
	if o != nil && !common.IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *ExchangeInformationResponse) SetTimezone(v string) {
	o.Timezone = &v
}

func (o ExchangeInformationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeInformationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ExchangeFilters) {
		toSerialize["exchangeFilters"] = o.ExchangeFilters
	}
	if !common.IsNil(o.RateLimits) {
		toSerialize["rateLimits"] = o.RateLimits
	}
	if !common.IsNil(o.ServerTime) {
		toSerialize["serverTime"] = o.ServerTime
	}
	if !common.IsNil(o.Assets) {
		toSerialize["assets"] = o.Assets
	}
	if !common.IsNil(o.Symbols) {
		toSerialize["symbols"] = o.Symbols
	}
	if !common.IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExchangeInformationResponse) UnmarshalJSON(data []byte) (err error) {
	varExchangeInformationResponse := _ExchangeInformationResponse{}

	err = json.Unmarshal(data, &varExchangeInformationResponse)

	if err != nil {
		return err
	}

	*o = ExchangeInformationResponse(varExchangeInformationResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "exchangeFilters")
		delete(additionalProperties, "rateLimits")
		delete(additionalProperties, "serverTime")
		delete(additionalProperties, "assets")
		delete(additionalProperties, "symbols")
		delete(additionalProperties, "timezone")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExchangeInformationResponse struct {
	value *ExchangeInformationResponse
	isSet bool
}

func (v NullableExchangeInformationResponse) Get() *ExchangeInformationResponse {
	return v.value
}

func (v *NullableExchangeInformationResponse) Set(val *ExchangeInformationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeInformationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeInformationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeInformationResponse(val *ExchangeInformationResponse) *NullableExchangeInformationResponse {
	return &NullableExchangeInformationResponse{value: val, isSet: true}
}

func (v NullableExchangeInformationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeInformationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
