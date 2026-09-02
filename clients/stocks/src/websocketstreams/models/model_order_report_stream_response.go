/*
Stocks Trading WebSocket Streams

WebSocket stream definitions for Binance Stocks Trading. Base URL: wss://nbstream.binance.com/equity
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the OrderReportStreamResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderReportStreamResponse{}

// OrderReportStreamResponse struct for OrderReportStreamResponse
type OrderReportStreamResponse struct {
	// Event type, always `\"orderReport\"`.
	Smalle *string `json:"e,omitempty"`
	// Event time (epoch milliseconds); server push time.
	E *int64 `json:"E,omitempty"`
	// Execution type: `\"ORDER_UPDATE\"` (still open) or `\"ORDER_TERMINAL\"` (reached terminal state).
	Smallx *string `json:"x,omitempty"`
	// Order ID (UUID).
	Smalli *string `json:"i,omitempty"`
	// Asset ID (internal identifier).
	Smallai *string `json:"ai,omitempty"`
	// Base asset — the internal asset code with an `EQ_` prefix (e.g. `\"EQ_AAPL\"`), not the bare ticker used in REST responses / order input. Strip the `EQ_` prefix to match a symbol used elsewhere.
	Smallb *string `json:"b,omitempty"`
	// Quote currency, e.g. `\"USD\"`.
	Smallq *string `json:"q,omitempty"`
	// Order side: `\"buy\"` or `\"sell\"`. Note: lowercase, unlike REST responses.
	S *string `json:"S,omitempty"`
	// Order type: `\"market\"` / `\"limit\"` / `\"stop\"` / `\"stop_limit\"` / `\"trailing_stop\"`. Note: lowercase.
	Smallo *string `json:"o,omitempty"`
	// Limit price; null for market orders.
	Smallp *float64 `json:"p,omitempty"`
	// Order quantity (shares); `0` when the order was submitted as notional.
	Q *float64 `json:"Q,omitempty"`
	// Order notional; set when the order was submitted as notional (market buy), null when submitted as quantity.
	N *float64 `json:"N,omitempty"`
	// Filled quantity.
	Smallfq *float64 `json:"fq,omitempty"`
	// Filled notional (= filledQty × filledAvgPrice).
	FN *float64 `json:"FN,omitempty"`
	// Total cost — cumulative buy-in cost including the commission fee.
	Smalltc *float64 `json:"tc,omitempty"`
	// Fill progress percentage (0–100, 2 dp). By notional: `FN / N × 100`. By qty: `fq / Q × 100`.
	Z *float64 `json:"Z,omitempty"`
	// Trading session label, e.g. `\"Regular\"`, `\"24 Hours Trading\"`.
	Smalln *string `json:"n,omitempty"`
	// Order status, e.g. `\"accepted\"`, `\"partially_filled\"`, `\"filled\"`, `\"canceled\"`. Note: lowercase, unlike REST responses.
	Smalls *string `json:"s,omitempty"`
	// Order create time (epoch milliseconds).
	T *int64 `json:"T,omitempty"`
	// Order update time (epoch milliseconds).
	U                    *int64 `json:"U,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderReportStreamResponse OrderReportStreamResponse

// NewOrderReportStreamResponse instantiates a new OrderReportStreamResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderReportStreamResponse() *OrderReportStreamResponse {
	this := OrderReportStreamResponse{}
	return &this
}

// NewOrderReportStreamResponseWithDefaults instantiates a new OrderReportStreamResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderReportStreamResponseWithDefaults() *OrderReportStreamResponse {
	this := OrderReportStreamResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *OrderReportStreamResponse) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderReportStreamResponse) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *OrderReportStreamResponse) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *OrderReportStreamResponse) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *OrderReportStreamResponse) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderReportStreamResponse) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *OrderReportStreamResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *OrderReportStreamResponse) SetE(v int64) {
	o.E = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *OrderReportStreamResponse) GetSmallx() string {
	if o == nil || common.IsNil(o.Smallx) {
		var ret string
		return ret
	}
	return *o.Smallx
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderReportStreamResponse) GetSmallxOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallx) {
		return nil, false
	}
	return o.Smallx, true
}

// HasX returns a boolean if a field has been set.
func (o *OrderReportStreamResponse) HasSmallx() bool {
	if o != nil && !common.IsNil(o.Smallx) {
		return true
	}

	return false
}

// SetX gets a reference to the given string and assigns it to the X field.
func (o *OrderReportStreamResponse) SetSmallx(v string) {
	o.Smallx = &v
}

// GetI returns the I field value if set, zero value otherwise.
func (o *OrderReportStreamResponse) GetSmalli() string {
	if o == nil || common.IsNil(o.Smalli) {
		var ret string
		return ret
	}
	return *o.Smalli
}

// GetIOk returns a tuple with the I field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderReportStreamResponse) GetSmalliOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalli) {
		return nil, false
	}
	return o.Smalli, true
}

// HasI returns a boolean if a field has been set.
func (o *OrderReportStreamResponse) HasSmalli() bool {
	if o != nil && !common.IsNil(o.Smalli) {
		return true
	}

	return false
}

// SetI gets a reference to the given string and assigns it to the I field.
func (o *OrderReportStreamResponse) SetSmalli(v string) {
	o.Smalli = &v
}

// GetAi returns the Ai field value if set, zero value otherwise.
func (o *OrderReportStreamResponse) GetSmallai() string {
	if o == nil || common.IsNil(o.Smallai) {
		var ret string
		return ret
	}
	return *o.Smallai
}

// GetAiOk returns a tuple with the Ai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderReportStreamResponse) GetSmallaiOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallai) {
		return nil, false
	}
	return o.Smallai, true
}

// HasAi returns a boolean if a field has been set.
func (o *OrderReportStreamResponse) HasSmallai() bool {
	if o != nil && !common.IsNil(o.Smallai) {
		return true
	}

	return false
}

// SetAi gets a reference to the given string and assigns it to the Ai field.
func (o *OrderReportStreamResponse) SetSmallai(v string) {
	o.Smallai = &v
}

// GetB returns the B field value if set, zero value otherwise.
func (o *OrderReportStreamResponse) GetSmallb() string {
	if o == nil || common.IsNil(o.Smallb) {
		var ret string
		return ret
	}
	return *o.Smallb
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderReportStreamResponse) GetSmallbOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallb) {
		return nil, false
	}
	return o.Smallb, true
}

// HasB returns a boolean if a field has been set.
func (o *OrderReportStreamResponse) HasSmallb() bool {
	if o != nil && !common.IsNil(o.Smallb) {
		return true
	}

	return false
}

// SetB gets a reference to the given string and assigns it to the B field.
func (o *OrderReportStreamResponse) SetSmallb(v string) {
	o.Smallb = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *OrderReportStreamResponse) GetSmallq() string {
	if o == nil || common.IsNil(o.Smallq) {
		var ret string
		return ret
	}
	return *o.Smallq
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderReportStreamResponse) GetSmallqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallq) {
		return nil, false
	}
	return o.Smallq, true
}

// HasQ returns a boolean if a field has been set.
func (o *OrderReportStreamResponse) HasSmallq() bool {
	if o != nil && !common.IsNil(o.Smallq) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *OrderReportStreamResponse) SetSmallq(v string) {
	o.Smallq = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *OrderReportStreamResponse) GetS() string {
	if o == nil || common.IsNil(o.S) {
		var ret string
		return ret
	}
	return *o.S
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderReportStreamResponse) GetSOk() (*string, bool) {
	if o == nil || common.IsNil(o.S) {
		return nil, false
	}
	return o.S, true
}

// HasS returns a boolean if a field has been set.
func (o *OrderReportStreamResponse) HasS() bool {
	if o != nil && !common.IsNil(o.S) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *OrderReportStreamResponse) SetS(v string) {
	o.S = &v
}

// GetO returns the O field value if set, zero value otherwise.
func (o *OrderReportStreamResponse) GetSmallo() string {
	if o == nil || common.IsNil(o.Smallo) {
		var ret string
		return ret
	}
	return *o.Smallo
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderReportStreamResponse) GetSmalloOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallo) {
		return nil, false
	}
	return o.Smallo, true
}

// HasO returns a boolean if a field has been set.
func (o *OrderReportStreamResponse) HasSmallo() bool {
	if o != nil && !common.IsNil(o.Smallo) {
		return true
	}

	return false
}

// SetO gets a reference to the given string and assigns it to the O field.
func (o *OrderReportStreamResponse) SetSmallo(v string) {
	o.Smallo = &v
}

// GetP returns the P field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderReportStreamResponse) GetSmallp() float64 {
	if o == nil || common.IsNil(o.Smallp) {
		var ret float64
		return ret
	}
	return *o.Smallp
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderReportStreamResponse) GetSmallpOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Smallp, true
}

// HasP returns a boolean if a field has been set.
func (o *OrderReportStreamResponse) HasSmallp() bool {
	if o != nil && common.IsNil(o.Smallp) {
		return true
	}

	return false
}

// SetP gets a reference to the given NullableFloat64 and assigns it to the P field.
func (o *OrderReportStreamResponse) SetSmallp(v float64) {
	o.Smallp = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *OrderReportStreamResponse) GetQ() float64 {
	if o == nil || common.IsNil(o.Q) {
		var ret float64
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderReportStreamResponse) GetQOk() (*float64, bool) {
	if o == nil || common.IsNil(o.Q) {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *OrderReportStreamResponse) HasQ() bool {
	if o != nil && !common.IsNil(o.Q) {
		return true
	}

	return false
}

// SetQ gets a reference to the given float64 and assigns it to the Q field.
func (o *OrderReportStreamResponse) SetQ(v float64) {
	o.Q = &v
}

// GetN returns the N field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderReportStreamResponse) GetN() float64 {
	if o == nil || common.IsNil(o.N) {
		var ret float64
		return ret
	}
	return *o.N
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderReportStreamResponse) GetNOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.N, true
}

// HasN returns a boolean if a field has been set.
func (o *OrderReportStreamResponse) HasN() bool {
	if o != nil && common.IsNil(o.N) {
		return true
	}

	return false
}

// SetN gets a reference to the given NullableFloat64 and assigns it to the N field.
func (o *OrderReportStreamResponse) SetN(v float64) {
	o.N = &v
}

// GetFq returns the Fq field value if set, zero value otherwise.
func (o *OrderReportStreamResponse) GetSmallfq() float64 {
	if o == nil || common.IsNil(o.Smallfq) {
		var ret float64
		return ret
	}
	return *o.Smallfq
}

// GetFqOk returns a tuple with the Fq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderReportStreamResponse) GetSmallfqOk() (*float64, bool) {
	if o == nil || common.IsNil(o.Smallfq) {
		return nil, false
	}
	return o.Smallfq, true
}

// HasFq returns a boolean if a field has been set.
func (o *OrderReportStreamResponse) HasSmallfq() bool {
	if o != nil && !common.IsNil(o.Smallfq) {
		return true
	}

	return false
}

// SetFq gets a reference to the given float64 and assigns it to the Fq field.
func (o *OrderReportStreamResponse) SetSmallfq(v float64) {
	o.Smallfq = &v
}

// GetFN returns the FN field value if set, zero value otherwise.
func (o *OrderReportStreamResponse) GetFN() float64 {
	if o == nil || common.IsNil(o.FN) {
		var ret float64
		return ret
	}
	return *o.FN
}

// GetFNOk returns a tuple with the FN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderReportStreamResponse) GetFNOk() (*float64, bool) {
	if o == nil || common.IsNil(o.FN) {
		return nil, false
	}
	return o.FN, true
}

// HasFN returns a boolean if a field has been set.
func (o *OrderReportStreamResponse) HasFN() bool {
	if o != nil && !common.IsNil(o.FN) {
		return true
	}

	return false
}

// SetFN gets a reference to the given float64 and assigns it to the FN field.
func (o *OrderReportStreamResponse) SetFN(v float64) {
	o.FN = &v
}

// GetTc returns the Tc field value if set, zero value otherwise.
func (o *OrderReportStreamResponse) GetSmalltc() float64 {
	if o == nil || common.IsNil(o.Smalltc) {
		var ret float64
		return ret
	}
	return *o.Smalltc
}

// GetTcOk returns a tuple with the Tc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderReportStreamResponse) GetSmalltcOk() (*float64, bool) {
	if o == nil || common.IsNil(o.Smalltc) {
		return nil, false
	}
	return o.Smalltc, true
}

// HasTc returns a boolean if a field has been set.
func (o *OrderReportStreamResponse) HasSmalltc() bool {
	if o != nil && !common.IsNil(o.Smalltc) {
		return true
	}

	return false
}

// SetTc gets a reference to the given float64 and assigns it to the Tc field.
func (o *OrderReportStreamResponse) SetSmalltc(v float64) {
	o.Smalltc = &v
}

// GetZ returns the Z field value if set, zero value otherwise.
func (o *OrderReportStreamResponse) GetZ() float64 {
	if o == nil || common.IsNil(o.Z) {
		var ret float64
		return ret
	}
	return *o.Z
}

// GetZOk returns a tuple with the Z field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderReportStreamResponse) GetZOk() (*float64, bool) {
	if o == nil || common.IsNil(o.Z) {
		return nil, false
	}
	return o.Z, true
}

// HasZ returns a boolean if a field has been set.
func (o *OrderReportStreamResponse) HasZ() bool {
	if o != nil && !common.IsNil(o.Z) {
		return true
	}

	return false
}

// SetZ gets a reference to the given float64 and assigns it to the Z field.
func (o *OrderReportStreamResponse) SetZ(v float64) {
	o.Z = &v
}

// GetN returns the N field value if set, zero value otherwise.
func (o *OrderReportStreamResponse) GetSmalln() string {
	if o == nil || common.IsNil(o.Smalln) {
		var ret string
		return ret
	}
	return *o.Smalln
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderReportStreamResponse) GetSmallnOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalln) {
		return nil, false
	}
	return o.Smalln, true
}

// HasN returns a boolean if a field has been set.
func (o *OrderReportStreamResponse) HasSmalln() bool {
	if o != nil && !common.IsNil(o.Smalln) {
		return true
	}

	return false
}

// SetN gets a reference to the given string and assigns it to the N field.
func (o *OrderReportStreamResponse) SetSmalln(v string) {
	o.Smalln = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *OrderReportStreamResponse) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderReportStreamResponse) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *OrderReportStreamResponse) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *OrderReportStreamResponse) SetSmalls(v string) {
	o.Smalls = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *OrderReportStreamResponse) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderReportStreamResponse) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *OrderReportStreamResponse) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *OrderReportStreamResponse) SetT(v int64) {
	o.T = &v
}

// GetU returns the U field value if set, zero value otherwise.
func (o *OrderReportStreamResponse) GetU() int64 {
	if o == nil || common.IsNil(o.U) {
		var ret int64
		return ret
	}
	return *o.U
}

// GetUOk returns a tuple with the U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderReportStreamResponse) GetUOk() (*int64, bool) {
	if o == nil || common.IsNil(o.U) {
		return nil, false
	}
	return o.U, true
}

// HasU returns a boolean if a field has been set.
func (o *OrderReportStreamResponse) HasU() bool {
	if o != nil && !common.IsNil(o.U) {
		return true
	}

	return false
}

// SetU gets a reference to the given int64 and assigns it to the U field.
func (o *OrderReportStreamResponse) SetU(v int64) {
	o.U = &v
}

func (o OrderReportStreamResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderReportStreamResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalle) {
		toSerialize["e"] = o.Smalle
	}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.Smallx) {
		toSerialize["x"] = o.Smallx
	}
	if !common.IsNil(o.Smalli) {
		toSerialize["i"] = o.Smalli
	}
	if !common.IsNil(o.Smallai) {
		toSerialize["ai"] = o.Smallai
	}
	if !common.IsNil(o.Smallb) {
		toSerialize["b"] = o.Smallb
	}
	if !common.IsNil(o.Smallq) {
		toSerialize["q"] = o.Smallq
	}
	if !common.IsNil(o.S) {
		toSerialize["S"] = o.S
	}
	if !common.IsNil(o.Smallo) {
		toSerialize["o"] = o.Smallo
	}
	if !common.IsNil(o.Smallp) {
		toSerialize["p"] = o.Smallp
	}
	if !common.IsNil(o.Q) {
		toSerialize["Q"] = o.Q
	}
	if !common.IsNil(o.N) {
		toSerialize["N"] = o.N
	}
	if !common.IsNil(o.Smallfq) {
		toSerialize["fq"] = o.Smallfq
	}
	if !common.IsNil(o.FN) {
		toSerialize["FN"] = o.FN
	}
	if !common.IsNil(o.Smalltc) {
		toSerialize["tc"] = o.Smalltc
	}
	if !common.IsNil(o.Z) {
		toSerialize["Z"] = o.Z
	}
	if !common.IsNil(o.Smalln) {
		toSerialize["n"] = o.Smalln
	}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}
	if !common.IsNil(o.U) {
		toSerialize["U"] = o.U
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderReportStreamResponse) UnmarshalJSON(data []byte) (err error) {
	varOrderReportStreamResponse := _OrderReportStreamResponse{}

	err = json.Unmarshal(data, &varOrderReportStreamResponse)

	if err != nil {
		return err
	}

	*o = OrderReportStreamResponse(varOrderReportStreamResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "E")
		delete(additionalProperties, "x")
		delete(additionalProperties, "i")
		delete(additionalProperties, "ai")
		delete(additionalProperties, "b")
		delete(additionalProperties, "q")
		delete(additionalProperties, "S")
		delete(additionalProperties, "o")
		delete(additionalProperties, "p")
		delete(additionalProperties, "Q")
		delete(additionalProperties, "N")
		delete(additionalProperties, "fq")
		delete(additionalProperties, "FN")
		delete(additionalProperties, "tc")
		delete(additionalProperties, "Z")
		delete(additionalProperties, "n")
		delete(additionalProperties, "s")
		delete(additionalProperties, "T")
		delete(additionalProperties, "U")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrderReportStreamResponse struct {
	value *OrderReportStreamResponse
	isSet bool
}

func (v NullableOrderReportStreamResponse) Get() *OrderReportStreamResponse {
	return v.value
}

func (v *NullableOrderReportStreamResponse) Set(val *OrderReportStreamResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderReportStreamResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderReportStreamResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderReportStreamResponse(val *OrderReportStreamResponse) *NullableOrderReportStreamResponse {
	return &NullableOrderReportStreamResponse{value: val, isSet: true}
}

func (v NullableOrderReportStreamResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderReportStreamResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
