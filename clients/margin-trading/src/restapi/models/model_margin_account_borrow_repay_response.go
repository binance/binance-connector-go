/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the MarginAccountBorrowRepayResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MarginAccountBorrowRepayResponse{}

// MarginAccountBorrowRepayResponse struct for MarginAccountBorrowRepayResponse
type MarginAccountBorrowRepayResponse struct {
	TranId               *int64 `json:"tranId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarginAccountBorrowRepayResponse MarginAccountBorrowRepayResponse

// NewMarginAccountBorrowRepayResponse instantiates a new MarginAccountBorrowRepayResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginAccountBorrowRepayResponse() *MarginAccountBorrowRepayResponse {
	this := MarginAccountBorrowRepayResponse{}
	return &this
}

// NewMarginAccountBorrowRepayResponseWithDefaults instantiates a new MarginAccountBorrowRepayResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginAccountBorrowRepayResponseWithDefaults() *MarginAccountBorrowRepayResponse {
	this := MarginAccountBorrowRepayResponse{}
	return &this
}

// GetTranId returns the TranId field value if set, zero value otherwise.
func (o *MarginAccountBorrowRepayResponse) GetTranId() int64 {
	if o == nil || common.IsNil(o.TranId) {
		var ret int64
		return ret
	}
	return *o.TranId
}

// GetTranIdOk returns a tuple with the TranId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountBorrowRepayResponse) GetTranIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TranId) {
		return nil, false
	}
	return o.TranId, true
}

// HasTranId returns a boolean if a field has been set.
func (o *MarginAccountBorrowRepayResponse) HasTranId() bool {
	if o != nil && !common.IsNil(o.TranId) {
		return true
	}

	return false
}

// SetTranId gets a reference to the given int64 and assigns it to the TranId field.
func (o *MarginAccountBorrowRepayResponse) SetTranId(v int64) {
	o.TranId = &v
}

func (o MarginAccountBorrowRepayResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarginAccountBorrowRepayResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TranId) {
		toSerialize["tranId"] = o.TranId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MarginAccountBorrowRepayResponse) UnmarshalJSON(data []byte) (err error) {
	varMarginAccountBorrowRepayResponse := _MarginAccountBorrowRepayResponse{}

	err = json.Unmarshal(data, &varMarginAccountBorrowRepayResponse)

	if err != nil {
		return err
	}

	*o = MarginAccountBorrowRepayResponse(varMarginAccountBorrowRepayResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tranId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarginAccountBorrowRepayResponse struct {
	value *MarginAccountBorrowRepayResponse
	isSet bool
}

func (v NullableMarginAccountBorrowRepayResponse) Get() *MarginAccountBorrowRepayResponse {
	return v.value
}

func (v *NullableMarginAccountBorrowRepayResponse) Set(val *MarginAccountBorrowRepayResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginAccountBorrowRepayResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginAccountBorrowRepayResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginAccountBorrowRepayResponse(val *MarginAccountBorrowRepayResponse) *NullableMarginAccountBorrowRepayResponse {
	return &NullableMarginAccountBorrowRepayResponse{value: val, isSet: true}
}

func (v NullableMarginAccountBorrowRepayResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginAccountBorrowRepayResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
