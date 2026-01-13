/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the HistoricalTradesResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &HistoricalTradesResponse{}

// HistoricalTradesResponse struct for HistoricalTradesResponse
type HistoricalTradesResponse struct {
	Items []HistoricalTradesResponseInner
}

// NewHistoricalTradesResponse instantiates a new HistoricalTradesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoricalTradesResponse() *HistoricalTradesResponse {
	this := HistoricalTradesResponse{}
	return &this
}

// NewHistoricalTradesResponseWithDefaults instantiates a new HistoricalTradesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoricalTradesResponseWithDefaults() *HistoricalTradesResponse {
	this := HistoricalTradesResponse{}
	return &this
}

func (o HistoricalTradesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoricalTradesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *HistoricalTradesResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableHistoricalTradesResponse struct {
	value HistoricalTradesResponse
	isSet bool
}

func (v NullableHistoricalTradesResponse) Get() HistoricalTradesResponse {
	return v.value
}

func (v *NullableHistoricalTradesResponse) Set(val HistoricalTradesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoricalTradesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoricalTradesResponse) Unset() {
	v.value = HistoricalTradesResponse{}
	v.isSet = false
}

func NewNullableHistoricalTradesResponse(val HistoricalTradesResponse) *NullableHistoricalTradesResponse {
	return &NullableHistoricalTradesResponse{value: val, isSet: true}
}

func (v NullableHistoricalTradesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoricalTradesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
