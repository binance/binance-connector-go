/*
Simple Earn REST API

Earn rewards by subscribing to flexible or locked Simple Earn products.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetYieldArenaActivitiesResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetYieldArenaActivitiesResponse{}

// GetYieldArenaActivitiesResponse struct for GetYieldArenaActivitiesResponse
type GetYieldArenaActivitiesResponse struct {
	Activities           []GetYieldArenaActivitiesResponseActivitiesInner `json:"activities,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetYieldArenaActivitiesResponse GetYieldArenaActivitiesResponse

// NewGetYieldArenaActivitiesResponse instantiates a new GetYieldArenaActivitiesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetYieldArenaActivitiesResponse() *GetYieldArenaActivitiesResponse {
	this := GetYieldArenaActivitiesResponse{}
	return &this
}

// NewGetYieldArenaActivitiesResponseWithDefaults instantiates a new GetYieldArenaActivitiesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetYieldArenaActivitiesResponseWithDefaults() *GetYieldArenaActivitiesResponse {
	this := GetYieldArenaActivitiesResponse{}
	return &this
}

// GetActivities returns the Activities field value if set, zero value otherwise.
func (o *GetYieldArenaActivitiesResponse) GetActivities() []GetYieldArenaActivitiesResponseActivitiesInner {
	if o == nil || common.IsNil(o.Activities) {
		var ret []GetYieldArenaActivitiesResponseActivitiesInner
		return ret
	}
	return o.Activities
}

// GetActivitiesOk returns a tuple with the Activities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetYieldArenaActivitiesResponse) GetActivitiesOk() ([]GetYieldArenaActivitiesResponseActivitiesInner, bool) {
	if o == nil || common.IsNil(o.Activities) {
		return nil, false
	}
	return o.Activities, true
}

// HasActivities returns a boolean if a field has been set.
func (o *GetYieldArenaActivitiesResponse) HasActivities() bool {
	if o != nil && !common.IsNil(o.Activities) {
		return true
	}

	return false
}

// SetActivities gets a reference to the given []GetYieldArenaActivitiesResponseActivitiesInner and assigns it to the Activities field.
func (o *GetYieldArenaActivitiesResponse) SetActivities(v []GetYieldArenaActivitiesResponseActivitiesInner) {
	o.Activities = v
}

func (o GetYieldArenaActivitiesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetYieldArenaActivitiesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Activities) {
		toSerialize["activities"] = o.Activities
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetYieldArenaActivitiesResponse) UnmarshalJSON(data []byte) (err error) {
	varGetYieldArenaActivitiesResponse := _GetYieldArenaActivitiesResponse{}

	err = json.Unmarshal(data, &varGetYieldArenaActivitiesResponse)

	if err != nil {
		return err
	}

	*o = GetYieldArenaActivitiesResponse(varGetYieldArenaActivitiesResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "activities")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetYieldArenaActivitiesResponse struct {
	value *GetYieldArenaActivitiesResponse
	isSet bool
}

func (v NullableGetYieldArenaActivitiesResponse) Get() *GetYieldArenaActivitiesResponse {
	return v.value
}

func (v *NullableGetYieldArenaActivitiesResponse) Set(val *GetYieldArenaActivitiesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetYieldArenaActivitiesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetYieldArenaActivitiesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetYieldArenaActivitiesResponse(val *GetYieldArenaActivitiesResponse) *NullableGetYieldArenaActivitiesResponse {
	return &NullableGetYieldArenaActivitiesResponse{value: val, isSet: true}
}

func (v NullableGetYieldArenaActivitiesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetYieldArenaActivitiesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
