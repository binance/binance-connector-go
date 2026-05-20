/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the HistoricalBlockTradesResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &HistoricalBlockTradesResponse{}

// HistoricalBlockTradesResponse struct for HistoricalBlockTradesResponse
type HistoricalBlockTradesResponse struct {
	Items []HistoricalBlockTradesResponseInner
}

// NewHistoricalBlockTradesResponse instantiates a new HistoricalBlockTradesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoricalBlockTradesResponse() *HistoricalBlockTradesResponse {
	this := HistoricalBlockTradesResponse{}
	return &this
}

// NewHistoricalBlockTradesResponseWithDefaults instantiates a new HistoricalBlockTradesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoricalBlockTradesResponseWithDefaults() *HistoricalBlockTradesResponse {
	this := HistoricalBlockTradesResponse{}
	return &this
}

func (o HistoricalBlockTradesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoricalBlockTradesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *HistoricalBlockTradesResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableHistoricalBlockTradesResponse struct {
	value HistoricalBlockTradesResponse
	isSet bool
}

func (v NullableHistoricalBlockTradesResponse) Get() HistoricalBlockTradesResponse {
	return v.value
}

func (v *NullableHistoricalBlockTradesResponse) Set(val HistoricalBlockTradesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoricalBlockTradesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoricalBlockTradesResponse) Unset() {
	v.value = HistoricalBlockTradesResponse{}
	v.isSet = false
}

func NewNullableHistoricalBlockTradesResponse(val HistoricalBlockTradesResponse) *NullableHistoricalBlockTradesResponse {
	return &NullableHistoricalBlockTradesResponse{value: val, isSet: true}
}

func (v NullableHistoricalBlockTradesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoricalBlockTradesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
