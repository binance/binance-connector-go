/*
Binance Alpha REST API

OpenAPI Specification for the Binance Alpha REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetExchangeInfoResponseDataSymbolsInnerFiltersInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetExchangeInfoResponseDataSymbolsInnerFiltersInner{}

// GetExchangeInfoResponseDataSymbolsInnerFiltersInner struct for GetExchangeInfoResponseDataSymbolsInnerFiltersInner
type GetExchangeInfoResponseDataSymbolsInnerFiltersInner struct {
	FilterType           *string `json:"filterType,omitempty"`
	MinPrice             *string `json:"minPrice,omitempty"`
	MaxPrice             *string `json:"maxPrice,omitempty"`
	TickSize             *string `json:"tickSize,omitempty"`
	StepSize             *string `json:"stepSize,omitempty"`
	MaxQty               *string `json:"maxQty,omitempty"`
	MinQty               *string `json:"minQty,omitempty"`
	Limit                *int64  `json:"limit,omitempty"`
	MinNotional          *string `json:"minNotional,omitempty"`
	MaxNotional          *string `json:"maxNotional,omitempty"`
	MultiplierDown       *string `json:"multiplierDown,omitempty"`
	MultiplierUp         *string `json:"multiplierUp,omitempty"`
	BidMultiplierUp      *string `json:"bidMultiplierUp,omitempty"`
	AskMultiplierUp      *string `json:"askMultiplierUp,omitempty"`
	BidMultiplierDown    *string `json:"bidMultiplierDown,omitempty"`
	AskMultiplierDown    *string `json:"askMultiplierDown,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetExchangeInfoResponseDataSymbolsInnerFiltersInner GetExchangeInfoResponseDataSymbolsInnerFiltersInner

// NewGetExchangeInfoResponseDataSymbolsInnerFiltersInner instantiates a new GetExchangeInfoResponseDataSymbolsInnerFiltersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetExchangeInfoResponseDataSymbolsInnerFiltersInner() *GetExchangeInfoResponseDataSymbolsInnerFiltersInner {
	this := GetExchangeInfoResponseDataSymbolsInnerFiltersInner{}
	return &this
}

// NewGetExchangeInfoResponseDataSymbolsInnerFiltersInnerWithDefaults instantiates a new GetExchangeInfoResponseDataSymbolsInnerFiltersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetExchangeInfoResponseDataSymbolsInnerFiltersInnerWithDefaults() *GetExchangeInfoResponseDataSymbolsInnerFiltersInner {
	this := GetExchangeInfoResponseDataSymbolsInnerFiltersInner{}
	return &this
}

// GetFilterType returns the FilterType field value if set, zero value otherwise.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) GetFilterType() string {
	if o == nil || common.IsNil(o.FilterType) {
		var ret string
		return ret
	}
	return *o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) GetFilterTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.FilterType) {
		return nil, false
	}
	return o.FilterType, true
}

// HasFilterType returns a boolean if a field has been set.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) HasFilterType() bool {
	if o != nil && !common.IsNil(o.FilterType) {
		return true
	}

	return false
}

// SetFilterType gets a reference to the given string and assigns it to the FilterType field.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) SetFilterType(v string) {
	o.FilterType = &v
}

// GetMinPrice returns the MinPrice field value if set, zero value otherwise.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) GetMinPrice() string {
	if o == nil || common.IsNil(o.MinPrice) {
		var ret string
		return ret
	}
	return *o.MinPrice
}

// GetMinPriceOk returns a tuple with the MinPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) GetMinPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.MinPrice) {
		return nil, false
	}
	return o.MinPrice, true
}

// HasMinPrice returns a boolean if a field has been set.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) HasMinPrice() bool {
	if o != nil && !common.IsNil(o.MinPrice) {
		return true
	}

	return false
}

// SetMinPrice gets a reference to the given string and assigns it to the MinPrice field.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) SetMinPrice(v string) {
	o.MinPrice = &v
}

// GetMaxPrice returns the MaxPrice field value if set, zero value otherwise.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) GetMaxPrice() string {
	if o == nil || common.IsNil(o.MaxPrice) {
		var ret string
		return ret
	}
	return *o.MaxPrice
}

// GetMaxPriceOk returns a tuple with the MaxPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) GetMaxPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxPrice) {
		return nil, false
	}
	return o.MaxPrice, true
}

// HasMaxPrice returns a boolean if a field has been set.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) HasMaxPrice() bool {
	if o != nil && !common.IsNil(o.MaxPrice) {
		return true
	}

	return false
}

// SetMaxPrice gets a reference to the given string and assigns it to the MaxPrice field.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) SetMaxPrice(v string) {
	o.MaxPrice = &v
}

// GetTickSize returns the TickSize field value if set, zero value otherwise.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) GetTickSize() string {
	if o == nil || common.IsNil(o.TickSize) {
		var ret string
		return ret
	}
	return *o.TickSize
}

// GetTickSizeOk returns a tuple with the TickSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) GetTickSizeOk() (*string, bool) {
	if o == nil || common.IsNil(o.TickSize) {
		return nil, false
	}
	return o.TickSize, true
}

// HasTickSize returns a boolean if a field has been set.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) HasTickSize() bool {
	if o != nil && !common.IsNil(o.TickSize) {
		return true
	}

	return false
}

// SetTickSize gets a reference to the given string and assigns it to the TickSize field.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) SetTickSize(v string) {
	o.TickSize = &v
}

// GetStepSize returns the StepSize field value if set, zero value otherwise.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) GetStepSize() string {
	if o == nil || common.IsNil(o.StepSize) {
		var ret string
		return ret
	}
	return *o.StepSize
}

// GetStepSizeOk returns a tuple with the StepSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) GetStepSizeOk() (*string, bool) {
	if o == nil || common.IsNil(o.StepSize) {
		return nil, false
	}
	return o.StepSize, true
}

// HasStepSize returns a boolean if a field has been set.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) HasStepSize() bool {
	if o != nil && !common.IsNil(o.StepSize) {
		return true
	}

	return false
}

// SetStepSize gets a reference to the given string and assigns it to the StepSize field.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) SetStepSize(v string) {
	o.StepSize = &v
}

// GetMaxQty returns the MaxQty field value if set, zero value otherwise.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) GetMaxQty() string {
	if o == nil || common.IsNil(o.MaxQty) {
		var ret string
		return ret
	}
	return *o.MaxQty
}

// GetMaxQtyOk returns a tuple with the MaxQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) GetMaxQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxQty) {
		return nil, false
	}
	return o.MaxQty, true
}

// HasMaxQty returns a boolean if a field has been set.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) HasMaxQty() bool {
	if o != nil && !common.IsNil(o.MaxQty) {
		return true
	}

	return false
}

// SetMaxQty gets a reference to the given string and assigns it to the MaxQty field.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) SetMaxQty(v string) {
	o.MaxQty = &v
}

// GetMinQty returns the MinQty field value if set, zero value otherwise.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) GetMinQty() string {
	if o == nil || common.IsNil(o.MinQty) {
		var ret string
		return ret
	}
	return *o.MinQty
}

// GetMinQtyOk returns a tuple with the MinQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) GetMinQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.MinQty) {
		return nil, false
	}
	return o.MinQty, true
}

// HasMinQty returns a boolean if a field has been set.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) HasMinQty() bool {
	if o != nil && !common.IsNil(o.MinQty) {
		return true
	}

	return false
}

// SetMinQty gets a reference to the given string and assigns it to the MinQty field.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) SetMinQty(v string) {
	o.MinQty = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) GetLimit() int64 {
	if o == nil || common.IsNil(o.Limit) {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) GetLimitOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) HasLimit() bool {
	if o != nil && !common.IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) SetLimit(v int64) {
	o.Limit = &v
}

// GetMinNotional returns the MinNotional field value if set, zero value otherwise.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) GetMinNotional() string {
	if o == nil || common.IsNil(o.MinNotional) {
		var ret string
		return ret
	}
	return *o.MinNotional
}

// GetMinNotionalOk returns a tuple with the MinNotional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) GetMinNotionalOk() (*string, bool) {
	if o == nil || common.IsNil(o.MinNotional) {
		return nil, false
	}
	return o.MinNotional, true
}

// HasMinNotional returns a boolean if a field has been set.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) HasMinNotional() bool {
	if o != nil && !common.IsNil(o.MinNotional) {
		return true
	}

	return false
}

// SetMinNotional gets a reference to the given string and assigns it to the MinNotional field.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) SetMinNotional(v string) {
	o.MinNotional = &v
}

// GetMaxNotional returns the MaxNotional field value if set, zero value otherwise.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) GetMaxNotional() string {
	if o == nil || common.IsNil(o.MaxNotional) {
		var ret string
		return ret
	}
	return *o.MaxNotional
}

// GetMaxNotionalOk returns a tuple with the MaxNotional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) GetMaxNotionalOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxNotional) {
		return nil, false
	}
	return o.MaxNotional, true
}

// HasMaxNotional returns a boolean if a field has been set.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) HasMaxNotional() bool {
	if o != nil && !common.IsNil(o.MaxNotional) {
		return true
	}

	return false
}

// SetMaxNotional gets a reference to the given string and assigns it to the MaxNotional field.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) SetMaxNotional(v string) {
	o.MaxNotional = &v
}

// GetMultiplierDown returns the MultiplierDown field value if set, zero value otherwise.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) GetMultiplierDown() string {
	if o == nil || common.IsNil(o.MultiplierDown) {
		var ret string
		return ret
	}
	return *o.MultiplierDown
}

// GetMultiplierDownOk returns a tuple with the MultiplierDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) GetMultiplierDownOk() (*string, bool) {
	if o == nil || common.IsNil(o.MultiplierDown) {
		return nil, false
	}
	return o.MultiplierDown, true
}

// HasMultiplierDown returns a boolean if a field has been set.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) HasMultiplierDown() bool {
	if o != nil && !common.IsNil(o.MultiplierDown) {
		return true
	}

	return false
}

// SetMultiplierDown gets a reference to the given string and assigns it to the MultiplierDown field.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) SetMultiplierDown(v string) {
	o.MultiplierDown = &v
}

// GetMultiplierUp returns the MultiplierUp field value if set, zero value otherwise.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) GetMultiplierUp() string {
	if o == nil || common.IsNil(o.MultiplierUp) {
		var ret string
		return ret
	}
	return *o.MultiplierUp
}

// GetMultiplierUpOk returns a tuple with the MultiplierUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) GetMultiplierUpOk() (*string, bool) {
	if o == nil || common.IsNil(o.MultiplierUp) {
		return nil, false
	}
	return o.MultiplierUp, true
}

// HasMultiplierUp returns a boolean if a field has been set.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) HasMultiplierUp() bool {
	if o != nil && !common.IsNil(o.MultiplierUp) {
		return true
	}

	return false
}

// SetMultiplierUp gets a reference to the given string and assigns it to the MultiplierUp field.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) SetMultiplierUp(v string) {
	o.MultiplierUp = &v
}

// GetBidMultiplierUp returns the BidMultiplierUp field value if set, zero value otherwise.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) GetBidMultiplierUp() string {
	if o == nil || common.IsNil(o.BidMultiplierUp) {
		var ret string
		return ret
	}
	return *o.BidMultiplierUp
}

// GetBidMultiplierUpOk returns a tuple with the BidMultiplierUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) GetBidMultiplierUpOk() (*string, bool) {
	if o == nil || common.IsNil(o.BidMultiplierUp) {
		return nil, false
	}
	return o.BidMultiplierUp, true
}

// HasBidMultiplierUp returns a boolean if a field has been set.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) HasBidMultiplierUp() bool {
	if o != nil && !common.IsNil(o.BidMultiplierUp) {
		return true
	}

	return false
}

// SetBidMultiplierUp gets a reference to the given string and assigns it to the BidMultiplierUp field.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) SetBidMultiplierUp(v string) {
	o.BidMultiplierUp = &v
}

// GetAskMultiplierUp returns the AskMultiplierUp field value if set, zero value otherwise.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) GetAskMultiplierUp() string {
	if o == nil || common.IsNil(o.AskMultiplierUp) {
		var ret string
		return ret
	}
	return *o.AskMultiplierUp
}

// GetAskMultiplierUpOk returns a tuple with the AskMultiplierUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) GetAskMultiplierUpOk() (*string, bool) {
	if o == nil || common.IsNil(o.AskMultiplierUp) {
		return nil, false
	}
	return o.AskMultiplierUp, true
}

// HasAskMultiplierUp returns a boolean if a field has been set.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) HasAskMultiplierUp() bool {
	if o != nil && !common.IsNil(o.AskMultiplierUp) {
		return true
	}

	return false
}

// SetAskMultiplierUp gets a reference to the given string and assigns it to the AskMultiplierUp field.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) SetAskMultiplierUp(v string) {
	o.AskMultiplierUp = &v
}

// GetBidMultiplierDown returns the BidMultiplierDown field value if set, zero value otherwise.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) GetBidMultiplierDown() string {
	if o == nil || common.IsNil(o.BidMultiplierDown) {
		var ret string
		return ret
	}
	return *o.BidMultiplierDown
}

// GetBidMultiplierDownOk returns a tuple with the BidMultiplierDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) GetBidMultiplierDownOk() (*string, bool) {
	if o == nil || common.IsNil(o.BidMultiplierDown) {
		return nil, false
	}
	return o.BidMultiplierDown, true
}

// HasBidMultiplierDown returns a boolean if a field has been set.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) HasBidMultiplierDown() bool {
	if o != nil && !common.IsNil(o.BidMultiplierDown) {
		return true
	}

	return false
}

// SetBidMultiplierDown gets a reference to the given string and assigns it to the BidMultiplierDown field.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) SetBidMultiplierDown(v string) {
	o.BidMultiplierDown = &v
}

// GetAskMultiplierDown returns the AskMultiplierDown field value if set, zero value otherwise.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) GetAskMultiplierDown() string {
	if o == nil || common.IsNil(o.AskMultiplierDown) {
		var ret string
		return ret
	}
	return *o.AskMultiplierDown
}

// GetAskMultiplierDownOk returns a tuple with the AskMultiplierDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) GetAskMultiplierDownOk() (*string, bool) {
	if o == nil || common.IsNil(o.AskMultiplierDown) {
		return nil, false
	}
	return o.AskMultiplierDown, true
}

// HasAskMultiplierDown returns a boolean if a field has been set.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) HasAskMultiplierDown() bool {
	if o != nil && !common.IsNil(o.AskMultiplierDown) {
		return true
	}

	return false
}

// SetAskMultiplierDown gets a reference to the given string and assigns it to the AskMultiplierDown field.
func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) SetAskMultiplierDown(v string) {
	o.AskMultiplierDown = &v
}

func (o GetExchangeInfoResponseDataSymbolsInnerFiltersInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetExchangeInfoResponseDataSymbolsInnerFiltersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.FilterType) {
		toSerialize["filterType"] = o.FilterType
	}
	if !common.IsNil(o.MinPrice) {
		toSerialize["minPrice"] = o.MinPrice
	}
	if !common.IsNil(o.MaxPrice) {
		toSerialize["maxPrice"] = o.MaxPrice
	}
	if !common.IsNil(o.TickSize) {
		toSerialize["tickSize"] = o.TickSize
	}
	if !common.IsNil(o.StepSize) {
		toSerialize["stepSize"] = o.StepSize
	}
	if !common.IsNil(o.MaxQty) {
		toSerialize["maxQty"] = o.MaxQty
	}
	if !common.IsNil(o.MinQty) {
		toSerialize["minQty"] = o.MinQty
	}
	if !common.IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !common.IsNil(o.MinNotional) {
		toSerialize["minNotional"] = o.MinNotional
	}
	if !common.IsNil(o.MaxNotional) {
		toSerialize["maxNotional"] = o.MaxNotional
	}
	if !common.IsNil(o.MultiplierDown) {
		toSerialize["multiplierDown"] = o.MultiplierDown
	}
	if !common.IsNil(o.MultiplierUp) {
		toSerialize["multiplierUp"] = o.MultiplierUp
	}
	if !common.IsNil(o.BidMultiplierUp) {
		toSerialize["bidMultiplierUp"] = o.BidMultiplierUp
	}
	if !common.IsNil(o.AskMultiplierUp) {
		toSerialize["askMultiplierUp"] = o.AskMultiplierUp
	}
	if !common.IsNil(o.BidMultiplierDown) {
		toSerialize["bidMultiplierDown"] = o.BidMultiplierDown
	}
	if !common.IsNil(o.AskMultiplierDown) {
		toSerialize["askMultiplierDown"] = o.AskMultiplierDown
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) UnmarshalJSON(data []byte) (err error) {
	varGetExchangeInfoResponseDataSymbolsInnerFiltersInner := _GetExchangeInfoResponseDataSymbolsInnerFiltersInner{}

	err = json.Unmarshal(data, &varGetExchangeInfoResponseDataSymbolsInnerFiltersInner)

	if err != nil {
		return err
	}

	*o = GetExchangeInfoResponseDataSymbolsInnerFiltersInner(varGetExchangeInfoResponseDataSymbolsInnerFiltersInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filterType")
		delete(additionalProperties, "minPrice")
		delete(additionalProperties, "maxPrice")
		delete(additionalProperties, "tickSize")
		delete(additionalProperties, "stepSize")
		delete(additionalProperties, "maxQty")
		delete(additionalProperties, "minQty")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "minNotional")
		delete(additionalProperties, "maxNotional")
		delete(additionalProperties, "multiplierDown")
		delete(additionalProperties, "multiplierUp")
		delete(additionalProperties, "bidMultiplierUp")
		delete(additionalProperties, "askMultiplierUp")
		delete(additionalProperties, "bidMultiplierDown")
		delete(additionalProperties, "askMultiplierDown")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetExchangeInfoResponseDataSymbolsInnerFiltersInner struct {
	value *GetExchangeInfoResponseDataSymbolsInnerFiltersInner
	isSet bool
}

func (v NullableGetExchangeInfoResponseDataSymbolsInnerFiltersInner) Get() *GetExchangeInfoResponseDataSymbolsInnerFiltersInner {
	return v.value
}

func (v *NullableGetExchangeInfoResponseDataSymbolsInnerFiltersInner) Set(val *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetExchangeInfoResponseDataSymbolsInnerFiltersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetExchangeInfoResponseDataSymbolsInnerFiltersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetExchangeInfoResponseDataSymbolsInnerFiltersInner(val *GetExchangeInfoResponseDataSymbolsInnerFiltersInner) *NullableGetExchangeInfoResponseDataSymbolsInnerFiltersInner {
	return &NullableGetExchangeInfoResponseDataSymbolsInnerFiltersInner{value: val, isSet: true}
}

func (v NullableGetExchangeInfoResponseDataSymbolsInnerFiltersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetExchangeInfoResponseDataSymbolsInnerFiltersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
