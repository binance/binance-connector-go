/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the CheckQuestionnaireRequirementsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CheckQuestionnaireRequirementsResponse{}

// CheckQuestionnaireRequirementsResponse struct for CheckQuestionnaireRequirementsResponse
type CheckQuestionnaireRequirementsResponse struct {
	QuestionnaireCountryCode *string `json:"questionnaireCountryCode,omitempty"`
	AdditionalProperties     map[string]interface{}
}

type _CheckQuestionnaireRequirementsResponse CheckQuestionnaireRequirementsResponse

// NewCheckQuestionnaireRequirementsResponse instantiates a new CheckQuestionnaireRequirementsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckQuestionnaireRequirementsResponse() *CheckQuestionnaireRequirementsResponse {
	this := CheckQuestionnaireRequirementsResponse{}
	return &this
}

// NewCheckQuestionnaireRequirementsResponseWithDefaults instantiates a new CheckQuestionnaireRequirementsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckQuestionnaireRequirementsResponseWithDefaults() *CheckQuestionnaireRequirementsResponse {
	this := CheckQuestionnaireRequirementsResponse{}
	return &this
}

// GetQuestionnaireCountryCode returns the QuestionnaireCountryCode field value if set, zero value otherwise.
func (o *CheckQuestionnaireRequirementsResponse) GetQuestionnaireCountryCode() string {
	if o == nil || common.IsNil(o.QuestionnaireCountryCode) {
		var ret string
		return ret
	}
	return *o.QuestionnaireCountryCode
}

// GetQuestionnaireCountryCodeOk returns a tuple with the QuestionnaireCountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckQuestionnaireRequirementsResponse) GetQuestionnaireCountryCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.QuestionnaireCountryCode) {
		return nil, false
	}
	return o.QuestionnaireCountryCode, true
}

// HasQuestionnaireCountryCode returns a boolean if a field has been set.
func (o *CheckQuestionnaireRequirementsResponse) HasQuestionnaireCountryCode() bool {
	if o != nil && !common.IsNil(o.QuestionnaireCountryCode) {
		return true
	}

	return false
}

// SetQuestionnaireCountryCode gets a reference to the given string and assigns it to the QuestionnaireCountryCode field.
func (o *CheckQuestionnaireRequirementsResponse) SetQuestionnaireCountryCode(v string) {
	o.QuestionnaireCountryCode = &v
}

func (o CheckQuestionnaireRequirementsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckQuestionnaireRequirementsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.QuestionnaireCountryCode) {
		toSerialize["questionnaireCountryCode"] = o.QuestionnaireCountryCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CheckQuestionnaireRequirementsResponse) UnmarshalJSON(data []byte) (err error) {
	varCheckQuestionnaireRequirementsResponse := _CheckQuestionnaireRequirementsResponse{}

	err = json.Unmarshal(data, &varCheckQuestionnaireRequirementsResponse)

	if err != nil {
		return err
	}

	*o = CheckQuestionnaireRequirementsResponse(varCheckQuestionnaireRequirementsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "questionnaireCountryCode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCheckQuestionnaireRequirementsResponse struct {
	value *CheckQuestionnaireRequirementsResponse
	isSet bool
}

func (v NullableCheckQuestionnaireRequirementsResponse) Get() *CheckQuestionnaireRequirementsResponse {
	return v.value
}

func (v *NullableCheckQuestionnaireRequirementsResponse) Set(val *CheckQuestionnaireRequirementsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckQuestionnaireRequirementsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckQuestionnaireRequirementsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckQuestionnaireRequirementsResponse(val *CheckQuestionnaireRequirementsResponse) *NullableCheckQuestionnaireRequirementsResponse {
	return &NullableCheckQuestionnaireRequirementsResponse{value: val, isSet: true}
}

func (v NullableCheckQuestionnaireRequirementsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckQuestionnaireRequirementsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
