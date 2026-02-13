/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryCmModifyOrderHistoryResponseInnerAmendmentPrice type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryCmModifyOrderHistoryResponseInnerAmendmentPrice{}

// QueryCmModifyOrderHistoryResponseInnerAmendmentPrice struct for QueryCmModifyOrderHistoryResponseInnerAmendmentPrice
type QueryCmModifyOrderHistoryResponseInnerAmendmentPrice struct {
	Before               *string `json:"before,omitempty"`
	After                *string `json:"after,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryCmModifyOrderHistoryResponseInnerAmendmentPrice QueryCmModifyOrderHistoryResponseInnerAmendmentPrice

// NewQueryCmModifyOrderHistoryResponseInnerAmendmentPrice instantiates a new QueryCmModifyOrderHistoryResponseInnerAmendmentPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryCmModifyOrderHistoryResponseInnerAmendmentPrice() *QueryCmModifyOrderHistoryResponseInnerAmendmentPrice {
	this := QueryCmModifyOrderHistoryResponseInnerAmendmentPrice{}
	return &this
}

// NewQueryCmModifyOrderHistoryResponseInnerAmendmentPriceWithDefaults instantiates a new QueryCmModifyOrderHistoryResponseInnerAmendmentPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryCmModifyOrderHistoryResponseInnerAmendmentPriceWithDefaults() *QueryCmModifyOrderHistoryResponseInnerAmendmentPrice {
	this := QueryCmModifyOrderHistoryResponseInnerAmendmentPrice{}
	return &this
}

// GetBefore returns the Before field value if set, zero value otherwise.
func (o *QueryCmModifyOrderHistoryResponseInnerAmendmentPrice) GetBefore() string {
	if o == nil || common.IsNil(o.Before) {
		var ret string
		return ret
	}
	return *o.Before
}

// GetBeforeOk returns a tuple with the Before field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCmModifyOrderHistoryResponseInnerAmendmentPrice) GetBeforeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Before) {
		return nil, false
	}
	return o.Before, true
}

// HasBefore returns a boolean if a field has been set.
func (o *QueryCmModifyOrderHistoryResponseInnerAmendmentPrice) HasBefore() bool {
	if o != nil && !common.IsNil(o.Before) {
		return true
	}

	return false
}

// SetBefore gets a reference to the given string and assigns it to the Before field.
func (o *QueryCmModifyOrderHistoryResponseInnerAmendmentPrice) SetBefore(v string) {
	o.Before = &v
}

// GetAfter returns the After field value if set, zero value otherwise.
func (o *QueryCmModifyOrderHistoryResponseInnerAmendmentPrice) GetAfter() string {
	if o == nil || common.IsNil(o.After) {
		var ret string
		return ret
	}
	return *o.After
}

// GetAfterOk returns a tuple with the After field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCmModifyOrderHistoryResponseInnerAmendmentPrice) GetAfterOk() (*string, bool) {
	if o == nil || common.IsNil(o.After) {
		return nil, false
	}
	return o.After, true
}

// HasAfter returns a boolean if a field has been set.
func (o *QueryCmModifyOrderHistoryResponseInnerAmendmentPrice) HasAfter() bool {
	if o != nil && !common.IsNil(o.After) {
		return true
	}

	return false
}

// SetAfter gets a reference to the given string and assigns it to the After field.
func (o *QueryCmModifyOrderHistoryResponseInnerAmendmentPrice) SetAfter(v string) {
	o.After = &v
}

func (o QueryCmModifyOrderHistoryResponseInnerAmendmentPrice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryCmModifyOrderHistoryResponseInnerAmendmentPrice) ToMap() (map[string]interface{}, error) {
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

func (o *QueryCmModifyOrderHistoryResponseInnerAmendmentPrice) UnmarshalJSON(data []byte) (err error) {
	varQueryCmModifyOrderHistoryResponseInnerAmendmentPrice := _QueryCmModifyOrderHistoryResponseInnerAmendmentPrice{}

	err = json.Unmarshal(data, &varQueryCmModifyOrderHistoryResponseInnerAmendmentPrice)

	if err != nil {
		return err
	}

	*o = QueryCmModifyOrderHistoryResponseInnerAmendmentPrice(varQueryCmModifyOrderHistoryResponseInnerAmendmentPrice)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "before")
		delete(additionalProperties, "after")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryCmModifyOrderHistoryResponseInnerAmendmentPrice struct {
	value *QueryCmModifyOrderHistoryResponseInnerAmendmentPrice
	isSet bool
}

func (v NullableQueryCmModifyOrderHistoryResponseInnerAmendmentPrice) Get() *QueryCmModifyOrderHistoryResponseInnerAmendmentPrice {
	return v.value
}

func (v *NullableQueryCmModifyOrderHistoryResponseInnerAmendmentPrice) Set(val *QueryCmModifyOrderHistoryResponseInnerAmendmentPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryCmModifyOrderHistoryResponseInnerAmendmentPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryCmModifyOrderHistoryResponseInnerAmendmentPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryCmModifyOrderHistoryResponseInnerAmendmentPrice(val *QueryCmModifyOrderHistoryResponseInnerAmendmentPrice) *NullableQueryCmModifyOrderHistoryResponseInnerAmendmentPrice {
	return &NullableQueryCmModifyOrderHistoryResponseInnerAmendmentPrice{value: val, isSet: true}
}

func (v NullableQueryCmModifyOrderHistoryResponseInnerAmendmentPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryCmModifyOrderHistoryResponseInnerAmendmentPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
