/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the MarginAccountRepayResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MarginAccountRepayResponse{}

// MarginAccountRepayResponse struct for MarginAccountRepayResponse
type MarginAccountRepayResponse struct {
	TranId               *int64 `json:"tranId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarginAccountRepayResponse MarginAccountRepayResponse

// NewMarginAccountRepayResponse instantiates a new MarginAccountRepayResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginAccountRepayResponse() *MarginAccountRepayResponse {
	this := MarginAccountRepayResponse{}
	return &this
}

// NewMarginAccountRepayResponseWithDefaults instantiates a new MarginAccountRepayResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginAccountRepayResponseWithDefaults() *MarginAccountRepayResponse {
	this := MarginAccountRepayResponse{}
	return &this
}

// GetTranId returns the TranId field value if set, zero value otherwise.
func (o *MarginAccountRepayResponse) GetTranId() int64 {
	if o == nil || common.IsNil(o.TranId) {
		var ret int64
		return ret
	}
	return *o.TranId
}

// GetTranIdOk returns a tuple with the TranId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountRepayResponse) GetTranIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TranId) {
		return nil, false
	}
	return o.TranId, true
}

// HasTranId returns a boolean if a field has been set.
func (o *MarginAccountRepayResponse) HasTranId() bool {
	if o != nil && !common.IsNil(o.TranId) {
		return true
	}

	return false
}

// SetTranId gets a reference to the given int64 and assigns it to the TranId field.
func (o *MarginAccountRepayResponse) SetTranId(v int64) {
	o.TranId = &v
}

func (o MarginAccountRepayResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarginAccountRepayResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TranId) {
		toSerialize["tranId"] = o.TranId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MarginAccountRepayResponse) UnmarshalJSON(data []byte) (err error) {
	varMarginAccountRepayResponse := _MarginAccountRepayResponse{}

	err = json.Unmarshal(data, &varMarginAccountRepayResponse)

	if err != nil {
		return err
	}

	*o = MarginAccountRepayResponse(varMarginAccountRepayResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tranId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarginAccountRepayResponse struct {
	value *MarginAccountRepayResponse
	isSet bool
}

func (v NullableMarginAccountRepayResponse) Get() *MarginAccountRepayResponse {
	return v.value
}

func (v *NullableMarginAccountRepayResponse) Set(val *MarginAccountRepayResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginAccountRepayResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginAccountRepayResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginAccountRepayResponse(val *MarginAccountRepayResponse) *NullableMarginAccountRepayResponse {
	return &NullableMarginAccountRepayResponse{value: val, isSet: true}
}

func (v NullableMarginAccountRepayResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginAccountRepayResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
