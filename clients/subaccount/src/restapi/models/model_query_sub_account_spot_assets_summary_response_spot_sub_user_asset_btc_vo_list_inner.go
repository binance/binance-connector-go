/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner{}

// QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner struct for QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner
type QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner struct {
	Email                *string `json:"email,omitempty"`
	TotalAsset           *string `json:"totalAsset,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner

// NewQuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner instantiates a new QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner() *QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner {
	this := QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner{}
	return &this
}

// NewQuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInnerWithDefaults instantiates a new QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInnerWithDefaults() *QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner {
	this := QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner) GetEmail() string {
	if o == nil || common.IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner) GetEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner) HasEmail() bool {
	if o != nil && !common.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner) SetEmail(v string) {
	o.Email = &v
}

// GetTotalAsset returns the TotalAsset field value if set, zero value otherwise.
func (o *QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner) GetTotalAsset() string {
	if o == nil || common.IsNil(o.TotalAsset) {
		var ret string
		return ret
	}
	return *o.TotalAsset
}

// GetTotalAssetOk returns a tuple with the TotalAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner) GetTotalAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalAsset) {
		return nil, false
	}
	return o.TotalAsset, true
}

// HasTotalAsset returns a boolean if a field has been set.
func (o *QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner) HasTotalAsset() bool {
	if o != nil && !common.IsNil(o.TotalAsset) {
		return true
	}

	return false
}

// SetTotalAsset gets a reference to the given string and assigns it to the TotalAsset field.
func (o *QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner) SetTotalAsset(v string) {
	o.TotalAsset = &v
}

func (o QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !common.IsNil(o.TotalAsset) {
		toSerialize["totalAsset"] = o.TotalAsset
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner) UnmarshalJSON(data []byte) (err error) {
	varQuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner := _QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner{}

	err = json.Unmarshal(data, &varQuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner)

	if err != nil {
		return err
	}

	*o = QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner(varQuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		delete(additionalProperties, "totalAsset")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner struct {
	value *QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner
	isSet bool
}

func (v NullableQuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner) Get() *QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner {
	return v.value
}

func (v *NullableQuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner) Set(val *QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner(val *QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner) *NullableQuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner {
	return &NullableQuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner{value: val, isSet: true}
}

func (v NullableQuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
