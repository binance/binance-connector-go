/*
Stocks Trading WebSocket Streams

WebSocket stream definitions for Binance Stocks Trading. Base URL: wss://nbstream.binance.com/equity
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the TradabilityStreamResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TradabilityStreamResponse{}

// TradabilityStreamResponse struct for TradabilityStreamResponse
type TradabilityStreamResponse struct {
	// Event type, always `\"tradability\"`.
	E *string `json:"e,omitempty"`
	// Symbol (UPPERCASE ticker).
	Symbol *string `json:"symbol,omitempty"`
	// Internal asset code `EQ_{symbol}`; reference only. May be null for symbols not yet in the symbol dictionary.
	AssetCode *string `json:"assetCode,omitempty"`
	// The new tradability value. `OFFMARKET` and `DELISTING` are protected states — trading-status events will not transition a symbol out of them.
	Tradability *string `json:"tradability,omitempty"`
	// Push time (epoch milliseconds UTC).
	T                    *int64 `json:"t,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TradabilityStreamResponse TradabilityStreamResponse

// NewTradabilityStreamResponse instantiates a new TradabilityStreamResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTradabilityStreamResponse() *TradabilityStreamResponse {
	this := TradabilityStreamResponse{}
	return &this
}

// NewTradabilityStreamResponseWithDefaults instantiates a new TradabilityStreamResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTradabilityStreamResponseWithDefaults() *TradabilityStreamResponse {
	this := TradabilityStreamResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *TradabilityStreamResponse) GetE() string {
	if o == nil || common.IsNil(o.E) {
		var ret string
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradabilityStreamResponse) GetEOk() (*string, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *TradabilityStreamResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *TradabilityStreamResponse) SetE(v string) {
	o.E = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *TradabilityStreamResponse) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradabilityStreamResponse) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *TradabilityStreamResponse) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *TradabilityStreamResponse) SetSymbol(v string) {
	o.Symbol = &v
}

// GetAssetCode returns the AssetCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TradabilityStreamResponse) GetAssetCode() string {
	if o == nil || common.IsNil(o.AssetCode) {
		var ret string
		return ret
	}
	return *o.AssetCode
}

// GetAssetCodeOk returns a tuple with the AssetCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TradabilityStreamResponse) GetAssetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetCode, true
}

// HasAssetCode returns a boolean if a field has been set.
func (o *TradabilityStreamResponse) HasAssetCode() bool {
	if o != nil && common.IsNil(o.AssetCode) {
		return true
	}

	return false
}

// SetAssetCode gets a reference to the given NullableString and assigns it to the AssetCode field.
func (o *TradabilityStreamResponse) SetAssetCode(v string) {
	o.AssetCode = &v
}

// GetTradability returns the Tradability field value if set, zero value otherwise.
func (o *TradabilityStreamResponse) GetTradability() string {
	if o == nil || common.IsNil(o.Tradability) {
		var ret string
		return ret
	}
	return *o.Tradability
}

// GetTradabilityOk returns a tuple with the Tradability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradabilityStreamResponse) GetTradabilityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Tradability) {
		return nil, false
	}
	return o.Tradability, true
}

// HasTradability returns a boolean if a field has been set.
func (o *TradabilityStreamResponse) HasTradability() bool {
	if o != nil && !common.IsNil(o.Tradability) {
		return true
	}

	return false
}

// SetTradability gets a reference to the given string and assigns it to the Tradability field.
func (o *TradabilityStreamResponse) SetTradability(v string) {
	o.Tradability = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *TradabilityStreamResponse) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradabilityStreamResponse) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *TradabilityStreamResponse) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *TradabilityStreamResponse) SetT(v int64) {
	o.T = &v
}

func (o TradabilityStreamResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TradabilityStreamResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.E) {
		toSerialize["e"] = o.E
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.AssetCode) {
		toSerialize["assetCode"] = o.AssetCode
	}
	if !common.IsNil(o.Tradability) {
		toSerialize["tradability"] = o.Tradability
	}
	if !common.IsNil(o.T) {
		toSerialize["t"] = o.T
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TradabilityStreamResponse) UnmarshalJSON(data []byte) (err error) {
	varTradabilityStreamResponse := _TradabilityStreamResponse{}

	err = json.Unmarshal(data, &varTradabilityStreamResponse)

	if err != nil {
		return err
	}

	*o = TradabilityStreamResponse(varTradabilityStreamResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "assetCode")
		delete(additionalProperties, "tradability")
		delete(additionalProperties, "t")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTradabilityStreamResponse struct {
	value *TradabilityStreamResponse
	isSet bool
}

func (v NullableTradabilityStreamResponse) Get() *TradabilityStreamResponse {
	return v.value
}

func (v *NullableTradabilityStreamResponse) Set(val *TradabilityStreamResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTradabilityStreamResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTradabilityStreamResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTradabilityStreamResponse(val *TradabilityStreamResponse) *NullableTradabilityStreamResponse {
	return &NullableTradabilityStreamResponse{value: val, isSet: true}
}

func (v NullableTradabilityStreamResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTradabilityStreamResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
