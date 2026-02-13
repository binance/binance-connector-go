/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetCurrentMultiAssetsModeResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetCurrentMultiAssetsModeResponse{}

// GetCurrentMultiAssetsModeResponse struct for GetCurrentMultiAssetsModeResponse
type GetCurrentMultiAssetsModeResponse struct {
	MultiAssetsMargin    *bool `json:"multiAssetsMargin,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetCurrentMultiAssetsModeResponse GetCurrentMultiAssetsModeResponse

// NewGetCurrentMultiAssetsModeResponse instantiates a new GetCurrentMultiAssetsModeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCurrentMultiAssetsModeResponse() *GetCurrentMultiAssetsModeResponse {
	this := GetCurrentMultiAssetsModeResponse{}
	return &this
}

// NewGetCurrentMultiAssetsModeResponseWithDefaults instantiates a new GetCurrentMultiAssetsModeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCurrentMultiAssetsModeResponseWithDefaults() *GetCurrentMultiAssetsModeResponse {
	this := GetCurrentMultiAssetsModeResponse{}
	return &this
}

// GetMultiAssetsMargin returns the MultiAssetsMargin field value if set, zero value otherwise.
func (o *GetCurrentMultiAssetsModeResponse) GetMultiAssetsMargin() bool {
	if o == nil || common.IsNil(o.MultiAssetsMargin) {
		var ret bool
		return ret
	}
	return *o.MultiAssetsMargin
}

// GetMultiAssetsMarginOk returns a tuple with the MultiAssetsMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrentMultiAssetsModeResponse) GetMultiAssetsMarginOk() (*bool, bool) {
	if o == nil || common.IsNil(o.MultiAssetsMargin) {
		return nil, false
	}
	return o.MultiAssetsMargin, true
}

// HasMultiAssetsMargin returns a boolean if a field has been set.
func (o *GetCurrentMultiAssetsModeResponse) HasMultiAssetsMargin() bool {
	if o != nil && !common.IsNil(o.MultiAssetsMargin) {
		return true
	}

	return false
}

// SetMultiAssetsMargin gets a reference to the given bool and assigns it to the MultiAssetsMargin field.
func (o *GetCurrentMultiAssetsModeResponse) SetMultiAssetsMargin(v bool) {
	o.MultiAssetsMargin = &v
}

func (o GetCurrentMultiAssetsModeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCurrentMultiAssetsModeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.MultiAssetsMargin) {
		toSerialize["multiAssetsMargin"] = o.MultiAssetsMargin
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetCurrentMultiAssetsModeResponse) UnmarshalJSON(data []byte) (err error) {
	varGetCurrentMultiAssetsModeResponse := _GetCurrentMultiAssetsModeResponse{}

	err = json.Unmarshal(data, &varGetCurrentMultiAssetsModeResponse)

	if err != nil {
		return err
	}

	*o = GetCurrentMultiAssetsModeResponse(varGetCurrentMultiAssetsModeResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "multiAssetsMargin")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetCurrentMultiAssetsModeResponse struct {
	value *GetCurrentMultiAssetsModeResponse
	isSet bool
}

func (v NullableGetCurrentMultiAssetsModeResponse) Get() *GetCurrentMultiAssetsModeResponse {
	return v.value
}

func (v *NullableGetCurrentMultiAssetsModeResponse) Set(val *GetCurrentMultiAssetsModeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCurrentMultiAssetsModeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCurrentMultiAssetsModeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCurrentMultiAssetsModeResponse(val *GetCurrentMultiAssetsModeResponse) *NullableGetCurrentMultiAssetsModeResponse {
	return &NullableGetCurrentMultiAssetsModeResponse{value: val, isSet: true}
}

func (v NullableGetCurrentMultiAssetsModeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCurrentMultiAssetsModeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
