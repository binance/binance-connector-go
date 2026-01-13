/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryCmModifyOrderHistoryResponseInnerAmendment type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryCmModifyOrderHistoryResponseInnerAmendment{}

// QueryCmModifyOrderHistoryResponseInnerAmendment struct for QueryCmModifyOrderHistoryResponseInnerAmendment
type QueryCmModifyOrderHistoryResponseInnerAmendment struct {
	Price                *QueryCmModifyOrderHistoryResponseInnerAmendmentPrice   `json:"price,omitempty"`
	OrigQty              *QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty `json:"origQty,omitempty"`
	Count                *int64                                                  `json:"count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryCmModifyOrderHistoryResponseInnerAmendment QueryCmModifyOrderHistoryResponseInnerAmendment

// NewQueryCmModifyOrderHistoryResponseInnerAmendment instantiates a new QueryCmModifyOrderHistoryResponseInnerAmendment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryCmModifyOrderHistoryResponseInnerAmendment() *QueryCmModifyOrderHistoryResponseInnerAmendment {
	this := QueryCmModifyOrderHistoryResponseInnerAmendment{}
	return &this
}

// NewQueryCmModifyOrderHistoryResponseInnerAmendmentWithDefaults instantiates a new QueryCmModifyOrderHistoryResponseInnerAmendment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryCmModifyOrderHistoryResponseInnerAmendmentWithDefaults() *QueryCmModifyOrderHistoryResponseInnerAmendment {
	this := QueryCmModifyOrderHistoryResponseInnerAmendment{}
	return &this
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *QueryCmModifyOrderHistoryResponseInnerAmendment) GetPrice() QueryCmModifyOrderHistoryResponseInnerAmendmentPrice {
	if o == nil || common.IsNil(o.Price) {
		var ret QueryCmModifyOrderHistoryResponseInnerAmendmentPrice
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCmModifyOrderHistoryResponseInnerAmendment) GetPriceOk() (*QueryCmModifyOrderHistoryResponseInnerAmendmentPrice, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *QueryCmModifyOrderHistoryResponseInnerAmendment) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given QueryCmModifyOrderHistoryResponseInnerAmendmentPrice and assigns it to the Price field.
func (o *QueryCmModifyOrderHistoryResponseInnerAmendment) SetPrice(v QueryCmModifyOrderHistoryResponseInnerAmendmentPrice) {
	o.Price = &v
}

// GetOrigQty returns the OrigQty field value if set, zero value otherwise.
func (o *QueryCmModifyOrderHistoryResponseInnerAmendment) GetOrigQty() QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty {
	if o == nil || common.IsNil(o.OrigQty) {
		var ret QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty
		return ret
	}
	return *o.OrigQty
}

// GetOrigQtyOk returns a tuple with the OrigQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCmModifyOrderHistoryResponseInnerAmendment) GetOrigQtyOk() (*QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty, bool) {
	if o == nil || common.IsNil(o.OrigQty) {
		return nil, false
	}
	return o.OrigQty, true
}

// HasOrigQty returns a boolean if a field has been set.
func (o *QueryCmModifyOrderHistoryResponseInnerAmendment) HasOrigQty() bool {
	if o != nil && !common.IsNil(o.OrigQty) {
		return true
	}

	return false
}

// SetOrigQty gets a reference to the given QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty and assigns it to the OrigQty field.
func (o *QueryCmModifyOrderHistoryResponseInnerAmendment) SetOrigQty(v QueryCmModifyOrderHistoryResponseInnerAmendmentOrigQty) {
	o.OrigQty = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *QueryCmModifyOrderHistoryResponseInnerAmendment) GetCount() int64 {
	if o == nil || common.IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCmModifyOrderHistoryResponseInnerAmendment) GetCountOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *QueryCmModifyOrderHistoryResponseInnerAmendment) HasCount() bool {
	if o != nil && !common.IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *QueryCmModifyOrderHistoryResponseInnerAmendment) SetCount(v int64) {
	o.Count = &v
}

func (o QueryCmModifyOrderHistoryResponseInnerAmendment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryCmModifyOrderHistoryResponseInnerAmendment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.OrigQty) {
		toSerialize["origQty"] = o.OrigQty
	}
	if !common.IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryCmModifyOrderHistoryResponseInnerAmendment) UnmarshalJSON(data []byte) (err error) {
	varQueryCmModifyOrderHistoryResponseInnerAmendment := _QueryCmModifyOrderHistoryResponseInnerAmendment{}

	err = json.Unmarshal(data, &varQueryCmModifyOrderHistoryResponseInnerAmendment)

	if err != nil {
		return err
	}

	*o = QueryCmModifyOrderHistoryResponseInnerAmendment(varQueryCmModifyOrderHistoryResponseInnerAmendment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "price")
		delete(additionalProperties, "origQty")
		delete(additionalProperties, "count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryCmModifyOrderHistoryResponseInnerAmendment struct {
	value *QueryCmModifyOrderHistoryResponseInnerAmendment
	isSet bool
}

func (v NullableQueryCmModifyOrderHistoryResponseInnerAmendment) Get() *QueryCmModifyOrderHistoryResponseInnerAmendment {
	return v.value
}

func (v *NullableQueryCmModifyOrderHistoryResponseInnerAmendment) Set(val *QueryCmModifyOrderHistoryResponseInnerAmendment) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryCmModifyOrderHistoryResponseInnerAmendment) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryCmModifyOrderHistoryResponseInnerAmendment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryCmModifyOrderHistoryResponseInnerAmendment(val *QueryCmModifyOrderHistoryResponseInnerAmendment) *NullableQueryCmModifyOrderHistoryResponseInnerAmendment {
	return &NullableQueryCmModifyOrderHistoryResponseInnerAmendment{value: val, isSet: true}
}

func (v NullableQueryCmModifyOrderHistoryResponseInnerAmendment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryCmModifyOrderHistoryResponseInnerAmendment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
