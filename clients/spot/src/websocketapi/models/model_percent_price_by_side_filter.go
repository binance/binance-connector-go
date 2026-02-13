/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the PercentPriceBySideFilter type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PercentPriceBySideFilter{}

// PercentPriceBySideFilter struct for PercentPriceBySideFilter
type PercentPriceBySideFilter struct {
	FilterType           *string `json:"filterType,omitempty"`
	MultiplierExponent   *int32  `json:"multiplierExponent,omitempty"`
	BidMultiplierUp      *string `json:"bidMultiplierUp,omitempty"`
	BidMultiplierDown    *string `json:"bidMultiplierDown,omitempty"`
	AskMultiplierUp      *string `json:"askMultiplierUp,omitempty"`
	AskMultiplierDown    *string `json:"askMultiplierDown,omitempty"`
	AvgPriceMins         *int32  `json:"avgPriceMins,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PercentPriceBySideFilter PercentPriceBySideFilter

// NewPercentPriceBySideFilter instantiates a new PercentPriceBySideFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPercentPriceBySideFilter() *PercentPriceBySideFilter {
	this := PercentPriceBySideFilter{}
	return &this
}

// NewPercentPriceBySideFilterWithDefaults instantiates a new PercentPriceBySideFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPercentPriceBySideFilterWithDefaults() *PercentPriceBySideFilter {
	this := PercentPriceBySideFilter{}
	return &this
}

// GetFilterType returns the FilterType field value if set, zero value otherwise.
func (o *PercentPriceBySideFilter) GetFilterType() string {
	if o == nil || common.IsNil(o.FilterType) {
		var ret string
		return ret
	}
	return *o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PercentPriceBySideFilter) GetFilterTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.FilterType) {
		return nil, false
	}
	return o.FilterType, true
}

// HasFilterType returns a boolean if a field has been set.
func (o *PercentPriceBySideFilter) HasFilterType() bool {
	if o != nil && !common.IsNil(o.FilterType) {
		return true
	}

	return false
}

// SetFilterType gets a reference to the given string and assigns it to the FilterType field.
func (o *PercentPriceBySideFilter) SetFilterType(v string) {
	o.FilterType = &v
}

// GetMultiplierExponent returns the MultiplierExponent field value if set, zero value otherwise.
func (o *PercentPriceBySideFilter) GetMultiplierExponent() int32 {
	if o == nil || common.IsNil(o.MultiplierExponent) {
		var ret int32
		return ret
	}
	return *o.MultiplierExponent
}

// GetMultiplierExponentOk returns a tuple with the MultiplierExponent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PercentPriceBySideFilter) GetMultiplierExponentOk() (*int32, bool) {
	if o == nil || common.IsNil(o.MultiplierExponent) {
		return nil, false
	}
	return o.MultiplierExponent, true
}

// HasMultiplierExponent returns a boolean if a field has been set.
func (o *PercentPriceBySideFilter) HasMultiplierExponent() bool {
	if o != nil && !common.IsNil(o.MultiplierExponent) {
		return true
	}

	return false
}

// SetMultiplierExponent gets a reference to the given int32 and assigns it to the MultiplierExponent field.
func (o *PercentPriceBySideFilter) SetMultiplierExponent(v int32) {
	o.MultiplierExponent = &v
}

// GetBidMultiplierUp returns the BidMultiplierUp field value if set, zero value otherwise.
func (o *PercentPriceBySideFilter) GetBidMultiplierUp() string {
	if o == nil || common.IsNil(o.BidMultiplierUp) {
		var ret string
		return ret
	}
	return *o.BidMultiplierUp
}

// GetBidMultiplierUpOk returns a tuple with the BidMultiplierUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PercentPriceBySideFilter) GetBidMultiplierUpOk() (*string, bool) {
	if o == nil || common.IsNil(o.BidMultiplierUp) {
		return nil, false
	}
	return o.BidMultiplierUp, true
}

// HasBidMultiplierUp returns a boolean if a field has been set.
func (o *PercentPriceBySideFilter) HasBidMultiplierUp() bool {
	if o != nil && !common.IsNil(o.BidMultiplierUp) {
		return true
	}

	return false
}

// SetBidMultiplierUp gets a reference to the given string and assigns it to the BidMultiplierUp field.
func (o *PercentPriceBySideFilter) SetBidMultiplierUp(v string) {
	o.BidMultiplierUp = &v
}

// GetBidMultiplierDown returns the BidMultiplierDown field value if set, zero value otherwise.
func (o *PercentPriceBySideFilter) GetBidMultiplierDown() string {
	if o == nil || common.IsNil(o.BidMultiplierDown) {
		var ret string
		return ret
	}
	return *o.BidMultiplierDown
}

// GetBidMultiplierDownOk returns a tuple with the BidMultiplierDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PercentPriceBySideFilter) GetBidMultiplierDownOk() (*string, bool) {
	if o == nil || common.IsNil(o.BidMultiplierDown) {
		return nil, false
	}
	return o.BidMultiplierDown, true
}

// HasBidMultiplierDown returns a boolean if a field has been set.
func (o *PercentPriceBySideFilter) HasBidMultiplierDown() bool {
	if o != nil && !common.IsNil(o.BidMultiplierDown) {
		return true
	}

	return false
}

// SetBidMultiplierDown gets a reference to the given string and assigns it to the BidMultiplierDown field.
func (o *PercentPriceBySideFilter) SetBidMultiplierDown(v string) {
	o.BidMultiplierDown = &v
}

// GetAskMultiplierUp returns the AskMultiplierUp field value if set, zero value otherwise.
func (o *PercentPriceBySideFilter) GetAskMultiplierUp() string {
	if o == nil || common.IsNil(o.AskMultiplierUp) {
		var ret string
		return ret
	}
	return *o.AskMultiplierUp
}

// GetAskMultiplierUpOk returns a tuple with the AskMultiplierUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PercentPriceBySideFilter) GetAskMultiplierUpOk() (*string, bool) {
	if o == nil || common.IsNil(o.AskMultiplierUp) {
		return nil, false
	}
	return o.AskMultiplierUp, true
}

// HasAskMultiplierUp returns a boolean if a field has been set.
func (o *PercentPriceBySideFilter) HasAskMultiplierUp() bool {
	if o != nil && !common.IsNil(o.AskMultiplierUp) {
		return true
	}

	return false
}

// SetAskMultiplierUp gets a reference to the given string and assigns it to the AskMultiplierUp field.
func (o *PercentPriceBySideFilter) SetAskMultiplierUp(v string) {
	o.AskMultiplierUp = &v
}

// GetAskMultiplierDown returns the AskMultiplierDown field value if set, zero value otherwise.
func (o *PercentPriceBySideFilter) GetAskMultiplierDown() string {
	if o == nil || common.IsNil(o.AskMultiplierDown) {
		var ret string
		return ret
	}
	return *o.AskMultiplierDown
}

// GetAskMultiplierDownOk returns a tuple with the AskMultiplierDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PercentPriceBySideFilter) GetAskMultiplierDownOk() (*string, bool) {
	if o == nil || common.IsNil(o.AskMultiplierDown) {
		return nil, false
	}
	return o.AskMultiplierDown, true
}

// HasAskMultiplierDown returns a boolean if a field has been set.
func (o *PercentPriceBySideFilter) HasAskMultiplierDown() bool {
	if o != nil && !common.IsNil(o.AskMultiplierDown) {
		return true
	}

	return false
}

// SetAskMultiplierDown gets a reference to the given string and assigns it to the AskMultiplierDown field.
func (o *PercentPriceBySideFilter) SetAskMultiplierDown(v string) {
	o.AskMultiplierDown = &v
}

// GetAvgPriceMins returns the AvgPriceMins field value if set, zero value otherwise.
func (o *PercentPriceBySideFilter) GetAvgPriceMins() int32 {
	if o == nil || common.IsNil(o.AvgPriceMins) {
		var ret int32
		return ret
	}
	return *o.AvgPriceMins
}

// GetAvgPriceMinsOk returns a tuple with the AvgPriceMins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PercentPriceBySideFilter) GetAvgPriceMinsOk() (*int32, bool) {
	if o == nil || common.IsNil(o.AvgPriceMins) {
		return nil, false
	}
	return o.AvgPriceMins, true
}

// HasAvgPriceMins returns a boolean if a field has been set.
func (o *PercentPriceBySideFilter) HasAvgPriceMins() bool {
	if o != nil && !common.IsNil(o.AvgPriceMins) {
		return true
	}

	return false
}

// SetAvgPriceMins gets a reference to the given int32 and assigns it to the AvgPriceMins field.
func (o *PercentPriceBySideFilter) SetAvgPriceMins(v int32) {
	o.AvgPriceMins = &v
}

func (o PercentPriceBySideFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PercentPriceBySideFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.FilterType) {
		toSerialize["filterType"] = o.FilterType
	}
	if !common.IsNil(o.MultiplierExponent) {
		toSerialize["multiplierExponent"] = o.MultiplierExponent
	}
	if !common.IsNil(o.BidMultiplierUp) {
		toSerialize["bidMultiplierUp"] = o.BidMultiplierUp
	}
	if !common.IsNil(o.BidMultiplierDown) {
		toSerialize["bidMultiplierDown"] = o.BidMultiplierDown
	}
	if !common.IsNil(o.AskMultiplierUp) {
		toSerialize["askMultiplierUp"] = o.AskMultiplierUp
	}
	if !common.IsNil(o.AskMultiplierDown) {
		toSerialize["askMultiplierDown"] = o.AskMultiplierDown
	}
	if !common.IsNil(o.AvgPriceMins) {
		toSerialize["avgPriceMins"] = o.AvgPriceMins
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PercentPriceBySideFilter) UnmarshalJSON(data []byte) (err error) {
	varPercentPriceBySideFilter := _PercentPriceBySideFilter{}

	err = json.Unmarshal(data, &varPercentPriceBySideFilter)

	if err != nil {
		return err
	}

	*o = PercentPriceBySideFilter(varPercentPriceBySideFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filterType")
		delete(additionalProperties, "multiplierExponent")
		delete(additionalProperties, "bidMultiplierUp")
		delete(additionalProperties, "bidMultiplierDown")
		delete(additionalProperties, "askMultiplierUp")
		delete(additionalProperties, "askMultiplierDown")
		delete(additionalProperties, "avgPriceMins")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePercentPriceBySideFilter struct {
	value *PercentPriceBySideFilter
	isSet bool
}

func (v NullablePercentPriceBySideFilter) Get() *PercentPriceBySideFilter {
	return v.value
}

func (v *NullablePercentPriceBySideFilter) Set(val *PercentPriceBySideFilter) {
	v.value = val
	v.isSet = true
}

func (v NullablePercentPriceBySideFilter) IsSet() bool {
	return v.isSet
}

func (v *NullablePercentPriceBySideFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePercentPriceBySideFilter(val *PercentPriceBySideFilter) *NullablePercentPriceBySideFilter {
	return &NullablePercentPriceBySideFilter{value: val, isSet: true}
}

func (v NullablePercentPriceBySideFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePercentPriceBySideFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
