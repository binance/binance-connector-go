/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the UmFuturesSymbolConfigurationResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UmFuturesSymbolConfigurationResponse{}

// UmFuturesSymbolConfigurationResponse struct for UmFuturesSymbolConfigurationResponse
type UmFuturesSymbolConfigurationResponse struct {
	Items []UmFuturesSymbolConfigurationResponseInner
}

// NewUmFuturesSymbolConfigurationResponse instantiates a new UmFuturesSymbolConfigurationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUmFuturesSymbolConfigurationResponse() *UmFuturesSymbolConfigurationResponse {
	this := UmFuturesSymbolConfigurationResponse{}
	return &this
}

// NewUmFuturesSymbolConfigurationResponseWithDefaults instantiates a new UmFuturesSymbolConfigurationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUmFuturesSymbolConfigurationResponseWithDefaults() *UmFuturesSymbolConfigurationResponse {
	this := UmFuturesSymbolConfigurationResponse{}
	return &this
}

func (o UmFuturesSymbolConfigurationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UmFuturesSymbolConfigurationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *UmFuturesSymbolConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableUmFuturesSymbolConfigurationResponse struct {
	value UmFuturesSymbolConfigurationResponse
	isSet bool
}

func (v NullableUmFuturesSymbolConfigurationResponse) Get() UmFuturesSymbolConfigurationResponse {
	return v.value
}

func (v *NullableUmFuturesSymbolConfigurationResponse) Set(val UmFuturesSymbolConfigurationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUmFuturesSymbolConfigurationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUmFuturesSymbolConfigurationResponse) Unset() {
	v.value = UmFuturesSymbolConfigurationResponse{}
	v.isSet = false
}

func NewNullableUmFuturesSymbolConfigurationResponse(val UmFuturesSymbolConfigurationResponse) *NullableUmFuturesSymbolConfigurationResponse {
	return &NullableUmFuturesSymbolConfigurationResponse{value: val, isSet: true}
}

func (v NullableUmFuturesSymbolConfigurationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUmFuturesSymbolConfigurationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
