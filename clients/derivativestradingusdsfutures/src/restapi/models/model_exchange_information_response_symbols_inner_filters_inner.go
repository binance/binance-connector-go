/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ExchangeInformationResponseSymbolsInnerFiltersInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ExchangeInformationResponseSymbolsInnerFiltersInner{}

// ExchangeInformationResponseSymbolsInnerFiltersInner struct for ExchangeInformationResponseSymbolsInnerFiltersInner
type ExchangeInformationResponseSymbolsInnerFiltersInner struct {
	FilterType           *string `json:"filterType,omitempty"`
	MaxPrice             *string `json:"maxPrice,omitempty"`
	MinPrice             *string `json:"minPrice,omitempty"`
	TickSize             *string `json:"tickSize,omitempty"`
	MaxQty               *string `json:"maxQty,omitempty"`
	MinQty               *string `json:"minQty,omitempty"`
	StepSize             *string `json:"stepSize,omitempty"`
	Limit                *int64  `json:"limit,omitempty"`
	Notional             *string `json:"notional,omitempty"`
	MultiplierUp         *string `json:"multiplierUp,omitempty"`
	MultiplierDown       *string `json:"multiplierDown,omitempty"`
	MultiplierDecimal    *string `json:"multiplierDecimal,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExchangeInformationResponseSymbolsInnerFiltersInner ExchangeInformationResponseSymbolsInnerFiltersInner

// NewExchangeInformationResponseSymbolsInnerFiltersInner instantiates a new ExchangeInformationResponseSymbolsInnerFiltersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeInformationResponseSymbolsInnerFiltersInner() *ExchangeInformationResponseSymbolsInnerFiltersInner {
	this := ExchangeInformationResponseSymbolsInnerFiltersInner{}
	return &this
}

// NewExchangeInformationResponseSymbolsInnerFiltersInnerWithDefaults instantiates a new ExchangeInformationResponseSymbolsInnerFiltersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeInformationResponseSymbolsInnerFiltersInnerWithDefaults() *ExchangeInformationResponseSymbolsInnerFiltersInner {
	this := ExchangeInformationResponseSymbolsInnerFiltersInner{}
	return &this
}

// GetFilterType returns the FilterType field value if set, zero value otherwise.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) GetFilterType() string {
	if o == nil || common.IsNil(o.FilterType) {
		var ret string
		return ret
	}
	return *o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) GetFilterTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.FilterType) {
		return nil, false
	}
	return o.FilterType, true
}

// HasFilterType returns a boolean if a field has been set.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) HasFilterType() bool {
	if o != nil && !common.IsNil(o.FilterType) {
		return true
	}

	return false
}

// SetFilterType gets a reference to the given string and assigns it to the FilterType field.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) SetFilterType(v string) {
	o.FilterType = &v
}

// GetMaxPrice returns the MaxPrice field value if set, zero value otherwise.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) GetMaxPrice() string {
	if o == nil || common.IsNil(o.MaxPrice) {
		var ret string
		return ret
	}
	return *o.MaxPrice
}

// GetMaxPriceOk returns a tuple with the MaxPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) GetMaxPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxPrice) {
		return nil, false
	}
	return o.MaxPrice, true
}

// HasMaxPrice returns a boolean if a field has been set.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) HasMaxPrice() bool {
	if o != nil && !common.IsNil(o.MaxPrice) {
		return true
	}

	return false
}

// SetMaxPrice gets a reference to the given string and assigns it to the MaxPrice field.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) SetMaxPrice(v string) {
	o.MaxPrice = &v
}

// GetMinPrice returns the MinPrice field value if set, zero value otherwise.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) GetMinPrice() string {
	if o == nil || common.IsNil(o.MinPrice) {
		var ret string
		return ret
	}
	return *o.MinPrice
}

// GetMinPriceOk returns a tuple with the MinPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) GetMinPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.MinPrice) {
		return nil, false
	}
	return o.MinPrice, true
}

// HasMinPrice returns a boolean if a field has been set.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) HasMinPrice() bool {
	if o != nil && !common.IsNil(o.MinPrice) {
		return true
	}

	return false
}

// SetMinPrice gets a reference to the given string and assigns it to the MinPrice field.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) SetMinPrice(v string) {
	o.MinPrice = &v
}

// GetTickSize returns the TickSize field value if set, zero value otherwise.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) GetTickSize() string {
	if o == nil || common.IsNil(o.TickSize) {
		var ret string
		return ret
	}
	return *o.TickSize
}

// GetTickSizeOk returns a tuple with the TickSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) GetTickSizeOk() (*string, bool) {
	if o == nil || common.IsNil(o.TickSize) {
		return nil, false
	}
	return o.TickSize, true
}

// HasTickSize returns a boolean if a field has been set.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) HasTickSize() bool {
	if o != nil && !common.IsNil(o.TickSize) {
		return true
	}

	return false
}

// SetTickSize gets a reference to the given string and assigns it to the TickSize field.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) SetTickSize(v string) {
	o.TickSize = &v
}

// GetMaxQty returns the MaxQty field value if set, zero value otherwise.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) GetMaxQty() string {
	if o == nil || common.IsNil(o.MaxQty) {
		var ret string
		return ret
	}
	return *o.MaxQty
}

// GetMaxQtyOk returns a tuple with the MaxQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) GetMaxQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxQty) {
		return nil, false
	}
	return o.MaxQty, true
}

// HasMaxQty returns a boolean if a field has been set.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) HasMaxQty() bool {
	if o != nil && !common.IsNil(o.MaxQty) {
		return true
	}

	return false
}

// SetMaxQty gets a reference to the given string and assigns it to the MaxQty field.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) SetMaxQty(v string) {
	o.MaxQty = &v
}

// GetMinQty returns the MinQty field value if set, zero value otherwise.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) GetMinQty() string {
	if o == nil || common.IsNil(o.MinQty) {
		var ret string
		return ret
	}
	return *o.MinQty
}

// GetMinQtyOk returns a tuple with the MinQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) GetMinQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.MinQty) {
		return nil, false
	}
	return o.MinQty, true
}

// HasMinQty returns a boolean if a field has been set.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) HasMinQty() bool {
	if o != nil && !common.IsNil(o.MinQty) {
		return true
	}

	return false
}

// SetMinQty gets a reference to the given string and assigns it to the MinQty field.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) SetMinQty(v string) {
	o.MinQty = &v
}

// GetStepSize returns the StepSize field value if set, zero value otherwise.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) GetStepSize() string {
	if o == nil || common.IsNil(o.StepSize) {
		var ret string
		return ret
	}
	return *o.StepSize
}

// GetStepSizeOk returns a tuple with the StepSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) GetStepSizeOk() (*string, bool) {
	if o == nil || common.IsNil(o.StepSize) {
		return nil, false
	}
	return o.StepSize, true
}

// HasStepSize returns a boolean if a field has been set.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) HasStepSize() bool {
	if o != nil && !common.IsNil(o.StepSize) {
		return true
	}

	return false
}

// SetStepSize gets a reference to the given string and assigns it to the StepSize field.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) SetStepSize(v string) {
	o.StepSize = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) GetLimit() int64 {
	if o == nil || common.IsNil(o.Limit) {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) GetLimitOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) HasLimit() bool {
	if o != nil && !common.IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) SetLimit(v int64) {
	o.Limit = &v
}

// GetNotional returns the Notional field value if set, zero value otherwise.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) GetNotional() string {
	if o == nil || common.IsNil(o.Notional) {
		var ret string
		return ret
	}
	return *o.Notional
}

// GetNotionalOk returns a tuple with the Notional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) GetNotionalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Notional) {
		return nil, false
	}
	return o.Notional, true
}

// HasNotional returns a boolean if a field has been set.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) HasNotional() bool {
	if o != nil && !common.IsNil(o.Notional) {
		return true
	}

	return false
}

// SetNotional gets a reference to the given string and assigns it to the Notional field.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) SetNotional(v string) {
	o.Notional = &v
}

// GetMultiplierUp returns the MultiplierUp field value if set, zero value otherwise.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) GetMultiplierUp() string {
	if o == nil || common.IsNil(o.MultiplierUp) {
		var ret string
		return ret
	}
	return *o.MultiplierUp
}

// GetMultiplierUpOk returns a tuple with the MultiplierUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) GetMultiplierUpOk() (*string, bool) {
	if o == nil || common.IsNil(o.MultiplierUp) {
		return nil, false
	}
	return o.MultiplierUp, true
}

// HasMultiplierUp returns a boolean if a field has been set.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) HasMultiplierUp() bool {
	if o != nil && !common.IsNil(o.MultiplierUp) {
		return true
	}

	return false
}

// SetMultiplierUp gets a reference to the given string and assigns it to the MultiplierUp field.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) SetMultiplierUp(v string) {
	o.MultiplierUp = &v
}

// GetMultiplierDown returns the MultiplierDown field value if set, zero value otherwise.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) GetMultiplierDown() string {
	if o == nil || common.IsNil(o.MultiplierDown) {
		var ret string
		return ret
	}
	return *o.MultiplierDown
}

// GetMultiplierDownOk returns a tuple with the MultiplierDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) GetMultiplierDownOk() (*string, bool) {
	if o == nil || common.IsNil(o.MultiplierDown) {
		return nil, false
	}
	return o.MultiplierDown, true
}

// HasMultiplierDown returns a boolean if a field has been set.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) HasMultiplierDown() bool {
	if o != nil && !common.IsNil(o.MultiplierDown) {
		return true
	}

	return false
}

// SetMultiplierDown gets a reference to the given string and assigns it to the MultiplierDown field.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) SetMultiplierDown(v string) {
	o.MultiplierDown = &v
}

// GetMultiplierDecimal returns the MultiplierDecimal field value if set, zero value otherwise.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) GetMultiplierDecimal() string {
	if o == nil || common.IsNil(o.MultiplierDecimal) {
		var ret string
		return ret
	}
	return *o.MultiplierDecimal
}

// GetMultiplierDecimalOk returns a tuple with the MultiplierDecimal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) GetMultiplierDecimalOk() (*string, bool) {
	if o == nil || common.IsNil(o.MultiplierDecimal) {
		return nil, false
	}
	return o.MultiplierDecimal, true
}

// HasMultiplierDecimal returns a boolean if a field has been set.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) HasMultiplierDecimal() bool {
	if o != nil && !common.IsNil(o.MultiplierDecimal) {
		return true
	}

	return false
}

// SetMultiplierDecimal gets a reference to the given string and assigns it to the MultiplierDecimal field.
func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) SetMultiplierDecimal(v string) {
	o.MultiplierDecimal = &v
}

func (o ExchangeInformationResponseSymbolsInnerFiltersInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeInformationResponseSymbolsInnerFiltersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.FilterType) {
		toSerialize["filterType"] = o.FilterType
	}
	if !common.IsNil(o.MaxPrice) {
		toSerialize["maxPrice"] = o.MaxPrice
	}
	if !common.IsNil(o.MinPrice) {
		toSerialize["minPrice"] = o.MinPrice
	}
	if !common.IsNil(o.TickSize) {
		toSerialize["tickSize"] = o.TickSize
	}
	if !common.IsNil(o.MaxQty) {
		toSerialize["maxQty"] = o.MaxQty
	}
	if !common.IsNil(o.MinQty) {
		toSerialize["minQty"] = o.MinQty
	}
	if !common.IsNil(o.StepSize) {
		toSerialize["stepSize"] = o.StepSize
	}
	if !common.IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !common.IsNil(o.Notional) {
		toSerialize["notional"] = o.Notional
	}
	if !common.IsNil(o.MultiplierUp) {
		toSerialize["multiplierUp"] = o.MultiplierUp
	}
	if !common.IsNil(o.MultiplierDown) {
		toSerialize["multiplierDown"] = o.MultiplierDown
	}
	if !common.IsNil(o.MultiplierDecimal) {
		toSerialize["multiplierDecimal"] = o.MultiplierDecimal
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExchangeInformationResponseSymbolsInnerFiltersInner) UnmarshalJSON(data []byte) (err error) {
	varExchangeInformationResponseSymbolsInnerFiltersInner := _ExchangeInformationResponseSymbolsInnerFiltersInner{}

	err = json.Unmarshal(data, &varExchangeInformationResponseSymbolsInnerFiltersInner)

	if err != nil {
		return err
	}

	*o = ExchangeInformationResponseSymbolsInnerFiltersInner(varExchangeInformationResponseSymbolsInnerFiltersInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filterType")
		delete(additionalProperties, "maxPrice")
		delete(additionalProperties, "minPrice")
		delete(additionalProperties, "tickSize")
		delete(additionalProperties, "maxQty")
		delete(additionalProperties, "minQty")
		delete(additionalProperties, "stepSize")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "notional")
		delete(additionalProperties, "multiplierUp")
		delete(additionalProperties, "multiplierDown")
		delete(additionalProperties, "multiplierDecimal")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExchangeInformationResponseSymbolsInnerFiltersInner struct {
	value *ExchangeInformationResponseSymbolsInnerFiltersInner
	isSet bool
}

func (v NullableExchangeInformationResponseSymbolsInnerFiltersInner) Get() *ExchangeInformationResponseSymbolsInnerFiltersInner {
	return v.value
}

func (v *NullableExchangeInformationResponseSymbolsInnerFiltersInner) Set(val *ExchangeInformationResponseSymbolsInnerFiltersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeInformationResponseSymbolsInnerFiltersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeInformationResponseSymbolsInnerFiltersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeInformationResponseSymbolsInnerFiltersInner(val *ExchangeInformationResponseSymbolsInnerFiltersInner) *NullableExchangeInformationResponseSymbolsInnerFiltersInner {
	return &NullableExchangeInformationResponseSymbolsInnerFiltersInner{value: val, isSet: true}
}

func (v NullableExchangeInformationResponseSymbolsInnerFiltersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeInformationResponseSymbolsInnerFiltersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
