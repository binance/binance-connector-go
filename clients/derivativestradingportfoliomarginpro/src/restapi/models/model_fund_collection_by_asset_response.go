/*
Binance Derivatives Trading Portfolio Margin Pro REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin Pro REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the FundCollectionByAssetResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FundCollectionByAssetResponse{}

// FundCollectionByAssetResponse struct for FundCollectionByAssetResponse
type FundCollectionByAssetResponse struct {
	Msg                  *string `json:"msg,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FundCollectionByAssetResponse FundCollectionByAssetResponse

// NewFundCollectionByAssetResponse instantiates a new FundCollectionByAssetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFundCollectionByAssetResponse() *FundCollectionByAssetResponse {
	this := FundCollectionByAssetResponse{}
	return &this
}

// NewFundCollectionByAssetResponseWithDefaults instantiates a new FundCollectionByAssetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFundCollectionByAssetResponseWithDefaults() *FundCollectionByAssetResponse {
	this := FundCollectionByAssetResponse{}
	return &this
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *FundCollectionByAssetResponse) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundCollectionByAssetResponse) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *FundCollectionByAssetResponse) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *FundCollectionByAssetResponse) SetMsg(v string) {
	o.Msg = &v
}

func (o FundCollectionByAssetResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FundCollectionByAssetResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FundCollectionByAssetResponse) UnmarshalJSON(data []byte) (err error) {
	varFundCollectionByAssetResponse := _FundCollectionByAssetResponse{}

	err = json.Unmarshal(data, &varFundCollectionByAssetResponse)

	if err != nil {
		return err
	}

	*o = FundCollectionByAssetResponse(varFundCollectionByAssetResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "msg")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFundCollectionByAssetResponse struct {
	value *FundCollectionByAssetResponse
	isSet bool
}

func (v NullableFundCollectionByAssetResponse) Get() *FundCollectionByAssetResponse {
	return v.value
}

func (v *NullableFundCollectionByAssetResponse) Set(val *FundCollectionByAssetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFundCollectionByAssetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFundCollectionByAssetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFundCollectionByAssetResponse(val *FundCollectionByAssetResponse) *NullableFundCollectionByAssetResponse {
	return &NullableFundCollectionByAssetResponse{value: val, isSet: true}
}

func (v NullableFundCollectionByAssetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFundCollectionByAssetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
