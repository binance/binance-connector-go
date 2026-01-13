/*
Binance Spot WebSocket Streams

OpenAPI Specifications for the Binance Spot WebSocket Streams  API documents:   - [Github web-socket-streams documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-streams.md)   - [General API information for web-socket-streams on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-streams)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AllMarketRollingWindowTickerResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AllMarketRollingWindowTickerResponse{}

// AllMarketRollingWindowTickerResponse struct for AllMarketRollingWindowTickerResponse
type AllMarketRollingWindowTickerResponse struct {
	Items []AllMarketRollingWindowTickerResponseInner
}

// NewAllMarketRollingWindowTickerResponse instantiates a new AllMarketRollingWindowTickerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllMarketRollingWindowTickerResponse() *AllMarketRollingWindowTickerResponse {
	this := AllMarketRollingWindowTickerResponse{}
	return &this
}

// NewAllMarketRollingWindowTickerResponseWithDefaults instantiates a new AllMarketRollingWindowTickerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllMarketRollingWindowTickerResponseWithDefaults() *AllMarketRollingWindowTickerResponse {
	this := AllMarketRollingWindowTickerResponse{}
	return &this
}

func (o AllMarketRollingWindowTickerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllMarketRollingWindowTickerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *AllMarketRollingWindowTickerResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableAllMarketRollingWindowTickerResponse struct {
	value AllMarketRollingWindowTickerResponse
	isSet bool
}

func (v NullableAllMarketRollingWindowTickerResponse) Get() AllMarketRollingWindowTickerResponse {
	return v.value
}

func (v *NullableAllMarketRollingWindowTickerResponse) Set(val AllMarketRollingWindowTickerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAllMarketRollingWindowTickerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAllMarketRollingWindowTickerResponse) Unset() {
	v.value = AllMarketRollingWindowTickerResponse{}
	v.isSet = false
}

func NewNullableAllMarketRollingWindowTickerResponse(val AllMarketRollingWindowTickerResponse) *NullableAllMarketRollingWindowTickerResponse {
	return &NullableAllMarketRollingWindowTickerResponse{value: val, isSet: true}
}

func (v NullableAllMarketRollingWindowTickerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllMarketRollingWindowTickerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
