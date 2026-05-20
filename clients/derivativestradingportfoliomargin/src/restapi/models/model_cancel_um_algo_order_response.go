/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the CancelUmAlgoOrderResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CancelUmAlgoOrderResponse{}

// CancelUmAlgoOrderResponse struct for CancelUmAlgoOrderResponse
type CancelUmAlgoOrderResponse struct {
	Complete             *bool `json:"complete,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CancelUmAlgoOrderResponse CancelUmAlgoOrderResponse

// NewCancelUmAlgoOrderResponse instantiates a new CancelUmAlgoOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelUmAlgoOrderResponse() *CancelUmAlgoOrderResponse {
	this := CancelUmAlgoOrderResponse{}
	return &this
}

// NewCancelUmAlgoOrderResponseWithDefaults instantiates a new CancelUmAlgoOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelUmAlgoOrderResponseWithDefaults() *CancelUmAlgoOrderResponse {
	this := CancelUmAlgoOrderResponse{}
	return &this
}

// GetComplete returns the Complete field value if set, zero value otherwise.
func (o *CancelUmAlgoOrderResponse) GetComplete() bool {
	if o == nil || common.IsNil(o.Complete) {
		var ret bool
		return ret
	}
	return *o.Complete
}

// GetCompleteOk returns a tuple with the Complete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelUmAlgoOrderResponse) GetCompleteOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Complete) {
		return nil, false
	}
	return o.Complete, true
}

// HasComplete returns a boolean if a field has been set.
func (o *CancelUmAlgoOrderResponse) HasComplete() bool {
	if o != nil && !common.IsNil(o.Complete) {
		return true
	}

	return false
}

// SetComplete gets a reference to the given bool and assigns it to the Complete field.
func (o *CancelUmAlgoOrderResponse) SetComplete(v bool) {
	o.Complete = &v
}

func (o CancelUmAlgoOrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelUmAlgoOrderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Complete) {
		toSerialize["complete"] = o.Complete
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CancelUmAlgoOrderResponse) UnmarshalJSON(data []byte) (err error) {
	varCancelUmAlgoOrderResponse := _CancelUmAlgoOrderResponse{}

	err = json.Unmarshal(data, &varCancelUmAlgoOrderResponse)

	if err != nil {
		return err
	}

	*o = CancelUmAlgoOrderResponse(varCancelUmAlgoOrderResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "complete")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCancelUmAlgoOrderResponse struct {
	value *CancelUmAlgoOrderResponse
	isSet bool
}

func (v NullableCancelUmAlgoOrderResponse) Get() *CancelUmAlgoOrderResponse {
	return v.value
}

func (v *NullableCancelUmAlgoOrderResponse) Set(val *CancelUmAlgoOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelUmAlgoOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelUmAlgoOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelUmAlgoOrderResponse(val *CancelUmAlgoOrderResponse) *NullableCancelUmAlgoOrderResponse {
	return &NullableCancelUmAlgoOrderResponse{value: val, isSet: true}
}

func (v NullableCancelUmAlgoOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelUmAlgoOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
