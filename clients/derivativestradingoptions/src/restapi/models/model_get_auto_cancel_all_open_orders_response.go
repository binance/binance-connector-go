/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetAutoCancelAllOpenOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetAutoCancelAllOpenOrdersResponse{}

// GetAutoCancelAllOpenOrdersResponse struct for GetAutoCancelAllOpenOrdersResponse
type GetAutoCancelAllOpenOrdersResponse struct {
	Underlying           *string `json:"underlying,omitempty"`
	CountdownTime        *int64  `json:"countdownTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetAutoCancelAllOpenOrdersResponse GetAutoCancelAllOpenOrdersResponse

// NewGetAutoCancelAllOpenOrdersResponse instantiates a new GetAutoCancelAllOpenOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAutoCancelAllOpenOrdersResponse() *GetAutoCancelAllOpenOrdersResponse {
	this := GetAutoCancelAllOpenOrdersResponse{}
	return &this
}

// NewGetAutoCancelAllOpenOrdersResponseWithDefaults instantiates a new GetAutoCancelAllOpenOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAutoCancelAllOpenOrdersResponseWithDefaults() *GetAutoCancelAllOpenOrdersResponse {
	this := GetAutoCancelAllOpenOrdersResponse{}
	return &this
}

// GetUnderlying returns the Underlying field value if set, zero value otherwise.
func (o *GetAutoCancelAllOpenOrdersResponse) GetUnderlying() string {
	if o == nil || common.IsNil(o.Underlying) {
		var ret string
		return ret
	}
	return *o.Underlying
}

// GetUnderlyingOk returns a tuple with the Underlying field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAutoCancelAllOpenOrdersResponse) GetUnderlyingOk() (*string, bool) {
	if o == nil || common.IsNil(o.Underlying) {
		return nil, false
	}
	return o.Underlying, true
}

// HasUnderlying returns a boolean if a field has been set.
func (o *GetAutoCancelAllOpenOrdersResponse) HasUnderlying() bool {
	if o != nil && !common.IsNil(o.Underlying) {
		return true
	}

	return false
}

// SetUnderlying gets a reference to the given string and assigns it to the Underlying field.
func (o *GetAutoCancelAllOpenOrdersResponse) SetUnderlying(v string) {
	o.Underlying = &v
}

// GetCountdownTime returns the CountdownTime field value if set, zero value otherwise.
func (o *GetAutoCancelAllOpenOrdersResponse) GetCountdownTime() int64 {
	if o == nil || common.IsNil(o.CountdownTime) {
		var ret int64
		return ret
	}
	return *o.CountdownTime
}

// GetCountdownTimeOk returns a tuple with the CountdownTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAutoCancelAllOpenOrdersResponse) GetCountdownTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CountdownTime) {
		return nil, false
	}
	return o.CountdownTime, true
}

// HasCountdownTime returns a boolean if a field has been set.
func (o *GetAutoCancelAllOpenOrdersResponse) HasCountdownTime() bool {
	if o != nil && !common.IsNil(o.CountdownTime) {
		return true
	}

	return false
}

// SetCountdownTime gets a reference to the given int64 and assigns it to the CountdownTime field.
func (o *GetAutoCancelAllOpenOrdersResponse) SetCountdownTime(v int64) {
	o.CountdownTime = &v
}

func (o GetAutoCancelAllOpenOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAutoCancelAllOpenOrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Underlying) {
		toSerialize["underlying"] = o.Underlying
	}
	if !common.IsNil(o.CountdownTime) {
		toSerialize["countdownTime"] = o.CountdownTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetAutoCancelAllOpenOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	varGetAutoCancelAllOpenOrdersResponse := _GetAutoCancelAllOpenOrdersResponse{}

	err = json.Unmarshal(data, &varGetAutoCancelAllOpenOrdersResponse)

	if err != nil {
		return err
	}

	*o = GetAutoCancelAllOpenOrdersResponse(varGetAutoCancelAllOpenOrdersResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "underlying")
		delete(additionalProperties, "countdownTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetAutoCancelAllOpenOrdersResponse struct {
	value *GetAutoCancelAllOpenOrdersResponse
	isSet bool
}

func (v NullableGetAutoCancelAllOpenOrdersResponse) Get() *GetAutoCancelAllOpenOrdersResponse {
	return v.value
}

func (v *NullableGetAutoCancelAllOpenOrdersResponse) Set(val *GetAutoCancelAllOpenOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAutoCancelAllOpenOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAutoCancelAllOpenOrdersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAutoCancelAllOpenOrdersResponse(val *GetAutoCancelAllOpenOrdersResponse) *NullableGetAutoCancelAllOpenOrdersResponse {
	return &NullableGetAutoCancelAllOpenOrdersResponse{value: val, isSet: true}
}

func (v NullableGetAutoCancelAllOpenOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAutoCancelAllOpenOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
