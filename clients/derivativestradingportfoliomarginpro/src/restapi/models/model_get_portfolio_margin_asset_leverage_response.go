/*
Binance Derivatives Trading Portfolio Margin Pro REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin Pro REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetPortfolioMarginAssetLeverageResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetPortfolioMarginAssetLeverageResponse{}

// GetPortfolioMarginAssetLeverageResponse struct for GetPortfolioMarginAssetLeverageResponse
type GetPortfolioMarginAssetLeverageResponse struct {
	Items []GetPortfolioMarginAssetLeverageResponseInner
}

// NewGetPortfolioMarginAssetLeverageResponse instantiates a new GetPortfolioMarginAssetLeverageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPortfolioMarginAssetLeverageResponse() *GetPortfolioMarginAssetLeverageResponse {
	this := GetPortfolioMarginAssetLeverageResponse{}
	return &this
}

// NewGetPortfolioMarginAssetLeverageResponseWithDefaults instantiates a new GetPortfolioMarginAssetLeverageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPortfolioMarginAssetLeverageResponseWithDefaults() *GetPortfolioMarginAssetLeverageResponse {
	this := GetPortfolioMarginAssetLeverageResponse{}
	return &this
}

func (o GetPortfolioMarginAssetLeverageResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPortfolioMarginAssetLeverageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *GetPortfolioMarginAssetLeverageResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableGetPortfolioMarginAssetLeverageResponse struct {
	value GetPortfolioMarginAssetLeverageResponse
	isSet bool
}

func (v NullableGetPortfolioMarginAssetLeverageResponse) Get() GetPortfolioMarginAssetLeverageResponse {
	return v.value
}

func (v *NullableGetPortfolioMarginAssetLeverageResponse) Set(val GetPortfolioMarginAssetLeverageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPortfolioMarginAssetLeverageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPortfolioMarginAssetLeverageResponse) Unset() {
	v.value = GetPortfolioMarginAssetLeverageResponse{}
	v.isSet = false
}

func NewNullableGetPortfolioMarginAssetLeverageResponse(val GetPortfolioMarginAssetLeverageResponse) *NullableGetPortfolioMarginAssetLeverageResponse {
	return &NullableGetPortfolioMarginAssetLeverageResponse{value: val, isSet: true}
}

func (v NullableGetPortfolioMarginAssetLeverageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPortfolioMarginAssetLeverageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
