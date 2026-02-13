/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the UserExerciseRecordResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UserExerciseRecordResponseInner{}

// UserExerciseRecordResponseInner struct for UserExerciseRecordResponseInner
type UserExerciseRecordResponseInner struct {
	Id                   *string `json:"id,omitempty"`
	Currency             *string `json:"currency,omitempty"`
	Symbol               *string `json:"symbol,omitempty"`
	ExercisePrice        *string `json:"exercisePrice,omitempty"`
	Quantity             *string `json:"quantity,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	Fee                  *string `json:"fee,omitempty"`
	CreateDate           *int64  `json:"createDate,omitempty"`
	PriceScale           *int64  `json:"priceScale,omitempty"`
	QuantityScale        *int64  `json:"quantityScale,omitempty"`
	OptionSide           *string `json:"optionSide,omitempty"`
	PositionSide         *string `json:"positionSide,omitempty"`
	QuoteAsset           *string `json:"quoteAsset,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserExerciseRecordResponseInner UserExerciseRecordResponseInner

// NewUserExerciseRecordResponseInner instantiates a new UserExerciseRecordResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserExerciseRecordResponseInner() *UserExerciseRecordResponseInner {
	this := UserExerciseRecordResponseInner{}
	return &this
}

// NewUserExerciseRecordResponseInnerWithDefaults instantiates a new UserExerciseRecordResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserExerciseRecordResponseInnerWithDefaults() *UserExerciseRecordResponseInner {
	this := UserExerciseRecordResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserExerciseRecordResponseInner) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserExerciseRecordResponseInner) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserExerciseRecordResponseInner) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserExerciseRecordResponseInner) SetId(v string) {
	o.Id = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UserExerciseRecordResponseInner) GetCurrency() string {
	if o == nil || common.IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserExerciseRecordResponseInner) GetCurrencyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UserExerciseRecordResponseInner) HasCurrency() bool {
	if o != nil && !common.IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UserExerciseRecordResponseInner) SetCurrency(v string) {
	o.Currency = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *UserExerciseRecordResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserExerciseRecordResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *UserExerciseRecordResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *UserExerciseRecordResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetExercisePrice returns the ExercisePrice field value if set, zero value otherwise.
func (o *UserExerciseRecordResponseInner) GetExercisePrice() string {
	if o == nil || common.IsNil(o.ExercisePrice) {
		var ret string
		return ret
	}
	return *o.ExercisePrice
}

// GetExercisePriceOk returns a tuple with the ExercisePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserExerciseRecordResponseInner) GetExercisePriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExercisePrice) {
		return nil, false
	}
	return o.ExercisePrice, true
}

// HasExercisePrice returns a boolean if a field has been set.
func (o *UserExerciseRecordResponseInner) HasExercisePrice() bool {
	if o != nil && !common.IsNil(o.ExercisePrice) {
		return true
	}

	return false
}

// SetExercisePrice gets a reference to the given string and assigns it to the ExercisePrice field.
func (o *UserExerciseRecordResponseInner) SetExercisePrice(v string) {
	o.ExercisePrice = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *UserExerciseRecordResponseInner) GetQuantity() string {
	if o == nil || common.IsNil(o.Quantity) {
		var ret string
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserExerciseRecordResponseInner) GetQuantityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *UserExerciseRecordResponseInner) HasQuantity() bool {
	if o != nil && !common.IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given string and assigns it to the Quantity field.
func (o *UserExerciseRecordResponseInner) SetQuantity(v string) {
	o.Quantity = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *UserExerciseRecordResponseInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserExerciseRecordResponseInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *UserExerciseRecordResponseInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *UserExerciseRecordResponseInner) SetAmount(v string) {
	o.Amount = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *UserExerciseRecordResponseInner) GetFee() string {
	if o == nil || common.IsNil(o.Fee) {
		var ret string
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserExerciseRecordResponseInner) GetFeeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *UserExerciseRecordResponseInner) HasFee() bool {
	if o != nil && !common.IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given string and assigns it to the Fee field.
func (o *UserExerciseRecordResponseInner) SetFee(v string) {
	o.Fee = &v
}

// GetCreateDate returns the CreateDate field value if set, zero value otherwise.
func (o *UserExerciseRecordResponseInner) GetCreateDate() int64 {
	if o == nil || common.IsNil(o.CreateDate) {
		var ret int64
		return ret
	}
	return *o.CreateDate
}

// GetCreateDateOk returns a tuple with the CreateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserExerciseRecordResponseInner) GetCreateDateOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CreateDate) {
		return nil, false
	}
	return o.CreateDate, true
}

