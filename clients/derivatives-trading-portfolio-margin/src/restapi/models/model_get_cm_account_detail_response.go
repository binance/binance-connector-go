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

// checks if the GetCmAccountDetailResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetCmAccountDetailResponse{}

// GetCmAccountDetailResponse struct for GetCmAccountDetailResponse
type GetCmAccountDetailResponse struct {
	Assets               []GetCmAccountDetailResponseAssetsInner    `json:"assets,omitempty"`
	Positions            []GetCmAccountDetailResponsePositionsInner `json:"positions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetCmAccountDetailResponse GetCmAccountDetailResponse

// NewGetCmAccountDetailResponse instantiates a new GetCmAccountDetailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCmAccountDetailResponse() *GetCmAccountDetailResponse {
	this := GetCmAccountDetailResponse{}
	return &this
}

// NewGetCmAccountDetailResponseWithDefaults instantiates a new GetCmAccountDetailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCmAccountDetailResponseWithDefaults() *GetCmAccountDetailResponse {
	this := GetCmAccountDetailResponse{}
	return &this
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *GetCmAccountDetailResponse) GetAssets() []GetCmAccountDetailResponseAssetsInner {
	if o == nil || common.IsNil(o.Assets) {
		var ret []GetCmAccountDetailResponseAssetsInner
		return ret
	}
	return o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCmAccountDetailResponse) GetAssetsOk() ([]GetCmAccountDetailResponseAssetsInner, bool) {
	if o == nil || common.IsNil(o.Assets) {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *GetCmAccountDetailResponse) HasAssets() bool {
	if o != nil && !common.IsNil(o.Assets) {
		return true
	}

	return false
}

// SetAssets gets a reference to the given []GetCmAccountDetailResponseAssetsInner and assigns it to the Assets field.
func (o *GetCmAccountDetailResponse) SetAssets(v []GetCmAccountDetailResponseAssetsInner) {
	o.Assets = v
}

// GetPositions returns the Positions field value if set, zero value otherwise.
func (o *GetCmAccountDetailResponse) GetPositions() []GetCmAccountDetailResponsePositionsInner {
	if o == nil || common.IsNil(o.Positions) {
		var ret []GetCmAccountDetailResponsePositionsInner
		return ret
	}
	return o.Positions
}

// GetPositionsOk returns a tuple with the Positions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCmAccountDetailResponse) GetPositionsOk() ([]GetCmAccountDetailResponsePositionsInner, bool) {
	if o == nil || common.IsNil(o.Positions) {
		return nil, false
	}
	return o.Positions, true
}

// HasPositions returns a boolean if a field has been set.
func (o *GetCmAccountDetailResponse) HasPositions() bool {
	if o != nil && !common.IsNil(o.Positions) {
		return true
	}

	return false
}

// SetPositions gets a reference to the given []GetCmAccountDetailResponsePositionsInner and assigns it to the Positions field.
func (o *GetCmAccountDetailResponse) SetPositions(v []GetCmAccountDetailResponsePositionsInner) {
	o.Positions = v
}

func (o GetCmAccountDetailResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCmAccountDetailResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetCmAccountDetailResponse) UnmarshalJSON(data []byte) (err error) {
	varGetCmAccountDetailResponse := _GetCmAccountDetailResponse{}

	err = json.Unmarshal(data, &varGetCmAccountDetailResponse)

	if err != nil {
		return err
	}

	*o = GetCmAccountDetailResponse(varGetCmAccountDetailResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "assets")
		delete(additionalProperties, "positions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetCmAccountDetailResponse struct {
	value *GetCmAccountDetailResponse
	isSet bool
}

func (v NullableGetCmAccountDetailResponse) Get() *GetCmAccountDetailResponse {
	return v.value
}

func (v *NullableGetCmAccountDetailResponse) Set(val *GetCmAccountDetailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCmAccountDetailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCmAccountDetailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCmAccountDetailResponse(val *GetCmAccountDetailResponse) *NullableGetCmAccountDetailResponse {
	return &NullableGetCmAccountDetailResponse{value: val, isSet: true}
}

func (v NullableGetCmAccountDetailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCmAccountDetailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
