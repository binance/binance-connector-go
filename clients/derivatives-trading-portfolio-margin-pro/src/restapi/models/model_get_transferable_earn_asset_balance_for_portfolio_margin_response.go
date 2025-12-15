/*
Binance Derivatives Trading Portfolio Margin Pro REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin Pro REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetTransferableEarnAssetBalanceForPortfolioMarginResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetTransferableEarnAssetBalanceForPortfolioMarginResponse{}

// GetTransferableEarnAssetBalanceForPortfolioMarginResponse struct for GetTransferableEarnAssetBalanceForPortfolioMarginResponse
type GetTransferableEarnAssetBalanceForPortfolioMarginResponse struct {
	Asset                *string `json:"asset,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetTransferableEarnAssetBalanceForPortfolioMarginResponse GetTransferableEarnAssetBalanceForPortfolioMarginResponse

// NewGetTransferableEarnAssetBalanceForPortfolioMarginResponse instantiates a new GetTransferableEarnAssetBalanceForPortfolioMarginResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransferableEarnAssetBalanceForPortfolioMarginResponse() *GetTransferableEarnAssetBalanceForPortfolioMarginResponse {
	this := GetTransferableEarnAssetBalanceForPortfolioMarginResponse{}
	return &this
}

// NewGetTransferableEarnAssetBalanceForPortfolioMarginResponseWithDefaults instantiates a new GetTransferableEarnAssetBalanceForPortfolioMarginResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransferableEarnAssetBalanceForPortfolioMarginResponseWithDefaults() *GetTransferableEarnAssetBalanceForPortfolioMarginResponse {
	this := GetTransferableEarnAssetBalanceForPortfolioMarginResponse{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetTransferableEarnAssetBalanceForPortfolioMarginResponse) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransferableEarnAssetBalanceForPortfolioMarginResponse) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetTransferableEarnAssetBalanceForPortfolioMarginResponse) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetTransferableEarnAssetBalanceForPortfolioMarginResponse) SetAsset(v string) {
	o.Asset = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetTransferableEarnAssetBalanceForPortfolioMarginResponse) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransferableEarnAssetBalanceForPortfolioMarginResponse) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetTransferableEarnAssetBalanceForPortfolioMarginResponse) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *GetTransferableEarnAssetBalanceForPortfolioMarginResponse) SetAmount(v string) {
	o.Amount = &v
}

func (o GetTransferableEarnAssetBalanceForPortfolioMarginResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTransferableEarnAssetBalanceForPortfolioMarginResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetTransferableEarnAssetBalanceForPortfolioMarginResponse) UnmarshalJSON(data []byte) (err error) {
	varGetTransferableEarnAssetBalanceForPortfolioMarginResponse := _GetTransferableEarnAssetBalanceForPortfolioMarginResponse{}

	err = json.Unmarshal(data, &varGetTransferableEarnAssetBalanceForPortfolioMarginResponse)

	if err != nil {
		return err
	}

	*o = GetTransferableEarnAssetBalanceForPortfolioMarginResponse(varGetTransferableEarnAssetBalanceForPortfolioMarginResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "amount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetTransferableEarnAssetBalanceForPortfolioMarginResponse struct {
	value *GetTransferableEarnAssetBalanceForPortfolioMarginResponse
	isSet bool
}

func (v NullableGetTransferableEarnAssetBalanceForPortfolioMarginResponse) Get() *GetTransferableEarnAssetBalanceForPortfolioMarginResponse {
	return v.value
}

func (v *NullableGetTransferableEarnAssetBalanceForPortfolioMarginResponse) Set(val *GetTransferableEarnAssetBalanceForPortfolioMarginResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransferableEarnAssetBalanceForPortfolioMarginResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransferableEarnAssetBalanceForPortfolioMarginResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransferableEarnAssetBalanceForPortfolioMarginResponse(val *GetTransferableEarnAssetBalanceForPortfolioMarginResponse) *NullableGetTransferableEarnAssetBalanceForPortfolioMarginResponse {
	return &NullableGetTransferableEarnAssetBalanceForPortfolioMarginResponse{value: val, isSet: true}
}

func (v NullableGetTransferableEarnAssetBalanceForPortfolioMarginResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransferableEarnAssetBalanceForPortfolioMarginResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
