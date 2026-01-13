/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the DepositHistoryV2ResponseInnerQuestionnaire type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DepositHistoryV2ResponseInnerQuestionnaire{}

// DepositHistoryV2ResponseInnerQuestionnaire struct for DepositHistoryV2ResponseInnerQuestionnaire
type DepositHistoryV2ResponseInnerQuestionnaire struct {
	VaspName             *string `json:"vaspName,omitempty"`
	DepositOriginator    *int64  `json:"depositOriginator,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DepositHistoryV2ResponseInnerQuestionnaire DepositHistoryV2ResponseInnerQuestionnaire

// NewDepositHistoryV2ResponseInnerQuestionnaire instantiates a new DepositHistoryV2ResponseInnerQuestionnaire object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepositHistoryV2ResponseInnerQuestionnaire() *DepositHistoryV2ResponseInnerQuestionnaire {
	this := DepositHistoryV2ResponseInnerQuestionnaire{}
	return &this
}

// NewDepositHistoryV2ResponseInnerQuestionnaireWithDefaults instantiates a new DepositHistoryV2ResponseInnerQuestionnaire object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositHistoryV2ResponseInnerQuestionnaireWithDefaults() *DepositHistoryV2ResponseInnerQuestionnaire {
	this := DepositHistoryV2ResponseInnerQuestionnaire{}
	return &this
}

// GetVaspName returns the VaspName field value if set, zero value otherwise.
func (o *DepositHistoryV2ResponseInnerQuestionnaire) GetVaspName() string {
	if o == nil || common.IsNil(o.VaspName) {
		var ret string
		return ret
	}
	return *o.VaspName
}

// GetVaspNameOk returns a tuple with the VaspName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistoryV2ResponseInnerQuestionnaire) GetVaspNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.VaspName) {
		return nil, false
	}
	return o.VaspName, true
}

// HasVaspName returns a boolean if a field has been set.
func (o *DepositHistoryV2ResponseInnerQuestionnaire) HasVaspName() bool {
	if o != nil && !common.IsNil(o.VaspName) {
		return true
	}

	return false
}

// SetVaspName gets a reference to the given string and assigns it to the VaspName field.
func (o *DepositHistoryV2ResponseInnerQuestionnaire) SetVaspName(v string) {
	o.VaspName = &v
}

// GetDepositOriginator returns the DepositOriginator field value if set, zero value otherwise.
func (o *DepositHistoryV2ResponseInnerQuestionnaire) GetDepositOriginator() int64 {
	if o == nil || common.IsNil(o.DepositOriginator) {
		var ret int64
		return ret
	}
	return *o.DepositOriginator
}

// GetDepositOriginatorOk returns a tuple with the DepositOriginator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistoryV2ResponseInnerQuestionnaire) GetDepositOriginatorOk() (*int64, bool) {
	if o == nil || common.IsNil(o.DepositOriginator) {
		return nil, false
	}
	return o.DepositOriginator, true
}

// HasDepositOriginator returns a boolean if a field has been set.
func (o *DepositHistoryV2ResponseInnerQuestionnaire) HasDepositOriginator() bool {
	if o != nil && !common.IsNil(o.DepositOriginator) {
		return true
	}

	return false
}

// SetDepositOriginator gets a reference to the given int64 and assigns it to the DepositOriginator field.
func (o *DepositHistoryV2ResponseInnerQuestionnaire) SetDepositOriginator(v int64) {
	o.DepositOriginator = &v
}

func (o DepositHistoryV2ResponseInnerQuestionnaire) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DepositHistoryV2ResponseInnerQuestionnaire) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.VaspName) {
		toSerialize["vaspName"] = o.VaspName
	}
	if !common.IsNil(o.DepositOriginator) {
		toSerialize["depositOriginator"] = o.DepositOriginator
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DepositHistoryV2ResponseInnerQuestionnaire) UnmarshalJSON(data []byte) (err error) {
	varDepositHistoryV2ResponseInnerQuestionnaire := _DepositHistoryV2ResponseInnerQuestionnaire{}

	err = json.Unmarshal(data, &varDepositHistoryV2ResponseInnerQuestionnaire)

	if err != nil {
		return err
	}

	*o = DepositHistoryV2ResponseInnerQuestionnaire(varDepositHistoryV2ResponseInnerQuestionnaire)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "vaspName")
		delete(additionalProperties, "depositOriginator")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDepositHistoryV2ResponseInnerQuestionnaire struct {
	value *DepositHistoryV2ResponseInnerQuestionnaire
	isSet bool
}

func (v NullableDepositHistoryV2ResponseInnerQuestionnaire) Get() *DepositHistoryV2ResponseInnerQuestionnaire {
	return v.value
}

func (v *NullableDepositHistoryV2ResponseInnerQuestionnaire) Set(val *DepositHistoryV2ResponseInnerQuestionnaire) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositHistoryV2ResponseInnerQuestionnaire) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositHistoryV2ResponseInnerQuestionnaire) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepositHistoryV2ResponseInnerQuestionnaire(val *DepositHistoryV2ResponseInnerQuestionnaire) *NullableDepositHistoryV2ResponseInnerQuestionnaire {
	return &NullableDepositHistoryV2ResponseInnerQuestionnaire{value: val, isSet: true}
}

func (v NullableDepositHistoryV2ResponseInnerQuestionnaire) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositHistoryV2ResponseInnerQuestionnaire) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
