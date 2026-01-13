/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the MyFiltersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MyFiltersResponse{}

// MyFiltersResponse struct for MyFiltersResponse
type MyFiltersResponse struct {
	ExchangeFilters      []ExchangeFilters `json:"exchangeFilters,omitempty"`
	SymbolFilters        []SymbolFilters   `json:"symbolFilters,omitempty"`
	AssetFilters         []AssetFilters    `json:"assetFilters,omitempty"`
	RateLimits           []RateLimits      `json:"rateLimits,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MyFiltersResponse MyFiltersResponse

// NewMyFiltersResponse instantiates a new MyFiltersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMyFiltersResponse() *MyFiltersResponse {
	this := MyFiltersResponse{}
	return &this
}

// NewMyFiltersResponseWithDefaults instantiates a new MyFiltersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMyFiltersResponseWithDefaults() *MyFiltersResponse {
	this := MyFiltersResponse{}
	return &this
}

// GetExchangeFilters returns the ExchangeFilters field value if set, zero value otherwise.
func (o *MyFiltersResponse) GetExchangeFilters() []ExchangeFilters {
	if o == nil || common.IsNil(o.ExchangeFilters) {
		var ret []ExchangeFilters
		return ret
	}
	return o.ExchangeFilters
}

// GetExchangeFiltersOk returns a tuple with the ExchangeFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyFiltersResponse) GetExchangeFiltersOk() ([]ExchangeFilters, bool) {
	if o == nil || common.IsNil(o.ExchangeFilters) {
		return nil, false
	}
	return o.ExchangeFilters, true
}

// HasExchangeFilters returns a boolean if a field has been set.
func (o *MyFiltersResponse) HasExchangeFilters() bool {
	if o != nil && !common.IsNil(o.ExchangeFilters) {
		return true
	}

	return false
}

// SetExchangeFilters gets a reference to the given []ExchangeFilters and assigns it to the ExchangeFilters field.
func (o *MyFiltersResponse) SetExchangeFilters(v []ExchangeFilters) {
	o.ExchangeFilters = v
}

// GetSymbolFilters returns the SymbolFilters field value if set, zero value otherwise.
func (o *MyFiltersResponse) GetSymbolFilters() []SymbolFilters {
	if o == nil || common.IsNil(o.SymbolFilters) {
		var ret []SymbolFilters
		return ret
	}
	return o.SymbolFilters
}

// GetSymbolFiltersOk returns a tuple with the SymbolFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyFiltersResponse) GetSymbolFiltersOk() ([]SymbolFilters, bool) {
	if o == nil || common.IsNil(o.SymbolFilters) {
		return nil, false
	}
	return o.SymbolFilters, true
}

// HasSymbolFilters returns a boolean if a field has been set.
func (o *MyFiltersResponse) HasSymbolFilters() bool {
	if o != nil && !common.IsNil(o.SymbolFilters) {
		return true
	}

	return false
}

// SetSymbolFilters gets a reference to the given []SymbolFilters and assigns it to the SymbolFilters field.
func (o *MyFiltersResponse) SetSymbolFilters(v []SymbolFilters) {
	o.SymbolFilters = v
}

// GetAssetFilters returns the AssetFilters field value if set, zero value otherwise.
func (o *MyFiltersResponse) GetAssetFilters() []AssetFilters {
	if o == nil || common.IsNil(o.AssetFilters) {
		var ret []AssetFilters
		return ret
	}
	return o.AssetFilters
}

// GetAssetFiltersOk returns a tuple with the AssetFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyFiltersResponse) GetAssetFiltersOk() ([]AssetFilters, bool) {
	if o == nil || common.IsNil(o.AssetFilters) {
		return nil, false
	}
	return o.AssetFilters, true
}

// HasAssetFilters returns a boolean if a field has been set.
func (o *MyFiltersResponse) HasAssetFilters() bool {
	if o != nil && !common.IsNil(o.AssetFilters) {
		return true
	}

	return false
}

// SetAssetFilters gets a reference to the given []AssetFilters and assigns it to the AssetFilters field.
func (o *MyFiltersResponse) SetAssetFilters(v []AssetFilters) {
	o.AssetFilters = v
}

// GetRateLimits returns the RateLimits field value if set, zero value otherwise.
func (o *MyFiltersResponse) GetRateLimits() []RateLimits {
	if o == nil || common.IsNil(o.RateLimits) {
		var ret []RateLimits
		return ret
	}
	return o.RateLimits
}

// GetRateLimitsOk returns a tuple with the RateLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyFiltersResponse) GetRateLimitsOk() ([]RateLimits, bool) {
	if o == nil || common.IsNil(o.RateLimits) {
		return nil, false
	}
	return o.RateLimits, true
}

// HasRateLimits returns a boolean if a field has been set.
func (o *MyFiltersResponse) HasRateLimits() bool {
	if o != nil && !common.IsNil(o.RateLimits) {
		return true
	}

	return false
}

// SetRateLimits gets a reference to the given []RateLimits and assigns it to the RateLimits field.
func (o *MyFiltersResponse) SetRateLimits(v []RateLimits) {
	o.RateLimits = v
}

func (o MyFiltersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MyFiltersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ExchangeFilters) {
		toSerialize["exchangeFilters"] = o.ExchangeFilters
	}
	if !common.IsNil(o.SymbolFilters) {
		toSerialize["symbolFilters"] = o.SymbolFilters
	}
	if !common.IsNil(o.AssetFilters) {
		toSerialize["assetFilters"] = o.AssetFilters
	}
	if !common.IsNil(o.RateLimits) {
		toSerialize["rateLimits"] = o.RateLimits
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MyFiltersResponse) UnmarshalJSON(data []byte) (err error) {
	varMyFiltersResponse := _MyFiltersResponse{}

	err = json.Unmarshal(data, &varMyFiltersResponse)

	if err != nil {
		return err
	}

	*o = MyFiltersResponse(varMyFiltersResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "exchangeFilters")
		delete(additionalProperties, "symbolFilters")
		delete(additionalProperties, "assetFilters")
		delete(additionalProperties, "rateLimits")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMyFiltersResponse struct {
	value *MyFiltersResponse
	isSet bool
}

func (v NullableMyFiltersResponse) Get() *MyFiltersResponse {
	return v.value
}

func (v *NullableMyFiltersResponse) Set(val *MyFiltersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMyFiltersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMyFiltersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMyFiltersResponse(val *MyFiltersResponse) *NullableMyFiltersResponse {
	return &NullableMyFiltersResponse{value: val, isSet: true}
}

func (v NullableMyFiltersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMyFiltersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
