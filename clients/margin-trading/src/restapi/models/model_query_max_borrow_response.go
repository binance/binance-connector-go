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

// checks if the QueryMaxBorrowResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryMaxBorrowResponse{}

// QueryMaxBorrowResponse struct for QueryMaxBorrowResponse
type QueryMaxBorrowResponse struct {
	Amount               *string `json:"amount,omitempty"`
	BorrowLimit          *string `json:"borrowLimit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryMaxBorrowResponse QueryMaxBorrowResponse

// NewQueryMaxBorrowResponse instantiates a new QueryMaxBorrowResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryMaxBorrowResponse() *QueryMaxBorrowResponse {
	this := QueryMaxBorrowResponse{}
	return &this
}

// NewQueryMaxBorrowResponseWithDefaults instantiates a new QueryMaxBorrowResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryMaxBorrowResponseWithDefaults() *QueryMaxBorrowResponse {
	this := QueryMaxBorrowResponse{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *QueryMaxBorrowResponse) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMaxBorrowResponse) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *QueryMaxBorrowResponse) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *QueryMaxBorrowResponse) SetAmount(v string) {
	o.Amount = &v
}

// GetBorrowLimit returns the BorrowLimit field value if set, zero value otherwise.
func (o *QueryMaxBorrowResponse) GetBorrowLimit() string {
	if o == nil || common.IsNil(o.BorrowLimit) {
		var ret string
		return ret
	}
	return *o.BorrowLimit
}

// GetBorrowLimitOk returns a tuple with the BorrowLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMaxBorrowResponse) GetBorrowLimitOk() (*string, bool) {
	if o == nil || common.IsNil(o.BorrowLimit) {
		return nil, false
	}
	return o.BorrowLimit, true
}

// HasBorrowLimit returns a boolean if a field has been set.
func (o *QueryMaxBorrowResponse) HasBorrowLimit() bool {
	if o != nil && !common.IsNil(o.BorrowLimit) {
		return true
	}

	return false
}

// SetBorrowLimit gets a reference to the given string and assigns it to the BorrowLimit field.
func (o *QueryMaxBorrowResponse) SetBorrowLimit(v string) {
	o.BorrowLimit = &v
}

func (o QueryMaxBorrowResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryMaxBorrowResponse) ToMap() (map[string]interface{}, error) {
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

func (o *QueryMaxBorrowResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryMaxBorrowResponse := _QueryMaxBorrowResponse{}

	err = json.Unmarshal(data, &varQueryMaxBorrowResponse)

	if err != nil {
		return err
	}

	*o = QueryMaxBorrowResponse(varQueryMaxBorrowResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "borrowLimit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryMaxBorrowResponse struct {
	value *QueryMaxBorrowResponse
	isSet bool
}

func (v NullableQueryMaxBorrowResponse) Get() *QueryMaxBorrowResponse {
	return v.value
}

func (v *NullableQueryMaxBorrowResponse) Set(val *QueryMaxBorrowResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryMaxBorrowResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryMaxBorrowResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryMaxBorrowResponse(val *QueryMaxBorrowResponse) *NullableQueryMaxBorrowResponse {
	return &NullableQueryMaxBorrowResponse{value: val, isSet: true}
}

func (v NullableQueryMaxBorrowResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryMaxBorrowResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
