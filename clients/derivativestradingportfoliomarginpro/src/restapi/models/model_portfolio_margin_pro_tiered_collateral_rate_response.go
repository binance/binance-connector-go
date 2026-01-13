/*
Binance Derivatives Trading Portfolio Margin Pro REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin Pro REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the PortfolioMarginProTieredCollateralRateResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PortfolioMarginProTieredCollateralRateResponse{}

// PortfolioMarginProTieredCollateralRateResponse struct for PortfolioMarginProTieredCollateralRateResponse
type PortfolioMarginProTieredCollateralRateResponse struct {
	Items []PortfolioMarginProTieredCollateralRateResponseInner
}

// NewPortfolioMarginProTieredCollateralRateResponse instantiates a new PortfolioMarginProTieredCollateralRateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioMarginProTieredCollateralRateResponse() *PortfolioMarginProTieredCollateralRateResponse {
	this := PortfolioMarginProTieredCollateralRateResponse{}
	return &this
}

// NewPortfolioMarginProTieredCollateralRateResponseWithDefaults instantiates a new PortfolioMarginProTieredCollateralRateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioMarginProTieredCollateralRateResponseWithDefaults() *PortfolioMarginProTieredCollateralRateResponse {
	this := PortfolioMarginProTieredCollateralRateResponse{}
	return &this
}

func (o PortfolioMarginProTieredCollateralRateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortfolioMarginProTieredCollateralRateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *PortfolioMarginProTieredCollateralRateResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullablePortfolioMarginProTieredCollateralRateResponse struct {
	value PortfolioMarginProTieredCollateralRateResponse
	isSet bool
}

func (v NullablePortfolioMarginProTieredCollateralRateResponse) Get() PortfolioMarginProTieredCollateralRateResponse {
	return v.value
}

func (v *NullablePortfolioMarginProTieredCollateralRateResponse) Set(val PortfolioMarginProTieredCollateralRateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioMarginProTieredCollateralRateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioMarginProTieredCollateralRateResponse) Unset() {
	v.value = PortfolioMarginProTieredCollateralRateResponse{}
	v.isSet = false
}

func NewNullablePortfolioMarginProTieredCollateralRateResponse(val PortfolioMarginProTieredCollateralRateResponse) *NullablePortfolioMarginProTieredCollateralRateResponse {
	return &NullablePortfolioMarginProTieredCollateralRateResponse{value: val, isSet: true}
}

func (v NullablePortfolioMarginProTieredCollateralRateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortfolioMarginProTieredCollateralRateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
