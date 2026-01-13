/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the MovePositionForSubAccountResponseMovePositionOrdersInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MovePositionForSubAccountResponseMovePositionOrdersInner{}

// MovePositionForSubAccountResponseMovePositionOrdersInner struct for MovePositionForSubAccountResponseMovePositionOrdersInner
type MovePositionForSubAccountResponseMovePositionOrdersInner struct {
	FromUserEmail        *string `json:"fromUserEmail,omitempty"`
	ToUserEmail          *string `json:"toUserEmail,omitempty"`
	ProductType          *string `json:"productType,omitempty"`
	Symbol               *string `json:"symbol,omitempty"`
	PriceType            *string `json:"priceType,omitempty"`
	Price                *string `json:"price,omitempty"`
	Quantity             *string `json:"quantity,omitempty"`
	PositionSide         *string `json:"positionSide,omitempty"`
	Side                 *string `json:"side,omitempty"`
	Success              *bool   `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MovePositionForSubAccountResponseMovePositionOrdersInner MovePositionForSubAccountResponseMovePositionOrdersInner

// NewMovePositionForSubAccountResponseMovePositionOrdersInner instantiates a new MovePositionForSubAccountResponseMovePositionOrdersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMovePositionForSubAccountResponseMovePositionOrdersInner() *MovePositionForSubAccountResponseMovePositionOrdersInner {
	this := MovePositionForSubAccountResponseMovePositionOrdersInner{}
	return &this
}

// NewMovePositionForSubAccountResponseMovePositionOrdersInnerWithDefaults instantiates a new MovePositionForSubAccountResponseMovePositionOrdersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMovePositionForSubAccountResponseMovePositionOrdersInnerWithDefaults() *MovePositionForSubAccountResponseMovePositionOrdersInner {
	this := MovePositionForSubAccountResponseMovePositionOrdersInner{}
	return &this
}

// GetFromUserEmail returns the FromUserEmail field value if set, zero value otherwise.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) GetFromUserEmail() string {
	if o == nil || common.IsNil(o.FromUserEmail) {
		var ret string
		return ret
	}
	return *o.FromUserEmail
}

// GetFromUserEmailOk returns a tuple with the FromUserEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) GetFromUserEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.FromUserEmail) {
		return nil, false
	}
	return o.FromUserEmail, true
}

// HasFromUserEmail returns a boolean if a field has been set.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) HasFromUserEmail() bool {
	if o != nil && !common.IsNil(o.FromUserEmail) {
		return true
	}

	return false
}

// SetFromUserEmail gets a reference to the given string and assigns it to the FromUserEmail field.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) SetFromUserEmail(v string) {
	o.FromUserEmail = &v
}

// GetToUserEmail returns the ToUserEmail field value if set, zero value otherwise.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) GetToUserEmail() string {
	if o == nil || common.IsNil(o.ToUserEmail) {
		var ret string
		return ret
	}
	return *o.ToUserEmail
}

// GetToUserEmailOk returns a tuple with the ToUserEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) GetToUserEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.ToUserEmail) {
		return nil, false
	}
	return o.ToUserEmail, true
}

// HasToUserEmail returns a boolean if a field has been set.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) HasToUserEmail() bool {
	if o != nil && !common.IsNil(o.ToUserEmail) {
		return true
	}

	return false
}

// SetToUserEmail gets a reference to the given string and assigns it to the ToUserEmail field.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) SetToUserEmail(v string) {
	o.ToUserEmail = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) GetProductType() string {
	if o == nil || common.IsNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) GetProductTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ProductType) {
		return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) HasProductType() bool {
	if o != nil && !common.IsNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) SetProductType(v string) {
	o.ProductType = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetPriceType returns the PriceType field value if set, zero value otherwise.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) GetPriceType() string {
	if o == nil || common.IsNil(o.PriceType) {
		var ret string
		return ret
	}
	return *o.PriceType
}

// GetPriceTypeOk returns a tuple with the PriceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) GetPriceTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.PriceType) {
		return nil, false
	}
	return o.PriceType, true
}

// HasPriceType returns a boolean if a field has been set.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) HasPriceType() bool {
	if o != nil && !common.IsNil(o.PriceType) {
		return true
	}

	return false
}

// SetPriceType gets a reference to the given string and assigns it to the PriceType field.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) SetPriceType(v string) {
	o.PriceType = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) SetPrice(v string) {
	o.Price = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) GetQuantity() string {
	if o == nil || common.IsNil(o.Quantity) {
		var ret string
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) GetQuantityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) HasQuantity() bool {
	if o != nil && !common.IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given string and assigns it to the Quantity field.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) SetQuantity(v string) {
	o.Quantity = &v
}

// GetPositionSide returns the PositionSide field value if set, zero value otherwise.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) GetPositionSide() string {
	if o == nil || common.IsNil(o.PositionSide) {
		var ret string
		return ret
	}
	return *o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) GetPositionSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionSide) {
		return nil, false
	}
	return o.PositionSide, true
}

// HasPositionSide returns a boolean if a field has been set.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) HasPositionSide() bool {
	if o != nil && !common.IsNil(o.PositionSide) {
		return true
	}

	return false
}

// SetPositionSide gets a reference to the given string and assigns it to the PositionSide field.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) SetPositionSide(v string) {
	o.PositionSide = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) SetSide(v string) {
	o.Side = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) SetSuccess(v bool) {
	o.Success = &v
}

func (o MovePositionForSubAccountResponseMovePositionOrdersInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MovePositionForSubAccountResponseMovePositionOrdersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.FromUserEmail) {
		toSerialize["fromUserEmail"] = o.FromUserEmail
	}
	if !common.IsNil(o.ToUserEmail) {
		toSerialize["toUserEmail"] = o.ToUserEmail
	}
	if !common.IsNil(o.ProductType) {
		toSerialize["productType"] = o.ProductType
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.PriceType) {
		toSerialize["priceType"] = o.PriceType
	}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !common.IsNil(o.PositionSide) {
		toSerialize["positionSide"] = o.PositionSide
	}
	if !common.IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !common.IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MovePositionForSubAccountResponseMovePositionOrdersInner) UnmarshalJSON(data []byte) (err error) {
	varMovePositionForSubAccountResponseMovePositionOrdersInner := _MovePositionForSubAccountResponseMovePositionOrdersInner{}

	err = json.Unmarshal(data, &varMovePositionForSubAccountResponseMovePositionOrdersInner)

	if err != nil {
		return err
	}

	*o = MovePositionForSubAccountResponseMovePositionOrdersInner(varMovePositionForSubAccountResponseMovePositionOrdersInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "fromUserEmail")
		delete(additionalProperties, "toUserEmail")
		delete(additionalProperties, "productType")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "priceType")
		delete(additionalProperties, "price")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "positionSide")
		delete(additionalProperties, "side")
		delete(additionalProperties, "success")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMovePositionForSubAccountResponseMovePositionOrdersInner struct {
	value *MovePositionForSubAccountResponseMovePositionOrdersInner
	isSet bool
}

func (v NullableMovePositionForSubAccountResponseMovePositionOrdersInner) Get() *MovePositionForSubAccountResponseMovePositionOrdersInner {
	return v.value
}

func (v *NullableMovePositionForSubAccountResponseMovePositionOrdersInner) Set(val *MovePositionForSubAccountResponseMovePositionOrdersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMovePositionForSubAccountResponseMovePositionOrdersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMovePositionForSubAccountResponseMovePositionOrdersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMovePositionForSubAccountResponseMovePositionOrdersInner(val *MovePositionForSubAccountResponseMovePositionOrdersInner) *NullableMovePositionForSubAccountResponseMovePositionOrdersInner {
	return &NullableMovePositionForSubAccountResponseMovePositionOrdersInner{value: val, isSet: true}
}

func (v NullableMovePositionForSubAccountResponseMovePositionOrdersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMovePositionForSubAccountResponseMovePositionOrdersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
