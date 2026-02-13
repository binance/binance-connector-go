/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ExchangeInformationResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ExchangeInformationResponse{}

// ExchangeInformationResponse struct for ExchangeInformationResponse
type ExchangeInformationResponse struct {
	Timezone             *string                                           `json:"timezone,omitempty"`
	ServerTime           *int64                                            `json:"serverTime,omitempty"`
	OptionContracts      []ExchangeInformationResponseOptionContractsInner `json:"optionContracts,omitempty"`
	OptionAssets         []ExchangeInformationResponseOptionAssetsInner    `json:"optionAssets,omitempty"`
	OptionSymbols        []ExchangeInformationResponseOptionSymbolsInner   `json:"optionSymbols,omitempty"`
	RateLimits           []ExchangeInformationResponseRateLimitsInner      `json:"rateLimits,omitempty"`
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

// GetOptionContracts returns the OptionContracts field value if set, zero value otherwise.
func (o *ExchangeInformationResponse) GetOptionContracts() []ExchangeInformationResponseOptionContractsInner {
	if o == nil || common.IsNil(o.OptionContracts) {
		var ret []ExchangeInformationResponseOptionContractsInner
		return ret
	}
	return o.OptionContracts
}

// GetOptionContractsOk returns a tuple with the OptionContracts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponse) GetOptionContractsOk() ([]ExchangeInformationResponseOptionContractsInner, bool) {
	if o == nil || common.IsNil(o.OptionContracts) {
		return nil, false
	}
	return o.OptionContracts, true
}

// HasOptionContracts returns a boolean if a field has been set.
func (o *ExchangeInformationResponse) HasOptionContracts() bool {
	if o != nil && !common.IsNil(o.OptionContracts) {
		return true
	}

	return false
}

// SetOptionContracts gets a reference to the given []ExchangeInformationResponseOptionContractsInner and assigns it to the OptionContracts field.
func (o *ExchangeInformationResponse) SetOptionContracts(v []ExchangeInformationResponseOptionContractsInner) {
	o.OptionContracts = v
}

// GetOptionAssets returns the OptionAssets field value if set, zero value otherwise.
func (o *ExchangeInformationResponse) GetOptionAssets() []ExchangeInformationResponseOptionAssetsInner {
	if o == nil || common.IsNil(o.OptionAssets) {
		var ret []ExchangeInformationResponseOptionAssetsInner
		return ret
	}
	return o.OptionAssets
}

// GetOptionAssetsOk returns a tuple with the OptionAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponse) GetOptionAssetsOk() ([]ExchangeInformationResponseOptionAssetsInner, bool) {
	if o == nil || common.IsNil(o.OptionAssets) {
		return nil, false
	}
	return o.OptionAssets, true
}

// HasOptionAssets returns a boolean if a field has been set.
func (o *ExchangeInformationResponse) HasOptionAssets() bool {
	if o != nil && !common.IsNil(o.OptionAssets) {
		return true
	}

	return false
}

// SetOptionAssets gets a reference to the given []ExchangeInformationResponseOptionAssetsInner and assigns it to the OptionAssets field.
func (o *ExchangeInformationResponse) SetOptionAssets(v []ExchangeInformationResponseOptionAssetsInner) {
	o.OptionAssets = v
}

// GetOptionSymbols returns the OptionSymbols field value if set, zero value otherwise.
func (o *ExchangeInformationResponse) GetOptionSymbols() []ExchangeInformationResponseOptionSymbolsInner {
	if o == nil || common.IsNil(o.OptionSymbols) {
		var ret []ExchangeInformationResponseOptionSymbolsInner
		return ret
	}
	return o.OptionSymbols
}

// GetOptionSymbolsOk returns a tuple with the OptionSymbols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponse) GetOptionSymbolsOk() ([]ExchangeInformationResponseOptionSymbolsInner, bool) {
	if o == nil || common.IsNil(o.OptionSymbols) {
		return nil, false
	}
	return o.OptionSymbols, true
}

// HasOptionSymbols returns a boolean if a field has been set.
func (o *ExchangeInformationResponse) HasOptionSymbols() bool {
	if o != nil && !common.IsNil(o.OptionSymbols) {
		return true
	}

	return false
}

// SetOptionSymbols gets a reference to the given []ExchangeInformationResponseOptionSymbolsInner and assigns it to the OptionSymbols field.
func (o *ExchangeInformationResponse) SetOptionSymbols(v []ExchangeInformationResponseOptionSymbolsInner) {
	o.OptionSymbols = v
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

func (o ExchangeInformationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeInformationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	if !common.IsNil(o.ServerTime) {
		toSerialize["serverTime"] = o.ServerTime
	}
	if !common.IsNil(o.OptionContracts) {
		toSerialize["optionContracts"] = o.OptionContracts
	}
	if !common.IsNil(o.OptionAssets) {
		toSerialize["optionAssets"] = o.OptionAssets
	}
	if !common.IsNil(o.OptionSymbols) {
		toSerialize["optionSymbols"] = o.OptionSymbols
	}
	if !common.IsNil(o.RateLimits) {
		toSerialize["rateLimits"] = o.RateLimits
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
		delete(additionalProperties, "timezone")
		delete(additionalProperties, "serverTime")
		delete(additionalProperties, "optionContracts")
		delete(additionalProperties, "optionAssets")
		delete(additionalProperties, "optionSymbols")
		delete(additionalProperties, "rateLimits")
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
