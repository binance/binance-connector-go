/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the RecentTradesListResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RecentTradesListResponseInner{}

// RecentTradesListResponseInner struct for RecentTradesListResponseInner
type RecentTradesListResponseInner struct {
	Id                   *int64  `json:"id,omitempty"`
	Price                *string `json:"price,omitempty"`
	Qty                  *string `json:"qty,omitempty"`
	QuoteQty             *string `json:"quoteQty,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	IsBuyerMaker         *bool   `json:"isBuyerMaker,omitempty"`
	IsRPITrade           *bool   `json:"isRPITrade,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecentTradesListResponseInner RecentTradesListResponseInner

// NewRecentTradesListResponseInner instantiates a new RecentTradesListResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecentTradesListResponseInner() *RecentTradesListResponseInner {
	this := RecentTradesListResponseInner{}
	return &this
}

// NewRecentTradesListResponseInnerWithDefaults instantiates a new RecentTradesListResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecentTradesListResponseInnerWithDefaults() *RecentTradesListResponseInner {
	this := RecentTradesListResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RecentTradesListResponseInner) GetId() int64 {
	if o == nil || common.IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecentTradesListResponseInner) GetIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RecentTradesListResponseInner) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *RecentTradesListResponseInner) SetId(v int64) {
	o.Id = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *RecentTradesListResponseInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecentTradesListResponseInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *RecentTradesListResponseInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *RecentTradesListResponseInner) SetPrice(v string) {
	o.Price = &v
}

// GetQty returns the Qty field value if set, zero value otherwise.
func (o *RecentTradesListResponseInner) GetQty() string {
	if o == nil || common.IsNil(o.Qty) {
		var ret string
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecentTradesListResponseInner) GetQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Qty) {
		return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *RecentTradesListResponseInner) HasQty() bool {
	if o != nil && !common.IsNil(o.Qty) {
		return true
	}

	return false
}

// SetQty gets a reference to the given string and assigns it to the Qty field.
func (o *RecentTradesListResponseInner) SetQty(v string) {
	o.Qty = &v
}

// GetQuoteQty returns the QuoteQty field value if set, zero value otherwise.
func (o *RecentTradesListResponseInner) GetQuoteQty() string {
	if o == nil || common.IsNil(o.QuoteQty) {
		var ret string
		return ret
	}
	return *o.QuoteQty
}

// GetQuoteQtyOk returns a tuple with the QuoteQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecentTradesListResponseInner) GetQuoteQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.QuoteQty) {
		return nil, false
	}
	return o.QuoteQty, true
}

// HasQuoteQty returns a boolean if a field has been set.
func (o *RecentTradesListResponseInner) HasQuoteQty() bool {
	if o != nil && !common.IsNil(o.QuoteQty) {
		return true
	}

	return false
}

// SetQuoteQty gets a reference to the given string and assigns it to the QuoteQty field.
func (o *RecentTradesListResponseInner) SetQuoteQty(v string) {
	o.QuoteQty = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *RecentTradesListResponseInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecentTradesListResponseInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *RecentTradesListResponseInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *RecentTradesListResponseInner) SetTime(v int64) {
	o.Time = &v
}

// GetIsBuyerMaker returns the IsBuyerMaker field value if set, zero value otherwise.
func (o *RecentTradesListResponseInner) GetIsBuyerMaker() bool {
	if o == nil || common.IsNil(o.IsBuyerMaker) {
		var ret bool
		return ret
	}
	return *o.IsBuyerMaker
}

// GetIsBuyerMakerOk returns a tuple with the IsBuyerMaker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecentTradesListResponseInner) GetIsBuyerMakerOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsBuyerMaker) {
		return nil, false
	}
	return o.IsBuyerMaker, true
}

// HasIsBuyerMaker returns a boolean if a field has been set.
func (o *RecentTradesListResponseInner) HasIsBuyerMaker() bool {
	if o != nil && !common.IsNil(o.IsBuyerMaker) {
		return true
	}

	return false
}

// SetIsBuyerMaker gets a reference to the given bool and assigns it to the IsBuyerMaker field.
func (o *RecentTradesListResponseInner) SetIsBuyerMaker(v bool) {
	o.IsBuyerMaker = &v
}

// GetIsRPITrade returns the IsRPITrade field value if set, zero value otherwise.
func (o *RecentTradesListResponseInner) GetIsRPITrade() bool {
	if o == nil || common.IsNil(o.IsRPITrade) {
		var ret bool
		return ret
	}
	return *o.IsRPITrade
}

// GetIsRPITradeOk returns a tuple with the IsRPITrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecentTradesListResponseInner) GetIsRPITradeOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsRPITrade) {
		return nil, false
	}
	return o.IsRPITrade, true
}

// HasIsRPITrade returns a boolean if a field has been set.
func (o *RecentTradesListResponseInner) HasIsRPITrade() bool {
	if o != nil && !common.IsNil(o.IsRPITrade) {
		return true
	}

	return false
}

// SetIsRPITrade gets a reference to the given bool and assigns it to the IsRPITrade field.
func (o *RecentTradesListResponseInner) SetIsRPITrade(v bool) {
	o.IsRPITrade = &v
}

func (o RecentTradesListResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecentTradesListResponseInner) ToMap() (map[string]interface{}, error) {
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

func (o *RecentTradesListResponseInner) UnmarshalJSON(data []byte) (err error) {
	varRecentTradesListResponseInner := _RecentTradesListResponseInner{}

	err = json.Unmarshal(data, &varRecentTradesListResponseInner)

	if err != nil {
		return err
	}

	*o = RecentTradesListResponseInner(varRecentTradesListResponseInner)

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

type NullableRecentTradesListResponseInner struct {
	value *RecentTradesListResponseInner
	isSet bool
}

func (v NullableRecentTradesListResponseInner) Get() *RecentTradesListResponseInner {
	return v.value
}

func (v *NullableRecentTradesListResponseInner) Set(val *RecentTradesListResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRecentTradesListResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRecentTradesListResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecentTradesListResponseInner(val *RecentTradesListResponseInner) *NullableRecentTradesListResponseInner {
	return &NullableRecentTradesListResponseInner{value: val, isSet: true}
}

func (v NullableRecentTradesListResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecentTradesListResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
