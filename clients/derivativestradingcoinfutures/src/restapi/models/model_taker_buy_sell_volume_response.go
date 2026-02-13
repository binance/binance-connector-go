/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the TakerBuySellVolumeResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TakerBuySellVolumeResponse{}

// TakerBuySellVolumeResponse struct for TakerBuySellVolumeResponse
type TakerBuySellVolumeResponse struct {
	Items []TakerBuySellVolumeResponseInner
}

// NewTakerBuySellVolumeResponse instantiates a new TakerBuySellVolumeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTakerBuySellVolumeResponse() *TakerBuySellVolumeResponse {
	this := TakerBuySellVolumeResponse{}
	return &this
}

// NewTakerBuySellVolumeResponseWithDefaults instantiates a new TakerBuySellVolumeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTakerBuySellVolumeResponseWithDefaults() *TakerBuySellVolumeResponse {
	this := TakerBuySellVolumeResponse{}
	return &this
}

func (o TakerBuySellVolumeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TakerBuySellVolumeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *TakerBuySellVolumeResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableTakerBuySellVolumeResponse struct {
	value TakerBuySellVolumeResponse
	isSet bool
}

func (v NullableTakerBuySellVolumeResponse) Get() TakerBuySellVolumeResponse {
	return v.value
}

func (v *NullableTakerBuySellVolumeResponse) Set(val TakerBuySellVolumeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTakerBuySellVolumeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTakerBuySellVolumeResponse) Unset() {
	v.value = TakerBuySellVolumeResponse{}
	v.isSet = false
}

func NewNullableTakerBuySellVolumeResponse(val TakerBuySellVolumeResponse) *NullableTakerBuySellVolumeResponse {
	return &NullableTakerBuySellVolumeResponse{value: val, isSet: true}
}

func (v NullableTakerBuySellVolumeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTakerBuySellVolumeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
