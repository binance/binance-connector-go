/*
Wallet REST API

Query balances, manage assets, and perform wallet operations via the Binance Wallet API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetSpotAssetTagsResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSpotAssetTagsResponseInner{}

// GetSpotAssetTagsResponseInner struct for GetSpotAssetTagsResponseInner
type GetSpotAssetTagsResponseInner struct {
	// Asset code.
	AssetCode *string `json:"assetCode,omitempty"`
	// Asset name.
	AssetName *string `json:"assetName,omitempty"`
	// Asset trading status. This endpoint only returns `true`.
	Trading *bool `json:"trading,omitempty"`
	// Tags configured for the asset. Returns an empty array when the asset has no tags.
	Tags                 []string `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetSpotAssetTagsResponseInner GetSpotAssetTagsResponseInner

// NewGetSpotAssetTagsResponseInner instantiates a new GetSpotAssetTagsResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSpotAssetTagsResponseInner() *GetSpotAssetTagsResponseInner {
	this := GetSpotAssetTagsResponseInner{}
	return &this
}

// NewGetSpotAssetTagsResponseInnerWithDefaults instantiates a new GetSpotAssetTagsResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSpotAssetTagsResponseInnerWithDefaults() *GetSpotAssetTagsResponseInner {
	this := GetSpotAssetTagsResponseInner{}
	return &this
}

// GetAssetCode returns the AssetCode field value if set, zero value otherwise.
func (o *GetSpotAssetTagsResponseInner) GetAssetCode() string {
	if o == nil || common.IsNil(o.AssetCode) {
		var ret string
		return ret
	}
	return *o.AssetCode
}

// GetAssetCodeOk returns a tuple with the AssetCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSpotAssetTagsResponseInner) GetAssetCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.AssetCode) {
		return nil, false
	}
	return o.AssetCode, true
}

// HasAssetCode returns a boolean if a field has been set.
func (o *GetSpotAssetTagsResponseInner) HasAssetCode() bool {
	if o != nil && !common.IsNil(o.AssetCode) {
		return true
	}

	return false
}

// SetAssetCode gets a reference to the given string and assigns it to the AssetCode field.
func (o *GetSpotAssetTagsResponseInner) SetAssetCode(v string) {
	o.AssetCode = &v
}

// GetAssetName returns the AssetName field value if set, zero value otherwise.
func (o *GetSpotAssetTagsResponseInner) GetAssetName() string {
	if o == nil || common.IsNil(o.AssetName) {
		var ret string
		return ret
	}
	return *o.AssetName
}

// GetAssetNameOk returns a tuple with the AssetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSpotAssetTagsResponseInner) GetAssetNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.AssetName) {
		return nil, false
	}
	return o.AssetName, true
}

// HasAssetName returns a boolean if a field has been set.
func (o *GetSpotAssetTagsResponseInner) HasAssetName() bool {
	if o != nil && !common.IsNil(o.AssetName) {
		return true
	}

	return false
}

// SetAssetName gets a reference to the given string and assigns it to the AssetName field.
func (o *GetSpotAssetTagsResponseInner) SetAssetName(v string) {
	o.AssetName = &v
}

// GetTrading returns the Trading field value if set, zero value otherwise.
func (o *GetSpotAssetTagsResponseInner) GetTrading() bool {
	if o == nil || common.IsNil(o.Trading) {
		var ret bool
		return ret
	}
	return *o.Trading
}

// GetTradingOk returns a tuple with the Trading field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSpotAssetTagsResponseInner) GetTradingOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Trading) {
		return nil, false
	}
	return o.Trading, true
}

// HasTrading returns a boolean if a field has been set.
func (o *GetSpotAssetTagsResponseInner) HasTrading() bool {
	if o != nil && !common.IsNil(o.Trading) {
		return true
	}

	return false
}

// SetTrading gets a reference to the given bool and assigns it to the Trading field.
func (o *GetSpotAssetTagsResponseInner) SetTrading(v bool) {
	o.Trading = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GetSpotAssetTagsResponseInner) GetTags() []string {
	if o == nil || common.IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSpotAssetTagsResponseInner) GetTagsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GetSpotAssetTagsResponseInner) HasTags() bool {
	if o != nil && !common.IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GetSpotAssetTagsResponseInner) SetTags(v []string) {
	o.Tags = v
}

func (o GetSpotAssetTagsResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSpotAssetTagsResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AssetCode) {
		toSerialize["assetCode"] = o.AssetCode
	}
	if !common.IsNil(o.AssetName) {
		toSerialize["assetName"] = o.AssetName
	}
	if !common.IsNil(o.Trading) {
		toSerialize["trading"] = o.Trading
	}
	if !common.IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetSpotAssetTagsResponseInner) UnmarshalJSON(data []byte) (err error) {
	varGetSpotAssetTagsResponseInner := _GetSpotAssetTagsResponseInner{}

	err = json.Unmarshal(data, &varGetSpotAssetTagsResponseInner)

	if err != nil {
		return err
	}

	*o = GetSpotAssetTagsResponseInner(varGetSpotAssetTagsResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "assetCode")
		delete(additionalProperties, "assetName")
		delete(additionalProperties, "trading")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSpotAssetTagsResponseInner struct {
	value *GetSpotAssetTagsResponseInner
	isSet bool
}

func (v NullableGetSpotAssetTagsResponseInner) Get() *GetSpotAssetTagsResponseInner {
	return v.value
}

func (v *NullableGetSpotAssetTagsResponseInner) Set(val *GetSpotAssetTagsResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSpotAssetTagsResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSpotAssetTagsResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSpotAssetTagsResponseInner(val *GetSpotAssetTagsResponseInner) *NullableGetSpotAssetTagsResponseInner {
	return &NullableGetSpotAssetTagsResponseInner{value: val, isSet: true}
}

func (v NullableGetSpotAssetTagsResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSpotAssetTagsResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
