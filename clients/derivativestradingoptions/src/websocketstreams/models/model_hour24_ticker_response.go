/*
Options WebSocket Market Streams

Access market data, manage accounts, and trade Binance Options.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the Hour24TickerResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Hour24TickerResponse{}

// Hour24TickerResponse struct for Hour24TickerResponse
type Hour24TickerResponse struct {
	// Event type
	Smalle *string `json:"e,omitempty"`
	// Event time
	E *int64 `json:"E,omitempty"`
	// Symbol
	Smalls *string `json:"s,omitempty"`
	// Price change
	Smallp *string `json:"p,omitempty"`
	// Price change percent
	P *string `json:"P,omitempty"`
	// Weighted average price
	Smallw *string `json:"w,omitempty"`
	// Last price
	Smallc *string `json:"c,omitempty"`
	// Last quantity
	Q *string `json:"Q,omitempty"`
	// Open price
	Smallo *string `json:"o,omitempty"`
	// High price
	Smallh *string `json:"h,omitempty"`
	// Low price
	Smalll *string `json:"l,omitempty"`
	// Trading volume(in contracts)
	Smallv *string `json:"v,omitempty"`
	// trade amount(in quote asset)
	Smallq *string `json:"q,omitempty"`
	// Statistics open time
	O *int64 `json:"O,omitempty"`
	// Statistics close time
	C *int64 `json:"C,omitempty"`
	// First trade ID
	F *int64 `json:"F,omitempty"`
	// Last trade Id
	L *int64 `json:"L,omitempty"`
	// Total number of trade
	Smalln               *int64 `json:"n,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Hour24TickerResponse Hour24TickerResponse

// NewHour24TickerResponse instantiates a new Hour24TickerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHour24TickerResponse() *Hour24TickerResponse {
	this := Hour24TickerResponse{}
	return &this
}

// NewHour24TickerResponseWithDefaults instantiates a new Hour24TickerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHour24TickerResponseWithDefaults() *Hour24TickerResponse {
	this := Hour24TickerResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *Hour24TickerResponse) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hour24TickerResponse) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *Hour24TickerResponse) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *Hour24TickerResponse) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *Hour24TickerResponse) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hour24TickerResponse) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *Hour24TickerResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *Hour24TickerResponse) SetE(v int64) {
	o.E = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *Hour24TickerResponse) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hour24TickerResponse) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *Hour24TickerResponse) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *Hour24TickerResponse) SetSmalls(v string) {
	o.Smalls = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *Hour24TickerResponse) GetSmallp() string {
	if o == nil || common.IsNil(o.Smallp) {
		var ret string
		return ret
	}
	return *o.Smallp
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hour24TickerResponse) GetSmallpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallp) {
		return nil, false
	}
	return o.Smallp, true
}

// HasP returns a boolean if a field has been set.
func (o *Hour24TickerResponse) HasSmallp() bool {
	if o != nil && !common.IsNil(o.Smallp) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *Hour24TickerResponse) SetSmallp(v string) {
	o.Smallp = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *Hour24TickerResponse) GetP() string {
	if o == nil || common.IsNil(o.P) {
		var ret string
		return ret
	}
	return *o.P
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hour24TickerResponse) GetPOk() (*string, bool) {
	if o == nil || common.IsNil(o.P) {
		return nil, false
	}
	return o.P, true
}

// HasP returns a boolean if a field has been set.
func (o *Hour24TickerResponse) HasP() bool {
	if o != nil && !common.IsNil(o.P) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *Hour24TickerResponse) SetP(v string) {
	o.P = &v
}

// GetW returns the W field value if set, zero value otherwise.
func (o *Hour24TickerResponse) GetSmallw() string {
	if o == nil || common.IsNil(o.Smallw) {
		var ret string
		return ret
	}
	return *o.Smallw
}

// GetWOk returns a tuple with the W field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hour24TickerResponse) GetSmallwOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallw) {
		return nil, false
	}
	return o.Smallw, true
}

// HasW returns a boolean if a field has been set.
func (o *Hour24TickerResponse) HasSmallw() bool {
	if o != nil && !common.IsNil(o.Smallw) {
		return true
	}

	return false
}

// SetW gets a reference to the given string and assigns it to the W field.
func (o *Hour24TickerResponse) SetSmallw(v string) {
	o.Smallw = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *Hour24TickerResponse) GetSmallc() string {
	if o == nil || common.IsNil(o.Smallc) {
		var ret string
		return ret
	}
	return *o.Smallc
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hour24TickerResponse) GetSmallcOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallc) {
		return nil, false
	}
	return o.Smallc, true
}

// HasC returns a boolean if a field has been set.
func (o *Hour24TickerResponse) HasSmallc() bool {
	if o != nil && !common.IsNil(o.Smallc) {
		return true
	}

	return false
}

// SetC gets a reference to the given string and assigns it to the C field.
func (o *Hour24TickerResponse) SetSmallc(v string) {
	o.Smallc = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *Hour24TickerResponse) GetQ() string {
	if o == nil || common.IsNil(o.Q) {
		var ret string
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hour24TickerResponse) GetQOk() (*string, bool) {
	if o == nil || common.IsNil(o.Q) {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *Hour24TickerResponse) HasQ() bool {
	if o != nil && !common.IsNil(o.Q) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *Hour24TickerResponse) SetQ(v string) {
	o.Q = &v
}

// GetO returns the O field value if set, zero value otherwise.
func (o *Hour24TickerResponse) GetSmallo() string {
	if o == nil || common.IsNil(o.Smallo) {
		var ret string
		return ret
	}
	return *o.Smallo
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hour24TickerResponse) GetSmalloOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallo) {
		return nil, false
	}
	return o.Smallo, true
}

// HasO returns a boolean if a field has been set.
func (o *Hour24TickerResponse) HasSmallo() bool {
	if o != nil && !common.IsNil(o.Smallo) {
		return true
	}

	return false
}

// SetO gets a reference to the given string and assigns it to the O field.
func (o *Hour24TickerResponse) SetSmallo(v string) {
	o.Smallo = &v
}

// GetH returns the H field value if set, zero value otherwise.
func (o *Hour24TickerResponse) GetSmallh() string {
	if o == nil || common.IsNil(o.Smallh) {
		var ret string
		return ret
	}
	return *o.Smallh
}

// GetHOk returns a tuple with the H field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hour24TickerResponse) GetSmallhOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallh) {
		return nil, false
	}
	return o.Smallh, true
}

// HasH returns a boolean if a field has been set.
func (o *Hour24TickerResponse) HasSmallh() bool {
	if o != nil && !common.IsNil(o.Smallh) {
		return true
	}

	return false
}

// SetH gets a reference to the given string and assigns it to the H field.
func (o *Hour24TickerResponse) SetSmallh(v string) {
	o.Smallh = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *Hour24TickerResponse) GetSmalll() string {
	if o == nil || common.IsNil(o.Smalll) {
		var ret string
		return ret
	}
	return *o.Smalll
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hour24TickerResponse) GetSmalllOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalll) {
		return nil, false
	}
	return o.Smalll, true
}

// HasL returns a boolean if a field has been set.
func (o *Hour24TickerResponse) HasSmalll() bool {
	if o != nil && !common.IsNil(o.Smalll) {
		return true
	}

	return false
}

// SetL gets a reference to the given string and assigns it to the L field.
func (o *Hour24TickerResponse) SetSmalll(v string) {
	o.Smalll = &v
}

// GetV returns the V field value if set, zero value otherwise.
func (o *Hour24TickerResponse) GetSmallv() string {
	if o == nil || common.IsNil(o.Smallv) {
		var ret string
		return ret
	}
	return *o.Smallv
}

// GetVOk returns a tuple with the V field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hour24TickerResponse) GetSmallvOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallv) {
		return nil, false
	}
	return o.Smallv, true
}

// HasV returns a boolean if a field has been set.
func (o *Hour24TickerResponse) HasSmallv() bool {
	if o != nil && !common.IsNil(o.Smallv) {
		return true
	}

	return false
}

// SetV gets a reference to the given string and assigns it to the V field.
func (o *Hour24TickerResponse) SetSmallv(v string) {
	o.Smallv = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *Hour24TickerResponse) GetSmallq() string {
	if o == nil || common.IsNil(o.Smallq) {
		var ret string
		return ret
	}
	return *o.Smallq
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hour24TickerResponse) GetSmallqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallq) {
		return nil, false
	}
	return o.Smallq, true
}

// HasQ returns a boolean if a field has been set.
func (o *Hour24TickerResponse) HasSmallq() bool {
	if o != nil && !common.IsNil(o.Smallq) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *Hour24TickerResponse) SetSmallq(v string) {
	o.Smallq = &v
}

// GetO returns the O field value if set, zero value otherwise.
func (o *Hour24TickerResponse) GetO() int64 {
	if o == nil || common.IsNil(o.O) {
		var ret int64
		return ret
	}
	return *o.O
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hour24TickerResponse) GetOOk() (*int64, bool) {
	if o == nil || common.IsNil(o.O) {
		return nil, false
	}
	return o.O, true
}

// HasO returns a boolean if a field has been set.
func (o *Hour24TickerResponse) HasO() bool {
	if o != nil && !common.IsNil(o.O) {
		return true
	}

	return false
}

// SetO gets a reference to the given int64 and assigns it to the O field.
func (o *Hour24TickerResponse) SetO(v int64) {
	o.O = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *Hour24TickerResponse) GetC() int64 {
	if o == nil || common.IsNil(o.C) {
		var ret int64
		return ret
	}
	return *o.C
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hour24TickerResponse) GetCOk() (*int64, bool) {
	if o == nil || common.IsNil(o.C) {
		return nil, false
	}
	return o.C, true
}

// HasC returns a boolean if a field has been set.
func (o *Hour24TickerResponse) HasC() bool {
	if o != nil && !common.IsNil(o.C) {
		return true
	}

	return false
}

// SetC gets a reference to the given int64 and assigns it to the C field.
func (o *Hour24TickerResponse) SetC(v int64) {
	o.C = &v
}

// GetF returns the F field value if set, zero value otherwise.
func (o *Hour24TickerResponse) GetF() int64 {
	if o == nil || common.IsNil(o.F) {
		var ret int64
		return ret
	}
	return *o.F
}

// GetFOk returns a tuple with the F field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hour24TickerResponse) GetFOk() (*int64, bool) {
	if o == nil || common.IsNil(o.F) {
		return nil, false
	}
	return o.F, true
}

// HasF returns a boolean if a field has been set.
func (o *Hour24TickerResponse) HasF() bool {
	if o != nil && !common.IsNil(o.F) {
		return true
	}

	return false
}

// SetF gets a reference to the given int64 and assigns it to the F field.
func (o *Hour24TickerResponse) SetF(v int64) {
	o.F = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *Hour24TickerResponse) GetL() int64 {
	if o == nil || common.IsNil(o.L) {
		var ret int64
		return ret
	}
	return *o.L
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hour24TickerResponse) GetLOk() (*int64, bool) {
	if o == nil || common.IsNil(o.L) {
		return nil, false
	}
	return o.L, true
}

// HasL returns a boolean if a field has been set.
func (o *Hour24TickerResponse) HasL() bool {
	if o != nil && !common.IsNil(o.L) {
		return true
	}

	return false
}

// SetL gets a reference to the given int64 and assigns it to the L field.
func (o *Hour24TickerResponse) SetL(v int64) {
	o.L = &v
}

// GetN returns the N field value if set, zero value otherwise.
func (o *Hour24TickerResponse) GetSmalln() int64 {
	if o == nil || common.IsNil(o.Smalln) {
		var ret int64
		return ret
	}
	return *o.Smalln
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hour24TickerResponse) GetSmallnOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smalln) {
		return nil, false
	}
	return o.Smalln, true
}

// HasN returns a boolean if a field has been set.
func (o *Hour24TickerResponse) HasSmalln() bool {
	if o != nil && !common.IsNil(o.Smalln) {
		return true
	}

	return false
}

// SetN gets a reference to the given int64 and assigns it to the N field.
func (o *Hour24TickerResponse) SetSmalln(v int64) {
	o.Smalln = &v
}

func (o Hour24TickerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Hour24TickerResponse) ToMap() (map[string]interface{}, error) {
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

func (o *Hour24TickerResponse) UnmarshalJSON(data []byte) (err error) {
	varHour24TickerResponse := _Hour24TickerResponse{}

	err = json.Unmarshal(data, &varHour24TickerResponse)

	if err != nil {
		return err
	}

	*o = Hour24TickerResponse(varHour24TickerResponse)

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

type NullableHour24TickerResponse struct {
	value *Hour24TickerResponse
	isSet bool
}

func (v NullableHour24TickerResponse) Get() *Hour24TickerResponse {
	return v.value
}

func (v *NullableHour24TickerResponse) Set(val *Hour24TickerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHour24TickerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHour24TickerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHour24TickerResponse(val *Hour24TickerResponse) *NullableHour24TickerResponse {
	return &NullableHour24TickerResponse{value: val, isSet: true}
}

func (v NullableHour24TickerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHour24TickerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
