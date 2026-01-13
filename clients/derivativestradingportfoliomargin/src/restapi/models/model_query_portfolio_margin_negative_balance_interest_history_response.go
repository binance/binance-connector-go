/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryPortfolioMarginNegativeBalanceInterestHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryPortfolioMarginNegativeBalanceInterestHistoryResponse{}

// QueryPortfolioMarginNegativeBalanceInterestHistoryResponse struct for QueryPortfolioMarginNegativeBalanceInterestHistoryResponse
type QueryPortfolioMarginNegativeBalanceInterestHistoryResponse struct {
	Items []QueryPortfolioMarginNegativeBalanceInterestHistoryResponseInner
}

// NewQueryPortfolioMarginNegativeBalanceInterestHistoryResponse instantiates a new QueryPortfolioMarginNegativeBalanceInterestHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryPortfolioMarginNegativeBalanceInterestHistoryResponse() *QueryPortfolioMarginNegativeBalanceInterestHistoryResponse {
	this := QueryPortfolioMarginNegativeBalanceInterestHistoryResponse{}
	return &this
}

// NewQueryPortfolioMarginNegativeBalanceInterestHistoryResponseWithDefaults instantiates a new QueryPortfolioMarginNegativeBalanceInterestHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryPortfolioMarginNegativeBalanceInterestHistoryResponseWithDefaults() *QueryPortfolioMarginNegativeBalanceInterestHistoryResponse {
	this := QueryPortfolioMarginNegativeBalanceInterestHistoryResponse{}
	return &this
}

func (o QueryPortfolioMarginNegativeBalanceInterestHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryPortfolioMarginNegativeBalanceInterestHistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryPortfolioMarginNegativeBalanceInterestHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryPortfolioMarginNegativeBalanceInterestHistoryResponse struct {
	value QueryPortfolioMarginNegativeBalanceInterestHistoryResponse
	isSet bool
}

func (v NullableQueryPortfolioMarginNegativeBalanceInterestHistoryResponse) Get() QueryPortfolioMarginNegativeBalanceInterestHistoryResponse {
	return v.value
}

func (v *NullableQueryPortfolioMarginNegativeBalanceInterestHistoryResponse) Set(val QueryPortfolioMarginNegativeBalanceInterestHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryPortfolioMarginNegativeBalanceInterestHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryPortfolioMarginNegativeBalanceInterestHistoryResponse) Unset() {
	v.value = QueryPortfolioMarginNegativeBalanceInterestHistoryResponse{}
	v.isSet = false
}

func NewNullableQueryPortfolioMarginNegativeBalanceInterestHistoryResponse(val QueryPortfolioMarginNegativeBalanceInterestHistoryResponse) *NullableQueryPortfolioMarginNegativeBalanceInterestHistoryResponse {
	return &NullableQueryPortfolioMarginNegativeBalanceInterestHistoryResponse{value: val, isSet: true}
}

func (v NullableQueryPortfolioMarginNegativeBalanceInterestHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryPortfolioMarginNegativeBalanceInterestHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
