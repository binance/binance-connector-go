/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetUmAccountDetailResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetUmAccountDetailResponse{}

// GetUmAccountDetailResponse struct for GetUmAccountDetailResponse
type GetUmAccountDetailResponse struct {
	Assets               []GetUmAccountDetailV2ResponseAssetsInner  `json:"assets,omitempty"`
	Positions            []GetUmAccountDetailResponsePositionsInner `json:"positions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetUmAccountDetailResponse GetUmAccountDetailResponse

// NewGetUmAccountDetailResponse instantiates a new GetUmAccountDetailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUmAccountDetailResponse() *GetUmAccountDetailResponse {
	this := GetUmAccountDetailResponse{}
	return &this
}

// NewGetUmAccountDetailResponseWithDefaults instantiates a new GetUmAccountDetailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUmAccountDetailResponseWithDefaults() *GetUmAccountDetailResponse {
	this := GetUmAccountDetailResponse{}
	return &this
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *GetUmAccountDetailResponse) GetAssets() []GetUmAccountDetailV2ResponseAssetsInner {
	if o == nil || common.IsNil(o.Assets) {
		var ret []GetUmAccountDetailV2ResponseAssetsInner
		return ret
	}
	return o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailResponse) GetAssetsOk() ([]GetUmAccountDetailV2ResponseAssetsInner, bool) {
	if o == nil || common.IsNil(o.Assets) {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *GetUmAccountDetailResponse) HasAssets() bool {
	if o != nil && !common.IsNil(o.Assets) {
		return true
	}

	return false
}

// SetAssets gets a reference to the given []GetUmAccountDetailV2ResponseAssetsInner and assigns it to the Assets field.
func (o *GetUmAccountDetailResponse) SetAssets(v []GetUmAccountDetailV2ResponseAssetsInner) {
	o.Assets = v
}

// GetPositions returns the Positions field value if set, zero value otherwise.
func (o *GetUmAccountDetailResponse) GetPositions() []GetUmAccountDetailResponsePositionsInner {
	if o == nil || common.IsNil(o.Positions) {
		var ret []GetUmAccountDetailResponsePositionsInner
		return ret
	}
	return o.Positions
}

// GetPositionsOk returns a tuple with the Positions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUmAccountDetailResponse) GetPositionsOk() ([]GetUmAccountDetailResponsePositionsInner, bool) {
	if o == nil || common.IsNil(o.Positions) {
		return nil, false
	}
	return o.Positions, true
}

// HasPositions returns a boolean if a field has been set.
func (o *GetUmAccountDetailResponse) HasPositions() bool {
	if o != nil && !common.IsNil(o.Positions) {
		return true
	}

	return false
}

// SetPositions gets a reference to the given []GetUmAccountDetailResponsePositionsInner and assigns it to the Positions field.
func (o *GetUmAccountDetailResponse) SetPositions(v []GetUmAccountDetailResponsePositionsInner) {
	o.Positions = v
}

func (o GetUmAccountDetailResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUmAccountDetailResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Assets) {
		toSerialize["assets"] = o.Assets
	}
	if !common.IsNil(o.Positions) {
		toSerialize["positions"] = o.Positions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetUmAccountDetailResponse) UnmarshalJSON(data []byte) (err error) {
	varGetUmAccountDetailResponse := _GetUmAccountDetailResponse{}

	err = json.Unmarshal(data, &varGetUmAccountDetailResponse)

	if err != nil {
		return err
	}

	*o = GetUmAccountDetailResponse(varGetUmAccountDetailResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "assets")
		delete(additionalProperties, "positions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetUmAccountDetailResponse struct {
	value *GetUmAccountDetailResponse
	isSet bool
}

func (v NullableGetUmAccountDetailResponse) Get() *GetUmAccountDetailResponse {
	return v.value
}

func (v *NullableGetUmAccountDetailResponse) Set(val *GetUmAccountDetailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUmAccountDetailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUmAccountDetailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUmAccountDetailResponse(val *GetUmAccountDetailResponse) *NullableGetUmAccountDetailResponse {
	return &NullableGetUmAccountDetailResponse{value: val, isSet: true}
}

func (v NullableGetUmAccountDetailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUmAccountDetailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
