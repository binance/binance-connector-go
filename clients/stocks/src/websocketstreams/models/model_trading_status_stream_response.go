/*
Stocks Trading WebSocket Streams

WebSocket stream definitions for Binance Stocks Trading. Base URL: wss://nbstream.binance.com/equity
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the TradingStatusStreamResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TradingStatusStreamResponse{}

// TradingStatusStreamResponse struct for TradingStatusStreamResponse
type TradingStatusStreamResponse struct {
	// Event type, always `\"tradingStatus\"`.
	E *string `json:"e,omitempty"`
	// Symbol (UPPERCASE ticker), e.g. `\"AAPL\"`.
	Symbol *string `json:"symbol,omitempty"`
	// Internal asset code `EQ_{symbol}`; reference only.
	AssetCode *string `json:"assetCode,omitempty"`
	// Trading status.
	Status *string `json:"status,omitempty"`
	// Reason code.
	Msg *string `json:"msg,omitempty"`
	// Status-effective time (epoch milliseconds UTC). May be in the future.
	Time *int64 `json:"time,omitempty"`
	// Tape designation: `C` = CTA (NYSE / AMEX), `N` = UTP (Nasdaq).
	Z *string `json:"z,omitempty"`
	// Tradability after the status change: `BUY_SELL` / `BUY` / `SELL` / `NONE`.
	Tradability          *string `json:"tradability,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TradingStatusStreamResponse TradingStatusStreamResponse

// NewTradingStatusStreamResponse instantiates a new TradingStatusStreamResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTradingStatusStreamResponse() *TradingStatusStreamResponse {
	this := TradingStatusStreamResponse{}
	return &this
}

// NewTradingStatusStreamResponseWithDefaults instantiates a new TradingStatusStreamResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTradingStatusStreamResponseWithDefaults() *TradingStatusStreamResponse {
	this := TradingStatusStreamResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *TradingStatusStreamResponse) GetE() string {
	if o == nil || common.IsNil(o.E) {
		var ret string
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradingStatusStreamResponse) GetEOk() (*string, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *TradingStatusStreamResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *TradingStatusStreamResponse) SetE(v string) {
	o.E = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *TradingStatusStreamResponse) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradingStatusStreamResponse) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *TradingStatusStreamResponse) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *TradingStatusStreamResponse) SetSymbol(v string) {
	o.Symbol = &v
}

// GetAssetCode returns the AssetCode field value if set, zero value otherwise.
func (o *TradingStatusStreamResponse) GetAssetCode() string {
	if o == nil || common.IsNil(o.AssetCode) {
		var ret string
		return ret
	}
	return *o.AssetCode
}

// GetAssetCodeOk returns a tuple with the AssetCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradingStatusStreamResponse) GetAssetCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.AssetCode) {
		return nil, false
	}
	return o.AssetCode, true
}

// HasAssetCode returns a boolean if a field has been set.
func (o *TradingStatusStreamResponse) HasAssetCode() bool {
	if o != nil && !common.IsNil(o.AssetCode) {
		return true
	}

	return false
}

// SetAssetCode gets a reference to the given string and assigns it to the AssetCode field.
func (o *TradingStatusStreamResponse) SetAssetCode(v string) {
	o.AssetCode = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TradingStatusStreamResponse) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradingStatusStreamResponse) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TradingStatusStreamResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TradingStatusStreamResponse) SetStatus(v string) {
	o.Status = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *TradingStatusStreamResponse) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradingStatusStreamResponse) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *TradingStatusStreamResponse) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *TradingStatusStreamResponse) SetMsg(v string) {
	o.Msg = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *TradingStatusStreamResponse) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradingStatusStreamResponse) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *TradingStatusStreamResponse) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *TradingStatusStreamResponse) SetTime(v int64) {
	o.Time = &v
}

// GetZ returns the Z field value if set, zero value otherwise.
func (o *TradingStatusStreamResponse) GetZ() string {
	if o == nil || common.IsNil(o.Z) {
		var ret string
		return ret
	}
	return *o.Z
}

// GetZOk returns a tuple with the Z field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradingStatusStreamResponse) GetZOk() (*string, bool) {
	if o == nil || common.IsNil(o.Z) {
		return nil, false
	}
	return o.Z, true
}

// HasZ returns a boolean if a field has been set.
func (o *TradingStatusStreamResponse) HasZ() bool {
	if o != nil && !common.IsNil(o.Z) {
		return true
	}

	return false
}

// SetZ gets a reference to the given string and assigns it to the Z field.
func (o *TradingStatusStreamResponse) SetZ(v string) {
	o.Z = &v
}

// GetTradability returns the Tradability field value if set, zero value otherwise.
func (o *TradingStatusStreamResponse) GetTradability() string {
	if o == nil || common.IsNil(o.Tradability) {
		var ret string
		return ret
	}
	return *o.Tradability
}

// GetTradabilityOk returns a tuple with the Tradability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradingStatusStreamResponse) GetTradabilityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Tradability) {
		return nil, false
	}
	return o.Tradability, true
}

// HasTradability returns a boolean if a field has been set.
func (o *TradingStatusStreamResponse) HasTradability() bool {
	if o != nil && !common.IsNil(o.Tradability) {
		return true
	}

	return false
}

// SetTradability gets a reference to the given string and assigns it to the Tradability field.
func (o *TradingStatusStreamResponse) SetTradability(v string) {
	o.Tradability = &v
}

func (o TradingStatusStreamResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TradingStatusStreamResponse) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !common.IsNil(o.Z) {
		toSerialize["z"] = o.Z
	}
	if !common.IsNil(o.Tradability) {
		toSerialize["tradability"] = o.Tradability
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TradingStatusStreamResponse) UnmarshalJSON(data []byte) (err error) {
	varTradingStatusStreamResponse := _TradingStatusStreamResponse{}

	err = json.Unmarshal(data, &varTradingStatusStreamResponse)

	if err != nil {
		return err
	}

	*o = TradingStatusStreamResponse(varTradingStatusStreamResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "assetCode")
		delete(additionalProperties, "status")
		delete(additionalProperties, "msg")
		delete(additionalProperties, "time")
		delete(additionalProperties, "z")
		delete(additionalProperties, "tradability")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTradingStatusStreamResponse struct {
	value *TradingStatusStreamResponse
	isSet bool
}

func (v NullableTradingStatusStreamResponse) Get() *TradingStatusStreamResponse {
	return v.value
}

func (v *NullableTradingStatusStreamResponse) Set(val *TradingStatusStreamResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTradingStatusStreamResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTradingStatusStreamResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTradingStatusStreamResponse(val *TradingStatusStreamResponse) *NullableTradingStatusStreamResponse {
	return &NullableTradingStatusStreamResponse{value: val, isSet: true}
}

func (v NullableTradingStatusStreamResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTradingStatusStreamResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
