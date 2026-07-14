/*
Wallet REST API

Query balances, manage assets, and perform wallet operations via the Binance Wallet API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetCountryListResponseCountriesInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetCountryListResponseCountriesInner{}

// GetCountryListResponseCountriesInner struct for GetCountryListResponseCountriesInner
type GetCountryListResponseCountriesInner struct {
	// ISO 2-digit country code, lowercase.
	CountryCode *string `json:"countryCode,omitempty"`
	// Country display name.
	CountryName *string `json:"countryName,omitempty"`
	// `supported`, `limited`, or `blocked`.
	BlockType *string `json:"blockType,omitempty"`
	// Whether deposit is allowed for this country.
	DepositAllowed *bool `json:"depositAllowed,omitempty"`
	// Whether withdrawal is allowed for this country.
	WithdrawalAllowed *bool `json:"withdrawalAllowed,omitempty"`
	// Whether this country has region-level restrictions.
	HasRegionRestrictions *bool `json:"hasRegionRestrictions,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _GetCountryListResponseCountriesInner GetCountryListResponseCountriesInner

// NewGetCountryListResponseCountriesInner instantiates a new GetCountryListResponseCountriesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCountryListResponseCountriesInner() *GetCountryListResponseCountriesInner {
	this := GetCountryListResponseCountriesInner{}
	return &this
}

// NewGetCountryListResponseCountriesInnerWithDefaults instantiates a new GetCountryListResponseCountriesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCountryListResponseCountriesInnerWithDefaults() *GetCountryListResponseCountriesInner {
	this := GetCountryListResponseCountriesInner{}
	return &this
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *GetCountryListResponseCountriesInner) GetCountryCode() string {
	if o == nil || common.IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCountryListResponseCountriesInner) GetCountryCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *GetCountryListResponseCountriesInner) HasCountryCode() bool {
	if o != nil && !common.IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *GetCountryListResponseCountriesInner) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetCountryName returns the CountryName field value if set, zero value otherwise.
func (o *GetCountryListResponseCountriesInner) GetCountryName() string {
	if o == nil || common.IsNil(o.CountryName) {
		var ret string
		return ret
	}
	return *o.CountryName
}

// GetCountryNameOk returns a tuple with the CountryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCountryListResponseCountriesInner) GetCountryNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.CountryName) {
		return nil, false
	}
	return o.CountryName, true
}

// HasCountryName returns a boolean if a field has been set.
func (o *GetCountryListResponseCountriesInner) HasCountryName() bool {
	if o != nil && !common.IsNil(o.CountryName) {
		return true
	}

	return false
}

// SetCountryName gets a reference to the given string and assigns it to the CountryName field.
func (o *GetCountryListResponseCountriesInner) SetCountryName(v string) {
	o.CountryName = &v
}

// GetBlockType returns the BlockType field value if set, zero value otherwise.
func (o *GetCountryListResponseCountriesInner) GetBlockType() string {
	if o == nil || common.IsNil(o.BlockType) {
		var ret string
		return ret
	}
	return *o.BlockType
}

// GetBlockTypeOk returns a tuple with the BlockType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCountryListResponseCountriesInner) GetBlockTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.BlockType) {
		return nil, false
	}
	return o.BlockType, true
}

// HasBlockType returns a boolean if a field has been set.
func (o *GetCountryListResponseCountriesInner) HasBlockType() bool {
	if o != nil && !common.IsNil(o.BlockType) {
		return true
	}

	return false
}

// SetBlockType gets a reference to the given string and assigns it to the BlockType field.
func (o *GetCountryListResponseCountriesInner) SetBlockType(v string) {
	o.BlockType = &v
}

// GetDepositAllowed returns the DepositAllowed field value if set, zero value otherwise.
func (o *GetCountryListResponseCountriesInner) GetDepositAllowed() bool {
	if o == nil || common.IsNil(o.DepositAllowed) {
		var ret bool
		return ret
	}
	return *o.DepositAllowed
}

// GetDepositAllowedOk returns a tuple with the DepositAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCountryListResponseCountriesInner) GetDepositAllowedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.DepositAllowed) {
		return nil, false
	}
	return o.DepositAllowed, true
}

// HasDepositAllowed returns a boolean if a field has been set.
func (o *GetCountryListResponseCountriesInner) HasDepositAllowed() bool {
	if o != nil && !common.IsNil(o.DepositAllowed) {
		return true
	}

	return false
}

// SetDepositAllowed gets a reference to the given bool and assigns it to the DepositAllowed field.
func (o *GetCountryListResponseCountriesInner) SetDepositAllowed(v bool) {
	o.DepositAllowed = &v
}

// GetWithdrawalAllowed returns the WithdrawalAllowed field value if set, zero value otherwise.
func (o *GetCountryListResponseCountriesInner) GetWithdrawalAllowed() bool {
	if o == nil || common.IsNil(o.WithdrawalAllowed) {
		var ret bool
		return ret
	}
	return *o.WithdrawalAllowed
}

// GetWithdrawalAllowedOk returns a tuple with the WithdrawalAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCountryListResponseCountriesInner) GetWithdrawalAllowedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.WithdrawalAllowed) {
		return nil, false
	}
	return o.WithdrawalAllowed, true
}

// HasWithdrawalAllowed returns a boolean if a field has been set.
func (o *GetCountryListResponseCountriesInner) HasWithdrawalAllowed() bool {
	if o != nil && !common.IsNil(o.WithdrawalAllowed) {
		return true
	}

	return false
}

// SetWithdrawalAllowed gets a reference to the given bool and assigns it to the WithdrawalAllowed field.
func (o *GetCountryListResponseCountriesInner) SetWithdrawalAllowed(v bool) {
	o.WithdrawalAllowed = &v
}

// GetHasRegionRestrictions returns the HasRegionRestrictions field value if set, zero value otherwise.
func (o *GetCountryListResponseCountriesInner) GetHasRegionRestrictions() bool {
	if o == nil || common.IsNil(o.HasRegionRestrictions) {
		var ret bool
		return ret
	}
	return *o.HasRegionRestrictions
}

// GetHasRegionRestrictionsOk returns a tuple with the HasRegionRestrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCountryListResponseCountriesInner) GetHasRegionRestrictionsOk() (*bool, bool) {
	if o == nil || common.IsNil(o.HasRegionRestrictions) {
		return nil, false
	}
	return o.HasRegionRestrictions, true
}

// HasHasRegionRestrictions returns a boolean if a field has been set.
func (o *GetCountryListResponseCountriesInner) HasHasRegionRestrictions() bool {
	if o != nil && !common.IsNil(o.HasRegionRestrictions) {
		return true
	}

	return false
}

// SetHasRegionRestrictions gets a reference to the given bool and assigns it to the HasRegionRestrictions field.
func (o *GetCountryListResponseCountriesInner) SetHasRegionRestrictions(v bool) {
	o.HasRegionRestrictions = &v
}

func (o GetCountryListResponseCountriesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCountryListResponseCountriesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	if !common.IsNil(o.CountryName) {
		toSerialize["countryName"] = o.CountryName
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
	if !common.IsNil(o.HasRegionRestrictions) {
		toSerialize["hasRegionRestrictions"] = o.HasRegionRestrictions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetCountryListResponseCountriesInner) UnmarshalJSON(data []byte) (err error) {
	varGetCountryListResponseCountriesInner := _GetCountryListResponseCountriesInner{}

	err = json.Unmarshal(data, &varGetCountryListResponseCountriesInner)

	if err != nil {
		return err
	}

	*o = GetCountryListResponseCountriesInner(varGetCountryListResponseCountriesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "countryCode")
		delete(additionalProperties, "countryName")
		delete(additionalProperties, "blockType")
		delete(additionalProperties, "depositAllowed")
		delete(additionalProperties, "withdrawalAllowed")
		delete(additionalProperties, "hasRegionRestrictions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetCountryListResponseCountriesInner struct {
	value *GetCountryListResponseCountriesInner
	isSet bool
}

func (v NullableGetCountryListResponseCountriesInner) Get() *GetCountryListResponseCountriesInner {
	return v.value
}

func (v *NullableGetCountryListResponseCountriesInner) Set(val *GetCountryListResponseCountriesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCountryListResponseCountriesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCountryListResponseCountriesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCountryListResponseCountriesInner(val *GetCountryListResponseCountriesInner) *NullableGetCountryListResponseCountriesInner {
	return &NullableGetCountryListResponseCountriesInner{value: val, isSet: true}
}

func (v NullableGetCountryListResponseCountriesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCountryListResponseCountriesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
