/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the RepayFuturesNegativeBalanceResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RepayFuturesNegativeBalanceResponse{}

// RepayFuturesNegativeBalanceResponse struct for RepayFuturesNegativeBalanceResponse
type RepayFuturesNegativeBalanceResponse struct {
	Msg                  *string `json:"msg,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RepayFuturesNegativeBalanceResponse RepayFuturesNegativeBalanceResponse

// NewRepayFuturesNegativeBalanceResponse instantiates a new RepayFuturesNegativeBalanceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepayFuturesNegativeBalanceResponse() *RepayFuturesNegativeBalanceResponse {
	this := RepayFuturesNegativeBalanceResponse{}
	return &this
}

// NewRepayFuturesNegativeBalanceResponseWithDefaults instantiates a new RepayFuturesNegativeBalanceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepayFuturesNegativeBalanceResponseWithDefaults() *RepayFuturesNegativeBalanceResponse {
	this := RepayFuturesNegativeBalanceResponse{}
	return &this
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *RepayFuturesNegativeBalanceResponse) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepayFuturesNegativeBalanceResponse) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *RepayFuturesNegativeBalanceResponse) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *RepayFuturesNegativeBalanceResponse) SetMsg(v string) {
	o.Msg = &v
}

func (o RepayFuturesNegativeBalanceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepayFuturesNegativeBalanceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RepayFuturesNegativeBalanceResponse) UnmarshalJSON(data []byte) (err error) {
	varRepayFuturesNegativeBalanceResponse := _RepayFuturesNegativeBalanceResponse{}

	err = json.Unmarshal(data, &varRepayFuturesNegativeBalanceResponse)

	if err != nil {
		return err
	}

	*o = RepayFuturesNegativeBalanceResponse(varRepayFuturesNegativeBalanceResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "msg")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRepayFuturesNegativeBalanceResponse struct {
	value *RepayFuturesNegativeBalanceResponse
	isSet bool
}

func (v NullableRepayFuturesNegativeBalanceResponse) Get() *RepayFuturesNegativeBalanceResponse {
	return v.value
}

func (v *NullableRepayFuturesNegativeBalanceResponse) Set(val *RepayFuturesNegativeBalanceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRepayFuturesNegativeBalanceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRepayFuturesNegativeBalanceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepayFuturesNegativeBalanceResponse(val *RepayFuturesNegativeBalanceResponse) *NullableRepayFuturesNegativeBalanceResponse {
	return &NullableRepayFuturesNegativeBalanceResponse{value: val, isSet: true}
}

func (v NullableRepayFuturesNegativeBalanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepayFuturesNegativeBalanceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
