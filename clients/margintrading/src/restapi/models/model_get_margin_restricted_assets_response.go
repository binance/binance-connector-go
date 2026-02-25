/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetMarginRestrictedAssetsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetMarginRestrictedAssetsResponse{}

// GetMarginRestrictedAssetsResponse struct for GetMarginRestrictedAssetsResponse
type GetMarginRestrictedAssetsResponse struct {
	OpenLongRestrictedAsset    []string `json:"openLongRestrictedAsset,omitempty"`
	MaxCollateralExceededAsset []string `json:"maxCollateralExceededAsset,omitempty"`
	AdditionalProperties       map[string]interface{}
}

type _GetMarginRestrictedAssetsResponse GetMarginRestrictedAssetsResponse

// NewGetMarginRestrictedAssetsResponse instantiates a new GetMarginRestrictedAssetsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMarginRestrictedAssetsResponse() *GetMarginRestrictedAssetsResponse {
	this := GetMarginRestrictedAssetsResponse{}
	return &this
}

// NewGetMarginRestrictedAssetsResponseWithDefaults instantiates a new GetMarginRestrictedAssetsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMarginRestrictedAssetsResponseWithDefaults() *GetMarginRestrictedAssetsResponse {
	this := GetMarginRestrictedAssetsResponse{}
	return &this
}

// GetOpenLongRestrictedAsset returns the OpenLongRestrictedAsset field value if set, zero value otherwise.
func (o *GetMarginRestrictedAssetsResponse) GetOpenLongRestrictedAsset() []string {
	if o == nil || common.IsNil(o.OpenLongRestrictedAsset) {
		var ret []string
		return ret
	}
	return o.OpenLongRestrictedAsset
}

// GetOpenLongRestrictedAssetOk returns a tuple with the OpenLongRestrictedAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginRestrictedAssetsResponse) GetOpenLongRestrictedAssetOk() ([]string, bool) {
	if o == nil || common.IsNil(o.OpenLongRestrictedAsset) {
		return nil, false
	}
	return o.OpenLongRestrictedAsset, true
}

// HasOpenLongRestrictedAsset returns a boolean if a field has been set.
func (o *GetMarginRestrictedAssetsResponse) HasOpenLongRestrictedAsset() bool {
	if o != nil && !common.IsNil(o.OpenLongRestrictedAsset) {
		return true
	}

	return false
}

// SetOpenLongRestrictedAsset gets a reference to the given []string and assigns it to the OpenLongRestrictedAsset field.
func (o *GetMarginRestrictedAssetsResponse) SetOpenLongRestrictedAsset(v []string) {
	o.OpenLongRestrictedAsset = v
}

// GetMaxCollateralExceededAsset returns the MaxCollateralExceededAsset field value if set, zero value otherwise.
func (o *GetMarginRestrictedAssetsResponse) GetMaxCollateralExceededAsset() []string {
	if o == nil || common.IsNil(o.MaxCollateralExceededAsset) {
		var ret []string
		return ret
	}
	return o.MaxCollateralExceededAsset
}

// GetMaxCollateralExceededAssetOk returns a tuple with the MaxCollateralExceededAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginRestrictedAssetsResponse) GetMaxCollateralExceededAssetOk() ([]string, bool) {
	if o == nil || common.IsNil(o.MaxCollateralExceededAsset) {
		return nil, false
	}
	return o.MaxCollateralExceededAsset, true
}

// HasMaxCollateralExceededAsset returns a boolean if a field has been set.
func (o *GetMarginRestrictedAssetsResponse) HasMaxCollateralExceededAsset() bool {
	if o != nil && !common.IsNil(o.MaxCollateralExceededAsset) {
		return true
	}

	return false
}

// SetMaxCollateralExceededAsset gets a reference to the given []string and assigns it to the MaxCollateralExceededAsset field.
func (o *GetMarginRestrictedAssetsResponse) SetMaxCollateralExceededAsset(v []string) {
	o.MaxCollateralExceededAsset = v
}

func (o GetMarginRestrictedAssetsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMarginRestrictedAssetsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.OpenLongRestrictedAsset) {
		toSerialize["openLongRestrictedAsset"] = o.OpenLongRestrictedAsset
	}
	if !common.IsNil(o.MaxCollateralExceededAsset) {
		toSerialize["maxCollateralExceededAsset"] = o.MaxCollateralExceededAsset
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetMarginRestrictedAssetsResponse) UnmarshalJSON(data []byte) (err error) {
	varGetMarginRestrictedAssetsResponse := _GetMarginRestrictedAssetsResponse{}

	err = json.Unmarshal(data, &varGetMarginRestrictedAssetsResponse)

	if err != nil {
		return err
	}

	*o = GetMarginRestrictedAssetsResponse(varGetMarginRestrictedAssetsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "openLongRestrictedAsset")
		delete(additionalProperties, "maxCollateralExceededAsset")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetMarginRestrictedAssetsResponse struct {
	value *GetMarginRestrictedAssetsResponse
	isSet bool
}

func (v NullableGetMarginRestrictedAssetsResponse) Get() *GetMarginRestrictedAssetsResponse {
	return v.value
}

func (v *NullableGetMarginRestrictedAssetsResponse) Set(val *GetMarginRestrictedAssetsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMarginRestrictedAssetsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMarginRestrictedAssetsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMarginRestrictedAssetsResponse(val *GetMarginRestrictedAssetsResponse) *NullableGetMarginRestrictedAssetsResponse {
	return &NullableGetMarginRestrictedAssetsResponse{value: val, isSet: true}
}

func (v NullableGetMarginRestrictedAssetsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMarginRestrictedAssetsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