// HasCreateDate returns a boolean if a field has been set.
func (o *UserExerciseRecordResponseInner) HasCreateDate() bool {
	if o != nil && !common.IsNil(o.CreateDate) {
		return true
	}

	return false
}

// SetCreateDate gets a reference to the given int64 and assigns it to the CreateDate field.
func (o *UserExerciseRecordResponseInner) SetCreateDate(v int64) {
	o.CreateDate = &v
}

// GetPriceScale returns the PriceScale field value if set, zero value otherwise.
func (o *UserExerciseRecordResponseInner) GetPriceScale() int64 {
	if o == nil || common.IsNil(o.PriceScale) {
		var ret int64
		return ret
	}
	return *o.PriceScale
}

// GetPriceScaleOk returns a tuple with the PriceScale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserExerciseRecordResponseInner) GetPriceScaleOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PriceScale) {
		return nil, false
	}
	return o.PriceScale, true
}

// HasPriceScale returns a boolean if a field has been set.
func (o *UserExerciseRecordResponseInner) HasPriceScale() bool {
	if o != nil && !common.IsNil(o.PriceScale) {
		return true
	}

	return false
}

// SetPriceScale gets a reference to the given int64 and assigns it to the PriceScale field.
func (o *UserExerciseRecordResponseInner) SetPriceScale(v int64) {
	o.PriceScale = &v
}

// GetQuantityScale returns the QuantityScale field value if set, zero value otherwise.
func (o *UserExerciseRecordResponseInner) GetQuantityScale() int64 {
	if o == nil || common.IsNil(o.QuantityScale) {
		var ret int64
		return ret
	}
	return *o.QuantityScale
}

// GetQuantityScaleOk returns a tuple with the QuantityScale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserExerciseRecordResponseInner) GetQuantityScaleOk() (*int64, bool) {
	if o == nil || common.IsNil(o.QuantityScale) {
		return nil, false
	}
	return o.QuantityScale, true
}

// HasQuantityScale returns a boolean if a field has been set.
func (o *UserExerciseRecordResponseInner) HasQuantityScale() bool {
	if o != nil && !common.IsNil(o.QuantityScale) {
		return true
	}

	return false
}

// SetQuantityScale gets a reference to the given int64 and assigns it to the QuantityScale field.
func (o *UserExerciseRecordResponseInner) SetQuantityScale(v int64) {
	o.QuantityScale = &v
}

// GetOptionSide returns the OptionSide field value if set, zero value otherwise.
func (o *UserExerciseRecordResponseInner) GetOptionSide() string {
	if o == nil || common.IsNil(o.OptionSide) {
		var ret string
		return ret
	}
	return *o.OptionSide
}

// GetOptionSideOk returns a tuple with the OptionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserExerciseRecordResponseInner) GetOptionSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.OptionSide) {
		return nil, false
	}
	return o.OptionSide, true
}

// HasOptionSide returns a boolean if a field has been set.
func (o *UserExerciseRecordResponseInner) HasOptionSide() bool {
	if o != nil && !common.IsNil(o.OptionSide) {
		return true
	}

	return false
}

// SetOptionSide gets a reference to the given string and assigns it to the OptionSide field.
func (o *UserExerciseRecordResponseInner) SetOptionSide(v string) {
	o.OptionSide = &v
}

