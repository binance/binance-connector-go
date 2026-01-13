/*
Binance Derivatives Trading Portfolio Margin Pro REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin Pro REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the TransferLdusdtRwusdForPortfolioMarginResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TransferLdusdtRwusdForPortfolioMarginResponse{}

// TransferLdusdtRwusdForPortfolioMarginResponse struct for TransferLdusdtRwusdForPortfolioMarginResponse
type TransferLdusdtRwusdForPortfolioMarginResponse struct {
	Msg                  *string `json:"msg,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TransferLdusdtRwusdForPortfolioMarginResponse TransferLdusdtRwusdForPortfolioMarginResponse

// NewTransferLdusdtRwusdForPortfolioMarginResponse instantiates a new TransferLdusdtRwusdForPortfolioMarginResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferLdusdtRwusdForPortfolioMarginResponse() *TransferLdusdtRwusdForPortfolioMarginResponse {
	this := TransferLdusdtRwusdForPortfolioMarginResponse{}
	return &this
}

// NewTransferLdusdtRwusdForPortfolioMarginResponseWithDefaults instantiates a new TransferLdusdtRwusdForPortfolioMarginResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferLdusdtRwusdForPortfolioMarginResponseWithDefaults() *TransferLdusdtRwusdForPortfolioMarginResponse {
	this := TransferLdusdtRwusdForPortfolioMarginResponse{}
	return &this
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *TransferLdusdtRwusdForPortfolioMarginResponse) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferLdusdtRwusdForPortfolioMarginResponse) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *TransferLdusdtRwusdForPortfolioMarginResponse) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *TransferLdusdtRwusdForPortfolioMarginResponse) SetMsg(v string) {
	o.Msg = &v
}

func (o TransferLdusdtRwusdForPortfolioMarginResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferLdusdtRwusdForPortfolioMarginResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TransferLdusdtRwusdForPortfolioMarginResponse) UnmarshalJSON(data []byte) (err error) {
	varTransferLdusdtRwusdForPortfolioMarginResponse := _TransferLdusdtRwusdForPortfolioMarginResponse{}

	err = json.Unmarshal(data, &varTransferLdusdtRwusdForPortfolioMarginResponse)

	if err != nil {
		return err
	}

	*o = TransferLdusdtRwusdForPortfolioMarginResponse(varTransferLdusdtRwusdForPortfolioMarginResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "msg")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferLdusdtRwusdForPortfolioMarginResponse struct {
	value *TransferLdusdtRwusdForPortfolioMarginResponse
	isSet bool
}

func (v NullableTransferLdusdtRwusdForPortfolioMarginResponse) Get() *TransferLdusdtRwusdForPortfolioMarginResponse {
	return v.value
}

func (v *NullableTransferLdusdtRwusdForPortfolioMarginResponse) Set(val *TransferLdusdtRwusdForPortfolioMarginResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferLdusdtRwusdForPortfolioMarginResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferLdusdtRwusdForPortfolioMarginResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferLdusdtRwusdForPortfolioMarginResponse(val *TransferLdusdtRwusdForPortfolioMarginResponse) *NullableTransferLdusdtRwusdForPortfolioMarginResponse {
	return &NullableTransferLdusdtRwusdForPortfolioMarginResponse{value: val, isSet: true}
}

func (v NullableTransferLdusdtRwusdForPortfolioMarginResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferLdusdtRwusdForPortfolioMarginResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
