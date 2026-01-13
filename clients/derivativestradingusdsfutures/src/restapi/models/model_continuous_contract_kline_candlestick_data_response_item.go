/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the ContinuousContractKlineCandlestickDataResponseItem type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ContinuousContractKlineCandlestickDataResponseItem{}

// ContinuousContractKlineCandlestickDataResponseItem struct for ContinuousContractKlineCandlestickDataResponseItem
type ContinuousContractKlineCandlestickDataResponseItem struct {
	Items []ContinuousContractKlineCandlestickDataResponseItemInner
}

// NewContinuousContractKlineCandlestickDataResponseItem instantiates a new ContinuousContractKlineCandlestickDataResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContinuousContractKlineCandlestickDataResponseItem() *ContinuousContractKlineCandlestickDataResponseItem {
	this := ContinuousContractKlineCandlestickDataResponseItem{}
	return &this
}

// NewContinuousContractKlineCandlestickDataResponseItemWithDefaults instantiates a new ContinuousContractKlineCandlestickDataResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContinuousContractKlineCandlestickDataResponseItemWithDefaults() *ContinuousContractKlineCandlestickDataResponseItem {
	this := ContinuousContractKlineCandlestickDataResponseItem{}
	return &this
}

func (o ContinuousContractKlineCandlestickDataResponseItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContinuousContractKlineCandlestickDataResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *ContinuousContractKlineCandlestickDataResponseItem) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableContinuousContractKlineCandlestickDataResponseItem struct {
	value ContinuousContractKlineCandlestickDataResponseItem
	isSet bool
}

func (v NullableContinuousContractKlineCandlestickDataResponseItem) Get() ContinuousContractKlineCandlestickDataResponseItem {
	return v.value
}

func (v *NullableContinuousContractKlineCandlestickDataResponseItem) Set(val ContinuousContractKlineCandlestickDataResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableContinuousContractKlineCandlestickDataResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableContinuousContractKlineCandlestickDataResponseItem) Unset() {
	v.value = ContinuousContractKlineCandlestickDataResponseItem{}
	v.isSet = false
}

func NewNullableContinuousContractKlineCandlestickDataResponseItem(val ContinuousContractKlineCandlestickDataResponseItem) *NullableContinuousContractKlineCandlestickDataResponseItem {
	return &NullableContinuousContractKlineCandlestickDataResponseItem{value: val, isSet: true}
}

func (v NullableContinuousContractKlineCandlestickDataResponseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContinuousContractKlineCandlestickDataResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
