/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetMarginAssetRiskBasedLiquidationRatioResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetMarginAssetRiskBasedLiquidationRatioResponse{}

// GetMarginAssetRiskBasedLiquidationRatioResponse struct for GetMarginAssetRiskBasedLiquidationRatioResponse
type GetMarginAssetRiskBasedLiquidationRatioResponse struct {
	Items []GetMarginAssetRiskBasedLiquidationRatioResponseInner
}

// NewGetMarginAssetRiskBasedLiquidationRatioResponse instantiates a new GetMarginAssetRiskBasedLiquidationRatioResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMarginAssetRiskBasedLiquidationRatioResponse() *GetMarginAssetRiskBasedLiquidationRatioResponse {
	this := GetMarginAssetRiskBasedLiquidationRatioResponse{}
	return &this
}

// NewGetMarginAssetRiskBasedLiquidationRatioResponseWithDefaults instantiates a new GetMarginAssetRiskBasedLiquidationRatioResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMarginAssetRiskBasedLiquidationRatioResponseWithDefaults() *GetMarginAssetRiskBasedLiquidationRatioResponse {
	this := GetMarginAssetRiskBasedLiquidationRatioResponse{}
	return &this
}

func (o GetMarginAssetRiskBasedLiquidationRatioResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMarginAssetRiskBasedLiquidationRatioResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *GetMarginAssetRiskBasedLiquidationRatioResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableGetMarginAssetRiskBasedLiquidationRatioResponse struct {
	value GetMarginAssetRiskBasedLiquidationRatioResponse
	isSet bool
}

func (v NullableGetMarginAssetRiskBasedLiquidationRatioResponse) Get() GetMarginAssetRiskBasedLiquidationRatioResponse {
	return v.value
}

func (v *NullableGetMarginAssetRiskBasedLiquidationRatioResponse) Set(val GetMarginAssetRiskBasedLiquidationRatioResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMarginAssetRiskBasedLiquidationRatioResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMarginAssetRiskBasedLiquidationRatioResponse) Unset() {
	v.value = GetMarginAssetRiskBasedLiquidationRatioResponse{}
	v.isSet = false
}

func NewNullableGetMarginAssetRiskBasedLiquidationRatioResponse(val GetMarginAssetRiskBasedLiquidationRatioResponse) *NullableGetMarginAssetRiskBasedLiquidationRatioResponse {
	return &NullableGetMarginAssetRiskBasedLiquidationRatioResponse{value: val, isSet: true}
}

func (v NullableGetMarginAssetRiskBasedLiquidationRatioResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMarginAssetRiskBasedLiquidationRatioResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
