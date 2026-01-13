/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the SorOrderTestResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SorOrderTestResponse{}

// SorOrderTestResponse struct for SorOrderTestResponse
type SorOrderTestResponse struct {
	StandardCommissionForOrder *OrderTestResponseStandardCommissionForOrder `json:"standardCommissionForOrder,omitempty"`
	TaxCommissionForOrder      *OrderTestResponseStandardCommissionForOrder `json:"taxCommissionForOrder,omitempty"`
	Discount                   *OrderTestResponseDiscount                   `json:"discount,omitempty"`
	AdditionalProperties       map[string]interface{}
}

type _SorOrderTestResponse SorOrderTestResponse

// NewSorOrderTestResponse instantiates a new SorOrderTestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSorOrderTestResponse() *SorOrderTestResponse {
	this := SorOrderTestResponse{}
	return &this
}

// NewSorOrderTestResponseWithDefaults instantiates a new SorOrderTestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSorOrderTestResponseWithDefaults() *SorOrderTestResponse {
	this := SorOrderTestResponse{}
	return &this
}

// GetStandardCommissionForOrder returns the StandardCommissionForOrder field value if set, zero value otherwise.
func (o *SorOrderTestResponse) GetStandardCommissionForOrder() OrderTestResponseStandardCommissionForOrder {
	if o == nil || common.IsNil(o.StandardCommissionForOrder) {
		var ret OrderTestResponseStandardCommissionForOrder
		return ret
	}
	return *o.StandardCommissionForOrder
}

// GetStandardCommissionForOrderOk returns a tuple with the StandardCommissionForOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SorOrderTestResponse) GetStandardCommissionForOrderOk() (*OrderTestResponseStandardCommissionForOrder, bool) {
	if o == nil || common.IsNil(o.StandardCommissionForOrder) {
		return nil, false
	}
	return o.StandardCommissionForOrder, true
}

// HasStandardCommissionForOrder returns a boolean if a field has been set.
func (o *SorOrderTestResponse) HasStandardCommissionForOrder() bool {
	if o != nil && !common.IsNil(o.StandardCommissionForOrder) {
		return true
	}

	return false
}

// SetStandardCommissionForOrder gets a reference to the given OrderTestResponseStandardCommissionForOrder and assigns it to the StandardCommissionForOrder field.
func (o *SorOrderTestResponse) SetStandardCommissionForOrder(v OrderTestResponseStandardCommissionForOrder) {
	o.StandardCommissionForOrder = &v
}

// GetTaxCommissionForOrder returns the TaxCommissionForOrder field value if set, zero value otherwise.
func (o *SorOrderTestResponse) GetTaxCommissionForOrder() OrderTestResponseStandardCommissionForOrder {
	if o == nil || common.IsNil(o.TaxCommissionForOrder) {
		var ret OrderTestResponseStandardCommissionForOrder
		return ret
	}
	return *o.TaxCommissionForOrder
}

// GetTaxCommissionForOrderOk returns a tuple with the TaxCommissionForOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SorOrderTestResponse) GetTaxCommissionForOrderOk() (*OrderTestResponseStandardCommissionForOrder, bool) {
	if o == nil || common.IsNil(o.TaxCommissionForOrder) {
		return nil, false
	}
	return o.TaxCommissionForOrder, true
}

// HasTaxCommissionForOrder returns a boolean if a field has been set.
func (o *SorOrderTestResponse) HasTaxCommissionForOrder() bool {
	if o != nil && !common.IsNil(o.TaxCommissionForOrder) {
		return true
	}

	return false
}

// SetTaxCommissionForOrder gets a reference to the given OrderTestResponseStandardCommissionForOrder and assigns it to the TaxCommissionForOrder field.
func (o *SorOrderTestResponse) SetTaxCommissionForOrder(v OrderTestResponseStandardCommissionForOrder) {
	o.TaxCommissionForOrder = &v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *SorOrderTestResponse) GetDiscount() OrderTestResponseDiscount {
	if o == nil || common.IsNil(o.Discount) {
		var ret OrderTestResponseDiscount
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SorOrderTestResponse) GetDiscountOk() (*OrderTestResponseDiscount, bool) {
	if o == nil || common.IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *SorOrderTestResponse) HasDiscount() bool {
	if o != nil && !common.IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given OrderTestResponseDiscount and assigns it to the Discount field.
func (o *SorOrderTestResponse) SetDiscount(v OrderTestResponseDiscount) {
	o.Discount = &v
}

func (o SorOrderTestResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SorOrderTestResponse) ToMap() (map[string]interface{}, error) {
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

func (o *SorOrderTestResponse) UnmarshalJSON(data []byte) (err error) {
	varSorOrderTestResponse := _SorOrderTestResponse{}

	err = json.Unmarshal(data, &varSorOrderTestResponse)

	if err != nil {
		return err
	}

	*o = SorOrderTestResponse(varSorOrderTestResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "standardCommissionForOrder")
		delete(additionalProperties, "taxCommissionForOrder")
		delete(additionalProperties, "discount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSorOrderTestResponse struct {
	value *SorOrderTestResponse
	isSet bool
}

func (v NullableSorOrderTestResponse) Get() *SorOrderTestResponse {
	return v.value
}

func (v *NullableSorOrderTestResponse) Set(val *SorOrderTestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSorOrderTestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSorOrderTestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSorOrderTestResponse(val *SorOrderTestResponse) *NullableSorOrderTestResponse {
	return &NullableSorOrderTestResponse{value: val, isSet: true}
}

func (v NullableSorOrderTestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSorOrderTestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
