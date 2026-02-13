/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the CmAccountTradeListResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CmAccountTradeListResponse{}

// CmAccountTradeListResponse struct for CmAccountTradeListResponse
type CmAccountTradeListResponse struct {
	Items []CmAccountTradeListResponseInner
}

// NewCmAccountTradeListResponse instantiates a new CmAccountTradeListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCmAccountTradeListResponse() *CmAccountTradeListResponse {
	this := CmAccountTradeListResponse{}
	return &this
}

// NewCmAccountTradeListResponseWithDefaults instantiates a new CmAccountTradeListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCmAccountTradeListResponseWithDefaults() *CmAccountTradeListResponse {
	this := CmAccountTradeListResponse{}
	return &this
}

func (o CmAccountTradeListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CmAccountTradeListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *CmAccountTradeListResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableCmAccountTradeListResponse struct {
	value CmAccountTradeListResponse
	isSet bool
}

func (v NullableCmAccountTradeListResponse) Get() CmAccountTradeListResponse {
	return v.value
}

func (v *NullableCmAccountTradeListResponse) Set(val CmAccountTradeListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCmAccountTradeListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCmAccountTradeListResponse) Unset() {
	v.value = CmAccountTradeListResponse{}
	v.isSet = false
}

func NewNullableCmAccountTradeListResponse(val CmAccountTradeListResponse) *NullableCmAccountTradeListResponse {
	return &NullableCmAccountTradeListResponse{value: val, isSet: true}
}

func (v NullableCmAccountTradeListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCmAccountTradeListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
