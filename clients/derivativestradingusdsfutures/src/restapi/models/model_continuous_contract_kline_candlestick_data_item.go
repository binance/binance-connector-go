/*
Futures (USDⓈ-M) REST API

Access market data, manage accounts, and trade USDⓈ-M perpetual futures.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ContinuousContractKlineCandlestickDataItem type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ContinuousContractKlineCandlestickDataItem{}

// ContinuousContractKlineCandlestickDataItem struct for ContinuousContractKlineCandlestickDataItem
type ContinuousContractKlineCandlestickDataItem struct {
	Items []ContinuousContractKlineCandlestickDataItemInner
}

// NewContinuousContractKlineCandlestickDataItem instantiates a new ContinuousContractKlineCandlestickDataItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContinuousContractKlineCandlestickDataItem() *ContinuousContractKlineCandlestickDataItem {
	this := ContinuousContractKlineCandlestickDataItem{}
	return &this
}

// NewContinuousContractKlineCandlestickDataItemWithDefaults instantiates a new ContinuousContractKlineCandlestickDataItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContinuousContractKlineCandlestickDataItemWithDefaults() *ContinuousContractKlineCandlestickDataItem {
	this := ContinuousContractKlineCandlestickDataItem{}
	return &this
}

func (o ContinuousContractKlineCandlestickDataItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContinuousContractKlineCandlestickDataItem) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *ContinuousContractKlineCandlestickDataItem) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableContinuousContractKlineCandlestickDataItem struct {
	value ContinuousContractKlineCandlestickDataItem
	isSet bool
}

func (v NullableContinuousContractKlineCandlestickDataItem) Get() ContinuousContractKlineCandlestickDataItem {
	return v.value
}

func (v *NullableContinuousContractKlineCandlestickDataItem) Set(val ContinuousContractKlineCandlestickDataItem) {
	v.value = val
	v.isSet = true
}

func (v NullableContinuousContractKlineCandlestickDataItem) IsSet() bool {
	return v.isSet
}

func (v *NullableContinuousContractKlineCandlestickDataItem) Unset() {
	v.value = ContinuousContractKlineCandlestickDataItem{}
	v.isSet = false
}

func NewNullableContinuousContractKlineCandlestickDataItem(val ContinuousContractKlineCandlestickDataItem) *NullableContinuousContractKlineCandlestickDataItem {
	return &NullableContinuousContractKlineCandlestickDataItem{value: val, isSet: true}
}

func (v NullableContinuousContractKlineCandlestickDataItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContinuousContractKlineCandlestickDataItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
