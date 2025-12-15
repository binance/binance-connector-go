/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetOrderModifyHistoryResponseInnerAmendmentOrigQty type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetOrderModifyHistoryResponseInnerAmendmentOrigQty{}

// GetOrderModifyHistoryResponseInnerAmendmentOrigQty struct for GetOrderModifyHistoryResponseInnerAmendmentOrigQty
type GetOrderModifyHistoryResponseInnerAmendmentOrigQty struct {
	Before               *string `json:"before,omitempty"`
	After                *string `json:"after,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetOrderModifyHistoryResponseInnerAmendmentOrigQty GetOrderModifyHistoryResponseInnerAmendmentOrigQty

// NewGetOrderModifyHistoryResponseInnerAmendmentOrigQty instantiates a new GetOrderModifyHistoryResponseInnerAmendmentOrigQty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrderModifyHistoryResponseInnerAmendmentOrigQty() *GetOrderModifyHistoryResponseInnerAmendmentOrigQty {
	this := GetOrderModifyHistoryResponseInnerAmendmentOrigQty{}
	return &this
}

// NewGetOrderModifyHistoryResponseInnerAmendmentOrigQtyWithDefaults instantiates a new GetOrderModifyHistoryResponseInnerAmendmentOrigQty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrderModifyHistoryResponseInnerAmendmentOrigQtyWithDefaults() *GetOrderModifyHistoryResponseInnerAmendmentOrigQty {
	this := GetOrderModifyHistoryResponseInnerAmendmentOrigQty{}
	return &this
}

// GetBefore returns the Before field value if set, zero value otherwise.
func (o *GetOrderModifyHistoryResponseInnerAmendmentOrigQty) GetBefore() string {
	if o == nil || common.IsNil(o.Before) {
		var ret string
		return ret
	}
	return *o.Before
}

// GetBeforeOk returns a tuple with the Before field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderModifyHistoryResponseInnerAmendmentOrigQty) GetBeforeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Before) {
		return nil, false
	}
	return o.Before, true
}

// HasBefore returns a boolean if a field has been set.
func (o *GetOrderModifyHistoryResponseInnerAmendmentOrigQty) HasBefore() bool {
	if o != nil && !common.IsNil(o.Before) {
		return true
	}

	return false
}

// SetBefore gets a reference to the given string and assigns it to the Before field.
func (o *GetOrderModifyHistoryResponseInnerAmendmentOrigQty) SetBefore(v string) {
	o.Before = &v
}

// GetAfter returns the After field value if set, zero value otherwise.
func (o *GetOrderModifyHistoryResponseInnerAmendmentOrigQty) GetAfter() string {
	if o == nil || common.IsNil(o.After) {
		var ret string
		return ret
	}
	return *o.After
}

// GetAfterOk returns a tuple with the After field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderModifyHistoryResponseInnerAmendmentOrigQty) GetAfterOk() (*string, bool) {
	if o == nil || common.IsNil(o.After) {
		return nil, false
	}
	return o.After, true
}

// HasAfter returns a boolean if a field has been set.
func (o *GetOrderModifyHistoryResponseInnerAmendmentOrigQty) HasAfter() bool {
	if o != nil && !common.IsNil(o.After) {
		return true
	}

	return false
}

// SetAfter gets a reference to the given string and assigns it to the After field.
func (o *GetOrderModifyHistoryResponseInnerAmendmentOrigQty) SetAfter(v string) {
	o.After = &v
}

func (o GetOrderModifyHistoryResponseInnerAmendmentOrigQty) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrderModifyHistoryResponseInnerAmendmentOrigQty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Before) {
		toSerialize["before"] = o.Before
	}
	if !common.IsNil(o.After) {
		toSerialize["after"] = o.After
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetOrderModifyHistoryResponseInnerAmendmentOrigQty) UnmarshalJSON(data []byte) (err error) {
	varGetOrderModifyHistoryResponseInnerAmendmentOrigQty := _GetOrderModifyHistoryResponseInnerAmendmentOrigQty{}

	err = json.Unmarshal(data, &varGetOrderModifyHistoryResponseInnerAmendmentOrigQty)

	if err != nil {
		return err
	}

	*o = GetOrderModifyHistoryResponseInnerAmendmentOrigQty(varGetOrderModifyHistoryResponseInnerAmendmentOrigQty)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "before")
		delete(additionalProperties, "after")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetOrderModifyHistoryResponseInnerAmendmentOrigQty struct {
	value *GetOrderModifyHistoryResponseInnerAmendmentOrigQty
	isSet bool
}

func (v NullableGetOrderModifyHistoryResponseInnerAmendmentOrigQty) Get() *GetOrderModifyHistoryResponseInnerAmendmentOrigQty {
	return v.value
}

func (v *NullableGetOrderModifyHistoryResponseInnerAmendmentOrigQty) Set(val *GetOrderModifyHistoryResponseInnerAmendmentOrigQty) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrderModifyHistoryResponseInnerAmendmentOrigQty) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrderModifyHistoryResponseInnerAmendmentOrigQty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrderModifyHistoryResponseInnerAmendmentOrigQty(val *GetOrderModifyHistoryResponseInnerAmendmentOrigQty) *NullableGetOrderModifyHistoryResponseInnerAmendmentOrigQty {
	return &NullableGetOrderModifyHistoryResponseInnerAmendmentOrigQty{value: val, isSet: true}
}

func (v NullableGetOrderModifyHistoryResponseInnerAmendmentOrigQty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrderModifyHistoryResponseInnerAmendmentOrigQty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
