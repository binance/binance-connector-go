/*
Wallet REST API

Query balances, manage assets, and perform wallet operations via the Binance Wallet API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetRegionListResponseRegionsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetRegionListResponseRegionsInner{}

// GetRegionListResponseRegionsInner struct for GetRegionListResponseRegionsInner
type GetRegionListResponseRegionsInner struct {
	// Region/city display name (use this value in questionnaire answers).
	RegionName *string `json:"regionName,omitempty"`
	// `supported`, `limited`, or `blocked`.
	BlockType *string `json:"blockType,omitempty"`
	// Whether deposit is allowed for this region.
	DepositAllowed *bool `json:"depositAllowed,omitempty"`
	// Whether withdrawal is allowed for this region.
	WithdrawalAllowed    *bool `json:"withdrawalAllowed,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetRegionListResponseRegionsInner GetRegionListResponseRegionsInner

// NewGetRegionListResponseRegionsInner instantiates a new GetRegionListResponseRegionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRegionListResponseRegionsInner() *GetRegionListResponseRegionsInner {
	this := GetRegionListResponseRegionsInner{}
	return &this
}

// NewGetRegionListResponseRegionsInnerWithDefaults instantiates a new GetRegionListResponseRegionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRegionListResponseRegionsInnerWithDefaults() *GetRegionListResponseRegionsInner {
	this := GetRegionListResponseRegionsInner{}
	return &this
}

// GetRegionName returns the RegionName field value if set, zero value otherwise.
func (o *GetRegionListResponseRegionsInner) GetRegionName() string {
	if o == nil || common.IsNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRegionListResponseRegionsInner) GetRegionNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.RegionName) {
		return nil, false
	}
	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *GetRegionListResponseRegionsInner) HasRegionName() bool {
	if o != nil && !common.IsNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *GetRegionListResponseRegionsInner) SetRegionName(v string) {
	o.RegionName = &v
}

// GetBlockType returns the BlockType field value if set, zero value otherwise.
func (o *GetRegionListResponseRegionsInner) GetBlockType() string {
	if o == nil || common.IsNil(o.BlockType) {
		var ret string
		return ret
	}
	return *o.BlockType
}

// GetBlockTypeOk returns a tuple with the BlockType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRegionListResponseRegionsInner) GetBlockTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.BlockType) {
		return nil, false
	}
	return o.BlockType, true
}

// HasBlockType returns a boolean if a field has been set.
func (o *GetRegionListResponseRegionsInner) HasBlockType() bool {
	if o != nil && !common.IsNil(o.BlockType) {
		return true
	}

	return false
}

// SetBlockType gets a reference to the given string and assigns it to the BlockType field.
func (o *GetRegionListResponseRegionsInner) SetBlockType(v string) {
	o.BlockType = &v
}

// GetDepositAllowed returns the DepositAllowed field value if set, zero value otherwise.
func (o *GetRegionListResponseRegionsInner) GetDepositAllowed() bool {
	if o == nil || common.IsNil(o.DepositAllowed) {
		var ret bool
		return ret
	}
	return *o.DepositAllowed
}

// GetDepositAllowedOk returns a tuple with the DepositAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRegionListResponseRegionsInner) GetDepositAllowedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.DepositAllowed) {
		return nil, false
	}
	return o.DepositAllowed, true
}

// HasDepositAllowed returns a boolean if a field has been set.
func (o *GetRegionListResponseRegionsInner) HasDepositAllowed() bool {
	if o != nil && !common.IsNil(o.DepositAllowed) {
		return true
	}

	return false
}

// SetDepositAllowed gets a reference to the given bool and assigns it to the DepositAllowed field.
func (o *GetRegionListResponseRegionsInner) SetDepositAllowed(v bool) {
	o.DepositAllowed = &v
}

// GetWithdrawalAllowed returns the WithdrawalAllowed field value if set, zero value otherwise.
func (o *GetRegionListResponseRegionsInner) GetWithdrawalAllowed() bool {
	if o == nil || common.IsNil(o.WithdrawalAllowed) {
		var ret bool
		return ret
	}
	return *o.WithdrawalAllowed
}

// GetWithdrawalAllowedOk returns a tuple with the WithdrawalAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRegionListResponseRegionsInner) GetWithdrawalAllowedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.WithdrawalAllowed) {
		return nil, false
	}
	return o.WithdrawalAllowed, true
}

// HasWithdrawalAllowed returns a boolean if a field has been set.
func (o *GetRegionListResponseRegionsInner) HasWithdrawalAllowed() bool {
	if o != nil && !common.IsNil(o.WithdrawalAllowed) {
		return true
	}

	return false
}

// SetWithdrawalAllowed gets a reference to the given bool and assigns it to the WithdrawalAllowed field.
func (o *GetRegionListResponseRegionsInner) SetWithdrawalAllowed(v bool) {
	o.WithdrawalAllowed = &v
}

func (o GetRegionListResponseRegionsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRegionListResponseRegionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.RegionName) {
		toSerialize["regionName"] = o.RegionName
	}
	if !common.IsNil(o.BlockType) {
		toSerialize["blockType"] = o.BlockType
	}
	if !common.IsNil(o.DepositAllowed) {
		toSerialize["depositAllowed"] = o.DepositAllowed
	}
	if !common.IsNil(o.WithdrawalAllowed) {
		toSerialize["withdrawalAllowed"] = o.WithdrawalAllowed
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetRegionListResponseRegionsInner) UnmarshalJSON(data []byte) (err error) {
	varGetRegionListResponseRegionsInner := _GetRegionListResponseRegionsInner{}

	err = json.Unmarshal(data, &varGetRegionListResponseRegionsInner)

	if err != nil {
		return err
	}

	*o = GetRegionListResponseRegionsInner(varGetRegionListResponseRegionsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "regionName")
		delete(additionalProperties, "blockType")
		delete(additionalProperties, "depositAllowed")
		delete(additionalProperties, "withdrawalAllowed")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetRegionListResponseRegionsInner struct {
	value *GetRegionListResponseRegionsInner
	isSet bool
}

func (v NullableGetRegionListResponseRegionsInner) Get() *GetRegionListResponseRegionsInner {
	return v.value
}

func (v *NullableGetRegionListResponseRegionsInner) Set(val *GetRegionListResponseRegionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRegionListResponseRegionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRegionListResponseRegionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRegionListResponseRegionsInner(val *GetRegionListResponseRegionsInner) *NullableGetRegionListResponseRegionsInner {
	return &NullableGetRegionListResponseRegionsInner{value: val, isSet: true}
}

func (v NullableGetRegionListResponseRegionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRegionListResponseRegionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
