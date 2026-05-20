/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the BlockTradesHistoricalResponseResultInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BlockTradesHistoricalResponseResultInner{}

// BlockTradesHistoricalResponseResultInner struct for BlockTradesHistoricalResponseResultInner
type BlockTradesHistoricalResponseResultInner struct {
	Id                   *int64  `json:"id,omitempty"`
	Price                *string `json:"price,omitempty"`
	Qty                  *string `json:"qty,omitempty"`
	QuoteQty             *string `json:"quoteQty,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	IsBuyerMaker         *bool   `json:"isBuyerMaker,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BlockTradesHistoricalResponseResultInner BlockTradesHistoricalResponseResultInner

// NewBlockTradesHistoricalResponseResultInner instantiates a new BlockTradesHistoricalResponseResultInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockTradesHistoricalResponseResultInner() *BlockTradesHistoricalResponseResultInner {
	this := BlockTradesHistoricalResponseResultInner{}
	return &this
}

// NewBlockTradesHistoricalResponseResultInnerWithDefaults instantiates a new BlockTradesHistoricalResponseResultInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockTradesHistoricalResponseResultInnerWithDefaults() *BlockTradesHistoricalResponseResultInner {
	this := BlockTradesHistoricalResponseResultInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BlockTradesHistoricalResponseResultInner) GetId() int64 {
	if o == nil || common.IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTradesHistoricalResponseResultInner) GetIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BlockTradesHistoricalResponseResultInner) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *BlockTradesHistoricalResponseResultInner) SetId(v int64) {
	o.Id = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *BlockTradesHistoricalResponseResultInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTradesHistoricalResponseResultInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *BlockTradesHistoricalResponseResultInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *BlockTradesHistoricalResponseResultInner) SetPrice(v string) {
	o.Price = &v
}

// GetQty returns the Qty field value if set, zero value otherwise.
func (o *BlockTradesHistoricalResponseResultInner) GetQty() string {
	if o == nil || common.IsNil(o.Qty) {
		var ret string
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTradesHistoricalResponseResultInner) GetQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Qty) {
		return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *BlockTradesHistoricalResponseResultInner) HasQty() bool {
	if o != nil && !common.IsNil(o.Qty) {
		return true
	}

	return false
}

// SetQty gets a reference to the given string and assigns it to the Qty field.
func (o *BlockTradesHistoricalResponseResultInner) SetQty(v string) {
	o.Qty = &v
}

// GetQuoteQty returns the QuoteQty field value if set, zero value otherwise.
func (o *BlockTradesHistoricalResponseResultInner) GetQuoteQty() string {
	if o == nil || common.IsNil(o.QuoteQty) {
		var ret string
		return ret
	}
	return *o.QuoteQty
}

// GetQuoteQtyOk returns a tuple with the QuoteQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTradesHistoricalResponseResultInner) GetQuoteQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.QuoteQty) {
		return nil, false
	}
	return o.QuoteQty, true
}

// HasQuoteQty returns a boolean if a field has been set.
func (o *BlockTradesHistoricalResponseResultInner) HasQuoteQty() bool {
	if o != nil && !common.IsNil(o.QuoteQty) {
		return true
	}

	return false
}

// SetQuoteQty gets a reference to the given string and assigns it to the QuoteQty field.
func (o *BlockTradesHistoricalResponseResultInner) SetQuoteQty(v string) {
	o.QuoteQty = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *BlockTradesHistoricalResponseResultInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTradesHistoricalResponseResultInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *BlockTradesHistoricalResponseResultInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *BlockTradesHistoricalResponseResultInner) SetTime(v int64) {
	o.Time = &v
}

// GetIsBuyerMaker returns the IsBuyerMaker field value if set, zero value otherwise.
func (o *BlockTradesHistoricalResponseResultInner) GetIsBuyerMaker() bool {
	if o == nil || common.IsNil(o.IsBuyerMaker) {
		var ret bool
		return ret
	}
	return *o.IsBuyerMaker
}

// GetIsBuyerMakerOk returns a tuple with the IsBuyerMaker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTradesHistoricalResponseResultInner) GetIsBuyerMakerOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsBuyerMaker) {
		return nil, false
	}
	return o.IsBuyerMaker, true
}

// HasIsBuyerMaker returns a boolean if a field has been set.
func (o *BlockTradesHistoricalResponseResultInner) HasIsBuyerMaker() bool {
	if o != nil && !common.IsNil(o.IsBuyerMaker) {
		return true
	}

	return false
}

// SetIsBuyerMaker gets a reference to the given bool and assigns it to the IsBuyerMaker field.
func (o *BlockTradesHistoricalResponseResultInner) SetIsBuyerMaker(v bool) {
	o.IsBuyerMaker = &v
}

func (o BlockTradesHistoricalResponseResultInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockTradesHistoricalResponseResultInner) ToMap() (map[string]interface{}, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BlockTradesHistoricalResponseResultInner) UnmarshalJSON(data []byte) (err error) {
	varBlockTradesHistoricalResponseResultInner := _BlockTradesHistoricalResponseResultInner{}

	err = json.Unmarshal(data, &varBlockTradesHistoricalResponseResultInner)

	if err != nil {
		return err
	}

	*o = BlockTradesHistoricalResponseResultInner(varBlockTradesHistoricalResponseResultInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "price")
		delete(additionalProperties, "qty")
		delete(additionalProperties, "quoteQty")
		delete(additionalProperties, "time")
		delete(additionalProperties, "isBuyerMaker")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBlockTradesHistoricalResponseResultInner struct {
	value *BlockTradesHistoricalResponseResultInner
	isSet bool
}

func (v NullableBlockTradesHistoricalResponseResultInner) Get() *BlockTradesHistoricalResponseResultInner {
	return v.value
}

func (v *NullableBlockTradesHistoricalResponseResultInner) Set(val *BlockTradesHistoricalResponseResultInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockTradesHistoricalResponseResultInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockTradesHistoricalResponseResultInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockTradesHistoricalResponseResultInner(val *BlockTradesHistoricalResponseResultInner) *NullableBlockTradesHistoricalResponseResultInner {
	return &NullableBlockTradesHistoricalResponseResultInner{value: val, isSet: true}
}

func (v NullableBlockTradesHistoricalResponseResultInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockTradesHistoricalResponseResultInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
