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

// checks if the GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner{}

// GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner struct for GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner
type GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner struct {
	Email                *string `json:"email,omitempty"`
	TotalAssetOfBtc      *string `json:"totalAssetOfBtc,omitempty"`
	TotalLiabilityOfBtc  *string `json:"totalLiabilityOfBtc,omitempty"`
	TotalNetAssetOfBtc   *string `json:"totalNetAssetOfBtc,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner

// NewGetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner instantiates a new GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner() *GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner {
	this := GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner{}
	return &this
}

// NewGetSummaryOfSubAccountsMarginAccountResponseSubAccountListInnerWithDefaults instantiates a new GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSummaryOfSubAccountsMarginAccountResponseSubAccountListInnerWithDefaults() *GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner {
	this := GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner) GetEmail() string {
	if o == nil || common.IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner) GetEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner) HasEmail() bool {
	if o != nil && !common.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner) SetEmail(v string) {
	o.Email = &v
}

// GetTotalAssetOfBtc returns the TotalAssetOfBtc field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner) GetTotalAssetOfBtc() string {
	if o == nil || common.IsNil(o.TotalAssetOfBtc) {
		var ret string
		return ret
	}
	return *o.TotalAssetOfBtc
}

// GetTotalAssetOfBtcOk returns a tuple with the TotalAssetOfBtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner) GetTotalAssetOfBtcOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalAssetOfBtc) {
		return nil, false
	}
	return o.TotalAssetOfBtc, true
}

// HasTotalAssetOfBtc returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner) HasTotalAssetOfBtc() bool {
	if o != nil && !common.IsNil(o.TotalAssetOfBtc) {
		return true
	}

	return false
}

// SetTotalAssetOfBtc gets a reference to the given string and assigns it to the TotalAssetOfBtc field.
func (o *GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner) SetTotalAssetOfBtc(v string) {
	o.TotalAssetOfBtc = &v
}

// GetTotalLiabilityOfBtc returns the TotalLiabilityOfBtc field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner) GetTotalLiabilityOfBtc() string {
	if o == nil || common.IsNil(o.TotalLiabilityOfBtc) {
		var ret string
		return ret
	}
	return *o.TotalLiabilityOfBtc
}

// GetTotalLiabilityOfBtcOk returns a tuple with the TotalLiabilityOfBtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner) GetTotalLiabilityOfBtcOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalLiabilityOfBtc) {
		return nil, false
	}
	return o.TotalLiabilityOfBtc, true
}

// HasTotalLiabilityOfBtc returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner) HasTotalLiabilityOfBtc() bool {
	if o != nil && !common.IsNil(o.TotalLiabilityOfBtc) {
		return true
	}

	return false
}

// SetTotalLiabilityOfBtc gets a reference to the given string and assigns it to the TotalLiabilityOfBtc field.
func (o *GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner) SetTotalLiabilityOfBtc(v string) {
	o.TotalLiabilityOfBtc = &v
}

// GetTotalNetAssetOfBtc returns the TotalNetAssetOfBtc field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner) GetTotalNetAssetOfBtc() string {
	if o == nil || common.IsNil(o.TotalNetAssetOfBtc) {
		var ret string
		return ret
	}
	return *o.TotalNetAssetOfBtc
}

// GetTotalNetAssetOfBtcOk returns a tuple with the TotalNetAssetOfBtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner) GetTotalNetAssetOfBtcOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalNetAssetOfBtc) {
		return nil, false
	}
	return o.TotalNetAssetOfBtc, true
}

// HasTotalNetAssetOfBtc returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner) HasTotalNetAssetOfBtc() bool {
	if o != nil && !common.IsNil(o.TotalNetAssetOfBtc) {
		return true
	}

	return false
}

// SetTotalNetAssetOfBtc gets a reference to the given string and assigns it to the TotalNetAssetOfBtc field.
func (o *GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner) SetTotalNetAssetOfBtc(v string) {
	o.TotalNetAssetOfBtc = &v
}

func (o GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !common.IsNil(o.TotalAssetOfBtc) {
		toSerialize["totalAssetOfBtc"] = o.TotalAssetOfBtc
	}
	if !common.IsNil(o.TotalLiabilityOfBtc) {
		toSerialize["totalLiabilityOfBtc"] = o.TotalLiabilityOfBtc
	}
	if !common.IsNil(o.TotalNetAssetOfBtc) {
		toSerialize["totalNetAssetOfBtc"] = o.TotalNetAssetOfBtc
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner) UnmarshalJSON(data []byte) (err error) {
	varGetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner := _GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner{}

	err = json.Unmarshal(data, &varGetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner)

	if err != nil {
		return err
	}

	*o = GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner(varGetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		delete(additionalProperties, "totalAssetOfBtc")
		delete(additionalProperties, "totalLiabilityOfBtc")
		delete(additionalProperties, "totalNetAssetOfBtc")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner struct {
	value *GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner
	isSet bool
}

func (v NullableGetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner) Get() *GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner {
	return v.value
}

func (v *NullableGetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner) Set(val *GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner(val *GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner) *NullableGetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner {
	return &NullableGetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner{value: val, isSet: true}
}

func (v NullableGetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
