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

// checks if the OptionAccountInformationResponseAssetInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OptionAccountInformationResponseAssetInner{}

// OptionAccountInformationResponseAssetInner struct for OptionAccountInformationResponseAssetInner
type OptionAccountInformationResponseAssetInner struct {
	Asset                *string `json:"asset,omitempty"`
	MarginBalance        *string `json:"marginBalance,omitempty"`
	Equity               *string `json:"equity,omitempty"`
	Available            *string `json:"available,omitempty"`
	Locked               *string `json:"locked,omitempty"`
	UnrealizedPNL        *string `json:"unrealizedPNL,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OptionAccountInformationResponseAssetInner OptionAccountInformationResponseAssetInner

// NewOptionAccountInformationResponseAssetInner instantiates a new OptionAccountInformationResponseAssetInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionAccountInformationResponseAssetInner() *OptionAccountInformationResponseAssetInner {
	this := OptionAccountInformationResponseAssetInner{}
	return &this
}

// NewOptionAccountInformationResponseAssetInnerWithDefaults instantiates a new OptionAccountInformationResponseAssetInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionAccountInformationResponseAssetInnerWithDefaults() *OptionAccountInformationResponseAssetInner {
	this := OptionAccountInformationResponseAssetInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *OptionAccountInformationResponseAssetInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionAccountInformationResponseAssetInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *OptionAccountInformationResponseAssetInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *OptionAccountInformationResponseAssetInner) SetAsset(v string) {
	o.Asset = &v
}

// GetMarginBalance returns the MarginBalance field value if set, zero value otherwise.
func (o *OptionAccountInformationResponseAssetInner) GetMarginBalance() string {
	if o == nil || common.IsNil(o.MarginBalance) {
		var ret string
		return ret
	}
	return *o.MarginBalance
}

// GetMarginBalanceOk returns a tuple with the MarginBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionAccountInformationResponseAssetInner) GetMarginBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarginBalance) {
		return nil, false
	}
	return o.MarginBalance, true
}

// HasMarginBalance returns a boolean if a field has been set.
func (o *OptionAccountInformationResponseAssetInner) HasMarginBalance() bool {
	if o != nil && !common.IsNil(o.MarginBalance) {
		return true
	}

	return false
}

// SetMarginBalance gets a reference to the given string and assigns it to the MarginBalance field.
func (o *OptionAccountInformationResponseAssetInner) SetMarginBalance(v string) {
	o.MarginBalance = &v
}

// GetEquity returns the Equity field value if set, zero value otherwise.
func (o *OptionAccountInformationResponseAssetInner) GetEquity() string {
	if o == nil || common.IsNil(o.Equity) {
		var ret string
		return ret
	}
	return *o.Equity
}

// GetEquityOk returns a tuple with the Equity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionAccountInformationResponseAssetInner) GetEquityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Equity) {
		return nil, false
	}
	return o.Equity, true
}

// HasEquity returns a boolean if a field has been set.
func (o *OptionAccountInformationResponseAssetInner) HasEquity() bool {
	if o != nil && !common.IsNil(o.Equity) {
		return true
	}

	return false
}

// SetEquity gets a reference to the given string and assigns it to the Equity field.
func (o *OptionAccountInformationResponseAssetInner) SetEquity(v string) {
	o.Equity = &v
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *OptionAccountInformationResponseAssetInner) GetAvailable() string {
	if o == nil || common.IsNil(o.Available) {
		var ret string
		return ret
	}
	return *o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionAccountInformationResponseAssetInner) GetAvailableOk() (*string, bool) {
	if o == nil || common.IsNil(o.Available) {
		return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *OptionAccountInformationResponseAssetInner) HasAvailable() bool {
	if o != nil && !common.IsNil(o.Available) {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given string and assigns it to the Available field.
func (o *OptionAccountInformationResponseAssetInner) SetAvailable(v string) {
	o.Available = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *OptionAccountInformationResponseAssetInner) GetLocked() string {
	if o == nil || common.IsNil(o.Locked) {
		var ret string
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionAccountInformationResponseAssetInner) GetLockedOk() (*string, bool) {
	if o == nil || common.IsNil(o.Locked) {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *OptionAccountInformationResponseAssetInner) HasLocked() bool {
	if o != nil && !common.IsNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given string and assigns it to the Locked field.
func (o *OptionAccountInformationResponseAssetInner) SetLocked(v string) {
	o.Locked = &v
}

// GetUnrealizedPNL returns the UnrealizedPNL field value if set, zero value otherwise.
func (o *OptionAccountInformationResponseAssetInner) GetUnrealizedPNL() string {
	if o == nil || common.IsNil(o.UnrealizedPNL) {
		var ret string
		return ret
	}
	return *o.UnrealizedPNL
}

// GetUnrealizedPNLOk returns a tuple with the UnrealizedPNL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionAccountInformationResponseAssetInner) GetUnrealizedPNLOk() (*string, bool) {
	if o == nil || common.IsNil(o.UnrealizedPNL) {
		return nil, false
	}
	return o.UnrealizedPNL, true
}

// HasUnrealizedPNL returns a boolean if a field has been set.
func (o *OptionAccountInformationResponseAssetInner) HasUnrealizedPNL() bool {
	if o != nil && !common.IsNil(o.UnrealizedPNL) {
		return true
	}

	return false
}

// SetUnrealizedPNL gets a reference to the given string and assigns it to the UnrealizedPNL field.
func (o *OptionAccountInformationResponseAssetInner) SetUnrealizedPNL(v string) {
	o.UnrealizedPNL = &v
}

func (o OptionAccountInformationResponseAssetInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptionAccountInformationResponseAssetInner) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.Locked) {
		toSerialize["locked"] = o.Locked
	}
	if !common.IsNil(o.UnrealizedPNL) {
		toSerialize["unrealizedPNL"] = o.UnrealizedPNL
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OptionAccountInformationResponseAssetInner) UnmarshalJSON(data []byte) (err error) {
	varOptionAccountInformationResponseAssetInner := _OptionAccountInformationResponseAssetInner{}

	err = json.Unmarshal(data, &varOptionAccountInformationResponseAssetInner)

	if err != nil {
		return err
	}

	*o = OptionAccountInformationResponseAssetInner(varOptionAccountInformationResponseAssetInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "marginBalance")
		delete(additionalProperties, "equity")
		delete(additionalProperties, "available")
		delete(additionalProperties, "locked")
		delete(additionalProperties, "unrealizedPNL")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOptionAccountInformationResponseAssetInner struct {
	value *OptionAccountInformationResponseAssetInner
	isSet bool
}

func (v NullableOptionAccountInformationResponseAssetInner) Get() *OptionAccountInformationResponseAssetInner {
	return v.value
}

func (v *NullableOptionAccountInformationResponseAssetInner) Set(val *OptionAccountInformationResponseAssetInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionAccountInformationResponseAssetInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionAccountInformationResponseAssetInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionAccountInformationResponseAssetInner(val *OptionAccountInformationResponseAssetInner) *NullableOptionAccountInformationResponseAssetInner {
	return &NullableOptionAccountInformationResponseAssetInner{value: val, isSet: true}
}

func (v NullableOptionAccountInformationResponseAssetInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionAccountInformationResponseAssetInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
