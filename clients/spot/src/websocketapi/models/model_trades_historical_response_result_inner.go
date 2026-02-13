/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the TradesHistoricalResponseResultInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TradesHistoricalResponseResultInner{}

// TradesHistoricalResponseResultInner struct for TradesHistoricalResponseResultInner
type TradesHistoricalResponseResultInner struct {
	Id                   *int64  `json:"id,omitempty"`
	Price                *string `json:"price,omitempty"`
	Qty                  *string `json:"qty,omitempty"`
	QuoteQty             *string `json:"quoteQty,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	IsBuyerMaker         *bool   `json:"isBuyerMaker,omitempty"`
	IsBestMatch          *bool   `json:"isBestMatch,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TradesHistoricalResponseResultInner TradesHistoricalResponseResultInner

// NewTradesHistoricalResponseResultInner instantiates a new TradesHistoricalResponseResultInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTradesHistoricalResponseResultInner() *TradesHistoricalResponseResultInner {
	this := TradesHistoricalResponseResultInner{}
	return &this
}

// NewTradesHistoricalResponseResultInnerWithDefaults instantiates a new TradesHistoricalResponseResultInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTradesHistoricalResponseResultInnerWithDefaults() *TradesHistoricalResponseResultInner {
	this := TradesHistoricalResponseResultInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TradesHistoricalResponseResultInner) GetId() int64 {
	if o == nil || common.IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradesHistoricalResponseResultInner) GetIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TradesHistoricalResponseResultInner) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *TradesHistoricalResponseResultInner) SetId(v int64) {
	o.Id = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *TradesHistoricalResponseResultInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradesHistoricalResponseResultInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *TradesHistoricalResponseResultInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *TradesHistoricalResponseResultInner) SetPrice(v string) {
	o.Price = &v
}

// GetQty returns the Qty field value if set, zero value otherwise.
func (o *TradesHistoricalResponseResultInner) GetQty() string {
	if o == nil || common.IsNil(o.Qty) {
		var ret string
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradesHistoricalResponseResultInner) GetQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Qty) {
		return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *TradesHistoricalResponseResultInner) HasQty() bool {
	if o != nil && !common.IsNil(o.Qty) {
		return true
	}

	return false
}

// SetQty gets a reference to the given string and assigns it to the Qty field.
func (o *TradesHistoricalResponseResultInner) SetQty(v string) {
	o.Qty = &v
}

// GetQuoteQty returns the QuoteQty field value if set, zero value otherwise.
func (o *TradesHistoricalResponseResultInner) GetQuoteQty() string {
	if o == nil || common.IsNil(o.QuoteQty) {
		var ret string
		return ret
	}
	return *o.QuoteQty
}

// GetQuoteQtyOk returns a tuple with the QuoteQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradesHistoricalResponseResultInner) GetQuoteQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.QuoteQty) {
		return nil, false
	}
	return o.QuoteQty, true
}

// HasQuoteQty returns a boolean if a field has been set.
func (o *TradesHistoricalResponseResultInner) HasQuoteQty() bool {
	if o != nil && !common.IsNil(o.QuoteQty) {
		return true
	}

	return false
}

// SetQuoteQty gets a reference to the given string and assigns it to the QuoteQty field.
func (o *TradesHistoricalResponseResultInner) SetQuoteQty(v string) {
	o.QuoteQty = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *TradesHistoricalResponseResultInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradesHistoricalResponseResultInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *TradesHistoricalResponseResultInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *TradesHistoricalResponseResultInner) SetTime(v int64) {
	o.Time = &v
}

// GetIsBuyerMaker returns the IsBuyerMaker field value if set, zero value otherwise.
func (o *TradesHistoricalResponseResultInner) GetIsBuyerMaker() bool {
	if o == nil || common.IsNil(o.IsBuyerMaker) {
		var ret bool
		return ret
	}
	return *o.IsBuyerMaker
}

// GetIsBuyerMakerOk returns a tuple with the IsBuyerMaker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradesHistoricalResponseResultInner) GetIsBuyerMakerOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsBuyerMaker) {
		return nil, false
	}
	return o.IsBuyerMaker, true
}

// HasIsBuyerMaker returns a boolean if a field has been set.
func (o *TradesHistoricalResponseResultInner) HasIsBuyerMaker() bool {
	if o != nil && !common.IsNil(o.IsBuyerMaker) {
		return true
	}

	return false
}

// SetIsBuyerMaker gets a reference to the given bool and assigns it to the IsBuyerMaker field.
func (o *TradesHistoricalResponseResultInner) SetIsBuyerMaker(v bool) {
	o.IsBuyerMaker = &v
}

// GetIsBestMatch returns the IsBestMatch field value if set, zero value otherwise.
func (o *TradesHistoricalResponseResultInner) GetIsBestMatch() bool {
	if o == nil || common.IsNil(o.IsBestMatch) {
		var ret bool
		return ret
	}
	return *o.IsBestMatch
}

// GetIsBestMatchOk returns a tuple with the IsBestMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradesHistoricalResponseResultInner) GetIsBestMatchOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsBestMatch) {
		return nil, false
	}
	return o.IsBestMatch, true
}

// HasIsBestMatch returns a boolean if a field has been set.
func (o *TradesHistoricalResponseResultInner) HasIsBestMatch() bool {
	if o != nil && !common.IsNil(o.IsBestMatch) {
		return true
	}

	return false
}

// SetIsBestMatch gets a reference to the given bool and assigns it to the IsBestMatch field.
func (o *TradesHistoricalResponseResultInner) SetIsBestMatch(v bool) {
	o.IsBestMatch = &v
}

func (o TradesHistoricalResponseResultInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TradesHistoricalResponseResultInner) ToMap() (map[string]interface{}, error) {
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

func (o *TradesHistoricalResponseResultInner) UnmarshalJSON(data []byte) (err error) {
	varTradesHistoricalResponseResultInner := _TradesHistoricalResponseResultInner{}

	err = json.Unmarshal(data, &varTradesHistoricalResponseResultInner)

	if err != nil {
		return err
	}

	*o = TradesHistoricalResponseResultInner(varTradesHistoricalResponseResultInner)

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

type NullableTradesHistoricalResponseResultInner struct {
	value *TradesHistoricalResponseResultInner
	isSet bool
}

func (v NullableTradesHistoricalResponseResultInner) Get() *TradesHistoricalResponseResultInner {
	return v.value
}

func (v *NullableTradesHistoricalResponseResultInner) Set(val *TradesHistoricalResponseResultInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTradesHistoricalResponseResultInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTradesHistoricalResponseResultInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTradesHistoricalResponseResultInner(val *TradesHistoricalResponseResultInner) *NullableTradesHistoricalResponseResultInner {
	return &NullableTradesHistoricalResponseResultInner{value: val, isSet: true}
}

func (v NullableTradesHistoricalResponseResultInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTradesHistoricalResponseResultInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
