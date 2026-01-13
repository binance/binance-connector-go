/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryMarginMaxWithdrawResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryMarginMaxWithdrawResponse{}

// QueryMarginMaxWithdrawResponse struct for QueryMarginMaxWithdrawResponse
type QueryMarginMaxWithdrawResponse struct {
	Amount               *string `json:"amount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryMarginMaxWithdrawResponse QueryMarginMaxWithdrawResponse

// NewQueryMarginMaxWithdrawResponse instantiates a new QueryMarginMaxWithdrawResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryMarginMaxWithdrawResponse() *QueryMarginMaxWithdrawResponse {
	this := QueryMarginMaxWithdrawResponse{}
	return &this
}

// NewQueryMarginMaxWithdrawResponseWithDefaults instantiates a new QueryMarginMaxWithdrawResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryMarginMaxWithdrawResponseWithDefaults() *QueryMarginMaxWithdrawResponse {
	this := QueryMarginMaxWithdrawResponse{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *QueryMarginMaxWithdrawResponse) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginMaxWithdrawResponse) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *QueryMarginMaxWithdrawResponse) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *QueryMarginMaxWithdrawResponse) SetAmount(v string) {
	o.Amount = &v
}

func (o QueryMarginMaxWithdrawResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryMarginMaxWithdrawResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryMarginMaxWithdrawResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryMarginMaxWithdrawResponse := _QueryMarginMaxWithdrawResponse{}

	err = json.Unmarshal(data, &varQueryMarginMaxWithdrawResponse)

	if err != nil {
		return err
	}

	*o = QueryMarginMaxWithdrawResponse(varQueryMarginMaxWithdrawResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryMarginMaxWithdrawResponse struct {
	value *QueryMarginMaxWithdrawResponse
	isSet bool
}

func (v NullableQueryMarginMaxWithdrawResponse) Get() *QueryMarginMaxWithdrawResponse {
	return v.value
}

func (v *NullableQueryMarginMaxWithdrawResponse) Set(val *QueryMarginMaxWithdrawResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryMarginMaxWithdrawResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryMarginMaxWithdrawResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryMarginMaxWithdrawResponse(val *QueryMarginMaxWithdrawResponse) *NullableQueryMarginMaxWithdrawResponse {
	return &NullableQueryMarginMaxWithdrawResponse{value: val, isSet: true}
}

func (v NullableQueryMarginMaxWithdrawResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryMarginMaxWithdrawResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
