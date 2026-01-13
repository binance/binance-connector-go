/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the MarginAccountBorrowResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MarginAccountBorrowResponse{}

// MarginAccountBorrowResponse struct for MarginAccountBorrowResponse
type MarginAccountBorrowResponse struct {
	TranId               *int64 `json:"tranId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarginAccountBorrowResponse MarginAccountBorrowResponse

// NewMarginAccountBorrowResponse instantiates a new MarginAccountBorrowResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginAccountBorrowResponse() *MarginAccountBorrowResponse {
	this := MarginAccountBorrowResponse{}
	return &this
}

// NewMarginAccountBorrowResponseWithDefaults instantiates a new MarginAccountBorrowResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginAccountBorrowResponseWithDefaults() *MarginAccountBorrowResponse {
	this := MarginAccountBorrowResponse{}
	return &this
}

// GetTranId returns the TranId field value if set, zero value otherwise.
func (o *MarginAccountBorrowResponse) GetTranId() int64 {
	if o == nil || common.IsNil(o.TranId) {
		var ret int64
		return ret
	}
	return *o.TranId
}

// GetTranIdOk returns a tuple with the TranId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountBorrowResponse) GetTranIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TranId) {
		return nil, false
	}
	return o.TranId, true
}

// HasTranId returns a boolean if a field has been set.
func (o *MarginAccountBorrowResponse) HasTranId() bool {
	if o != nil && !common.IsNil(o.TranId) {
		return true
	}

	return false
}

// SetTranId gets a reference to the given int64 and assigns it to the TranId field.
func (o *MarginAccountBorrowResponse) SetTranId(v int64) {
	o.TranId = &v
}

func (o MarginAccountBorrowResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarginAccountBorrowResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TranId) {
		toSerialize["tranId"] = o.TranId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MarginAccountBorrowResponse) UnmarshalJSON(data []byte) (err error) {
	varMarginAccountBorrowResponse := _MarginAccountBorrowResponse{}

	err = json.Unmarshal(data, &varMarginAccountBorrowResponse)

	if err != nil {
		return err
	}

	*o = MarginAccountBorrowResponse(varMarginAccountBorrowResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tranId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarginAccountBorrowResponse struct {
	value *MarginAccountBorrowResponse
	isSet bool
}

func (v NullableMarginAccountBorrowResponse) Get() *MarginAccountBorrowResponse {
	return v.value
}

func (v *NullableMarginAccountBorrowResponse) Set(val *MarginAccountBorrowResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginAccountBorrowResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginAccountBorrowResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginAccountBorrowResponse(val *MarginAccountBorrowResponse) *NullableMarginAccountBorrowResponse {
	return &NullableMarginAccountBorrowResponse{value: val, isSet: true}
}

func (v NullableMarginAccountBorrowResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginAccountBorrowResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
