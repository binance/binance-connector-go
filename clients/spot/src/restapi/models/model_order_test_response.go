/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OrderTestResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderTestResponse{}

// OrderTestResponse struct for OrderTestResponse
type OrderTestResponse struct {
	StandardCommissionForOrder *OrderTestResponseStandardCommissionForOrder `json:"standardCommissionForOrder,omitempty"`
	SpecialCommissionForOrder  *OrderTestResponseSpecialCommissionForOrder  `json:"specialCommissionForOrder,omitempty"`
	TaxCommissionForOrder      *OrderTestResponseStandardCommissionForOrder `json:"taxCommissionForOrder,omitempty"`
	Discount                   *OrderTestResponseDiscount                   `json:"discount,omitempty"`
	AdditionalProperties       map[string]interface{}
}

type _OrderTestResponse OrderTestResponse

// NewOrderTestResponse instantiates a new OrderTestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderTestResponse() *OrderTestResponse {
	this := OrderTestResponse{}
	return &this
}

// NewOrderTestResponseWithDefaults instantiates a new OrderTestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderTestResponseWithDefaults() *OrderTestResponse {
	this := OrderTestResponse{}
	return &this
}

// GetStandardCommissionForOrder returns the StandardCommissionForOrder field value if set, zero value otherwise.
func (o *OrderTestResponse) GetStandardCommissionForOrder() OrderTestResponseStandardCommissionForOrder {
	if o == nil || common.IsNil(o.StandardCommissionForOrder) {
		var ret OrderTestResponseStandardCommissionForOrder
		return ret
	}
	return *o.StandardCommissionForOrder
}

// GetStandardCommissionForOrderOk returns a tuple with the StandardCommissionForOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTestResponse) GetStandardCommissionForOrderOk() (*OrderTestResponseStandardCommissionForOrder, bool) {
	if o == nil || common.IsNil(o.StandardCommissionForOrder) {
		return nil, false
	}
	return o.StandardCommissionForOrder, true
}

// HasStandardCommissionForOrder returns a boolean if a field has been set.
func (o *OrderTestResponse) HasStandardCommissionForOrder() bool {
	if o != nil && !common.IsNil(o.StandardCommissionForOrder) {
		return true
	}

	return false
}

// SetStandardCommissionForOrder gets a reference to the given OrderTestResponseStandardCommissionForOrder and assigns it to the StandardCommissionForOrder field.
func (o *OrderTestResponse) SetStandardCommissionForOrder(v OrderTestResponseStandardCommissionForOrder) {
	o.StandardCommissionForOrder = &v
}

// GetSpecialCommissionForOrder returns the SpecialCommissionForOrder field value if set, zero value otherwise.
func (o *OrderTestResponse) GetSpecialCommissionForOrder() OrderTestResponseSpecialCommissionForOrder {
	if o == nil || common.IsNil(o.SpecialCommissionForOrder) {
		var ret OrderTestResponseSpecialCommissionForOrder
		return ret
	}
	return *o.SpecialCommissionForOrder
}

// GetSpecialCommissionForOrderOk returns a tuple with the SpecialCommissionForOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTestResponse) GetSpecialCommissionForOrderOk() (*OrderTestResponseSpecialCommissionForOrder, bool) {
	if o == nil || common.IsNil(o.SpecialCommissionForOrder) {
		return nil, false
	}
	return o.SpecialCommissionForOrder, true
}

// HasSpecialCommissionForOrder returns a boolean if a field has been set.
func (o *OrderTestResponse) HasSpecialCommissionForOrder() bool {
	if o != nil && !common.IsNil(o.SpecialCommissionForOrder) {
		return true
	}

	return false
}

// SetSpecialCommissionForOrder gets a reference to the given OrderTestResponseSpecialCommissionForOrder and assigns it to the SpecialCommissionForOrder field.
func (o *OrderTestResponse) SetSpecialCommissionForOrder(v OrderTestResponseSpecialCommissionForOrder) {
	o.SpecialCommissionForOrder = &v
}

// GetTaxCommissionForOrder returns the TaxCommissionForOrder field value if set, zero value otherwise.
func (o *OrderTestResponse) GetTaxCommissionForOrder() OrderTestResponseStandardCommissionForOrder {
	if o == nil || common.IsNil(o.TaxCommissionForOrder) {
		var ret OrderTestResponseStandardCommissionForOrder
		return ret
	}
	return *o.TaxCommissionForOrder
}

// GetTaxCommissionForOrderOk returns a tuple with the TaxCommissionForOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTestResponse) GetTaxCommissionForOrderOk() (*OrderTestResponseStandardCommissionForOrder, bool) {
	if o == nil || common.IsNil(o.TaxCommissionForOrder) {
		return nil, false
	}
	return o.TaxCommissionForOrder, true
}

// HasTaxCommissionForOrder returns a boolean if a field has been set.
func (o *OrderTestResponse) HasTaxCommissionForOrder() bool {
	if o != nil && !common.IsNil(o.TaxCommissionForOrder) {
		return true
	}

	return false
}

// SetTaxCommissionForOrder gets a reference to the given OrderTestResponseStandardCommissionForOrder and assigns it to the TaxCommissionForOrder field.
func (o *OrderTestResponse) SetTaxCommissionForOrder(v OrderTestResponseStandardCommissionForOrder) {
	o.TaxCommissionForOrder = &v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *OrderTestResponse) GetDiscount() OrderTestResponseDiscount {
	if o == nil || common.IsNil(o.Discount) {
		var ret OrderTestResponseDiscount
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTestResponse) GetDiscountOk() (*OrderTestResponseDiscount, bool) {
	if o == nil || common.IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *OrderTestResponse) HasDiscount() bool {
	if o != nil && !common.IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given OrderTestResponseDiscount and assigns it to the Discount field.
func (o *OrderTestResponse) SetDiscount(v OrderTestResponseDiscount) {
	o.Discount = &v
}

func (o OrderTestResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderTestResponse) ToMap() (map[string]interface{}, error) {
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

func (o *OrderTestResponse) UnmarshalJSON(data []byte) (err error) {
	varOrderTestResponse := _OrderTestResponse{}

	err = json.Unmarshal(data, &varOrderTestResponse)

	if err != nil {
		return err
	}

	*o = OrderTestResponse(varOrderTestResponse)

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

type NullableOrderTestResponse struct {
	value *OrderTestResponse
	isSet bool
}

func (v NullableOrderTestResponse) Get() *OrderTestResponse {
	return v.value
}

func (v *NullableOrderTestResponse) Set(val *OrderTestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderTestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderTestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderTestResponse(val *OrderTestResponse) *NullableOrderTestResponse {
	return &NullableOrderTestResponse{value: val, isSet: true}
}

func (v NullableOrderTestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderTestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
