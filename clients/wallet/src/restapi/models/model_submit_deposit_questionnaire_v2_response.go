/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the SubmitDepositQuestionnaireV2Response type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SubmitDepositQuestionnaireV2Response{}

// SubmitDepositQuestionnaireV2Response struct for SubmitDepositQuestionnaireV2Response
type SubmitDepositQuestionnaireV2Response struct {
	TrId                 *int64  `json:"trId,omitempty"`
	Accepted             *bool   `json:"accepted,omitempty"`
	Info                 *string `json:"info,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SubmitDepositQuestionnaireV2Response SubmitDepositQuestionnaireV2Response

// NewSubmitDepositQuestionnaireV2Response instantiates a new SubmitDepositQuestionnaireV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitDepositQuestionnaireV2Response() *SubmitDepositQuestionnaireV2Response {
	this := SubmitDepositQuestionnaireV2Response{}
	return &this
}

// NewSubmitDepositQuestionnaireV2ResponseWithDefaults instantiates a new SubmitDepositQuestionnaireV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitDepositQuestionnaireV2ResponseWithDefaults() *SubmitDepositQuestionnaireV2Response {
	this := SubmitDepositQuestionnaireV2Response{}
	return &this
}

// GetTrId returns the TrId field value if set, zero value otherwise.
func (o *SubmitDepositQuestionnaireV2Response) GetTrId() int64 {
	if o == nil || common.IsNil(o.TrId) {
		var ret int64
		return ret
	}
	return *o.TrId
}

// GetTrIdOk returns a tuple with the TrId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitDepositQuestionnaireV2Response) GetTrIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TrId) {
		return nil, false
	}
	return o.TrId, true
}

// HasTrId returns a boolean if a field has been set.
func (o *SubmitDepositQuestionnaireV2Response) HasTrId() bool {
	if o != nil && !common.IsNil(o.TrId) {
		return true
	}

	return false
}

// SetTrId gets a reference to the given int64 and assigns it to the TrId field.
func (o *SubmitDepositQuestionnaireV2Response) SetTrId(v int64) {
	o.TrId = &v
}

// GetAccepted returns the Accepted field value if set, zero value otherwise.
func (o *SubmitDepositQuestionnaireV2Response) GetAccepted() bool {
	if o == nil || common.IsNil(o.Accepted) {
		var ret bool
		return ret
	}
	return *o.Accepted
}

// GetAcceptedOk returns a tuple with the Accepted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitDepositQuestionnaireV2Response) GetAcceptedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Accepted) {
		return nil, false
	}
	return o.Accepted, true
}

// HasAccepted returns a boolean if a field has been set.
func (o *SubmitDepositQuestionnaireV2Response) HasAccepted() bool {
	if o != nil && !common.IsNil(o.Accepted) {
		return true
	}

	return false
}

// SetAccepted gets a reference to the given bool and assigns it to the Accepted field.
func (o *SubmitDepositQuestionnaireV2Response) SetAccepted(v bool) {
	o.Accepted = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *SubmitDepositQuestionnaireV2Response) GetInfo() string {
	if o == nil || common.IsNil(o.Info) {
		var ret string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitDepositQuestionnaireV2Response) GetInfoOk() (*string, bool) {
	if o == nil || common.IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *SubmitDepositQuestionnaireV2Response) HasInfo() bool {
	if o != nil && !common.IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given string and assigns it to the Info field.
func (o *SubmitDepositQuestionnaireV2Response) SetInfo(v string) {
	o.Info = &v
}

func (o SubmitDepositQuestionnaireV2Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubmitDepositQuestionnaireV2Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TrId) {
		toSerialize["trId"] = o.TrId
	}
	if !common.IsNil(o.Accepted) {
		toSerialize["accepted"] = o.Accepted
	}
	if !common.IsNil(o.Info) {
		toSerialize["info"] = o.Info
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SubmitDepositQuestionnaireV2Response) UnmarshalJSON(data []byte) (err error) {
	varSubmitDepositQuestionnaireV2Response := _SubmitDepositQuestionnaireV2Response{}

	err = json.Unmarshal(data, &varSubmitDepositQuestionnaireV2Response)

	if err != nil {
		return err
	}

	*o = SubmitDepositQuestionnaireV2Response(varSubmitDepositQuestionnaireV2Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "trId")
		delete(additionalProperties, "accepted")
		delete(additionalProperties, "info")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubmitDepositQuestionnaireV2Response struct {
	value *SubmitDepositQuestionnaireV2Response
	isSet bool
}

func (v NullableSubmitDepositQuestionnaireV2Response) Get() *SubmitDepositQuestionnaireV2Response {
	return v.value
}

func (v *NullableSubmitDepositQuestionnaireV2Response) Set(val *SubmitDepositQuestionnaireV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitDepositQuestionnaireV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitDepositQuestionnaireV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitDepositQuestionnaireV2Response(val *SubmitDepositQuestionnaireV2Response) *NullableSubmitDepositQuestionnaireV2Response {
	return &NullableSubmitDepositQuestionnaireV2Response{value: val, isSet: true}
}

func (v NullableSubmitDepositQuestionnaireV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmitDepositQuestionnaireV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
