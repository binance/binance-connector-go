/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner{}

// QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner struct for QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner
type QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner struct {
	Symbol               *string  `json:"symbol,omitempty"`
	EntryPrice           *int64   `json:"entryPrice,omitempty"`
	MarkPrice            *int64   `json:"markPrice,omitempty"`
	PositionAmt          *float32 `json:"positionAmt,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner

// NewQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner instantiates a new QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner() *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner {
	this := QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner{}
	return &this
}

// NewQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInnerWithDefaults instantiates a new QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInnerWithDefaults() *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner {
	this := QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetEntryPrice returns the EntryPrice field value if set, zero value otherwise.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner) GetEntryPrice() int64 {
	if o == nil || common.IsNil(o.EntryPrice) {
		var ret int64
		return ret
	}
	return *o.EntryPrice
}

// GetEntryPriceOk returns a tuple with the EntryPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner) GetEntryPriceOk() (*int64, bool) {
	if o == nil || common.IsNil(o.EntryPrice) {
		return nil, false
	}
	return o.EntryPrice, true
}

// HasEntryPrice returns a boolean if a field has been set.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner) HasEntryPrice() bool {
	if o != nil && !common.IsNil(o.EntryPrice) {
		return true
	}

	return false
}

// SetEntryPrice gets a reference to the given int64 and assigns it to the EntryPrice field.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner) SetEntryPrice(v int64) {
	o.EntryPrice = &v
}

// GetMarkPrice returns the MarkPrice field value if set, zero value otherwise.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner) GetMarkPrice() int64 {
	if o == nil || common.IsNil(o.MarkPrice) {
		var ret int64
		return ret
	}
	return *o.MarkPrice
}

// GetMarkPriceOk returns a tuple with the MarkPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner) GetMarkPriceOk() (*int64, bool) {
	if o == nil || common.IsNil(o.MarkPrice) {
		return nil, false
	}
	return o.MarkPrice, true
}

// HasMarkPrice returns a boolean if a field has been set.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner) HasMarkPrice() bool {
	if o != nil && !common.IsNil(o.MarkPrice) {
		return true
	}

	return false
}

// SetMarkPrice gets a reference to the given int64 and assigns it to the MarkPrice field.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner) SetMarkPrice(v int64) {
	o.MarkPrice = &v
}

// GetPositionAmt returns the PositionAmt field value if set, zero value otherwise.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner) GetPositionAmt() float32 {
	if o == nil || common.IsNil(o.PositionAmt) {
		var ret float32
		return ret
	}
	return *o.PositionAmt
}

// GetPositionAmtOk returns a tuple with the PositionAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner) GetPositionAmtOk() (*float32, bool) {
	if o == nil || common.IsNil(o.PositionAmt) {
		return nil, false
	}
	return o.PositionAmt, true
}

// HasPositionAmt returns a boolean if a field has been set.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner) HasPositionAmt() bool {
	if o != nil && !common.IsNil(o.PositionAmt) {
		return true
	}

	return false
}

// SetPositionAmt gets a reference to the given float32 and assigns it to the PositionAmt field.
func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner) SetPositionAmt(v float32) {
	o.PositionAmt = &v
}

func (o QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.EntryPrice) {
		toSerialize["entryPrice"] = o.EntryPrice
	}
	if !common.IsNil(o.MarkPrice) {
		toSerialize["markPrice"] = o.MarkPrice
	}
	if !common.IsNil(o.PositionAmt) {
		toSerialize["positionAmt"] = o.PositionAmt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner) UnmarshalJSON(data []byte) (err error) {
	varQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner := _QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner{}

	err = json.Unmarshal(data, &varQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner)

	if err != nil {
		return err
	}

	*o = QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner(varQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "entryPrice")
		delete(additionalProperties, "markPrice")
		delete(additionalProperties, "positionAmt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner struct {
	value *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner
	isSet bool
}

func (v NullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner) Get() *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner {
	return v.value
}

func (v *NullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner) Set(val *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner(val *QueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner) *NullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner {
	return &NullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner{value: val, isSet: true}
}

func (v NullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryManagedSubAccountFuturesAssetDetailsResponseSnapshotVosInnerDataPositionInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
