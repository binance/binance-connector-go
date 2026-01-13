/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AdjustCrossMarginMaxLeverageResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AdjustCrossMarginMaxLeverageResponse{}

// AdjustCrossMarginMaxLeverageResponse struct for AdjustCrossMarginMaxLeverageResponse
type AdjustCrossMarginMaxLeverageResponse struct {
	Success              *bool `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AdjustCrossMarginMaxLeverageResponse AdjustCrossMarginMaxLeverageResponse

// NewAdjustCrossMarginMaxLeverageResponse instantiates a new AdjustCrossMarginMaxLeverageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdjustCrossMarginMaxLeverageResponse() *AdjustCrossMarginMaxLeverageResponse {
	this := AdjustCrossMarginMaxLeverageResponse{}
	return &this
}

// NewAdjustCrossMarginMaxLeverageResponseWithDefaults instantiates a new AdjustCrossMarginMaxLeverageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdjustCrossMarginMaxLeverageResponseWithDefaults() *AdjustCrossMarginMaxLeverageResponse {
	this := AdjustCrossMarginMaxLeverageResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *AdjustCrossMarginMaxLeverageResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdjustCrossMarginMaxLeverageResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *AdjustCrossMarginMaxLeverageResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *AdjustCrossMarginMaxLeverageResponse) SetSuccess(v bool) {
	o.Success = &v
}

func (o AdjustCrossMarginMaxLeverageResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdjustCrossMarginMaxLeverageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AdjustCrossMarginMaxLeverageResponse) UnmarshalJSON(data []byte) (err error) {
	varAdjustCrossMarginMaxLeverageResponse := _AdjustCrossMarginMaxLeverageResponse{}

	err = json.Unmarshal(data, &varAdjustCrossMarginMaxLeverageResponse)

	if err != nil {
		return err
	}

	*o = AdjustCrossMarginMaxLeverageResponse(varAdjustCrossMarginMaxLeverageResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAdjustCrossMarginMaxLeverageResponse struct {
	value *AdjustCrossMarginMaxLeverageResponse
	isSet bool
}

func (v NullableAdjustCrossMarginMaxLeverageResponse) Get() *AdjustCrossMarginMaxLeverageResponse {
	return v.value
}

func (v *NullableAdjustCrossMarginMaxLeverageResponse) Set(val *AdjustCrossMarginMaxLeverageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAdjustCrossMarginMaxLeverageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAdjustCrossMarginMaxLeverageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdjustCrossMarginMaxLeverageResponse(val *AdjustCrossMarginMaxLeverageResponse) *NullableAdjustCrossMarginMaxLeverageResponse {
	return &NullableAdjustCrossMarginMaxLeverageResponse{value: val, isSet: true}
}

func (v NullableAdjustCrossMarginMaxLeverageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdjustCrossMarginMaxLeverageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
