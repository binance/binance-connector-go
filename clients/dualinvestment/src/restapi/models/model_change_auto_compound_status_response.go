/*
Binance Dual Investment REST API

OpenAPI Specification for the Binance Dual Investment REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ChangeAutoCompoundStatusResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ChangeAutoCompoundStatusResponse{}

// ChangeAutoCompoundStatusResponse struct for ChangeAutoCompoundStatusResponse
type ChangeAutoCompoundStatusResponse struct {
	PositionId           *string `json:"positionId,omitempty"`
	AutoCompoundPlan     *string `json:"autoCompoundPlan,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChangeAutoCompoundStatusResponse ChangeAutoCompoundStatusResponse

// NewChangeAutoCompoundStatusResponse instantiates a new ChangeAutoCompoundStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeAutoCompoundStatusResponse() *ChangeAutoCompoundStatusResponse {
	this := ChangeAutoCompoundStatusResponse{}
	return &this
}

// NewChangeAutoCompoundStatusResponseWithDefaults instantiates a new ChangeAutoCompoundStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeAutoCompoundStatusResponseWithDefaults() *ChangeAutoCompoundStatusResponse {
	this := ChangeAutoCompoundStatusResponse{}
	return &this
}

// GetPositionId returns the PositionId field value if set, zero value otherwise.
func (o *ChangeAutoCompoundStatusResponse) GetPositionId() string {
	if o == nil || common.IsNil(o.PositionId) {
		var ret string
		return ret
	}
	return *o.PositionId
}

// GetPositionIdOk returns a tuple with the PositionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeAutoCompoundStatusResponse) GetPositionIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionId) {
		return nil, false
	}
	return o.PositionId, true
}

// HasPositionId returns a boolean if a field has been set.
func (o *ChangeAutoCompoundStatusResponse) HasPositionId() bool {
	if o != nil && !common.IsNil(o.PositionId) {
		return true
	}

	return false
}

// SetPositionId gets a reference to the given string and assigns it to the PositionId field.
func (o *ChangeAutoCompoundStatusResponse) SetPositionId(v string) {
	o.PositionId = &v
}

// GetAutoCompoundPlan returns the AutoCompoundPlan field value if set, zero value otherwise.
func (o *ChangeAutoCompoundStatusResponse) GetAutoCompoundPlan() string {
	if o == nil || common.IsNil(o.AutoCompoundPlan) {
		var ret string
		return ret
	}
	return *o.AutoCompoundPlan
}

// GetAutoCompoundPlanOk returns a tuple with the AutoCompoundPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeAutoCompoundStatusResponse) GetAutoCompoundPlanOk() (*string, bool) {
	if o == nil || common.IsNil(o.AutoCompoundPlan) {
		return nil, false
	}
	return o.AutoCompoundPlan, true
}

// HasAutoCompoundPlan returns a boolean if a field has been set.
func (o *ChangeAutoCompoundStatusResponse) HasAutoCompoundPlan() bool {
	if o != nil && !common.IsNil(o.AutoCompoundPlan) {
		return true
	}

	return false
}

// SetAutoCompoundPlan gets a reference to the given string and assigns it to the AutoCompoundPlan field.
func (o *ChangeAutoCompoundStatusResponse) SetAutoCompoundPlan(v string) {
	o.AutoCompoundPlan = &v
}

func (o ChangeAutoCompoundStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeAutoCompoundStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.PositionId) {
		toSerialize["positionId"] = o.PositionId
	}
	if !common.IsNil(o.AutoCompoundPlan) {
		toSerialize["autoCompoundPlan"] = o.AutoCompoundPlan
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ChangeAutoCompoundStatusResponse) UnmarshalJSON(data []byte) (err error) {
	varChangeAutoCompoundStatusResponse := _ChangeAutoCompoundStatusResponse{}

	err = json.Unmarshal(data, &varChangeAutoCompoundStatusResponse)

	if err != nil {
		return err
	}

	*o = ChangeAutoCompoundStatusResponse(varChangeAutoCompoundStatusResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "positionId")
		delete(additionalProperties, "autoCompoundPlan")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChangeAutoCompoundStatusResponse struct {
	value *ChangeAutoCompoundStatusResponse
	isSet bool
}

func (v NullableChangeAutoCompoundStatusResponse) Get() *ChangeAutoCompoundStatusResponse {
	return v.value
}

func (v *NullableChangeAutoCompoundStatusResponse) Set(val *ChangeAutoCompoundStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeAutoCompoundStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeAutoCompoundStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeAutoCompoundStatusResponse(val *ChangeAutoCompoundStatusResponse) *NullableChangeAutoCompoundStatusResponse {
	return &NullableChangeAutoCompoundStatusResponse{value: val, isSet: true}
}

func (v NullableChangeAutoCompoundStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeAutoCompoundStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
