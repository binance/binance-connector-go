/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the SorOrderPlaceResponseResultInnerFillsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SorOrderPlaceResponseResultInnerFillsInner{}

// SorOrderPlaceResponseResultInnerFillsInner struct for SorOrderPlaceResponseResultInnerFillsInner
type SorOrderPlaceResponseResultInnerFillsInner struct {
	MatchType            *string `json:"matchType,omitempty"`
	Price                *string `json:"price,omitempty"`
	Qty                  *string `json:"qty,omitempty"`
	Commission           *string `json:"commission,omitempty"`
	CommissionAsset      *string `json:"commissionAsset,omitempty"`
	TradeId              *int64  `json:"tradeId,omitempty"`
	AllocId              *int64  `json:"allocId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SorOrderPlaceResponseResultInnerFillsInner SorOrderPlaceResponseResultInnerFillsInner

// NewSorOrderPlaceResponseResultInnerFillsInner instantiates a new SorOrderPlaceResponseResultInnerFillsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSorOrderPlaceResponseResultInnerFillsInner() *SorOrderPlaceResponseResultInnerFillsInner {
	this := SorOrderPlaceResponseResultInnerFillsInner{}
	return &this
}

// NewSorOrderPlaceResponseResultInnerFillsInnerWithDefaults instantiates a new SorOrderPlaceResponseResultInnerFillsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSorOrderPlaceResponseResultInnerFillsInnerWithDefaults() *SorOrderPlaceResponseResultInnerFillsInner {
	this := SorOrderPlaceResponseResultInnerFillsInner{}
	return &this
}

// GetMatchType returns the MatchType field value if set, zero value otherwise.
func (o *SorOrderPlaceResponseResultInnerFillsInner) GetMatchType() string {
	if o == nil || common.IsNil(o.MatchType) {
		var ret string
		return ret
	}
	return *o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SorOrderPlaceResponseResultInnerFillsInner) GetMatchTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.MatchType) {
		return nil, false
	}
	return o.MatchType, true
}

// HasMatchType returns a boolean if a field has been set.
func (o *SorOrderPlaceResponseResultInnerFillsInner) HasMatchType() bool {
	if o != nil && !common.IsNil(o.MatchType) {
		return true
	}

	return false
}

// SetMatchType gets a reference to the given string and assigns it to the MatchType field.
func (o *SorOrderPlaceResponseResultInnerFillsInner) SetMatchType(v string) {
	o.MatchType = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *SorOrderPlaceResponseResultInnerFillsInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SorOrderPlaceResponseResultInnerFillsInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *SorOrderPlaceResponseResultInnerFillsInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *SorOrderPlaceResponseResultInnerFillsInner) SetPrice(v string) {
	o.Price = &v
}

// GetQty returns the Qty field value if set, zero value otherwise.
func (o *SorOrderPlaceResponseResultInnerFillsInner) GetQty() string {
	if o == nil || common.IsNil(o.Qty) {
		var ret string
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SorOrderPlaceResponseResultInnerFillsInner) GetQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Qty) {
		return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *SorOrderPlaceResponseResultInnerFillsInner) HasQty() bool {
	if o != nil && !common.IsNil(o.Qty) {
		return true
	}

	return false
}

// SetQty gets a reference to the given string and assigns it to the Qty field.
func (o *SorOrderPlaceResponseResultInnerFillsInner) SetQty(v string) {
	o.Qty = &v
}

// GetCommission returns the Commission field value if set, zero value otherwise.
func (o *SorOrderPlaceResponseResultInnerFillsInner) GetCommission() string {
	if o == nil || common.IsNil(o.Commission) {
		var ret string
		return ret
	}
	return *o.Commission
}

// GetCommissionOk returns a tuple with the Commission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SorOrderPlaceResponseResultInnerFillsInner) GetCommissionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Commission) {
		return nil, false
	}
	return o.Commission, true
}

// HasCommission returns a boolean if a field has been set.
func (o *SorOrderPlaceResponseResultInnerFillsInner) HasCommission() bool {
	if o != nil && !common.IsNil(o.Commission) {
		return true
	}

	return false
}

// SetCommission gets a reference to the given string and assigns it to the Commission field.
func (o *SorOrderPlaceResponseResultInnerFillsInner) SetCommission(v string) {
	o.Commission = &v
}

// GetCommissionAsset returns the CommissionAsset field value if set, zero value otherwise.
func (o *SorOrderPlaceResponseResultInnerFillsInner) GetCommissionAsset() string {
	if o == nil || common.IsNil(o.CommissionAsset) {
		var ret string
		return ret
	}
	return *o.CommissionAsset
}

// GetCommissionAssetOk returns a tuple with the CommissionAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SorOrderPlaceResponseResultInnerFillsInner) GetCommissionAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.CommissionAsset) {
		return nil, false
	}
	return o.CommissionAsset, true
}

// HasCommissionAsset returns a boolean if a field has been set.
func (o *SorOrderPlaceResponseResultInnerFillsInner) HasCommissionAsset() bool {
	if o != nil && !common.IsNil(o.CommissionAsset) {
		return true
	}

	return false
}

// SetCommissionAsset gets a reference to the given string and assigns it to the CommissionAsset field.
func (o *SorOrderPlaceResponseResultInnerFillsInner) SetCommissionAsset(v string) {
	o.CommissionAsset = &v
}

// GetTradeId returns the TradeId field value if set, zero value otherwise.
func (o *SorOrderPlaceResponseResultInnerFillsInner) GetTradeId() int64 {
	if o == nil || common.IsNil(o.TradeId) {
		var ret int64
		return ret
	}
	return *o.TradeId
}

// GetTradeIdOk returns a tuple with the TradeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SorOrderPlaceResponseResultInnerFillsInner) GetTradeIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TradeId) {
		return nil, false
	}
	return o.TradeId, true
}

// HasTradeId returns a boolean if a field has been set.
func (o *SorOrderPlaceResponseResultInnerFillsInner) HasTradeId() bool {
	if o != nil && !common.IsNil(o.TradeId) {
		return true
	}

	return false
}

// SetTradeId gets a reference to the given int64 and assigns it to the TradeId field.
func (o *SorOrderPlaceResponseResultInnerFillsInner) SetTradeId(v int64) {
	o.TradeId = &v
}

// GetAllocId returns the AllocId field value if set, zero value otherwise.
func (o *SorOrderPlaceResponseResultInnerFillsInner) GetAllocId() int64 {
	if o == nil || common.IsNil(o.AllocId) {
		var ret int64
		return ret
	}
	return *o.AllocId
}

// GetAllocIdOk returns a tuple with the AllocId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SorOrderPlaceResponseResultInnerFillsInner) GetAllocIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.AllocId) {
		return nil, false
	}
	return o.AllocId, true
}

// HasAllocId returns a boolean if a field has been set.
func (o *SorOrderPlaceResponseResultInnerFillsInner) HasAllocId() bool {
	if o != nil && !common.IsNil(o.AllocId) {
		return true
	}

	return false
}

// SetAllocId gets a reference to the given int64 and assigns it to the AllocId field.
func (o *SorOrderPlaceResponseResultInnerFillsInner) SetAllocId(v int64) {
	o.AllocId = &v
}

func (o SorOrderPlaceResponseResultInnerFillsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SorOrderPlaceResponseResultInnerFillsInner) ToMap() (map[string]interface{}, error) {
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

func (o *SorOrderPlaceResponseResultInnerFillsInner) UnmarshalJSON(data []byte) (err error) {
	varSorOrderPlaceResponseResultInnerFillsInner := _SorOrderPlaceResponseResultInnerFillsInner{}

	err = json.Unmarshal(data, &varSorOrderPlaceResponseResultInnerFillsInner)

	if err != nil {
		return err
	}

	*o = SorOrderPlaceResponseResultInnerFillsInner(varSorOrderPlaceResponseResultInnerFillsInner)

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

type NullableSorOrderPlaceResponseResultInnerFillsInner struct {
	value *SorOrderPlaceResponseResultInnerFillsInner
	isSet bool
}

func (v NullableSorOrderPlaceResponseResultInnerFillsInner) Get() *SorOrderPlaceResponseResultInnerFillsInner {
	return v.value
}

func (v *NullableSorOrderPlaceResponseResultInnerFillsInner) Set(val *SorOrderPlaceResponseResultInnerFillsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSorOrderPlaceResponseResultInnerFillsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSorOrderPlaceResponseResultInnerFillsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSorOrderPlaceResponseResultInnerFillsInner(val *SorOrderPlaceResponseResultInnerFillsInner) *NullableSorOrderPlaceResponseResultInnerFillsInner {
	return &NullableSorOrderPlaceResponseResultInnerFillsInner{value: val, isSet: true}
}

func (v NullableSorOrderPlaceResponseResultInnerFillsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSorOrderPlaceResponseResultInnerFillsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
