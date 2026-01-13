/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner{}

// GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner struct for GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner
type GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner struct {
	FromUserEmail        *string `json:"fromUserEmail,omitempty"`
	ToUserEmail          *string `json:"toUserEmail,omitempty"`
	ProductType          *string `json:"productType,omitempty"`
	Symbol               *string `json:"symbol,omitempty"`
	Price                *string `json:"price,omitempty"`
	Quantity             *string `json:"quantity,omitempty"`
	PositionSide         *string `json:"positionSide,omitempty"`
	Side                 *string `json:"side,omitempty"`
	TimeStamp            *int64  `json:"timeStamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner

// NewGetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner instantiates a new GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner() *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner {
	this := GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner{}
	return &this
}

// NewGetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInnerWithDefaults instantiates a new GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInnerWithDefaults() *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner {
	this := GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner{}
	return &this
}

// GetFromUserEmail returns the FromUserEmail field value if set, zero value otherwise.
func (o *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) GetFromUserEmail() string {
	if o == nil || common.IsNil(o.FromUserEmail) {
		var ret string
		return ret
	}
	return *o.FromUserEmail
}

// GetFromUserEmailOk returns a tuple with the FromUserEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) GetFromUserEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.FromUserEmail) {
		return nil, false
	}
	return o.FromUserEmail, true
}

// HasFromUserEmail returns a boolean if a field has been set.
func (o *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) HasFromUserEmail() bool {
	if o != nil && !common.IsNil(o.FromUserEmail) {
		return true
	}

	return false
}

// SetFromUserEmail gets a reference to the given string and assigns it to the FromUserEmail field.
func (o *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) SetFromUserEmail(v string) {
	o.FromUserEmail = &v
}

// GetToUserEmail returns the ToUserEmail field value if set, zero value otherwise.
func (o *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) GetToUserEmail() string {
	if o == nil || common.IsNil(o.ToUserEmail) {
		var ret string
		return ret
	}
	return *o.ToUserEmail
}

// GetToUserEmailOk returns a tuple with the ToUserEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) GetToUserEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.ToUserEmail) {
		return nil, false
	}
	return o.ToUserEmail, true
}

// HasToUserEmail returns a boolean if a field has been set.
func (o *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) HasToUserEmail() bool {
	if o != nil && !common.IsNil(o.ToUserEmail) {
		return true
	}

	return false
}

// SetToUserEmail gets a reference to the given string and assigns it to the ToUserEmail field.
func (o *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) SetToUserEmail(v string) {
	o.ToUserEmail = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) GetProductType() string {
	if o == nil || common.IsNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) GetProductTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ProductType) {
		return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) HasProductType() bool {
	if o != nil && !common.IsNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) SetProductType(v string) {
	o.ProductType = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) SetPrice(v string) {
	o.Price = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) GetQuantity() string {
	if o == nil || common.IsNil(o.Quantity) {
		var ret string
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) GetQuantityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) HasQuantity() bool {
	if o != nil && !common.IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given string and assigns it to the Quantity field.
func (o *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) SetQuantity(v string) {
	o.Quantity = &v
}

// GetPositionSide returns the PositionSide field value if set, zero value otherwise.
func (o *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) GetPositionSide() string {
	if o == nil || common.IsNil(o.PositionSide) {
		var ret string
		return ret
	}
	return *o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) GetPositionSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionSide) {
		return nil, false
	}
	return o.PositionSide, true
}

// HasPositionSide returns a boolean if a field has been set.
func (o *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) HasPositionSide() bool {
	if o != nil && !common.IsNil(o.PositionSide) {
		return true
	}

	return false
}

// SetPositionSide gets a reference to the given string and assigns it to the PositionSide field.
func (o *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) SetPositionSide(v string) {
	o.PositionSide = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) SetSide(v string) {
	o.Side = &v
}

// GetTimeStamp returns the TimeStamp field value if set, zero value otherwise.
func (o *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) GetTimeStamp() int64 {
	if o == nil || common.IsNil(o.TimeStamp) {
		var ret int64
		return ret
	}
	return *o.TimeStamp
}

// GetTimeStampOk returns a tuple with the TimeStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) GetTimeStampOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TimeStamp) {
		return nil, false
	}
	return o.TimeStamp, true
}

// HasTimeStamp returns a boolean if a field has been set.
func (o *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) HasTimeStamp() bool {
	if o != nil && !common.IsNil(o.TimeStamp) {
		return true
	}

	return false
}

// SetTimeStamp gets a reference to the given int64 and assigns it to the TimeStamp field.
func (o *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) SetTimeStamp(v int64) {
	o.TimeStamp = &v
}

func (o GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.TimeStamp) {
		toSerialize["timeStamp"] = o.TimeStamp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) UnmarshalJSON(data []byte) (err error) {
	varGetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner := _GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner{}

	err = json.Unmarshal(data, &varGetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner)

	if err != nil {
		return err
	}

	*o = GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner(varGetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "fromUserEmail")
		delete(additionalProperties, "toUserEmail")
		delete(additionalProperties, "productType")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "price")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "positionSide")
		delete(additionalProperties, "side")
		delete(additionalProperties, "timeStamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner struct {
	value *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner
	isSet bool
}

func (v NullableGetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) Get() *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner {
	return v.value
}

func (v *NullableGetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) Set(val *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner(val *GetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) *NullableGetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner {
	return &NullableGetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner{value: val, isSet: true}
}

func (v NullableGetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMovePositionHistoryForSubAccountResponseFutureMovePositionOrderVoListInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
