/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetFundingRateHistoryOfPerpetualFuturesResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFundingRateHistoryOfPerpetualFuturesResponse{}

// GetFundingRateHistoryOfPerpetualFuturesResponse struct for GetFundingRateHistoryOfPerpetualFuturesResponse
type GetFundingRateHistoryOfPerpetualFuturesResponse struct {
	Items []GetFundingRateHistoryOfPerpetualFuturesResponseInner
}

// NewGetFundingRateHistoryOfPerpetualFuturesResponse instantiates a new GetFundingRateHistoryOfPerpetualFuturesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFundingRateHistoryOfPerpetualFuturesResponse() *GetFundingRateHistoryOfPerpetualFuturesResponse {
	this := GetFundingRateHistoryOfPerpetualFuturesResponse{}
	return &this
}

// NewGetFundingRateHistoryOfPerpetualFuturesResponseWithDefaults instantiates a new GetFundingRateHistoryOfPerpetualFuturesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFundingRateHistoryOfPerpetualFuturesResponseWithDefaults() *GetFundingRateHistoryOfPerpetualFuturesResponse {
	this := GetFundingRateHistoryOfPerpetualFuturesResponse{}
	return &this
}

func (o GetFundingRateHistoryOfPerpetualFuturesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFundingRateHistoryOfPerpetualFuturesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *GetFundingRateHistoryOfPerpetualFuturesResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableGetFundingRateHistoryOfPerpetualFuturesResponse struct {
	value GetFundingRateHistoryOfPerpetualFuturesResponse
	isSet bool
}

func (v NullableGetFundingRateHistoryOfPerpetualFuturesResponse) Get() GetFundingRateHistoryOfPerpetualFuturesResponse {
	return v.value
}

func (v *NullableGetFundingRateHistoryOfPerpetualFuturesResponse) Set(val GetFundingRateHistoryOfPerpetualFuturesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFundingRateHistoryOfPerpetualFuturesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFundingRateHistoryOfPerpetualFuturesResponse) Unset() {
	v.value = GetFundingRateHistoryOfPerpetualFuturesResponse{}
	v.isSet = false
}

func NewNullableGetFundingRateHistoryOfPerpetualFuturesResponse(val GetFundingRateHistoryOfPerpetualFuturesResponse) *NullableGetFundingRateHistoryOfPerpetualFuturesResponse {
	return &NullableGetFundingRateHistoryOfPerpetualFuturesResponse{value: val, isSet: true}
}

func (v NullableGetFundingRateHistoryOfPerpetualFuturesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFundingRateHistoryOfPerpetualFuturesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
