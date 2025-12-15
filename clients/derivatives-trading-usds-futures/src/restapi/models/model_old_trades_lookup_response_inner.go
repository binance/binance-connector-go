/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OldTradesLookupResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OldTradesLookupResponseInner{}

// OldTradesLookupResponseInner struct for OldTradesLookupResponseInner
type OldTradesLookupResponseInner struct {
	Id                   *int64  `json:"id,omitempty"`
	Price                *string `json:"price,omitempty"`
	Qty                  *string `json:"qty,omitempty"`
	QuoteQty             *string `json:"quoteQty,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	IsBuyerMaker         *bool   `json:"isBuyerMaker,omitempty"`
	IsRPITrade           *bool   `json:"isRPITrade,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OldTradesLookupResponseInner OldTradesLookupResponseInner

// NewOldTradesLookupResponseInner instantiates a new OldTradesLookupResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOldTradesLookupResponseInner() *OldTradesLookupResponseInner {
	this := OldTradesLookupResponseInner{}
	return &this
}

// NewOldTradesLookupResponseInnerWithDefaults instantiates a new OldTradesLookupResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOldTradesLookupResponseInnerWithDefaults() *OldTradesLookupResponseInner {
	this := OldTradesLookupResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OldTradesLookupResponseInner) GetId() int64 {
	if o == nil || common.IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OldTradesLookupResponseInner) GetIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OldTradesLookupResponseInner) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *OldTradesLookupResponseInner) SetId(v int64) {
	o.Id = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *OldTradesLookupResponseInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OldTradesLookupResponseInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *OldTradesLookupResponseInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *OldTradesLookupResponseInner) SetPrice(v string) {
	o.Price = &v
}

// GetQty returns the Qty field value if set, zero value otherwise.
func (o *OldTradesLookupResponseInner) GetQty() string {
	if o == nil || common.IsNil(o.Qty) {
		var ret string
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OldTradesLookupResponseInner) GetQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Qty) {
		return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *OldTradesLookupResponseInner) HasQty() bool {
	if o != nil && !common.IsNil(o.Qty) {
		return true
	}

	return false
}

// SetQty gets a reference to the given string and assigns it to the Qty field.
func (o *OldTradesLookupResponseInner) SetQty(v string) {
	o.Qty = &v
}

// GetQuoteQty returns the QuoteQty field value if set, zero value otherwise.
func (o *OldTradesLookupResponseInner) GetQuoteQty() string {
	if o == nil || common.IsNil(o.QuoteQty) {
		var ret string
		return ret
	}
	return *o.QuoteQty
}

// GetQuoteQtyOk returns a tuple with the QuoteQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OldTradesLookupResponseInner) GetQuoteQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.QuoteQty) {
		return nil, false
	}
	return o.QuoteQty, true
}

// HasQuoteQty returns a boolean if a field has been set.
func (o *OldTradesLookupResponseInner) HasQuoteQty() bool {
	if o != nil && !common.IsNil(o.QuoteQty) {
		return true
	}

	return false
}

// SetQuoteQty gets a reference to the given string and assigns it to the QuoteQty field.
func (o *OldTradesLookupResponseInner) SetQuoteQty(v string) {
	o.QuoteQty = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *OldTradesLookupResponseInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OldTradesLookupResponseInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *OldTradesLookupResponseInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *OldTradesLookupResponseInner) SetTime(v int64) {
	o.Time = &v
}

// GetIsBuyerMaker returns the IsBuyerMaker field value if set, zero value otherwise.
func (o *OldTradesLookupResponseInner) GetIsBuyerMaker() bool {
	if o == nil || common.IsNil(o.IsBuyerMaker) {
		var ret bool
		return ret
	}
	return *o.IsBuyerMaker
}

// GetIsBuyerMakerOk returns a tuple with the IsBuyerMaker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OldTradesLookupResponseInner) GetIsBuyerMakerOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsBuyerMaker) {
		return nil, false
	}
	return o.IsBuyerMaker, true
}

// HasIsBuyerMaker returns a boolean if a field has been set.
func (o *OldTradesLookupResponseInner) HasIsBuyerMaker() bool {
	if o != nil && !common.IsNil(o.IsBuyerMaker) {
		return true
	}

	return false
}

// SetIsBuyerMaker gets a reference to the given bool and assigns it to the IsBuyerMaker field.
func (o *OldTradesLookupResponseInner) SetIsBuyerMaker(v bool) {
	o.IsBuyerMaker = &v
}

// GetIsRPITrade returns the IsRPITrade field value if set, zero value otherwise.
func (o *OldTradesLookupResponseInner) GetIsRPITrade() bool {
	if o == nil || common.IsNil(o.IsRPITrade) {
		var ret bool
		return ret
	}
	return *o.IsRPITrade
}

// GetIsRPITradeOk returns a tuple with the IsRPITrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OldTradesLookupResponseInner) GetIsRPITradeOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsRPITrade) {
		return nil, false
	}
	return o.IsRPITrade, true
}

// HasIsRPITrade returns a boolean if a field has been set.
func (o *OldTradesLookupResponseInner) HasIsRPITrade() bool {
	if o != nil && !common.IsNil(o.IsRPITrade) {
		return true
	}

	return false
}

// SetIsRPITrade gets a reference to the given bool and assigns it to the IsRPITrade field.
func (o *OldTradesLookupResponseInner) SetIsRPITrade(v bool) {
	o.IsRPITrade = &v
}

func (o OldTradesLookupResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OldTradesLookupResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.Qty) {
		toSerialize["qty"] = o.Qty
	}
	if !common.IsNil(o.QuoteQty) {
		toSerialize["quoteQty"] = o.QuoteQty
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !common.IsNil(o.IsBuyerMaker) {
		toSerialize["isBuyerMaker"] = o.IsBuyerMaker
	}
	if !common.IsNil(o.IsRPITrade) {
		toSerialize["isRPITrade"] = o.IsRPITrade
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OldTradesLookupResponseInner) UnmarshalJSON(data []byte) (err error) {
	varOldTradesLookupResponseInner := _OldTradesLookupResponseInner{}

	err = json.Unmarshal(data, &varOldTradesLookupResponseInner)

	if err != nil {
		return err
	}

	*o = OldTradesLookupResponseInner(varOldTradesLookupResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "price")
		delete(additionalProperties, "qty")
		delete(additionalProperties, "quoteQty")
		delete(additionalProperties, "time")
		delete(additionalProperties, "isBuyerMaker")
		delete(additionalProperties, "isRPITrade")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOldTradesLookupResponseInner struct {
	value *OldTradesLookupResponseInner
	isSet bool
}

func (v NullableOldTradesLookupResponseInner) Get() *OldTradesLookupResponseInner {
	return v.value
}

func (v *NullableOldTradesLookupResponseInner) Set(val *OldTradesLookupResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOldTradesLookupResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOldTradesLookupResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOldTradesLookupResponseInner(val *OldTradesLookupResponseInner) *NullableOldTradesLookupResponseInner {
	return &NullableOldTradesLookupResponseInner{value: val, isSet: true}
}

func (v NullableOldTradesLookupResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOldTradesLookupResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
