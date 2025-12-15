/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryCrossIsolatedMarginCapitalFlowResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryCrossIsolatedMarginCapitalFlowResponseInner{}

// QueryCrossIsolatedMarginCapitalFlowResponseInner struct for QueryCrossIsolatedMarginCapitalFlowResponseInner
type QueryCrossIsolatedMarginCapitalFlowResponseInner struct {
	Id                   *int64  `json:"id,omitempty"`
	TranId               *int64  `json:"tranId,omitempty"`
	Timestamp            *int64  `json:"timestamp,omitempty"`
	Asset                *string `json:"asset,omitempty"`
	Symbol               *string `json:"symbol,omitempty"`
	Type                 *string `json:"type,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryCrossIsolatedMarginCapitalFlowResponseInner QueryCrossIsolatedMarginCapitalFlowResponseInner

// NewQueryCrossIsolatedMarginCapitalFlowResponseInner instantiates a new QueryCrossIsolatedMarginCapitalFlowResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryCrossIsolatedMarginCapitalFlowResponseInner() *QueryCrossIsolatedMarginCapitalFlowResponseInner {
	this := QueryCrossIsolatedMarginCapitalFlowResponseInner{}
	return &this
}

// NewQueryCrossIsolatedMarginCapitalFlowResponseInnerWithDefaults instantiates a new QueryCrossIsolatedMarginCapitalFlowResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryCrossIsolatedMarginCapitalFlowResponseInnerWithDefaults() *QueryCrossIsolatedMarginCapitalFlowResponseInner {
	this := QueryCrossIsolatedMarginCapitalFlowResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *QueryCrossIsolatedMarginCapitalFlowResponseInner) GetId() int64 {
	if o == nil || common.IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCrossIsolatedMarginCapitalFlowResponseInner) GetIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *QueryCrossIsolatedMarginCapitalFlowResponseInner) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *QueryCrossIsolatedMarginCapitalFlowResponseInner) SetId(v int64) {
	o.Id = &v
}

// GetTranId returns the TranId field value if set, zero value otherwise.
func (o *QueryCrossIsolatedMarginCapitalFlowResponseInner) GetTranId() int64 {
	if o == nil || common.IsNil(o.TranId) {
		var ret int64
		return ret
	}
	return *o.TranId
}

// GetTranIdOk returns a tuple with the TranId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCrossIsolatedMarginCapitalFlowResponseInner) GetTranIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TranId) {
		return nil, false
	}
	return o.TranId, true
}

// HasTranId returns a boolean if a field has been set.
func (o *QueryCrossIsolatedMarginCapitalFlowResponseInner) HasTranId() bool {
	if o != nil && !common.IsNil(o.TranId) {
		return true
	}

	return false
}

// SetTranId gets a reference to the given int64 and assigns it to the TranId field.
func (o *QueryCrossIsolatedMarginCapitalFlowResponseInner) SetTranId(v int64) {
	o.TranId = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *QueryCrossIsolatedMarginCapitalFlowResponseInner) GetTimestamp() int64 {
	if o == nil || common.IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCrossIsolatedMarginCapitalFlowResponseInner) GetTimestampOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *QueryCrossIsolatedMarginCapitalFlowResponseInner) HasTimestamp() bool {
	if o != nil && !common.IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *QueryCrossIsolatedMarginCapitalFlowResponseInner) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *QueryCrossIsolatedMarginCapitalFlowResponseInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCrossIsolatedMarginCapitalFlowResponseInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *QueryCrossIsolatedMarginCapitalFlowResponseInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *QueryCrossIsolatedMarginCapitalFlowResponseInner) SetAsset(v string) {
	o.Asset = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *QueryCrossIsolatedMarginCapitalFlowResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCrossIsolatedMarginCapitalFlowResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *QueryCrossIsolatedMarginCapitalFlowResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *QueryCrossIsolatedMarginCapitalFlowResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *QueryCrossIsolatedMarginCapitalFlowResponseInner) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCrossIsolatedMarginCapitalFlowResponseInner) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *QueryCrossIsolatedMarginCapitalFlowResponseInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *QueryCrossIsolatedMarginCapitalFlowResponseInner) SetType(v string) {
	o.Type = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *QueryCrossIsolatedMarginCapitalFlowResponseInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCrossIsolatedMarginCapitalFlowResponseInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *QueryCrossIsolatedMarginCapitalFlowResponseInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *QueryCrossIsolatedMarginCapitalFlowResponseInner) SetAmount(v string) {
	o.Amount = &v
}

func (o QueryCrossIsolatedMarginCapitalFlowResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryCrossIsolatedMarginCapitalFlowResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.TranId) {
		toSerialize["tranId"] = o.TranId
	}
	if !common.IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryCrossIsolatedMarginCapitalFlowResponseInner) UnmarshalJSON(data []byte) (err error) {
	varQueryCrossIsolatedMarginCapitalFlowResponseInner := _QueryCrossIsolatedMarginCapitalFlowResponseInner{}

	err = json.Unmarshal(data, &varQueryCrossIsolatedMarginCapitalFlowResponseInner)

	if err != nil {
		return err
	}

	*o = QueryCrossIsolatedMarginCapitalFlowResponseInner(varQueryCrossIsolatedMarginCapitalFlowResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "tranId")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "type")
		delete(additionalProperties, "amount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryCrossIsolatedMarginCapitalFlowResponseInner struct {
	value *QueryCrossIsolatedMarginCapitalFlowResponseInner
	isSet bool
}

func (v NullableQueryCrossIsolatedMarginCapitalFlowResponseInner) Get() *QueryCrossIsolatedMarginCapitalFlowResponseInner {
	return v.value
}

func (v *NullableQueryCrossIsolatedMarginCapitalFlowResponseInner) Set(val *QueryCrossIsolatedMarginCapitalFlowResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryCrossIsolatedMarginCapitalFlowResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryCrossIsolatedMarginCapitalFlowResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryCrossIsolatedMarginCapitalFlowResponseInner(val *QueryCrossIsolatedMarginCapitalFlowResponseInner) *NullableQueryCrossIsolatedMarginCapitalFlowResponseInner {
	return &NullableQueryCrossIsolatedMarginCapitalFlowResponseInner{value: val, isSet: true}
}

func (v NullableQueryCrossIsolatedMarginCapitalFlowResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryCrossIsolatedMarginCapitalFlowResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
