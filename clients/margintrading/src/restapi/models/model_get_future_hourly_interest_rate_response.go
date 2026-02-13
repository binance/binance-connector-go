/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetFutureHourlyInterestRateResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFutureHourlyInterestRateResponse{}

// GetFutureHourlyInterestRateResponse struct for GetFutureHourlyInterestRateResponse
type GetFutureHourlyInterestRateResponse struct {
	Items []GetFutureHourlyInterestRateResponseInner
}

// NewGetFutureHourlyInterestRateResponse instantiates a new GetFutureHourlyInterestRateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFutureHourlyInterestRateResponse() *GetFutureHourlyInterestRateResponse {
	this := GetFutureHourlyInterestRateResponse{}
	return &this
}

// NewGetFutureHourlyInterestRateResponseWithDefaults instantiates a new GetFutureHourlyInterestRateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFutureHourlyInterestRateResponseWithDefaults() *GetFutureHourlyInterestRateResponse {
	this := GetFutureHourlyInterestRateResponse{}
	return &this
}

func (o GetFutureHourlyInterestRateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFutureHourlyInterestRateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *GetFutureHourlyInterestRateResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableGetFutureHourlyInterestRateResponse struct {
	value GetFutureHourlyInterestRateResponse
	isSet bool
}

func (v NullableGetFutureHourlyInterestRateResponse) Get() GetFutureHourlyInterestRateResponse {
	return v.value
}

func (v *NullableGetFutureHourlyInterestRateResponse) Set(val GetFutureHourlyInterestRateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFutureHourlyInterestRateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFutureHourlyInterestRateResponse) Unset() {
	v.value = GetFutureHourlyInterestRateResponse{}
	v.isSet = false
}

func NewNullableGetFutureHourlyInterestRateResponse(val GetFutureHourlyInterestRateResponse) *NullableGetFutureHourlyInterestRateResponse {
	return &NullableGetFutureHourlyInterestRateResponse{value: val, isSet: true}
}

func (v NullableGetFutureHourlyInterestRateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFutureHourlyInterestRateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
