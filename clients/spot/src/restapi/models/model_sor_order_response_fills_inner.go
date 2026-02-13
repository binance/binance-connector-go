/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the SorOrderResponseFillsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SorOrderResponseFillsInner{}

// SorOrderResponseFillsInner struct for SorOrderResponseFillsInner
type SorOrderResponseFillsInner struct {
	MatchType            *string `json:"matchType,omitempty"`
	Price                *string `json:"price,omitempty"`
	Qty                  *string `json:"qty,omitempty"`
	Commission           *string `json:"commission,omitempty"`
	CommissionAsset      *string `json:"commissionAsset,omitempty"`
	TradeId              *int64  `json:"tradeId,omitempty"`
	AllocId              *int64  `json:"allocId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SorOrderResponseFillsInner SorOrderResponseFillsInner

// NewSorOrderResponseFillsInner instantiates a new SorOrderResponseFillsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSorOrderResponseFillsInner() *SorOrderResponseFillsInner {
	this := SorOrderResponseFillsInner{}
	return &this
}

// NewSorOrderResponseFillsInnerWithDefaults instantiates a new SorOrderResponseFillsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSorOrderResponseFillsInnerWithDefaults() *SorOrderResponseFillsInner {
	this := SorOrderResponseFillsInner{}
	return &this
}

// GetMatchType returns the MatchType field value if set, zero value otherwise.
func (o *SorOrderResponseFillsInner) GetMatchType() string {
	if o == nil || common.IsNil(o.MatchType) {
		var ret string
		return ret
	}
	return *o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SorOrderResponseFillsInner) GetMatchTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.MatchType) {
		return nil, false
	}
	return o.MatchType, true
}

// HasMatchType returns a boolean if a field has been set.
func (o *SorOrderResponseFillsInner) HasMatchType() bool {
	if o != nil && !common.IsNil(o.MatchType) {
		return true
	}

	return false
}

// SetMatchType gets a reference to the given string and assigns it to the MatchType field.
func (o *SorOrderResponseFillsInner) SetMatchType(v string) {
	o.MatchType = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *SorOrderResponseFillsInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SorOrderResponseFillsInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *SorOrderResponseFillsInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *SorOrderResponseFillsInner) SetPrice(v string) {
	o.Price = &v
}

// GetQty returns the Qty field value if set, zero value otherwise.
func (o *SorOrderResponseFillsInner) GetQty() string {
	if o == nil || common.IsNil(o.Qty) {
		var ret string
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SorOrderResponseFillsInner) GetQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Qty) {
		return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *SorOrderResponseFillsInner) HasQty() bool {
	if o != nil && !common.IsNil(o.Qty) {
		return true
	}

	return false
}

// SetQty gets a reference to the given string and assigns it to the Qty field.
func (o *SorOrderResponseFillsInner) SetQty(v string) {
	o.Qty = &v
}

// GetCommission returns the Commission field value if set, zero value otherwise.
func (o *SorOrderResponseFillsInner) GetCommission() string {
	if o == nil || common.IsNil(o.Commission) {
		var ret string
		return ret
	}
	return *o.Commission
}

// GetCommissionOk returns a tuple with the Commission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SorOrderResponseFillsInner) GetCommissionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Commission) {
		return nil, false
	}
	return o.Commission, true
}

// HasCommission returns a boolean if a field has been set.
func (o *SorOrderResponseFillsInner) HasCommission() bool {
	if o != nil && !common.IsNil(o.Commission) {
		return true
	}

	return false
}

// SetCommission gets a reference to the given string and assigns it to the Commission field.
func (o *SorOrderResponseFillsInner) SetCommission(v string) {
	o.Commission = &v
}

// GetCommissionAsset returns the CommissionAsset field value if set, zero value otherwise.
func (o *SorOrderResponseFillsInner) GetCommissionAsset() string {
	if o == nil || common.IsNil(o.CommissionAsset) {
		var ret string
		return ret
	}
	return *o.CommissionAsset
}

// GetCommissionAssetOk returns a tuple with the CommissionAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SorOrderResponseFillsInner) GetCommissionAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.CommissionAsset) {
		return nil, false
	}
	return o.CommissionAsset, true
}

// HasCommissionAsset returns a boolean if a field has been set.
func (o *SorOrderResponseFillsInner) HasCommissionAsset() bool {
	if o != nil && !common.IsNil(o.CommissionAsset) {
		return true
	}

	return false
}

// SetCommissionAsset gets a reference to the given string and assigns it to the CommissionAsset field.
func (o *SorOrderResponseFillsInner) SetCommissionAsset(v string) {
	o.CommissionAsset = &v
}

// GetTradeId returns the TradeId field value if set, zero value otherwise.
func (o *SorOrderResponseFillsInner) GetTradeId() int64 {
	if o == nil || common.IsNil(o.TradeId) {
		var ret int64
		return ret
	}
	return *o.TradeId
}

// GetTradeIdOk returns a tuple with the TradeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SorOrderResponseFillsInner) GetTradeIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TradeId) {
		return nil, false
	}
	return o.TradeId, true
}

// HasTradeId returns a boolean if a field has been set.
func (o *SorOrderResponseFillsInner) HasTradeId() bool {
	if o != nil && !common.IsNil(o.TradeId) {
		return true
	}

	return false
}

// SetTradeId gets a reference to the given int64 and assigns it to the TradeId field.
func (o *SorOrderResponseFillsInner) SetTradeId(v int64) {
	o.TradeId = &v
}

// GetAllocId returns the AllocId field value if set, zero value otherwise.
func (o *SorOrderResponseFillsInner) GetAllocId() int64 {
	if o == nil || common.IsNil(o.AllocId) {
		var ret int64
		return ret
	}
	return *o.AllocId
}

// GetAllocIdOk returns a tuple with the AllocId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SorOrderResponseFillsInner) GetAllocIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.AllocId) {
		return nil, false
	}
	return o.AllocId, true
}

// HasAllocId returns a boolean if a field has been set.
func (o *SorOrderResponseFillsInner) HasAllocId() bool {
	if o != nil && !common.IsNil(o.AllocId) {
		return true
	}

	return false
}

// SetAllocId gets a reference to the given int64 and assigns it to the AllocId field.
func (o *SorOrderResponseFillsInner) SetAllocId(v int64) {
	o.AllocId = &v
}

func (o SorOrderResponseFillsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SorOrderResponseFillsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.MatchType) {
		toSerialize["matchType"] = o.MatchType
	}
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
	if !common.IsNil(o.TradeId) {
		toSerialize["tradeId"] = o.TradeId
	}
	if !common.IsNil(o.AllocId) {
		toSerialize["allocId"] = o.AllocId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SorOrderResponseFillsInner) UnmarshalJSON(data []byte) (err error) {
	varSorOrderResponseFillsInner := _SorOrderResponseFillsInner{}

	err = json.Unmarshal(data, &varSorOrderResponseFillsInner)

	if err != nil {
		return err
	}

	*o = SorOrderResponseFillsInner(varSorOrderResponseFillsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "matchType")
		delete(additionalProperties, "price")
		delete(additionalProperties, "qty")
		delete(additionalProperties, "commission")
		delete(additionalProperties, "commissionAsset")
		delete(additionalProperties, "tradeId")
		delete(additionalProperties, "allocId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSorOrderResponseFillsInner struct {
	value *SorOrderResponseFillsInner
	isSet bool
}

func (v NullableSorOrderResponseFillsInner) Get() *SorOrderResponseFillsInner {
	return v.value
}

func (v *NullableSorOrderResponseFillsInner) Set(val *SorOrderResponseFillsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSorOrderResponseFillsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSorOrderResponseFillsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSorOrderResponseFillsInner(val *SorOrderResponseFillsInner) *NullableSorOrderResponseFillsInner {
	return &NullableSorOrderResponseFillsInner{value: val, isSet: true}
}

func (v NullableSorOrderResponseFillsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSorOrderResponseFillsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
