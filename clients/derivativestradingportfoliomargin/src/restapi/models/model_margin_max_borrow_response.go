/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the MarginMaxBorrowResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MarginMaxBorrowResponse{}

// MarginMaxBorrowResponse struct for MarginMaxBorrowResponse
type MarginMaxBorrowResponse struct {
	Amount               *string `json:"amount,omitempty"`
	BorrowLimit          *string `json:"borrowLimit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarginMaxBorrowResponse MarginMaxBorrowResponse

// NewMarginMaxBorrowResponse instantiates a new MarginMaxBorrowResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginMaxBorrowResponse() *MarginMaxBorrowResponse {
	this := MarginMaxBorrowResponse{}
	return &this
}

// NewMarginMaxBorrowResponseWithDefaults instantiates a new MarginMaxBorrowResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginMaxBorrowResponseWithDefaults() *MarginMaxBorrowResponse {
	this := MarginMaxBorrowResponse{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *MarginMaxBorrowResponse) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginMaxBorrowResponse) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *MarginMaxBorrowResponse) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *MarginMaxBorrowResponse) SetAmount(v string) {
	o.Amount = &v
}

// GetBorrowLimit returns the BorrowLimit field value if set, zero value otherwise.
func (o *MarginMaxBorrowResponse) GetBorrowLimit() string {
	if o == nil || common.IsNil(o.BorrowLimit) {
		var ret string
		return ret
	}
	return *o.BorrowLimit
}

// GetBorrowLimitOk returns a tuple with the BorrowLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginMaxBorrowResponse) GetBorrowLimitOk() (*string, bool) {
	if o == nil || common.IsNil(o.BorrowLimit) {
		return nil, false
	}
	return o.BorrowLimit, true
}

// HasBorrowLimit returns a boolean if a field has been set.
func (o *MarginMaxBorrowResponse) HasBorrowLimit() bool {
	if o != nil && !common.IsNil(o.BorrowLimit) {
		return true
	}

	return false
}

// SetBorrowLimit gets a reference to the given string and assigns it to the BorrowLimit field.
func (o *MarginMaxBorrowResponse) SetBorrowLimit(v string) {
	o.BorrowLimit = &v
}

func (o MarginMaxBorrowResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarginMaxBorrowResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.BorrowLimit) {
		toSerialize["borrowLimit"] = o.BorrowLimit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MarginMaxBorrowResponse) UnmarshalJSON(data []byte) (err error) {
	varMarginMaxBorrowResponse := _MarginMaxBorrowResponse{}

	err = json.Unmarshal(data, &varMarginMaxBorrowResponse)

	if err != nil {
		return err
	}

	*o = MarginMaxBorrowResponse(varMarginMaxBorrowResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "borrowLimit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarginMaxBorrowResponse struct {
	value *MarginMaxBorrowResponse
	isSet bool
}

func (v NullableMarginMaxBorrowResponse) Get() *MarginMaxBorrowResponse {
	return v.value
}

func (v *NullableMarginMaxBorrowResponse) Set(val *MarginMaxBorrowResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginMaxBorrowResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginMaxBorrowResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginMaxBorrowResponse(val *MarginMaxBorrowResponse) *NullableMarginMaxBorrowResponse {
	return &NullableMarginMaxBorrowResponse{value: val, isSet: true}
}

func (v NullableMarginMaxBorrowResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginMaxBorrowResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
