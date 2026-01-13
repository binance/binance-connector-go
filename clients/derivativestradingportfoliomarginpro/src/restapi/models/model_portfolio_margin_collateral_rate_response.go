/*
Binance Derivatives Trading Portfolio Margin Pro REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin Pro REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the PortfolioMarginCollateralRateResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PortfolioMarginCollateralRateResponse{}

// PortfolioMarginCollateralRateResponse struct for PortfolioMarginCollateralRateResponse
type PortfolioMarginCollateralRateResponse struct {
	Items []PortfolioMarginCollateralRateResponseInner
}

// NewPortfolioMarginCollateralRateResponse instantiates a new PortfolioMarginCollateralRateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioMarginCollateralRateResponse() *PortfolioMarginCollateralRateResponse {
	this := PortfolioMarginCollateralRateResponse{}
	return &this
}

// NewPortfolioMarginCollateralRateResponseWithDefaults instantiates a new PortfolioMarginCollateralRateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioMarginCollateralRateResponseWithDefaults() *PortfolioMarginCollateralRateResponse {
	this := PortfolioMarginCollateralRateResponse{}
	return &this
}

func (o PortfolioMarginCollateralRateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortfolioMarginCollateralRateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *PortfolioMarginCollateralRateResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullablePortfolioMarginCollateralRateResponse struct {
	value PortfolioMarginCollateralRateResponse
	isSet bool
}

func (v NullablePortfolioMarginCollateralRateResponse) Get() PortfolioMarginCollateralRateResponse {
	return v.value
}

func (v *NullablePortfolioMarginCollateralRateResponse) Set(val PortfolioMarginCollateralRateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioMarginCollateralRateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioMarginCollateralRateResponse) Unset() {
	v.value = PortfolioMarginCollateralRateResponse{}
	v.isSet = false
}

func NewNullablePortfolioMarginCollateralRateResponse(val PortfolioMarginCollateralRateResponse) *NullablePortfolioMarginCollateralRateResponse {
	return &NullablePortfolioMarginCollateralRateResponse{value: val, isSet: true}
}

func (v NullablePortfolioMarginCollateralRateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortfolioMarginCollateralRateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
