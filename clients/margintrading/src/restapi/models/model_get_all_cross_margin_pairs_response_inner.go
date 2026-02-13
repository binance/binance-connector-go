/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetAllCrossMarginPairsResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetAllCrossMarginPairsResponseInner{}

// GetAllCrossMarginPairsResponseInner struct for GetAllCrossMarginPairsResponseInner
type GetAllCrossMarginPairsResponseInner struct {
	Base                 *string `json:"base,omitempty"`
	Id                   *int64  `json:"id,omitempty"`
	IsBuyAllowed         *bool   `json:"isBuyAllowed,omitempty"`
	IsMarginTrade        *bool   `json:"isMarginTrade,omitempty"`
	IsSellAllowed        *bool   `json:"isSellAllowed,omitempty"`
	Quote                *string `json:"quote,omitempty"`
	Symbol               *string `json:"symbol,omitempty"`
	DelistTime           *int64  `json:"delistTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetAllCrossMarginPairsResponseInner GetAllCrossMarginPairsResponseInner

// NewGetAllCrossMarginPairsResponseInner instantiates a new GetAllCrossMarginPairsResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllCrossMarginPairsResponseInner() *GetAllCrossMarginPairsResponseInner {
	this := GetAllCrossMarginPairsResponseInner{}
	return &this
}

// NewGetAllCrossMarginPairsResponseInnerWithDefaults instantiates a new GetAllCrossMarginPairsResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllCrossMarginPairsResponseInnerWithDefaults() *GetAllCrossMarginPairsResponseInner {
	this := GetAllCrossMarginPairsResponseInner{}
	return &this
}

// GetBase returns the Base field value if set, zero value otherwise.
func (o *GetAllCrossMarginPairsResponseInner) GetBase() string {
	if o == nil || common.IsNil(o.Base) {
		var ret string
		return ret
	}
	return *o.Base
}

// GetBaseOk returns a tuple with the Base field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllCrossMarginPairsResponseInner) GetBaseOk() (*string, bool) {
	if o == nil || common.IsNil(o.Base) {
		return nil, false
	}
	return o.Base, true
}

// HasBase returns a boolean if a field has been set.
func (o *GetAllCrossMarginPairsResponseInner) HasBase() bool {
	if o != nil && !common.IsNil(o.Base) {
		return true
	}

	return false
}

// SetBase gets a reference to the given string and assigns it to the Base field.
func (o *GetAllCrossMarginPairsResponseInner) SetBase(v string) {
	o.Base = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetAllCrossMarginPairsResponseInner) GetId() int64 {
	if o == nil || common.IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllCrossMarginPairsResponseInner) GetIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetAllCrossMarginPairsResponseInner) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *GetAllCrossMarginPairsResponseInner) SetId(v int64) {
	o.Id = &v
}

// GetIsBuyAllowed returns the IsBuyAllowed field value if set, zero value otherwise.
func (o *GetAllCrossMarginPairsResponseInner) GetIsBuyAllowed() bool {
	if o == nil || common.IsNil(o.IsBuyAllowed) {
		var ret bool
		return ret
	}
	return *o.IsBuyAllowed
}

// GetIsBuyAllowedOk returns a tuple with the IsBuyAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllCrossMarginPairsResponseInner) GetIsBuyAllowedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsBuyAllowed) {
		return nil, false
	}
	return o.IsBuyAllowed, true
}

// HasIsBuyAllowed returns a boolean if a field has been set.
func (o *GetAllCrossMarginPairsResponseInner) HasIsBuyAllowed() bool {
	if o != nil && !common.IsNil(o.IsBuyAllowed) {
		return true
	}

	return false
}

// SetIsBuyAllowed gets a reference to the given bool and assigns it to the IsBuyAllowed field.
func (o *GetAllCrossMarginPairsResponseInner) SetIsBuyAllowed(v bool) {
	o.IsBuyAllowed = &v
}

// GetIsMarginTrade returns the IsMarginTrade field value if set, zero value otherwise.
func (o *GetAllCrossMarginPairsResponseInner) GetIsMarginTrade() bool {
	if o == nil || common.IsNil(o.IsMarginTrade) {
		var ret bool
		return ret
	}
	return *o.IsMarginTrade
}

// GetIsMarginTradeOk returns a tuple with the IsMarginTrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllCrossMarginPairsResponseInner) GetIsMarginTradeOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsMarginTrade) {
		return nil, false
	}
	return o.IsMarginTrade, true
}

// HasIsMarginTrade returns a boolean if a field has been set.
func (o *GetAllCrossMarginPairsResponseInner) HasIsMarginTrade() bool {
	if o != nil && !common.IsNil(o.IsMarginTrade) {
		return true
	}

	return false
}

// SetIsMarginTrade gets a reference to the given bool and assigns it to the IsMarginTrade field.
func (o *GetAllCrossMarginPairsResponseInner) SetIsMarginTrade(v bool) {
	o.IsMarginTrade = &v
}

// GetIsSellAllowed returns the IsSellAllowed field value if set, zero value otherwise.
func (o *GetAllCrossMarginPairsResponseInner) GetIsSellAllowed() bool {
	if o == nil || common.IsNil(o.IsSellAllowed) {
		var ret bool
		return ret
	}
	return *o.IsSellAllowed
}

// GetIsSellAllowedOk returns a tuple with the IsSellAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllCrossMarginPairsResponseInner) GetIsSellAllowedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsSellAllowed) {
		return nil, false
	}
	return o.IsSellAllowed, true
}

// HasIsSellAllowed returns a boolean if a field has been set.
func (o *GetAllCrossMarginPairsResponseInner) HasIsSellAllowed() bool {
	if o != nil && !common.IsNil(o.IsSellAllowed) {
		return true
	}

	return false
}

// SetIsSellAllowed gets a reference to the given bool and assigns it to the IsSellAllowed field.
func (o *GetAllCrossMarginPairsResponseInner) SetIsSellAllowed(v bool) {
	o.IsSellAllowed = &v
}

// GetQuote returns the Quote field value if set, zero value otherwise.
func (o *GetAllCrossMarginPairsResponseInner) GetQuote() string {
	if o == nil || common.IsNil(o.Quote) {
		var ret string
		return ret
	}
	return *o.Quote
}

// GetQuoteOk returns a tuple with the Quote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllCrossMarginPairsResponseInner) GetQuoteOk() (*string, bool) {
	if o == nil || common.IsNil(o.Quote) {
		return nil, false
	}
	return o.Quote, true
}

// HasQuote returns a boolean if a field has been set.
func (o *GetAllCrossMarginPairsResponseInner) HasQuote() bool {
	if o != nil && !common.IsNil(o.Quote) {
		return true
	}

	return false
}

// SetQuote gets a reference to the given string and assigns it to the Quote field.
func (o *GetAllCrossMarginPairsResponseInner) SetQuote(v string) {
	o.Quote = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetAllCrossMarginPairsResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllCrossMarginPairsResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetAllCrossMarginPairsResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetAllCrossMarginPairsResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetDelistTime returns the DelistTime field value if set, zero value otherwise.
func (o *GetAllCrossMarginPairsResponseInner) GetDelistTime() int64 {
	if o == nil || common.IsNil(o.DelistTime) {
		var ret int64
		return ret
	}
	return *o.DelistTime
}

// GetDelistTimeOk returns a tuple with the DelistTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllCrossMarginPairsResponseInner) GetDelistTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.DelistTime) {
		return nil, false
	}
	return o.DelistTime, true
}

// HasDelistTime returns a boolean if a field has been set.
func (o *GetAllCrossMarginPairsResponseInner) HasDelistTime() bool {
	if o != nil && !common.IsNil(o.DelistTime) {
		return true
	}

	return false
}

// SetDelistTime gets a reference to the given int64 and assigns it to the DelistTime field.
func (o *GetAllCrossMarginPairsResponseInner) SetDelistTime(v int64) {
	o.DelistTime = &v
}

func (o GetAllCrossMarginPairsResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllCrossMarginPairsResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Base) {
		toSerialize["base"] = o.Base
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.IsBuyAllowed) {
		toSerialize["isBuyAllowed"] = o.IsBuyAllowed
	}
	if !common.IsNil(o.IsMarginTrade) {
		toSerialize["isMarginTrade"] = o.IsMarginTrade
	}
	if !common.IsNil(o.IsSellAllowed) {
		toSerialize["isSellAllowed"] = o.IsSellAllowed
	}
	if !common.IsNil(o.Quote) {
		toSerialize["quote"] = o.Quote
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.DelistTime) {
		toSerialize["delistTime"] = o.DelistTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetAllCrossMarginPairsResponseInner) UnmarshalJSON(data []byte) (err error) {
	varGetAllCrossMarginPairsResponseInner := _GetAllCrossMarginPairsResponseInner{}

	err = json.Unmarshal(data, &varGetAllCrossMarginPairsResponseInner)

	if err != nil {
		return err
	}

	*o = GetAllCrossMarginPairsResponseInner(varGetAllCrossMarginPairsResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "base")
		delete(additionalProperties, "id")
		delete(additionalProperties, "isBuyAllowed")
		delete(additionalProperties, "isMarginTrade")
		delete(additionalProperties, "isSellAllowed")
		delete(additionalProperties, "quote")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "delistTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetAllCrossMarginPairsResponseInner struct {
	value *GetAllCrossMarginPairsResponseInner
	isSet bool
}

func (v NullableGetAllCrossMarginPairsResponseInner) Get() *GetAllCrossMarginPairsResponseInner {
	return v.value
}

func (v *NullableGetAllCrossMarginPairsResponseInner) Set(val *GetAllCrossMarginPairsResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllCrossMarginPairsResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllCrossMarginPairsResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllCrossMarginPairsResponseInner(val *GetAllCrossMarginPairsResponseInner) *NullableGetAllCrossMarginPairsResponseInner {
	return &NullableGetAllCrossMarginPairsResponseInner{value: val, isSet: true}
}

func (v NullableGetAllCrossMarginPairsResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllCrossMarginPairsResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
