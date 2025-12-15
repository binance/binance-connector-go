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

// checks if the OptionMarkPriceResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OptionMarkPriceResponseInner{}

// OptionMarkPriceResponseInner struct for OptionMarkPriceResponseInner
type OptionMarkPriceResponseInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	MarkPrice            *string `json:"markPrice,omitempty"`
	BidIV                *string `json:"bidIV,omitempty"`
	AskIV                *string `json:"askIV,omitempty"`
	MarkIV               *string `json:"markIV,omitempty"`
	Delta                *string `json:"delta,omitempty"`
	Theta                *string `json:"theta,omitempty"`
	Gamma                *string `json:"gamma,omitempty"`
	Vega                 *string `json:"vega,omitempty"`
	HighPriceLimit       *string `json:"highPriceLimit,omitempty"`
	LowPriceLimit        *string `json:"lowPriceLimit,omitempty"`
	RiskFreeInterest     *string `json:"riskFreeInterest,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OptionMarkPriceResponseInner OptionMarkPriceResponseInner

// NewOptionMarkPriceResponseInner instantiates a new OptionMarkPriceResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionMarkPriceResponseInner() *OptionMarkPriceResponseInner {
	this := OptionMarkPriceResponseInner{}
	return &this
}

// NewOptionMarkPriceResponseInnerWithDefaults instantiates a new OptionMarkPriceResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionMarkPriceResponseInnerWithDefaults() *OptionMarkPriceResponseInner {
	this := OptionMarkPriceResponseInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *OptionMarkPriceResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarkPriceResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *OptionMarkPriceResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *OptionMarkPriceResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetMarkPrice returns the MarkPrice field value if set, zero value otherwise.
func (o *OptionMarkPriceResponseInner) GetMarkPrice() string {
	if o == nil || common.IsNil(o.MarkPrice) {
		var ret string
		return ret
	}
	return *o.MarkPrice
}

// GetMarkPriceOk returns a tuple with the MarkPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarkPriceResponseInner) GetMarkPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarkPrice) {
		return nil, false
	}
	return o.MarkPrice, true
}

// HasMarkPrice returns a boolean if a field has been set.
func (o *OptionMarkPriceResponseInner) HasMarkPrice() bool {
	if o != nil && !common.IsNil(o.MarkPrice) {
		return true
	}

	return false
}

// SetMarkPrice gets a reference to the given string and assigns it to the MarkPrice field.
func (o *OptionMarkPriceResponseInner) SetMarkPrice(v string) {
	o.MarkPrice = &v
}

// GetBidIV returns the BidIV field value if set, zero value otherwise.
func (o *OptionMarkPriceResponseInner) GetBidIV() string {
	if o == nil || common.IsNil(o.BidIV) {
		var ret string
		return ret
	}
	return *o.BidIV
}

// GetBidIVOk returns a tuple with the BidIV field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarkPriceResponseInner) GetBidIVOk() (*string, bool) {
	if o == nil || common.IsNil(o.BidIV) {
		return nil, false
	}
	return o.BidIV, true
}

// HasBidIV returns a boolean if a field has been set.
func (o *OptionMarkPriceResponseInner) HasBidIV() bool {
	if o != nil && !common.IsNil(o.BidIV) {
		return true
	}

	return false
}

// SetBidIV gets a reference to the given string and assigns it to the BidIV field.
func (o *OptionMarkPriceResponseInner) SetBidIV(v string) {
	o.BidIV = &v
}

// GetAskIV returns the AskIV field value if set, zero value otherwise.
func (o *OptionMarkPriceResponseInner) GetAskIV() string {
	if o == nil || common.IsNil(o.AskIV) {
		var ret string
		return ret
	}
	return *o.AskIV
}

// GetAskIVOk returns a tuple with the AskIV field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarkPriceResponseInner) GetAskIVOk() (*string, bool) {
	if o == nil || common.IsNil(o.AskIV) {
		return nil, false
	}
	return o.AskIV, true
}

// HasAskIV returns a boolean if a field has been set.
func (o *OptionMarkPriceResponseInner) HasAskIV() bool {
	if o != nil && !common.IsNil(o.AskIV) {
		return true
	}

	return false
}

// SetAskIV gets a reference to the given string and assigns it to the AskIV field.
func (o *OptionMarkPriceResponseInner) SetAskIV(v string) {
	o.AskIV = &v
}

// GetMarkIV returns the MarkIV field value if set, zero value otherwise.
func (o *OptionMarkPriceResponseInner) GetMarkIV() string {
	if o == nil || common.IsNil(o.MarkIV) {
		var ret string
		return ret
	}
	return *o.MarkIV
}

// GetMarkIVOk returns a tuple with the MarkIV field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarkPriceResponseInner) GetMarkIVOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarkIV) {
		return nil, false
	}
	return o.MarkIV, true
}

// HasMarkIV returns a boolean if a field has been set.
func (o *OptionMarkPriceResponseInner) HasMarkIV() bool {
	if o != nil && !common.IsNil(o.MarkIV) {
		return true
	}

	return false
}

// SetMarkIV gets a reference to the given string and assigns it to the MarkIV field.
func (o *OptionMarkPriceResponseInner) SetMarkIV(v string) {
	o.MarkIV = &v
}

// GetDelta returns the Delta field value if set, zero value otherwise.
func (o *OptionMarkPriceResponseInner) GetDelta() string {
	if o == nil || common.IsNil(o.Delta) {
		var ret string
		return ret
	}
	return *o.Delta
}

// GetDeltaOk returns a tuple with the Delta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarkPriceResponseInner) GetDeltaOk() (*string, bool) {
	if o == nil || common.IsNil(o.Delta) {
		return nil, false
	}
	return o.Delta, true
}

// HasDelta returns a boolean if a field has been set.
func (o *OptionMarkPriceResponseInner) HasDelta() bool {
	if o != nil && !common.IsNil(o.Delta) {
		return true
	}

	return false
}

// SetDelta gets a reference to the given string and assigns it to the Delta field.
func (o *OptionMarkPriceResponseInner) SetDelta(v string) {
	o.Delta = &v
}

// GetTheta returns the Theta field value if set, zero value otherwise.
func (o *OptionMarkPriceResponseInner) GetTheta() string {
	if o == nil || common.IsNil(o.Theta) {
		var ret string
		return ret
	}
	return *o.Theta
}

// GetThetaOk returns a tuple with the Theta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarkPriceResponseInner) GetThetaOk() (*string, bool) {
	if o == nil || common.IsNil(o.Theta) {
		return nil, false
	}
	return o.Theta, true
}

// HasTheta returns a boolean if a field has been set.
func (o *OptionMarkPriceResponseInner) HasTheta() bool {
	if o != nil && !common.IsNil(o.Theta) {
		return true
	}

	return false
}

// SetTheta gets a reference to the given string and assigns it to the Theta field.
func (o *OptionMarkPriceResponseInner) SetTheta(v string) {
	o.Theta = &v
}

// GetGamma returns the Gamma field value if set, zero value otherwise.
func (o *OptionMarkPriceResponseInner) GetGamma() string {
	if o == nil || common.IsNil(o.Gamma) {
		var ret string
		return ret
	}
	return *o.Gamma
}

// GetGammaOk returns a tuple with the Gamma field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarkPriceResponseInner) GetGammaOk() (*string, bool) {
	if o == nil || common.IsNil(o.Gamma) {
		return nil, false
	}
	return o.Gamma, true
}

// HasGamma returns a boolean if a field has been set.
func (o *OptionMarkPriceResponseInner) HasGamma() bool {
	if o != nil && !common.IsNil(o.Gamma) {
		return true
	}

	return false
}

// SetGamma gets a reference to the given string and assigns it to the Gamma field.
func (o *OptionMarkPriceResponseInner) SetGamma(v string) {
	o.Gamma = &v
}

// GetVega returns the Vega field value if set, zero value otherwise.
func (o *OptionMarkPriceResponseInner) GetVega() string {
	if o == nil || common.IsNil(o.Vega) {
		var ret string
		return ret
	}
	return *o.Vega
}

// GetVegaOk returns a tuple with the Vega field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarkPriceResponseInner) GetVegaOk() (*string, bool) {
	if o == nil || common.IsNil(o.Vega) {
		return nil, false
	}
	return o.Vega, true
}

// HasVega returns a boolean if a field has been set.
func (o *OptionMarkPriceResponseInner) HasVega() bool {
	if o != nil && !common.IsNil(o.Vega) {
		return true
	}

	return false
}

// SetVega gets a reference to the given string and assigns it to the Vega field.
func (o *OptionMarkPriceResponseInner) SetVega(v string) {
	o.Vega = &v
}

// GetHighPriceLimit returns the HighPriceLimit field value if set, zero value otherwise.
func (o *OptionMarkPriceResponseInner) GetHighPriceLimit() string {
	if o == nil || common.IsNil(o.HighPriceLimit) {
		var ret string
		return ret
	}
	return *o.HighPriceLimit
}

// GetHighPriceLimitOk returns a tuple with the HighPriceLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarkPriceResponseInner) GetHighPriceLimitOk() (*string, bool) {
	if o == nil || common.IsNil(o.HighPriceLimit) {
		return nil, false
	}
	return o.HighPriceLimit, true
}

// HasHighPriceLimit returns a boolean if a field has been set.
func (o *OptionMarkPriceResponseInner) HasHighPriceLimit() bool {
	if o != nil && !common.IsNil(o.HighPriceLimit) {
		return true
	}

	return false
}

// SetHighPriceLimit gets a reference to the given string and assigns it to the HighPriceLimit field.
func (o *OptionMarkPriceResponseInner) SetHighPriceLimit(v string) {
	o.HighPriceLimit = &v
}

// GetLowPriceLimit returns the LowPriceLimit field value if set, zero value otherwise.
func (o *OptionMarkPriceResponseInner) GetLowPriceLimit() string {
	if o == nil || common.IsNil(o.LowPriceLimit) {
		var ret string
		return ret
	}
	return *o.LowPriceLimit
}

// GetLowPriceLimitOk returns a tuple with the LowPriceLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarkPriceResponseInner) GetLowPriceLimitOk() (*string, bool) {
	if o == nil || common.IsNil(o.LowPriceLimit) {
		return nil, false
	}
	return o.LowPriceLimit, true
}

// HasLowPriceLimit returns a boolean if a field has been set.
func (o *OptionMarkPriceResponseInner) HasLowPriceLimit() bool {
	if o != nil && !common.IsNil(o.LowPriceLimit) {
		return true
	}

	return false
}

// SetLowPriceLimit gets a reference to the given string and assigns it to the LowPriceLimit field.
func (o *OptionMarkPriceResponseInner) SetLowPriceLimit(v string) {
	o.LowPriceLimit = &v
}

// GetRiskFreeInterest returns the RiskFreeInterest field value if set, zero value otherwise.
func (o *OptionMarkPriceResponseInner) GetRiskFreeInterest() string {
	if o == nil || common.IsNil(o.RiskFreeInterest) {
		var ret string
		return ret
	}
	return *o.RiskFreeInterest
}

// GetRiskFreeInterestOk returns a tuple with the RiskFreeInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarkPriceResponseInner) GetRiskFreeInterestOk() (*string, bool) {
	if o == nil || common.IsNil(o.RiskFreeInterest) {
		return nil, false
	}
	return o.RiskFreeInterest, true
}

// HasRiskFreeInterest returns a boolean if a field has been set.
func (o *OptionMarkPriceResponseInner) HasRiskFreeInterest() bool {
	if o != nil && !common.IsNil(o.RiskFreeInterest) {
		return true
	}

	return false
}

// SetRiskFreeInterest gets a reference to the given string and assigns it to the RiskFreeInterest field.
func (o *OptionMarkPriceResponseInner) SetRiskFreeInterest(v string) {
	o.RiskFreeInterest = &v
}

func (o OptionMarkPriceResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptionMarkPriceResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.MarkPrice) {
		toSerialize["markPrice"] = o.MarkPrice
	}
	if !common.IsNil(o.BidIV) {
		toSerialize["bidIV"] = o.BidIV
	}
	if !common.IsNil(o.AskIV) {
		toSerialize["askIV"] = o.AskIV
	}
	if !common.IsNil(o.MarkIV) {
		toSerialize["markIV"] = o.MarkIV
	}
	if !common.IsNil(o.Delta) {
		toSerialize["delta"] = o.Delta
	}
	if !common.IsNil(o.Theta) {
		toSerialize["theta"] = o.Theta
	}
	if !common.IsNil(o.Gamma) {
		toSerialize["gamma"] = o.Gamma
	}
	if !common.IsNil(o.Vega) {
		toSerialize["vega"] = o.Vega
	}
	if !common.IsNil(o.HighPriceLimit) {
		toSerialize["highPriceLimit"] = o.HighPriceLimit
	}
	if !common.IsNil(o.LowPriceLimit) {
		toSerialize["lowPriceLimit"] = o.LowPriceLimit
	}
	if !common.IsNil(o.RiskFreeInterest) {
		toSerialize["riskFreeInterest"] = o.RiskFreeInterest
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OptionMarkPriceResponseInner) UnmarshalJSON(data []byte) (err error) {
	varOptionMarkPriceResponseInner := _OptionMarkPriceResponseInner{}

	err = json.Unmarshal(data, &varOptionMarkPriceResponseInner)

	if err != nil {
		return err
	}

	*o = OptionMarkPriceResponseInner(varOptionMarkPriceResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "markPrice")
		delete(additionalProperties, "bidIV")
		delete(additionalProperties, "askIV")
		delete(additionalProperties, "markIV")
		delete(additionalProperties, "delta")
		delete(additionalProperties, "theta")
		delete(additionalProperties, "gamma")
		delete(additionalProperties, "vega")
		delete(additionalProperties, "highPriceLimit")
		delete(additionalProperties, "lowPriceLimit")
		delete(additionalProperties, "riskFreeInterest")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOptionMarkPriceResponseInner struct {
	value *OptionMarkPriceResponseInner
	isSet bool
}

func (v NullableOptionMarkPriceResponseInner) Get() *OptionMarkPriceResponseInner {
	return v.value
}

func (v *NullableOptionMarkPriceResponseInner) Set(val *OptionMarkPriceResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionMarkPriceResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionMarkPriceResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionMarkPriceResponseInner(val *OptionMarkPriceResponseInner) *NullableOptionMarkPriceResponseInner {
	return &NullableOptionMarkPriceResponseInner{value: val, isSet: true}
}

func (v NullableOptionMarkPriceResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionMarkPriceResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
