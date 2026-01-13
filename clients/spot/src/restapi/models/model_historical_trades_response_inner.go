/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the HistoricalTradesResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &HistoricalTradesResponseInner{}

// HistoricalTradesResponseInner struct for HistoricalTradesResponseInner
type HistoricalTradesResponseInner struct {
	Id                   *int64  `json:"id,omitempty"`
	Price                *string `json:"price,omitempty"`
	Qty                  *string `json:"qty,omitempty"`
	QuoteQty             *string `json:"quoteQty,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	IsBuyerMaker         *bool   `json:"isBuyerMaker,omitempty"`
	IsBestMatch          *bool   `json:"isBestMatch,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HistoricalTradesResponseInner HistoricalTradesResponseInner

// NewHistoricalTradesResponseInner instantiates a new HistoricalTradesResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoricalTradesResponseInner() *HistoricalTradesResponseInner {
	this := HistoricalTradesResponseInner{}
	return &this
}

// NewHistoricalTradesResponseInnerWithDefaults instantiates a new HistoricalTradesResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoricalTradesResponseInnerWithDefaults() *HistoricalTradesResponseInner {
	this := HistoricalTradesResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HistoricalTradesResponseInner) GetId() int64 {
	if o == nil || common.IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalTradesResponseInner) GetIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HistoricalTradesResponseInner) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *HistoricalTradesResponseInner) SetId(v int64) {
	o.Id = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *HistoricalTradesResponseInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalTradesResponseInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *HistoricalTradesResponseInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *HistoricalTradesResponseInner) SetPrice(v string) {
	o.Price = &v
}

// GetQty returns the Qty field value if set, zero value otherwise.
func (o *HistoricalTradesResponseInner) GetQty() string {
	if o == nil || common.IsNil(o.Qty) {
		var ret string
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalTradesResponseInner) GetQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Qty) {
		return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *HistoricalTradesResponseInner) HasQty() bool {
	if o != nil && !common.IsNil(o.Qty) {
		return true
	}

	return false
}

// SetQty gets a reference to the given string and assigns it to the Qty field.
func (o *HistoricalTradesResponseInner) SetQty(v string) {
	o.Qty = &v
}

// GetQuoteQty returns the QuoteQty field value if set, zero value otherwise.
func (o *HistoricalTradesResponseInner) GetQuoteQty() string {
	if o == nil || common.IsNil(o.QuoteQty) {
		var ret string
		return ret
	}
	return *o.QuoteQty
}

// GetQuoteQtyOk returns a tuple with the QuoteQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalTradesResponseInner) GetQuoteQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.QuoteQty) {
		return nil, false
	}
	return o.QuoteQty, true
}

// HasQuoteQty returns a boolean if a field has been set.
func (o *HistoricalTradesResponseInner) HasQuoteQty() bool {
	if o != nil && !common.IsNil(o.QuoteQty) {
		return true
	}

	return false
}

// SetQuoteQty gets a reference to the given string and assigns it to the QuoteQty field.
func (o *HistoricalTradesResponseInner) SetQuoteQty(v string) {
	o.QuoteQty = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *HistoricalTradesResponseInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalTradesResponseInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *HistoricalTradesResponseInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *HistoricalTradesResponseInner) SetTime(v int64) {
	o.Time = &v
}

// GetIsBuyerMaker returns the IsBuyerMaker field value if set, zero value otherwise.
func (o *HistoricalTradesResponseInner) GetIsBuyerMaker() bool {
	if o == nil || common.IsNil(o.IsBuyerMaker) {
		var ret bool
		return ret
	}
	return *o.IsBuyerMaker
}

// GetIsBuyerMakerOk returns a tuple with the IsBuyerMaker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalTradesResponseInner) GetIsBuyerMakerOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsBuyerMaker) {
		return nil, false
	}
	return o.IsBuyerMaker, true
}

// HasIsBuyerMaker returns a boolean if a field has been set.
func (o *HistoricalTradesResponseInner) HasIsBuyerMaker() bool {
	if o != nil && !common.IsNil(o.IsBuyerMaker) {
		return true
	}

	return false
}

// SetIsBuyerMaker gets a reference to the given bool and assigns it to the IsBuyerMaker field.
func (o *HistoricalTradesResponseInner) SetIsBuyerMaker(v bool) {
	o.IsBuyerMaker = &v
}

// GetIsBestMatch returns the IsBestMatch field value if set, zero value otherwise.
func (o *HistoricalTradesResponseInner) GetIsBestMatch() bool {
	if o == nil || common.IsNil(o.IsBestMatch) {
		var ret bool
		return ret
	}
	return *o.IsBestMatch
}

// GetIsBestMatchOk returns a tuple with the IsBestMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalTradesResponseInner) GetIsBestMatchOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsBestMatch) {
		return nil, false
	}
	return o.IsBestMatch, true
}

// HasIsBestMatch returns a boolean if a field has been set.
func (o *HistoricalTradesResponseInner) HasIsBestMatch() bool {
	if o != nil && !common.IsNil(o.IsBestMatch) {
		return true
	}

	return false
}

// SetIsBestMatch gets a reference to the given bool and assigns it to the IsBestMatch field.
func (o *HistoricalTradesResponseInner) SetIsBestMatch(v bool) {
	o.IsBestMatch = &v
}

func (o HistoricalTradesResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoricalTradesResponseInner) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.IsBestMatch) {
		toSerialize["isBestMatch"] = o.IsBestMatch
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HistoricalTradesResponseInner) UnmarshalJSON(data []byte) (err error) {
	varHistoricalTradesResponseInner := _HistoricalTradesResponseInner{}

	err = json.Unmarshal(data, &varHistoricalTradesResponseInner)

	if err != nil {
		return err
	}

	*o = HistoricalTradesResponseInner(varHistoricalTradesResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "price")
		delete(additionalProperties, "qty")
		delete(additionalProperties, "quoteQty")
		delete(additionalProperties, "time")
		delete(additionalProperties, "isBuyerMaker")
		delete(additionalProperties, "isBestMatch")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHistoricalTradesResponseInner struct {
	value *HistoricalTradesResponseInner
	isSet bool
}

func (v NullableHistoricalTradesResponseInner) Get() *HistoricalTradesResponseInner {
	return v.value
}

func (v *NullableHistoricalTradesResponseInner) Set(val *HistoricalTradesResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoricalTradesResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoricalTradesResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoricalTradesResponseInner(val *HistoricalTradesResponseInner) *NullableHistoricalTradesResponseInner {
	return &NullableHistoricalTradesResponseInner{value: val, isSet: true}
}

func (v NullableHistoricalTradesResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoricalTradesResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
