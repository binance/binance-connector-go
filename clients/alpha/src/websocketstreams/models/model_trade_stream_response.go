/*
Alpha WebSocket Market Streams

Access Alpha market streams over WebSocket.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the TradeStreamResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TradeStreamResponse{}

// TradeStreamResponse struct for TradeStreamResponse
type TradeStreamResponse struct {
	// eventType
	Smalle *string `json:"e,omitempty"`
	// eventTime
	E *int64 `json:"E,omitempty"`
	// tradeTime
	T *int64 `json:"T,omitempty"`
	// symbol
	Smalls *string `json:"s,omitempty"`
	// tradeId
	Smallt *int64 `json:"t,omitempty"`
	// fillPrice
	Smallp *string `json:"p,omitempty"`
	// fillQty
	Smallq *string `json:"q,omitempty"`
	// isBuyerMaker
	Smallm               *bool `json:"m,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TradeStreamResponse TradeStreamResponse

// NewTradeStreamResponse instantiates a new TradeStreamResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTradeStreamResponse() *TradeStreamResponse {
	this := TradeStreamResponse{}
	return &this
}

// NewTradeStreamResponseWithDefaults instantiates a new TradeStreamResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTradeStreamResponseWithDefaults() *TradeStreamResponse {
	this := TradeStreamResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *TradeStreamResponse) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeStreamResponse) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *TradeStreamResponse) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *TradeStreamResponse) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *TradeStreamResponse) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeStreamResponse) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *TradeStreamResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *TradeStreamResponse) SetE(v int64) {
	o.E = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *TradeStreamResponse) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeStreamResponse) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *TradeStreamResponse) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *TradeStreamResponse) SetT(v int64) {
	o.T = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *TradeStreamResponse) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeStreamResponse) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *TradeStreamResponse) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *TradeStreamResponse) SetSmalls(v string) {
	o.Smalls = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *TradeStreamResponse) GetSmallt() int64 {
	if o == nil || common.IsNil(o.Smallt) {
		var ret int64
		return ret
	}
	return *o.Smallt
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeStreamResponse) GetSmalltOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallt) {
		return nil, false
	}
	return o.Smallt, true
}

// HasT returns a boolean if a field has been set.
func (o *TradeStreamResponse) HasSmallt() bool {
	if o != nil && !common.IsNil(o.Smallt) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *TradeStreamResponse) SetSmallt(v int64) {
	o.Smallt = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *TradeStreamResponse) GetSmallp() string {
	if o == nil || common.IsNil(o.Smallp) {
		var ret string
		return ret
	}
	return *o.Smallp
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeStreamResponse) GetSmallpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallp) {
		return nil, false
	}
	return o.Smallp, true
}

// HasP returns a boolean if a field has been set.
func (o *TradeStreamResponse) HasSmallp() bool {
	if o != nil && !common.IsNil(o.Smallp) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *TradeStreamResponse) SetSmallp(v string) {
	o.Smallp = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *TradeStreamResponse) GetSmallq() string {
	if o == nil || common.IsNil(o.Smallq) {
		var ret string
		return ret
	}
	return *o.Smallq
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeStreamResponse) GetSmallqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallq) {
		return nil, false
	}
	return o.Smallq, true
}

// HasQ returns a boolean if a field has been set.
func (o *TradeStreamResponse) HasSmallq() bool {
	if o != nil && !common.IsNil(o.Smallq) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *TradeStreamResponse) SetSmallq(v string) {
	o.Smallq = &v
}

// GetM returns the M field value if set, zero value otherwise.
func (o *TradeStreamResponse) GetSmallm() bool {
	if o == nil || common.IsNil(o.Smallm) {
		var ret bool
		return ret
	}
	return *o.Smallm
}

// GetMOk returns a tuple with the M field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeStreamResponse) GetSmallmOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Smallm) {
		return nil, false
	}
	return o.Smallm, true
}

// HasM returns a boolean if a field has been set.
func (o *TradeStreamResponse) HasSmallm() bool {
	if o != nil && !common.IsNil(o.Smallm) {
		return true
	}

	return false
}

// SetM gets a reference to the given bool and assigns it to the M field.
func (o *TradeStreamResponse) SetSmallm(v bool) {
	o.Smallm = &v
}

func (o TradeStreamResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TradeStreamResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalle) {
		toSerialize["e"] = o.Smalle
	}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}
	if !common.IsNil(o.Smallt) {
		toSerialize["t"] = o.Smallt
	}
	if !common.IsNil(o.Smallp) {
		toSerialize["p"] = o.Smallp
	}
	if !common.IsNil(o.Smallq) {
		toSerialize["q"] = o.Smallq
	}
	if !common.IsNil(o.Smallm) {
		toSerialize["m"] = o.Smallm
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TradeStreamResponse) UnmarshalJSON(data []byte) (err error) {
	varTradeStreamResponse := _TradeStreamResponse{}

	err = json.Unmarshal(data, &varTradeStreamResponse)

	if err != nil {
		return err
	}

	*o = TradeStreamResponse(varTradeStreamResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "E")
		delete(additionalProperties, "T")
		delete(additionalProperties, "s")
		delete(additionalProperties, "t")
		delete(additionalProperties, "p")
		delete(additionalProperties, "q")
		delete(additionalProperties, "m")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTradeStreamResponse struct {
	value *TradeStreamResponse
	isSet bool
}

func (v NullableTradeStreamResponse) Get() *TradeStreamResponse {
	return v.value
}

func (v *NullableTradeStreamResponse) Set(val *TradeStreamResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTradeStreamResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTradeStreamResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTradeStreamResponse(val *TradeStreamResponse) *NullableTradeStreamResponse {
	return &NullableTradeStreamResponse{value: val, isSet: true}
}

func (v NullableTradeStreamResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTradeStreamResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
