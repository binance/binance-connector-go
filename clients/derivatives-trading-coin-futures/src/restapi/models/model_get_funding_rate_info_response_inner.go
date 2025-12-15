/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetFundingRateInfoResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFundingRateInfoResponseInner{}

// GetFundingRateInfoResponseInner struct for GetFundingRateInfoResponseInner
type GetFundingRateInfoResponseInner struct {
	Symbol                   *string `json:"symbol,omitempty"`
	AdjustedFundingRateCap   *string `json:"adjustedFundingRateCap,omitempty"`
	AdjustedFundingRateFloor *string `json:"adjustedFundingRateFloor,omitempty"`
	FundingIntervalHours     *int64  `json:"fundingIntervalHours,omitempty"`
	Disclaimer               *bool   `json:"disclaimer,omitempty"`
	AdditionalProperties     map[string]interface{}
}

type _GetFundingRateInfoResponseInner GetFundingRateInfoResponseInner

// NewGetFundingRateInfoResponseInner instantiates a new GetFundingRateInfoResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFundingRateInfoResponseInner() *GetFundingRateInfoResponseInner {
	this := GetFundingRateInfoResponseInner{}
	return &this
}

// NewGetFundingRateInfoResponseInnerWithDefaults instantiates a new GetFundingRateInfoResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFundingRateInfoResponseInnerWithDefaults() *GetFundingRateInfoResponseInner {
	this := GetFundingRateInfoResponseInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetFundingRateInfoResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFundingRateInfoResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetFundingRateInfoResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetFundingRateInfoResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetAdjustedFundingRateCap returns the AdjustedFundingRateCap field value if set, zero value otherwise.
func (o *GetFundingRateInfoResponseInner) GetAdjustedFundingRateCap() string {
	if o == nil || common.IsNil(o.AdjustedFundingRateCap) {
		var ret string
		return ret
	}
	return *o.AdjustedFundingRateCap
}

// GetAdjustedFundingRateCapOk returns a tuple with the AdjustedFundingRateCap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFundingRateInfoResponseInner) GetAdjustedFundingRateCapOk() (*string, bool) {
	if o == nil || common.IsNil(o.AdjustedFundingRateCap) {
		return nil, false
	}
	return o.AdjustedFundingRateCap, true
}

// HasAdjustedFundingRateCap returns a boolean if a field has been set.
func (o *GetFundingRateInfoResponseInner) HasAdjustedFundingRateCap() bool {
	if o != nil && !common.IsNil(o.AdjustedFundingRateCap) {
		return true
	}

	return false
}

// SetAdjustedFundingRateCap gets a reference to the given string and assigns it to the AdjustedFundingRateCap field.
func (o *GetFundingRateInfoResponseInner) SetAdjustedFundingRateCap(v string) {
	o.AdjustedFundingRateCap = &v
}

// GetAdjustedFundingRateFloor returns the AdjustedFundingRateFloor field value if set, zero value otherwise.
func (o *GetFundingRateInfoResponseInner) GetAdjustedFundingRateFloor() string {
	if o == nil || common.IsNil(o.AdjustedFundingRateFloor) {
		var ret string
		return ret
	}
	return *o.AdjustedFundingRateFloor
}

// GetAdjustedFundingRateFloorOk returns a tuple with the AdjustedFundingRateFloor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFundingRateInfoResponseInner) GetAdjustedFundingRateFloorOk() (*string, bool) {
	if o == nil || common.IsNil(o.AdjustedFundingRateFloor) {
		return nil, false
	}
	return o.AdjustedFundingRateFloor, true
}

// HasAdjustedFundingRateFloor returns a boolean if a field has been set.
func (o *GetFundingRateInfoResponseInner) HasAdjustedFundingRateFloor() bool {
	if o != nil && !common.IsNil(o.AdjustedFundingRateFloor) {
		return true
	}

	return false
}

// SetAdjustedFundingRateFloor gets a reference to the given string and assigns it to the AdjustedFundingRateFloor field.
func (o *GetFundingRateInfoResponseInner) SetAdjustedFundingRateFloor(v string) {
	o.AdjustedFundingRateFloor = &v
}

