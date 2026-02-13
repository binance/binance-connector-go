/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetOrderModifyHistoryResponseInnerAmendmentPrice type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetOrderModifyHistoryResponseInnerAmendmentPrice{}

// GetOrderModifyHistoryResponseInnerAmendmentPrice struct for GetOrderModifyHistoryResponseInnerAmendmentPrice
type GetOrderModifyHistoryResponseInnerAmendmentPrice struct {
	Before               *string `json:"before,omitempty"`
	After                *string `json:"after,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetOrderModifyHistoryResponseInnerAmendmentPrice GetOrderModifyHistoryResponseInnerAmendmentPrice

// NewGetOrderModifyHistoryResponseInnerAmendmentPrice instantiates a new GetOrderModifyHistoryResponseInnerAmendmentPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrderModifyHistoryResponseInnerAmendmentPrice() *GetOrderModifyHistoryResponseInnerAmendmentPrice {
	this := GetOrderModifyHistoryResponseInnerAmendmentPrice{}
	return &this
}

// NewGetOrderModifyHistoryResponseInnerAmendmentPriceWithDefaults instantiates a new GetOrderModifyHistoryResponseInnerAmendmentPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrderModifyHistoryResponseInnerAmendmentPriceWithDefaults() *GetOrderModifyHistoryResponseInnerAmendmentPrice {
	this := GetOrderModifyHistoryResponseInnerAmendmentPrice{}
	return &this
}

// GetBefore returns the Before field value if set, zero value otherwise.
func (o *GetOrderModifyHistoryResponseInnerAmendmentPrice) GetBefore() string {
	if o == nil || common.IsNil(o.Before) {
		var ret string
		return ret
	}
	return *o.Before
}

// GetBeforeOk returns a tuple with the Before field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderModifyHistoryResponseInnerAmendmentPrice) GetBeforeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Before) {
		return nil, false
	}
	return o.Before, true
}

// HasBefore returns a boolean if a field has been set.
func (o *GetOrderModifyHistoryResponseInnerAmendmentPrice) HasBefore() bool {
	if o != nil && !common.IsNil(o.Before) {
		return true
	}

	return false
}

// SetBefore gets a reference to the given string and assigns it to the Before field.
func (o *GetOrderModifyHistoryResponseInnerAmendmentPrice) SetBefore(v string) {
	o.Before = &v
}

// GetAfter returns the After field value if set, zero value otherwise.
func (o *GetOrderModifyHistoryResponseInnerAmendmentPrice) GetAfter() string {
	if o == nil || common.IsNil(o.After) {
		var ret string
		return ret
	}
	return *o.After
}

// GetAfterOk returns a tuple with the After field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderModifyHistoryResponseInnerAmendmentPrice) GetAfterOk() (*string, bool) {
	if o == nil || common.IsNil(o.After) {
		return nil, false
	}
	return o.After, true
}

// HasAfter returns a boolean if a field has been set.
func (o *GetOrderModifyHistoryResponseInnerAmendmentPrice) HasAfter() bool {
	if o != nil && !common.IsNil(o.After) {
		return true
	}

	return false
}

// SetAfter gets a reference to the given string and assigns it to the After field.
func (o *GetOrderModifyHistoryResponseInnerAmendmentPrice) SetAfter(v string) {
	o.After = &v
}

func (o GetOrderModifyHistoryResponseInnerAmendmentPrice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrderModifyHistoryResponseInnerAmendmentPrice) ToMap() (map[string]interface{}, error) {
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

func (o *GetOrderModifyHistoryResponseInnerAmendmentPrice) UnmarshalJSON(data []byte) (err error) {
	varGetOrderModifyHistoryResponseInnerAmendmentPrice := _GetOrderModifyHistoryResponseInnerAmendmentPrice{}

	err = json.Unmarshal(data, &varGetOrderModifyHistoryResponseInnerAmendmentPrice)

	if err != nil {
		return err
	}

	*o = GetOrderModifyHistoryResponseInnerAmendmentPrice(varGetOrderModifyHistoryResponseInnerAmendmentPrice)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "before")
		delete(additionalProperties, "after")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetOrderModifyHistoryResponseInnerAmendmentPrice struct {
	value *GetOrderModifyHistoryResponseInnerAmendmentPrice
	isSet bool
}

func (v NullableGetOrderModifyHistoryResponseInnerAmendmentPrice) Get() *GetOrderModifyHistoryResponseInnerAmendmentPrice {
	return v.value
}

func (v *NullableGetOrderModifyHistoryResponseInnerAmendmentPrice) Set(val *GetOrderModifyHistoryResponseInnerAmendmentPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrderModifyHistoryResponseInnerAmendmentPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrderModifyHistoryResponseInnerAmendmentPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrderModifyHistoryResponseInnerAmendmentPrice(val *GetOrderModifyHistoryResponseInnerAmendmentPrice) *NullableGetOrderModifyHistoryResponseInnerAmendmentPrice {
	return &NullableGetOrderModifyHistoryResponseInnerAmendmentPrice{value: val, isSet: true}
}

func (v NullableGetOrderModifyHistoryResponseInnerAmendmentPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrderModifyHistoryResponseInnerAmendmentPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
