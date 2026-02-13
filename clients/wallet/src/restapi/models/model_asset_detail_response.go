/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AssetDetailResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AssetDetailResponse{}

// AssetDetailResponse struct for AssetDetailResponse
type AssetDetailResponse struct {
	CTR                  *AssetDetailResponseCTR `json:"CTR,omitempty"`
	SKY                  *AssetDetailResponseSKY `json:"SKY,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetDetailResponse AssetDetailResponse

// NewAssetDetailResponse instantiates a new AssetDetailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetDetailResponse() *AssetDetailResponse {
	this := AssetDetailResponse{}
	return &this
}

// NewAssetDetailResponseWithDefaults instantiates a new AssetDetailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetDetailResponseWithDefaults() *AssetDetailResponse {
	this := AssetDetailResponse{}
	return &this
}

// GetCTR returns the CTR field value if set, zero value otherwise.
func (o *AssetDetailResponse) GetCTR() AssetDetailResponseCTR {
	if o == nil || common.IsNil(o.CTR) {
		var ret AssetDetailResponseCTR
		return ret
	}
	return *o.CTR
}

// GetCTROk returns a tuple with the CTR field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDetailResponse) GetCTROk() (*AssetDetailResponseCTR, bool) {
	if o == nil || common.IsNil(o.CTR) {
		return nil, false
	}
	return o.CTR, true
}

// HasCTR returns a boolean if a field has been set.
func (o *AssetDetailResponse) HasCTR() bool {
	if o != nil && !common.IsNil(o.CTR) {
		return true
	}

	return false
}

// SetCTR gets a reference to the given AssetDetailResponseCTR and assigns it to the CTR field.
func (o *AssetDetailResponse) SetCTR(v AssetDetailResponseCTR) {
	o.CTR = &v
}

// GetSKY returns the SKY field value if set, zero value otherwise.
func (o *AssetDetailResponse) GetSKY() AssetDetailResponseSKY {
	if o == nil || common.IsNil(o.SKY) {
		var ret AssetDetailResponseSKY
		return ret
	}
	return *o.SKY
}

// GetSKYOk returns a tuple with the SKY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDetailResponse) GetSKYOk() (*AssetDetailResponseSKY, bool) {
	if o == nil || common.IsNil(o.SKY) {
		return nil, false
	}
	return o.SKY, true
}

// HasSKY returns a boolean if a field has been set.
func (o *AssetDetailResponse) HasSKY() bool {
	if o != nil && !common.IsNil(o.SKY) {
		return true
	}

	return false
}

// SetSKY gets a reference to the given AssetDetailResponseSKY and assigns it to the SKY field.
func (o *AssetDetailResponse) SetSKY(v AssetDetailResponseSKY) {
	o.SKY = &v
}

func (o AssetDetailResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetDetailResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CTR) {
		toSerialize["CTR"] = o.CTR
	}
	if !common.IsNil(o.SKY) {
		toSerialize["SKY"] = o.SKY
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AssetDetailResponse) UnmarshalJSON(data []byte) (err error) {
	varAssetDetailResponse := _AssetDetailResponse{}

	err = json.Unmarshal(data, &varAssetDetailResponse)

	if err != nil {
		return err
	}

	*o = AssetDetailResponse(varAssetDetailResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "CTR")
		delete(additionalProperties, "SKY")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetDetailResponse struct {
	value *AssetDetailResponse
	isSet bool
}

func (v NullableAssetDetailResponse) Get() *AssetDetailResponse {
	return v.value
}

func (v *NullableAssetDetailResponse) Set(val *AssetDetailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetDetailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetDetailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetDetailResponse(val *AssetDetailResponse) *NullableAssetDetailResponse {
	return &NullableAssetDetailResponse{value: val, isSet: true}
}

func (v NullableAssetDetailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetDetailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
