/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the ContinuousContractKlineCandlestickDataResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ContinuousContractKlineCandlestickDataResponse{}

// ContinuousContractKlineCandlestickDataResponse struct for ContinuousContractKlineCandlestickDataResponse
type ContinuousContractKlineCandlestickDataResponse struct {
	Items []ContinuousContractKlineCandlestickDataResponseItem
}

// NewContinuousContractKlineCandlestickDataResponse instantiates a new ContinuousContractKlineCandlestickDataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContinuousContractKlineCandlestickDataResponse() *ContinuousContractKlineCandlestickDataResponse {
	this := ContinuousContractKlineCandlestickDataResponse{}
	return &this
}

// NewContinuousContractKlineCandlestickDataResponseWithDefaults instantiates a new ContinuousContractKlineCandlestickDataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContinuousContractKlineCandlestickDataResponseWithDefaults() *ContinuousContractKlineCandlestickDataResponse {
	this := ContinuousContractKlineCandlestickDataResponse{}
	return &this
}

func (o ContinuousContractKlineCandlestickDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContinuousContractKlineCandlestickDataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *ContinuousContractKlineCandlestickDataResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableContinuousContractKlineCandlestickDataResponse struct {
	value ContinuousContractKlineCandlestickDataResponse
	isSet bool
}

func (v NullableContinuousContractKlineCandlestickDataResponse) Get() ContinuousContractKlineCandlestickDataResponse {
	return v.value
}

func (v *NullableContinuousContractKlineCandlestickDataResponse) Set(val ContinuousContractKlineCandlestickDataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContinuousContractKlineCandlestickDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContinuousContractKlineCandlestickDataResponse) Unset() {
	v.value = ContinuousContractKlineCandlestickDataResponse{}
	v.isSet = false
}

func NewNullableContinuousContractKlineCandlestickDataResponse(val ContinuousContractKlineCandlestickDataResponse) *NullableContinuousContractKlineCandlestickDataResponse {
	return &NullableContinuousContractKlineCandlestickDataResponse{value: val, isSet: true}
}

func (v NullableContinuousContractKlineCandlestickDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContinuousContractKlineCandlestickDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
