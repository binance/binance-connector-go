/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the FundAutoCollectionResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FundAutoCollectionResponse{}

// FundAutoCollectionResponse struct for FundAutoCollectionResponse
type FundAutoCollectionResponse struct {
	Msg                  *string `json:"msg,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FundAutoCollectionResponse FundAutoCollectionResponse

// NewFundAutoCollectionResponse instantiates a new FundAutoCollectionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFundAutoCollectionResponse() *FundAutoCollectionResponse {
	this := FundAutoCollectionResponse{}
	return &this
}

// NewFundAutoCollectionResponseWithDefaults instantiates a new FundAutoCollectionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFundAutoCollectionResponseWithDefaults() *FundAutoCollectionResponse {
	this := FundAutoCollectionResponse{}
	return &this
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *FundAutoCollectionResponse) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundAutoCollectionResponse) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *FundAutoCollectionResponse) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *FundAutoCollectionResponse) SetMsg(v string) {
	o.Msg = &v
}

func (o FundAutoCollectionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FundAutoCollectionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FundAutoCollectionResponse) UnmarshalJSON(data []byte) (err error) {
	varFundAutoCollectionResponse := _FundAutoCollectionResponse{}

	err = json.Unmarshal(data, &varFundAutoCollectionResponse)

	if err != nil {
		return err
	}

	*o = FundAutoCollectionResponse(varFundAutoCollectionResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "msg")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFundAutoCollectionResponse struct {
	value *FundAutoCollectionResponse
	isSet bool
}

func (v NullableFundAutoCollectionResponse) Get() *FundAutoCollectionResponse {
	return v.value
}

func (v *NullableFundAutoCollectionResponse) Set(val *FundAutoCollectionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFundAutoCollectionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFundAutoCollectionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFundAutoCollectionResponse(val *FundAutoCollectionResponse) *NullableFundAutoCollectionResponse {
	return &NullableFundAutoCollectionResponse{value: val, isSet: true}
}

func (v NullableFundAutoCollectionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFundAutoCollectionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
