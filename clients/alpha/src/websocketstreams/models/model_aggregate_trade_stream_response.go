/*
Alpha WebSocket Market Streams

Access Alpha market streams over WebSocket.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AggregateTradeStreamResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AggregateTradeStreamResponse{}

// AggregateTradeStreamResponse struct for AggregateTradeStreamResponse
type AggregateTradeStreamResponse struct {
	// Event type
	Smalle *string `json:"e,omitempty"`
	// Event time (ms)
	E *int64 `json:"E,omitempty"`
	// Trade time (ms)
	T *int64 `json:"T,omitempty"`
	// Aggregated trade ID
	Smalla *int64 `json:"a,omitempty"`
	// First trade ID in the aggregation
	Smallf *int64 `json:"f,omitempty"`
	// Last trade ID in the aggregation
	Smalll *int64 `json:"l,omitempty"`
	// Is the buyer the market maker
	Smallm *bool `json:"m,omitempty"`
	// Price
	Smallp *string `json:"p,omitempty"`
	// Quantity
	Smallq *string `json:"q,omitempty"`
	// Symbol
	Smalls               *string `json:"s,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AggregateTradeStreamResponse AggregateTradeStreamResponse

// NewAggregateTradeStreamResponse instantiates a new AggregateTradeStreamResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggregateTradeStreamResponse() *AggregateTradeStreamResponse {
	this := AggregateTradeStreamResponse{}
	return &this
}

// NewAggregateTradeStreamResponseWithDefaults instantiates a new AggregateTradeStreamResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggregateTradeStreamResponseWithDefaults() *AggregateTradeStreamResponse {
	this := AggregateTradeStreamResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *AggregateTradeStreamResponse) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateTradeStreamResponse) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *AggregateTradeStreamResponse) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *AggregateTradeStreamResponse) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *AggregateTradeStreamResponse) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateTradeStreamResponse) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *AggregateTradeStreamResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *AggregateTradeStreamResponse) SetE(v int64) {
	o.E = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *AggregateTradeStreamResponse) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateTradeStreamResponse) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *AggregateTradeStreamResponse) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *AggregateTradeStreamResponse) SetT(v int64) {
	o.T = &v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *AggregateTradeStreamResponse) GetSmalla() int64 {
	if o == nil || common.IsNil(o.Smalla) {
		var ret int64
		return ret
	}
	return *o.Smalla
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateTradeStreamResponse) GetSmallaOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smalla) {
		return nil, false
	}
	return o.Smalla, true
}

// HasA returns a boolean if a field has been set.
func (o *AggregateTradeStreamResponse) HasSmalla() bool {
	if o != nil && !common.IsNil(o.Smalla) {
		return true
	}

	return false
}

// SetA gets a reference to the given int64 and assigns it to the A field.
func (o *AggregateTradeStreamResponse) SetSmalla(v int64) {
	o.Smalla = &v
}

// GetF returns the F field value if set, zero value otherwise.
func (o *AggregateTradeStreamResponse) GetSmallf() int64 {
	if o == nil || common.IsNil(o.Smallf) {
		var ret int64
		return ret
	}
	return *o.Smallf
}

// GetFOk returns a tuple with the F field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateTradeStreamResponse) GetSmallfOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallf) {
		return nil, false
	}
	return o.Smallf, true
}

// HasF returns a boolean if a field has been set.
func (o *AggregateTradeStreamResponse) HasSmallf() bool {
	if o != nil && !common.IsNil(o.Smallf) {
		return true
	}

	return false
}

// SetF gets a reference to the given int64 and assigns it to the F field.
func (o *AggregateTradeStreamResponse) SetSmallf(v int64) {
	o.Smallf = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *AggregateTradeStreamResponse) GetSmalll() int64 {
	if o == nil || common.IsNil(o.Smalll) {
		var ret int64
		return ret
	}
	return *o.Smalll
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateTradeStreamResponse) GetSmalllOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smalll) {
		return nil, false
	}
	return o.Smalll, true
}

// HasL returns a boolean if a field has been set.
func (o *AggregateTradeStreamResponse) HasSmalll() bool {
	if o != nil && !common.IsNil(o.Smalll) {
		return true
	}

	return false
}

// SetL gets a reference to the given int64 and assigns it to the L field.
func (o *AggregateTradeStreamResponse) SetSmalll(v int64) {
	o.Smalll = &v
}

// GetM returns the M field value if set, zero value otherwise.
func (o *AggregateTradeStreamResponse) GetSmallm() bool {
	if o == nil || common.IsNil(o.Smallm) {
		var ret bool
		return ret
	}
	return *o.Smallm
}

// GetMOk returns a tuple with the M field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateTradeStreamResponse) GetSmallmOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Smallm) {
		return nil, false
	}
	return o.Smallm, true
}

// HasM returns a boolean if a field has been set.
func (o *AggregateTradeStreamResponse) HasSmallm() bool {
	if o != nil && !common.IsNil(o.Smallm) {
		return true
	}

	return false
}

// SetM gets a reference to the given bool and assigns it to the M field.
func (o *AggregateTradeStreamResponse) SetSmallm(v bool) {
	o.Smallm = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *AggregateTradeStreamResponse) GetSmallp() string {
	if o == nil || common.IsNil(o.Smallp) {
		var ret string
		return ret
	}
	return *o.Smallp
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateTradeStreamResponse) GetSmallpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallp) {
		return nil, false
	}
	return o.Smallp, true
}

// HasP returns a boolean if a field has been set.
func (o *AggregateTradeStreamResponse) HasSmallp() bool {
	if o != nil && !common.IsNil(o.Smallp) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *AggregateTradeStreamResponse) SetSmallp(v string) {
	o.Smallp = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *AggregateTradeStreamResponse) GetSmallq() string {
	if o == nil || common.IsNil(o.Smallq) {
		var ret string
		return ret
	}
	return *o.Smallq
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateTradeStreamResponse) GetSmallqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallq) {
		return nil, false
	}
	return o.Smallq, true
}

// HasQ returns a boolean if a field has been set.
func (o *AggregateTradeStreamResponse) HasSmallq() bool {
	if o != nil && !common.IsNil(o.Smallq) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *AggregateTradeStreamResponse) SetSmallq(v string) {
	o.Smallq = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *AggregateTradeStreamResponse) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateTradeStreamResponse) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *AggregateTradeStreamResponse) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *AggregateTradeStreamResponse) SetSmalls(v string) {
	o.Smalls = &v
}

func (o AggregateTradeStreamResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AggregateTradeStreamResponse) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.Smalla) {
		toSerialize["a"] = o.Smalla
	}
	if !common.IsNil(o.Smallf) {
		toSerialize["f"] = o.Smallf
	}
	if !common.IsNil(o.Smalll) {
		toSerialize["l"] = o.Smalll
	}
	if !common.IsNil(o.Smallm) {
		toSerialize["m"] = o.Smallm
	}
	if !common.IsNil(o.Smallp) {
		toSerialize["p"] = o.Smallp
	}
	if !common.IsNil(o.Smallq) {
		toSerialize["q"] = o.Smallq
	}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AggregateTradeStreamResponse) UnmarshalJSON(data []byte) (err error) {
	varAggregateTradeStreamResponse := _AggregateTradeStreamResponse{}

	err = json.Unmarshal(data, &varAggregateTradeStreamResponse)

	if err != nil {
		return err
	}

	*o = AggregateTradeStreamResponse(varAggregateTradeStreamResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "E")
		delete(additionalProperties, "T")
		delete(additionalProperties, "a")
		delete(additionalProperties, "f")
		delete(additionalProperties, "l")
		delete(additionalProperties, "m")
		delete(additionalProperties, "p")
		delete(additionalProperties, "q")
		delete(additionalProperties, "s")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAggregateTradeStreamResponse struct {
	value *AggregateTradeStreamResponse
	isSet bool
}

func (v NullableAggregateTradeStreamResponse) Get() *AggregateTradeStreamResponse {
	return v.value
}

func (v *NullableAggregateTradeStreamResponse) Set(val *AggregateTradeStreamResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAggregateTradeStreamResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAggregateTradeStreamResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggregateTradeStreamResponse(val *AggregateTradeStreamResponse) *NullableAggregateTradeStreamResponse {
	return &NullableAggregateTradeStreamResponse{value: val, isSet: true}
}

func (v NullableAggregateTradeStreamResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggregateTradeStreamResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
