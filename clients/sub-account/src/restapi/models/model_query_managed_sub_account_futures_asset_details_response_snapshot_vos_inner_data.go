/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData{}

// QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData struct for QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData
type QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData struct {
	Assets               []QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner   `json:"assets,omitempty"`
	Position             []QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner `json:"position,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData

// NewQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData instantiates a new QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData() *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData {
	this := QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData{}
	return &this
}

// NewQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataWithDefaults instantiates a new QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataWithDefaults() *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData {
	this := QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData{}
	return &this
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData) GetAssets() []QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner {
	if o == nil || common.IsNil(o.Assets) {
		var ret []QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner
		return ret
	}
	return o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData) GetAssetsOk() ([]QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner, bool) {
	if o == nil || common.IsNil(o.Assets) {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData) HasAssets() bool {
	if o != nil && !common.IsNil(o.Assets) {
		return true
	}

	return false
}

// SetAssets gets a reference to the given []QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner and assigns it to the Assets field.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData) SetAssets(v []QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataAssetsInner) {
	o.Assets = v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData) GetPosition() []QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner {
	if o == nil || common.IsNil(o.Position) {
		var ret []QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner
		return ret
	}
	return o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData) GetPositionOk() ([]QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner, bool) {
	if o == nil || common.IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData) HasPosition() bool {
	if o != nil && !common.IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given []QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner and assigns it to the Position field.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData) SetPosition(v []QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner) {
	o.Position = v
}

func (o QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Assets) {
		toSerialize["assets"] = o.Assets
	}
	if !common.IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData) UnmarshalJSON(data []byte) (err error) {
	varQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData := _QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData{}

	err = json.Unmarshal(data, &varQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData)

	if err != nil {
		return err
	}

	*o = QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData(varQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "assets")
		delete(additionalProperties, "position")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData struct {
	value *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData
	isSet bool
}

func (v NullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData) Get() *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData {
	return v.value
}

func (v *NullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData) Set(val *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData(val *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData) *NullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData {
	return &NullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData{value: val, isSet: true}
}

func (v NullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
