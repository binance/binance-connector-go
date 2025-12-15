/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty{}

// QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty struct for QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty
type QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty struct {
	Before               *string `json:"before,omitempty"`
	After                *string `json:"after,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty

// NewQueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty instantiates a new QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty() *QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty {
	this := QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty{}
	return &this
}

// NewQueryCmModifyOrderHistoryResponseInnerAmendmentOrigQtyWithDefaults instantiates a new QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryCmModifyOrderHistoryResponseInnerAmendmentOrigQtyWithDefaults() *QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty {
	this := QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty{}
	return &this
}

// GetBefore returns the Before field value if set, zero value otherwise.
func (o *QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty) GetBefore() string {
	if o == nil || common.IsNil(o.Before) {
		var ret string
		return ret
	}
	return *o.Before
}

// GetBeforeOk returns a tuple with the Before field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty) GetBeforeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Before) {
		return nil, false
	}
	return o.Before, true
}

// HasBefore returns a boolean if a field has been set.
func (o *QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty) HasBefore() bool {
	if o != nil && !common.IsNil(o.Before) {
		return true
	}

	return false
}

// SetBefore gets a reference to the given string and assigns it to the Before field.
func (o *QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty) SetBefore(v string) {
	o.Before = &v
}

// GetAfter returns the After field value if set, zero value otherwise.
func (o *QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty) GetAfter() string {
	if o == nil || common.IsNil(o.After) {
		var ret string
		return ret
	}
	return *o.After
}

// GetAfterOk returns a tuple with the After field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty) GetAfterOk() (*string, bool) {
	if o == nil || common.IsNil(o.After) {
		return nil, false
	}
	return o.After, true
}

// HasAfter returns a boolean if a field has been set.
func (o *QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty) HasAfter() bool {
	if o != nil && !common.IsNil(o.After) {
		return true
	}

	return false
}

// SetAfter gets a reference to the given string and assigns it to the After field.
func (o *QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty) SetAfter(v string) {
	o.After = &v
}

func (o QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty) ToMap() (map[string]interface{}, error) {
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

func (o *QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty) UnmarshalJSON(data []byte) (err error) {
	varQueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty := _QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty{}

	err = json.Unmarshal(data, &varQueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty)

	if err != nil {
		return err
	}

	*o = QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty(varQueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "before")
		delete(additionalProperties, "after")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty struct {
	value *QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty
	isSet bool
}

func (v NullableQueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty) Get() *QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty {
	return v.value
}

func (v *NullableQueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty) Set(val *QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty(val *QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty) *NullableQueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty {
	return &NullableQueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty{value: val, isSet: true}
}

func (v NullableQueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
