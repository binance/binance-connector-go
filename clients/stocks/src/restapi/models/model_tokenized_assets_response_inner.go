/*
Stocks Trading REST API

REST APIs for Binance Stocks Trading. All endpoints under `/sapi/v1/equity/_*`.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the TokenizedAssetsResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TokenizedAssetsResponseInner{}

// TokenizedAssetsResponseInner struct for TokenizedAssetsResponseInner
type TokenizedAssetsResponseInner struct {
	// Tokenized asset code, e.g. `AAPLB`.
	AssetCode *string `json:"assetCode,omitempty"`
	// Human-readable display name of the tokenized asset.
	AssetName *string `json:"assetName,omitempty"`
	// Underlying US-equity ticker, e.g. `AAPL`. Use this when calling `/tokenized/mint` or `/tokenized/redeem`.
	UnderlyingEquitySymbol *string `json:"underlyingEquitySymbol,omitempty"`
	// Conversion multiplier between the underlying equity and the tokenized asset.
	Multiplier *string `json:"multiplier,omitempty"`
	// Whether the multiplier is currently valid and applicable.
	MultiplierValid      *bool `json:"multiplierValid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TokenizedAssetsResponseInner TokenizedAssetsResponseInner

// NewTokenizedAssetsResponseInner instantiates a new TokenizedAssetsResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenizedAssetsResponseInner() *TokenizedAssetsResponseInner {
	this := TokenizedAssetsResponseInner{}
	return &this
}

// NewTokenizedAssetsResponseInnerWithDefaults instantiates a new TokenizedAssetsResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenizedAssetsResponseInnerWithDefaults() *TokenizedAssetsResponseInner {
	this := TokenizedAssetsResponseInner{}
	return &this
}

// GetAssetCode returns the AssetCode field value if set, zero value otherwise.
func (o *TokenizedAssetsResponseInner) GetAssetCode() string {
	if o == nil || common.IsNil(o.AssetCode) {
		var ret string
		return ret
	}
	return *o.AssetCode
}

// GetAssetCodeOk returns a tuple with the AssetCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizedAssetsResponseInner) GetAssetCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.AssetCode) {
		return nil, false
	}
	return o.AssetCode, true
}

// HasAssetCode returns a boolean if a field has been set.
func (o *TokenizedAssetsResponseInner) HasAssetCode() bool {
	if o != nil && !common.IsNil(o.AssetCode) {
		return true
	}

	return false
}

// SetAssetCode gets a reference to the given string and assigns it to the AssetCode field.
func (o *TokenizedAssetsResponseInner) SetAssetCode(v string) {
	o.AssetCode = &v
}

// GetAssetName returns the AssetName field value if set, zero value otherwise.
func (o *TokenizedAssetsResponseInner) GetAssetName() string {
	if o == nil || common.IsNil(o.AssetName) {
		var ret string
		return ret
	}
	return *o.AssetName
}

// GetAssetNameOk returns a tuple with the AssetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizedAssetsResponseInner) GetAssetNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.AssetName) {
		return nil, false
	}
	return o.AssetName, true
}

// HasAssetName returns a boolean if a field has been set.
func (o *TokenizedAssetsResponseInner) HasAssetName() bool {
	if o != nil && !common.IsNil(o.AssetName) {
		return true
	}

	return false
}

// SetAssetName gets a reference to the given string and assigns it to the AssetName field.
func (o *TokenizedAssetsResponseInner) SetAssetName(v string) {
	o.AssetName = &v
}

// GetUnderlyingEquitySymbol returns the UnderlyingEquitySymbol field value if set, zero value otherwise.
func (o *TokenizedAssetsResponseInner) GetUnderlyingEquitySymbol() string {
	if o == nil || common.IsNil(o.UnderlyingEquitySymbol) {
		var ret string
		return ret
	}
	return *o.UnderlyingEquitySymbol
}

// GetUnderlyingEquitySymbolOk returns a tuple with the UnderlyingEquitySymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizedAssetsResponseInner) GetUnderlyingEquitySymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.UnderlyingEquitySymbol) {
		return nil, false
	}
	return o.UnderlyingEquitySymbol, true
}

// HasUnderlyingEquitySymbol returns a boolean if a field has been set.
func (o *TokenizedAssetsResponseInner) HasUnderlyingEquitySymbol() bool {
	if o != nil && !common.IsNil(o.UnderlyingEquitySymbol) {
		return true
	}

	return false
}

// SetUnderlyingEquitySymbol gets a reference to the given string and assigns it to the UnderlyingEquitySymbol field.
func (o *TokenizedAssetsResponseInner) SetUnderlyingEquitySymbol(v string) {
	o.UnderlyingEquitySymbol = &v
}

// GetMultiplier returns the Multiplier field value if set, zero value otherwise.
func (o *TokenizedAssetsResponseInner) GetMultiplier() string {
	if o == nil || common.IsNil(o.Multiplier) {
		var ret string
		return ret
	}
	return *o.Multiplier
}

// GetMultiplierOk returns a tuple with the Multiplier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizedAssetsResponseInner) GetMultiplierOk() (*string, bool) {
	if o == nil || common.IsNil(o.Multiplier) {
		return nil, false
	}
	return o.Multiplier, true
}

// HasMultiplier returns a boolean if a field has been set.
func (o *TokenizedAssetsResponseInner) HasMultiplier() bool {
	if o != nil && !common.IsNil(o.Multiplier) {
		return true
	}

	return false
}

// SetMultiplier gets a reference to the given string and assigns it to the Multiplier field.
func (o *TokenizedAssetsResponseInner) SetMultiplier(v string) {
	o.Multiplier = &v
}

// GetMultiplierValid returns the MultiplierValid field value if set, zero value otherwise.
func (o *TokenizedAssetsResponseInner) GetMultiplierValid() bool {
	if o == nil || common.IsNil(o.MultiplierValid) {
		var ret bool
		return ret
	}
	return *o.MultiplierValid
}

// GetMultiplierValidOk returns a tuple with the MultiplierValid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizedAssetsResponseInner) GetMultiplierValidOk() (*bool, bool) {
	if o == nil || common.IsNil(o.MultiplierValid) {
		return nil, false
	}
	return o.MultiplierValid, true
}

// HasMultiplierValid returns a boolean if a field has been set.
func (o *TokenizedAssetsResponseInner) HasMultiplierValid() bool {
	if o != nil && !common.IsNil(o.MultiplierValid) {
		return true
	}

	return false
}

// SetMultiplierValid gets a reference to the given bool and assigns it to the MultiplierValid field.
func (o *TokenizedAssetsResponseInner) SetMultiplierValid(v bool) {
	o.MultiplierValid = &v
}

func (o TokenizedAssetsResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenizedAssetsResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AssetCode) {
		toSerialize["assetCode"] = o.AssetCode
	}
	if !common.IsNil(o.AssetName) {
		toSerialize["assetName"] = o.AssetName
	}
	if !common.IsNil(o.UnderlyingEquitySymbol) {
		toSerialize["underlyingEquitySymbol"] = o.UnderlyingEquitySymbol
	}
	if !common.IsNil(o.Multiplier) {
		toSerialize["multiplier"] = o.Multiplier
	}
	if !common.IsNil(o.MultiplierValid) {
		toSerialize["multiplierValid"] = o.MultiplierValid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TokenizedAssetsResponseInner) UnmarshalJSON(data []byte) (err error) {
	varTokenizedAssetsResponseInner := _TokenizedAssetsResponseInner{}

	err = json.Unmarshal(data, &varTokenizedAssetsResponseInner)

	if err != nil {
		return err
	}

	*o = TokenizedAssetsResponseInner(varTokenizedAssetsResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "assetCode")
		delete(additionalProperties, "assetName")
		delete(additionalProperties, "underlyingEquitySymbol")
		delete(additionalProperties, "multiplier")
		delete(additionalProperties, "multiplierValid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTokenizedAssetsResponseInner struct {
	value *TokenizedAssetsResponseInner
	isSet bool
}

func (v NullableTokenizedAssetsResponseInner) Get() *TokenizedAssetsResponseInner {
	return v.value
}

func (v *NullableTokenizedAssetsResponseInner) Set(val *TokenizedAssetsResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenizedAssetsResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenizedAssetsResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenizedAssetsResponseInner(val *TokenizedAssetsResponseInner) *NullableTokenizedAssetsResponseInner {
	return &NullableTokenizedAssetsResponseInner{value: val, isSet: true}
}

func (v NullableTokenizedAssetsResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenizedAssetsResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
