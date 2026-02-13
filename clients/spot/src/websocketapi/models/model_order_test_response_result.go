/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the OrderTestResponseResult type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderTestResponseResult{}

// OrderTestResponseResult struct for OrderTestResponseResult
type OrderTestResponseResult struct {
	StandardCommissionForOrder *OrderTestResponseResultStandardCommissionForOrder `json:"standardCommissionForOrder,omitempty"`
	SpecialCommissionForOrder  *OrderTestResponseResultSpecialCommissionForOrder  `json:"specialCommissionForOrder,omitempty"`
	TaxCommissionForOrder      *OrderTestResponseResultStandardCommissionForOrder `json:"taxCommissionForOrder,omitempty"`
	Discount                   *OrderTestResponseResultDiscount                   `json:"discount,omitempty"`
	AdditionalProperties       map[string]interface{}
}

type _OrderTestResponseResult OrderTestResponseResult

// NewOrderTestResponseResult instantiates a new OrderTestResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderTestResponseResult() *OrderTestResponseResult {
	this := OrderTestResponseResult{}
	return &this
}

// NewOrderTestResponseResultWithDefaults instantiates a new OrderTestResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderTestResponseResultWithDefaults() *OrderTestResponseResult {
	this := OrderTestResponseResult{}
	return &this
}

// GetStandardCommissionForOrder returns the StandardCommissionForOrder field value if set, zero value otherwise.
func (o *OrderTestResponseResult) GetStandardCommissionForOrder() OrderTestResponseResultStandardCommissionForOrder {
	if o == nil || common.IsNil(o.StandardCommissionForOrder) {
		var ret OrderTestResponseResultStandardCommissionForOrder
		return ret
	}
	return *o.StandardCommissionForOrder
}

// GetStandardCommissionForOrderOk returns a tuple with the StandardCommissionForOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTestResponseResult) GetStandardCommissionForOrderOk() (*OrderTestResponseResultStandardCommissionForOrder, bool) {
	if o == nil || common.IsNil(o.StandardCommissionForOrder) {
		return nil, false
	}
	return o.StandardCommissionForOrder, true
}

// HasStandardCommissionForOrder returns a boolean if a field has been set.
func (o *OrderTestResponseResult) HasStandardCommissionForOrder() bool {
	if o != nil && !common.IsNil(o.StandardCommissionForOrder) {
		return true
	}

	return false
}

// SetStandardCommissionForOrder gets a reference to the given OrderTestResponseResultStandardCommissionForOrder and assigns it to the StandardCommissionForOrder field.
func (o *OrderTestResponseResult) SetStandardCommissionForOrder(v OrderTestResponseResultStandardCommissionForOrder) {
	o.StandardCommissionForOrder = &v
}

// GetSpecialCommissionForOrder returns the SpecialCommissionForOrder field value if set, zero value otherwise.
func (o *OrderTestResponseResult) GetSpecialCommissionForOrder() OrderTestResponseResultSpecialCommissionForOrder {
	if o == nil || common.IsNil(o.SpecialCommissionForOrder) {
		var ret OrderTestResponseResultSpecialCommissionForOrder
		return ret
	}
	return *o.SpecialCommissionForOrder
}

// GetSpecialCommissionForOrderOk returns a tuple with the SpecialCommissionForOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTestResponseResult) GetSpecialCommissionForOrderOk() (*OrderTestResponseResultSpecialCommissionForOrder, bool) {
	if o == nil || common.IsNil(o.SpecialCommissionForOrder) {
		return nil, false
	}
	return o.SpecialCommissionForOrder, true
}

// HasSpecialCommissionForOrder returns a boolean if a field has been set.
func (o *OrderTestResponseResult) HasSpecialCommissionForOrder() bool {
	if o != nil && !common.IsNil(o.SpecialCommissionForOrder) {
		return true
	}

	return false
}

// SetSpecialCommissionForOrder gets a reference to the given OrderTestResponseResultSpecialCommissionForOrder and assigns it to the SpecialCommissionForOrder field.
func (o *OrderTestResponseResult) SetSpecialCommissionForOrder(v OrderTestResponseResultSpecialCommissionForOrder) {
	o.SpecialCommissionForOrder = &v
}

// GetTaxCommissionForOrder returns the TaxCommissionForOrder field value if set, zero value otherwise.
func (o *OrderTestResponseResult) GetTaxCommissionForOrder() OrderTestResponseResultStandardCommissionForOrder {
	if o == nil || common.IsNil(o.TaxCommissionForOrder) {
		var ret OrderTestResponseResultStandardCommissionForOrder
		return ret
	}
	return *o.TaxCommissionForOrder
}

// GetTaxCommissionForOrderOk returns a tuple with the TaxCommissionForOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTestResponseResult) GetTaxCommissionForOrderOk() (*OrderTestResponseResultStandardCommissionForOrder, bool) {
	if o == nil || common.IsNil(o.TaxCommissionForOrder) {
		return nil, false
	}
	return o.TaxCommissionForOrder, true
}

// HasTaxCommissionForOrder returns a boolean if a field has been set.
func (o *OrderTestResponseResult) HasTaxCommissionForOrder() bool {
	if o != nil && !common.IsNil(o.TaxCommissionForOrder) {
		return true
	}

	return false
}

// SetTaxCommissionForOrder gets a reference to the given OrderTestResponseResultStandardCommissionForOrder and assigns it to the TaxCommissionForOrder field.
func (o *OrderTestResponseResult) SetTaxCommissionForOrder(v OrderTestResponseResultStandardCommissionForOrder) {
	o.TaxCommissionForOrder = &v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *OrderTestResponseResult) GetDiscount() OrderTestResponseResultDiscount {
	if o == nil || common.IsNil(o.Discount) {
		var ret OrderTestResponseResultDiscount
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTestResponseResult) GetDiscountOk() (*OrderTestResponseResultDiscount, bool) {
	if o == nil || common.IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *OrderTestResponseResult) HasDiscount() bool {
	if o != nil && !common.IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given OrderTestResponseResultDiscount and assigns it to the Discount field.
func (o *OrderTestResponseResult) SetDiscount(v OrderTestResponseResultDiscount) {
	o.Discount = &v
}

func (o OrderTestResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderTestResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.StandardCommissionForOrder) {
		toSerialize["standardCommissionForOrder"] = o.StandardCommissionForOrder
	}
	if !common.IsNil(o.SpecialCommissionForOrder) {
		toSerialize["specialCommissionForOrder"] = o.SpecialCommissionForOrder
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

func (o *OrderTestResponseResult) UnmarshalJSON(data []byte) (err error) {
	varOrderTestResponseResult := _OrderTestResponseResult{}

	err = json.Unmarshal(data, &varOrderTestResponseResult)

	if err != nil {
		return err
	}

	*o = OrderTestResponseResult(varOrderTestResponseResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "standardCommissionForOrder")
		delete(additionalProperties, "specialCommissionForOrder")
		delete(additionalProperties, "taxCommissionForOrder")
		delete(additionalProperties, "discount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrderTestResponseResult struct {
	value *OrderTestResponseResult
	isSet bool
}

func (v NullableOrderTestResponseResult) Get() *OrderTestResponseResult {
	return v.value
}

func (v *NullableOrderTestResponseResult) Set(val *OrderTestResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderTestResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderTestResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderTestResponseResult(val *OrderTestResponseResult) *NullableOrderTestResponseResult {
	return &NullableOrderTestResponseResult{value: val, isSet: true}
}

func (v NullableOrderTestResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderTestResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
