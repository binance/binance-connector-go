/*
Binance Spot WebSocket Streams

OpenAPI Specifications for the Binance Spot WebSocket Streams  API documents:   - [Github web-socket-streams documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-streams.md)   - [General API information for web-socket-streams on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-streams)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the PercentPriceFilter type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PercentPriceFilter{}

// PercentPriceFilter struct for PercentPriceFilter
type PercentPriceFilter struct {
	FilterType           *string `json:"filterType,omitempty"`
	MultiplierExponent   *int32  `json:"multiplierExponent,omitempty"`
	MultiplierUp         *string `json:"multiplierUp,omitempty"`
	MultiplierDown       *string `json:"multiplierDown,omitempty"`
	AvgPriceMins         *int32  `json:"avgPriceMins,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PercentPriceFilter PercentPriceFilter

// NewPercentPriceFilter instantiates a new PercentPriceFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPercentPriceFilter() *PercentPriceFilter {
	this := PercentPriceFilter{}
	return &this
}

// NewPercentPriceFilterWithDefaults instantiates a new PercentPriceFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPercentPriceFilterWithDefaults() *PercentPriceFilter {
	this := PercentPriceFilter{}
	return &this
}

// GetFilterType returns the FilterType field value if set, zero value otherwise.
func (o *PercentPriceFilter) GetFilterType() string {
	if o == nil || common.IsNil(o.FilterType) {
		var ret string
		return ret
	}
	return *o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PercentPriceFilter) GetFilterTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.FilterType) {
		return nil, false
	}
	return o.FilterType, true
}

// HasFilterType returns a boolean if a field has been set.
func (o *PercentPriceFilter) HasFilterType() bool {
	if o != nil && !common.IsNil(o.FilterType) {
		return true
	}

	return false
}

// SetFilterType gets a reference to the given string and assigns it to the FilterType field.
func (o *PercentPriceFilter) SetFilterType(v string) {
	o.FilterType = &v
}

// GetMultiplierExponent returns the MultiplierExponent field value if set, zero value otherwise.
func (o *PercentPriceFilter) GetMultiplierExponent() int32 {
	if o == nil || common.IsNil(o.MultiplierExponent) {
		var ret int32
		return ret
	}
	return *o.MultiplierExponent
}

// GetMultiplierExponentOk returns a tuple with the MultiplierExponent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PercentPriceFilter) GetMultiplierExponentOk() (*int32, bool) {
	if o == nil || common.IsNil(o.MultiplierExponent) {
		return nil, false
	}
	return o.MultiplierExponent, true
}

// HasMultiplierExponent returns a boolean if a field has been set.
func (o *PercentPriceFilter) HasMultiplierExponent() bool {
	if o != nil && !common.IsNil(o.MultiplierExponent) {
		return true
	}

	return false
}

// SetMultiplierExponent gets a reference to the given int32 and assigns it to the MultiplierExponent field.
func (o *PercentPriceFilter) SetMultiplierExponent(v int32) {
	o.MultiplierExponent = &v
}

// GetMultiplierUp returns the MultiplierUp field value if set, zero value otherwise.
func (o *PercentPriceFilter) GetMultiplierUp() string {
	if o == nil || common.IsNil(o.MultiplierUp) {
		var ret string
		return ret
	}
	return *o.MultiplierUp
}

// GetMultiplierUpOk returns a tuple with the MultiplierUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PercentPriceFilter) GetMultiplierUpOk() (*string, bool) {
	if o == nil || common.IsNil(o.MultiplierUp) {
		return nil, false
	}
	return o.MultiplierUp, true
}

// HasMultiplierUp returns a boolean if a field has been set.
func (o *PercentPriceFilter) HasMultiplierUp() bool {
	if o != nil && !common.IsNil(o.MultiplierUp) {
		return true
	}

	return false
}

// SetMultiplierUp gets a reference to the given string and assigns it to the MultiplierUp field.
func (o *PercentPriceFilter) SetMultiplierUp(v string) {
	o.MultiplierUp = &v
}

// GetMultiplierDown returns the MultiplierDown field value if set, zero value otherwise.
func (o *PercentPriceFilter) GetMultiplierDown() string {
	if o == nil || common.IsNil(o.MultiplierDown) {
		var ret string
		return ret
	}
	return *o.MultiplierDown
}

// GetMultiplierDownOk returns a tuple with the MultiplierDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PercentPriceFilter) GetMultiplierDownOk() (*string, bool) {
	if o == nil || common.IsNil(o.MultiplierDown) {
		return nil, false
	}
	return o.MultiplierDown, true
}

// HasMultiplierDown returns a boolean if a field has been set.
func (o *PercentPriceFilter) HasMultiplierDown() bool {
	if o != nil && !common.IsNil(o.MultiplierDown) {
		return true
	}

	return false
}

// SetMultiplierDown gets a reference to the given string and assigns it to the MultiplierDown field.
func (o *PercentPriceFilter) SetMultiplierDown(v string) {
	o.MultiplierDown = &v
}

// GetAvgPriceMins returns the AvgPriceMins field value if set, zero value otherwise.
func (o *PercentPriceFilter) GetAvgPriceMins() int32 {
	if o == nil || common.IsNil(o.AvgPriceMins) {
		var ret int32
		return ret
	}
	return *o.AvgPriceMins
}

// GetAvgPriceMinsOk returns a tuple with the AvgPriceMins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PercentPriceFilter) GetAvgPriceMinsOk() (*int32, bool) {
	if o == nil || common.IsNil(o.AvgPriceMins) {
		return nil, false
	}
	return o.AvgPriceMins, true
}

// HasAvgPriceMins returns a boolean if a field has been set.
func (o *PercentPriceFilter) HasAvgPriceMins() bool {
	if o != nil && !common.IsNil(o.AvgPriceMins) {
		return true
	}

	return false
}

// SetAvgPriceMins gets a reference to the given int32 and assigns it to the AvgPriceMins field.
func (o *PercentPriceFilter) SetAvgPriceMins(v int32) {
	o.AvgPriceMins = &v
}

func (o PercentPriceFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PercentPriceFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.FilterType) {
		toSerialize["filterType"] = o.FilterType
	}
	if !common.IsNil(o.MultiplierExponent) {
		toSerialize["multiplierExponent"] = o.MultiplierExponent
	}
	if !common.IsNil(o.MultiplierUp) {
		toSerialize["multiplierUp"] = o.MultiplierUp
	}
	if !common.IsNil(o.MultiplierDown) {
		toSerialize["multiplierDown"] = o.MultiplierDown
	}
	if !common.IsNil(o.AvgPriceMins) {
		toSerialize["avgPriceMins"] = o.AvgPriceMins
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PercentPriceFilter) UnmarshalJSON(data []byte) (err error) {
	varPercentPriceFilter := _PercentPriceFilter{}

	err = json.Unmarshal(data, &varPercentPriceFilter)

	if err != nil {
		return err
	}

	*o = PercentPriceFilter(varPercentPriceFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filterType")
		delete(additionalProperties, "multiplierExponent")
		delete(additionalProperties, "multiplierUp")
		delete(additionalProperties, "multiplierDown")
		delete(additionalProperties, "avgPriceMins")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePercentPriceFilter struct {
	value *PercentPriceFilter
	isSet bool
}

func (v NullablePercentPriceFilter) Get() *PercentPriceFilter {
	return v.value
}

func (v *NullablePercentPriceFilter) Set(val *PercentPriceFilter) {
	v.value = val
	v.isSet = true
}

func (v NullablePercentPriceFilter) IsSet() bool {
	return v.isSet
}

func (v *NullablePercentPriceFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePercentPriceFilter(val *PercentPriceFilter) *NullablePercentPriceFilter {
	return &NullablePercentPriceFilter{value: val, isSet: true}
}

func (v NullablePercentPriceFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePercentPriceFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
