/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryMaxTransferOutAmountResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryMaxTransferOutAmountResponse{}

// QueryMaxTransferOutAmountResponse struct for QueryMaxTransferOutAmountResponse
type QueryMaxTransferOutAmountResponse struct {
	Amount               *string `json:"amount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryMaxTransferOutAmountResponse QueryMaxTransferOutAmountResponse

// NewQueryMaxTransferOutAmountResponse instantiates a new QueryMaxTransferOutAmountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryMaxTransferOutAmountResponse() *QueryMaxTransferOutAmountResponse {
	this := QueryMaxTransferOutAmountResponse{}
	return &this
}

// NewQueryMaxTransferOutAmountResponseWithDefaults instantiates a new QueryMaxTransferOutAmountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryMaxTransferOutAmountResponseWithDefaults() *QueryMaxTransferOutAmountResponse {
	this := QueryMaxTransferOutAmountResponse{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *QueryMaxTransferOutAmountResponse) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMaxTransferOutAmountResponse) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *QueryMaxTransferOutAmountResponse) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *QueryMaxTransferOutAmountResponse) SetAmount(v string) {
	o.Amount = &v
}

func (o QueryMaxTransferOutAmountResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryMaxTransferOutAmountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryMaxTransferOutAmountResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryMaxTransferOutAmountResponse := _QueryMaxTransferOutAmountResponse{}

	err = json.Unmarshal(data, &varQueryMaxTransferOutAmountResponse)

	if err != nil {
		return err
	}

	*o = QueryMaxTransferOutAmountResponse(varQueryMaxTransferOutAmountResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryMaxTransferOutAmountResponse struct {
	value *QueryMaxTransferOutAmountResponse
	isSet bool
}

func (v NullableQueryMaxTransferOutAmountResponse) Get() *QueryMaxTransferOutAmountResponse {
	return v.value
}

func (v *NullableQueryMaxTransferOutAmountResponse) Set(val *QueryMaxTransferOutAmountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryMaxTransferOutAmountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryMaxTransferOutAmountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryMaxTransferOutAmountResponse(val *QueryMaxTransferOutAmountResponse) *NullableQueryMaxTransferOutAmountResponse {
	return &NullableQueryMaxTransferOutAmountResponse{value: val, isSet: true}
}

func (v NullableQueryMaxTransferOutAmountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryMaxTransferOutAmountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
