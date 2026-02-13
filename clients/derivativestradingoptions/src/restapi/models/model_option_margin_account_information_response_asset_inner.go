/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the OptionMarginAccountInformationResponseAssetInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OptionMarginAccountInformationResponseAssetInner{}

// OptionMarginAccountInformationResponseAssetInner struct for OptionMarginAccountInformationResponseAssetInner
type OptionMarginAccountInformationResponseAssetInner struct {
	Asset                *string `json:"asset,omitempty"`
	MarginBalance        *string `json:"marginBalance,omitempty"`
	Equity               *string `json:"equity,omitempty"`
	Available            *string `json:"available,omitempty"`
	InitialMargin        *string `json:"initialMargin,omitempty"`
	MaintMargin          *string `json:"maintMargin,omitempty"`
	UnrealizedPNL        *string `json:"unrealizedPNL,omitempty"`
	AdjustedEquity       *string `json:"adjustedEquity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OptionMarginAccountInformationResponseAssetInner OptionMarginAccountInformationResponseAssetInner

// NewOptionMarginAccountInformationResponseAssetInner instantiates a new OptionMarginAccountInformationResponseAssetInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionMarginAccountInformationResponseAssetInner() *OptionMarginAccountInformationResponseAssetInner {
	this := OptionMarginAccountInformationResponseAssetInner{}
	return &this
}

// NewOptionMarginAccountInformationResponseAssetInnerWithDefaults instantiates a new OptionMarginAccountInformationResponseAssetInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionMarginAccountInformationResponseAssetInnerWithDefaults() *OptionMarginAccountInformationResponseAssetInner {
	this := OptionMarginAccountInformationResponseAssetInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *OptionMarginAccountInformationResponseAssetInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarginAccountInformationResponseAssetInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *OptionMarginAccountInformationResponseAssetInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *OptionMarginAccountInformationResponseAssetInner) SetAsset(v string) {
	o.Asset = &v
}

// GetMarginBalance returns the MarginBalance field value if set, zero value otherwise.
func (o *OptionMarginAccountInformationResponseAssetInner) GetMarginBalance() string {
	if o == nil || common.IsNil(o.MarginBalance) {
		var ret string
		return ret
	}
	return *o.MarginBalance
}

// GetMarginBalanceOk returns a tuple with the MarginBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarginAccountInformationResponseAssetInner) GetMarginBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarginBalance) {
		return nil, false
	}
	return o.MarginBalance, true
}

// HasMarginBalance returns a boolean if a field has been set.
func (o *OptionMarginAccountInformationResponseAssetInner) HasMarginBalance() bool {
	if o != nil && !common.IsNil(o.MarginBalance) {
		return true
	}

	return false
}

// SetMarginBalance gets a reference to the given string and assigns it to the MarginBalance field.
func (o *OptionMarginAccountInformationResponseAssetInner) SetMarginBalance(v string) {
	o.MarginBalance = &v
}

// GetEquity returns the Equity field value if set, zero value otherwise.
func (o *OptionMarginAccountInformationResponseAssetInner) GetEquity() string {
	if o == nil || common.IsNil(o.Equity) {
		var ret string
		return ret
	}
	return *o.Equity
}

// GetEquityOk returns a tuple with the Equity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarginAccountInformationResponseAssetInner) GetEquityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Equity) {
		return nil, false
	}
	return o.Equity, true
}

// HasEquity returns a boolean if a field has been set.
func (o *OptionMarginAccountInformationResponseAssetInner) HasEquity() bool {
	if o != nil && !common.IsNil(o.Equity) {
		return true
	}

	return false
}

// SetEquity gets a reference to the given string and assigns it to the Equity field.
func (o *OptionMarginAccountInformationResponseAssetInner) SetEquity(v string) {
	o.Equity = &v
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *OptionMarginAccountInformationResponseAssetInner) GetAvailable() string {
	if o == nil || common.IsNil(o.Available) {
		var ret string
		return ret
	}
	return *o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarginAccountInformationResponseAssetInner) GetAvailableOk() (*string, bool) {
	if o == nil || common.IsNil(o.Available) {
		return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *OptionMarginAccountInformationResponseAssetInner) HasAvailable() bool {
	if o != nil && !common.IsNil(o.Available) {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given string and assigns it to the Available field.
func (o *OptionMarginAccountInformationResponseAssetInner) SetAvailable(v string) {
	o.Available = &v
}

// GetInitialMargin returns the InitialMargin field value if set, zero value otherwise.
func (o *OptionMarginAccountInformationResponseAssetInner) GetInitialMargin() string {
	if o == nil || common.IsNil(o.InitialMargin) {
		var ret string
		return ret
	}
	return *o.InitialMargin
}

// GetInitialMarginOk returns a tuple with the InitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarginAccountInformationResponseAssetInner) GetInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.InitialMargin) {
		return nil, false
	}
	return o.InitialMargin, true
}

// HasInitialMargin returns a boolean if a field has been set.
func (o *OptionMarginAccountInformationResponseAssetInner) HasInitialMargin() bool {
	if o != nil && !common.IsNil(o.InitialMargin) {
		return true
	}

	return false
}

// SetInitialMargin gets a reference to the given string and assigns it to the InitialMargin field.
func (o *OptionMarginAccountInformationResponseAssetInner) SetInitialMargin(v string) {
	o.InitialMargin = &v
}

// GetMaintMargin returns the MaintMargin field value if set, zero value otherwise.
func (o *OptionMarginAccountInformationResponseAssetInner) GetMaintMargin() string {
	if o == nil || common.IsNil(o.MaintMargin) {
		var ret string
		return ret
	}
	return *o.MaintMargin
}

// GetMaintMarginOk returns a tuple with the MaintMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarginAccountInformationResponseAssetInner) GetMaintMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaintMargin) {
		return nil, false
	}
	return o.MaintMargin, true
}

// HasMaintMargin returns a boolean if a field has been set.
func (o *OptionMarginAccountInformationResponseAssetInner) HasMaintMargin() bool {
	if o != nil && !common.IsNil(o.MaintMargin) {
		return true
	}

	return false
}

// SetMaintMargin gets a reference to the given string and assigns it to the MaintMargin field.
func (o *OptionMarginAccountInformationResponseAssetInner) SetMaintMargin(v string) {
	o.MaintMargin = &v
}

// GetUnrealizedPNL returns the UnrealizedPNL field value if set, zero value otherwise.
func (o *OptionMarginAccountInformationResponseAssetInner) GetUnrealizedPNL() string {
	if o == nil || common.IsNil(o.UnrealizedPNL) {
		var ret string
		return ret
	}
	return *o.UnrealizedPNL
}

// GetUnrealizedPNLOk returns a tuple with the UnrealizedPNL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarginAccountInformationResponseAssetInner) GetUnrealizedPNLOk() (*string, bool) {
	if o == nil || common.IsNil(o.UnrealizedPNL) {
		return nil, false
	}
	return o.UnrealizedPNL, true
}

// HasUnrealizedPNL returns a boolean if a field has been set.
func (o *OptionMarginAccountInformationResponseAssetInner) HasUnrealizedPNL() bool {
	if o != nil && !common.IsNil(o.UnrealizedPNL) {
		return true
	}

	return false
}

// SetUnrealizedPNL gets a reference to the given string and assigns it to the UnrealizedPNL field.
func (o *OptionMarginAccountInformationResponseAssetInner) SetUnrealizedPNL(v string) {
	o.UnrealizedPNL = &v
}

// GetAdjustedEquity returns the AdjustedEquity field value if set, zero value otherwise.
func (o *OptionMarginAccountInformationResponseAssetInner) GetAdjustedEquity() string {
	if o == nil || common.IsNil(o.AdjustedEquity) {
		var ret string
		return ret
	}
	return *o.AdjustedEquity
}

// GetAdjustedEquityOk returns a tuple with the AdjustedEquity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarginAccountInformationResponseAssetInner) GetAdjustedEquityOk() (*string, bool) {
	if o == nil || common.IsNil(o.AdjustedEquity) {
		return nil, false
	}
	return o.AdjustedEquity, true
}

// HasAdjustedEquity returns a boolean if a field has been set.
func (o *OptionMarginAccountInformationResponseAssetInner) HasAdjustedEquity() bool {
	if o != nil && !common.IsNil(o.AdjustedEquity) {
		return true
	}

	return false
}

// SetAdjustedEquity gets a reference to the given string and assigns it to the AdjustedEquity field.
func (o *OptionMarginAccountInformationResponseAssetInner) SetAdjustedEquity(v string) {
	o.AdjustedEquity = &v
}

func (o OptionMarginAccountInformationResponseAssetInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptionMarginAccountInformationResponseAssetInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.MarginBalance) {
		toSerialize["marginBalance"] = o.MarginBalance
	}
	if !common.IsNil(o.Equity) {
		toSerialize["equity"] = o.Equity
	}
	if !common.IsNil(o.Available) {
		toSerialize["available"] = o.Available
	}
	if !common.IsNil(o.InitialMargin) {
		toSerialize["initialMargin"] = o.InitialMargin
	}
	if !common.IsNil(o.MaintMargin) {
		toSerialize["maintMargin"] = o.MaintMargin
	}
	if !common.IsNil(o.UnrealizedPNL) {
		toSerialize["unrealizedPNL"] = o.UnrealizedPNL
	}
	if !common.IsNil(o.AdjustedEquity) {
		toSerialize["adjustedEquity"] = o.AdjustedEquity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OptionMarginAccountInformationResponseAssetInner) UnmarshalJSON(data []byte) (err error) {
	varOptionMarginAccountInformationResponseAssetInner := _OptionMarginAccountInformationResponseAssetInner{}

	err = json.Unmarshal(data, &varOptionMarginAccountInformationResponseAssetInner)

	if err != nil {
		return err
	}

	*o = OptionMarginAccountInformationResponseAssetInner(varOptionMarginAccountInformationResponseAssetInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "marginBalance")
		delete(additionalProperties, "equity")
		delete(additionalProperties, "available")
		delete(additionalProperties, "initialMargin")
		delete(additionalProperties, "maintMargin")
		delete(additionalProperties, "unrealizedPNL")
		delete(additionalProperties, "adjustedEquity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOptionMarginAccountInformationResponseAssetInner struct {
	value *OptionMarginAccountInformationResponseAssetInner
	isSet bool
}

func (v NullableOptionMarginAccountInformationResponseAssetInner) Get() *OptionMarginAccountInformationResponseAssetInner {
	return v.value
}

func (v *NullableOptionMarginAccountInformationResponseAssetInner) Set(val *OptionMarginAccountInformationResponseAssetInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionMarginAccountInformationResponseAssetInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionMarginAccountInformationResponseAssetInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionMarginAccountInformationResponseAssetInner(val *OptionMarginAccountInformationResponseAssetInner) *NullableOptionMarginAccountInformationResponseAssetInner {
	return &NullableOptionMarginAccountInformationResponseAssetInner{value: val, isSet: true}
}

func (v NullableOptionMarginAccountInformationResponseAssetInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionMarginAccountInformationResponseAssetInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