// GetPositionSide returns the PositionSide field value if set, zero value otherwise.
func (o *UserExerciseRecordResponseInner) GetPositionSide() string {
	if o == nil || common.IsNil(o.PositionSide) {
		var ret string
		return ret
	}
	return *o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserExerciseRecordResponseInner) GetPositionSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionSide) {
		return nil, false
	}
	return o.PositionSide, true
}

// HasPositionSide returns a boolean if a field has been set.
func (o *UserExerciseRecordResponseInner) HasPositionSide() bool {
	if o != nil && !common.IsNil(o.PositionSide) {
		return true
	}

	return false
}

// SetPositionSide gets a reference to the given string and assigns it to the PositionSide field.
func (o *UserExerciseRecordResponseInner) SetPositionSide(v string) {
	o.PositionSide = &v
}

// GetQuoteAsset returns the QuoteAsset field value if set, zero value otherwise.
func (o *UserExerciseRecordResponseInner) GetQuoteAsset() string {
	if o == nil || common.IsNil(o.QuoteAsset) {
		var ret string
		return ret
	}
	return *o.QuoteAsset
}

// GetQuoteAssetOk returns a tuple with the QuoteAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserExerciseRecordResponseInner) GetQuoteAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.QuoteAsset) {
		return nil, false
	}
	return o.QuoteAsset, true
}

// HasQuoteAsset returns a boolean if a field has been set.
func (o *UserExerciseRecordResponseInner) HasQuoteAsset() bool {
	if o != nil && !common.IsNil(o.QuoteAsset) {
		return true
	}

	return false
}

// SetQuoteAsset gets a reference to the given string and assigns it to the QuoteAsset field.
func (o *UserExerciseRecordResponseInner) SetQuoteAsset(v string) {
	o.QuoteAsset = &v
}

func (o UserExerciseRecordResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserExerciseRecordResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.ExercisePrice) {
		toSerialize["exercisePrice"] = o.ExercisePrice
	}
	if !common.IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	if !common.IsNil(o.CreateDate) {
		toSerialize["createDate"] = o.CreateDate
	}
	if !common.IsNil(o.PriceScale) {
		toSerialize["priceScale"] = o.PriceScale
	}
	if !common.IsNil(o.QuantityScale) {
		toSerialize["quantityScale"] = o.QuantityScale
	}
	if !common.IsNil(o.OptionSide) {
		toSerialize["optionSide"] = o.OptionSide
	}
	if !common.IsNil(o.PositionSide) {
		toSerialize["positionSide"] = o.PositionSide
	}
	if !common.IsNil(o.QuoteAsset) {
		toSerialize["quoteAsset"] = o.QuoteAsset
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserExerciseRecordResponseInner) UnmarshalJSON(data []byte) (err error) {
	varUserExerciseRecordResponseInner := _UserExerciseRecordResponseInner{}

	err = json.Unmarshal(data, &varUserExerciseRecordResponseInner)

	if err != nil {
		return err
	}

	*o = UserExerciseRecordResponseInner(varUserExerciseRecordResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "exercisePrice")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "fee")
		delete(additionalProperties, "createDate")
		delete(additionalProperties, "priceScale")
		delete(additionalProperties, "quantityScale")
		delete(additionalProperties, "optionSide")
		delete(additionalProperties, "positionSide")
		delete(additionalProperties, "quoteAsset")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserExerciseRecordResponseInner struct {
	value *UserExerciseRecordResponseInner
	isSet bool
}

func (v NullableUserExerciseRecordResponseInner) Get() *UserExerciseRecordResponseInner {
	return v.value
}

func (v *NullableUserExerciseRecordResponseInner) Set(val *UserExerciseRecordResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUserExerciseRecordResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUserExerciseRecordResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserExerciseRecordResponseInner(val *UserExerciseRecordResponseInner) *NullableUserExerciseRecordResponseInner {
	return &NullableUserExerciseRecordResponseInner{value: val, isSet: true}
}

func (v NullableUserExerciseRecordResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserExerciseRecordResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
