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

// checks if the GetSummaryOfMarginAccountResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSummaryOfMarginAccountResponse{}

// GetSummaryOfMarginAccountResponse struct for GetSummaryOfMarginAccountResponse
type GetSummaryOfMarginAccountResponse struct {
	NormalBar            *string `json:"normalBar,omitempty"`
	MarginCallBar        *string `json:"marginCallBar,omitempty"`
	ForceLiquidationBar  *string `json:"forceLiquidationBar,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetSummaryOfMarginAccountResponse GetSummaryOfMarginAccountResponse

// NewGetSummaryOfMarginAccountResponse instantiates a new GetSummaryOfMarginAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSummaryOfMarginAccountResponse() *GetSummaryOfMarginAccountResponse {
	this := GetSummaryOfMarginAccountResponse{}
	return &this
}

// NewGetSummaryOfMarginAccountResponseWithDefaults instantiates a new GetSummaryOfMarginAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSummaryOfMarginAccountResponseWithDefaults() *GetSummaryOfMarginAccountResponse {
	this := GetSummaryOfMarginAccountResponse{}
	return &this
}

// GetNormalBar returns the NormalBar field value if set, zero value otherwise.
func (o *GetSummaryOfMarginAccountResponse) GetNormalBar() string {
	if o == nil || common.IsNil(o.NormalBar) {
		var ret string
		return ret
	}
	return *o.NormalBar
}

// GetNormalBarOk returns a tuple with the NormalBar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfMarginAccountResponse) GetNormalBarOk() (*string, bool) {
	if o == nil || common.IsNil(o.NormalBar) {
		return nil, false
	}
	return o.NormalBar, true
}

// HasNormalBar returns a boolean if a field has been set.
func (o *GetSummaryOfMarginAccountResponse) HasNormalBar() bool {
	if o != nil && !common.IsNil(o.NormalBar) {
		return true
	}

	return false
}

// SetNormalBar gets a reference to the given string and assigns it to the NormalBar field.
func (o *GetSummaryOfMarginAccountResponse) SetNormalBar(v string) {
	o.NormalBar = &v
}

// GetMarginCallBar returns the MarginCallBar field value if set, zero value otherwise.
func (o *GetSummaryOfMarginAccountResponse) GetMarginCallBar() string {
	if o == nil || common.IsNil(o.MarginCallBar) {
		var ret string
		return ret
	}
	return *o.MarginCallBar
}

// GetMarginCallBarOk returns a tuple with the MarginCallBar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfMarginAccountResponse) GetMarginCallBarOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarginCallBar) {
		return nil, false
	}
	return o.MarginCallBar, true
}

// HasMarginCallBar returns a boolean if a field has been set.
func (o *GetSummaryOfMarginAccountResponse) HasMarginCallBar() bool {
	if o != nil && !common.IsNil(o.MarginCallBar) {
		return true
	}

	return false
}

// SetMarginCallBar gets a reference to the given string and assigns it to the MarginCallBar field.
func (o *GetSummaryOfMarginAccountResponse) SetMarginCallBar(v string) {
	o.MarginCallBar = &v
}

// GetForceLiquidationBar returns the ForceLiquidationBar field value if set, zero value otherwise.
func (o *GetSummaryOfMarginAccountResponse) GetForceLiquidationBar() string {
	if o == nil || common.IsNil(o.ForceLiquidationBar) {
		var ret string
		return ret
	}
	return *o.ForceLiquidationBar
}

// GetForceLiquidationBarOk returns a tuple with the ForceLiquidationBar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfMarginAccountResponse) GetForceLiquidationBarOk() (*string, bool) {
	if o == nil || common.IsNil(o.ForceLiquidationBar) {
		return nil, false
	}
	return o.ForceLiquidationBar, true
}

// HasForceLiquidationBar returns a boolean if a field has been set.
func (o *GetSummaryOfMarginAccountResponse) HasForceLiquidationBar() bool {
	if o != nil && !common.IsNil(o.ForceLiquidationBar) {
		return true
	}

	return false
}

// SetForceLiquidationBar gets a reference to the given string and assigns it to the ForceLiquidationBar field.
func (o *GetSummaryOfMarginAccountResponse) SetForceLiquidationBar(v string) {
	o.ForceLiquidationBar = &v
}

func (o GetSummaryOfMarginAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSummaryOfMarginAccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.NormalBar) {
		toSerialize["normalBar"] = o.NormalBar
	}
	if !common.IsNil(o.MarginCallBar) {
		toSerialize["marginCallBar"] = o.MarginCallBar
	}
	if !common.IsNil(o.ForceLiquidationBar) {
		toSerialize["forceLiquidationBar"] = o.ForceLiquidationBar
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetSummaryOfMarginAccountResponse) UnmarshalJSON(data []byte) (err error) {
	varGetSummaryOfMarginAccountResponse := _GetSummaryOfMarginAccountResponse{}

	err = json.Unmarshal(data, &varGetSummaryOfMarginAccountResponse)

	if err != nil {
		return err
	}

	*o = GetSummaryOfMarginAccountResponse(varGetSummaryOfMarginAccountResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "normalBar")
		delete(additionalProperties, "marginCallBar")
		delete(additionalProperties, "forceLiquidationBar")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSummaryOfMarginAccountResponse struct {
	value *GetSummaryOfMarginAccountResponse
	isSet bool
}

func (v NullableGetSummaryOfMarginAccountResponse) Get() *GetSummaryOfMarginAccountResponse {
	return v.value
}

func (v *NullableGetSummaryOfMarginAccountResponse) Set(val *GetSummaryOfMarginAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSummaryOfMarginAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSummaryOfMarginAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSummaryOfMarginAccountResponse(val *GetSummaryOfMarginAccountResponse) *NullableGetSummaryOfMarginAccountResponse {
	return &NullableGetSummaryOfMarginAccountResponse{value: val, isSet: true}
}

func (v NullableGetSummaryOfMarginAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSummaryOfMarginAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
