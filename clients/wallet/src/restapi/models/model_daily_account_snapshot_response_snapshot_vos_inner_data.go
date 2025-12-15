/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the DailyAccountSnapshotResponseSnapshotVosInnerData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DailyAccountSnapshotResponseSnapshotVosInnerData{}

// DailyAccountSnapshotResponseSnapshotVosInnerData struct for DailyAccountSnapshotResponseSnapshotVosInnerData
type DailyAccountSnapshotResponseSnapshotVosInnerData struct {
	Balances             []DailyAccountSnapshotResponseSnapshotVosInnerDataBalancesInner   `json:"balances,omitempty"`
	TotalAssetOfBtc      *string                                                           `json:"totalAssetOfBtc,omitempty"`
	MarginLevel          *string                                                           `json:"marginLevel,omitempty"`
	TotalLiabilityOfBtc  *string                                                           `json:"totalLiabilityOfBtc,omitempty"`
	TotalNetAssetOfBtc   *string                                                           `json:"totalNetAssetOfBtc,omitempty"`
	UserAssets           []DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner `json:"userAssets,omitempty"`
	Assets               []DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner     `json:"assets,omitempty"`
	Position             []DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner   `json:"position,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DailyAccountSnapshotResponseSnapshotVosInnerData DailyAccountSnapshotResponseSnapshotVosInnerData

// NewDailyAccountSnapshotResponseSnapshotVosInnerData instantiates a new DailyAccountSnapshotResponseSnapshotVosInnerData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDailyAccountSnapshotResponseSnapshotVosInnerData() *DailyAccountSnapshotResponseSnapshotVosInnerData {
	this := DailyAccountSnapshotResponseSnapshotVosInnerData{}
	return &this
}

// NewDailyAccountSnapshotResponseSnapshotVosInnerDataWithDefaults instantiates a new DailyAccountSnapshotResponseSnapshotVosInnerData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDailyAccountSnapshotResponseSnapshotVosInnerDataWithDefaults() *DailyAccountSnapshotResponseSnapshotVosInnerData {
	this := DailyAccountSnapshotResponseSnapshotVosInnerData{}
	return &this
}

// GetBalances returns the Balances field value if set, zero value otherwise.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) GetBalances() []DailyAccountSnapshotResponseSnapshotVosInnerDataBalancesInner {
	if o == nil || common.IsNil(o.Balances) {
		var ret []DailyAccountSnapshotResponseSnapshotVosInnerDataBalancesInner
		return ret
	}
	return o.Balances
}

// GetBalancesOk returns a tuple with the Balances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) GetBalancesOk() ([]DailyAccountSnapshotResponseSnapshotVosInnerDataBalancesInner, bool) {
	if o == nil || common.IsNil(o.Balances) {
		return nil, false
	}
	return o.Balances, true
}

// HasBalances returns a boolean if a field has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) HasBalances() bool {
	if o != nil && !common.IsNil(o.Balances) {
		return true
	}

	return false
}

// SetBalances gets a reference to the given []DailyAccountSnapshotResponseSnapshotVosInnerDataBalancesInner and assigns it to the Balances field.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) SetBalances(v []DailyAccountSnapshotResponseSnapshotVosInnerDataBalancesInner) {
	o.Balances = v
}

// GetTotalAssetOfBtc returns the TotalAssetOfBtc field value if set, zero value otherwise.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) GetTotalAssetOfBtc() string {
	if o == nil || common.IsNil(o.TotalAssetOfBtc) {
		var ret string
		return ret
	}
	return *o.TotalAssetOfBtc
}

// GetTotalAssetOfBtcOk returns a tuple with the TotalAssetOfBtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) GetTotalAssetOfBtcOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalAssetOfBtc) {
		return nil, false
	}
	return o.TotalAssetOfBtc, true
}

// HasTotalAssetOfBtc returns a boolean if a field has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) HasTotalAssetOfBtc() bool {
	if o != nil && !common.IsNil(o.TotalAssetOfBtc) {
		return true
	}

	return false
}

// SetTotalAssetOfBtc gets a reference to the given string and assigns it to the TotalAssetOfBtc field.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) SetTotalAssetOfBtc(v string) {
	o.TotalAssetOfBtc = &v
}

// GetMarginLevel returns the MarginLevel field value if set, zero value otherwise.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) GetMarginLevel() string {
	if o == nil || common.IsNil(o.MarginLevel) {
		var ret string
		return ret
	}
	return *o.MarginLevel
}

// GetMarginLevelOk returns a tuple with the MarginLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) GetMarginLevelOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarginLevel) {
		return nil, false
	}
	return o.MarginLevel, true
}

// HasMarginLevel returns a boolean if a field has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) HasMarginLevel() bool {
	if o != nil && !common.IsNil(o.MarginLevel) {
		return true
	}

	return false
}

// SetMarginLevel gets a reference to the given string and assigns it to the MarginLevel field.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) SetMarginLevel(v string) {
	o.MarginLevel = &v
}

// GetTotalLiabilityOfBtc returns the TotalLiabilityOfBtc field value if set, zero value otherwise.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) GetTotalLiabilityOfBtc() string {
	if o == nil || common.IsNil(o.TotalLiabilityOfBtc) {
		var ret string
		return ret
	}
	return *o.TotalLiabilityOfBtc
}

// GetTotalLiabilityOfBtcOk returns a tuple with the TotalLiabilityOfBtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) GetTotalLiabilityOfBtcOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalLiabilityOfBtc) {
		return nil, false
	}
	return o.TotalLiabilityOfBtc, true
}

// HasTotalLiabilityOfBtc returns a boolean if a field has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) HasTotalLiabilityOfBtc() bool {
	if o != nil && !common.IsNil(o.TotalLiabilityOfBtc) {
		return true
	}

	return false
}

// SetTotalLiabilityOfBtc gets a reference to the given string and assigns it to the TotalLiabilityOfBtc field.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) SetTotalLiabilityOfBtc(v string) {
	o.TotalLiabilityOfBtc = &v
}

// GetTotalNetAssetOfBtc returns the TotalNetAssetOfBtc field value if set, zero value otherwise.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) GetTotalNetAssetOfBtc() string {
	if o == nil || common.IsNil(o.TotalNetAssetOfBtc) {
		var ret string
		return ret
	}
	return *o.TotalNetAssetOfBtc
}

// GetTotalNetAssetOfBtcOk returns a tuple with the TotalNetAssetOfBtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) GetTotalNetAssetOfBtcOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalNetAssetOfBtc) {
		return nil, false
	}
	return o.TotalNetAssetOfBtc, true
}

// HasTotalNetAssetOfBtc returns a boolean if a field has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) HasTotalNetAssetOfBtc() bool {
	if o != nil && !common.IsNil(o.TotalNetAssetOfBtc) {
		return true
	}

	return false
}

// SetTotalNetAssetOfBtc gets a reference to the given string and assigns it to the TotalNetAssetOfBtc field.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) SetTotalNetAssetOfBtc(v string) {
	o.TotalNetAssetOfBtc = &v
}

// GetUserAssets returns the UserAssets field value if set, zero value otherwise.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) GetUserAssets() []DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner {
	if o == nil || common.IsNil(o.UserAssets) {
		var ret []DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner
		return ret
	}
	return o.UserAssets
}

// GetUserAssetsOk returns a tuple with the UserAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) GetUserAssetsOk() ([]DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner, bool) {
	if o == nil || common.IsNil(o.UserAssets) {
		return nil, false
	}
	return o.UserAssets, true
}

// HasUserAssets returns a boolean if a field has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) HasUserAssets() bool {
	if o != nil && !common.IsNil(o.UserAssets) {
		return true
	}

	return false
}

// SetUserAssets gets a reference to the given []DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner and assigns it to the UserAssets field.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) SetUserAssets(v []DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner) {
	o.UserAssets = v
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) GetAssets() []DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner {
	if o == nil || common.IsNil(o.Assets) {
		var ret []DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner
		return ret
	}
	return o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) GetAssetsOk() ([]DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner, bool) {
	if o == nil || common.IsNil(o.Assets) {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) HasAssets() bool {
	if o != nil && !common.IsNil(o.Assets) {
		return true
	}

	return false
}

// SetAssets gets a reference to the given []DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner and assigns it to the Assets field.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) SetAssets(v []DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner) {
	o.Assets = v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) GetPosition() []DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner {
	if o == nil || common.IsNil(o.Position) {
		var ret []DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner
		return ret
	}
	return o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) GetPositionOk() ([]DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner, bool) {
	if o == nil || common.IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) HasPosition() bool {
	if o != nil && !common.IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given []DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner and assigns it to the Position field.
func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) SetPosition(v []DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner) {
	o.Position = v
}

func (o DailyAccountSnapshotResponseSnapshotVosInnerData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DailyAccountSnapshotResponseSnapshotVosInnerData) ToMap() (map[string]interface{}, error) {
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

func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) UnmarshalJSON(data []byte) (err error) {
	varDailyAccountSnapshotResponseSnapshotVosInnerData := _DailyAccountSnapshotResponseSnapshotVosInnerData{}

	err = json.Unmarshal(data, &varDailyAccountSnapshotResponseSnapshotVosInnerData)

	if err != nil {
		return err
	}

	*o = DailyAccountSnapshotResponseSnapshotVosInnerData(varDailyAccountSnapshotResponseSnapshotVosInnerData)

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

type NullableDailyAccountSnapshotResponseSnapshotVosInnerData struct {
	value *DailyAccountSnapshotResponseSnapshotVosInnerData
	isSet bool
}

func (v NullableDailyAccountSnapshotResponseSnapshotVosInnerData) Get() *DailyAccountSnapshotResponseSnapshotVosInnerData {
	return v.value
}

func (v *NullableDailyAccountSnapshotResponseSnapshotVosInnerData) Set(val *DailyAccountSnapshotResponseSnapshotVosInnerData) {
	v.value = val
	v.isSet = true
}

func (v NullableDailyAccountSnapshotResponseSnapshotVosInnerData) IsSet() bool {
	return v.isSet
}

func (v *NullableDailyAccountSnapshotResponseSnapshotVosInnerData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDailyAccountSnapshotResponseSnapshotVosInnerData(val *DailyAccountSnapshotResponseSnapshotVosInnerData) *NullableDailyAccountSnapshotResponseSnapshotVosInnerData {
	return &NullableDailyAccountSnapshotResponseSnapshotVosInnerData{value: val, isSet: true}
}

func (v NullableDailyAccountSnapshotResponseSnapshotVosInnerData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDailyAccountSnapshotResponseSnapshotVosInnerData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