// GetFundingIntervalHours returns the FundingIntervalHours field value if set, zero value otherwise.
func (o *GetFundingRateInfoResponseInner) GetFundingIntervalHours() int64 {
	if o == nil || common.IsNil(o.FundingIntervalHours) {
		var ret int64
		return ret
	}
	return *o.FundingIntervalHours
}

// GetFundingIntervalHoursOk returns a tuple with the FundingIntervalHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFundingRateInfoResponseInner) GetFundingIntervalHoursOk() (*int64, bool) {
	if o == nil || common.IsNil(o.FundingIntervalHours) {
		return nil, false
	}
	return o.FundingIntervalHours, true
}

// HasFundingIntervalHours returns a boolean if a field has been set.
func (o *GetFundingRateInfoResponseInner) HasFundingIntervalHours() bool {
	if o != nil && !common.IsNil(o.FundingIntervalHours) {
		return true
	}

	return false
}

// SetFundingIntervalHours gets a reference to the given int64 and assigns it to the FundingIntervalHours field.
func (o *GetFundingRateInfoResponseInner) SetFundingIntervalHours(v int64) {
	o.FundingIntervalHours = &v
}

// GetDisclaimer returns the Disclaimer field value if set, zero value otherwise.
func (o *GetFundingRateInfoResponseInner) GetDisclaimer() bool {
	if o == nil || common.IsNil(o.Disclaimer) {
		var ret bool
		return ret
	}
	return *o.Disclaimer
}

// GetDisclaimerOk returns a tuple with the Disclaimer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFundingRateInfoResponseInner) GetDisclaimerOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Disclaimer) {
		return nil, false
	}
	return o.Disclaimer, true
}

// HasDisclaimer returns a boolean if a field has been set.
func (o *GetFundingRateInfoResponseInner) HasDisclaimer() bool {
	if o != nil && !common.IsNil(o.Disclaimer) {
		return true
	}

	return false
}

// SetDisclaimer gets a reference to the given bool and assigns it to the Disclaimer field.
func (o *GetFundingRateInfoResponseInner) SetDisclaimer(v bool) {
	o.Disclaimer = &v
}

func (o GetFundingRateInfoResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFundingRateInfoResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.AdjustedFundingRateCap) {
		toSerialize["adjustedFundingRateCap"] = o.AdjustedFundingRateCap
	}
	if !common.IsNil(o.AdjustedFundingRateFloor) {
		toSerialize["adjustedFundingRateFloor"] = o.AdjustedFundingRateFloor
	}
	if !common.IsNil(o.FundingIntervalHours) {
		toSerialize["fundingIntervalHours"] = o.FundingIntervalHours
	}
	if !common.IsNil(o.Disclaimer) {
		toSerialize["disclaimer"] = o.Disclaimer
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetFundingRateInfoResponseInner) UnmarshalJSON(data []byte) (err error) {
	varGetFundingRateInfoResponseInner := _GetFundingRateInfoResponseInner{}

	err = json.Unmarshal(data, &varGetFundingRateInfoResponseInner)

	if err != nil {
		return err
	}

	*o = GetFundingRateInfoResponseInner(varGetFundingRateInfoResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "adjustedFundingRateCap")
		delete(additionalProperties, "adjustedFundingRateFloor")
		delete(additionalProperties, "fundingIntervalHours")
		delete(additionalProperties, "disclaimer")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFundingRateInfoResponseInner struct {
	value *GetFundingRateInfoResponseInner
	isSet bool
}

func (v NullableGetFundingRateInfoResponseInner) Get() *GetFundingRateInfoResponseInner {
	return v.value
}

func (v *NullableGetFundingRateInfoResponseInner) Set(val *GetFundingRateInfoResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFundingRateInfoResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFundingRateInfoResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFundingRateInfoResponseInner(val *GetFundingRateInfoResponseInner) *NullableGetFundingRateInfoResponseInner {
	return &NullableGetFundingRateInfoResponseInner{value: val, isSet: true}
}

func (v NullableGetFundingRateInfoResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFundingRateInfoResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
