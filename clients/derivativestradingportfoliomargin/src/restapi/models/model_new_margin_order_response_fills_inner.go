/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the NewMarginOrderResponseFillsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &NewMarginOrderResponseFillsInner{}

// NewMarginOrderResponseFillsInner struct for NewMarginOrderResponseFillsInner
type NewMarginOrderResponseFillsInner struct {
	Price                *string `json:"price,omitempty"`
	Qty                  *string `json:"qty,omitempty"`
	Commission           *string `json:"commission,omitempty"`
	CommissionAsset      *string `json:"commissionAsset,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NewMarginOrderResponseFillsInner NewMarginOrderResponseFillsInner

// NewNewMarginOrderResponseFillsInner instantiates a new NewMarginOrderResponseFillsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewMarginOrderResponseFillsInner() *NewMarginOrderResponseFillsInner {
	this := NewMarginOrderResponseFillsInner{}
	return &this
}

// NewNewMarginOrderResponseFillsInnerWithDefaults instantiates a new NewMarginOrderResponseFillsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewMarginOrderResponseFillsInnerWithDefaults() *NewMarginOrderResponseFillsInner {
	this := NewMarginOrderResponseFillsInner{}
	return &this
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *NewMarginOrderResponseFillsInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewMarginOrderResponseFillsInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *NewMarginOrderResponseFillsInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *NewMarginOrderResponseFillsInner) SetPrice(v string) {
	o.Price = &v
}

// GetQty returns the Qty field value if set, zero value otherwise.
func (o *NewMarginOrderResponseFillsInner) GetQty() string {
	if o == nil || common.IsNil(o.Qty) {
		var ret string
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewMarginOrderResponseFillsInner) GetQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Qty) {
		return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *NewMarginOrderResponseFillsInner) HasQty() bool {
	if o != nil && !common.IsNil(o.Qty) {
		return true
	}

	return false
}

// SetQty gets a reference to the given string and assigns it to the Qty field.
func (o *NewMarginOrderResponseFillsInner) SetQty(v string) {
	o.Qty = &v
}

// GetCommission returns the Commission field value if set, zero value otherwise.
func (o *NewMarginOrderResponseFillsInner) GetCommission() string {
	if o == nil || common.IsNil(o.Commission) {
		var ret string
		return ret
	}
	return *o.Commission
}

// GetCommissionOk returns a tuple with the Commission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewMarginOrderResponseFillsInner) GetCommissionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Commission) {
		return nil, false
	}
	return o.Commission, true
}

// HasCommission returns a boolean if a field has been set.
func (o *NewMarginOrderResponseFillsInner) HasCommission() bool {
	if o != nil && !common.IsNil(o.Commission) {
		return true
	}

	return false
}

// SetCommission gets a reference to the given string and assigns it to the Commission field.
func (o *NewMarginOrderResponseFillsInner) SetCommission(v string) {
	o.Commission = &v
}

// GetCommissionAsset returns the CommissionAsset field value if set, zero value otherwise.
func (o *NewMarginOrderResponseFillsInner) GetCommissionAsset() string {
	if o == nil || common.IsNil(o.CommissionAsset) {
		var ret string
		return ret
	}
	return *o.CommissionAsset
}

// GetCommissionAssetOk returns a tuple with the CommissionAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewMarginOrderResponseFillsInner) GetCommissionAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.CommissionAsset) {
		return nil, false
	}
	return o.CommissionAsset, true
}

// HasCommissionAsset returns a boolean if a field has been set.
func (o *NewMarginOrderResponseFillsInner) HasCommissionAsset() bool {
	if o != nil && !common.IsNil(o.CommissionAsset) {
		return true
	}

	return false
}

// SetCommissionAsset gets a reference to the given string and assigns it to the CommissionAsset field.
func (o *NewMarginOrderResponseFillsInner) SetCommissionAsset(v string) {
	o.CommissionAsset = &v
}

func (o NewMarginOrderResponseFillsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewMarginOrderResponseFillsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.Qty) {
		toSerialize["qty"] = o.Qty
	}
	if !common.IsNil(o.Commission) {
		toSerialize["commission"] = o.Commission
	}
	if !common.IsNil(o.CommissionAsset) {
		toSerialize["commissionAsset"] = o.CommissionAsset
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NewMarginOrderResponseFillsInner) UnmarshalJSON(data []byte) (err error) {
	varNewMarginOrderResponseFillsInner := _NewMarginOrderResponseFillsInner{}

	err = json.Unmarshal(data, &varNewMarginOrderResponseFillsInner)

	if err != nil {
		return err
	}

	*o = NewMarginOrderResponseFillsInner(varNewMarginOrderResponseFillsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "price")
		delete(additionalProperties, "qty")
		delete(additionalProperties, "commission")
		delete(additionalProperties, "commissionAsset")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNewMarginOrderResponseFillsInner struct {
	value *NewMarginOrderResponseFillsInner
	isSet bool
}

func (v NullableNewMarginOrderResponseFillsInner) Get() *NewMarginOrderResponseFillsInner {
	return v.value
}

func (v *NullableNewMarginOrderResponseFillsInner) Set(val *NewMarginOrderResponseFillsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableNewMarginOrderResponseFillsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableNewMarginOrderResponseFillsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewMarginOrderResponseFillsInner(val *NewMarginOrderResponseFillsInner) *NullableNewMarginOrderResponseFillsInner {
	return &NullableNewMarginOrderResponseFillsInner{value: val, isSet: true}
}

func (v NullableNewMarginOrderResponseFillsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewMarginOrderResponseFillsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
