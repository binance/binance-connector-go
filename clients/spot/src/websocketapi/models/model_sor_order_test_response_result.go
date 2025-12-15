/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the SorOrderTestResponseResult type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SorOrderTestResponseResult{}

// SorOrderTestResponseResult struct for SorOrderTestResponseResult
type SorOrderTestResponseResult struct {
	StandardCommissionForOrder *OrderTestResponseResultStandardCommissionForOrder `json:"standardCommissionForOrder,omitempty"`
	TaxCommissionForOrder      *OrderTestResponseResultStandardCommissionForOrder `json:"taxCommissionForOrder,omitempty"`
	Discount                   *OrderTestResponseResultDiscount                   `json:"discount,omitempty"`
	AdditionalProperties       map[string]interface{}
}

type _SorOrderTestResponseResult SorOrderTestResponseResult

// NewSorOrderTestResponseResult instantiates a new SorOrderTestResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSorOrderTestResponseResult() *SorOrderTestResponseResult {
	this := SorOrderTestResponseResult{}
	return &this
}

// NewSorOrderTestResponseResultWithDefaults instantiates a new SorOrderTestResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSorOrderTestResponseResultWithDefaults() *SorOrderTestResponseResult {
	this := SorOrderTestResponseResult{}
	return &this
}

// GetStandardCommissionForOrder returns the StandardCommissionForOrder field value if set, zero value otherwise.
func (o *SorOrderTestResponseResult) GetStandardCommissionForOrder() OrderTestResponseResultStandardCommissionForOrder {
	if o == nil || common.IsNil(o.StandardCommissionForOrder) {
		var ret OrderTestResponseResultStandardCommissionForOrder
		return ret
	}
	return *o.StandardCommissionForOrder
}

// GetStandardCommissionForOrderOk returns a tuple with the StandardCommissionForOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SorOrderTestResponseResult) GetStandardCommissionForOrderOk() (*OrderTestResponseResultStandardCommissionForOrder, bool) {
	if o == nil || common.IsNil(o.StandardCommissionForOrder) {
		return nil, false
	}
	return o.StandardCommissionForOrder, true
}

// HasStandardCommissionForOrder returns a boolean if a field has been set.
func (o *SorOrderTestResponseResult) HasStandardCommissionForOrder() bool {
	if o != nil && !common.IsNil(o.StandardCommissionForOrder) {
		return true
	}

	return false
}

// SetStandardCommissionForOrder gets a reference to the given OrderTestResponseResultStandardCommissionForOrder and assigns it to the StandardCommissionForOrder field.
func (o *SorOrderTestResponseResult) SetStandardCommissionForOrder(v OrderTestResponseResultStandardCommissionForOrder) {
	o.StandardCommissionForOrder = &v
}

// GetTaxCommissionForOrder returns the TaxCommissionForOrder field value if set, zero value otherwise.
func (o *SorOrderTestResponseResult) GetTaxCommissionForOrder() OrderTestResponseResultStandardCommissionForOrder {
	if o == nil || common.IsNil(o.TaxCommissionForOrder) {
		var ret OrderTestResponseResultStandardCommissionForOrder
		return ret
	}
	return *o.TaxCommissionForOrder
}

// GetTaxCommissionForOrderOk returns a tuple with the TaxCommissionForOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SorOrderTestResponseResult) GetTaxCommissionForOrderOk() (*OrderTestResponseResultStandardCommissionForOrder, bool) {
	if o == nil || common.IsNil(o.TaxCommissionForOrder) {
		return nil, false
	}
	return o.TaxCommissionForOrder, true
}

// HasTaxCommissionForOrder returns a boolean if a field has been set.
func (o *SorOrderTestResponseResult) HasTaxCommissionForOrder() bool {
	if o != nil && !common.IsNil(o.TaxCommissionForOrder) {
		return true
	}

	return false
}

// SetTaxCommissionForOrder gets a reference to the given OrderTestResponseResultStandardCommissionForOrder and assigns it to the TaxCommissionForOrder field.
func (o *SorOrderTestResponseResult) SetTaxCommissionForOrder(v OrderTestResponseResultStandardCommissionForOrder) {
	o.TaxCommissionForOrder = &v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *SorOrderTestResponseResult) GetDiscount() OrderTestResponseResultDiscount {
	if o == nil || common.IsNil(o.Discount) {
		var ret OrderTestResponseResultDiscount
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SorOrderTestResponseResult) GetDiscountOk() (*OrderTestResponseResultDiscount, bool) {
	if o == nil || common.IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *SorOrderTestResponseResult) HasDiscount() bool {
	if o != nil && !common.IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given OrderTestResponseResultDiscount and assigns it to the Discount field.
func (o *SorOrderTestResponseResult) SetDiscount(v OrderTestResponseResultDiscount) {
	o.Discount = &v
}

func (o SorOrderTestResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SorOrderTestResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.StandardCommissionForOrder) {
		toSerialize["standardCommissionForOrder"] = o.StandardCommissionForOrder
	}
	if !common.IsNil(o.TaxCommissionForOrder) {
		toSerialize["taxCommissionForOrder"] = o.TaxCommissionForOrder
	}
	if !common.IsNil(o.Discount) {
		toSerialize["discount"] = o.Discount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SorOrderTestResponseResult) UnmarshalJSON(data []byte) (err error) {
	varSorOrderTestResponseResult := _SorOrderTestResponseResult{}

	err = json.Unmarshal(data, &varSorOrderTestResponseResult)

	if err != nil {
		return err
	}

	*o = SorOrderTestResponseResult(varSorOrderTestResponseResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "standardCommissionForOrder")
		delete(additionalProperties, "taxCommissionForOrder")
		delete(additionalProperties, "discount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSorOrderTestResponseResult struct {
	value *SorOrderTestResponseResult
	isSet bool
}

func (v NullableSorOrderTestResponseResult) Get() *SorOrderTestResponseResult {
	return v.value
}

func (v *NullableSorOrderTestResponseResult) Set(val *SorOrderTestResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSorOrderTestResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSorOrderTestResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSorOrderTestResponseResult(val *SorOrderTestResponseResult) *NullableSorOrderTestResponseResult {
	return &NullableSorOrderTestResponseResult{value: val, isSet: true}
}

func (v NullableSorOrderTestResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSorOrderTestResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
