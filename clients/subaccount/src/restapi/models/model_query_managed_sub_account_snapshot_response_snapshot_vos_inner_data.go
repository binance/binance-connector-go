/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData{}

// QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData struct for QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData
type QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData struct {
	Balances             []QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner   `json:"balances,omitempty"`
	TotalAssetOfBtc      *string                                                                     `json:"totalAssetOfBtc,omitempty"`
	MarginLevel          *string                                                                     `json:"marginLevel,omitempty"`
	TotalLiabilityOfBtc  *string                                                                     `json:"totalLiabilityOfBtc,omitempty"`
	TotalNetAssetOfBtc   *string                                                                     `json:"totalNetAssetOfBtc,omitempty"`
	UserAssets           []QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner `json:"userAssets,omitempty"`
	Assets               []QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner     `json:"assets,omitempty"`
	Position             []QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataPositionInner   `json:"position,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData

// NewQueryManagedSubAccountSnapshotResponseSnapshotVosInnerData instantiates a new QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryManagedSubAccountSnapshotResponseSnapshotVosInnerData() *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData {
	this := QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData{}
	return &this
}

// NewQueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataWithDefaults instantiates a new QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataWithDefaults() *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData {
	this := QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData{}
	return &this
}

// GetBalances returns the Balances field value if set, zero value otherwise.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) GetBalances() []QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner {
	if o == nil || common.IsNil(o.Balances) {
		var ret []QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner
		return ret
	}
	return o.Balances
}

// GetBalancesOk returns a tuple with the Balances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) GetBalancesOk() ([]QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner, bool) {
	if o == nil || common.IsNil(o.Balances) {
		return nil, false
	}
	return o.Balances, true
}

// HasBalances returns a boolean if a field has been set.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) HasBalances() bool {
	if o != nil && !common.IsNil(o.Balances) {
		return true
	}

	return false
}

// SetBalances gets a reference to the given []QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner and assigns it to the Balances field.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) SetBalances(v []QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner) {
	o.Balances = v
}

// GetTotalAssetOfBtc returns the TotalAssetOfBtc field value if set, zero value otherwise.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) GetTotalAssetOfBtc() string {
	if o == nil || common.IsNil(o.TotalAssetOfBtc) {
		var ret string
		return ret
	}
	return *o.TotalAssetOfBtc
}

// GetTotalAssetOfBtcOk returns a tuple with the TotalAssetOfBtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) GetTotalAssetOfBtcOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalAssetOfBtc) {
		return nil, false
	}
	return o.TotalAssetOfBtc, true
}

// HasTotalAssetOfBtc returns a boolean if a field has been set.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) HasTotalAssetOfBtc() bool {
	if o != nil && !common.IsNil(o.TotalAssetOfBtc) {
		return true
	}

	return false
}

// SetTotalAssetOfBtc gets a reference to the given string and assigns it to the TotalAssetOfBtc field.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) SetTotalAssetOfBtc(v string) {
	o.TotalAssetOfBtc = &v
}

// GetMarginLevel returns the MarginLevel field value if set, zero value otherwise.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) GetMarginLevel() string {
	if o == nil || common.IsNil(o.MarginLevel) {
		var ret string
		return ret
	}
	return *o.MarginLevel
}

// GetMarginLevelOk returns a tuple with the MarginLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) GetMarginLevelOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarginLevel) {
		return nil, false
	}
	return o.MarginLevel, true
}

// HasMarginLevel returns a boolean if a field has been set.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) HasMarginLevel() bool {
	if o != nil && !common.IsNil(o.MarginLevel) {
		return true
	}

	return false
}

// SetMarginLevel gets a reference to the given string and assigns it to the MarginLevel field.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) SetMarginLevel(v string) {
	o.MarginLevel = &v
}

// GetTotalLiabilityOfBtc returns the TotalLiabilityOfBtc field value if set, zero value otherwise.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) GetTotalLiabilityOfBtc() string {
	if o == nil || common.IsNil(o.TotalLiabilityOfBtc) {
		var ret string
		return ret
	}
	return *o.TotalLiabilityOfBtc
}

// GetTotalLiabilityOfBtcOk returns a tuple with the TotalLiabilityOfBtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) GetTotalLiabilityOfBtcOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalLiabilityOfBtc) {
		return nil, false
	}
	return o.TotalLiabilityOfBtc, true
}

// HasTotalLiabilityOfBtc returns a boolean if a field has been set.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) HasTotalLiabilityOfBtc() bool {
	if o != nil && !common.IsNil(o.TotalLiabilityOfBtc) {
		return true
	}

	return false
}

// SetTotalLiabilityOfBtc gets a reference to the given string and assigns it to the TotalLiabilityOfBtc field.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) SetTotalLiabilityOfBtc(v string) {
	o.TotalLiabilityOfBtc = &v
}

// GetTotalNetAssetOfBtc returns the TotalNetAssetOfBtc field value if set, zero value otherwise.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) GetTotalNetAssetOfBtc() string {
	if o == nil || common.IsNil(o.TotalNetAssetOfBtc) {
		var ret string
		return ret
	}
	return *o.TotalNetAssetOfBtc
}

// GetTotalNetAssetOfBtcOk returns a tuple with the TotalNetAssetOfBtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) GetTotalNetAssetOfBtcOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalNetAssetOfBtc) {
		return nil, false
	}
	return o.TotalNetAssetOfBtc, true
}

// HasTotalNetAssetOfBtc returns a boolean if a field has been set.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) HasTotalNetAssetOfBtc() bool {
	if o != nil && !common.IsNil(o.TotalNetAssetOfBtc) {
		return true
	}

	return false
}

// SetTotalNetAssetOfBtc gets a reference to the given string and assigns it to the TotalNetAssetOfBtc field.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) SetTotalNetAssetOfBtc(v string) {
	o.TotalNetAssetOfBtc = &v
}

// GetUserAssets returns the UserAssets field value if set, zero value otherwise.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) GetUserAssets() []QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner {
	if o == nil || common.IsNil(o.UserAssets) {
		var ret []QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner
		return ret
	}
	return o.UserAssets
}

// GetUserAssetsOk returns a tuple with the UserAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) GetUserAssetsOk() ([]QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner, bool) {
	if o == nil || common.IsNil(o.UserAssets) {
		return nil, false
	}
	return o.UserAssets, true
}

// HasUserAssets returns a boolean if a field has been set.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) HasUserAssets() bool {
	if o != nil && !common.IsNil(o.UserAssets) {
		return true
	}

	return false
}

// SetUserAssets gets a reference to the given []QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner and assigns it to the UserAssets field.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) SetUserAssets(v []QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner) {
	o.UserAssets = v
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) GetAssets() []QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner {
	if o == nil || common.IsNil(o.Assets) {
		var ret []QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner
		return ret
	}
	return o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) GetAssetsOk() ([]QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner, bool) {
	if o == nil || common.IsNil(o.Assets) {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) HasAssets() bool {
	if o != nil && !common.IsNil(o.Assets) {
		return true
	}

	return false
}

// SetAssets gets a reference to the given []QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner and assigns it to the Assets field.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) SetAssets(v []QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) {
	o.Assets = v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) GetPosition() []QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataPositionInner {
	if o == nil || common.IsNil(o.Position) {
		var ret []QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataPositionInner
		return ret
	}
	return o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) GetPositionOk() ([]QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataPositionInner, bool) {
	if o == nil || common.IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) HasPosition() bool {
	if o != nil && !common.IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given []QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataPositionInner and assigns it to the Position field.
func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) SetPosition(v []QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataPositionInner) {
	o.Position = v
}

func (o QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Balances) {
		toSerialize["balances"] = o.Balances
	}
	if !common.IsNil(o.TotalAssetOfBtc) {
		toSerialize["totalAssetOfBtc"] = o.TotalAssetOfBtc
	}
	if !common.IsNil(o.MarginLevel) {
		toSerialize["marginLevel"] = o.MarginLevel
	}
	if !common.IsNil(o.TotalLiabilityOfBtc) {
		toSerialize["totalLiabilityOfBtc"] = o.TotalLiabilityOfBtc
	}
	if !common.IsNil(o.TotalNetAssetOfBtc) {
		toSerialize["totalNetAssetOfBtc"] = o.TotalNetAssetOfBtc
	}
	if !common.IsNil(o.UserAssets) {
		toSerialize["userAssets"] = o.UserAssets
	}
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

func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) UnmarshalJSON(data []byte) (err error) {
	varQueryManagedSubAccountSnapshotResponseSnapshotVosInnerData := _QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData{}

	err = json.Unmarshal(data, &varQueryManagedSubAccountSnapshotResponseSnapshotVosInnerData)

	if err != nil {
		return err
	}

	*o = QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData(varQueryManagedSubAccountSnapshotResponseSnapshotVosInnerData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "balances")
		delete(additionalProperties, "totalAssetOfBtc")
		delete(additionalProperties, "marginLevel")
		delete(additionalProperties, "totalLiabilityOfBtc")
		delete(additionalProperties, "totalNetAssetOfBtc")
		delete(additionalProperties, "userAssets")
		delete(additionalProperties, "assets")
		delete(additionalProperties, "position")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryManagedSubAccountSnapshotResponseSnapshotVosInnerData struct {
	value *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData
	isSet bool
}

func (v NullableQueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) Get() *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData {
	return v.value
}

func (v *NullableQueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) Set(val *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryManagedSubAccountSnapshotResponseSnapshotVosInnerData(val *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) *NullableQueryManagedSubAccountSnapshotResponseSnapshotVosInnerData {
	return &NullableQueryManagedSubAccountSnapshotResponseSnapshotVosInnerData{value: val, isSet: true}
}

func (v NullableQueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
