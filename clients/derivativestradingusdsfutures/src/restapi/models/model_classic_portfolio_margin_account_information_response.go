/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the ClassicPortfolioMarginAccountInformationResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ClassicPortfolioMarginAccountInformationResponse{}

// ClassicPortfolioMarginAccountInformationResponse struct for ClassicPortfolioMarginAccountInformationResponse
type ClassicPortfolioMarginAccountInformationResponse struct {
	MaxWithdrawAmountUSD *string `json:"maxWithdrawAmountUSD,omitempty"`
	Asset                *string `json:"asset,omitempty"`
	MaxWithdrawAmount    *string `json:"maxWithdrawAmount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ClassicPortfolioMarginAccountInformationResponse ClassicPortfolioMarginAccountInformationResponse

// NewClassicPortfolioMarginAccountInformationResponse instantiates a new ClassicPortfolioMarginAccountInformationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClassicPortfolioMarginAccountInformationResponse() *ClassicPortfolioMarginAccountInformationResponse {
	this := ClassicPortfolioMarginAccountInformationResponse{}
	return &this
}

// NewClassicPortfolioMarginAccountInformationResponseWithDefaults instantiates a new ClassicPortfolioMarginAccountInformationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClassicPortfolioMarginAccountInformationResponseWithDefaults() *ClassicPortfolioMarginAccountInformationResponse {
	this := ClassicPortfolioMarginAccountInformationResponse{}
	return &this
}

// GetMaxWithdrawAmountUSD returns the MaxWithdrawAmountUSD field value if set, zero value otherwise.
func (o *ClassicPortfolioMarginAccountInformationResponse) GetMaxWithdrawAmountUSD() string {
	if o == nil || common.IsNil(o.MaxWithdrawAmountUSD) {
		var ret string
		return ret
	}
	return *o.MaxWithdrawAmountUSD
}

// GetMaxWithdrawAmountUSDOk returns a tuple with the MaxWithdrawAmountUSD field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassicPortfolioMarginAccountInformationResponse) GetMaxWithdrawAmountUSDOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxWithdrawAmountUSD) {
		return nil, false
	}
	return o.MaxWithdrawAmountUSD, true
}

// HasMaxWithdrawAmountUSD returns a boolean if a field has been set.
func (o *ClassicPortfolioMarginAccountInformationResponse) HasMaxWithdrawAmountUSD() bool {
	if o != nil && !common.IsNil(o.MaxWithdrawAmountUSD) {
		return true
	}

	return false
}

// SetMaxWithdrawAmountUSD gets a reference to the given string and assigns it to the MaxWithdrawAmountUSD field.
func (o *ClassicPortfolioMarginAccountInformationResponse) SetMaxWithdrawAmountUSD(v string) {
	o.MaxWithdrawAmountUSD = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *ClassicPortfolioMarginAccountInformationResponse) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassicPortfolioMarginAccountInformationResponse) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *ClassicPortfolioMarginAccountInformationResponse) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *ClassicPortfolioMarginAccountInformationResponse) SetAsset(v string) {
	o.Asset = &v
}

// GetMaxWithdrawAmount returns the MaxWithdrawAmount field value if set, zero value otherwise.
func (o *ClassicPortfolioMarginAccountInformationResponse) GetMaxWithdrawAmount() string {
	if o == nil || common.IsNil(o.MaxWithdrawAmount) {
		var ret string
		return ret
	}
	return *o.MaxWithdrawAmount
}

// GetMaxWithdrawAmountOk returns a tuple with the MaxWithdrawAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassicPortfolioMarginAccountInformationResponse) GetMaxWithdrawAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.MaxWithdrawAmount) {
		return nil, false
	}
	return o.MaxWithdrawAmount, true
}

// HasMaxWithdrawAmount returns a boolean if a field has been set.
func (o *ClassicPortfolioMarginAccountInformationResponse) HasMaxWithdrawAmount() bool {
	if o != nil && !common.IsNil(o.MaxWithdrawAmount) {
		return true
	}

	return false
}

// SetMaxWithdrawAmount gets a reference to the given string and assigns it to the MaxWithdrawAmount field.
func (o *ClassicPortfolioMarginAccountInformationResponse) SetMaxWithdrawAmount(v string) {
	o.MaxWithdrawAmount = &v
}

func (o ClassicPortfolioMarginAccountInformationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClassicPortfolioMarginAccountInformationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.MaxWithdrawAmountUSD) {
		toSerialize["maxWithdrawAmountUSD"] = o.MaxWithdrawAmountUSD
	}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.MaxWithdrawAmount) {
		toSerialize["maxWithdrawAmount"] = o.MaxWithdrawAmount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ClassicPortfolioMarginAccountInformationResponse) UnmarshalJSON(data []byte) (err error) {
	varClassicPortfolioMarginAccountInformationResponse := _ClassicPortfolioMarginAccountInformationResponse{}

	err = json.Unmarshal(data, &varClassicPortfolioMarginAccountInformationResponse)

	if err != nil {
		return err
	}

	*o = ClassicPortfolioMarginAccountInformationResponse(varClassicPortfolioMarginAccountInformationResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "maxWithdrawAmountUSD")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "maxWithdrawAmount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClassicPortfolioMarginAccountInformationResponse struct {
	value *ClassicPortfolioMarginAccountInformationResponse
	isSet bool
}

func (v NullableClassicPortfolioMarginAccountInformationResponse) Get() *ClassicPortfolioMarginAccountInformationResponse {
	return v.value
}

func (v *NullableClassicPortfolioMarginAccountInformationResponse) Set(val *ClassicPortfolioMarginAccountInformationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableClassicPortfolioMarginAccountInformationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableClassicPortfolioMarginAccountInformationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClassicPortfolioMarginAccountInformationResponse(val *ClassicPortfolioMarginAccountInformationResponse) *NullableClassicPortfolioMarginAccountInformationResponse {
	return &NullableClassicPortfolioMarginAccountInformationResponse{value: val, isSet: true}
}

func (v NullableClassicPortfolioMarginAccountInformationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClassicPortfolioMarginAccountInformationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
