/*
Binance VIP Loan REST API

OpenAPI Specification for the Binance VIP Loan REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetBorrowInterestRateResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetBorrowInterestRateResponse{}

// GetBorrowInterestRateResponse struct for GetBorrowInterestRateResponse
type GetBorrowInterestRateResponse struct {
	Items []GetBorrowInterestRateResponseInner
}

// NewGetBorrowInterestRateResponse instantiates a new GetBorrowInterestRateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBorrowInterestRateResponse() *GetBorrowInterestRateResponse {
	this := GetBorrowInterestRateResponse{}
	return &this
}

// NewGetBorrowInterestRateResponseWithDefaults instantiates a new GetBorrowInterestRateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBorrowInterestRateResponseWithDefaults() *GetBorrowInterestRateResponse {
	this := GetBorrowInterestRateResponse{}
	return &this
}

func (o GetBorrowInterestRateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBorrowInterestRateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *GetBorrowInterestRateResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableGetBorrowInterestRateResponse struct {
	value GetBorrowInterestRateResponse
	isSet bool
}

func (v NullableGetBorrowInterestRateResponse) Get() GetBorrowInterestRateResponse {
	return v.value
}

func (v *NullableGetBorrowInterestRateResponse) Set(val GetBorrowInterestRateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBorrowInterestRateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBorrowInterestRateResponse) Unset() {
	v.value = GetBorrowInterestRateResponse{}
	v.isSet = false
}

func NewNullableGetBorrowInterestRateResponse(val GetBorrowInterestRateResponse) *NullableGetBorrowInterestRateResponse {
	return &NullableGetBorrowInterestRateResponse{value: val, isSet: true}
}

func (v NullableGetBorrowInterestRateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBorrowInterestRateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
