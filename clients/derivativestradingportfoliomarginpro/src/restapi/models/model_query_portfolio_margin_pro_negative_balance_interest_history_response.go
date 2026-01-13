/*
Binance Derivatives Trading Portfolio Margin Pro REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin Pro REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryPortfolioMarginProNegativeBalanceInterestHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryPortfolioMarginProNegativeBalanceInterestHistoryResponse{}

// QueryPortfolioMarginProNegativeBalanceInterestHistoryResponse struct for QueryPortfolioMarginProNegativeBalanceInterestHistoryResponse
type QueryPortfolioMarginProNegativeBalanceInterestHistoryResponse struct {
	Items []QueryPortfolioMarginProNegativeBalanceInterestHistoryResponseInner
}

// NewQueryPortfolioMarginProNegativeBalanceInterestHistoryResponse instantiates a new QueryPortfolioMarginProNegativeBalanceInterestHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryPortfolioMarginProNegativeBalanceInterestHistoryResponse() *QueryPortfolioMarginProNegativeBalanceInterestHistoryResponse {
	this := QueryPortfolioMarginProNegativeBalanceInterestHistoryResponse{}
	return &this
}

// NewQueryPortfolioMarginProNegativeBalanceInterestHistoryResponseWithDefaults instantiates a new QueryPortfolioMarginProNegativeBalanceInterestHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryPortfolioMarginProNegativeBalanceInterestHistoryResponseWithDefaults() *QueryPortfolioMarginProNegativeBalanceInterestHistoryResponse {
	this := QueryPortfolioMarginProNegativeBalanceInterestHistoryResponse{}
	return &this
}

func (o QueryPortfolioMarginProNegativeBalanceInterestHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryPortfolioMarginProNegativeBalanceInterestHistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryPortfolioMarginProNegativeBalanceInterestHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryPortfolioMarginProNegativeBalanceInterestHistoryResponse struct {
	value QueryPortfolioMarginProNegativeBalanceInterestHistoryResponse
	isSet bool
}

func (v NullableQueryPortfolioMarginProNegativeBalanceInterestHistoryResponse) Get() QueryPortfolioMarginProNegativeBalanceInterestHistoryResponse {
	return v.value
}

func (v *NullableQueryPortfolioMarginProNegativeBalanceInterestHistoryResponse) Set(val QueryPortfolioMarginProNegativeBalanceInterestHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryPortfolioMarginProNegativeBalanceInterestHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryPortfolioMarginProNegativeBalanceInterestHistoryResponse) Unset() {
	v.value = QueryPortfolioMarginProNegativeBalanceInterestHistoryResponse{}
	v.isSet = false
}

func NewNullableQueryPortfolioMarginProNegativeBalanceInterestHistoryResponse(val QueryPortfolioMarginProNegativeBalanceInterestHistoryResponse) *NullableQueryPortfolioMarginProNegativeBalanceInterestHistoryResponse {
	return &NullableQueryPortfolioMarginProNegativeBalanceInterestHistoryResponse{value: val, isSet: true}
}

func (v NullableQueryPortfolioMarginProNegativeBalanceInterestHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryPortfolioMarginProNegativeBalanceInterestHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
