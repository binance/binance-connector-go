/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OptionAccountInformationResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OptionAccountInformationResponse{}

// OptionAccountInformationResponse struct for OptionAccountInformationResponse
type OptionAccountInformationResponse struct {
	Asset                []OptionAccountInformationResponseAssetInner `json:"asset,omitempty"`
	Greek                []OptionAccountInformationResponseGreekInner `json:"greek,omitempty"`
	Time                 *int64                                       `json:"time,omitempty"`
	RiskLevel            *string                                      `json:"riskLevel,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OptionAccountInformationResponse OptionAccountInformationResponse

// NewOptionAccountInformationResponse instantiates a new OptionAccountInformationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionAccountInformationResponse() *OptionAccountInformationResponse {
	this := OptionAccountInformationResponse{}
	return &this
}

// NewOptionAccountInformationResponseWithDefaults instantiates a new OptionAccountInformationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionAccountInformationResponseWithDefaults() *OptionAccountInformationResponse {
	this := OptionAccountInformationResponse{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *OptionAccountInformationResponse) GetAsset() []OptionAccountInformationResponseAssetInner {
	if o == nil || common.IsNil(o.Asset) {
		var ret []OptionAccountInformationResponseAssetInner
		return ret
	}
	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionAccountInformationResponse) GetAssetOk() ([]OptionAccountInformationResponseAssetInner, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *OptionAccountInformationResponse) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given []OptionAccountInformationResponseAssetInner and assigns it to the Asset field.
func (o *OptionAccountInformationResponse) SetAsset(v []OptionAccountInformationResponseAssetInner) {
	o.Asset = v
}

// GetGreek returns the Greek field value if set, zero value otherwise.
func (o *OptionAccountInformationResponse) GetGreek() []OptionAccountInformationResponseGreekInner {
	if o == nil || common.IsNil(o.Greek) {
		var ret []OptionAccountInformationResponseGreekInner
		return ret
	}
	return o.Greek
}

// GetGreekOk returns a tuple with the Greek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionAccountInformationResponse) GetGreekOk() ([]OptionAccountInformationResponseGreekInner, bool) {
	if o == nil || common.IsNil(o.Greek) {
		return nil, false
	}
	return o.Greek, true
}

// HasGreek returns a boolean if a field has been set.
func (o *OptionAccountInformationResponse) HasGreek() bool {
	if o != nil && !common.IsNil(o.Greek) {
		return true
	}

	return false
}

// SetGreek gets a reference to the given []OptionAccountInformationResponseGreekInner and assigns it to the Greek field.
func (o *OptionAccountInformationResponse) SetGreek(v []OptionAccountInformationResponseGreekInner) {
	o.Greek = v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *OptionAccountInformationResponse) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionAccountInformationResponse) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *OptionAccountInformationResponse) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *OptionAccountInformationResponse) SetTime(v int64) {
	o.Time = &v
}

// GetRiskLevel returns the RiskLevel field value if set, zero value otherwise.
func (o *OptionAccountInformationResponse) GetRiskLevel() string {
	if o == nil || common.IsNil(o.RiskLevel) {
		var ret string
		return ret
	}
	return *o.RiskLevel
}

// GetRiskLevelOk returns a tuple with the RiskLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionAccountInformationResponse) GetRiskLevelOk() (*string, bool) {
	if o == nil || common.IsNil(o.RiskLevel) {
		return nil, false
	}
	return o.RiskLevel, true
}

// HasRiskLevel returns a boolean if a field has been set.
func (o *OptionAccountInformationResponse) HasRiskLevel() bool {
	if o != nil && !common.IsNil(o.RiskLevel) {
		return true
	}

	return false
}

// SetRiskLevel gets a reference to the given string and assigns it to the RiskLevel field.
func (o *OptionAccountInformationResponse) SetRiskLevel(v string) {
	o.RiskLevel = &v
}

func (o OptionAccountInformationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptionAccountInformationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Greek) {
		toSerialize["greek"] = o.Greek
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !common.IsNil(o.RiskLevel) {
		toSerialize["riskLevel"] = o.RiskLevel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OptionAccountInformationResponse) UnmarshalJSON(data []byte) (err error) {
	varOptionAccountInformationResponse := _OptionAccountInformationResponse{}

	err = json.Unmarshal(data, &varOptionAccountInformationResponse)

	if err != nil {
		return err
	}

	*o = OptionAccountInformationResponse(varOptionAccountInformationResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "greek")
		delete(additionalProperties, "time")
		delete(additionalProperties, "riskLevel")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOptionAccountInformationResponse struct {
	value *OptionAccountInformationResponse
	isSet bool
}

func (v NullableOptionAccountInformationResponse) Get() *OptionAccountInformationResponse {
	return v.value
}

func (v *NullableOptionAccountInformationResponse) Set(val *OptionAccountInformationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionAccountInformationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionAccountInformationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionAccountInformationResponse(val *OptionAccountInformationResponse) *NullableOptionAccountInformationResponse {
	return &NullableOptionAccountInformationResponse{value: val, isSet: true}
}

func (v NullableOptionAccountInformationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionAccountInformationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
