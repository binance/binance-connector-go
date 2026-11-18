/*
Portfolio Margin Pro REST API

Access advanced account management and high-frequency trading with Binance Portfolio Margin Pro.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetPortfolioMarginTieredCollateralRatesResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetPortfolioMarginTieredCollateralRatesResponse{}

// GetPortfolioMarginTieredCollateralRatesResponse struct for GetPortfolioMarginTieredCollateralRatesResponse
type GetPortfolioMarginTieredCollateralRatesResponse struct {
	Items []GetPortfolioMarginTieredCollateralRatesResponseInner
}

// NewGetPortfolioMarginTieredCollateralRatesResponse instantiates a new GetPortfolioMarginTieredCollateralRatesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPortfolioMarginTieredCollateralRatesResponse() *GetPortfolioMarginTieredCollateralRatesResponse {
	this := GetPortfolioMarginTieredCollateralRatesResponse{}
	return &this
}

// NewGetPortfolioMarginTieredCollateralRatesResponseWithDefaults instantiates a new GetPortfolioMarginTieredCollateralRatesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPortfolioMarginTieredCollateralRatesResponseWithDefaults() *GetPortfolioMarginTieredCollateralRatesResponse {
	this := GetPortfolioMarginTieredCollateralRatesResponse{}
	return &this
}

func (o GetPortfolioMarginTieredCollateralRatesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPortfolioMarginTieredCollateralRatesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *GetPortfolioMarginTieredCollateralRatesResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableGetPortfolioMarginTieredCollateralRatesResponse struct {
	value GetPortfolioMarginTieredCollateralRatesResponse
	isSet bool
}

func (v NullableGetPortfolioMarginTieredCollateralRatesResponse) Get() GetPortfolioMarginTieredCollateralRatesResponse {
	return v.value
}

func (v *NullableGetPortfolioMarginTieredCollateralRatesResponse) Set(val GetPortfolioMarginTieredCollateralRatesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPortfolioMarginTieredCollateralRatesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPortfolioMarginTieredCollateralRatesResponse) Unset() {
	v.value = GetPortfolioMarginTieredCollateralRatesResponse{}
	v.isSet = false
}

func NewNullableGetPortfolioMarginTieredCollateralRatesResponse(val GetPortfolioMarginTieredCollateralRatesResponse) *NullableGetPortfolioMarginTieredCollateralRatesResponse {
	return &NullableGetPortfolioMarginTieredCollateralRatesResponse{value: val, isSet: true}
}

func (v NullableGetPortfolioMarginTieredCollateralRatesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPortfolioMarginTieredCollateralRatesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
