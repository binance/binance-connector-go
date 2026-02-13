/*
Binance Derivatives Trading Portfolio Margin Pro REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin Pro REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryPortfolioMarginAssetIndexPriceResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryPortfolioMarginAssetIndexPriceResponse{}

// QueryPortfolioMarginAssetIndexPriceResponse struct for QueryPortfolioMarginAssetIndexPriceResponse
type QueryPortfolioMarginAssetIndexPriceResponse struct {
	Items []QueryPortfolioMarginAssetIndexPriceResponseInner
}

// NewQueryPortfolioMarginAssetIndexPriceResponse instantiates a new QueryPortfolioMarginAssetIndexPriceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryPortfolioMarginAssetIndexPriceResponse() *QueryPortfolioMarginAssetIndexPriceResponse {
	this := QueryPortfolioMarginAssetIndexPriceResponse{}
	return &this
}

// NewQueryPortfolioMarginAssetIndexPriceResponseWithDefaults instantiates a new QueryPortfolioMarginAssetIndexPriceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryPortfolioMarginAssetIndexPriceResponseWithDefaults() *QueryPortfolioMarginAssetIndexPriceResponse {
	this := QueryPortfolioMarginAssetIndexPriceResponse{}
	return &this
}

func (o QueryPortfolioMarginAssetIndexPriceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryPortfolioMarginAssetIndexPriceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryPortfolioMarginAssetIndexPriceResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryPortfolioMarginAssetIndexPriceResponse struct {
	value QueryPortfolioMarginAssetIndexPriceResponse
	isSet bool
}

func (v NullableQueryPortfolioMarginAssetIndexPriceResponse) Get() QueryPortfolioMarginAssetIndexPriceResponse {
	return v.value
}

func (v *NullableQueryPortfolioMarginAssetIndexPriceResponse) Set(val QueryPortfolioMarginAssetIndexPriceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryPortfolioMarginAssetIndexPriceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryPortfolioMarginAssetIndexPriceResponse) Unset() {
	v.value = QueryPortfolioMarginAssetIndexPriceResponse{}
	v.isSet = false
}

func NewNullableQueryPortfolioMarginAssetIndexPriceResponse(val QueryPortfolioMarginAssetIndexPriceResponse) *NullableQueryPortfolioMarginAssetIndexPriceResponse {
	return &NullableQueryPortfolioMarginAssetIndexPriceResponse{value: val, isSet: true}
}

func (v NullableQueryPortfolioMarginAssetIndexPriceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryPortfolioMarginAssetIndexPriceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
