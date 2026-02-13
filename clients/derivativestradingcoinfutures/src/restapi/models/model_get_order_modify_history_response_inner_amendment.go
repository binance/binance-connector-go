/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetOrderModifyHistoryResponseInnerAmendment type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetOrderModifyHistoryResponseInnerAmendment{}

// GetOrderModifyHistoryResponseInnerAmendment struct for GetOrderModifyHistoryResponseInnerAmendment
type GetOrderModifyHistoryResponseInnerAmendment struct {
	Price                *GetOrderModifyHistoryResponseInnerAmendmentPrice   `json:"price,omitempty"`
	OrigQty              *GetOrderModifyHistoryResponseInnerAmendmentOrigQty `json:"origQty,omitempty"`
	Count                *int64                                              `json:"count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetOrderModifyHistoryResponseInnerAmendment GetOrderModifyHistoryResponseInnerAmendment

// NewGetOrderModifyHistoryResponseInnerAmendment instantiates a new GetOrderModifyHistoryResponseInnerAmendment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrderModifyHistoryResponseInnerAmendment() *GetOrderModifyHistoryResponseInnerAmendment {
	this := GetOrderModifyHistoryResponseInnerAmendment{}
	return &this
}

// NewGetOrderModifyHistoryResponseInnerAmendmentWithDefaults instantiates a new GetOrderModifyHistoryResponseInnerAmendment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrderModifyHistoryResponseInnerAmendmentWithDefaults() *GetOrderModifyHistoryResponseInnerAmendment {
	this := GetOrderModifyHistoryResponseInnerAmendment{}
	return &this
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *GetOrderModifyHistoryResponseInnerAmendment) GetPrice() GetOrderModifyHistoryResponseInnerAmendmentPrice {
	if o == nil || common.IsNil(o.Price) {
		var ret GetOrderModifyHistoryResponseInnerAmendmentPrice
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderModifyHistoryResponseInnerAmendment) GetPriceOk() (*GetOrderModifyHistoryResponseInnerAmendmentPrice, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *GetOrderModifyHistoryResponseInnerAmendment) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given GetOrderModifyHistoryResponseInnerAmendmentPrice and assigns it to the Price field.
func (o *GetOrderModifyHistoryResponseInnerAmendment) SetPrice(v GetOrderModifyHistoryResponseInnerAmendmentPrice) {
	o.Price = &v
}

// GetOrigQty returns the OrigQty field value if set, zero value otherwise.
func (o *GetOrderModifyHistoryResponseInnerAmendment) GetOrigQty() GetOrderModifyHistoryResponseInnerAmendmentOrigQty {
	if o == nil || common.IsNil(o.OrigQty) {
		var ret GetOrderModifyHistoryResponseInnerAmendmentOrigQty
		return ret
	}
	return *o.OrigQty
}

// GetOrigQtyOk returns a tuple with the OrigQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderModifyHistoryResponseInnerAmendment) GetOrigQtyOk() (*GetOrderModifyHistoryResponseInnerAmendmentOrigQty, bool) {
	if o == nil || common.IsNil(o.OrigQty) {
		return nil, false
	}
	return o.OrigQty, true
}

// HasOrigQty returns a boolean if a field has been set.
func (o *GetOrderModifyHistoryResponseInnerAmendment) HasOrigQty() bool {
	if o != nil && !common.IsNil(o.OrigQty) {
		return true
	}

	return false
}

// SetOrigQty gets a reference to the given GetOrderModifyHistoryResponseInnerAmendmentOrigQty and assigns it to the OrigQty field.
func (o *GetOrderModifyHistoryResponseInnerAmendment) SetOrigQty(v GetOrderModifyHistoryResponseInnerAmendmentOrigQty) {
	o.OrigQty = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *GetOrderModifyHistoryResponseInnerAmendment) GetCount() int64 {
	if o == nil || common.IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderModifyHistoryResponseInnerAmendment) GetCountOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *GetOrderModifyHistoryResponseInnerAmendment) HasCount() bool {
	if o != nil && !common.IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *GetOrderModifyHistoryResponseInnerAmendment) SetCount(v int64) {
	o.Count = &v
}

func (o GetOrderModifyHistoryResponseInnerAmendment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrderModifyHistoryResponseInnerAmendment) ToMap() (map[string]interface{}, error) {
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

func (o *GetOrderModifyHistoryResponseInnerAmendment) UnmarshalJSON(data []byte) (err error) {
	varGetOrderModifyHistoryResponseInnerAmendment := _GetOrderModifyHistoryResponseInnerAmendment{}

	err = json.Unmarshal(data, &varGetOrderModifyHistoryResponseInnerAmendment)

	if err != nil {
		return err
	}

	*o = GetOrderModifyHistoryResponseInnerAmendment(varGetOrderModifyHistoryResponseInnerAmendment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "price")
		delete(additionalProperties, "origQty")
		delete(additionalProperties, "count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetOrderModifyHistoryResponseInnerAmendment struct {
	value *GetOrderModifyHistoryResponseInnerAmendment
	isSet bool
}

func (v NullableGetOrderModifyHistoryResponseInnerAmendment) Get() *GetOrderModifyHistoryResponseInnerAmendment {
	return v.value
}

func (v *NullableGetOrderModifyHistoryResponseInnerAmendment) Set(val *GetOrderModifyHistoryResponseInnerAmendment) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrderModifyHistoryResponseInnerAmendment) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrderModifyHistoryResponseInnerAmendment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrderModifyHistoryResponseInnerAmendment(val *GetOrderModifyHistoryResponseInnerAmendment) *NullableGetOrderModifyHistoryResponseInnerAmendment {
	return &NullableGetOrderModifyHistoryResponseInnerAmendment{value: val, isSet: true}
}

func (v NullableGetOrderModifyHistoryResponseInnerAmendment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrderModifyHistoryResponseInnerAmendment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
