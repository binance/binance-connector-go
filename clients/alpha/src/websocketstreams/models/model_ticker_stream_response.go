/*
Alpha WebSocket Market Streams

Access Alpha market streams over WebSocket.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the TickerStreamResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TickerStreamResponse{}

// TickerStreamResponse struct for TickerStreamResponse
type TickerStreamResponse struct {
	// eventType
	Smalle *string `json:"e,omitempty"`
	// eventTime
	E *int64 `json:"E,omitempty"`
	// symbol
	Smalls *string `json:"s,omitempty"`
	// priceChange
	Smallp *string `json:"p,omitempty"`
	// priceChangePercent
	P *string `json:"P,omitempty"`
	// averagePrice
	Smallw *string `json:"w,omitempty"`
	// closePrice
	Smallc *string `json:"c,omitempty"`
	// lastTradeVolume
	Q *string `json:"Q,omitempty"`
	// openPrice
	Smallo *string `json:"o,omitempty"`
	// highPrice
	Smallh *string `json:"h,omitempty"`
	// lowPrice
	Smalll *string `json:"l,omitempty"`
	// volume
	Smallv *string `json:"v,omitempty"`
	// quoteVolume
	Smallq *string `json:"q,omitempty"`
	// startTime
	O *int64 `json:"O,omitempty"`
	// endTime
	C *int64 `json:"C,omitempty"`
	// firstTradeId
	F *int64 `json:"F,omitempty"`
	// lastTradeId
	L *int64 `json:"L,omitempty"`
	// tradeNum
	Smalln               *int64 `json:"n,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TickerStreamResponse TickerStreamResponse

// NewTickerStreamResponse instantiates a new TickerStreamResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTickerStreamResponse() *TickerStreamResponse {
	this := TickerStreamResponse{}
	return &this
}

// NewTickerStreamResponseWithDefaults instantiates a new TickerStreamResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTickerStreamResponseWithDefaults() *TickerStreamResponse {
	this := TickerStreamResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *TickerStreamResponse) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerStreamResponse) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *TickerStreamResponse) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *TickerStreamResponse) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *TickerStreamResponse) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerStreamResponse) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *TickerStreamResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *TickerStreamResponse) SetE(v int64) {
	o.E = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *TickerStreamResponse) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerStreamResponse) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *TickerStreamResponse) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *TickerStreamResponse) SetSmalls(v string) {
	o.Smalls = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *TickerStreamResponse) GetSmallp() string {
	if o == nil || common.IsNil(o.Smallp) {
		var ret string
		return ret
	}
	return *o.Smallp
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerStreamResponse) GetSmallpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallp) {
		return nil, false
	}
	return o.Smallp, true
}

// HasP returns a boolean if a field has been set.
func (o *TickerStreamResponse) HasSmallp() bool {
	if o != nil && !common.IsNil(o.Smallp) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *TickerStreamResponse) SetSmallp(v string) {
	o.Smallp = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *TickerStreamResponse) GetP() string {
	if o == nil || common.IsNil(o.P) {
		var ret string
		return ret
	}
	return *o.P
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerStreamResponse) GetPOk() (*string, bool) {
	if o == nil || common.IsNil(o.P) {
		return nil, false
	}
	return o.P, true
}

// HasP returns a boolean if a field has been set.
func (o *TickerStreamResponse) HasP() bool {
	if o != nil && !common.IsNil(o.P) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *TickerStreamResponse) SetP(v string) {
	o.P = &v
}

// GetW returns the W field value if set, zero value otherwise.
func (o *TickerStreamResponse) GetSmallw() string {
	if o == nil || common.IsNil(o.Smallw) {
		var ret string
		return ret
	}
	return *o.Smallw
}

// GetWOk returns a tuple with the W field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerStreamResponse) GetSmallwOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallw) {
		return nil, false
	}
	return o.Smallw, true
}

// HasW returns a boolean if a field has been set.
func (o *TickerStreamResponse) HasSmallw() bool {
	if o != nil && !common.IsNil(o.Smallw) {
		return true
	}

	return false
}

// SetW gets a reference to the given string and assigns it to the W field.
func (o *TickerStreamResponse) SetSmallw(v string) {
	o.Smallw = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *TickerStreamResponse) GetSmallc() string {
	if o == nil || common.IsNil(o.Smallc) {
		var ret string
		return ret
	}
	return *o.Smallc
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerStreamResponse) GetSmallcOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallc) {
		return nil, false
	}
	return o.Smallc, true
}

// HasC returns a boolean if a field has been set.
func (o *TickerStreamResponse) HasSmallc() bool {
	if o != nil && !common.IsNil(o.Smallc) {
		return true
	}

	return false
}

// SetC gets a reference to the given string and assigns it to the C field.
func (o *TickerStreamResponse) SetSmallc(v string) {
	o.Smallc = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *TickerStreamResponse) GetQ() string {
	if o == nil || common.IsNil(o.Q) {
		var ret string
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerStreamResponse) GetQOk() (*string, bool) {
	if o == nil || common.IsNil(o.Q) {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *TickerStreamResponse) HasQ() bool {
	if o != nil && !common.IsNil(o.Q) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *TickerStreamResponse) SetQ(v string) {
	o.Q = &v
}

// GetO returns the O field value if set, zero value otherwise.
func (o *TickerStreamResponse) GetSmallo() string {
	if o == nil || common.IsNil(o.Smallo) {
		var ret string
		return ret
	}
	return *o.Smallo
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerStreamResponse) GetSmalloOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallo) {
		return nil, false
	}
	return o.Smallo, true
}

// HasO returns a boolean if a field has been set.
func (o *TickerStreamResponse) HasSmallo() bool {
	if o != nil && !common.IsNil(o.Smallo) {
		return true
	}

	return false
}

// SetO gets a reference to the given string and assigns it to the O field.
func (o *TickerStreamResponse) SetSmallo(v string) {
	o.Smallo = &v
}

// GetH returns the H field value if set, zero value otherwise.
func (o *TickerStreamResponse) GetSmallh() string {
	if o == nil || common.IsNil(o.Smallh) {
		var ret string
		return ret
	}
	return *o.Smallh
}

// GetHOk returns a tuple with the H field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerStreamResponse) GetSmallhOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallh) {
		return nil, false
	}
	return o.Smallh, true
}

// HasH returns a boolean if a field has been set.
func (o *TickerStreamResponse) HasSmallh() bool {
	if o != nil && !common.IsNil(o.Smallh) {
		return true
	}

	return false
}

// SetH gets a reference to the given string and assigns it to the H field.
func (o *TickerStreamResponse) SetSmallh(v string) {
	o.Smallh = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *TickerStreamResponse) GetSmalll() string {
	if o == nil || common.IsNil(o.Smalll) {
		var ret string
		return ret
	}
	return *o.Smalll
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerStreamResponse) GetSmalllOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalll) {
		return nil, false
	}
	return o.Smalll, true
}

// HasL returns a boolean if a field has been set.
func (o *TickerStreamResponse) HasSmalll() bool {
	if o != nil && !common.IsNil(o.Smalll) {
		return true
	}

	return false
}

// SetL gets a reference to the given string and assigns it to the L field.
func (o *TickerStreamResponse) SetSmalll(v string) {
	o.Smalll = &v
}

// GetV returns the V field value if set, zero value otherwise.
func (o *TickerStreamResponse) GetSmallv() string {
	if o == nil || common.IsNil(o.Smallv) {
		var ret string
		return ret
	}
	return *o.Smallv
}

// GetVOk returns a tuple with the V field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerStreamResponse) GetSmallvOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallv) {
		return nil, false
	}
	return o.Smallv, true
}

// HasV returns a boolean if a field has been set.
func (o *TickerStreamResponse) HasSmallv() bool {
	if o != nil && !common.IsNil(o.Smallv) {
		return true
	}

	return false
}

// SetV gets a reference to the given string and assigns it to the V field.
func (o *TickerStreamResponse) SetSmallv(v string) {
	o.Smallv = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *TickerStreamResponse) GetSmallq() string {
	if o == nil || common.IsNil(o.Smallq) {
		var ret string
		return ret
	}
	return *o.Smallq
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerStreamResponse) GetSmallqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallq) {
		return nil, false
	}
	return o.Smallq, true
}

// HasQ returns a boolean if a field has been set.
func (o *TickerStreamResponse) HasSmallq() bool {
	if o != nil && !common.IsNil(o.Smallq) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *TickerStreamResponse) SetSmallq(v string) {
	o.Smallq = &v
}

// GetO returns the O field value if set, zero value otherwise.
func (o *TickerStreamResponse) GetO() int64 {
	if o == nil || common.IsNil(o.O) {
		var ret int64
		return ret
	}
	return *o.O
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerStreamResponse) GetOOk() (*int64, bool) {
	if o == nil || common.IsNil(o.O) {
		return nil, false
	}
	return o.O, true
}

// HasO returns a boolean if a field has been set.
func (o *TickerStreamResponse) HasO() bool {
	if o != nil && !common.IsNil(o.O) {
		return true
	}

	return false
}

// SetO gets a reference to the given int64 and assigns it to the O field.
func (o *TickerStreamResponse) SetO(v int64) {
	o.O = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *TickerStreamResponse) GetC() int64 {
	if o == nil || common.IsNil(o.C) {
		var ret int64
		return ret
	}
	return *o.C
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerStreamResponse) GetCOk() (*int64, bool) {
	if o == nil || common.IsNil(o.C) {
		return nil, false
	}
	return o.C, true
}

// HasC returns a boolean if a field has been set.
func (o *TickerStreamResponse) HasC() bool {
	if o != nil && !common.IsNil(o.C) {
		return true
	}

	return false
}

// SetC gets a reference to the given int64 and assigns it to the C field.
func (o *TickerStreamResponse) SetC(v int64) {
	o.C = &v
}

// GetF returns the F field value if set, zero value otherwise.
func (o *TickerStreamResponse) GetF() int64 {
	if o == nil || common.IsNil(o.F) {
		var ret int64
		return ret
	}
	return *o.F
}

// GetFOk returns a tuple with the F field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerStreamResponse) GetFOk() (*int64, bool) {
	if o == nil || common.IsNil(o.F) {
		return nil, false
	}
	return o.F, true
}

// HasF returns a boolean if a field has been set.
func (o *TickerStreamResponse) HasF() bool {
	if o != nil && !common.IsNil(o.F) {
		return true
	}

	return false
}

// SetF gets a reference to the given int64 and assigns it to the F field.
func (o *TickerStreamResponse) SetF(v int64) {
	o.F = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *TickerStreamResponse) GetL() int64 {
	if o == nil || common.IsNil(o.L) {
		var ret int64
		return ret
	}
	return *o.L
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerStreamResponse) GetLOk() (*int64, bool) {
	if o == nil || common.IsNil(o.L) {
		return nil, false
	}
	return o.L, true
}

// HasL returns a boolean if a field has been set.
func (o *TickerStreamResponse) HasL() bool {
	if o != nil && !common.IsNil(o.L) {
		return true
	}

	return false
}

// SetL gets a reference to the given int64 and assigns it to the L field.
func (o *TickerStreamResponse) SetL(v int64) {
	o.L = &v
}

// GetN returns the N field value if set, zero value otherwise.
func (o *TickerStreamResponse) GetSmalln() int64 {
	if o == nil || common.IsNil(o.Smalln) {
		var ret int64
		return ret
	}
	return *o.Smalln
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerStreamResponse) GetSmallnOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smalln) {
		return nil, false
	}
	return o.Smalln, true
}

// HasN returns a boolean if a field has been set.
func (o *TickerStreamResponse) HasSmalln() bool {
	if o != nil && !common.IsNil(o.Smalln) {
		return true
	}

	return false
}

// SetN gets a reference to the given int64 and assigns it to the N field.
func (o *TickerStreamResponse) SetSmalln(v int64) {
	o.Smalln = &v
}

func (o TickerStreamResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TickerStreamResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalle) {
		toSerialize["e"] = o.Smalle
	}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}
	if !common.IsNil(o.Smallp) {
		toSerialize["p"] = o.Smallp
	}
	if !common.IsNil(o.P) {
		toSerialize["P"] = o.P
	}
	if !common.IsNil(o.Smallw) {
		toSerialize["w"] = o.Smallw
	}
	if !common.IsNil(o.Smallc) {
		toSerialize["c"] = o.Smallc
	}
	if !common.IsNil(o.Q) {
		toSerialize["Q"] = o.Q
	}
	if !common.IsNil(o.Smallo) {
		toSerialize["o"] = o.Smallo
	}
	if !common.IsNil(o.Smallh) {
		toSerialize["h"] = o.Smallh
	}
	if !common.IsNil(o.Smalll) {
		toSerialize["l"] = o.Smalll
	}
	if !common.IsNil(o.Smallv) {
		toSerialize["v"] = o.Smallv
	}
	if !common.IsNil(o.Smallq) {
		toSerialize["q"] = o.Smallq
	}
	if !common.IsNil(o.O) {
		toSerialize["O"] = o.O
	}
	if !common.IsNil(o.C) {
		toSerialize["C"] = o.C
	}
	if !common.IsNil(o.F) {
		toSerialize["F"] = o.F
	}
	if !common.IsNil(o.L) {
		toSerialize["L"] = o.L
	}
	if !common.IsNil(o.Smalln) {
		toSerialize["n"] = o.Smalln
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TickerStreamResponse) UnmarshalJSON(data []byte) (err error) {
	varTickerStreamResponse := _TickerStreamResponse{}

	err = json.Unmarshal(data, &varTickerStreamResponse)

	if err != nil {
		return err
	}

	*o = TickerStreamResponse(varTickerStreamResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "E")
		delete(additionalProperties, "s")
		delete(additionalProperties, "p")
		delete(additionalProperties, "P")
		delete(additionalProperties, "w")
		delete(additionalProperties, "c")
		delete(additionalProperties, "Q")
		delete(additionalProperties, "o")
		delete(additionalProperties, "h")
		delete(additionalProperties, "l")
		delete(additionalProperties, "v")
		delete(additionalProperties, "q")
		delete(additionalProperties, "O")
		delete(additionalProperties, "C")
		delete(additionalProperties, "F")
		delete(additionalProperties, "L")
		delete(additionalProperties, "n")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTickerStreamResponse struct {
	value *TickerStreamResponse
	isSet bool
}

func (v NullableTickerStreamResponse) Get() *TickerStreamResponse {
	return v.value
}

func (v *NullableTickerStreamResponse) Set(val *TickerStreamResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTickerStreamResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTickerStreamResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTickerStreamResponse(val *TickerStreamResponse) *NullableTickerStreamResponse {
	return &NullableTickerStreamResponse{value: val, isSet: true}
}

func (v NullableTickerStreamResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTickerStreamResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
