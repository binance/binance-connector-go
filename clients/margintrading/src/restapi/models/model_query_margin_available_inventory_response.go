/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryMarginAvailableInventoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryMarginAvailableInventoryResponse{}

// QueryMarginAvailableInventoryResponse struct for QueryMarginAvailableInventoryResponse
type QueryMarginAvailableInventoryResponse struct {
	Assets               *QueryMarginAvailableInventoryResponseAssets `json:"assets,omitempty"`
	UpdateTime           *int64                                       `json:"updateTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryMarginAvailableInventoryResponse QueryMarginAvailableInventoryResponse

// NewQueryMarginAvailableInventoryResponse instantiates a new QueryMarginAvailableInventoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryMarginAvailableInventoryResponse() *QueryMarginAvailableInventoryResponse {
	this := QueryMarginAvailableInventoryResponse{}
	return &this
}

// NewQueryMarginAvailableInventoryResponseWithDefaults instantiates a new QueryMarginAvailableInventoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryMarginAvailableInventoryResponseWithDefaults() *QueryMarginAvailableInventoryResponse {
	this := QueryMarginAvailableInventoryResponse{}
	return &this
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *QueryMarginAvailableInventoryResponse) GetAssets() QueryMarginAvailableInventoryResponseAssets {
	if o == nil || common.IsNil(o.Assets) {
		var ret QueryMarginAvailableInventoryResponseAssets
		return ret
	}
	return *o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAvailableInventoryResponse) GetAssetsOk() (*QueryMarginAvailableInventoryResponseAssets, bool) {
	if o == nil || common.IsNil(o.Assets) {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *QueryMarginAvailableInventoryResponse) HasAssets() bool {
	if o != nil && !common.IsNil(o.Assets) {
		return true
	}

	return false
}

// SetAssets gets a reference to the given QueryMarginAvailableInventoryResponseAssets and assigns it to the Assets field.
func (o *QueryMarginAvailableInventoryResponse) SetAssets(v QueryMarginAvailableInventoryResponseAssets) {
	o.Assets = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *QueryMarginAvailableInventoryResponse) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginAvailableInventoryResponse) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *QueryMarginAvailableInventoryResponse) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *QueryMarginAvailableInventoryResponse) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o QueryMarginAvailableInventoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryMarginAvailableInventoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Assets) {
		toSerialize["assets"] = o.Assets
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryMarginAvailableInventoryResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryMarginAvailableInventoryResponse := _QueryMarginAvailableInventoryResponse{}

	err = json.Unmarshal(data, &varQueryMarginAvailableInventoryResponse)

	if err != nil {
		return err
	}

	*o = QueryMarginAvailableInventoryResponse(varQueryMarginAvailableInventoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "assets")
		delete(additionalProperties, "updateTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryMarginAvailableInventoryResponse struct {
	value *QueryMarginAvailableInventoryResponse
	isSet bool
}

func (v NullableQueryMarginAvailableInventoryResponse) Get() *QueryMarginAvailableInventoryResponse {
	return v.value
}

func (v *NullableQueryMarginAvailableInventoryResponse) Set(val *QueryMarginAvailableInventoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryMarginAvailableInventoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryMarginAvailableInventoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryMarginAvailableInventoryResponse(val *QueryMarginAvailableInventoryResponse) *NullableQueryMarginAvailableInventoryResponse {
	return &NullableQueryMarginAvailableInventoryResponse{value: val, isSet: true}
}

func (v NullableQueryMarginAvailableInventoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryMarginAvailableInventoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
